// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLargeModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *AddLargeModelRequest
	GetAuthCode() *string
	SetBaseModel(v []*string) *AddLargeModelRequest
	GetBaseModel() []*string
	SetModelName(v string) *AddLargeModelRequest
	GetModelName() *string
	SetModelUrl(v string) *AddLargeModelRequest
	GetModelUrl() *string
	SetOwnerId(v int64) *AddLargeModelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddLargeModelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddLargeModelRequest
	GetResourceOwnerId() *int64
	SetTemperature(v float64) *AddLargeModelRequest
	GetTemperature() *float64
	SetTopK(v int64) *AddLargeModelRequest
	GetTopK() *int64
	SetTopP(v float64) *AddLargeModelRequest
	GetTopP() *float64
}

type AddLargeModelRequest struct {
	// 授权码
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// 基础模型
	BaseModel []*string `json:"BaseModel,omitempty" xml:"BaseModel,omitempty" type:"Repeated"`
	// 模型名称
	//
	// example:
	//
	// Test Model Name
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// 模型地址
	//
	// example:
	//
	// 示例值示例值
	ModelUrl             *string `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 温度
	//
	// example:
	//
	// 16.46
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// topK
	//
	// example:
	//
	// 87
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// topP
	//
	// example:
	//
	// 73.64386
	TopP *float64 `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s AddLargeModelRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLargeModelRequest) GoString() string {
	return s.String()
}

func (s *AddLargeModelRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *AddLargeModelRequest) GetBaseModel() []*string {
	return s.BaseModel
}

func (s *AddLargeModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *AddLargeModelRequest) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *AddLargeModelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLargeModelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddLargeModelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddLargeModelRequest) GetTemperature() *float64 {
	return s.Temperature
}

func (s *AddLargeModelRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *AddLargeModelRequest) GetTopP() *float64 {
	return s.TopP
}

func (s *AddLargeModelRequest) SetAuthCode(v string) *AddLargeModelRequest {
	s.AuthCode = &v
	return s
}

func (s *AddLargeModelRequest) SetBaseModel(v []*string) *AddLargeModelRequest {
	s.BaseModel = v
	return s
}

func (s *AddLargeModelRequest) SetModelName(v string) *AddLargeModelRequest {
	s.ModelName = &v
	return s
}

func (s *AddLargeModelRequest) SetModelUrl(v string) *AddLargeModelRequest {
	s.ModelUrl = &v
	return s
}

func (s *AddLargeModelRequest) SetOwnerId(v int64) *AddLargeModelRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLargeModelRequest) SetResourceOwnerAccount(v string) *AddLargeModelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddLargeModelRequest) SetResourceOwnerId(v int64) *AddLargeModelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddLargeModelRequest) SetTemperature(v float64) *AddLargeModelRequest {
	s.Temperature = &v
	return s
}

func (s *AddLargeModelRequest) SetTopK(v int64) *AddLargeModelRequest {
	s.TopK = &v
	return s
}

func (s *AddLargeModelRequest) SetTopP(v float64) *AddLargeModelRequest {
	s.TopP = &v
	return s
}

func (s *AddLargeModelRequest) Validate() error {
	return dara.Validate(s)
}
