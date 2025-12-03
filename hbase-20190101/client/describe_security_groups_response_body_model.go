// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityGroupsResponseBody
	GetRequestId() *string
	SetSecurityGroupIds(v *DescribeSecurityGroupsResponseBodySecurityGroupIds) *DescribeSecurityGroupsResponseBody
	GetSecurityGroupIds() *DescribeSecurityGroupsResponseBodySecurityGroupIds
}

type DescribeSecurityGroupsResponseBody struct {
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroupIds *DescribeSecurityGroupsResponseBodySecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
}

func (s DescribeSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityGroupsResponseBody) GetSecurityGroupIds() *DescribeSecurityGroupsResponseBodySecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeSecurityGroupsResponseBody) SetRequestId(v string) *DescribeSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetSecurityGroupIds(v *DescribeSecurityGroupsResponseBodySecurityGroupIds) *DescribeSecurityGroupsResponseBody {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) Validate() error {
	if s.SecurityGroupIds != nil {
		if err := s.SecurityGroupIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityGroupsResponseBodySecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeSecurityGroupsResponseBodySecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupIds) Validate() error {
	return dara.Validate(s)
}
