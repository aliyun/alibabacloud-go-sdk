// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupUserSaveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseCityCode(v string) *GroupUserSaveRequest
	GetBaseCityCode() *string
	SetBirthday(v string) *GroupUserSaveRequest
	GetBirthday() *string
	SetCertList(v []*GroupUserSaveRequestCertList) *GroupUserSaveRequest
	GetCertList() []*GroupUserSaveRequestCertList
	SetGender(v string) *GroupUserSaveRequest
	GetGender() *string
	SetJobNo(v string) *GroupUserSaveRequest
	GetJobNo() *string
	SetPhone(v string) *GroupUserSaveRequest
	GetPhone() *string
	SetRealNameEn(v string) *GroupUserSaveRequest
	GetRealNameEn() *string
	SetSubCorpIdList(v []*GroupUserSaveRequestSubCorpIdList) *GroupUserSaveRequest
	GetSubCorpIdList() []*GroupUserSaveRequestSubCorpIdList
	SetUserId(v string) *GroupUserSaveRequest
	GetUserId() *string
	SetUserName(v string) *GroupUserSaveRequest
	GetUserName() *string
}

type GroupUserSaveRequest struct {
	BaseCityCode *string                         `json:"base_city_code,omitempty" xml:"base_city_code,omitempty"`
	Birthday     *string                         `json:"birthday,omitempty" xml:"birthday,omitempty"`
	CertList     []*GroupUserSaveRequestCertList `json:"cert_list,omitempty" xml:"cert_list,omitempty" type:"Repeated"`
	Gender       *string                         `json:"gender,omitempty" xml:"gender,omitempty"`
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
	SubCorpIdList []*GroupUserSaveRequestSubCorpIdList `json:"sub_corp_id_list,omitempty" xml:"sub_corp_id_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// This parameter is required.
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s GroupUserSaveRequest) String() string {
	return dara.Prettify(s)
}

func (s GroupUserSaveRequest) GoString() string {
	return s.String()
}

func (s *GroupUserSaveRequest) GetBaseCityCode() *string {
	return s.BaseCityCode
}

func (s *GroupUserSaveRequest) GetBirthday() *string {
	return s.Birthday
}

func (s *GroupUserSaveRequest) GetCertList() []*GroupUserSaveRequestCertList {
	return s.CertList
}

func (s *GroupUserSaveRequest) GetGender() *string {
	return s.Gender
}

func (s *GroupUserSaveRequest) GetJobNo() *string {
	return s.JobNo
}

func (s *GroupUserSaveRequest) GetPhone() *string {
	return s.Phone
}

func (s *GroupUserSaveRequest) GetRealNameEn() *string {
	return s.RealNameEn
}

func (s *GroupUserSaveRequest) GetSubCorpIdList() []*GroupUserSaveRequestSubCorpIdList {
	return s.SubCorpIdList
}

func (s *GroupUserSaveRequest) GetUserId() *string {
	return s.UserId
}

func (s *GroupUserSaveRequest) GetUserName() *string {
	return s.UserName
}

func (s *GroupUserSaveRequest) SetBaseCityCode(v string) *GroupUserSaveRequest {
	s.BaseCityCode = &v
	return s
}

func (s *GroupUserSaveRequest) SetBirthday(v string) *GroupUserSaveRequest {
	s.Birthday = &v
	return s
}

func (s *GroupUserSaveRequest) SetCertList(v []*GroupUserSaveRequestCertList) *GroupUserSaveRequest {
	s.CertList = v
	return s
}

func (s *GroupUserSaveRequest) SetGender(v string) *GroupUserSaveRequest {
	s.Gender = &v
	return s
}

func (s *GroupUserSaveRequest) SetJobNo(v string) *GroupUserSaveRequest {
	s.JobNo = &v
	return s
}

func (s *GroupUserSaveRequest) SetPhone(v string) *GroupUserSaveRequest {
	s.Phone = &v
	return s
}

func (s *GroupUserSaveRequest) SetRealNameEn(v string) *GroupUserSaveRequest {
	s.RealNameEn = &v
	return s
}

func (s *GroupUserSaveRequest) SetSubCorpIdList(v []*GroupUserSaveRequestSubCorpIdList) *GroupUserSaveRequest {
	s.SubCorpIdList = v
	return s
}

func (s *GroupUserSaveRequest) SetUserId(v string) *GroupUserSaveRequest {
	s.UserId = &v
	return s
}

func (s *GroupUserSaveRequest) SetUserName(v string) *GroupUserSaveRequest {
	s.UserName = &v
	return s
}

func (s *GroupUserSaveRequest) Validate() error {
	return dara.Validate(s)
}

type GroupUserSaveRequestCertList struct {
	CertExpiredTime *string `json:"cert_expired_time,omitempty" xml:"cert_expired_time,omitempty"`
	CertNation      *string `json:"cert_nation,omitempty" xml:"cert_nation,omitempty"`
	CertNo          *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	CertType        *int32  `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	Nationality     *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
}

func (s GroupUserSaveRequestCertList) String() string {
	return dara.Prettify(s)
}

func (s GroupUserSaveRequestCertList) GoString() string {
	return s.String()
}

func (s *GroupUserSaveRequestCertList) GetCertExpiredTime() *string {
	return s.CertExpiredTime
}

func (s *GroupUserSaveRequestCertList) GetCertNation() *string {
	return s.CertNation
}

func (s *GroupUserSaveRequestCertList) GetCertNo() *string {
	return s.CertNo
}

func (s *GroupUserSaveRequestCertList) GetCertType() *int32 {
	return s.CertType
}

func (s *GroupUserSaveRequestCertList) GetNationality() *string {
	return s.Nationality
}

func (s *GroupUserSaveRequestCertList) SetCertExpiredTime(v string) *GroupUserSaveRequestCertList {
	s.CertExpiredTime = &v
	return s
}

func (s *GroupUserSaveRequestCertList) SetCertNation(v string) *GroupUserSaveRequestCertList {
	s.CertNation = &v
	return s
}

func (s *GroupUserSaveRequestCertList) SetCertNo(v string) *GroupUserSaveRequestCertList {
	s.CertNo = &v
	return s
}

func (s *GroupUserSaveRequestCertList) SetCertType(v int32) *GroupUserSaveRequestCertList {
	s.CertType = &v
	return s
}

func (s *GroupUserSaveRequestCertList) SetNationality(v string) *GroupUserSaveRequestCertList {
	s.Nationality = &v
	return s
}

func (s *GroupUserSaveRequestCertList) Validate() error {
	return dara.Validate(s)
}

type GroupUserSaveRequestSubCorpIdList struct {
	DepartIds []*string `json:"depart_ids,omitempty" xml:"depart_ids,omitempty" type:"Repeated"`
	Email     *string   `json:"email,omitempty" xml:"email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	LeaveStatus *int32 `json:"leave_status,omitempty" xml:"leave_status,omitempty"`
	// example:
	//
	// 123
	ManagerUserId *string `json:"manager_user_id,omitempty" xml:"manager_user_id,omitempty"`
	// example:
	//
	// 10
	PositionLevel *string `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// btrip123
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
}

func (s GroupUserSaveRequestSubCorpIdList) String() string {
	return dara.Prettify(s)
}

func (s GroupUserSaveRequestSubCorpIdList) GoString() string {
	return s.String()
}

func (s *GroupUserSaveRequestSubCorpIdList) GetDepartIds() []*string {
	return s.DepartIds
}

func (s *GroupUserSaveRequestSubCorpIdList) GetEmail() *string {
	return s.Email
}

func (s *GroupUserSaveRequestSubCorpIdList) GetLeaveStatus() *int32 {
	return s.LeaveStatus
}

func (s *GroupUserSaveRequestSubCorpIdList) GetManagerUserId() *string {
	return s.ManagerUserId
}

func (s *GroupUserSaveRequestSubCorpIdList) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *GroupUserSaveRequestSubCorpIdList) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *GroupUserSaveRequestSubCorpIdList) SetDepartIds(v []*string) *GroupUserSaveRequestSubCorpIdList {
	s.DepartIds = v
	return s
}

func (s *GroupUserSaveRequestSubCorpIdList) SetEmail(v string) *GroupUserSaveRequestSubCorpIdList {
	s.Email = &v
	return s
}

func (s *GroupUserSaveRequestSubCorpIdList) SetLeaveStatus(v int32) *GroupUserSaveRequestSubCorpIdList {
	s.LeaveStatus = &v
	return s
}

func (s *GroupUserSaveRequestSubCorpIdList) SetManagerUserId(v string) *GroupUserSaveRequestSubCorpIdList {
	s.ManagerUserId = &v
	return s
}

func (s *GroupUserSaveRequestSubCorpIdList) SetPositionLevel(v string) *GroupUserSaveRequestSubCorpIdList {
	s.PositionLevel = &v
	return s
}

func (s *GroupUserSaveRequestSubCorpIdList) SetSubCorpId(v string) *GroupUserSaveRequestSubCorpIdList {
	s.SubCorpId = &v
	return s
}

func (s *GroupUserSaveRequestSubCorpIdList) Validate() error {
	return dara.Validate(s)
}
