// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecognitionSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateRecognitionSampleRequest
	GetAlgorithm() *string
	SetEntityId(v string) *CreateRecognitionSampleRequest
	GetEntityId() *string
	SetImageUrl(v string) *CreateRecognitionSampleRequest
	GetImageUrl() *string
	SetLabelPrompt(v string) *CreateRecognitionSampleRequest
	GetLabelPrompt() *string
	SetLibId(v string) *CreateRecognitionSampleRequest
	GetLibId() *string
	SetOwnerAccount(v string) *CreateRecognitionSampleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRecognitionSampleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateRecognitionSampleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRecognitionSampleRequest
	GetResourceOwnerId() *int64
}

type CreateRecognitionSampleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// **************544cb84754************
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// https://example.com/sample.png
	ImageUrl    *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	LabelPrompt *string `json:"LabelPrompt,omitempty" xml:"LabelPrompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *************24b47865c6**************
	LibId                *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateRecognitionSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecognitionSampleRequest) GoString() string {
	return s.String()
}

func (s *CreateRecognitionSampleRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateRecognitionSampleRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateRecognitionSampleRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateRecognitionSampleRequest) GetLabelPrompt() *string {
	return s.LabelPrompt
}

func (s *CreateRecognitionSampleRequest) GetLibId() *string {
	return s.LibId
}

func (s *CreateRecognitionSampleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRecognitionSampleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRecognitionSampleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRecognitionSampleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRecognitionSampleRequest) SetAlgorithm(v string) *CreateRecognitionSampleRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateRecognitionSampleRequest) SetEntityId(v string) *CreateRecognitionSampleRequest {
	s.EntityId = &v
	return s
}

func (s *CreateRecognitionSampleRequest) SetImageUrl(v string) *CreateRecognitionSampleRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateRecognitionSampleRequest) SetLabelPrompt(v string) *CreateRecognitionSampleRequest {
	s.LabelPrompt = &v
	return s
}

func (s *CreateRecognitionSampleRequest) SetLibId(v string) *CreateRecognitionSampleRequest {
	s.LibId = &v
	return s
}

func (s *CreateRecognitionSampleRequest) SetOwnerAccount(v string) *CreateRecognitionSampleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRecognitionSampleRequest) SetOwnerId(v int64) *CreateRecognitionSampleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRecognitionSampleRequest) SetResourceOwnerAccount(v string) *CreateRecognitionSampleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRecognitionSampleRequest) SetResourceOwnerId(v int64) *CreateRecognitionSampleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRecognitionSampleRequest) Validate() error {
	return dara.Validate(s)
}
