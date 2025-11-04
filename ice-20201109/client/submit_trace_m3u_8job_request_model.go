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
	Output *SubmitTraceM3u8JobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
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
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTraceM3u8JobRequestOutput struct {
	// The OSS path where the output file is saved. You can specify the path in one of the following formats:
	//
	// 1\\. oss://bucket/object
	//
	// 2\\. http(s)://bucket.oss-[regionId].aliyuncs.com/object where bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object path in OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output file. Valid value:
	//
	// 1.  OSS: an OSS object.
	//
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
