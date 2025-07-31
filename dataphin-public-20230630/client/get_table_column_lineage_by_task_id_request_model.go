// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnLineageByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetTableColumnLineageByTaskIdRequest
	GetOpTenantId() *int64
	SetTableColumnLineageByTaskIdQuery(v *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) *GetTableColumnLineageByTaskIdRequest
	GetTableColumnLineageByTaskIdQuery() *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery
}

type GetTableColumnLineageByTaskIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	TableColumnLineageByTaskIdQuery *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery `json:"TableColumnLineageByTaskIdQuery,omitempty" xml:"TableColumnLineageByTaskIdQuery,omitempty" type:"Struct"`
}

func (s GetTableColumnLineageByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineageByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineageByTaskIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableColumnLineageByTaskIdRequest) GetTableColumnLineageByTaskIdQuery() *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery {
	return s.TableColumnLineageByTaskIdQuery
}

func (s *GetTableColumnLineageByTaskIdRequest) SetOpTenantId(v int64) *GetTableColumnLineageByTaskIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdRequest) SetTableColumnLineageByTaskIdQuery(v *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) *GetTableColumnLineageByTaskIdRequest {
	s.TableColumnLineageByTaskIdQuery = v
	return s
}

func (s *GetTableColumnLineageByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}

type GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery struct {
	NeedNotExistObject *bool `json:"NeedNotExistObject,omitempty" xml:"NeedNotExistObject,omitempty"`
	// example:
	//
	// DEV
	TaskEnv *string `json:"TaskEnv,omitempty" xml:"TaskEnv,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// n_123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) GetNeedNotExistObject() *bool {
	return s.NeedNotExistObject
}

func (s *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) GetTaskEnv() *string {
	return s.TaskEnv
}

func (s *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) SetNeedNotExistObject(v bool) *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery {
	s.NeedNotExistObject = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) SetTaskEnv(v string) *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery {
	s.TaskEnv = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) SetTaskId(v string) *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery {
	s.TaskId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdRequestTableColumnLineageByTaskIdQuery) Validate() error {
	return dara.Validate(s)
}
