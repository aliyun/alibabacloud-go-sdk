// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListOrganizationMembersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListOrganizationMembersResponseBody
	GetErrorMessage() *string
	SetMembers(v []*ListOrganizationMembersResponseBodyMembers) *ListOrganizationMembersResponseBody
	GetMembers() []*ListOrganizationMembersResponseBodyMembers
	SetNextToken(v string) *ListOrganizationMembersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListOrganizationMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOrganizationMembersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListOrganizationMembersResponseBody
	GetTotalCount() *int64
}

type ListOrganizationMembersResponseBody struct {
	// example:
	//
	// null
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error info
	ErrorMessage *string                                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Members      []*ListOrganizationMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// FC93CE1A-8D7A-13A9-8306-7465DE2E5C0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListOrganizationMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListOrganizationMembersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListOrganizationMembersResponseBody) GetMembers() []*ListOrganizationMembersResponseBodyMembers {
	return s.Members
}

func (s *ListOrganizationMembersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOrganizationMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrganizationMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOrganizationMembersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOrganizationMembersResponseBody) SetErrorCode(v string) *ListOrganizationMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetErrorMessage(v string) *ListOrganizationMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetMembers(v []*ListOrganizationMembersResponseBodyMembers) *ListOrganizationMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetNextToken(v string) *ListOrganizationMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetRequestId(v string) *ListOrganizationMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetSuccess(v bool) *ListOrganizationMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) SetTotalCount(v int64) *ListOrganizationMembersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrganizationMembersResponseBody) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOrganizationMembersResponseBodyMembers struct {
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
	HiredDate  *int64                                                `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	Identities *ListOrganizationMembersResponseBodyMembersIdentities `json:"identities,omitempty" xml:"identities,omitempty" type:"Struct"`
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

func (s ListOrganizationMembersResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBodyMembers) GetAccountId() *string {
	return s.AccountId
}

func (s *ListOrganizationMembersResponseBodyMembers) GetBirthday() *int64 {
	return s.Birthday
}

func (s *ListOrganizationMembersResponseBodyMembers) GetDeptLists() []*string {
	return s.DeptLists
}

func (s *ListOrganizationMembersResponseBodyMembers) GetEmail() *string {
	return s.Email
}

func (s *ListOrganizationMembersResponseBodyMembers) GetHiredDate() *int64 {
	return s.HiredDate
}

func (s *ListOrganizationMembersResponseBodyMembers) GetIdentities() *ListOrganizationMembersResponseBodyMembersIdentities {
	return s.Identities
}

func (s *ListOrganizationMembersResponseBodyMembers) GetJobNumber() *string {
	return s.JobNumber
}

func (s *ListOrganizationMembersResponseBodyMembers) GetJoinTime() *int64 {
	return s.JoinTime
}

func (s *ListOrganizationMembersResponseBodyMembers) GetLastVisitTime() *int64 {
	return s.LastVisitTime
}

func (s *ListOrganizationMembersResponseBodyMembers) GetMobile() *string {
	return s.Mobile
}

func (s *ListOrganizationMembersResponseBodyMembers) GetOrganizationMemberName() *string {
	return s.OrganizationMemberName
}

func (s *ListOrganizationMembersResponseBodyMembers) GetOrganizationRoleId() *string {
	return s.OrganizationRoleId
}

func (s *ListOrganizationMembersResponseBodyMembers) GetOrganizationRoleName() *string {
	return s.OrganizationRoleName
}

func (s *ListOrganizationMembersResponseBodyMembers) GetState() *string {
	return s.State
}

func (s *ListOrganizationMembersResponseBodyMembers) SetAccountId(v string) *ListOrganizationMembersResponseBodyMembers {
	s.AccountId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetBirthday(v int64) *ListOrganizationMembersResponseBodyMembers {
	s.Birthday = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetDeptLists(v []*string) *ListOrganizationMembersResponseBodyMembers {
	s.DeptLists = v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetEmail(v string) *ListOrganizationMembersResponseBodyMembers {
	s.Email = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetHiredDate(v int64) *ListOrganizationMembersResponseBodyMembers {
	s.HiredDate = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetIdentities(v *ListOrganizationMembersResponseBodyMembersIdentities) *ListOrganizationMembersResponseBodyMembers {
	s.Identities = v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetJobNumber(v string) *ListOrganizationMembersResponseBodyMembers {
	s.JobNumber = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetJoinTime(v int64) *ListOrganizationMembersResponseBodyMembers {
	s.JoinTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetLastVisitTime(v int64) *ListOrganizationMembersResponseBodyMembers {
	s.LastVisitTime = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetMobile(v string) *ListOrganizationMembersResponseBodyMembers {
	s.Mobile = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetOrganizationMemberName(v string) *ListOrganizationMembersResponseBodyMembers {
	s.OrganizationMemberName = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetOrganizationRoleId(v string) *ListOrganizationMembersResponseBodyMembers {
	s.OrganizationRoleId = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetOrganizationRoleName(v string) *ListOrganizationMembersResponseBodyMembers {
	s.OrganizationRoleName = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) SetState(v string) *ListOrganizationMembersResponseBodyMembers {
	s.State = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembers) Validate() error {
	if s.Identities != nil {
		if err := s.Identities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOrganizationMembersResponseBodyMembersIdentities struct {
	// example:
	//
	// 1236666
	ExternUid *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	// example:
	//
	// Dingtalk
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ListOrganizationMembersResponseBodyMembersIdentities) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersResponseBodyMembersIdentities) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersResponseBodyMembersIdentities) GetExternUid() *string {
	return s.ExternUid
}

func (s *ListOrganizationMembersResponseBodyMembersIdentities) GetProvider() *string {
	return s.Provider
}

func (s *ListOrganizationMembersResponseBodyMembersIdentities) SetExternUid(v string) *ListOrganizationMembersResponseBodyMembersIdentities {
	s.ExternUid = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembersIdentities) SetProvider(v string) *ListOrganizationMembersResponseBodyMembersIdentities {
	s.Provider = &v
	return s
}

func (s *ListOrganizationMembersResponseBodyMembersIdentities) Validate() error {
	return dara.Validate(s)
}
