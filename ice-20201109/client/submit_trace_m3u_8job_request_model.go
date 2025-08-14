// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceM3u8JobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyUri(v string) *SubmitTraceM3u8JobRequest
	GetKeyUri() *string
	SetOutput(v *SubmitTraceM3u8JobRequestOutput) *SubmitTraceM3u8JobRequest
	GetOutput() *SubmitTraceM3u8JobRequestOutput
	SetParams(v string) *SubmitTraceM3u8JobRequest
	GetParams() *string
	SetTrace(v string) *SubmitTraceM3u8JobRequest
	GetTrace() *string
	SetTraceMediaId(v string) *SubmitTraceM3u8JobRequest
	GetTraceMediaId() *string
}

type SubmitTraceM3u8JobRequest struct {
	// example:
	//
	// https://cipher.abc.com
	KeyUri *string `json:"KeyUri,omitempty" xml:"KeyUri,omitempty"`
	// This parameter is required.
	Output *SubmitTraceM3u8JobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// example:
	//
	// {"m3u8Type":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	Trace  *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
	// example:
	//
	// 437bd2b516ffda105d07b12a9a82****
	TraceMediaId *string `json:"TraceMediaId,omitempty" xml:"TraceMediaId,omitempty"`
}

func (s SubmitTraceM3u8JobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceM3u8JobRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceM3u8JobRequest) GetKeyUri() *string {
	return s.KeyUri
}

func (s *SubmitTraceM3u8JobRequest) GetOutput() *SubmitTraceM3u8JobRequestOutput {
	return s.Output
}

func (s *SubmitTraceM3u8JobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitTraceM3u8JobRequest) GetTrace() *string {
	return s.Trace
}

func (s *SubmitTraceM3u8JobRequest) GetTraceMediaId() *string {
	return s.TraceMediaId
}

func (s *SubmitTraceM3u8JobRequest) SetKeyUri(v string) *SubmitTraceM3u8JobRequest {
	s.KeyUri = &v
	return s
}

func (s *SubmitTraceM3u8JobRequest) SetOutput(v *SubmitTraceM3u8JobRequestOutput) *SubmitTraceM3u8JobRequest {
	s.Output = v
	return s
}

func (s *SubmitTraceM3u8JobRequest) SetParams(v string) *SubmitTraceM3u8JobRequest {
	s.Params = &v
	return s
}

func (s *SubmitTraceM3u8JobRequest) SetTrace(v string) *SubmitTraceM3u8JobRequest {
	s.Trace = &v
	return s
}

func (s *SubmitTraceM3u8JobRequest) SetTraceMediaId(v string) *SubmitTraceM3u8JobRequest {
	s.TraceMediaId = &v
	return s
}

func (s *SubmitTraceM3u8JobRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitTraceM3u8JobRequestOutput struct {
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTraceM3u8JobRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceM3u8JobRequestOutput) GoString() string {
	return s.String()
}

func (s *SubmitTraceM3u8JobRequestOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitTraceM3u8JobRequestOutput) GetType() *string {
	return s.Type
}

func (s *SubmitTraceM3u8JobRequestOutput) SetMedia(v string) *SubmitTraceM3u8JobRequestOutput {
	s.Media = &v
	return s
}

func (s *SubmitTraceM3u8JobRequestOutput) SetType(v string) *SubmitTraceM3u8JobRequestOutput {
	s.Type = &v
	return s
}

func (s *SubmitTraceM3u8JobRequestOutput) Validate() error {
	return dara.Validate(s)
}
