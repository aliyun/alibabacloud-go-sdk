// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetScheduleTaskRequestDeviceInfo) *GetScheduleTaskRequest
	GetDeviceInfo() *GetScheduleTaskRequestDeviceInfo
	SetPayload(v *GetScheduleTaskRequestPayload) *GetScheduleTaskRequest
	GetPayload() *GetScheduleTaskRequestPayload
	SetUserInfo(v *GetScheduleTaskRequestUserInfo) *GetScheduleTaskRequest
	GetUserInfo() *GetScheduleTaskRequestUserInfo
}

type GetScheduleTaskRequest struct {
	// This parameter is required.
	DeviceInfo *GetScheduleTaskRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *GetScheduleTaskRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *GetScheduleTaskRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequest) GetDeviceInfo() *GetScheduleTaskRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetScheduleTaskRequest) GetPayload() *GetScheduleTaskRequestPayload {
	return s.Payload
}

func (s *GetScheduleTaskRequest) GetUserInfo() *GetScheduleTaskRequestUserInfo {
	return s.UserInfo
}

func (s *GetScheduleTaskRequest) SetDeviceInfo(v *GetScheduleTaskRequestDeviceInfo) *GetScheduleTaskRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetScheduleTaskRequest) SetPayload(v *GetScheduleTaskRequestPayload) *GetScheduleTaskRequest {
	s.Payload = v
	return s
}

func (s *GetScheduleTaskRequest) SetUserInfo(v *GetScheduleTaskRequestUserInfo) *GetScheduleTaskRequest {
	s.UserInfo = v
	return s
}

func (s *GetScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}

type GetScheduleTaskRequestDeviceInfo struct {
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

func (s GetScheduleTaskRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetScheduleTaskRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetScheduleTaskRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetScheduleTaskRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetScheduleTaskRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetScheduleTaskRequestDeviceInfo) SetEncodeKey(v string) *GetScheduleTaskRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetEncodeType(v string) *GetScheduleTaskRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetId(v string) *GetScheduleTaskRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetIdType(v string) *GetScheduleTaskRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetOrganizationId(v string) *GetScheduleTaskRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetScheduleTaskRequestPayload struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetScheduleTaskRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskRequestPayload) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequestPayload) GetId() *int64 {
	return s.Id
}

func (s *GetScheduleTaskRequestPayload) SetId(v int64) *GetScheduleTaskRequestPayload {
	s.Id = &v
	return s
}

func (s *GetScheduleTaskRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetScheduleTaskRequestUserInfo struct {
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

func (s GetScheduleTaskRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetScheduleTaskRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetScheduleTaskRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetScheduleTaskRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetScheduleTaskRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetScheduleTaskRequestUserInfo) SetEncodeKey(v string) *GetScheduleTaskRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetEncodeType(v string) *GetScheduleTaskRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetId(v string) *GetScheduleTaskRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetIdType(v string) *GetScheduleTaskRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetOrganizationId(v string) *GetScheduleTaskRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
