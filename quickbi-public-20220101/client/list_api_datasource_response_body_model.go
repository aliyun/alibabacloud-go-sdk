// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListApiDatasourceResponseBody
	GetRequestId() *string
	SetResult(v *ListApiDatasourceResponseBodyResult) *ListApiDatasourceResponseBody
	GetResult() *ListApiDatasourceResponseBodyResult
	SetSuccess(v bool) *ListApiDatasourceResponseBody
	GetSuccess() *bool
}

type ListApiDatasourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The query results are returned.
	Result *ListApiDatasourceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListApiDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiDatasourceResponseBody) GetResult() *ListApiDatasourceResponseBodyResult {
	return s.Result
}

func (s *ListApiDatasourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApiDatasourceResponseBody) SetRequestId(v string) *ListApiDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiDatasourceResponseBody) SetResult(v *ListApiDatasourceResponseBodyResult) *ListApiDatasourceResponseBody {
	s.Result = v
	return s
}

func (s *ListApiDatasourceResponseBody) SetSuccess(v bool) *ListApiDatasourceResponseBody {
	s.Success = &v
	return s
}

func (s *ListApiDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApiDatasourceResponseBodyResult struct {
	// The list of API data sources that were queried.
	Data []*ListApiDatasourceResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page set when the interface is requested.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of rows.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListApiDatasourceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListApiDatasourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceResponseBodyResult) GetData() []*ListApiDatasourceResponseBodyResultData {
	return s.Data
}

func (s *ListApiDatasourceResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListApiDatasourceResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiDatasourceResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListApiDatasourceResponseBodyResult) SetData(v []*ListApiDatasourceResponseBodyResultData) *ListApiDatasourceResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListApiDatasourceResponseBodyResult) SetPageNum(v int32) *ListApiDatasourceResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResult) SetPageSize(v int32) *ListApiDatasourceResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResult) SetTotalNum(v int32) *ListApiDatasourceResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListApiDatasourceResponseBodyResultData struct {
	// The ID of the API data source.
	//
	// example:
	//
	// 0f2c3c6409be4dc0810f2a5785e816a8
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The parameter configuration of the query statement in JSON format. You can customize the parameter configuration.
	//
	// example:
	//
	// {"key1":"value1"}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The data volume of the API data source.
	//
	// 	- Unit: Kbit/s
	//
	// example:
	//
	// 0.39746094
	DataSize *float32 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The last synchronization time of the API data source.
	//
	// example:
	//
	// 2022-05-25 16:19:43
	DateUpdateTime *string `json:"DateUpdateTime,omitempty" xml:"DateUpdateTime,omitempty"`
	// The time when the quota plan was created.
	//
	// example:
	//
	// 2022-05-25 16:19:43
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the optimization job was modified.
	//
	// example:
	//
	// 2022-05-25 16:19:43
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The job ID.
	//
	// example:
	//
	// REST_API_SYNC_0f2c3c6409be4dc0810f2a5785e816a8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The parameter configurations in the JSONArray format.
	//
	// 	- name: parameter name
	//
	// 	- value: the parameter value
	//
	// example:
	//
	// [{"name":"token","value":"xxxxxxxxxxxx"},{"name":"pageSize","value":100}]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The name of the API data source.
	//
	// example:
	//
	// test data source
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The status of the API data source synchronization task.
	//
	// Valid values:
	//
	// 	- 0: the to be run.
	//
	// 	- 1: The is running.
	//
	// 	- 2: The is successfully.
	//
	// 	- 3: failed.
	//
	// example:
	//
	// 2
	StatusType *int32 `json:"StatusType,omitempty" xml:"StatusType,omitempty"`
}

func (s ListApiDatasourceResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListApiDatasourceResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceResponseBodyResultData) GetApiId() *string {
	return s.ApiId
}

func (s *ListApiDatasourceResponseBodyResultData) GetBody() *string {
	return s.Body
}

func (s *ListApiDatasourceResponseBodyResultData) GetDataSize() *float32 {
	return s.DataSize
}

func (s *ListApiDatasourceResponseBodyResultData) GetDateUpdateTime() *string {
	return s.DateUpdateTime
}

func (s *ListApiDatasourceResponseBodyResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListApiDatasourceResponseBodyResultData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListApiDatasourceResponseBodyResultData) GetJobId() *string {
	return s.JobId
}

func (s *ListApiDatasourceResponseBodyResultData) GetParameters() *string {
	return s.Parameters
}

func (s *ListApiDatasourceResponseBodyResultData) GetShowName() *string {
	return s.ShowName
}

func (s *ListApiDatasourceResponseBodyResultData) GetStatusType() *int32 {
	return s.StatusType
}

func (s *ListApiDatasourceResponseBodyResultData) SetApiId(v string) *ListApiDatasourceResponseBodyResultData {
	s.ApiId = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetBody(v string) *ListApiDatasourceResponseBodyResultData {
	s.Body = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetDataSize(v float32) *ListApiDatasourceResponseBodyResultData {
	s.DataSize = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetDateUpdateTime(v string) *ListApiDatasourceResponseBodyResultData {
	s.DateUpdateTime = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetGmtCreate(v string) *ListApiDatasourceResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetGmtModified(v string) *ListApiDatasourceResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetJobId(v string) *ListApiDatasourceResponseBodyResultData {
	s.JobId = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetParameters(v string) *ListApiDatasourceResponseBodyResultData {
	s.Parameters = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetShowName(v string) *ListApiDatasourceResponseBodyResultData {
	s.ShowName = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetStatusType(v int32) *ListApiDatasourceResponseBodyResultData {
	s.StatusType = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}
