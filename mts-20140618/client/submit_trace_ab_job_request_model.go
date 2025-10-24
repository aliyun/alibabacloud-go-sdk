// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTraceAbJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallBack(v string) *SubmitTraceAbJobRequest
	GetCallBack() *string
	SetCipherBase64ed(v string) *SubmitTraceAbJobRequest
	GetCipherBase64ed() *string
	SetInput(v string) *SubmitTraceAbJobRequest
	GetInput() *string
	SetLevel(v int64) *SubmitTraceAbJobRequest
	GetLevel() *int64
	SetOutput(v string) *SubmitTraceAbJobRequest
	GetOutput() *string
	SetStartTime(v int64) *SubmitTraceAbJobRequest
	GetStartTime() *int64
	SetTotalTime(v int64) *SubmitTraceAbJobRequest
	GetTotalTime() *int64
	SetUrl(v string) *SubmitTraceAbJobRequest
	GetUrl() *string
	SetUserData(v string) *SubmitTraceAbJobRequest
	GetUserData() *string
}

type SubmitTraceAbJobRequest struct {
	// example:
	//
	// http://example.com/callback
	CallBack *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	// example:
	//
	// Qh6OdgIMcliQSI1fReOw****
	CipherBase64ed *string `json:"CipherBase64ed,omitempty" xml:"CipherBase64ed,omitempty"`
	// example:
	//
	// {"Bucket":"ivison-test","Location":"oss-cn-shanghai","Object":"test.mp4"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// 2
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"ivison-test","Location":"oss-cn-shanghai","Dir":"out/"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 360
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

func (s SubmitTraceAbJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTraceAbJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitTraceAbJobRequest) GetCallBack() *string {
	return s.CallBack
}

func (s *SubmitTraceAbJobRequest) GetCipherBase64ed() *string {
	return s.CipherBase64ed
}

func (s *SubmitTraceAbJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitTraceAbJobRequest) GetLevel() *int64 {
	return s.Level
}

func (s *SubmitTraceAbJobRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitTraceAbJobRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SubmitTraceAbJobRequest) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *SubmitTraceAbJobRequest) GetUrl() *string {
	return s.Url
}

func (s *SubmitTraceAbJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTraceAbJobRequest) SetCallBack(v string) *SubmitTraceAbJobRequest {
	s.CallBack = &v
	return s
}

func (s *SubmitTraceAbJobRequest) SetCipherBase64ed(v string) *SubmitTraceAbJobRequest {
	s.CipherBase64ed = &v
	return s
}

func (s *SubmitTraceAbJobRequest) SetInput(v string) *SubmitTraceAbJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitTraceAbJobRequest) SetLevel(v int64) *SubmitTraceAbJobRequest {
	s.Level = &v
	return s
}

func (s *SubmitTraceAbJobRequest) SetOutput(v string) *SubmitTraceAbJobRequest {
	s.Output = &v
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

func (s *SubmitTraceAbJobRequest) SetUrl(v string) *SubmitTraceAbJobRequest {
	s.Url = &v
	return s
}

func (s *SubmitTraceAbJobRequest) SetUserData(v string) *SubmitTraceAbJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTraceAbJobRequest) Validate() error {
	return dara.Validate(s)
}
