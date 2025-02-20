// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConfirmDjbhReportRequest struct {
	// Primary key ID of the report.
	//
	// example:
	//
	// 24563
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ConfirmDjbhReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDjbhReportRequest) GoString() string {
	return s.String()
}

func (s *ConfirmDjbhReportRequest) SetId(v int64) *ConfirmDjbhReportRequest {
	s.Id = &v
	return s
}

type ConfirmDjbhReportResponseBody struct {
	// API response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EF801DD1-D934-51B3-92D4-776CE17B184F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfirmDjbhReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDjbhReportResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmDjbhReportResponseBody) SetCode(v string) *ConfirmDjbhReportResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmDjbhReportResponseBody) SetHttpStatusCode(v string) *ConfirmDjbhReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfirmDjbhReportResponseBody) SetMessage(v string) *ConfirmDjbhReportResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmDjbhReportResponseBody) SetRequestId(v string) *ConfirmDjbhReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmDjbhReportResponseBody) SetSuccess(v string) *ConfirmDjbhReportResponseBody {
	s.Success = &v
	return s
}

type ConfirmDjbhReportResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmDjbhReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmDjbhReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDjbhReportResponse) GoString() string {
	return s.String()
}

func (s *ConfirmDjbhReportResponse) SetHeaders(v map[string]*string) *ConfirmDjbhReportResponse {
	s.Headers = v
	return s
}

func (s *ConfirmDjbhReportResponse) SetStatusCode(v int32) *ConfirmDjbhReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmDjbhReportResponse) SetBody(v *ConfirmDjbhReportResponseBody) *ConfirmDjbhReportResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	// Language.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetLang(v string) *CreateServiceLinkedRoleRequest {
	s.Lang = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CreateServiceLinkedRoleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 592B80F0-7674-56A4-9027-8A0A9ACDBD56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateServiceWorkOrderRequest struct {
	// Creator.
	//
	// This parameter is required.
	//
	// example:
	//
	// 426556
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Customer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1477832102462645
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// Duration in days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DurationDay *string `json:"DurationDay,omitempty" xml:"DurationDay,omitempty"`
	// Attachment requirement.
	//
	// This parameter is required.
	//
	// example:
	//
	// Y
	IsAttachment *string `json:"IsAttachment,omitempty" xml:"IsAttachment,omitempty"`
	// Is milestone.
	//
	// example:
	//
	// Y
	IsMilestone *string `json:"IsMilestone,omitempty" xml:"IsMilestone,omitempty"`
	// Whether a reminder is needed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Y
	IsWorkOrderNotify *string `json:"IsWorkOrderNotify,omitempty" xml:"IsWorkOrderNotify,omitempty"`
	// Number of days for advance notification.
	//
	// example:
	//
	// 5
	NotifyDay *string `json:"NotifyDay,omitempty" xml:"NotifyDay,omitempty"`
	// Notification ID.
	//
	// example:
	//
	// 10
	NotifyId *int64 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// Operation remarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// newly built
	OperateRemark *string `json:"OperateRemark,omitempty" xml:"OperateRemark,omitempty"`
	// Operation type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREATE
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// Operator.
	//
	// This parameter is required.
	//
	// example:
	//
	// 426556
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1734788109000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Work order details.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"questionDetail":"测试工单","answerDetail":""}
	WorkOrderDetail *string `json:"WorkOrderDetail,omitempty" xml:"WorkOrderDetail,omitempty"`
	// Work order name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Delivery task of safety monthly report
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
	// Work order source.
	//
	// This parameter is required.
	//
	// example:
	//
	// Work order migration
	WorkOrderSource *string `json:"WorkOrderSource,omitempty" xml:"WorkOrderSource,omitempty"`
	// Work order status.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNPROCESSED
	WorkOrderStatus *string `json:"WorkOrderStatus,omitempty" xml:"WorkOrderStatus,omitempty"`
	// Work order type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MONTH_REPORT
	WorkOrderType *string `json:"WorkOrderType,omitempty" xml:"WorkOrderType,omitempty"`
}

func (s CreateServiceWorkOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceWorkOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderRequest) SetCreator(v string) *CreateServiceWorkOrderRequest {
	s.Creator = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetCustomerId(v string) *CreateServiceWorkOrderRequest {
	s.CustomerId = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetDurationDay(v string) *CreateServiceWorkOrderRequest {
	s.DurationDay = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetIsAttachment(v string) *CreateServiceWorkOrderRequest {
	s.IsAttachment = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetIsMilestone(v string) *CreateServiceWorkOrderRequest {
	s.IsMilestone = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetIsWorkOrderNotify(v string) *CreateServiceWorkOrderRequest {
	s.IsWorkOrderNotify = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetNotifyDay(v string) *CreateServiceWorkOrderRequest {
	s.NotifyDay = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetNotifyId(v int64) *CreateServiceWorkOrderRequest {
	s.NotifyId = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOperateRemark(v string) *CreateServiceWorkOrderRequest {
	s.OperateRemark = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOperateType(v string) *CreateServiceWorkOrderRequest {
	s.OperateType = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOperator(v string) *CreateServiceWorkOrderRequest {
	s.Operator = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOwnerId(v string) *CreateServiceWorkOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetStartTime(v int64) *CreateServiceWorkOrderRequest {
	s.StartTime = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderDetail(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderDetail = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderName(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderName = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderSource(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderSource = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderStatus(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderStatus = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderType(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderType = &v
	return s
}

type CreateServiceWorkOrderResponseBody struct {
	// Interface status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *CreateServiceWorkOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7DC44321-7AAE-51CD-8E5F-CEB968569042
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceWorkOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceWorkOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderResponseBody) SetCode(v string) *CreateServiceWorkOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetData(v *CreateServiceWorkOrderResponseBodyData) *CreateServiceWorkOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetHttpStatusCode(v int32) *CreateServiceWorkOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetMessage(v string) *CreateServiceWorkOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetRequestId(v string) *CreateServiceWorkOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetSuccess(v bool) *CreateServiceWorkOrderResponseBody {
	s.Success = &v
	return s
}

type CreateServiceWorkOrderResponseBodyData struct {
	// Completion time.
	//
	// example:
	//
	// 1734788109000
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1734788109000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Creator.
	//
	// example:
	//
	// 426556
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Customer ID.
	//
	// example:
	//
	// 1477832102462645
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// End time.
	//
	// example:
	//
	// 1734788109000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Primary key.
	//
	// example:
	//
	// 1978941
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Owner.
	//
	// example:
	//
	// 426556
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Start time.
	//
	// example:
	//
	// 1734788109000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Work order details.
	//
	// example:
	//
	// {"questionDetail":"测试工单","answerDetail":""}
	WorkOrderDetail *string `json:"WorkOrderDetail,omitempty" xml:"WorkOrderDetail,omitempty"`
	// Work order name.
	//
	// example:
	//
	// Delivery task of safety monthly report
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
	// Work order source.
	//
	// example:
	//
	// Work order migration
	WorkOrderSource *string `json:"WorkOrderSource,omitempty" xml:"WorkOrderSource,omitempty"`
	// Work order status.
	//
	// example:
	//
	// UNPROCESSED
	WorkOrderStatus *string `json:"WorkOrderStatus,omitempty" xml:"WorkOrderStatus,omitempty"`
	// Work order type.
	//
	// example:
	//
	// MONTH_REPORT
	WorkOrderType *string `json:"WorkOrderType,omitempty" xml:"WorkOrderType,omitempty"`
}

func (s CreateServiceWorkOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceWorkOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCompleteTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.CompleteTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCreateTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCreator(v string) *CreateServiceWorkOrderResponseBodyData {
	s.Creator = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCustomerId(v string) *CreateServiceWorkOrderResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetEndTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetId(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetOwnerId(v string) *CreateServiceWorkOrderResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetStartTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderDetail(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderDetail = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderName(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderName = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderSource(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderSource = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderStatus(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderStatus = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderType(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderType = &v
	return s
}

type CreateServiceWorkOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceWorkOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceWorkOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceWorkOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderResponse) SetHeaders(v map[string]*string) *CreateServiceWorkOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceWorkOrderResponse) SetStatusCode(v int32) *CreateServiceWorkOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceWorkOrderResponse) SetBody(v *CreateServiceWorkOrderResponseBody) *CreateServiceWorkOrderResponse {
	s.Body = v
	return s
}

type DeleteDjbhReportRequest struct {
	// Primary key ID of the report.
	//
	// This parameter is required.
	//
	// example:
	//
	// 26579
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDjbhReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDjbhReportRequest) GoString() string {
	return s.String()
}

func (s *DeleteDjbhReportRequest) SetId(v int64) *DeleteDjbhReportRequest {
	s.Id = &v
	return s
}

type DeleteDjbhReportResponseBody struct {
	// API response code.
	//
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 86786E4C-6416-55CF-9AB6-5E275B68801D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDjbhReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDjbhReportResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDjbhReportResponseBody) SetCode(v string) *DeleteDjbhReportResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDjbhReportResponseBody) SetHttpStatusCode(v string) *DeleteDjbhReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDjbhReportResponseBody) SetMessage(v string) *DeleteDjbhReportResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDjbhReportResponseBody) SetRequestId(v string) *DeleteDjbhReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDjbhReportResponseBody) SetSuccess(v string) *DeleteDjbhReportResponseBody {
	s.Success = &v
	return s
}

type DeleteDjbhReportResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDjbhReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDjbhReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDjbhReportResponse) GoString() string {
	return s.String()
}

func (s *DeleteDjbhReportResponse) SetHeaders(v map[string]*string) *DeleteDjbhReportResponse {
	s.Headers = v
	return s
}

func (s *DeleteDjbhReportResponse) SetStatusCode(v int32) *DeleteDjbhReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDjbhReportResponse) SetBody(v *DeleteDjbhReportResponseBody) *DeleteDjbhReportResponse {
	s.Body = v
	return s
}

type DescribeServiceLinkedRoleRequest struct {
	// Language.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleRequest) SetLang(v string) *DescribeServiceLinkedRoleRequest {
	s.Lang = &v
	return s
}

func (s *DescribeServiceLinkedRoleRequest) SetRegionId(v string) *DescribeServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type DescribeServiceLinkedRoleResponseBody struct {
	// Whether the service-linked role permission is granted:
	//
	// - true: Granted.
	//
	// - false: Not granted.
	//
	// example:
	//
	// true
	EntityRoleGrant *bool `json:"EntityRoleGrant,omitempty" xml:"EntityRoleGrant,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 02F8BBF3-2D61-5982-8911-EEB387BE3AF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleResponseBody) SetEntityRoleGrant(v bool) *DescribeServiceLinkedRoleResponseBody {
	s.EntityRoleGrant = &v
	return s
}

func (s *DescribeServiceLinkedRoleResponseBody) SetRequestId(v string) *DescribeServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type DescribeServiceLinkedRoleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *DescribeServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceLinkedRoleResponse) SetStatusCode(v int32) *DescribeServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceLinkedRoleResponse) SetBody(v *DescribeServiceLinkedRoleResponseBody) *DescribeServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type DisposeServiceWorkOrderRequest struct {
	// Attachment name.
	//
	// example:
	//
	// bbaa133c-0ac2-489f-9fc8-39f91c2e770c_20230301-20240403-服务工单列表.xlsx
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// End time.
	//
	// example:
	//
	// 2024-04-14 00:00:00
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Forward to owner.
	//
	// example:
	//
	// 405639
	ForwardOwnerId *string `json:"ForwardOwnerId,omitempty" xml:"ForwardOwnerId,omitempty"`
	// Work order ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23172
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Attachment requirement.
	//
	// example:
	//
	// Y
	IsAttachment *string `json:"IsAttachment,omitempty" xml:"IsAttachment,omitempty"`
	// Work order notification.
	//
	// example:
	//
	// Y
	IsWorkOrderNotify *string `json:"IsWorkOrderNotify,omitempty" xml:"IsWorkOrderNotify,omitempty"`
	// Notification ID.
	//
	// example:
	//
	// 10
	NotifyId *int64 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// Operation remarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 处理完成
	OperateRemark *string `json:"OperateRemark,omitempty" xml:"OperateRemark,omitempty"`
	// Processing type.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROCESSED
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// Operator.
	//
	// This parameter is required.
	//
	// example:
	//
	// 396120
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Start time.
	//
	// example:
	//
	// 2024-04-02 00:00:00
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Upgrade owner.
	//
	// example:
	//
	// 336333
	UpgradeOwnerId *string `json:"UpgradeOwnerId,omitempty" xml:"UpgradeOwnerId,omitempty"`
	// Work order details.
	//
	// example:
	//
	// {"questionDetail":"测试工单","answerDetail":""}
	WorkOrderDetail *string `json:"WorkOrderDetail,omitempty" xml:"WorkOrderDetail,omitempty"`
	// Work order name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 安全产品配置问题与超量提醒
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
	// Work order status.
	//
	// example:
	//
	// PROCESSED
	WorkOrderStatus *string `json:"WorkOrderStatus,omitempty" xml:"WorkOrderStatus,omitempty"`
}

func (s DisposeServiceWorkOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DisposeServiceWorkOrderRequest) GoString() string {
	return s.String()
}

func (s *DisposeServiceWorkOrderRequest) SetAttachmentName(v string) *DisposeServiceWorkOrderRequest {
	s.AttachmentName = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetEndTime(v int64) *DisposeServiceWorkOrderRequest {
	s.EndTime = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetForwardOwnerId(v string) *DisposeServiceWorkOrderRequest {
	s.ForwardOwnerId = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetId(v int64) *DisposeServiceWorkOrderRequest {
	s.Id = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetIsAttachment(v string) *DisposeServiceWorkOrderRequest {
	s.IsAttachment = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetIsWorkOrderNotify(v string) *DisposeServiceWorkOrderRequest {
	s.IsWorkOrderNotify = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetNotifyId(v int64) *DisposeServiceWorkOrderRequest {
	s.NotifyId = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetOperateRemark(v string) *DisposeServiceWorkOrderRequest {
	s.OperateRemark = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetOperateType(v string) *DisposeServiceWorkOrderRequest {
	s.OperateType = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetOperator(v string) *DisposeServiceWorkOrderRequest {
	s.Operator = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetStartTime(v int64) *DisposeServiceWorkOrderRequest {
	s.StartTime = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetUpgradeOwnerId(v string) *DisposeServiceWorkOrderRequest {
	s.UpgradeOwnerId = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetWorkOrderDetail(v string) *DisposeServiceWorkOrderRequest {
	s.WorkOrderDetail = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetWorkOrderName(v string) *DisposeServiceWorkOrderRequest {
	s.WorkOrderName = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetWorkOrderStatus(v string) *DisposeServiceWorkOrderRequest {
	s.WorkOrderStatus = &v
	return s
}

type DisposeServiceWorkOrderResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ED520610-6231-5D80-BADD-A8CDC7BBC809
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisposeServiceWorkOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisposeServiceWorkOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DisposeServiceWorkOrderResponseBody) SetCode(v string) *DisposeServiceWorkOrderResponseBody {
	s.Code = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetHttpStatusCode(v int32) *DisposeServiceWorkOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetMessage(v string) *DisposeServiceWorkOrderResponseBody {
	s.Message = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetRequestId(v string) *DisposeServiceWorkOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetSuccess(v bool) *DisposeServiceWorkOrderResponseBody {
	s.Success = &v
	return s
}

type DisposeServiceWorkOrderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisposeServiceWorkOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisposeServiceWorkOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DisposeServiceWorkOrderResponse) GoString() string {
	return s.String()
}

func (s *DisposeServiceWorkOrderResponse) SetHeaders(v map[string]*string) *DisposeServiceWorkOrderResponse {
	s.Headers = v
	return s
}

func (s *DisposeServiceWorkOrderResponse) SetStatusCode(v int32) *DisposeServiceWorkOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DisposeServiceWorkOrderResponse) SetBody(v *DisposeServiceWorkOrderResponseBody) *DisposeServiceWorkOrderResponse {
	s.Body = v
	return s
}

type DisposeWorkTaskRequest struct {
	// Operator.
	//
	// This parameter is required.
	//
	// example:
	//
	// WB01089929
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Operation remarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 处理完成
	OptRemark *string `json:"OptRemark,omitempty" xml:"OptRemark,omitempty"`
	// Work order status.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Work order ID, multiple IDs separated by commas.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10310
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s DisposeWorkTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DisposeWorkTaskRequest) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskRequest) SetOperator(v string) *DisposeWorkTaskRequest {
	s.Operator = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetOptRemark(v string) *DisposeWorkTaskRequest {
	s.OptRemark = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetStatus(v int32) *DisposeWorkTaskRequest {
	s.Status = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetTaskIds(v string) *DisposeWorkTaskRequest {
	s.TaskIds = &v
	return s
}

type DisposeWorkTaskResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86786E4C-6416-55CF-9AB6-5E275B68801D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisposeWorkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisposeWorkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskResponseBody) SetCode(v string) *DisposeWorkTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetHttpStatusCode(v int32) *DisposeWorkTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetMessage(v string) *DisposeWorkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetRequestId(v string) *DisposeWorkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetSuccess(v bool) *DisposeWorkTaskResponseBody {
	s.Success = &v
	return s
}

type DisposeWorkTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisposeWorkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisposeWorkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DisposeWorkTaskResponse) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskResponse) SetHeaders(v map[string]*string) *DisposeWorkTaskResponse {
	s.Headers = v
	return s
}

func (s *DisposeWorkTaskResponse) SetStatusCode(v int32) *DisposeWorkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DisposeWorkTaskResponse) SetBody(v *DisposeWorkTaskResponseBody) *DisposeWorkTaskResponse {
	s.Body = v
	return s
}

type GetAlarmDetailByIdRequest struct {
	// Primary key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20077761
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetAlarmDetailByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdRequest) SetId(v int64) *GetAlarmDetailByIdRequest {
	s.Id = &v
	return s
}

type GetAlarmDetailByIdResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetAlarmDetailByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5C1B0668-442C-57AE-9668-D894B0B012EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful: - true: Success. - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAlarmDetailByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdResponseBody) SetCode(v string) *GetAlarmDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetData(v *GetAlarmDetailByIdResponseBodyData) *GetAlarmDetailByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetHttpStatusCode(v int32) *GetAlarmDetailByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetMessage(v string) *GetAlarmDetailByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetRequestId(v string) *GetAlarmDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBody) SetSuccess(v bool) *GetAlarmDetailByIdResponseBody {
	s.Success = &v
	return s
}

type GetAlarmDetailByIdResponseBodyData struct {
	// Alarm event type.
	//
	// example:
	//
	// Unusual Logon
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// Alarm event type.
	//
	// example:
	//
	// Login with unusual location
	AlarmEventTypeDisplay *string `json:"AlarmEventTypeDisplay,omitempty" xml:"AlarmEventTypeDisplay,omitempty"`
	// Alarm ID.
	//
	// example:
	//
	// 202427220
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// Alarm name.
	//
	// example:
	//
	// 负载均衡可挂载服务器数量告警
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// Alarm source.
	//
	// example:
	//
	// SUSP_EVENT
	AlarmSource *string `json:"AlarmSource,omitempty" xml:"AlarmSource,omitempty"`
	// Latest alarm time.
	//
	// example:
	//
	// 2018-09-26 01:51:01
	AlarmTime *string `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	// Analysis process.
	//
	// example:
	//
	// [{"value":"服务器可能已被黑客攻击，存在恶意进程在运行。 分析过程：告警显示，服务端存在一个名为”dns.exe”的进程在访问”polling.burpcollaborator.net”，这是一个被黑名单列出的恶意域名。在正常情况下,”dns.exe”不应该单独存在于系统的路径下，并且也不应该访问这类恶意域名。因此，这个进程可能是黑客留下的恶意进程。","key":"结论"},{"value":"尽快对服务器进行全面扫描，清除恶意进程。同时，联系网络安全专家进行深入调查，以确定是否有其他潜在的安全威胁。","key":"处置建议"}]
	AnalysisResult *string `json:"AnalysisResult,omitempty" xml:"AnalysisResult,omitempty"`
	// Whether high-protection mode is enabled. true means enabled, false means not enabled.
	//
	// example:
	//
	// false
	ContainHwMode *bool `json:"ContainHwMode,omitempty" xml:"ContainHwMode,omitempty"`
	// Alarm handling time.
	//
	// example:
	//
	// 2018-09-26 01:51:01
	DealTime *string `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// Description.
	//
	// example:
	//
	// webshell
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// Event details information.
	EventDetails []*GetAlarmDetailByIdResponseBodyDataEventDetails `json:"EventDetails,omitempty" xml:"EventDetails,omitempty" type:"Repeated"`
	// Alarm level.
	//
	// example:
	//
	// suspicious
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// Primary key ID of the work order.
	//
	// example:
	//
	// 9772
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Affected asset.
	//
	// example:
	//
	// nginx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Public IP.
	//
	// example:
	//
	// 47.116.126.79
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Private IP.
	//
	// example:
	//
	// 172.19.195.176
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// First occurrence time
	//
	// example:
	//
	// 2018-09-26 01:51:01
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	// Owner.
	//
	// example:
	//
	// 324546
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Disposal method.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Handling status.
	//
	// example:
	//
	// 要查询的告警事件状态。取值：
	//
	// 0：全部
	//
	// 1：待处理
	//
	// 2：已忽略
	//
	// 4：已确认
	//
	// 8：已标记为误报
	//
	// 16：处理中
	//
	// 32：处理完毕
	//
	// 64：已经过期
	//
	// 128：已经删除
	//
	// 512：自动拦截中
	//
	// 513：自动拦截完毕
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// ATT&CK tactic name.
	//
	// example:
	//
	// Malicious scripts-Malicious script code execution
	TacticDisplayName *string `json:"TacticDisplayName,omitempty" xml:"TacticDisplayName,omitempty"`
}

func (s GetAlarmDetailByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmDetailByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmEventType(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmEventType = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmEventTypeDisplay(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmEventTypeDisplay = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmId(v int64) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmId = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmName(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmName = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmSource(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmSource = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAlarmTime(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AlarmTime = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetAnalysisResult(v string) *GetAlarmDetailByIdResponseBodyData {
	s.AnalysisResult = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetContainHwMode(v bool) *GetAlarmDetailByIdResponseBodyData {
	s.ContainHwMode = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetDealTime(v string) *GetAlarmDetailByIdResponseBodyData {
	s.DealTime = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetDesc(v string) *GetAlarmDetailByIdResponseBodyData {
	s.Desc = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetEventDetails(v []*GetAlarmDetailByIdResponseBodyDataEventDetails) *GetAlarmDetailByIdResponseBodyData {
	s.EventDetails = v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetEventLevel(v string) *GetAlarmDetailByIdResponseBodyData {
	s.EventLevel = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetId(v int64) *GetAlarmDetailByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetInstanceName(v string) *GetAlarmDetailByIdResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetInternetIp(v string) *GetAlarmDetailByIdResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetIntranetIp(v string) *GetAlarmDetailByIdResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetOccurrenceTime(v string) *GetAlarmDetailByIdResponseBodyData {
	s.OccurrenceTime = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetOwnerId(v string) *GetAlarmDetailByIdResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetRemark(v string) *GetAlarmDetailByIdResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetStatus(v string) *GetAlarmDetailByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyData) SetTacticDisplayName(v string) *GetAlarmDetailByIdResponseBodyData {
	s.TacticDisplayName = &v
	return s
}

type GetAlarmDetailByIdResponseBodyDataEventDetails struct {
	// Alarm event display name.
	//
	// example:
	//
	// Login with unusual location
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// Alarm event type.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Path where the alarm event occurred.
	//
	// example:
	//
	// /etc/crontab
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// Path where the alarm event occurred.
	//
	// example:
	//
	// /etc/crontab
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s GetAlarmDetailByIdResponseBodyDataEventDetails) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmDetailByIdResponseBodyDataEventDetails) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) SetNameDisplay(v string) *GetAlarmDetailByIdResponseBodyDataEventDetails {
	s.NameDisplay = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) SetType(v string) *GetAlarmDetailByIdResponseBodyDataEventDetails {
	s.Type = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) SetValue(v string) *GetAlarmDetailByIdResponseBodyDataEventDetails {
	s.Value = &v
	return s
}

func (s *GetAlarmDetailByIdResponseBodyDataEventDetails) SetValueDisplay(v string) *GetAlarmDetailByIdResponseBodyDataEventDetails {
	s.ValueDisplay = &v
	return s
}

type GetAlarmDetailByIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlarmDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlarmDetailByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlarmDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetAlarmDetailByIdResponse) SetHeaders(v map[string]*string) *GetAlarmDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetAlarmDetailByIdResponse) SetStatusCode(v int32) *GetAlarmDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlarmDetailByIdResponse) SetBody(v *GetAlarmDetailByIdResponseBody) *GetAlarmDetailByIdResponse {
	s.Body = v
	return s
}

type GetAttackedAssetDealRequest struct {
	// Time filter type, supporting filtering by the last 7 days, the last 30 days, the last half year, or custom time periods.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732268720000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732268720000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Source of the alert event.
	//
	// example:
	//
	// 暂时无需传参，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetAttackedAssetDealRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealRequest) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealRequest) SetDateType(v string) *GetAttackedAssetDealRequest {
	s.DateType = &v
	return s
}

func (s *GetAttackedAssetDealRequest) SetEndDate(v int64) *GetAttackedAssetDealRequest {
	s.EndDate = &v
	return s
}

func (s *GetAttackedAssetDealRequest) SetStartDate(v int64) *GetAttackedAssetDealRequest {
	s.StartDate = &v
	return s
}

func (s *GetAttackedAssetDealRequest) SetSuspEventSource(v string) *GetAttackedAssetDealRequest {
	s.SuspEventSource = &v
	return s
}

type GetAttackedAssetDealResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetAttackedAssetDealResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1E74F11C-B4A8-5774-962C-02003BA8504E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the query was successful.<br />
	//
	// **Enum values:**
	//
	// 	- true: Success.
	//
	// 	- false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAttackedAssetDealResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponseBody) SetCode(v string) *GetAttackedAssetDealResponseBody {
	s.Code = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetData(v *GetAttackedAssetDealResponseBodyData) *GetAttackedAssetDealResponseBody {
	s.Data = v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetHttpStatusCode(v int32) *GetAttackedAssetDealResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetMessage(v string) *GetAttackedAssetDealResponseBody {
	s.Message = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetRequestId(v string) *GetAttackedAssetDealResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetSuccess(v bool) *GetAttackedAssetDealResponseBody {
	s.Success = &v
	return s
}

type GetAttackedAssetDealResponseBodyData struct {
	// Collection of attacked asset convergence trends.
	EcsTrendList []*GetAttackedAssetDealResponseBodyDataEcsTrendList `json:"EcsTrendList,omitempty" xml:"EcsTrendList,omitempty" type:"Repeated"`
}

func (s GetAttackedAssetDealResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponseBodyData) SetEcsTrendList(v []*GetAttackedAssetDealResponseBodyDataEcsTrendList) *GetAttackedAssetDealResponseBodyData {
	s.EcsTrendList = v
	return s
}

type GetAttackedAssetDealResponseBodyDataEcsTrendList struct {
	// Date point.
	//
	// example:
	//
	// 202312或20231205
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Number of processed items.
	//
	// example:
	//
	// 2
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// Number of discovered items.
	//
	// example:
	//
	// 暂时无值，有疑问请联系管理员
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetAttackedAssetDealResponseBodyDataEcsTrendList) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealResponseBodyDataEcsTrendList) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) SetDate(v string) *GetAttackedAssetDealResponseBodyDataEcsTrendList {
	s.Date = &v
	return s
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) SetDealCount(v int64) *GetAttackedAssetDealResponseBodyDataEcsTrendList {
	s.DealCount = &v
	return s
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) SetFindCount(v int64) *GetAttackedAssetDealResponseBodyDataEcsTrendList {
	s.FindCount = &v
	return s
}

type GetAttackedAssetDealResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackedAssetDealResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackedAssetDealResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealResponse) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponse) SetHeaders(v map[string]*string) *GetAttackedAssetDealResponse {
	s.Headers = v
	return s
}

func (s *GetAttackedAssetDealResponse) SetStatusCode(v int32) *GetAttackedAssetDealResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackedAssetDealResponse) SetBody(v *GetAttackedAssetDealResponseBody) *GetAttackedAssetDealResponse {
	s.Body = v
	return s
}

type GetBaselineSummaryRequest struct {
	// Time filter type, supports filtering by the last 7 days, the last 30 days, the last half year, or custom time periods.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetBaselineSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryRequest) SetDateType(v string) *GetBaselineSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetBaselineSummaryRequest) SetEndDate(v int64) *GetBaselineSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetBaselineSummaryRequest) SetStartDate(v int64) *GetBaselineSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetBaselineSummaryRequest) SetSuspEventSource(v string) *GetBaselineSummaryRequest {
	s.SuspEventSource = &v
	return s
}

type GetBaselineSummaryResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetBaselineSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 67D61738-5E38-5164-947A-34E3850D493A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Values: true: success; false: failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBaselineSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponseBody) SetCode(v string) *GetBaselineSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetData(v *GetBaselineSummaryResponseBodyData) *GetBaselineSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetHttpStatusCode(v int32) *GetBaselineSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetMessage(v string) *GetBaselineSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetRequestId(v string) *GetBaselineSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetSuccess(v bool) *GetBaselineSummaryResponseBody {
	s.Success = &v
	return s
}

type GetBaselineSummaryResponseBodyData struct {
	// Collection of baseline statistical data.
	TrendDTOList []*GetBaselineSummaryResponseBodyDataTrendDTOList `json:"TrendDTOList,omitempty" xml:"TrendDTOList,omitempty" type:"Repeated"`
}

func (s GetBaselineSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponseBodyData) SetTrendDTOList(v []*GetBaselineSummaryResponseBodyDataTrendDTOList) *GetBaselineSummaryResponseBodyData {
	s.TrendDTOList = v
	return s
}

type GetBaselineSummaryResponseBodyDataTrendDTOList struct {
	// Date point.
	//
	// example:
	//
	// 202408或者20240801
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Number of processed items.
	//
	// example:
	//
	// 10
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// Number of discovered items.
	//
	// example:
	//
	// 12
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetBaselineSummaryResponseBodyDataTrendDTOList) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryResponseBodyDataTrendDTOList) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) SetDate(v string) *GetBaselineSummaryResponseBodyDataTrendDTOList {
	s.Date = &v
	return s
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) SetDealCount(v int64) *GetBaselineSummaryResponseBodyDataTrendDTOList {
	s.DealCount = &v
	return s
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) SetFindCount(v int64) *GetBaselineSummaryResponseBodyDataTrendDTOList {
	s.FindCount = &v
	return s
}

type GetBaselineSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBaselineSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBaselineSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponse) SetHeaders(v map[string]*string) *GetBaselineSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetBaselineSummaryResponse) SetStatusCode(v int32) *GetBaselineSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBaselineSummaryResponse) SetBody(v *GetBaselineSummaryResponseBody) *GetBaselineSummaryResponse {
	s.Body = v
	return s
}

type GetConsoleScoreRequest struct {
	// Filter time type, supports filtering by the last 7 days, last 30 days, last half year, or custom. If empty, it represents the last 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Source of alert events.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetConsoleScoreRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConsoleScoreRequest) GoString() string {
	return s.String()
}

func (s *GetConsoleScoreRequest) SetDateType(v string) *GetConsoleScoreRequest {
	s.DateType = &v
	return s
}

func (s *GetConsoleScoreRequest) SetEndDate(v int64) *GetConsoleScoreRequest {
	s.EndDate = &v
	return s
}

func (s *GetConsoleScoreRequest) SetStartDate(v int64) *GetConsoleScoreRequest {
	s.StartDate = &v
	return s
}

func (s *GetConsoleScoreRequest) SetSuspEventSource(v string) *GetConsoleScoreRequest {
	s.SuspEventSource = &v
	return s
}

type GetConsoleScoreResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	//
	// example:
	//
	// {
	//
	//     "score": "94.00",
	//
	//     "consoleScoreTrendDTOS": [
	//
	//         {
	//
	//             "date": "20241009",
	//
	//             "score": "100.0"
	//
	//         }
	//
	//     ],
	//
	//     "cyclicYearOverYear": "-6.00",
	//
	//     "recordDate": "20241209",
	//
	//     "weeklyYearOverYear": "1.62",
	//
	//     "aboveWholeNetworkUserRatio": "6.25",
	//
	//     "aliUid": "1601097845544644",
	//
	//     "detailJson": "[{\\"detailDTO\\":[{\\"count\\":0,\\"itemName\\":\\"应用漏洞POC验证\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"未授权访问漏洞（公网暴露）\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"后台弱口令漏洞（公网暴露）\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"文件上传漏洞（公网暴露）\\",\\"mark\\":\\"1\\"}],\\"markRate\\":\\"0.5\\",\\"markType\\":\\"vul\\"},{\\"detailDTO\\":[{\\"count\\":12,\\"itemName\\":\\"WAF3.0回源配置不正确\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"AK泄露检查未开启\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"DNAT管理端口开放\\",\\"mark\\":\\"1\\"},{\\"count\\":0,\\"itemName\\":\\"高危端口暴露\\",\\"mark\\":\\"0.5\\"}],\\"markRate\\":\\"0.5\\",\\"markType\\":\\"risk\\"}]"
	//
	// }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the result returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// D0937B0F-9180-5F70-B6ED-0BA22591627F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. true means success, false means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConsoleScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsoleScoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsoleScoreResponseBody) SetCode(v string) *GetConsoleScoreResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsoleScoreResponseBody) SetData(v interface{}) *GetConsoleScoreResponseBody {
	s.Data = v
	return s
}

func (s *GetConsoleScoreResponseBody) SetHttpStatusCode(v int32) *GetConsoleScoreResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsoleScoreResponseBody) SetMessage(v string) *GetConsoleScoreResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsoleScoreResponseBody) SetRequestId(v string) *GetConsoleScoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsoleScoreResponseBody) SetSuccess(v bool) *GetConsoleScoreResponseBody {
	s.Success = &v
	return s
}

type GetConsoleScoreResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsoleScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsoleScoreResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsoleScoreResponse) GoString() string {
	return s.String()
}

func (s *GetConsoleScoreResponse) SetHeaders(v map[string]*string) *GetConsoleScoreResponse {
	s.Headers = v
	return s
}

func (s *GetConsoleScoreResponse) SetStatusCode(v int32) *GetConsoleScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsoleScoreResponse) SetBody(v *GetConsoleScoreResponseBody) *GetConsoleScoreResponse {
	s.Body = v
	return s
}

type GetDetailByIdRequest struct {
	// Primary key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetailByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetDetailByIdRequest) SetId(v int64) *GetDetailByIdRequest {
	s.Id = &v
	return s
}

type GetDetailByIdResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// 404
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetDetailByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAB46EC5-3746-59C4-B6D2-469F442EC73F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Values: - **true**: indicates a successful call. - **false**: indicates a failed call.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDetailByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponseBody) SetCode(v string) *GetDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetData(v *GetDetailByIdResponseBodyData) *GetDetailByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetDetailByIdResponseBody) SetHttpStatusCode(v int32) *GetDetailByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetMessage(v string) *GetDetailByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetRequestId(v string) *GetDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetSuccess(v bool) *GetDetailByIdResponseBody {
	s.Success = &v
	return s
}

type GetDetailByIdResponseBodyData struct {
	// Vulnerability details.
	VulDetails []*GetDetailByIdResponseBodyDataVulDetails `json:"VulDetails,omitempty" xml:"VulDetails,omitempty" type:"Repeated"`
}

func (s GetDetailByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponseBodyData) SetVulDetails(v []*GetDetailByIdResponseBodyDataVulDetails) *GetDetailByIdResponseBodyData {
	s.VulDetails = v
	return s
}

type GetDetailByIdResponseBodyDataVulDetails struct {
	// CVE ID.
	//
	// example:
	//
	// CVE-2022-21291
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The CVSS score of the vulnerability in the Alibaba Cloud vulnerability database.
	//
	// example:
	//
	// 10.0
	CvssScore *string `json:"CvssScore,omitempty" xml:"CvssScore,omitempty"`
	// Fix suggestion.
	//
	// example:
	//
	// https://avd.aliyun.com/detail/CVE-2022-21291
	FixSuggestion *string `json:"FixSuggestion,omitempty" xml:"FixSuggestion,omitempty"`
	// Title of the vulnerability announcement.
	//
	// example:
	//
	// Chanjet T-Plus SetupAccount/Upload. Aspx file upload vulnerability(CNVD-2022-60632)
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetDetailByIdResponseBodyDataVulDetails) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdResponseBodyDataVulDetails) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetCveId(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.CveId = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetCvssScore(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.CvssScore = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetFixSuggestion(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.FixSuggestion = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetTitle(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.Title = &v
	return s
}

type GetDetailByIdResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetailByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponse) SetHeaders(v map[string]*string) *GetDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetDetailByIdResponse) SetStatusCode(v int32) *GetDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetailByIdResponse) SetBody(v *GetDetailByIdResponseBody) *GetDetailByIdResponse {
	s.Body = v
	return s
}

type GetDocumentDownloadUrlRequest struct {
	// Document management ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 175815
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Report type.
	//
	// example:
	//
	// 5
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetDocumentDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlRequest) SetId(v int64) *GetDocumentDownloadUrlRequest {
	s.Id = &v
	return s
}

func (s *GetDocumentDownloadUrlRequest) SetReportType(v string) *GetDocumentDownloadUrlRequest {
	s.ReportType = &v
	return s
}

type GetDocumentDownloadUrlResponseBody struct {
	// API status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// OSS file access URL.
	//
	// example:
	//
	// https://oos-cn.ctyunapi.cn/example-bucket/test/1.jpg
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Message of the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// C7BE80B4-7692-54FA-AB22-2A7DF08C4754
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful: - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlResponseBody) SetCode(v string) *GetDocumentDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetData(v string) *GetDocumentDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetHttpStatusCode(v int32) *GetDocumentDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetMessage(v string) *GetDocumentDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetRequestId(v string) *GetDocumentDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetSuccess(v bool) *GetDocumentDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type GetDocumentDownloadUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlResponse) SetHeaders(v map[string]*string) *GetDocumentDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentDownloadUrlResponse) SetStatusCode(v int32) *GetDocumentDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentDownloadUrlResponse) SetBody(v *GetDocumentDownloadUrlResponseBody) *GetDocumentDownloadUrlResponse {
	s.Body = v
	return s
}

type GetDocumentPageRequest struct {
	// Current page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Delivered by.
	//
	// example:
	//
	// luna
	DeliveredBy *string `json:"DeliveredBy,omitempty" xml:"DeliveredBy,omitempty"`
	// Document name.
	//
	// example:
	//
	// month report
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// Document type.
	//
	// example:
	//
	// 0
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// Page size.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Report type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetDocumentPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentPageRequest) SetCurrentPage(v int32) *GetDocumentPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentPageRequest) SetDeliveredBy(v string) *GetDocumentPageRequest {
	s.DeliveredBy = &v
	return s
}

func (s *GetDocumentPageRequest) SetDocumentName(v string) *GetDocumentPageRequest {
	s.DocumentName = &v
	return s
}

func (s *GetDocumentPageRequest) SetDocumentType(v string) *GetDocumentPageRequest {
	s.DocumentType = &v
	return s
}

func (s *GetDocumentPageRequest) SetPageSize(v int32) *GetDocumentPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetDocumentPageRequest) SetReportType(v string) *GetDocumentPageRequest {
	s.ReportType = &v
	return s
}

type GetDocumentPageResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response data.
	Data []*GetDocumentPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *GetDocumentPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04DAD7B4-E1DA-5C2C-8E5C-A1EDC880CF60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponseBody) SetCode(v string) *GetDocumentPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetData(v []*GetDocumentPageResponseBodyData) *GetDocumentPageResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentPageResponseBody) SetHttpStatusCode(v int32) *GetDocumentPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetMessage(v string) *GetDocumentPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetPageInfo(v *GetDocumentPageResponseBodyPageInfo) *GetDocumentPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetDocumentPageResponseBody) SetRequestId(v string) *GetDocumentPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetSuccess(v bool) *GetDocumentPageResponseBody {
	s.Success = &v
	return s
}

type GetDocumentPageResponseBodyData struct {
	// Delivered by.
	//
	// example:
	//
	// luna
	DeliveredBy *string `json:"DeliveredBy,omitempty" xml:"DeliveredBy,omitempty"`
	// Report name.
	//
	// example:
	//
	// month report
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// Service report type.
	//
	// example:
	//
	// 3
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// Primary key ID of the document.
	//
	// example:
	//
	// 346409
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Report status.
	//
	// example:
	//
	// uploaded
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// Report generation time.
	//
	// example:
	//
	// 2023-03-21 17:26:34
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s GetDocumentPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponseBodyData) SetDeliveredBy(v string) *GetDocumentPageResponseBodyData {
	s.DeliveredBy = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetDocumentName(v string) *GetDocumentPageResponseBodyData {
	s.DocumentName = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetDocumentType(v string) *GetDocumentPageResponseBodyData {
	s.DocumentType = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetId(v int64) *GetDocumentPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetReportStatus(v string) *GetDocumentPageResponseBodyData {
	s.ReportStatus = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetUploadTime(v string) *GetDocumentPageResponseBodyData {
	s.UploadTime = &v
	return s
}

type GetDocumentPageResponseBodyPageInfo struct {
	// The current page number in pagination queries.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page in the returned data.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of queried items.
	//
	// example:
	//
	// 3149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDocumentPageResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponseBodyPageInfo) SetCurrentPage(v int32) *GetDocumentPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentPageResponseBodyPageInfo) SetPageSize(v int32) *GetDocumentPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetDocumentPageResponseBodyPageInfo) SetTotalCount(v int32) *GetDocumentPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type GetDocumentPageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponse) SetHeaders(v map[string]*string) *GetDocumentPageResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentPageResponse) SetStatusCode(v int32) *GetDocumentPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentPageResponse) SetBody(v *GetDocumentPageResponseBody) *GetDocumentPageResponse {
	s.Body = v
	return s
}

type GetDocumentSummaryRequest struct {
	// Type of service report.
	//
	// example:
	//
	// 1
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetDocumentSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryRequest) SetReportType(v string) *GetDocumentSummaryRequest {
	s.ReportType = &v
	return s
}

type GetDocumentSummaryResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetDocumentSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7903F2DE-D9EE-5D16-8A08-E9223E54B281
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. Values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryResponseBody) SetCode(v string) *GetDocumentSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetData(v *GetDocumentSummaryResponseBodyData) *GetDocumentSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetHttpStatusCode(v int32) *GetDocumentSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetMessage(v string) *GetDocumentSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetRequestId(v string) *GetDocumentSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetSuccess(v bool) *GetDocumentSummaryResponseBody {
	s.Success = &v
	return s
}

type GetDocumentSummaryResponseBodyData struct {
	// Number of documents.
	//
	// example:
	//
	// 10
	DocumentCount *int64 `json:"DocumentCount,omitempty" xml:"DocumentCount,omitempty"`
	// Number of services or days.
	//
	// example:
	//
	// 10
	Frequency *int64 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
}

func (s GetDocumentSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryResponseBodyData) SetDocumentCount(v int64) *GetDocumentSummaryResponseBodyData {
	s.DocumentCount = &v
	return s
}

func (s *GetDocumentSummaryResponseBodyData) SetFrequency(v int64) *GetDocumentSummaryResponseBodyData {
	s.Frequency = &v
	return s
}

type GetDocumentSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryResponse) SetHeaders(v map[string]*string) *GetDocumentSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentSummaryResponse) SetStatusCode(v int32) *GetDocumentSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentSummaryResponse) SetBody(v *GetDocumentSummaryResponseBody) *GetDocumentSummaryResponse {
	s.Body = v
	return s
}

type GetRecentDocumentRequest struct {
	// Filter time type, supports filtering by the last 7 days, the last 30 days, the last half year, or custom time ranges.
	//
	// This parameter is required.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetRecentDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecentDocumentRequest) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentRequest) SetDateType(v string) *GetRecentDocumentRequest {
	s.DateType = &v
	return s
}

func (s *GetRecentDocumentRequest) SetEndDate(v int64) *GetRecentDocumentRequest {
	s.EndDate = &v
	return s
}

func (s *GetRecentDocumentRequest) SetStartDate(v int64) *GetRecentDocumentRequest {
	s.StartDate = &v
	return s
}

func (s *GetRecentDocumentRequest) SetSuspEventSource(v string) *GetRecentDocumentRequest {
	s.SuspEventSource = &v
	return s
}

type GetRecentDocumentResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetRecentDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4916FA8D-F294-518D-B373-8B59D63CAB19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRecentDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecentDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentResponseBody) SetCode(v string) *GetRecentDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetData(v []*GetRecentDocumentResponseBodyData) *GetRecentDocumentResponseBody {
	s.Data = v
	return s
}

func (s *GetRecentDocumentResponseBody) SetHttpStatusCode(v int32) *GetRecentDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetMessage(v string) *GetRecentDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetRequestId(v string) *GetRecentDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetSuccess(v bool) *GetRecentDocumentResponseBody {
	s.Success = &v
	return s
}

type GetRecentDocumentResponseBodyData struct {
	// Primary key ID.
	//
	// example:
	//
	// 360491
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Document name
	//
	// example:
	//
	// 文档名称测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Upload time.
	//
	// example:
	//
	// 2023-03-20 14:30:38
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s GetRecentDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRecentDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentResponseBodyData) SetId(v int64) *GetRecentDocumentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetRecentDocumentResponseBodyData) SetName(v string) *GetRecentDocumentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetRecentDocumentResponseBodyData) SetUploadTime(v string) *GetRecentDocumentResponseBodyData {
	s.UploadTime = &v
	return s
}

type GetRecentDocumentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecentDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecentDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecentDocumentResponse) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentResponse) SetHeaders(v map[string]*string) *GetRecentDocumentResponse {
	s.Headers = v
	return s
}

func (s *GetRecentDocumentResponse) SetStatusCode(v int32) *GetRecentDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecentDocumentResponse) SetBody(v *GetRecentDocumentResponseBody) *GetRecentDocumentResponse {
	s.Body = v
	return s
}

type GetSafetyCoverRequest struct {
	// Filter time type, supports filtering by the last 7 days, the last 30 days, the last half year, or custom time periods.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732268720000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732255620000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// 该接口不用传
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetSafetyCoverRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverRequest) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverRequest) SetDateType(v string) *GetSafetyCoverRequest {
	s.DateType = &v
	return s
}

func (s *GetSafetyCoverRequest) SetEndDate(v int64) *GetSafetyCoverRequest {
	s.EndDate = &v
	return s
}

func (s *GetSafetyCoverRequest) SetStartDate(v int64) *GetSafetyCoverRequest {
	s.StartDate = &v
	return s
}

func (s *GetSafetyCoverRequest) SetSuspEventSource(v string) *GetSafetyCoverRequest {
	s.SuspEventSource = &v
	return s
}

type GetSafetyCoverResponseBody struct {
	// API return code.
	//
	// example:
	//
	// 404
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetSafetyCoverResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Message of the response result.
	//
	// example:
	//
	// system error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 564f8bb9-df3c-42a0-877a-b35d48f66603
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful:
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSafetyCoverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBody) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBody) SetCode(v string) *GetSafetyCoverResponseBody {
	s.Code = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetData(v *GetSafetyCoverResponseBodyData) *GetSafetyCoverResponseBody {
	s.Data = v
	return s
}

func (s *GetSafetyCoverResponseBody) SetHttpStatusCode(v int32) *GetSafetyCoverResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetMessage(v string) *GetSafetyCoverResponseBody {
	s.Message = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetRequestId(v string) *GetSafetyCoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetSuccess(v bool) *GetSafetyCoverResponseBody {
	s.Success = &v
	return s
}

type GetSafetyCoverResponseBodyData struct {
	// CFW protection coverage.
	CfwProtection *GetSafetyCoverResponseBodyDataCfwProtection `json:"CfwProtection,omitempty" xml:"CfwProtection,omitempty" type:"Struct"`
	// ECS protection coverage.
	EcsProtection *GetSafetyCoverResponseBodyDataEcsProtection `json:"EcsProtection,omitempty" xml:"EcsProtection,omitempty" type:"Struct"`
	// WAF protection coverage.
	WafProtection *GetSafetyCoverResponseBodyDataWafProtection `json:"WafProtection,omitempty" xml:"WafProtection,omitempty" type:"Struct"`
}

func (s GetSafetyCoverResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyData) SetCfwProtection(v *GetSafetyCoverResponseBodyDataCfwProtection) *GetSafetyCoverResponseBodyData {
	s.CfwProtection = v
	return s
}

func (s *GetSafetyCoverResponseBodyData) SetEcsProtection(v *GetSafetyCoverResponseBodyDataEcsProtection) *GetSafetyCoverResponseBodyData {
	s.EcsProtection = v
	return s
}

func (s *GetSafetyCoverResponseBodyData) SetWafProtection(v *GetSafetyCoverResponseBodyDataWafProtection) *GetSafetyCoverResponseBodyData {
	s.WafProtection = v
	return s
}

type GetSafetyCoverResponseBodyDataCfwProtection struct {
	// Number of unprotected items.
	//
	// example:
	//
	// 5
	NoProtectionCount *int64 `json:"NoProtectionCount,omitempty" xml:"NoProtectionCount,omitempty"`
	// Number of protected items.
	//
	// example:
	//
	// 5
	ProtectionCount *int64 `json:"ProtectionCount,omitempty" xml:"ProtectionCount,omitempty"`
	// Year-over-year protection rate.
	//
	// example:
	//
	// 35.00
	ProtectionGrowthRate *string `json:"ProtectionGrowthRate,omitempty" xml:"ProtectionGrowthRate,omitempty"`
	// Protection rate.
	//
	// example:
	//
	// 50.00
	ProtectionRate *string `json:"ProtectionRate,omitempty" xml:"ProtectionRate,omitempty"`
	// Total quantity.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSafetyCoverResponseBodyDataCfwProtection) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBodyDataCfwProtection) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetNoProtectionCount(v int64) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.NoProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetProtectionCount(v int64) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.ProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetProtectionGrowthRate(v string) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.ProtectionGrowthRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetProtectionRate(v string) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.ProtectionRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetTotalCount(v int64) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.TotalCount = &v
	return s
}

type GetSafetyCoverResponseBodyDataEcsProtection struct {
	// Number of unprotected items.
	//
	// example:
	//
	// 5
	NoProtectionCount *int64 `json:"NoProtectionCount,omitempty" xml:"NoProtectionCount,omitempty"`
	// Number of protected items.
	//
	// example:
	//
	// 5
	ProtectionCount *int64 `json:"ProtectionCount,omitempty" xml:"ProtectionCount,omitempty"`
	// Year-over-year growth in protection rate.
	//
	// example:
	//
	// 35.00
	ProtectionGrowthRate *string `json:"ProtectionGrowthRate,omitempty" xml:"ProtectionGrowthRate,omitempty"`
	// Protection rate.
	//
	// example:
	//
	// 50.00
	ProtectionRate *string `json:"ProtectionRate,omitempty" xml:"ProtectionRate,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSafetyCoverResponseBodyDataEcsProtection) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBodyDataEcsProtection) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetNoProtectionCount(v int64) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.NoProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetProtectionCount(v int64) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.ProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetProtectionGrowthRate(v string) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.ProtectionGrowthRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetProtectionRate(v string) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.ProtectionRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetTotalCount(v int64) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.TotalCount = &v
	return s
}

type GetSafetyCoverResponseBodyDataWafProtection struct {
	// Number of unprotected items.
	//
	// example:
	//
	// 5
	NoProtectionCount *int64 `json:"NoProtectionCount,omitempty" xml:"NoProtectionCount,omitempty"`
	// Number of protected items.
	//
	// example:
	//
	// 5
	ProtectionCount *int64 `json:"ProtectionCount,omitempty" xml:"ProtectionCount,omitempty"`
	// Year-over-year growth in protection rate.
	//
	// example:
	//
	// 35.00
	ProtectionGrowthRate *string `json:"ProtectionGrowthRate,omitempty" xml:"ProtectionGrowthRate,omitempty"`
	// Protection rate.
	//
	// example:
	//
	// 50.00
	ProtectionRate *string `json:"ProtectionRate,omitempty" xml:"ProtectionRate,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSafetyCoverResponseBodyDataWafProtection) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBodyDataWafProtection) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetNoProtectionCount(v int64) *GetSafetyCoverResponseBodyDataWafProtection {
	s.NoProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetProtectionCount(v int64) *GetSafetyCoverResponseBodyDataWafProtection {
	s.ProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetProtectionGrowthRate(v string) *GetSafetyCoverResponseBodyDataWafProtection {
	s.ProtectionGrowthRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetProtectionRate(v string) *GetSafetyCoverResponseBodyDataWafProtection {
	s.ProtectionRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetTotalCount(v int64) *GetSafetyCoverResponseBodyDataWafProtection {
	s.TotalCount = &v
	return s
}

type GetSafetyCoverResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSafetyCoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSafetyCoverResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponse) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponse) SetHeaders(v map[string]*string) *GetSafetyCoverResponse {
	s.Headers = v
	return s
}

func (s *GetSafetyCoverResponse) SetStatusCode(v int32) *GetSafetyCoverResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSafetyCoverResponse) SetBody(v *GetSafetyCoverResponseBody) *GetSafetyCoverResponse {
	s.Body = v
	return s
}

type GetSowListRequest struct {
	// Filter time type, supports filtering by the last 7 days, the last 30 days, the last half year, or custom time ranges.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetSowListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSowListRequest) GoString() string {
	return s.String()
}

func (s *GetSowListRequest) SetDateType(v string) *GetSowListRequest {
	s.DateType = &v
	return s
}

func (s *GetSowListRequest) SetEndDate(v int64) *GetSowListRequest {
	s.EndDate = &v
	return s
}

func (s *GetSowListRequest) SetStartDate(v int64) *GetSowListRequest {
	s.StartDate = &v
	return s
}

func (s *GetSowListRequest) SetSuspEventSource(v string) *GetSowListRequest {
	s.SuspEventSource = &v
	return s
}

type GetSowListResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetSowListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt information for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// FA8883BC-CB18-5E28-A113-8249917CA05E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSowListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSowListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSowListResponseBody) SetCode(v string) *GetSowListResponseBody {
	s.Code = &v
	return s
}

func (s *GetSowListResponseBody) SetData(v []*GetSowListResponseBodyData) *GetSowListResponseBody {
	s.Data = v
	return s
}

func (s *GetSowListResponseBody) SetHttpStatusCode(v int32) *GetSowListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSowListResponseBody) SetMessage(v string) *GetSowListResponseBody {
	s.Message = &v
	return s
}

func (s *GetSowListResponseBody) SetRequestId(v string) *GetSowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSowListResponseBody) SetSuccess(v bool) *GetSowListResponseBody {
	s.Success = &v
	return s
}

type GetSowListResponseBodyData struct {
	// Completion time.
	//
	// example:
	//
	// 2024-03-28 16:19:35
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// Operation remarks.
	//
	// example:
	//
	// 新建
	OperateRemark *string `json:"OperateRemark,omitempty" xml:"OperateRemark,omitempty"`
	// Progress.
	//
	// example:
	//
	// IN_PREPARATION
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Record count.
	//
	// example:
	//
	// 173
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// Start time.
	//
	// example:
	//
	// 2023-03-24 16:51:26
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Task type.
	//
	// example:
	//
	// 安全风险评估
	TaskTypeName *string `json:"TaskTypeName,omitempty" xml:"TaskTypeName,omitempty"`
	// Work order name.
	//
	// example:
	//
	// 安全产品配置问题与超量提醒
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
}

func (s GetSowListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSowListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSowListResponseBodyData) SetCompleteTime(v string) *GetSowListResponseBodyData {
	s.CompleteTime = &v
	return s
}

func (s *GetSowListResponseBodyData) SetOperateRemark(v string) *GetSowListResponseBodyData {
	s.OperateRemark = &v
	return s
}

func (s *GetSowListResponseBodyData) SetProgress(v string) *GetSowListResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetSowListResponseBodyData) SetRecordCount(v int32) *GetSowListResponseBodyData {
	s.RecordCount = &v
	return s
}

func (s *GetSowListResponseBodyData) SetStartTime(v string) *GetSowListResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetSowListResponseBodyData) SetTaskTypeName(v string) *GetSowListResponseBodyData {
	s.TaskTypeName = &v
	return s
}

func (s *GetSowListResponseBodyData) SetWorkOrderName(v string) *GetSowListResponseBodyData {
	s.WorkOrderName = &v
	return s
}

type GetSowListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSowListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSowListResponse) GoString() string {
	return s.String()
}

func (s *GetSowListResponse) SetHeaders(v map[string]*string) *GetSowListResponse {
	s.Headers = v
	return s
}

func (s *GetSowListResponse) SetStatusCode(v int32) *GetSowListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSowListResponse) SetBody(v *GetSowListResponseBody) *GetSowListResponse {
	s.Body = v
	return s
}

type GetSuspEventPageRequest struct {
	// Alarm end time.
	//
	// example:
	//
	// 1732515522000
	AlarmEndTime *int64 `json:"AlarmEndTime,omitempty" xml:"AlarmEndTime,omitempty"`
	// Alarm start time.
	//
	// example:
	//
	// 1722515522000
	AlarmStartTime *int64 `json:"AlarmStartTime,omitempty" xml:"AlarmStartTime,omitempty"`
	// Current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Alarm source.
	//
	// example:
	//
	// SUSP_EVENT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Disposal status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSuspEventPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageRequest) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageRequest) SetAlarmEndTime(v int64) *GetSuspEventPageRequest {
	s.AlarmEndTime = &v
	return s
}

func (s *GetSuspEventPageRequest) SetAlarmStartTime(v int64) *GetSuspEventPageRequest {
	s.AlarmStartTime = &v
	return s
}

func (s *GetSuspEventPageRequest) SetCurrentPage(v int32) *GetSuspEventPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSuspEventPageRequest) SetPageSize(v int32) *GetSuspEventPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetSuspEventPageRequest) SetSource(v string) *GetSuspEventPageRequest {
	s.Source = &v
	return s
}

func (s *GetSuspEventPageRequest) SetStatus(v int32) *GetSuspEventPageRequest {
	s.Status = &v
	return s
}

type GetSuspEventPageResponseBody struct {
	// API response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetSuspEventPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// system error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *GetSuspEventPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AFA6F7B7-7C4B-58BB-B8FB-E0FFA4483561
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSuspEventPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponseBody) SetCode(v string) *GetSuspEventPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetData(v []*GetSuspEventPageResponseBodyData) *GetSuspEventPageResponseBody {
	s.Data = v
	return s
}

func (s *GetSuspEventPageResponseBody) SetHttpStatusCode(v int32) *GetSuspEventPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetMessage(v string) *GetSuspEventPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetPageInfo(v *GetSuspEventPageResponseBodyPageInfo) *GetSuspEventPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetSuspEventPageResponseBody) SetRequestId(v string) *GetSuspEventPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetSuccess(v bool) *GetSuspEventPageResponseBody {
	s.Success = &v
	return s
}

type GetSuspEventPageResponseBodyData struct {
	// Alarm event type.
	//
	// example:
	//
	// 精准防御
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// Alarm ID.
	//
	// example:
	//
	// 5b1eeebe4f22daa2b177298234214fa3
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// Alarm name.
	//
	// example:
	//
	// Web服务漏洞利用
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// Alarm source.
	//
	// example:
	//
	// SUSP_EVENT
	AlarmSource *string `json:"AlarmSource,omitempty" xml:"AlarmSource,omitempty"`
	// Latest alarm time.
	//
	// example:
	//
	// 1722515522000
	AlarmTime *string `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	// Analysis process.
	//
	// example:
	//
	// [{"value":"服务器可能已被黑客攻击，存在恶意进程在运行。 分析过程：告警显示，服务端存在一个名为”dns.exe”的进程在访问”polling.burpcollaborator.net”，这是一个被黑名单列出的恶意域名。在正常情况下,”dns.exe”不应该单独存在于系统的路径下，并且也不应该访问这类恶意域名。因此，这个进程可能是黑客留下的恶意进程。","key":"结论"},{"value":"尽快对服务器进行全面扫描，清除恶意进程。同时，联系网络安全专家进行深入调查，以确定是否有其他潜在的安全威胁。","key":"处置建议"}]
	AnalysisResult *string `json:"AnalysisResult,omitempty" xml:"AnalysisResult,omitempty"`
	// Alarm handling time.
	//
	// example:
	//
	// 1732515522000
	DealTime *string `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// Alarm level.
	//
	// example:
	//
	// suspicious
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// Ticket primary key id.
	//
	// example:
	//
	// 9947
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Affected asset.
	//
	// example:
	//
	// shells-azhou
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Public IP address.
	//
	// example:
	//
	// 47.99.188.31
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Private IP address.
	//
	// example:
	//
	// 172.16.109.130
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// First occurrence time.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	// Owner ID.
	//
	// example:
	//
	// 张三
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Disposal method.
	//
	// example:
	//
	// 处理完成
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Handling status.
	//
	// example:
	//
	// 未处理
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSuspEventPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmEventType(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmEventType = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmId(v int64) *GetSuspEventPageResponseBodyData {
	s.AlarmId = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmName(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmName = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmSource(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmSource = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmTime(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmTime = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAnalysisResult(v string) *GetSuspEventPageResponseBodyData {
	s.AnalysisResult = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetDealTime(v string) *GetSuspEventPageResponseBodyData {
	s.DealTime = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetEventLevel(v string) *GetSuspEventPageResponseBodyData {
	s.EventLevel = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetId(v int64) *GetSuspEventPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetInstanceName(v string) *GetSuspEventPageResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetInternetIp(v string) *GetSuspEventPageResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetIntranetIp(v string) *GetSuspEventPageResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetOccurrenceTime(v string) *GetSuspEventPageResponseBodyData {
	s.OccurrenceTime = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetOwnerId(v string) *GetSuspEventPageResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetRemark(v string) *GetSuspEventPageResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetStatus(v string) *GetSuspEventPageResponseBodyData {
	s.Status = &v
	return s
}

type GetSuspEventPageResponseBodyPageInfo struct {
	// The current page number in pagination queries.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of items displayed per page in the returned data.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of query results.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSuspEventPageResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponseBodyPageInfo) SetCurrentPage(v int32) *GetSuspEventPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetSuspEventPageResponseBodyPageInfo) SetPageSize(v int32) *GetSuspEventPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetSuspEventPageResponseBodyPageInfo) SetTotalCount(v int32) *GetSuspEventPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type GetSuspEventPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspEventPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspEventPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageResponse) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponse) SetHeaders(v map[string]*string) *GetSuspEventPageResponse {
	s.Headers = v
	return s
}

func (s *GetSuspEventPageResponse) SetStatusCode(v int32) *GetSuspEventPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspEventPageResponse) SetBody(v *GetSuspEventPageResponseBody) *GetSuspEventPageResponse {
	s.Body = v
	return s
}

type GetSuspEventSummaryRequest struct {
	// Filter time type. Supports filtering by the last 7 days, the last 30 days, the last half year, or custom time ranges.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// SUSP_EVENT
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetSuspEventSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryRequest) SetDateType(v string) *GetSuspEventSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetSuspEventSummaryRequest) SetEndDate(v int64) *GetSuspEventSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetSuspEventSummaryRequest) SetStartDate(v int64) *GetSuspEventSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetSuspEventSummaryRequest) SetSuspEventSource(v string) *GetSuspEventSummaryRequest {
	s.SuspEventSource = &v
	return s
}

type GetSuspEventSummaryResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetSuspEventSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9B2DAE9B-B901-5818-AFEF-E5637D938280
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// - true: Call succeeded.
	//
	// - false: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSuspEventSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBody) SetCode(v string) *GetSuspEventSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetData(v *GetSuspEventSummaryResponseBodyData) *GetSuspEventSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetHttpStatusCode(v int32) *GetSuspEventSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetMessage(v string) *GetSuspEventSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetRequestId(v string) *GetSuspEventSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetSuccess(v bool) *GetSuspEventSummaryResponseBody {
	s.Success = &v
	return s
}

type GetSuspEventSummaryResponseBodyData struct {
	// Network attack trend.
	NetworkAttackTrendDTO *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO `json:"NetworkAttackTrendDTO,omitempty" xml:"NetworkAttackTrendDTO,omitempty" type:"Struct"`
	// Overview of alert handling.
	SuspEventDealSummaryDTO *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO `json:"SuspEventDealSummaryDTO,omitempty" xml:"SuspEventDealSummaryDTO,omitempty" type:"Struct"`
	// Top 10 alerts before handling.
	SuspEventTopDTO *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO `json:"SuspEventTopDTO,omitempty" xml:"SuspEventTopDTO,omitempty" type:"Struct"`
	// Trend of alert responses.
	SuspEventTrendDTO *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO `json:"SuspEventTrendDTO,omitempty" xml:"SuspEventTrendDTO,omitempty" type:"Struct"`
}

func (s GetSuspEventSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyData) SetNetworkAttackTrendDTO(v *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) *GetSuspEventSummaryResponseBodyData {
	s.NetworkAttackTrendDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) SetSuspEventDealSummaryDTO(v *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) *GetSuspEventSummaryResponseBodyData {
	s.SuspEventDealSummaryDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) SetSuspEventTopDTO(v *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) *GetSuspEventSummaryResponseBodyData {
	s.SuspEventTopDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) SetSuspEventTrendDTO(v *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) *GetSuspEventSummaryResponseBodyData {
	s.SuspEventTrendDTO = v
	return s
}

type GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO struct {
	// Collection of trend nodes for each attack item.
	TrendList []*GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) SetTrendList(v []*GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO {
	s.TrendList = v
	return s
}

type GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList struct {
	// Date.
	//
	// example:
	//
	// 202409或20240901
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// DDoS count.
	//
	// example:
	//
	// 10
	DdosCount *int64 `json:"DdosCount,omitempty" xml:"DdosCount,omitempty"`
	// EIP count.
	//
	// example:
	//
	// 10
	EipCount *int64 `json:"EipCount,omitempty" xml:"EipCount,omitempty"`
	// WAF count.
	//
	// example:
	//
	// 10
	WafCount *int64 `json:"WafCount,omitempty" xml:"WafCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetDate(v string) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.Date = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetDdosCount(v int64) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.DdosCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetEipCount(v int64) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.EipCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetWafCount(v int64) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.WafCount = &v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO struct {
	// Completed.
	//
	// example:
	//
	// 20
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// In progress.
	//
	// example:
	//
	// 5
	HandingCount *int64 `json:"HandingCount,omitempty" xml:"HandingCount,omitempty"`
	// Alert handling rate.
	//
	// example:
	//
	// 90
	HandingRate *string `json:"HandingRate,omitempty" xml:"HandingRate,omitempty"`
	// Total number of alerts.
	//
	// example:
	//
	// 35
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Year-over-year comparison of alerts.
	//
	// example:
	//
	// 10
	TotalGrowthRate *string `json:"TotalGrowthRate,omitempty" xml:"TotalGrowthRate,omitempty"`
	// Number of unhandled alerts.
	//
	// example:
	//
	// 10
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetCompletedCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.CompletedCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetHandingCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.HandingCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetHandingRate(v string) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.HandingRate = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetTotalCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.TotalCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetTotalGrowthRate(v string) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.TotalGrowthRate = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetWaitHandleCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.WaitHandleCount = &v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventTopDTO struct {
	// Top 10 before handling alarms
	SuspEventList []*GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList `json:"SuspEventList,omitempty" xml:"SuspEventList,omitempty" type:"Repeated"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) SetSuspEventList(v []*GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO {
	s.SuspEventList = v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList struct {
	// Alert name.
	//
	// example:
	//
	// 主动外连风险 IP
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// Count.
	//
	// example:
	//
	// 7
	TaskCount *int64 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) SetEventName(v string) *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList {
	s.EventName = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) SetTaskCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList {
	s.TaskCount = &v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO struct {
	// Trend of alerts.
	TrendList []*GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) SetTrendList(v []*GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO {
	s.TrendList = v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList struct {
	// Time point.
	//
	// example:
	//
	// 202405或者20240501
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Number of handled alerts.
	//
	// example:
	//
	// 10
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// Number of discovered alerts.
	//
	// example:
	//
	// 15
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) SetDate(v string) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	s.Date = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) SetDealCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	s.DealCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) SetFindCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	s.FindCount = &v
	return s
}

type GetSuspEventSummaryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspEventSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspEventSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponse) SetHeaders(v map[string]*string) *GetSuspEventSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetSuspEventSummaryResponse) SetStatusCode(v int32) *GetSuspEventSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspEventSummaryResponse) SetBody(v *GetSuspEventSummaryResponseBody) *GetSuspEventSummaryResponse {
	s.Body = v
	return s
}

type GetSuspPageSummaryResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetSuspPageSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the result returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EF801DD1-D934-51B3-92D4-776CE17B184F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSuspPageSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuspPageSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspPageSummaryResponseBody) SetCode(v string) *GetSuspPageSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetData(v *GetSuspPageSummaryResponseBodyData) *GetSuspPageSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetHttpStatusCode(v int32) *GetSuspPageSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetMessage(v string) *GetSuspPageSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetRequestId(v string) *GetSuspPageSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetSuccess(v bool) *GetSuspPageSummaryResponseBody {
	s.Success = &v
	return s
}

type GetSuspPageSummaryResponseBodyData struct {
	// Number of completed items.
	//
	// example:
	//
	// 10
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// Number of items being processed.
	//
	// example:
	//
	// 10
	HandingCount *int64 `json:"HandingCount,omitempty" xml:"HandingCount,omitempty"`
	// Number of high-risk items.
	//
	// example:
	//
	// 10
	HighCount *int64 `json:"HighCount,omitempty" xml:"HighCount,omitempty"`
	// Number of low-risk items.
	//
	// example:
	//
	// 10
	LowCount *int64 `json:"LowCount,omitempty" xml:"LowCount,omitempty"`
	// Number of medium-risk items.
	//
	// example:
	//
	// 10
	MediumCount *int64 `json:"MediumCount,omitempty" xml:"MediumCount,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 30
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Number of unhandled items.
	//
	// example:
	//
	// 10
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetSuspPageSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSuspPageSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSuspPageSummaryResponseBodyData) SetCompletedCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetHandingCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.HandingCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetHighCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.HighCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetLowCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.LowCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetMediumCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.MediumCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetTotalCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetWaitHandleCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.WaitHandleCount = &v
	return s
}

type GetSuspPageSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspPageSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspPageSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuspPageSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetSuspPageSummaryResponse) SetHeaders(v map[string]*string) *GetSuspPageSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetSuspPageSummaryResponse) SetStatusCode(v int32) *GetSuspPageSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspPageSummaryResponse) SetBody(v *GetSuspPageSummaryResponseBody) *GetSuspPageSummaryResponse {
	s.Body = v
	return s
}

type GetUserStatusResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetUserStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// D8DBD769-613E-5E6B-A9FD-B622375B152D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserStatusResponseBody) SetCode(v string) *GetUserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserStatusResponseBody) SetData(v *GetUserStatusResponseBodyData) *GetUserStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetUserStatusResponseBody) SetHttpStatusCode(v int32) *GetUserStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserStatusResponseBody) SetMessage(v string) *GetUserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserStatusResponseBody) SetRequestId(v string) *GetUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserStatusResponseBody) SetSuccess(v bool) *GetUserStatusResponseBody {
	s.Success = &v
	return s
}

type GetUserStatusResponseBodyData struct {
	// Customer type.
	//
	// example:
	//
	// official
	CustomerType *string `json:"CustomerType,omitempty" xml:"CustomerType,omitempty"`
	// End date.
	//
	// example:
	//
	// 2023-09-28 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// 726cec3c-4887-4354-8c21-c0ad12e10fc2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Start date.
	//
	// example:
	//
	// 2023-09-20 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Status.
	//
	// example:
	//
	// FirstLogin
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Version.
	//
	// example:
	//
	// mdrjichu
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetUserStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserStatusResponseBodyData) SetCustomerType(v string) *GetUserStatusResponseBodyData {
	s.CustomerType = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetEndDate(v string) *GetUserStatusResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetInstanceId(v string) *GetUserStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetStartDate(v string) *GetUserStatusResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetStatus(v string) *GetUserStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetVersion(v string) *GetUserStatusResponseBodyData {
	s.Version = &v
	return s
}

type GetUserStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserStatusResponse) SetHeaders(v map[string]*string) *GetUserStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserStatusResponse) SetStatusCode(v int32) *GetUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserStatusResponse) SetBody(v *GetUserStatusResponseBody) *GetUserStatusResponse {
	s.Body = v
	return s
}

type GetVulItemPageRequest struct {
	// Vulnerability alias.
	//
	// example:
	//
	// RHSA-2018:3665-Important: NetworkManager security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Processing status. y: processed; n: unprocessed; h: processing.
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// Risk level.
	//
	// example:
	//
	// later
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Vulnerability name.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20183665
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Number of items to display per page in the returned data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Vulnerability type.
	//
	// example:
	//
	// sca
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
}

func (s GetVulItemPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageRequest) GoString() string {
	return s.String()
}

func (s *GetVulItemPageRequest) SetAliasName(v string) *GetVulItemPageRequest {
	s.AliasName = &v
	return s
}

func (s *GetVulItemPageRequest) SetCurrentPage(v int32) *GetVulItemPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetVulItemPageRequest) SetDealed(v string) *GetVulItemPageRequest {
	s.Dealed = &v
	return s
}

func (s *GetVulItemPageRequest) SetLevel(v string) *GetVulItemPageRequest {
	s.Level = &v
	return s
}

func (s *GetVulItemPageRequest) SetName(v string) *GetVulItemPageRequest {
	s.Name = &v
	return s
}

func (s *GetVulItemPageRequest) SetPageSize(v int32) *GetVulItemPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetVulItemPageRequest) SetScanType(v string) *GetVulItemPageRequest {
	s.ScanType = &v
	return s
}

type GetVulItemPageResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetVulItemPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *GetVulItemPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02F8BBF3-2D61-5982-8911-EEB387BE3AF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// true: Call succeeded. false: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulItemPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponseBody) SetCode(v string) *GetVulItemPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetData(v []*GetVulItemPageResponseBodyData) *GetVulItemPageResponseBody {
	s.Data = v
	return s
}

func (s *GetVulItemPageResponseBody) SetHttpStatusCode(v int32) *GetVulItemPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetMessage(v string) *GetVulItemPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetPageInfo(v *GetVulItemPageResponseBodyPageInfo) *GetVulItemPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetVulItemPageResponseBody) SetRequestId(v string) *GetVulItemPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetSuccess(v bool) *GetVulItemPageResponseBody {
	s.Success = &v
	return s
}

type GetVulItemPageResponseBodyData struct {
	// Vulnerability alias.
	//
	// example:
	//
	// RHSA-2024:4620: libndp 安全更新
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Number of high-priority vulnerabilities to be fixed.
	//
	// example:
	//
	// 74
	AsapCount *int32 `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	// User ID.
	//
	// example:
	//
	// 1940494487193744
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// Prefix for the CVE remediation advice URL.
	//
	// example:
	//
	// https://avd.aliyun.com/detail/
	CveUrlPrefix *string `json:"CveUrlPrefix,omitempty" xml:"CveUrlPrefix,omitempty"`
	// Processing status.
	//
	// example:
	//
	// y
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// Timestamp of the last discovery of the vulnerability.
	//
	// example:
	//
	// 2023-04-23 14:47:34
	FindTime *string `json:"FindTime,omitempty" xml:"FindTime,omitempty"`
	// Number of processed vulnerabilities.
	//
	// example:
	//
	// 20
	HandledCount *int32 `json:"HandledCount,omitempty" xml:"HandledCount,omitempty"`
	// Primary key ID.
	//
	// example:
	//
	// 353845
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Number of medium-priority vulnerabilities to be fixed.
	//
	// example:
	//
	// 10
	LaterCount *int32 `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	// Risk level
	//
	// example:
	//
	// later
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Vulnerability name.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20205002
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Number of low-priority vulnerabilities to be fixed.
	//
	// example:
	//
	// 8
	NntfCount *int32 `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	// CVE number.
	//
	// example:
	//
	// CVE-2019-20907
	Related *string `json:"Related,omitempty" xml:"Related,omitempty"`
	// Number of related CVE numbers.
	//
	// example:
	//
	// 20
	RelatedCveCount *int32 `json:"RelatedCveCount,omitempty" xml:"RelatedCveCount,omitempty"`
	// Vulnerability type.
	//
	// example:
	//
	// sca
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// Tags.
	//
	// example:
	//
	// Elevation of Privilege
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Total number of fixed vulnerabilities.
	//
	// example:
	//
	// 50
	TotalFixCount *int64 `json:"TotalFixCount,omitempty" xml:"TotalFixCount,omitempty"`
}

func (s GetVulItemPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponseBodyData) SetAliasName(v string) *GetVulItemPageResponseBodyData {
	s.AliasName = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetAsapCount(v int32) *GetVulItemPageResponseBodyData {
	s.AsapCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetCustomerId(v string) *GetVulItemPageResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetCveUrlPrefix(v string) *GetVulItemPageResponseBodyData {
	s.CveUrlPrefix = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetDealed(v string) *GetVulItemPageResponseBodyData {
	s.Dealed = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetFindTime(v string) *GetVulItemPageResponseBodyData {
	s.FindTime = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetHandledCount(v int32) *GetVulItemPageResponseBodyData {
	s.HandledCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetId(v int64) *GetVulItemPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetLaterCount(v int32) *GetVulItemPageResponseBodyData {
	s.LaterCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetLevel(v string) *GetVulItemPageResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetName(v string) *GetVulItemPageResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetNntfCount(v int32) *GetVulItemPageResponseBodyData {
	s.NntfCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetRelated(v string) *GetVulItemPageResponseBodyData {
	s.Related = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetRelatedCveCount(v int32) *GetVulItemPageResponseBodyData {
	s.RelatedCveCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetScanType(v string) *GetVulItemPageResponseBodyData {
	s.ScanType = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetTags(v string) *GetVulItemPageResponseBodyData {
	s.Tags = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetTotalFixCount(v int64) *GetVulItemPageResponseBodyData {
	s.TotalFixCount = &v
	return s
}

type GetVulItemPageResponseBodyPageInfo struct {
	// The current page number for pagination queries.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items to display per page in the returned data.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of records in the query result.
	//
	// example:
	//
	// 163
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetVulItemPageResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponseBodyPageInfo) SetCurrentPage(v int32) *GetVulItemPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetVulItemPageResponseBodyPageInfo) SetPageSize(v int32) *GetVulItemPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetVulItemPageResponseBodyPageInfo) SetTotalCount(v int32) *GetVulItemPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type GetVulItemPageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulItemPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulItemPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageResponse) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponse) SetHeaders(v map[string]*string) *GetVulItemPageResponse {
	s.Headers = v
	return s
}

func (s *GetVulItemPageResponse) SetStatusCode(v int32) *GetVulItemPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulItemPageResponse) SetBody(v *GetVulItemPageResponseBody) *GetVulItemPageResponse {
	s.Body = v
	return s
}

type GetVulListByIdRequest struct {
	// Current page
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Whether it has been processed; y: processed; n: not processed
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 4209205
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Risk level
	//
	// example:
	//
	// asap,later,nntf
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Asset information of the vulnerability to be queried, which can be set as asset name, public IP, or private IP.
	//
	// example:
	//
	// production_nat_cn-hangzhou_zone_105
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// UUID of the server with the vulnerability to be queried. Multiple UUIDs should be separated by a comma (,).
	//
	// example:
	//
	// 3615b908-995a-4edb-bc85-1981b4e94ba0,9c52cf9a-d8ba-4e31-ae06-500b879ee4e6,4b7de3cf-c4ac-42fc-8804-35070493dc29,f3c01525-0777-4c97-88d9-bec11afd4a6a,a80bd516-c4f3-4c27-a169-c8abfaf9e89e
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s GetVulListByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVulListByIdRequest) GoString() string {
	return s.String()
}

func (s *GetVulListByIdRequest) SetCurrentPage(v int32) *GetVulListByIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetVulListByIdRequest) SetDealed(v string) *GetVulListByIdRequest {
	s.Dealed = &v
	return s
}

func (s *GetVulListByIdRequest) SetId(v int64) *GetVulListByIdRequest {
	s.Id = &v
	return s
}

func (s *GetVulListByIdRequest) SetNecessity(v string) *GetVulListByIdRequest {
	s.Necessity = &v
	return s
}

func (s *GetVulListByIdRequest) SetPageSize(v int32) *GetVulListByIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetVulListByIdRequest) SetRemark(v string) *GetVulListByIdRequest {
	s.Remark = &v
	return s
}

func (s *GetVulListByIdRequest) SetUuids(v string) *GetVulListByIdRequest {
	s.Uuids = &v
	return s
}

type GetVulListByIdResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetVulListByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *GetVulListByIdResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// D38B3D2F-67FD-57FF-87D1-C431D2C70F76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Values: - **true**: Yes. - **false**: No.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulListByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVulListByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponseBody) SetCode(v string) *GetVulListByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulListByIdResponseBody) SetData(v []*GetVulListByIdResponseBodyData) *GetVulListByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetVulListByIdResponseBody) SetHttpStatusCode(v int32) *GetVulListByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulListByIdResponseBody) SetMessage(v string) *GetVulListByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulListByIdResponseBody) SetPageInfo(v *GetVulListByIdResponseBodyPageInfo) *GetVulListByIdResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetVulListByIdResponseBody) SetRequestId(v string) *GetVulListByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulListByIdResponseBody) SetSuccess(v bool) *GetVulListByIdResponseBody {
	s.Success = &v
	return s
}

type GetVulListByIdResponseBodyData struct {
	// Vulnerability Alias
	//
	// example:
	//
	// Tomcat websocket 拒绝服务漏洞利用代码披露（CVE-2020-13935）
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Impact description
	EffectMsgDTOS []*GetVulListByIdResponseBodyDataEffectMsgDTOS `json:"EffectMsgDTOS,omitempty" xml:"EffectMsgDTOS,omitempty" type:"Repeated"`
	// Timestamp of the first time the vulnerability was detected
	//
	// example:
	//
	// 1620404763000
	FirstTs *string `json:"FirstTs,omitempty" xml:"FirstTs,omitempty"`
	// Instance name of the asset
	//
	// example:
	//
	// 凌星-CentOS
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Public IP of the asset
	//
	// example:
	//
	// 39.101.73.28
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Private IP of the asset
	//
	// example:
	//
	// 172.22.216.17
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// Timestamp of the last time the vulnerability was detected
	//
	// example:
	//
	// 1620404763000
	LastTs *string `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	// Vulnerability name
	//
	// example:
	//
	// SCA:ACSV-2020-111301
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Necessity level of vulnerability repair
	//
	// example:
	//
	// later,asap,nntf
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// List of associated CVEs for the vulnerability, separated by commas (,) if there are multiple values.
	//
	// example:
	//
	// CVE-2020-13935
	Related *string `json:"Related,omitempty" xml:"Related,omitempty"`
	// Repair command
	//
	// example:
	//
	// **	- update python-perf
	RepairCmd *string `json:"RepairCmd,omitempty" xml:"RepairCmd,omitempty"`
	// Timestamp of vulnerability repair
	//
	// example:
	//
	// 1541207563000
	RepairTs *string `json:"RepairTs,omitempty" xml:"RepairTs,omitempty"`
	// Vulnerability status:
	//
	// 1: Not fixed
	//
	// 2: Fix failed
	//
	// 3: Rollback failed
	//
	// 4: Fixing
	//
	// 5: Rolling back
	//
	// 6: Verifying
	//
	// 7: Fixed successfully
	//
	// 8: Fixed successfully, pending reboot
	//
	// 9: Rolled back successfully
	//
	// 10: Ignored
	//
	// 11: Rolled back successfully, pending reboot
	//
	// 12: Vulnerability does not exist
	//
	// 20: Expired
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Vulnerability tag
	//
	// example:
	//
	// Restart Required
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// UUID of the asset instance.
	//
	// example:
	//
	// hdm_5cf2eaf263c021b354877943f181956d
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetVulListByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVulListByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponseBodyData) SetAliasName(v string) *GetVulListByIdResponseBodyData {
	s.AliasName = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetEffectMsgDTOS(v []*GetVulListByIdResponseBodyDataEffectMsgDTOS) *GetVulListByIdResponseBodyData {
	s.EffectMsgDTOS = v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetFirstTs(v string) *GetVulListByIdResponseBodyData {
	s.FirstTs = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetInstanceName(v string) *GetVulListByIdResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetInternetIp(v string) *GetVulListByIdResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetIntranetIp(v string) *GetVulListByIdResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetLastTs(v string) *GetVulListByIdResponseBodyData {
	s.LastTs = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetName(v string) *GetVulListByIdResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetNecessity(v string) *GetVulListByIdResponseBodyData {
	s.Necessity = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetRelated(v string) *GetVulListByIdResponseBodyData {
	s.Related = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetRepairCmd(v string) *GetVulListByIdResponseBodyData {
	s.RepairCmd = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetRepairTs(v string) *GetVulListByIdResponseBodyData {
	s.RepairTs = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetStatus(v string) *GetVulListByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetTag(v string) *GetVulListByIdResponseBodyData {
	s.Tag = &v
	return s
}

func (s *GetVulListByIdResponseBodyData) SetUuid(v string) *GetVulListByIdResponseBodyData {
	s.Uuid = &v
	return s
}

type GetVulListByIdResponseBodyDataEffectMsgDTOS struct {
	// Hit
	//
	// example:
	//
	// fastjson(jar) extendField.safemode equals false
	MatchList *string `json:"MatchList,omitempty" xml:"MatchList,omitempty"`
	// Path
	//
	// example:
	//
	// /uat6/qry/enquiry/policy/yrtPolicyList
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Software name
	//
	// example:
	//
	// python-perf 3.10.0
	SoftName *string `json:"SoftName,omitempty" xml:"SoftName,omitempty"`
}

func (s GetVulListByIdResponseBodyDataEffectMsgDTOS) String() string {
	return tea.Prettify(s)
}

func (s GetVulListByIdResponseBodyDataEffectMsgDTOS) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) SetMatchList(v string) *GetVulListByIdResponseBodyDataEffectMsgDTOS {
	s.MatchList = &v
	return s
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) SetPath(v string) *GetVulListByIdResponseBodyDataEffectMsgDTOS {
	s.Path = &v
	return s
}

func (s *GetVulListByIdResponseBodyDataEffectMsgDTOS) SetSoftName(v string) *GetVulListByIdResponseBodyDataEffectMsgDTOS {
	s.SoftName = &v
	return s
}

type GetVulListByIdResponseBodyPageInfo struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page in the returned data.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of records in the query result.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetVulListByIdResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVulListByIdResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponseBodyPageInfo) SetCurrentPage(v int32) *GetVulListByIdResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetVulListByIdResponseBodyPageInfo) SetPageSize(v int32) *GetVulListByIdResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetVulListByIdResponseBodyPageInfo) SetTotalCount(v int32) *GetVulListByIdResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type GetVulListByIdResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulListByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulListByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVulListByIdResponse) GoString() string {
	return s.String()
}

func (s *GetVulListByIdResponse) SetHeaders(v map[string]*string) *GetVulListByIdResponse {
	s.Headers = v
	return s
}

func (s *GetVulListByIdResponse) SetStatusCode(v int32) *GetVulListByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulListByIdResponse) SetBody(v *GetVulListByIdResponseBody) *GetVulListByIdResponse {
	s.Body = v
	return s
}

type GetVulPageSummaryResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetVulPageSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// A3A575C8-80F9-5F04-AA24-CCAC246884A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulPageSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVulPageSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulPageSummaryResponseBody) SetCode(v string) *GetVulPageSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetData(v *GetVulPageSummaryResponseBodyData) *GetVulPageSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetHttpStatusCode(v int32) *GetVulPageSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetMessage(v string) *GetVulPageSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetRequestId(v string) *GetVulPageSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetSuccess(v bool) *GetVulPageSummaryResponseBody {
	s.Success = &v
	return s
}

type GetVulPageSummaryResponseBodyData struct {
	// Number of completed items.
	//
	// example:
	//
	// 1990
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// Number of items being handled.
	//
	// example:
	//
	// 6
	HandingCount *int64 `json:"HandingCount,omitempty" xml:"HandingCount,omitempty"`
	// Number of high-risk items.
	//
	// example:
	//
	// 500
	HighCount *int64 `json:"HighCount,omitempty" xml:"HighCount,omitempty"`
	// Number of low-risk items.
	//
	// example:
	//
	// 1000
	LowCount *int64 `json:"LowCount,omitempty" xml:"LowCount,omitempty"`
	// Number of medium-risk items.
	//
	// example:
	//
	// 500
	MediumCount *int64 `json:"MediumCount,omitempty" xml:"MediumCount,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 2000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Number of unhandled items.
	//
	// example:
	//
	// 4
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetVulPageSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVulPageSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulPageSummaryResponseBodyData) SetCompletedCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetHandingCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.HandingCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetHighCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.HighCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetLowCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.LowCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetMediumCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.MediumCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetTotalCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetWaitHandleCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.WaitHandleCount = &v
	return s
}

type GetVulPageSummaryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulPageSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulPageSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVulPageSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetVulPageSummaryResponse) SetHeaders(v map[string]*string) *GetVulPageSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetVulPageSummaryResponse) SetStatusCode(v int32) *GetVulPageSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulPageSummaryResponse) SetBody(v *GetVulPageSummaryResponseBody) *GetVulPageSummaryResponse {
	s.Body = v
	return s
}

type GetVulSummaryRequest struct {
	// Filter time type. Supports filtering by the last 7 days, the last 30 days, the last half year, or a custom time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// 该字段暂未使用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetVulSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetVulSummaryRequest) SetDateType(v string) *GetVulSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetVulSummaryRequest) SetEndDate(v int64) *GetVulSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetVulSummaryRequest) SetStartDate(v int64) *GetVulSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetVulSummaryRequest) SetSuspEventSource(v string) *GetVulSummaryRequest {
	s.SuspEventSource = &v
	return s
}

type GetVulSummaryResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetVulSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the response result.
	//
	// example:
	//
	// system error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EF801DD1-D934-51B3-92D4-776CE17B184F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponseBody) SetCode(v string) *GetVulSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetData(v *GetVulSummaryResponseBodyData) *GetVulSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetVulSummaryResponseBody) SetHttpStatusCode(v int32) *GetVulSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetMessage(v string) *GetVulSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetRequestId(v string) *GetVulSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetSuccess(v bool) *GetVulSummaryResponseBody {
	s.Success = &v
	return s
}

type GetVulSummaryResponseBodyData struct {
	// Number of completed items.
	//
	// example:
	//
	// 10
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// Risk convergence rate.
	//
	// example:
	//
	// 50
	DealRate *string `json:"DealRate,omitempty" xml:"DealRate,omitempty"`
	// Collection of vulnerability trend nodes.
	TrendList []*GetVulSummaryResponseBodyDataTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
	// Number of unhandled items.
	//
	// example:
	//
	// 5
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetVulSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponseBodyData) SetCompletedCount(v int64) *GetVulSummaryResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetVulSummaryResponseBodyData) SetDealRate(v string) *GetVulSummaryResponseBodyData {
	s.DealRate = &v
	return s
}

func (s *GetVulSummaryResponseBodyData) SetTrendList(v []*GetVulSummaryResponseBodyDataTrendList) *GetVulSummaryResponseBodyData {
	s.TrendList = v
	return s
}

func (s *GetVulSummaryResponseBodyData) SetWaitHandleCount(v int64) *GetVulSummaryResponseBodyData {
	s.WaitHandleCount = &v
	return s
}

type GetVulSummaryResponseBodyDataTrendList struct {
	// Time point.
	//
	// example:
	//
	// 202407或者20240701
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Number of handled items.
	//
	// example:
	//
	// 10
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// Number of discovered items.
	//
	// example:
	//
	// 15
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetVulSummaryResponseBodyDataTrendList) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryResponseBodyDataTrendList) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponseBodyDataTrendList) SetDate(v string) *GetVulSummaryResponseBodyDataTrendList {
	s.Date = &v
	return s
}

func (s *GetVulSummaryResponseBodyDataTrendList) SetDealCount(v int64) *GetVulSummaryResponseBodyDataTrendList {
	s.DealCount = &v
	return s
}

func (s *GetVulSummaryResponseBodyDataTrendList) SetFindCount(v int64) *GetVulSummaryResponseBodyDataTrendList {
	s.FindCount = &v
	return s
}

type GetVulSummaryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponse) SetHeaders(v map[string]*string) *GetVulSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetVulSummaryResponse) SetStatusCode(v int32) *GetVulSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulSummaryResponse) SetBody(v *GetVulSummaryResponseBody) *GetVulSummaryResponse {
	s.Body = v
	return s
}

type GetWorkTaskSummaryRequest struct {
	// Filter time type, supports filtering by the last 7 days, the last 30 days, the last half year, or custom time periods.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Alert event source.
	//
	// example:
	//
	// 该字段暂时未用，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetWorkTaskSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkTaskSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryRequest) SetDateType(v string) *GetWorkTaskSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) SetEndDate(v int64) *GetWorkTaskSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) SetStartDate(v int64) *GetWorkTaskSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) SetSuspEventSource(v string) *GetWorkTaskSummaryRequest {
	s.SuspEventSource = &v
	return s
}

type GetWorkTaskSummaryResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetWorkTaskSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the response result.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// EF801DD1-D934-51B3-92D4-776CE17B184F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkTaskSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkTaskSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryResponseBody) SetCode(v string) *GetWorkTaskSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetData(v *GetWorkTaskSummaryResponseBodyData) *GetWorkTaskSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetHttpStatusCode(v int32) *GetWorkTaskSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetMessage(v string) *GetWorkTaskSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetRequestId(v string) *GetWorkTaskSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetSuccess(v bool) *GetWorkTaskSummaryResponseBody {
	s.Success = &v
	return s
}

type GetWorkTaskSummaryResponseBodyData struct {
	// Average response time (in minutes).
	//
	// example:
	//
	// 60
	DealAverageDuration *int64 `json:"DealAverageDuration,omitempty" xml:"DealAverageDuration,omitempty"`
	// Year-over-year growth rate of average response time.
	//
	// example:
	//
	// 20
	DealAverageDurationGrowthRate *string `json:"DealAverageDurationGrowthRate,omitempty" xml:"DealAverageDurationGrowthRate,omitempty"`
	// Number of work orders responded to.
	//
	// example:
	//
	// 100
	DealWorkTaskCount *int64 `json:"DealWorkTaskCount,omitempty" xml:"DealWorkTaskCount,omitempty"`
	// Year-over-year growth rate of the number of work orders responded to.
	//
	// example:
	//
	// 20
	DealWorkTaskCountRate *string `json:"DealWorkTaskCountRate,omitempty" xml:"DealWorkTaskCountRate,omitempty"`
	// Number of service responses.
	//
	// example:
	//
	// 10
	WorkTaskCount *int64 `json:"WorkTaskCount,omitempty" xml:"WorkTaskCount,omitempty"`
	// Problem closure rate.
	//
	// example:
	//
	// 90
	WorkTaskDealRate *string `json:"WorkTaskDealRate,omitempty" xml:"WorkTaskDealRate,omitempty"`
	// Year-over-year growth rate of problem closure rate.
	//
	// example:
	//
	// 20
	WorkTaskDealRateGrowthRate *string `json:"WorkTaskDealRateGrowthRate,omitempty" xml:"WorkTaskDealRateGrowthRate,omitempty"`
	// Year-over-year growth rate of service responses.
	//
	// example:
	//
	// 20
	WorkTaskGrowthRate *string `json:"WorkTaskGrowthRate,omitempty" xml:"WorkTaskGrowthRate,omitempty"`
}

func (s GetWorkTaskSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWorkTaskSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealAverageDuration(v int64) *GetWorkTaskSummaryResponseBodyData {
	s.DealAverageDuration = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealAverageDurationGrowthRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.DealAverageDurationGrowthRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealWorkTaskCount(v int64) *GetWorkTaskSummaryResponseBodyData {
	s.DealWorkTaskCount = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealWorkTaskCountRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.DealWorkTaskCountRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskCount(v int64) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskCount = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskDealRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskDealRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskDealRateGrowthRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskDealRateGrowthRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskGrowthRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskGrowthRate = &v
	return s
}

type GetWorkTaskSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkTaskSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkTaskSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkTaskSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryResponse) SetHeaders(v map[string]*string) *GetWorkTaskSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetWorkTaskSummaryResponse) SetStatusCode(v int32) *GetWorkTaskSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkTaskSummaryResponse) SetBody(v *GetWorkTaskSummaryResponseBody) *GetWorkTaskSummaryResponse {
	s.Body = v
	return s
}

type PageServiceCustomerRequest struct {
	// Authorization status.
	//
	// example:
	//
	// 1
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Cloud Monitoring - Alert authorization status.
	//
	// example:
	//
	// 1
	CmAuthStatus *int32 `json:"CmAuthStatus,omitempty" xml:"CmAuthStatus,omitempty"`
	// The page number of the query result, default is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// End time. The format is a Unix timestamp, which is the number of milliseconds since January 1, 1970.
	//
	// example:
	//
	// 1710641101123
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Cloud Security - Alert authorization status.
	//
	// example:
	//
	// 1
	MonitorAuthStatus *int32 `json:"MonitorAuthStatus,omitempty" xml:"MonitorAuthStatus,omitempty"`
	// Number of records per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start time. The format is a Unix timestamp, which is the number of milliseconds since January 1, 1970.
	//
	// example:
	//
	// 1710641101000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PageServiceCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerRequest) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerRequest) SetAuthStatus(v int32) *PageServiceCustomerRequest {
	s.AuthStatus = &v
	return s
}

func (s *PageServiceCustomerRequest) SetCmAuthStatus(v int32) *PageServiceCustomerRequest {
	s.CmAuthStatus = &v
	return s
}

func (s *PageServiceCustomerRequest) SetCurrentPage(v int32) *PageServiceCustomerRequest {
	s.CurrentPage = &v
	return s
}

func (s *PageServiceCustomerRequest) SetEndTime(v int64) *PageServiceCustomerRequest {
	s.EndTime = &v
	return s
}

func (s *PageServiceCustomerRequest) SetMonitorAuthStatus(v int32) *PageServiceCustomerRequest {
	s.MonitorAuthStatus = &v
	return s
}

func (s *PageServiceCustomerRequest) SetPageSize(v int32) *PageServiceCustomerRequest {
	s.PageSize = &v
	return s
}

func (s *PageServiceCustomerRequest) SetStartTime(v int64) *PageServiceCustomerRequest {
	s.StartTime = &v
	return s
}

type PageServiceCustomerResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// System error or openapi error
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query results.
	Data []*PageServiceCustomerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message. When the request is successful, it returns a success message; when the request fails, it returns the reason for the failure.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *PageServiceCustomerResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 808A307F-9513-5099-AAA5-98D4EF199140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request return status.
	//
	// - true: Success.
	//
	// - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PageServiceCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponseBody) SetCode(v string) *PageServiceCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetData(v []*PageServiceCustomerResponseBodyData) *PageServiceCustomerResponseBody {
	s.Data = v
	return s
}

func (s *PageServiceCustomerResponseBody) SetHttpStatusCode(v int32) *PageServiceCustomerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetMessage(v string) *PageServiceCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetPageInfo(v *PageServiceCustomerResponseBodyPageInfo) *PageServiceCustomerResponseBody {
	s.PageInfo = v
	return s
}

func (s *PageServiceCustomerResponseBody) SetRequestId(v string) *PageServiceCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetSuccess(v bool) *PageServiceCustomerResponseBody {
	s.Success = &v
	return s
}

type PageServiceCustomerResponseBodyData struct {
	// Customer UID.
	//
	// example:
	//
	// 1667751131382856
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// Authorization status.
	//
	// example:
	//
	// 1
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Cloud Monitoring - Alert authorization status.
	//
	// example:
	//
	// 0
	CmAuthStatus *int32 `json:"CmAuthStatus,omitempty" xml:"CmAuthStatus,omitempty"`
	// End time. The format is a Unix timestamp, which is the number of milliseconds since January 1, 1970.
	//
	// example:
	//
	// 1710123149222
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Customer level.
	//
	// example:
	//
	// GC1
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Cloud Security - Alert authorization status.
	//
	// example:
	//
	// 1
	MonitorAuthStatus *int32 `json:"MonitorAuthStatus,omitempty" xml:"MonitorAuthStatus,omitempty"`
	// Customer name.
	//
	// example:
	//
	// 中国工程院
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Owner name.
	//
	// example:
	//
	// 常温
	OwnId *string `json:"OwnId,omitempty" xml:"OwnId,omitempty"`
	// Start time. The format is a Unix timestamp, which is the number of milliseconds since January 1, 1970.
	//
	// example:
	//
	// 1710123149000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Customer ID.
	//
	// example:
	//
	// 1667751131382856
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Version information.
	//
	// example:
	//
	// 企业版
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PageServiceCustomerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerResponseBodyData) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponseBodyData) SetAliuid(v string) *PageServiceCustomerResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetAuthStatus(v int32) *PageServiceCustomerResponseBodyData {
	s.AuthStatus = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetCmAuthStatus(v int32) *PageServiceCustomerResponseBodyData {
	s.CmAuthStatus = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetEndTime(v int64) *PageServiceCustomerResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetLevel(v string) *PageServiceCustomerResponseBodyData {
	s.Level = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetMonitorAuthStatus(v int32) *PageServiceCustomerResponseBodyData {
	s.MonitorAuthStatus = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetName(v string) *PageServiceCustomerResponseBodyData {
	s.Name = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetOwnId(v string) *PageServiceCustomerResponseBodyData {
	s.OwnId = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetStartTime(v int64) *PageServiceCustomerResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetUserId(v string) *PageServiceCustomerResponseBodyData {
	s.UserId = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetVersion(v string) *PageServiceCustomerResponseBodyData {
	s.Version = &v
	return s
}

type PageServiceCustomerResponseBodyPageInfo struct {
	// The current page number in pagination queries.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of query results.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageServiceCustomerResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponseBodyPageInfo) SetCurrentPage(v int32) *PageServiceCustomerResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *PageServiceCustomerResponseBodyPageInfo) SetPageSize(v int32) *PageServiceCustomerResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *PageServiceCustomerResponseBodyPageInfo) SetTotalCount(v int32) *PageServiceCustomerResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type PageServiceCustomerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageServiceCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageServiceCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerResponse) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponse) SetHeaders(v map[string]*string) *PageServiceCustomerResponse {
	s.Headers = v
	return s
}

func (s *PageServiceCustomerResponse) SetStatusCode(v int32) *PageServiceCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *PageServiceCustomerResponse) SetBody(v *PageServiceCustomerResponseBody) *PageServiceCustomerResponse {
	s.Body = v
	return s
}

type SendCustomEventRequest struct {
	// User ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1214484929940219
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// Data source.
	//
	// example:
	//
	// aegis_suspicious_event
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// Event details.
	//
	// example:
	//
	// 疑似病毒木马启动运行。
	EventDescription *string `json:"EventDescription,omitempty" xml:"EventDescription,omitempty"`
	// Alert event details.
	//
	// example:
	//
	// [{"name":"提示","type":"text","value":"在您的系统上发现可疑进程启动行为，通常与病毒木马或入侵事件相关"},{"name":"ATT&CK攻击阶段","type":"text","value":"代码执行"},{"name":"恶意行为","type":"text","value":"可疑的漏洞利用代码执行"},{"name":"规则类型","type":"text","value":"进程启动"},{"name":"规则引擎","type":"text","value":"精准攻击识别引擎"},{"name":"处置动作","type":"text","value":"阻断行为"},{"name":"进程路径","type":"text","value":"/usr/bin/python3.9"},{"name":"命令行","type":"text","value":"python3 /root/poc/python/cve-2018-15473.py --username root --port 22"},{"name":"父进程路径","type":"text","value":"/bin/gunkit"},{"name":"父进程命令行","type":"text","value":"gunkit serve-grpc --addr unix:///data/gunkit-grpc.sock"},{"name":"进程ID","type":"text","value":"22714"},{"name":"父进程ID","type":"text","value":"2986"},{"name":"描述","type":"text","value":"主动防御检测到可疑进程启动行为，这类可疑进程通常存在于特殊的系统目录，或通过后缀伪装成文档/音频/图片等文件诱导用户运行，该异常行为已被成功拦截"},{"name":"处置建议","type":"text","value":"请您及时排查是否是正常的业务操作，如果您觉得此次拦截是非预期的，那您可以在主动防御 - 恶意行为防御页面中，关闭“可疑进程启动“规则集或者将受影响机器从管理主机中移除"},{"name":"父进程关系","type":"processChain","value":"1:::/usr/lib/systemd/systemd --switched-root --system --deserialize 22 &&& 2939:::/usr/local/bin/containerd-shim-runc-v2 -namespace moby -id 270f164903b47d4e219b410b8d11d9079a7ad1bac8133aea604598300d3b03d5 -address /run/containerd/containerd.sock &&& 2962:::/usr/bin/python3 /usr/bin/supervisord -n &&& 2986:::gunkit serve-grpc --addr unix:///data/gunkit-grpc.sock"}]
	EventDetails  *string `json:"EventDetails,omitempty" xml:"EventDetails,omitempty"`
	EventMarkdown *string `json:"EventMarkdown,omitempty" xml:"EventMarkdown,omitempty"`
	// Event name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 客户端离线
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// Event type.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUSP_CUSTOM_CFW
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Alert discovery time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-23 14:47:34
	FindTime *int64 `json:"FindTime,omitempty" xml:"FindTime,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// i-uf60h3ns25bzq9eyf8ps
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name.
	//
	// example:
	//
	// 猫吉-售卖-MDR扫描器集群1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Whether to send.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	IsSend *string `json:"IsSend,omitempty" xml:"IsSend,omitempty"`
	// Event level.
	//
	// This parameter is required.
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The first occurrence time of the alert event.
	//
	// example:
	//
	// 1724956996000
	OccurrenceTime *int64  `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Product name.
	//
	// example:
	//
	// CloudSecCenter
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// Unique ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68888f02-98f2-492b-a2b2-5b13295755b7
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// UUID.
	//
	// example:
	//
	// 93B6CDAB-7D2E-33D2-9EBA-25D561A2E95F
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s SendCustomEventRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomEventRequest) GoString() string {
	return s.String()
}

func (s *SendCustomEventRequest) SetCustomerId(v string) *SendCustomEventRequest {
	s.CustomerId = &v
	return s
}

func (s *SendCustomEventRequest) SetDataSource(v string) *SendCustomEventRequest {
	s.DataSource = &v
	return s
}

func (s *SendCustomEventRequest) SetEventDescription(v string) *SendCustomEventRequest {
	s.EventDescription = &v
	return s
}

func (s *SendCustomEventRequest) SetEventDetails(v string) *SendCustomEventRequest {
	s.EventDetails = &v
	return s
}

func (s *SendCustomEventRequest) SetEventMarkdown(v string) *SendCustomEventRequest {
	s.EventMarkdown = &v
	return s
}

func (s *SendCustomEventRequest) SetEventName(v string) *SendCustomEventRequest {
	s.EventName = &v
	return s
}

func (s *SendCustomEventRequest) SetEventType(v string) *SendCustomEventRequest {
	s.EventType = &v
	return s
}

func (s *SendCustomEventRequest) SetFindTime(v int64) *SendCustomEventRequest {
	s.FindTime = &v
	return s
}

func (s *SendCustomEventRequest) SetInstanceId(v string) *SendCustomEventRequest {
	s.InstanceId = &v
	return s
}

func (s *SendCustomEventRequest) SetInstanceName(v string) *SendCustomEventRequest {
	s.InstanceName = &v
	return s
}

func (s *SendCustomEventRequest) SetIsSend(v string) *SendCustomEventRequest {
	s.IsSend = &v
	return s
}

func (s *SendCustomEventRequest) SetLevel(v string) *SendCustomEventRequest {
	s.Level = &v
	return s
}

func (s *SendCustomEventRequest) SetOccurrenceTime(v int64) *SendCustomEventRequest {
	s.OccurrenceTime = &v
	return s
}

func (s *SendCustomEventRequest) SetOwnerId(v string) *SendCustomEventRequest {
	s.OwnerId = &v
	return s
}

func (s *SendCustomEventRequest) SetProduct(v string) *SendCustomEventRequest {
	s.Product = &v
	return s
}

func (s *SendCustomEventRequest) SetUniqueId(v string) *SendCustomEventRequest {
	s.UniqueId = &v
	return s
}

func (s *SendCustomEventRequest) SetUuid(v string) *SendCustomEventRequest {
	s.Uuid = &v
	return s
}

type SendCustomEventResponseBody struct {
	// Interface response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Interface return data.
	Data *SendCustomEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message. When the request is successful, it returns a success message; when the request fails, it returns the reason for the failure.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 606EB377-155D-5AEB-AC4F-F013444A4C45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// - true: Call succeeded.
	//
	// - false: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendCustomEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCustomEventResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomEventResponseBody) SetCode(v string) *SendCustomEventResponseBody {
	s.Code = &v
	return s
}

func (s *SendCustomEventResponseBody) SetData(v *SendCustomEventResponseBodyData) *SendCustomEventResponseBody {
	s.Data = v
	return s
}

func (s *SendCustomEventResponseBody) SetHttpStatusCode(v int32) *SendCustomEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendCustomEventResponseBody) SetMessage(v string) *SendCustomEventResponseBody {
	s.Message = &v
	return s
}

func (s *SendCustomEventResponseBody) SetRequestId(v string) *SendCustomEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomEventResponseBody) SetSuccess(v bool) *SendCustomEventResponseBody {
	s.Success = &v
	return s
}

type SendCustomEventResponseBodyData struct {
	// Service UID.
	//
	// example:
	//
	// 1601097845544644
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// Customer name.
	//
	// example:
	//
	// 天津瑞鹏昇科技发展有限公司
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// Alert ID.
	//
	// example:
	//
	// c0dc71d1-8a1d-4043-9767-f6c420e34901-81bd
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// Alert type.
	//
	// example:
	//
	// SUSP_CUSTOM_WAF
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Work order ID.
	//
	// example:
	//
	// 1914348
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Owner ID.
	//
	// example:
	//
	// 352675
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Owner name.
	//
	// example:
	//
	// 乐牙
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Work order name.
	//
	// example:
	//
	// 22端口禁止任意IP访问
	WorkTaskName *string `json:"WorkTaskName,omitempty" xml:"WorkTaskName,omitempty"`
}

func (s SendCustomEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendCustomEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendCustomEventResponseBodyData) SetCustomerId(v string) *SendCustomEventResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetCustomerName(v string) *SendCustomEventResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetEventId(v string) *SendCustomEventResponseBodyData {
	s.EventId = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetEventType(v string) *SendCustomEventResponseBodyData {
	s.EventType = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetId(v int64) *SendCustomEventResponseBodyData {
	s.Id = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetOwnerId(v string) *SendCustomEventResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetOwnerName(v string) *SendCustomEventResponseBodyData {
	s.OwnerName = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetWorkTaskName(v string) *SendCustomEventResponseBodyData {
	s.WorkTaskName = &v
	return s
}

type SendCustomEventResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCustomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendCustomEventResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCustomEventResponse) GoString() string {
	return s.String()
}

func (s *SendCustomEventResponse) SetHeaders(v map[string]*string) *SendCustomEventResponse {
	s.Headers = v
	return s
}

func (s *SendCustomEventResponse) SetStatusCode(v int32) *SendCustomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCustomEventResponse) SetBody(v *SendCustomEventResponseBody) *SendCustomEventResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("mssp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Confirm Receipt of Security Assessment Report
//
// @param request - ConfirmDjbhReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmDjbhReportResponse
func (client *Client) ConfirmDjbhReportWithOptions(request *ConfirmDjbhReportRequest, runtime *util.RuntimeOptions) (_result *ConfirmDjbhReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmDjbhReport"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConfirmDjbhReportResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConfirmDjbhReportResponse{}
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
// Confirm Receipt of Security Assessment Report
//
// @param request - ConfirmDjbhReportRequest
//
// @return ConfirmDjbhReportResponse
func (client *Client) ConfirmDjbhReport(request *ConfirmDjbhReportRequest) (_result *ConfirmDjbhReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmDjbhReportResponse{}
	_body, _err := client.ConfirmDjbhReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create Service-Linked Role
//
// @param request - CreateServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceLinkedRole"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateServiceLinkedRoleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateServiceLinkedRoleResponse{}
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
// Create Service-Linked Role
//
// @param request - CreateServiceLinkedRoleRequest
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create Service Work Order
//
// @param request - CreateServiceWorkOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceWorkOrderResponse
func (client *Client) CreateServiceWorkOrderWithOptions(request *CreateServiceWorkOrderRequest, runtime *util.RuntimeOptions) (_result *CreateServiceWorkOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		body["Creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		body["CustomerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.DurationDay)) {
		body["DurationDay"] = request.DurationDay
	}

	if !tea.BoolValue(util.IsUnset(request.IsAttachment)) {
		body["IsAttachment"] = request.IsAttachment
	}

	if !tea.BoolValue(util.IsUnset(request.IsMilestone)) {
		body["IsMilestone"] = request.IsMilestone
	}

	if !tea.BoolValue(util.IsUnset(request.IsWorkOrderNotify)) {
		body["IsWorkOrderNotify"] = request.IsWorkOrderNotify
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyDay)) {
		body["NotifyDay"] = request.NotifyDay
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyId)) {
		body["NotifyId"] = request.NotifyId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateRemark)) {
		body["OperateRemark"] = request.OperateRemark
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		body["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderDetail)) {
		body["WorkOrderDetail"] = request.WorkOrderDetail
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderName)) {
		body["WorkOrderName"] = request.WorkOrderName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderSource)) {
		body["WorkOrderSource"] = request.WorkOrderSource
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderStatus)) {
		body["WorkOrderStatus"] = request.WorkOrderStatus
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderType)) {
		body["WorkOrderType"] = request.WorkOrderType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceWorkOrder"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateServiceWorkOrderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateServiceWorkOrderResponse{}
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
// Create Service Work Order
//
// @param request - CreateServiceWorkOrderRequest
//
// @return CreateServiceWorkOrderResponse
func (client *Client) CreateServiceWorkOrder(request *CreateServiceWorkOrderRequest) (_result *CreateServiceWorkOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceWorkOrderResponse{}
	_body, _err := client.CreateServiceWorkOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete Security Assessment Report
//
// @param request - DeleteDjbhReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDjbhReportResponse
func (client *Client) DeleteDjbhReportWithOptions(request *DeleteDjbhReportRequest, runtime *util.RuntimeOptions) (_result *DeleteDjbhReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDjbhReport"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDjbhReportResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDjbhReportResponse{}
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
// Delete Security Assessment Report
//
// @param request - DeleteDjbhReportRequest
//
// @return DeleteDjbhReportResponse
func (client *Client) DeleteDjbhReport(request *DeleteDjbhReportRequest) (_result *DeleteDjbhReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDjbhReportResponse{}
	_body, _err := client.DeleteDjbhReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query Service-Linked Role
//
// @param request - DescribeServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceLinkedRoleResponse
func (client *Client) DescribeServiceLinkedRoleWithOptions(request *DescribeServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceLinkedRole"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeServiceLinkedRoleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeServiceLinkedRoleResponse{}
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
// Query Service-Linked Role
//
// @param request - DescribeServiceLinkedRoleRequest
//
// @return DescribeServiceLinkedRoleResponse
func (client *Client) DescribeServiceLinkedRole(request *DescribeServiceLinkedRoleRequest) (_result *DescribeServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceLinkedRoleResponse{}
	_body, _err := client.DescribeServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Process Service Work Order
//
// @param request - DisposeServiceWorkOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisposeServiceWorkOrderResponse
func (client *Client) DisposeServiceWorkOrderWithOptions(request *DisposeServiceWorkOrderRequest, runtime *util.RuntimeOptions) (_result *DisposeServiceWorkOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachmentName)) {
		body["AttachmentName"] = request.AttachmentName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardOwnerId)) {
		body["ForwardOwnerId"] = request.ForwardOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IsAttachment)) {
		body["IsAttachment"] = request.IsAttachment
	}

	if !tea.BoolValue(util.IsUnset(request.IsWorkOrderNotify)) {
		body["IsWorkOrderNotify"] = request.IsWorkOrderNotify
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyId)) {
		body["NotifyId"] = request.NotifyId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateRemark)) {
		body["OperateRemark"] = request.OperateRemark
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeOwnerId)) {
		body["UpgradeOwnerId"] = request.UpgradeOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderDetail)) {
		body["WorkOrderDetail"] = request.WorkOrderDetail
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderName)) {
		body["WorkOrderName"] = request.WorkOrderName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderStatus)) {
		body["WorkOrderStatus"] = request.WorkOrderStatus
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisposeServiceWorkOrder"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DisposeServiceWorkOrderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DisposeServiceWorkOrderResponse{}
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
// Process Service Work Order
//
// @param request - DisposeServiceWorkOrderRequest
//
// @return DisposeServiceWorkOrderResponse
func (client *Client) DisposeServiceWorkOrder(request *DisposeServiceWorkOrderRequest) (_result *DisposeServiceWorkOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisposeServiceWorkOrderResponse{}
	_body, _err := client.DisposeServiceWorkOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Handle Alert Work Order
//
// @param request - DisposeWorkTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisposeWorkTaskResponse
func (client *Client) DisposeWorkTaskWithOptions(request *DisposeWorkTaskRequest, runtime *util.RuntimeOptions) (_result *DisposeWorkTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.OptRemark)) {
		body["OptRemark"] = request.OptRemark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		body["TaskIds"] = request.TaskIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisposeWorkTask"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DisposeWorkTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DisposeWorkTaskResponse{}
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
// Handle Alert Work Order
//
// @param request - DisposeWorkTaskRequest
//
// @return DisposeWorkTaskResponse
func (client *Client) DisposeWorkTask(request *DisposeWorkTaskRequest) (_result *DisposeWorkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisposeWorkTaskResponse{}
	_body, _err := client.DisposeWorkTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query Alarm Details
//
// @param request - GetAlarmDetailByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlarmDetailByIdResponse
func (client *Client) GetAlarmDetailByIdWithOptions(request *GetAlarmDetailByIdRequest, runtime *util.RuntimeOptions) (_result *GetAlarmDetailByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlarmDetailById"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAlarmDetailByIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAlarmDetailByIdResponse{}
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
// Query Alarm Details
//
// @param request - GetAlarmDetailByIdRequest
//
// @return GetAlarmDetailByIdResponse
func (client *Client) GetAlarmDetailById(request *GetAlarmDetailByIdRequest) (_result *GetAlarmDetailByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlarmDetailByIdResponse{}
	_body, _err := client.GetAlarmDetailByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Trend of Attacked Asset Convergence
//
// @param request - GetAttackedAssetDealRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttackedAssetDealResponse
func (client *Client) GetAttackedAssetDealWithOptions(request *GetAttackedAssetDealRequest, runtime *util.RuntimeOptions) (_result *GetAttackedAssetDealResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAttackedAssetDeal"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAttackedAssetDealResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAttackedAssetDealResponse{}
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
// Trend of Attacked Asset Convergence
//
// @param request - GetAttackedAssetDealRequest
//
// @return GetAttackedAssetDealResponse
func (client *Client) GetAttackedAssetDeal(request *GetAttackedAssetDealRequest) (_result *GetAttackedAssetDealResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAttackedAssetDealResponse{}
	_body, _err := client.GetAttackedAssetDealWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Compliance Risk Convergence Trend
//
// @param request - GetBaselineSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBaselineSummaryResponse
func (client *Client) GetBaselineSummaryWithOptions(request *GetBaselineSummaryRequest, runtime *util.RuntimeOptions) (_result *GetBaselineSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBaselineSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetBaselineSummaryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetBaselineSummaryResponse{}
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
// Compliance Risk Convergence Trend
//
// @param request - GetBaselineSummaryRequest
//
// @return GetBaselineSummaryResponse
func (client *Client) GetBaselineSummary(request *GetBaselineSummaryRequest) (_result *GetBaselineSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBaselineSummaryResponse{}
	_body, _err := client.GetBaselineSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get Console Score
//
// @param request - GetConsoleScoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsoleScoreResponse
func (client *Client) GetConsoleScoreWithOptions(request *GetConsoleScoreRequest, runtime *util.RuntimeOptions) (_result *GetConsoleScoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConsoleScore"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetConsoleScoreResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetConsoleScoreResponse{}
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
// Get Console Score
//
// @param request - GetConsoleScoreRequest
//
// @return GetConsoleScoreResponse
func (client *Client) GetConsoleScore(request *GetConsoleScoreRequest) (_result *GetConsoleScoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConsoleScoreResponse{}
	_body, _err := client.GetConsoleScoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query Risk Details
//
// @param request - GetDetailByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetailByIdResponse
func (client *Client) GetDetailByIdWithOptions(request *GetDetailByIdRequest, runtime *util.RuntimeOptions) (_result *GetDetailByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDetailById"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDetailByIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDetailByIdResponse{}
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
// Query Risk Details
//
// @param request - GetDetailByIdRequest
//
// @return GetDetailByIdResponse
func (client *Client) GetDetailById(request *GetDetailByIdRequest) (_result *GetDetailByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetailByIdResponse{}
	_body, _err := client.GetDetailByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Single Service Report Download
//
// @param request - GetDocumentDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentDownloadUrlResponse
func (client *Client) GetDocumentDownloadUrlWithOptions(request *GetDocumentDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetDocumentDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ReportType)) {
		body["ReportType"] = request.ReportType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentDownloadUrl"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocumentDownloadUrlResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocumentDownloadUrlResponse{}
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
// Single Service Report Download
//
// @param request - GetDocumentDownloadUrlRequest
//
// @return GetDocumentDownloadUrlResponse
func (client *Client) GetDocumentDownloadUrl(request *GetDocumentDownloadUrlRequest) (_result *GetDocumentDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocumentDownloadUrlResponse{}
	_body, _err := client.GetDocumentDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Service Report Query
//
// @param request - GetDocumentPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentPageResponse
func (client *Client) GetDocumentPageWithOptions(request *GetDocumentPageRequest, runtime *util.RuntimeOptions) (_result *GetDocumentPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveredBy)) {
		body["DeliveredBy"] = request.DeliveredBy
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentName)) {
		body["DocumentName"] = request.DocumentName
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentType)) {
		body["DocumentType"] = request.DocumentType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReportType)) {
		body["ReportType"] = request.ReportType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentPage"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocumentPageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocumentPageResponse{}
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
// Service Report Query
//
// @param request - GetDocumentPageRequest
//
// @return GetDocumentPageResponse
func (client *Client) GetDocumentPage(request *GetDocumentPageRequest) (_result *GetDocumentPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocumentPageResponse{}
	_body, _err := client.GetDocumentPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Service Report Home Page Statistics Acquisition
//
// @param request - GetDocumentSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentSummaryResponse
func (client *Client) GetDocumentSummaryWithOptions(request *GetDocumentSummaryRequest, runtime *util.RuntimeOptions) (_result *GetDocumentSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportType)) {
		body["ReportType"] = request.ReportType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocumentSummaryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocumentSummaryResponse{}
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
// Service Report Home Page Statistics Acquisition
//
// @param request - GetDocumentSummaryRequest
//
// @return GetDocumentSummaryResponse
func (client *Client) GetDocumentSummary(request *GetDocumentSummaryRequest) (_result *GetDocumentSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocumentSummaryResponse{}
	_body, _err := client.GetDocumentSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get Recently Uploaded Service Reports
//
// @param request - GetRecentDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecentDocumentResponse
func (client *Client) GetRecentDocumentWithOptions(request *GetRecentDocumentRequest, runtime *util.RuntimeOptions) (_result *GetRecentDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecentDocument"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRecentDocumentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRecentDocumentResponse{}
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
// Get Recently Uploaded Service Reports
//
// @param request - GetRecentDocumentRequest
//
// @return GetRecentDocumentResponse
func (client *Client) GetRecentDocument(request *GetRecentDocumentRequest) (_result *GetRecentDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecentDocumentResponse{}
	_body, _err := client.GetRecentDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get Safety Coverage
//
// @param request - GetSafetyCoverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSafetyCoverResponse
func (client *Client) GetSafetyCoverWithOptions(request *GetSafetyCoverRequest, runtime *util.RuntimeOptions) (_result *GetSafetyCoverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSafetyCover"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSafetyCoverResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSafetyCoverResponse{}
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
// Get Safety Coverage
//
// @param request - GetSafetyCoverRequest
//
// @return GetSafetyCoverResponse
func (client *Client) GetSafetyCover(request *GetSafetyCoverRequest) (_result *GetSafetyCoverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSafetyCoverResponse{}
	_body, _err := client.GetSafetyCoverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get SOW List
//
// @param request - GetSowListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSowListResponse
func (client *Client) GetSowListWithOptions(request *GetSowListRequest, runtime *util.RuntimeOptions) (_result *GetSowListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSowList"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSowListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSowListResponse{}
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
// Get SOW List
//
// @param request - GetSowListRequest
//
// @return GetSowListResponse
func (client *Client) GetSowList(request *GetSowListRequest) (_result *GetSowListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSowListResponse{}
	_body, _err := client.GetSowListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Alarm Disposal Query
//
// @param request - GetSuspEventPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuspEventPageResponse
func (client *Client) GetSuspEventPageWithOptions(request *GetSuspEventPageRequest, runtime *util.RuntimeOptions) (_result *GetSuspEventPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmEndTime)) {
		body["AlarmEndTime"] = request.AlarmEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmStartTime)) {
		body["AlarmStartTime"] = request.AlarmStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSuspEventPage"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSuspEventPageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSuspEventPageResponse{}
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
// Alarm Disposal Query
//
// @param request - GetSuspEventPageRequest
//
// @return GetSuspEventPageResponse
func (client *Client) GetSuspEventPage(request *GetSuspEventPageRequest) (_result *GetSuspEventPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSuspEventPageResponse{}
	_body, _err := client.GetSuspEventPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get Alert Statistics
//
// @param request - GetSuspEventSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuspEventSummaryResponse
func (client *Client) GetSuspEventSummaryWithOptions(request *GetSuspEventSummaryRequest, runtime *util.RuntimeOptions) (_result *GetSuspEventSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSuspEventSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSuspEventSummaryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSuspEventSummaryResponse{}
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
// Get Alert Statistics
//
// @param request - GetSuspEventSummaryRequest
//
// @return GetSuspEventSummaryResponse
func (client *Client) GetSuspEventSummary(request *GetSuspEventSummaryRequest) (_result *GetSuspEventSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSuspEventSummaryResponse{}
	_body, _err := client.GetSuspEventSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Alarm Page Statistics
//
// @param request - GetSuspPageSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuspPageSummaryResponse
func (client *Client) GetSuspPageSummaryWithOptions(runtime *util.RuntimeOptions) (_result *GetSuspPageSummaryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetSuspPageSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSuspPageSummaryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSuspPageSummaryResponse{}
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
// Alarm Page Statistics
//
// @return GetSuspPageSummaryResponse
func (client *Client) GetSuspPageSummary() (_result *GetSuspPageSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSuspPageSummaryResponse{}
	_body, _err := client.GetSuspPageSummaryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query User Activation Status
//
// @param request - GetUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserStatusResponse
func (client *Client) GetUserStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetUserStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetUserStatus"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetUserStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetUserStatusResponse{}
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
// Query User Activation Status
//
// @return GetUserStatusResponse
func (client *Client) GetUserStatus() (_result *GetUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserStatusResponse{}
	_body, _err := client.GetUserStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Risk Query
//
// @param request - GetVulItemPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulItemPageResponse
func (client *Client) GetVulItemPageWithOptions(request *GetVulItemPageRequest, runtime *util.RuntimeOptions) (_result *GetVulItemPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["AliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Dealed)) {
		body["Dealed"] = request.Dealed
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ScanType)) {
		body["ScanType"] = request.ScanType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVulItemPage"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVulItemPageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVulItemPageResponse{}
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
// Risk Query
//
// @param request - GetVulItemPageRequest
//
// @return GetVulItemPageResponse
func (client *Client) GetVulItemPage(request *GetVulItemPageRequest) (_result *GetVulItemPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVulItemPageResponse{}
	_body, _err := client.GetVulItemPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query processed details
//
// @param request - GetVulListByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulListByIdResponse
func (client *Client) GetVulListByIdWithOptions(request *GetVulListByIdRequest, runtime *util.RuntimeOptions) (_result *GetVulListByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Dealed)) {
		body["Dealed"] = request.Dealed
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Necessity)) {
		body["Necessity"] = request.Necessity
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Uuids)) {
		body["Uuids"] = request.Uuids
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVulListById"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVulListByIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVulListByIdResponse{}
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
// Query processed details
//
// @param request - GetVulListByIdRequest
//
// @return GetVulListByIdResponse
func (client *Client) GetVulListById(request *GetVulListByIdRequest) (_result *GetVulListByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVulListByIdResponse{}
	_body, _err := client.GetVulListByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Risk Page Statistics
//
// @param request - GetVulPageSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulPageSummaryResponse
func (client *Client) GetVulPageSummaryWithOptions(runtime *util.RuntimeOptions) (_result *GetVulPageSummaryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetVulPageSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVulPageSummaryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVulPageSummaryResponse{}
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
// Risk Page Statistics
//
// @return GetVulPageSummaryResponse
func (client *Client) GetVulPageSummary() (_result *GetVulPageSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVulPageSummaryResponse{}
	_body, _err := client.GetVulPageSummaryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get Risk Statistics
//
// @param request - GetVulSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulSummaryResponse
func (client *Client) GetVulSummaryWithOptions(request *GetVulSummaryRequest, runtime *util.RuntimeOptions) (_result *GetVulSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVulSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetVulSummaryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetVulSummaryResponse{}
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
// Get Risk Statistics
//
// @param request - GetVulSummaryRequest
//
// @return GetVulSummaryResponse
func (client *Client) GetVulSummary(request *GetVulSummaryRequest) (_result *GetVulSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVulSummaryResponse{}
	_body, _err := client.GetVulSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get the First Line Work Order Statistics
//
// @param request - GetWorkTaskSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkTaskSummaryResponse
func (client *Client) GetWorkTaskSummaryWithOptions(request *GetWorkTaskSummaryRequest, runtime *util.RuntimeOptions) (_result *GetWorkTaskSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkTaskSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetWorkTaskSummaryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetWorkTaskSummaryResponse{}
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
// Get the First Line Work Order Statistics
//
// @param request - GetWorkTaskSummaryRequest
//
// @return GetWorkTaskSummaryResponse
func (client *Client) GetWorkTaskSummary(request *GetWorkTaskSummaryRequest) (_result *GetWorkTaskSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkTaskSummaryResponse{}
	_body, _err := client.GetWorkTaskSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Service Customer Information Query
//
// @param request - PageServiceCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageServiceCustomerResponse
func (client *Client) PageServiceCustomerWithOptions(request *PageServiceCustomerRequest, runtime *util.RuntimeOptions) (_result *PageServiceCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthStatus)) {
		body["AuthStatus"] = request.AuthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CmAuthStatus)) {
		body["CmAuthStatus"] = request.CmAuthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorAuthStatus)) {
		body["MonitorAuthStatus"] = request.MonitorAuthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PageServiceCustomer"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PageServiceCustomerResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PageServiceCustomerResponse{}
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
// Service Customer Information Query
//
// @param request - PageServiceCustomerRequest
//
// @return PageServiceCustomerResponse
func (client *Client) PageServiceCustomer(request *PageServiceCustomerRequest) (_result *PageServiceCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageServiceCustomerResponse{}
	_body, _err := client.PageServiceCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Send Custom Alert Event
//
// @param request - SendCustomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendCustomEventResponse
func (client *Client) SendCustomEventWithOptions(request *SendCustomEventRequest, runtime *util.RuntimeOptions) (_result *SendCustomEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		body["CustomerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		body["DataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.EventDescription)) {
		body["EventDescription"] = request.EventDescription
	}

	if !tea.BoolValue(util.IsUnset(request.EventDetails)) {
		body["EventDetails"] = request.EventDetails
	}

	if !tea.BoolValue(util.IsUnset(request.EventMarkdown)) {
		body["EventMarkdown"] = request.EventMarkdown
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		body["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		body["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.FindTime)) {
		body["FindTime"] = request.FindTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.IsSend)) {
		body["IsSend"] = request.IsSend
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.OccurrenceTime)) {
		body["OccurrenceTime"] = request.OccurrenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		body["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		body["UniqueId"] = request.UniqueId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCustomEvent"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SendCustomEventResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SendCustomEventResponse{}
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
// Send Custom Alert Event
//
// @param request - SendCustomEventRequest
//
// @return SendCustomEventResponse
func (client *Client) SendCustomEvent(request *SendCustomEventRequest) (_result *SendCustomEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCustomEventResponse{}
	_body, _err := client.SendCustomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
