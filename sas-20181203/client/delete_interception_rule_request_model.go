// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInterceptionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteInterceptionRuleRequest
	GetClusterId() *string
	SetRuleIds(v []*int64) *DeleteInterceptionRuleRequest
	GetRuleIds() []*int64
}

type DeleteInterceptionRuleRequest struct {
	// The ID of the cluster that you want to query.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// cdf629147cc3747d292a3f587xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The IDs of the rules that you want to delete.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DeleteInterceptionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterceptionRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteInterceptionRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteInterceptionRuleRequest) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *DeleteInterceptionRuleRequest) SetClusterId(v string) *DeleteInterceptionRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteInterceptionRuleRequest) SetRuleIds(v []*int64) *DeleteInterceptionRuleRequest {
	s.RuleIds = v
	return s
}

func (s *DeleteInterceptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
