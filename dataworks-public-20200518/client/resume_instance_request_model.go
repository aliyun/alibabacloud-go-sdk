// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *ResumeInstanceRequest
	GetInstanceId() *int64
	SetProjectEnv(v string) *ResumeInstanceRequest
	GetProjectEnv() *string
}

type ResumeInstanceRequest struct {
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

func (s ResumeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ResumeInstanceRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ResumeInstanceRequest) SetInstanceId(v int64) *ResumeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ResumeInstanceRequest) SetProjectEnv(v string) *ResumeInstanceRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ResumeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
