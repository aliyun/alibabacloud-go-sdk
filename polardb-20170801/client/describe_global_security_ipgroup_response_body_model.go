// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalSecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalSecurityIPGroup(v []*DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) *DescribeGlobalSecurityIPGroupResponseBody
	GetGlobalSecurityIPGroup() []*DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup
	SetRequestId(v string) *DescribeGlobalSecurityIPGroupResponseBody
	GetRequestId() *string
}

type DescribeGlobalSecurityIPGroupResponseBody struct {
	// The details of the global IP whitelist template.
	GlobalSecurityIPGroup []*DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGlobalSecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupResponseBody) GetGlobalSecurityIPGroup() []*DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	return s.GlobalSecurityIPGroup
}

func (s *DescribeGlobalSecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalSecurityIPGroupResponseBody) SetGlobalSecurityIPGroup(v []*DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) *DescribeGlobalSecurityIPGroupResponseBody {
	s.GlobalSecurityIPGroup = v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponseBody) SetRequestId(v string) *DescribeGlobalSecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup struct {
	// The details of the clusters that are associated with the global IP address whitelist template.
	DBInstances []*string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The IP address in the global IP whitelist template.
	//
	// >  Separate multiple IP addresses with commas (,). You can add up to 1,000 IP addresses or CIDR blocks to all IP whitelists.
	//
	// example:
	//
	// 192.168.0.1
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the global IP whitelist template. The name must meet the following requirements:
	//
	// 	- The name can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The name must start with a letter and end with a letter or a digit.
	//
	// 	- The name must be 2 to 120 characters in length.
	//
	// example:
	//
	// test_123
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the global IP whitelist template.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetDBInstances() []*string {
	return s.DBInstances
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGIpList() *string {
	return s.GIpList
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetDBInstances(v []*string) *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.DBInstances = v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGIpList(v string) *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GIpList = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGlobalIgName(v string) *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GlobalIgName = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGlobalSecurityGroupId(v string) *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetRegionId(v string) *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
