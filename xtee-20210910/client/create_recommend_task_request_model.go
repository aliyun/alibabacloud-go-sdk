// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecommendTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateRecommendTaskRequest
	GetLang() *string
	SetName(v string) *CreateRecommendTaskRequest
	GetName() *string
	SetRegId(v string) *CreateRecommendTaskRequest
	GetRegId() *string
	SetSampleId(v int64) *CreateRecommendTaskRequest
	GetSampleId() *int64
	SetVariablesStr(v string) *CreateRecommendTaskRequest
	GetVariablesStr() *string
	SetVelocitiesStr(v string) *CreateRecommendTaskRequest
	GetVelocitiesStr() *string
}

type CreateRecommendTaskRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Name
	//
	// This parameter is required.
	//
	// example:
	//
	// 注册样本
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5467
	SampleId *int64 `json:"sampleId,omitempty" xml:"sampleId,omitempty"`
	// Variables to be calculated, variables
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"1112\\",\\"1113\\"]
	VariablesStr *string `json:"variablesStr,omitempty" xml:"variablesStr,omitempty"`
	// Indicator effect
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"dxkkLw8fe18\\",\\"dxUxSCM26d7\\"]
	VelocitiesStr *string `json:"velocitiesStr,omitempty" xml:"velocitiesStr,omitempty"`
}

func (s CreateRecommendTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecommendTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRecommendTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateRecommendTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateRecommendTaskRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateRecommendTaskRequest) GetSampleId() *int64 {
	return s.SampleId
}

func (s *CreateRecommendTaskRequest) GetVariablesStr() *string {
	return s.VariablesStr
}

func (s *CreateRecommendTaskRequest) GetVelocitiesStr() *string {
	return s.VelocitiesStr
}

func (s *CreateRecommendTaskRequest) SetLang(v string) *CreateRecommendTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateRecommendTaskRequest) SetName(v string) *CreateRecommendTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateRecommendTaskRequest) SetRegId(v string) *CreateRecommendTaskRequest {
	s.RegId = &v
	return s
}

func (s *CreateRecommendTaskRequest) SetSampleId(v int64) *CreateRecommendTaskRequest {
	s.SampleId = &v
	return s
}

func (s *CreateRecommendTaskRequest) SetVariablesStr(v string) *CreateRecommendTaskRequest {
	s.VariablesStr = &v
	return s
}

func (s *CreateRecommendTaskRequest) SetVelocitiesStr(v string) *CreateRecommendTaskRequest {
	s.VelocitiesStr = &v
	return s
}

func (s *CreateRecommendTaskRequest) Validate() error {
	return dara.Validate(s)
}
