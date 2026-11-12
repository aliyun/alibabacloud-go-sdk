// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFirewallRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallRuleId(v string) *ListFirewallRulesRequest
	GetFirewallRuleId() *string
	SetInstanceId(v string) *ListFirewallRulesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListFirewallRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFirewallRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListFirewallRulesRequest
	GetRegionId() *string
	SetTag(v []*ListFirewallRulesRequestTag) *ListFirewallRulesRequest
	GetTag() []*ListFirewallRulesRequestTag
}

type ListFirewallRulesRequest struct {
	// The ID of the firewall rule.
	//
	// example:
	//
	// 1a16263ab0f541288312a15fa64280de
	FirewallRuleId *string `json:"FirewallRuleId,omitempty" xml:"FirewallRuleId,omitempty"`
	// The instance ID of the specified simple application server.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the firewall rule list.
	//
	// Starting value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The paging Settings for the number of entries per page in a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags of the firewall rules.
	Tag []*ListFirewallRulesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListFirewallRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesRequest) GetFirewallRuleId() *string {
	return s.FirewallRuleId
}

func (s *ListFirewallRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFirewallRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFirewallRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFirewallRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListFirewallRulesRequest) GetTag() []*ListFirewallRulesRequestTag {
	return s.Tag
}

func (s *ListFirewallRulesRequest) SetFirewallRuleId(v string) *ListFirewallRulesRequest {
	s.FirewallRuleId = &v
	return s
}

func (s *ListFirewallRulesRequest) SetInstanceId(v string) *ListFirewallRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFirewallRulesRequest) SetPageNumber(v int32) *ListFirewallRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFirewallRulesRequest) SetPageSize(v int32) *ListFirewallRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFirewallRulesRequest) SetRegionId(v string) *ListFirewallRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListFirewallRulesRequest) SetTag(v []*ListFirewallRulesRequestTag) *ListFirewallRulesRequest {
	s.Tag = v
	return s
}

func (s *ListFirewallRulesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFirewallRulesRequestTag struct {
	// The tag key of the firewall rule. The tag key must be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the firewall rule. The tag value must be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFirewallRulesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListFirewallRulesRequestTag) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListFirewallRulesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListFirewallRulesRequestTag) SetKey(v string) *ListFirewallRulesRequestTag {
	s.Key = &v
	return s
}

func (s *ListFirewallRulesRequestTag) SetValue(v string) *ListFirewallRulesRequestTag {
	s.Value = &v
	return s
}

func (s *ListFirewallRulesRequestTag) Validate() error {
	return dara.Validate(s)
}
