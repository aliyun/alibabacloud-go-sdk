// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteArtifactSubscriptionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteArtifactSubscriptionRuleRequest
	GetInstanceId() *string
	SetRuleId(v string) *DeleteArtifactSubscriptionRuleRequest
	GetRuleId() *string
}

type DeleteArtifactSubscriptionRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-c0o11woew0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteArtifactSubscriptionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteArtifactSubscriptionRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteArtifactSubscriptionRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteArtifactSubscriptionRuleRequest) SetInstanceId(v string) *DeleteArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleRequest) SetRuleId(v string) *DeleteArtifactSubscriptionRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
