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
	// example:
	//
	// Qh6OdgIMcliQSI1fReOw****
	CipherBase64ed *string `json:"CipherBase64ed,omitempty" xml:"CipherBase64ed,omitempty"`
	// This parameter is required.
	Input *SubmitTraceAbJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// This parameter is required.
	Output *SubmitTraceAbJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type SubmitTraceAbJobRequestInput struct {
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
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/dir/
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
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
