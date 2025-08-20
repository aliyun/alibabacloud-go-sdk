// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactSubscriptionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetArtifactSubscriptionTaskRequest
	GetInstanceId() *string
	SetTaskId(v string) *GetArtifactSubscriptionTaskRequest
	GetTaskId() *string
}

type GetArtifactSubscriptionTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crast-40le4es9yh0p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetArtifactSubscriptionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionTaskRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactSubscriptionTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetArtifactSubscriptionTaskRequest) SetInstanceId(v string) *GetArtifactSubscriptionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskRequest) SetTaskId(v string) *GetArtifactSubscriptionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskRequest) Validate() error {
	return dara.Validate(s)
}
