// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnreadMessageCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfo(v *GetUnreadMessageCountRequestUserInfo) *GetUnreadMessageCountRequest
	GetUserInfo() *GetUnreadMessageCountRequestUserInfo
}

type GetUnreadMessageCountRequest struct {
	// User Identifier information
	UserInfo *GetUnreadMessageCountRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetUnreadMessageCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUnreadMessageCountRequest) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountRequest) GetUserInfo() *GetUnreadMessageCountRequestUserInfo {
	return s.UserInfo
}

func (s *GetUnreadMessageCountRequest) SetUserInfo(v *GetUnreadMessageCountRequestUserInfo) *GetUnreadMessageCountRequest {
	s.UserInfo = v
	return s
}

func (s *GetUnreadMessageCountRequest) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUnreadMessageCountRequestUserInfo struct {
	// Value corresponding to the encoding type. When the encoding type is SKILLID, this value is the application\\"s Skill ID. When the encoding type is PACKAGENAME, this value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier from Maojing, and each method corresponds to a different encoding type: PACKAGENAME for the APK package name used in the Android application client link, and SKILLID for the cloud-based link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier (userOpenId or userUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of User ID: OPENID is the default user ID identity; UNIONID is the organization-dimension user ID identity, which is available only after an organization has been requested on the Maojing Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required when IdType is UNION_ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetUnreadMessageCountRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUnreadMessageCountRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetUnreadMessageCountRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetUnreadMessageCountRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetUnreadMessageCountRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetUnreadMessageCountRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetUnreadMessageCountRequestUserInfo) SetEncodeKey(v string) *GetUnreadMessageCountRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetUnreadMessageCountRequestUserInfo) SetEncodeType(v string) *GetUnreadMessageCountRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetUnreadMessageCountRequestUserInfo) SetId(v string) *GetUnreadMessageCountRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetUnreadMessageCountRequestUserInfo) SetIdType(v string) *GetUnreadMessageCountRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetUnreadMessageCountRequestUserInfo) SetOrganizationId(v string) *GetUnreadMessageCountRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetUnreadMessageCountRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
