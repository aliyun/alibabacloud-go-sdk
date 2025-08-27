// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupUserSaveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseCityCode(v string) *GroupUserSaveShrinkRequest
	GetBaseCityCode() *string
	SetBirthday(v string) *GroupUserSaveShrinkRequest
	GetBirthday() *string
	SetCertListShrink(v string) *GroupUserSaveShrinkRequest
	GetCertListShrink() *string
	SetGender(v string) *GroupUserSaveShrinkRequest
	GetGender() *string
	SetJobNo(v string) *GroupUserSaveShrinkRequest
	GetJobNo() *string
	SetPhone(v string) *GroupUserSaveShrinkRequest
	GetPhone() *string
	SetRealNameEn(v string) *GroupUserSaveShrinkRequest
	GetRealNameEn() *string
	SetSubCorpIdListShrink(v string) *GroupUserSaveShrinkRequest
	GetSubCorpIdListShrink() *string
	SetUserId(v string) *GroupUserSaveShrinkRequest
	GetUserId() *string
	SetUserName(v string) *GroupUserSaveShrinkRequest
	GetUserName() *string
}

type GroupUserSaveShrinkRequest struct {
	BaseCityCode   *string `json:"base_city_code,omitempty" xml:"base_city_code,omitempty"`
	Birthday       *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	CertListShrink *string `json:"cert_list,omitempty" xml:"cert_list,omitempty"`
	Gender         *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// 1001
	JobNo *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// example:
	//
	// 18000000000
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// ce/shi
	RealNameEn *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
	// This parameter is required.
	SubCorpIdListShrink *string `json:"sub_corp_id_list,omitempty" xml:"sub_corp_id_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// This parameter is required.
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s GroupUserSaveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GroupUserSaveShrinkRequest) GoString() string {
	return s.String()
}

func (s *GroupUserSaveShrinkRequest) GetBaseCityCode() *string {
	return s.BaseCityCode
}

func (s *GroupUserSaveShrinkRequest) GetBirthday() *string {
	return s.Birthday
}

func (s *GroupUserSaveShrinkRequest) GetCertListShrink() *string {
	return s.CertListShrink
}

func (s *GroupUserSaveShrinkRequest) GetGender() *string {
	return s.Gender
}

func (s *GroupUserSaveShrinkRequest) GetJobNo() *string {
	return s.JobNo
}

func (s *GroupUserSaveShrinkRequest) GetPhone() *string {
	return s.Phone
}

func (s *GroupUserSaveShrinkRequest) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *GroupUserSaveShrinkRequest) GetSubCorpIdListShrink() *string {
	return s.SubCorpIdListShrink
}

func (s *GroupUserSaveShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *GroupUserSaveShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *GroupUserSaveShrinkRequest) SetBaseCityCode(v string) *GroupUserSaveShrinkRequest {
	s.BaseCityCode = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) SetBirthday(v string) *GroupUserSaveShrinkRequest {
	s.Birthday = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) SetCertListShrink(v string) *GroupUserSaveShrinkRequest {
	s.CertListShrink = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) SetGender(v string) *GroupUserSaveShrinkRequest {
	s.Gender = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) SetJobNo(v string) *GroupUserSaveShrinkRequest {
	s.JobNo = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) SetPhone(v string) *GroupUserSaveShrinkRequest {
	s.Phone = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) SetRealNameEn(v string) *GroupUserSaveShrinkRequest {
	s.RealNameEn = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) SetSubCorpIdListShrink(v string) *GroupUserSaveShrinkRequest {
	s.SubCorpIdListShrink = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) SetUserId(v string) *GroupUserSaveShrinkRequest {
	s.UserId = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) SetUserName(v string) *GroupUserSaveShrinkRequest {
	s.UserName = &v
	return s
}

func (s *GroupUserSaveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
