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
	// This parameter is required.
	Input *SubmitTraceExtractJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// example:
	//
	// {"m3u8Type":"v1"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
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
	return dara.Validate(s)
}

type SubmitTraceExtractJobRequestInput struct {
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
