// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuccessInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *SetSuccessInstanceRequest
	GetInstanceId() *int64
	SetProjectEnv(v string) *SetSuccessInstanceRequest
	GetProjectEnv() *string
}

type SetSuccessInstanceRequest struct {
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

func (s SetSuccessInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSuccessInstanceRequest) GoString() string {
	return s.String()
}

func (s *SetSuccessInstanceRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *SetSuccessInstanceRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *SetSuccessInstanceRequest) SetInstanceId(v int64) *SetSuccessInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *SetSuccessInstanceRequest) SetProjectEnv(v string) *SetSuccessInstanceRequest {
	s.ProjectEnv = &v
	return s
}

func (s *SetSuccessInstanceRequest) Validate() error {
	return dara.Validate(s)
}
