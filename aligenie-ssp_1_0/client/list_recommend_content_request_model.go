// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecommendContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListRecommendContentRequestDeviceInfo) *ListRecommendContentRequest
	GetDeviceInfo() *ListRecommendContentRequestDeviceInfo
	SetRequest(v *ListRecommendContentRequestRequest) *ListRecommendContentRequest
	GetRequest() *ListRecommendContentRequestRequest
	SetUserInfo(v *ListRecommendContentRequestUserInfo) *ListRecommendContentRequest
	GetUserInfo() *ListRecommendContentRequestUserInfo
}

type ListRecommendContentRequest struct {
	// This parameter is required.
	DeviceInfo *ListRecommendContentRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	Request *ListRecommendContentRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ListRecommendContentRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListRecommendContentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentRequest) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequest) GetDeviceInfo() *ListRecommendContentRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListRecommendContentRequest) GetRequest() *ListRecommendContentRequestRequest {
	return s.Request
}

func (s *ListRecommendContentRequest) GetUserInfo() *ListRecommendContentRequestUserInfo {
	return s.UserInfo
}

func (s *ListRecommendContentRequest) SetDeviceInfo(v *ListRecommendContentRequestDeviceInfo) *ListRecommendContentRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListRecommendContentRequest) SetRequest(v *ListRecommendContentRequestRequest) *ListRecommendContentRequest {
	s.Request = v
	return s
}

func (s *ListRecommendContentRequest) SetUserInfo(v *ListRecommendContentRequestUserInfo) *ListRecommendContentRequest {
	s.UserInfo = v
	return s
}

func (s *ListRecommendContentRequest) Validate() error {
	return dara.Validate(s)
}

type ListRecommendContentRequestDeviceInfo struct {
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

func (s ListRecommendContentRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListRecommendContentRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListRecommendContentRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListRecommendContentRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListRecommendContentRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRecommendContentRequestDeviceInfo) SetEncodeKey(v string) *ListRecommendContentRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetEncodeType(v string) *ListRecommendContentRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetId(v string) *ListRecommendContentRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetIdType(v string) *ListRecommendContentRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetOrganizationId(v string) *ListRecommendContentRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListRecommendContentRequestRequest struct {
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// song
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecommendContentRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentRequestRequest) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequestRequest) GetCount() *int32 {
	return s.Count
}

func (s *ListRecommendContentRequestRequest) GetType() *string {
	return s.Type
}

func (s *ListRecommendContentRequestRequest) SetCount(v int32) *ListRecommendContentRequestRequest {
	s.Count = &v
	return s
}

func (s *ListRecommendContentRequestRequest) SetType(v string) *ListRecommendContentRequestRequest {
	s.Type = &v
	return s
}

func (s *ListRecommendContentRequestRequest) Validate() error {
	return dara.Validate(s)
}

type ListRecommendContentRequestUserInfo struct {
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

func (s ListRecommendContentRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListRecommendContentRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListRecommendContentRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListRecommendContentRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListRecommendContentRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRecommendContentRequestUserInfo) SetEncodeKey(v string) *ListRecommendContentRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetEncodeType(v string) *ListRecommendContentRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetId(v string) *ListRecommendContentRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetIdType(v string) *ListRecommendContentRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetOrganizationId(v string) *ListRecommendContentRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
