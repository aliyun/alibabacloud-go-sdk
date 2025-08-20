// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoBuildRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRuleId(v string) *DeleteRepoBuildRuleRequest
	GetBuildRuleId() *string
	SetInstanceId(v string) *DeleteRepoBuildRuleRequest
	GetInstanceId() *string
	SetRepoId(v string) *DeleteRepoBuildRuleRequest
	GetRepoId() *string
}

type DeleteRepoBuildRuleRequest struct {
	// The ID of the image building rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crbr-36tffn0kouvi****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s DeleteRepoBuildRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepoBuildRuleRequest) GetBuildRuleId() *string {
	return s.BuildRuleId
}

func (s *DeleteRepoBuildRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRepoBuildRuleRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DeleteRepoBuildRuleRequest) SetBuildRuleId(v string) *DeleteRepoBuildRuleRequest {
	s.BuildRuleId = &v
	return s
}

func (s *DeleteRepoBuildRuleRequest) SetInstanceId(v string) *DeleteRepoBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepoBuildRuleRequest) SetRepoId(v string) *DeleteRepoBuildRuleRequest {
	s.RepoId = &v
	return s
}

func (s *DeleteRepoBuildRuleRequest) Validate() error {
	return dara.Validate(s)
}
