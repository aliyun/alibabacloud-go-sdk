// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *StopInstanceRequest
	GetInstanceId() *int64
	SetProjectEnv(v string) *StopInstanceRequest
	GetProjectEnv() *string
}

type StopInstanceRequest struct {
	// The instance ID. You can call the [ListInstances](https://help.aliyun.com/document_detail/173982.html) operation to obtain the ID.
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

func (s StopInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *StopInstanceRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *StopInstanceRequest) SetInstanceId(v int64) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetProjectEnv(v string) *StopInstanceRequest {
	s.ProjectEnv = &v
	return s
}

func (s *StopInstanceRequest) Validate() error {
	return dara.Validate(s)
}
