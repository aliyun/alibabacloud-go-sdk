// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceAbJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCipherBase64ed(v string) *SubmitTraceAbJobRequest
	GetCipherBase64ed() *string
	SetInput(v *SubmitTraceAbJobRequestInput) *SubmitTraceAbJobRequest
	GetInput() *SubmitTraceAbJobRequestInput
	SetLevel(v int64) *SubmitTraceAbJobRequest
	GetLevel() *int64
	SetOutput(v *SubmitTraceAbJobRequestOutput) *SubmitTraceAbJobRequest
	GetOutput() *SubmitTraceAbJobRequestOutput
	SetStartTime(v int64) *SubmitTraceAbJobRequest
	GetStartTime() *int64
	SetTotalTime(v int64) *SubmitTraceAbJobRequest
	GetTotalTime() *int64
	SetUserData(v string) *SubmitTraceAbJobRequest
	GetUserData() *string
}

type SubmitTraceAbJobRequest struct {
	// The Base64-encoded encryption key.
	//
	// example:
	//
	// Qh6OdgIMcliQSI1fReOw****
	CipherBase64ed *string `json:"CipherBase64ed,omitempty" xml:"CipherBase64ed,omitempty"`
	// The input video for the A/B stream forensic watermarking job.
	//
	// > - The Object Storage Service (OSS) file or media asset must be in the same region where Intelligent Media Services (IMS) is deployed.
	//
	// >
	//
	// > - This operation supports only videos that are three minutes or longer. Using a shorter video may cause the API call to fail or produce no output.
	//
	// This parameter is required.
	Input *SubmitTraceAbJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The watermark level, which specifies the embedding channel. Valid values: `0` (U channel), `1` (UV channels), and `2` (YUV channels).
	//
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The output location for the A/B stream job. This must be an OSS directory.
	//
	// This parameter is required.
	Output *SubmitTraceAbJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The start time for watermark embedding, in seconds.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total duration for watermark embedding, in seconds.
	//
	// example:
	//
	// 360
	TotalTime *int64 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// User data to include in the request. The maximum length is 1,024 bytes.
	//
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTraceAbJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceAbJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbJobRequest) GetCipherBase64ed() *string {
	return s.CipherBase64ed
}

func (s *SubmitTraceAbJobRequest) GetInput() *SubmitTraceAbJobRequestInput {
	return s.Input
}

func (s *SubmitTraceAbJobRequest) GetLevel() *int64 {
	return s.Level
}

func (s *SubmitTraceAbJobRequest) GetOutput() *SubmitTraceAbJobRequestOutput {
	return s.Output
}

func (s *SubmitTraceAbJobRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SubmitTraceAbJobRequest) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *SubmitTraceAbJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTraceAbJobRequest) SetCipherBase64ed(v string) *SubmitTraceAbJobRequest {
	s.CipherBase64ed = &v
	return s
}

func (s *SubmitTraceAbJobRequest) SetInput(v *SubmitTraceAbJobRequestInput) *SubmitTraceAbJobRequest {
	s.Input = v
	return s
}

func (s *SubmitTraceAbJobRequest) SetLevel(v int64) *SubmitTraceAbJobRequest {
	s.Level = &v
	return s
}

func (s *SubmitTraceAbJobRequest) SetOutput(v *SubmitTraceAbJobRequestOutput) *SubmitTraceAbJobRequest {
	s.Output = v
	return s
}

func (s *SubmitTraceAbJobRequest) SetStartTime(v int64) *SubmitTraceAbJobRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitTraceAbJobRequest) SetTotalTime(v int64) *SubmitTraceAbJobRequest {
	s.TotalTime = &v
	return s
}

func (s *SubmitTraceAbJobRequest) SetUserData(v string) *SubmitTraceAbJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTraceAbJobRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTraceAbJobRequestInput struct {
	// The location of the input file. This can be an OSS URL or a media asset ID.
	//
	// The supported OSS URL formats are:
	//
	// 1\\. `oss://<bucket>/<object>`
	//
	// 2\\. `http(s)://<bucket>.oss-[regionId].aliyuncs.com/<object>`
	//
	// In these formats, `<bucket>` is the name of the OSS bucket, which must be in the same region as your project, and `<object>` is the path to the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid values:
	//
	// 1. `OSS`: The file is specified by an OSS URL.
	//
	// 2. `Media`: The file is specified by a media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTraceAbJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceAbJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitTraceAbJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitTraceAbJobRequestInput) SetMedia(v string) *SubmitTraceAbJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitTraceAbJobRequestInput) SetType(v string) *SubmitTraceAbJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitTraceAbJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitTraceAbJobRequestOutput struct {
	// The output path. Specify an OSS directory URL or a media asset ID. If you specify an OSS URL, use one of the following formats for the output directory:
	//
	// 1\\. `oss://<bucket>/<dir>/`
	//
	// 2\\. `http(s)://<bucket>.oss-[regionId].aliyuncs.com/<dir>/`
	//
	// In these formats, `<bucket>` is the name of the OSS bucket, which must be in the same region as your project, and `<dir>` is the output directory path.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/dir/
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output object.
	//
	// Valid values:
	//
	// - `OSS`: An OSS object.
	//
	// - `Media`: A media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTraceAbJobRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceAbJobRequestOutput) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbJobRequestOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitTraceAbJobRequestOutput) GetType() *string {
	return s.Type
}

func (s *SubmitTraceAbJobRequestOutput) SetMedia(v string) *SubmitTraceAbJobRequestOutput {
	s.Media = &v
	return s
}

func (s *SubmitTraceAbJobRequestOutput) SetType(v string) *SubmitTraceAbJobRequestOutput {
	s.Type = &v
	return s
}

func (s *SubmitTraceAbJobRequestOutput) Validate() error {
	return dara.Validate(s)
}
