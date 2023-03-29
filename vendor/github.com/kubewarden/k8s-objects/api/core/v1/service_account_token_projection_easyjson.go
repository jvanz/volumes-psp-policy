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

func easyjson91088bddDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *ServiceAccountTokenProjection) {
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
		case "audience":
			out.Audience = string(in.String())
		case "expirationSeconds":
			out.ExpirationSeconds = int64(in.Int64())
		case "path":
			if in.IsNull() {
				in.Skip()
				out.Path = nil
			} else {
				if out.Path == nil {
					out.Path = new(string)
				}
				*out.Path = string(in.String())
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
func easyjson91088bddEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in ServiceAccountTokenProjection) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Audience != "" {
		const prefix string = ",\"audience\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Audience))
	}
	if in.ExpirationSeconds != 0 {
		const prefix string = ",\"expirationSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ExpirationSeconds))
	}
	{
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Path == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Path))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceAccountTokenProjection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson91088bddEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceAccountTokenProjection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson91088bddEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceAccountTokenProjection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson91088bddDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceAccountTokenProjection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson91088bddDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}