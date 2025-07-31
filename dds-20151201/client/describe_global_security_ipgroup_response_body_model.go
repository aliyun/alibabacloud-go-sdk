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
	// The global IP whitelist templates.
	GlobalSecurityIPGroup []*DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
	// The unique ID of the request. If the request fails, provide this ID for technical support to troubleshoot the failure.
	//
	// example:
	//
	// 72651AF9-7897-41A7-8B67-6C15C7F26A0A
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
	// The instances associated with the global whitelist template.
	DBInstances []*string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas (,). You can create up to 1,000 IP addresses or CIDR blocks for all IP address whitelists.
	//
	// example:
	//
	// 117.12.202.19
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP whitelist template.
	//
	// example:
	//
	// dev_baoxian_k8s_bj
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP whitelist template.
	//
	// example:
	//
	// g-sdgwqyp4f5j1x3qk7yvm
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID of the instance.
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
