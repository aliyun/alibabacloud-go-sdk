// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *SuspendInstanceRequest
	GetInstanceId() *int64
	SetProjectEnv(v string) *SuspendInstanceRequest
	GetProjectEnv() *string
}

type SuspendInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s SuspendInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendInstanceRequest) GoString() string {
	return s.String()
}

func (s *SuspendInstanceRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *SuspendInstanceRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *SuspendInstanceRequest) SetInstanceId(v int64) *SuspendInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *SuspendInstanceRequest) SetProjectEnv(v string) *SuspendInstanceRequest {
	s.ProjectEnv = &v
	return s
}

func (s *SuspendInstanceRequest) Validate() error {
	return dara.Validate(s)
}
