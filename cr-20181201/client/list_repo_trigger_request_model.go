// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRepoTriggerRequest
	GetInstanceId() *string
	SetRepoId(v string) *ListRepoTriggerRequest
	GetRepoId() *string
}

type ListRepoTriggerRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTriggerRequest) GoString() string {
	return s.String()
}

func (s *ListRepoTriggerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoTriggerRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListRepoTriggerRequest) SetInstanceId(v string) *ListRepoTriggerRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoTriggerRequest) SetRepoId(v string) *ListRepoTriggerRequest {
	s.RepoId = &v
	return s
}

func (s *ListRepoTriggerRequest) Validate() error {
	return dara.Validate(s)
}
