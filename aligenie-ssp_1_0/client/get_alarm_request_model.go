// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetAlarmRequestDeviceInfo) *GetAlarmRequest
	GetDeviceInfo() *GetAlarmRequestDeviceInfo
	SetPayload(v *GetAlarmRequestPayload) *GetAlarmRequest
	GetPayload() *GetAlarmRequestPayload
	SetUserInfo(v *GetAlarmRequestUserInfo) *GetAlarmRequest
	GetUserInfo() *GetAlarmRequestUserInfo
}

type GetAlarmRequest struct {
	// This parameter is required.
	DeviceInfo *GetAlarmRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *GetAlarmRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *GetAlarmRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmRequest) GetDeviceInfo() *GetAlarmRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetAlarmRequest) GetPayload() *GetAlarmRequestPayload {
	return s.Payload
}

func (s *GetAlarmRequest) GetUserInfo() *GetAlarmRequestUserInfo {
	return s.UserInfo
}

func (s *GetAlarmRequest) SetDeviceInfo(v *GetAlarmRequestDeviceInfo) *GetAlarmRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetAlarmRequest) SetPayload(v *GetAlarmRequestPayload) *GetAlarmRequest {
	s.Payload = v
	return s
}

func (s *GetAlarmRequest) SetUserInfo(v *GetAlarmRequestUserInfo) *GetAlarmRequest {
	s.UserInfo = v
	return s
}

func (s *GetAlarmRequest) Validate() error {
	return dara.Validate(s)
}

type GetAlarmRequestDeviceInfo struct {
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

func (s GetAlarmRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetAlarmRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetAlarmRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetAlarmRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetAlarmRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetAlarmRequestDeviceInfo) SetEncodeKey(v string) *GetAlarmRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetEncodeType(v string) *GetAlarmRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetId(v string) *GetAlarmRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetIdType(v string) *GetAlarmRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetOrganizationId(v string) *GetAlarmRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetAlarmRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
}

func (s GetAlarmRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmRequestPayload) GoString() string {
	return s.String()
}

func (s *GetAlarmRequestPayload) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *GetAlarmRequestPayload) SetAlarmId(v int64) *GetAlarmRequestPayload {
	s.AlarmId = &v
	return s
}

func (s *GetAlarmRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetAlarmRequestUserInfo struct {
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

func (s GetAlarmRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetAlarmRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetAlarmRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetAlarmRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetAlarmRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetAlarmRequestUserInfo) SetEncodeKey(v string) *GetAlarmRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetEncodeType(v string) *GetAlarmRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetId(v string) *GetAlarmRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetIdType(v string) *GetAlarmRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetOrganizationId(v string) *GetAlarmRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetAlarmRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
