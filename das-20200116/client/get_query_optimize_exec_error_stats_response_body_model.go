// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeExecErrorStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueryOptimizeExecErrorStatsResponseBody
	GetCode() *string
	SetData(v *GetQueryOptimizeExecErrorStatsResponseBodyData) *GetQueryOptimizeExecErrorStatsResponseBody
	GetData() *GetQueryOptimizeExecErrorStatsResponseBodyData
	SetMessage(v string) *GetQueryOptimizeExecErrorStatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueryOptimizeExecErrorStatsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueryOptimizeExecErrorStatsResponseBody
	GetSuccess() *string
}

type GetQueryOptimizeExecErrorStatsResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *GetQueryOptimizeExecErrorStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeExecErrorStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) GetData() *GetQueryOptimizeExecErrorStatsResponseBodyData {
	return s.Data
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetCode(v string) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetData(v *GetQueryOptimizeExecErrorStatsResponseBodyData) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetMessage(v string) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetRequestId(v string) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) SetSuccess(v string) *GetQueryOptimizeExecErrorStatsResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQueryOptimizeExecErrorStatsResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The information about the SQL templates that failed to execute.
	List []*GetQueryOptimizeExecErrorStatsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 19
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeExecErrorStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) GetList() []*GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	return s.List
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetExtra(v string) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetList(v []*GetQueryOptimizeExecErrorStatsResponseBodyDataList) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) SetTotal(v int64) *GetQueryOptimizeExecErrorStatsResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetQueryOptimizeExecErrorStatsResponseBodyDataList struct {
	// The name of the database.
	//
	// example:
	//
	// testdb01
	Dbname *string `json:"Dbname,omitempty" xml:"Dbname,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 1146
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The number of errors.
	//
	// example:
	//
	// 10
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The alias of the database instance.
	//
	// example:
	//
	// test01
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The SQL template ID.
	//
	// example:
	//
	// 2e8147b5ca2dfc640dfd5e43d96a****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The content of the SQL template.
	//
	// example:
	//
	// select 	- from test1
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
}

func (s GetQueryOptimizeExecErrorStatsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorStatsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) GetDbname() *string {
	return s.Dbname
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) GetSqlId() *string {
	return s.SqlId
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) GetSqlText() *string {
	return s.SqlText
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetDbname(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.Dbname = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetErrorCode(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.ErrorCode = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetErrorCount(v int64) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.ErrorCount = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetInstanceId(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetInstanceName(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.InstanceName = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetSqlId(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) SetSqlText(v string) *GetQueryOptimizeExecErrorStatsResponseBodyDataList {
	s.SqlText = &v
	return s
}

func (s *GetQueryOptimizeExecErrorStatsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
