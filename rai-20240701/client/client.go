// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchContentAsyncDetectRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName          *string                                               `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceParameterList []*BatchContentAsyncDetectRequestServiceParameterList `json:"serviceParameterList,omitempty" xml:"serviceParameterList,omitempty" type:"Repeated"`
}

func (s BatchContentAsyncDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchContentAsyncDetectRequest) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectRequest) SetRegionId(v string) *BatchContentAsyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *BatchContentAsyncDetectRequest) SetSceneName(v string) *BatchContentAsyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *BatchContentAsyncDetectRequest) SetServiceName(v string) *BatchContentAsyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *BatchContentAsyncDetectRequest) SetServiceParameterList(v []*BatchContentAsyncDetectRequestServiceParameterList) *BatchContentAsyncDetectRequest {
	s.ServiceParameterList = v
	return s
}

type BatchContentAsyncDetectRequestServiceParameterList struct {
	// example:
	//
	// "XXXXXXXX"
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s BatchContentAsyncDetectRequestServiceParameterList) String() string {
	return tea.Prettify(s)
}

func (s BatchContentAsyncDetectRequestServiceParameterList) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectRequestServiceParameterList) SetContent(v string) *BatchContentAsyncDetectRequestServiceParameterList {
	s.Content = &v
	return s
}

type BatchContentAsyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BatchContentAsyncDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9736C44E-F718-566B-821F-710216aAAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****F68692
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchContentAsyncDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchContentAsyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectResponseBody) SetCode(v string) *BatchContentAsyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetData(v *BatchContentAsyncDetectResponseBodyData) *BatchContentAsyncDetectResponseBody {
	s.Data = v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetHttpStatusCode(v int32) *BatchContentAsyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetMessage(v string) *BatchContentAsyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetRequestId(v string) *BatchContentAsyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetSuccess(v bool) *BatchContentAsyncDetectResponseBody {
	s.Success = &v
	return s
}

type BatchContentAsyncDetectResponseBodyData struct {
	// example:
	//
	// 19657954336
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchContentAsyncDetectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchContentAsyncDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectResponseBodyData) SetTaskId(v string) *BatchContentAsyncDetectResponseBodyData {
	s.TaskId = &v
	return s
}

type BatchContentAsyncDetectResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchContentAsyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchContentAsyncDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchContentAsyncDetectResponse) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectResponse) SetHeaders(v map[string]*string) *BatchContentAsyncDetectResponse {
	s.Headers = v
	return s
}

func (s *BatchContentAsyncDetectResponse) SetStatusCode(v int32) *BatchContentAsyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchContentAsyncDetectResponse) SetBody(v *BatchContentAsyncDetectResponseBody) *BatchContentAsyncDetectResponse {
	s.Body = v
	return s
}

type BatchContentSyncDetectRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	//
	// imageDetection
	ServiceName          *string                                              `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceParameterList []*BatchContentSyncDetectRequestServiceParameterList `json:"serviceParameterList,omitempty" xml:"serviceParameterList,omitempty" type:"Repeated"`
}

func (s BatchContentSyncDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchContentSyncDetectRequest) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectRequest) SetRegionId(v string) *BatchContentSyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *BatchContentSyncDetectRequest) SetSceneName(v string) *BatchContentSyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *BatchContentSyncDetectRequest) SetServiceName(v string) *BatchContentSyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *BatchContentSyncDetectRequest) SetServiceParameterList(v []*BatchContentSyncDetectRequestServiceParameterList) *BatchContentSyncDetectRequest {
	s.ServiceParameterList = v
	return s
}

type BatchContentSyncDetectRequestServiceParameterList struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s BatchContentSyncDetectRequestServiceParameterList) String() string {
	return tea.Prettify(s)
}

func (s BatchContentSyncDetectRequestServiceParameterList) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectRequestServiceParameterList) SetContent(v string) *BatchContentSyncDetectRequestServiceParameterList {
	s.Content = &v
	return s
}

type BatchContentSyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BatchContentSyncDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchContentSyncDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchContentSyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectResponseBody) SetCode(v string) *BatchContentSyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetData(v *BatchContentSyncDetectResponseBodyData) *BatchContentSyncDetectResponseBody {
	s.Data = v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetHttpStatusCode(v int32) *BatchContentSyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetMessage(v string) *BatchContentSyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetRequestId(v string) *BatchContentSyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchContentSyncDetectResponseBody) SetSuccess(v bool) *BatchContentSyncDetectResponseBody {
	s.Success = &v
	return s
}

type BatchContentSyncDetectResponseBodyData struct {
	DetectResultList []*BatchContentSyncDetectResponseBodyDataDetectResultList `json:"DetectResultList,omitempty" xml:"DetectResultList,omitempty" type:"Repeated"`
}

func (s BatchContentSyncDetectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchContentSyncDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectResponseBodyData) SetDetectResultList(v []*BatchContentSyncDetectResponseBodyDataDetectResultList) *BatchContentSyncDetectResponseBodyData {
	s.DetectResultList = v
	return s
}

type BatchContentSyncDetectResponseBodyDataDetectResultList struct {
	RiskInfo *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty"`
	// example:
	//
	// 1
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s BatchContentSyncDetectResponseBodyDataDetectResultList) String() string {
	return tea.Prettify(s)
}

func (s BatchContentSyncDetectResponseBodyDataDetectResultList) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectResponseBodyDataDetectResultList) SetRiskInfo(v string) *BatchContentSyncDetectResponseBodyDataDetectResultList {
	s.RiskInfo = &v
	return s
}

func (s *BatchContentSyncDetectResponseBodyDataDetectResultList) SetRiskResult(v int32) *BatchContentSyncDetectResponseBodyDataDetectResultList {
	s.RiskResult = &v
	return s
}

type BatchContentSyncDetectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchContentSyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchContentSyncDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchContentSyncDetectResponse) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectResponse) SetHeaders(v map[string]*string) *BatchContentSyncDetectResponse {
	s.Headers = v
	return s
}

func (s *BatchContentSyncDetectResponse) SetStatusCode(v int32) *BatchContentSyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchContentSyncDetectResponse) SetBody(v *BatchContentSyncDetectResponseBody) *BatchContentSyncDetectResponse {
	s.Body = v
	return s
}

type CheckAccountRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountRequest) SetRegionId(v string) *CheckAccountRequest {
	s.RegionId = &v
	return s
}

type CheckAccountResponseBody struct {
	// example:
	//
	// 00000
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CheckAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountResponseBody) SetCode(v string) *CheckAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAccountResponseBody) SetData(v *CheckAccountResponseBodyData) *CheckAccountResponseBody {
	s.Data = v
	return s
}

func (s *CheckAccountResponseBody) SetHttpStatusCode(v int32) *CheckAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckAccountResponseBody) SetMessage(v string) *CheckAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAccountResponseBody) SetRequestId(v string) *CheckAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAccountResponseBody) SetSuccess(v bool) *CheckAccountResponseBody {
	s.Success = &v
	return s
}

type CheckAccountResponseBodyData struct {
	// example:
	//
	// 1
	CheckResult *int32 `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
}

func (s CheckAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckAccountResponseBodyData) SetCheckResult(v int32) *CheckAccountResponseBodyData {
	s.CheckResult = &v
	return s
}

type CheckAccountResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountResponse) SetHeaders(v map[string]*string) *CheckAccountResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountResponse) SetStatusCode(v int32) *CheckAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountResponse) SetBody(v *CheckAccountResponseBody) *CheckAccountResponse {
	s.Body = v
	return s
}

type ContentAsyncDetectRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName      *string                                    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceParameter *ContentAsyncDetectRequestServiceParameter `json:"serviceParameter,omitempty" xml:"serviceParameter,omitempty" type:"Struct"`
}

func (s ContentAsyncDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s ContentAsyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectRequest) SetRegionId(v string) *ContentAsyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ContentAsyncDetectRequest) SetSceneName(v string) *ContentAsyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ContentAsyncDetectRequest) SetServiceName(v string) *ContentAsyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *ContentAsyncDetectRequest) SetServiceParameter(v *ContentAsyncDetectRequestServiceParameter) *ContentAsyncDetectRequest {
	s.ServiceParameter = v
	return s
}

type ContentAsyncDetectRequestServiceParameter struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s ContentAsyncDetectRequestServiceParameter) String() string {
	return tea.Prettify(s)
}

func (s ContentAsyncDetectRequestServiceParameter) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectRequestServiceParameter) SetContent(v string) *ContentAsyncDetectRequestServiceParameter {
	s.Content = &v
	return s
}

type ContentAsyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ContentAsyncDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ContentAsyncDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContentAsyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectResponseBody) SetCode(v string) *ContentAsyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetData(v *ContentAsyncDetectResponseBodyData) *ContentAsyncDetectResponseBody {
	s.Data = v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetHttpStatusCode(v int32) *ContentAsyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetMessage(v string) *ContentAsyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetRequestId(v string) *ContentAsyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContentAsyncDetectResponseBody) SetSuccess(v bool) *ContentAsyncDetectResponseBody {
	s.Success = &v
	return s
}

type ContentAsyncDetectResponseBodyData struct {
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ContentAsyncDetectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ContentAsyncDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectResponseBodyData) SetTaskId(v string) *ContentAsyncDetectResponseBodyData {
	s.TaskId = &v
	return s
}

type ContentAsyncDetectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContentAsyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContentAsyncDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s ContentAsyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ContentAsyncDetectResponse) SetHeaders(v map[string]*string) *ContentAsyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ContentAsyncDetectResponse) SetStatusCode(v int32) *ContentAsyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ContentAsyncDetectResponse) SetBody(v *ContentAsyncDetectResponseBody) *ContentAsyncDetectResponse {
	s.Body = v
	return s
}

type ContentSyncDetectRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName      *string                                   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceParameter *ContentSyncDetectRequestServiceParameter `json:"serviceParameter,omitempty" xml:"serviceParameter,omitempty" type:"Struct"`
}

func (s ContentSyncDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s ContentSyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectRequest) SetRegionId(v string) *ContentSyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ContentSyncDetectRequest) SetSceneName(v string) *ContentSyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ContentSyncDetectRequest) SetServiceName(v string) *ContentSyncDetectRequest {
	s.ServiceName = &v
	return s
}

func (s *ContentSyncDetectRequest) SetServiceParameter(v *ContentSyncDetectRequestServiceParameter) *ContentSyncDetectRequest {
	s.ServiceParameter = v
	return s
}

type ContentSyncDetectRequestServiceParameter struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s ContentSyncDetectRequestServiceParameter) String() string {
	return tea.Prettify(s)
}

func (s ContentSyncDetectRequestServiceParameter) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectRequestServiceParameter) SetContent(v string) *ContentSyncDetectRequestServiceParameter {
	s.Content = &v
	return s
}

type ContentSyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ContentSyncDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ContentSyncDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContentSyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectResponseBody) SetCode(v string) *ContentSyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ContentSyncDetectResponseBody) SetData(v *ContentSyncDetectResponseBodyData) *ContentSyncDetectResponseBody {
	s.Data = v
	return s
}

func (s *ContentSyncDetectResponseBody) SetHttpStatusCode(v int32) *ContentSyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ContentSyncDetectResponseBody) SetMessage(v string) *ContentSyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ContentSyncDetectResponseBody) SetRequestId(v string) *ContentSyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContentSyncDetectResponseBody) SetSuccess(v bool) *ContentSyncDetectResponseBody {
	s.Success = &v
	return s
}

type ContentSyncDetectResponseBodyData struct {
	RiskInfo *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty"`
	// example:
	//
	// 1
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ContentSyncDetectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ContentSyncDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectResponseBodyData) SetRiskInfo(v string) *ContentSyncDetectResponseBodyData {
	s.RiskInfo = &v
	return s
}

func (s *ContentSyncDetectResponseBodyData) SetRiskResult(v int32) *ContentSyncDetectResponseBodyData {
	s.RiskResult = &v
	return s
}

type ContentSyncDetectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContentSyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContentSyncDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s ContentSyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectResponse) SetHeaders(v map[string]*string) *ContentSyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ContentSyncDetectResponse) SetStatusCode(v int32) *ContentSyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ContentSyncDetectResponse) SetBody(v *ContentSyncDetectResponseBody) *ContentSyncDetectResponse {
	s.Body = v
	return s
}

type GetContentDetectResultRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetContentDetectResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetContentDetectResultRequest) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultRequest) SetRegionId(v string) *GetContentDetectResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetContentDetectResultRequest) SetTaskId(v string) *GetContentDetectResultRequest {
	s.TaskId = &v
	return s
}

type GetContentDetectResultResponseBody struct {
	// example:
	//
	// 00000
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetContentDetectResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContentDetectResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContentDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultResponseBody) SetCode(v string) *GetContentDetectResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetContentDetectResultResponseBody) SetData(v *GetContentDetectResultResponseBodyData) *GetContentDetectResultResponseBody {
	s.Data = v
	return s
}

func (s *GetContentDetectResultResponseBody) SetHttpStatusCode(v int32) *GetContentDetectResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContentDetectResultResponseBody) SetMessage(v string) *GetContentDetectResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetContentDetectResultResponseBody) SetRequestId(v string) *GetContentDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContentDetectResultResponseBody) SetSuccess(v bool) *GetContentDetectResultResponseBody {
	s.Success = &v
	return s
}

type GetContentDetectResultResponseBodyData struct {
	DetectResultList []*GetContentDetectResultResponseBodyDataDetectResultList `json:"DetectResultList,omitempty" xml:"DetectResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	ProcessedCount *int32  `json:"ProcessedCount,omitempty" xml:"ProcessedCount,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetContentDetectResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetContentDetectResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultResponseBodyData) SetDetectResultList(v []*GetContentDetectResultResponseBodyDataDetectResultList) *GetContentDetectResultResponseBodyData {
	s.DetectResultList = v
	return s
}

func (s *GetContentDetectResultResponseBodyData) SetProcessedCount(v int32) *GetContentDetectResultResponseBodyData {
	s.ProcessedCount = &v
	return s
}

func (s *GetContentDetectResultResponseBodyData) SetTaskId(v string) *GetContentDetectResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetContentDetectResultResponseBodyData) SetTaskStatus(v int32) *GetContentDetectResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetContentDetectResultResponseBodyData) SetTotalCount(v int32) *GetContentDetectResultResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetContentDetectResultResponseBodyDataDetectResultList struct {
	RiskInfo *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty"`
	// example:
	//
	// 1
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetContentDetectResultResponseBodyDataDetectResultList) String() string {
	return tea.Prettify(s)
}

func (s GetContentDetectResultResponseBodyDataDetectResultList) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) SetRiskInfo(v string) *GetContentDetectResultResponseBodyDataDetectResultList {
	s.RiskInfo = &v
	return s
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) SetRiskResult(v int32) *GetContentDetectResultResponseBodyDataDetectResultList {
	s.RiskResult = &v
	return s
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) SetStatus(v int32) *GetContentDetectResultResponseBodyDataDetectResultList {
	s.Status = &v
	return s
}

type GetContentDetectResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContentDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContentDetectResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContentDetectResultResponse) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultResponse) SetHeaders(v map[string]*string) *GetContentDetectResultResponse {
	s.Headers = v
	return s
}

func (s *GetContentDetectResultResponse) SetStatusCode(v int32) *GetContentDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContentDetectResultResponse) SetBody(v *GetContentDetectResultResponseBody) *GetContentDetectResultResponse {
	s.Body = v
	return s
}

type GetModelInputContentDetectResultRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetModelInputContentDetectResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultRequest) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultRequest) SetRegionId(v string) *GetModelInputContentDetectResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetModelInputContentDetectResultRequest) SetTaskId(v string) *GetModelInputContentDetectResultRequest {
	s.TaskId = &v
	return s
}

type GetModelInputContentDetectResultResponseBody struct {
	// example:
	//
	// 00000
	Code             *string                                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	DetectResultList []*GetModelInputContentDetectResultResponseBodyDetectResultList `json:"DetectResultList,omitempty" xml:"DetectResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	ProcessedCount *int32 `json:"ProcessedCount,omitempty" xml:"ProcessedCount,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBody) SetCode(v string) *GetModelInputContentDetectResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetDetectResultList(v []*GetModelInputContentDetectResultResponseBodyDetectResultList) *GetModelInputContentDetectResultResponseBody {
	s.DetectResultList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetHttpStatusCode(v int32) *GetModelInputContentDetectResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetMessage(v string) *GetModelInputContentDetectResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetProcessedCount(v int32) *GetModelInputContentDetectResultResponseBody {
	s.ProcessedCount = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetRequestId(v string) *GetModelInputContentDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetSuccess(v bool) *GetModelInputContentDetectResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetTaskId(v string) *GetModelInputContentDetectResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetTaskStatus(v int32) *GetModelInputContentDetectResultResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBody) SetTotalCount(v int32) *GetModelInputContentDetectResultResponseBody {
	s.TotalCount = &v
	return s
}

type GetModelInputContentDetectResultResponseBodyDetectResultList struct {
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 2
	Status    *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	TraceInfo *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Struct"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultList) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultList {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) SetStatus(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultList {
	s.Status = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultList) SetTraceInfo(v *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) *GetModelInputContentDetectResultResponseBodyDetectResultList {
	s.TraceInfo = v
	return s
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo struct {
	BlockWord         *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord         `json:"BlockWord,omitempty" xml:"BlockWord,omitempty" type:"Struct"`
	HarmfulCategories *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories `json:"HarmfulCategories,omitempty" xml:"HarmfulCategories,omitempty" type:"Struct"`
	PromptAttack      *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack      `json:"PromptAttack,omitempty" xml:"PromptAttack,omitempty" type:"Struct"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) SetBlockWord(v *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.BlockWord = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) SetHarmfulCategories(v *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.HarmfulCategories = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo) SetPromptAttack(v *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.PromptAttack = v
	return s
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord struct {
	BlockWordGroupInfoList []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList `json:"BlockWordGroupInfoList,omitempty" xml:"BlockWordGroupInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) SetBlockWordGroupInfoList(v []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	s.BlockWordGroupInfoList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	s.RiskResult = &v
	return s
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList struct {
	BlockWordList []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList `json:"BlockWordList,omitempty" xml:"BlockWordList,omitempty" type:"Repeated"`
	GroupName     *string                                                                                                              `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) SetBlockWordList(v []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	s.BlockWordList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) SetGroupName(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	s.GroupName = &v
	return s
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList struct {
	Word      *string `json:"Word,omitempty" xml:"Word,omitempty"`
	WordLabel *string `json:"WordLabel,omitempty" xml:"WordLabel,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWord(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.Word = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWordLabel(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.WordLabel = &v
	return s
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories struct {
	// example:
	//
	// 0.85
	ConfidenceScore         *float64                                                                                                         `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	HarmfulCategoryInfoList []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList `json:"HarmfulCategoryInfoList,omitempty" xml:"HarmfulCategoryInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetConfidenceScore(v float64) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.ConfidenceScore = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetHarmfulCategoryInfoList(v []*GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.HarmfulCategoryInfoList = v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.RiskResult = &v
	return s
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList struct {
	CategoryLabel *string `json:"CategoryLabel,omitempty" xml:"CategoryLabel,omitempty"`
	// example:
	//
	// 0
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 0
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryLabel(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryType(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSecurityLevel(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SecurityLevel = &v
	return s
}

type GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack struct {
	PromptAttackInfo *string `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 0
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetPromptAttackInfo(v string) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.PromptAttackInfo = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetRiskResult(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.RiskResult = &v
	return s
}

func (s *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetSecurityLevel(v int32) *GetModelInputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.SecurityLevel = &v
	return s
}

type GetModelInputContentDetectResultResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelInputContentDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelInputContentDetectResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelInputContentDetectResultResponse) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultResponse) SetHeaders(v map[string]*string) *GetModelInputContentDetectResultResponse {
	s.Headers = v
	return s
}

func (s *GetModelInputContentDetectResultResponse) SetStatusCode(v int32) *GetModelInputContentDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelInputContentDetectResultResponse) SetBody(v *GetModelInputContentDetectResultResponseBody) *GetModelInputContentDetectResultResponse {
	s.Body = v
	return s
}

type GetModelOutputContentDetectResultRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetModelOutputContentDetectResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultRequest) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultRequest) SetRegionId(v string) *GetModelOutputContentDetectResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetModelOutputContentDetectResultRequest) SetTaskId(v string) *GetModelOutputContentDetectResultRequest {
	s.TaskId = &v
	return s
}

type GetModelOutputContentDetectResultResponseBody struct {
	// example:
	//
	// 00000
	Code             *string                                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	DetectResultList []*GetModelOutputContentDetectResultResponseBodyDetectResultList `json:"DetectResultList,omitempty" xml:"DetectResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	ProcessedCount *int32 `json:"ProcessedCount,omitempty" xml:"ProcessedCount,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBody) SetCode(v string) *GetModelOutputContentDetectResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetDetectResultList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultList) *GetModelOutputContentDetectResultResponseBody {
	s.DetectResultList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetHttpStatusCode(v int32) *GetModelOutputContentDetectResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetMessage(v string) *GetModelOutputContentDetectResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetProcessedCount(v int32) *GetModelOutputContentDetectResultResponseBody {
	s.ProcessedCount = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetRequestId(v string) *GetModelOutputContentDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetSuccess(v bool) *GetModelOutputContentDetectResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetTaskId(v string) *GetModelOutputContentDetectResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetTaskStatus(v int32) *GetModelOutputContentDetectResultResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBody) SetTotalCount(v int32) *GetModelOutputContentDetectResultResponseBody {
	s.TotalCount = &v
	return s
}

type GetModelOutputContentDetectResultResponseBodyDetectResultList struct {
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 2
	Status    *int32                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TraceInfo *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Struct"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultList) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultList {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) SetStatus(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultList {
	s.Status = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultList) SetTraceInfo(v *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) *GetModelOutputContentDetectResultResponseBodyDetectResultList {
	s.TraceInfo = v
	return s
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo struct {
	BlockWord         *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord         `json:"BlockWord,omitempty" xml:"BlockWord,omitempty" type:"Struct"`
	HarmfulCategories *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories `json:"HarmfulCategories,omitempty" xml:"HarmfulCategories,omitempty" type:"Struct"`
	PromptAttack      *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack      `json:"PromptAttack,omitempty" xml:"PromptAttack,omitempty" type:"Struct"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) SetBlockWord(v *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.BlockWord = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) SetHarmfulCategories(v *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.HarmfulCategories = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo) SetPromptAttack(v *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfo {
	s.PromptAttack = v
	return s
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord struct {
	BlockWordGroupInfoList []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList `json:"BlockWordGroupInfoList,omitempty" xml:"BlockWordGroupInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) SetBlockWordGroupInfoList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	s.BlockWordGroupInfoList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWord {
	s.RiskResult = &v
	return s
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList struct {
	BlockWordList []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList `json:"BlockWordList,omitempty" xml:"BlockWordList,omitempty" type:"Repeated"`
	GroupName     *string                                                                                                               `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) SetBlockWordList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	s.BlockWordList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList) SetGroupName(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoList {
	s.GroupName = &v
	return s
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList struct {
	Word      *string `json:"Word,omitempty" xml:"Word,omitempty"`
	WordLabel *string `json:"WordLabel,omitempty" xml:"WordLabel,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWord(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.Word = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWordLabel(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.WordLabel = &v
	return s
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories struct {
	// example:
	//
	// 0.85
	ConfidenceScore         *float64                                                                                                          `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	HarmfulCategoryInfoList []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList `json:"HarmfulCategoryInfoList,omitempty" xml:"HarmfulCategoryInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetConfidenceScore(v float64) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.ConfidenceScore = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetHarmfulCategoryInfoList(v []*GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.HarmfulCategoryInfoList = v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategories {
	s.RiskResult = &v
	return s
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList struct {
	CategoryLabel *string `json:"CategoryLabel,omitempty" xml:"CategoryLabel,omitempty"`
	// example:
	//
	// 0
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryLabel(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryType(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryType = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSecurityLevel(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SecurityLevel = &v
	return s
}

type GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack struct {
	PromptAttackInfo *string `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetPromptAttackInfo(v string) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.PromptAttackInfo = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetRiskResult(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.RiskResult = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack) SetSecurityLevel(v int32) *GetModelOutputContentDetectResultResponseBodyDetectResultListTraceInfoPromptAttack {
	s.SecurityLevel = &v
	return s
}

type GetModelOutputContentDetectResultResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelOutputContentDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelOutputContentDetectResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelOutputContentDetectResultResponse) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultResponse) SetHeaders(v map[string]*string) *GetModelOutputContentDetectResultResponse {
	s.Headers = v
	return s
}

func (s *GetModelOutputContentDetectResultResponse) SetStatusCode(v int32) *GetModelOutputContentDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelOutputContentDetectResultResponse) SetBody(v *GetModelOutputContentDetectResultResponseBody) *GetModelOutputContentDetectResultResponse {
	s.Body = v
	return s
}

type ListSensitiveWordRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordRequest) SetPageNumber(v int32) *ListSensitiveWordRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSensitiveWordRequest) SetPageSize(v int32) *ListSensitiveWordRequest {
	s.PageSize = &v
	return s
}

func (s *ListSensitiveWordRequest) SetRegionId(v string) *ListSensitiveWordRequest {
	s.RegionId = &v
	return s
}

type ListSensitiveWordResponseBody struct {
	// example:
	//
	// 00000
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListSensitiveWordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponseBody) SetCode(v string) *ListSensitiveWordResponseBody {
	s.Code = &v
	return s
}

func (s *ListSensitiveWordResponseBody) SetData(v *ListSensitiveWordResponseBodyData) *ListSensitiveWordResponseBody {
	s.Data = v
	return s
}

func (s *ListSensitiveWordResponseBody) SetHttpStatusCode(v int32) *ListSensitiveWordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSensitiveWordResponseBody) SetMessage(v string) *ListSensitiveWordResponseBody {
	s.Message = &v
	return s
}

func (s *ListSensitiveWordResponseBody) SetRequestId(v string) *ListSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitiveWordResponseBody) SetSuccess(v bool) *ListSensitiveWordResponseBody {
	s.Success = &v
	return s
}

type ListSensitiveWordResponseBodyData struct {
	// example:
	//
	// 10000
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize          *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SensitiveWordList []*ListSensitiveWordResponseBodyDataSensitiveWordList `json:"SensitiveWordList,omitempty" xml:"SensitiveWordList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSensitiveWordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponseBodyData) SetMaxCount(v int32) *ListSensitiveWordResponseBodyData {
	s.MaxCount = &v
	return s
}

func (s *ListSensitiveWordResponseBodyData) SetPageNumber(v int32) *ListSensitiveWordResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSensitiveWordResponseBodyData) SetPageSize(v int32) *ListSensitiveWordResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSensitiveWordResponseBodyData) SetSensitiveWordList(v []*ListSensitiveWordResponseBodyDataSensitiveWordList) *ListSensitiveWordResponseBodyData {
	s.SensitiveWordList = v
	return s
}

func (s *ListSensitiveWordResponseBodyData) SetTotalCount(v int32) *ListSensitiveWordResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSensitiveWordResponseBodyDataSensitiveWordList struct {
	// example:
	//
	// 387907
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Word  *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s ListSensitiveWordResponseBodyDataSensitiveWordList) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponseBodyDataSensitiveWordList) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponseBodyDataSensitiveWordList) SetId(v int64) *ListSensitiveWordResponseBodyDataSensitiveWordList {
	s.Id = &v
	return s
}

func (s *ListSensitiveWordResponseBodyDataSensitiveWordList) SetLabel(v string) *ListSensitiveWordResponseBodyDataSensitiveWordList {
	s.Label = &v
	return s
}

func (s *ListSensitiveWordResponseBodyDataSensitiveWordList) SetWord(v string) *ListSensitiveWordResponseBodyDataSensitiveWordList {
	s.Word = &v
	return s
}

type ListSensitiveWordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponse) SetHeaders(v map[string]*string) *ListSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *ListSensitiveWordResponse) SetStatusCode(v int32) *ListSensitiveWordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSensitiveWordResponse) SetBody(v *ListSensitiveWordResponseBody) *ListSensitiveWordResponse {
	s.Body = v
	return s
}

type ModelInputContentAsyncDetectRequest struct {
	BodyData *ModelInputContentAsyncDetectRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModelInputContentAsyncDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentAsyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectRequest) SetBodyData(v *ModelInputContentAsyncDetectRequestBodyData) *ModelInputContentAsyncDetectRequest {
	s.BodyData = v
	return s
}

func (s *ModelInputContentAsyncDetectRequest) SetPolicyIdentifier(v string) *ModelInputContentAsyncDetectRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelInputContentAsyncDetectRequest) SetRegionId(v string) *ModelInputContentAsyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ModelInputContentAsyncDetectRequest) SetSceneName(v string) *ModelInputContentAsyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ModelInputContentAsyncDetectRequest) SetServiceName(v string) *ModelInputContentAsyncDetectRequest {
	s.ServiceName = &v
	return s
}

type ModelInputContentAsyncDetectRequestBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ModelInputContentAsyncDetectRequestBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentAsyncDetectRequestBodyData) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectRequestBodyData) SetContent(v string) *ModelInputContentAsyncDetectRequestBodyData {
	s.Content = &v
	return s
}

type ModelInputContentAsyncDetectShrinkRequest struct {
	BodyDataShrink *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModelInputContentAsyncDetectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentAsyncDetectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetBodyDataShrink(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetPolicyIdentifier(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetRegionId(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetSceneName(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.SceneName = &v
	return s
}

func (s *ModelInputContentAsyncDetectShrinkRequest) SetServiceName(v string) *ModelInputContentAsyncDetectShrinkRequest {
	s.ServiceName = &v
	return s
}

type ModelInputContentAsyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModelInputContentAsyncDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentAsyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectResponseBody) SetCode(v string) *ModelInputContentAsyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetHttpStatusCode(v int32) *ModelInputContentAsyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetMessage(v string) *ModelInputContentAsyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetRequestId(v string) *ModelInputContentAsyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetSuccess(v bool) *ModelInputContentAsyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetTaskId(v string) *ModelInputContentAsyncDetectResponseBody {
	s.TaskId = &v
	return s
}

type ModelInputContentAsyncDetectResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelInputContentAsyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelInputContentAsyncDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentAsyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectResponse) SetHeaders(v map[string]*string) *ModelInputContentAsyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ModelInputContentAsyncDetectResponse) SetStatusCode(v int32) *ModelInputContentAsyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponse) SetBody(v *ModelInputContentAsyncDetectResponseBody) *ModelInputContentAsyncDetectResponse {
	s.Body = v
	return s
}

type ModelInputContentSyncDetectRequest struct {
	BodyData         *ModelInputContentSyncDetectRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
	PolicyIdentifier *string                                     `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	RegionId         *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SceneName        *string                                     `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	ServiceName      *string                                     `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModelInputContentSyncDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectRequest) SetBodyData(v *ModelInputContentSyncDetectRequestBodyData) *ModelInputContentSyncDetectRequest {
	s.BodyData = v
	return s
}

func (s *ModelInputContentSyncDetectRequest) SetPolicyIdentifier(v string) *ModelInputContentSyncDetectRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelInputContentSyncDetectRequest) SetRegionId(v string) *ModelInputContentSyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ModelInputContentSyncDetectRequest) SetSceneName(v string) *ModelInputContentSyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ModelInputContentSyncDetectRequest) SetServiceName(v string) *ModelInputContentSyncDetectRequest {
	s.ServiceName = &v
	return s
}

type ModelInputContentSyncDetectRequestBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ModelInputContentSyncDetectRequestBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectRequestBodyData) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectRequestBodyData) SetContent(v string) *ModelInputContentSyncDetectRequestBodyData {
	s.Content = &v
	return s
}

type ModelInputContentSyncDetectShrinkRequest struct {
	BodyDataShrink   *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SceneName        *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	ServiceName      *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModelInputContentSyncDetectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetBodyDataShrink(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetPolicyIdentifier(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetRegionId(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetSceneName(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.SceneName = &v
	return s
}

func (s *ModelInputContentSyncDetectShrinkRequest) SetServiceName(v string) *ModelInputContentSyncDetectShrinkRequest {
	s.ServiceName = &v
	return s
}

type ModelInputContentSyncDetectResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskResult *int32                                            `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	Success    *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceInfo  *ModelInputContentSyncDetectResponseBodyTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Struct"`
}

func (s ModelInputContentSyncDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBody) SetCode(v string) *ModelInputContentSyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetHttpStatusCode(v int32) *ModelInputContentSyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetMessage(v string) *ModelInputContentSyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetRequestId(v string) *ModelInputContentSyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBody {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetSuccess(v bool) *ModelInputContentSyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBody) SetTraceInfo(v *ModelInputContentSyncDetectResponseBodyTraceInfo) *ModelInputContentSyncDetectResponseBody {
	s.TraceInfo = v
	return s
}

type ModelInputContentSyncDetectResponseBodyTraceInfo struct {
	BlockWord         *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord         `json:"BlockWord,omitempty" xml:"BlockWord,omitempty" type:"Struct"`
	HarmfulCategories *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories `json:"HarmfulCategories,omitempty" xml:"HarmfulCategories,omitempty" type:"Struct"`
	PromptAttack      *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack      `json:"PromptAttack,omitempty" xml:"PromptAttack,omitempty" type:"Struct"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfo) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfo) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) SetBlockWord(v *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) *ModelInputContentSyncDetectResponseBodyTraceInfo {
	s.BlockWord = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) SetHarmfulCategories(v *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) *ModelInputContentSyncDetectResponseBodyTraceInfo {
	s.HarmfulCategories = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfo) SetPromptAttack(v *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) *ModelInputContentSyncDetectResponseBodyTraceInfo {
	s.PromptAttack = v
	return s
}

type ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord struct {
	BlockWordGroupInfoList []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList `json:"BlockWordGroupInfoList,omitempty" xml:"BlockWordGroupInfoList,omitempty" type:"Repeated"`
	RiskResult             *int32                                                                             `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) SetBlockWordGroupInfoList(v []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord {
	s.BlockWordGroupInfoList = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWord {
	s.RiskResult = &v
	return s
}

type ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList struct {
	BlockWordList []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList `json:"BlockWordList,omitempty" xml:"BlockWordList,omitempty" type:"Repeated"`
	GroupName     *string                                                                                         `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) SetBlockWordList(v []*ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	s.BlockWordList = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) SetGroupName(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	s.GroupName = &v
	return s
}

type ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList struct {
	Word      *string `json:"Word,omitempty" xml:"Word,omitempty"`
	WordLabel *string `json:"WordLabel,omitempty" xml:"WordLabel,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWord(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.Word = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWordLabel(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.WordLabel = &v
	return s
}

type ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories struct {
	ConfidenceScore         *float64                                                                                    `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	HarmfulCategoryInfoList []*ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList `json:"HarmfulCategoryInfoList,omitempty" xml:"HarmfulCategoryInfoList,omitempty" type:"Repeated"`
	RiskResult              *int32                                                                                      `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetConfidenceScore(v float64) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.ConfidenceScore = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetHarmfulCategoryInfoList(v []*ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.HarmfulCategoryInfoList = v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.RiskResult = &v
	return s
}

type ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList struct {
	CategoryLabel *string `json:"CategoryLabel,omitempty" xml:"CategoryLabel,omitempty"`
	CategoryType  *int32  `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	RiskResult    *int32  `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	SecurityLevel *int32  `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryLabel(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryType(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryType = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSecurityLevel(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SecurityLevel = &v
	return s
}

type ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack struct {
	ConfidenceScore  *float64 `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	PromptAttackInfo *string  `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty"`
	RiskResult       *int32   `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	SecurityLevel    *int32   `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetConfidenceScore(v float64) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.ConfidenceScore = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetPromptAttackInfo(v string) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.PromptAttackInfo = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetRiskResult(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.RiskResult = &v
	return s
}

func (s *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetSecurityLevel(v int32) *ModelInputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.SecurityLevel = &v
	return s
}

type ModelInputContentSyncDetectResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelInputContentSyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelInputContentSyncDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModelInputContentSyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponse) SetHeaders(v map[string]*string) *ModelInputContentSyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ModelInputContentSyncDetectResponse) SetStatusCode(v int32) *ModelInputContentSyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelInputContentSyncDetectResponse) SetBody(v *ModelInputContentSyncDetectResponseBody) *ModelInputContentSyncDetectResponse {
	s.Body = v
	return s
}

type ModelOutputContentAsyncDetectRequest struct {
	BodyData *ModelOutputContentAsyncDetectRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModelOutputContentAsyncDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentAsyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectRequest) SetBodyData(v *ModelOutputContentAsyncDetectRequestBodyData) *ModelOutputContentAsyncDetectRequest {
	s.BodyData = v
	return s
}

func (s *ModelOutputContentAsyncDetectRequest) SetPolicyIdentifier(v string) *ModelOutputContentAsyncDetectRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelOutputContentAsyncDetectRequest) SetRegionId(v string) *ModelOutputContentAsyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ModelOutputContentAsyncDetectRequest) SetSceneName(v string) *ModelOutputContentAsyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ModelOutputContentAsyncDetectRequest) SetServiceName(v string) *ModelOutputContentAsyncDetectRequest {
	s.ServiceName = &v
	return s
}

type ModelOutputContentAsyncDetectRequestBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ModelOutputContentAsyncDetectRequestBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentAsyncDetectRequestBodyData) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectRequestBodyData) SetContent(v string) *ModelOutputContentAsyncDetectRequestBodyData {
	s.Content = &v
	return s
}

type ModelOutputContentAsyncDetectShrinkRequest struct {
	BodyDataShrink *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModelOutputContentAsyncDetectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentAsyncDetectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetBodyDataShrink(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetPolicyIdentifier(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetRegionId(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetSceneName(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.SceneName = &v
	return s
}

func (s *ModelOutputContentAsyncDetectShrinkRequest) SetServiceName(v string) *ModelOutputContentAsyncDetectShrinkRequest {
	s.ServiceName = &v
	return s
}

type ModelOutputContentAsyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModelOutputContentAsyncDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentAsyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetCode(v string) *ModelOutputContentAsyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetHttpStatusCode(v int32) *ModelOutputContentAsyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetMessage(v string) *ModelOutputContentAsyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetRequestId(v string) *ModelOutputContentAsyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetSuccess(v bool) *ModelOutputContentAsyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetTaskId(v string) *ModelOutputContentAsyncDetectResponseBody {
	s.TaskId = &v
	return s
}

type ModelOutputContentAsyncDetectResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelOutputContentAsyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelOutputContentAsyncDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentAsyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectResponse) SetHeaders(v map[string]*string) *ModelOutputContentAsyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ModelOutputContentAsyncDetectResponse) SetStatusCode(v int32) *ModelOutputContentAsyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponse) SetBody(v *ModelOutputContentAsyncDetectResponseBody) *ModelOutputContentAsyncDetectResponse {
	s.Body = v
	return s
}

type ModelOutputContentSyncDetectRequest struct {
	BodyData *ModelOutputContentSyncDetectRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModelOutputContentSyncDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectRequest) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectRequest) SetBodyData(v *ModelOutputContentSyncDetectRequestBodyData) *ModelOutputContentSyncDetectRequest {
	s.BodyData = v
	return s
}

func (s *ModelOutputContentSyncDetectRequest) SetPolicyIdentifier(v string) *ModelOutputContentSyncDetectRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelOutputContentSyncDetectRequest) SetRegionId(v string) *ModelOutputContentSyncDetectRequest {
	s.RegionId = &v
	return s
}

func (s *ModelOutputContentSyncDetectRequest) SetSceneName(v string) *ModelOutputContentSyncDetectRequest {
	s.SceneName = &v
	return s
}

func (s *ModelOutputContentSyncDetectRequest) SetServiceName(v string) *ModelOutputContentSyncDetectRequest {
	s.ServiceName = &v
	return s
}

type ModelOutputContentSyncDetectRequestBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ModelOutputContentSyncDetectRequestBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectRequestBodyData) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectRequestBodyData) SetContent(v string) *ModelOutputContentSyncDetectRequestBodyData {
	s.Content = &v
	return s
}

type ModelOutputContentSyncDetectShrinkRequest struct {
	BodyDataShrink *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ""
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// textDetection
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModelOutputContentSyncDetectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetBodyDataShrink(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetPolicyIdentifier(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.PolicyIdentifier = &v
	return s
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetRegionId(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetSceneName(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.SceneName = &v
	return s
}

func (s *ModelOutputContentSyncDetectShrinkRequest) SetServiceName(v string) *ModelOutputContentSyncDetectShrinkRequest {
	s.ServiceName = &v
	return s
}

type ModelOutputContentSyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// “”
	RiskInfo *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// True
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceInfo *ModelOutputContentSyncDetectResponseBodyTraceInfo `json:"TraceInfo,omitempty" xml:"TraceInfo,omitempty" type:"Struct"`
}

func (s ModelOutputContentSyncDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBody) SetCode(v string) *ModelOutputContentSyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetHttpStatusCode(v int32) *ModelOutputContentSyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetMessage(v string) *ModelOutputContentSyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetRequestId(v string) *ModelOutputContentSyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetRiskInfo(v string) *ModelOutputContentSyncDetectResponseBody {
	s.RiskInfo = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBody {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetSuccess(v bool) *ModelOutputContentSyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBody) SetTraceInfo(v *ModelOutputContentSyncDetectResponseBodyTraceInfo) *ModelOutputContentSyncDetectResponseBody {
	s.TraceInfo = v
	return s
}

type ModelOutputContentSyncDetectResponseBodyTraceInfo struct {
	BlockWord         *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord         `json:"BlockWord,omitempty" xml:"BlockWord,omitempty" type:"Struct"`
	HarmfulCategories *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories `json:"HarmfulCategories,omitempty" xml:"HarmfulCategories,omitempty" type:"Struct"`
	PromptAttack      *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack      `json:"PromptAttack,omitempty" xml:"PromptAttack,omitempty" type:"Struct"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfo) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfo) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) SetBlockWord(v *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) *ModelOutputContentSyncDetectResponseBodyTraceInfo {
	s.BlockWord = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) SetHarmfulCategories(v *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) *ModelOutputContentSyncDetectResponseBodyTraceInfo {
	s.HarmfulCategories = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfo) SetPromptAttack(v *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) *ModelOutputContentSyncDetectResponseBodyTraceInfo {
	s.PromptAttack = v
	return s
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord struct {
	BlockWordGroupInfoList []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList `json:"BlockWordGroupInfoList,omitempty" xml:"BlockWordGroupInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) SetBlockWordGroupInfoList(v []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord {
	s.BlockWordGroupInfoList = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWord {
	s.RiskResult = &v
	return s
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList struct {
	BlockWordList []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList `json:"BlockWordList,omitempty" xml:"BlockWordList,omitempty" type:"Repeated"`
	GroupName     *string                                                                                          `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) SetBlockWordList(v []*ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	s.BlockWordList = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList) SetGroupName(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoList {
	s.GroupName = &v
	return s
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList struct {
	Word      *string `json:"Word,omitempty" xml:"Word,omitempty"`
	WordLabel *string `json:"WordLabel,omitempty" xml:"WordLabel,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWord(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.Word = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList) SetWordLabel(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoBlockWordBlockWordGroupInfoListBlockWordList {
	s.WordLabel = &v
	return s
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories struct {
	// example:
	//
	// 0.85
	ConfidenceScore         *float64                                                                                     `json:"ConfidenceScore,omitempty" xml:"ConfidenceScore,omitempty"`
	HarmfulCategoryInfoList []*ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList `json:"HarmfulCategoryInfoList,omitempty" xml:"HarmfulCategoryInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetConfidenceScore(v float64) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.ConfidenceScore = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetHarmfulCategoryInfoList(v []*ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.HarmfulCategoryInfoList = v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategories {
	s.RiskResult = &v
	return s
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList struct {
	CategoryLabel *string `json:"CategoryLabel,omitempty" xml:"CategoryLabel,omitempty"`
	// example:
	//
	// 0
	CategoryType *int32 `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 1
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryLabel(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryLabel = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetCategoryType(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.CategoryType = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList) SetSecurityLevel(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoHarmfulCategoriesHarmfulCategoryInfoList {
	s.SecurityLevel = &v
	return s
}

type ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack struct {
	PromptAttackInfo *string `json:"PromptAttackInfo,omitempty" xml:"PromptAttackInfo,omitempty"`
	// example:
	//
	// 0
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 0
	SecurityLevel *int32 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetPromptAttackInfo(v string) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.PromptAttackInfo = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetRiskResult(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.RiskResult = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack) SetSecurityLevel(v int32) *ModelOutputContentSyncDetectResponseBodyTraceInfoPromptAttack {
	s.SecurityLevel = &v
	return s
}

type ModelOutputContentSyncDetectResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelOutputContentSyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelOutputContentSyncDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponse) SetHeaders(v map[string]*string) *ModelOutputContentSyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ModelOutputContentSyncDetectResponse) SetStatusCode(v int32) *ModelOutputContentSyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponse) SetBody(v *ModelOutputContentSyncDetectResponseBody) *ModelOutputContentSyncDetectResponse {
	s.Body = v
	return s
}

type RegisterAccountRequest struct {
	// example:
	//
	// "user api register"
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RegisterAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterAccountRequest) GoString() string {
	return s.String()
}

func (s *RegisterAccountRequest) SetMemo(v string) *RegisterAccountRequest {
	s.Memo = &v
	return s
}

func (s *RegisterAccountRequest) SetRegionId(v string) *RegisterAccountRequest {
	s.RegionId = &v
	return s
}

type RegisterAccountResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterAccountResponseBody) SetCode(v string) *RegisterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterAccountResponseBody) SetHttpStatusCode(v int32) *RegisterAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RegisterAccountResponseBody) SetMessage(v string) *RegisterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterAccountResponseBody) SetRequestId(v string) *RegisterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterAccountResponseBody) SetSuccess(v bool) *RegisterAccountResponseBody {
	s.Success = &v
	return s
}

type RegisterAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterAccountResponse) GoString() string {
	return s.String()
}

func (s *RegisterAccountResponse) SetHeaders(v map[string]*string) *RegisterAccountResponse {
	s.Headers = v
	return s
}

func (s *RegisterAccountResponse) SetStatusCode(v int32) *RegisterAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterAccountResponse) SetBody(v *RegisterAccountResponseBody) *RegisterAccountResponse {
	s.Body = v
	return s
}

type SyncSensitiveWordRequest struct {
	BodyData *SyncSensitiveWordRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
	// example:
	//
	// true
	Commit *bool `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SyncSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *SyncSensitiveWordRequest) SetBodyData(v *SyncSensitiveWordRequestBodyData) *SyncSensitiveWordRequest {
	s.BodyData = v
	return s
}

func (s *SyncSensitiveWordRequest) SetCommit(v bool) *SyncSensitiveWordRequest {
	s.Commit = &v
	return s
}

func (s *SyncSensitiveWordRequest) SetRegionId(v string) *SyncSensitiveWordRequest {
	s.RegionId = &v
	return s
}

type SyncSensitiveWordRequestBodyData struct {
	SensitiveWordList []*SyncSensitiveWordRequestBodyDataSensitiveWordList `json:"SensitiveWordList,omitempty" xml:"SensitiveWordList,omitempty" type:"Repeated"`
}

func (s SyncSensitiveWordRequestBodyData) String() string {
	return tea.Prettify(s)
}

func (s SyncSensitiveWordRequestBodyData) GoString() string {
	return s.String()
}

func (s *SyncSensitiveWordRequestBodyData) SetSensitiveWordList(v []*SyncSensitiveWordRequestBodyDataSensitiveWordList) *SyncSensitiveWordRequestBodyData {
	s.SensitiveWordList = v
	return s
}

type SyncSensitiveWordRequestBodyDataSensitiveWordList struct {
	// example:
	//
	// 341806
	Id     *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Label  *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Word   *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s SyncSensitiveWordRequestBodyDataSensitiveWordList) String() string {
	return tea.Prettify(s)
}

func (s SyncSensitiveWordRequestBodyDataSensitiveWordList) GoString() string {
	return s.String()
}

func (s *SyncSensitiveWordRequestBodyDataSensitiveWordList) SetId(v int32) *SyncSensitiveWordRequestBodyDataSensitiveWordList {
	s.Id = &v
	return s
}

func (s *SyncSensitiveWordRequestBodyDataSensitiveWordList) SetLabel(v string) *SyncSensitiveWordRequestBodyDataSensitiveWordList {
	s.Label = &v
	return s
}

func (s *SyncSensitiveWordRequestBodyDataSensitiveWordList) SetStatus(v int32) *SyncSensitiveWordRequestBodyDataSensitiveWordList {
	s.Status = &v
	return s
}

func (s *SyncSensitiveWordRequestBodyDataSensitiveWordList) SetWord(v string) *SyncSensitiveWordRequestBodyDataSensitiveWordList {
	s.Word = &v
	return s
}

type SyncSensitiveWordShrinkRequest struct {
	BodyDataShrink *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
	// example:
	//
	// true
	Commit *bool `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SyncSensitiveWordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncSensitiveWordShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncSensitiveWordShrinkRequest) SetBodyDataShrink(v string) *SyncSensitiveWordShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *SyncSensitiveWordShrinkRequest) SetCommit(v bool) *SyncSensitiveWordShrinkRequest {
	s.Commit = &v
	return s
}

func (s *SyncSensitiveWordShrinkRequest) SetRegionId(v string) *SyncSensitiveWordShrinkRequest {
	s.RegionId = &v
	return s
}

type SyncSensitiveWordResponseBody struct {
	// example:
	//
	// 00000
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SyncSensitiveWordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *SyncSensitiveWordResponseBody) SetCode(v string) *SyncSensitiveWordResponseBody {
	s.Code = &v
	return s
}

func (s *SyncSensitiveWordResponseBody) SetData(v *SyncSensitiveWordResponseBodyData) *SyncSensitiveWordResponseBody {
	s.Data = v
	return s
}

func (s *SyncSensitiveWordResponseBody) SetHttpStatusCode(v int32) *SyncSensitiveWordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncSensitiveWordResponseBody) SetMessage(v string) *SyncSensitiveWordResponseBody {
	s.Message = &v
	return s
}

func (s *SyncSensitiveWordResponseBody) SetRequestId(v string) *SyncSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncSensitiveWordResponseBody) SetSuccess(v bool) *SyncSensitiveWordResponseBody {
	s.Success = &v
	return s
}

type SyncSensitiveWordResponseBodyData struct {
	SensitiveWordErrorList []*SyncSensitiveWordResponseBodyDataSensitiveWordErrorList `json:"SensitiveWordErrorList,omitempty" xml:"SensitiveWordErrorList,omitempty" type:"Repeated"`
}

func (s SyncSensitiveWordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SyncSensitiveWordResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyncSensitiveWordResponseBodyData) SetSensitiveWordErrorList(v []*SyncSensitiveWordResponseBodyDataSensitiveWordErrorList) *SyncSensitiveWordResponseBodyData {
	s.SensitiveWordErrorList = v
	return s
}

type SyncSensitiveWordResponseBodyDataSensitiveWordErrorList struct {
	// example:
	//
	// ""
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 1
	ErrStatus *int32 `json:"ErrStatus,omitempty" xml:"ErrStatus,omitempty"`
	// example:
	//
	// 0
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// contraband
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Word  *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s SyncSensitiveWordResponseBodyDataSensitiveWordErrorList) String() string {
	return tea.Prettify(s)
}

func (s SyncSensitiveWordResponseBodyDataSensitiveWordErrorList) GoString() string {
	return s.String()
}

func (s *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList) SetErrMessage(v string) *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList {
	s.ErrMessage = &v
	return s
}

func (s *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList) SetErrStatus(v int32) *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList {
	s.ErrStatus = &v
	return s
}

func (s *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList) SetIndex(v int64) *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList {
	s.Index = &v
	return s
}

func (s *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList) SetLabel(v string) *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList {
	s.Label = &v
	return s
}

func (s *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList) SetWord(v string) *SyncSensitiveWordResponseBodyDataSensitiveWordErrorList {
	s.Word = &v
	return s
}

type SyncSensitiveWordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *SyncSensitiveWordResponse) SetHeaders(v map[string]*string) *SyncSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *SyncSensitiveWordResponse) SetStatusCode(v int32) *SyncSensitiveWordResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncSensitiveWordResponse) SetBody(v *SyncSensitiveWordResponseBody) *SyncSensitiveWordResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("rai"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # BatchContentAsyncDetect
//
// @param request - BatchContentAsyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchContentAsyncDetectResponse
func (client *Client) BatchContentAsyncDetectWithOptions(request *BatchContentAsyncDetectRequest, runtime *util.RuntimeOptions) (_result *BatchContentAsyncDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		query["SceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceParameterList)) {
		body["serviceParameterList"] = request.ServiceParameterList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchContentAsyncDetect"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BatchContentAsyncDetectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BatchContentAsyncDetectResponse{}
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
// # BatchContentAsyncDetect
//
// @param request - BatchContentAsyncDetectRequest
//
// @return BatchContentAsyncDetectResponse
func (client *Client) BatchContentAsyncDetect(request *BatchContentAsyncDetectRequest) (_result *BatchContentAsyncDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchContentAsyncDetectResponse{}
	_body, _err := client.BatchContentAsyncDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # BatchContentSyncDetect
//
// @param request - BatchContentSyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchContentSyncDetectResponse
func (client *Client) BatchContentSyncDetectWithOptions(request *BatchContentSyncDetectRequest, runtime *util.RuntimeOptions) (_result *BatchContentSyncDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		query["SceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceParameterList)) {
		body["serviceParameterList"] = request.ServiceParameterList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchContentSyncDetect"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BatchContentSyncDetectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BatchContentSyncDetectResponse{}
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
// # BatchContentSyncDetect
//
// @param request - BatchContentSyncDetectRequest
//
// @return BatchContentSyncDetectResponse
func (client *Client) BatchContentSyncDetect(request *BatchContentSyncDetectRequest) (_result *BatchContentSyncDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchContentSyncDetectResponse{}
	_body, _err := client.BatchContentSyncDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查用户是否开通RAI服务
//
// @param request - CheckAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountResponse
func (client *Client) CheckAccountWithOptions(request *CheckAccountRequest, runtime *util.RuntimeOptions) (_result *CheckAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAccount"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CheckAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CheckAccountResponse{}
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
// 检查用户是否开通RAI服务
//
// @param request - CheckAccountRequest
//
// @return CheckAccountResponse
func (client *Client) CheckAccount(request *CheckAccountRequest) (_result *CheckAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAccountResponse{}
	_body, _err := client.CheckAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ContentAsyncDetect
//
// @param request - ContentAsyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContentAsyncDetectResponse
func (client *Client) ContentAsyncDetectWithOptions(request *ContentAsyncDetectRequest, runtime *util.RuntimeOptions) (_result *ContentAsyncDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		query["SceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceParameter)) {
		body["serviceParameter"] = request.ServiceParameter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ContentAsyncDetect"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ContentAsyncDetectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ContentAsyncDetectResponse{}
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
// # ContentAsyncDetect
//
// @param request - ContentAsyncDetectRequest
//
// @return ContentAsyncDetectResponse
func (client *Client) ContentAsyncDetect(request *ContentAsyncDetectRequest) (_result *ContentAsyncDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContentAsyncDetectResponse{}
	_body, _err := client.ContentAsyncDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ContentSyncDetect
//
// @param request - ContentSyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContentSyncDetectResponse
func (client *Client) ContentSyncDetectWithOptions(request *ContentSyncDetectRequest, runtime *util.RuntimeOptions) (_result *ContentSyncDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		query["SceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceParameter)) {
		body["serviceParameter"] = request.ServiceParameter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ContentSyncDetect"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ContentSyncDetectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ContentSyncDetectResponse{}
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
// # ContentSyncDetect
//
// @param request - ContentSyncDetectRequest
//
// @return ContentSyncDetectResponse
func (client *Client) ContentSyncDetect(request *ContentSyncDetectRequest) (_result *ContentSyncDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContentSyncDetectResponse{}
	_body, _err := client.ContentSyncDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetContentDetectResult
//
// @param request - GetContentDetectResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContentDetectResultResponse
func (client *Client) GetContentDetectResultWithOptions(request *GetContentDetectResultRequest, runtime *util.RuntimeOptions) (_result *GetContentDetectResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetContentDetectResult"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetContentDetectResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetContentDetectResultResponse{}
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
// # GetContentDetectResult
//
// @param request - GetContentDetectResultRequest
//
// @return GetContentDetectResultResponse
func (client *Client) GetContentDetectResult(request *GetContentDetectResultRequest) (_result *GetContentDetectResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetContentDetectResultResponse{}
	_body, _err := client.GetContentDetectResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetModelInputContentDetectResult
//
// @param request - GetModelInputContentDetectResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelInputContentDetectResultResponse
func (client *Client) GetModelInputContentDetectResultWithOptions(request *GetModelInputContentDetectResultRequest, runtime *util.RuntimeOptions) (_result *GetModelInputContentDetectResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetModelInputContentDetectResult"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetModelInputContentDetectResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetModelInputContentDetectResultResponse{}
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
// # GetModelInputContentDetectResult
//
// @param request - GetModelInputContentDetectResultRequest
//
// @return GetModelInputContentDetectResultResponse
func (client *Client) GetModelInputContentDetectResult(request *GetModelInputContentDetectResultRequest) (_result *GetModelInputContentDetectResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModelInputContentDetectResultResponse{}
	_body, _err := client.GetModelInputContentDetectResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetModelOutputContentDetectResult
//
// @param request - GetModelOutputContentDetectResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelOutputContentDetectResultResponse
func (client *Client) GetModelOutputContentDetectResultWithOptions(request *GetModelOutputContentDetectResultRequest, runtime *util.RuntimeOptions) (_result *GetModelOutputContentDetectResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetModelOutputContentDetectResult"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetModelOutputContentDetectResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetModelOutputContentDetectResultResponse{}
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
// # GetModelOutputContentDetectResult
//
// @param request - GetModelOutputContentDetectResultRequest
//
// @return GetModelOutputContentDetectResultResponse
func (client *Client) GetModelOutputContentDetectResult(request *GetModelOutputContentDetectResultRequest) (_result *GetModelOutputContentDetectResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModelOutputContentDetectResultResponse{}
	_body, _err := client.GetModelOutputContentDetectResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListSensitiveWord
//
// @param request - ListSensitiveWordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSensitiveWordResponse
func (client *Client) ListSensitiveWordWithOptions(request *ListSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *ListSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSensitiveWord"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSensitiveWordResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSensitiveWordResponse{}
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
// # ListSensitiveWord
//
// @param request - ListSensitiveWordRequest
//
// @return ListSensitiveWordResponse
func (client *Client) ListSensitiveWord(request *ListSensitiveWordRequest) (_result *ListSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSensitiveWordResponse{}
	_body, _err := client.ListSensitiveWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ModelInputContentAsyncDetect
//
// @param tmpReq - ModelInputContentAsyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelInputContentAsyncDetectResponse
func (client *Client) ModelInputContentAsyncDetectWithOptions(tmpReq *ModelInputContentAsyncDetectRequest, runtime *util.RuntimeOptions) (_result *ModelInputContentAsyncDetectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModelInputContentAsyncDetectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BodyData)) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, tea.String("BodyData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyIdentifier)) {
		query["PolicyIdentifier"] = request.PolicyIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		query["SceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyDataShrink)) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModelInputContentAsyncDetect"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModelInputContentAsyncDetectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModelInputContentAsyncDetectResponse{}
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
// # ModelInputContentAsyncDetect
//
// @param request - ModelInputContentAsyncDetectRequest
//
// @return ModelInputContentAsyncDetectResponse
func (client *Client) ModelInputContentAsyncDetect(request *ModelInputContentAsyncDetectRequest) (_result *ModelInputContentAsyncDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModelInputContentAsyncDetectResponse{}
	_body, _err := client.ModelInputContentAsyncDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ModelInputContentSyncDetect
//
// @param tmpReq - ModelInputContentSyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelInputContentSyncDetectResponse
func (client *Client) ModelInputContentSyncDetectWithOptions(tmpReq *ModelInputContentSyncDetectRequest, runtime *util.RuntimeOptions) (_result *ModelInputContentSyncDetectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModelInputContentSyncDetectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BodyData)) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, tea.String("BodyData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyIdentifier)) {
		query["PolicyIdentifier"] = request.PolicyIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		query["SceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyDataShrink)) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModelInputContentSyncDetect"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModelInputContentSyncDetectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModelInputContentSyncDetectResponse{}
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
// # ModelInputContentSyncDetect
//
// @param request - ModelInputContentSyncDetectRequest
//
// @return ModelInputContentSyncDetectResponse
func (client *Client) ModelInputContentSyncDetect(request *ModelInputContentSyncDetectRequest) (_result *ModelInputContentSyncDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModelInputContentSyncDetectResponse{}
	_body, _err := client.ModelInputContentSyncDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ModelOutputContentAsyncDetect
//
// @param tmpReq - ModelOutputContentAsyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelOutputContentAsyncDetectResponse
func (client *Client) ModelOutputContentAsyncDetectWithOptions(tmpReq *ModelOutputContentAsyncDetectRequest, runtime *util.RuntimeOptions) (_result *ModelOutputContentAsyncDetectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModelOutputContentAsyncDetectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BodyData)) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, tea.String("BodyData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyIdentifier)) {
		query["PolicyIdentifier"] = request.PolicyIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		query["SceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyDataShrink)) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModelOutputContentAsyncDetect"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModelOutputContentAsyncDetectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModelOutputContentAsyncDetectResponse{}
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
// # ModelOutputContentAsyncDetect
//
// @param request - ModelOutputContentAsyncDetectRequest
//
// @return ModelOutputContentAsyncDetectResponse
func (client *Client) ModelOutputContentAsyncDetect(request *ModelOutputContentAsyncDetectRequest) (_result *ModelOutputContentAsyncDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModelOutputContentAsyncDetectResponse{}
	_body, _err := client.ModelOutputContentAsyncDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ModelOutputContentSyncDetect
//
// @param tmpReq - ModelOutputContentSyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelOutputContentSyncDetectResponse
func (client *Client) ModelOutputContentSyncDetectWithOptions(tmpReq *ModelOutputContentSyncDetectRequest, runtime *util.RuntimeOptions) (_result *ModelOutputContentSyncDetectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModelOutputContentSyncDetectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BodyData)) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, tea.String("BodyData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyIdentifier)) {
		query["PolicyIdentifier"] = request.PolicyIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		query["SceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyDataShrink)) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModelOutputContentSyncDetect"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModelOutputContentSyncDetectResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModelOutputContentSyncDetectResponse{}
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
// # ModelOutputContentSyncDetect
//
// @param request - ModelOutputContentSyncDetectRequest
//
// @return ModelOutputContentSyncDetectResponse
func (client *Client) ModelOutputContentSyncDetect(request *ModelOutputContentSyncDetectRequest) (_result *ModelOutputContentSyncDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModelOutputContentSyncDetectResponse{}
	_body, _err := client.ModelOutputContentSyncDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 注册RAI账号
//
// @param request - RegisterAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterAccountResponse
func (client *Client) RegisterAccountWithOptions(request *RegisterAccountRequest, runtime *util.RuntimeOptions) (_result *RegisterAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Memo)) {
		query["Memo"] = request.Memo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterAccount"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RegisterAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RegisterAccountResponse{}
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
// 注册RAI账号
//
// @param request - RegisterAccountRequest
//
// @return RegisterAccountResponse
func (client *Client) RegisterAccount(request *RegisterAccountRequest) (_result *RegisterAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterAccountResponse{}
	_body, _err := client.RegisterAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SyncSensitiveWord
//
// @param tmpReq - SyncSensitiveWordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncSensitiveWordResponse
func (client *Client) SyncSensitiveWordWithOptions(tmpReq *SyncSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *SyncSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SyncSensitiveWordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BodyData)) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, tea.String("BodyData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Commit)) {
		query["Commit"] = request.Commit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyDataShrink)) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncSensitiveWord"),
		Version:     tea.String("2024-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SyncSensitiveWordResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SyncSensitiveWordResponse{}
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
// # SyncSensitiveWord
//
// @param request - SyncSensitiveWordRequest
//
// @return SyncSensitiveWordResponse
func (client *Client) SyncSensitiveWord(request *SyncSensitiveWordRequest) (_result *SyncSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncSensitiveWordResponse{}
	_body, _err := client.SyncSensitiveWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
