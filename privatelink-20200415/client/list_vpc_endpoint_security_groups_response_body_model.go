// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcEndpointSecurityGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointSecurityGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcEndpointSecurityGroupsResponseBody
	GetRequestId() *string
	SetSecurityGroups(v []*ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) *ListVpcEndpointSecurityGroupsResponseBody
	GetSecurityGroups() []*ListVpcEndpointSecurityGroupsResponseBodySecurityGroups
	SetTotalCount(v int32) *ListVpcEndpointSecurityGroupsResponseBody
	GetTotalCount() *int32
}

type ListVpcEndpointSecurityGroupsResponseBody struct {
	// The number of entries returned per page.
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
	// The information about the security groups.
	SecurityGroups []*ListVpcEndpointSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcEndpointSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) GetSecurityGroups() []*ListVpcEndpointSecurityGroupsResponseBodySecurityGroups {
	return s.SecurityGroups
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetMaxResults(v int32) *ListVpcEndpointSecurityGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetNextToken(v string) *ListVpcEndpointSecurityGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetRequestId(v string) *ListVpcEndpointSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetSecurityGroups(v []*ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) *ListVpcEndpointSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) SetTotalCount(v int32) *ListVpcEndpointSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBody) Validate() error {
	if s.SecurityGroups != nil {
		for _, item := range s.SecurityGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcEndpointSecurityGroupsResponseBodySecurityGroups struct {
	// The ID of the security group that is associated with the endpoint.
	//
	// example:
	//
	// sg-hp33bw6ynvm2yb0e****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The associate status of the security group, valid values:
	//
	// - Attaching: The security group is being attached.
	//
	// - Attached: The security group is attached.
	//
	// - Detaching: The security group is being detached.
	//
	// example:
	//
	// Attached
	SecurityGroupStatus *string `json:"SecurityGroupStatus,omitempty" xml:"SecurityGroupStatus,omitempty"`
}

func (s ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) GetSecurityGroupStatus() *string {
	return s.SecurityGroupStatus
}

func (s *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) SetSecurityGroupId(v string) *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) SetSecurityGroupStatus(v string) *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroupStatus = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponseBodySecurityGroups) Validate() error {
	return dara.Validate(s)
}
