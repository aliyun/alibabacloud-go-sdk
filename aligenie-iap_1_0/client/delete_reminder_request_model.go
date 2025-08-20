// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReminderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *DeleteReminderRequestDeviceInfo) *DeleteReminderRequest
	GetDeviceInfo() *DeleteReminderRequestDeviceInfo
	SetPayload(v *DeleteReminderRequestPayload) *DeleteReminderRequest
	GetPayload() *DeleteReminderRequestPayload
	SetUserInfo(v *DeleteReminderRequestUserInfo) *DeleteReminderRequest
	GetUserInfo() *DeleteReminderRequestUserInfo
}

type DeleteReminderRequest struct {
	// This parameter is required.
	DeviceInfo *DeleteReminderRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *DeleteReminderRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *DeleteReminderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeleteReminderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReminderRequest) GoString() string {
	return s.String()
}

func (s *DeleteReminderRequest) GetDeviceInfo() *DeleteReminderRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *DeleteReminderRequest) GetPayload() *DeleteReminderRequestPayload {
	return s.Payload
}

func (s *DeleteReminderRequest) GetUserInfo() *DeleteReminderRequestUserInfo {
	return s.UserInfo
}

func (s *DeleteReminderRequest) SetDeviceInfo(v *DeleteReminderRequestDeviceInfo) *DeleteReminderRequest {
	s.DeviceInfo = v
	return s
}

func (s *DeleteReminderRequest) SetPayload(v *DeleteReminderRequestPayload) *DeleteReminderRequest {
	s.Payload = v
	return s
}

func (s *DeleteReminderRequest) SetUserInfo(v *DeleteReminderRequestUserInfo) *DeleteReminderRequest {
	s.UserInfo = v
	return s
}

func (s *DeleteReminderRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteReminderRequestDeviceInfo struct {
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
	// SKILL_ID
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

func (s DeleteReminderRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteReminderRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *DeleteReminderRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeleteReminderRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeleteReminderRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *DeleteReminderRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeleteReminderRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteReminderRequestDeviceInfo) SetEncodeKey(v string) *DeleteReminderRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteReminderRequestDeviceInfo) SetEncodeType(v string) *DeleteReminderRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteReminderRequestDeviceInfo) SetId(v string) *DeleteReminderRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *DeleteReminderRequestDeviceInfo) SetIdType(v string) *DeleteReminderRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *DeleteReminderRequestDeviceInfo) SetOrganizationId(v string) *DeleteReminderRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeleteReminderRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type DeleteReminderRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// 20****34
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
}

func (s DeleteReminderRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s DeleteReminderRequestPayload) GoString() string {
	return s.String()
}

func (s *DeleteReminderRequestPayload) GetId() *int64 {
	return s.Id
}

func (s *DeleteReminderRequestPayload) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *DeleteReminderRequestPayload) SetId(v int64) *DeleteReminderRequestPayload {
	s.Id = &v
	return s
}

func (s *DeleteReminderRequestPayload) SetIsDebug(v bool) *DeleteReminderRequestPayload {
	s.IsDebug = &v
	return s
}

func (s *DeleteReminderRequestPayload) Validate() error {
	return dara.Validate(s)
}

type DeleteReminderRequestUserInfo struct {
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
	// SKILL_ID
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

func (s DeleteReminderRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteReminderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeleteReminderRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeleteReminderRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeleteReminderRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *DeleteReminderRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeleteReminderRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteReminderRequestUserInfo) SetEncodeKey(v string) *DeleteReminderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteReminderRequestUserInfo) SetEncodeType(v string) *DeleteReminderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteReminderRequestUserInfo) SetId(v string) *DeleteReminderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeleteReminderRequestUserInfo) SetIdType(v string) *DeleteReminderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeleteReminderRequestUserInfo) SetOrganizationId(v string) *DeleteReminderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeleteReminderRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
