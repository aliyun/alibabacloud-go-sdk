// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *DeleteScheduleTaskRequestDeviceInfo) *DeleteScheduleTaskRequest
	GetDeviceInfo() *DeleteScheduleTaskRequestDeviceInfo
	SetPayload(v *DeleteScheduleTaskRequestPayload) *DeleteScheduleTaskRequest
	GetPayload() *DeleteScheduleTaskRequestPayload
	SetUserInfo(v *DeleteScheduleTaskRequestUserInfo) *DeleteScheduleTaskRequest
	GetUserInfo() *DeleteScheduleTaskRequestUserInfo
}

type DeleteScheduleTaskRequest struct {
	// This parameter is required.
	DeviceInfo *DeleteScheduleTaskRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *DeleteScheduleTaskRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *DeleteScheduleTaskRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeleteScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequest) GetDeviceInfo() *DeleteScheduleTaskRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *DeleteScheduleTaskRequest) GetPayload() *DeleteScheduleTaskRequestPayload {
	return s.Payload
}

func (s *DeleteScheduleTaskRequest) GetUserInfo() *DeleteScheduleTaskRequestUserInfo {
	return s.UserInfo
}

func (s *DeleteScheduleTaskRequest) SetDeviceInfo(v *DeleteScheduleTaskRequestDeviceInfo) *DeleteScheduleTaskRequest {
	s.DeviceInfo = v
	return s
}

func (s *DeleteScheduleTaskRequest) SetPayload(v *DeleteScheduleTaskRequestPayload) *DeleteScheduleTaskRequest {
	s.Payload = v
	return s
}

func (s *DeleteScheduleTaskRequest) SetUserInfo(v *DeleteScheduleTaskRequestUserInfo) *DeleteScheduleTaskRequest {
	s.UserInfo = v
	return s
}

func (s *DeleteScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteScheduleTaskRequestDeviceInfo struct {
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteScheduleTaskRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetEncodeKey(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetEncodeType(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetId(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetIdType(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetOrganizationId(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type DeleteScheduleTaskRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteScheduleTaskRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskRequestPayload) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequestPayload) GetId() *int64 {
	return s.Id
}

func (s *DeleteScheduleTaskRequestPayload) SetId(v int64) *DeleteScheduleTaskRequestPayload {
	s.Id = &v
	return s
}

func (s *DeleteScheduleTaskRequestPayload) Validate() error {
	return dara.Validate(s)
}

type DeleteScheduleTaskRequestUserInfo struct {
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteScheduleTaskRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeleteScheduleTaskRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeleteScheduleTaskRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *DeleteScheduleTaskRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeleteScheduleTaskRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteScheduleTaskRequestUserInfo) SetEncodeKey(v string) *DeleteScheduleTaskRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetEncodeType(v string) *DeleteScheduleTaskRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetId(v string) *DeleteScheduleTaskRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetIdType(v string) *DeleteScheduleTaskRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetOrganizationId(v string) *DeleteScheduleTaskRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
