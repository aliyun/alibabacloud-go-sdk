// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAppsResponseBody
	GetCode() *int32
	SetData(v *ListAppsResponseBodyData) *ListAppsResponseBody
	GetData() *ListAppsResponseBodyData
	SetMessage(v string) *ListAppsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAppsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAppsResponseBody
	GetSuccess() *bool
}

type ListAppsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2C3E52FF-CBE9-5C0E-8252-37ACFF1F5EFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAppsResponseBody) GetData() *ListAppsResponseBodyData {
	return s.Data
}

func (s *ListAppsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAppsResponseBody) SetCode(v int32) *ListAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppsResponseBody) SetData(v *ListAppsResponseBodyData) *ListAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppsResponseBody) SetMessage(v string) *ListAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetSuccess(v bool) *ListAppsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAppsResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListAppsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAppsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppsResponseBodyData) GetRecords() []*ListAppsResponseBodyDataRecords {
	return s.Records
}

func (s *ListAppsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListAppsResponseBodyData) SetPageNumber(v int32) *ListAppsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAppsResponseBodyData) SetPageSize(v int32) *ListAppsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAppsResponseBodyData) SetRecords(v []*ListAppsResponseBodyDataRecords) *ListAppsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListAppsResponseBodyData) SetTotal(v int32) *ListAppsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListAppsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAppsResponseBodyDataRecords struct {
	// AccessToken
	//
	// example:
	//
	// 2f4ddeab8e344ed68e0402cf9b8502ffv3
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 1827811800555555
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// true
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// example:
	//
	// 1
	ExecutorNum *int64 `json:"ExecutorNum,omitempty" xml:"ExecutorNum,omitempty"`
	// example:
	//
	// 43885
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10
	JobNum             *int32 `json:"JobNum,omitempty" xml:"JobNum,omitempty"`
	LabelRouteStrategy *int32 `json:"LabelRouteStrategy,omitempty" xml:"LabelRouteStrategy,omitempty"`
	// example:
	//
	// http://28.5.128.3:80
	Leader *string `json:"Leader,omitempty" xml:"Leader,omitempty"`
	// example:
	//
	// 100
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// example:
	//
	// 1000
	MaxJobs *int32  `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1827811800555555
	Updater *string `json:"Updater,omitempty" xml:"Updater,omitempty"`
}

func (s ListAppsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyDataRecords) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListAppsResponseBodyDataRecords) GetAppName() *string {
	return s.AppName
}

func (s *ListAppsResponseBodyDataRecords) GetAppType() *int32 {
	return s.AppType
}

func (s *ListAppsResponseBodyDataRecords) GetCreator() *string {
	return s.Creator
}

func (s *ListAppsResponseBodyDataRecords) GetEnableLog() *bool {
	return s.EnableLog
}

func (s *ListAppsResponseBodyDataRecords) GetExecutorNum() *int64 {
	return s.ExecutorNum
}

func (s *ListAppsResponseBodyDataRecords) GetId() *int64 {
	return s.Id
}

func (s *ListAppsResponseBodyDataRecords) GetJobNum() *int32 {
	return s.JobNum
}

func (s *ListAppsResponseBodyDataRecords) GetLabelRouteStrategy() *int32 {
	return s.LabelRouteStrategy
}

func (s *ListAppsResponseBodyDataRecords) GetLeader() *string {
	return s.Leader
}

func (s *ListAppsResponseBodyDataRecords) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *ListAppsResponseBodyDataRecords) GetMaxJobs() *int32 {
	return s.MaxJobs
}

func (s *ListAppsResponseBodyDataRecords) GetTitle() *string {
	return s.Title
}

func (s *ListAppsResponseBodyDataRecords) GetUpdater() *string {
	return s.Updater
}

func (s *ListAppsResponseBodyDataRecords) SetAccessToken(v string) *ListAppsResponseBodyDataRecords {
	s.AccessToken = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetAppName(v string) *ListAppsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetAppType(v int32) *ListAppsResponseBodyDataRecords {
	s.AppType = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetCreator(v string) *ListAppsResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetEnableLog(v bool) *ListAppsResponseBodyDataRecords {
	s.EnableLog = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetExecutorNum(v int64) *ListAppsResponseBodyDataRecords {
	s.ExecutorNum = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetId(v int64) *ListAppsResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetJobNum(v int32) *ListAppsResponseBodyDataRecords {
	s.JobNum = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetLabelRouteStrategy(v int32) *ListAppsResponseBodyDataRecords {
	s.LabelRouteStrategy = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetLeader(v string) *ListAppsResponseBodyDataRecords {
	s.Leader = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetMaxConcurrency(v int32) *ListAppsResponseBodyDataRecords {
	s.MaxConcurrency = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetMaxJobs(v int32) *ListAppsResponseBodyDataRecords {
	s.MaxJobs = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetTitle(v string) *ListAppsResponseBodyDataRecords {
	s.Title = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) SetUpdater(v string) *ListAppsResponseBodyDataRecords {
	s.Updater = &v
	return s
}

func (s *ListAppsResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
