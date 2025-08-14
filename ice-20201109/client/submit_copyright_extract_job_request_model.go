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
	// This parameter is required.
	Input *SubmitCopyrightExtractJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// example:
	//
	// {"algoType":"v2"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
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
	return dara.Validate(s)
}

type SubmitCopyrightExtractJobRequestInput struct {
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
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
