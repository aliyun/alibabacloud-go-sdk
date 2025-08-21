// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfo(v *ListDeviceByUserIdRequestUserInfo) *ListDeviceByUserIdRequest
	GetUserInfo() *ListDeviceByUserIdRequestUserInfo
}

type ListDeviceByUserIdRequest struct {
	// This parameter is required.
	UserInfo *ListDeviceByUserIdRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListDeviceByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdRequest) GetUserInfo() *ListDeviceByUserIdRequestUserInfo {
	return s.UserInfo
}

func (s *ListDeviceByUserIdRequest) SetUserInfo(v *ListDeviceByUserIdRequestUserInfo) *ListDeviceByUserIdRequest {
	s.UserInfo = v
	return s
}

func (s *ListDeviceByUserIdRequest) Validate() error {
	return dara.Validate(s)
}

type ListDeviceByUserIdRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
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
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListDeviceByUserIdRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListDeviceByUserIdRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListDeviceByUserIdRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListDeviceByUserIdRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListDeviceByUserIdRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListDeviceByUserIdRequestUserInfo) SetEncodeKey(v string) *ListDeviceByUserIdRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListDeviceByUserIdRequestUserInfo) SetEncodeType(v string) *ListDeviceByUserIdRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListDeviceByUserIdRequestUserInfo) SetId(v string) *ListDeviceByUserIdRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListDeviceByUserIdRequestUserInfo) SetIdType(v string) *ListDeviceByUserIdRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListDeviceByUserIdRequestUserInfo) SetOrganizationId(v string) *ListDeviceByUserIdRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListDeviceByUserIdRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
