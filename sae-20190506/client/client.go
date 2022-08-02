// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AbortAndRollbackChangeOrderRequest struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortAndRollbackChangeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortAndRollbackChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderRequest) SetChangeOrderId(v string) *AbortAndRollbackChangeOrderRequest {
	s.ChangeOrderId = &v
	return s
}

type AbortAndRollbackChangeOrderResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AbortAndRollbackChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                      `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s AbortAndRollbackChangeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetCode(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetData(v *AbortAndRollbackChangeOrderResponseBodyData) *AbortAndRollbackChangeOrderResponseBody {
	s.Data = v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetErrorCode(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetMessage(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetRequestId(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetSuccess(v bool) *AbortAndRollbackChangeOrderResponseBody {
	s.Success = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetTraceId(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

type AbortAndRollbackChangeOrderResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortAndRollbackChangeOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponseBodyData) SetChangeOrderId(v string) *AbortAndRollbackChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type AbortAndRollbackChangeOrderResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AbortAndRollbackChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortAndRollbackChangeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponse) SetHeaders(v map[string]*string) *AbortAndRollbackChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *AbortAndRollbackChangeOrderResponse) SetStatusCode(v int32) *AbortAndRollbackChangeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponse) SetBody(v *AbortAndRollbackChangeOrderResponseBody) *AbortAndRollbackChangeOrderResponse {
	s.Body = v
	return s
}

type AbortChangeOrderRequest struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortChangeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderRequest) SetChangeOrderId(v string) *AbortChangeOrderRequest {
	s.ChangeOrderId = &v
	return s
}

type AbortChangeOrderResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AbortChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s AbortChangeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponseBody) SetCode(v string) *AbortChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetData(v *AbortChangeOrderResponseBodyData) *AbortChangeOrderResponseBody {
	s.Data = v
	return s
}

func (s *AbortChangeOrderResponseBody) SetErrorCode(v string) *AbortChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetMessage(v string) *AbortChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetRequestId(v string) *AbortChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetSuccess(v bool) *AbortChangeOrderResponseBody {
	s.Success = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetTraceId(v string) *AbortChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

type AbortChangeOrderResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortChangeOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponseBodyData) SetChangeOrderId(v string) *AbortChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type AbortChangeOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AbortChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortChangeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponse) SetHeaders(v map[string]*string) *AbortChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *AbortChangeOrderResponse) SetStatusCode(v int32) *AbortChangeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortChangeOrderResponse) SetBody(v *AbortChangeOrderResponseBody) *AbortChangeOrderResponse {
	s.Body = v
	return s
}

type BatchStartApplicationsRequest struct {
	AppIds      *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s BatchStartApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStartApplicationsRequest) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsRequest) SetAppIds(v string) *BatchStartApplicationsRequest {
	s.AppIds = &v
	return s
}

func (s *BatchStartApplicationsRequest) SetNamespaceId(v string) *BatchStartApplicationsRequest {
	s.NamespaceId = &v
	return s
}

type BatchStartApplicationsResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchStartApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                 `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BatchStartApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStartApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsResponseBody) SetCode(v string) *BatchStartApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetData(v *BatchStartApplicationsResponseBodyData) *BatchStartApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetErrorCode(v string) *BatchStartApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetMessage(v string) *BatchStartApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetRequestId(v string) *BatchStartApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetSuccess(v bool) *BatchStartApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetTraceId(v string) *BatchStartApplicationsResponseBody {
	s.TraceId = &v
	return s
}

type BatchStartApplicationsResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BatchStartApplicationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchStartApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsResponseBodyData) SetChangeOrderId(v string) *BatchStartApplicationsResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type BatchStartApplicationsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchStartApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStartApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStartApplicationsResponse) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsResponse) SetHeaders(v map[string]*string) *BatchStartApplicationsResponse {
	s.Headers = v
	return s
}

func (s *BatchStartApplicationsResponse) SetStatusCode(v int32) *BatchStartApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStartApplicationsResponse) SetBody(v *BatchStartApplicationsResponseBody) *BatchStartApplicationsResponse {
	s.Body = v
	return s
}

type BatchStopApplicationsRequest struct {
	AppIds      *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s BatchStopApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStopApplicationsRequest) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsRequest) SetAppIds(v string) *BatchStopApplicationsRequest {
	s.AppIds = &v
	return s
}

func (s *BatchStopApplicationsRequest) SetNamespaceId(v string) *BatchStopApplicationsRequest {
	s.NamespaceId = &v
	return s
}

type BatchStopApplicationsResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchStopApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BatchStopApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStopApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsResponseBody) SetCode(v string) *BatchStopApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetData(v *BatchStopApplicationsResponseBodyData) *BatchStopApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetErrorCode(v string) *BatchStopApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetMessage(v string) *BatchStopApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetRequestId(v string) *BatchStopApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetSuccess(v bool) *BatchStopApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetTraceId(v string) *BatchStopApplicationsResponseBody {
	s.TraceId = &v
	return s
}

type BatchStopApplicationsResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BatchStopApplicationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchStopApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsResponseBodyData) SetChangeOrderId(v string) *BatchStopApplicationsResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type BatchStopApplicationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchStopApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStopApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStopApplicationsResponse) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsResponse) SetHeaders(v map[string]*string) *BatchStopApplicationsResponse {
	s.Headers = v
	return s
}

func (s *BatchStopApplicationsResponse) SetStatusCode(v int32) *BatchStopApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStopApplicationsResponse) SetBody(v *BatchStopApplicationsResponseBody) *BatchStopApplicationsResponse {
	s.Body = v
	return s
}

type BindSlbRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Internet      *string `json:"Internet,omitempty" xml:"Internet,omitempty"`
	InternetSlbId *string `json:"InternetSlbId,omitempty" xml:"InternetSlbId,omitempty"`
	Intranet      *string `json:"Intranet,omitempty" xml:"Intranet,omitempty"`
	IntranetSlbId *string `json:"IntranetSlbId,omitempty" xml:"IntranetSlbId,omitempty"`
}

func (s BindSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s BindSlbRequest) GoString() string {
	return s.String()
}

func (s *BindSlbRequest) SetAppId(v string) *BindSlbRequest {
	s.AppId = &v
	return s
}

func (s *BindSlbRequest) SetInternet(v string) *BindSlbRequest {
	s.Internet = &v
	return s
}

func (s *BindSlbRequest) SetInternetSlbId(v string) *BindSlbRequest {
	s.InternetSlbId = &v
	return s
}

func (s *BindSlbRequest) SetIntranet(v string) *BindSlbRequest {
	s.Intranet = &v
	return s
}

func (s *BindSlbRequest) SetIntranetSlbId(v string) *BindSlbRequest {
	s.IntranetSlbId = &v
	return s
}

type BindSlbResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BindSlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BindSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindSlbResponseBody) GoString() string {
	return s.String()
}

func (s *BindSlbResponseBody) SetCode(v string) *BindSlbResponseBody {
	s.Code = &v
	return s
}

func (s *BindSlbResponseBody) SetData(v *BindSlbResponseBodyData) *BindSlbResponseBody {
	s.Data = v
	return s
}

func (s *BindSlbResponseBody) SetErrorCode(v string) *BindSlbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BindSlbResponseBody) SetMessage(v string) *BindSlbResponseBody {
	s.Message = &v
	return s
}

func (s *BindSlbResponseBody) SetRequestId(v string) *BindSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindSlbResponseBody) SetSuccess(v bool) *BindSlbResponseBody {
	s.Success = &v
	return s
}

func (s *BindSlbResponseBody) SetTraceId(v string) *BindSlbResponseBody {
	s.TraceId = &v
	return s
}

type BindSlbResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BindSlbResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindSlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindSlbResponseBodyData) SetChangeOrderId(v string) *BindSlbResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type BindSlbResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s BindSlbResponse) GoString() string {
	return s.String()
}

func (s *BindSlbResponse) SetHeaders(v map[string]*string) *BindSlbResponse {
	s.Headers = v
	return s
}

func (s *BindSlbResponse) SetStatusCode(v int32) *BindSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSlbResponse) SetBody(v *BindSlbResponseBody) *BindSlbResponse {
	s.Body = v
	return s
}

type ConfirmPipelineBatchRequest struct {
	Confirm    *bool   `json:"Confirm,omitempty" xml:"Confirm,omitempty"`
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s ConfirmPipelineBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmPipelineBatchRequest) GoString() string {
	return s.String()
}

func (s *ConfirmPipelineBatchRequest) SetConfirm(v bool) *ConfirmPipelineBatchRequest {
	s.Confirm = &v
	return s
}

func (s *ConfirmPipelineBatchRequest) SetPipelineId(v string) *ConfirmPipelineBatchRequest {
	s.PipelineId = &v
	return s
}

type ConfirmPipelineBatchResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ConfirmPipelineBatchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                               `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ConfirmPipelineBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmPipelineBatchResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmPipelineBatchResponseBody) SetCode(v string) *ConfirmPipelineBatchResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetData(v *ConfirmPipelineBatchResponseBodyData) *ConfirmPipelineBatchResponseBody {
	s.Data = v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetErrorCode(v string) *ConfirmPipelineBatchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetMessage(v string) *ConfirmPipelineBatchResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetRequestId(v string) *ConfirmPipelineBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetSuccess(v bool) *ConfirmPipelineBatchResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmPipelineBatchResponseBody) SetTraceId(v string) *ConfirmPipelineBatchResponseBody {
	s.TraceId = &v
	return s
}

type ConfirmPipelineBatchResponseBodyData struct {
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s ConfirmPipelineBatchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ConfirmPipelineBatchResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfirmPipelineBatchResponseBodyData) SetPipelineId(v string) *ConfirmPipelineBatchResponseBodyData {
	s.PipelineId = &v
	return s
}

type ConfirmPipelineBatchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfirmPipelineBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmPipelineBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmPipelineBatchResponse) GoString() string {
	return s.String()
}

func (s *ConfirmPipelineBatchResponse) SetHeaders(v map[string]*string) *ConfirmPipelineBatchResponse {
	s.Headers = v
	return s
}

func (s *ConfirmPipelineBatchResponse) SetStatusCode(v int32) *ConfirmPipelineBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmPipelineBatchResponse) SetBody(v *ConfirmPipelineBatchResponseBody) *ConfirmPipelineBatchResponse {
	s.Body = v
	return s
}

type CreateApplicationRequest struct {
	AcrAssumeRoleArn              *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	AcrInstanceId                 *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	AppDescription                *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	AppName                       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AssociateEip                  *bool   `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	AutoConfig                    *bool   `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	Command                       *string `json:"Command,omitempty" xml:"Command,omitempty"`
	CommandArgs                   *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	ConfigMapMountDesc            *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	Cpu                           *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CustomHostAlias               *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	Deploy                        *bool   `json:"Deploy,omitempty" xml:"Deploy,omitempty"`
	EdasContainerVersion          *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	Envs                          *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	ImageUrl                      *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	JarStartArgs                  *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	JarStartOptions               *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	Jdk                           *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	KafkaConfigs                  *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	Liveness                      *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	Memory                        *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	MountDesc                     *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	MountHost                     *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	NamespaceId                   *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NasId                         *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	OssAkId                       *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	OssAkSecret                   *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	OssMountDescs                 *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	PackageType                   *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PackageUrl                    *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	PackageVersion                *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	PhpArmsConfigLocation         *string `json:"PhpArmsConfigLocation,omitempty" xml:"PhpArmsConfigLocation,omitempty"`
	PhpConfig                     *string `json:"PhpConfig,omitempty" xml:"PhpConfig,omitempty"`
	PhpConfigLocation             *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	PostStart                     *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	PreStop                       *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	ProgrammingLanguage           *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	Readiness                     *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	Replicas                      *int32  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	SecurityGroupId               *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SlsConfigs                    *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	TerminationGracePeriodSeconds *int32  `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	Timezone                      *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	TomcatConfig                  *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	VSwitchId                     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WarStartOptions               *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	WebContainer                  *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetAcrAssumeRoleArn(v string) *CreateApplicationRequest {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *CreateApplicationRequest) SetAcrInstanceId(v string) *CreateApplicationRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateApplicationRequest) SetAppDescription(v string) *CreateApplicationRequest {
	s.AppDescription = &v
	return s
}

func (s *CreateApplicationRequest) SetAppName(v string) *CreateApplicationRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationRequest) SetAssociateEip(v bool) *CreateApplicationRequest {
	s.AssociateEip = &v
	return s
}

func (s *CreateApplicationRequest) SetAutoConfig(v bool) *CreateApplicationRequest {
	s.AutoConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetCommand(v string) *CreateApplicationRequest {
	s.Command = &v
	return s
}

func (s *CreateApplicationRequest) SetCommandArgs(v string) *CreateApplicationRequest {
	s.CommandArgs = &v
	return s
}

func (s *CreateApplicationRequest) SetConfigMapMountDesc(v string) *CreateApplicationRequest {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *CreateApplicationRequest) SetCpu(v int32) *CreateApplicationRequest {
	s.Cpu = &v
	return s
}

func (s *CreateApplicationRequest) SetCustomHostAlias(v string) *CreateApplicationRequest {
	s.CustomHostAlias = &v
	return s
}

func (s *CreateApplicationRequest) SetDeploy(v bool) *CreateApplicationRequest {
	s.Deploy = &v
	return s
}

func (s *CreateApplicationRequest) SetEdasContainerVersion(v string) *CreateApplicationRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetEnvs(v string) *CreateApplicationRequest {
	s.Envs = &v
	return s
}

func (s *CreateApplicationRequest) SetImageUrl(v string) *CreateApplicationRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateApplicationRequest) SetJarStartArgs(v string) *CreateApplicationRequest {
	s.JarStartArgs = &v
	return s
}

func (s *CreateApplicationRequest) SetJarStartOptions(v string) *CreateApplicationRequest {
	s.JarStartOptions = &v
	return s
}

func (s *CreateApplicationRequest) SetJdk(v string) *CreateApplicationRequest {
	s.Jdk = &v
	return s
}

func (s *CreateApplicationRequest) SetKafkaConfigs(v string) *CreateApplicationRequest {
	s.KafkaConfigs = &v
	return s
}

func (s *CreateApplicationRequest) SetLiveness(v string) *CreateApplicationRequest {
	s.Liveness = &v
	return s
}

func (s *CreateApplicationRequest) SetMemory(v int32) *CreateApplicationRequest {
	s.Memory = &v
	return s
}

func (s *CreateApplicationRequest) SetMountDesc(v string) *CreateApplicationRequest {
	s.MountDesc = &v
	return s
}

func (s *CreateApplicationRequest) SetMountHost(v string) *CreateApplicationRequest {
	s.MountHost = &v
	return s
}

func (s *CreateApplicationRequest) SetNamespaceId(v string) *CreateApplicationRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateApplicationRequest) SetNasId(v string) *CreateApplicationRequest {
	s.NasId = &v
	return s
}

func (s *CreateApplicationRequest) SetOssAkId(v string) *CreateApplicationRequest {
	s.OssAkId = &v
	return s
}

func (s *CreateApplicationRequest) SetOssAkSecret(v string) *CreateApplicationRequest {
	s.OssAkSecret = &v
	return s
}

func (s *CreateApplicationRequest) SetOssMountDescs(v string) *CreateApplicationRequest {
	s.OssMountDescs = &v
	return s
}

func (s *CreateApplicationRequest) SetPackageType(v string) *CreateApplicationRequest {
	s.PackageType = &v
	return s
}

func (s *CreateApplicationRequest) SetPackageUrl(v string) *CreateApplicationRequest {
	s.PackageUrl = &v
	return s
}

func (s *CreateApplicationRequest) SetPackageVersion(v string) *CreateApplicationRequest {
	s.PackageVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetPhpArmsConfigLocation(v string) *CreateApplicationRequest {
	s.PhpArmsConfigLocation = &v
	return s
}

func (s *CreateApplicationRequest) SetPhpConfig(v string) *CreateApplicationRequest {
	s.PhpConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetPhpConfigLocation(v string) *CreateApplicationRequest {
	s.PhpConfigLocation = &v
	return s
}

func (s *CreateApplicationRequest) SetPostStart(v string) *CreateApplicationRequest {
	s.PostStart = &v
	return s
}

func (s *CreateApplicationRequest) SetPreStop(v string) *CreateApplicationRequest {
	s.PreStop = &v
	return s
}

func (s *CreateApplicationRequest) SetProgrammingLanguage(v string) *CreateApplicationRequest {
	s.ProgrammingLanguage = &v
	return s
}

func (s *CreateApplicationRequest) SetReadiness(v string) *CreateApplicationRequest {
	s.Readiness = &v
	return s
}

func (s *CreateApplicationRequest) SetReplicas(v int32) *CreateApplicationRequest {
	s.Replicas = &v
	return s
}

func (s *CreateApplicationRequest) SetSecurityGroupId(v string) *CreateApplicationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetSlsConfigs(v string) *CreateApplicationRequest {
	s.SlsConfigs = &v
	return s
}

func (s *CreateApplicationRequest) SetTerminationGracePeriodSeconds(v int32) *CreateApplicationRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *CreateApplicationRequest) SetTimezone(v string) *CreateApplicationRequest {
	s.Timezone = &v
	return s
}

func (s *CreateApplicationRequest) SetTomcatConfig(v string) *CreateApplicationRequest {
	s.TomcatConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetVSwitchId(v string) *CreateApplicationRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateApplicationRequest) SetVpcId(v string) *CreateApplicationRequest {
	s.VpcId = &v
	return s
}

func (s *CreateApplicationRequest) SetWarStartOptions(v string) *CreateApplicationRequest {
	s.WarStartOptions = &v
	return s
}

func (s *CreateApplicationRequest) SetWebContainer(v string) *CreateApplicationRequest {
	s.WebContainer = &v
	return s
}

type CreateApplicationResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                            `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) SetCode(v string) *CreateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationResponseBody) SetData(v *CreateApplicationResponseBodyData) *CreateApplicationResponseBody {
	s.Data = v
	return s
}

func (s *CreateApplicationResponseBody) SetErrorCode(v string) *CreateApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateApplicationResponseBody) SetMessage(v string) *CreateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetSuccess(v bool) *CreateApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApplicationResponseBody) SetTraceId(v string) *CreateApplicationResponseBody {
	s.TraceId = &v
	return s
}

type CreateApplicationResponseBodyData struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s CreateApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyData) SetAppId(v string) *CreateApplicationResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetChangeOrderId(v string) *CreateApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type CreateApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetStatusCode(v int32) *CreateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

type CreateApplicationScalingRuleRequest struct {
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	MinReadyInstanceRatio *int32  `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	MinReadyInstances     *int32  `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	ScalingRuleEnable     *bool   `json:"ScalingRuleEnable,omitempty" xml:"ScalingRuleEnable,omitempty"`
	ScalingRuleMetric     *string `json:"ScalingRuleMetric,omitempty" xml:"ScalingRuleMetric,omitempty"`
	ScalingRuleName       *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	ScalingRuleTimer      *string `json:"ScalingRuleTimer,omitempty" xml:"ScalingRuleTimer,omitempty"`
	ScalingRuleType       *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
}

func (s CreateApplicationScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleRequest) SetAppId(v string) *CreateApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetMinReadyInstanceRatio(v int32) *CreateApplicationScalingRuleRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetMinReadyInstances(v int32) *CreateApplicationScalingRuleRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleEnable(v bool) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleEnable = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleMetric(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleMetric = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleName(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleTimer(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleTimer = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleType(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleType = &v
	return s
}

type CreateApplicationScalingRuleResponseBody struct {
	Data      *CreateApplicationScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string                                       `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBody) SetData(v *CreateApplicationScalingRuleResponseBodyData) *CreateApplicationScalingRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetRequestId(v string) *CreateApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetTraceId(v string) *CreateApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

type CreateApplicationScalingRuleResponseBodyData struct {
	AppId            *string                                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime       *int64                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastDisableTime  *int64                                              `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	Metric           *CreateApplicationScalingRuleResponseBodyDataMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	ScaleRuleEnabled *bool                                               `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	ScaleRuleName    *string                                             `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	ScaleRuleType    *string                                             `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	Timer            *CreateApplicationScalingRuleResponseBodyDataTimer  `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	UpdateTime       *int64                                              `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetAppId(v string) *CreateApplicationScalingRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetCreateTime(v int64) *CreateApplicationScalingRuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetLastDisableTime(v int64) *CreateApplicationScalingRuleResponseBodyData {
	s.LastDisableTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetMetric(v *CreateApplicationScalingRuleResponseBodyDataMetric) *CreateApplicationScalingRuleResponseBodyData {
	s.Metric = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetScaleRuleEnabled(v bool) *CreateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetScaleRuleName(v string) *CreateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleName = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetScaleRuleType(v string) *CreateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleType = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetTimer(v *CreateApplicationScalingRuleResponseBodyDataTimer) *CreateApplicationScalingRuleResponseBodyData {
	s.Timer = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetUpdateTime(v int64) *CreateApplicationScalingRuleResponseBodyData {
	s.UpdateTime = &v
	return s
}

type CreateApplicationScalingRuleResponseBodyDataMetric struct {
	MaxReplicas *int32                                                       `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	Metrics     []*CreateApplicationScalingRuleResponseBodyDataMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	MinReplicas *int32                                                       `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyDataMetric) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyDataMetric) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) SetMaxReplicas(v int32) *CreateApplicationScalingRuleResponseBodyDataMetric {
	s.MaxReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) SetMetrics(v []*CreateApplicationScalingRuleResponseBodyDataMetricMetrics) *CreateApplicationScalingRuleResponseBodyDataMetric {
	s.Metrics = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) SetMinReplicas(v int32) *CreateApplicationScalingRuleResponseBodyDataMetric {
	s.MinReplicas = &v
	return s
}

type CreateApplicationScalingRuleResponseBodyDataMetricMetrics struct {
	MetricTargetAverageUtilization *int32  `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	MetricType                     *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyDataMetricMetrics) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyDataMetricMetrics) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricTargetAverageUtilization(v int32) *CreateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricType(v string) *CreateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricType = &v
	return s
}

type CreateApplicationScalingRuleResponseBodyDataTimer struct {
	BeginDate *string                                                       `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate   *string                                                       `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Period    *string                                                       `json:"Period,omitempty" xml:"Period,omitempty"`
	Schedules []*CreateApplicationScalingRuleResponseBodyDataTimerSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
}

func (s CreateApplicationScalingRuleResponseBodyDataTimer) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyDataTimer) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) SetBeginDate(v string) *CreateApplicationScalingRuleResponseBodyDataTimer {
	s.BeginDate = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) SetEndDate(v string) *CreateApplicationScalingRuleResponseBodyDataTimer {
	s.EndDate = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) SetPeriod(v string) *CreateApplicationScalingRuleResponseBodyDataTimer {
	s.Period = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) SetSchedules(v []*CreateApplicationScalingRuleResponseBodyDataTimerSchedules) *CreateApplicationScalingRuleResponseBodyDataTimer {
	s.Schedules = v
	return s
}

type CreateApplicationScalingRuleResponseBodyDataTimerSchedules struct {
	AtTime         *string `json:"AtTime,omitempty" xml:"AtTime,omitempty"`
	TargetReplicas *int32  `json:"TargetReplicas,omitempty" xml:"TargetReplicas,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyDataTimerSchedules) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyDataTimerSchedules) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) SetAtTime(v string) *CreateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.AtTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) SetTargetReplicas(v int32) *CreateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.TargetReplicas = &v
	return s
}

type CreateApplicationScalingRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *CreateApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationScalingRuleResponse) SetStatusCode(v int32) *CreateApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationScalingRuleResponse) SetBody(v *CreateApplicationScalingRuleResponseBody) *CreateApplicationScalingRuleResponse {
	s.Body = v
	return s
}

type CreateConfigMapRequest struct {
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s CreateConfigMapRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigMapRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigMapRequest) SetData(v string) *CreateConfigMapRequest {
	s.Data = &v
	return s
}

func (s *CreateConfigMapRequest) SetDescription(v string) *CreateConfigMapRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigMapRequest) SetName(v string) *CreateConfigMapRequest {
	s.Name = &v
	return s
}

func (s *CreateConfigMapRequest) SetNamespaceId(v string) *CreateConfigMapRequest {
	s.NamespaceId = &v
	return s
}

type CreateConfigMapResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateConfigMapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                          `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateConfigMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigMapResponseBody) SetCode(v string) *CreateConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetData(v *CreateConfigMapResponseBodyData) *CreateConfigMapResponseBody {
	s.Data = v
	return s
}

func (s *CreateConfigMapResponseBody) SetErrorCode(v string) *CreateConfigMapResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetMessage(v string) *CreateConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetRequestId(v string) *CreateConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetSuccess(v bool) *CreateConfigMapResponseBody {
	s.Success = &v
	return s
}

func (s *CreateConfigMapResponseBody) SetTraceId(v string) *CreateConfigMapResponseBody {
	s.TraceId = &v
	return s
}

type CreateConfigMapResponseBodyData struct {
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s CreateConfigMapResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigMapResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateConfigMapResponseBodyData) SetConfigMapId(v int64) *CreateConfigMapResponseBodyData {
	s.ConfigMapId = &v
	return s
}

type CreateConfigMapResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConfigMapResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigMapResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigMapResponse) SetHeaders(v map[string]*string) *CreateConfigMapResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigMapResponse) SetStatusCode(v int32) *CreateConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigMapResponse) SetBody(v *CreateConfigMapResponseBody) *CreateConfigMapResponse {
	s.Body = v
	return s
}

type CreateGreyTagRouteRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DubboRules  *string `json:"DubboRules,omitempty" xml:"DubboRules,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ScRules     *string `json:"ScRules,omitempty" xml:"ScRules,omitempty"`
}

func (s CreateGreyTagRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateGreyTagRouteRequest) SetAppId(v string) *CreateGreyTagRouteRequest {
	s.AppId = &v
	return s
}

func (s *CreateGreyTagRouteRequest) SetDescription(v string) *CreateGreyTagRouteRequest {
	s.Description = &v
	return s
}

func (s *CreateGreyTagRouteRequest) SetDubboRules(v string) *CreateGreyTagRouteRequest {
	s.DubboRules = &v
	return s
}

func (s *CreateGreyTagRouteRequest) SetName(v string) *CreateGreyTagRouteRequest {
	s.Name = &v
	return s
}

func (s *CreateGreyTagRouteRequest) SetScRules(v string) *CreateGreyTagRouteRequest {
	s.ScRules = &v
	return s
}

type CreateGreyTagRouteResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                             `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateGreyTagRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGreyTagRouteResponseBody) SetCode(v string) *CreateGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetData(v *CreateGreyTagRouteResponseBodyData) *CreateGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetErrorCode(v string) *CreateGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetMessage(v string) *CreateGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetRequestId(v string) *CreateGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetSuccess(v bool) *CreateGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGreyTagRouteResponseBody) SetTraceId(v string) *CreateGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

type CreateGreyTagRouteResponseBodyData struct {
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s CreateGreyTagRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateGreyTagRouteResponseBodyData) SetGreyTagRouteId(v int64) *CreateGreyTagRouteResponseBodyData {
	s.GreyTagRouteId = &v
	return s
}

type CreateGreyTagRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGreyTagRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateGreyTagRouteResponse) SetHeaders(v map[string]*string) *CreateGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateGreyTagRouteResponse) SetStatusCode(v int32) *CreateGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGreyTagRouteResponse) SetBody(v *CreateGreyTagRouteResponseBody) *CreateGreyTagRouteResponse {
	s.Body = v
	return s
}

type CreateIngressRequest struct {
	CertId           *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	DefaultRule      *string `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ListenerPort     *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalanceType  *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	NamespaceId      *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Rules            *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	SlbId            *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
}

func (s CreateIngressRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIngressRequest) GoString() string {
	return s.String()
}

func (s *CreateIngressRequest) SetCertId(v string) *CreateIngressRequest {
	s.CertId = &v
	return s
}

func (s *CreateIngressRequest) SetDefaultRule(v string) *CreateIngressRequest {
	s.DefaultRule = &v
	return s
}

func (s *CreateIngressRequest) SetDescription(v string) *CreateIngressRequest {
	s.Description = &v
	return s
}

func (s *CreateIngressRequest) SetListenerPort(v int32) *CreateIngressRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateIngressRequest) SetListenerProtocol(v string) *CreateIngressRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateIngressRequest) SetLoadBalanceType(v string) *CreateIngressRequest {
	s.LoadBalanceType = &v
	return s
}

func (s *CreateIngressRequest) SetNamespaceId(v string) *CreateIngressRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateIngressRequest) SetRules(v string) *CreateIngressRequest {
	s.Rules = &v
	return s
}

func (s *CreateIngressRequest) SetSlbId(v string) *CreateIngressRequest {
	s.SlbId = &v
	return s
}

type CreateIngressResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                        `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateIngressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIngressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIngressResponseBody) SetCode(v string) *CreateIngressResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIngressResponseBody) SetData(v *CreateIngressResponseBodyData) *CreateIngressResponseBody {
	s.Data = v
	return s
}

func (s *CreateIngressResponseBody) SetErrorCode(v string) *CreateIngressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateIngressResponseBody) SetMessage(v string) *CreateIngressResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIngressResponseBody) SetRequestId(v string) *CreateIngressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIngressResponseBody) SetSuccess(v bool) *CreateIngressResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIngressResponseBody) SetTraceId(v string) *CreateIngressResponseBody {
	s.TraceId = &v
	return s
}

type CreateIngressResponseBodyData struct {
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s CreateIngressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateIngressResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIngressResponseBodyData) SetIngressId(v int64) *CreateIngressResponseBodyData {
	s.IngressId = &v
	return s
}

type CreateIngressResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIngressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIngressResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIngressResponse) GoString() string {
	return s.String()
}

func (s *CreateIngressResponse) SetHeaders(v map[string]*string) *CreateIngressResponse {
	s.Headers = v
	return s
}

func (s *CreateIngressResponse) SetStatusCode(v int32) *CreateIngressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIngressResponse) SetBody(v *CreateIngressResponseBody) *CreateIngressResponse {
	s.Body = v
	return s
}

type CreateNamespaceRequest struct {
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	NamespaceId          *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName        *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) SetNamespaceDescription(v string) *CreateNamespaceRequest {
	s.NamespaceDescription = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceId(v string) *CreateNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceName(v string) *CreateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type CreateNamespaceResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                          `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) SetCode(v string) *CreateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateNamespaceResponseBody) SetErrorCode(v string) *CreateNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetMessage(v string) *CreateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetSuccess(v bool) *CreateNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetTraceId(v string) *CreateNamespaceResponseBody {
	s.TraceId = &v
	return s
}

type CreateNamespaceResponseBodyData struct {
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	NamespaceId          *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName        *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateNamespaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBodyData) SetNamespaceDescription(v string) *CreateNamespaceResponseBodyData {
	s.NamespaceDescription = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) SetNamespaceId(v string) *CreateNamespaceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) SetNamespaceName(v string) *CreateNamespaceResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) SetRegionId(v string) *CreateNamespaceResponseBodyData {
	s.RegionId = &v
	return s
}

type CreateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponse) SetHeaders(v map[string]*string) *CreateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateNamespaceResponse) SetStatusCode(v int32) *CreateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNamespaceResponse) SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse {
	s.Body = v
	return s
}

type DeleteApplicationRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetAppId(v string) *DeleteApplicationRequest {
	s.AppId = &v
	return s
}

type DeleteApplicationResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                            `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) SetCode(v string) *DeleteApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetData(v *DeleteApplicationResponseBodyData) *DeleteApplicationResponseBody {
	s.Data = v
	return s
}

func (s *DeleteApplicationResponseBody) SetErrorCode(v string) *DeleteApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetMessage(v string) *DeleteApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetSuccess(v bool) *DeleteApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetTraceId(v string) *DeleteApplicationResponseBody {
	s.TraceId = &v
	return s
}

type DeleteApplicationResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s DeleteApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBodyData) SetChangeOrderId(v string) *DeleteApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type DeleteApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetStatusCode(v int32) *DeleteApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

type DeleteApplicationScalingRuleRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s DeleteApplicationScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationScalingRuleRequest) SetAppId(v string) *DeleteApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *DeleteApplicationScalingRuleRequest) SetScalingRuleName(v string) *DeleteApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

type DeleteApplicationScalingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteApplicationScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationScalingRuleResponseBody) SetRequestId(v string) *DeleteApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) SetTraceId(v string) *DeleteApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

type DeleteApplicationScalingRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplicationScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *DeleteApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationScalingRuleResponse) SetStatusCode(v int32) *DeleteApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponse) SetBody(v *DeleteApplicationScalingRuleResponseBody) *DeleteApplicationScalingRuleResponse {
	s.Body = v
	return s
}

type DeleteConfigMapRequest struct {
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s DeleteConfigMapRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigMapRequest) SetConfigMapId(v int64) *DeleteConfigMapRequest {
	s.ConfigMapId = &v
	return s
}

type DeleteConfigMapResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteConfigMapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                          `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteConfigMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigMapResponseBody) SetCode(v string) *DeleteConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetData(v *DeleteConfigMapResponseBodyData) *DeleteConfigMapResponseBody {
	s.Data = v
	return s
}

func (s *DeleteConfigMapResponseBody) SetErrorCode(v string) *DeleteConfigMapResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetMessage(v string) *DeleteConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetRequestId(v string) *DeleteConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetSuccess(v bool) *DeleteConfigMapResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConfigMapResponseBody) SetTraceId(v string) *DeleteConfigMapResponseBody {
	s.TraceId = &v
	return s
}

type DeleteConfigMapResponseBodyData struct {
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s DeleteConfigMapResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigMapResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteConfigMapResponseBodyData) SetConfigMapId(v int64) *DeleteConfigMapResponseBodyData {
	s.ConfigMapId = &v
	return s
}

type DeleteConfigMapResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConfigMapResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigMapResponse) SetHeaders(v map[string]*string) *DeleteConfigMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigMapResponse) SetStatusCode(v int32) *DeleteConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigMapResponse) SetBody(v *DeleteConfigMapResponseBody) *DeleteConfigMapResponse {
	s.Body = v
	return s
}

type DeleteGreyTagRouteRequest struct {
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s DeleteGreyTagRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteGreyTagRouteRequest) SetGreyTagRouteId(v int64) *DeleteGreyTagRouteRequest {
	s.GreyTagRouteId = &v
	return s
}

type DeleteGreyTagRouteResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                             `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteGreyTagRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGreyTagRouteResponseBody) SetCode(v string) *DeleteGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetData(v *DeleteGreyTagRouteResponseBodyData) *DeleteGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetErrorCode(v string) *DeleteGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetMessage(v string) *DeleteGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetRequestId(v string) *DeleteGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetSuccess(v bool) *DeleteGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGreyTagRouteResponseBody) SetTraceId(v string) *DeleteGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

type DeleteGreyTagRouteResponseBodyData struct {
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s DeleteGreyTagRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteGreyTagRouteResponseBodyData) SetGreyTagRouteId(v int64) *DeleteGreyTagRouteResponseBodyData {
	s.GreyTagRouteId = &v
	return s
}

type DeleteGreyTagRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGreyTagRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteGreyTagRouteResponse) SetHeaders(v map[string]*string) *DeleteGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteGreyTagRouteResponse) SetStatusCode(v int32) *DeleteGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGreyTagRouteResponse) SetBody(v *DeleteGreyTagRouteResponseBody) *DeleteGreyTagRouteResponse {
	s.Body = v
	return s
}

type DeleteIngressRequest struct {
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s DeleteIngressRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIngressRequest) GoString() string {
	return s.String()
}

func (s *DeleteIngressRequest) SetIngressId(v int64) *DeleteIngressRequest {
	s.IngressId = &v
	return s
}

type DeleteIngressResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeleteIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                        `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteIngressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIngressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIngressResponseBody) SetCode(v string) *DeleteIngressResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIngressResponseBody) SetData(v *DeleteIngressResponseBodyData) *DeleteIngressResponseBody {
	s.Data = v
	return s
}

func (s *DeleteIngressResponseBody) SetErrorCode(v string) *DeleteIngressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteIngressResponseBody) SetMessage(v string) *DeleteIngressResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIngressResponseBody) SetRequestId(v string) *DeleteIngressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIngressResponseBody) SetSuccess(v bool) *DeleteIngressResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIngressResponseBody) SetTraceId(v string) *DeleteIngressResponseBody {
	s.TraceId = &v
	return s
}

type DeleteIngressResponseBodyData struct {
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s DeleteIngressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteIngressResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteIngressResponseBodyData) SetIngressId(v int64) *DeleteIngressResponseBodyData {
	s.IngressId = &v
	return s
}

type DeleteIngressResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIngressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIngressResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIngressResponse) GoString() string {
	return s.String()
}

func (s *DeleteIngressResponse) SetHeaders(v map[string]*string) *DeleteIngressResponse {
	s.Headers = v
	return s
}

func (s *DeleteIngressResponse) SetStatusCode(v int32) *DeleteIngressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIngressResponse) SetBody(v *DeleteIngressResponseBody) *DeleteIngressResponse {
	s.Body = v
	return s
}

type DeleteNamespaceRequest struct {
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) SetNamespaceId(v string) *DeleteNamespaceRequest {
	s.NamespaceId = &v
	return s
}

type DeleteNamespaceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) SetCode(v string) *DeleteNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetErrorCode(v string) *DeleteNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetMessage(v string) *DeleteNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetSuccess(v bool) *DeleteNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetTraceId(v string) *DeleteNamespaceResponseBody {
	s.TraceId = &v
	return s
}

type DeleteNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponse) SetHeaders(v map[string]*string) *DeleteNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNamespaceResponse) SetStatusCode(v int32) *DeleteNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNamespaceResponse) SetBody(v *DeleteNamespaceResponseBody) *DeleteNamespaceResponse {
	s.Body = v
	return s
}

type DeployApplicationRequest struct {
	AcrAssumeRoleArn                 *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	AcrInstanceId                    *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	AppId                            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AssociateEip                     *bool   `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	AutoEnableApplicationScalingRule *bool   `json:"AutoEnableApplicationScalingRule,omitempty" xml:"AutoEnableApplicationScalingRule,omitempty"`
	BatchWaitTime                    *int32  `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	ChangeOrderDesc                  *string `json:"ChangeOrderDesc,omitempty" xml:"ChangeOrderDesc,omitempty"`
	Command                          *string `json:"Command,omitempty" xml:"Command,omitempty"`
	CommandArgs                      *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	ConfigMapMountDesc               *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	CustomHostAlias                  *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	EdasContainerVersion             *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	EnableAhas                       *string `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	EnableGreyTagRoute               *bool   `json:"EnableGreyTagRoute,omitempty" xml:"EnableGreyTagRoute,omitempty"`
	Envs                             *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	ImageUrl                         *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	JarStartArgs                     *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	JarStartOptions                  *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	Jdk                              *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	KafkaConfigs                     *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	Liveness                         *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	MinReadyInstanceRatio            *int32  `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	MinReadyInstances                *int32  `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	MountDesc                        *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	MountHost                        *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	NasId                            *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	OssAkId                          *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	OssAkSecret                      *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	OssMountDescs                    *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	PackageUrl                       *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	PackageVersion                   *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	PhpArmsConfigLocation            *string `json:"PhpArmsConfigLocation,omitempty" xml:"PhpArmsConfigLocation,omitempty"`
	PhpConfig                        *string `json:"PhpConfig,omitempty" xml:"PhpConfig,omitempty"`
	PhpConfigLocation                *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	PostStart                        *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	PreStop                          *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	Readiness                        *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	SlsConfigs                       *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	TerminationGracePeriodSeconds    *int32  `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	Timezone                         *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	TomcatConfig                     *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	UpdateStrategy                   *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	WarStartOptions                  *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	WebContainer                     *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s DeployApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeployApplicationRequest) SetAcrAssumeRoleArn(v string) *DeployApplicationRequest {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *DeployApplicationRequest) SetAcrInstanceId(v string) *DeployApplicationRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *DeployApplicationRequest) SetAppId(v string) *DeployApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeployApplicationRequest) SetAssociateEip(v bool) *DeployApplicationRequest {
	s.AssociateEip = &v
	return s
}

func (s *DeployApplicationRequest) SetAutoEnableApplicationScalingRule(v bool) *DeployApplicationRequest {
	s.AutoEnableApplicationScalingRule = &v
	return s
}

func (s *DeployApplicationRequest) SetBatchWaitTime(v int32) *DeployApplicationRequest {
	s.BatchWaitTime = &v
	return s
}

func (s *DeployApplicationRequest) SetChangeOrderDesc(v string) *DeployApplicationRequest {
	s.ChangeOrderDesc = &v
	return s
}

func (s *DeployApplicationRequest) SetCommand(v string) *DeployApplicationRequest {
	s.Command = &v
	return s
}

func (s *DeployApplicationRequest) SetCommandArgs(v string) *DeployApplicationRequest {
	s.CommandArgs = &v
	return s
}

func (s *DeployApplicationRequest) SetConfigMapMountDesc(v string) *DeployApplicationRequest {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *DeployApplicationRequest) SetCustomHostAlias(v string) *DeployApplicationRequest {
	s.CustomHostAlias = &v
	return s
}

func (s *DeployApplicationRequest) SetEdasContainerVersion(v string) *DeployApplicationRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *DeployApplicationRequest) SetEnableAhas(v string) *DeployApplicationRequest {
	s.EnableAhas = &v
	return s
}

func (s *DeployApplicationRequest) SetEnableGreyTagRoute(v bool) *DeployApplicationRequest {
	s.EnableGreyTagRoute = &v
	return s
}

func (s *DeployApplicationRequest) SetEnvs(v string) *DeployApplicationRequest {
	s.Envs = &v
	return s
}

func (s *DeployApplicationRequest) SetImageUrl(v string) *DeployApplicationRequest {
	s.ImageUrl = &v
	return s
}

func (s *DeployApplicationRequest) SetJarStartArgs(v string) *DeployApplicationRequest {
	s.JarStartArgs = &v
	return s
}

func (s *DeployApplicationRequest) SetJarStartOptions(v string) *DeployApplicationRequest {
	s.JarStartOptions = &v
	return s
}

func (s *DeployApplicationRequest) SetJdk(v string) *DeployApplicationRequest {
	s.Jdk = &v
	return s
}

func (s *DeployApplicationRequest) SetKafkaConfigs(v string) *DeployApplicationRequest {
	s.KafkaConfigs = &v
	return s
}

func (s *DeployApplicationRequest) SetLiveness(v string) *DeployApplicationRequest {
	s.Liveness = &v
	return s
}

func (s *DeployApplicationRequest) SetMinReadyInstanceRatio(v int32) *DeployApplicationRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *DeployApplicationRequest) SetMinReadyInstances(v int32) *DeployApplicationRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *DeployApplicationRequest) SetMountDesc(v string) *DeployApplicationRequest {
	s.MountDesc = &v
	return s
}

func (s *DeployApplicationRequest) SetMountHost(v string) *DeployApplicationRequest {
	s.MountHost = &v
	return s
}

func (s *DeployApplicationRequest) SetNasId(v string) *DeployApplicationRequest {
	s.NasId = &v
	return s
}

func (s *DeployApplicationRequest) SetOssAkId(v string) *DeployApplicationRequest {
	s.OssAkId = &v
	return s
}

func (s *DeployApplicationRequest) SetOssAkSecret(v string) *DeployApplicationRequest {
	s.OssAkSecret = &v
	return s
}

func (s *DeployApplicationRequest) SetOssMountDescs(v string) *DeployApplicationRequest {
	s.OssMountDescs = &v
	return s
}

func (s *DeployApplicationRequest) SetPackageUrl(v string) *DeployApplicationRequest {
	s.PackageUrl = &v
	return s
}

func (s *DeployApplicationRequest) SetPackageVersion(v string) *DeployApplicationRequest {
	s.PackageVersion = &v
	return s
}

func (s *DeployApplicationRequest) SetPhpArmsConfigLocation(v string) *DeployApplicationRequest {
	s.PhpArmsConfigLocation = &v
	return s
}

func (s *DeployApplicationRequest) SetPhpConfig(v string) *DeployApplicationRequest {
	s.PhpConfig = &v
	return s
}

func (s *DeployApplicationRequest) SetPhpConfigLocation(v string) *DeployApplicationRequest {
	s.PhpConfigLocation = &v
	return s
}

func (s *DeployApplicationRequest) SetPostStart(v string) *DeployApplicationRequest {
	s.PostStart = &v
	return s
}

func (s *DeployApplicationRequest) SetPreStop(v string) *DeployApplicationRequest {
	s.PreStop = &v
	return s
}

func (s *DeployApplicationRequest) SetReadiness(v string) *DeployApplicationRequest {
	s.Readiness = &v
	return s
}

func (s *DeployApplicationRequest) SetSlsConfigs(v string) *DeployApplicationRequest {
	s.SlsConfigs = &v
	return s
}

func (s *DeployApplicationRequest) SetTerminationGracePeriodSeconds(v int32) *DeployApplicationRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *DeployApplicationRequest) SetTimezone(v string) *DeployApplicationRequest {
	s.Timezone = &v
	return s
}

func (s *DeployApplicationRequest) SetTomcatConfig(v string) *DeployApplicationRequest {
	s.TomcatConfig = &v
	return s
}

func (s *DeployApplicationRequest) SetUpdateStrategy(v string) *DeployApplicationRequest {
	s.UpdateStrategy = &v
	return s
}

func (s *DeployApplicationRequest) SetWarStartOptions(v string) *DeployApplicationRequest {
	s.WarStartOptions = &v
	return s
}

func (s *DeployApplicationRequest) SetWebContainer(v string) *DeployApplicationRequest {
	s.WebContainer = &v
	return s
}

type DeployApplicationResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DeployApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                            `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeployApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponseBody) SetCode(v string) *DeployApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeployApplicationResponseBody) SetData(v *DeployApplicationResponseBodyData) *DeployApplicationResponseBody {
	s.Data = v
	return s
}

func (s *DeployApplicationResponseBody) SetErrorCode(v string) *DeployApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeployApplicationResponseBody) SetMessage(v string) *DeployApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeployApplicationResponseBody) SetRequestId(v string) *DeployApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployApplicationResponseBody) SetSuccess(v bool) *DeployApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *DeployApplicationResponseBody) SetTraceId(v string) *DeployApplicationResponseBody {
	s.TraceId = &v
	return s
}

type DeployApplicationResponseBodyData struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChangeOrderId  *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	IsNeedApproval *bool   `json:"IsNeedApproval,omitempty" xml:"IsNeedApproval,omitempty"`
}

func (s DeployApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponseBodyData) SetAppId(v string) *DeployApplicationResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DeployApplicationResponseBodyData) SetChangeOrderId(v string) *DeployApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *DeployApplicationResponseBodyData) SetIsNeedApproval(v bool) *DeployApplicationResponseBodyData {
	s.IsNeedApproval = &v
	return s
}

type DeployApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeployApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponse) SetHeaders(v map[string]*string) *DeployApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeployApplicationResponse) SetStatusCode(v int32) *DeployApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployApplicationResponse) SetBody(v *DeployApplicationResponseBody) *DeployApplicationResponse {
	s.Body = v
	return s
}

type DescribeAppServiceDetailRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ServiceGroup   *string `json:"ServiceGroup,omitempty" xml:"ServiceGroup,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceType    *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s DescribeAppServiceDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppServiceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailRequest) SetAppId(v string) *DescribeAppServiceDetailRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetServiceGroup(v string) *DescribeAppServiceDetailRequest {
	s.ServiceGroup = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetServiceName(v string) *DescribeAppServiceDetailRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetServiceType(v string) *DescribeAppServiceDetailRequest {
	s.ServiceType = &v
	return s
}

func (s *DescribeAppServiceDetailRequest) SetServiceVersion(v string) *DescribeAppServiceDetailRequest {
	s.ServiceVersion = &v
	return s
}

type DescribeAppServiceDetailResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeAppServiceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                   `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeAppServiceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppServiceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponseBody) SetCode(v string) *DescribeAppServiceDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetData(v *DescribeAppServiceDetailResponseBodyData) *DescribeAppServiceDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetErrorCode(v string) *DescribeAppServiceDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetMessage(v string) *DescribeAppServiceDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetRequestId(v string) *DescribeAppServiceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetSuccess(v bool) *DescribeAppServiceDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetTraceId(v string) *DescribeAppServiceDetailResponseBody {
	s.TraceId = &v
	return s
}

type DescribeAppServiceDetailResponseBodyData struct {
	DubboApplicationName  *string                                            `json:"DubboApplicationName,omitempty" xml:"DubboApplicationName,omitempty"`
	EdasAppName           *string                                            `json:"EdasAppName,omitempty" xml:"EdasAppName,omitempty"`
	Group                 *string                                            `json:"Group,omitempty" xml:"Group,omitempty"`
	Metadata              map[string]interface{}                             `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Methods               []*DescribeAppServiceDetailResponseBodyDataMethods `json:"Methods,omitempty" xml:"Methods,omitempty" type:"Repeated"`
	ServiceName           *string                                            `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceType           *string                                            `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SpringApplicationName *string                                            `json:"SpringApplicationName,omitempty" xml:"SpringApplicationName,omitempty"`
	Version               *string                                            `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAppServiceDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppServiceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponseBodyData) SetDubboApplicationName(v string) *DescribeAppServiceDetailResponseBodyData {
	s.DubboApplicationName = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetEdasAppName(v string) *DescribeAppServiceDetailResponseBodyData {
	s.EdasAppName = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetGroup(v string) *DescribeAppServiceDetailResponseBodyData {
	s.Group = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetMetadata(v map[string]interface{}) *DescribeAppServiceDetailResponseBodyData {
	s.Metadata = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetMethods(v []*DescribeAppServiceDetailResponseBodyDataMethods) *DescribeAppServiceDetailResponseBodyData {
	s.Methods = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetServiceName(v string) *DescribeAppServiceDetailResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetServiceType(v string) *DescribeAppServiceDetailResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetSpringApplicationName(v string) *DescribeAppServiceDetailResponseBodyData {
	s.SpringApplicationName = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetVersion(v string) *DescribeAppServiceDetailResponseBodyData {
	s.Version = &v
	return s
}

type DescribeAppServiceDetailResponseBodyDataMethods struct {
	MethodController     *string                                                                `json:"MethodController,omitempty" xml:"MethodController,omitempty"`
	Name                 *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	NameDetail           *string                                                                `json:"NameDetail,omitempty" xml:"NameDetail,omitempty"`
	ParameterDefinitions []*DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions `json:"ParameterDefinitions,omitempty" xml:"ParameterDefinitions,omitempty" type:"Repeated"`
	ParameterDetails     []*string                                                              `json:"ParameterDetails,omitempty" xml:"ParameterDetails,omitempty" type:"Repeated"`
	ParameterTypes       []*string                                                              `json:"ParameterTypes,omitempty" xml:"ParameterTypes,omitempty" type:"Repeated"`
	Paths                []*string                                                              `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	RequestMethods       []*string                                                              `json:"RequestMethods,omitempty" xml:"RequestMethods,omitempty" type:"Repeated"`
	ReturnDetails        *string                                                                `json:"ReturnDetails,omitempty" xml:"ReturnDetails,omitempty"`
	ReturnType           *string                                                                `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
}

func (s DescribeAppServiceDetailResponseBodyDataMethods) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppServiceDetailResponseBodyDataMethods) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetMethodController(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.MethodController = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetName(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.Name = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetNameDetail(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.NameDetail = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetParameterDefinitions(v []*DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ParameterDefinitions = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetParameterDetails(v []*string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ParameterDetails = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetParameterTypes(v []*string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ParameterTypes = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetPaths(v []*string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.Paths = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetRequestMethods(v []*string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.RequestMethods = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetReturnDetails(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ReturnDetails = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetReturnType(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ReturnType = &v
	return s
}

type DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) SetDescription(v string) *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions {
	s.Description = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) SetName(v string) *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions {
	s.Name = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) SetType(v string) *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions {
	s.Type = &v
	return s
}

type DescribeAppServiceDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppServiceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppServiceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppServiceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponse) SetHeaders(v map[string]*string) *DescribeAppServiceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppServiceDetailResponse) SetStatusCode(v int32) *DescribeAppServiceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppServiceDetailResponse) SetBody(v *DescribeAppServiceDetailResponseBody) *DescribeAppServiceDetailResponse {
	s.Body = v
	return s
}

type DescribeApplicationConfigRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DescribeApplicationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigRequest) SetAppId(v string) *DescribeApplicationConfigRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationConfigRequest) SetVersionId(v string) *DescribeApplicationConfigRequest {
	s.VersionId = &v
	return s
}

type DescribeApplicationConfigResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeApplicationConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                    `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBody) SetCode(v string) *DescribeApplicationConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetData(v *DescribeApplicationConfigResponseBodyData) *DescribeApplicationConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetErrorCode(v string) *DescribeApplicationConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetMessage(v string) *DescribeApplicationConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetRequestId(v string) *DescribeApplicationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetSuccess(v bool) *DescribeApplicationConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetTraceId(v string) *DescribeApplicationConfigResponseBody {
	s.TraceId = &v
	return s
}

type DescribeApplicationConfigResponseBodyData struct {
	AcrAssumeRoleArn              *string                                                        `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	AcrInstanceId                 *string                                                        `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	AppDescription                *string                                                        `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	AppId                         *string                                                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName                       *string                                                        `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AssociateEip                  *bool                                                          `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	BatchWaitTime                 *int32                                                         `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	Command                       *string                                                        `json:"Command,omitempty" xml:"Command,omitempty"`
	CommandArgs                   *string                                                        `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	ConfigMapMountDesc            []*DescribeApplicationConfigResponseBodyDataConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	Cpu                           *int32                                                         `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CustomHostAlias               *string                                                        `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	EdasContainerVersion          *string                                                        `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	EnableAhas                    *string                                                        `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	EnableGreyTagRoute            *bool                                                          `json:"EnableGreyTagRoute,omitempty" xml:"EnableGreyTagRoute,omitempty"`
	Envs                          *string                                                        `json:"Envs,omitempty" xml:"Envs,omitempty"`
	ImageUrl                      *string                                                        `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	JarStartArgs                  *string                                                        `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	JarStartOptions               *string                                                        `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	Jdk                           *string                                                        `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	KafkaConfigs                  *string                                                        `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	Liveness                      *string                                                        `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	Memory                        *int32                                                         `json:"Memory,omitempty" xml:"Memory,omitempty"`
	MinReadyInstanceRatio         *int32                                                         `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	MinReadyInstances             *int32                                                         `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	MountDesc                     []*DescribeApplicationConfigResponseBodyDataMountDesc          `json:"MountDesc,omitempty" xml:"MountDesc,omitempty" type:"Repeated"`
	MountHost                     *string                                                        `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	MseApplicationId              *string                                                        `json:"MseApplicationId,omitempty" xml:"MseApplicationId,omitempty"`
	MseFeatureConfig              *string                                                        `json:"MseFeatureConfig,omitempty" xml:"MseFeatureConfig,omitempty"`
	NamespaceId                   *string                                                        `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NasId                         *string                                                        `json:"NasId,omitempty" xml:"NasId,omitempty"`
	OssAkId                       *string                                                        `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	OssAkSecret                   *string                                                        `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	OssMountDescs                 []*DescribeApplicationConfigResponseBodyDataOssMountDescs      `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty" type:"Repeated"`
	PackageType                   *string                                                        `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PackageUrl                    *string                                                        `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	PackageVersion                *string                                                        `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	PhpArmsConfigLocation         *string                                                        `json:"PhpArmsConfigLocation,omitempty" xml:"PhpArmsConfigLocation,omitempty"`
	PhpConfig                     *string                                                        `json:"PhpConfig,omitempty" xml:"PhpConfig,omitempty"`
	PhpConfigLocation             *string                                                        `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	PostStart                     *string                                                        `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	PreStop                       *string                                                        `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	ProgrammingLanguage           *string                                                        `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	Readiness                     *string                                                        `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	RegionId                      *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Replicas                      *int32                                                         `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	SecurityGroupId               *string                                                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SlsConfigs                    *string                                                        `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	Tags                          []*DescribeApplicationConfigResponseBodyDataTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TerminationGracePeriodSeconds *int32                                                         `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	Timezone                      *string                                                        `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	TomcatConfig                  *string                                                        `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	UpdateStrategy                *string                                                        `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	VSwitchId                     *string                                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                         *string                                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WarStartOptions               *string                                                        `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	WebContainer                  *string                                                        `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyData) SetAcrAssumeRoleArn(v string) *DescribeApplicationConfigResponseBodyData {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAcrInstanceId(v string) *DescribeApplicationConfigResponseBodyData {
	s.AcrInstanceId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAppDescription(v string) *DescribeApplicationConfigResponseBodyData {
	s.AppDescription = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAppId(v string) *DescribeApplicationConfigResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAppName(v string) *DescribeApplicationConfigResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAssociateEip(v bool) *DescribeApplicationConfigResponseBodyData {
	s.AssociateEip = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetBatchWaitTime(v int32) *DescribeApplicationConfigResponseBodyData {
	s.BatchWaitTime = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCommand(v string) *DescribeApplicationConfigResponseBodyData {
	s.Command = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCommandArgs(v string) *DescribeApplicationConfigResponseBodyData {
	s.CommandArgs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetConfigMapMountDesc(v []*DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) *DescribeApplicationConfigResponseBodyData {
	s.ConfigMapMountDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCpu(v int32) *DescribeApplicationConfigResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCustomHostAlias(v string) *DescribeApplicationConfigResponseBodyData {
	s.CustomHostAlias = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEdasContainerVersion(v string) *DescribeApplicationConfigResponseBodyData {
	s.EdasContainerVersion = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnableAhas(v string) *DescribeApplicationConfigResponseBodyData {
	s.EnableAhas = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnableGreyTagRoute(v bool) *DescribeApplicationConfigResponseBodyData {
	s.EnableGreyTagRoute = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnvs(v string) *DescribeApplicationConfigResponseBodyData {
	s.Envs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetImageUrl(v string) *DescribeApplicationConfigResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetJarStartArgs(v string) *DescribeApplicationConfigResponseBodyData {
	s.JarStartArgs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetJarStartOptions(v string) *DescribeApplicationConfigResponseBodyData {
	s.JarStartOptions = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetJdk(v string) *DescribeApplicationConfigResponseBodyData {
	s.Jdk = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetKafkaConfigs(v string) *DescribeApplicationConfigResponseBodyData {
	s.KafkaConfigs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetLiveness(v string) *DescribeApplicationConfigResponseBodyData {
	s.Liveness = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMemory(v int32) *DescribeApplicationConfigResponseBodyData {
	s.Memory = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMinReadyInstanceRatio(v int32) *DescribeApplicationConfigResponseBodyData {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMinReadyInstances(v int32) *DescribeApplicationConfigResponseBodyData {
	s.MinReadyInstances = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMountDesc(v []*DescribeApplicationConfigResponseBodyDataMountDesc) *DescribeApplicationConfigResponseBodyData {
	s.MountDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMountHost(v string) *DescribeApplicationConfigResponseBodyData {
	s.MountHost = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMseApplicationId(v string) *DescribeApplicationConfigResponseBodyData {
	s.MseApplicationId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMseFeatureConfig(v string) *DescribeApplicationConfigResponseBodyData {
	s.MseFeatureConfig = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetNamespaceId(v string) *DescribeApplicationConfigResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetNasId(v string) *DescribeApplicationConfigResponseBodyData {
	s.NasId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetOssAkId(v string) *DescribeApplicationConfigResponseBodyData {
	s.OssAkId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetOssAkSecret(v string) *DescribeApplicationConfigResponseBodyData {
	s.OssAkSecret = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetOssMountDescs(v []*DescribeApplicationConfigResponseBodyDataOssMountDescs) *DescribeApplicationConfigResponseBodyData {
	s.OssMountDescs = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPackageType(v string) *DescribeApplicationConfigResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPackageUrl(v string) *DescribeApplicationConfigResponseBodyData {
	s.PackageUrl = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPackageVersion(v string) *DescribeApplicationConfigResponseBodyData {
	s.PackageVersion = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPhpArmsConfigLocation(v string) *DescribeApplicationConfigResponseBodyData {
	s.PhpArmsConfigLocation = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPhpConfig(v string) *DescribeApplicationConfigResponseBodyData {
	s.PhpConfig = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPhpConfigLocation(v string) *DescribeApplicationConfigResponseBodyData {
	s.PhpConfigLocation = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPostStart(v string) *DescribeApplicationConfigResponseBodyData {
	s.PostStart = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPreStop(v string) *DescribeApplicationConfigResponseBodyData {
	s.PreStop = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetProgrammingLanguage(v string) *DescribeApplicationConfigResponseBodyData {
	s.ProgrammingLanguage = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetReadiness(v string) *DescribeApplicationConfigResponseBodyData {
	s.Readiness = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetRegionId(v string) *DescribeApplicationConfigResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetReplicas(v int32) *DescribeApplicationConfigResponseBodyData {
	s.Replicas = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetSecurityGroupId(v string) *DescribeApplicationConfigResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetSlsConfigs(v string) *DescribeApplicationConfigResponseBodyData {
	s.SlsConfigs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetTags(v []*DescribeApplicationConfigResponseBodyDataTags) *DescribeApplicationConfigResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetTerminationGracePeriodSeconds(v int32) *DescribeApplicationConfigResponseBodyData {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetTimezone(v string) *DescribeApplicationConfigResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetTomcatConfig(v string) *DescribeApplicationConfigResponseBodyData {
	s.TomcatConfig = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetUpdateStrategy(v string) *DescribeApplicationConfigResponseBodyData {
	s.UpdateStrategy = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetVSwitchId(v string) *DescribeApplicationConfigResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetVpcId(v string) *DescribeApplicationConfigResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetWarStartOptions(v string) *DescribeApplicationConfigResponseBodyData {
	s.WarStartOptions = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetWebContainer(v string) *DescribeApplicationConfigResponseBodyData {
	s.WebContainer = &v
	return s
}

type DescribeApplicationConfigResponseBodyDataConfigMapMountDesc struct {
	ConfigMapId   *int64  `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	ConfigMapName *string `json:"ConfigMapName,omitempty" xml:"ConfigMapName,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	MountPath     *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) SetConfigMapId(v int64) *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc {
	s.ConfigMapId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) SetConfigMapName(v string) *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc {
	s.ConfigMapName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) SetKey(v string) *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc {
	s.Key = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc {
	s.MountPath = &v
	return s
}

type DescribeApplicationConfigResponseBodyDataMountDesc struct {
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	NasPath   *string `json:"NasPath,omitempty" xml:"NasPath,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataMountDesc) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataMountDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataMountDesc) SetNasPath(v string) *DescribeApplicationConfigResponseBodyDataMountDesc {
	s.NasPath = &v
	return s
}

type DescribeApplicationConfigResponseBodyDataOssMountDescs struct {
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	MountPath  *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataOssMountDescs) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataOssMountDescs) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) SetBucketName(v string) *DescribeApplicationConfigResponseBodyDataOssMountDescs {
	s.BucketName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) SetBucketPath(v string) *DescribeApplicationConfigResponseBodyDataOssMountDescs {
	s.BucketPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataOssMountDescs {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) SetReadOnly(v bool) *DescribeApplicationConfigResponseBodyDataOssMountDescs {
	s.ReadOnly = &v
	return s
}

type DescribeApplicationConfigResponseBodyDataTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataTags) SetKey(v string) *DescribeApplicationConfigResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataTags) SetValue(v string) *DescribeApplicationConfigResponseBodyDataTags {
	s.Value = &v
	return s
}

type DescribeApplicationConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApplicationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponse) SetHeaders(v map[string]*string) *DescribeApplicationConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationConfigResponse) SetStatusCode(v int32) *DescribeApplicationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationConfigResponse) SetBody(v *DescribeApplicationConfigResponseBody) *DescribeApplicationConfigResponse {
	s.Body = v
	return s
}

type DescribeApplicationGroupsRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeApplicationGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupsRequest) SetAppId(v string) *DescribeApplicationGroupsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationGroupsRequest) SetCurrentPage(v int32) *DescribeApplicationGroupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationGroupsRequest) SetPageSize(v int32) *DescribeApplicationGroupsRequest {
	s.PageSize = &v
	return s
}

type DescribeApplicationGroupsResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeApplicationGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                      `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupsResponseBody) SetCode(v string) *DescribeApplicationGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetData(v []*DescribeApplicationGroupsResponseBodyData) *DescribeApplicationGroupsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetErrorCode(v string) *DescribeApplicationGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetMessage(v string) *DescribeApplicationGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetRequestId(v string) *DescribeApplicationGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetSuccess(v bool) *DescribeApplicationGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetTraceId(v string) *DescribeApplicationGroupsResponseBody {
	s.TraceId = &v
	return s
}

type DescribeApplicationGroupsResponseBodyData struct {
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType            *int32  `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	ImageUrl             *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Jdk                  *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	PackageType          *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PackageUrl           *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	PackageVersion       *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	Replicas             *int32  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	RunningInstances     *int32  `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	WebContainer         *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s DescribeApplicationGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupsResponseBodyData) SetEdasContainerVersion(v string) *DescribeApplicationGroupsResponseBodyData {
	s.EdasContainerVersion = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetGroupId(v string) *DescribeApplicationGroupsResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetGroupName(v string) *DescribeApplicationGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetGroupType(v int32) *DescribeApplicationGroupsResponseBodyData {
	s.GroupType = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetImageUrl(v string) *DescribeApplicationGroupsResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetJdk(v string) *DescribeApplicationGroupsResponseBodyData {
	s.Jdk = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetPackageType(v string) *DescribeApplicationGroupsResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetPackageUrl(v string) *DescribeApplicationGroupsResponseBodyData {
	s.PackageUrl = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetPackageVersion(v string) *DescribeApplicationGroupsResponseBodyData {
	s.PackageVersion = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetReplicas(v int32) *DescribeApplicationGroupsResponseBodyData {
	s.Replicas = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetRunningInstances(v int32) *DescribeApplicationGroupsResponseBodyData {
	s.RunningInstances = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetWebContainer(v string) *DescribeApplicationGroupsResponseBodyData {
	s.WebContainer = &v
	return s
}

type DescribeApplicationGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApplicationGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupsResponse) SetHeaders(v map[string]*string) *DescribeApplicationGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationGroupsResponse) SetStatusCode(v int32) *DescribeApplicationGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationGroupsResponse) SetBody(v *DescribeApplicationGroupsResponseBody) *DescribeApplicationGroupsResponse {
	s.Body = v
	return s
}

type DescribeApplicationImageRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s DescribeApplicationImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationImageRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationImageRequest) SetAppId(v string) *DescribeApplicationImageRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationImageRequest) SetImageUrl(v string) *DescribeApplicationImageRequest {
	s.ImageUrl = &v
	return s
}

type DescribeApplicationImageResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeApplicationImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                   `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationImageResponseBody) SetCode(v string) *DescribeApplicationImageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetData(v *DescribeApplicationImageResponseBodyData) *DescribeApplicationImageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetErrorCode(v string) *DescribeApplicationImageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetMessage(v string) *DescribeApplicationImageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetRequestId(v string) *DescribeApplicationImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetSuccess(v bool) *DescribeApplicationImageResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationImageResponseBody) SetTraceId(v string) *DescribeApplicationImageResponseBody {
	s.TraceId = &v
	return s
}

type DescribeApplicationImageResponseBodyData struct {
	CrUrl          *string `json:"CrUrl,omitempty" xml:"CrUrl,omitempty"`
	Logo           *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepoName       *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespace  *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	RepoOriginType *string `json:"RepoOriginType,omitempty" xml:"RepoOriginType,omitempty"`
	RepoTag        *string `json:"RepoTag,omitempty" xml:"RepoTag,omitempty"`
	RepoType       *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
}

func (s DescribeApplicationImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationImageResponseBodyData) SetCrUrl(v string) *DescribeApplicationImageResponseBodyData {
	s.CrUrl = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetLogo(v string) *DescribeApplicationImageResponseBodyData {
	s.Logo = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRegionId(v string) *DescribeApplicationImageResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoName(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoName = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoNamespace(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoOriginType(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoOriginType = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoTag(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoTag = &v
	return s
}

func (s *DescribeApplicationImageResponseBodyData) SetRepoType(v string) *DescribeApplicationImageResponseBodyData {
	s.RepoType = &v
	return s
}

type DescribeApplicationImageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApplicationImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationImageResponse) SetHeaders(v map[string]*string) *DescribeApplicationImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationImageResponse) SetStatusCode(v int32) *DescribeApplicationImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationImageResponse) SetBody(v *DescribeApplicationImageResponseBody) *DescribeApplicationImageResponse {
	s.Body = v
	return s
}

type DescribeApplicationInstancesRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse     *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s DescribeApplicationInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesRequest) SetAppId(v string) *DescribeApplicationInstancesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetCurrentPage(v int32) *DescribeApplicationInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetGroupId(v string) *DescribeApplicationInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetPageSize(v int32) *DescribeApplicationInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationInstancesRequest) SetReverse(v bool) *DescribeApplicationInstancesRequest {
	s.Reverse = &v
	return s
}

type DescribeApplicationInstancesResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeApplicationInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                       `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesResponseBody) SetCode(v string) *DescribeApplicationInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetData(v *DescribeApplicationInstancesResponseBodyData) *DescribeApplicationInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetErrorCode(v string) *DescribeApplicationInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetMessage(v string) *DescribeApplicationInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetRequestId(v string) *DescribeApplicationInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetSuccess(v bool) *DescribeApplicationInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBody) SetTraceId(v string) *DescribeApplicationInstancesResponseBody {
	s.TraceId = &v
	return s
}

type DescribeApplicationInstancesResponseBodyData struct {
	CurrentPage *int32                                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Instances   []*DescribeApplicationInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	PageSize    *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize   *int32                                                   `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeApplicationInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesResponseBodyData) SetCurrentPage(v int32) *DescribeApplicationInstancesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyData) SetInstances(v []*DescribeApplicationInstancesResponseBodyDataInstances) *DescribeApplicationInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyData) SetPageSize(v int32) *DescribeApplicationInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyData) SetTotalSize(v int32) *DescribeApplicationInstancesResponseBodyData {
	s.TotalSize = &v
	return s
}

type DescribeApplicationInstancesResponseBodyDataInstances struct {
	CreateTimeStamp           *int64  `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	DebugStatus               *bool   `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	Eip                       *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	FinishTimeStamp           *int64  `json:"FinishTimeStamp,omitempty" xml:"FinishTimeStamp,omitempty"`
	GroupId                   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ImageUrl                  *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	InstanceContainerIp       *string `json:"InstanceContainerIp,omitempty" xml:"InstanceContainerIp,omitempty"`
	InstanceContainerRestarts *int64  `json:"InstanceContainerRestarts,omitempty" xml:"InstanceContainerRestarts,omitempty"`
	InstanceContainerStatus   *string `json:"InstanceContainerStatus,omitempty" xml:"InstanceContainerStatus,omitempty"`
	InstanceHealthStatus      *string `json:"InstanceHealthStatus,omitempty" xml:"InstanceHealthStatus,omitempty"`
	InstanceId                *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PackageVersion            *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	VSwitchId                 *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeApplicationInstancesResponseBodyDataInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetCreateTimeStamp(v int64) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetDebugStatus(v bool) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.DebugStatus = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetEip(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.Eip = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetFinishTimeStamp(v int64) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.FinishTimeStamp = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetGroupId(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.GroupId = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetImageUrl(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceContainerIp(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceContainerIp = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceContainerRestarts(v int64) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceContainerRestarts = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceContainerStatus(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceContainerStatus = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceHealthStatus(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceHealthStatus = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetInstanceId(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetPackageVersion(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.PackageVersion = &v
	return s
}

func (s *DescribeApplicationInstancesResponseBodyDataInstances) SetVSwitchId(v string) *DescribeApplicationInstancesResponseBodyDataInstances {
	s.VSwitchId = &v
	return s
}

type DescribeApplicationInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApplicationInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationInstancesResponse) SetHeaders(v map[string]*string) *DescribeApplicationInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationInstancesResponse) SetStatusCode(v int32) *DescribeApplicationInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationInstancesResponse) SetBody(v *DescribeApplicationInstancesResponseBody) *DescribeApplicationInstancesResponse {
	s.Body = v
	return s
}

type DescribeApplicationScalingRuleRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s DescribeApplicationScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleRequest) SetAppId(v string) *DescribeApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationScalingRuleRequest) SetScalingRuleName(v string) *DescribeApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

type DescribeApplicationScalingRuleResponseBody struct {
	Data      *DescribeApplicationScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string                                         `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBody) SetData(v *DescribeApplicationScalingRuleResponseBodyData) *DescribeApplicationScalingRuleResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBody) SetRequestId(v string) *DescribeApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBody) SetTraceId(v string) *DescribeApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

type DescribeApplicationScalingRuleResponseBodyData struct {
	AppId            *string                                               `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime       *int64                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastDisableTime  *int64                                                `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	Metric           *DescribeApplicationScalingRuleResponseBodyDataMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	ScaleRuleEnabled *bool                                                 `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	ScaleRuleName    *string                                               `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	ScaleRuleType    *string                                               `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	Timer            *DescribeApplicationScalingRuleResponseBodyDataTimer  `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	UpdateTime       *int64                                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetAppId(v string) *DescribeApplicationScalingRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetCreateTime(v int64) *DescribeApplicationScalingRuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetLastDisableTime(v int64) *DescribeApplicationScalingRuleResponseBodyData {
	s.LastDisableTime = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetMetric(v *DescribeApplicationScalingRuleResponseBodyDataMetric) *DescribeApplicationScalingRuleResponseBodyData {
	s.Metric = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetScaleRuleEnabled(v bool) *DescribeApplicationScalingRuleResponseBodyData {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetScaleRuleName(v string) *DescribeApplicationScalingRuleResponseBodyData {
	s.ScaleRuleName = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetScaleRuleType(v string) *DescribeApplicationScalingRuleResponseBodyData {
	s.ScaleRuleType = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetTimer(v *DescribeApplicationScalingRuleResponseBodyDataTimer) *DescribeApplicationScalingRuleResponseBodyData {
	s.Timer = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetUpdateTime(v int64) *DescribeApplicationScalingRuleResponseBodyData {
	s.UpdateTime = &v
	return s
}

type DescribeApplicationScalingRuleResponseBodyDataMetric struct {
	MaxReplicas    *int32                                                              `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	Metrics        []*DescribeApplicationScalingRuleResponseBodyDataMetricMetrics      `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	MetricsStatus  *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus  `json:"MetricsStatus,omitempty" xml:"MetricsStatus,omitempty" type:"Struct"`
	MinReplicas    *int32                                                              `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	ScaleDownRules *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules `json:"ScaleDownRules,omitempty" xml:"ScaleDownRules,omitempty" type:"Struct"`
	ScaleUpRules   *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules   `json:"ScaleUpRules,omitempty" xml:"ScaleUpRules,omitempty" type:"Struct"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetric) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetric) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetMaxReplicas(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetMetrics(v []*DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.Metrics = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetMetricsStatus(v *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.MetricsStatus = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetMinReplicas(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetScaleDownRules(v *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.ScaleDownRules = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetScaleUpRules(v *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.ScaleUpRules = v
	return s
}

type DescribeApplicationScalingRuleResponseBodyDataMetricMetrics struct {
	MetricTargetAverageUtilization *int32  `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	MetricType                     *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricTargetAverageUtilization(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricType(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricType = &v
	return s
}

type DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus struct {
	CurrentMetrics      []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics   `json:"CurrentMetrics,omitempty" xml:"CurrentMetrics,omitempty" type:"Repeated"`
	CurrentReplicas     *int64                                                                               `json:"CurrentReplicas,omitempty" xml:"CurrentReplicas,omitempty"`
	DesiredReplicas     *int64                                                                               `json:"DesiredReplicas,omitempty" xml:"DesiredReplicas,omitempty"`
	LastScaleTime       *string                                                                              `json:"LastScaleTime,omitempty" xml:"LastScaleTime,omitempty"`
	NextScaleMetrics    []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics `json:"NextScaleMetrics,omitempty" xml:"NextScaleMetrics,omitempty" type:"Repeated"`
	NextScaleTimePeriod *int32                                                                               `json:"NextScaleTimePeriod,omitempty" xml:"NextScaleTimePeriod,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetCurrentMetrics(v []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.CurrentMetrics = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetCurrentReplicas(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.CurrentReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetDesiredReplicas(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.DesiredReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetLastScaleTime(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.LastScaleTime = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetNextScaleMetrics(v []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.NextScaleMetrics = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetNextScaleTimePeriod(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.NextScaleTimePeriod = &v
	return s
}

type DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics struct {
	CurrentValue *int64  `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) SetCurrentValue(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics {
	s.CurrentValue = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) SetName(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics {
	s.Name = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) SetType(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics {
	s.Type = &v
	return s
}

type DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics struct {
	Name                           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextScaleInAverageUtilization  *int32  `json:"NextScaleInAverageUtilization,omitempty" xml:"NextScaleInAverageUtilization,omitempty"`
	NextScaleOutAverageUtilization *int32  `json:"NextScaleOutAverageUtilization,omitempty" xml:"NextScaleOutAverageUtilization,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) SetName(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics {
	s.Name = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) SetNextScaleInAverageUtilization(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics {
	s.NextScaleInAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) SetNextScaleOutAverageUtilization(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics {
	s.NextScaleOutAverageUtilization = &v
	return s
}

type DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules struct {
	Disabled                   *bool  `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	Step                       *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) SetDisabled(v bool) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules {
	s.Disabled = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) SetStabilizationWindowSeconds(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) SetStep(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules {
	s.Step = &v
	return s
}

type DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules struct {
	Disabled                   *bool  `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	Step                       *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) SetDisabled(v bool) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules {
	s.Disabled = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) SetStabilizationWindowSeconds(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) SetStep(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules {
	s.Step = &v
	return s
}

type DescribeApplicationScalingRuleResponseBodyDataTimer struct {
	BeginDate *string                                                         `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate   *string                                                         `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Period    *string                                                         `json:"Period,omitempty" xml:"Period,omitempty"`
	Schedules []*DescribeApplicationScalingRuleResponseBodyDataTimerSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataTimer) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataTimer) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) SetBeginDate(v string) *DescribeApplicationScalingRuleResponseBodyDataTimer {
	s.BeginDate = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) SetEndDate(v string) *DescribeApplicationScalingRuleResponseBodyDataTimer {
	s.EndDate = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) SetPeriod(v string) *DescribeApplicationScalingRuleResponseBodyDataTimer {
	s.Period = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) SetSchedules(v []*DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) *DescribeApplicationScalingRuleResponseBodyDataTimer {
	s.Schedules = v
	return s
}

type DescribeApplicationScalingRuleResponseBodyDataTimerSchedules struct {
	AtTime         *string `json:"AtTime,omitempty" xml:"AtTime,omitempty"`
	TargetReplicas *int32  `json:"TargetReplicas,omitempty" xml:"TargetReplicas,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) SetAtTime(v string) *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.AtTime = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) SetTargetReplicas(v int32) *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.TargetReplicas = &v
	return s
}

type DescribeApplicationScalingRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *DescribeApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationScalingRuleResponse) SetStatusCode(v int32) *DescribeApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponse) SetBody(v *DescribeApplicationScalingRuleResponseBody) *DescribeApplicationScalingRuleResponse {
	s.Body = v
	return s
}

type DescribeApplicationScalingRulesRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeApplicationScalingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesRequest) SetAppId(v string) *DescribeApplicationScalingRulesRequest {
	s.AppId = &v
	return s
}

type DescribeApplicationScalingRulesResponseBody struct {
	Data      *DescribeApplicationScalingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string                                          `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBody) SetData(v *DescribeApplicationScalingRulesResponseBodyData) *DescribeApplicationScalingRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetRequestId(v string) *DescribeApplicationScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetTraceId(v string) *DescribeApplicationScalingRulesResponseBody {
	s.TraceId = &v
	return s
}

type DescribeApplicationScalingRulesResponseBodyData struct {
	ApplicationScalingRules []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules `json:"ApplicationScalingRules,omitempty" xml:"ApplicationScalingRules,omitempty" type:"Repeated"`
	CurrentPage             *int32                                                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize                *int32                                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize               *int32                                                                    `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyData) SetApplicationScalingRules(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) *DescribeApplicationScalingRulesResponseBodyData {
	s.ApplicationScalingRules = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyData) SetCurrentPage(v int32) *DescribeApplicationScalingRulesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyData) SetPageSize(v int32) *DescribeApplicationScalingRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyData) SetTotalSize(v int32) *DescribeApplicationScalingRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules struct {
	AppId            *string                                                                       `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime       *int64                                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastDisableTime  *int64                                                                        `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	Metric           *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	ScaleRuleEnabled *bool                                                                         `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	ScaleRuleName    *string                                                                       `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	ScaleRuleType    *string                                                                       `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	Timer            *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer  `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	UpdateTime       *int64                                                                        `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetAppId(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetCreateTime(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetLastDisableTime(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.LastDisableTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetMetric(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.Metric = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetScaleRuleEnabled(v bool) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetScaleRuleName(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.ScaleRuleName = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetScaleRuleType(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.ScaleRuleType = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetTimer(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.Timer = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetUpdateTime(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.UpdateTime = &v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric struct {
	MaxReplicas    *int32                                                                                      `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	Metrics        []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics      `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	MetricsStatus  *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus  `json:"MetricsStatus,omitempty" xml:"MetricsStatus,omitempty" type:"Struct"`
	MinReplicas    *int32                                                                                      `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	ScaleDownRules *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules `json:"ScaleDownRules,omitempty" xml:"ScaleDownRules,omitempty" type:"Struct"`
	ScaleUpRules   *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules   `json:"ScaleUpRules,omitempty" xml:"ScaleUpRules,omitempty" type:"Struct"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetMaxReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetMetrics(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.Metrics = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetMetricsStatus(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.MetricsStatus = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetMinReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetScaleDownRules(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.ScaleDownRules = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetScaleUpRules(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.ScaleUpRules = v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics struct {
	MetricTargetAverageUtilization *int32  `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	MetricType                     *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) SetMetricTargetAverageUtilization(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) SetMetricType(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics {
	s.MetricType = &v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus struct {
	CurrentMetrics      []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics   `json:"CurrentMetrics,omitempty" xml:"CurrentMetrics,omitempty" type:"Repeated"`
	CurrentReplicas     *int64                                                                                                       `json:"CurrentReplicas,omitempty" xml:"CurrentReplicas,omitempty"`
	DesiredReplicas     *int64                                                                                                       `json:"DesiredReplicas,omitempty" xml:"DesiredReplicas,omitempty"`
	LastScaleTime       *string                                                                                                      `json:"LastScaleTime,omitempty" xml:"LastScaleTime,omitempty"`
	MaxReplicas         *int64                                                                                                       `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	MinReplicas         *int64                                                                                                       `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	NextScaleMetrics    []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics `json:"NextScaleMetrics,omitempty" xml:"NextScaleMetrics,omitempty" type:"Repeated"`
	NextScaleTimePeriod *int32                                                                                                       `json:"NextScaleTimePeriod,omitempty" xml:"NextScaleTimePeriod,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetCurrentMetrics(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.CurrentMetrics = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetCurrentReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.CurrentReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetDesiredReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.DesiredReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetLastScaleTime(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.LastScaleTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetMaxReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetMinReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetNextScaleMetrics(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.NextScaleMetrics = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetNextScaleTimePeriod(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.NextScaleTimePeriod = &v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics struct {
	CurrentValue *int64  `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) SetCurrentValue(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics {
	s.CurrentValue = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) SetName(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics {
	s.Name = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) SetType(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics {
	s.Type = &v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics struct {
	Name                           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextScaleInAverageUtilization  *int32  `json:"NextScaleInAverageUtilization,omitempty" xml:"NextScaleInAverageUtilization,omitempty"`
	NextScaleOutAverageUtilization *int32  `json:"NextScaleOutAverageUtilization,omitempty" xml:"NextScaleOutAverageUtilization,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) SetName(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics {
	s.Name = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) SetNextScaleInAverageUtilization(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics {
	s.NextScaleInAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) SetNextScaleOutAverageUtilization(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics {
	s.NextScaleOutAverageUtilization = &v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules struct {
	Disabled                   *bool  `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	Step                       *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) SetDisabled(v bool) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules {
	s.Disabled = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) SetStabilizationWindowSeconds(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) SetStep(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules {
	s.Step = &v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules struct {
	Disabled                   *bool  `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	Step                       *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) SetDisabled(v bool) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules {
	s.Disabled = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) SetStabilizationWindowSeconds(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) SetStep(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules {
	s.Step = &v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer struct {
	BeginDate *string                                                                                 `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate   *string                                                                                 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Period    *string                                                                                 `json:"Period,omitempty" xml:"Period,omitempty"`
	Schedules []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) SetBeginDate(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	s.BeginDate = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) SetEndDate(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	s.EndDate = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) SetPeriod(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	s.Period = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) SetSchedules(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	s.Schedules = v
	return s
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules struct {
	AtTime         *string `json:"AtTime,omitempty" xml:"AtTime,omitempty"`
	MaxReplicas    *int64  `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	MinReplicas    *int64  `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	TargetReplicas *int32  `json:"TargetReplicas,omitempty" xml:"TargetReplicas,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) SetAtTime(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules {
	s.AtTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) SetMaxReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) SetMinReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) SetTargetReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules {
	s.TargetReplicas = &v
	return s
}

type DescribeApplicationScalingRulesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApplicationScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationScalingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponse) SetHeaders(v map[string]*string) *DescribeApplicationScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationScalingRulesResponse) SetStatusCode(v int32) *DescribeApplicationScalingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponse) SetBody(v *DescribeApplicationScalingRulesResponseBody) *DescribeApplicationScalingRulesResponse {
	s.Body = v
	return s
}

type DescribeApplicationSlbsRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeApplicationSlbsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationSlbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsRequest) SetAppId(v string) *DescribeApplicationSlbsRequest {
	s.AppId = &v
	return s
}

type DescribeApplicationSlbsResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeApplicationSlbsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationSlbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationSlbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponseBody) SetCode(v string) *DescribeApplicationSlbsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetData(v *DescribeApplicationSlbsResponseBodyData) *DescribeApplicationSlbsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetErrorCode(v string) *DescribeApplicationSlbsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetMessage(v string) *DescribeApplicationSlbsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetRequestId(v string) *DescribeApplicationSlbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetSuccess(v bool) *DescribeApplicationSlbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetTraceId(v string) *DescribeApplicationSlbsResponseBody {
	s.TraceId = &v
	return s
}

type DescribeApplicationSlbsResponseBodyData struct {
	Internet      []*DescribeApplicationSlbsResponseBodyDataInternet `json:"Internet,omitempty" xml:"Internet,omitempty" type:"Repeated"`
	InternetIp    *string                                            `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	InternetSlbId *string                                            `json:"InternetSlbId,omitempty" xml:"InternetSlbId,omitempty"`
	Intranet      []*DescribeApplicationSlbsResponseBodyDataIntranet `json:"Intranet,omitempty" xml:"Intranet,omitempty" type:"Repeated"`
	IntranetIp    *string                                            `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	IntranetSlbId *string                                            `json:"IntranetSlbId,omitempty" xml:"IntranetSlbId,omitempty"`
}

func (s DescribeApplicationSlbsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationSlbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponseBodyData) SetInternet(v []*DescribeApplicationSlbsResponseBodyDataInternet) *DescribeApplicationSlbsResponseBodyData {
	s.Internet = v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetInternetIp(v string) *DescribeApplicationSlbsResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetInternetSlbId(v string) *DescribeApplicationSlbsResponseBodyData {
	s.InternetSlbId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetIntranet(v []*DescribeApplicationSlbsResponseBodyDataIntranet) *DescribeApplicationSlbsResponseBodyData {
	s.Intranet = v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetIntranetIp(v string) *DescribeApplicationSlbsResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetIntranetSlbId(v string) *DescribeApplicationSlbsResponseBodyData {
	s.IntranetSlbId = &v
	return s
}

type DescribeApplicationSlbsResponseBodyDataInternet struct {
	HttpsCertId *string `json:"HttpsCertId,omitempty" xml:"HttpsCertId,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TargetPort  *int32  `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s DescribeApplicationSlbsResponseBodyDataInternet) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationSlbsResponseBodyDataInternet) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetHttpsCertId(v string) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.HttpsCertId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetPort(v int32) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.Port = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetProtocol(v string) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.Protocol = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetTargetPort(v int32) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.TargetPort = &v
	return s
}

type DescribeApplicationSlbsResponseBodyDataIntranet struct {
	HttpsCertId *string `json:"HttpsCertId,omitempty" xml:"HttpsCertId,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TargetPort  *int32  `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s DescribeApplicationSlbsResponseBodyDataIntranet) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationSlbsResponseBodyDataIntranet) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetHttpsCertId(v string) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.HttpsCertId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetPort(v int32) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.Port = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetProtocol(v string) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.Protocol = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetTargetPort(v int32) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.TargetPort = &v
	return s
}

type DescribeApplicationSlbsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApplicationSlbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationSlbsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationSlbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponse) SetHeaders(v map[string]*string) *DescribeApplicationSlbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationSlbsResponse) SetStatusCode(v int32) *DescribeApplicationSlbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationSlbsResponse) SetBody(v *DescribeApplicationSlbsResponseBody) *DescribeApplicationSlbsResponse {
	s.Body = v
	return s
}

type DescribeApplicationStatusRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeApplicationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatusRequest) SetAppId(v string) *DescribeApplicationStatusRequest {
	s.AppId = &v
	return s
}

type DescribeApplicationStatusResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeApplicationStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                    `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatusResponseBody) SetCode(v string) *DescribeApplicationStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetData(v *DescribeApplicationStatusResponseBodyData) *DescribeApplicationStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetErrorCode(v string) *DescribeApplicationStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetMessage(v string) *DescribeApplicationStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetRequestId(v string) *DescribeApplicationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetSuccess(v bool) *DescribeApplicationStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetTraceId(v string) *DescribeApplicationStatusResponseBody {
	s.TraceId = &v
	return s
}

type DescribeApplicationStatusResponseBodyData struct {
	AppId                  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ArmsAdvancedEnabled    *string `json:"ArmsAdvancedEnabled,omitempty" xml:"ArmsAdvancedEnabled,omitempty"`
	ArmsApmInfo            *string `json:"ArmsApmInfo,omitempty" xml:"ArmsApmInfo,omitempty"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentStatus          *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	EnableAgent            *bool   `json:"EnableAgent,omitempty" xml:"EnableAgent,omitempty"`
	FileSizeLimit          *int64  `json:"FileSizeLimit,omitempty" xml:"FileSizeLimit,omitempty"`
	LastChangeOrderId      *string `json:"LastChangeOrderId,omitempty" xml:"LastChangeOrderId,omitempty"`
	LastChangeOrderRunning *bool   `json:"LastChangeOrderRunning,omitempty" xml:"LastChangeOrderRunning,omitempty"`
	LastChangeOrderStatus  *string `json:"LastChangeOrderStatus,omitempty" xml:"LastChangeOrderStatus,omitempty"`
	RunningInstances       *int32  `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	SubStatus              *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
}

func (s DescribeApplicationStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatusResponseBodyData) SetAppId(v string) *DescribeApplicationStatusResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetArmsAdvancedEnabled(v string) *DescribeApplicationStatusResponseBodyData {
	s.ArmsAdvancedEnabled = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetArmsApmInfo(v string) *DescribeApplicationStatusResponseBodyData {
	s.ArmsApmInfo = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetCreateTime(v string) *DescribeApplicationStatusResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetCurrentStatus(v string) *DescribeApplicationStatusResponseBodyData {
	s.CurrentStatus = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetEnableAgent(v bool) *DescribeApplicationStatusResponseBodyData {
	s.EnableAgent = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetFileSizeLimit(v int64) *DescribeApplicationStatusResponseBodyData {
	s.FileSizeLimit = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetLastChangeOrderId(v string) *DescribeApplicationStatusResponseBodyData {
	s.LastChangeOrderId = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetLastChangeOrderRunning(v bool) *DescribeApplicationStatusResponseBodyData {
	s.LastChangeOrderRunning = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetLastChangeOrderStatus(v string) *DescribeApplicationStatusResponseBodyData {
	s.LastChangeOrderStatus = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetRunningInstances(v int32) *DescribeApplicationStatusResponseBodyData {
	s.RunningInstances = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetSubStatus(v string) *DescribeApplicationStatusResponseBodyData {
	s.SubStatus = &v
	return s
}

type DescribeApplicationStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApplicationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatusResponse) SetHeaders(v map[string]*string) *DescribeApplicationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationStatusResponse) SetStatusCode(v int32) *DescribeApplicationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationStatusResponse) SetBody(v *DescribeApplicationStatusResponseBody) *DescribeApplicationStatusResponse {
	s.Body = v
	return s
}

type DescribeChangeOrderRequest struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s DescribeChangeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderRequest) SetChangeOrderId(v string) *DescribeChangeOrderRequest {
	s.ChangeOrderId = &v
	return s
}

type DescribeChangeOrderResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                              `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeChangeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderResponseBody) SetCode(v string) *DescribeChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetData(v *DescribeChangeOrderResponseBodyData) *DescribeChangeOrderResponseBody {
	s.Data = v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetErrorCode(v string) *DescribeChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetMessage(v string) *DescribeChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetRequestId(v string) *DescribeChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetSuccess(v bool) *DescribeChangeOrderResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeChangeOrderResponseBody) SetTraceId(v string) *DescribeChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

type DescribeChangeOrderResponseBodyData struct {
	AppId             *string                                         `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName           *string                                         `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ApprovalId        *string                                         `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	Auto              *bool                                           `json:"Auto,omitempty" xml:"Auto,omitempty"`
	BatchCount        *int32                                          `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	BatchType         *string                                         `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	BatchWaitTime     *int32                                          `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	ChangeOrderId     *string                                         `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	CoType            *string                                         `json:"CoType,omitempty" xml:"CoType,omitempty"`
	CoTypeCode        *string                                         `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	CreateTime        *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentPipelineId *string                                         `json:"CurrentPipelineId,omitempty" xml:"CurrentPipelineId,omitempty"`
	Description       *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorMessage      *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Pipelines         []*DescribeChangeOrderResponseBodyDataPipelines `json:"Pipelines,omitempty" xml:"Pipelines,omitempty" type:"Repeated"`
	Status            *int32                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus         *int32                                          `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	SupportRollback   *bool                                           `json:"SupportRollback,omitempty" xml:"SupportRollback,omitempty"`
}

func (s DescribeChangeOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderResponseBodyData) SetAppId(v string) *DescribeChangeOrderResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetAppName(v string) *DescribeChangeOrderResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetApprovalId(v string) *DescribeChangeOrderResponseBodyData {
	s.ApprovalId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetAuto(v bool) *DescribeChangeOrderResponseBodyData {
	s.Auto = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetBatchCount(v int32) *DescribeChangeOrderResponseBodyData {
	s.BatchCount = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetBatchType(v string) *DescribeChangeOrderResponseBodyData {
	s.BatchType = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetBatchWaitTime(v int32) *DescribeChangeOrderResponseBodyData {
	s.BatchWaitTime = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetChangeOrderId(v string) *DescribeChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetCoType(v string) *DescribeChangeOrderResponseBodyData {
	s.CoType = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetCoTypeCode(v string) *DescribeChangeOrderResponseBodyData {
	s.CoTypeCode = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetCreateTime(v string) *DescribeChangeOrderResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetCurrentPipelineId(v string) *DescribeChangeOrderResponseBodyData {
	s.CurrentPipelineId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetDescription(v string) *DescribeChangeOrderResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetErrorMessage(v string) *DescribeChangeOrderResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetPipelines(v []*DescribeChangeOrderResponseBodyDataPipelines) *DescribeChangeOrderResponseBodyData {
	s.Pipelines = v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetStatus(v int32) *DescribeChangeOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetSubStatus(v int32) *DescribeChangeOrderResponseBodyData {
	s.SubStatus = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyData) SetSupportRollback(v bool) *DescribeChangeOrderResponseBodyData {
	s.SupportRollback = &v
	return s
}

type DescribeChangeOrderResponseBodyDataPipelines struct {
	BatchType     *int32  `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	ParallelCount *int32  `json:"ParallelCount,omitempty" xml:"ParallelCount,omitempty"`
	PipelineId    *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	PipelineName  *string `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeChangeOrderResponseBodyDataPipelines) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeOrderResponseBodyDataPipelines) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetBatchType(v int32) *DescribeChangeOrderResponseBodyDataPipelines {
	s.BatchType = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetParallelCount(v int32) *DescribeChangeOrderResponseBodyDataPipelines {
	s.ParallelCount = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetPipelineId(v string) *DescribeChangeOrderResponseBodyDataPipelines {
	s.PipelineId = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetPipelineName(v string) *DescribeChangeOrderResponseBodyDataPipelines {
	s.PipelineName = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetStartTime(v int64) *DescribeChangeOrderResponseBodyDataPipelines {
	s.StartTime = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetStatus(v int32) *DescribeChangeOrderResponseBodyDataPipelines {
	s.Status = &v
	return s
}

func (s *DescribeChangeOrderResponseBodyDataPipelines) SetUpdateTime(v int64) *DescribeChangeOrderResponseBodyDataPipelines {
	s.UpdateTime = &v
	return s
}

type DescribeChangeOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeChangeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderResponse) SetHeaders(v map[string]*string) *DescribeChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeChangeOrderResponse) SetStatusCode(v int32) *DescribeChangeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChangeOrderResponse) SetBody(v *DescribeChangeOrderResponseBody) *DescribeChangeOrderResponse {
	s.Body = v
	return s
}

type DescribeComponentsRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentsRequest) SetAppId(v string) *DescribeComponentsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeComponentsRequest) SetType(v string) *DescribeComponentsRequest {
	s.Type = &v
	return s
}

type DescribeComponentsResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeComponentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                               `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentsResponseBody) SetCode(v string) *DescribeComponentsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetData(v []*DescribeComponentsResponseBodyData) *DescribeComponentsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeComponentsResponseBody) SetErrorCode(v string) *DescribeComponentsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetMessage(v string) *DescribeComponentsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetRequestId(v string) *DescribeComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetSuccess(v bool) *DescribeComponentsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetTraceId(v string) *DescribeComponentsResponseBody {
	s.TraceId = &v
	return s
}

type DescribeComponentsResponseBodyData struct {
	ComponentDescription *string `json:"ComponentDescription,omitempty" xml:"ComponentDescription,omitempty"`
	ComponentKey         *string `json:"ComponentKey,omitempty" xml:"ComponentKey,omitempty"`
	Expired              *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeComponentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeComponentsResponseBodyData) SetComponentDescription(v string) *DescribeComponentsResponseBodyData {
	s.ComponentDescription = &v
	return s
}

func (s *DescribeComponentsResponseBodyData) SetComponentKey(v string) *DescribeComponentsResponseBodyData {
	s.ComponentKey = &v
	return s
}

func (s *DescribeComponentsResponseBodyData) SetExpired(v bool) *DescribeComponentsResponseBodyData {
	s.Expired = &v
	return s
}

func (s *DescribeComponentsResponseBodyData) SetType(v string) *DescribeComponentsResponseBodyData {
	s.Type = &v
	return s
}

type DescribeComponentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComponentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentsResponse) SetHeaders(v map[string]*string) *DescribeComponentsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentsResponse) SetStatusCode(v int32) *DescribeComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentsResponse) SetBody(v *DescribeComponentsResponseBody) *DescribeComponentsResponse {
	s.Body = v
	return s
}

type DescribeConfigMapRequest struct {
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s DescribeConfigMapRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigMapRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapRequest) SetConfigMapId(v int64) *DescribeConfigMapRequest {
	s.ConfigMapId = &v
	return s
}

type DescribeConfigMapResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeConfigMapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                            `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeConfigMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapResponseBody) SetCode(v string) *DescribeConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetData(v *DescribeConfigMapResponseBodyData) *DescribeConfigMapResponseBody {
	s.Data = v
	return s
}

func (s *DescribeConfigMapResponseBody) SetErrorCode(v string) *DescribeConfigMapResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetMessage(v string) *DescribeConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetRequestId(v string) *DescribeConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetSuccess(v bool) *DescribeConfigMapResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConfigMapResponseBody) SetTraceId(v string) *DescribeConfigMapResponseBody {
	s.TraceId = &v
	return s
}

type DescribeConfigMapResponseBodyData struct {
	ConfigMapId *int64                                         `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	CreateTime  *int64                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Data        map[string]interface{}                         `json:"Data,omitempty" xml:"Data,omitempty"`
	Description *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId *string                                        `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	RelateApps  []*DescribeConfigMapResponseBodyDataRelateApps `json:"RelateApps,omitempty" xml:"RelateApps,omitempty" type:"Repeated"`
	UpdateTime  *int64                                         `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeConfigMapResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigMapResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapResponseBodyData) SetConfigMapId(v int64) *DescribeConfigMapResponseBodyData {
	s.ConfigMapId = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetCreateTime(v int64) *DescribeConfigMapResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetData(v map[string]interface{}) *DescribeConfigMapResponseBodyData {
	s.Data = v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetDescription(v string) *DescribeConfigMapResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetName(v string) *DescribeConfigMapResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetNamespaceId(v string) *DescribeConfigMapResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetRelateApps(v []*DescribeConfigMapResponseBodyDataRelateApps) *DescribeConfigMapResponseBodyData {
	s.RelateApps = v
	return s
}

func (s *DescribeConfigMapResponseBodyData) SetUpdateTime(v int64) *DescribeConfigMapResponseBodyData {
	s.UpdateTime = &v
	return s
}

type DescribeConfigMapResponseBodyDataRelateApps struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s DescribeConfigMapResponseBodyDataRelateApps) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigMapResponseBodyDataRelateApps) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapResponseBodyDataRelateApps) SetAppId(v string) *DescribeConfigMapResponseBodyDataRelateApps {
	s.AppId = &v
	return s
}

func (s *DescribeConfigMapResponseBodyDataRelateApps) SetAppName(v string) *DescribeConfigMapResponseBodyDataRelateApps {
	s.AppName = &v
	return s
}

type DescribeConfigMapResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigMapResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigMapResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigMapResponse) SetHeaders(v map[string]*string) *DescribeConfigMapResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigMapResponse) SetStatusCode(v int32) *DescribeConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigMapResponse) SetBody(v *DescribeConfigMapResponseBody) *DescribeConfigMapResponse {
	s.Body = v
	return s
}

type DescribeConfigurationPriceRequest struct {
	Cpu      *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory   *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Workload *string `json:"Workload,omitempty" xml:"Workload,omitempty"`
}

func (s DescribeConfigurationPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceRequest) SetCpu(v int32) *DescribeConfigurationPriceRequest {
	s.Cpu = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) SetMemory(v int32) *DescribeConfigurationPriceRequest {
	s.Memory = &v
	return s
}

func (s *DescribeConfigurationPriceRequest) SetWorkload(v string) *DescribeConfigurationPriceRequest {
	s.Workload = &v
	return s
}

type DescribeConfigurationPriceResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeConfigurationPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                     `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeConfigurationPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBody) SetCode(v string) *DescribeConfigurationPriceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetData(v *DescribeConfigurationPriceResponseBodyData) *DescribeConfigurationPriceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetErrorCode(v string) *DescribeConfigurationPriceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetMessage(v string) *DescribeConfigurationPriceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetRequestId(v string) *DescribeConfigurationPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetSuccess(v bool) *DescribeConfigurationPriceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetTraceId(v string) *DescribeConfigurationPriceResponseBody {
	s.TraceId = &v
	return s
}

type DescribeConfigurationPriceResponseBodyData struct {
	BagUsage *DescribeConfigurationPriceResponseBodyDataBagUsage `json:"BagUsage,omitempty" xml:"BagUsage,omitempty" type:"Struct"`
	Order    *DescribeConfigurationPriceResponseBodyDataOrder    `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	Rules    []*DescribeConfigurationPriceResponseBodyDataRules  `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeConfigurationPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyData) SetBagUsage(v *DescribeConfigurationPriceResponseBodyDataBagUsage) *DescribeConfigurationPriceResponseBodyData {
	s.BagUsage = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyData) SetOrder(v *DescribeConfigurationPriceResponseBodyDataOrder) *DescribeConfigurationPriceResponseBodyData {
	s.Order = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyData) SetRules(v []*DescribeConfigurationPriceResponseBodyDataRules) *DescribeConfigurationPriceResponseBodyData {
	s.Rules = v
	return s
}

type DescribeConfigurationPriceResponseBodyDataBagUsage struct {
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Mem *float32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataBagUsage) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataBagUsage) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataBagUsage) SetCpu(v float32) *DescribeConfigurationPriceResponseBodyDataBagUsage {
	s.Cpu = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataBagUsage) SetMem(v float32) *DescribeConfigurationPriceResponseBodyDataBagUsage {
	s.Mem = &v
	return s
}

type DescribeConfigurationPriceResponseBodyDataOrder struct {
	DiscountAmount *float32  `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	OriginalAmount *float32  `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	RuleIds        []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	TradeAmount    *float32  `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataOrder) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataOrder) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) SetDiscountAmount(v float32) *DescribeConfigurationPriceResponseBodyDataOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) SetOriginalAmount(v float32) *DescribeConfigurationPriceResponseBodyDataOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) SetRuleIds(v []*string) *DescribeConfigurationPriceResponseBodyDataOrder {
	s.RuleIds = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) SetTradeAmount(v float32) *DescribeConfigurationPriceResponseBodyDataOrder {
	s.TradeAmount = &v
	return s
}

type DescribeConfigurationPriceResponseBodyDataRules struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleDescId *int64  `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataRules) SetName(v string) *DescribeConfigurationPriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRules) SetRuleDescId(v int64) *DescribeConfigurationPriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

type DescribeConfigurationPriceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeConfigurationPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigurationPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponse) SetHeaders(v map[string]*string) *DescribeConfigurationPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigurationPriceResponse) SetStatusCode(v int32) *DescribeConfigurationPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigurationPriceResponse) SetBody(v *DescribeConfigurationPriceResponseBody) *DescribeConfigurationPriceResponse {
	s.Body = v
	return s
}

type DescribeEdasContainersResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeEdasContainersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                   `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeEdasContainersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEdasContainersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdasContainersResponseBody) SetCode(v string) *DescribeEdasContainersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetData(v []*DescribeEdasContainersResponseBodyData) *DescribeEdasContainersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetErrorCode(v string) *DescribeEdasContainersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetMessage(v string) *DescribeEdasContainersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetRequestId(v string) *DescribeEdasContainersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetSuccess(v bool) *DescribeEdasContainersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetTraceId(v string) *DescribeEdasContainersResponseBody {
	s.TraceId = &v
	return s
}

type DescribeEdasContainersResponseBodyData struct {
	Disabled             *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
}

func (s DescribeEdasContainersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEdasContainersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEdasContainersResponseBodyData) SetDisabled(v bool) *DescribeEdasContainersResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *DescribeEdasContainersResponseBodyData) SetEdasContainerVersion(v string) *DescribeEdasContainersResponseBodyData {
	s.EdasContainerVersion = &v
	return s
}

type DescribeEdasContainersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEdasContainersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEdasContainersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEdasContainersResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdasContainersResponse) SetHeaders(v map[string]*string) *DescribeEdasContainersResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdasContainersResponse) SetStatusCode(v int32) *DescribeEdasContainersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdasContainersResponse) SetBody(v *DescribeEdasContainersResponseBody) *DescribeEdasContainersResponse {
	s.Body = v
	return s
}

type DescribeGreyTagRouteRequest struct {
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s DescribeGreyTagRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteRequest) SetGreyTagRouteId(v int64) *DescribeGreyTagRouteRequest {
	s.GreyTagRouteId = &v
	return s
}

type DescribeGreyTagRouteResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                               `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeGreyTagRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBody) SetCode(v string) *DescribeGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetData(v *DescribeGreyTagRouteResponseBodyData) *DescribeGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetErrorCode(v string) *DescribeGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetMessage(v string) *DescribeGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetRequestId(v string) *DescribeGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetSuccess(v bool) *DescribeGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBody) SetTraceId(v string) *DescribeGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

type DescribeGreyTagRouteResponseBodyData struct {
	AppId          *string                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime     *int64                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	DubboRules     []*DescribeGreyTagRouteResponseBodyDataDubboRules `json:"DubboRules,omitempty" xml:"DubboRules,omitempty" type:"Repeated"`
	GreyTagRouteId *int64                                            `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
	Name           *string                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	ScRules        []*DescribeGreyTagRouteResponseBodyDataScRules    `json:"ScRules,omitempty" xml:"ScRules,omitempty" type:"Repeated"`
	UpdateTime     *int64                                            `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeGreyTagRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyData) SetAppId(v string) *DescribeGreyTagRouteResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetCreateTime(v int64) *DescribeGreyTagRouteResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetDescription(v string) *DescribeGreyTagRouteResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetDubboRules(v []*DescribeGreyTagRouteResponseBodyDataDubboRules) *DescribeGreyTagRouteResponseBodyData {
	s.DubboRules = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetGreyTagRouteId(v int64) *DescribeGreyTagRouteResponseBodyData {
	s.GreyTagRouteId = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetName(v string) *DescribeGreyTagRouteResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetScRules(v []*DescribeGreyTagRouteResponseBodyDataScRules) *DescribeGreyTagRouteResponseBodyData {
	s.ScRules = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyData) SetUpdateTime(v int64) *DescribeGreyTagRouteResponseBodyData {
	s.UpdateTime = &v
	return s
}

type DescribeGreyTagRouteResponseBodyDataDubboRules struct {
	Condition   *string                                                `json:"condition,omitempty" xml:"condition,omitempty"`
	Group       *string                                                `json:"group,omitempty" xml:"group,omitempty"`
	Items       []*DescribeGreyTagRouteResponseBodyDataDubboRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	MethodName  *string                                                `json:"methodName,omitempty" xml:"methodName,omitempty"`
	ServiceName *string                                                `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	Version     *string                                                `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeGreyTagRouteResponseBodyDataDubboRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataDubboRules) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetCondition(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.Condition = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetGroup(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.Group = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetItems(v []*DescribeGreyTagRouteResponseBodyDataDubboRulesItems) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.Items = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetMethodName(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.MethodName = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetServiceName(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.ServiceName = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRules) SetVersion(v string) *DescribeGreyTagRouteResponseBodyDataDubboRules {
	s.Version = &v
	return s
}

type DescribeGreyTagRouteResponseBodyDataDubboRulesItems struct {
	Cond     *string `json:"cond,omitempty" xml:"cond,omitempty"`
	Expr     *string `json:"expr,omitempty" xml:"expr,omitempty"`
	Index    *int32  `json:"index,omitempty" xml:"index,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeGreyTagRouteResponseBodyDataDubboRulesItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataDubboRulesItems) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetCond(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Cond = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetExpr(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Expr = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetIndex(v int32) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Index = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetName(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Name = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetOperator(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Operator = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetType(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Type = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataDubboRulesItems) SetValue(v string) *DescribeGreyTagRouteResponseBodyDataDubboRulesItems {
	s.Value = &v
	return s
}

type DescribeGreyTagRouteResponseBodyDataScRules struct {
	Condition *string                                             `json:"condition,omitempty" xml:"condition,omitempty"`
	Items     []*DescribeGreyTagRouteResponseBodyDataScRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Path      *string                                             `json:"path,omitempty" xml:"path,omitempty"`
}

func (s DescribeGreyTagRouteResponseBodyDataScRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataScRules) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) SetCondition(v string) *DescribeGreyTagRouteResponseBodyDataScRules {
	s.Condition = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) SetItems(v []*DescribeGreyTagRouteResponseBodyDataScRulesItems) *DescribeGreyTagRouteResponseBodyDataScRules {
	s.Items = v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRules) SetPath(v string) *DescribeGreyTagRouteResponseBodyDataScRules {
	s.Path = &v
	return s
}

type DescribeGreyTagRouteResponseBodyDataScRulesItems struct {
	Cond     *string `json:"cond,omitempty" xml:"cond,omitempty"`
	Expr     *string `json:"expr,omitempty" xml:"expr,omitempty"`
	Index    *int32  `json:"index,omitempty" xml:"index,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeGreyTagRouteResponseBodyDataScRulesItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreyTagRouteResponseBodyDataScRulesItems) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetCond(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Cond = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetExpr(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Expr = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetIndex(v int32) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Index = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetName(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Name = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetOperator(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Operator = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetType(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Type = &v
	return s
}

func (s *DescribeGreyTagRouteResponseBodyDataScRulesItems) SetValue(v string) *DescribeGreyTagRouteResponseBodyDataScRulesItems {
	s.Value = &v
	return s
}

type DescribeGreyTagRouteResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGreyTagRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponse) SetHeaders(v map[string]*string) *DescribeGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *DescribeGreyTagRouteResponse) SetStatusCode(v int32) *DescribeGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGreyTagRouteResponse) SetBody(v *DescribeGreyTagRouteResponseBody) *DescribeGreyTagRouteResponse {
	s.Body = v
	return s
}

type DescribeIngressRequest struct {
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s DescribeIngressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIngressRequest) GoString() string {
	return s.String()
}

func (s *DescribeIngressRequest) SetIngressId(v int64) *DescribeIngressRequest {
	s.IngressId = &v
	return s
}

type DescribeIngressResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                          `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeIngressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIngressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBody) SetCode(v string) *DescribeIngressResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeIngressResponseBody) SetData(v *DescribeIngressResponseBodyData) *DescribeIngressResponseBody {
	s.Data = v
	return s
}

func (s *DescribeIngressResponseBody) SetErrorCode(v string) *DescribeIngressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeIngressResponseBody) SetMessage(v string) *DescribeIngressResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeIngressResponseBody) SetRequestId(v string) *DescribeIngressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIngressResponseBody) SetSuccess(v bool) *DescribeIngressResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeIngressResponseBody) SetTraceId(v string) *DescribeIngressResponseBody {
	s.TraceId = &v
	return s
}

type DescribeIngressResponseBodyData struct {
	CertId           *string                                     `json:"CertId,omitempty" xml:"CertId,omitempty"`
	DefaultRule      *DescribeIngressResponseBodyDataDefaultRule `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty" type:"Struct"`
	Description      *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	Id               *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	ListenerPort     *int32                                      `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol *string                                     `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalanceType  *string                                     `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	Name             *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId      *string                                     `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Rules            []*DescribeIngressResponseBodyDataRules     `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	SlbId            *string                                     `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbType          *string                                     `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s DescribeIngressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeIngressResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBodyData) SetCertId(v string) *DescribeIngressResponseBodyData {
	s.CertId = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetDefaultRule(v *DescribeIngressResponseBodyDataDefaultRule) *DescribeIngressResponseBodyData {
	s.DefaultRule = v
	return s
}

func (s *DescribeIngressResponseBodyData) SetDescription(v string) *DescribeIngressResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetId(v int64) *DescribeIngressResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetListenerPort(v int32) *DescribeIngressResponseBodyData {
	s.ListenerPort = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetListenerProtocol(v string) *DescribeIngressResponseBodyData {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetLoadBalanceType(v string) *DescribeIngressResponseBodyData {
	s.LoadBalanceType = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetName(v string) *DescribeIngressResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetNamespaceId(v string) *DescribeIngressResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetRules(v []*DescribeIngressResponseBodyDataRules) *DescribeIngressResponseBodyData {
	s.Rules = v
	return s
}

func (s *DescribeIngressResponseBodyData) SetSlbId(v string) *DescribeIngressResponseBodyData {
	s.SlbId = &v
	return s
}

func (s *DescribeIngressResponseBodyData) SetSlbType(v string) *DescribeIngressResponseBodyData {
	s.SlbType = &v
	return s
}

type DescribeIngressResponseBodyDataDefaultRule struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BackendProtocol *string `json:"BackendProtocol,omitempty" xml:"BackendProtocol,omitempty"`
	ContainerPort   *int32  `json:"ContainerPort,omitempty" xml:"ContainerPort,omitempty"`
}

func (s DescribeIngressResponseBodyDataDefaultRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeIngressResponseBodyDataDefaultRule) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBodyDataDefaultRule) SetAppId(v string) *DescribeIngressResponseBodyDataDefaultRule {
	s.AppId = &v
	return s
}

func (s *DescribeIngressResponseBodyDataDefaultRule) SetAppName(v string) *DescribeIngressResponseBodyDataDefaultRule {
	s.AppName = &v
	return s
}

func (s *DescribeIngressResponseBodyDataDefaultRule) SetBackendProtocol(v string) *DescribeIngressResponseBodyDataDefaultRule {
	s.BackendProtocol = &v
	return s
}

func (s *DescribeIngressResponseBodyDataDefaultRule) SetContainerPort(v int32) *DescribeIngressResponseBodyDataDefaultRule {
	s.ContainerPort = &v
	return s
}

type DescribeIngressResponseBodyDataRules struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BackendProtocol *string `json:"BackendProtocol,omitempty" xml:"BackendProtocol,omitempty"`
	ContainerPort   *int32  `json:"ContainerPort,omitempty" xml:"ContainerPort,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeIngressResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeIngressResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponseBodyDataRules) SetAppId(v string) *DescribeIngressResponseBodyDataRules {
	s.AppId = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetAppName(v string) *DescribeIngressResponseBodyDataRules {
	s.AppName = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetBackendProtocol(v string) *DescribeIngressResponseBodyDataRules {
	s.BackendProtocol = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetContainerPort(v int32) *DescribeIngressResponseBodyDataRules {
	s.ContainerPort = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetDomain(v string) *DescribeIngressResponseBodyDataRules {
	s.Domain = &v
	return s
}

func (s *DescribeIngressResponseBodyDataRules) SetPath(v string) *DescribeIngressResponseBodyDataRules {
	s.Path = &v
	return s
}

type DescribeIngressResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIngressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIngressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIngressResponse) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponse) SetHeaders(v map[string]*string) *DescribeIngressResponse {
	s.Headers = v
	return s
}

func (s *DescribeIngressResponse) SetStatusCode(v int32) *DescribeIngressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIngressResponse) SetBody(v *DescribeIngressResponseBody) *DescribeIngressResponse {
	s.Body = v
	return s
}

type DescribeInstanceLogRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLogRequest) SetInstanceId(v string) *DescribeInstanceLogRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceLogResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeInstanceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLogResponseBody) SetCode(v string) *DescribeInstanceLogResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetData(v string) *DescribeInstanceLogResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetErrorCode(v string) *DescribeInstanceLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetMessage(v string) *DescribeInstanceLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetRequestId(v string) *DescribeInstanceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetSuccess(v bool) *DescribeInstanceLogResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceLogResponseBody) SetTraceId(v string) *DescribeInstanceLogResponseBody {
	s.TraceId = &v
	return s
}

type DescribeInstanceLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLogResponse) SetHeaders(v map[string]*string) *DescribeInstanceLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceLogResponse) SetStatusCode(v int32) *DescribeInstanceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceLogResponse) SetBody(v *DescribeInstanceLogResponseBody) *DescribeInstanceLogResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecificationsResponseBody struct {
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeInstanceSpecificationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeInstanceSpecificationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecificationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecificationsResponseBody) SetCode(v string) *DescribeInstanceSpecificationsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetData(v []*DescribeInstanceSpecificationsResponseBodyData) *DescribeInstanceSpecificationsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetErrorCode(v string) *DescribeInstanceSpecificationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetMessage(v string) *DescribeInstanceSpecificationsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetRequestId(v string) *DescribeInstanceSpecificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetSuccess(v bool) *DescribeInstanceSpecificationsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBody) SetTraceId(v string) *DescribeInstanceSpecificationsResponseBody {
	s.TraceId = &v
	return s
}

type DescribeInstanceSpecificationsResponseBodyData struct {
	Cpu      *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Enable   *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Id       *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Memory   *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	SpecInfo *string `json:"SpecInfo,omitempty" xml:"SpecInfo,omitempty"`
	Version  *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeInstanceSpecificationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecificationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetCpu(v int32) *DescribeInstanceSpecificationsResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetEnable(v bool) *DescribeInstanceSpecificationsResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetId(v int32) *DescribeInstanceSpecificationsResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetMemory(v int32) *DescribeInstanceSpecificationsResponseBodyData {
	s.Memory = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetSpecInfo(v string) *DescribeInstanceSpecificationsResponseBodyData {
	s.SpecInfo = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponseBodyData) SetVersion(v int32) *DescribeInstanceSpecificationsResponseBodyData {
	s.Version = &v
	return s
}

type DescribeInstanceSpecificationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceSpecificationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSpecificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecificationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecificationsResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecificationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecificationsResponse) SetStatusCode(v int32) *DescribeInstanceSpecificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSpecificationsResponse) SetBody(v *DescribeInstanceSpecificationsResponseBody) *DescribeInstanceSpecificationsResponse {
	s.Body = v
	return s
}

type DescribeJobStatusRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusRequest) SetAppId(v string) *DescribeJobStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeJobStatusRequest) SetJobId(v string) *DescribeJobStatusRequest {
	s.JobId = &v
	return s
}

type DescribeJobStatusResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                            `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponseBody) SetCode(v string) *DescribeJobStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetData(v *DescribeJobStatusResponseBodyData) *DescribeJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeJobStatusResponseBody) SetErrorCode(v string) *DescribeJobStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetMessage(v string) *DescribeJobStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetRequestId(v string) *DescribeJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetSuccess(v bool) *DescribeJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetTraceId(v string) *DescribeJobStatusResponseBody {
	s.TraceId = &v
	return s
}

type DescribeJobStatusResponseBodyData struct {
	Active         *int64  `json:"Active,omitempty" xml:"Active,omitempty"`
	CompletionTime *int64  `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	Failed         *int64  `json:"Failed,omitempty" xml:"Failed,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State          *string `json:"State,omitempty" xml:"State,omitempty"`
	Succeeded      *int64  `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s DescribeJobStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponseBodyData) SetActive(v int64) *DescribeJobStatusResponseBodyData {
	s.Active = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetCompletionTime(v int64) *DescribeJobStatusResponseBodyData {
	s.CompletionTime = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetFailed(v int64) *DescribeJobStatusResponseBodyData {
	s.Failed = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetJobId(v string) *DescribeJobStatusResponseBodyData {
	s.JobId = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetMessage(v string) *DescribeJobStatusResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetStartTime(v int64) *DescribeJobStatusResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetState(v string) *DescribeJobStatusResponseBodyData {
	s.State = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetSucceeded(v int64) *DescribeJobStatusResponseBodyData {
	s.Succeeded = &v
	return s
}

type DescribeJobStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponse) SetHeaders(v map[string]*string) *DescribeJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobStatusResponse) SetStatusCode(v int32) *DescribeJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobStatusResponse) SetBody(v *DescribeJobStatusResponseBody) *DescribeJobStatusResponse {
	s.Body = v
	return s
}

type DescribeNamespaceRequest struct {
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceRequest) SetNamespaceId(v string) *DescribeNamespaceRequest {
	s.NamespaceId = &v
	return s
}

type DescribeNamespaceResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                            `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponseBody) SetCode(v string) *DescribeNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetData(v *DescribeNamespaceResponseBodyData) *DescribeNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNamespaceResponseBody) SetErrorCode(v string) *DescribeNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetMessage(v string) *DescribeNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetRequestId(v string) *DescribeNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetSuccess(v bool) *DescribeNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetTraceId(v string) *DescribeNamespaceResponseBody {
	s.TraceId = &v
	return s
}

type DescribeNamespaceResponseBodyData struct {
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	NamespaceId          *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName        *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeNamespaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponseBodyData) SetNamespaceDescription(v string) *DescribeNamespaceResponseBodyData {
	s.NamespaceDescription = &v
	return s
}

func (s *DescribeNamespaceResponseBodyData) SetNamespaceId(v string) *DescribeNamespaceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespaceResponseBodyData) SetNamespaceName(v string) *DescribeNamespaceResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespaceResponseBodyData) SetRegionId(v string) *DescribeNamespaceResponseBodyData {
	s.RegionId = &v
	return s
}

type DescribeNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponse) SetHeaders(v map[string]*string) *DescribeNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceResponse) SetStatusCode(v int32) *DescribeNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceResponse) SetBody(v *DescribeNamespaceResponseBody) *DescribeNamespaceResponse {
	s.Body = v
	return s
}

type DescribeNamespaceListRequest struct {
	ContainCustom      *bool `json:"ContainCustom,omitempty" xml:"ContainCustom,omitempty"`
	HybridCloudExclude *bool `json:"HybridCloudExclude,omitempty" xml:"HybridCloudExclude,omitempty"`
}

func (s DescribeNamespaceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceListRequest) SetContainCustom(v bool) *DescribeNamespaceListRequest {
	s.ContainCustom = &v
	return s
}

func (s *DescribeNamespaceListRequest) SetHybridCloudExclude(v bool) *DescribeNamespaceListRequest {
	s.HybridCloudExclude = &v
	return s
}

type DescribeNamespaceListResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeNamespaceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeNamespaceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceListResponseBody) SetCode(v string) *DescribeNamespaceListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetData(v []*DescribeNamespaceListResponseBodyData) *DescribeNamespaceListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetErrorCode(v string) *DescribeNamespaceListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetMessage(v string) *DescribeNamespaceListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetRequestId(v string) *DescribeNamespaceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetSuccess(v bool) *DescribeNamespaceListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespaceListResponseBody) SetTraceId(v string) *DescribeNamespaceListResponseBody {
	s.TraceId = &v
	return s
}

type DescribeNamespaceListResponseBodyData struct {
	AgentInstall      *string `json:"AgentInstall,omitempty" xml:"AgentInstall,omitempty"`
	Current           *bool   `json:"Current,omitempty" xml:"Current,omitempty"`
	Custom            *bool   `json:"Custom,omitempty" xml:"Custom,omitempty"`
	HybridCloudEnable *bool   `json:"HybridCloudEnable,omitempty" xml:"HybridCloudEnable,omitempty"`
	NamespaceId       *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName     *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId         *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNamespaceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceListResponseBodyData) SetAgentInstall(v string) *DescribeNamespaceListResponseBodyData {
	s.AgentInstall = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetCurrent(v bool) *DescribeNamespaceListResponseBodyData {
	s.Current = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetCustom(v bool) *DescribeNamespaceListResponseBodyData {
	s.Custom = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetHybridCloudEnable(v bool) *DescribeNamespaceListResponseBodyData {
	s.HybridCloudEnable = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetNamespaceId(v string) *DescribeNamespaceListResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetNamespaceName(v string) *DescribeNamespaceListResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetRegionId(v string) *DescribeNamespaceListResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetSecurityGroupId(v string) *DescribeNamespaceListResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetVSwitchId(v string) *DescribeNamespaceListResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNamespaceListResponseBodyData) SetVpcId(v string) *DescribeNamespaceListResponseBodyData {
	s.VpcId = &v
	return s
}

type DescribeNamespaceListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNamespaceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNamespaceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceListResponse) SetHeaders(v map[string]*string) *DescribeNamespaceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceListResponse) SetStatusCode(v int32) *DescribeNamespaceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceListResponse) SetBody(v *DescribeNamespaceListResponseBody) *DescribeNamespaceListResponse {
	s.Body = v
	return s
}

type DescribeNamespaceResourcesRequest struct {
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeNamespaceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResourcesRequest) SetNamespaceId(v string) *DescribeNamespaceResourcesRequest {
	s.NamespaceId = &v
	return s
}

type DescribeNamespaceResourcesResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeNamespaceResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                     `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeNamespaceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResourcesResponseBody) SetCode(v string) *DescribeNamespaceResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetData(v *DescribeNamespaceResourcesResponseBodyData) *DescribeNamespaceResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetErrorCode(v string) *DescribeNamespaceResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetMessage(v string) *DescribeNamespaceResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetRequestId(v string) *DescribeNamespaceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetSuccess(v bool) *DescribeNamespaceResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBody) SetTraceId(v string) *DescribeNamespaceResourcesResponseBody {
	s.TraceId = &v
	return s
}

type DescribeNamespaceResourcesResponseBodyData struct {
	AppCount               *int64  `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	BelongRegion           *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	JumpServerAppId        *string `json:"JumpServerAppId,omitempty" xml:"JumpServerAppId,omitempty"`
	JumpServerIp           *string `json:"JumpServerIp,omitempty" xml:"JumpServerIp,omitempty"`
	LastChangeOrderId      *string `json:"LastChangeOrderId,omitempty" xml:"LastChangeOrderId,omitempty"`
	LastChangeOrderRunning *bool   `json:"LastChangeOrderRunning,omitempty" xml:"LastChangeOrderRunning,omitempty"`
	LastChangeOrderStatus  *string `json:"LastChangeOrderStatus,omitempty" xml:"LastChangeOrderStatus,omitempty"`
	NamespaceId            *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName          *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	NotificationExpired    *bool   `json:"NotificationExpired,omitempty" xml:"NotificationExpired,omitempty"`
	SecurityGroupId        *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	TenantId               *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VSwitchId              *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName            *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	VpcId                  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName                *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeNamespaceResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetAppCount(v int64) *DescribeNamespaceResourcesResponseBodyData {
	s.AppCount = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetBelongRegion(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.BelongRegion = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetDescription(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetJumpServerAppId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.JumpServerAppId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetJumpServerIp(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.JumpServerIp = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetLastChangeOrderId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.LastChangeOrderId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetLastChangeOrderRunning(v bool) *DescribeNamespaceResourcesResponseBodyData {
	s.LastChangeOrderRunning = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetLastChangeOrderStatus(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.LastChangeOrderStatus = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetNamespaceId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetNamespaceName(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetNotificationExpired(v bool) *DescribeNamespaceResourcesResponseBodyData {
	s.NotificationExpired = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetSecurityGroupId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetTenantId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetUserId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetVSwitchId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetVSwitchName(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.VSwitchName = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetVpcId(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeNamespaceResourcesResponseBodyData) SetVpcName(v string) *DescribeNamespaceResourcesResponseBodyData {
	s.VpcName = &v
	return s
}

type DescribeNamespaceResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNamespaceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNamespaceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespaceResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResourcesResponse) SetHeaders(v map[string]*string) *DescribeNamespaceResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespaceResourcesResponse) SetStatusCode(v int32) *DescribeNamespaceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespaceResourcesResponse) SetBody(v *DescribeNamespaceResourcesResponseBody) *DescribeNamespaceResourcesResponse {
	s.Body = v
	return s
}

type DescribeNamespacesRequest struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequest) SetCurrentPage(v int32) *DescribeNamespacesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNamespacesRequest) SetPageSize(v int32) *DescribeNamespacesRequest {
	s.PageSize = &v
	return s
}

type DescribeNamespacesResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                             `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBody) SetCode(v string) *DescribeNamespacesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetData(v *DescribeNamespacesResponseBodyData) *DescribeNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNamespacesResponseBody) SetErrorCode(v string) *DescribeNamespacesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetMessage(v string) *DescribeNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetRequestId(v string) *DescribeNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetSuccess(v bool) *DescribeNamespacesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetTraceId(v string) *DescribeNamespacesResponseBody {
	s.TraceId = &v
	return s
}

type DescribeNamespacesResponseBodyData struct {
	CurrentPage *int32                                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Namespaces  []*DescribeNamespacesResponseBodyDataNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	PageSize    *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize   *int32                                          `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeNamespacesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyData) SetCurrentPage(v int32) *DescribeNamespacesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNamespacesResponseBodyData) SetNamespaces(v []*DescribeNamespacesResponseBodyDataNamespaces) *DescribeNamespacesResponseBodyData {
	s.Namespaces = v
	return s
}

func (s *DescribeNamespacesResponseBodyData) SetPageSize(v int32) *DescribeNamespacesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesResponseBodyData) SetTotalSize(v int32) *DescribeNamespacesResponseBodyData {
	s.TotalSize = &v
	return s
}

type DescribeNamespacesResponseBodyDataNamespaces struct {
	AccessKey            *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	NamespaceId          *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName        *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecretKey            *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeNamespacesResponseBodyDataNamespaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponseBodyDataNamespaces) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetAccessKey(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.AccessKey = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetNamespaceDescription(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.NamespaceDescription = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetNamespaceId(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetNamespaceName(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetRegionId(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.RegionId = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetSecretKey(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.SecretKey = &v
	return s
}

func (s *DescribeNamespacesResponseBodyDataNamespaces) SetTenantId(v string) *DescribeNamespacesResponseBodyDataNamespaces {
	s.TenantId = &v
	return s
}

type DescribeNamespacesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponse) SetHeaders(v map[string]*string) *DescribeNamespacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespacesResponse) SetStatusCode(v int32) *DescribeNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespacesResponse) SetBody(v *DescribeNamespacesResponseBody) *DescribeNamespacesResponse {
	s.Body = v
	return s
}

type DescribePipelineRequest struct {
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s DescribePipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineRequest) GoString() string {
	return s.String()
}

func (s *DescribePipelineRequest) SetPipelineId(v string) *DescribePipelineRequest {
	s.PipelineId = &v
	return s
}

type DescribePipelineResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePipelineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBody) SetCode(v string) *DescribePipelineResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePipelineResponseBody) SetData(v *DescribePipelineResponseBodyData) *DescribePipelineResponseBody {
	s.Data = v
	return s
}

func (s *DescribePipelineResponseBody) SetErrorCode(v string) *DescribePipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribePipelineResponseBody) SetMessage(v string) *DescribePipelineResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePipelineResponseBody) SetRequestId(v string) *DescribePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePipelineResponseBody) SetSuccess(v bool) *DescribePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePipelineResponseBody) SetTraceId(v string) *DescribePipelineResponseBody {
	s.TraceId = &v
	return s
}

type DescribePipelineResponseBodyData struct {
	CoStatus       *string                                      `json:"CoStatus,omitempty" xml:"CoStatus,omitempty"`
	CurrentStageId *string                                      `json:"CurrentStageId,omitempty" xml:"CurrentStageId,omitempty"`
	NextPipelineId *string                                      `json:"NextPipelineId,omitempty" xml:"NextPipelineId,omitempty"`
	PipelineId     *string                                      `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	PipelineName   *string                                      `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	PipelineStatus *int32                                       `json:"PipelineStatus,omitempty" xml:"PipelineStatus,omitempty"`
	ShowBatch      *bool                                        `json:"ShowBatch,omitempty" xml:"ShowBatch,omitempty"`
	StageList      []*DescribePipelineResponseBodyDataStageList `json:"StageList,omitempty" xml:"StageList,omitempty" type:"Repeated"`
}

func (s DescribePipelineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBodyData) SetCoStatus(v string) *DescribePipelineResponseBodyData {
	s.CoStatus = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetCurrentStageId(v string) *DescribePipelineResponseBodyData {
	s.CurrentStageId = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetNextPipelineId(v string) *DescribePipelineResponseBodyData {
	s.NextPipelineId = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetPipelineId(v string) *DescribePipelineResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetPipelineName(v string) *DescribePipelineResponseBodyData {
	s.PipelineName = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetPipelineStatus(v int32) *DescribePipelineResponseBodyData {
	s.PipelineStatus = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetShowBatch(v bool) *DescribePipelineResponseBodyData {
	s.ShowBatch = &v
	return s
}

func (s *DescribePipelineResponseBodyData) SetStageList(v []*DescribePipelineResponseBodyDataStageList) *DescribePipelineResponseBodyData {
	s.StageList = v
	return s
}

type DescribePipelineResponseBodyDataStageList struct {
	ExecutorType *int32                                               `json:"ExecutorType,omitempty" xml:"ExecutorType,omitempty"`
	StageId      *string                                              `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName    *string                                              `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Status       *int32                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskList     []*DescribePipelineResponseBodyDataStageListTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s DescribePipelineResponseBodyDataStageList) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineResponseBodyDataStageList) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBodyDataStageList) SetExecutorType(v int32) *DescribePipelineResponseBodyDataStageList {
	s.ExecutorType = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageList) SetStageId(v string) *DescribePipelineResponseBodyDataStageList {
	s.StageId = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageList) SetStageName(v string) *DescribePipelineResponseBodyDataStageList {
	s.StageName = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageList) SetStatus(v int32) *DescribePipelineResponseBodyDataStageList {
	s.Status = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageList) SetTaskList(v []*DescribePipelineResponseBodyDataStageListTaskList) *DescribePipelineResponseBodyDataStageList {
	s.TaskList = v
	return s
}

type DescribePipelineResponseBodyDataStageListTaskList struct {
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorIgnore      *int32  `json:"ErrorIgnore,omitempty" xml:"ErrorIgnore,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ShowManualIgnore *bool   `json:"ShowManualIgnore,omitempty" xml:"ShowManualIgnore,omitempty"`
	StageId          *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId           *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName         *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribePipelineResponseBodyDataStageListTaskList) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineResponseBodyDataStageListTaskList) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetErrorCode(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.ErrorCode = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetErrorIgnore(v int32) *DescribePipelineResponseBodyDataStageListTaskList {
	s.ErrorIgnore = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetErrorMessage(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.ErrorMessage = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetMessage(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.Message = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetShowManualIgnore(v bool) *DescribePipelineResponseBodyDataStageListTaskList {
	s.ShowManualIgnore = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetStageId(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.StageId = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetStatus(v int32) *DescribePipelineResponseBodyDataStageListTaskList {
	s.Status = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetTaskId(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.TaskId = &v
	return s
}

func (s *DescribePipelineResponseBodyDataStageListTaskList) SetTaskName(v string) *DescribePipelineResponseBodyDataStageListTaskList {
	s.TaskName = &v
	return s
}

type DescribePipelineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePipelineResponse) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponse) SetHeaders(v map[string]*string) *DescribePipelineResponse {
	s.Headers = v
	return s
}

func (s *DescribePipelineResponse) SetStatusCode(v int32) *DescribePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePipelineResponse) SetBody(v *DescribePipelineResponseBody) *DescribePipelineResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v int32) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName      *string                                                 `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RecommendZones *DescribeRegionsResponseBodyRegionsRegionRecommendZones `json:"RecommendZones,omitempty" xml:"RecommendZones,omitempty" type:"Struct"`
	RegionEndpoint *string                                                 `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRecommendZones(v *DescribeRegionsResponseBodyRegionsRegionRecommendZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.RecommendZones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionRecommendZones struct {
	RecommendZone []*string `json:"RecommendZone,omitempty" xml:"RecommendZone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionRecommendZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionRecommendZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionRecommendZones) SetRecommendZone(v []*string) *DescribeRegionsResponseBodyRegionsRegionRecommendZones {
	s.RecommendZone = v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DisableApplicationScalingRuleRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s DisableApplicationScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleRequest) SetAppId(v string) *DisableApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *DisableApplicationScalingRuleRequest) SetScalingRuleName(v string) *DisableApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

type DisableApplicationScalingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DisableApplicationScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponseBody) SetRequestId(v string) *DisableApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) SetTraceId(v string) *DisableApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

type DisableApplicationScalingRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableApplicationScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *DisableApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationScalingRuleResponse) SetStatusCode(v int32) *DisableApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationScalingRuleResponse) SetBody(v *DisableApplicationScalingRuleResponseBody) *DisableApplicationScalingRuleResponse {
	s.Body = v
	return s
}

type EnableApplicationScalingRuleRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s EnableApplicationScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableApplicationScalingRuleRequest) SetAppId(v string) *EnableApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *EnableApplicationScalingRuleRequest) SetScalingRuleName(v string) *EnableApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

type EnableApplicationScalingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s EnableApplicationScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableApplicationScalingRuleResponseBody) SetRequestId(v string) *EnableApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableApplicationScalingRuleResponseBody) SetTraceId(v string) *EnableApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

type EnableApplicationScalingRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableApplicationScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *EnableApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableApplicationScalingRuleResponse) SetStatusCode(v int32) *EnableApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableApplicationScalingRuleResponse) SetBody(v *EnableApplicationScalingRuleResponseBody) *EnableApplicationScalingRuleResponse {
	s.Body = v
	return s
}

type ExecJobRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Command         *string `json:"Command,omitempty" xml:"Command,omitempty"`
	CommandArgs     *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	Envs            *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	EventId         *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	JarStartArgs    *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
}

func (s ExecJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecJobRequest) GoString() string {
	return s.String()
}

func (s *ExecJobRequest) SetAppId(v string) *ExecJobRequest {
	s.AppId = &v
	return s
}

func (s *ExecJobRequest) SetCommand(v string) *ExecJobRequest {
	s.Command = &v
	return s
}

func (s *ExecJobRequest) SetCommandArgs(v string) *ExecJobRequest {
	s.CommandArgs = &v
	return s
}

func (s *ExecJobRequest) SetEnvs(v string) *ExecJobRequest {
	s.Envs = &v
	return s
}

func (s *ExecJobRequest) SetEventId(v string) *ExecJobRequest {
	s.EventId = &v
	return s
}

func (s *ExecJobRequest) SetJarStartArgs(v string) *ExecJobRequest {
	s.JarStartArgs = &v
	return s
}

func (s *ExecJobRequest) SetJarStartOptions(v string) *ExecJobRequest {
	s.JarStartOptions = &v
	return s
}

func (s *ExecJobRequest) SetWarStartOptions(v string) *ExecJobRequest {
	s.WarStartOptions = &v
	return s
}

type ExecJobResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ExecJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ExecJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecJobResponseBody) GoString() string {
	return s.String()
}

func (s *ExecJobResponseBody) SetCode(v string) *ExecJobResponseBody {
	s.Code = &v
	return s
}

func (s *ExecJobResponseBody) SetData(v *ExecJobResponseBodyData) *ExecJobResponseBody {
	s.Data = v
	return s
}

func (s *ExecJobResponseBody) SetErrorCode(v string) *ExecJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ExecJobResponseBody) SetMessage(v string) *ExecJobResponseBody {
	s.Message = &v
	return s
}

func (s *ExecJobResponseBody) SetRequestId(v string) *ExecJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecJobResponseBody) SetSuccess(v bool) *ExecJobResponseBody {
	s.Success = &v
	return s
}

func (s *ExecJobResponseBody) SetTraceId(v string) *ExecJobResponseBody {
	s.TraceId = &v
	return s
}

type ExecJobResponseBodyData struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg     *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecJobResponseBodyData) SetCode(v string) *ExecJobResponseBodyData {
	s.Code = &v
	return s
}

func (s *ExecJobResponseBodyData) SetData(v string) *ExecJobResponseBodyData {
	s.Data = &v
	return s
}

func (s *ExecJobResponseBodyData) SetMsg(v string) *ExecJobResponseBodyData {
	s.Msg = &v
	return s
}

func (s *ExecJobResponseBodyData) SetSuccess(v string) *ExecJobResponseBodyData {
	s.Success = &v
	return s
}

type ExecJobResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecJobResponse) GoString() string {
	return s.String()
}

func (s *ExecJobResponse) SetHeaders(v map[string]*string) *ExecJobResponse {
	s.Headers = v
	return s
}

func (s *ExecJobResponse) SetStatusCode(v int32) *ExecJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecJobResponse) SetBody(v *ExecJobResponseBody) *ExecJobResponse {
	s.Body = v
	return s
}

type ListAppEventsRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EventType   *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ObjectKind  *string `json:"ObjectKind,omitempty" xml:"ObjectKind,omitempty"`
	ObjectName  *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ListAppEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppEventsRequest) GoString() string {
	return s.String()
}

func (s *ListAppEventsRequest) SetAppId(v string) *ListAppEventsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppEventsRequest) SetCurrentPage(v int32) *ListAppEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAppEventsRequest) SetEventType(v string) *ListAppEventsRequest {
	s.EventType = &v
	return s
}

func (s *ListAppEventsRequest) SetNamespace(v string) *ListAppEventsRequest {
	s.Namespace = &v
	return s
}

func (s *ListAppEventsRequest) SetObjectKind(v string) *ListAppEventsRequest {
	s.ObjectKind = &v
	return s
}

func (s *ListAppEventsRequest) SetObjectName(v string) *ListAppEventsRequest {
	s.ObjectName = &v
	return s
}

func (s *ListAppEventsRequest) SetPageSize(v int32) *ListAppEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppEventsRequest) SetReason(v string) *ListAppEventsRequest {
	s.Reason = &v
	return s
}

type ListAppEventsResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListAppEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppEventsResponseBody) SetCode(v string) *ListAppEventsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppEventsResponseBody) SetData(v *ListAppEventsResponseBodyData) *ListAppEventsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppEventsResponseBody) SetErrorCode(v string) *ListAppEventsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAppEventsResponseBody) SetMessage(v string) *ListAppEventsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppEventsResponseBody) SetRequestId(v string) *ListAppEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppEventsResponseBody) SetSuccess(v bool) *ListAppEventsResponseBody {
	s.Success = &v
	return s
}

type ListAppEventsResponseBodyData struct {
	AppEventEntity []*ListAppEventsResponseBodyDataAppEventEntity `json:"AppEventEntity,omitempty" xml:"AppEventEntity,omitempty" type:"Repeated"`
	CurrentPage    *int32                                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize       *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize      *int32                                         `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListAppEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppEventsResponseBodyData) SetAppEventEntity(v []*ListAppEventsResponseBodyDataAppEventEntity) *ListAppEventsResponseBodyData {
	s.AppEventEntity = v
	return s
}

func (s *ListAppEventsResponseBodyData) SetCurrentPage(v int32) *ListAppEventsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListAppEventsResponseBodyData) SetPageSize(v int32) *ListAppEventsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAppEventsResponseBodyData) SetTotalSize(v int32) *ListAppEventsResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListAppEventsResponseBodyDataAppEventEntity struct {
	EventType      *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	LastTimestamp  *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ObjectKind     *string `json:"ObjectKind,omitempty" xml:"ObjectKind,omitempty"`
	ObjectName     *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ListAppEventsResponseBodyDataAppEventEntity) String() string {
	return tea.Prettify(s)
}

func (s ListAppEventsResponseBodyDataAppEventEntity) GoString() string {
	return s.String()
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetEventType(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.EventType = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetFirstTimestamp(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.FirstTimestamp = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetLastTimestamp(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.LastTimestamp = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetMessage(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.Message = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetObjectKind(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.ObjectKind = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetObjectName(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.ObjectName = &v
	return s
}

func (s *ListAppEventsResponseBodyDataAppEventEntity) SetReason(v string) *ListAppEventsResponseBodyDataAppEventEntity {
	s.Reason = &v
	return s
}

type ListAppEventsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppEventsResponse) GoString() string {
	return s.String()
}

func (s *ListAppEventsResponse) SetHeaders(v map[string]*string) *ListAppEventsResponse {
	s.Headers = v
	return s
}

func (s *ListAppEventsResponse) SetStatusCode(v int32) *ListAppEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppEventsResponse) SetBody(v *ListAppEventsResponseBody) *ListAppEventsResponse {
	s.Body = v
	return s
}

type ListAppServicesPageRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListAppServicesPageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppServicesPageRequest) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageRequest) SetAppId(v string) *ListAppServicesPageRequest {
	s.AppId = &v
	return s
}

func (s *ListAppServicesPageRequest) SetPageNumber(v int32) *ListAppServicesPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppServicesPageRequest) SetPageSize(v int32) *ListAppServicesPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppServicesPageRequest) SetServiceType(v string) *ListAppServicesPageRequest {
	s.ServiceType = &v
	return s
}

type ListAppServicesPageResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListAppServicesPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListAppServicesPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppServicesPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageResponseBody) SetCode(v string) *ListAppServicesPageResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetData(v []*ListAppServicesPageResponseBodyData) *ListAppServicesPageResponseBody {
	s.Data = v
	return s
}

func (s *ListAppServicesPageResponseBody) SetErrorCode(v string) *ListAppServicesPageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetMessage(v string) *ListAppServicesPageResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetRequestId(v string) *ListAppServicesPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetSuccess(v bool) *ListAppServicesPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppServicesPageResponseBody) SetTraceId(v string) *ListAppServicesPageResponseBody {
	s.TraceId = &v
	return s
}

type ListAppServicesPageResponseBodyData struct {
	CurrentPage *string                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageNumber  *string                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result      []*ListAppServicesPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	TotalSize   *string                                      `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListAppServicesPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppServicesPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageResponseBodyData) SetCurrentPage(v string) *ListAppServicesPageResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListAppServicesPageResponseBodyData) SetPageNumber(v string) *ListAppServicesPageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAppServicesPageResponseBodyData) SetPageSize(v string) *ListAppServicesPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAppServicesPageResponseBodyData) SetResult(v []*ListAppServicesPageResponseBodyDataResult) *ListAppServicesPageResponseBodyData {
	s.Result = v
	return s
}

func (s *ListAppServicesPageResponseBodyData) SetTotalSize(v string) *ListAppServicesPageResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListAppServicesPageResponseBodyDataResult struct {
	EdasAppId   *string `json:"EdasAppId,omitempty" xml:"EdasAppId,omitempty"`
	EdasAppName *string `json:"EdasAppName,omitempty" xml:"EdasAppName,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	InstanceNum *int64  `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAppServicesPageResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppServicesPageResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageResponseBodyDataResult) SetEdasAppId(v string) *ListAppServicesPageResponseBodyDataResult {
	s.EdasAppId = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetEdasAppName(v string) *ListAppServicesPageResponseBodyDataResult {
	s.EdasAppName = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetGroup(v string) *ListAppServicesPageResponseBodyDataResult {
	s.Group = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetInstanceNum(v int64) *ListAppServicesPageResponseBodyDataResult {
	s.InstanceNum = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetServiceName(v string) *ListAppServicesPageResponseBodyDataResult {
	s.ServiceName = &v
	return s
}

func (s *ListAppServicesPageResponseBodyDataResult) SetVersion(v string) *ListAppServicesPageResponseBodyDataResult {
	s.Version = &v
	return s
}

type ListAppServicesPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppServicesPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppServicesPageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppServicesPageResponse) GoString() string {
	return s.String()
}

func (s *ListAppServicesPageResponse) SetHeaders(v map[string]*string) *ListAppServicesPageResponse {
	s.Headers = v
	return s
}

func (s *ListAppServicesPageResponse) SetStatusCode(v int32) *ListAppServicesPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppServicesPageResponse) SetBody(v *ListAppServicesPageResponseBody) *ListAppServicesPageResponse {
	s.Body = v
	return s
}

type ListAppVersionsRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListAppVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAppVersionsRequest) SetAppId(v string) *ListAppVersionsRequest {
	s.AppId = &v
	return s
}

type ListAppVersionsResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListAppVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppVersionsResponseBody) SetCode(v string) *ListAppVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppVersionsResponseBody) SetData(v []*ListAppVersionsResponseBodyData) *ListAppVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppVersionsResponseBody) SetErrorCode(v string) *ListAppVersionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAppVersionsResponseBody) SetMessage(v string) *ListAppVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppVersionsResponseBody) SetRequestId(v string) *ListAppVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppVersionsResponseBody) SetSuccess(v bool) *ListAppVersionsResponseBody {
	s.Success = &v
	return s
}

type ListAppVersionsResponseBodyData struct {
	BuildPackageUrl *string `json:"BuildPackageUrl,omitempty" xml:"BuildPackageUrl,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WarUrl          *string `json:"WarUrl,omitempty" xml:"WarUrl,omitempty"`
}

func (s ListAppVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppVersionsResponseBodyData) SetBuildPackageUrl(v string) *ListAppVersionsResponseBodyData {
	s.BuildPackageUrl = &v
	return s
}

func (s *ListAppVersionsResponseBodyData) SetCreateTime(v string) *ListAppVersionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAppVersionsResponseBodyData) SetId(v string) *ListAppVersionsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAppVersionsResponseBodyData) SetType(v string) *ListAppVersionsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListAppVersionsResponseBodyData) SetWarUrl(v string) *ListAppVersionsResponseBodyData {
	s.WarUrl = &v
	return s
}

type ListAppVersionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAppVersionsResponse) SetHeaders(v map[string]*string) *ListAppVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAppVersionsResponse) SetStatusCode(v int32) *ListAppVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppVersionsResponse) SetBody(v *ListAppVersionsResponseBody) *ListAppVersionsResponse {
	s.Body = v
	return s
}

type ListApplicationsRequest struct {
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	FieldType   *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	FieldValue  *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	OrderBy     *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse     *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) SetAppName(v string) *ListApplicationsRequest {
	s.AppName = &v
	return s
}

func (s *ListApplicationsRequest) SetCurrentPage(v int32) *ListApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListApplicationsRequest) SetFieldType(v string) *ListApplicationsRequest {
	s.FieldType = &v
	return s
}

func (s *ListApplicationsRequest) SetFieldValue(v string) *ListApplicationsRequest {
	s.FieldValue = &v
	return s
}

func (s *ListApplicationsRequest) SetNamespaceId(v string) *ListApplicationsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListApplicationsRequest) SetOrderBy(v string) *ListApplicationsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v int32) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) SetReverse(v bool) *ListApplicationsRequest {
	s.Reverse = &v
	return s
}

func (s *ListApplicationsRequest) SetTags(v string) *ListApplicationsRequest {
	s.Tags = &v
	return s
}

type ListApplicationsResponseBody struct {
	Code        *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        *ListApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode   *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message     *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalSize   *int32                            `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) SetCode(v string) *ListApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationsResponseBody) SetCurrentPage(v int32) *ListApplicationsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListApplicationsResponseBody) SetData(v *ListApplicationsResponseBodyData) *ListApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationsResponseBody) SetErrorCode(v string) *ListApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListApplicationsResponseBody) SetMessage(v string) *ListApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationsResponseBody) SetPageSize(v int32) *ListApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetSuccess(v bool) *ListApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalSize(v int32) *ListApplicationsResponseBody {
	s.TotalSize = &v
	return s
}

type ListApplicationsResponseBodyData struct {
	Applications []*ListApplicationsResponseBodyDataApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	CurrentPage  *int32                                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize    *int32                                          `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListApplicationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyData) SetApplications(v []*ListApplicationsResponseBodyDataApplications) *ListApplicationsResponseBodyData {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBodyData) SetCurrentPage(v int32) *ListApplicationsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListApplicationsResponseBodyData) SetPageSize(v int32) *ListApplicationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponseBodyData) SetTotalSize(v int32) *ListApplicationsResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListApplicationsResponseBodyDataApplications struct {
	AppDeletingStatus *bool                                               `json:"AppDeletingStatus,omitempty" xml:"AppDeletingStatus,omitempty"`
	AppDescription    *string                                             `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	AppId             *string                                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName           *string                                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Instances         *int32                                              `json:"Instances,omitempty" xml:"Instances,omitempty"`
	NamespaceId       *string                                             `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	RegionId          *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RunningInstances  *int32                                              `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	Tags              []*ListApplicationsResponseBodyDataApplicationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyDataApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyDataApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyDataApplications) SetAppDeletingStatus(v bool) *ListApplicationsResponseBodyDataApplications {
	s.AppDeletingStatus = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetAppDescription(v string) *ListApplicationsResponseBodyDataApplications {
	s.AppDescription = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetAppId(v string) *ListApplicationsResponseBodyDataApplications {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetAppName(v string) *ListApplicationsResponseBodyDataApplications {
	s.AppName = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetInstances(v int32) *ListApplicationsResponseBodyDataApplications {
	s.Instances = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetNamespaceId(v string) *ListApplicationsResponseBodyDataApplications {
	s.NamespaceId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetRegionId(v string) *ListApplicationsResponseBodyDataApplications {
	s.RegionId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetRunningInstances(v int32) *ListApplicationsResponseBodyDataApplications {
	s.RunningInstances = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetTags(v []*ListApplicationsResponseBodyDataApplicationsTags) *ListApplicationsResponseBodyDataApplications {
	s.Tags = v
	return s
}

type ListApplicationsResponseBodyDataApplicationsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListApplicationsResponseBodyDataApplicationsTags) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyDataApplicationsTags) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyDataApplicationsTags) SetKey(v string) *ListApplicationsResponseBodyDataApplicationsTags {
	s.Key = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsTags) SetValue(v string) *ListApplicationsResponseBodyDataApplicationsTags {
	s.Value = &v
	return s
}

type ListApplicationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) SetHeaders(v map[string]*string) *ListApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsResponse) SetStatusCode(v int32) *ListApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsResponse) SetBody(v *ListApplicationsResponseBody) *ListApplicationsResponse {
	s.Body = v
	return s
}

type ListChangeOrdersRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CoStatus    *string `json:"CoStatus,omitempty" xml:"CoStatus,omitempty"`
	CoType      *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListChangeOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChangeOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersRequest) SetAppId(v string) *ListChangeOrdersRequest {
	s.AppId = &v
	return s
}

func (s *ListChangeOrdersRequest) SetCoStatus(v string) *ListChangeOrdersRequest {
	s.CoStatus = &v
	return s
}

func (s *ListChangeOrdersRequest) SetCoType(v string) *ListChangeOrdersRequest {
	s.CoType = &v
	return s
}

func (s *ListChangeOrdersRequest) SetCurrentPage(v int32) *ListChangeOrdersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListChangeOrdersRequest) SetKey(v string) *ListChangeOrdersRequest {
	s.Key = &v
	return s
}

func (s *ListChangeOrdersRequest) SetPageSize(v int32) *ListChangeOrdersRequest {
	s.PageSize = &v
	return s
}

type ListChangeOrdersResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListChangeOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListChangeOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChangeOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersResponseBody) SetCode(v string) *ListChangeOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetData(v *ListChangeOrdersResponseBodyData) *ListChangeOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ListChangeOrdersResponseBody) SetErrorCode(v string) *ListChangeOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetMessage(v string) *ListChangeOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetRequestId(v string) *ListChangeOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetSuccess(v bool) *ListChangeOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ListChangeOrdersResponseBody) SetTraceId(v string) *ListChangeOrdersResponseBody {
	s.TraceId = &v
	return s
}

type ListChangeOrdersResponseBodyData struct {
	ChangeOrderList []*ListChangeOrdersResponseBodyDataChangeOrderList `json:"ChangeOrderList,omitempty" xml:"ChangeOrderList,omitempty" type:"Repeated"`
	CurrentPage     *int32                                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize       *int32                                             `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListChangeOrdersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListChangeOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersResponseBodyData) SetChangeOrderList(v []*ListChangeOrdersResponseBodyDataChangeOrderList) *ListChangeOrdersResponseBodyData {
	s.ChangeOrderList = v
	return s
}

func (s *ListChangeOrdersResponseBodyData) SetCurrentPage(v int32) *ListChangeOrdersResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListChangeOrdersResponseBodyData) SetPageSize(v int32) *ListChangeOrdersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListChangeOrdersResponseBodyData) SetTotalSize(v int32) *ListChangeOrdersResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListChangeOrdersResponseBodyDataChangeOrderList struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BatchCount    *int32  `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	BatchType     *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	CoType        *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	CoTypeCode    *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId  *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FinishTime    *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListChangeOrdersResponseBodyDataChangeOrderList) String() string {
	return tea.Prettify(s)
}

func (s ListChangeOrdersResponseBodyDataChangeOrderList) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetAppId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.AppId = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetBatchCount(v int32) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.BatchCount = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetBatchType(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.BatchType = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetChangeOrderId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.ChangeOrderId = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetCoType(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.CoType = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetCoTypeCode(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.CoTypeCode = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetCreateTime(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.CreateTime = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetCreateUserId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.CreateUserId = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetDescription(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.Description = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetFinishTime(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.FinishTime = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetGroupId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.GroupId = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetSource(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.Source = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetStatus(v int32) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.Status = &v
	return s
}

func (s *ListChangeOrdersResponseBodyDataChangeOrderList) SetUserId(v string) *ListChangeOrdersResponseBodyDataChangeOrderList {
	s.UserId = &v
	return s
}

type ListChangeOrdersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListChangeOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChangeOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChangeOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListChangeOrdersResponse) SetHeaders(v map[string]*string) *ListChangeOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListChangeOrdersResponse) SetStatusCode(v int32) *ListChangeOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChangeOrdersResponse) SetBody(v *ListChangeOrdersResponseBody) *ListChangeOrdersResponse {
	s.Body = v
	return s
}

type ListConsumedServicesRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListConsumedServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesRequest) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesRequest) SetAppId(v string) *ListConsumedServicesRequest {
	s.AppId = &v
	return s
}

type ListConsumedServicesResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListConsumedServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                 `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListConsumedServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBody) SetCode(v string) *ListConsumedServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetData(v []*ListConsumedServicesResponseBodyData) *ListConsumedServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumedServicesResponseBody) SetErrorCode(v string) *ListConsumedServicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetMessage(v string) *ListConsumedServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetRequestId(v string) *ListConsumedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetSuccess(v bool) *ListConsumedServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetTraceId(v string) *ListConsumedServicesResponseBody {
	s.TraceId = &v
	return s
}

type ListConsumedServicesResponseBodyData struct {
	AppId    *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Group2Ip *string   `json:"Group2Ip,omitempty" xml:"Group2Ip,omitempty"`
	Groups   []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	Ips      []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	Name     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Type     *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Version  *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListConsumedServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyData) SetAppId(v string) *ListConsumedServicesResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetGroup2Ip(v string) *ListConsumedServicesResponseBodyData {
	s.Group2Ip = &v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetGroups(v []*string) *ListConsumedServicesResponseBodyData {
	s.Groups = v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetIps(v []*string) *ListConsumedServicesResponseBodyData {
	s.Ips = v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetName(v string) *ListConsumedServicesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetType(v string) *ListConsumedServicesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetVersion(v string) *ListConsumedServicesResponseBodyData {
	s.Version = &v
	return s
}

type ListConsumedServicesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConsumedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConsumedServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesResponse) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponse) SetHeaders(v map[string]*string) *ListConsumedServicesResponse {
	s.Headers = v
	return s
}

func (s *ListConsumedServicesResponse) SetStatusCode(v int32) *ListConsumedServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumedServicesResponse) SetBody(v *ListConsumedServicesResponseBody) *ListConsumedServicesResponse {
	s.Body = v
	return s
}

type ListGreyTagRouteRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListGreyTagRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteRequest) SetAppId(v string) *ListGreyTagRouteRequest {
	s.AppId = &v
	return s
}

type ListGreyTagRouteResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListGreyTagRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBody) SetCode(v string) *ListGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetData(v *ListGreyTagRouteResponseBodyData) *ListGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetErrorCode(v string) *ListGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetMessage(v string) *ListGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetRequestId(v string) *ListGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetSuccess(v bool) *ListGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *ListGreyTagRouteResponseBody) SetTraceId(v string) *ListGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

type ListGreyTagRouteResponseBodyData struct {
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result      []*ListGreyTagRouteResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	TotalSize   *int64                                    `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGreyTagRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyData) SetCurrentPage(v int32) *ListGreyTagRouteResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyData) SetPageSize(v int32) *ListGreyTagRouteResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyData) SetResult(v []*ListGreyTagRouteResponseBodyDataResult) *ListGreyTagRouteResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGreyTagRouteResponseBodyData) SetTotalSize(v int64) *ListGreyTagRouteResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListGreyTagRouteResponseBodyDataResult struct {
	CreateTime     *int64                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DubboRules     []*ListGreyTagRouteResponseBodyDataResultDubboRules `json:"DubboRules,omitempty" xml:"DubboRules,omitempty" type:"Repeated"`
	GreyTagRouteId *int64                                              `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
	Name           *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	ScRules        []*ListGreyTagRouteResponseBodyDataResultScRules    `json:"ScRules,omitempty" xml:"ScRules,omitempty" type:"Repeated"`
	UpdateTime     *int64                                              `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetCreateTime(v int64) *ListGreyTagRouteResponseBodyDataResult {
	s.CreateTime = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetDescription(v string) *ListGreyTagRouteResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetDubboRules(v []*ListGreyTagRouteResponseBodyDataResultDubboRules) *ListGreyTagRouteResponseBodyDataResult {
	s.DubboRules = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetGreyTagRouteId(v int64) *ListGreyTagRouteResponseBodyDataResult {
	s.GreyTagRouteId = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetName(v string) *ListGreyTagRouteResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetScRules(v []*ListGreyTagRouteResponseBodyDataResultScRules) *ListGreyTagRouteResponseBodyDataResult {
	s.ScRules = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResult) SetUpdateTime(v int64) *ListGreyTagRouteResponseBodyDataResult {
	s.UpdateTime = &v
	return s
}

type ListGreyTagRouteResponseBodyDataResultDubboRules struct {
	Condition   *string                                                  `json:"condition,omitempty" xml:"condition,omitempty"`
	Group       *string                                                  `json:"group,omitempty" xml:"group,omitempty"`
	Items       []*ListGreyTagRouteResponseBodyDataResultDubboRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	MethodName  *string                                                  `json:"methodName,omitempty" xml:"methodName,omitempty"`
	ServiceName *string                                                  `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	Version     *string                                                  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultDubboRules) String() string {
	return tea.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultDubboRules) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetCondition(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.Condition = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetGroup(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.Group = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetItems(v []*ListGreyTagRouteResponseBodyDataResultDubboRulesItems) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.Items = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetMethodName(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.MethodName = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetServiceName(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.ServiceName = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRules) SetVersion(v string) *ListGreyTagRouteResponseBodyDataResultDubboRules {
	s.Version = &v
	return s
}

type ListGreyTagRouteResponseBodyDataResultDubboRulesItems struct {
	Cond     *string `json:"cond,omitempty" xml:"cond,omitempty"`
	Expr     *string `json:"expr,omitempty" xml:"expr,omitempty"`
	Index    *int32  `json:"index,omitempty" xml:"index,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultDubboRulesItems) String() string {
	return tea.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultDubboRulesItems) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetCond(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Cond = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetExpr(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Expr = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetIndex(v int32) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Index = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetName(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Name = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetOperator(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Operator = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetType(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Type = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultDubboRulesItems) SetValue(v string) *ListGreyTagRouteResponseBodyDataResultDubboRulesItems {
	s.Value = &v
	return s
}

type ListGreyTagRouteResponseBodyDataResultScRules struct {
	Condition *string                                               `json:"condition,omitempty" xml:"condition,omitempty"`
	Items     []*ListGreyTagRouteResponseBodyDataResultScRulesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Path      *string                                               `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultScRules) String() string {
	return tea.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultScRules) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) SetCondition(v string) *ListGreyTagRouteResponseBodyDataResultScRules {
	s.Condition = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) SetItems(v []*ListGreyTagRouteResponseBodyDataResultScRulesItems) *ListGreyTagRouteResponseBodyDataResultScRules {
	s.Items = v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRules) SetPath(v string) *ListGreyTagRouteResponseBodyDataResultScRules {
	s.Path = &v
	return s
}

type ListGreyTagRouteResponseBodyDataResultScRulesItems struct {
	Cond     *string `json:"cond,omitempty" xml:"cond,omitempty"`
	Expr     *string `json:"expr,omitempty" xml:"expr,omitempty"`
	Index    *int32  `json:"index,omitempty" xml:"index,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGreyTagRouteResponseBodyDataResultScRulesItems) String() string {
	return tea.Prettify(s)
}

func (s ListGreyTagRouteResponseBodyDataResultScRulesItems) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetCond(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Cond = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetExpr(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Expr = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetIndex(v int32) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Index = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetName(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Name = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetOperator(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Operator = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetType(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Type = &v
	return s
}

func (s *ListGreyTagRouteResponseBodyDataResultScRulesItems) SetValue(v string) *ListGreyTagRouteResponseBodyDataResultScRulesItems {
	s.Value = &v
	return s
}

type ListGreyTagRouteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGreyTagRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *ListGreyTagRouteResponse) SetHeaders(v map[string]*string) *ListGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *ListGreyTagRouteResponse) SetStatusCode(v int32) *ListGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGreyTagRouteResponse) SetBody(v *ListGreyTagRouteResponseBody) *ListGreyTagRouteResponse {
	s.Body = v
	return s
}

type ListIngressesRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListIngressesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIngressesRequest) GoString() string {
	return s.String()
}

func (s *ListIngressesRequest) SetAppId(v string) *ListIngressesRequest {
	s.AppId = &v
	return s
}

func (s *ListIngressesRequest) SetNamespaceId(v string) *ListIngressesRequest {
	s.NamespaceId = &v
	return s
}

type ListIngressesResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListIngressesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                        `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListIngressesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIngressesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBody) SetCode(v string) *ListIngressesResponseBody {
	s.Code = &v
	return s
}

func (s *ListIngressesResponseBody) SetData(v *ListIngressesResponseBodyData) *ListIngressesResponseBody {
	s.Data = v
	return s
}

func (s *ListIngressesResponseBody) SetErrorCode(v string) *ListIngressesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListIngressesResponseBody) SetMessage(v string) *ListIngressesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIngressesResponseBody) SetRequestId(v string) *ListIngressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIngressesResponseBody) SetSuccess(v bool) *ListIngressesResponseBody {
	s.Success = &v
	return s
}

func (s *ListIngressesResponseBody) SetTraceId(v string) *ListIngressesResponseBody {
	s.TraceId = &v
	return s
}

type ListIngressesResponseBodyData struct {
	IngressList []*ListIngressesResponseBodyDataIngressList `json:"IngressList,omitempty" xml:"IngressList,omitempty" type:"Repeated"`
}

func (s ListIngressesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIngressesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBodyData) SetIngressList(v []*ListIngressesResponseBodyDataIngressList) *ListIngressesResponseBodyData {
	s.IngressList = v
	return s
}

type ListIngressesResponseBodyDataIngressList struct {
	CertId           *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ListenerPort     *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalanceType  *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId      *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	SlbId            *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbType          *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s ListIngressesResponseBodyDataIngressList) String() string {
	return tea.Prettify(s)
}

func (s ListIngressesResponseBodyDataIngressList) GoString() string {
	return s.String()
}

func (s *ListIngressesResponseBodyDataIngressList) SetCertId(v string) *ListIngressesResponseBodyDataIngressList {
	s.CertId = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetDescription(v string) *ListIngressesResponseBodyDataIngressList {
	s.Description = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetId(v int64) *ListIngressesResponseBodyDataIngressList {
	s.Id = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetListenerPort(v string) *ListIngressesResponseBodyDataIngressList {
	s.ListenerPort = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetListenerProtocol(v string) *ListIngressesResponseBodyDataIngressList {
	s.ListenerProtocol = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetLoadBalanceType(v string) *ListIngressesResponseBodyDataIngressList {
	s.LoadBalanceType = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetName(v string) *ListIngressesResponseBodyDataIngressList {
	s.Name = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetNamespaceId(v string) *ListIngressesResponseBodyDataIngressList {
	s.NamespaceId = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetSlbId(v string) *ListIngressesResponseBodyDataIngressList {
	s.SlbId = &v
	return s
}

func (s *ListIngressesResponseBodyDataIngressList) SetSlbType(v string) *ListIngressesResponseBodyDataIngressList {
	s.SlbType = &v
	return s
}

type ListIngressesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIngressesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIngressesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIngressesResponse) GoString() string {
	return s.String()
}

func (s *ListIngressesResponse) SetHeaders(v map[string]*string) *ListIngressesResponse {
	s.Headers = v
	return s
}

func (s *ListIngressesResponse) SetStatusCode(v int32) *ListIngressesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIngressesResponse) SetBody(v *ListIngressesResponseBody) *ListIngressesResponse {
	s.Body = v
	return s
}

type ListLogConfigsRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLogConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListLogConfigsRequest) SetAppId(v string) *ListLogConfigsRequest {
	s.AppId = &v
	return s
}

func (s *ListLogConfigsRequest) SetCurrentPage(v int32) *ListLogConfigsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListLogConfigsRequest) SetPageSize(v int32) *ListLogConfigsRequest {
	s.PageSize = &v
	return s
}

type ListLogConfigsResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListLogConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                         `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListLogConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogConfigsResponseBody) SetCode(v string) *ListLogConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetData(v *ListLogConfigsResponseBodyData) *ListLogConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListLogConfigsResponseBody) SetErrorCode(v string) *ListLogConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetMessage(v string) *ListLogConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetRequestId(v string) *ListLogConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetSuccess(v bool) *ListLogConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListLogConfigsResponseBody) SetTraceId(v string) *ListLogConfigsResponseBody {
	s.TraceId = &v
	return s
}

type ListLogConfigsResponseBodyData struct {
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	LogConfigs  []*ListLogConfigsResponseBodyDataLogConfigs `json:"LogConfigs,omitempty" xml:"LogConfigs,omitempty" type:"Repeated"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize   *int32                                      `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListLogConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLogConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLogConfigsResponseBodyData) SetCurrentPage(v int32) *ListLogConfigsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListLogConfigsResponseBodyData) SetLogConfigs(v []*ListLogConfigsResponseBodyDataLogConfigs) *ListLogConfigsResponseBodyData {
	s.LogConfigs = v
	return s
}

func (s *ListLogConfigsResponseBodyData) SetPageSize(v int32) *ListLogConfigsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListLogConfigsResponseBodyData) SetTotalSize(v int32) *ListLogConfigsResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListLogConfigsResponseBodyDataLogConfigs struct {
	ConfigName  *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LogDir      *string `json:"LogDir,omitempty" xml:"LogDir,omitempty"`
	LogType     *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SlsLogStore *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
	SlsProject  *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	StoreType   *string `json:"StoreType,omitempty" xml:"StoreType,omitempty"`
}

func (s ListLogConfigsResponseBodyDataLogConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListLogConfigsResponseBodyDataLogConfigs) GoString() string {
	return s.String()
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetConfigName(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.ConfigName = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetCreateTime(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.CreateTime = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetLogDir(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.LogDir = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetLogType(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.LogType = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetRegionId(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.RegionId = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetSlsLogStore(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.SlsLogStore = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetSlsProject(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.SlsProject = &v
	return s
}

func (s *ListLogConfigsResponseBodyDataLogConfigs) SetStoreType(v string) *ListLogConfigsResponseBodyDataLogConfigs {
	s.StoreType = &v
	return s
}

type ListLogConfigsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLogConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListLogConfigsResponse) SetHeaders(v map[string]*string) *ListLogConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListLogConfigsResponse) SetStatusCode(v int32) *ListLogConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogConfigsResponse) SetBody(v *ListLogConfigsResponseBody) *ListLogConfigsResponse {
	s.Body = v
	return s
}

type ListNamespaceChangeOrdersRequest struct {
	CoStatus    *string `json:"CoStatus,omitempty" xml:"CoStatus,omitempty"`
	CoType      *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNamespaceChangeOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNamespaceChangeOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersRequest) SetCoStatus(v string) *ListNamespaceChangeOrdersRequest {
	s.CoStatus = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetCoType(v string) *ListNamespaceChangeOrdersRequest {
	s.CoType = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetCurrentPage(v int32) *ListNamespaceChangeOrdersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetKey(v string) *ListNamespaceChangeOrdersRequest {
	s.Key = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetNamespaceId(v string) *ListNamespaceChangeOrdersRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespaceChangeOrdersRequest) SetPageSize(v int32) *ListNamespaceChangeOrdersRequest {
	s.PageSize = &v
	return s
}

type ListNamespaceChangeOrdersResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListNamespaceChangeOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                    `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListNamespaceChangeOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNamespaceChangeOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersResponseBody) SetCode(v string) *ListNamespaceChangeOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetData(v *ListNamespaceChangeOrdersResponseBodyData) *ListNamespaceChangeOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetErrorCode(v string) *ListNamespaceChangeOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetMessage(v string) *ListNamespaceChangeOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetRequestId(v string) *ListNamespaceChangeOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetSuccess(v bool) *ListNamespaceChangeOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBody) SetTraceId(v string) *ListNamespaceChangeOrdersResponseBody {
	s.TraceId = &v
	return s
}

type ListNamespaceChangeOrdersResponseBodyData struct {
	ChangeOrderList []*ListNamespaceChangeOrdersResponseBodyDataChangeOrderList `json:"ChangeOrderList,omitempty" xml:"ChangeOrderList,omitempty" type:"Repeated"`
	CurrentPage     *int32                                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize       *int32                                                      `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListNamespaceChangeOrdersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNamespaceChangeOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersResponseBodyData) SetChangeOrderList(v []*ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) *ListNamespaceChangeOrdersResponseBodyData {
	s.ChangeOrderList = v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyData) SetCurrentPage(v int32) *ListNamespaceChangeOrdersResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyData) SetPageSize(v int32) *ListNamespaceChangeOrdersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyData) SetTotalSize(v int32) *ListNamespaceChangeOrdersResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListNamespaceChangeOrdersResponseBodyDataChangeOrderList struct {
	BatchCount    *int32  `json:"BatchCount,omitempty" xml:"BatchCount,omitempty"`
	BatchType     *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	CoType        *string `json:"CoType,omitempty" xml:"CoType,omitempty"`
	CoTypeCode    *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId  *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FinishTime    *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NamespaceId   *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Pipelines     *string `json:"Pipelines,omitempty" xml:"Pipelines,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) String() string {
	return tea.Prettify(s)
}

func (s ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetBatchCount(v int32) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.BatchCount = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetBatchType(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.BatchType = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetChangeOrderId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.ChangeOrderId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetCoType(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.CoType = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetCoTypeCode(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.CoTypeCode = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetCreateTime(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.CreateTime = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetCreateUserId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.CreateUserId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetDescription(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.Description = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetFinishTime(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.FinishTime = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetGroupId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.GroupId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetNamespaceId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetPipelines(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.Pipelines = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetSource(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.Source = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetStatus(v int32) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.Status = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList) SetUserId(v string) *ListNamespaceChangeOrdersResponseBodyDataChangeOrderList {
	s.UserId = &v
	return s
}

type ListNamespaceChangeOrdersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNamespaceChangeOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNamespaceChangeOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNamespaceChangeOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListNamespaceChangeOrdersResponse) SetHeaders(v map[string]*string) *ListNamespaceChangeOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListNamespaceChangeOrdersResponse) SetStatusCode(v int32) *ListNamespaceChangeOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespaceChangeOrdersResponse) SetBody(v *ListNamespaceChangeOrdersResponseBody) *ListNamespaceChangeOrdersResponse {
	s.Body = v
	return s
}

type ListNamespacedConfigMapsRequest struct {
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListNamespacedConfigMapsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacedConfigMapsRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsRequest) SetNamespaceId(v string) *ListNamespacedConfigMapsRequest {
	s.NamespaceId = &v
	return s
}

type ListNamespacedConfigMapsResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListNamespacedConfigMapsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                   `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListNamespacedConfigMapsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacedConfigMapsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponseBody) SetCode(v string) *ListNamespacedConfigMapsResponseBody {
	s.Code = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetData(v *ListNamespacedConfigMapsResponseBodyData) *ListNamespacedConfigMapsResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetErrorCode(v string) *ListNamespacedConfigMapsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetMessage(v string) *ListNamespacedConfigMapsResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetRequestId(v string) *ListNamespacedConfigMapsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetSuccess(v bool) *ListNamespacedConfigMapsResponseBody {
	s.Success = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBody) SetTraceId(v string) *ListNamespacedConfigMapsResponseBody {
	s.TraceId = &v
	return s
}

type ListNamespacedConfigMapsResponseBodyData struct {
	ConfigMaps []*ListNamespacedConfigMapsResponseBodyDataConfigMaps `json:"ConfigMaps,omitempty" xml:"ConfigMaps,omitempty" type:"Repeated"`
}

func (s ListNamespacedConfigMapsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacedConfigMapsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponseBodyData) SetConfigMaps(v []*ListNamespacedConfigMapsResponseBodyDataConfigMaps) *ListNamespacedConfigMapsResponseBodyData {
	s.ConfigMaps = v
	return s
}

type ListNamespacedConfigMapsResponseBodyDataConfigMaps struct {
	ConfigMapId *int64                                                          `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	CreateTime  *int64                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Data        map[string]interface{}                                          `json:"Data,omitempty" xml:"Data,omitempty"`
	Description *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId *string                                                         `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	RelateApps  []*ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps `json:"RelateApps,omitempty" xml:"RelateApps,omitempty" type:"Repeated"`
	UpdateTime  *int64                                                          `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListNamespacedConfigMapsResponseBodyDataConfigMaps) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacedConfigMapsResponseBodyDataConfigMaps) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetConfigMapId(v int64) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.ConfigMapId = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetCreateTime(v int64) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.CreateTime = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetData(v map[string]interface{}) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.Data = v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetDescription(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.Description = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetName(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.Name = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetNamespaceId(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetRelateApps(v []*ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.RelateApps = v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMaps) SetUpdateTime(v int64) *ListNamespacedConfigMapsResponseBodyDataConfigMaps {
	s.UpdateTime = &v
	return s
}

type ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) SetAppId(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps {
	s.AppId = &v
	return s
}

func (s *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps) SetAppName(v string) *ListNamespacedConfigMapsResponseBodyDataConfigMapsRelateApps {
	s.AppName = &v
	return s
}

type ListNamespacedConfigMapsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNamespacedConfigMapsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNamespacedConfigMapsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacedConfigMapsResponse) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsResponse) SetHeaders(v map[string]*string) *ListNamespacedConfigMapsResponse {
	s.Headers = v
	return s
}

func (s *ListNamespacedConfigMapsResponse) SetStatusCode(v int32) *ListNamespacedConfigMapsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespacedConfigMapsResponse) SetBody(v *ListNamespacedConfigMapsResponseBody) *ListNamespacedConfigMapsResponse {
	s.Body = v
	return s
}

type ListPublishedServicesRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListPublishedServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesRequest) SetAppId(v string) *ListPublishedServicesRequest {
	s.AppId = &v
	return s
}

type ListPublishedServicesResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListPublishedServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListPublishedServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBody) SetCode(v string) *ListPublishedServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetData(v []*ListPublishedServicesResponseBodyData) *ListPublishedServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListPublishedServicesResponseBody) SetErrorCode(v string) *ListPublishedServicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetMessage(v string) *ListPublishedServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetRequestId(v string) *ListPublishedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetSuccess(v bool) *ListPublishedServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetTraceId(v string) *ListPublishedServicesResponseBody {
	s.TraceId = &v
	return s
}

type ListPublishedServicesResponseBodyData struct {
	AppId    *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Group2Ip *string   `json:"Group2Ip,omitempty" xml:"Group2Ip,omitempty"`
	Groups   []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	Ips      []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	Name     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Type     *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Version  *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPublishedServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyData) SetAppId(v string) *ListPublishedServicesResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetGroup2Ip(v string) *ListPublishedServicesResponseBodyData {
	s.Group2Ip = &v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetGroups(v []*string) *ListPublishedServicesResponseBodyData {
	s.Groups = v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetIps(v []*string) *ListPublishedServicesResponseBodyData {
	s.Ips = v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetName(v string) *ListPublishedServicesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetType(v string) *ListPublishedServicesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetVersion(v string) *ListPublishedServicesResponseBodyData {
	s.Version = &v
	return s
}

type ListPublishedServicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPublishedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPublishedServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponse) SetHeaders(v map[string]*string) *ListPublishedServicesResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedServicesResponse) SetStatusCode(v int32) *ListPublishedServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedServicesResponse) SetBody(v *ListPublishedServicesResponseBody) *ListPublishedServicesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIds  *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceIds(v string) *ListTagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v string) *ListTagResourcesRequest {
	s.Tags = &v
	return s
}

type ListTagResourcesResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListTagResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetCode(v string) *ListTagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetData(v *ListTagResourcesResponseBodyData) *ListTagResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListTagResourcesResponseBody) SetErrorCode(v string) *ListTagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetMessage(v string) *ListTagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTraceId(v string) *ListTagResourcesResponseBody {
	s.TraceId = &v
	return s
}

type ListTagResourcesResponseBodyData struct {
	NextToken    *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TagResources []*ListTagResourcesResponseBodyDataTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyData) SetNextToken(v string) *ListTagResourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetTagResources(v []*ListTagResourcesResponseBodyDataTagResources) *ListTagResourcesResponseBodyData {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyDataTagResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyDataTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyDataTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type OpenSaeServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenSaeServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenSaeServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenSaeServiceResponseBody) SetOrderId(v string) *OpenSaeServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenSaeServiceResponseBody) SetRequestId(v string) *OpenSaeServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenSaeServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenSaeServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenSaeServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenSaeServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenSaeServiceResponse) SetHeaders(v map[string]*string) *OpenSaeServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenSaeServiceResponse) SetStatusCode(v int32) *OpenSaeServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenSaeServiceResponse) SetBody(v *OpenSaeServiceResponseBody) *OpenSaeServiceResponse {
	s.Body = v
	return s
}

type QueryResourceStaticsRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s QueryResourceStaticsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceStaticsRequest) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsRequest) SetAppId(v string) *QueryResourceStaticsRequest {
	s.AppId = &v
	return s
}

type QueryResourceStaticsResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryResourceStaticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                               `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s QueryResourceStaticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceStaticsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponseBody) SetCode(v string) *QueryResourceStaticsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetData(v *QueryResourceStaticsResponseBodyData) *QueryResourceStaticsResponseBody {
	s.Data = v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetErrorCode(v string) *QueryResourceStaticsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetMessage(v string) *QueryResourceStaticsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetRequestId(v string) *QueryResourceStaticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetSuccess(v bool) *QueryResourceStaticsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryResourceStaticsResponseBody) SetTraceId(v string) *QueryResourceStaticsResponseBody {
	s.TraceId = &v
	return s
}

type QueryResourceStaticsResponseBodyData struct {
	RealTimeRes *QueryResourceStaticsResponseBodyDataRealTimeRes `json:"RealTimeRes,omitempty" xml:"RealTimeRes,omitempty" type:"Struct"`
	Summary     *QueryResourceStaticsResponseBodyDataSummary     `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
}

func (s QueryResourceStaticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceStaticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponseBodyData) SetRealTimeRes(v *QueryResourceStaticsResponseBodyDataRealTimeRes) *QueryResourceStaticsResponseBodyData {
	s.RealTimeRes = v
	return s
}

func (s *QueryResourceStaticsResponseBodyData) SetSummary(v *QueryResourceStaticsResponseBodyDataSummary) *QueryResourceStaticsResponseBodyData {
	s.Summary = v
	return s
}

type QueryResourceStaticsResponseBodyDataRealTimeRes struct {
	Cpu    *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s QueryResourceStaticsResponseBodyDataRealTimeRes) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceStaticsResponseBodyDataRealTimeRes) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponseBodyDataRealTimeRes) SetCpu(v float32) *QueryResourceStaticsResponseBodyDataRealTimeRes {
	s.Cpu = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataRealTimeRes) SetMemory(v float32) *QueryResourceStaticsResponseBodyDataRealTimeRes {
	s.Memory = &v
	return s
}

type QueryResourceStaticsResponseBodyDataSummary struct {
	Cpu    *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s QueryResourceStaticsResponseBodyDataSummary) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceStaticsResponseBodyDataSummary) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetCpu(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.Cpu = &v
	return s
}

func (s *QueryResourceStaticsResponseBodyDataSummary) SetMemory(v float32) *QueryResourceStaticsResponseBodyDataSummary {
	s.Memory = &v
	return s
}

type QueryResourceStaticsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryResourceStaticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryResourceStaticsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceStaticsResponse) GoString() string {
	return s.String()
}

func (s *QueryResourceStaticsResponse) SetHeaders(v map[string]*string) *QueryResourceStaticsResponse {
	s.Headers = v
	return s
}

func (s *QueryResourceStaticsResponse) SetStatusCode(v int32) *QueryResourceStaticsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResourceStaticsResponse) SetBody(v *QueryResourceStaticsResponseBody) *QueryResourceStaticsResponse {
	s.Body = v
	return s
}

type ReduceApplicationCapacityByInstanceIdsRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s ReduceApplicationCapacityByInstanceIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ReduceApplicationCapacityByInstanceIdsRequest) GoString() string {
	return s.String()
}

func (s *ReduceApplicationCapacityByInstanceIdsRequest) SetAppId(v string) *ReduceApplicationCapacityByInstanceIdsRequest {
	s.AppId = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsRequest) SetInstanceIds(v string) *ReduceApplicationCapacityByInstanceIdsRequest {
	s.InstanceIds = &v
	return s
}

type ReduceApplicationCapacityByInstanceIdsResponseBody struct {
	Code      *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReduceApplicationCapacityByInstanceIdsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                                 `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ReduceApplicationCapacityByInstanceIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReduceApplicationCapacityByInstanceIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetCode(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.Code = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetData(v *ReduceApplicationCapacityByInstanceIdsResponseBodyData) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.Data = v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetErrorCode(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetMessage(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.Message = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetRequestId(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetSuccess(v bool) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.Success = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetTraceId(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.TraceId = &v
	return s
}

type ReduceApplicationCapacityByInstanceIdsResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s ReduceApplicationCapacityByInstanceIdsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReduceApplicationCapacityByInstanceIdsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBodyData) SetChangeOrderId(v string) *ReduceApplicationCapacityByInstanceIdsResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type ReduceApplicationCapacityByInstanceIdsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReduceApplicationCapacityByInstanceIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReduceApplicationCapacityByInstanceIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ReduceApplicationCapacityByInstanceIdsResponse) GoString() string {
	return s.String()
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) SetHeaders(v map[string]*string) *ReduceApplicationCapacityByInstanceIdsResponse {
	s.Headers = v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) SetStatusCode(v int32) *ReduceApplicationCapacityByInstanceIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponse) SetBody(v *ReduceApplicationCapacityByInstanceIdsResponseBody) *ReduceApplicationCapacityByInstanceIdsResponse {
	s.Body = v
	return s
}

type RescaleApplicationRequest struct {
	AppId                            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AutoEnableApplicationScalingRule *bool   `json:"AutoEnableApplicationScalingRule,omitempty" xml:"AutoEnableApplicationScalingRule,omitempty"`
	MinReadyInstanceRatio            *int32  `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	MinReadyInstances                *int32  `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	Replicas                         *int32  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s RescaleApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationRequest) GoString() string {
	return s.String()
}

func (s *RescaleApplicationRequest) SetAppId(v string) *RescaleApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RescaleApplicationRequest) SetAutoEnableApplicationScalingRule(v bool) *RescaleApplicationRequest {
	s.AutoEnableApplicationScalingRule = &v
	return s
}

func (s *RescaleApplicationRequest) SetMinReadyInstanceRatio(v int32) *RescaleApplicationRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *RescaleApplicationRequest) SetMinReadyInstances(v int32) *RescaleApplicationRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *RescaleApplicationRequest) SetReplicas(v int32) *RescaleApplicationRequest {
	s.Replicas = &v
	return s
}

type RescaleApplicationResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RescaleApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RescaleApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponseBody) SetCode(v string) *RescaleApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RescaleApplicationResponseBody) SetData(v *RescaleApplicationResponseBodyData) *RescaleApplicationResponseBody {
	s.Data = v
	return s
}

func (s *RescaleApplicationResponseBody) SetErrorCode(v string) *RescaleApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RescaleApplicationResponseBody) SetMessage(v string) *RescaleApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RescaleApplicationResponseBody) SetRequestId(v string) *RescaleApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RescaleApplicationResponseBody) SetSuccess(v bool) *RescaleApplicationResponseBody {
	s.Success = &v
	return s
}

type RescaleApplicationResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RescaleApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponseBodyData) SetChangeOrderId(v string) *RescaleApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type RescaleApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RescaleApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RescaleApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationResponse) GoString() string {
	return s.String()
}

func (s *RescaleApplicationResponse) SetHeaders(v map[string]*string) *RescaleApplicationResponse {
	s.Headers = v
	return s
}

func (s *RescaleApplicationResponse) SetStatusCode(v int32) *RescaleApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *RescaleApplicationResponse) SetBody(v *RescaleApplicationResponseBody) *RescaleApplicationResponse {
	s.Body = v
	return s
}

type RescaleApplicationVerticallyRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Cpu    *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s RescaleApplicationVerticallyRequest) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationVerticallyRequest) GoString() string {
	return s.String()
}

func (s *RescaleApplicationVerticallyRequest) SetAppId(v string) *RescaleApplicationVerticallyRequest {
	s.AppId = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetCpu(v string) *RescaleApplicationVerticallyRequest {
	s.Cpu = &v
	return s
}

func (s *RescaleApplicationVerticallyRequest) SetMemory(v string) *RescaleApplicationVerticallyRequest {
	s.Memory = &v
	return s
}

type RescaleApplicationVerticallyResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RescaleApplicationVerticallyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                       `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RescaleApplicationVerticallyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationVerticallyResponseBody) GoString() string {
	return s.String()
}

func (s *RescaleApplicationVerticallyResponseBody) SetCode(v string) *RescaleApplicationVerticallyResponseBody {
	s.Code = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetData(v *RescaleApplicationVerticallyResponseBodyData) *RescaleApplicationVerticallyResponseBody {
	s.Data = v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetErrorCode(v string) *RescaleApplicationVerticallyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetMessage(v string) *RescaleApplicationVerticallyResponseBody {
	s.Message = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetRequestId(v string) *RescaleApplicationVerticallyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetSuccess(v bool) *RescaleApplicationVerticallyResponseBody {
	s.Success = &v
	return s
}

func (s *RescaleApplicationVerticallyResponseBody) SetTraceId(v string) *RescaleApplicationVerticallyResponseBody {
	s.TraceId = &v
	return s
}

type RescaleApplicationVerticallyResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RescaleApplicationVerticallyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationVerticallyResponseBodyData) GoString() string {
	return s.String()
}

func (s *RescaleApplicationVerticallyResponseBodyData) SetChangeOrderId(v string) *RescaleApplicationVerticallyResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type RescaleApplicationVerticallyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RescaleApplicationVerticallyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RescaleApplicationVerticallyResponse) String() string {
	return tea.Prettify(s)
}

func (s RescaleApplicationVerticallyResponse) GoString() string {
	return s.String()
}

func (s *RescaleApplicationVerticallyResponse) SetHeaders(v map[string]*string) *RescaleApplicationVerticallyResponse {
	s.Headers = v
	return s
}

func (s *RescaleApplicationVerticallyResponse) SetStatusCode(v int32) *RescaleApplicationVerticallyResponse {
	s.StatusCode = &v
	return s
}

func (s *RescaleApplicationVerticallyResponse) SetBody(v *RescaleApplicationVerticallyResponseBody) *RescaleApplicationVerticallyResponse {
	s.Body = v
	return s
}

type RestartApplicationRequest struct {
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	MinReadyInstanceRatio *int32  `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	MinReadyInstances     *int32  `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
}

func (s RestartApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartApplicationRequest) GoString() string {
	return s.String()
}

func (s *RestartApplicationRequest) SetAppId(v string) *RestartApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RestartApplicationRequest) SetMinReadyInstanceRatio(v int32) *RestartApplicationRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *RestartApplicationRequest) SetMinReadyInstances(v int32) *RestartApplicationRequest {
	s.MinReadyInstances = &v
	return s
}

type RestartApplicationResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RestartApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                             `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RestartApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RestartApplicationResponseBody) SetCode(v string) *RestartApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RestartApplicationResponseBody) SetData(v *RestartApplicationResponseBodyData) *RestartApplicationResponseBody {
	s.Data = v
	return s
}

func (s *RestartApplicationResponseBody) SetErrorCode(v string) *RestartApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartApplicationResponseBody) SetMessage(v string) *RestartApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RestartApplicationResponseBody) SetRequestId(v string) *RestartApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartApplicationResponseBody) SetSuccess(v bool) *RestartApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *RestartApplicationResponseBody) SetTraceId(v string) *RestartApplicationResponseBody {
	s.TraceId = &v
	return s
}

type RestartApplicationResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RestartApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RestartApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartApplicationResponseBodyData) SetChangeOrderId(v string) *RestartApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type RestartApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartApplicationResponse) GoString() string {
	return s.String()
}

func (s *RestartApplicationResponse) SetHeaders(v map[string]*string) *RestartApplicationResponse {
	s.Headers = v
	return s
}

func (s *RestartApplicationResponse) SetStatusCode(v int32) *RestartApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartApplicationResponse) SetBody(v *RestartApplicationResponseBody) *RestartApplicationResponse {
	s.Body = v
	return s
}

type RestartInstancesRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s RestartInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartInstancesRequest) GoString() string {
	return s.String()
}

func (s *RestartInstancesRequest) SetAppId(v string) *RestartInstancesRequest {
	s.AppId = &v
	return s
}

func (s *RestartInstancesRequest) SetInstanceIds(v string) *RestartInstancesRequest {
	s.InstanceIds = &v
	return s
}

type RestartInstancesResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RestartInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RestartInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstancesResponseBody) SetCode(v string) *RestartInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *RestartInstancesResponseBody) SetData(v *RestartInstancesResponseBodyData) *RestartInstancesResponseBody {
	s.Data = v
	return s
}

func (s *RestartInstancesResponseBody) SetErrorCode(v string) *RestartInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartInstancesResponseBody) SetMessage(v string) *RestartInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *RestartInstancesResponseBody) SetRequestId(v string) *RestartInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstancesResponseBody) SetSuccess(v bool) *RestartInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *RestartInstancesResponseBody) SetTraceId(v string) *RestartInstancesResponseBody {
	s.TraceId = &v
	return s
}

type RestartInstancesResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RestartInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RestartInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartInstancesResponseBodyData) SetChangeOrderId(v string) *RestartInstancesResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type RestartInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstancesResponse) GoString() string {
	return s.String()
}

func (s *RestartInstancesResponse) SetHeaders(v map[string]*string) *RestartInstancesResponse {
	s.Headers = v
	return s
}

func (s *RestartInstancesResponse) SetStatusCode(v int32) *RestartInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartInstancesResponse) SetBody(v *RestartInstancesResponseBody) *RestartInstancesResponse {
	s.Body = v
	return s
}

type RollbackApplicationRequest struct {
	AppId                            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AutoEnableApplicationScalingRule *string `json:"AutoEnableApplicationScalingRule,omitempty" xml:"AutoEnableApplicationScalingRule,omitempty"`
	BatchWaitTime                    *int32  `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	MinReadyInstanceRatio            *int32  `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	MinReadyInstances                *int32  `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	UpdateStrategy                   *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	VersionId                        *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s RollbackApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationRequest) GoString() string {
	return s.String()
}

func (s *RollbackApplicationRequest) SetAppId(v string) *RollbackApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RollbackApplicationRequest) SetAutoEnableApplicationScalingRule(v string) *RollbackApplicationRequest {
	s.AutoEnableApplicationScalingRule = &v
	return s
}

func (s *RollbackApplicationRequest) SetBatchWaitTime(v int32) *RollbackApplicationRequest {
	s.BatchWaitTime = &v
	return s
}

func (s *RollbackApplicationRequest) SetMinReadyInstanceRatio(v int32) *RollbackApplicationRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *RollbackApplicationRequest) SetMinReadyInstances(v int32) *RollbackApplicationRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *RollbackApplicationRequest) SetUpdateStrategy(v string) *RollbackApplicationRequest {
	s.UpdateStrategy = &v
	return s
}

func (s *RollbackApplicationRequest) SetVersionId(v string) *RollbackApplicationRequest {
	s.VersionId = &v
	return s
}

type RollbackApplicationResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RollbackApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                              `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RollbackApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponseBody) SetCode(v string) *RollbackApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetData(v *RollbackApplicationResponseBodyData) *RollbackApplicationResponseBody {
	s.Data = v
	return s
}

func (s *RollbackApplicationResponseBody) SetErrorCode(v string) *RollbackApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetMessage(v string) *RollbackApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetRequestId(v string) *RollbackApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetSuccess(v bool) *RollbackApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetTraceId(v string) *RollbackApplicationResponseBody {
	s.TraceId = &v
	return s
}

type RollbackApplicationResponseBodyData struct {
	ChangeOrderId  *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	IsNeedApproval *bool   `json:"IsNeedApproval,omitempty" xml:"IsNeedApproval,omitempty"`
}

func (s RollbackApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponseBodyData) SetChangeOrderId(v string) *RollbackApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *RollbackApplicationResponseBodyData) SetIsNeedApproval(v bool) *RollbackApplicationResponseBodyData {
	s.IsNeedApproval = &v
	return s
}

type RollbackApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RollbackApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationResponse) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponse) SetHeaders(v map[string]*string) *RollbackApplicationResponse {
	s.Headers = v
	return s
}

func (s *RollbackApplicationResponse) SetStatusCode(v int32) *RollbackApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackApplicationResponse) SetBody(v *RollbackApplicationResponseBody) *RollbackApplicationResponse {
	s.Body = v
	return s
}

type StartApplicationRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s StartApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s StartApplicationRequest) GoString() string {
	return s.String()
}

func (s *StartApplicationRequest) SetAppId(v string) *StartApplicationRequest {
	s.AppId = &v
	return s
}

type StartApplicationResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *StartApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s StartApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StartApplicationResponseBody) SetCode(v string) *StartApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StartApplicationResponseBody) SetData(v *StartApplicationResponseBodyData) *StartApplicationResponseBody {
	s.Data = v
	return s
}

func (s *StartApplicationResponseBody) SetErrorCode(v string) *StartApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartApplicationResponseBody) SetMessage(v string) *StartApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StartApplicationResponseBody) SetRequestId(v string) *StartApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartApplicationResponseBody) SetSuccess(v bool) *StartApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *StartApplicationResponseBody) SetTraceId(v string) *StartApplicationResponseBody {
	s.TraceId = &v
	return s
}

type StartApplicationResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s StartApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartApplicationResponseBodyData) SetChangeOrderId(v string) *StartApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type StartApplicationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s StartApplicationResponse) GoString() string {
	return s.String()
}

func (s *StartApplicationResponse) SetHeaders(v map[string]*string) *StartApplicationResponse {
	s.Headers = v
	return s
}

func (s *StartApplicationResponse) SetStatusCode(v int32) *StartApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartApplicationResponse) SetBody(v *StartApplicationResponseBody) *StartApplicationResponse {
	s.Body = v
	return s
}

type StopApplicationRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s StopApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopApplicationRequest) GoString() string {
	return s.String()
}

func (s *StopApplicationRequest) SetAppId(v string) *StopApplicationRequest {
	s.AppId = &v
	return s
}

type StopApplicationResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *StopApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                          `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s StopApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StopApplicationResponseBody) SetCode(v string) *StopApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StopApplicationResponseBody) SetData(v *StopApplicationResponseBodyData) *StopApplicationResponseBody {
	s.Data = v
	return s
}

func (s *StopApplicationResponseBody) SetErrorCode(v string) *StopApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopApplicationResponseBody) SetMessage(v string) *StopApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StopApplicationResponseBody) SetRequestId(v string) *StopApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopApplicationResponseBody) SetSuccess(v bool) *StopApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *StopApplicationResponseBody) SetTraceId(v string) *StopApplicationResponseBody {
	s.TraceId = &v
	return s
}

type StopApplicationResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s StopApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StopApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopApplicationResponseBodyData) SetChangeOrderId(v string) *StopApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type StopApplicationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopApplicationResponse) GoString() string {
	return s.String()
}

func (s *StopApplicationResponse) SetHeaders(v map[string]*string) *StopApplicationResponse {
	s.Headers = v
	return s
}

func (s *StopApplicationResponse) SetStatusCode(v int32) *StopApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopApplicationResponse) SetBody(v *StopApplicationResponseBody) *StopApplicationResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIds  *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v string) *TagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v string) *TagResourcesRequest {
	s.Tags = &v
	return s
}

type TagResourcesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetCode(v string) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetData(v bool) *TagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *TagResourcesResponseBody) SetErrorCode(v string) *TagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TagResourcesResponseBody) SetMessage(v string) *TagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesResponseBody) SetTraceId(v string) *TagResourcesResponseBody {
	s.TraceId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnbindSlbRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Internet *bool   `json:"Internet,omitempty" xml:"Internet,omitempty"`
	Intranet *bool   `json:"Intranet,omitempty" xml:"Intranet,omitempty"`
}

func (s UnbindSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindSlbRequest) GoString() string {
	return s.String()
}

func (s *UnbindSlbRequest) SetAppId(v string) *UnbindSlbRequest {
	s.AppId = &v
	return s
}

func (s *UnbindSlbRequest) SetInternet(v bool) *UnbindSlbRequest {
	s.Internet = &v
	return s
}

func (s *UnbindSlbRequest) SetIntranet(v bool) *UnbindSlbRequest {
	s.Intranet = &v
	return s
}

type UnbindSlbResponseBody struct {
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UnbindSlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                    `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UnbindSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindSlbResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSlbResponseBody) SetCode(v string) *UnbindSlbResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindSlbResponseBody) SetData(v *UnbindSlbResponseBodyData) *UnbindSlbResponseBody {
	s.Data = v
	return s
}

func (s *UnbindSlbResponseBody) SetErrorCode(v string) *UnbindSlbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnbindSlbResponseBody) SetMessage(v string) *UnbindSlbResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindSlbResponseBody) SetRequestId(v string) *UnbindSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindSlbResponseBody) SetSuccess(v bool) *UnbindSlbResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindSlbResponseBody) SetTraceId(v string) *UnbindSlbResponseBody {
	s.TraceId = &v
	return s
}

type UnbindSlbResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s UnbindSlbResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnbindSlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnbindSlbResponseBodyData) SetChangeOrderId(v string) *UnbindSlbResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type UnbindSlbResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindSlbResponse) GoString() string {
	return s.String()
}

func (s *UnbindSlbResponse) SetHeaders(v map[string]*string) *UnbindSlbResponse {
	s.Headers = v
	return s
}

func (s *UnbindSlbResponse) SetStatusCode(v int32) *UnbindSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindSlbResponse) SetBody(v *UnbindSlbResponseBody) *UnbindSlbResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	DeleteAll    *bool   `json:"DeleteAll,omitempty" xml:"DeleteAll,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceIds  *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKeys      *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetDeleteAll(v bool) *UntagResourcesRequest {
	s.DeleteAll = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v string) *UntagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v string) *UntagResourcesRequest {
	s.TagKeys = &v
	return s
}

type UntagResourcesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetData(v bool) *UntagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *UntagResourcesResponseBody) SetErrorCode(v string) *UntagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetMessage(v string) *UntagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) SetTraceId(v string) *UntagResourcesResponseBody {
	s.TraceId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAppSecurityGroupRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s UpdateAppSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppSecurityGroupRequest) SetAppId(v string) *UpdateAppSecurityGroupRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppSecurityGroupRequest) SetSecurityGroupId(v string) *UpdateAppSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

type UpdateAppSecurityGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateAppSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppSecurityGroupResponseBody) SetCode(v string) *UpdateAppSecurityGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetErrorCode(v string) *UpdateAppSecurityGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetMessage(v string) *UpdateAppSecurityGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetRequestId(v string) *UpdateAppSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetSuccess(v bool) *UpdateAppSecurityGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAppSecurityGroupResponseBody) SetTraceId(v string) *UpdateAppSecurityGroupResponseBody {
	s.TraceId = &v
	return s
}

type UpdateAppSecurityGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppSecurityGroupResponse) SetHeaders(v map[string]*string) *UpdateAppSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppSecurityGroupResponse) SetStatusCode(v int32) *UpdateAppSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppSecurityGroupResponse) SetBody(v *UpdateAppSecurityGroupResponseBody) *UpdateAppSecurityGroupResponse {
	s.Body = v
	return s
}

type UpdateApplicationDescriptionRequest struct {
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s UpdateApplicationDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionRequest) SetAppDescription(v string) *UpdateApplicationDescriptionRequest {
	s.AppDescription = &v
	return s
}

func (s *UpdateApplicationDescriptionRequest) SetAppId(v string) *UpdateApplicationDescriptionRequest {
	s.AppId = &v
	return s
}

type UpdateApplicationDescriptionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateApplicationDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionResponseBody) SetCode(v string) *UpdateApplicationDescriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetErrorCode(v string) *UpdateApplicationDescriptionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetMessage(v string) *UpdateApplicationDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetRequestId(v string) *UpdateApplicationDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetSuccess(v bool) *UpdateApplicationDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApplicationDescriptionResponseBody) SetTraceId(v string) *UpdateApplicationDescriptionResponseBody {
	s.TraceId = &v
	return s
}

type UpdateApplicationDescriptionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateApplicationDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicationDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionResponse) SetHeaders(v map[string]*string) *UpdateApplicationDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationDescriptionResponse) SetStatusCode(v int32) *UpdateApplicationDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationDescriptionResponse) SetBody(v *UpdateApplicationDescriptionResponseBody) *UpdateApplicationDescriptionResponse {
	s.Body = v
	return s
}

type UpdateApplicationScalingRuleRequest struct {
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	MinReadyInstanceRatio *int32  `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	MinReadyInstances     *int32  `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	ScalingRuleMetric     *string `json:"ScalingRuleMetric,omitempty" xml:"ScalingRuleMetric,omitempty"`
	ScalingRuleName       *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	ScalingRuleTimer      *string `json:"ScalingRuleTimer,omitempty" xml:"ScalingRuleTimer,omitempty"`
}

func (s UpdateApplicationScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleRequest) SetAppId(v string) *UpdateApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetMinReadyInstanceRatio(v int32) *UpdateApplicationScalingRuleRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetMinReadyInstances(v int32) *UpdateApplicationScalingRuleRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingRuleMetric(v string) *UpdateApplicationScalingRuleRequest {
	s.ScalingRuleMetric = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingRuleName(v string) *UpdateApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *UpdateApplicationScalingRuleRequest) SetScalingRuleTimer(v string) *UpdateApplicationScalingRuleRequest {
	s.ScalingRuleTimer = &v
	return s
}

type UpdateApplicationScalingRuleResponseBody struct {
	Data      *UpdateApplicationScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string                                       `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBody) SetData(v *UpdateApplicationScalingRuleResponseBodyData) *UpdateApplicationScalingRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetRequestId(v string) *UpdateApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetTraceId(v string) *UpdateApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

type UpdateApplicationScalingRuleResponseBodyData struct {
	AppId            *string                                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime       *int64                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastDisableTime  *int64                                              `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	Metric           *UpdateApplicationScalingRuleResponseBodyDataMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	ScaleRuleEnabled *bool                                               `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	ScaleRuleName    *string                                             `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	ScaleRuleType    *string                                             `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	Timer            *UpdateApplicationScalingRuleResponseBodyDataTimer  `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	UpdateTime       *int64                                              `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetAppId(v string) *UpdateApplicationScalingRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetCreateTime(v int64) *UpdateApplicationScalingRuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetLastDisableTime(v int64) *UpdateApplicationScalingRuleResponseBodyData {
	s.LastDisableTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetMetric(v *UpdateApplicationScalingRuleResponseBodyDataMetric) *UpdateApplicationScalingRuleResponseBodyData {
	s.Metric = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetScaleRuleEnabled(v bool) *UpdateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetScaleRuleName(v string) *UpdateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleName = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetScaleRuleType(v string) *UpdateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleType = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetTimer(v *UpdateApplicationScalingRuleResponseBodyDataTimer) *UpdateApplicationScalingRuleResponseBodyData {
	s.Timer = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetUpdateTime(v int64) *UpdateApplicationScalingRuleResponseBodyData {
	s.UpdateTime = &v
	return s
}

type UpdateApplicationScalingRuleResponseBodyDataMetric struct {
	MaxReplicas *int32                                                       `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	Metrics     []*UpdateApplicationScalingRuleResponseBodyDataMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	MinReplicas *int32                                                       `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyDataMetric) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyDataMetric) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) SetMaxReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyDataMetric {
	s.MaxReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) SetMetrics(v []*UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) *UpdateApplicationScalingRuleResponseBodyDataMetric {
	s.Metrics = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) SetMinReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyDataMetric {
	s.MinReplicas = &v
	return s
}

type UpdateApplicationScalingRuleResponseBodyDataMetricMetrics struct {
	MetricTargetAverageUtilization *int32  `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	MetricType                     *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricTargetAverageUtilization(v int32) *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricType(v string) *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricType = &v
	return s
}

type UpdateApplicationScalingRuleResponseBodyDataTimer struct {
	BeginDate *string                                                       `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate   *string                                                       `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Period    *string                                                       `json:"Period,omitempty" xml:"Period,omitempty"`
	Schedules []*UpdateApplicationScalingRuleResponseBodyDataTimerSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
}

func (s UpdateApplicationScalingRuleResponseBodyDataTimer) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyDataTimer) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) SetBeginDate(v string) *UpdateApplicationScalingRuleResponseBodyDataTimer {
	s.BeginDate = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) SetEndDate(v string) *UpdateApplicationScalingRuleResponseBodyDataTimer {
	s.EndDate = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) SetPeriod(v string) *UpdateApplicationScalingRuleResponseBodyDataTimer {
	s.Period = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) SetSchedules(v []*UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) *UpdateApplicationScalingRuleResponseBodyDataTimer {
	s.Schedules = v
	return s
}

type UpdateApplicationScalingRuleResponseBodyDataTimerSchedules struct {
	AtTime         *string `json:"AtTime,omitempty" xml:"AtTime,omitempty"`
	TargetReplicas *int32  `json:"TargetReplicas,omitempty" xml:"TargetReplicas,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) SetAtTime(v string) *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.AtTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) SetTargetReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.TargetReplicas = &v
	return s
}

type UpdateApplicationScalingRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicationScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *UpdateApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationScalingRuleResponse) SetStatusCode(v int32) *UpdateApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponse) SetBody(v *UpdateApplicationScalingRuleResponseBody) *UpdateApplicationScalingRuleResponse {
	s.Body = v
	return s
}

type UpdateApplicationVswitchesRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UpdateApplicationVswitchesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationVswitchesRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVswitchesRequest) SetAppId(v string) *UpdateApplicationVswitchesRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationVswitchesRequest) SetVSwitchId(v string) *UpdateApplicationVswitchesRequest {
	s.VSwitchId = &v
	return s
}

type UpdateApplicationVswitchesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateApplicationVswitchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationVswitchesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVswitchesResponseBody) SetCode(v string) *UpdateApplicationVswitchesResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetErrorCode(v string) *UpdateApplicationVswitchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetMessage(v string) *UpdateApplicationVswitchesResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetRequestId(v string) *UpdateApplicationVswitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetSuccess(v bool) *UpdateApplicationVswitchesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetTraceId(v string) *UpdateApplicationVswitchesResponseBody {
	s.TraceId = &v
	return s
}

type UpdateApplicationVswitchesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateApplicationVswitchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicationVswitchesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationVswitchesResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVswitchesResponse) SetHeaders(v map[string]*string) *UpdateApplicationVswitchesResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationVswitchesResponse) SetStatusCode(v int32) *UpdateApplicationVswitchesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationVswitchesResponse) SetBody(v *UpdateApplicationVswitchesResponseBody) *UpdateApplicationVswitchesResponse {
	s.Body = v
	return s
}

type UpdateConfigMapRequest struct {
	ConfigMapId *int64  `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateConfigMapRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigMapRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigMapRequest) SetConfigMapId(v int64) *UpdateConfigMapRequest {
	s.ConfigMapId = &v
	return s
}

func (s *UpdateConfigMapRequest) SetData(v string) *UpdateConfigMapRequest {
	s.Data = &v
	return s
}

func (s *UpdateConfigMapRequest) SetDescription(v string) *UpdateConfigMapRequest {
	s.Description = &v
	return s
}

type UpdateConfigMapResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UpdateConfigMapResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                          `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateConfigMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigMapResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigMapResponseBody) SetCode(v string) *UpdateConfigMapResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetData(v *UpdateConfigMapResponseBodyData) *UpdateConfigMapResponseBody {
	s.Data = v
	return s
}

func (s *UpdateConfigMapResponseBody) SetErrorCode(v string) *UpdateConfigMapResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetMessage(v string) *UpdateConfigMapResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetRequestId(v string) *UpdateConfigMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetSuccess(v bool) *UpdateConfigMapResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConfigMapResponseBody) SetTraceId(v string) *UpdateConfigMapResponseBody {
	s.TraceId = &v
	return s
}

type UpdateConfigMapResponseBodyData struct {
	ConfigMapId *string `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
}

func (s UpdateConfigMapResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigMapResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateConfigMapResponseBodyData) SetConfigMapId(v string) *UpdateConfigMapResponseBodyData {
	s.ConfigMapId = &v
	return s
}

type UpdateConfigMapResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConfigMapResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigMapResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigMapResponse) SetHeaders(v map[string]*string) *UpdateConfigMapResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigMapResponse) SetStatusCode(v int32) *UpdateConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigMapResponse) SetBody(v *UpdateConfigMapResponseBody) *UpdateConfigMapResponse {
	s.Body = v
	return s
}

type UpdateGreyTagRouteRequest struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DubboRules     *string `json:"DubboRules,omitempty" xml:"DubboRules,omitempty"`
	GreyTagRouteId *int64  `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
	ScRules        *string `json:"ScRules,omitempty" xml:"ScRules,omitempty"`
}

func (s UpdateGreyTagRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGreyTagRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateGreyTagRouteRequest) SetDescription(v string) *UpdateGreyTagRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateGreyTagRouteRequest) SetDubboRules(v string) *UpdateGreyTagRouteRequest {
	s.DubboRules = &v
	return s
}

func (s *UpdateGreyTagRouteRequest) SetGreyTagRouteId(v int64) *UpdateGreyTagRouteRequest {
	s.GreyTagRouteId = &v
	return s
}

func (s *UpdateGreyTagRouteRequest) SetScRules(v string) *UpdateGreyTagRouteRequest {
	s.ScRules = &v
	return s
}

type UpdateGreyTagRouteResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UpdateGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                             `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateGreyTagRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGreyTagRouteResponseBody) SetCode(v string) *UpdateGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetData(v *UpdateGreyTagRouteResponseBodyData) *UpdateGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetErrorCode(v string) *UpdateGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetMessage(v string) *UpdateGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetRequestId(v string) *UpdateGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetSuccess(v bool) *UpdateGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetTraceId(v string) *UpdateGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

type UpdateGreyTagRouteResponseBodyData struct {
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s UpdateGreyTagRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateGreyTagRouteResponseBodyData) SetGreyTagRouteId(v int64) *UpdateGreyTagRouteResponseBodyData {
	s.GreyTagRouteId = &v
	return s
}

type UpdateGreyTagRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGreyTagRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateGreyTagRouteResponse) SetHeaders(v map[string]*string) *UpdateGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateGreyTagRouteResponse) SetStatusCode(v int32) *UpdateGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGreyTagRouteResponse) SetBody(v *UpdateGreyTagRouteResponseBody) *UpdateGreyTagRouteResponse {
	s.Body = v
	return s
}

type UpdateIngressRequest struct {
	CertId           *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	DefaultRule      *string `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IngressId        *int64  `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
	ListenerPort     *string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalanceType  *string `json:"LoadBalanceType,omitempty" xml:"LoadBalanceType,omitempty"`
	Rules            *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s UpdateIngressRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIngressRequest) GoString() string {
	return s.String()
}

func (s *UpdateIngressRequest) SetCertId(v string) *UpdateIngressRequest {
	s.CertId = &v
	return s
}

func (s *UpdateIngressRequest) SetDefaultRule(v string) *UpdateIngressRequest {
	s.DefaultRule = &v
	return s
}

func (s *UpdateIngressRequest) SetDescription(v string) *UpdateIngressRequest {
	s.Description = &v
	return s
}

func (s *UpdateIngressRequest) SetIngressId(v int64) *UpdateIngressRequest {
	s.IngressId = &v
	return s
}

func (s *UpdateIngressRequest) SetListenerPort(v string) *UpdateIngressRequest {
	s.ListenerPort = &v
	return s
}

func (s *UpdateIngressRequest) SetListenerProtocol(v string) *UpdateIngressRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *UpdateIngressRequest) SetLoadBalanceType(v string) *UpdateIngressRequest {
	s.LoadBalanceType = &v
	return s
}

func (s *UpdateIngressRequest) SetRules(v string) *UpdateIngressRequest {
	s.Rules = &v
	return s
}

type UpdateIngressResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UpdateIngressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                        `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateIngressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIngressResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIngressResponseBody) SetCode(v string) *UpdateIngressResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateIngressResponseBody) SetData(v *UpdateIngressResponseBodyData) *UpdateIngressResponseBody {
	s.Data = v
	return s
}

func (s *UpdateIngressResponseBody) SetErrorCode(v string) *UpdateIngressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateIngressResponseBody) SetMessage(v string) *UpdateIngressResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateIngressResponseBody) SetRequestId(v string) *UpdateIngressResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIngressResponseBody) SetSuccess(v bool) *UpdateIngressResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateIngressResponseBody) SetTraceId(v string) *UpdateIngressResponseBody {
	s.TraceId = &v
	return s
}

type UpdateIngressResponseBodyData struct {
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s UpdateIngressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateIngressResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateIngressResponseBodyData) SetIngressId(v int64) *UpdateIngressResponseBodyData {
	s.IngressId = &v
	return s
}

type UpdateIngressResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIngressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIngressResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIngressResponse) GoString() string {
	return s.String()
}

func (s *UpdateIngressResponse) SetHeaders(v map[string]*string) *UpdateIngressResponse {
	s.Headers = v
	return s
}

func (s *UpdateIngressResponse) SetStatusCode(v int32) *UpdateIngressResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIngressResponse) SetBody(v *UpdateIngressResponseBody) *UpdateIngressResponse {
	s.Body = v
	return s
}

type UpdateNamespaceRequest struct {
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	NamespaceId          *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName        *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceRequest) SetNamespaceDescription(v string) *UpdateNamespaceRequest {
	s.NamespaceDescription = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceId(v string) *UpdateNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceName(v string) *UpdateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type UpdateNamespaceResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UpdateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                          `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBody) SetCode(v string) *UpdateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetData(v *UpdateNamespaceResponseBodyData) *UpdateNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateNamespaceResponseBody) SetErrorCode(v string) *UpdateNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetMessage(v string) *UpdateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetRequestId(v string) *UpdateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetSuccess(v bool) *UpdateNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetTraceId(v string) *UpdateNamespaceResponseBody {
	s.TraceId = &v
	return s
}

type UpdateNamespaceResponseBodyData struct {
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	NamespaceId          *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName        *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateNamespaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBodyData) SetNamespaceDescription(v string) *UpdateNamespaceResponseBodyData {
	s.NamespaceDescription = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) SetNamespaceId(v string) *UpdateNamespaceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) SetNamespaceName(v string) *UpdateNamespaceResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) SetRegionId(v string) *UpdateNamespaceResponseBodyData {
	s.RegionId = &v
	return s
}

type UpdateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponse) SetHeaders(v map[string]*string) *UpdateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceResponse) SetStatusCode(v int32) *UpdateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceResponse) SetBody(v *UpdateNamespaceResponseBody) *UpdateNamespaceResponse {
	s.Body = v
	return s
}

type UpdateNamespaceVpcRequest struct {
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateNamespaceVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceVpcRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceVpcRequest) SetNamespaceId(v string) *UpdateNamespaceVpcRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceVpcRequest) SetVpcId(v string) *UpdateNamespaceVpcRequest {
	s.VpcId = &v
	return s
}

type UpdateNamespaceVpcResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateNamespaceVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceVpcResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceVpcResponseBody) SetCode(v string) *UpdateNamespaceVpcResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetErrorCode(v string) *UpdateNamespaceVpcResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetMessage(v string) *UpdateNamespaceVpcResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetRequestId(v string) *UpdateNamespaceVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetSuccess(v bool) *UpdateNamespaceVpcResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNamespaceVpcResponseBody) SetTraceId(v string) *UpdateNamespaceVpcResponseBody {
	s.TraceId = &v
	return s
}

type UpdateNamespaceVpcResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateNamespaceVpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNamespaceVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceVpcResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceVpcResponse) SetHeaders(v map[string]*string) *UpdateNamespaceVpcResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceVpcResponse) SetStatusCode(v int32) *UpdateNamespaceVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceVpcResponse) SetBody(v *UpdateNamespaceVpcResponseBody) *UpdateNamespaceVpcResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sae"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AbortAndRollbackChangeOrder(request *AbortAndRollbackChangeOrderRequest) (_result *AbortAndRollbackChangeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AbortAndRollbackChangeOrderResponse{}
	_body, _err := client.AbortAndRollbackChangeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortAndRollbackChangeOrderWithOptions(request *AbortAndRollbackChangeOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AbortAndRollbackChangeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeOrderId)) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbortAndRollbackChangeOrder"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/changeorder/AbortAndRollbackChangeOrder"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AbortAndRollbackChangeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbortChangeOrder(request *AbortChangeOrderRequest) (_result *AbortChangeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AbortChangeOrderResponse{}
	_body, _err := client.AbortChangeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortChangeOrderWithOptions(request *AbortChangeOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AbortChangeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeOrderId)) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbortChangeOrder"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/changeorder/AbortChangeOrder"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AbortChangeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStartApplications(request *BatchStartApplicationsRequest) (_result *BatchStartApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchStartApplicationsResponse{}
	_body, _err := client.BatchStartApplicationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStartApplicationsWithOptions(request *BatchStartApplicationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchStartApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		query["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStartApplications"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/batchStartApplications"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStartApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStopApplications(request *BatchStopApplicationsRequest) (_result *BatchStopApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchStopApplicationsResponse{}
	_body, _err := client.BatchStopApplicationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStopApplicationsWithOptions(request *BatchStopApplicationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchStopApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		query["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStopApplications"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/batchStopApplications"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStopApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindSlb(request *BindSlbRequest) (_result *BindSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindSlbResponse{}
	_body, _err := client.BindSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindSlbWithOptions(request *BindSlbRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BindSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Internet)) {
		query["Internet"] = request.Internet
	}

	if !tea.BoolValue(util.IsUnset(request.InternetSlbId)) {
		query["InternetSlbId"] = request.InternetSlbId
	}

	if !tea.BoolValue(util.IsUnset(request.Intranet)) {
		query["Intranet"] = request.Intranet
	}

	if !tea.BoolValue(util.IsUnset(request.IntranetSlbId)) {
		query["IntranetSlbId"] = request.IntranetSlbId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindSlb"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/slb"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BindSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmPipelineBatch(request *ConfirmPipelineBatchRequest) (_result *ConfirmPipelineBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConfirmPipelineBatchResponse{}
	_body, _err := client.ConfirmPipelineBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmPipelineBatchWithOptions(request *ConfirmPipelineBatchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConfirmPipelineBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Confirm)) {
		query["Confirm"] = request.Confirm
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmPipelineBatch"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/changeorder/ConfirmPipelineBatch"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmPipelineBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrAssumeRoleArn)) {
		query["AcrAssumeRoleArn"] = request.AcrAssumeRoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.AppDescription)) {
		query["AppDescription"] = request.AppDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AutoConfig)) {
		query["AutoConfig"] = request.AutoConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Command)) {
		query["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.CommandArgs)) {
		query["CommandArgs"] = request.CommandArgs
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.CustomHostAlias)) {
		query["CustomHostAlias"] = request.CustomHostAlias
	}

	if !tea.BoolValue(util.IsUnset(request.Deploy)) {
		query["Deploy"] = request.Deploy
	}

	if !tea.BoolValue(util.IsUnset(request.EdasContainerVersion)) {
		query["EdasContainerVersion"] = request.EdasContainerVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Envs)) {
		query["Envs"] = request.Envs
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.JarStartArgs)) {
		query["JarStartArgs"] = request.JarStartArgs
	}

	if !tea.BoolValue(util.IsUnset(request.JarStartOptions)) {
		query["JarStartOptions"] = request.JarStartOptions
	}

	if !tea.BoolValue(util.IsUnset(request.Jdk)) {
		query["Jdk"] = request.Jdk
	}

	if !tea.BoolValue(util.IsUnset(request.KafkaConfigs)) {
		query["KafkaConfigs"] = request.KafkaConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.Liveness)) {
		query["Liveness"] = request.Liveness
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.MountDesc)) {
		query["MountDesc"] = request.MountDesc
	}

	if !tea.BoolValue(util.IsUnset(request.MountHost)) {
		query["MountHost"] = request.MountHost
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.NasId)) {
		query["NasId"] = request.NasId
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		query["PackageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.PackageUrl)) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PackageVersion)) {
		query["PackageVersion"] = request.PackageVersion
	}

	if !tea.BoolValue(util.IsUnset(request.PhpArmsConfigLocation)) {
		query["PhpArmsConfigLocation"] = request.PhpArmsConfigLocation
	}

	if !tea.BoolValue(util.IsUnset(request.PhpConfigLocation)) {
		query["PhpConfigLocation"] = request.PhpConfigLocation
	}

	if !tea.BoolValue(util.IsUnset(request.PostStart)) {
		query["PostStart"] = request.PostStart
	}

	if !tea.BoolValue(util.IsUnset(request.PreStop)) {
		query["PreStop"] = request.PreStop
	}

	if !tea.BoolValue(util.IsUnset(request.ProgrammingLanguage)) {
		query["ProgrammingLanguage"] = request.ProgrammingLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.Readiness)) {
		query["Readiness"] = request.Readiness
	}

	if !tea.BoolValue(util.IsUnset(request.Replicas)) {
		query["Replicas"] = request.Replicas
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SlsConfigs)) {
		query["SlsConfigs"] = request.SlsConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationGracePeriodSeconds)) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Timezone)) {
		query["Timezone"] = request.Timezone
	}

	if !tea.BoolValue(util.IsUnset(request.TomcatConfig)) {
		query["TomcatConfig"] = request.TomcatConfig
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.WarStartOptions)) {
		query["WarStartOptions"] = request.WarStartOptions
	}

	if !tea.BoolValue(util.IsUnset(request.WebContainer)) {
		query["WebContainer"] = request.WebContainer
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrInstanceId)) {
		body["AcrInstanceId"] = request.AcrInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AssociateEip)) {
		body["AssociateEip"] = request.AssociateEip
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigMapMountDesc)) {
		body["ConfigMapMountDesc"] = request.ConfigMapMountDesc
	}

	if !tea.BoolValue(util.IsUnset(request.OssAkId)) {
		body["OssAkId"] = request.OssAkId
	}

	if !tea.BoolValue(util.IsUnset(request.OssAkSecret)) {
		body["OssAkSecret"] = request.OssAkSecret
	}

	if !tea.BoolValue(util.IsUnset(request.OssMountDescs)) {
		body["OssMountDescs"] = request.OssMountDescs
	}

	if !tea.BoolValue(util.IsUnset(request.PhpConfig)) {
		body["PhpConfig"] = request.PhpConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplication"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/createApplication"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplicationScalingRule(request *CreateApplicationScalingRuleRequest) (_result *CreateApplicationScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateApplicationScalingRuleResponse{}
	_body, _err := client.CreateApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationScalingRuleWithOptions(request *CreateApplicationScalingRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateApplicationScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstanceRatio)) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstances)) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleEnable)) {
		query["ScalingRuleEnable"] = request.ScalingRuleEnable
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleMetric)) {
		query["ScalingRuleMetric"] = request.ScalingRuleMetric
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleTimer)) {
		query["ScalingRuleTimer"] = request.ScalingRuleTimer
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleType)) {
		query["ScalingRuleType"] = request.ScalingRuleType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplicationScalingRule"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/scale/applicationScalingRule"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConfigMap(request *CreateConfigMapRequest) (_result *CreateConfigMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConfigMapResponse{}
	_body, _err := client.CreateConfigMapWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConfigMapWithOptions(request *CreateConfigMapRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConfigMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConfigMap"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/configmap/configMap"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConfigMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGreyTagRoute(request *CreateGreyTagRouteRequest) (_result *CreateGreyTagRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGreyTagRouteResponse{}
	_body, _err := client.CreateGreyTagRouteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGreyTagRouteWithOptions(request *CreateGreyTagRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGreyTagRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DubboRules)) {
		query["DubboRules"] = request.DubboRules
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ScRules)) {
		query["ScRules"] = request.ScRules
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGreyTagRoute"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/tagroute/greyTagRoute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGreyTagRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIngress(request *CreateIngressRequest) (_result *CreateIngressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIngressResponse{}
	_body, _err := client.CreateIngressWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIngressWithOptions(request *CreateIngressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIngressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertId)) {
		query["CertId"] = request.CertId
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRule)) {
		query["DefaultRule"] = request.DefaultRule
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalanceType)) {
		query["LoadBalanceType"] = request.LoadBalanceType
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.SlbId)) {
		query["SlbId"] = request.SlbId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Rules)) {
		body["Rules"] = request.Rules
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIngress"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/ingress/Ingress"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIngressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceDescription)) {
		query["NamespaceDescription"] = request.NamespaceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNamespace"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/paas/namespace"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplication"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/deleteApplication"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplicationScalingRule(request *DeleteApplicationScalingRuleRequest) (_result *DeleteApplicationScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApplicationScalingRuleResponse{}
	_body, _err := client.DeleteApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApplicationScalingRuleWithOptions(request *DeleteApplicationScalingRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteApplicationScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplicationScalingRule"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/scale/applicationScalingRule"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConfigMap(request *DeleteConfigMapRequest) (_result *DeleteConfigMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigMapResponse{}
	_body, _err := client.DeleteConfigMapWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConfigMapWithOptions(request *DeleteConfigMapRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConfigMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigMapId)) {
		query["ConfigMapId"] = request.ConfigMapId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConfigMap"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/configmap/configMap"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConfigMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGreyTagRoute(request *DeleteGreyTagRouteRequest) (_result *DeleteGreyTagRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGreyTagRouteResponse{}
	_body, _err := client.DeleteGreyTagRouteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGreyTagRouteWithOptions(request *DeleteGreyTagRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGreyTagRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GreyTagRouteId)) {
		query["GreyTagRouteId"] = request.GreyTagRouteId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGreyTagRoute"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/tagroute/greyTagRoute"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGreyTagRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIngress(request *DeleteIngressRequest) (_result *DeleteIngressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIngressResponse{}
	_body, _err := client.DeleteIngressWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIngressWithOptions(request *DeleteIngressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIngressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IngressId)) {
		query["IngressId"] = request.IngressId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIngress"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/ingress/Ingress"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIngressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNamespace"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/paas/namespace"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployApplication(request *DeployApplicationRequest) (_result *DeployApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployApplicationResponse{}
	_body, _err := client.DeployApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployApplicationWithOptions(request *DeployApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeployApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrAssumeRoleArn)) {
		query["AcrAssumeRoleArn"] = request.AcrAssumeRoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AutoEnableApplicationScalingRule)) {
		query["AutoEnableApplicationScalingRule"] = request.AutoEnableApplicationScalingRule
	}

	if !tea.BoolValue(util.IsUnset(request.BatchWaitTime)) {
		query["BatchWaitTime"] = request.BatchWaitTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeOrderDesc)) {
		query["ChangeOrderDesc"] = request.ChangeOrderDesc
	}

	if !tea.BoolValue(util.IsUnset(request.Command)) {
		query["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.CommandArgs)) {
		query["CommandArgs"] = request.CommandArgs
	}

	if !tea.BoolValue(util.IsUnset(request.CustomHostAlias)) {
		query["CustomHostAlias"] = request.CustomHostAlias
	}

	if !tea.BoolValue(util.IsUnset(request.EdasContainerVersion)) {
		query["EdasContainerVersion"] = request.EdasContainerVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAhas)) {
		query["EnableAhas"] = request.EnableAhas
	}

	if !tea.BoolValue(util.IsUnset(request.EnableGreyTagRoute)) {
		query["EnableGreyTagRoute"] = request.EnableGreyTagRoute
	}

	if !tea.BoolValue(util.IsUnset(request.Envs)) {
		query["Envs"] = request.Envs
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.JarStartArgs)) {
		query["JarStartArgs"] = request.JarStartArgs
	}

	if !tea.BoolValue(util.IsUnset(request.JarStartOptions)) {
		query["JarStartOptions"] = request.JarStartOptions
	}

	if !tea.BoolValue(util.IsUnset(request.Jdk)) {
		query["Jdk"] = request.Jdk
	}

	if !tea.BoolValue(util.IsUnset(request.KafkaConfigs)) {
		query["KafkaConfigs"] = request.KafkaConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.Liveness)) {
		query["Liveness"] = request.Liveness
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstanceRatio)) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstances)) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !tea.BoolValue(util.IsUnset(request.MountDesc)) {
		query["MountDesc"] = request.MountDesc
	}

	if !tea.BoolValue(util.IsUnset(request.MountHost)) {
		query["MountHost"] = request.MountHost
	}

	if !tea.BoolValue(util.IsUnset(request.NasId)) {
		query["NasId"] = request.NasId
	}

	if !tea.BoolValue(util.IsUnset(request.PackageUrl)) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PackageVersion)) {
		query["PackageVersion"] = request.PackageVersion
	}

	if !tea.BoolValue(util.IsUnset(request.PhpArmsConfigLocation)) {
		query["PhpArmsConfigLocation"] = request.PhpArmsConfigLocation
	}

	if !tea.BoolValue(util.IsUnset(request.PhpConfigLocation)) {
		query["PhpConfigLocation"] = request.PhpConfigLocation
	}

	if !tea.BoolValue(util.IsUnset(request.PostStart)) {
		query["PostStart"] = request.PostStart
	}

	if !tea.BoolValue(util.IsUnset(request.PreStop)) {
		query["PreStop"] = request.PreStop
	}

	if !tea.BoolValue(util.IsUnset(request.Readiness)) {
		query["Readiness"] = request.Readiness
	}

	if !tea.BoolValue(util.IsUnset(request.SlsConfigs)) {
		query["SlsConfigs"] = request.SlsConfigs
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationGracePeriodSeconds)) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Timezone)) {
		query["Timezone"] = request.Timezone
	}

	if !tea.BoolValue(util.IsUnset(request.TomcatConfig)) {
		query["TomcatConfig"] = request.TomcatConfig
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateStrategy)) {
		query["UpdateStrategy"] = request.UpdateStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.WarStartOptions)) {
		query["WarStartOptions"] = request.WarStartOptions
	}

	if !tea.BoolValue(util.IsUnset(request.WebContainer)) {
		query["WebContainer"] = request.WebContainer
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrInstanceId)) {
		body["AcrInstanceId"] = request.AcrInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AssociateEip)) {
		body["AssociateEip"] = request.AssociateEip
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigMapMountDesc)) {
		body["ConfigMapMountDesc"] = request.ConfigMapMountDesc
	}

	if !tea.BoolValue(util.IsUnset(request.OssAkId)) {
		body["OssAkId"] = request.OssAkId
	}

	if !tea.BoolValue(util.IsUnset(request.OssAkSecret)) {
		body["OssAkSecret"] = request.OssAkSecret
	}

	if !tea.BoolValue(util.IsUnset(request.OssMountDescs)) {
		body["OssMountDescs"] = request.OssMountDescs
	}

	if !tea.BoolValue(util.IsUnset(request.PhpConfig)) {
		body["PhpConfig"] = request.PhpConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployApplication"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/deployApplication"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppServiceDetail(request *DescribeAppServiceDetailRequest) (_result *DescribeAppServiceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppServiceDetailResponse{}
	_body, _err := client.DescribeAppServiceDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppServiceDetailWithOptions(request *DescribeAppServiceDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppServiceDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceGroup)) {
		query["ServiceGroup"] = request.ServiceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppServiceDetail"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/service/describeAppServiceDetail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppServiceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationConfig(request *DescribeApplicationConfigRequest) (_result *DescribeApplicationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationConfigResponse{}
	_body, _err := client.DescribeApplicationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationConfigWithOptions(request *DescribeApplicationConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApplicationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationConfig"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/describeApplicationConfig"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationGroups(request *DescribeApplicationGroupsRequest) (_result *DescribeApplicationGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationGroupsResponse{}
	_body, _err := client.DescribeApplicationGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationGroupsWithOptions(request *DescribeApplicationGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApplicationGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationGroups"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/describeApplicationGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationImage(request *DescribeApplicationImageRequest) (_result *DescribeApplicationImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationImageResponse{}
	_body, _err := client.DescribeApplicationImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationImageWithOptions(request *DescribeApplicationImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApplicationImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		query["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationImage"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/container/describeApplicationImage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationInstances(request *DescribeApplicationInstancesRequest) (_result *DescribeApplicationInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationInstancesResponse{}
	_body, _err := client.DescribeApplicationInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationInstancesWithOptions(request *DescribeApplicationInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApplicationInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationInstances"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/describeApplicationInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationScalingRule(request *DescribeApplicationScalingRuleRequest) (_result *DescribeApplicationScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationScalingRuleResponse{}
	_body, _err := client.DescribeApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationScalingRuleWithOptions(request *DescribeApplicationScalingRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApplicationScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationScalingRule"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/scale/applicationScalingRule"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationScalingRules(request *DescribeApplicationScalingRulesRequest) (_result *DescribeApplicationScalingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationScalingRulesResponse{}
	_body, _err := client.DescribeApplicationScalingRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationScalingRulesWithOptions(request *DescribeApplicationScalingRulesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApplicationScalingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationScalingRules"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/scale/applicationScalingRules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationScalingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationSlbs(request *DescribeApplicationSlbsRequest) (_result *DescribeApplicationSlbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationSlbsResponse{}
	_body, _err := client.DescribeApplicationSlbsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationSlbsWithOptions(request *DescribeApplicationSlbsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApplicationSlbsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationSlbs"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/slb"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationSlbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationStatus(request *DescribeApplicationStatusRequest) (_result *DescribeApplicationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationStatusResponse{}
	_body, _err := client.DescribeApplicationStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationStatusWithOptions(request *DescribeApplicationStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeApplicationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationStatus"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/describeApplicationStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChangeOrder(request *DescribeChangeOrderRequest) (_result *DescribeChangeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeChangeOrderResponse{}
	_body, _err := client.DescribeChangeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChangeOrderWithOptions(request *DescribeChangeOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeChangeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeOrderId)) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChangeOrder"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/changeorder/DescribeChangeOrder"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChangeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComponents(request *DescribeComponentsRequest) (_result *DescribeComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeComponentsResponse{}
	_body, _err := client.DescribeComponentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComponentsWithOptions(request *DescribeComponentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComponents"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/resource/components"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComponentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigMap(request *DescribeConfigMapRequest) (_result *DescribeConfigMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConfigMapResponse{}
	_body, _err := client.DescribeConfigMapWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigMapWithOptions(request *DescribeConfigMapRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeConfigMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigMapId)) {
		query["ConfigMapId"] = request.ConfigMapId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConfigMap"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/configmap/configMap"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConfigMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigurationPrice(request *DescribeConfigurationPriceRequest) (_result *DescribeConfigurationPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConfigurationPriceResponse{}
	_body, _err := client.DescribeConfigurationPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigurationPriceWithOptions(request *DescribeConfigurationPriceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeConfigurationPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.Workload)) {
		query["Workload"] = request.Workload
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConfigurationPrice"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/paas/configurationPrice"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConfigurationPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEdasContainers() (_result *DescribeEdasContainersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeEdasContainersResponse{}
	_body, _err := client.DescribeEdasContainersWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEdasContainersWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeEdasContainersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEdasContainers"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/resource/edasContainers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEdasContainersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGreyTagRoute(request *DescribeGreyTagRouteRequest) (_result *DescribeGreyTagRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeGreyTagRouteResponse{}
	_body, _err := client.DescribeGreyTagRouteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGreyTagRouteWithOptions(request *DescribeGreyTagRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeGreyTagRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GreyTagRouteId)) {
		query["GreyTagRouteId"] = request.GreyTagRouteId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGreyTagRoute"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/tagroute/greyTagRoute"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGreyTagRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIngress(request *DescribeIngressRequest) (_result *DescribeIngressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeIngressResponse{}
	_body, _err := client.DescribeIngressWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIngressWithOptions(request *DescribeIngressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeIngressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IngressId)) {
		query["IngressId"] = request.IngressId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIngress"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/ingress/Ingress"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIngressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceLog(request *DescribeInstanceLogRequest) (_result *DescribeInstanceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceLogResponse{}
	_body, _err := client.DescribeInstanceLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceLogWithOptions(request *DescribeInstanceLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeInstanceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceLog"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/instance/describeInstanceLog"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSpecifications() (_result *DescribeInstanceSpecificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceSpecificationsResponse{}
	_body, _err := client.DescribeInstanceSpecificationsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSpecificationsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecificationsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSpecifications"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/paas/quota/instanceSpecifications"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSpecificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJobStatus(request *DescribeJobStatusRequest) (_result *DescribeJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeJobStatusResponse{}
	_body, _err := client.DescribeJobStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJobStatusWithOptions(request *DescribeJobStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeJobStatus"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/job/describeJobStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespace(request *DescribeNamespaceRequest) (_result *DescribeNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeNamespaceResponse{}
	_body, _err := client.DescribeNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespaceWithOptions(request *DescribeNamespaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespace"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/paas/namespace"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespaceList(request *DescribeNamespaceListRequest) (_result *DescribeNamespaceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeNamespaceListResponse{}
	_body, _err := client.DescribeNamespaceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespaceListWithOptions(request *DescribeNamespaceListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeNamespaceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainCustom)) {
		query["ContainCustom"] = request.ContainCustom
	}

	if !tea.BoolValue(util.IsUnset(request.HybridCloudExclude)) {
		query["HybridCloudExclude"] = request.HybridCloudExclude
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespaceList"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/namespace/describeNamespaceList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespaceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespaceResources(request *DescribeNamespaceResourcesRequest) (_result *DescribeNamespaceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeNamespaceResourcesResponse{}
	_body, _err := client.DescribeNamespaceResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespaceResourcesWithOptions(request *DescribeNamespaceResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeNamespaceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespaceResources"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/namespace/describeNamespaceResources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespaceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (_result *DescribeNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.DescribeNamespacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNamespacesWithOptions(request *DescribeNamespacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNamespaces"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/paas/namespaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePipeline(request *DescribePipelineRequest) (_result *DescribePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribePipelineResponse{}
	_body, _err := client.DescribePipelineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePipelineWithOptions(request *DescribePipelineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribePipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePipeline"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/changeorder/DescribePipeline"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/paas/regionConfig"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableApplicationScalingRule(request *DisableApplicationScalingRuleRequest) (_result *DisableApplicationScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableApplicationScalingRuleResponse{}
	_body, _err := client.DisableApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableApplicationScalingRuleWithOptions(request *DisableApplicationScalingRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableApplicationScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableApplicationScalingRule"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/scale/disableApplicationScalingRule"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableApplicationScalingRule(request *EnableApplicationScalingRuleRequest) (_result *EnableApplicationScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableApplicationScalingRuleResponse{}
	_body, _err := client.EnableApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableApplicationScalingRuleWithOptions(request *EnableApplicationScalingRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableApplicationScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableApplicationScalingRule"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/scale/enableApplicationScalingRule"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecJob(request *ExecJobRequest) (_result *ExecJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecJobResponse{}
	_body, _err := client.ExecJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecJobWithOptions(request *ExecJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Command)) {
		query["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.CommandArgs)) {
		query["CommandArgs"] = request.CommandArgs
	}

	if !tea.BoolValue(util.IsUnset(request.Envs)) {
		query["Envs"] = request.Envs
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		query["EventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.JarStartArgs)) {
		query["JarStartArgs"] = request.JarStartArgs
	}

	if !tea.BoolValue(util.IsUnset(request.JarStartOptions)) {
		query["JarStartOptions"] = request.JarStartOptions
	}

	if !tea.BoolValue(util.IsUnset(request.WarStartOptions)) {
		query["WarStartOptions"] = request.WarStartOptions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecJob"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/job/execJob"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppEvents(request *ListAppEventsRequest) (_result *ListAppEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppEventsResponse{}
	_body, _err := client.ListAppEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppEventsWithOptions(request *ListAppEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectKind)) {
		query["ObjectKind"] = request.ObjectKind
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectName)) {
		query["ObjectName"] = request.ObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		query["Reason"] = request.Reason
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppEvents"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/listAppEvents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppServicesPage(request *ListAppServicesPageRequest) (_result *ListAppServicesPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppServicesPageResponse{}
	_body, _err := client.ListAppServicesPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppServicesPageWithOptions(request *ListAppServicesPageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppServicesPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppServicesPage"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/service/listAppServicesPage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppServicesPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppVersions(request *ListAppVersionsRequest) (_result *ListAppVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppVersionsResponse{}
	_body, _err := client.ListAppVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppVersionsWithOptions(request *ListAppVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppVersions"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/listAppVersions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationsWithOptions(request *ListApplicationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FieldType)) {
		query["FieldType"] = request.FieldType
	}

	if !tea.BoolValue(util.IsUnset(request.FieldValue)) {
		query["FieldValue"] = request.FieldValue
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplications"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/listApplications"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListChangeOrders(request *ListChangeOrdersRequest) (_result *ListChangeOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListChangeOrdersResponse{}
	_body, _err := client.ListChangeOrdersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListChangeOrdersWithOptions(request *ListChangeOrdersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListChangeOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CoStatus)) {
		query["CoStatus"] = request.CoStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CoType)) {
		query["CoType"] = request.CoType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChangeOrders"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/changeorder/ListChangeOrders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChangeOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConsumedServices(request *ListConsumedServicesRequest) (_result *ListConsumedServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumedServicesResponse{}
	_body, _err := client.ListConsumedServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConsumedServicesWithOptions(request *ListConsumedServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConsumedServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConsumedServices"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/service/listConsumedServices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConsumedServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGreyTagRoute(request *ListGreyTagRouteRequest) (_result *ListGreyTagRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGreyTagRouteResponse{}
	_body, _err := client.ListGreyTagRouteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGreyTagRouteWithOptions(request *ListGreyTagRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGreyTagRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGreyTagRoute"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/tagroute/greyTagRouteList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGreyTagRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIngresses(request *ListIngressesRequest) (_result *ListIngressesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIngressesResponse{}
	_body, _err := client.ListIngressesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIngressesWithOptions(request *ListIngressesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIngressesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIngresses"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/ingress/IngressList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIngressesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogConfigs(request *ListLogConfigsRequest) (_result *ListLogConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogConfigsResponse{}
	_body, _err := client.ListLogConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogConfigsWithOptions(request *ListLogConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLogConfigs"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/log/listLogConfigs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLogConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNamespaceChangeOrders(request *ListNamespaceChangeOrdersRequest) (_result *ListNamespaceChangeOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNamespaceChangeOrdersResponse{}
	_body, _err := client.ListNamespaceChangeOrdersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNamespaceChangeOrdersWithOptions(request *ListNamespaceChangeOrdersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNamespaceChangeOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoStatus)) {
		query["CoStatus"] = request.CoStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CoType)) {
		query["CoType"] = request.CoType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNamespaceChangeOrders"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/changeorder/listNamespaceChangeOrders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNamespaceChangeOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNamespacedConfigMaps(request *ListNamespacedConfigMapsRequest) (_result *ListNamespacedConfigMapsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNamespacedConfigMapsResponse{}
	_body, _err := client.ListNamespacedConfigMapsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNamespacedConfigMapsWithOptions(request *ListNamespacedConfigMapsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNamespacedConfigMapsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNamespacedConfigMaps"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/configmap/listNamespacedConfigMaps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNamespacedConfigMapsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublishedServices(request *ListPublishedServicesRequest) (_result *ListPublishedServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPublishedServicesResponse{}
	_body, _err := client.ListPublishedServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublishedServicesWithOptions(request *ListPublishedServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPublishedServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublishedServices"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/service/listPublishedServices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublishedServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenSaeService() (_result *OpenSaeServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenSaeServiceResponse{}
	_body, _err := client.OpenSaeServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenSaeServiceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *OpenSaeServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("OpenSaeService"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/service/open"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenSaeServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryResourceStatics(request *QueryResourceStaticsRequest) (_result *QueryResourceStaticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryResourceStaticsResponse{}
	_body, _err := client.QueryResourceStaticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryResourceStaticsWithOptions(request *QueryResourceStaticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryResourceStaticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryResourceStatics"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/paas/quota/queryResourceStatics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryResourceStaticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReduceApplicationCapacityByInstanceIds(request *ReduceApplicationCapacityByInstanceIdsRequest) (_result *ReduceApplicationCapacityByInstanceIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReduceApplicationCapacityByInstanceIdsResponse{}
	_body, _err := client.ReduceApplicationCapacityByInstanceIdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReduceApplicationCapacityByInstanceIdsWithOptions(request *ReduceApplicationCapacityByInstanceIdsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReduceApplicationCapacityByInstanceIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReduceApplicationCapacityByInstanceIds"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/ScaleInApplicationWithInstanceIds"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReduceApplicationCapacityByInstanceIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RescaleApplication(request *RescaleApplicationRequest) (_result *RescaleApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RescaleApplicationResponse{}
	_body, _err := client.RescaleApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RescaleApplicationWithOptions(request *RescaleApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RescaleApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AutoEnableApplicationScalingRule)) {
		query["AutoEnableApplicationScalingRule"] = request.AutoEnableApplicationScalingRule
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstanceRatio)) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstances)) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !tea.BoolValue(util.IsUnset(request.Replicas)) {
		query["Replicas"] = request.Replicas
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RescaleApplication"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/rescaleApplication"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RescaleApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RescaleApplicationVertically(request *RescaleApplicationVerticallyRequest) (_result *RescaleApplicationVerticallyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RescaleApplicationVerticallyResponse{}
	_body, _err := client.RescaleApplicationVerticallyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RescaleApplicationVerticallyWithOptions(request *RescaleApplicationVerticallyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RescaleApplicationVerticallyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RescaleApplicationVertically"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/rescaleApplicationVertically"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RescaleApplicationVerticallyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartApplication(request *RestartApplicationRequest) (_result *RestartApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartApplicationResponse{}
	_body, _err := client.RestartApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartApplicationWithOptions(request *RestartApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstanceRatio)) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstances)) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartApplication"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/restartApplication"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartInstances(request *RestartInstancesRequest) (_result *RestartInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartInstancesResponse{}
	_body, _err := client.RestartInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartInstancesWithOptions(request *RestartInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartInstances"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/restartInstances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackApplication(request *RollbackApplicationRequest) (_result *RollbackApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RollbackApplicationResponse{}
	_body, _err := client.RollbackApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackApplicationWithOptions(request *RollbackApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RollbackApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AutoEnableApplicationScalingRule)) {
		query["AutoEnableApplicationScalingRule"] = request.AutoEnableApplicationScalingRule
	}

	if !tea.BoolValue(util.IsUnset(request.BatchWaitTime)) {
		query["BatchWaitTime"] = request.BatchWaitTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstanceRatio)) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstances)) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateStrategy)) {
		query["UpdateStrategy"] = request.UpdateStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackApplication"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/rollbackApplication"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartApplication(request *StartApplicationRequest) (_result *StartApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartApplicationResponse{}
	_body, _err := client.StartApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartApplicationWithOptions(request *StartApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartApplication"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/startApplication"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopApplication(request *StopApplicationRequest) (_result *StopApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopApplicationResponse{}
	_body, _err := client.StopApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopApplicationWithOptions(request *StopApplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopApplication"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/stopApplication"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindSlb(request *UnbindSlbRequest) (_result *UnbindSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnbindSlbResponse{}
	_body, _err := client.UnbindSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindSlbWithOptions(request *UnbindSlbRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnbindSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Internet)) {
		query["Internet"] = request.Internet
	}

	if !tea.BoolValue(util.IsUnset(request.Intranet)) {
		query["Intranet"] = request.Intranet
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindSlb"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/slb"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteAll)) {
		query["DeleteAll"] = request.DeleteAll
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeys)) {
		query["TagKeys"] = request.TagKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tags"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppSecurityGroup(request *UpdateAppSecurityGroupRequest) (_result *UpdateAppSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAppSecurityGroupResponse{}
	_body, _err := client.UpdateAppSecurityGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppSecurityGroupWithOptions(request *UpdateAppSecurityGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAppSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppSecurityGroup"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/updateAppSecurityGroup"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplicationDescription(request *UpdateApplicationDescriptionRequest) (_result *UpdateApplicationDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApplicationDescriptionResponse{}
	_body, _err := client.UpdateApplicationDescriptionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicationDescriptionWithOptions(request *UpdateApplicationDescriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateApplicationDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppDescription)) {
		query["AppDescription"] = request.AppDescription
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicationDescription"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/updateAppDescription"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplicationScalingRule(request *UpdateApplicationScalingRuleRequest) (_result *UpdateApplicationScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApplicationScalingRuleResponse{}
	_body, _err := client.UpdateApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicationScalingRuleWithOptions(request *UpdateApplicationScalingRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateApplicationScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstanceRatio)) {
		query["MinReadyInstanceRatio"] = request.MinReadyInstanceRatio
	}

	if !tea.BoolValue(util.IsUnset(request.MinReadyInstances)) {
		query["MinReadyInstances"] = request.MinReadyInstances
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleMetric)) {
		query["ScalingRuleMetric"] = request.ScalingRuleMetric
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleTimer)) {
		query["ScalingRuleTimer"] = request.ScalingRuleTimer
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicationScalingRule"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/scale/applicationScalingRule"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplicationVswitches(request *UpdateApplicationVswitchesRequest) (_result *UpdateApplicationVswitchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApplicationVswitchesResponse{}
	_body, _err := client.UpdateApplicationVswitchesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicationVswitchesWithOptions(request *UpdateApplicationVswitchesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateApplicationVswitchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicationVswitches"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/app/updateAppVswitches"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationVswitchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfigMap(request *UpdateConfigMapRequest) (_result *UpdateConfigMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConfigMapResponse{}
	_body, _err := client.UpdateConfigMapWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigMapWithOptions(request *UpdateConfigMapRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateConfigMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigMapId)) {
		query["ConfigMapId"] = request.ConfigMapId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConfigMap"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/configmap/configMap"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConfigMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGreyTagRoute(request *UpdateGreyTagRouteRequest) (_result *UpdateGreyTagRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGreyTagRouteResponse{}
	_body, _err := client.UpdateGreyTagRouteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGreyTagRouteWithOptions(request *UpdateGreyTagRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGreyTagRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DubboRules)) {
		query["DubboRules"] = request.DubboRules
	}

	if !tea.BoolValue(util.IsUnset(request.GreyTagRouteId)) {
		query["GreyTagRouteId"] = request.GreyTagRouteId
	}

	if !tea.BoolValue(util.IsUnset(request.ScRules)) {
		query["ScRules"] = request.ScRules
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGreyTagRoute"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/tagroute/greyTagRoute"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGreyTagRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIngress(request *UpdateIngressRequest) (_result *UpdateIngressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIngressResponse{}
	_body, _err := client.UpdateIngressWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIngressWithOptions(request *UpdateIngressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateIngressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertId)) {
		query["CertId"] = request.CertId
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRule)) {
		query["DefaultRule"] = request.DefaultRule
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IngressId)) {
		query["IngressId"] = request.IngressId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalanceType)) {
		query["LoadBalanceType"] = request.LoadBalanceType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Rules)) {
		body["Rules"] = request.Rules
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIngress"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/ingress/Ingress"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIngressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNamespace(request *UpdateNamespaceRequest) (_result *UpdateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.UpdateNamespaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNamespaceWithOptions(request *UpdateNamespaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceDescription)) {
		query["NamespaceDescription"] = request.NamespaceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNamespace"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/paas/namespace"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNamespaceVpc(request *UpdateNamespaceVpcRequest) (_result *UpdateNamespaceVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateNamespaceVpcResponse{}
	_body, _err := client.UpdateNamespaceVpcWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNamespaceVpcWithOptions(request *UpdateNamespaceVpcRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateNamespaceVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNamespaceVpc"),
		Version:     tea.String("2019-05-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/v1/sam/namespace/updateNamespaceVpc"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNamespaceVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
