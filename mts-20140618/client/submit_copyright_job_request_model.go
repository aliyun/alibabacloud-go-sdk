// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallBack(v string) *SubmitCopyrightJobRequest
	GetCallBack() *string
	SetDescription(v string) *SubmitCopyrightJobRequest
	GetDescription() *string
	SetInput(v string) *SubmitCopyrightJobRequest
	GetInput() *string
	SetLevel(v int64) *SubmitCopyrightJobRequest
	GetLevel() *int64
	SetMessage(v string) *SubmitCopyrightJobRequest
	GetMessage() *string
	SetOutput(v string) *SubmitCopyrightJobRequest
	GetOutput() *string
	SetParams(v string) *SubmitCopyrightJobRequest
	GetParams() *string
	SetStartTime(v int64) *SubmitCopyrightJobRequest
	GetStartTime() *int64
	SetTotalTime(v int64) *SubmitCopyrightJobRequest
	GetTotalTime() *int64
	SetUrl(v string) *SubmitCopyrightJobRequest
	GetUrl() *string
	SetUserData(v string) *SubmitCopyrightJobRequest
	GetUserData() *string
}

type SubmitCopyrightJobRequest struct {
	// example:
	//
	// http://example.com/callback
	CallBack    *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {"Bucket":"example-bucket","Location":"oss-cn-shanghai","Object":"example.mp4"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// 2
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// This parameter is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"example-bucket","Location":"oss-cn-shanghai","Object":"example_result.mp4"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
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
	// http://www.example.com/video/test.mp4
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

func (s *SubmitCopyrightJobRequest) GetCallBack() *string {
	return s.CallBack
}

func (s *SubmitCopyrightJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitCopyrightJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitCopyrightJobRequest) GetLevel() *int64 {
	return s.Level
}

func (s *SubmitCopyrightJobRequest) GetMessage() *string {
	return s.Message
}

func (s *SubmitCopyrightJobRequest) GetOutput() *string {
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

func (s *SubmitCopyrightJobRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitCopyrightJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitCopyrightJobRequest) SetCallBack(v string) *SubmitCopyrightJobRequest {
	s.CallBack = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetDescription(v string) *SubmitCopyrightJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetInput(v string) *SubmitCopyrightJobRequest {
	s.Input = &v
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

func (s *SubmitCopyrightJobRequest) SetOutput(v string) *SubmitCopyrightJobRequest {
	s.Output = &v
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

func (s *SubmitCopyrightJobRequest) SetUrl(v string) *SubmitCopyrightJobRequest {
	s.Url = &v
	return s
}

func (s *SubmitCopyrightJobRequest) SetUserData(v string) *SubmitCopyrightJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCopyrightJobRequest) Validate() error {
	return dara.Validate(s)
}
