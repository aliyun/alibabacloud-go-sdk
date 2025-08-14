// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceM3u8JobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyUri(v string) *SubmitTraceM3u8JobShrinkRequest
	GetKeyUri() *string
	SetOutputShrink(v string) *SubmitTraceM3u8JobShrinkRequest
	GetOutputShrink() *string
	SetParams(v string) *SubmitTraceM3u8JobShrinkRequest
	GetParams() *string
	SetTrace(v string) *SubmitTraceM3u8JobShrinkRequest
	GetTrace() *string
	SetTraceMediaId(v string) *SubmitTraceM3u8JobShrinkRequest
	GetTraceMediaId() *string
}

type SubmitTraceM3u8JobShrinkRequest struct {
	// example:
	//
	// https://cipher.abc.com
	KeyUri *string `json:"KeyUri,omitempty" xml:"KeyUri,omitempty"`
	// This parameter is required.
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
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

func (s SubmitTraceM3u8JobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceM3u8JobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceM3u8JobShrinkRequest) GetKeyUri() *string {
	return s.KeyUri
}

func (s *SubmitTraceM3u8JobShrinkRequest) GetOutputShrink() *string {
	return s.OutputShrink
}

func (s *SubmitTraceM3u8JobShrinkRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitTraceM3u8JobShrinkRequest) GetTrace() *string {
	return s.Trace
}

func (s *SubmitTraceM3u8JobShrinkRequest) GetTraceMediaId() *string {
	return s.TraceMediaId
}

func (s *SubmitTraceM3u8JobShrinkRequest) SetKeyUri(v string) *SubmitTraceM3u8JobShrinkRequest {
	s.KeyUri = &v
	return s
}

func (s *SubmitTraceM3u8JobShrinkRequest) SetOutputShrink(v string) *SubmitTraceM3u8JobShrinkRequest {
	s.OutputShrink = &v
	return s
}

func (s *SubmitTraceM3u8JobShrinkRequest) SetParams(v string) *SubmitTraceM3u8JobShrinkRequest {
	s.Params = &v
	return s
}

func (s *SubmitTraceM3u8JobShrinkRequest) SetTrace(v string) *SubmitTraceM3u8JobShrinkRequest {
	s.Trace = &v
	return s
}

func (s *SubmitTraceM3u8JobShrinkRequest) SetTraceMediaId(v string) *SubmitTraceM3u8JobShrinkRequest {
	s.TraceMediaId = &v
	return s
}

func (s *SubmitTraceM3u8JobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
