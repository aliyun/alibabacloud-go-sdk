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
	// The source video file from which to extract the watermark.
	//
	// > The OSS object or media asset must reside in the same region as the IMS service region.
	//
	// This parameter is required.
	Input *SubmitTraceExtractJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// Additional parameters for the watermark job, provided as a JSON string. Supported parameter:
	//
	// 	- m3u8Type: The extraction algorithm type. Defaults to v1.
	//
	//     	- v1: Extracts from an M3U8 with absolute paths.
	//
	//     	- v2: Extracts from an M3U8 with relative paths.
	//
	// example:
	//
	// {"m3u8Type":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The custom data, which can be up to 1,024 bytes in size.
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
	// The specific information for the source file, which can be an OSS URL or a media asset ID. OSS URL formats:
	//
	// 1\\. oss://bucket/object
	//
	// 2\\. http(s)://bucket.oss-[regionId].aliyuncs.com/object
	//
	// where bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object path in OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the source file. Valid values:
	//
	// 	- OSS: an OSS object.
	//
	// 	- Media: a media asset.
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
