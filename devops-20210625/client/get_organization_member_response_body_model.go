// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetOrganizationMemberResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetOrganizationMemberResponseBody
	GetErrorMessage() *string
	SetMember(v *GetOrganizationMemberResponseBodyMember) *GetOrganizationMemberResponseBody
	GetMember() *GetOrganizationMemberResponseBodyMember
	SetRequestId(v string) *GetOrganizationMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOrganizationMemberResponseBody
	GetSuccess() *bool
}

type GetOrganizationMemberResponseBody struct {
	// example:
	//
	// null
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error info
	ErrorMessage *string                                  `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Member       *GetOrganizationMemberResponseBodyMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	// example:
	//
	// HC93CE1A-8D7A-13A9-8306-7465DE2E5C0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetOrganizationMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOrganizationMemberResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOrganizationMemberResponseBody) GetMember() *GetOrganizationMemberResponseBodyMember {
	return s.Member
}

func (s *GetOrganizationMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrganizationMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOrganizationMemberResponseBody) SetErrorCode(v string) *GetOrganizationMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOrganizationMemberResponseBody) SetErrorMessage(v string) *GetOrganizationMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOrganizationMemberResponseBody) SetMember(v *GetOrganizationMemberResponseBodyMember) *GetOrganizationMemberResponseBody {
	s.Member = v
	return s
}

func (s *GetOrganizationMemberResponseBody) SetRequestId(v string) *GetOrganizationMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrganizationMemberResponseBody) SetSuccess(v bool) *GetOrganizationMemberResponseBody {
	s.Success = &v
	return s
}

func (s *GetOrganizationMemberResponseBody) Validate() error {
	if s.Member != nil {
		if err := s.Member.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrganizationMemberResponseBodyMember struct {
	// example:
	//
	// 123456677888
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 1631845101798
	Birthday  *int64    `json:"birthday,omitempty" xml:"birthday,omitempty"`
	DeptLists []*string `json:"deptLists,omitempty" xml:"deptLists,omitempty" type:"Repeated"`
	// example:
	//
	// (敏感字段，暂不返回)
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 1631845101798
	HiredDate  *int64                                             `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	Identities *GetOrganizationMemberResponseBodyMemberIdentities `json:"identities,omitempty" xml:"identities,omitempty" type:"Struct"`
	// example:
	//
	// 373***
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// example:
	//
	// 1631845101798
	JoinTime *int64 `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	// example:
	//
	// 1631845101798
	LastVisitTime *int64 `json:"lastVisitTime,omitempty" xml:"lastVisitTime,omitempty"`
	// example:
	//
	// (敏感字段，暂不返回)
	Mobile                 *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	OrganizationMemberName *string `json:"organizationMemberName,omitempty" xml:"organizationMemberName,omitempty"`
	// example:
	//
	// 8fc0c9ff039711dd64acd000
	OrganizationRoleId   *string `json:"organizationRoleId,omitempty" xml:"organizationRoleId,omitempty"`
	OrganizationRoleName *string `json:"organizationRoleName,omitempty" xml:"organizationRoleName,omitempty"`
	// example:
	//
	// normal
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s GetOrganizationMemberResponseBodyMember) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationMemberResponseBodyMember) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberResponseBodyMember) GetAccountId() *string {
	return s.AccountId
}

func (s *GetOrganizationMemberResponseBodyMember) GetBirthday() *int64 {
	return s.Birthday
}

func (s *GetOrganizationMemberResponseBodyMember) GetDeptLists() []*string {
	return s.DeptLists
}

func (s *GetOrganizationMemberResponseBodyMember) GetEmail() *string {
	return s.Email
}

func (s *GetOrganizationMemberResponseBodyMember) GetHiredDate() *int64 {
	return s.HiredDate
}

func (s *GetOrganizationMemberResponseBodyMember) GetIdentities() *GetOrganizationMemberResponseBodyMemberIdentities {
	return s.Identities
}

func (s *GetOrganizationMemberResponseBodyMember) GetJobNumber() *string {
	return s.JobNumber
}

func (s *GetOrganizationMemberResponseBodyMember) GetJoinTime() *int64 {
	return s.JoinTime
}

func (s *GetOrganizationMemberResponseBodyMember) GetLastVisitTime() *int64 {
	return s.LastVisitTime
}

func (s *GetOrganizationMemberResponseBodyMember) GetMobile() *string {
	return s.Mobile
}

func (s *GetOrganizationMemberResponseBodyMember) GetOrganizationMemberName() *string {
	return s.OrganizationMemberName
}

func (s *GetOrganizationMemberResponseBodyMember) GetOrganizationRoleId() *string {
	return s.OrganizationRoleId
}

func (s *GetOrganizationMemberResponseBodyMember) GetOrganizationRoleName() *string {
	return s.OrganizationRoleName
}

func (s *GetOrganizationMemberResponseBodyMember) GetState() *string {
	return s.State
}

func (s *GetOrganizationMemberResponseBodyMember) SetAccountId(v string) *GetOrganizationMemberResponseBodyMember {
	s.AccountId = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetBirthday(v int64) *GetOrganizationMemberResponseBodyMember {
	s.Birthday = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetDeptLists(v []*string) *GetOrganizationMemberResponseBodyMember {
	s.DeptLists = v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetEmail(v string) *GetOrganizationMemberResponseBodyMember {
	s.Email = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetHiredDate(v int64) *GetOrganizationMemberResponseBodyMember {
	s.HiredDate = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetIdentities(v *GetOrganizationMemberResponseBodyMemberIdentities) *GetOrganizationMemberResponseBodyMember {
	s.Identities = v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetJobNumber(v string) *GetOrganizationMemberResponseBodyMember {
	s.JobNumber = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetJoinTime(v int64) *GetOrganizationMemberResponseBodyMember {
	s.JoinTime = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetLastVisitTime(v int64) *GetOrganizationMemberResponseBodyMember {
	s.LastVisitTime = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetMobile(v string) *GetOrganizationMemberResponseBodyMember {
	s.Mobile = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetOrganizationMemberName(v string) *GetOrganizationMemberResponseBodyMember {
	s.OrganizationMemberName = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetOrganizationRoleId(v string) *GetOrganizationMemberResponseBodyMember {
	s.OrganizationRoleId = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetOrganizationRoleName(v string) *GetOrganizationMemberResponseBodyMember {
	s.OrganizationRoleName = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) SetState(v string) *GetOrganizationMemberResponseBodyMember {
	s.State = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMember) Validate() error {
	if s.Identities != nil {
		if err := s.Identities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrganizationMemberResponseBodyMemberIdentities struct {
	// example:
	//
	// 1236666
	ExternUid *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	// example:
	//
	// Dingtalk
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s GetOrganizationMemberResponseBodyMemberIdentities) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationMemberResponseBodyMemberIdentities) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberResponseBodyMemberIdentities) GetExternUid() *string {
	return s.ExternUid
}

func (s *GetOrganizationMemberResponseBodyMemberIdentities) GetProvider() *string {
	return s.Provider
}

func (s *GetOrganizationMemberResponseBodyMemberIdentities) SetExternUid(v string) *GetOrganizationMemberResponseBodyMemberIdentities {
	s.ExternUid = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMemberIdentities) SetProvider(v string) *GetOrganizationMemberResponseBodyMemberIdentities {
	s.Provider = &v
	return s
}

func (s *GetOrganizationMemberResponseBodyMemberIdentities) Validate() error {
	return dara.Validate(s)
}
