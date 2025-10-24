// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagCustomPersonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryDescription(v string) *TagCustomPersonRequest
	GetCategoryDescription() *string
	SetCategoryId(v string) *TagCustomPersonRequest
	GetCategoryId() *string
	SetCategoryName(v string) *TagCustomPersonRequest
	GetCategoryName() *string
	SetOwnerAccount(v string) *TagCustomPersonRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TagCustomPersonRequest
	GetOwnerId() *int64
	SetPersonDescription(v string) *TagCustomPersonRequest
	GetPersonDescription() *string
	SetPersonId(v string) *TagCustomPersonRequest
	GetPersonId() *string
	SetPersonName(v string) *TagCustomPersonRequest
	GetPersonName() *string
	SetResourceOwnerAccount(v string) *TagCustomPersonRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TagCustomPersonRequest
	GetResourceOwnerId() *int64
}

type TagCustomPersonRequest struct {
	// example:
	//
	// CategoryDescription001-****
	CategoryDescription *string `json:"CategoryDescription,omitempty" xml:"CategoryDescription,omitempty"`
	// example:
	//
	// CategoryId001-****
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// CategoryNametest-****
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// PersonDescriptiontest-****
	PersonDescription *string `json:"PersonDescription,omitempty" xml:"PersonDescription,omitempty"`
	// example:
	//
	// PersonId001-****
	PersonId *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	// example:
	//
	// PersonNametest-****
	PersonName           *string `json:"PersonName,omitempty" xml:"PersonName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s TagCustomPersonRequest) String() string {
	return dara.Prettify(s)
}

func (s TagCustomPersonRequest) GoString() string {
	return s.String()
}

func (s *TagCustomPersonRequest) GetCategoryDescription() *string {
	return s.CategoryDescription
}

func (s *TagCustomPersonRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *TagCustomPersonRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *TagCustomPersonRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TagCustomPersonRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TagCustomPersonRequest) GetPersonDescription() *string {
	return s.PersonDescription
}

func (s *TagCustomPersonRequest) GetPersonId() *string {
	return s.PersonId
}

func (s *TagCustomPersonRequest) GetPersonName() *string {
	return s.PersonName
}

func (s *TagCustomPersonRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TagCustomPersonRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TagCustomPersonRequest) SetCategoryDescription(v string) *TagCustomPersonRequest {
	s.CategoryDescription = &v
	return s
}

func (s *TagCustomPersonRequest) SetCategoryId(v string) *TagCustomPersonRequest {
	s.CategoryId = &v
	return s
}

func (s *TagCustomPersonRequest) SetCategoryName(v string) *TagCustomPersonRequest {
	s.CategoryName = &v
	return s
}

func (s *TagCustomPersonRequest) SetOwnerAccount(v string) *TagCustomPersonRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagCustomPersonRequest) SetOwnerId(v int64) *TagCustomPersonRequest {
	s.OwnerId = &v
	return s
}

func (s *TagCustomPersonRequest) SetPersonDescription(v string) *TagCustomPersonRequest {
	s.PersonDescription = &v
	return s
}

func (s *TagCustomPersonRequest) SetPersonId(v string) *TagCustomPersonRequest {
	s.PersonId = &v
	return s
}

func (s *TagCustomPersonRequest) SetPersonName(v string) *TagCustomPersonRequest {
	s.PersonName = &v
	return s
}

func (s *TagCustomPersonRequest) SetResourceOwnerAccount(v string) *TagCustomPersonRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagCustomPersonRequest) SetResourceOwnerId(v int64) *TagCustomPersonRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagCustomPersonRequest) Validate() error {
	return dara.Validate(s)
}
