// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInnerIpWhitelistGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInnerIpWhitelistGroupsRequest
	GetInstanceId() *string
}

type DescribeInnerIpWhitelistGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInnerIpWhitelistGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInnerIpWhitelistGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInnerIpWhitelistGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInnerIpWhitelistGroupsRequest) SetInstanceId(v string) *DescribeInnerIpWhitelistGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInnerIpWhitelistGroupsRequest) Validate() error {
	return dara.Validate(s)
}
