// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayClientUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSmartAccessGatewayClientUsersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSmartAccessGatewayClientUsersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSmartAccessGatewayClientUsersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSmartAccessGatewayClientUsersResponseBody
	GetTotalCount() *int32
	SetUsers(v *DescribeSmartAccessGatewayClientUsersResponseBodyUsers) *DescribeSmartAccessGatewayClientUsersResponseBody
	GetUsers() *DescribeSmartAccessGatewayClientUsersResponseBodyUsers
}

type DescribeSmartAccessGatewayClientUsersResponseBody struct {
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 62F4CF10-F909-487E-8E95-BC35457C5F50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users      *DescribeSmartAccessGatewayClientUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s DescribeSmartAccessGatewayClientUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayClientUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) GetUsers() *DescribeSmartAccessGatewayClientUsersResponseBodyUsers {
	return s.Users
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) SetPageNumber(v int32) *DescribeSmartAccessGatewayClientUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) SetPageSize(v int32) *DescribeSmartAccessGatewayClientUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) SetRequestId(v string) *DescribeSmartAccessGatewayClientUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) SetTotalCount(v int32) *DescribeSmartAccessGatewayClientUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) SetUsers(v *DescribeSmartAccessGatewayClientUsersResponseBodyUsers) *DescribeSmartAccessGatewayClientUsersResponseBody {
	s.Users = v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBody) Validate() error {
	if s.Users != nil {
		if err := s.Users.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSmartAccessGatewayClientUsersResponseBodyUsers struct {
	User []*DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewayClientUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayClientUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsers) GetUser() []*DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser {
	return s.User
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsers) SetUser(v []*DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) *DescribeSmartAccessGatewayClientUsersResponseBodyUsers {
	s.User = v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsers) Validate() error {
	if s.User != nil {
		for _, item := range s.User {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser struct {
	AccelerateBandwidth *int64  `json:"AccelerateBandwidth,omitempty" xml:"AccelerateBandwidth,omitempty"`
	Bandwidth           *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientIp            *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	IsStaticIp          *int64  `json:"IsStaticIp,omitempty" xml:"IsStaticIp,omitempty"`
	State               *int32  `json:"State,omitempty" xml:"State,omitempty"`
	UserMail            *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	UserName            *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) GetAccelerateBandwidth() *int64 {
	return s.AccelerateBandwidth
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) GetIsStaticIp() *int64 {
	return s.IsStaticIp
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) GetState() *int32 {
	return s.State
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) GetUserMail() *string {
	return s.UserMail
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) SetAccelerateBandwidth(v int64) *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser {
	s.AccelerateBandwidth = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) SetBandwidth(v int32) *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser {
	s.Bandwidth = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) SetClientIp(v string) *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser {
	s.ClientIp = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) SetIsStaticIp(v int64) *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser {
	s.IsStaticIp = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) SetState(v int32) *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser {
	s.State = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) SetUserMail(v string) *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser {
	s.UserMail = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) SetUserName(v string) *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser {
	s.UserName = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponseBodyUsersUser) Validate() error {
	return dara.Validate(s)
}
