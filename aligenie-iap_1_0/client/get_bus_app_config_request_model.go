// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetBusAppConfigRequestDeviceInfo) *GetBusAppConfigRequest
	GetDeviceInfo() *GetBusAppConfigRequestDeviceInfo
	SetPayload(v *GetBusAppConfigRequestPayload) *GetBusAppConfigRequest
	GetPayload() *GetBusAppConfigRequestPayload
	SetUserInfo(v *GetBusAppConfigRequestUserInfo) *GetBusAppConfigRequest
	GetUserInfo() *GetBusAppConfigRequestUserInfo
}

type GetBusAppConfigRequest struct {
	DeviceInfo *GetBusAppConfigRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *GetBusAppConfigRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *GetBusAppConfigRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetBusAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigRequest) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequest) GetDeviceInfo() *GetBusAppConfigRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetBusAppConfigRequest) GetPayload() *GetBusAppConfigRequestPayload {
	return s.Payload
}

func (s *GetBusAppConfigRequest) GetUserInfo() *GetBusAppConfigRequestUserInfo {
	return s.UserInfo
}

func (s *GetBusAppConfigRequest) SetDeviceInfo(v *GetBusAppConfigRequestDeviceInfo) *GetBusAppConfigRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetBusAppConfigRequest) SetPayload(v *GetBusAppConfigRequestPayload) *GetBusAppConfigRequest {
	s.Payload = v
	return s
}

func (s *GetBusAppConfigRequest) SetUserInfo(v *GetBusAppConfigRequestUserInfo) *GetBusAppConfigRequest {
	s.UserInfo = v
	return s
}

func (s *GetBusAppConfigRequest) Validate() error {
	return dara.Validate(s)
}

type GetBusAppConfigRequestDeviceInfo struct {
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// SKILL_ID
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

func (s GetBusAppConfigRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetBusAppConfigRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetBusAppConfigRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetBusAppConfigRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetBusAppConfigRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetBusAppConfigRequestDeviceInfo) SetEncodeKey(v string) *GetBusAppConfigRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetEncodeType(v string) *GetBusAppConfigRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetId(v string) *GetBusAppConfigRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetIdType(v string) *GetBusAppConfigRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetOrganizationId(v string) *GetBusAppConfigRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetBusAppConfigRequestPayload struct {
	// example:
	//
	// 731D5F********DC3B
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
	// example:
	//
	// 136****1111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s GetBusAppConfigRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigRequestPayload) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequestPayload) GetOriginUuid() *string {
	return s.OriginUuid
}

func (s *GetBusAppConfigRequestPayload) GetPhone() *string {
	return s.Phone
}

func (s *GetBusAppConfigRequestPayload) SetOriginUuid(v string) *GetBusAppConfigRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *GetBusAppConfigRequestPayload) SetPhone(v string) *GetBusAppConfigRequestPayload {
	s.Phone = &v
	return s
}

func (s *GetBusAppConfigRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetBusAppConfigRequestUserInfo struct {
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

func (s GetBusAppConfigRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetBusAppConfigRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetBusAppConfigRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetBusAppConfigRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetBusAppConfigRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetBusAppConfigRequestUserInfo) SetEncodeKey(v string) *GetBusAppConfigRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetEncodeType(v string) *GetBusAppConfigRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetId(v string) *GetBusAppConfigRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetIdType(v string) *GetBusAppConfigRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetOrganizationId(v string) *GetBusAppConfigRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
