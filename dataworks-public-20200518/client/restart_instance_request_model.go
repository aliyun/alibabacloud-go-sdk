// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *RestartInstanceRequest
	GetInstanceId() *int64
	SetProjectEnv(v string) *RestartInstanceRequest
	GetProjectEnv() *string
}

type RestartInstanceRequest struct {
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

func (s RestartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *RestartInstanceRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *RestartInstanceRequest) SetInstanceId(v int64) *RestartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartInstanceRequest) SetProjectEnv(v string) *RestartInstanceRequest {
	s.ProjectEnv = &v
	return s
}

func (s *RestartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
