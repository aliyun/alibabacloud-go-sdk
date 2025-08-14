// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllPublicMediaTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *ListAllPublicMediaTagsRequest
	GetBusinessType() *string
	SetEntityId(v string) *ListAllPublicMediaTagsRequest
	GetEntityId() *string
}

type ListAllPublicMediaTagsRequest struct {
	// The business type of the media asset.
	//
	// example:
	//
	// "sticker"
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The entity ID, which is used to distinguish between media assets of different types in the public domain.
	//
	// Set this parameter to Copyright_Music, which indicates music in the public domain.
	//
	// example:
	//
	// Copyright_Music
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s ListAllPublicMediaTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllPublicMediaTagsRequest) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListAllPublicMediaTagsRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *ListAllPublicMediaTagsRequest) SetBusinessType(v string) *ListAllPublicMediaTagsRequest {
	s.BusinessType = &v
	return s
}

func (s *ListAllPublicMediaTagsRequest) SetEntityId(v string) *ListAllPublicMediaTagsRequest {
	s.EntityId = &v
	return s
}

func (s *ListAllPublicMediaTagsRequest) Validate() error {
	return dara.Validate(s)
}
