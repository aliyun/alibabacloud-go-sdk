// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type DeleteSymRecordsRequest struct {
	// This parameter is required.
	AppVersions []*string `json:"appVersions,omitempty" xml:"appVersions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	FileType *int32 `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s DeleteSymRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSymRecordsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSymRecordsRequest) SetAppVersions(v []*string) *DeleteSymRecordsRequest {
	s.AppVersions = v
	return s
}

func (s *DeleteSymRecordsRequest) SetDataSourceId(v string) *DeleteSymRecordsRequest {
	s.DataSourceId = &v
	return s
}

func (s *DeleteSymRecordsRequest) SetFileType(v int32) *DeleteSymRecordsRequest {
	s.FileType = &v
	return s
}

type DeleteSymRecordsShrinkRequest struct {
	// This parameter is required.
	AppVersionsShrink *string `json:"appVersions,omitempty" xml:"appVersions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	FileType *int32 `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s DeleteSymRecordsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSymRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSymRecordsShrinkRequest) SetAppVersionsShrink(v string) *DeleteSymRecordsShrinkRequest {
	s.AppVersionsShrink = &v
	return s
}

func (s *DeleteSymRecordsShrinkRequest) SetDataSourceId(v string) *DeleteSymRecordsShrinkRequest {
	s.DataSourceId = &v
	return s
}

func (s *DeleteSymRecordsShrinkRequest) SetFileType(v int32) *DeleteSymRecordsShrinkRequest {
	s.FileType = &v
	return s
}

type DeleteSymRecordsResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// 1
	Num *int32 `json:"num,omitempty" xml:"num,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f07c516457690916816858d94ea
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeleteSymRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSymRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSymRecordsResponseBody) SetCode(v int64) *DeleteSymRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSymRecordsResponseBody) SetMsg(v string) *DeleteSymRecordsResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteSymRecordsResponseBody) SetNum(v int32) *DeleteSymRecordsResponseBody {
	s.Num = &v
	return s
}

func (s *DeleteSymRecordsResponseBody) SetSuccess(v bool) *DeleteSymRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSymRecordsResponseBody) SetTraceId(v string) *DeleteSymRecordsResponseBody {
	s.TraceId = &v
	return s
}

type DeleteSymRecordsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSymRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSymRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSymRecordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSymRecordsResponse) SetHeaders(v map[string]*string) *DeleteSymRecordsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSymRecordsResponse) SetStatusCode(v int32) *DeleteSymRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSymRecordsResponse) SetBody(v *DeleteSymRecordsResponseBody) *DeleteSymRecordsResponse {
	s.Body = v
	return s
}

type GetErrorMinuteStatTrendRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-27 09:07
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetErrorMinuteStatTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetErrorMinuteStatTrendRequest) GoString() string {
	return s.String()
}

func (s *GetErrorMinuteStatTrendRequest) SetDataSourceId(v string) *GetErrorMinuteStatTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetErrorMinuteStatTrendRequest) SetStartTime(v string) *GetErrorMinuteStatTrendRequest {
	s.StartTime = &v
	return s
}

func (s *GetErrorMinuteStatTrendRequest) SetType(v int32) *GetErrorMinuteStatTrendRequest {
	s.Type = &v
	return s
}

type GetErrorMinuteStatTrendResponseBody struct {
	// example:
	//
	// 200
	Code *int64                                     `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetErrorMinuteStatTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetErrorMinuteStatTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErrorMinuteStatTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetErrorMinuteStatTrendResponseBody) SetCode(v int64) *GetErrorMinuteStatTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetErrorMinuteStatTrendResponseBody) SetData(v []*GetErrorMinuteStatTrendResponseBodyData) *GetErrorMinuteStatTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetErrorMinuteStatTrendResponseBody) SetMsg(v string) *GetErrorMinuteStatTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetErrorMinuteStatTrendResponseBody) SetSuccess(v bool) *GetErrorMinuteStatTrendResponseBody {
	s.Success = &v
	return s
}

type GetErrorMinuteStatTrendResponseBodyData struct {
	// example:
	//
	// 120
	ErrorCount *int64 `json:"errorCount,omitempty" xml:"errorCount,omitempty"`
	// example:
	//
	// 1200
	LaunchCount *int64 `json:"launchCount,omitempty" xml:"launchCount,omitempty"`
	// example:
	//
	// 2023-05-20 13:01
	TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
}

func (s GetErrorMinuteStatTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetErrorMinuteStatTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetErrorMinuteStatTrendResponseBodyData) SetErrorCount(v int64) *GetErrorMinuteStatTrendResponseBodyData {
	s.ErrorCount = &v
	return s
}

func (s *GetErrorMinuteStatTrendResponseBodyData) SetLaunchCount(v int64) *GetErrorMinuteStatTrendResponseBodyData {
	s.LaunchCount = &v
	return s
}

func (s *GetErrorMinuteStatTrendResponseBodyData) SetTimePoint(v string) *GetErrorMinuteStatTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

type GetErrorMinuteStatTrendResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErrorMinuteStatTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErrorMinuteStatTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetErrorMinuteStatTrendResponse) GoString() string {
	return s.String()
}

func (s *GetErrorMinuteStatTrendResponse) SetHeaders(v map[string]*string) *GetErrorMinuteStatTrendResponse {
	s.Headers = v
	return s
}

func (s *GetErrorMinuteStatTrendResponse) SetStatusCode(v int32) *GetErrorMinuteStatTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErrorMinuteStatTrendResponse) SetBody(v *GetErrorMinuteStatTrendResponseBody) *GetErrorMinuteStatTrendResponse {
	s.Body = v
	return s
}

type GetH5PageTrendRequest struct {
	// example:
	//
	// 1.0.2
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-03
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// day
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s GetH5PageTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetH5PageTrendRequest) GoString() string {
	return s.String()
}

func (s *GetH5PageTrendRequest) SetAppVersion(v string) *GetH5PageTrendRequest {
	s.AppVersion = &v
	return s
}

func (s *GetH5PageTrendRequest) SetDataSourceId(v string) *GetH5PageTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetH5PageTrendRequest) SetEndDate(v string) *GetH5PageTrendRequest {
	s.EndDate = &v
	return s
}

func (s *GetH5PageTrendRequest) SetStartDate(v string) *GetH5PageTrendRequest {
	s.StartDate = &v
	return s
}

func (s *GetH5PageTrendRequest) SetTimeUnit(v string) *GetH5PageTrendRequest {
	s.TimeUnit = &v
	return s
}

type GetH5PageTrendResponseBody struct {
	// example:
	//
	// 200
	Code *int64                            `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetH5PageTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetH5PageTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetH5PageTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetH5PageTrendResponseBody) SetCode(v int64) *GetH5PageTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetH5PageTrendResponseBody) SetData(v []*GetH5PageTrendResponseBodyData) *GetH5PageTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetH5PageTrendResponseBody) SetMsg(v string) *GetH5PageTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetH5PageTrendResponseBody) SetSuccess(v bool) *GetH5PageTrendResponseBody {
	s.Success = &v
	return s
}

type GetH5PageTrendResponseBodyData struct {
	// example:
	//
	// 504.89
	AnalyzeDOM *float64 `json:"analyzeDOM,omitempty" xml:"analyzeDOM,omitempty"`
	// example:
	//
	// 49.48
	AppCache *float64 `json:"appCache,omitempty" xml:"appCache,omitempty"`
	// example:
	//
	// 979.83
	ContentTrans *float64 `json:"contentTrans,omitempty" xml:"contentTrans,omitempty"`
	// example:
	//
	// 50.16
	Dns *float64 `json:"dns,omitempty" xml:"dns,omitempty"`
	// example:
	//
	// 1881.96
	DomReady *float64 `json:"domReady,omitempty" xml:"domReady,omitempty"`
	// example:
	//
	// 190.69
	Fcp *float64 `json:"fcp,omitempty" xml:"fcp,omitempty"`
	// example:
	//
	// 472.57
	FirstByte *float64 `json:"firstByte,omitempty" xml:"firstByte,omitempty"`
	// example:
	//
	// 44.67
	FiveSecondRate *float64 `json:"fiveSecondRate,omitempty" xml:"fiveSecondRate,omitempty"`
	// example:
	//
	// 50.25
	Fp *float64 `json:"fp,omitempty" xml:"fp,omitempty"`
	// example:
	//
	// 492.86
	LoadEvent *float64 `json:"loadEvent,omitempty" xml:"loadEvent,omitempty"`
	// example:
	//
	// 4741.44
	LoadFinish *float64 `json:"loadFinish,omitempty" xml:"loadFinish,omitempty"`
	// example:
	//
	// 2549.46
	LoadResource *float64 `json:"loadResource,omitempty" xml:"loadResource,omitempty"`
	// example:
	//
	// 2062
	LogCnt *int64 `json:"logCnt,omitempty" xml:"logCnt,omitempty"`
	// example:
	//
	// 0.19
	OneSecondRate *float64 `json:"oneSecondRate,omitempty" xml:"oneSecondRate,omitempty"`
	// example:
	//
	// 100.93
	Redirect *float64 `json:"redirect,omitempty" xml:"redirect,omitempty"`
	// example:
	//
	// 71.02
	Ssl *float64 `json:"ssl,omitempty" xml:"ssl,omitempty"`
	// example:
	//
	// 150.18
	Tcp *float64 `json:"tcp,omitempty" xml:"tcp,omitempty"`
	// example:
	//
	// 2023-05-20
	TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
	// example:
	//
	// 249.55
	Ttfb *float64 `json:"ttfb,omitempty" xml:"ttfb,omitempty"`
	// example:
	//
	// 2126.61
	Tti *float64 `json:"tti,omitempty" xml:"tti,omitempty"`
	// example:
	//
	// 4.9
	TwoSecondRate *float64 `json:"twoSecondRate,omitempty" xml:"twoSecondRate,omitempty"`
	// example:
	//
	// 98.26
	Unload *float64 `json:"unload,omitempty" xml:"unload,omitempty"`
}

func (s GetH5PageTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetH5PageTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetH5PageTrendResponseBodyData) SetAnalyzeDOM(v float64) *GetH5PageTrendResponseBodyData {
	s.AnalyzeDOM = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetAppCache(v float64) *GetH5PageTrendResponseBodyData {
	s.AppCache = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetContentTrans(v float64) *GetH5PageTrendResponseBodyData {
	s.ContentTrans = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetDns(v float64) *GetH5PageTrendResponseBodyData {
	s.Dns = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetDomReady(v float64) *GetH5PageTrendResponseBodyData {
	s.DomReady = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetFcp(v float64) *GetH5PageTrendResponseBodyData {
	s.Fcp = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetFirstByte(v float64) *GetH5PageTrendResponseBodyData {
	s.FirstByte = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetFiveSecondRate(v float64) *GetH5PageTrendResponseBodyData {
	s.FiveSecondRate = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetFp(v float64) *GetH5PageTrendResponseBodyData {
	s.Fp = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetLoadEvent(v float64) *GetH5PageTrendResponseBodyData {
	s.LoadEvent = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetLoadFinish(v float64) *GetH5PageTrendResponseBodyData {
	s.LoadFinish = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetLoadResource(v float64) *GetH5PageTrendResponseBodyData {
	s.LoadResource = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetLogCnt(v int64) *GetH5PageTrendResponseBodyData {
	s.LogCnt = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetOneSecondRate(v float64) *GetH5PageTrendResponseBodyData {
	s.OneSecondRate = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetRedirect(v float64) *GetH5PageTrendResponseBodyData {
	s.Redirect = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetSsl(v float64) *GetH5PageTrendResponseBodyData {
	s.Ssl = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetTcp(v float64) *GetH5PageTrendResponseBodyData {
	s.Tcp = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetTimePoint(v string) *GetH5PageTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetTtfb(v float64) *GetH5PageTrendResponseBodyData {
	s.Ttfb = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetTti(v float64) *GetH5PageTrendResponseBodyData {
	s.Tti = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetTwoSecondRate(v float64) *GetH5PageTrendResponseBodyData {
	s.TwoSecondRate = &v
	return s
}

func (s *GetH5PageTrendResponseBodyData) SetUnload(v float64) *GetH5PageTrendResponseBodyData {
	s.Unload = &v
	return s
}

type GetH5PageTrendResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetH5PageTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetH5PageTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetH5PageTrendResponse) GoString() string {
	return s.String()
}

func (s *GetH5PageTrendResponse) SetHeaders(v map[string]*string) *GetH5PageTrendResponse {
	s.Headers = v
	return s
}

func (s *GetH5PageTrendResponse) SetStatusCode(v int32) *GetH5PageTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetH5PageTrendResponse) SetBody(v *GetH5PageTrendResponseBody) *GetH5PageTrendResponse {
	s.Body = v
	return s
}

type GetLaunchTrendRequest struct {
	// example:
	//
	// 1.0.2
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-03
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// day
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s GetLaunchTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLaunchTrendRequest) GoString() string {
	return s.String()
}

func (s *GetLaunchTrendRequest) SetAppVersion(v string) *GetLaunchTrendRequest {
	s.AppVersion = &v
	return s
}

func (s *GetLaunchTrendRequest) SetDataSourceId(v string) *GetLaunchTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetLaunchTrendRequest) SetEndDate(v string) *GetLaunchTrendRequest {
	s.EndDate = &v
	return s
}

func (s *GetLaunchTrendRequest) SetStartDate(v string) *GetLaunchTrendRequest {
	s.StartDate = &v
	return s
}

func (s *GetLaunchTrendRequest) SetTimeUnit(v string) *GetLaunchTrendRequest {
	s.TimeUnit = &v
	return s
}

type GetLaunchTrendResponseBody struct {
	// example:
	//
	// 200
	Code *int64                            `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetLaunchTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetLaunchTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLaunchTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetLaunchTrendResponseBody) SetCode(v int64) *GetLaunchTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetLaunchTrendResponseBody) SetData(v []*GetLaunchTrendResponseBodyData) *GetLaunchTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetLaunchTrendResponseBody) SetMsg(v string) *GetLaunchTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetLaunchTrendResponseBody) SetSuccess(v bool) *GetLaunchTrendResponseBody {
	s.Success = &v
	return s
}

type GetLaunchTrendResponseBodyData struct {
	// example:
	//
	// 2495
	ColdLaunchCount *int64 `json:"coldLaunchCount,omitempty" xml:"coldLaunchCount,omitempty"`
	// example:
	//
	// 3784.5
	ColdLaunchDuration *float64 `json:"coldLaunchDuration,omitempty" xml:"coldLaunchDuration,omitempty"`
	// example:
	//
	// 2495
	FirstLaunchCount *int64 `json:"firstLaunchCount,omitempty" xml:"firstLaunchCount,omitempty"`
	// example:
	//
	// 3740.5
	FirstLaunchDuration *float64 `json:"firstLaunchDuration,omitempty" xml:"firstLaunchDuration,omitempty"`
	// example:
	//
	// 2495
	HotLaunchCount *int64 `json:"hotLaunchCount,omitempty" xml:"hotLaunchCount,omitempty"`
	// example:
	//
	// 1400.5
	HotLaunchDuration *float64 `json:"hotLaunchDuration,omitempty" xml:"hotLaunchDuration,omitempty"`
	// example:
	//
	// 2023-05-20
	TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
}

func (s GetLaunchTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLaunchTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLaunchTrendResponseBodyData) SetColdLaunchCount(v int64) *GetLaunchTrendResponseBodyData {
	s.ColdLaunchCount = &v
	return s
}

func (s *GetLaunchTrendResponseBodyData) SetColdLaunchDuration(v float64) *GetLaunchTrendResponseBodyData {
	s.ColdLaunchDuration = &v
	return s
}

func (s *GetLaunchTrendResponseBodyData) SetFirstLaunchCount(v int64) *GetLaunchTrendResponseBodyData {
	s.FirstLaunchCount = &v
	return s
}

func (s *GetLaunchTrendResponseBodyData) SetFirstLaunchDuration(v float64) *GetLaunchTrendResponseBodyData {
	s.FirstLaunchDuration = &v
	return s
}

func (s *GetLaunchTrendResponseBodyData) SetHotLaunchCount(v int64) *GetLaunchTrendResponseBodyData {
	s.HotLaunchCount = &v
	return s
}

func (s *GetLaunchTrendResponseBodyData) SetHotLaunchDuration(v float64) *GetLaunchTrendResponseBodyData {
	s.HotLaunchDuration = &v
	return s
}

func (s *GetLaunchTrendResponseBodyData) SetTimePoint(v string) *GetLaunchTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

type GetLaunchTrendResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLaunchTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLaunchTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLaunchTrendResponse) GoString() string {
	return s.String()
}

func (s *GetLaunchTrendResponse) SetHeaders(v map[string]*string) *GetLaunchTrendResponse {
	s.Headers = v
	return s
}

func (s *GetLaunchTrendResponse) SetStatusCode(v int32) *GetLaunchTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLaunchTrendResponse) SetBody(v *GetLaunchTrendResponseBody) *GetLaunchTrendResponse {
	s.Body = v
	return s
}

type GetNativePageTrendRequest struct {
	// example:
	//
	// 1.0.2
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-03
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// day
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s GetNativePageTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNativePageTrendRequest) GoString() string {
	return s.String()
}

func (s *GetNativePageTrendRequest) SetAppVersion(v string) *GetNativePageTrendRequest {
	s.AppVersion = &v
	return s
}

func (s *GetNativePageTrendRequest) SetDataSourceId(v string) *GetNativePageTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetNativePageTrendRequest) SetEndDate(v string) *GetNativePageTrendRequest {
	s.EndDate = &v
	return s
}

func (s *GetNativePageTrendRequest) SetStartDate(v string) *GetNativePageTrendRequest {
	s.StartDate = &v
	return s
}

func (s *GetNativePageTrendRequest) SetTimeUnit(v string) *GetNativePageTrendRequest {
	s.TimeUnit = &v
	return s
}

type GetNativePageTrendResponseBody struct {
	// example:
	//
	// 200
	Code *int64                                `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetNativePageTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetNativePageTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNativePageTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetNativePageTrendResponseBody) SetCode(v int64) *GetNativePageTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetNativePageTrendResponseBody) SetData(v []*GetNativePageTrendResponseBodyData) *GetNativePageTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetNativePageTrendResponseBody) SetMsg(v string) *GetNativePageTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetNativePageTrendResponseBody) SetSuccess(v bool) *GetNativePageTrendResponseBody {
	s.Success = &v
	return s
}

type GetNativePageTrendResponseBodyData struct {
	// example:
	//
	// 75.9
	AvgLoadDuration *float64 `json:"avgLoadDuration,omitempty" xml:"avgLoadDuration,omitempty"`
	// example:
	//
	// 37.317
	CrashRate *float64 `json:"crashRate,omitempty" xml:"crashRate,omitempty"`
	// example:
	//
	// 2460
	LoadCnt *int64 `json:"loadCnt,omitempty" xml:"loadCnt,omitempty"`
	// example:
	//
	// 99.837
	SlowLoadRate *float64 `json:"slowLoadRate,omitempty" xml:"slowLoadRate,omitempty"`
	// example:
	//
	// 2023-05-20
	TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
}

func (s GetNativePageTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNativePageTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNativePageTrendResponseBodyData) SetAvgLoadDuration(v float64) *GetNativePageTrendResponseBodyData {
	s.AvgLoadDuration = &v
	return s
}

func (s *GetNativePageTrendResponseBodyData) SetCrashRate(v float64) *GetNativePageTrendResponseBodyData {
	s.CrashRate = &v
	return s
}

func (s *GetNativePageTrendResponseBodyData) SetLoadCnt(v int64) *GetNativePageTrendResponseBodyData {
	s.LoadCnt = &v
	return s
}

func (s *GetNativePageTrendResponseBodyData) SetSlowLoadRate(v float64) *GetNativePageTrendResponseBodyData {
	s.SlowLoadRate = &v
	return s
}

func (s *GetNativePageTrendResponseBodyData) SetTimePoint(v string) *GetNativePageTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

type GetNativePageTrendResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNativePageTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNativePageTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNativePageTrendResponse) GoString() string {
	return s.String()
}

func (s *GetNativePageTrendResponse) SetHeaders(v map[string]*string) *GetNativePageTrendResponse {
	s.Headers = v
	return s
}

func (s *GetNativePageTrendResponse) SetStatusCode(v int32) *GetNativePageTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNativePageTrendResponse) SetBody(v *GetNativePageTrendResponseBody) *GetNativePageTrendResponse {
	s.Body = v
	return s
}

type GetNetworkMinuteTrendRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-27 09:07
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetNetworkMinuteTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkMinuteTrendRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkMinuteTrendRequest) SetDataSourceId(v string) *GetNetworkMinuteTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetNetworkMinuteTrendRequest) SetStartTime(v string) *GetNetworkMinuteTrendRequest {
	s.StartTime = &v
	return s
}

type GetNetworkMinuteTrendResponseBody struct {
	// example:
	//
	// 200
	Code *int64                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetNetworkMinuteTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetNetworkMinuteTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkMinuteTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkMinuteTrendResponseBody) SetCode(v int64) *GetNetworkMinuteTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetNetworkMinuteTrendResponseBody) SetData(v []*GetNetworkMinuteTrendResponseBodyData) *GetNetworkMinuteTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetNetworkMinuteTrendResponseBody) SetMsg(v string) *GetNetworkMinuteTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetNetworkMinuteTrendResponseBody) SetSuccess(v bool) *GetNetworkMinuteTrendResponseBody {
	s.Success = &v
	return s
}

type GetNetworkMinuteTrendResponseBodyData struct {
	// example:
	//
	// 120
	ErrorCount *int64 `json:"errorCount,omitempty" xml:"errorCount,omitempty"`
	// example:
	//
	// 1200
	RequestCount *int64 `json:"requestCount,omitempty" xml:"requestCount,omitempty"`
	// example:
	//
	// 2023-05-20 09:08
	TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
}

func (s GetNetworkMinuteTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkMinuteTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNetworkMinuteTrendResponseBodyData) SetErrorCount(v int64) *GetNetworkMinuteTrendResponseBodyData {
	s.ErrorCount = &v
	return s
}

func (s *GetNetworkMinuteTrendResponseBodyData) SetRequestCount(v int64) *GetNetworkMinuteTrendResponseBodyData {
	s.RequestCount = &v
	return s
}

func (s *GetNetworkMinuteTrendResponseBodyData) SetTimePoint(v string) *GetNetworkMinuteTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

type GetNetworkMinuteTrendResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkMinuteTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkMinuteTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkMinuteTrendResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkMinuteTrendResponse) SetHeaders(v map[string]*string) *GetNetworkMinuteTrendResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkMinuteTrendResponse) SetStatusCode(v int32) *GetNetworkMinuteTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkMinuteTrendResponse) SetBody(v *GetNetworkMinuteTrendResponseBody) *GetNetworkMinuteTrendResponse {
	s.Body = v
	return s
}

type GetNetworkTrendRequest struct {
	// example:
	//
	// 1.0.2
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-03
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// day
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s GetNetworkTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkTrendRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkTrendRequest) SetAppVersion(v string) *GetNetworkTrendRequest {
	s.AppVersion = &v
	return s
}

func (s *GetNetworkTrendRequest) SetDataSourceId(v string) *GetNetworkTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetNetworkTrendRequest) SetEndDate(v string) *GetNetworkTrendRequest {
	s.EndDate = &v
	return s
}

func (s *GetNetworkTrendRequest) SetStartDate(v string) *GetNetworkTrendRequest {
	s.StartDate = &v
	return s
}

func (s *GetNetworkTrendRequest) SetTimeUnit(v string) *GetNetworkTrendRequest {
	s.TimeUnit = &v
	return s
}

type GetNetworkTrendResponseBody struct {
	// example:
	//
	// 200
	Code *int64                             `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetNetworkTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetNetworkTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkTrendResponseBody) SetCode(v int64) *GetNetworkTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetNetworkTrendResponseBody) SetData(v []*GetNetworkTrendResponseBodyData) *GetNetworkTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetNetworkTrendResponseBody) SetMsg(v string) *GetNetworkTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetNetworkTrendResponseBody) SetSuccess(v bool) *GetNetworkTrendResponseBody {
	s.Success = &v
	return s
}

type GetNetworkTrendResponseBodyData struct {
	// example:
	//
	// 4402.8
	AvgCost *float64 `json:"avgCost,omitempty" xml:"avgCost,omitempty"`
	// example:
	//
	// 1654.51
	AvgResponseTime *float64 `json:"avgResponseTime,omitempty" xml:"avgResponseTime,omitempty"`
	// example:
	//
	// 3299.43
	AvgTransformBytes *float64 `json:"avgTransformBytes,omitempty" xml:"avgTransformBytes,omitempty"`
	// example:
	//
	// 1.61
	RequestPerMinute *float64 `json:"requestPerMinute,omitempty" xml:"requestPerMinute,omitempty"`
	// example:
	//
	// 2023-05-20
	TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
}

func (s GetNetworkTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNetworkTrendResponseBodyData) SetAvgCost(v float64) *GetNetworkTrendResponseBodyData {
	s.AvgCost = &v
	return s
}

func (s *GetNetworkTrendResponseBodyData) SetAvgResponseTime(v float64) *GetNetworkTrendResponseBodyData {
	s.AvgResponseTime = &v
	return s
}

func (s *GetNetworkTrendResponseBodyData) SetAvgTransformBytes(v float64) *GetNetworkTrendResponseBodyData {
	s.AvgTransformBytes = &v
	return s
}

func (s *GetNetworkTrendResponseBodyData) SetRequestPerMinute(v float64) *GetNetworkTrendResponseBodyData {
	s.RequestPerMinute = &v
	return s
}

func (s *GetNetworkTrendResponseBodyData) SetTimePoint(v string) *GetNetworkTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

type GetNetworkTrendResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkTrendResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkTrendResponse) SetHeaders(v map[string]*string) *GetNetworkTrendResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkTrendResponse) SetStatusCode(v int32) *GetNetworkTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkTrendResponse) SetBody(v *GetNetworkTrendResponseBody) *GetNetworkTrendResponse {
	s.Body = v
	return s
}

type GetStatTrendRequest struct {
	// example:
	//
	// 1.0
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// example:
	//
	// 2021-06-03
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 2021-06-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetStatTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStatTrendRequest) GoString() string {
	return s.String()
}

func (s *GetStatTrendRequest) SetAppVersion(v string) *GetStatTrendRequest {
	s.AppVersion = &v
	return s
}

func (s *GetStatTrendRequest) SetDataSourceId(v string) *GetStatTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetStatTrendRequest) SetEndDate(v string) *GetStatTrendRequest {
	s.EndDate = &v
	return s
}

func (s *GetStatTrendRequest) SetStartDate(v string) *GetStatTrendRequest {
	s.StartDate = &v
	return s
}

func (s *GetStatTrendRequest) SetType(v int32) *GetStatTrendRequest {
	s.Type = &v
	return s
}

type GetStatTrendResponseBody struct {
	// example:
	//
	// 200
	Code *int64                          `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetStatTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetStatTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStatTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetStatTrendResponseBody) SetCode(v int64) *GetStatTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetStatTrendResponseBody) SetData(v []*GetStatTrendResponseBodyData) *GetStatTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetStatTrendResponseBody) SetMsg(v string) *GetStatTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetStatTrendResponseBody) SetSuccess(v bool) *GetStatTrendResponseBody {
	s.Success = &v
	return s
}

type GetStatTrendResponseBodyData struct {
	// example:
	//
	// 52
	AffectedUserCount *int64 `json:"affectedUserCount,omitempty" xml:"affectedUserCount,omitempty"`
	// example:
	//
	// 10.3
	AffectedUserRate *float64 `json:"affectedUserRate,omitempty" xml:"affectedUserRate,omitempty"`
	// example:
	//
	// 120
	ErrorCount *int64 `json:"errorCount,omitempty" xml:"errorCount,omitempty"`
	// example:
	//
	// 25.6
	ErrorRate *float64 `json:"errorRate,omitempty" xml:"errorRate,omitempty"`
	// example:
	//
	// 2021-06-01
	TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
}

func (s GetStatTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStatTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStatTrendResponseBodyData) SetAffectedUserCount(v int64) *GetStatTrendResponseBodyData {
	s.AffectedUserCount = &v
	return s
}

func (s *GetStatTrendResponseBodyData) SetAffectedUserRate(v float64) *GetStatTrendResponseBodyData {
	s.AffectedUserRate = &v
	return s
}

func (s *GetStatTrendResponseBodyData) SetErrorCount(v int64) *GetStatTrendResponseBodyData {
	s.ErrorCount = &v
	return s
}

func (s *GetStatTrendResponseBodyData) SetErrorRate(v float64) *GetStatTrendResponseBodyData {
	s.ErrorRate = &v
	return s
}

func (s *GetStatTrendResponseBodyData) SetTimePoint(v string) *GetStatTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

type GetStatTrendResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStatTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStatTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStatTrendResponse) GoString() string {
	return s.String()
}

func (s *GetStatTrendResponse) SetHeaders(v map[string]*string) *GetStatTrendResponse {
	s.Headers = v
	return s
}

func (s *GetStatTrendResponse) SetStatusCode(v int32) *GetStatTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStatTrendResponse) SetBody(v *GetStatTrendResponseBody) *GetStatTrendResponse {
	s.Body = v
	return s
}

type GetSymUploadParamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.0.3
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// symbol.zip
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	FileType *int32 `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// my-flutter-app
	FlutterName *string `json:"flutterName,omitempty" xml:"flutterName,omitempty"`
}

func (s GetSymUploadParamRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSymUploadParamRequest) GoString() string {
	return s.String()
}

func (s *GetSymUploadParamRequest) SetAppVersion(v string) *GetSymUploadParamRequest {
	s.AppVersion = &v
	return s
}

func (s *GetSymUploadParamRequest) SetDataSourceId(v string) *GetSymUploadParamRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetSymUploadParamRequest) SetFileName(v string) *GetSymUploadParamRequest {
	s.FileName = &v
	return s
}

func (s *GetSymUploadParamRequest) SetFileType(v int32) *GetSymUploadParamRequest {
	s.FileType = &v
	return s
}

func (s *GetSymUploadParamRequest) SetFlutterName(v string) *GetSymUploadParamRequest {
	s.FlutterName = &v
	return s
}

type GetSymUploadParamResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *GetSymUploadParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f07c516457690916816858d94ea
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetSymUploadParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSymUploadParamResponseBody) GoString() string {
	return s.String()
}

func (s *GetSymUploadParamResponseBody) SetCode(v int64) *GetSymUploadParamResponseBody {
	s.Code = &v
	return s
}

func (s *GetSymUploadParamResponseBody) SetData(v *GetSymUploadParamResponseBodyData) *GetSymUploadParamResponseBody {
	s.Data = v
	return s
}

func (s *GetSymUploadParamResponseBody) SetMsg(v string) *GetSymUploadParamResponseBody {
	s.Msg = &v
	return s
}

func (s *GetSymUploadParamResponseBody) SetSuccess(v bool) *GetSymUploadParamResponseBody {
	s.Success = &v
	return s
}

func (s *GetSymUploadParamResponseBody) SetTraceId(v string) *GetSymUploadParamResponseBody {
	s.TraceId = &v
	return s
}

type GetSymUploadParamResponseBodyData struct {
	// example:
	//
	// LTAI5tM4ZXXXXX
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// example:
	//
	// eyJjYWxsYmFja1VybCI6Imh0dHBzOi8vYXBtLnVtZW5nLmNvbS9oc2Yvc3ltL29zcy9ub3RpZnlNc2ciLCJjYqc29uIn0=
	Callback *string `json:"callback,omitempty" xml:"callback,omitempty"`
	// example:
	//
	// tmp/20220428/5fb6001a73749c24fd9cb356_f49a08dc1225438188c109fcf92eb9f3/symbol.zip
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyMi0wNC0yOFQwNDoxMzo0MS43OTJaIiwiY29uZGl0aW9ucyI6W1siZXEiLCIka2V5IiwidG1WpveGZTSXNJbU5oYkd4aVlXTnJRbTlrZVZSNWNHVWlPaUpoY0hCc2FXTmhkR2x2Ymk5cWMyOXVJbjA9In1dfQ==
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// example:
	//
	// 3f67c435e08d164f41f6e522a2b5d1d7feb93000
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// example:
	//
	// https://quickbird.oss-cn-shanghai.aliyuncs.com
	UploadAddress *string `json:"uploadAddress,omitempty" xml:"uploadAddress,omitempty"`
}

func (s GetSymUploadParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSymUploadParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSymUploadParamResponseBodyData) SetAccessKeyId(v string) *GetSymUploadParamResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetCallback(v string) *GetSymUploadParamResponseBodyData {
	s.Callback = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetKey(v string) *GetSymUploadParamResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetPolicy(v string) *GetSymUploadParamResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetSignature(v string) *GetSymUploadParamResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetSymUploadParamResponseBodyData) SetUploadAddress(v string) *GetSymUploadParamResponseBodyData {
	s.UploadAddress = &v
	return s
}

type GetSymUploadParamResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSymUploadParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSymUploadParamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSymUploadParamResponse) GoString() string {
	return s.String()
}

func (s *GetSymUploadParamResponse) SetHeaders(v map[string]*string) *GetSymUploadParamResponse {
	s.Headers = v
	return s
}

func (s *GetSymUploadParamResponse) SetStatusCode(v int32) *GetSymUploadParamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSymUploadParamResponse) SetBody(v *GetSymUploadParamResponseBody) *GetSymUploadParamResponse {
	s.Body = v
	return s
}

type GetTodayStatTrendRequest struct {
	// example:
	//
	// 1.0
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTodayStatTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTodayStatTrendRequest) GoString() string {
	return s.String()
}

func (s *GetTodayStatTrendRequest) SetAppVersion(v string) *GetTodayStatTrendRequest {
	s.AppVersion = &v
	return s
}

func (s *GetTodayStatTrendRequest) SetDataSourceId(v string) *GetTodayStatTrendRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetTodayStatTrendRequest) SetType(v int32) *GetTodayStatTrendRequest {
	s.Type = &v
	return s
}

type GetTodayStatTrendResponseBody struct {
	// example:
	//
	// 200
	Code *int64                               `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetTodayStatTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTodayStatTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTodayStatTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodayStatTrendResponseBody) SetCode(v int64) *GetTodayStatTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetTodayStatTrendResponseBody) SetData(v []*GetTodayStatTrendResponseBodyData) *GetTodayStatTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetTodayStatTrendResponseBody) SetMsg(v string) *GetTodayStatTrendResponseBody {
	s.Msg = &v
	return s
}

func (s *GetTodayStatTrendResponseBody) SetSuccess(v bool) *GetTodayStatTrendResponseBody {
	s.Success = &v
	return s
}

type GetTodayStatTrendResponseBodyData struct {
	// example:
	//
	// 56
	AffectedUserCount *int64 `json:"affectedUserCount,omitempty" xml:"affectedUserCount,omitempty"`
	// example:
	//
	// 10.21
	AffectedUserRate *float64 `json:"affectedUserRate,omitempty" xml:"affectedUserRate,omitempty"`
	// example:
	//
	// 120
	ErrorCount *int64 `json:"errorCount,omitempty" xml:"errorCount,omitempty"`
	// example:
	//
	// 17.24
	ErrorRate *float64 `json:"errorRate,omitempty" xml:"errorRate,omitempty"`
	// example:
	//
	// 13:00
	TimePoint *string `json:"timePoint,omitempty" xml:"timePoint,omitempty"`
}

func (s GetTodayStatTrendResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTodayStatTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTodayStatTrendResponseBodyData) SetAffectedUserCount(v int64) *GetTodayStatTrendResponseBodyData {
	s.AffectedUserCount = &v
	return s
}

func (s *GetTodayStatTrendResponseBodyData) SetAffectedUserRate(v float64) *GetTodayStatTrendResponseBodyData {
	s.AffectedUserRate = &v
	return s
}

func (s *GetTodayStatTrendResponseBodyData) SetErrorCount(v int64) *GetTodayStatTrendResponseBodyData {
	s.ErrorCount = &v
	return s
}

func (s *GetTodayStatTrendResponseBodyData) SetErrorRate(v float64) *GetTodayStatTrendResponseBodyData {
	s.ErrorRate = &v
	return s
}

func (s *GetTodayStatTrendResponseBodyData) SetTimePoint(v string) *GetTodayStatTrendResponseBodyData {
	s.TimePoint = &v
	return s
}

type GetTodayStatTrendResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTodayStatTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTodayStatTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTodayStatTrendResponse) GoString() string {
	return s.String()
}

func (s *GetTodayStatTrendResponse) SetHeaders(v map[string]*string) *GetTodayStatTrendResponse {
	s.Headers = v
	return s
}

func (s *GetTodayStatTrendResponse) SetStatusCode(v int32) *GetTodayStatTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTodayStatTrendResponse) SetBody(v *GetTodayStatTrendResponseBody) *GetTodayStatTrendResponse {
	s.Body = v
	return s
}

type UpdateAlertPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18288
	PlanId *int64 `json:"planId,omitempty" xml:"planId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "1.1.0,1.2.0,1.3.0"
	Versions *string `json:"versions,omitempty" xml:"versions,omitempty"`
}

func (s UpdateAlertPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertPlanRequest) SetDataSourceId(v string) *UpdateAlertPlanRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdateAlertPlanRequest) SetPlanId(v int64) *UpdateAlertPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateAlertPlanRequest) SetVersions(v string) *UpdateAlertPlanRequest {
	s.Versions = &v
	return s
}

type UpdateAlertPlanResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// Success
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAlertPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertPlanResponseBody) SetCode(v int64) *UpdateAlertPlanResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAlertPlanResponseBody) SetMsg(v string) *UpdateAlertPlanResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateAlertPlanResponseBody) SetSuccess(v bool) *UpdateAlertPlanResponseBody {
	s.Success = &v
	return s
}

type UpdateAlertPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertPlanResponse) SetHeaders(v map[string]*string) *UpdateAlertPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertPlanResponse) SetStatusCode(v int32) *UpdateAlertPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertPlanResponse) SetBody(v *UpdateAlertPlanResponseBody) *UpdateAlertPlanResponse {
	s.Body = v
	return s
}

type UploadSymbolFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.0.3
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// symbol.zip
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	FileType *int32 `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// my-flutter-app
	FlutterName *string `json:"flutterName,omitempty" xml:"flutterName,omitempty"`
	// example:
	//
	// -
	OssUrl *string `json:"ossUrl,omitempty" xml:"ossUrl,omitempty"`
}

func (s UploadSymbolFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadSymbolFileRequest) GoString() string {
	return s.String()
}

func (s *UploadSymbolFileRequest) SetAppVersion(v string) *UploadSymbolFileRequest {
	s.AppVersion = &v
	return s
}

func (s *UploadSymbolFileRequest) SetDataSourceId(v string) *UploadSymbolFileRequest {
	s.DataSourceId = &v
	return s
}

func (s *UploadSymbolFileRequest) SetFileName(v string) *UploadSymbolFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadSymbolFileRequest) SetFileType(v int32) *UploadSymbolFileRequest {
	s.FileType = &v
	return s
}

func (s *UploadSymbolFileRequest) SetFlutterName(v string) *UploadSymbolFileRequest {
	s.FlutterName = &v
	return s
}

func (s *UploadSymbolFileRequest) SetOssUrl(v string) *UploadSymbolFileRequest {
	s.OssUrl = &v
	return s
}

type UploadSymbolFileAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.0.3
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5fb6001a73749c24fd9cb356
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// symbol.zip
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	FileType *int32 `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// my-flutter-app
	FlutterName *string `json:"flutterName,omitempty" xml:"flutterName,omitempty"`
	// example:
	//
	// -
	OssUrlObject io.Reader `json:"ossUrl,omitempty" xml:"ossUrl,omitempty"`
}

func (s UploadSymbolFileAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadSymbolFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadSymbolFileAdvanceRequest) SetAppVersion(v string) *UploadSymbolFileAdvanceRequest {
	s.AppVersion = &v
	return s
}

func (s *UploadSymbolFileAdvanceRequest) SetDataSourceId(v string) *UploadSymbolFileAdvanceRequest {
	s.DataSourceId = &v
	return s
}

func (s *UploadSymbolFileAdvanceRequest) SetFileName(v string) *UploadSymbolFileAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *UploadSymbolFileAdvanceRequest) SetFileType(v int32) *UploadSymbolFileAdvanceRequest {
	s.FileType = &v
	return s
}

func (s *UploadSymbolFileAdvanceRequest) SetFlutterName(v string) *UploadSymbolFileAdvanceRequest {
	s.FlutterName = &v
	return s
}

func (s *UploadSymbolFileAdvanceRequest) SetOssUrlObject(v io.Reader) *UploadSymbolFileAdvanceRequest {
	s.OssUrlObject = v
	return s
}

type UploadSymbolFileResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// 8B99488B-2B73-502E-A5F2-00B4746F4325
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f07c516457690916816858d94ea
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s UploadSymbolFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadSymbolFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadSymbolFileResponseBody) SetCode(v int64) *UploadSymbolFileResponseBody {
	s.Code = &v
	return s
}

func (s *UploadSymbolFileResponseBody) SetMsg(v string) *UploadSymbolFileResponseBody {
	s.Msg = &v
	return s
}

func (s *UploadSymbolFileResponseBody) SetRequestId(v string) *UploadSymbolFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadSymbolFileResponseBody) SetSuccess(v bool) *UploadSymbolFileResponseBody {
	s.Success = &v
	return s
}

func (s *UploadSymbolFileResponseBody) SetTraceId(v string) *UploadSymbolFileResponseBody {
	s.TraceId = &v
	return s
}

type UploadSymbolFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadSymbolFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadSymbolFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadSymbolFileResponse) GoString() string {
	return s.String()
}

func (s *UploadSymbolFileResponse) SetHeaders(v map[string]*string) *UploadSymbolFileResponse {
	s.Headers = v
	return s
}

func (s *UploadSymbolFileResponse) SetStatusCode(v int32) *UploadSymbolFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadSymbolFileResponse) SetBody(v *UploadSymbolFileResponseBody) *UploadSymbolFileResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("umeng-apm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除符号表记录
//
// @param tmpReq - DeleteSymRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSymRecordsResponse
func (client *Client) DeleteSymRecordsWithOptions(tmpReq *DeleteSymRecordsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSymRecordsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteSymRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AppVersions)) {
		request.AppVersionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppVersions, tea.String("appVersions"), tea.String("simple"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionsShrink)) {
		body["appVersions"] = request.AppVersionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		body["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["fileType"] = request.FileType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSymRecords"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/deleteSymRecords"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteSymRecordsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteSymRecordsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除符号表记录
//
// @param request - DeleteSymRecordsRequest
//
// @return DeleteSymRecordsResponse
func (client *Client) DeleteSymRecords(request *DeleteSymRecordsRequest) (_result *DeleteSymRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSymRecordsResponse{}
	_body, _err := client.DeleteSymRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取分钟粒度稳定性统计数据
//
// @param request - GetErrorMinuteStatTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErrorMinuteStatTrendResponse
func (client *Client) GetErrorMinuteStatTrendWithOptions(request *GetErrorMinuteStatTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetErrorMinuteStatTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetErrorMinuteStatTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/GetErrorMinuteStatTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetErrorMinuteStatTrendResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetErrorMinuteStatTrendResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取分钟粒度稳定性统计数据
//
// @param request - GetErrorMinuteStatTrendRequest
//
// @return GetErrorMinuteStatTrendResponse
func (client *Client) GetErrorMinuteStatTrend(request *GetErrorMinuteStatTrendRequest) (_result *GetErrorMinuteStatTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetErrorMinuteStatTrendResponse{}
	_body, _err := client.GetErrorMinuteStatTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取H5页面性能统计数据
//
// @param request - GetH5PageTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetH5PageTrendResponse
func (client *Client) GetH5PageTrendWithOptions(request *GetH5PageTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetH5PageTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TimeUnit)) {
		query["timeUnit"] = request.TimeUnit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetH5PageTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/getH5PageTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetH5PageTrendResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetH5PageTrendResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取H5页面性能统计数据
//
// @param request - GetH5PageTrendRequest
//
// @return GetH5PageTrendResponse
func (client *Client) GetH5PageTrend(request *GetH5PageTrendRequest) (_result *GetH5PageTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetH5PageTrendResponse{}
	_body, _err := client.GetH5PageTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取启动性能统计数据
//
// @param request - GetLaunchTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLaunchTrendResponse
func (client *Client) GetLaunchTrendWithOptions(request *GetLaunchTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLaunchTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TimeUnit)) {
		query["timeUnit"] = request.TimeUnit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLaunchTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/getLaunchTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLaunchTrendResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLaunchTrendResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取启动性能统计数据
//
// @param request - GetLaunchTrendRequest
//
// @return GetLaunchTrendResponse
func (client *Client) GetLaunchTrend(request *GetLaunchTrendRequest) (_result *GetLaunchTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLaunchTrendResponse{}
	_body, _err := client.GetLaunchTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取原生页面性能统计数据
//
// @param request - GetNativePageTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNativePageTrendResponse
func (client *Client) GetNativePageTrendWithOptions(request *GetNativePageTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNativePageTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TimeUnit)) {
		query["timeUnit"] = request.TimeUnit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNativePageTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/getNativePageTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetNativePageTrendResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetNativePageTrendResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取原生页面性能统计数据
//
// @param request - GetNativePageTrendRequest
//
// @return GetNativePageTrendResponse
func (client *Client) GetNativePageTrend(request *GetNativePageTrendRequest) (_result *GetNativePageTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNativePageTrendResponse{}
	_body, _err := client.GetNativePageTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取分钟粒度网络统计数据
//
// @param request - GetNetworkMinuteTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkMinuteTrendResponse
func (client *Client) GetNetworkMinuteTrendWithOptions(request *GetNetworkMinuteTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNetworkMinuteTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNetworkMinuteTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/getNetworkMinuteTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetNetworkMinuteTrendResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetNetworkMinuteTrendResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取分钟粒度网络统计数据
//
// @param request - GetNetworkMinuteTrendRequest
//
// @return GetNetworkMinuteTrendResponse
func (client *Client) GetNetworkMinuteTrend(request *GetNetworkMinuteTrendRequest) (_result *GetNetworkMinuteTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNetworkMinuteTrendResponse{}
	_body, _err := client.GetNetworkMinuteTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网络性能统计数据
//
// @param request - GetNetworkTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkTrendResponse
func (client *Client) GetNetworkTrendWithOptions(request *GetNetworkTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetNetworkTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TimeUnit)) {
		query["timeUnit"] = request.TimeUnit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNetworkTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/getNetworkTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetNetworkTrendResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetNetworkTrendResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取网络性能统计数据
//
// @param request - GetNetworkTrendRequest
//
// @return GetNetworkTrendResponse
func (client *Client) GetNetworkTrend(request *GetNetworkTrendRequest) (_result *GetNetworkTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNetworkTrendResponse{}
	_body, _err := client.GetNetworkTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取离线统计数据
//
// @param request - GetStatTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStatTrendResponse
func (client *Client) GetStatTrendWithOptions(request *GetStatTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetStatTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStatTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/getStatTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetStatTrendResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetStatTrendResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取离线统计数据
//
// @param request - GetStatTrendRequest
//
// @return GetStatTrendResponse
func (client *Client) GetStatTrend(request *GetStatTrendRequest) (_result *GetStatTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetStatTrendResponse{}
	_body, _err := client.GetStatTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取符号表文件上传参数
//
// @param request - GetSymUploadParamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSymUploadParamResponse
func (client *Client) GetSymUploadParamWithOptions(request *GetSymUploadParamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSymUploadParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["fileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.FlutterName)) {
		query["flutterName"] = request.FlutterName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSymUploadParam"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/getSymUploadParam"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSymUploadParamResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSymUploadParamResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取符号表文件上传参数
//
// @param request - GetSymUploadParamRequest
//
// @return GetSymUploadParamResponse
func (client *Client) GetSymUploadParam(request *GetSymUploadParamRequest) (_result *GetSymUploadParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSymUploadParamResponse{}
	_body, _err := client.GetSymUploadParamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取今日实时统计数据
//
// @param request - GetTodayStatTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTodayStatTrendResponse
func (client *Client) GetTodayStatTrendWithOptions(request *GetTodayStatTrendRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTodayStatTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTodayStatTrend"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/stat/getTodayStatTrend"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTodayStatTrendResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTodayStatTrendResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取今日实时统计数据
//
// @param request - GetTodayStatTrendRequest
//
// @return GetTodayStatTrendResponse
func (client *Client) GetTodayStatTrend(request *GetTodayStatTrendRequest) (_result *GetTodayStatTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTodayStatTrendResponse{}
	_body, _err := client.GetTodayStatTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新监控告警计划
//
// @param request - UpdateAlertPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertPlanResponse
func (client *Client) UpdateAlertPlanWithOptions(request *UpdateAlertPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAlertPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["planId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.Versions)) {
		query["versions"] = request.Versions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlertPlan"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/updateAlertPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateAlertPlanResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateAlertPlanResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新监控告警计划
//
// @param request - UpdateAlertPlanRequest
//
// @return UpdateAlertPlanResponse
func (client *Client) UpdateAlertPlan(request *UpdateAlertPlanRequest) (_result *UpdateAlertPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlertPlanResponse{}
	_body, _err := client.UpdateAlertPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传符号表文件
//
// @param request - UploadSymbolFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadSymbolFileResponse
func (client *Client) UploadSymbolFileWithOptions(request *UploadSymbolFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadSymbolFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["fileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.FlutterName)) {
		query["flutterName"] = request.FlutterName
	}

	if !tea.BoolValue(util.IsUnset(request.OssUrl)) {
		query["ossUrl"] = request.OssUrl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadSymbolFile"),
		Version:     tea.String("2022-02-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/uploadSymbolFile"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UploadSymbolFileResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UploadSymbolFileResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 上传符号表文件
//
// @param request - UploadSymbolFileRequest
//
// @return UploadSymbolFileResponse
func (client *Client) UploadSymbolFile(request *UploadSymbolFileRequest) (_result *UploadSymbolFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadSymbolFileResponse{}
	_body, _err := client.UploadSymbolFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadSymbolFileAdvance(request *UploadSymbolFileAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadSymbolFileResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("umeng-apm"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	uploadSymbolFileReq := &UploadSymbolFileRequest{}
	openapiutil.Convert(request, uploadSymbolFileReq)
	if !tea.BoolValue(util.IsUnset(request.OssUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.OssUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		uploadSymbolFileReq.OssUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	uploadSymbolFileResp, _err := client.UploadSymbolFileWithOptions(uploadSymbolFileReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = uploadSymbolFileResp
	return _result, _err
}
