// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteArtifactLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteArtifactLifecycleRuleRequest
	GetInstanceId() *string
	SetRuleId(v string) *DeleteArtifactLifecycleRuleRequest
	GetRuleId() *string
}

type DeleteArtifactLifecycleRuleRequest struct {
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-brlg4cbj2ylkrqqq
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cralr-3v8pao9k7chb8q62
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteArtifactLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteArtifactLifecycleRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteArtifactLifecycleRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteArtifactLifecycleRuleRequest) SetInstanceId(v string) *DeleteArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleRequest) SetRuleId(v string) *DeleteArtifactLifecycleRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
