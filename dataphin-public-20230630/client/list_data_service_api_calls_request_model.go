// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiCallsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataServiceApiCallsRequestListQuery) *ListDataServiceApiCallsRequest
	GetListQuery() *ListDataServiceApiCallsRequestListQuery
	SetOpTenantId(v int64) *ListDataServiceApiCallsRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceApiCallsRequest
	GetProjectId() *int32
}

type ListDataServiceApiCallsRequest struct {
	// This parameter is required.
	ListQuery *ListDataServiceApiCallsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataServiceApiCallsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallsRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallsRequest) GetListQuery() *ListDataServiceApiCallsRequestListQuery {
	return s.ListQuery
}

func (s *ListDataServiceApiCallsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceApiCallsRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceApiCallsRequest) SetListQuery(v *ListDataServiceApiCallsRequestListQuery) *ListDataServiceApiCallsRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataServiceApiCallsRequest) SetOpTenantId(v int64) *ListDataServiceApiCallsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceApiCallsRequest) SetProjectId(v int32) *ListDataServiceApiCallsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiCallsRequest) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApiCallsRequestListQuery struct {
	// example:
	//
	// 20122
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// test
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// appKey
	//
	// example:
	//
	// 1021
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 192.168.1.1
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30 20:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDataServiceApiCallsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallsRequestListQuery) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceApiCallsRequestListQuery) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServiceApiCallsRequestListQuery) GetAppKey() *int64 {
	return s.AppKey
}

func (s *ListDataServiceApiCallsRequestListQuery) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListDataServiceApiCallsRequestListQuery) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDataServiceApiCallsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDataServiceApiCallsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApiCallsRequestListQuery) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDataServiceApiCallsRequestListQuery) GetSuccessful() *bool {
	return s.Successful
}

func (s *ListDataServiceApiCallsRequestListQuery) SetApiId(v int64) *ListDataServiceApiCallsRequestListQuery {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceApiCallsRequestListQuery) SetApiName(v string) *ListDataServiceApiCallsRequestListQuery {
	s.ApiName = &v
	return s
}

func (s *ListDataServiceApiCallsRequestListQuery) SetAppKey(v int64) *ListDataServiceApiCallsRequestListQuery {
	s.AppKey = &v
	return s
}

func (s *ListDataServiceApiCallsRequestListQuery) SetClientIp(v string) *ListDataServiceApiCallsRequestListQuery {
	s.ClientIp = &v
	return s
}

func (s *ListDataServiceApiCallsRequestListQuery) SetEndTime(v string) *ListDataServiceApiCallsRequestListQuery {
	s.EndTime = &v
	return s
}

func (s *ListDataServiceApiCallsRequestListQuery) SetPageNo(v int32) *ListDataServiceApiCallsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListDataServiceApiCallsRequestListQuery) SetPageSize(v int32) *ListDataServiceApiCallsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApiCallsRequestListQuery) SetStartTime(v string) *ListDataServiceApiCallsRequestListQuery {
	s.StartTime = &v
	return s
}

func (s *ListDataServiceApiCallsRequestListQuery) SetSuccessful(v bool) *ListDataServiceApiCallsRequestListQuery {
	s.Successful = &v
	return s
}

func (s *ListDataServiceApiCallsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
