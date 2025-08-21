// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *UnbindDeviceRequestDeviceInfo) *UnbindDeviceRequest
	GetDeviceInfo() *UnbindDeviceRequestDeviceInfo
	SetUserInfo(v *UnbindDeviceRequestUserInfo) *UnbindDeviceRequest
	GetUserInfo() *UnbindDeviceRequestUserInfo
}

type UnbindDeviceRequest struct {
	// This parameter is required.
	DeviceInfo *UnbindDeviceRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *UnbindDeviceRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s UnbindDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceRequest) GetDeviceInfo() *UnbindDeviceRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *UnbindDeviceRequest) GetUserInfo() *UnbindDeviceRequestUserInfo {
	return s.UserInfo
}

func (s *UnbindDeviceRequest) SetDeviceInfo(v *UnbindDeviceRequestDeviceInfo) *UnbindDeviceRequest {
	s.DeviceInfo = v
	return s
}

func (s *UnbindDeviceRequest) SetUserInfo(v *UnbindDeviceRequestUserInfo) *UnbindDeviceRequest {
	s.UserInfo = v
	return s
}

func (s *UnbindDeviceRequest) Validate() error {
	return dara.Validate(s)
}

type UnbindDeviceRequestDeviceInfo struct {
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
	// DAFE****ce3ej=
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

func (s UnbindDeviceRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *UnbindDeviceRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *UnbindDeviceRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *UnbindDeviceRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *UnbindDeviceRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *UnbindDeviceRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UnbindDeviceRequestDeviceInfo) SetEncodeKey(v string) *UnbindDeviceRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *UnbindDeviceRequestDeviceInfo) SetEncodeType(v string) *UnbindDeviceRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *UnbindDeviceRequestDeviceInfo) SetId(v string) *UnbindDeviceRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *UnbindDeviceRequestDeviceInfo) SetIdType(v string) *UnbindDeviceRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *UnbindDeviceRequestDeviceInfo) SetOrganizationId(v string) *UnbindDeviceRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *UnbindDeviceRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type UnbindDeviceRequestUserInfo struct {
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

func (s UnbindDeviceRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceRequestUserInfo) GoString() string {
	return s.String()
}

func (s *UnbindDeviceRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *UnbindDeviceRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *UnbindDeviceRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *UnbindDeviceRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *UnbindDeviceRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UnbindDeviceRequestUserInfo) SetEncodeKey(v string) *UnbindDeviceRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *UnbindDeviceRequestUserInfo) SetEncodeType(v string) *UnbindDeviceRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *UnbindDeviceRequestUserInfo) SetId(v string) *UnbindDeviceRequestUserInfo {
	s.Id = &v
	return s
}

func (s *UnbindDeviceRequestUserInfo) SetIdType(v string) *UnbindDeviceRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *UnbindDeviceRequestUserInfo) SetOrganizationId(v string) *UnbindDeviceRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *UnbindDeviceRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
