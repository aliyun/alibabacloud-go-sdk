// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListAlarmsRequestDeviceInfo) *ListAlarmsRequest
	GetDeviceInfo() *ListAlarmsRequestDeviceInfo
	SetPayload(v *ListAlarmsRequestPayload) *ListAlarmsRequest
	GetPayload() *ListAlarmsRequestPayload
	SetUserInfo(v *ListAlarmsRequestUserInfo) *ListAlarmsRequest
	GetUserInfo() *ListAlarmsRequestUserInfo
}

type ListAlarmsRequest struct {
	// This parameter is required.
	DeviceInfo *ListAlarmsRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Payload *ListAlarmsRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListAlarmsRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListAlarmsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequest) GetDeviceInfo() *ListAlarmsRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListAlarmsRequest) GetPayload() *ListAlarmsRequestPayload {
	return s.Payload
}

func (s *ListAlarmsRequest) GetUserInfo() *ListAlarmsRequestUserInfo {
	return s.UserInfo
}

func (s *ListAlarmsRequest) SetDeviceInfo(v *ListAlarmsRequestDeviceInfo) *ListAlarmsRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListAlarmsRequest) SetPayload(v *ListAlarmsRequestPayload) *ListAlarmsRequest {
	s.Payload = v
	return s
}

func (s *ListAlarmsRequest) SetUserInfo(v *ListAlarmsRequestUserInfo) *ListAlarmsRequest {
	s.UserInfo = v
	return s
}

func (s *ListAlarmsRequest) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsRequestDeviceInfo struct {
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

func (s ListAlarmsRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListAlarmsRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListAlarmsRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListAlarmsRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListAlarmsRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListAlarmsRequestDeviceInfo) SetEncodeKey(v string) *ListAlarmsRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetEncodeType(v string) *ListAlarmsRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetId(v string) *ListAlarmsRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetIdType(v string) *ListAlarmsRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetOrganizationId(v string) *ListAlarmsRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsRequestPayload struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAlarmsRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsRequestPayload) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequestPayload) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAlarmsRequestPayload) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlarmsRequestPayload) SetCurrentPage(v int32) *ListAlarmsRequestPayload {
	s.CurrentPage = &v
	return s
}

func (s *ListAlarmsRequestPayload) SetPageSize(v int32) *ListAlarmsRequestPayload {
	s.PageSize = &v
	return s
}

func (s *ListAlarmsRequestPayload) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsRequestUserInfo struct {
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

func (s ListAlarmsRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListAlarmsRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListAlarmsRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListAlarmsRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListAlarmsRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListAlarmsRequestUserInfo) SetEncodeKey(v string) *ListAlarmsRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetEncodeType(v string) *ListAlarmsRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetId(v string) *ListAlarmsRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetIdType(v string) *ListAlarmsRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetOrganizationId(v string) *ListAlarmsRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
