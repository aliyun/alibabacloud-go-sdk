// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemindersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListRemindersRequestDeviceInfo) *ListRemindersRequest
	GetDeviceInfo() *ListRemindersRequestDeviceInfo
	SetPayload(v *ListRemindersRequestPayload) *ListRemindersRequest
	GetPayload() *ListRemindersRequestPayload
	SetUserInfo(v *ListRemindersRequestUserInfo) *ListRemindersRequest
	GetUserInfo() *ListRemindersRequestUserInfo
}

type ListRemindersRequest struct {
	// This parameter is required.
	DeviceInfo *ListRemindersRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *ListRemindersRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListRemindersRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListRemindersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersRequest) GoString() string {
	return s.String()
}

func (s *ListRemindersRequest) GetDeviceInfo() *ListRemindersRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListRemindersRequest) GetPayload() *ListRemindersRequestPayload {
	return s.Payload
}

func (s *ListRemindersRequest) GetUserInfo() *ListRemindersRequestUserInfo {
	return s.UserInfo
}

func (s *ListRemindersRequest) SetDeviceInfo(v *ListRemindersRequestDeviceInfo) *ListRemindersRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListRemindersRequest) SetPayload(v *ListRemindersRequestPayload) *ListRemindersRequest {
	s.Payload = v
	return s
}

func (s *ListRemindersRequest) SetUserInfo(v *ListRemindersRequestUserInfo) *ListRemindersRequest {
	s.UserInfo = v
	return s
}

func (s *ListRemindersRequest) Validate() error {
	return dara.Validate(s)
}

type ListRemindersRequestDeviceInfo struct {
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
	// PACKAGE_NAME
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

func (s ListRemindersRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListRemindersRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListRemindersRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListRemindersRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListRemindersRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListRemindersRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRemindersRequestDeviceInfo) SetEncodeKey(v string) *ListRemindersRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetEncodeType(v string) *ListRemindersRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetId(v string) *ListRemindersRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetIdType(v string) *ListRemindersRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetOrganizationId(v string) *ListRemindersRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListRemindersRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
}

func (s ListRemindersRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersRequestPayload) GoString() string {
	return s.String()
}

func (s *ListRemindersRequestPayload) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *ListRemindersRequestPayload) SetIsDebug(v bool) *ListRemindersRequestPayload {
	s.IsDebug = &v
	return s
}

func (s *ListRemindersRequestPayload) Validate() error {
	return dara.Validate(s)
}

type ListRemindersRequestUserInfo struct {
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
	// PACKAGE_NAME
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

func (s ListRemindersRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListRemindersRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListRemindersRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListRemindersRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListRemindersRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListRemindersRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRemindersRequestUserInfo) SetEncodeKey(v string) *ListRemindersRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetEncodeType(v string) *ListRemindersRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetId(v string) *ListRemindersRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetIdType(v string) *ListRemindersRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetOrganizationId(v string) *ListRemindersRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListRemindersRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
