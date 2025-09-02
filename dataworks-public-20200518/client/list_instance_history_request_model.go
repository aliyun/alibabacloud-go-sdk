// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *ListInstanceHistoryRequest
	GetInstanceId() *int64
	SetProjectEnv(v string) *ListInstanceHistoryRequest
	GetProjectEnv() *string
}

type ListInstanceHistoryRequest struct {
	// The instance ID. You can call the ListInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV. By default, data of instances in the production environment is queried.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s ListInstanceHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceHistoryRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListInstanceHistoryRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListInstanceHistoryRequest) SetInstanceId(v int64) *ListInstanceHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceHistoryRequest) SetProjectEnv(v string) *ListInstanceHistoryRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListInstanceHistoryRequest) Validate() error {
	return dara.Validate(s)
}
