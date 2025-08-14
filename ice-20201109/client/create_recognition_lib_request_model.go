// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecognitionLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateRecognitionLibRequest
	GetAlgorithm() *string
	SetLibDescription(v string) *CreateRecognitionLibRequest
	GetLibDescription() *string
	SetLibName(v string) *CreateRecognitionLibRequest
	GetLibName() *string
	SetOwnerAccount(v string) *CreateRecognitionLibRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRecognitionLibRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateRecognitionLibRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRecognitionLibRequest
	GetResourceOwnerId() *int64
}

type CreateRecognitionLibRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm      *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	LibDescription *string `json:"LibDescription,omitempty" xml:"LibDescription,omitempty"`
	// This parameter is required.
	LibName              *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateRecognitionLibRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecognitionLibRequest) GoString() string {
	return s.String()
}

func (s *CreateRecognitionLibRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateRecognitionLibRequest) GetLibDescription() *string {
	return s.LibDescription
}

func (s *CreateRecognitionLibRequest) GetLibName() *string {
	return s.LibName
}

func (s *CreateRecognitionLibRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRecognitionLibRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRecognitionLibRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRecognitionLibRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRecognitionLibRequest) SetAlgorithm(v string) *CreateRecognitionLibRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateRecognitionLibRequest) SetLibDescription(v string) *CreateRecognitionLibRequest {
	s.LibDescription = &v
	return s
}

func (s *CreateRecognitionLibRequest) SetLibName(v string) *CreateRecognitionLibRequest {
	s.LibName = &v
	return s
}

func (s *CreateRecognitionLibRequest) SetOwnerAccount(v string) *CreateRecognitionLibRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRecognitionLibRequest) SetOwnerId(v int64) *CreateRecognitionLibRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRecognitionLibRequest) SetResourceOwnerAccount(v string) *CreateRecognitionLibRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRecognitionLibRequest) SetResourceOwnerId(v int64) *CreateRecognitionLibRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRecognitionLibRequest) Validate() error {
	return dara.Validate(s)
}
