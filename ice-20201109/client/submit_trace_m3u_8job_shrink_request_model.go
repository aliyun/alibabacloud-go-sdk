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
	// The URI of the key server.
	//
	// example:
	//
	// https://cipher.abc.com
	KeyUri *string `json:"KeyUri,omitempty" xml:"KeyUri,omitempty"`
	// The OSS URL of the output M3U8 file.
	//
	// > The OSS bucket must reside in the same region as the service region.
	//
	// This parameter is required.
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// Additional parameters for the watermark job, provided as a JSON string. Supported parameter:
	//
	// 	- m3u8Type: The type of M3U8 to generate. Defaults to v1.
	//
	//     	- v1: Generates an M3U8 with absolute paths, playable directly. The signed URL for access is valid for 24 hours. If you need to use it after expiration, you must call this API again.
	//
	//     	- v2: Generates an M3U8 with relative paths. It must be placed in the same directory as the TS segment files to be playable.
	//
	// example:
	//
	// {"m3u8Type":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The specific trace watermark information.
	Trace *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
	// The media ID for the trace watermark. You can obtain this from the response of the SubmitTraceAbJob operation.
	//
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
