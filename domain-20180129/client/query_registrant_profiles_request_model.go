// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegistrantProfilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultRegistrantProfile(v bool) *QueryRegistrantProfilesRequest
	GetDefaultRegistrantProfile() *bool
	SetEmail(v string) *QueryRegistrantProfilesRequest
	GetEmail() *string
	SetLang(v string) *QueryRegistrantProfilesRequest
	GetLang() *string
	SetPageNum(v int32) *QueryRegistrantProfilesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryRegistrantProfilesRequest
	GetPageSize() *int32
	SetRealNameStatus(v string) *QueryRegistrantProfilesRequest
	GetRealNameStatus() *string
	SetRegistrantOrganization(v string) *QueryRegistrantProfilesRequest
	GetRegistrantOrganization() *string
	SetRegistrantProfileId(v int64) *QueryRegistrantProfilesRequest
	GetRegistrantProfileId() *int64
	SetRegistrantProfileType(v string) *QueryRegistrantProfilesRequest
	GetRegistrantProfileType() *string
	SetRegistrantType(v string) *QueryRegistrantProfilesRequest
	GetRegistrantType() *string
	SetRemark(v string) *QueryRegistrantProfilesRequest
	GetRemark() *string
	SetUserClientIp(v string) *QueryRegistrantProfilesRequest
	GetUserClientIp() *string
	SetZhRegistrantOrganization(v string) *QueryRegistrantProfilesRequest
	GetZhRegistrantOrganization() *string
}

type QueryRegistrantProfilesRequest struct {
	// Specifies whether to query the default profile. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	DefaultRegistrantProfile *bool `json:"DefaultRegistrantProfile,omitempty" xml:"DefaultRegistrantProfile,omitempty"`
	// The email address of the domain name registrant.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The language of the error message to return if the request fails. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Default value: **0**.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: **0**. Maximum value: **5000**.
	//
	// example:
	//
	// 500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The state of real-name verification for the domain name registrant. Valid values:
	//
	// 	- **FAILED**: Real-name verification for the domain name fails.
	//
	// 	- **SUCCEED**: Real-name verification for the domain name is successful.
	//
	// 	- **NONAUDIT**: Real-name verification for the domain name is not performed.
	//
	// 	- **AUDITING**: Real-name verification for the domain name is in progress.
	//
	// example:
	//
	// SUCCEED
	RealNameStatus *string `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	// The name of the domain name registrant.
	//
	// example:
	//
	// li si
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// The ID of the registrant profile that you want to query. The system generates an ID after you create a registrant profile.
	//
	// example:
	//
	// 1234567
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// The type of the registrant profile. Valid values:
	//
	// 	- **common**: common profile.
	//
	// 	- **cnnic**: China Internet Network Information Center (CNNIC) profile.
	//
	// >  Only the Alibaba Cloud international site (alibabacloud.com) supports CNNIC profiles. To register domain names provided by CNNIC such as the .cn and . domain names on the Alibaba Cloud international site, you must use a CNNIC profile. To register other domain names, use a common profile.
	//
	// example:
	//
	// common
	RegistrantProfileType *string `json:"RegistrantProfileType,omitempty" xml:"RegistrantProfileType,omitempty"`
	// The type of the domain name registrant. Valid values:
	//
	// 	- **1**: individual.
	//
	// 	- **2**: enterprise.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	RegistrantType *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Test domain name
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The IP address of the client. Set the value to 127.0.0.1.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The Chinese name of the domain name registrant.
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryRegistrantProfilesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfilesRequest) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesRequest) GetDefaultRegistrantProfile() *bool {
	return s.DefaultRegistrantProfile
}

func (s *QueryRegistrantProfilesRequest) GetEmail() *string {
	return s.Email
}

func (s *QueryRegistrantProfilesRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryRegistrantProfilesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryRegistrantProfilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRegistrantProfilesRequest) GetRealNameStatus() *string {
	return s.RealNameStatus
}

func (s *QueryRegistrantProfilesRequest) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *QueryRegistrantProfilesRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *QueryRegistrantProfilesRequest) GetRegistrantProfileType() *string {
	return s.RegistrantProfileType
}

func (s *QueryRegistrantProfilesRequest) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *QueryRegistrantProfilesRequest) GetRemark() *string {
	return s.Remark
}

func (s *QueryRegistrantProfilesRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryRegistrantProfilesRequest) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *QueryRegistrantProfilesRequest) SetDefaultRegistrantProfile(v bool) *QueryRegistrantProfilesRequest {
	s.DefaultRegistrantProfile = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetEmail(v string) *QueryRegistrantProfilesRequest {
	s.Email = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetLang(v string) *QueryRegistrantProfilesRequest {
	s.Lang = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetPageNum(v int32) *QueryRegistrantProfilesRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetPageSize(v int32) *QueryRegistrantProfilesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRealNameStatus(v string) *QueryRegistrantProfilesRequest {
	s.RealNameStatus = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRegistrantOrganization(v string) *QueryRegistrantProfilesRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRegistrantProfileId(v int64) *QueryRegistrantProfilesRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRegistrantProfileType(v string) *QueryRegistrantProfilesRequest {
	s.RegistrantProfileType = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRegistrantType(v string) *QueryRegistrantProfilesRequest {
	s.RegistrantType = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRemark(v string) *QueryRegistrantProfilesRequest {
	s.Remark = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetUserClientIp(v string) *QueryRegistrantProfilesRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetZhRegistrantOrganization(v string) *QueryRegistrantProfilesRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) Validate() error {
	return dara.Validate(s)
}
