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
	// The OSS destination for the output M3U8 file.
	//
	// > The OSS bucket must be in the same region as your MPS service.
	//
	// This parameter is required.
	Output *SubmitTraceM3u8JobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// A JSON string that contains parameters for the watermarking job. The following parameter is supported:
	//
	// - `m3u8Type`: The algorithm type. The default value is `v1`.
	//
	//   - `v1`: Generates an M3U8 file that uses an absolute path. The file can be played directly. The signature is valid for 24 hours. After expiration, you must submit a new job to get a new M3U8 file.
	//
	//   - `v2`: Generates an M3U8 file that uses a relative path. This file must be stored in the same directory as the TS files.
	//
	// example:
	//
	// {"m3u8Type":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The watermark content to embed.
	//
	// example:
	//
	// Trace watermark test
	Trace *string `json:"Trace,omitempty" xml:"Trace,omitempty"`
	// The media ID of the processed A/B stream for video watermarking for tracing. This ID is returned in the response when you submit the A/B stream job.
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
	// The output file path. Only OSS paths are supported. You can use one of the following formats:
	//
	// 1\\. `oss://bucket/object`
	//
	// 2\\. `http(s)://bucket.oss-[regionId].aliyuncs.com/object`. In these formats, `bucket` specifies the name of an OSS bucket in the same region as your project, and `object` specifies the file path.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output destination. Valid values:
	//
	// 1. `OSS`: an OSS file location.
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
