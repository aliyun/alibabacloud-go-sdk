// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostModelDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *ModelRouterQueryCostModelDetailRequest
	GetClientId() *int64
	SetEndTime(v int64) *ModelRouterQueryCostModelDetailRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *ModelRouterQueryCostModelDetailRequest
	GetMaxResults() *int32
	SetModelId(v int64) *ModelRouterQueryCostModelDetailRequest
	GetModelId() *int64
	SetNextToken(v string) *ModelRouterQueryCostModelDetailRequest
	GetNextToken() *string
	SetPage(v int32) *ModelRouterQueryCostModelDetailRequest
	GetPage() *int32
	SetPageIndex(v int32) *ModelRouterQueryCostModelDetailRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryCostModelDetailRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ModelRouterQueryCostModelDetailRequest
	GetStartTime() *int64
}

type ModelRouterQueryCostModelDetailRequest struct {
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1700086400
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// maxResults
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1700000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ModelRouterQueryCostModelDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostModelDetailRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostModelDetailRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryCostModelDetailRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModelRouterQueryCostModelDetailRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryCostModelDetailRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterQueryCostModelDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostModelDetailRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryCostModelDetailRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryCostModelDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryCostModelDetailRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModelRouterQueryCostModelDetailRequest) SetClientId(v int64) *ModelRouterQueryCostModelDetailRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetEndTime(v int64) *ModelRouterQueryCostModelDetailRequest {
	s.EndTime = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetMaxResults(v int32) *ModelRouterQueryCostModelDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetModelId(v int64) *ModelRouterQueryCostModelDetailRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetNextToken(v string) *ModelRouterQueryCostModelDetailRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetPage(v int32) *ModelRouterQueryCostModelDetailRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetPageIndex(v int32) *ModelRouterQueryCostModelDetailRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetPageSize(v int32) *ModelRouterQueryCostModelDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) SetStartTime(v int64) *ModelRouterQueryCostModelDetailRequest {
	s.StartTime = &v
	return s
}

func (s *ModelRouterQueryCostModelDetailRequest) Validate() error {
	return dara.Validate(s)
}
