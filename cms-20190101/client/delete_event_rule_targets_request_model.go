// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRuleTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*string) *DeleteEventRuleTargetsRequest
	GetIds() []*string
	SetRegionId(v string) *DeleteEventRuleTargetsRequest
	GetRegionId() *string
	SetRuleName(v string) *DeleteEventRuleTargetsRequest
	GetRuleName() *string
}

type DeleteEventRuleTargetsRequest struct {
	// The IDs of event-triggered alert rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Ids      []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the event-triggered alert rule.
	//
	// For information about how to obtain the name of an event-triggered alert rule, see [DescribeEventRuleList](https://help.aliyun.com/document_detail/114996.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// testRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteEventRuleTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleTargetsRequest) GetIds() []*string {
	return s.Ids
}

func (s *DeleteEventRuleTargetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEventRuleTargetsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteEventRuleTargetsRequest) SetIds(v []*string) *DeleteEventRuleTargetsRequest {
	s.Ids = v
	return s
}

func (s *DeleteEventRuleTargetsRequest) SetRegionId(v string) *DeleteEventRuleTargetsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEventRuleTargetsRequest) SetRuleName(v string) *DeleteEventRuleTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteEventRuleTargetsRequest) Validate() error {
	return dara.Validate(s)
}
