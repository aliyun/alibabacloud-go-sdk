// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanQueryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwner(v string) *DsgDesensPlanQueryListRequest
	GetOwner() *string
	SetPageNumber(v int32) *DsgDesensPlanQueryListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DsgDesensPlanQueryListRequest
	GetPageSize() *int32
	SetRuleName(v string) *DsgDesensPlanQueryListRequest
	GetRuleName() *string
	SetSceneId(v int64) *DsgDesensPlanQueryListRequest
	GetSceneId() *int64
	SetStatus(v int32) *DsgDesensPlanQueryListRequest
	GetStatus() *int32
	SetColumns(v []*DsgDesensPlanQueryListRequestColumns) *DsgDesensPlanQueryListRequest
	GetColumns() []*DsgDesensPlanQueryListRequestColumns
	SetDataType(v string) *DsgDesensPlanQueryListRequest
	GetDataType() *string
	SetEmptyNotDesesn(v string) *DsgDesensPlanQueryListRequest
	GetEmptyNotDesesn() *string
}

type DsgDesensPlanQueryListRequest struct {
	// The owner of the data masking rule.
	//
	// example:
	//
	// user1
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the sensitive field.
	//
	// example:
	//
	// phone
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The ID of the level-2 data masking scenario. You can call the [DsgSceneQuerySceneListByName](https://help.aliyun.com/document_detail/2786322.html) operation to query the list of IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The status of the data masking rule. Valid values:
	//
	// 	- 0: expired
	//
	// 	- 1: effective
	//
	// example:
	//
	// 1
	Status         *int32                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Columns        []*DsgDesensPlanQueryListRequestColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	DataType       *string                                 `json:"dataType,omitempty" xml:"dataType,omitempty"`
	EmptyNotDesesn *string                                 `json:"emptyNotDesesn,omitempty" xml:"emptyNotDesesn,omitempty"`
}

func (s DsgDesensPlanQueryListRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListRequest) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListRequest) GetOwner() *string {
	return s.Owner
}

func (s *DsgDesensPlanQueryListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgDesensPlanQueryListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgDesensPlanQueryListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DsgDesensPlanQueryListRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DsgDesensPlanQueryListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DsgDesensPlanQueryListRequest) GetColumns() []*DsgDesensPlanQueryListRequestColumns {
	return s.Columns
}

func (s *DsgDesensPlanQueryListRequest) GetDataType() *string {
	return s.DataType
}

func (s *DsgDesensPlanQueryListRequest) GetEmptyNotDesesn() *string {
	return s.EmptyNotDesesn
}

func (s *DsgDesensPlanQueryListRequest) SetOwner(v string) *DsgDesensPlanQueryListRequest {
	s.Owner = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetPageNumber(v int32) *DsgDesensPlanQueryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetPageSize(v int32) *DsgDesensPlanQueryListRequest {
	s.PageSize = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetRuleName(v string) *DsgDesensPlanQueryListRequest {
	s.RuleName = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetSceneId(v int64) *DsgDesensPlanQueryListRequest {
	s.SceneId = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetStatus(v int32) *DsgDesensPlanQueryListRequest {
	s.Status = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetColumns(v []*DsgDesensPlanQueryListRequestColumns) *DsgDesensPlanQueryListRequest {
	s.Columns = v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetDataType(v string) *DsgDesensPlanQueryListRequest {
	s.DataType = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetEmptyNotDesesn(v string) *DsgDesensPlanQueryListRequest {
	s.EmptyNotDesesn = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DsgDesensPlanQueryListRequestColumns struct {
	Column  *string `json:"column,omitempty" xml:"column,omitempty"`
	DbType  *string `json:"dbType,omitempty" xml:"dbType,omitempty"`
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	Table   *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s DsgDesensPlanQueryListRequestColumns) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListRequestColumns) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListRequestColumns) GetColumn() *string {
	return s.Column
}

func (s *DsgDesensPlanQueryListRequestColumns) GetDbType() *string {
	return s.DbType
}

func (s *DsgDesensPlanQueryListRequestColumns) GetProject() *string {
	return s.Project
}

func (s *DsgDesensPlanQueryListRequestColumns) GetTable() *string {
	return s.Table
}

func (s *DsgDesensPlanQueryListRequestColumns) SetColumn(v string) *DsgDesensPlanQueryListRequestColumns {
	s.Column = &v
	return s
}

func (s *DsgDesensPlanQueryListRequestColumns) SetDbType(v string) *DsgDesensPlanQueryListRequestColumns {
	s.DbType = &v
	return s
}

func (s *DsgDesensPlanQueryListRequestColumns) SetProject(v string) *DsgDesensPlanQueryListRequestColumns {
	s.Project = &v
	return s
}

func (s *DsgDesensPlanQueryListRequestColumns) SetTable(v string) *DsgDesensPlanQueryListRequestColumns {
	s.Table = &v
	return s
}

func (s *DsgDesensPlanQueryListRequestColumns) Validate() error {
	return dara.Validate(s)
}
