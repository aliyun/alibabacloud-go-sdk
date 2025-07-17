// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceSharedRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int64) *ListDataSourceSharedRulesRequest
	GetDataSourceId() *int64
	SetTargetProjectId(v int64) *ListDataSourceSharedRulesRequest
	GetTargetProjectId() *int64
}

type ListDataSourceSharedRulesRequest struct {
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The ID of the workspace to which the data source is shared. You cannot share the data source to the workspace with which the data source is associated.
	//
	// example:
	//
	// 1
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s ListDataSourceSharedRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceSharedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *ListDataSourceSharedRulesRequest) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *ListDataSourceSharedRulesRequest) SetDataSourceId(v int64) *ListDataSourceSharedRulesRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListDataSourceSharedRulesRequest) SetTargetProjectId(v int64) *ListDataSourceSharedRulesRequest {
	s.TargetProjectId = &v
	return s
}

func (s *ListDataSourceSharedRulesRequest) Validate() error {
	return dara.Validate(s)
}
