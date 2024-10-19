// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson3e29e0f0DecodeGithubComMrPixikUrlShortenerInternalAppModels(in *jlexer.Lexer, out *URLResponse) {
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
		case "result":
			out.URL = string(in.String())
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
func easyjson3e29e0f0EncodeGithubComMrPixikUrlShortenerInternalAppModels(out *jwriter.Writer, in URLResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix[1:])
		out.String(string(in.URL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v URLResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e29e0f0EncodeGithubComMrPixikUrlShortenerInternalAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v URLResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e29e0f0EncodeGithubComMrPixikUrlShortenerInternalAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *URLResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e29e0f0DecodeGithubComMrPixikUrlShortenerInternalAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *URLResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e29e0f0DecodeGithubComMrPixikUrlShortenerInternalAppModels(l, v)
}
func easyjson3e29e0f0DecodeGithubComMrPixikUrlShortenerInternalAppModels1(in *jlexer.Lexer, out *URLRequest) {
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
		case "url":
			out.URL = string(in.String())
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
func easyjson3e29e0f0EncodeGithubComMrPixikUrlShortenerInternalAppModels1(out *jwriter.Writer, in URLRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix[1:])
		out.String(string(in.URL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v URLRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e29e0f0EncodeGithubComMrPixikUrlShortenerInternalAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v URLRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e29e0f0EncodeGithubComMrPixikUrlShortenerInternalAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *URLRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e29e0f0DecodeGithubComMrPixikUrlShortenerInternalAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *URLRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e29e0f0DecodeGithubComMrPixikUrlShortenerInternalAppModels1(l, v)
}
