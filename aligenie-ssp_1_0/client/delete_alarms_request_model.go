// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlarmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *DeleteAlarmsRequestDeviceInfo) *DeleteAlarmsRequest
	GetDeviceInfo() *DeleteAlarmsRequestDeviceInfo
	SetPayload(v *DeleteAlarmsRequestPayload) *DeleteAlarmsRequest
	GetPayload() *DeleteAlarmsRequestPayload
	SetUserInfo(v *DeleteAlarmsRequestUserInfo) *DeleteAlarmsRequest
	GetUserInfo() *DeleteAlarmsRequestUserInfo
}

type DeleteAlarmsRequest struct {
	// This parameter is required.
	DeviceInfo *DeleteAlarmsRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *DeleteAlarmsRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *DeleteAlarmsRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeleteAlarmsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsRequest) GetDeviceInfo() *DeleteAlarmsRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *DeleteAlarmsRequest) GetPayload() *DeleteAlarmsRequestPayload {
	return s.Payload
}

func (s *DeleteAlarmsRequest) GetUserInfo() *DeleteAlarmsRequestUserInfo {
	return s.UserInfo
}

func (s *DeleteAlarmsRequest) SetDeviceInfo(v *DeleteAlarmsRequestDeviceInfo) *DeleteAlarmsRequest {
	s.DeviceInfo = v
	return s
}

func (s *DeleteAlarmsRequest) SetPayload(v *DeleteAlarmsRequestPayload) *DeleteAlarmsRequest {
	s.Payload = v
	return s
}

func (s *DeleteAlarmsRequest) SetUserInfo(v *DeleteAlarmsRequestUserInfo) *DeleteAlarmsRequest {
	s.UserInfo = v
	return s
}

func (s *DeleteAlarmsRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteAlarmsRequestDeviceInfo struct {
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

func (s DeleteAlarmsRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmsRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeleteAlarmsRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeleteAlarmsRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *DeleteAlarmsRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeleteAlarmsRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteAlarmsRequestDeviceInfo) SetEncodeKey(v string) *DeleteAlarmsRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteAlarmsRequestDeviceInfo) SetEncodeType(v string) *DeleteAlarmsRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteAlarmsRequestDeviceInfo) SetId(v string) *DeleteAlarmsRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *DeleteAlarmsRequestDeviceInfo) SetIdType(v string) *DeleteAlarmsRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *DeleteAlarmsRequestDeviceInfo) SetOrganizationId(v string) *DeleteAlarmsRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeleteAlarmsRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type DeleteAlarmsRequestPayload struct {
	// This parameter is required.
	AlarmIds []*int64 `json:"AlarmIds,omitempty" xml:"AlarmIds,omitempty" type:"Repeated"`
}

func (s DeleteAlarmsRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmsRequestPayload) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsRequestPayload) GetAlarmIds() []*int64 {
	return s.AlarmIds
}

func (s *DeleteAlarmsRequestPayload) SetAlarmIds(v []*int64) *DeleteAlarmsRequestPayload {
	s.AlarmIds = v
	return s
}

func (s *DeleteAlarmsRequestPayload) Validate() error {
	return dara.Validate(s)
}

type DeleteAlarmsRequestUserInfo struct {
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

func (s DeleteAlarmsRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmsRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeleteAlarmsRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeleteAlarmsRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *DeleteAlarmsRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeleteAlarmsRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteAlarmsRequestUserInfo) SetEncodeKey(v string) *DeleteAlarmsRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteAlarmsRequestUserInfo) SetEncodeType(v string) *DeleteAlarmsRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteAlarmsRequestUserInfo) SetId(v string) *DeleteAlarmsRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeleteAlarmsRequestUserInfo) SetIdType(v string) *DeleteAlarmsRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeleteAlarmsRequestUserInfo) SetOrganizationId(v string) *DeleteAlarmsRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeleteAlarmsRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
