// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegistrantProfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryRegistrantProfilesResponseBody
	GetCurrentPageNum() *int32
	SetNextPage(v bool) *QueryRegistrantProfilesResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryRegistrantProfilesResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryRegistrantProfilesResponseBody
	GetPrePage() *bool
	SetRegistrantProfiles(v *QueryRegistrantProfilesResponseBodyRegistrantProfiles) *QueryRegistrantProfilesResponseBody
	GetRegistrantProfiles() *QueryRegistrantProfilesResponseBodyRegistrantProfiles
	SetRequestId(v string) *QueryRegistrantProfilesResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryRegistrantProfilesResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryRegistrantProfilesResponseBody
	GetTotalPageNum() *int32
}

type QueryRegistrantProfilesResponseBody struct {
	CurrentPageNum     *int32                                                 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	NextPage           *bool                                                  `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize           *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage            *bool                                                  `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RegistrantProfiles *QueryRegistrantProfilesResponseBodyRegistrantProfiles `json:"RegistrantProfiles,omitempty" xml:"RegistrantProfiles,omitempty" type:"Struct"`
	RequestId          *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum       *int32                                                 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum       *int32                                                 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryRegistrantProfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryRegistrantProfilesResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryRegistrantProfilesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRegistrantProfilesResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryRegistrantProfilesResponseBody) GetRegistrantProfiles() *QueryRegistrantProfilesResponseBodyRegistrantProfiles {
	return s.RegistrantProfiles
}

func (s *QueryRegistrantProfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRegistrantProfilesResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryRegistrantProfilesResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryRegistrantProfilesResponseBody) SetCurrentPageNum(v int32) *QueryRegistrantProfilesResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetNextPage(v bool) *QueryRegistrantProfilesResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetPageSize(v int32) *QueryRegistrantProfilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetPrePage(v bool) *QueryRegistrantProfilesResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetRegistrantProfiles(v *QueryRegistrantProfilesResponseBodyRegistrantProfiles) *QueryRegistrantProfilesResponseBody {
	s.RegistrantProfiles = v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetRequestId(v string) *QueryRegistrantProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetTotalItemNum(v int32) *QueryRegistrantProfilesResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetTotalPageNum(v int32) *QueryRegistrantProfilesResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) Validate() error {
	if s.RegistrantProfiles != nil {
		if err := s.RegistrantProfiles.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRegistrantProfilesResponseBodyRegistrantProfiles struct {
	RegistrantProfile []*QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile `json:"RegistrantProfile,omitempty" xml:"RegistrantProfile,omitempty" type:"Repeated"`
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfiles) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfiles) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfiles) GetRegistrantProfile() []*QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	return s.RegistrantProfile
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfiles) SetRegistrantProfile(v []*QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) *QueryRegistrantProfilesResponseBodyRegistrantProfiles {
	s.RegistrantProfile = v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfiles) Validate() error {
	if s.RegistrantProfile != nil {
		for _, item := range s.RegistrantProfile {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile struct {
	Address                  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	City                     *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DefaultRegistrantProfile *bool   `json:"DefaultRegistrantProfile,omitempty" xml:"DefaultRegistrantProfile,omitempty"`
	Email                    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EmailVerificationStatus  *int32  `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	PostalCode               *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                 *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RealNameStatus           *string `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	RegistrantName           *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization   *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantProfileId      *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RegistrantProfileType    *string `json:"RegistrantProfileType,omitempty" xml:"RegistrantProfileType,omitempty"`
	RegistrantType           *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	TelArea                  *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                   *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	UpdateTime               *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetAddress() *string {
	return s.Address
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetCity() *string {
	return s.City
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetCountry() *string {
	return s.Country
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetDefaultRegistrantProfile() *bool {
	return s.DefaultRegistrantProfile
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetEmail() *string {
	return s.Email
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetEmailVerificationStatus() *int32 {
	return s.EmailVerificationStatus
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetPostalCode() *string {
	return s.PostalCode
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetProvince() *string {
	return s.Province
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRealNameStatus() *string {
	return s.RealNameStatus
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantProfileType() *string {
	return s.RegistrantProfileType
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetRegistrantType() *string {
	return s.RegistrantType
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetTelArea() *string {
	return s.TelArea
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetTelExt() *string {
	return s.TelExt
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetTelephone() *string {
	return s.Telephone
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetAddress(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Address = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCity(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.City = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCountry(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Country = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCreateTime(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.CreateTime = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetDefaultRegistrantProfile(v bool) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.DefaultRegistrantProfile = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetEmail(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Email = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetEmailVerificationStatus(v int32) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.EmailVerificationStatus = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetPostalCode(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.PostalCode = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetProvince(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Province = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRealNameStatus(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RealNameStatus = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantName(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantName = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantOrganization(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantProfileId(v int64) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantProfileType(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantProfileType = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantType(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantType = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetTelArea(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.TelArea = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetTelExt(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.TelExt = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetTelephone(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Telephone = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetUpdateTime(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.UpdateTime = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) Validate() error {
	return dara.Validate(s)
}
