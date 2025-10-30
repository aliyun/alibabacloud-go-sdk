// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServiceUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcEndpointServiceUsersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointServiceUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcEndpointServiceUsersResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListVpcEndpointServiceUsersResponseBody
	GetTotalCount() *string
	SetUserARNs(v []*ListVpcEndpointServiceUsersResponseBodyUserARNs) *ListVpcEndpointServiceUsersResponseBody
	GetUserARNs() []*ListVpcEndpointServiceUsersResponseBodyUserARNs
	SetUsers(v []*ListVpcEndpointServiceUsersResponseBodyUsers) *ListVpcEndpointServiceUsersResponseBody
	GetUsers() []*ListVpcEndpointServiceUsersResponseBodyUsers
}

type ListVpcEndpointServiceUsersResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The whitelists in the format of Aliyun Resource Name (ARN).
	UserARNs []*ListVpcEndpointServiceUsersResponseBodyUserARNs `json:"UserARNs,omitempty" xml:"UserARNs,omitempty" type:"Repeated"`
	// The Alibaba Cloud accounts in the whitelist of the endpoint service.
	Users []*ListVpcEndpointServiceUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServiceUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServiceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointServiceUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServiceUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointServiceUsersResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListVpcEndpointServiceUsersResponseBody) GetUserARNs() []*ListVpcEndpointServiceUsersResponseBodyUserARNs {
	return s.UserARNs
}

func (s *ListVpcEndpointServiceUsersResponseBody) GetUsers() []*ListVpcEndpointServiceUsersResponseBodyUsers {
	return s.Users
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetMaxResults(v int32) *ListVpcEndpointServiceUsersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetNextToken(v string) *ListVpcEndpointServiceUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetRequestId(v string) *ListVpcEndpointServiceUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetTotalCount(v string) *ListVpcEndpointServiceUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetUserARNs(v []*ListVpcEndpointServiceUsersResponseBodyUserARNs) *ListVpcEndpointServiceUsersResponseBody {
	s.UserARNs = v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) SetUsers(v []*ListVpcEndpointServiceUsersResponseBodyUsers) *ListVpcEndpointServiceUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBody) Validate() error {
	if s.UserARNs != nil {
		for _, item := range s.UserARNs {
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

type ListVpcEndpointServiceUsersResponseBodyUserARNs struct {
	// The whitelist in the format of ARN.
	//
	// example:
	//
	// acs:ram:*::*
	UserARN *string `json:"UserARN,omitempty" xml:"UserARN,omitempty"`
}

func (s ListVpcEndpointServiceUsersResponseBodyUserARNs) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServiceUsersResponseBodyUserARNs) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersResponseBodyUserARNs) GetUserARN() *string {
	return s.UserARN
}

func (s *ListVpcEndpointServiceUsersResponseBodyUserARNs) SetUserARN(v string) *ListVpcEndpointServiceUsersResponseBodyUserARNs {
	s.UserARN = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBodyUserARNs) Validate() error {
	return dara.Validate(s)
}

type ListVpcEndpointServiceUsersResponseBodyUsers struct {
	// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
	//
	// example:
	//
	// 12345678
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListVpcEndpointServiceUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServiceUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersResponseBodyUsers) GetUserId() *int64 {
	return s.UserId
}

func (s *ListVpcEndpointServiceUsersResponseBodyUsers) SetUserId(v int64) *ListVpcEndpointServiceUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponseBodyUsers) Validate() error {
	return dara.Validate(s)
}
