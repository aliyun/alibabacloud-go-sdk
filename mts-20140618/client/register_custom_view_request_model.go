// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *RegisterCustomViewRequest
	GetAlgorithm() *string
	SetCustomEntityId(v string) *RegisterCustomViewRequest
	GetCustomEntityId() *string
	SetCustomGroupId(v string) *RegisterCustomViewRequest
	GetCustomGroupId() *string
	SetImageUrl(v string) *RegisterCustomViewRequest
	GetImageUrl() *string
	SetLabelPrompt(v string) *RegisterCustomViewRequest
	GetLabelPrompt() *string
	SetOwnerAccount(v string) *RegisterCustomViewRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RegisterCustomViewRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RegisterCustomViewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RegisterCustomViewRequest
	GetResourceOwnerId() *int64
}

type RegisterCustomViewRequest struct {
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
	// 2
	CustomEntityId *string `json:"CustomEntityId,omitempty" xml:"CustomEntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomGroupId *string `json:"CustomGroupId,omitempty" xml:"CustomGroupId,omitempty"`
	// example:
	//
	// http://``127.66.**.**``/image.jpeg
	ImageUrl             *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	LabelPrompt          *string `json:"LabelPrompt,omitempty" xml:"LabelPrompt,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RegisterCustomViewRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomViewRequest) GoString() string {
	return s.String()
}

func (s *RegisterCustomViewRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *RegisterCustomViewRequest) GetCustomEntityId() *string {
	return s.CustomEntityId
}

func (s *RegisterCustomViewRequest) GetCustomGroupId() *string {
	return s.CustomGroupId
}

func (s *RegisterCustomViewRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *RegisterCustomViewRequest) GetLabelPrompt() *string {
	return s.LabelPrompt
}

func (s *RegisterCustomViewRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RegisterCustomViewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RegisterCustomViewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RegisterCustomViewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RegisterCustomViewRequest) SetAlgorithm(v string) *RegisterCustomViewRequest {
	s.Algorithm = &v
	return s
}

func (s *RegisterCustomViewRequest) SetCustomEntityId(v string) *RegisterCustomViewRequest {
	s.CustomEntityId = &v
	return s
}

func (s *RegisterCustomViewRequest) SetCustomGroupId(v string) *RegisterCustomViewRequest {
	s.CustomGroupId = &v
	return s
}

func (s *RegisterCustomViewRequest) SetImageUrl(v string) *RegisterCustomViewRequest {
	s.ImageUrl = &v
	return s
}

func (s *RegisterCustomViewRequest) SetLabelPrompt(v string) *RegisterCustomViewRequest {
	s.LabelPrompt = &v
	return s
}

func (s *RegisterCustomViewRequest) SetOwnerAccount(v string) *RegisterCustomViewRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RegisterCustomViewRequest) SetOwnerId(v int64) *RegisterCustomViewRequest {
	s.OwnerId = &v
	return s
}

func (s *RegisterCustomViewRequest) SetResourceOwnerAccount(v string) *RegisterCustomViewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RegisterCustomViewRequest) SetResourceOwnerId(v int64) *RegisterCustomViewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RegisterCustomViewRequest) Validate() error {
	return dara.Validate(s)
}
