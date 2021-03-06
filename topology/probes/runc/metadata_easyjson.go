// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package runc

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesRunc(in *jlexer.Lexer, out *Metadata) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ContainerID":
			out.ContainerID = string(in.String())
		case "Status":
			out.Status = string(in.String())
		case "Labels":
			(out.Labels).UnmarshalEasyJSON(in)
		case "CreateConfig":
			if in.IsNull() {
				in.Skip()
				out.CreateConfig = nil
			} else {
				if out.CreateConfig == nil {
					out.CreateConfig = new(CreateConfig)
				}
				(*out.CreateConfig).UnmarshalEasyJSON(in)
			}
		case "Hosts":
			if in.IsNull() {
				in.Skip()
				out.Hosts = nil
			} else {
				if out.Hosts == nil {
					out.Hosts = new(Hosts)
				}
				(*out.Hosts).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesRunc(out *jwriter.Writer, in Metadata) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ContainerID != "" {
		const prefix string = ",\"ContainerID\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ContainerID))
	}
	if in.Status != "" {
		const prefix string = ",\"Status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	if len(in.Labels) != 0 {
		const prefix string = ",\"Labels\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Labels).MarshalEasyJSON(out)
	}
	if in.CreateConfig != nil {
		const prefix string = ",\"CreateConfig\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.CreateConfig).MarshalEasyJSON(out)
	}
	if in.Hosts != nil {
		const prefix string = ",\"Hosts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Hosts).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Metadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesRunc(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Metadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesRunc(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Metadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesRunc(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Metadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesRunc(l, v)
}
func easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesRunc1(in *jlexer.Lexer, out *Hosts) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "IP":
			out.IP = string(in.String())
		case "Hostname":
			out.Hostname = string(in.String())
		case "ByIP":
			(out.ByIP).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesRunc1(out *jwriter.Writer, in Hosts) {
	out.RawByte('{')
	first := true
	_ = first
	if in.IP != "" {
		const prefix string = ",\"IP\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.IP))
	}
	if in.Hostname != "" {
		const prefix string = ",\"Hostname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Hostname))
	}
	if len(in.ByIP) != 0 {
		const prefix string = ",\"ByIP\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.ByIP).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Hosts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesRunc1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Hosts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesRunc1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Hosts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesRunc1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Hosts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesRunc1(l, v)
}
func easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesRunc2(in *jlexer.Lexer, out *CreateConfig) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Image":
			out.Image = string(in.String())
		case "ImageID":
			out.ImageID = string(in.String())
		case "Labels":
			(out.Labels).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesRunc2(out *jwriter.Writer, in CreateConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Image != "" {
		const prefix string = ",\"Image\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Image))
	}
	if in.ImageID != "" {
		const prefix string = ",\"ImageID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ImageID))
	}
	if len(in.Labels) != 0 {
		const prefix string = ",\"Labels\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Labels).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateConfig) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesRunc2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateConfig) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesRunc2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateConfig) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesRunc2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateConfig) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesRunc2(l, v)
}
