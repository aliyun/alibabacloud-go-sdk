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
	// The IP whitelist template information.
	GlobalSecurityIPGroup []*DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2FF6158E-3394-4A90-B634-79C49184****
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
	if s.GlobalSecurityIPGroup != nil {
		for _, item := range s.GlobalSecurityIPGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup struct {
	// The instance IDs associated with the IP whitelist template.
	DBInstances []*string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas (,).
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
