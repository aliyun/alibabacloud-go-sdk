// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SubmitCopyrightJobRequest
	GetDescription() *string
	SetInput(v *SubmitCopyrightJobRequestInput) *SubmitCopyrightJobRequest
	GetInput() *SubmitCopyrightJobRequestInput
	SetLevel(v int64) *SubmitCopyrightJobRequest
	GetLevel() *int64
	SetMessage(v string) *SubmitCopyrightJobRequest
	GetMessage() *string
	SetOutput(v *SubmitCopyrightJobRequestOutput) *SubmitCopyrightJobRequest
	GetOutput() *SubmitCopyrightJobRequestOutput
	SetParams(v string) *SubmitCopyrightJobRequest
	GetParams() *string
	SetStartTime(v int64) *SubmitCopyrightJobRequest
	GetStartTime() *int64
	SetTotalTime(v int64) *SubmitCopyrightJobRequest
	GetTotalTime() *int64
	SetUserData(v string) *SubmitCopyrightJobRequest
	GetUserData() *string
}

type SubmitCopyrightJobRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"example-bucket","Location":"oss-cn-shanghai","Object":"example.mp4"}
	Input *SubmitCopyrightJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// This parameter is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"example-bucket","Location":"oss-cn-shanghai","Object":"example_result.mp4"}
	Output *SubmitCopyrightJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// example:
	//
	// {"algoType":"v2"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 10
	TotalTime *int64 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// example:
	//
	// 123
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitCopyrightJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitCopyrightJobRequest) GetInput() *SubmitCopyrightJobRequestInput {
	return s.Input
}

func (s *SubmitCopyrightJobRequest) GetLevel() *int64 {
	return s.Level
}

func (s *SubmitCopyrightJobRequest) GetMessage() *string {
	return s.Message
}

func (s *SubmitCopyrightJobRequest) GetOutput() *SubmitCopyrightJobRequestOutput {
	return s.Output
}

func (s *SubmitCopyrightJobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitCopyrightJobRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SubmitCopyrightJobRequest) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *SubmitCopyrightJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitCopyrightJobRequest) SetDescription(v string) *SubmitCopyrightJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetInput(v *SubmitCopyrightJobRequestInput) *SubmitCopyrightJobRequest {
	s.Input = v
	return s
}

func (s *SubmitCopyrightJobRequest) SetLevel(v int64) *SubmitCopyrightJobRequest {
	s.Level = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetMessage(v string) *SubmitCopyrightJobRequest {
	s.Message = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetOutput(v *SubmitCopyrightJobRequestOutput) *SubmitCopyrightJobRequest {
	s.Output = v
	return s
}

func (s *SubmitCopyrightJobRequest) SetParams(v string) *SubmitCopyrightJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetStartTime(v int64) *SubmitCopyrightJobRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetTotalTime(v int64) *SubmitCopyrightJobRequest {
	s.TotalTime = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetUserData(v string) *SubmitCopyrightJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCopyrightJobRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitCopyrightJobRequestInput struct {
	// This parameter is required.
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitCopyrightJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitCopyrightJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitCopyrightJobRequestInput) SetMedia(v string) *SubmitCopyrightJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitCopyrightJobRequestInput) SetType(v string) *SubmitCopyrightJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitCopyrightJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitCopyrightJobRequestOutput struct {
	// This parameter is required.
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitCopyrightJobRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobRequestOutput) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobRequestOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitCopyrightJobRequestOutput) GetType() *string {
	return s.Type
}

func (s *SubmitCopyrightJobRequestOutput) SetMedia(v string) *SubmitCopyrightJobRequestOutput {
	s.Media = &v
	return s
}

func (s *SubmitCopyrightJobRequestOutput) SetType(v string) *SubmitCopyrightJobRequestOutput {
	s.Type = &v
	return s
}

func (s *SubmitCopyrightJobRequestOutput) Validate() error {
	return dara.Validate(s)
}
