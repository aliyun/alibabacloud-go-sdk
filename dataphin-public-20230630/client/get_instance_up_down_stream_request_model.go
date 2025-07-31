// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceUpDownStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownStreamDepth(v int32) *GetInstanceUpDownStreamRequest
	GetDownStreamDepth() *int32
	SetEnv(v string) *GetInstanceUpDownStreamRequest
	GetEnv() *string
	SetInstanceId(v *GetInstanceUpDownStreamRequestInstanceId) *GetInstanceUpDownStreamRequest
	GetInstanceId() *GetInstanceUpDownStreamRequestInstanceId
	SetOpTenantId(v int64) *GetInstanceUpDownStreamRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetInstanceUpDownStreamRequest
	GetProjectId() *int64
	SetUpStreamDepth(v int32) *GetInstanceUpDownStreamRequest
	GetUpStreamDepth() *int32
}

type GetInstanceUpDownStreamRequest struct {
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	InstanceId *GetInstanceUpDownStreamRequestInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1
	UpStreamDepth *int32 `json:"UpStreamDepth,omitempty" xml:"UpStreamDepth,omitempty"`
}

func (s GetInstanceUpDownStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUpDownStreamRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamRequest) GetDownStreamDepth() *int32 {
	return s.DownStreamDepth
}

func (s *GetInstanceUpDownStreamRequest) GetEnv() *string {
	return s.Env
}

func (s *GetInstanceUpDownStreamRequest) GetInstanceId() *GetInstanceUpDownStreamRequestInstanceId {
	return s.InstanceId
}

func (s *GetInstanceUpDownStreamRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetInstanceUpDownStreamRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetInstanceUpDownStreamRequest) GetUpStreamDepth() *int32 {
	return s.UpStreamDepth
}

func (s *GetInstanceUpDownStreamRequest) SetDownStreamDepth(v int32) *GetInstanceUpDownStreamRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetEnv(v string) *GetInstanceUpDownStreamRequest {
	s.Env = &v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetInstanceId(v *GetInstanceUpDownStreamRequestInstanceId) *GetInstanceUpDownStreamRequest {
	s.InstanceId = v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetOpTenantId(v int64) *GetInstanceUpDownStreamRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetProjectId(v int64) *GetInstanceUpDownStreamRequest {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetUpStreamDepth(v int32) *GetInstanceUpDownStreamRequest {
	s.UpStreamDepth = &v
	return s
}

func (s *GetInstanceUpDownStreamRequest) Validate() error {
	return dara.Validate(s)
}

type GetInstanceUpDownStreamRequestInstanceId struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// t_123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetInstanceUpDownStreamRequestInstanceId) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUpDownStreamRequestInstanceId) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamRequestInstanceId) GetFieldInstanceIdList() []*string {
	return s.FieldInstanceIdList
}

func (s *GetInstanceUpDownStreamRequestInstanceId) GetId() *string {
	return s.Id
}

func (s *GetInstanceUpDownStreamRequestInstanceId) SetFieldInstanceIdList(v []*string) *GetInstanceUpDownStreamRequestInstanceId {
	s.FieldInstanceIdList = v
	return s
}

func (s *GetInstanceUpDownStreamRequestInstanceId) SetId(v string) *GetInstanceUpDownStreamRequestInstanceId {
	s.Id = &v
	return s
}

func (s *GetInstanceUpDownStreamRequestInstanceId) Validate() error {
	return dara.Validate(s)
}
