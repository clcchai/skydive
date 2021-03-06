/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package analyzer

import (
	"github.com/skydive-project/skydive/config"
	"github.com/skydive-project/skydive/graffiti/graph"
	"github.com/skydive-project/skydive/logging"
	"github.com/skydive-project/skydive/probe"
	"github.com/skydive-project/skydive/topology/probes/fabric"
	"github.com/skydive-project/skydive/topology/probes/istio"
	"github.com/skydive-project/skydive/topology/probes/k8s"
	"github.com/skydive-project/skydive/topology/probes/nsm"
	"github.com/skydive-project/skydive/topology/probes/ovn"
	"github.com/skydive-project/skydive/topology/probes/peering"
)

// NewTopologyProbeBundleFromConfig creates a new topology server probes from configuration
func NewTopologyProbeBundleFromConfig(g *graph.Graph) (*probe.Bundle, error) {
	list := config.GetStringSlice("analyzer.topology.probes")

	var handler probe.Handler
	var err error

	bundle := probe.NewBundle()

	fabricProbe, err := fabric.NewProbe(g)
	if err != nil {
		return nil, err
	}
	bundle.AddHandler("fabric", fabricProbe)
	bundle.AddHandler("peering", peering.NewProbe(g))

	for _, t := range list {
		if bundle.GetHandler(t) != nil {
			continue
		}

		switch t {
		case "ovn":
			addr := config.GetString("analyzer.topology.ovn.address")
			handler, err = ovn.NewProbe(g, addr)
		case "k8s":
			handler, err = k8s.NewK8sProbe(g)
		case "istio":
			handler, err = istio.NewIstioProbe(g)
		case "nsm":
			handler, err = nsm.NewNsmProbe(g)
		default:
			logging.GetLogger().Errorf("unknown probe type: %s", t)
			continue
		}

		if err != nil {
			return nil, err
		}
		if handler != nil {
			bundle.AddHandler(t, handler)
		}
	}

	return bundle, nil
}
