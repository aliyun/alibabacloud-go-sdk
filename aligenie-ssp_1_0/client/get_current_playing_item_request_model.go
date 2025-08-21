// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetCurrentPlayingItemRequestDeviceInfo) *GetCurrentPlayingItemRequest
	GetDeviceInfo() *GetCurrentPlayingItemRequestDeviceInfo
	SetUserInfo(v *GetCurrentPlayingItemRequestUserInfo) *GetCurrentPlayingItemRequest
	GetUserInfo() *GetCurrentPlayingItemRequestUserInfo
}

type GetCurrentPlayingItemRequest struct {
	// This parameter is required.
	DeviceInfo *GetCurrentPlayingItemRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *GetCurrentPlayingItemRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetCurrentPlayingItemRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemRequest) GetDeviceInfo() *GetCurrentPlayingItemRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetCurrentPlayingItemRequest) GetUserInfo() *GetCurrentPlayingItemRequestUserInfo {
	return s.UserInfo
}

func (s *GetCurrentPlayingItemRequest) SetDeviceInfo(v *GetCurrentPlayingItemRequestDeviceInfo) *GetCurrentPlayingItemRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetCurrentPlayingItemRequest) SetUserInfo(v *GetCurrentPlayingItemRequestUserInfo) *GetCurrentPlayingItemRequest {
	s.UserInfo = v
	return s
}

func (s *GetCurrentPlayingItemRequest) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingItemRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCurrentPlayingItemRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetEncodeKey(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetEncodeType(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetId(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetIdType(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetOrganizationId(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingItemRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCurrentPlayingItemRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetEncodeKey(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetEncodeType(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetId(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetIdType(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetOrganizationId(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
