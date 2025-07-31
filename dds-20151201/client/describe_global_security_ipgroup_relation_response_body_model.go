// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalSecurityIPGroupRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBody
	GetDBClusterId() *string
	SetGlobalSecurityIPGroupRel(v []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) *DescribeGlobalSecurityIPGroupRelationResponseBody
	GetGlobalSecurityIPGroupRel() []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel
	SetRequestId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBody
	GetRequestId() *string
}

type DescribeGlobalSecurityIPGroupRelationResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// dds-2ze6069764423m0l
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The global IP whitelist templates associated with the instance.
	GlobalSecurityIPGroupRel []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel `json:"GlobalSecurityIPGroupRel,omitempty" xml:"GlobalSecurityIPGroupRel,omitempty" type:"Repeated"`
	// The unique ID of the request. If the request fails, provide this ID for technical support to troubleshoot the failure.
	//
	// example:
	//
	// F8CA8312-530A-413A-9129-F2BB32A8D404
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGlobalSecurityIPGroupRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) GetGlobalSecurityIPGroupRel() []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	return s.GlobalSecurityIPGroupRel
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) SetDBClusterId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) SetGlobalSecurityIPGroupRel(v []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) *DescribeGlobalSecurityIPGroupRelationResponseBody {
	s.GlobalSecurityIPGroupRel = v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) SetRequestId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel struct {
	// The IP addresses in the whitelist template.
	//
	// >  Separate multiple IP addresses with commas (,). You can create up to 1,000 IP addresses or CIDR blocks for all IP whitelists.
	//
	// example:
	//
	// 27.16.214.10,111.60.117.181
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
	// g-gfurfpsh4ycbrm2avst7
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetGIpList() *string {
	return s.GIpList
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetGIpList(v string) *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.GIpList = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetGlobalIgName(v string) *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.GlobalIgName = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetGlobalSecurityGroupId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetRegionId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) Validate() error {
	return dara.Validate(s)
}
