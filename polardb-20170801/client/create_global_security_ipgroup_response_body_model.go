// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalSecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalSecurityIPGroup(v []*CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) *CreateGlobalSecurityIPGroupResponseBody
	GetGlobalSecurityIPGroup() []*CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup
	SetRequestId(v string) *CreateGlobalSecurityIPGroupResponseBody
	GetRequestId() *string
}

type CreateGlobalSecurityIPGroupResponseBody struct {
	// The details of the global IP whitelist template.
	GlobalSecurityIPGroup []*CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGlobalSecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalSecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalSecurityIPGroupResponseBody) GetGlobalSecurityIPGroup() []*CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	return s.GlobalSecurityIPGroup
}

func (s *CreateGlobalSecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGlobalSecurityIPGroupResponseBody) SetGlobalSecurityIPGroup(v []*CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) *CreateGlobalSecurityIPGroupResponseBody {
	s.GlobalSecurityIPGroup = v
	return s
}

func (s *CreateGlobalSecurityIPGroupResponseBody) SetRequestId(v string) *CreateGlobalSecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup struct {
	// The IP address in the whitelist template.
	//
	// >  Separate multiple IP addresses with commas (,). You can add up to 1,000 IP addresses or CIDR blocks to all IP whitelists.
	//
	// example:
	//
	// 192.168.0.1
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP whitelist template. The name must meet the following requirements:
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
	// The ID of the IP whitelist template.
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

func (s CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGIpList() *string {
	return s.GIpList
}

func (s *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGIpList(v string) *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GIpList = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGlobalIgName(v string) *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GlobalIgName = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGlobalSecurityGroupId(v string) *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetRegionId(v string) *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.RegionId = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
