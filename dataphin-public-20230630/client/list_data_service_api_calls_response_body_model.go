// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiCallsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataServiceApiCallsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListDataServiceApiCallsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataServiceApiCallsResponseBody
	GetMessage() *string
	SetPageResult(v *ListDataServiceApiCallsResponseBodyPageResult) *ListDataServiceApiCallsResponseBody
	GetPageResult() *ListDataServiceApiCallsResponseBodyPageResult
	SetRequestId(v string) *ListDataServiceApiCallsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceApiCallsResponseBody
	GetSuccess() *bool
}

type ListDataServiceApiCallsResponseBody struct {
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
	Message    *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListDataServiceApiCallsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServiceApiCallsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataServiceApiCallsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceApiCallsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataServiceApiCallsResponseBody) GetPageResult() *ListDataServiceApiCallsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListDataServiceApiCallsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceApiCallsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceApiCallsResponseBody) SetCode(v string) *ListDataServiceApiCallsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBody) SetHttpStatusCode(v int32) *ListDataServiceApiCallsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBody) SetMessage(v string) *ListDataServiceApiCallsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBody) SetPageResult(v *ListDataServiceApiCallsResponseBodyPageResult) *ListDataServiceApiCallsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDataServiceApiCallsResponseBody) SetRequestId(v string) *ListDataServiceApiCallsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBody) SetSuccess(v bool) *ListDataServiceApiCallsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServiceApiCallsResponseBodyPageResult struct {
	CallLogList []*ListDataServiceApiCallsResponseBodyPageResultCallLogList `json:"CallLogList,omitempty" xml:"CallLogList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceApiCallsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallsResponseBodyPageResult) GetCallLogList() []*ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	return s.CallLogList
}

func (s *ListDataServiceApiCallsResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataServiceApiCallsResponseBodyPageResult) SetCallLogList(v []*ListDataServiceApiCallsResponseBodyPageResultCallLogList) *ListDataServiceApiCallsResponseBodyPageResult {
	s.CallLogList = v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResult) SetTotalCount(v int64) *ListDataServiceApiCallsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResult) Validate() error {
	if s.CallLogList != nil {
		for _, item := range s.CallLogList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServiceApiCallsResponseBodyPageResultCallLogList struct {
	// example:
	//
	// 102112
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// test
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// 201211
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// OK
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// OK
	BizCodeDescription *string `json:"BizCodeDescription,omitempty" xml:"BizCodeDescription,omitempty"`
	// example:
	//
	// 192.168.1.1
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// 2000
	CostTime *int32 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	// example:
	//
	// 2024-11-01 01:01:03.000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Env *int32 `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 1000
	ExecuteCostTime *int32 `json:"ExecuteCostTime,omitempty" xml:"ExecuteCostTime,omitempty"`
	// example:
	//
	// 1
	ExecuteMode *int32 `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	HttpStatusDescription *string `json:"HttpStatusDescription,omitempty" xml:"HttpStatusDescription,omitempty"`
	// example:
	//
	// 102356
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 102356
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1234567890-232sds-3e232-ae2e232
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {"name":"test"}
	RequestParameter *string `json:"RequestParameter,omitempty" xml:"RequestParameter,omitempty"`
	// example:
	//
	// 1024
	RequestSize *int32 `json:"RequestSize,omitempty" xml:"RequestSize,omitempty"`
	// example:
	//
	// {"code":"200","message":"success"}
	ResponseParameter *string `json:"ResponseParameter,omitempty" xml:"ResponseParameter,omitempty"`
	// example:
	//
	// 1024
	ResponseSize *int32 `json:"ResponseSize,omitempty" xml:"ResponseSize,omitempty"`
	// example:
	//
	// 1
	ResultCount *int32 `json:"ResultCount,omitempty" xml:"ResultCount,omitempty"`
	// example:
	//
	// select col1 from t_test1 limit 100;
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// example:
	//
	// 2024-11-01 01:01:01.000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status     *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	Successful *bool  `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDataServiceApiCallsResponseBodyPageResultCallLogList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallsResponseBodyPageResultCallLogList) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetAppKey() *int64 {
	return s.AppKey
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetAppName() *string {
	return s.AppName
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetBizCode() *string {
	return s.BizCode
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetBizCodeDescription() *string {
	return s.BizCodeDescription
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetCostTime() *int32 {
	return s.CostTime
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetEnv() *int32 {
	return s.Env
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetExecuteCostTime() *int32 {
	return s.ExecuteCostTime
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetExecuteMode() *int32 {
	return s.ExecuteMode
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetHttpStatusDescription() *string {
	return s.HttpStatusDescription
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetJobId() *string {
	return s.JobId
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetRequestParameter() *string {
	return s.RequestParameter
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetRequestSize() *int32 {
	return s.RequestSize
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetResponseParameter() *string {
	return s.ResponseParameter
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetResponseSize() *int32 {
	return s.ResponseSize
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetResultCount() *int32 {
	return s.ResultCount
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetSql() *string {
	return s.Sql
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetStatus() *int32 {
	return s.Status
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) GetSuccessful() *bool {
	return s.Successful
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetApiId(v int64) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetApiName(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ApiName = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetAppKey(v int64) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.AppKey = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetAppName(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.AppName = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetBizCode(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.BizCode = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetBizCodeDescription(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.BizCodeDescription = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetClientIp(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ClientIp = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetCostTime(v int32) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.CostTime = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetEndTime(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.EndTime = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetEnv(v int32) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.Env = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetExecuteCostTime(v int32) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ExecuteCostTime = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetExecuteMode(v int32) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ExecuteMode = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetHttpStatusCode(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetHttpStatusDescription(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.HttpStatusDescription = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetJobId(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.JobId = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetProjectId(v int32) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetProjectName(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ProjectName = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetRequestId(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetRequestParameter(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.RequestParameter = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetRequestSize(v int32) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.RequestSize = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetResponseParameter(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ResponseParameter = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetResponseSize(v int32) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ResponseSize = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetResultCount(v int32) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.ResultCount = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetSql(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.Sql = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetStartTime(v string) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.StartTime = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetStatus(v int32) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.Status = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) SetSuccessful(v bool) *ListDataServiceApiCallsResponseBodyPageResultCallLogList {
	s.Successful = &v
	return s
}

func (s *ListDataServiceApiCallsResponseBodyPageResultCallLogList) Validate() error {
	return dara.Validate(s)
}
