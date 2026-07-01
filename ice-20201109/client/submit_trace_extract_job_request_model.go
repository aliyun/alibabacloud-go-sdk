// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceExtractJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *SubmitTraceExtractJobRequestInput) *SubmitTraceExtractJobRequest
	GetInput() *SubmitTraceExtractJobRequestInput
	SetParams(v string) *SubmitTraceExtractJobRequest
	GetParams() *string
	SetUserData(v string) *SubmitTraceExtractJobRequest
	GetUserData() *string
}

type SubmitTraceExtractJobRequest struct {
	// The input video from which to extract the watermark.
	//
	// > - The OSS object or media asset must be in the same region as your IMS service.
	//
	// This parameter is required.
	Input *SubmitTraceExtractJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// Extraction job parameters, specified as a JSON string. The following parameters are supported:
	//
	// - `m3u8Type`: The algorithm type. The default value is `v1`.
	//
	//   - `v1`: Extracts an m3u8 playlist with absolute paths.
	//
	//   - `v2`: Extracts an m3u8 playlist with relative paths.
	//
	// example:
	//
	// {"m3u8Type":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The user-defined data. Maximum length: 1,024 bytes.
	//
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTraceExtractJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractJobRequest) GetInput() *SubmitTraceExtractJobRequestInput {
	return s.Input
}

func (s *SubmitTraceExtractJobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitTraceExtractJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTraceExtractJobRequest) SetInput(v *SubmitTraceExtractJobRequestInput) *SubmitTraceExtractJobRequest {
	s.Input = v
	return s
}

func (s *SubmitTraceExtractJobRequest) SetParams(v string) *SubmitTraceExtractJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitTraceExtractJobRequest) SetUserData(v string) *SubmitTraceExtractJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTraceExtractJobRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTraceExtractJobRequestInput struct {
	// The input source. Specify an OSS object URL or a media asset ID.
	//
	// An OSS object URL can be in one of the following formats:
	//
	// 1\\. oss\\://bucket/object
	//
	// In these formats, `bucket` is the name of an OSS bucket in the same region as your IMS service, and `object` is the path of the OSS object.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The input type. Valid values:
	//
	// - OSS: An OSS object URL.
	//
	// - Media: A media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTraceExtractJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceExtractJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitTraceExtractJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitTraceExtractJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitTraceExtractJobRequestInput) SetMedia(v string) *SubmitTraceExtractJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitTraceExtractJobRequestInput) SetType(v string) *SubmitTraceExtractJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitTraceExtractJobRequestInput) Validate() error {
	return dara.Validate(s)
}
