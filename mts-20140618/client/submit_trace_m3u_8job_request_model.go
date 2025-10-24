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
	SetMediaId(v string) *SubmitTraceM3u8JobRequest
	GetMediaId() *string
	SetOutput(v string) *SubmitTraceM3u8JobRequest
	GetOutput() *string
	SetParams(v string) *SubmitTraceM3u8JobRequest
	GetParams() *string
	SetTrace(v string) *SubmitTraceM3u8JobRequest
	GetTrace() *string
}

type SubmitTraceM3u8JobRequest struct {
	// example:
	//
	// https://cipher.abc.com
	KeyUri *string `json:"KeyUri,omitempty" xml:"KeyUri,omitempty"`
	// example:
	//
	// 437bd2b516ffda105d07b12a9a82****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// {"Bucket":"exampleBucket","Location":"oss-cn-shanghai","Object":"out.m3u8"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// {"m3u8Type":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	Trace  *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
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

func (s *SubmitTraceM3u8JobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitTraceM3u8JobRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitTraceM3u8JobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitTraceM3u8JobRequest) GetTrace() *string {
	return s.Trace
}

func (s *SubmitTraceM3u8JobRequest) SetKeyUri(v string) *SubmitTraceM3u8JobRequest {
	s.KeyUri = &v
	return s
}

func (s *SubmitTraceM3u8JobRequest) SetMediaId(v string) *SubmitTraceM3u8JobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitTraceM3u8JobRequest) SetOutput(v string) *SubmitTraceM3u8JobRequest {
	s.Output = &v
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

func (s *SubmitTraceM3u8JobRequest) Validate() error {
	return dara.Validate(s)
}
