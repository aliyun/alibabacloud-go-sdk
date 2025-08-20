// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactBuildTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildTaskId(v string) *GetArtifactBuildTaskRequest
	GetBuildTaskId() *string
	SetInstanceId(v string) *GetArtifactBuildTaskRequest
	GetInstanceId() *string
}

type GetArtifactBuildTaskRequest struct {
	// The ID of the artifact building task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i2a-1yu****
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetArtifactBuildTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildTaskRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskRequest) GetBuildTaskId() *string {
	return s.BuildTaskId
}

func (s *GetArtifactBuildTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactBuildTaskRequest) SetBuildTaskId(v string) *GetArtifactBuildTaskRequest {
	s.BuildTaskId = &v
	return s
}

func (s *GetArtifactBuildTaskRequest) SetInstanceId(v string) *GetArtifactBuildTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactBuildTaskRequest) Validate() error {
	return dara.Validate(s)
}
