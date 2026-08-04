// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *SendMessageRequest
	GetUrl() *string
	SetUserInfo(v *SendMessageRequestUserInfo) *SendMessageRequest
	GetUserInfo() *SendMessageRequestUserInfo
}

type SendMessageRequest struct {
	// Message URL
	//
	// example:
	//
	// http://xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	UserInfo *SendMessageRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s SendMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) GetUrl() *string {
	return s.Url
}

func (s *SendMessageRequest) GetUserInfo() *SendMessageRequestUserInfo {
	return s.UserInfo
}

func (s *SendMessageRequest) SetUrl(v string) *SendMessageRequest {
	s.Url = &v
	return s
}

func (s *SendMessageRequest) SetUserInfo(v *SendMessageRequestUserInfo) *SendMessageRequest {
	s.UserInfo = v
	return s
}

func (s *SendMessageRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendMessageRequestUserInfo struct {
	// Value corresponding to the encoding type. When the encoding type is SKILLID, this value is the application\\"s Skill ID. When the encoding type is PACKAGENAME, this value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123L
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the User Identifier for Maojing, and each method corresponds to a different encoding type: - PACKAGENAME: APK package name, used for Android application customer links - SKILLID: Skill ID, used for cloud-based links
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier (userOpenId or userUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// 123L
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of User ID: - OPENID: default User ID identity - UNIONID: organization-dimension User ID identity, available only after an organization has been requested on the Maojing Skill Application Open Platform
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID; required when IdType is UNION_ID
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SendMessageRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequestUserInfo) GoString() string {
	return s.String()
}

func (s *SendMessageRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *SendMessageRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *SendMessageRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *SendMessageRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *SendMessageRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *SendMessageRequestUserInfo) SetEncodeKey(v string) *SendMessageRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *SendMessageRequestUserInfo) SetEncodeType(v string) *SendMessageRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *SendMessageRequestUserInfo) SetId(v string) *SendMessageRequestUserInfo {
	s.Id = &v
	return s
}

func (s *SendMessageRequestUserInfo) SetIdType(v string) *SendMessageRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *SendMessageRequestUserInfo) SetOrganizationId(v string) *SendMessageRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *SendMessageRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
