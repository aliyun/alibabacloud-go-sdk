// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsDebugSampleLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPtsDebugSampleLogsResponseBody
	GetCode() *string
	SetMessage(v string) *GetPtsDebugSampleLogsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetPtsDebugSampleLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetPtsDebugSampleLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetPtsDebugSampleLogsResponseBody
	GetRequestId() *string
	SetSamplingLogs(v []*GetPtsDebugSampleLogsResponseBodySamplingLogs) *GetPtsDebugSampleLogsResponseBody
	GetSamplingLogs() []*GetPtsDebugSampleLogsResponseBodySamplingLogs
	SetSuccess(v bool) *GetPtsDebugSampleLogsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetPtsDebugSampleLogsResponseBody
	GetTotalCount() *int64
}

type GetPtsDebugSampleLogsResponseBody struct {
	// The system status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 4001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sampling logs.
	SamplingLogs []*GetPtsDebugSampleLogsResponseBodySamplingLogs `json:"SamplingLogs,omitempty" xml:"SamplingLogs,omitempty" type:"Repeated"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetPtsDebugSampleLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPtsDebugSampleLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsDebugSampleLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPtsDebugSampleLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPtsDebugSampleLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetPtsDebugSampleLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetPtsDebugSampleLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPtsDebugSampleLogsResponseBody) GetSamplingLogs() []*GetPtsDebugSampleLogsResponseBodySamplingLogs {
	return s.SamplingLogs
}

func (s *GetPtsDebugSampleLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPtsDebugSampleLogsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetPtsDebugSampleLogsResponseBody) SetCode(v string) *GetPtsDebugSampleLogsResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBody) SetMessage(v string) *GetPtsDebugSampleLogsResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBody) SetPageNumber(v int32) *GetPtsDebugSampleLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBody) SetPageSize(v int32) *GetPtsDebugSampleLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBody) SetRequestId(v string) *GetPtsDebugSampleLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBody) SetSamplingLogs(v []*GetPtsDebugSampleLogsResponseBodySamplingLogs) *GetPtsDebugSampleLogsResponseBody {
	s.SamplingLogs = v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBody) SetSuccess(v bool) *GetPtsDebugSampleLogsResponseBody {
	s.Success = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBody) SetTotalCount(v int64) *GetPtsDebugSampleLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBody) Validate() error {
	if s.SamplingLogs != nil {
		for _, item := range s.SamplingLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPtsDebugSampleLogsResponseBodySamplingLogs struct {
	// The ID of the session.
	//
	// example:
	//
	// 65354719
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The name of the session.
	ChainName *string `json:"ChainName,omitempty" xml:"ChainName,omitempty"`
	// The assertion check result.
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The parameter export configuration.
	//
	// example:
	//
	// {\\"skuId\\":\\"{R:json@$.page.list[0].skuId}\\"}
	ExportConfig *string `json:"ExportConfig,omitempty" xml:"ExportConfig,omitempty"`
	// The exported parameters.
	//
	// example:
	//
	// {"skuId":"1"}
	ExportContent *string `json:"ExportContent,omitempty" xml:"ExportContent,omitempty"`
	// The body of the request.
	//
	// example:
	//
	// {"loginacct":"acce"}
	HttpRequestBody *string `json:"HttpRequestBody,omitempty" xml:"HttpRequestBody,omitempty"`
	// The request headers.
	//
	// example:
	//
	// [{"name":"v2","sensitive":false,"value":"1"},{"name":"x-pts-test","sensitive":false,"value":"2"}]
	HttpRequestHeaders *string `json:"HttpRequestHeaders,omitempty" xml:"HttpRequestHeaders,omitempty"`
	// The request method.
	//
	// example:
	//
	// GET
	HttpRequestMethod *string `json:"HttpRequestMethod,omitempty" xml:"HttpRequestMethod,omitempty"`
	// The endpoint that specifies where the request is directed.
	//
	// example:
	//
	// http://www.example.com
	HttpRequestUrl *string `json:"HttpRequestUrl,omitempty" xml:"HttpRequestUrl,omitempty"`
	// The response body.
	//
	// example:
	//
	// {"timestamp":1679903049155,"status":404,"error":"Not Found","message":"No message available","path":"/"}
	HttpResponseBody *string `json:"HttpResponseBody,omitempty" xml:"HttpResponseBody,omitempty"`
	// The error message.
	//
	// example:
	//
	// ""
	HttpResponseFailMsg *string `json:"HttpResponseFailMsg,omitempty" xml:"HttpResponseFailMsg,omitempty"`
	// The response headers.
	//
	// example:
	//
	// [{"valuePos":18,"name":"transfer-encoding","buffer":{"empty":false,"full":false},"sensitive":false,"value":"chunked"},{"valuePos":13,"name":"Content-Type","buffer":{"empty":false,"full":false},"sensitive":false,"value":"application/json;charset=UTF-8"},{"valuePos":5,"name":"Date","buffer":{"empty":false,"full":false},"sensitive":false,"value":"Mon, 27 Mar 2023 07:44:08 GMT"}]
	HttpResponseHeaders *string `json:"HttpResponseHeaders,omitempty" xml:"HttpResponseHeaders,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpResponseStatus *string `json:"HttpResponseStatus,omitempty" xml:"HttpResponseStatus,omitempty"`
	// The time when the request was sent.
	//
	// example:
	//
	// 12
	HttpStartTime *int64 `json:"HttpStartTime,omitempty" xml:"HttpStartTime,omitempty"`
	// The HTTP timing information in a waterfall format.
	//
	// example:
	//
	// {"traceId":"0:1:10a94f66pts-2069351-allsparktask","requests":[{"lease":{"conn":{"duration":-1,"finish":-1,"operation":"conn","start":-1},"dns":{"duration":-1,"finish":-1,"operation":"dns","start":-1},"duration":-1,"finish":-1,"operation":"lease","start":32277914755},"recv":{"duration":225975,"finish":32283700284,"message":"","operation":"recv","start":32283474309},"sent":{"duration":594179,"finish":32278776504,"message":"","operation":"sent","start":32278182325},"tag":"GET http://tomcodemall.com:30080/api/product/skuinfo/list?key=2&vv=4&t4=%EF%BB%BF101"}],"message":""}
	HttpTiming *string `json:"HttpTiming,omitempty" xml:"HttpTiming,omitempty"`
	// The imported parameters.
	//
	// example:
	//
	// ""
	ImportContent *string `json:"ImportContent,omitempty" xml:"ImportContent,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 1345531
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The response time. Unit: ms.
	//
	// example:
	//
	// 230
	Rt *string `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// The timestamp. Unit: ms.
	//
	// example:
	//
	// 1650253024471
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetPtsDebugSampleLogsResponseBodySamplingLogs) String() string {
	return dara.Prettify(s)
}

func (s GetPtsDebugSampleLogsResponseBodySamplingLogs) GoString() string {
	return s.String()
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetChainId() *string {
	return s.ChainId
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetChainName() *string {
	return s.ChainName
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetCheckResult() *string {
	return s.CheckResult
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetExportConfig() *string {
	return s.ExportConfig
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetExportContent() *string {
	return s.ExportContent
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpRequestBody() *string {
	return s.HttpRequestBody
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpRequestHeaders() *string {
	return s.HttpRequestHeaders
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpRequestMethod() *string {
	return s.HttpRequestMethod
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpRequestUrl() *string {
	return s.HttpRequestUrl
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpResponseBody() *string {
	return s.HttpResponseBody
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpResponseFailMsg() *string {
	return s.HttpResponseFailMsg
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpResponseHeaders() *string {
	return s.HttpResponseHeaders
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpResponseStatus() *string {
	return s.HttpResponseStatus
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpStartTime() *int64 {
	return s.HttpStartTime
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetHttpTiming() *string {
	return s.HttpTiming
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetImportContent() *string {
	return s.ImportContent
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetRt() *string {
	return s.Rt
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetChainId(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.ChainId = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetChainName(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.ChainName = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetCheckResult(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.CheckResult = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetExportConfig(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.ExportConfig = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetExportContent(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.ExportContent = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpRequestBody(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpRequestBody = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpRequestHeaders(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpRequestHeaders = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpRequestMethod(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpRequestMethod = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpRequestUrl(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpRequestUrl = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpResponseBody(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpResponseBody = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpResponseFailMsg(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpResponseFailMsg = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpResponseHeaders(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpResponseHeaders = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpResponseStatus(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpResponseStatus = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpStartTime(v int64) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpStartTime = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetHttpTiming(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.HttpTiming = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetImportContent(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.ImportContent = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetNodeId(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.NodeId = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetRt(v string) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.Rt = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) SetTimestamp(v int64) *GetPtsDebugSampleLogsResponseBodySamplingLogs {
	s.Timestamp = &v
	return s
}

func (s *GetPtsDebugSampleLogsResponseBodySamplingLogs) Validate() error {
	return dara.Validate(s)
}
