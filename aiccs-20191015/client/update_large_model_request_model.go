// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLargeModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *UpdateLargeModelRequest
	GetAuthCode() *string
	SetBaseModel(v []*string) *UpdateLargeModelRequest
	GetBaseModel() []*string
	SetModelCode(v string) *UpdateLargeModelRequest
	GetModelCode() *string
	SetModelName(v string) *UpdateLargeModelRequest
	GetModelName() *string
	SetModelUrl(v string) *UpdateLargeModelRequest
	GetModelUrl() *string
	SetOwnerId(v int64) *UpdateLargeModelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateLargeModelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateLargeModelRequest
	GetResourceOwnerId() *int64
	SetTemperature(v float64) *UpdateLargeModelRequest
	GetTemperature() *float64
	SetTopK(v int64) *UpdateLargeModelRequest
	GetTopK() *int64
	SetTopP(v float64) *UpdateLargeModelRequest
	GetTopP() *float64
}

type UpdateLargeModelRequest struct {
	// 授权码
	//
	// example:
	//
	// 示例值示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// 基础模型
	BaseModel []*string `json:"BaseModel,omitempty" xml:"BaseModel,omitempty" type:"Repeated"`
	// 模型编码
	//
	// example:
	//
	// Test Model Name
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// 模型名称
	//
	// example:
	//
	// 示例值
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// 模型地址
	//
	// example:
	//
	// 示例值示例值示例值
	ModelUrl             *string `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 温度
	//
	// example:
	//
	// 49.29
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// topK
	//
	// example:
	//
	// 20
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// topP
	//
	// example:
	//
	// 84.38427
	TopP *float64 `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s UpdateLargeModelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLargeModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateLargeModelRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *UpdateLargeModelRequest) GetBaseModel() []*string {
	return s.BaseModel
}

func (s *UpdateLargeModelRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *UpdateLargeModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateLargeModelRequest) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *UpdateLargeModelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLargeModelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateLargeModelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateLargeModelRequest) GetTemperature() *float64 {
	return s.Temperature
}

func (s *UpdateLargeModelRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *UpdateLargeModelRequest) GetTopP() *float64 {
	return s.TopP
}

func (s *UpdateLargeModelRequest) SetAuthCode(v string) *UpdateLargeModelRequest {
	s.AuthCode = &v
	return s
}

func (s *UpdateLargeModelRequest) SetBaseModel(v []*string) *UpdateLargeModelRequest {
	s.BaseModel = v
	return s
}

func (s *UpdateLargeModelRequest) SetModelCode(v string) *UpdateLargeModelRequest {
	s.ModelCode = &v
	return s
}

func (s *UpdateLargeModelRequest) SetModelName(v string) *UpdateLargeModelRequest {
	s.ModelName = &v
	return s
}

func (s *UpdateLargeModelRequest) SetModelUrl(v string) *UpdateLargeModelRequest {
	s.ModelUrl = &v
	return s
}

func (s *UpdateLargeModelRequest) SetOwnerId(v int64) *UpdateLargeModelRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLargeModelRequest) SetResourceOwnerAccount(v string) *UpdateLargeModelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateLargeModelRequest) SetResourceOwnerId(v int64) *UpdateLargeModelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLargeModelRequest) SetTemperature(v float64) *UpdateLargeModelRequest {
	s.Temperature = &v
	return s
}

func (s *UpdateLargeModelRequest) SetTopK(v int64) *UpdateLargeModelRequest {
	s.TopK = &v
	return s
}

func (s *UpdateLargeModelRequest) SetTopP(v float64) *UpdateLargeModelRequest {
	s.TopP = &v
	return s
}

func (s *UpdateLargeModelRequest) Validate() error {
	return dara.Validate(s)
}
