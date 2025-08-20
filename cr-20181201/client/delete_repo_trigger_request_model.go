// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoTriggerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteRepoTriggerRequest
	GetInstanceId() *string
	SetRepoId(v string) *DeleteRepoTriggerRequest
	GetRepoId() *string
	SetTriggerId(v string) *DeleteRepoTriggerRequest
	GetTriggerId() *string
}

type DeleteRepoTriggerRequest struct {
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
	// The ID of the trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// crw-0z4pf81pgz35****
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
}

func (s DeleteRepoTriggerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepoTriggerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRepoTriggerRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DeleteRepoTriggerRequest) GetTriggerId() *string {
	return s.TriggerId
}

func (s *DeleteRepoTriggerRequest) SetInstanceId(v string) *DeleteRepoTriggerRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepoTriggerRequest) SetRepoId(v string) *DeleteRepoTriggerRequest {
	s.RepoId = &v
	return s
}

func (s *DeleteRepoTriggerRequest) SetTriggerId(v string) *DeleteRepoTriggerRequest {
	s.TriggerId = &v
	return s
}

func (s *DeleteRepoTriggerRequest) Validate() error {
	return dara.Validate(s)
}
