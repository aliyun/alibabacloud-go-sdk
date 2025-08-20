// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactSubscriptionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateArtifactSubscriptionTaskRequest
	GetInstanceId() *string
	SetRuleId(v string) *CreateArtifactSubscriptionTaskRequest
	GetRuleId() *string
}

type CreateArtifactSubscriptionTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4ec5xvj4j0l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crasr-88s7vmelc3m****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateArtifactSubscriptionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactSubscriptionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateArtifactSubscriptionTaskRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateArtifactSubscriptionTaskRequest) SetInstanceId(v string) *CreateArtifactSubscriptionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskRequest) SetRuleId(v string) *CreateArtifactSubscriptionTaskRequest {
	s.RuleId = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskRequest) Validate() error {
	return dara.Validate(s)
}
