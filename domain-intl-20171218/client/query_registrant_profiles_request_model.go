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
	SetUserClientIp(v string) *QueryRegistrantProfilesRequest
	GetUserClientIp() *string
}

type QueryRegistrantProfilesRequest struct {
	DefaultRegistrantProfile *bool   `json:"DefaultRegistrantProfile,omitempty" xml:"DefaultRegistrantProfile,omitempty"`
	Email                    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang                     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum                  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize                 *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RealNameStatus           *string `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	RegistrantOrganization   *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantProfileId      *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RegistrantProfileType    *string `json:"RegistrantProfileType,omitempty" xml:"RegistrantProfileType,omitempty"`
	RegistrantType           *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	UserClientIp             *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
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

func (s *QueryRegistrantProfilesRequest) GetUserClientIp() *string {
	return s.UserClientIp
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

func (s *QueryRegistrantProfilesRequest) SetUserClientIp(v string) *QueryRegistrantProfilesRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) Validate() error {
	return dara.Validate(s)
}
