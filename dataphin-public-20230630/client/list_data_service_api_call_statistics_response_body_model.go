// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiCallStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataServiceApiCallStatisticsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListDataServiceApiCallStatisticsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataServiceApiCallStatisticsResponseBody
	GetMessage() *string
	SetPageResult(v *ListDataServiceApiCallStatisticsResponseBodyPageResult) *ListDataServiceApiCallStatisticsResponseBody
	GetPageResult() *ListDataServiceApiCallStatisticsResponseBodyPageResult
	SetRequestId(v string) *ListDataServiceApiCallStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceApiCallStatisticsResponseBody
	GetSuccess() *bool
}

type ListDataServiceApiCallStatisticsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message    *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListDataServiceApiCallStatisticsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServiceApiCallStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataServiceApiCallStatisticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceApiCallStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataServiceApiCallStatisticsResponseBody) GetPageResult() *ListDataServiceApiCallStatisticsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListDataServiceApiCallStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceApiCallStatisticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceApiCallStatisticsResponseBody) SetCode(v string) *ListDataServiceApiCallStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBody) SetHttpStatusCode(v int32) *ListDataServiceApiCallStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBody) SetMessage(v string) *ListDataServiceApiCallStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBody) SetPageResult(v *ListDataServiceApiCallStatisticsResponseBodyPageResult) *ListDataServiceApiCallStatisticsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBody) SetRequestId(v string) *ListDataServiceApiCallStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBody) SetSuccess(v bool) *ListDataServiceApiCallStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServiceApiCallStatisticsResponseBodyPageResult struct {
	CallStatisticsList []*ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList `json:"CallStatisticsList,omitempty" xml:"CallStatisticsList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceApiCallStatisticsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallStatisticsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResult) GetCallStatisticsList() []*ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	return s.CallStatisticsList
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResult) SetCallStatisticsList(v []*ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) *ListDataServiceApiCallStatisticsResponseBodyPageResult {
	s.CallStatisticsList = v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResult) SetTotalCount(v int32) *ListDataServiceApiCallStatisticsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResult) Validate() error {
	if s.CallStatisticsList != nil {
		for _, item := range s.CallStatisticsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList struct {
	// example:
	//
	// 1003
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// test
	ApiName     *string   `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AppNameList []*string `json:"AppNameList,omitempty" xml:"AppNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	AuthorizedAppCount *int32 `json:"AuthorizedAppCount,omitempty" xml:"AuthorizedAppCount,omitempty"`
	// example:
	//
	// 11
	AvgResponseTime *float64 `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	// example:
	//
	// 100
	CallCount *int64 `json:"CallCount,omitempty" xml:"CallCount,omitempty"`
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 22
	ErrorCount *string `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// example:
	//
	// 23.2%
	ErrorRate *string `json:"ErrorRate,omitempty" xml:"ErrorRate,omitempty"`
	// example:
	//
	// 2025-06-30 08:00:00
	LastCallTime *string `json:"LastCallTime,omitempty" xml:"LastCallTime,omitempty"`
	// example:
	//
	// 23.2%
	OfflineRate *string `json:"OfflineRate,omitempty" xml:"OfflineRate,omitempty"`
	// example:
	//
	// 101201
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1121
	SqlId *int32 `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
}

func (s ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetAppNameList() []*string {
	return s.AppNameList
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetAuthorizedAppCount() *int32 {
	return s.AuthorizedAppCount
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetAvgResponseTime() *float64 {
	return s.AvgResponseTime
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetCallCount() *int64 {
	return s.CallCount
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetCreator() *string {
	return s.Creator
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetErrorCount() *string {
	return s.ErrorCount
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetErrorRate() *string {
	return s.ErrorRate
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetLastCallTime() *string {
	return s.LastCallTime
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetOfflineRate() *string {
	return s.OfflineRate
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) GetSqlId() *int32 {
	return s.SqlId
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetApiId(v int64) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetApiName(v string) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.ApiName = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetAppNameList(v []*string) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.AppNameList = v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetAuthorizedAppCount(v int32) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.AuthorizedAppCount = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetAvgResponseTime(v float64) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.AvgResponseTime = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetCallCount(v int64) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.CallCount = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetCreator(v string) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.Creator = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetErrorCount(v string) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.ErrorCount = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetErrorRate(v string) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.ErrorRate = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetLastCallTime(v string) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.LastCallTime = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetOfflineRate(v string) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.OfflineRate = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetProjectId(v int32) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetProjectName(v string) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.ProjectName = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) SetSqlId(v int32) *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList {
	s.SqlId = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponseBodyPageResultCallStatisticsList) Validate() error {
	return dara.Validate(s)
}
