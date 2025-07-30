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
	// r-t4n885e834f6****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information about the associated global IP whitelist template.
	GlobalSecurityIPGroupRel []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel `json:"GlobalSecurityIPGroupRel,omitempty" xml:"GlobalSecurityIPGroupRel,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
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
	// The IP address in the IP whitelist template.
	//
	// >  Multiple IP addresses are separated by commas (,).
	//
	// example:
	//
	// 192.168.0.1,10.10.10.10
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP whitelist template.
	//
	// example:
	//
	// test_123
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP whitelist template.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
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
