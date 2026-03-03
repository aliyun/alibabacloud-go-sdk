// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *SearchTracesByPageRequest
	GetEndTime() *string
	SetEnv(v string) *SearchTracesByPageRequest
	GetEnv() *string
	SetMinDuration(v int64) *SearchTracesByPageRequest
	GetMinDuration() *int64
	SetOperationName(v string) *SearchTracesByPageRequest
	GetOperationName() *string
	SetOrderBy(v string) *SearchTracesByPageRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *SearchTracesByPageRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *SearchTracesByPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchTracesByPageRequest
	GetPageSize() *int32
	SetServiceGroupId(v int64) *SearchTracesByPageRequest
	GetServiceGroupId() *int64
	SetServiceName(v string) *SearchTracesByPageRequest
	GetServiceName() *string
	SetStartTime(v string) *SearchTracesByPageRequest
	GetStartTime() *string
}

type SearchTracesByPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-08 15:03:21
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 100
	MinDuration *int64 `json:"minDuration,omitempty" xml:"minDuration,omitempty"`
	// example:
	//
	// /demo/queryNotExistDB/11
	OperationName *string `json:"operationName,omitempty" xml:"operationName,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// example:
	//
	// dev-sellercenter
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-31 11:10:22
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s SearchTracesByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *SearchTracesByPageRequest) GetEnv() *string {
	return s.Env
}

func (s *SearchTracesByPageRequest) GetMinDuration() *int64 {
	return s.MinDuration
}

func (s *SearchTracesByPageRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *SearchTracesByPageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchTracesByPageRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *SearchTracesByPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTracesByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTracesByPageRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *SearchTracesByPageRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *SearchTracesByPageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *SearchTracesByPageRequest) SetEndTime(v string) *SearchTracesByPageRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesByPageRequest) SetEnv(v string) *SearchTracesByPageRequest {
	s.Env = &v
	return s
}

func (s *SearchTracesByPageRequest) SetMinDuration(v int64) *SearchTracesByPageRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesByPageRequest) SetOperationName(v string) *SearchTracesByPageRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetOrderBy(v string) *SearchTracesByPageRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchTracesByPageRequest) SetOrderDirection(v string) *SearchTracesByPageRequest {
	s.OrderDirection = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageNumber(v int32) *SearchTracesByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageSize(v int32) *SearchTracesByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceGroupId(v int64) *SearchTracesByPageRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceName(v string) *SearchTracesByPageRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetStartTime(v string) *SearchTracesByPageRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesByPageRequest) Validate() error {
	return dara.Validate(s)
}
