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
	// The authorization code.
	//
	// example:
	//
	// sk-sxxxxx*********xx
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// A list of base models.
	BaseModel []*string `json:"BaseModel,omitempty" xml:"BaseModel,omitempty" type:"Repeated"`
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
