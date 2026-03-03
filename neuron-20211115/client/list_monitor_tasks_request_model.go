// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *ListMonitorTasksRequest
	GetAlertName() *string
	SetEnv(v string) *ListMonitorTasksRequest
	GetEnv() *string
	SetOrderBy(v string) *ListMonitorTasksRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMonitorTasksRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListMonitorTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMonitorTasksRequest
	GetPageSize() *int32
	SetPbcId(v string) *ListMonitorTasksRequest
	GetPbcId() *string
	SetServiceGroupId(v string) *ListMonitorTasksRequest
	GetServiceGroupId() *string
	SetType(v string) *ListMonitorTasksRequest
	GetType() *string
}

type ListMonitorTasksRequest struct {
	// example:
	//
	// alert
	AlertName *string `json:"alertName,omitempty" xml:"alertName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// id
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 2
	PbcId *string `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// example:
	//
	// 1
	ServiceGroupId *string `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// example:
	//
	// SLS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMonitorTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMonitorTasksRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *ListMonitorTasksRequest) GetEnv() *string {
	return s.Env
}

func (s *ListMonitorTasksRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMonitorTasksRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMonitorTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMonitorTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMonitorTasksRequest) GetPbcId() *string {
	return s.PbcId
}

func (s *ListMonitorTasksRequest) GetServiceGroupId() *string {
	return s.ServiceGroupId
}

func (s *ListMonitorTasksRequest) GetType() *string {
	return s.Type
}

func (s *ListMonitorTasksRequest) SetAlertName(v string) *ListMonitorTasksRequest {
	s.AlertName = &v
	return s
}

func (s *ListMonitorTasksRequest) SetEnv(v string) *ListMonitorTasksRequest {
	s.Env = &v
	return s
}

func (s *ListMonitorTasksRequest) SetOrderBy(v string) *ListMonitorTasksRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMonitorTasksRequest) SetOrderDirection(v string) *ListMonitorTasksRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMonitorTasksRequest) SetPageNumber(v int32) *ListMonitorTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMonitorTasksRequest) SetPageSize(v int32) *ListMonitorTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListMonitorTasksRequest) SetPbcId(v string) *ListMonitorTasksRequest {
	s.PbcId = &v
	return s
}

func (s *ListMonitorTasksRequest) SetServiceGroupId(v string) *ListMonitorTasksRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListMonitorTasksRequest) SetType(v string) *ListMonitorTasksRequest {
	s.Type = &v
	return s
}

func (s *ListMonitorTasksRequest) Validate() error {
	return dara.Validate(s)
}
