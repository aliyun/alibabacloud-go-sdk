// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceRegistrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *FaceRegistrationRequest
	GetCategory() *string
	SetImageIds(v string) *FaceRegistrationRequest
	GetImageIds() *string
	SetOwnerAccount(v string) *FaceRegistrationRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *FaceRegistrationRequest
	GetOwnerId() *string
	SetPersonId(v string) *FaceRegistrationRequest
	GetPersonId() *string
	SetPersonLibrary(v string) *FaceRegistrationRequest
	GetPersonLibrary() *string
	SetPersonName(v string) *FaceRegistrationRequest
	GetPersonName() *string
	SetResourceOwnerAccount(v string) *FaceRegistrationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *FaceRegistrationRequest
	GetResourceOwnerId() *string
}

type FaceRegistrationRequest struct {
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ImageIds      *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PersonId      *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	PersonLibrary *string `json:"PersonLibrary,omitempty" xml:"PersonLibrary,omitempty"`
	// This parameter is required.
	PersonName           *string `json:"PersonName,omitempty" xml:"PersonName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s FaceRegistrationRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceRegistrationRequest) GoString() string {
	return s.String()
}

func (s *FaceRegistrationRequest) GetCategory() *string {
	return s.Category
}

func (s *FaceRegistrationRequest) GetImageIds() *string {
	return s.ImageIds
}

func (s *FaceRegistrationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *FaceRegistrationRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *FaceRegistrationRequest) GetPersonId() *string {
	return s.PersonId
}

func (s *FaceRegistrationRequest) GetPersonLibrary() *string {
	return s.PersonLibrary
}

func (s *FaceRegistrationRequest) GetPersonName() *string {
	return s.PersonName
}

func (s *FaceRegistrationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FaceRegistrationRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *FaceRegistrationRequest) SetCategory(v string) *FaceRegistrationRequest {
	s.Category = &v
	return s
}

func (s *FaceRegistrationRequest) SetImageIds(v string) *FaceRegistrationRequest {
	s.ImageIds = &v
	return s
}

func (s *FaceRegistrationRequest) SetOwnerAccount(v string) *FaceRegistrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FaceRegistrationRequest) SetOwnerId(v string) *FaceRegistrationRequest {
	s.OwnerId = &v
	return s
}

func (s *FaceRegistrationRequest) SetPersonId(v string) *FaceRegistrationRequest {
	s.PersonId = &v
	return s
}

func (s *FaceRegistrationRequest) SetPersonLibrary(v string) *FaceRegistrationRequest {
	s.PersonLibrary = &v
	return s
}

func (s *FaceRegistrationRequest) SetPersonName(v string) *FaceRegistrationRequest {
	s.PersonName = &v
	return s
}

func (s *FaceRegistrationRequest) SetResourceOwnerAccount(v string) *FaceRegistrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FaceRegistrationRequest) SetResourceOwnerId(v string) *FaceRegistrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FaceRegistrationRequest) Validate() error {
	return dara.Validate(s)
}
