// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalSecurityIPGroup(v []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) *ModifyGlobalSecurityIPGroupNameResponseBody
	GetGlobalSecurityIPGroup() []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup
	SetRequestId(v string) *ModifyGlobalSecurityIPGroupNameResponseBody
	GetRequestId() *string
}

type ModifyGlobalSecurityIPGroupNameResponseBody struct {
	// The global IP whitelist templates.
	GlobalSecurityIPGroup []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F8CA8312-530A-413A-9129-F2BB32A8D404
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) GetGlobalSecurityIPGroup() []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	return s.GlobalSecurityIPGroup
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) SetGlobalSecurityIPGroup(v []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) *ModifyGlobalSecurityIPGroupNameResponseBody {
	s.GlobalSecurityIPGroup = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) SetRequestId(v string) *ModifyGlobalSecurityIPGroupNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup struct {
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas (,). You can create up to 1,000 IP addresses or CIDR blocks for all IP address whitelists.
	//
	// example:
	//
	// 222.70.197.187
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP whitelist template.
	//
	// example:
	//
	// def
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP whitelist template.
	//
	// example:
	//
	// g-qiawi8ec1urcx9swoy37
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GetGIpList() *string {
	return s.GIpList
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) SetGIpList(v string) *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	s.GIpList = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) SetGlobalIgName(v string) *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	s.GlobalIgName = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) SetRegionId(v string) *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
