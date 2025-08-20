// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelArtifactBuildTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildTaskId(v string) *CancelArtifactBuildTaskRequest
	GetBuildTaskId() *string
	SetInstanceId(v string) *CancelArtifactBuildTaskRequest
	GetInstanceId() *string
}

type CancelArtifactBuildTaskRequest struct {
	// The ID of the artifact building task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i2ei-12****
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

func (s CancelArtifactBuildTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelArtifactBuildTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelArtifactBuildTaskRequest) GetBuildTaskId() *string {
	return s.BuildTaskId
}

func (s *CancelArtifactBuildTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelArtifactBuildTaskRequest) SetBuildTaskId(v string) *CancelArtifactBuildTaskRequest {
	s.BuildTaskId = &v
	return s
}

func (s *CancelArtifactBuildTaskRequest) SetInstanceId(v string) *CancelArtifactBuildTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelArtifactBuildTaskRequest) Validate() error {
	return dara.Validate(s)
}
