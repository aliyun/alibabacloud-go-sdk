// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWeatherRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetWeatherRequestDeviceInfo) *GetWeatherRequest
	GetDeviceInfo() *GetWeatherRequestDeviceInfo
	SetPayload(v *GetWeatherRequestPayload) *GetWeatherRequest
	GetPayload() *GetWeatherRequestPayload
	SetUserInfo(v *GetWeatherRequestUserInfo) *GetWeatherRequest
	GetUserInfo() *GetWeatherRequestUserInfo
}

type GetWeatherRequest struct {
	// This parameter is required.
	DeviceInfo *GetWeatherRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// if can be null:
	// false
	Payload *GetWeatherRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *GetWeatherRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetWeatherRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherRequest) GoString() string {
	return s.String()
}

func (s *GetWeatherRequest) GetDeviceInfo() *GetWeatherRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetWeatherRequest) GetPayload() *GetWeatherRequestPayload {
	return s.Payload
}

func (s *GetWeatherRequest) GetUserInfo() *GetWeatherRequestUserInfo {
	return s.UserInfo
}

func (s *GetWeatherRequest) SetDeviceInfo(v *GetWeatherRequestDeviceInfo) *GetWeatherRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetWeatherRequest) SetPayload(v *GetWeatherRequestPayload) *GetWeatherRequest {
	s.Payload = v
	return s
}

func (s *GetWeatherRequest) SetUserInfo(v *GetWeatherRequestUserInfo) *GetWeatherRequest {
	s.UserInfo = v
	return s
}

func (s *GetWeatherRequest) Validate() error {
	return dara.Validate(s)
}

type GetWeatherRequestDeviceInfo struct {
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

func (s GetWeatherRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetWeatherRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetWeatherRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetWeatherRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetWeatherRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetWeatherRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetWeatherRequestDeviceInfo) SetEncodeKey(v string) *GetWeatherRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetEncodeType(v string) *GetWeatherRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetId(v string) *GetWeatherRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetIdType(v string) *GetWeatherRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetOrganizationId(v string) *GetWeatherRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetWeatherRequestPayload struct {
}

func (s GetWeatherRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherRequestPayload) GoString() string {
	return s.String()
}

func (s *GetWeatherRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetWeatherRequestUserInfo struct {
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

func (s GetWeatherRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetWeatherRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetWeatherRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetWeatherRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetWeatherRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetWeatherRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetWeatherRequestUserInfo) SetEncodeKey(v string) *GetWeatherRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetEncodeType(v string) *GetWeatherRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetId(v string) *GetWeatherRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetIdType(v string) *GetWeatherRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetOrganizationId(v string) *GetWeatherRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetWeatherRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
