// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAuthCodeBindForExtRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *CheckAuthCodeBindForExtRequest
	GetAuthCode() *string
	SetEncodeKey(v string) *CheckAuthCodeBindForExtRequest
	GetEncodeKey() *string
	SetEncodeType(v string) *CheckAuthCodeBindForExtRequest
	GetEncodeType() *string
	SetUserInfo(v *CheckAuthCodeBindForExtRequestUserInfo) *CheckAuthCodeBindForExtRequest
	GetUserInfo() *CheckAuthCodeBindForExtRequestUserInfo
}

type CheckAuthCodeBindForExtRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Aexfgc
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
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
	UserInfo *CheckAuthCodeBindForExtRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CheckAuthCodeBindForExtRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAuthCodeBindForExtRequest) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *CheckAuthCodeBindForExtRequest) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CheckAuthCodeBindForExtRequest) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CheckAuthCodeBindForExtRequest) GetUserInfo() *CheckAuthCodeBindForExtRequestUserInfo {
	return s.UserInfo
}

func (s *CheckAuthCodeBindForExtRequest) SetAuthCode(v string) *CheckAuthCodeBindForExtRequest {
	s.AuthCode = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequest) SetEncodeKey(v string) *CheckAuthCodeBindForExtRequest {
	s.EncodeKey = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequest) SetEncodeType(v string) *CheckAuthCodeBindForExtRequest {
	s.EncodeType = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequest) SetUserInfo(v *CheckAuthCodeBindForExtRequestUserInfo) *CheckAuthCodeBindForExtRequest {
	s.UserInfo = v
	return s
}

func (s *CheckAuthCodeBindForExtRequest) Validate() error {
	return dara.Validate(s)
}

type CheckAuthCodeBindForExtRequestUserInfo struct {
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
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1***2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CheckAuthCodeBindForExtRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckAuthCodeBindForExtRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetEncodeKey(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetEncodeType(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetId(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetIdType(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) SetOrganizationId(v string) *CheckAuthCodeBindForExtRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *CheckAuthCodeBindForExtRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
