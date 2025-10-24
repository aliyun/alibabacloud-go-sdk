// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *RegisterCustomFaceRequest
	GetCategoryId() *string
	SetImageUrl(v string) *RegisterCustomFaceRequest
	GetImageUrl() *string
	SetOwnerAccount(v string) *RegisterCustomFaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RegisterCustomFaceRequest
	GetOwnerId() *int64
	SetPersonId(v string) *RegisterCustomFaceRequest
	GetPersonId() *string
	SetPersonName(v string) *RegisterCustomFaceRequest
	GetPersonName() *string
	SetResourceOwnerAccount(v string) *RegisterCustomFaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RegisterCustomFaceRequest
	GetResourceOwnerId() *int64
}

type RegisterCustomFaceRequest struct {
	// The ID of the figure library in which you want to register a custom face. The ID is used to uniquely identify a figure library. You can specify the ID of a custom figure library. Make sure that the ID is unique and keep the ID for future API operation calls. If you set this parameter to the ID of a system figure library, the custom face is registered in the system figure library. The ID can be up to 120 characters in length and is not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// CategoryId001-****
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The URL of the facial image that you want to register for the specified figure. The image must contain only one face.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example-****.jpeg
	ImageUrl     *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the figure for which you want to register a custom face. The ID is used to uniquely identify a figure. You can specify a figure ID. Make sure that the ID is unique and keep the ID for future API operation calls. The ID can be up to 120 characters in length and is not case-sensitive. The value returned is of the String type.
	//
	// This parameter is required.
	//
	// example:
	//
	// PersonId001-****
	PersonId             *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	PersonName           *string `json:"PersonName,omitempty" xml:"PersonName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RegisterCustomFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomFaceRequest) GoString() string {
	return s.String()
}

func (s *RegisterCustomFaceRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *RegisterCustomFaceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *RegisterCustomFaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RegisterCustomFaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RegisterCustomFaceRequest) GetPersonId() *string {
	return s.PersonId
}

func (s *RegisterCustomFaceRequest) GetPersonName() *string {
	return s.PersonName
}

func (s *RegisterCustomFaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RegisterCustomFaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RegisterCustomFaceRequest) SetCategoryId(v string) *RegisterCustomFaceRequest {
	s.CategoryId = &v
	return s
}

func (s *RegisterCustomFaceRequest) SetImageUrl(v string) *RegisterCustomFaceRequest {
	s.ImageUrl = &v
	return s
}

func (s *RegisterCustomFaceRequest) SetOwnerAccount(v string) *RegisterCustomFaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RegisterCustomFaceRequest) SetOwnerId(v int64) *RegisterCustomFaceRequest {
	s.OwnerId = &v
	return s
}

func (s *RegisterCustomFaceRequest) SetPersonId(v string) *RegisterCustomFaceRequest {
	s.PersonId = &v
	return s
}

func (s *RegisterCustomFaceRequest) SetPersonName(v string) *RegisterCustomFaceRequest {
	s.PersonName = &v
	return s
}

func (s *RegisterCustomFaceRequest) SetResourceOwnerAccount(v string) *RegisterCustomFaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RegisterCustomFaceRequest) SetResourceOwnerId(v int64) *RegisterCustomFaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RegisterCustomFaceRequest) Validate() error {
	return dara.Validate(s)
}
