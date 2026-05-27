// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLargeModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *AddLargeModelShrinkRequest
	GetAuthCode() *string
	SetBaseModelShrink(v string) *AddLargeModelShrinkRequest
	GetBaseModelShrink() *string
	SetModelName(v string) *AddLargeModelShrinkRequest
	GetModelName() *string
	SetModelUrl(v string) *AddLargeModelShrinkRequest
	GetModelUrl() *string
	SetOwnerId(v int64) *AddLargeModelShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddLargeModelShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddLargeModelShrinkRequest
	GetResourceOwnerId() *int64
	SetTemperature(v float64) *AddLargeModelShrinkRequest
	GetTemperature() *float64
	SetTopK(v int64) *AddLargeModelShrinkRequest
	GetTopK() *int64
	SetTopP(v float64) *AddLargeModelShrinkRequest
	GetTopP() *float64
}

type AddLargeModelShrinkRequest struct {
	// 授权码
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// 基础模型
	BaseModelShrink *string `json:"BaseModel,omitempty" xml:"BaseModel,omitempty"`
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

func (s AddLargeModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLargeModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddLargeModelShrinkRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *AddLargeModelShrinkRequest) GetBaseModelShrink() *string {
	return s.BaseModelShrink
}

func (s *AddLargeModelShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *AddLargeModelShrinkRequest) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *AddLargeModelShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLargeModelShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddLargeModelShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddLargeModelShrinkRequest) GetTemperature() *float64 {
	return s.Temperature
}

func (s *AddLargeModelShrinkRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *AddLargeModelShrinkRequest) GetTopP() *float64 {
	return s.TopP
}

func (s *AddLargeModelShrinkRequest) SetAuthCode(v string) *AddLargeModelShrinkRequest {
	s.AuthCode = &v
	return s
}

func (s *AddLargeModelShrinkRequest) SetBaseModelShrink(v string) *AddLargeModelShrinkRequest {
	s.BaseModelShrink = &v
	return s
}

func (s *AddLargeModelShrinkRequest) SetModelName(v string) *AddLargeModelShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *AddLargeModelShrinkRequest) SetModelUrl(v string) *AddLargeModelShrinkRequest {
	s.ModelUrl = &v
	return s
}

func (s *AddLargeModelShrinkRequest) SetOwnerId(v int64) *AddLargeModelShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLargeModelShrinkRequest) SetResourceOwnerAccount(v string) *AddLargeModelShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddLargeModelShrinkRequest) SetResourceOwnerId(v int64) *AddLargeModelShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddLargeModelShrinkRequest) SetTemperature(v float64) *AddLargeModelShrinkRequest {
	s.Temperature = &v
	return s
}

func (s *AddLargeModelShrinkRequest) SetTopK(v int64) *AddLargeModelShrinkRequest {
	s.TopK = &v
	return s
}

func (s *AddLargeModelShrinkRequest) SetTopP(v float64) *AddLargeModelShrinkRequest {
	s.TopP = &v
	return s
}

func (s *AddLargeModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
