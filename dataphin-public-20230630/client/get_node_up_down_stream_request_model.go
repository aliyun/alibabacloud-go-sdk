// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeUpDownStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownStreamDepth(v int32) *GetNodeUpDownStreamRequest
	GetDownStreamDepth() *int32
	SetEnv(v string) *GetNodeUpDownStreamRequest
	GetEnv() *string
	SetNodeId(v *GetNodeUpDownStreamRequestNodeId) *GetNodeUpDownStreamRequest
	GetNodeId() *GetNodeUpDownStreamRequestNodeId
	SetOpTenantId(v int64) *GetNodeUpDownStreamRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetNodeUpDownStreamRequest
	GetProjectId() *int64
	SetUpStreamDepth(v int32) *GetNodeUpDownStreamRequest
	GetUpStreamDepth() *int32
}

type GetNodeUpDownStreamRequest struct {
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	NodeId *GetNodeUpDownStreamRequestNodeId `json:"NodeId,omitempty" xml:"NodeId,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 113123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1
	UpStreamDepth *int32 `json:"UpStreamDepth,omitempty" xml:"UpStreamDepth,omitempty"`
}

func (s GetNodeUpDownStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeUpDownStreamRequest) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamRequest) GetDownStreamDepth() *int32 {
	return s.DownStreamDepth
}

func (s *GetNodeUpDownStreamRequest) GetEnv() *string {
	return s.Env
}

func (s *GetNodeUpDownStreamRequest) GetNodeId() *GetNodeUpDownStreamRequestNodeId {
	return s.NodeId
}

func (s *GetNodeUpDownStreamRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetNodeUpDownStreamRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNodeUpDownStreamRequest) GetUpStreamDepth() *int32 {
	return s.UpStreamDepth
}

func (s *GetNodeUpDownStreamRequest) SetDownStreamDepth(v int32) *GetNodeUpDownStreamRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetEnv(v string) *GetNodeUpDownStreamRequest {
	s.Env = &v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetNodeId(v *GetNodeUpDownStreamRequestNodeId) *GetNodeUpDownStreamRequest {
	s.NodeId = v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetOpTenantId(v int64) *GetNodeUpDownStreamRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetProjectId(v int64) *GetNodeUpDownStreamRequest {
	s.ProjectId = &v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetUpStreamDepth(v int32) *GetNodeUpDownStreamRequest {
	s.UpStreamDepth = &v
	return s
}

func (s *GetNodeUpDownStreamRequest) Validate() error {
	return dara.Validate(s)
}

type GetNodeUpDownStreamRequestNodeId struct {
	// example:
	//
	// 12
	FieldIdList *string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11313
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetNodeUpDownStreamRequestNodeId) String() string {
	return dara.Prettify(s)
}

func (s GetNodeUpDownStreamRequestNodeId) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamRequestNodeId) GetFieldIdList() *string {
	return s.FieldIdList
}

func (s *GetNodeUpDownStreamRequestNodeId) GetId() *string {
	return s.Id
}

func (s *GetNodeUpDownStreamRequestNodeId) SetFieldIdList(v string) *GetNodeUpDownStreamRequestNodeId {
	s.FieldIdList = &v
	return s
}

func (s *GetNodeUpDownStreamRequestNodeId) SetId(v string) *GetNodeUpDownStreamRequestNodeId {
	s.Id = &v
	return s
}

func (s *GetNodeUpDownStreamRequestNodeId) Validate() error {
	return dara.Validate(s)
}
