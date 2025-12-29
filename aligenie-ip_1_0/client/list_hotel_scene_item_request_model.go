// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayload(v *ListHotelSceneItemRequestPayload) *ListHotelSceneItemRequest
	GetPayload() *ListHotelSceneItemRequestPayload
	SetUserInfo(v *ListHotelSceneItemRequestUserInfo) *ListHotelSceneItemRequest
	GetUserInfo() *ListHotelSceneItemRequestUserInfo
}

type ListHotelSceneItemRequest struct {
	// This parameter is required.
	Payload *ListHotelSceneItemRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListHotelSceneItemRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemRequest) GetPayload() *ListHotelSceneItemRequestPayload {
	return s.Payload
}

func (s *ListHotelSceneItemRequest) GetUserInfo() *ListHotelSceneItemRequestUserInfo {
	return s.UserInfo
}

func (s *ListHotelSceneItemRequest) SetPayload(v *ListHotelSceneItemRequestPayload) *ListHotelSceneItemRequest {
	s.Payload = v
	return s
}

func (s *ListHotelSceneItemRequest) SetUserInfo(v *ListHotelSceneItemRequestUserInfo) *ListHotelSceneItemRequest {
	s.UserInfo = v
	return s
}

func (s *ListHotelSceneItemRequest) Validate() error {
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelSceneItemRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneItemRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemRequestPayload) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemRequestPayload) GetType() *string {
	return s.Type
}

func (s *ListHotelSceneItemRequestPayload) SetType(v string) *ListHotelSceneItemRequestPayload {
	s.Type = &v
	return s
}

func (s *ListHotelSceneItemRequestPayload) Validate() error {
	return dara.Validate(s)
}

type ListHotelSceneItemRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 1248494721591392955
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mFU6VtVU+pgA8lx6rYMo7SPl11t+8b+8ALrn10MIPEdpK/HI9wELAEppYhPI1cYRDa4og8AMjAEBZKbLUwFjFA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListHotelSceneItemRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListHotelSceneItemRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListHotelSceneItemRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListHotelSceneItemRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListHotelSceneItemRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListHotelSceneItemRequestUserInfo) SetEncodeKey(v string) *ListHotelSceneItemRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetEncodeType(v string) *ListHotelSceneItemRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetId(v string) *ListHotelSceneItemRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetIdType(v string) *ListHotelSceneItemRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) SetOrganizationId(v string) *ListHotelSceneItemRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListHotelSceneItemRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
