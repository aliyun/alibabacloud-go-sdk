// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountForAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetAccountForAppRequestDeviceInfo) *GetAccountForAppRequest
	GetDeviceInfo() *GetAccountForAppRequestDeviceInfo
	SetPayload(v *GetAccountForAppRequestPayload) *GetAccountForAppRequest
	GetPayload() *GetAccountForAppRequestPayload
	SetUserInfo(v *GetAccountForAppRequestUserInfo) *GetAccountForAppRequest
	GetUserInfo() *GetAccountForAppRequestUserInfo
}

type GetAccountForAppRequest struct {
	DeviceInfo *GetAccountForAppRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *GetAccountForAppRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *GetAccountForAppRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetAccountForAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppRequest) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequest) GetDeviceInfo() *GetAccountForAppRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetAccountForAppRequest) GetPayload() *GetAccountForAppRequestPayload {
	return s.Payload
}

func (s *GetAccountForAppRequest) GetUserInfo() *GetAccountForAppRequestUserInfo {
	return s.UserInfo
}

func (s *GetAccountForAppRequest) SetDeviceInfo(v *GetAccountForAppRequestDeviceInfo) *GetAccountForAppRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetAccountForAppRequest) SetPayload(v *GetAccountForAppRequestPayload) *GetAccountForAppRequest {
	s.Payload = v
	return s
}

func (s *GetAccountForAppRequest) SetUserInfo(v *GetAccountForAppRequestUserInfo) *GetAccountForAppRequest {
	s.UserInfo = v
	return s
}

func (s *GetAccountForAppRequest) Validate() error {
	return dara.Validate(s)
}

type GetAccountForAppRequestDeviceInfo struct {
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// PACKAGE_NAME
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

func (s GetAccountForAppRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetAccountForAppRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetAccountForAppRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetAccountForAppRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetAccountForAppRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetAccountForAppRequestDeviceInfo) SetEncodeKey(v string) *GetAccountForAppRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetEncodeType(v string) *GetAccountForAppRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetId(v string) *GetAccountForAppRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetIdType(v string) *GetAccountForAppRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetOrganizationId(v string) *GetAccountForAppRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetAccountForAppRequestPayload struct {
	// example:
	//
	// 188*777
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// example:
	//
	// 731D5F********DC3B
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
}

func (s GetAccountForAppRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppRequestPayload) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequestPayload) GetPhone() *string {
	return s.Phone
}

func (s *GetAccountForAppRequestPayload) GetOriginUuid() *string {
	return s.OriginUuid
}

func (s *GetAccountForAppRequestPayload) SetPhone(v string) *GetAccountForAppRequestPayload {
	s.Phone = &v
	return s
}

func (s *GetAccountForAppRequestPayload) SetOriginUuid(v string) *GetAccountForAppRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *GetAccountForAppRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetAccountForAppRequestUserInfo struct {
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

func (s GetAccountForAppRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetAccountForAppRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetAccountForAppRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetAccountForAppRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetAccountForAppRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetAccountForAppRequestUserInfo) SetEncodeKey(v string) *GetAccountForAppRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetEncodeType(v string) *GetAccountForAppRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetId(v string) *GetAccountForAppRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetIdType(v string) *GetAccountForAppRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetOrganizationId(v string) *GetAccountForAppRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
