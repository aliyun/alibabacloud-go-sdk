// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopsPolicyGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *ModifyDesktopsPolicyGroupRequest
	GetDesktopId() []*string
	SetPolicyGroupId(v string) *ModifyDesktopsPolicyGroupRequest
	GetPolicyGroupId() *string
	SetPolicyGroupIds(v []*string) *ModifyDesktopsPolicyGroupRequest
	GetPolicyGroupIds() []*string
	SetRegionId(v string) *ModifyDesktopsPolicyGroupRequest
	GetRegionId() *string
}

type ModifyDesktopsPolicyGroupRequest struct {
	// The cloud computer IDs. You can specify one or more cloud computers IDs. The value is a JSON array.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-ia2zw38bi6cm7****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the cloud computer policy that you want to associate with cloud computers.
	//
	// >  If the `PolicyGroupIds` parameter is used, ignore the current parameter.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The IDs of the cloud computer policies that you want to associate with cloud computers.
	//
	// >  You can specify up to one cloud computer policy that takes effect globally, and up to four cloud computer policies that apply to specific IP addresses. If you specify more than one cloud computer policy that takes effect globally, only the policy first associate with the cloud computer can take effect.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDesktopsPolicyGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *ModifyDesktopsPolicyGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyDesktopsPolicyGroupRequest) GetPolicyGroupIds() []*string {
	return s.PolicyGroupIds
}

func (s *ModifyDesktopsPolicyGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDesktopsPolicyGroupRequest) SetDesktopId(v []*string) *ModifyDesktopsPolicyGroupRequest {
	s.DesktopId = v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetPolicyGroupIds(v []*string) *ModifyDesktopsPolicyGroupRequest {
	s.PolicyGroupIds = v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetRegionId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) Validate() error {
	return dara.Validate(s)
}
