// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTimeTriggerScalingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *DescribeTimeTriggerScalingRulesRequest
	GetNodeGroupId() *string
}

type DescribeTimeTriggerScalingRulesRequest struct {
	// example:
	//
	// ng-d332aa8bca48****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s DescribeTimeTriggerScalingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimeTriggerScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTimeTriggerScalingRulesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeTimeTriggerScalingRulesRequest) SetNodeGroupId(v string) *DescribeTimeTriggerScalingRulesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeTimeTriggerScalingRulesRequest) Validate() error {
	return dara.Validate(s)
}
