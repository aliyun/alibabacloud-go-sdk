// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLargeModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *UpdateLargeModelShrinkRequest
	GetAuthCode() *string
	SetBaseModelShrink(v string) *UpdateLargeModelShrinkRequest
	GetBaseModelShrink() *string
	SetModelCode(v string) *UpdateLargeModelShrinkRequest
	GetModelCode() *string
	SetModelName(v string) *UpdateLargeModelShrinkRequest
	GetModelName() *string
	SetModelUrl(v string) *UpdateLargeModelShrinkRequest
	GetModelUrl() *string
	SetOwnerId(v int64) *UpdateLargeModelShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateLargeModelShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateLargeModelShrinkRequest
	GetResourceOwnerId() *int64
	SetTemperature(v float64) *UpdateLargeModelShrinkRequest
	GetTemperature() *float64
	SetTopK(v int64) *UpdateLargeModelShrinkRequest
	GetTopK() *int64
	SetTopP(v float64) *UpdateLargeModelShrinkRequest
	GetTopP() *float64
}

type UpdateLargeModelShrinkRequest struct {
	// The authorization code.
	//
	// example:
	//
	// sk-sxxxxx*********xx
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// A list of base models.
	BaseModelShrink *string `json:"BaseModel,omitempty" xml:"BaseModel,omitempty"`
	// The model code.
	//
	// example:
	//
	// Test Model Name
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// The model name.
	//
	// example:
	//
	// Test model
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model URL.
	//
	// example:
	//
	// https://xxxxxxxxx
	ModelUrl             *string `json:"ModelUrl,omitempty" xml:"ModelUrl,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Controls the randomness of the model\\"s output. A higher value increases randomness, and a lower value makes the output more deterministic.
	//
	// example:
	//
	// 0.1
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// Restricts token selection to the top k most probable tokens.
	//
	// example:
	//
	// 2
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Controls output diversity by using nucleus sampling. It defines a cumulative probability threshold for token selection, considering only the most likely tokens.
	//
	// example:
	//
	// 0.1
	TopP *float64 `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s UpdateLargeModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLargeModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLargeModelShrinkRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *UpdateLargeModelShrinkRequest) GetBaseModelShrink() *string {
	return s.BaseModelShrink
}

func (s *UpdateLargeModelShrinkRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *UpdateLargeModelShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateLargeModelShrinkRequest) GetModelUrl() *string {
	return s.ModelUrl
}

func (s *UpdateLargeModelShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLargeModelShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateLargeModelShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateLargeModelShrinkRequest) GetTemperature() *float64 {
	return s.Temperature
}

func (s *UpdateLargeModelShrinkRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *UpdateLargeModelShrinkRequest) GetTopP() *float64 {
	return s.TopP
}

func (s *UpdateLargeModelShrinkRequest) SetAuthCode(v string) *UpdateLargeModelShrinkRequest {
	s.AuthCode = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetBaseModelShrink(v string) *UpdateLargeModelShrinkRequest {
	s.BaseModelShrink = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetModelCode(v string) *UpdateLargeModelShrinkRequest {
	s.ModelCode = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetModelName(v string) *UpdateLargeModelShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetModelUrl(v string) *UpdateLargeModelShrinkRequest {
	s.ModelUrl = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetOwnerId(v int64) *UpdateLargeModelShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetResourceOwnerAccount(v string) *UpdateLargeModelShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetResourceOwnerId(v int64) *UpdateLargeModelShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetTemperature(v float64) *UpdateLargeModelShrinkRequest {
	s.Temperature = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetTopK(v int64) *UpdateLargeModelShrinkRequest {
	s.TopK = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) SetTopP(v float64) *UpdateLargeModelShrinkRequest {
	s.TopP = &v
	return s
}

func (s *UpdateLargeModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
