// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*DescribeGroupUserResponseBodyGroups) *DescribeGroupUserResponseBody
	GetGroups() []*DescribeGroupUserResponseBodyGroups
	SetNextToken(v string) *DescribeGroupUserResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeGroupUserResponseBody
	GetRequestId() *string
	SetUsers(v []*DescribeGroupUserResponseBodyUsers) *DescribeGroupUserResponseBody
	GetUsers() []*DescribeGroupUserResponseBodyUsers
}

type DescribeGroupUserResponseBody struct {
	// >  This field is deprecated.
	Groups []*DescribeGroupUserResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The token for the next query. If NextToken is empty, all results have been queried.
	//
	// example:
	//
	// AAAAAV3MpHK****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AA8D67CB-345D-5CDA-986E-FFAC7D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The members.
	Users []*DescribeGroupUserResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeGroupUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserResponseBody) GetGroups() []*DescribeGroupUserResponseBodyGroups {
	return s.Groups
}

func (s *DescribeGroupUserResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGroupUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupUserResponseBody) GetUsers() []*DescribeGroupUserResponseBodyUsers {
	return s.Users
}

func (s *DescribeGroupUserResponseBody) SetGroups(v []*DescribeGroupUserResponseBodyGroups) *DescribeGroupUserResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeGroupUserResponseBody) SetNextToken(v string) *DescribeGroupUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeGroupUserResponseBody) SetRequestId(v string) *DescribeGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupUserResponseBody) SetUsers(v []*DescribeGroupUserResponseBodyUsers) *DescribeGroupUserResponseBody {
	s.Users = v
	return s
}

func (s *DescribeGroupUserResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGroupUserResponseBodyGroups struct {
	// The user group ID.
	//
	// example:
	//
	// ug-91mvbosdjsdfh****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the user group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The number of members in the user group.
	//
	// example:
	//
	// 1
	UserCount *string `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s DescribeGroupUserResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupUserResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupUserResponseBodyGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeGroupUserResponseBodyGroups) GetUserCount() *string {
	return s.UserCount
}

func (s *DescribeGroupUserResponseBodyGroups) SetGroupId(v string) *DescribeGroupUserResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupUserResponseBodyGroups) SetGroupName(v string) *DescribeGroupUserResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeGroupUserResponseBodyGroups) SetUserCount(v string) *DescribeGroupUserResponseBodyGroups {
	s.UserCount = &v
	return s
}

func (s *DescribeGroupUserResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupUserResponseBodyUsers struct {
	// >  This field is deprecated.
	//
	// example:
	//
	// xx-xx-xx
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// >  This field is deprecated.
	//
	// example:
	//
	// https://avatar.****.com
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// The email address.
	//
	// example:
	//
	// alex****@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The user name.
	//
	// example:
	//
	// alex****
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The time when the user was created.
	//
	// example:
	//
	// 2024-08-26T02:59:22.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the user was added to the user group.
	//
	// example:
	//
	// 2024-08-26T02:59:22.000+00:00
	GmtJoinGroup *string `json:"GmtJoinGroup,omitempty" xml:"GmtJoinGroup,omitempty"`
	// >  This field is deprecated.
	//
	// example:
	//
	// 123
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
	// The display name.
	//
	// example:
	//
	// alex
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The mobile number.
	//
	// example:
	//
	// 188888****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The remarks on the user.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeGroupUserResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupUserResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserResponseBodyUsers) GetAddress() *string {
	return s.Address
}

func (s *DescribeGroupUserResponseBodyUsers) GetAvatar() *string {
	return s.Avatar
}

func (s *DescribeGroupUserResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *DescribeGroupUserResponseBodyUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeGroupUserResponseBodyUsers) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeGroupUserResponseBodyUsers) GetGmtJoinGroup() *string {
	return s.GmtJoinGroup
}

func (s *DescribeGroupUserResponseBodyUsers) GetJobNumber() *string {
	return s.JobNumber
}

func (s *DescribeGroupUserResponseBodyUsers) GetNickName() *string {
	return s.NickName
}

func (s *DescribeGroupUserResponseBodyUsers) GetPhone() *string {
	return s.Phone
}

func (s *DescribeGroupUserResponseBodyUsers) GetRemark() *string {
	return s.Remark
}

func (s *DescribeGroupUserResponseBodyUsers) SetAddress(v string) *DescribeGroupUserResponseBodyUsers {
	s.Address = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetAvatar(v string) *DescribeGroupUserResponseBodyUsers {
	s.Avatar = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetEmail(v string) *DescribeGroupUserResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetEndUserId(v string) *DescribeGroupUserResponseBodyUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetGmtCreated(v string) *DescribeGroupUserResponseBodyUsers {
	s.GmtCreated = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetGmtJoinGroup(v string) *DescribeGroupUserResponseBodyUsers {
	s.GmtJoinGroup = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetJobNumber(v string) *DescribeGroupUserResponseBodyUsers {
	s.JobNumber = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetNickName(v string) *DescribeGroupUserResponseBodyUsers {
	s.NickName = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetPhone(v string) *DescribeGroupUserResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetRemark(v string) *DescribeGroupUserResponseBodyUsers {
	s.Remark = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
