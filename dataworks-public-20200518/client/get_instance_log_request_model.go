// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceHistoryId(v int64) *GetInstanceLogRequest
	GetInstanceHistoryId() *int64
	SetInstanceId(v int64) *GetInstanceLogRequest
	GetInstanceId() *int64
	SetProjectEnv(v string) *GetInstanceLogRequest
	GetProjectEnv() *string
}

type GetInstanceLogRequest struct {
	// The historical record number of the instance. You can call the ListInstanceHistory operation to query the ID.
	//
	// example:
	//
	// 1
	InstanceHistoryId *int64 `json:"InstanceHistoryId,omitempty" xml:"InstanceHistoryId,omitempty"`
	// The ID of the instance.
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

func (s GetInstanceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLogRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceLogRequest) GetInstanceHistoryId() *int64 {
	return s.InstanceHistoryId
}

func (s *GetInstanceLogRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetInstanceLogRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetInstanceLogRequest) SetInstanceHistoryId(v int64) *GetInstanceLogRequest {
	s.InstanceHistoryId = &v
	return s
}

func (s *GetInstanceLogRequest) SetInstanceId(v int64) *GetInstanceLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceLogRequest) SetProjectEnv(v string) *GetInstanceLogRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetInstanceLogRequest) Validate() error {
	return dara.Validate(s)
}
