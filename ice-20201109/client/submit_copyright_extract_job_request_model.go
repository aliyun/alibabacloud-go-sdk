// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightExtractJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *SubmitCopyrightExtractJobRequestInput) *SubmitCopyrightExtractJobRequest
	GetInput() *SubmitCopyrightExtractJobRequestInput
	SetParams(v string) *SubmitCopyrightExtractJobRequest
	GetParams() *string
	SetUserData(v string) *SubmitCopyrightExtractJobRequest
	GetUserData() *string
}

type SubmitCopyrightExtractJobRequest struct {
	// The source video file from which to extract the watermark.
	//
	// > The OSS object or media asset must reside in the same region as the IMS service region.
	//
	// This parameter is required.
	Input *SubmitCopyrightExtractJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// Additional parameters for the watermark job, provided as a JSON string. Supported parameter:
	//
	// 	- algoType: The algorithm type. Defaults to v1. The extraction algorithm must match the one used for embedding.
	//
	//     	- v1: Copyright watermark extraction algorithm for long videos.
	//
	//     	- v2: Copyright watermark extraction algorithm for short videos.
	//
	// example:
	//
	// {"algoType":"v2"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The custom data, which can be up to 1,024 bytes in size.
	//
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitCopyrightExtractJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractJobRequest) GetInput() *SubmitCopyrightExtractJobRequestInput {
	return s.Input
}

func (s *SubmitCopyrightExtractJobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitCopyrightExtractJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitCopyrightExtractJobRequest) SetInput(v *SubmitCopyrightExtractJobRequestInput) *SubmitCopyrightExtractJobRequest {
	s.Input = v
	return s
}

func (s *SubmitCopyrightExtractJobRequest) SetParams(v string) *SubmitCopyrightExtractJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitCopyrightExtractJobRequest) SetUserData(v string) *SubmitCopyrightExtractJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCopyrightExtractJobRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitCopyrightExtractJobRequestInput struct {
	// The specific information for the input file, which can be an OSS URL or a media asset ID. OSS URL formats:
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
	// 1.  OSS: an OSS object.
	//
	// 2.  Media: a media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitCopyrightExtractJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightExtractJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitCopyrightExtractJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitCopyrightExtractJobRequestInput) SetMedia(v string) *SubmitCopyrightExtractJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitCopyrightExtractJobRequestInput) SetType(v string) *SubmitCopyrightExtractJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitCopyrightExtractJobRequestInput) Validate() error {
	return dara.Validate(s)
}
