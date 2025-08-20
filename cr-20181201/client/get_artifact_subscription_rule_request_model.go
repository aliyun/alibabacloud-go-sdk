// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactSubscriptionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetArtifactSubscriptionRuleRequest
	GetInstanceId() *string
	SetRuleId(v string) *GetArtifactSubscriptionRuleRequest
	GetRuleId() *string
}

type GetArtifactSubscriptionRuleRequest struct {
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

func (s GetArtifactSubscriptionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactSubscriptionRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *GetArtifactSubscriptionRuleRequest) SetInstanceId(v string) *GetArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionRuleRequest) SetRuleId(v string) *GetArtifactSubscriptionRuleRequest {
	s.RuleId = &v
	return s
}

func (s *GetArtifactSubscriptionRuleRequest) Validate() error {
	return dara.Validate(s)
}
