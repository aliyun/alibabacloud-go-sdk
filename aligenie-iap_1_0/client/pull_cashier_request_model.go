// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullCashierRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *PullCashierRequestDeviceInfo) *PullCashierRequest
	GetDeviceInfo() *PullCashierRequestDeviceInfo
	SetPayload(v *PullCashierRequestPayload) *PullCashierRequest
	GetPayload() *PullCashierRequestPayload
	SetUserInfo(v *PullCashierRequestUserInfo) *PullCashierRequest
	GetUserInfo() *PullCashierRequestUserInfo
}

type PullCashierRequest struct {
	DeviceInfo *PullCashierRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *PullCashierRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *PullCashierRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PullCashierRequest) String() string {
	return dara.Prettify(s)
}

func (s PullCashierRequest) GoString() string {
	return s.String()
}

func (s *PullCashierRequest) GetDeviceInfo() *PullCashierRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *PullCashierRequest) GetPayload() *PullCashierRequestPayload {
	return s.Payload
}

func (s *PullCashierRequest) GetUserInfo() *PullCashierRequestUserInfo {
	return s.UserInfo
}

func (s *PullCashierRequest) SetDeviceInfo(v *PullCashierRequestDeviceInfo) *PullCashierRequest {
	s.DeviceInfo = v
	return s
}

func (s *PullCashierRequest) SetPayload(v *PullCashierRequestPayload) *PullCashierRequest {
	s.Payload = v
	return s
}

func (s *PullCashierRequest) SetUserInfo(v *PullCashierRequestUserInfo) *PullCashierRequest {
	s.UserInfo = v
	return s
}

func (s *PullCashierRequest) Validate() error {
	return dara.Validate(s)
}

type PullCashierRequestDeviceInfo struct {
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

func (s PullCashierRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s PullCashierRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PullCashierRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PullCashierRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PullCashierRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *PullCashierRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *PullCashierRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PullCashierRequestDeviceInfo) SetEncodeKey(v string) *PullCashierRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetEncodeType(v string) *PullCashierRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetId(v string) *PullCashierRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetIdType(v string) *PullCashierRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetOrganizationId(v string) *PullCashierRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type PullCashierRequestPayload struct {
	// example:
	//
	// 731D5F********DC3B
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
}

func (s PullCashierRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s PullCashierRequestPayload) GoString() string {
	return s.String()
}

func (s *PullCashierRequestPayload) GetOriginUuid() *string {
	return s.OriginUuid
}

func (s *PullCashierRequestPayload) SetOriginUuid(v string) *PullCashierRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *PullCashierRequestPayload) Validate() error {
	return dara.Validate(s)
}

type PullCashierRequestUserInfo struct {
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

func (s PullCashierRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s PullCashierRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PullCashierRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PullCashierRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PullCashierRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *PullCashierRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *PullCashierRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PullCashierRequestUserInfo) SetEncodeKey(v string) *PullCashierRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetEncodeType(v string) *PullCashierRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetId(v string) *PullCashierRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetIdType(v string) *PullCashierRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetOrganizationId(v string) *PullCashierRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *PullCashierRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
