// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalSecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalSecurityIPGroup(v []*DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) *DeleteGlobalSecurityIPGroupResponseBody
	GetGlobalSecurityIPGroup() []*DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup
	SetRequestId(v string) *DeleteGlobalSecurityIPGroupResponseBody
	GetRequestId() *string
}

type DeleteGlobalSecurityIPGroupResponseBody struct {
	// The details of the global IP whitelist template.
	GlobalSecurityIPGroup []*DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGlobalSecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalSecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalSecurityIPGroupResponseBody) GetGlobalSecurityIPGroup() []*DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	return s.GlobalSecurityIPGroup
}

func (s *DeleteGlobalSecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGlobalSecurityIPGroupResponseBody) SetGlobalSecurityIPGroup(v []*DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) *DeleteGlobalSecurityIPGroupResponseBody {
	s.GlobalSecurityIPGroup = v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponseBody) SetRequestId(v string) *DeleteGlobalSecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup struct {
	// The clusters that are associated with the IP address whitelist template.
	DBInstances []*string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The IP address in the whitelist template.
	//
	// >  Multiple IP addresses are separated by commas (,). You can create up to 1,000 IP addresses or CIDR blocks for all IP whitelists.
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

func (s DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetDBInstances() []*string {
	return s.DBInstances
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGIpList() *string {
	return s.GIpList
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetDBInstances(v []*string) *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.DBInstances = v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGIpList(v string) *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GIpList = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGlobalIgName(v string) *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GlobalIgName = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGlobalSecurityGroupId(v string) *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetRegionId(v string) *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.RegionId = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
