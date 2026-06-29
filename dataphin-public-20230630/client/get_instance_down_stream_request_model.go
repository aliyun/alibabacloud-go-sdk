// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceDownStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownStreamDepth(v int32) *GetInstanceDownStreamRequest
	GetDownStreamDepth() *int32
	SetEnv(v string) *GetInstanceDownStreamRequest
	GetEnv() *string
	SetInstanceGet(v *GetInstanceDownStreamRequestInstanceGet) *GetInstanceDownStreamRequest
	GetInstanceGet() *GetInstanceDownStreamRequestInstanceGet
	SetOpTenantId(v int64) *GetInstanceDownStreamRequest
	GetOpTenantId() *int64
	SetRunStatus(v string) *GetInstanceDownStreamRequest
	GetRunStatus() *string
}

type GetInstanceDownStreamRequest struct {
	// Number of levels to expand downstream in the DAG query. Valid values: 1 to 6.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// Environment identifier.
	//
	// - DEV: Development environment.
	//
	// - PROD (default): Production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The request body.
	//
	// This parameter is required.
	InstanceGet *GetInstanceDownStreamRequestInstanceGet `json:"InstanceGet,omitempty" xml:"InstanceGet,omitempty" type:"Struct"`
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// Run status of the instance.
	//
	// - INIT
	//
	// - WATING
	//
	// - RUNNING
	//
	// - SUCCESS
	//
	// - FAILED
	//
	// example:
	//
	// SUCCESS
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
}

func (s GetInstanceDownStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDownStreamRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamRequest) GetDownStreamDepth() *int32 {
	return s.DownStreamDepth
}

func (s *GetInstanceDownStreamRequest) GetEnv() *string {
	return s.Env
}

func (s *GetInstanceDownStreamRequest) GetInstanceGet() *GetInstanceDownStreamRequestInstanceGet {
	return s.InstanceGet
}

func (s *GetInstanceDownStreamRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetInstanceDownStreamRequest) GetRunStatus() *string {
	return s.RunStatus
}

func (s *GetInstanceDownStreamRequest) SetDownStreamDepth(v int32) *GetInstanceDownStreamRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceDownStreamRequest) SetEnv(v string) *GetInstanceDownStreamRequest {
	s.Env = &v
	return s
}

func (s *GetInstanceDownStreamRequest) SetInstanceGet(v *GetInstanceDownStreamRequestInstanceGet) *GetInstanceDownStreamRequest {
	s.InstanceGet = v
	return s
}

func (s *GetInstanceDownStreamRequest) SetOpTenantId(v int64) *GetInstanceDownStreamRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetInstanceDownStreamRequest) SetRunStatus(v string) *GetInstanceDownStreamRequest {
	s.RunStatus = &v
	return s
}

func (s *GetInstanceDownStreamRequest) Validate() error {
	if s.InstanceGet != nil {
		if err := s.InstanceGet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceDownStreamRequestInstanceGet struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t_5929472_20210411_9577721
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Node type of the instance.
	//
	// - BBOX_LOGIC_TABLE_NODE
	//
	// - BBOX_LOGIC_FIELD_NODE
	//
	// - BBOX_LOGIC_FIELD_GROUP_NODE
	//
	// - BBOX_INNER_TEMP_NODE
	//
	// - DATA_PROCESS
	//
	// - STREAM_TASK_NODE
	//
	// - FLINK_BATCH
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s GetInstanceDownStreamRequestInstanceGet) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDownStreamRequestInstanceGet) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamRequestInstanceGet) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceDownStreamRequestInstanceGet) GetNodeType() *string {
	return s.NodeType
}

func (s *GetInstanceDownStreamRequestInstanceGet) SetInstanceId(v string) *GetInstanceDownStreamRequestInstanceGet {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceDownStreamRequestInstanceGet) SetNodeType(v string) *GetInstanceDownStreamRequestInstanceGet {
	s.NodeType = &v
	return s
}

func (s *GetInstanceDownStreamRequestInstanceGet) Validate() error {
	return dara.Validate(s)
}
