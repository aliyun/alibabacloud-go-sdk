// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRequestLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryRequestLogsResponseBody
	GetRequestId() *string
	SetRequestLogs(v *QueryRequestLogsResponseBodyRequestLogs) *QueryRequestLogsResponseBody
	GetRequestLogs() *QueryRequestLogsResponseBodyRequestLogs
}

type QueryRequestLogsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CE5722A6-AE78-4741-A9B0-6C81********
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestLogs *QueryRequestLogsResponseBodyRequestLogs `json:"RequestLogs,omitempty" xml:"RequestLogs,omitempty" type:"Struct"`
}

func (s QueryRequestLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRequestLogsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRequestLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRequestLogsResponseBody) GetRequestLogs() *QueryRequestLogsResponseBodyRequestLogs {
	return s.RequestLogs
}

func (s *QueryRequestLogsResponseBody) SetRequestId(v string) *QueryRequestLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRequestLogsResponseBody) SetRequestLogs(v *QueryRequestLogsResponseBodyRequestLogs) *QueryRequestLogsResponseBody {
	s.RequestLogs = v
	return s
}

func (s *QueryRequestLogsResponseBody) Validate() error {
	if s.RequestLogs != nil {
		if err := s.RequestLogs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRequestLogsResponseBodyRequestLogs struct {
	RequestLog []*QueryRequestLogsResponseBodyRequestLogsRequestLog `json:"RequestLog,omitempty" xml:"RequestLog,omitempty" type:"Repeated"`
}

func (s QueryRequestLogsResponseBodyRequestLogs) String() string {
	return dara.Prettify(s)
}

func (s QueryRequestLogsResponseBodyRequestLogs) GoString() string {
	return s.String()
}

func (s *QueryRequestLogsResponseBodyRequestLogs) GetRequestLog() []*QueryRequestLogsResponseBodyRequestLogsRequestLog {
	return s.RequestLog
}

func (s *QueryRequestLogsResponseBodyRequestLogs) SetRequestLog(v []*QueryRequestLogsResponseBodyRequestLogsRequestLog) *QueryRequestLogsResponseBodyRequestLogs {
	s.RequestLog = v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogs) Validate() error {
	if s.RequestLog != nil {
		for _, item := range s.RequestLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRequestLogsResponseBodyRequestLogsRequestLog struct {
	ApiId                *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName              *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AppName              *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BackendRequestEnd    *int64  `json:"BackendRequestEnd,omitempty" xml:"BackendRequestEnd,omitempty"`
	BackendRequestStart  *int64  `json:"BackendRequestStart,omitempty" xml:"BackendRequestStart,omitempty"`
	BackendResponseEnd   *int64  `json:"BackendResponseEnd,omitempty" xml:"BackendResponseEnd,omitempty"`
	BackendResponseStart *int64  `json:"BackendResponseStart,omitempty" xml:"BackendResponseStart,omitempty"`
	ClientIp             *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientNonce          *string `json:"ClientNonce,omitempty" xml:"ClientNonce,omitempty"`
	ConsumerAppId        *string `json:"ConsumerAppId,omitempty" xml:"ConsumerAppId,omitempty"`
	ConsumerAppKey       *string `json:"ConsumerAppKey,omitempty" xml:"ConsumerAppKey,omitempty"`
	CustomTraceId        *string `json:"CustomTraceId,omitempty" xml:"CustomTraceId,omitempty"`
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ErrorCode            *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage         *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Exception            *string `json:"Exception,omitempty" xml:"Exception,omitempty"`
	FrontRequestEnd      *int64  `json:"FrontRequestEnd,omitempty" xml:"FrontRequestEnd,omitempty"`
	FrontRequestStart    *int64  `json:"FrontRequestStart,omitempty" xml:"FrontRequestStart,omitempty"`
	FrontResponseEnd     *int64  `json:"FrontResponseEnd,omitempty" xml:"FrontResponseEnd,omitempty"`
	FrontResponseStart   *int64  `json:"FrontResponseStart,omitempty" xml:"FrontResponseStart,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HttpMethod           *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpPath             *string `json:"HttpPath,omitempty" xml:"HttpPath,omitempty"`
	InitialRequestId     *string `json:"InitialRequestId,omitempty" xml:"InitialRequestId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JwtClaims            *string `json:"JwtClaims,omitempty" xml:"JwtClaims,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RequestBody          *string `json:"RequestBody,omitempty" xml:"RequestBody,omitempty"`
	RequestHeaders       *string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestProtocol      *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
	RequestQueryString   *string `json:"RequestQueryString,omitempty" xml:"RequestQueryString,omitempty"`
	RequestSize          *string `json:"RequestSize,omitempty" xml:"RequestSize,omitempty"`
	RequestTime          *string `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	ResponseBody         *string `json:"ResponseBody,omitempty" xml:"ResponseBody,omitempty"`
	ResponseHeaders      *string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty"`
	ResponseSize         *string `json:"ResponseSize,omitempty" xml:"ResponseSize,omitempty"`
	ServiceLatency       *string `json:"ServiceLatency,omitempty" xml:"ServiceLatency,omitempty"`
	StageId              *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName            *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	StatusCode           *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	TotalLatency         *string `json:"TotalLatency,omitempty" xml:"TotalLatency,omitempty"`
	Plugin               *string `json:"plugin,omitempty" xml:"plugin,omitempty"`
}

func (s QueryRequestLogsResponseBodyRequestLogsRequestLog) String() string {
	return dara.Prettify(s)
}

func (s QueryRequestLogsResponseBodyRequestLogsRequestLog) GoString() string {
	return s.String()
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetApiId() *string {
	return s.ApiId
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetApiName() *string {
	return s.ApiName
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetAppName() *string {
	return s.AppName
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetBackendRequestEnd() *int64 {
	return s.BackendRequestEnd
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetBackendRequestStart() *int64 {
	return s.BackendRequestStart
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetBackendResponseEnd() *int64 {
	return s.BackendResponseEnd
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetBackendResponseStart() *int64 {
	return s.BackendResponseStart
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetClientIp() *string {
	return s.ClientIp
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetClientNonce() *string {
	return s.ClientNonce
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetConsumerAppId() *string {
	return s.ConsumerAppId
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetConsumerAppKey() *string {
	return s.ConsumerAppKey
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetCustomTraceId() *string {
	return s.CustomTraceId
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetDomain() *string {
	return s.Domain
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetException() *string {
	return s.Exception
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetFrontRequestEnd() *int64 {
	return s.FrontRequestEnd
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetFrontRequestStart() *int64 {
	return s.FrontRequestStart
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetFrontResponseEnd() *int64 {
	return s.FrontResponseEnd
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetFrontResponseStart() *int64 {
	return s.FrontResponseStart
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetHttpPath() *string {
	return s.HttpPath
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetInitialRequestId() *string {
	return s.InitialRequestId
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetJwtClaims() *string {
	return s.JwtClaims
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetRegion() *string {
	return s.Region
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetRequestBody() *string {
	return s.RequestBody
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetRequestHeaders() *string {
	return s.RequestHeaders
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetRequestProtocol() *string {
	return s.RequestProtocol
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetRequestQueryString() *string {
	return s.RequestQueryString
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetRequestSize() *string {
	return s.RequestSize
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetRequestTime() *string {
	return s.RequestTime
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetResponseBody() *string {
	return s.ResponseBody
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetResponseHeaders() *string {
	return s.ResponseHeaders
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetResponseSize() *string {
	return s.ResponseSize
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetServiceLatency() *string {
	return s.ServiceLatency
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetStageId() *string {
	return s.StageId
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetStageName() *string {
	return s.StageName
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetStatusCode() *string {
	return s.StatusCode
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetTotalLatency() *string {
	return s.TotalLatency
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) GetPlugin() *string {
	return s.Plugin
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetApiId(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ApiId = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetApiName(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ApiName = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetAppName(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.AppName = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetBackendRequestEnd(v int64) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.BackendRequestEnd = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetBackendRequestStart(v int64) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.BackendRequestStart = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetBackendResponseEnd(v int64) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.BackendResponseEnd = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetBackendResponseStart(v int64) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.BackendResponseStart = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetClientIp(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ClientIp = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetClientNonce(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ClientNonce = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetConsumerAppId(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ConsumerAppId = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetConsumerAppKey(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ConsumerAppKey = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetCustomTraceId(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.CustomTraceId = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetDomain(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.Domain = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetErrorCode(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ErrorCode = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetErrorMessage(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetException(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.Exception = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetFrontRequestEnd(v int64) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.FrontRequestEnd = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetFrontRequestStart(v int64) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.FrontRequestStart = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetFrontResponseEnd(v int64) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.FrontResponseEnd = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetFrontResponseStart(v int64) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.FrontResponseStart = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetGroupId(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.GroupId = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetGroupName(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.GroupName = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetHttpMethod(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.HttpMethod = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetHttpPath(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.HttpPath = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetInitialRequestId(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.InitialRequestId = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetInstanceId(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.InstanceId = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetJwtClaims(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.JwtClaims = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetRegion(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.Region = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetRequestBody(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.RequestBody = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetRequestHeaders(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.RequestHeaders = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetRequestId(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.RequestId = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetRequestProtocol(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.RequestProtocol = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetRequestQueryString(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.RequestQueryString = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetRequestSize(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.RequestSize = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetRequestTime(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.RequestTime = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetResponseBody(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ResponseBody = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetResponseHeaders(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ResponseHeaders = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetResponseSize(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ResponseSize = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetServiceLatency(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.ServiceLatency = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetStageId(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.StageId = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetStageName(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.StageName = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetStatusCode(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.StatusCode = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetTotalLatency(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.TotalLatency = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) SetPlugin(v string) *QueryRequestLogsResponseBodyRequestLogsRequestLog {
	s.Plugin = &v
	return s
}

func (s *QueryRequestLogsResponseBodyRequestLogsRequestLog) Validate() error {
	return dara.Validate(s)
}
