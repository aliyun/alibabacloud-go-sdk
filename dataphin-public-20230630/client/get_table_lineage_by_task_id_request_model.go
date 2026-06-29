// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableLineageByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetTableLineageByTaskIdRequest
	GetOpTenantId() *int64
	SetTableLineageByTaskIdQuery(v *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) *GetTableLineageByTaskIdRequest
	GetTableLineageByTaskIdQuery() *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery
}

type GetTableLineageByTaskIdRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The data structure for querying table lineage.
	//
	// This parameter is required.
	TableLineageByTaskIdQuery *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery `json:"TableLineageByTaskIdQuery,omitempty" xml:"TableLineageByTaskIdQuery,omitempty" type:"Struct"`
}

func (s GetTableLineageByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineageByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetTableLineageByTaskIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTableLineageByTaskIdRequest) GetTableLineageByTaskIdQuery() *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery {
	return s.TableLineageByTaskIdQuery
}

func (s *GetTableLineageByTaskIdRequest) SetOpTenantId(v int64) *GetTableLineageByTaskIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTableLineageByTaskIdRequest) SetTableLineageByTaskIdQuery(v *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) *GetTableLineageByTaskIdRequest {
	s.TableLineageByTaskIdQuery = v
	return s
}

func (s *GetTableLineageByTaskIdRequest) Validate() error {
	if s.TableLineageByTaskIdQuery != nil {
		if err := s.TableLineageByTaskIdQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery struct {
	// Specifies whether to return tables that do not exist in the asset inventory. If this parameter is not specified, non-existent tables are not returned.
	NeedNotExistObject *bool `json:"NeedNotExistObject,omitempty" xml:"NeedNotExistObject,omitempty"`
	// The environment of the task to query: DEV or PROD.
	//
	// example:
	//
	// DEV
	TaskEnv *string `json:"TaskEnv,omitempty" xml:"TaskEnv,omitempty"`
	// The task (node) ID used to filter the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// n_123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) GoString() string {
	return s.String()
}

func (s *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) GetNeedNotExistObject() *bool {
	return s.NeedNotExistObject
}

func (s *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) GetTaskEnv() *string {
	return s.TaskEnv
}

func (s *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) SetNeedNotExistObject(v bool) *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery {
	s.NeedNotExistObject = &v
	return s
}

func (s *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) SetTaskEnv(v string) *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery {
	s.TaskEnv = &v
	return s
}

func (s *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) SetTaskId(v string) *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery {
	s.TaskId = &v
	return s
}

func (s *GetTableLineageByTaskIdRequestTableLineageByTaskIdQuery) Validate() error {
	return dara.Validate(s)
}
