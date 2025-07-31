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
	// The IP whitelist templates.
	GlobalSecurityIPGroup []*CreateGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F8CA8312-530A-413A-9129-F2BB32A8D404
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
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas (,). You can create up to 1,000 IP addresses or CIDR blocks for all IP address whitelists.
	//
	// example:
	//
	// 111.207.194.217,61.149.7.154
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP whitelist template. The name must meet the following requirements:
	//
	// 	- It can contain lowercase letters, digits, and underscores (_).
	//
	// 	- It must start with a letter and end with a letter or digit.
	//
	// 	- It must be 2 to 120 characters in length.
	//
	// example:
	//
	// developer
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP whitelist template.
	//
	// example:
	//
	// g-2ifynl1jfkdlhhxgva3q
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID of the instance.
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
