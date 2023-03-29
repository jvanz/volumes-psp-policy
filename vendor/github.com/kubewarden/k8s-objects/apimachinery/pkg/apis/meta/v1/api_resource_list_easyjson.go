// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjsonBd8685ecDecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(in *jlexer.Lexer, out *APIResourceList) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "groupVersion":
			if in.IsNull() {
				in.Skip()
				out.GroupVersion = nil
			} else {
				if out.GroupVersion == nil {
					out.GroupVersion = new(string)
				}
				*out.GroupVersion = string(in.String())
			}
		case "kind":
			out.Kind = string(in.String())
		case "resources":
			if in.IsNull() {
				in.Skip()
				out.Resources = nil
			} else {
				in.Delim('[')
				if out.Resources == nil {
					if !in.IsDelim(']') {
						out.Resources = make([]*APIResource, 0, 8)
					} else {
						out.Resources = []*APIResource{}
					}
				} else {
					out.Resources = (out.Resources)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *APIResource
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(APIResource)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Resources = append(out.Resources, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonBd8685ecEncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(out *jwriter.Writer, in APIResourceList) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	{
		const prefix string = ",\"groupVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.GroupVersion == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.GroupVersion))
		}
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	{
		const prefix string = ",\"resources\":"
		out.RawString(prefix)
		if in.Resources == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Resources {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v APIResourceList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBd8685ecEncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v APIResourceList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBd8685ecEncodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *APIResourceList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBd8685ecDecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *APIResourceList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBd8685ecDecodeGithubComKubewardenK8sObjectsApimachineryPkgApisMetaV1(l, v)
}
