// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactLifecycleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetArtifactLifecycleRuleRequest
	GetInstanceId() *string
	SetRuleId(v string) *GetArtifactLifecycleRuleRequest
	GetRuleId() *string
}

type GetArtifactLifecycleRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cralr-a18bkiajy81****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetArtifactLifecycleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactLifecycleRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactLifecycleRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *GetArtifactLifecycleRuleRequest) SetInstanceId(v string) *GetArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactLifecycleRuleRequest) SetRuleId(v string) *GetArtifactLifecycleRuleRequest {
	s.RuleId = &v
	return s
}

func (s *GetArtifactLifecycleRuleRequest) Validate() error {
	return dara.Validate(s)
}
