// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnregisterCustomFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *UnregisterCustomFaceRequest
	GetCategoryId() *string
	SetFaceId(v string) *UnregisterCustomFaceRequest
	GetFaceId() *string
	SetOwnerAccount(v string) *UnregisterCustomFaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnregisterCustomFaceRequest
	GetOwnerId() *int64
	SetPersonId(v string) *UnregisterCustomFaceRequest
	GetPersonId() *string
	SetResourceOwnerAccount(v string) *UnregisterCustomFaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnregisterCustomFaceRequest
	GetResourceOwnerId() *int64
}

type UnregisterCustomFaceRequest struct {
	// The ID of the figure library. The ID is used to uniquely identify a figure library. You can specify the ID of a custom figure library. Make sure that the ID is unique. If you set this parameter to the ID of a system figure library, the system figure library is used. The ID can be up to 120 characters in length and is not case-sensitive. You can call the [ListCustomPersons](https://help.aliyun.com/document_detail/187787.html) operation to query the figure library ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// CategoryId001-****
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The ID of the face. The ID is used to uniquely identify a face. Make sure that the ID is unique. The ID can be up to 120 characters in length and is not case-sensitive. You can call the [ListCustomPersons](https://help.aliyun.com/document_detail/187787.html) operation to query the face ID. If you set this parameter to ALL, all the faces associated with the specified figure are deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15****
	FaceId       *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the figure. The ID is used to uniquely identify a custom figure. Make sure that the ID is unique. The ID can be up to 120 characters in length and is not case-sensitive. You can call the [ListCustomPersons](https://help.aliyun.com/document_detail/187787.html) operation to query the figure ID. If you set this parameter to ALL, all the faces in the specified figure library are deleted, and the custom figure library is deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// PersonId001-****
	PersonId             *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnregisterCustomFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnregisterCustomFaceRequest) GoString() string {
	return s.String()
}

func (s *UnregisterCustomFaceRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *UnregisterCustomFaceRequest) GetFaceId() *string {
	return s.FaceId
}

func (s *UnregisterCustomFaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnregisterCustomFaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnregisterCustomFaceRequest) GetPersonId() *string {
	return s.PersonId
}

func (s *UnregisterCustomFaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnregisterCustomFaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnregisterCustomFaceRequest) SetCategoryId(v string) *UnregisterCustomFaceRequest {
	s.CategoryId = &v
	return s
}

func (s *UnregisterCustomFaceRequest) SetFaceId(v string) *UnregisterCustomFaceRequest {
	s.FaceId = &v
	return s
}

func (s *UnregisterCustomFaceRequest) SetOwnerAccount(v string) *UnregisterCustomFaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnregisterCustomFaceRequest) SetOwnerId(v int64) *UnregisterCustomFaceRequest {
	s.OwnerId = &v
	return s
}

func (s *UnregisterCustomFaceRequest) SetPersonId(v string) *UnregisterCustomFaceRequest {
	s.PersonId = &v
	return s
}

func (s *UnregisterCustomFaceRequest) SetResourceOwnerAccount(v string) *UnregisterCustomFaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnregisterCustomFaceRequest) SetResourceOwnerId(v int64) *UnregisterCustomFaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnregisterCustomFaceRequest) Validate() error {
	return dara.Validate(s)
}
