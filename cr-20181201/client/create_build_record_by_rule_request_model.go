// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBuildRecordByRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRuleId(v string) *CreateBuildRecordByRuleRequest
	GetBuildRuleId() *string
	SetInstanceId(v string) *CreateBuildRecordByRuleRequest
	GetInstanceId() *string
	SetRepoId(v string) *CreateBuildRecordByRuleRequest
	GetRepoId() *string
}

type CreateBuildRecordByRuleRequest struct {
	// The ID of the image building rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crbr-1j95g4bu2s1i****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-asd6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-8dz3aedjqlmk****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CreateBuildRecordByRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBuildRecordByRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRuleRequest) GetBuildRuleId() *string {
	return s.BuildRuleId
}

func (s *CreateBuildRecordByRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBuildRecordByRuleRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateBuildRecordByRuleRequest) SetBuildRuleId(v string) *CreateBuildRecordByRuleRequest {
	s.BuildRuleId = &v
	return s
}

func (s *CreateBuildRecordByRuleRequest) SetInstanceId(v string) *CreateBuildRecordByRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBuildRecordByRuleRequest) SetRepoId(v string) *CreateBuildRecordByRuleRequest {
	s.RepoId = &v
	return s
}

func (s *CreateBuildRecordByRuleRequest) Validate() error {
	return dara.Validate(s)
}
