// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLineageInfoParams interface {
	dara.Model
	String() string
	GoString() string
	SetDepth(v int64) *GetLineageInfoParams
	GetDepth() *int64
	SetDirection(v string) *GetLineageInfoParams
	GetDirection() *string
	SetId(v string) *GetLineageInfoParams
	GetId() *string
	SetIdType(v string) *GetLineageInfoParams
	GetIdType() *string
	SetIsColumnLevel(v bool) *GetLineageInfoParams
	GetIsColumnLevel() *bool
	SetIsTemporary(v bool) *GetLineageInfoParams
	GetIsTemporary() *bool
	SetNamespace(v string) *GetLineageInfoParams
	GetNamespace() *string
	SetWorkspace(v string) *GetLineageInfoParams
	GetWorkspace() *string
}

type GetLineageInfoParams struct {
	// The depth of the data lineage.
	//
	// example:
	//
	// 1
	Depth *int64 `json:"depth,omitempty" xml:"depth,omitempty"`
	// The direction of the lineage. Valid values:
	//
	// 	- UPSTREAM: searches for upstream operators.
	//
	// 	- DOWNSTREAM: searches for downstream operators.
	//
	// 	- BOTH: searches for both upstream and downstream operators.
	//
	// example:
	//
	// Both
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The job ID or the table ID. You can call the ListJobs operation to obtain the job ID or call the ListTables operation to obtain the table ID.
	//
	// example:
	//
	// 664cc64d-5dea-4ad3-9ee4-8432a874****
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The data type of the job or table. If the id parameter is set to the job ID, the value of this parameter is the data type of the job. If the id parameter is set to the table ID, the value of this parameter is the data type of the table.
	//
	// example:
	//
	// JOB
	IdType *string `json:"idType,omitempty" xml:"idType,omitempty"`
	// The lineage type. The value true indicates a field-level lineage. The value false indicates a table-level lineage.
	//
	// example:
	//
	// true
	IsColumnLevel *bool `json:"isColumnLevel,omitempty" xml:"isColumnLevel,omitempty"`
	// Indicates whether the table was a temporary table.
	//
	// example:
	//
	// true
	IsTemporary *bool `json:"isTemporary,omitempty" xml:"isTemporary,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetLineageInfoParams) String() string {
	return dara.Prettify(s)
}

func (s GetLineageInfoParams) GoString() string {
	return s.String()
}

func (s *GetLineageInfoParams) GetDepth() *int64 {
	return s.Depth
}

func (s *GetLineageInfoParams) GetDirection() *string {
	return s.Direction
}

func (s *GetLineageInfoParams) GetId() *string {
	return s.Id
}

func (s *GetLineageInfoParams) GetIdType() *string {
	return s.IdType
}

func (s *GetLineageInfoParams) GetIsColumnLevel() *bool {
	return s.IsColumnLevel
}

func (s *GetLineageInfoParams) GetIsTemporary() *bool {
	return s.IsTemporary
}

func (s *GetLineageInfoParams) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLineageInfoParams) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetLineageInfoParams) SetDepth(v int64) *GetLineageInfoParams {
	s.Depth = &v
	return s
}

func (s *GetLineageInfoParams) SetDirection(v string) *GetLineageInfoParams {
	s.Direction = &v
	return s
}

func (s *GetLineageInfoParams) SetId(v string) *GetLineageInfoParams {
	s.Id = &v
	return s
}

func (s *GetLineageInfoParams) SetIdType(v string) *GetLineageInfoParams {
	s.IdType = &v
	return s
}

func (s *GetLineageInfoParams) SetIsColumnLevel(v bool) *GetLineageInfoParams {
	s.IsColumnLevel = &v
	return s
}

func (s *GetLineageInfoParams) SetIsTemporary(v bool) *GetLineageInfoParams {
	s.IsTemporary = &v
	return s
}

func (s *GetLineageInfoParams) SetNamespace(v string) *GetLineageInfoParams {
	s.Namespace = &v
	return s
}

func (s *GetLineageInfoParams) SetWorkspace(v string) *GetLineageInfoParams {
	s.Workspace = &v
	return s
}

func (s *GetLineageInfoParams) Validate() error {
	return dara.Validate(s)
}
