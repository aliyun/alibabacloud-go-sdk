// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopUsers(v []*DescribeUsersPasswordResponseBodyDesktopUsers) *DescribeUsersPasswordResponseBody
	GetDesktopUsers() []*DescribeUsersPasswordResponseBodyDesktopUsers
	SetRequestId(v string) *DescribeUsersPasswordResponseBody
	GetRequestId() *string
}

type DescribeUsersPasswordResponseBody struct {
	// The authorized users of the cloud computer.
	DesktopUsers []*DescribeUsersPasswordResponseBodyDesktopUsers `json:"DesktopUsers,omitempty" xml:"DesktopUsers,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F7B4B17B-5C8A-514C-AA4D-F8090E3A63E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUsersPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersPasswordResponseBody) GetDesktopUsers() []*DescribeUsersPasswordResponseBodyDesktopUsers {
	return s.DesktopUsers
}

func (s *DescribeUsersPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUsersPasswordResponseBody) SetDesktopUsers(v []*DescribeUsersPasswordResponseBodyDesktopUsers) *DescribeUsersPasswordResponseBody {
	s.DesktopUsers = v
	return s
}

func (s *DescribeUsersPasswordResponseBody) SetRequestId(v string) *DescribeUsersPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersPasswordResponseBody) Validate() error {
	if s.DesktopUsers != nil {
		for _, item := range s.DesktopUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUsersPasswordResponseBodyDesktopUsers struct {
	// The display name of the end user.
	//
	// example:
	//
	// alice_1365*****
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The password of the end user.
	//
	// example:
	//
	// tes123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s DescribeUsersPasswordResponseBodyDesktopUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersPasswordResponseBodyDesktopUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersPasswordResponseBodyDesktopUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeUsersPasswordResponseBodyDesktopUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeUsersPasswordResponseBodyDesktopUsers) GetPassword() *string {
	return s.Password
}

func (s *DescribeUsersPasswordResponseBodyDesktopUsers) SetDisplayName(v string) *DescribeUsersPasswordResponseBodyDesktopUsers {
	s.DisplayName = &v
	return s
}

func (s *DescribeUsersPasswordResponseBodyDesktopUsers) SetEndUserId(v string) *DescribeUsersPasswordResponseBodyDesktopUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersPasswordResponseBodyDesktopUsers) SetPassword(v string) *DescribeUsersPasswordResponseBodyDesktopUsers {
	s.Password = &v
	return s
}

func (s *DescribeUsersPasswordResponseBodyDesktopUsers) Validate() error {
	return dara.Validate(s)
}
