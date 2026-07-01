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
	// The video file for watermark extraction.
	//
	// > - The region of the Object Storage Service (OSS) file or media asset must match the region of the current Intelligent Media Service (IMS) instance.
	//
	// This parameter is required.
	Input *SubmitCopyrightExtractJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The watermark job parameters, specified as a JSON string.
	//
	// - algoType: The algorithm type. Default value: v1. The extraction algorithm type must match the algorithm type used for embedding the watermark.
	//
	//   - v1: The copyright extraction algorithm for long-form videos.
	//
	//   - v2: The copyright extraction algorithm for short-form videos.
	//
	// example:
	//
	// {"algoType":"v2"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The user-defined data. The maximum length is 1,024 bytes.
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
	// Specifies the URL of an Object Storage Service (OSS) object or the ID of a media asset.
	//
	// OSS URLs can be in the following formats:
	//
	// 1\\. oss\\://bucket/object
	//
	// 2\\. http(s)://bucket.oss-[regionId].aliyuncs.com/object
	//
	// In these formats, `bucket` is the name of a bucket in the same region as your IMS instance, and `object` is the file path.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid values:
	//
	// 1. OSS: The URL of a file in Object Storage Service (OSS).
	//
	// 2. Media: The ID of a media asset.
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
