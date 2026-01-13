// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRecallManagementTablesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecallManagementTablesResponseBody
	GetNextToken() *string
	SetRecallManagementTables(v []*ListRecallManagementTablesResponseBodyRecallManagementTables) *ListRecallManagementTablesResponseBody
	GetRecallManagementTables() []*ListRecallManagementTablesResponseBodyRecallManagementTables
	SetRequestId(v string) *ListRecallManagementTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRecallManagementTablesResponseBody
	GetTotalCount() *string
}

type ListRecallManagementTablesResponseBody struct {
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ""
	NextToken              *string                                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RecallManagementTables []*ListRecallManagementTablesResponseBodyRecallManagementTables `json:"RecallManagementTables,omitempty" xml:"RecallManagementTables,omitempty" type:"Repeated"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecallManagementTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecallManagementTablesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecallManagementTablesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecallManagementTablesResponseBody) GetRecallManagementTables() []*ListRecallManagementTablesResponseBodyRecallManagementTables {
	return s.RecallManagementTables
}

func (s *ListRecallManagementTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecallManagementTablesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRecallManagementTablesResponseBody) SetMaxResults(v int32) *ListRecallManagementTablesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRecallManagementTablesResponseBody) SetNextToken(v string) *ListRecallManagementTablesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecallManagementTablesResponseBody) SetRecallManagementTables(v []*ListRecallManagementTablesResponseBodyRecallManagementTables) *ListRecallManagementTablesResponseBody {
	s.RecallManagementTables = v
	return s
}

func (s *ListRecallManagementTablesResponseBody) SetRequestId(v string) *ListRecallManagementTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecallManagementTablesResponseBody) SetTotalCount(v string) *ListRecallManagementTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecallManagementTablesResponseBody) Validate() error {
	if s.RecallManagementTables != nil {
		for _, item := range s.RecallManagementTables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecallManagementTablesResponseBodyRecallManagementTables struct {
	// example:
	//
	// false
	CanDelete *bool `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	// example:
	//
	// Api
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// example:
	//
	// this is a test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	IndexEffectiveTime *string `json:"IndexEffectiveTime,omitempty" xml:"IndexEffectiveTime,omitempty"`
	// example:
	//
	// 20250701
	IndexVersionId *string `json:"IndexVersionId,omitempty" xml:"IndexVersionId,omitempty"`
	// example:
	//
	// test
	MaxcomputeProjectName *string `json:"MaxcomputeProjectName,omitempty" xml:"MaxcomputeProjectName,omitempty"`
	// maxcompute schema。
	//
	// example:
	//
	// default
	MaxcomputeSchema *string `json:"MaxcomputeSchema,omitempty" xml:"MaxcomputeSchema,omitempty"`
	// example:
	//
	// table-1
	MaxcomputeTableName *string `json:"MaxcomputeTableName,omitempty" xml:"MaxcomputeTableName,omitempty"`
	// example:
	//
	// table-123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// dt
	PartitionFields *string `json:"PartitionFields,omitempty" xml:"PartitionFields,omitempty"`
	// example:
	//
	// 3
	RecallManagementTableId *string `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
	// example:
	//
	// X2I
	RecallType *string `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// example:
	//
	// Recall
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecallManagementTablesResponseBodyRecallManagementTables) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementTablesResponseBodyRecallManagementTables) GoString() string {
	return s.String()
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetDataSource() *string {
	return s.DataSource
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetDescription() *string {
	return s.Description
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetIndexEffectiveTime() *string {
	return s.IndexEffectiveTime
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetIndexVersionId() *string {
	return s.IndexVersionId
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetMaxcomputeProjectName() *string {
	return s.MaxcomputeProjectName
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetMaxcomputeSchema() *string {
	return s.MaxcomputeSchema
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetMaxcomputeTableName() *string {
	return s.MaxcomputeTableName
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetName() *string {
	return s.Name
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetPartitionFields() *string {
	return s.PartitionFields
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetRecallType() *string {
	return s.RecallType
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) GetType() *string {
	return s.Type
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetCanDelete(v bool) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.CanDelete = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetDataSource(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.DataSource = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetDescription(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.Description = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetGmtCreateTime(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.GmtCreateTime = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetGmtModifiedTime(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetIndexEffectiveTime(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.IndexEffectiveTime = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetIndexVersionId(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.IndexVersionId = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetMaxcomputeProjectName(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.MaxcomputeProjectName = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetMaxcomputeSchema(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.MaxcomputeSchema = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetMaxcomputeTableName(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.MaxcomputeTableName = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetName(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.Name = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetPartitionFields(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.PartitionFields = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetRecallManagementTableId(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.RecallManagementTableId = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetRecallType(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.RecallType = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) SetType(v string) *ListRecallManagementTablesResponseBodyRecallManagementTables {
	s.Type = &v
	return s
}

func (s *ListRecallManagementTablesResponseBodyRecallManagementTables) Validate() error {
	return dara.Validate(s)
}
