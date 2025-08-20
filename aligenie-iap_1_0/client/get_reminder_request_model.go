// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReminderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetReminderRequestDeviceInfo) *GetReminderRequest
	GetDeviceInfo() *GetReminderRequestDeviceInfo
	SetPayload(v *GetReminderRequestPayload) *GetReminderRequest
	GetPayload() *GetReminderRequestPayload
	SetUserInfo(v *GetReminderRequestUserInfo) *GetReminderRequest
	GetUserInfo() *GetReminderRequestUserInfo
}

type GetReminderRequest struct {
	// This parameter is required.
	DeviceInfo *GetReminderRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *GetReminderRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *GetReminderRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetReminderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReminderRequest) GoString() string {
	return s.String()
}

func (s *GetReminderRequest) GetDeviceInfo() *GetReminderRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetReminderRequest) GetPayload() *GetReminderRequestPayload {
	return s.Payload
}

func (s *GetReminderRequest) GetUserInfo() *GetReminderRequestUserInfo {
	return s.UserInfo
}

func (s *GetReminderRequest) SetDeviceInfo(v *GetReminderRequestDeviceInfo) *GetReminderRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetReminderRequest) SetPayload(v *GetReminderRequestPayload) *GetReminderRequest {
	s.Payload = v
	return s
}

func (s *GetReminderRequest) SetUserInfo(v *GetReminderRequestUserInfo) *GetReminderRequest {
	s.UserInfo = v
	return s
}

func (s *GetReminderRequest) Validate() error {
	return dara.Validate(s)
}

type GetReminderRequestDeviceInfo struct {
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

func (s GetReminderRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetReminderRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetReminderRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetReminderRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetReminderRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetReminderRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetReminderRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetReminderRequestDeviceInfo) SetEncodeKey(v string) *GetReminderRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetReminderRequestDeviceInfo) SetEncodeType(v string) *GetReminderRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetReminderRequestDeviceInfo) SetId(v string) *GetReminderRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetReminderRequestDeviceInfo) SetIdType(v string) *GetReminderRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetReminderRequestDeviceInfo) SetOrganizationId(v string) *GetReminderRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetReminderRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetReminderRequestPayload struct {
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

func (s GetReminderRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetReminderRequestPayload) GoString() string {
	return s.String()
}

func (s *GetReminderRequestPayload) GetId() *int64 {
	return s.Id
}

func (s *GetReminderRequestPayload) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *GetReminderRequestPayload) SetId(v int64) *GetReminderRequestPayload {
	s.Id = &v
	return s
}

func (s *GetReminderRequestPayload) SetIsDebug(v bool) *GetReminderRequestPayload {
	s.IsDebug = &v
	return s
}

func (s *GetReminderRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetReminderRequestUserInfo struct {
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

func (s GetReminderRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetReminderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetReminderRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetReminderRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetReminderRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetReminderRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetReminderRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetReminderRequestUserInfo) SetEncodeKey(v string) *GetReminderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetReminderRequestUserInfo) SetEncodeType(v string) *GetReminderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetReminderRequestUserInfo) SetId(v string) *GetReminderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetReminderRequestUserInfo) SetIdType(v string) *GetReminderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetReminderRequestUserInfo) SetOrganizationId(v string) *GetReminderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetReminderRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
