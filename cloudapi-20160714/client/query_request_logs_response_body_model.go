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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request logs.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type QueryRequestLogsResponseBodyRequestLogsRequestLog struct {
	// The API ID.
	//
	// example:
	//
	// 4b83229ebcab4ecd88956fb3********
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// ApiName
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The application name.
	//
	// example:
	//
	// VIPROOM_VIPROOM
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when API Gateway finished forwarding the request to the backend service.
	//
	// example:
	//
	// 1731487224969
	BackendRequestEnd *int64 `json:"BackendRequestEnd,omitempty" xml:"BackendRequestEnd,omitempty"`
	// The time when API Gateway started to forward the request to the backend service.
	//
	// example:
	//
	// 1731487224969
	BackendRequestStart *int64 `json:"BackendRequestStart,omitempty" xml:"BackendRequestStart,omitempty"`
	// The time when API Gateway finished receiving the response from the backend service.
	//
	// example:
	//
	// 1731487224989
	BackendResponseEnd *int64 `json:"BackendResponseEnd,omitempty" xml:"BackendResponseEnd,omitempty"`
	// The time when API Gateway started to receive the response from the backend service.
	//
	// example:
	//
	// 1731487224989
	BackendResponseStart *int64 `json:"BackendResponseStart,omitempty" xml:"BackendResponseStart,omitempty"`
	// The IP address of the client that sends the request.
	//
	// example:
	//
	// 21.237.XXX.XXX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The X-Ca-Nonce header included in the request from the client.
	//
	// example:
	//
	// d43df9db-3b05-4cd6-888a-1c0b********
	ClientNonce *string `json:"ClientNonce,omitempty" xml:"ClientNonce,omitempty"`
	// The application ID that is used by the caller.
	//
	// example:
	//
	// 11096****
	ConsumerAppId *string `json:"ConsumerAppId,omitempty" xml:"ConsumerAppId,omitempty"`
	// The App Key that is used by the caller.
	//
	// example:
	//
	// 20412****
	ConsumerAppKey *string `json:"ConsumerAppKey,omitempty" xml:"ConsumerAppKey,omitempty"`
	// The custom trace ID.
	//
	// example:
	//
	// 95657ED9-2F6F-426F-BD99-79C8********
	CustomTraceId *string `json:"CustomTraceId,omitempty" xml:"CustomTraceId,omitempty"`
	// The requested domain name in the request.
	//
	// example:
	//
	// 360bdd88695c48ae8085c7f2********-ap-southeast-1.alicloudapi.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// X500ER
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Backend service connect failed `Timeout connecting to [/1XX.20.0.XX:8080]`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The specific error message returned by the backend service.
	//
	// example:
	//
	// error msg
	Exception *string `json:"Exception,omitempty" xml:"Exception,omitempty"`
	// The time when API Gateway finished receiving the request.
	//
	// example:
	//
	// 1731487224968
	FrontRequestEnd *int64 `json:"FrontRequestEnd,omitempty" xml:"FrontRequestEnd,omitempty"`
	// The time when API Gateway started to receive the request.
	//
	// example:
	//
	// 1731487224968
	FrontRequestStart *int64 `json:"FrontRequestStart,omitempty" xml:"FrontRequestStart,omitempty"`
	// The time when API Gateway finished forwarding the response to the client.
	//
	// example:
	//
	// 1731487224989
	FrontResponseEnd *int64 `json:"FrontResponseEnd,omitempty" xml:"FrontResponseEnd,omitempty"`
	// The time when API Gateway started to forward the response to the client.
	//
	// example:
	//
	// 1731487224989
	FrontResponseStart *int64 `json:"FrontResponseStart,omitempty" xml:"FrontResponseStart,omitempty"`
	// The ID of the API group to which the API belongs.
	//
	// example:
	//
	// dc024277fe6c4cada79ba0bd6********
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group to which the API belongs.
	//
	// example:
	//
	// GroupName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The HTTP method that is used to send the request.
	//
	// example:
	//
	// POST
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The path of the request.
	//
	// example:
	//
	// /testPath
	HttpPath *string `json:"HttpPath,omitempty" xml:"HttpPath,omitempty"`
	// The initial request ID when API Gateway calls an API. For example, if API-1 calls API-2, the initialRequestId parameter in the log of API-2 indicates the ID of the request from API-1.
	//
	// example:
	//
	// 95657ED9-2F6F-426F-BD99-79C8********
	InitialRequestId *string `json:"InitialRequestId,omitempty" xml:"InitialRequestId,omitempty"`
	// The ID of the API Gateway instance to which the API belongs.
	//
	// example:
	//
	// apigateway-bj-ab2b********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The JSON web token (JWT) claims. The claims can be configured at the group level.
	//
	// example:
	//
	// {}
	JwtClaims *string `json:"JwtClaims,omitempty" xml:"JwtClaims,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request body. A request body cannot exceed 1,024 bytes in size.
	//
	// example:
	//
	// param=paramName
	RequestBody *string `json:"RequestBody,omitempty" xml:"RequestBody,omitempty"`
	// The request headers.
	//
	// example:
	//
	// content-type: application/x-www-form-urlencoded
	RequestHeaders *string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 95657ED9-2F6F-426F-BD99-79C8********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The protocol used by the client to send the request. Valid values: HTTP, HTTPS, and WS.
	//
	// example:
	//
	// HTTP
	RequestProtocol *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
	// The query string for the request.
	//
	// example:
	//
	// username=name
	RequestQueryString *string `json:"RequestQueryString,omitempty" xml:"RequestQueryString,omitempty"`
	// The size of the request. Unit: bytes.
	//
	// example:
	//
	// 1923
	RequestSize *string `json:"RequestSize,omitempty" xml:"RequestSize,omitempty"`
	// The request time, in UTC.
	//
	// example:
	//
	// 2022-10-29T03:59:59Z
	RequestTime *string `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// The response body. A response body cannot exceed 1,024 bytes in size.
	//
	// example:
	//
	// param=paramName
	ResponseBody *string `json:"ResponseBody,omitempty" xml:"ResponseBody,omitempty"`
	// The headers in the API response.
	//
	// example:
	//
	// content-type: application/x-www-form-urlencoded
	ResponseHeaders *string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty"`
	// The size of returned data. Unit: bytes.
	//
	// example:
	//
	// 23441
	ResponseSize *string `json:"ResponseSize,omitempty" xml:"ResponseSize,omitempty"`
	// The total time consumed to access the backend resources. The total time includes the time consumed to request a connection to the resources, the time consumed to establish the connection, and the time consumed to call the backend service. Unit: milliseconds.
	//
	// example:
	//
	// 324
	ServiceLatency *string `json:"ServiceLatency,omitempty" xml:"ServiceLatency,omitempty"`
	// The ID of the API environment.
	//
	// example:
	//
	// 8a305b7f10334052a52d9156********
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The name of the API environment.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The status code returned.
	//
	// example:
	//
	// 200
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The total time consumed by the request. Unit: milliseconds.
	//
	// example:
	//
	// 1345
	TotalLatency *string `json:"TotalLatency,omitempty" xml:"TotalLatency,omitempty"`
	// The plug-in hit by the request and the relevant context.
	//
	// example:
	//
	// []
	Plugin *string `json:"plugin,omitempty" xml:"plugin,omitempty"`
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
