// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *GetInstanceRequest
	GetInstanceId() *int64
	SetProjectEnv(v string) *GetInstanceRequest
	GetProjectEnv() *string
}

type GetInstanceRequest struct {
	// The instance ID. You can call the [ListInstances](https://help.aliyun.com/document_detail/173982.html) operation to query the ID.
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

func (s GetInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetInstanceRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetInstanceRequest) SetInstanceId(v int64) *GetInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceRequest) SetProjectEnv(v string) *GetInstanceRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetInstanceRequest) Validate() error {
	return dara.Validate(s)
}
