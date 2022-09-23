// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AcceptJoinPermissionAuthOrderRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s AcceptJoinPermissionAuthOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptJoinPermissionAuthOrderRequest) GoString() string {
	return s.String()
}

func (s *AcceptJoinPermissionAuthOrderRequest) SetOrderId(v string) *AcceptJoinPermissionAuthOrderRequest {
	s.OrderId = &v
	return s
}

type AcceptJoinPermissionAuthOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AcceptJoinPermissionAuthOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptJoinPermissionAuthOrderResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptJoinPermissionAuthOrderResponseBody) SetRequestId(v string) *AcceptJoinPermissionAuthOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptJoinPermissionAuthOrderResponseBody) SetSuccess(v bool) *AcceptJoinPermissionAuthOrderResponseBody {
	s.Success = &v
	return s
}

type AcceptJoinPermissionAuthOrderResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AcceptJoinPermissionAuthOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AcceptJoinPermissionAuthOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptJoinPermissionAuthOrderResponse) GoString() string {
	return s.String()
}

func (s *AcceptJoinPermissionAuthOrderResponse) SetHeaders(v map[string]*string) *AcceptJoinPermissionAuthOrderResponse {
	s.Headers = v
	return s
}

func (s *AcceptJoinPermissionAuthOrderResponse) SetStatusCode(v int32) *AcceptJoinPermissionAuthOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptJoinPermissionAuthOrderResponse) SetBody(v *AcceptJoinPermissionAuthOrderResponseBody) *AcceptJoinPermissionAuthOrderResponse {
	s.Body = v
	return s
}

type AddNodeToGroupRequest struct {
	DevEui      *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	PinCode     *string `json:"PinCode,omitempty" xml:"PinCode,omitempty"`
}

func (s AddNodeToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddNodeToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddNodeToGroupRequest) SetDevEui(v string) *AddNodeToGroupRequest {
	s.DevEui = &v
	return s
}

func (s *AddNodeToGroupRequest) SetNodeGroupId(v string) *AddNodeToGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *AddNodeToGroupRequest) SetPinCode(v string) *AddNodeToGroupRequest {
	s.PinCode = &v
	return s
}

type AddNodeToGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddNodeToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddNodeToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddNodeToGroupResponseBody) SetRequestId(v string) *AddNodeToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddNodeToGroupResponseBody) SetSuccess(v bool) *AddNodeToGroupResponseBody {
	s.Success = &v
	return s
}

type AddNodeToGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddNodeToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddNodeToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddNodeToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddNodeToGroupResponse) SetHeaders(v map[string]*string) *AddNodeToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddNodeToGroupResponse) SetStatusCode(v int32) *AddNodeToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddNodeToGroupResponse) SetBody(v *AddNodeToGroupResponseBody) *AddNodeToGroupResponse {
	s.Body = v
	return s
}

type ApplyRoamingJoinPermissionRequest struct {
	ClassMode           *string `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	DataRate            *int64  `json:"DataRate,omitempty" xml:"DataRate,omitempty"`
	FreqBandPlanGroupId *int64  `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	JoinPermissionName  *string `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	RxDelay             *int64  `json:"RxDelay,omitempty" xml:"RxDelay,omitempty"`
}

func (s ApplyRoamingJoinPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyRoamingJoinPermissionRequest) GoString() string {
	return s.String()
}

func (s *ApplyRoamingJoinPermissionRequest) SetClassMode(v string) *ApplyRoamingJoinPermissionRequest {
	s.ClassMode = &v
	return s
}

func (s *ApplyRoamingJoinPermissionRequest) SetDataRate(v int64) *ApplyRoamingJoinPermissionRequest {
	s.DataRate = &v
	return s
}

func (s *ApplyRoamingJoinPermissionRequest) SetFreqBandPlanGroupId(v int64) *ApplyRoamingJoinPermissionRequest {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *ApplyRoamingJoinPermissionRequest) SetJoinPermissionName(v string) *ApplyRoamingJoinPermissionRequest {
	s.JoinPermissionName = &v
	return s
}

func (s *ApplyRoamingJoinPermissionRequest) SetRxDelay(v int64) *ApplyRoamingJoinPermissionRequest {
	s.RxDelay = &v
	return s
}

type ApplyRoamingJoinPermissionResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyRoamingJoinPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyRoamingJoinPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyRoamingJoinPermissionResponseBody) SetData(v string) *ApplyRoamingJoinPermissionResponseBody {
	s.Data = &v
	return s
}

func (s *ApplyRoamingJoinPermissionResponseBody) SetRequestId(v string) *ApplyRoamingJoinPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyRoamingJoinPermissionResponseBody) SetSuccess(v bool) *ApplyRoamingJoinPermissionResponseBody {
	s.Success = &v
	return s
}

type ApplyRoamingJoinPermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyRoamingJoinPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyRoamingJoinPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyRoamingJoinPermissionResponse) GoString() string {
	return s.String()
}

func (s *ApplyRoamingJoinPermissionResponse) SetHeaders(v map[string]*string) *ApplyRoamingJoinPermissionResponse {
	s.Headers = v
	return s
}

func (s *ApplyRoamingJoinPermissionResponse) SetStatusCode(v int32) *ApplyRoamingJoinPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyRoamingJoinPermissionResponse) SetBody(v *ApplyRoamingJoinPermissionResponseBody) *ApplyRoamingJoinPermissionResponse {
	s.Body = v
	return s
}

type BindJoinPermissionToNodeGroupRequest struct {
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	NodeGroupId      *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s BindJoinPermissionToNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s BindJoinPermissionToNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *BindJoinPermissionToNodeGroupRequest) SetJoinPermissionId(v string) *BindJoinPermissionToNodeGroupRequest {
	s.JoinPermissionId = &v
	return s
}

func (s *BindJoinPermissionToNodeGroupRequest) SetNodeGroupId(v string) *BindJoinPermissionToNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

type BindJoinPermissionToNodeGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindJoinPermissionToNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindJoinPermissionToNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *BindJoinPermissionToNodeGroupResponseBody) SetRequestId(v string) *BindJoinPermissionToNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindJoinPermissionToNodeGroupResponseBody) SetSuccess(v bool) *BindJoinPermissionToNodeGroupResponseBody {
	s.Success = &v
	return s
}

type BindJoinPermissionToNodeGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindJoinPermissionToNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindJoinPermissionToNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s BindJoinPermissionToNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *BindJoinPermissionToNodeGroupResponse) SetHeaders(v map[string]*string) *BindJoinPermissionToNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *BindJoinPermissionToNodeGroupResponse) SetStatusCode(v int32) *BindJoinPermissionToNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *BindJoinPermissionToNodeGroupResponse) SetBody(v *BindJoinPermissionToNodeGroupResponseBody) *BindJoinPermissionToNodeGroupResponse {
	s.Body = v
	return s
}

type CancelJoinPermissionAuthOrderRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CancelJoinPermissionAuthOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelJoinPermissionAuthOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelJoinPermissionAuthOrderRequest) SetOrderId(v string) *CancelJoinPermissionAuthOrderRequest {
	s.OrderId = &v
	return s
}

type CancelJoinPermissionAuthOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelJoinPermissionAuthOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelJoinPermissionAuthOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelJoinPermissionAuthOrderResponseBody) SetRequestId(v string) *CancelJoinPermissionAuthOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelJoinPermissionAuthOrderResponseBody) SetSuccess(v bool) *CancelJoinPermissionAuthOrderResponseBody {
	s.Success = &v
	return s
}

type CancelJoinPermissionAuthOrderResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelJoinPermissionAuthOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelJoinPermissionAuthOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelJoinPermissionAuthOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelJoinPermissionAuthOrderResponse) SetHeaders(v map[string]*string) *CancelJoinPermissionAuthOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelJoinPermissionAuthOrderResponse) SetStatusCode(v int32) *CancelJoinPermissionAuthOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelJoinPermissionAuthOrderResponse) SetBody(v *CancelJoinPermissionAuthOrderResponseBody) *CancelJoinPermissionAuthOrderResponse {
	s.Body = v
	return s
}

type CheckCloudProductOpenStatusRequest struct {
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s CheckCloudProductOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudProductOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckCloudProductOpenStatusRequest) SetServiceCode(v string) *CheckCloudProductOpenStatusRequest {
	s.ServiceCode = &v
	return s
}

type CheckCloudProductOpenStatusResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckCloudProductOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudProductOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCloudProductOpenStatusResponseBody) SetData(v bool) *CheckCloudProductOpenStatusResponseBody {
	s.Data = &v
	return s
}

func (s *CheckCloudProductOpenStatusResponseBody) SetRequestId(v string) *CheckCloudProductOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCloudProductOpenStatusResponseBody) SetSuccess(v bool) *CheckCloudProductOpenStatusResponseBody {
	s.Success = &v
	return s
}

type CheckCloudProductOpenStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckCloudProductOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckCloudProductOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudProductOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckCloudProductOpenStatusResponse) SetHeaders(v map[string]*string) *CheckCloudProductOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckCloudProductOpenStatusResponse) SetStatusCode(v int32) *CheckCloudProductOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCloudProductOpenStatusResponse) SetBody(v *CheckCloudProductOpenStatusResponseBody) *CheckCloudProductOpenStatusResponse {
	s.Body = v
	return s
}

type CheckUserChargeStatusResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckUserChargeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserChargeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserChargeStatusResponseBody) SetData(v string) *CheckUserChargeStatusResponseBody {
	s.Data = &v
	return s
}

func (s *CheckUserChargeStatusResponseBody) SetRequestId(v string) *CheckUserChargeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUserChargeStatusResponseBody) SetSuccess(v bool) *CheckUserChargeStatusResponseBody {
	s.Success = &v
	return s
}

type CheckUserChargeStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckUserChargeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckUserChargeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserChargeStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckUserChargeStatusResponse) SetHeaders(v map[string]*string) *CheckUserChargeStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckUserChargeStatusResponse) SetStatusCode(v int32) *CheckUserChargeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserChargeStatusResponse) SetBody(v *CheckUserChargeStatusResponseBody) *CheckUserChargeStatusResponse {
	s.Body = v
	return s
}

type CountGatewayTupleOrdersRequest struct {
	States []*string `json:"States,omitempty" xml:"States,omitempty" type:"Repeated"`
}

func (s CountGatewayTupleOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s CountGatewayTupleOrdersRequest) GoString() string {
	return s.String()
}

func (s *CountGatewayTupleOrdersRequest) SetStates(v []*string) *CountGatewayTupleOrdersRequest {
	s.States = v
	return s
}

type CountGatewayTupleOrdersResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CountGatewayTupleOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountGatewayTupleOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *CountGatewayTupleOrdersResponseBody) SetData(v int64) *CountGatewayTupleOrdersResponseBody {
	s.Data = &v
	return s
}

func (s *CountGatewayTupleOrdersResponseBody) SetRequestId(v string) *CountGatewayTupleOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountGatewayTupleOrdersResponseBody) SetSuccess(v bool) *CountGatewayTupleOrdersResponseBody {
	s.Success = &v
	return s
}

type CountGatewayTupleOrdersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountGatewayTupleOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountGatewayTupleOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s CountGatewayTupleOrdersResponse) GoString() string {
	return s.String()
}

func (s *CountGatewayTupleOrdersResponse) SetHeaders(v map[string]*string) *CountGatewayTupleOrdersResponse {
	s.Headers = v
	return s
}

func (s *CountGatewayTupleOrdersResponse) SetStatusCode(v int32) *CountGatewayTupleOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *CountGatewayTupleOrdersResponse) SetBody(v *CountGatewayTupleOrdersResponseBody) *CountGatewayTupleOrdersResponse {
	s.Body = v
	return s
}

type CountGatewaysRequest struct {
	FreqBandPlanGroupId *int64  `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	FuzzyCity           *string `json:"FuzzyCity,omitempty" xml:"FuzzyCity,omitempty"`
	FuzzyGwEui          *string `json:"FuzzyGwEui,omitempty" xml:"FuzzyGwEui,omitempty"`
	FuzzyName           *string `json:"FuzzyName,omitempty" xml:"FuzzyName,omitempty"`
	IotInstanceId       *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IsEnabled           *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	OnlineState         *string `json:"OnlineState,omitempty" xml:"OnlineState,omitempty"`
}

func (s CountGatewaysRequest) String() string {
	return tea.Prettify(s)
}

func (s CountGatewaysRequest) GoString() string {
	return s.String()
}

func (s *CountGatewaysRequest) SetFreqBandPlanGroupId(v int64) *CountGatewaysRequest {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *CountGatewaysRequest) SetFuzzyCity(v string) *CountGatewaysRequest {
	s.FuzzyCity = &v
	return s
}

func (s *CountGatewaysRequest) SetFuzzyGwEui(v string) *CountGatewaysRequest {
	s.FuzzyGwEui = &v
	return s
}

func (s *CountGatewaysRequest) SetFuzzyName(v string) *CountGatewaysRequest {
	s.FuzzyName = &v
	return s
}

func (s *CountGatewaysRequest) SetIotInstanceId(v string) *CountGatewaysRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CountGatewaysRequest) SetIsEnabled(v bool) *CountGatewaysRequest {
	s.IsEnabled = &v
	return s
}

func (s *CountGatewaysRequest) SetOnlineState(v string) *CountGatewaysRequest {
	s.OnlineState = &v
	return s
}

type CountGatewaysResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CountGatewaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *CountGatewaysResponseBody) SetData(v int64) *CountGatewaysResponseBody {
	s.Data = &v
	return s
}

func (s *CountGatewaysResponseBody) SetRequestId(v string) *CountGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountGatewaysResponseBody) SetSuccess(v bool) *CountGatewaysResponseBody {
	s.Success = &v
	return s
}

type CountGatewaysResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountGatewaysResponse) String() string {
	return tea.Prettify(s)
}

func (s CountGatewaysResponse) GoString() string {
	return s.String()
}

func (s *CountGatewaysResponse) SetHeaders(v map[string]*string) *CountGatewaysResponse {
	s.Headers = v
	return s
}

func (s *CountGatewaysResponse) SetStatusCode(v int32) *CountGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *CountGatewaysResponse) SetBody(v *CountGatewaysResponseBody) *CountGatewaysResponse {
	s.Body = v
	return s
}

type CountNodeGroupsRequest struct {
	FuzzyDevEui   *string `json:"FuzzyDevEui,omitempty" xml:"FuzzyDevEui,omitempty"`
	FuzzyJoinEui  *string `json:"FuzzyJoinEui,omitempty" xml:"FuzzyJoinEui,omitempty"`
	FuzzyName     *string `json:"FuzzyName,omitempty" xml:"FuzzyName,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s CountNodeGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountNodeGroupsRequest) GoString() string {
	return s.String()
}

func (s *CountNodeGroupsRequest) SetFuzzyDevEui(v string) *CountNodeGroupsRequest {
	s.FuzzyDevEui = &v
	return s
}

func (s *CountNodeGroupsRequest) SetFuzzyJoinEui(v string) *CountNodeGroupsRequest {
	s.FuzzyJoinEui = &v
	return s
}

func (s *CountNodeGroupsRequest) SetFuzzyName(v string) *CountNodeGroupsRequest {
	s.FuzzyName = &v
	return s
}

func (s *CountNodeGroupsRequest) SetIotInstanceId(v string) *CountNodeGroupsRequest {
	s.IotInstanceId = &v
	return s
}

type CountNodeGroupsResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CountNodeGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountNodeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CountNodeGroupsResponseBody) SetData(v int64) *CountNodeGroupsResponseBody {
	s.Data = &v
	return s
}

func (s *CountNodeGroupsResponseBody) SetRequestId(v string) *CountNodeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountNodeGroupsResponseBody) SetSuccess(v bool) *CountNodeGroupsResponseBody {
	s.Success = &v
	return s
}

type CountNodeGroupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountNodeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountNodeGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountNodeGroupsResponse) GoString() string {
	return s.String()
}

func (s *CountNodeGroupsResponse) SetHeaders(v map[string]*string) *CountNodeGroupsResponse {
	s.Headers = v
	return s
}

func (s *CountNodeGroupsResponse) SetStatusCode(v int32) *CountNodeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountNodeGroupsResponse) SetBody(v *CountNodeGroupsResponseBody) *CountNodeGroupsResponse {
	s.Body = v
	return s
}

type CountNodeTupleOrdersRequest struct {
	IsKpm  *bool     `json:"IsKpm,omitempty" xml:"IsKpm,omitempty"`
	States []*string `json:"States,omitempty" xml:"States,omitempty" type:"Repeated"`
}

func (s CountNodeTupleOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s CountNodeTupleOrdersRequest) GoString() string {
	return s.String()
}

func (s *CountNodeTupleOrdersRequest) SetIsKpm(v bool) *CountNodeTupleOrdersRequest {
	s.IsKpm = &v
	return s
}

func (s *CountNodeTupleOrdersRequest) SetStates(v []*string) *CountNodeTupleOrdersRequest {
	s.States = v
	return s
}

type CountNodeTupleOrdersResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CountNodeTupleOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountNodeTupleOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *CountNodeTupleOrdersResponseBody) SetData(v int64) *CountNodeTupleOrdersResponseBody {
	s.Data = &v
	return s
}

func (s *CountNodeTupleOrdersResponseBody) SetRequestId(v string) *CountNodeTupleOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountNodeTupleOrdersResponseBody) SetSuccess(v bool) *CountNodeTupleOrdersResponseBody {
	s.Success = &v
	return s
}

type CountNodeTupleOrdersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountNodeTupleOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountNodeTupleOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s CountNodeTupleOrdersResponse) GoString() string {
	return s.String()
}

func (s *CountNodeTupleOrdersResponse) SetHeaders(v map[string]*string) *CountNodeTupleOrdersResponse {
	s.Headers = v
	return s
}

func (s *CountNodeTupleOrdersResponse) SetStatusCode(v int32) *CountNodeTupleOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *CountNodeTupleOrdersResponse) SetBody(v *CountNodeTupleOrdersResponseBody) *CountNodeTupleOrdersResponse {
	s.Body = v
	return s
}

type CountNodesByNodeGroupIdRequest struct {
	FuzzyDevEui *string `json:"FuzzyDevEui,omitempty" xml:"FuzzyDevEui,omitempty"`
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s CountNodesByNodeGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CountNodesByNodeGroupIdRequest) GoString() string {
	return s.String()
}

func (s *CountNodesByNodeGroupIdRequest) SetFuzzyDevEui(v string) *CountNodesByNodeGroupIdRequest {
	s.FuzzyDevEui = &v
	return s
}

func (s *CountNodesByNodeGroupIdRequest) SetNodeGroupId(v string) *CountNodesByNodeGroupIdRequest {
	s.NodeGroupId = &v
	return s
}

type CountNodesByNodeGroupIdResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CountNodesByNodeGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountNodesByNodeGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *CountNodesByNodeGroupIdResponseBody) SetData(v int64) *CountNodesByNodeGroupIdResponseBody {
	s.Data = &v
	return s
}

func (s *CountNodesByNodeGroupIdResponseBody) SetRequestId(v string) *CountNodesByNodeGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountNodesByNodeGroupIdResponseBody) SetSuccess(v bool) *CountNodesByNodeGroupIdResponseBody {
	s.Success = &v
	return s
}

type CountNodesByNodeGroupIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountNodesByNodeGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountNodesByNodeGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CountNodesByNodeGroupIdResponse) GoString() string {
	return s.String()
}

func (s *CountNodesByNodeGroupIdResponse) SetHeaders(v map[string]*string) *CountNodesByNodeGroupIdResponse {
	s.Headers = v
	return s
}

func (s *CountNodesByNodeGroupIdResponse) SetStatusCode(v int32) *CountNodesByNodeGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CountNodesByNodeGroupIdResponse) SetBody(v *CountNodesByNodeGroupIdResponseBody) *CountNodesByNodeGroupIdResponse {
	s.Body = v
	return s
}

type CountNodesByOwnedJoinPermissionIdRequest struct {
	FuzzyDevEui      *string `json:"FuzzyDevEui,omitempty" xml:"FuzzyDevEui,omitempty"`
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
}

func (s CountNodesByOwnedJoinPermissionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CountNodesByOwnedJoinPermissionIdRequest) GoString() string {
	return s.String()
}

func (s *CountNodesByOwnedJoinPermissionIdRequest) SetFuzzyDevEui(v string) *CountNodesByOwnedJoinPermissionIdRequest {
	s.FuzzyDevEui = &v
	return s
}

func (s *CountNodesByOwnedJoinPermissionIdRequest) SetJoinPermissionId(v string) *CountNodesByOwnedJoinPermissionIdRequest {
	s.JoinPermissionId = &v
	return s
}

type CountNodesByOwnedJoinPermissionIdResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CountNodesByOwnedJoinPermissionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountNodesByOwnedJoinPermissionIdResponseBody) GoString() string {
	return s.String()
}

func (s *CountNodesByOwnedJoinPermissionIdResponseBody) SetData(v int64) *CountNodesByOwnedJoinPermissionIdResponseBody {
	s.Data = &v
	return s
}

func (s *CountNodesByOwnedJoinPermissionIdResponseBody) SetRequestId(v string) *CountNodesByOwnedJoinPermissionIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountNodesByOwnedJoinPermissionIdResponseBody) SetSuccess(v bool) *CountNodesByOwnedJoinPermissionIdResponseBody {
	s.Success = &v
	return s
}

type CountNodesByOwnedJoinPermissionIdResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountNodesByOwnedJoinPermissionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountNodesByOwnedJoinPermissionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CountNodesByOwnedJoinPermissionIdResponse) GoString() string {
	return s.String()
}

func (s *CountNodesByOwnedJoinPermissionIdResponse) SetHeaders(v map[string]*string) *CountNodesByOwnedJoinPermissionIdResponse {
	s.Headers = v
	return s
}

func (s *CountNodesByOwnedJoinPermissionIdResponse) SetStatusCode(v int32) *CountNodesByOwnedJoinPermissionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CountNodesByOwnedJoinPermissionIdResponse) SetBody(v *CountNodesByOwnedJoinPermissionIdResponseBody) *CountNodesByOwnedJoinPermissionIdResponse {
	s.Body = v
	return s
}

type CountNotificationsRequest struct {
	BeginMillis *int64    `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	Category    []*string `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
	EndMillis   *int64    `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	HandleState *string   `json:"HandleState,omitempty" xml:"HandleState,omitempty"`
}

func (s CountNotificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountNotificationsRequest) GoString() string {
	return s.String()
}

func (s *CountNotificationsRequest) SetBeginMillis(v int64) *CountNotificationsRequest {
	s.BeginMillis = &v
	return s
}

func (s *CountNotificationsRequest) SetCategory(v []*string) *CountNotificationsRequest {
	s.Category = v
	return s
}

func (s *CountNotificationsRequest) SetEndMillis(v int64) *CountNotificationsRequest {
	s.EndMillis = &v
	return s
}

func (s *CountNotificationsRequest) SetHandleState(v string) *CountNotificationsRequest {
	s.HandleState = &v
	return s
}

type CountNotificationsResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CountNotificationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountNotificationsResponseBody) GoString() string {
	return s.String()
}

func (s *CountNotificationsResponseBody) SetData(v int64) *CountNotificationsResponseBody {
	s.Data = &v
	return s
}

func (s *CountNotificationsResponseBody) SetRequestId(v string) *CountNotificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountNotificationsResponseBody) SetSuccess(v bool) *CountNotificationsResponseBody {
	s.Success = &v
	return s
}

type CountNotificationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountNotificationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountNotificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountNotificationsResponse) GoString() string {
	return s.String()
}

func (s *CountNotificationsResponse) SetHeaders(v map[string]*string) *CountNotificationsResponse {
	s.Headers = v
	return s
}

func (s *CountNotificationsResponse) SetStatusCode(v int32) *CountNotificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountNotificationsResponse) SetBody(v *CountNotificationsResponseBody) *CountNotificationsResponse {
	s.Body = v
	return s
}

type CountOwnedJoinPermissionsRequest struct {
	Enabled                 *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FuzzyJoinEui            *string `json:"FuzzyJoinEui,omitempty" xml:"FuzzyJoinEui,omitempty"`
	FuzzyJoinPermissionName *string `json:"FuzzyJoinPermissionName,omitempty" xml:"FuzzyJoinPermissionName,omitempty"`
	FuzzyRenterAliyunId     *string `json:"FuzzyRenterAliyunId,omitempty" xml:"FuzzyRenterAliyunId,omitempty"`
}

func (s CountOwnedJoinPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountOwnedJoinPermissionsRequest) GoString() string {
	return s.String()
}

func (s *CountOwnedJoinPermissionsRequest) SetEnabled(v bool) *CountOwnedJoinPermissionsRequest {
	s.Enabled = &v
	return s
}

func (s *CountOwnedJoinPermissionsRequest) SetFuzzyJoinEui(v string) *CountOwnedJoinPermissionsRequest {
	s.FuzzyJoinEui = &v
	return s
}

func (s *CountOwnedJoinPermissionsRequest) SetFuzzyJoinPermissionName(v string) *CountOwnedJoinPermissionsRequest {
	s.FuzzyJoinPermissionName = &v
	return s
}

func (s *CountOwnedJoinPermissionsRequest) SetFuzzyRenterAliyunId(v string) *CountOwnedJoinPermissionsRequest {
	s.FuzzyRenterAliyunId = &v
	return s
}

type CountOwnedJoinPermissionsResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CountOwnedJoinPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountOwnedJoinPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *CountOwnedJoinPermissionsResponseBody) SetData(v int64) *CountOwnedJoinPermissionsResponseBody {
	s.Data = &v
	return s
}

func (s *CountOwnedJoinPermissionsResponseBody) SetRequestId(v string) *CountOwnedJoinPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountOwnedJoinPermissionsResponseBody) SetSuccess(v bool) *CountOwnedJoinPermissionsResponseBody {
	s.Success = &v
	return s
}

type CountOwnedJoinPermissionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountOwnedJoinPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountOwnedJoinPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountOwnedJoinPermissionsResponse) GoString() string {
	return s.String()
}

func (s *CountOwnedJoinPermissionsResponse) SetHeaders(v map[string]*string) *CountOwnedJoinPermissionsResponse {
	s.Headers = v
	return s
}

func (s *CountOwnedJoinPermissionsResponse) SetStatusCode(v int32) *CountOwnedJoinPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountOwnedJoinPermissionsResponse) SetBody(v *CountOwnedJoinPermissionsResponseBody) *CountOwnedJoinPermissionsResponse {
	s.Body = v
	return s
}

type CountRentedJoinPermissionsRequest struct {
	BoundNodeGroup          *bool   `json:"BoundNodeGroup,omitempty" xml:"BoundNodeGroup,omitempty"`
	Enabled                 *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FuzzyJoinEui            *string `json:"FuzzyJoinEui,omitempty" xml:"FuzzyJoinEui,omitempty"`
	FuzzyJoinPermissionName *string `json:"FuzzyJoinPermissionName,omitempty" xml:"FuzzyJoinPermissionName,omitempty"`
	FuzzyOwnerAliyunId      *string `json:"FuzzyOwnerAliyunId,omitempty" xml:"FuzzyOwnerAliyunId,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CountRentedJoinPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountRentedJoinPermissionsRequest) GoString() string {
	return s.String()
}

func (s *CountRentedJoinPermissionsRequest) SetBoundNodeGroup(v bool) *CountRentedJoinPermissionsRequest {
	s.BoundNodeGroup = &v
	return s
}

func (s *CountRentedJoinPermissionsRequest) SetEnabled(v bool) *CountRentedJoinPermissionsRequest {
	s.Enabled = &v
	return s
}

func (s *CountRentedJoinPermissionsRequest) SetFuzzyJoinEui(v string) *CountRentedJoinPermissionsRequest {
	s.FuzzyJoinEui = &v
	return s
}

func (s *CountRentedJoinPermissionsRequest) SetFuzzyJoinPermissionName(v string) *CountRentedJoinPermissionsRequest {
	s.FuzzyJoinPermissionName = &v
	return s
}

func (s *CountRentedJoinPermissionsRequest) SetFuzzyOwnerAliyunId(v string) *CountRentedJoinPermissionsRequest {
	s.FuzzyOwnerAliyunId = &v
	return s
}

func (s *CountRentedJoinPermissionsRequest) SetType(v string) *CountRentedJoinPermissionsRequest {
	s.Type = &v
	return s
}

type CountRentedJoinPermissionsResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CountRentedJoinPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountRentedJoinPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *CountRentedJoinPermissionsResponseBody) SetData(v int64) *CountRentedJoinPermissionsResponseBody {
	s.Data = &v
	return s
}

func (s *CountRentedJoinPermissionsResponseBody) SetRequestId(v string) *CountRentedJoinPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountRentedJoinPermissionsResponseBody) SetSuccess(v bool) *CountRentedJoinPermissionsResponseBody {
	s.Success = &v
	return s
}

type CountRentedJoinPermissionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CountRentedJoinPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CountRentedJoinPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountRentedJoinPermissionsResponse) GoString() string {
	return s.String()
}

func (s *CountRentedJoinPermissionsResponse) SetHeaders(v map[string]*string) *CountRentedJoinPermissionsResponse {
	s.Headers = v
	return s
}

func (s *CountRentedJoinPermissionsResponse) SetStatusCode(v int32) *CountRentedJoinPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountRentedJoinPermissionsResponse) SetBody(v *CountRentedJoinPermissionsResponseBody) *CountRentedJoinPermissionsResponse {
	s.Body = v
	return s
}

type CreateGatewayRequest struct {
	Address             *string  `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressCode         *int64   `json:"AddressCode,omitempty" xml:"AddressCode,omitempty"`
	City                *string  `json:"City,omitempty" xml:"City,omitempty"`
	CommunicationMode   *string  `json:"CommunicationMode,omitempty" xml:"CommunicationMode,omitempty"`
	Description         *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	District            *string  `json:"District,omitempty" xml:"District,omitempty"`
	FreqBandPlanGroupId *int64   `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	GisCoordinateSystem *string  `json:"GisCoordinateSystem,omitempty" xml:"GisCoordinateSystem,omitempty"`
	GwEui               *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId       *string  `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Latitude            *float32 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude           *float32 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name                *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	PinCode             *string  `json:"PinCode,omitempty" xml:"PinCode,omitempty"`
}

func (s CreateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequest) SetAddress(v string) *CreateGatewayRequest {
	s.Address = &v
	return s
}

func (s *CreateGatewayRequest) SetAddressCode(v int64) *CreateGatewayRequest {
	s.AddressCode = &v
	return s
}

func (s *CreateGatewayRequest) SetCity(v string) *CreateGatewayRequest {
	s.City = &v
	return s
}

func (s *CreateGatewayRequest) SetCommunicationMode(v string) *CreateGatewayRequest {
	s.CommunicationMode = &v
	return s
}

func (s *CreateGatewayRequest) SetDescription(v string) *CreateGatewayRequest {
	s.Description = &v
	return s
}

func (s *CreateGatewayRequest) SetDistrict(v string) *CreateGatewayRequest {
	s.District = &v
	return s
}

func (s *CreateGatewayRequest) SetFreqBandPlanGroupId(v int64) *CreateGatewayRequest {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *CreateGatewayRequest) SetGisCoordinateSystem(v string) *CreateGatewayRequest {
	s.GisCoordinateSystem = &v
	return s
}

func (s *CreateGatewayRequest) SetGwEui(v string) *CreateGatewayRequest {
	s.GwEui = &v
	return s
}

func (s *CreateGatewayRequest) SetIotInstanceId(v string) *CreateGatewayRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateGatewayRequest) SetLatitude(v float32) *CreateGatewayRequest {
	s.Latitude = &v
	return s
}

func (s *CreateGatewayRequest) SetLongitude(v float32) *CreateGatewayRequest {
	s.Longitude = &v
	return s
}

func (s *CreateGatewayRequest) SetName(v string) *CreateGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateGatewayRequest) SetPinCode(v string) *CreateGatewayRequest {
	s.PinCode = &v
	return s
}

type CreateGatewayResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBody) SetRequestId(v string) *CreateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetSuccess(v bool) *CreateGatewayResponseBody {
	s.Success = &v
	return s
}

type CreateGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponse) SetHeaders(v map[string]*string) *CreateGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayResponse) SetStatusCode(v int32) *CreateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayResponse) SetBody(v *CreateGatewayResponseBody) *CreateGatewayResponse {
	s.Body = v
	return s
}

type CreateLocalJoinPermissionRequest struct {
	ClassMode           *string `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	DataRate            *int64  `json:"DataRate,omitempty" xml:"DataRate,omitempty"`
	FreqBandPlanGroupId *int64  `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	IotInstanceId       *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JoinEui             *string `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionName  *string `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	RxDelay             *int64  `json:"RxDelay,omitempty" xml:"RxDelay,omitempty"`
	UseDefaultJoinEui   *bool   `json:"UseDefaultJoinEui,omitempty" xml:"UseDefaultJoinEui,omitempty"`
}

func (s CreateLocalJoinPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalJoinPermissionRequest) GoString() string {
	return s.String()
}

func (s *CreateLocalJoinPermissionRequest) SetClassMode(v string) *CreateLocalJoinPermissionRequest {
	s.ClassMode = &v
	return s
}

func (s *CreateLocalJoinPermissionRequest) SetDataRate(v int64) *CreateLocalJoinPermissionRequest {
	s.DataRate = &v
	return s
}

func (s *CreateLocalJoinPermissionRequest) SetFreqBandPlanGroupId(v int64) *CreateLocalJoinPermissionRequest {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *CreateLocalJoinPermissionRequest) SetIotInstanceId(v string) *CreateLocalJoinPermissionRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreateLocalJoinPermissionRequest) SetJoinEui(v string) *CreateLocalJoinPermissionRequest {
	s.JoinEui = &v
	return s
}

func (s *CreateLocalJoinPermissionRequest) SetJoinPermissionName(v string) *CreateLocalJoinPermissionRequest {
	s.JoinPermissionName = &v
	return s
}

func (s *CreateLocalJoinPermissionRequest) SetRxDelay(v int64) *CreateLocalJoinPermissionRequest {
	s.RxDelay = &v
	return s
}

func (s *CreateLocalJoinPermissionRequest) SetUseDefaultJoinEui(v bool) *CreateLocalJoinPermissionRequest {
	s.UseDefaultJoinEui = &v
	return s
}

type CreateLocalJoinPermissionResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLocalJoinPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalJoinPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLocalJoinPermissionResponseBody) SetData(v string) *CreateLocalJoinPermissionResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLocalJoinPermissionResponseBody) SetRequestId(v string) *CreateLocalJoinPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLocalJoinPermissionResponseBody) SetSuccess(v bool) *CreateLocalJoinPermissionResponseBody {
	s.Success = &v
	return s
}

type CreateLocalJoinPermissionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLocalJoinPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLocalJoinPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLocalJoinPermissionResponse) GoString() string {
	return s.String()
}

func (s *CreateLocalJoinPermissionResponse) SetHeaders(v map[string]*string) *CreateLocalJoinPermissionResponse {
	s.Headers = v
	return s
}

func (s *CreateLocalJoinPermissionResponse) SetStatusCode(v int32) *CreateLocalJoinPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLocalJoinPermissionResponse) SetBody(v *CreateLocalJoinPermissionResponseBody) *CreateLocalJoinPermissionResponse {
	s.Body = v
	return s
}

type CreateNodeGroupRequest struct {
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	NodeGroupName    *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
}

func (s CreateNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequest) SetJoinPermissionId(v string) *CreateNodeGroupRequest {
	s.JoinPermissionId = &v
	return s
}

func (s *CreateNodeGroupRequest) SetNodeGroupName(v string) *CreateNodeGroupRequest {
	s.NodeGroupName = &v
	return s
}

type CreateNodeGroupResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponseBody) SetData(v string) *CreateNodeGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateNodeGroupResponseBody) SetRequestId(v string) *CreateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeGroupResponseBody) SetSuccess(v bool) *CreateNodeGroupResponseBody {
	s.Success = &v
	return s
}

type CreateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponse) SetHeaders(v map[string]*string) *CreateNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeGroupResponse) SetStatusCode(v int32) *CreateNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeGroupResponse) SetBody(v *CreateNodeGroupResponseBody) *CreateNodeGroupResponse {
	s.Body = v
	return s
}

type DeleteGatewayRequest struct {
	GwEui         *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s DeleteGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRequest) SetGwEui(v string) *DeleteGatewayRequest {
	s.GwEui = &v
	return s
}

func (s *DeleteGatewayRequest) SetIotInstanceId(v string) *DeleteGatewayRequest {
	s.IotInstanceId = &v
	return s
}

type DeleteGatewayResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetSuccess(v bool) *DeleteGatewayResponseBody {
	s.Success = &v
	return s
}

type DeleteGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponse) SetHeaders(v map[string]*string) *DeleteGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayResponse) SetStatusCode(v int32) *DeleteGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayResponse) SetBody(v *DeleteGatewayResponseBody) *DeleteGatewayResponse {
	s.Body = v
	return s
}

type DeleteLocalJoinPermissionRequest struct {
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
}

func (s DeleteLocalJoinPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocalJoinPermissionRequest) GoString() string {
	return s.String()
}

func (s *DeleteLocalJoinPermissionRequest) SetIotInstanceId(v string) *DeleteLocalJoinPermissionRequest {
	s.IotInstanceId = &v
	return s
}

func (s *DeleteLocalJoinPermissionRequest) SetJoinPermissionId(v string) *DeleteLocalJoinPermissionRequest {
	s.JoinPermissionId = &v
	return s
}

type DeleteLocalJoinPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLocalJoinPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocalJoinPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLocalJoinPermissionResponseBody) SetRequestId(v string) *DeleteLocalJoinPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLocalJoinPermissionResponseBody) SetSuccess(v bool) *DeleteLocalJoinPermissionResponseBody {
	s.Success = &v
	return s
}

type DeleteLocalJoinPermissionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLocalJoinPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLocalJoinPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLocalJoinPermissionResponse) GoString() string {
	return s.String()
}

func (s *DeleteLocalJoinPermissionResponse) SetHeaders(v map[string]*string) *DeleteLocalJoinPermissionResponse {
	s.Headers = v
	return s
}

func (s *DeleteLocalJoinPermissionResponse) SetStatusCode(v int32) *DeleteLocalJoinPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLocalJoinPermissionResponse) SetBody(v *DeleteLocalJoinPermissionResponseBody) *DeleteLocalJoinPermissionResponse {
	s.Body = v
	return s
}

type DeleteNodeGroupRequest struct {
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s DeleteNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupRequest) SetNodeGroupId(v string) *DeleteNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

type DeleteNodeGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponseBody) SetRequestId(v string) *DeleteNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeGroupResponseBody) SetSuccess(v bool) *DeleteNodeGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponse) SetHeaders(v map[string]*string) *DeleteNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeGroupResponse) SetStatusCode(v int32) *DeleteNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodeGroupResponse) SetBody(v *DeleteNodeGroupResponseBody) *DeleteNodeGroupResponse {
	s.Body = v
	return s
}

type GetFreqBandPlanGroupRequest struct {
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetFreqBandPlanGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFreqBandPlanGroupRequest) GoString() string {
	return s.String()
}

func (s *GetFreqBandPlanGroupRequest) SetGroupId(v int64) *GetFreqBandPlanGroupRequest {
	s.GroupId = &v
	return s
}

type GetFreqBandPlanGroupResponseBody struct {
	Data      *GetFreqBandPlanGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFreqBandPlanGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFreqBandPlanGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetFreqBandPlanGroupResponseBody) SetData(v *GetFreqBandPlanGroupResponseBodyData) *GetFreqBandPlanGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetFreqBandPlanGroupResponseBody) SetRequestId(v string) *GetFreqBandPlanGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFreqBandPlanGroupResponseBody) SetSuccess(v bool) *GetFreqBandPlanGroupResponseBody {
	s.Success = &v
	return s
}

type GetFreqBandPlanGroupResponseBodyData struct {
	BeginFrequency    *int64  `json:"BeginFrequency,omitempty" xml:"BeginFrequency,omitempty"`
	EndFrequency      *int64  `json:"EndFrequency,omitempty" xml:"EndFrequency,omitempty"`
	FrequencyRegionId *string `json:"FrequencyRegionId,omitempty" xml:"FrequencyRegionId,omitempty"`
	FrequencyType     *string `json:"FrequencyType,omitempty" xml:"FrequencyType,omitempty"`
	GroupId           *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetFreqBandPlanGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFreqBandPlanGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFreqBandPlanGroupResponseBodyData) SetBeginFrequency(v int64) *GetFreqBandPlanGroupResponseBodyData {
	s.BeginFrequency = &v
	return s
}

func (s *GetFreqBandPlanGroupResponseBodyData) SetEndFrequency(v int64) *GetFreqBandPlanGroupResponseBodyData {
	s.EndFrequency = &v
	return s
}

func (s *GetFreqBandPlanGroupResponseBodyData) SetFrequencyRegionId(v string) *GetFreqBandPlanGroupResponseBodyData {
	s.FrequencyRegionId = &v
	return s
}

func (s *GetFreqBandPlanGroupResponseBodyData) SetFrequencyType(v string) *GetFreqBandPlanGroupResponseBodyData {
	s.FrequencyType = &v
	return s
}

func (s *GetFreqBandPlanGroupResponseBodyData) SetGroupId(v int64) *GetFreqBandPlanGroupResponseBodyData {
	s.GroupId = &v
	return s
}

type GetFreqBandPlanGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFreqBandPlanGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFreqBandPlanGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFreqBandPlanGroupResponse) GoString() string {
	return s.String()
}

func (s *GetFreqBandPlanGroupResponse) SetHeaders(v map[string]*string) *GetFreqBandPlanGroupResponse {
	s.Headers = v
	return s
}

func (s *GetFreqBandPlanGroupResponse) SetStatusCode(v int32) *GetFreqBandPlanGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFreqBandPlanGroupResponse) SetBody(v *GetFreqBandPlanGroupResponseBody) *GetFreqBandPlanGroupResponse {
	s.Body = v
	return s
}

type GetGatewayRequest struct {
	GwEui         *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s GetGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayRequest) SetGwEui(v string) *GetGatewayRequest {
	s.GwEui = &v
	return s
}

func (s *GetGatewayRequest) SetIotInstanceId(v string) *GetGatewayRequest {
	s.IotInstanceId = &v
	return s
}

type GetGatewayResponseBody struct {
	Data      *GetGatewayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBody) SetData(v *GetGatewayResponseBodyData) *GetGatewayResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayResponseBody) SetRequestId(v string) *GetGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayResponseBody) SetSuccess(v bool) *GetGatewayResponseBody {
	s.Success = &v
	return s
}

type GetGatewayResponseBodyData struct {
	Address                  *string  `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressCode              *int64   `json:"AddressCode,omitempty" xml:"AddressCode,omitempty"`
	AuthTypes                *string  `json:"AuthTypes,omitempty" xml:"AuthTypes,omitempty"`
	ChargeStatus             *string  `json:"ChargeStatus,omitempty" xml:"ChargeStatus,omitempty"`
	City                     *string  `json:"City,omitempty" xml:"City,omitempty"`
	ClassBSupported          *bool    `json:"ClassBSupported,omitempty" xml:"ClassBSupported,omitempty"`
	ClassBWorking            *bool    `json:"ClassBWorking,omitempty" xml:"ClassBWorking,omitempty"`
	CommunicationMode        *string  `json:"CommunicationMode,omitempty" xml:"CommunicationMode,omitempty"`
	Description              *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	District                 *string  `json:"District,omitempty" xml:"District,omitempty"`
	EmbeddedNsId             *string  `json:"EmbeddedNsId,omitempty" xml:"EmbeddedNsId,omitempty"`
	Enabled                  *bool    `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FreqBandPlanGroupId      *int64   `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	GisCoordinateSystem      *string  `json:"GisCoordinateSystem,omitempty" xml:"GisCoordinateSystem,omitempty"`
	GwEui                    *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	Latitude                 *float32 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude                *float32 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name                     *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OnlineState              *string  `json:"OnlineState,omitempty" xml:"OnlineState,omitempty"`
	OnlineStateChangedMillis *int64   `json:"OnlineStateChangedMillis,omitempty" xml:"OnlineStateChangedMillis,omitempty"`
	TimeCorrectable          *bool    `json:"TimeCorrectable,omitempty" xml:"TimeCorrectable,omitempty"`
}

func (s GetGatewayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyData) SetAddress(v string) *GetGatewayResponseBodyData {
	s.Address = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetAddressCode(v int64) *GetGatewayResponseBodyData {
	s.AddressCode = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetAuthTypes(v string) *GetGatewayResponseBodyData {
	s.AuthTypes = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetChargeStatus(v string) *GetGatewayResponseBodyData {
	s.ChargeStatus = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetCity(v string) *GetGatewayResponseBodyData {
	s.City = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetClassBSupported(v bool) *GetGatewayResponseBodyData {
	s.ClassBSupported = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetClassBWorking(v bool) *GetGatewayResponseBodyData {
	s.ClassBWorking = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetCommunicationMode(v string) *GetGatewayResponseBodyData {
	s.CommunicationMode = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetDescription(v string) *GetGatewayResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetDistrict(v string) *GetGatewayResponseBodyData {
	s.District = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetEmbeddedNsId(v string) *GetGatewayResponseBodyData {
	s.EmbeddedNsId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetEnabled(v bool) *GetGatewayResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetFreqBandPlanGroupId(v int64) *GetGatewayResponseBodyData {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGisCoordinateSystem(v string) *GetGatewayResponseBodyData {
	s.GisCoordinateSystem = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGwEui(v string) *GetGatewayResponseBodyData {
	s.GwEui = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetLatitude(v float32) *GetGatewayResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetLongitude(v float32) *GetGatewayResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetName(v string) *GetGatewayResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetOnlineState(v string) *GetGatewayResponseBodyData {
	s.OnlineState = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetOnlineStateChangedMillis(v int64) *GetGatewayResponseBodyData {
	s.OnlineStateChangedMillis = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetTimeCorrectable(v bool) *GetGatewayResponseBodyData {
	s.TimeCorrectable = &v
	return s
}

type GetGatewayResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayResponse) SetHeaders(v map[string]*string) *GetGatewayResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayResponse) SetStatusCode(v int32) *GetGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayResponse) SetBody(v *GetGatewayResponseBody) *GetGatewayResponse {
	s.Body = v
	return s
}

type GetGatewayPacketStatRequest struct {
	BeginMillis   *int64  `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	EndMillis     *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	GwEui         *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s GetGatewayPacketStatRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayPacketStatRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayPacketStatRequest) SetBeginMillis(v int64) *GetGatewayPacketStatRequest {
	s.BeginMillis = &v
	return s
}

func (s *GetGatewayPacketStatRequest) SetEndMillis(v int64) *GetGatewayPacketStatRequest {
	s.EndMillis = &v
	return s
}

func (s *GetGatewayPacketStatRequest) SetGwEui(v string) *GetGatewayPacketStatRequest {
	s.GwEui = &v
	return s
}

func (s *GetGatewayPacketStatRequest) SetIotInstanceId(v string) *GetGatewayPacketStatRequest {
	s.IotInstanceId = &v
	return s
}

type GetGatewayPacketStatResponseBody struct {
	Data      *GetGatewayPacketStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayPacketStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayPacketStatResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayPacketStatResponseBody) SetData(v *GetGatewayPacketStatResponseBodyData) *GetGatewayPacketStatResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayPacketStatResponseBody) SetRequestId(v string) *GetGatewayPacketStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayPacketStatResponseBody) SetSuccess(v bool) *GetGatewayPacketStatResponseBody {
	s.Success = &v
	return s
}

type GetGatewayPacketStatResponseBodyData struct {
	DownlinkInvalid *int32 `json:"DownlinkInvalid,omitempty" xml:"DownlinkInvalid,omitempty"`
	DownlinkValid   *int32 `json:"DownlinkValid,omitempty" xml:"DownlinkValid,omitempty"`
	UplinkInvalid   *int32 `json:"UplinkInvalid,omitempty" xml:"UplinkInvalid,omitempty"`
	UplinkValid     *int32 `json:"UplinkValid,omitempty" xml:"UplinkValid,omitempty"`
}

func (s GetGatewayPacketStatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayPacketStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayPacketStatResponseBodyData) SetDownlinkInvalid(v int32) *GetGatewayPacketStatResponseBodyData {
	s.DownlinkInvalid = &v
	return s
}

func (s *GetGatewayPacketStatResponseBodyData) SetDownlinkValid(v int32) *GetGatewayPacketStatResponseBodyData {
	s.DownlinkValid = &v
	return s
}

func (s *GetGatewayPacketStatResponseBodyData) SetUplinkInvalid(v int32) *GetGatewayPacketStatResponseBodyData {
	s.UplinkInvalid = &v
	return s
}

func (s *GetGatewayPacketStatResponseBodyData) SetUplinkValid(v int32) *GetGatewayPacketStatResponseBodyData {
	s.UplinkValid = &v
	return s
}

type GetGatewayPacketStatResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGatewayPacketStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGatewayPacketStatResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayPacketStatResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayPacketStatResponse) SetHeaders(v map[string]*string) *GetGatewayPacketStatResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayPacketStatResponse) SetStatusCode(v int32) *GetGatewayPacketStatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayPacketStatResponse) SetBody(v *GetGatewayPacketStatResponseBody) *GetGatewayPacketStatResponse {
	s.Body = v
	return s
}

type GetGatewayStatusStatRequest struct {
	GwEui         *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s GetGatewayStatusStatRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayStatusStatRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayStatusStatRequest) SetGwEui(v string) *GetGatewayStatusStatRequest {
	s.GwEui = &v
	return s
}

func (s *GetGatewayStatusStatRequest) SetIotInstanceId(v string) *GetGatewayStatusStatRequest {
	s.IotInstanceId = &v
	return s
}

type GetGatewayStatusStatResponseBody struct {
	Data      *GetGatewayStatusStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayStatusStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayStatusStatResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayStatusStatResponseBody) SetData(v *GetGatewayStatusStatResponseBodyData) *GetGatewayStatusStatResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayStatusStatResponseBody) SetRequestId(v string) *GetGatewayStatusStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayStatusStatResponseBody) SetSuccess(v bool) *GetGatewayStatusStatResponseBody {
	s.Success = &v
	return s
}

type GetGatewayStatusStatResponseBodyData struct {
	CpuRadio    *float32 `json:"CpuRadio,omitempty" xml:"CpuRadio,omitempty"`
	Enabled     *bool    `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GwEui       *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	MemoryRadio *float32 `json:"MemoryRadio,omitempty" xml:"MemoryRadio,omitempty"`
	OnlineHour  *int64   `json:"OnlineHour,omitempty" xml:"OnlineHour,omitempty"`
	OnlineState *string  `json:"OnlineState,omitempty" xml:"OnlineState,omitempty"`
	RxCount     *int64   `json:"RxCount,omitempty" xml:"RxCount,omitempty"`
	TxCount     *int64   `json:"TxCount,omitempty" xml:"TxCount,omitempty"`
}

func (s GetGatewayStatusStatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayStatusStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayStatusStatResponseBodyData) SetCpuRadio(v float32) *GetGatewayStatusStatResponseBodyData {
	s.CpuRadio = &v
	return s
}

func (s *GetGatewayStatusStatResponseBodyData) SetEnabled(v bool) *GetGatewayStatusStatResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetGatewayStatusStatResponseBodyData) SetGwEui(v string) *GetGatewayStatusStatResponseBodyData {
	s.GwEui = &v
	return s
}

func (s *GetGatewayStatusStatResponseBodyData) SetMemoryRadio(v float32) *GetGatewayStatusStatResponseBodyData {
	s.MemoryRadio = &v
	return s
}

func (s *GetGatewayStatusStatResponseBodyData) SetOnlineHour(v int64) *GetGatewayStatusStatResponseBodyData {
	s.OnlineHour = &v
	return s
}

func (s *GetGatewayStatusStatResponseBodyData) SetOnlineState(v string) *GetGatewayStatusStatResponseBodyData {
	s.OnlineState = &v
	return s
}

func (s *GetGatewayStatusStatResponseBodyData) SetRxCount(v int64) *GetGatewayStatusStatResponseBodyData {
	s.RxCount = &v
	return s
}

func (s *GetGatewayStatusStatResponseBodyData) SetTxCount(v int64) *GetGatewayStatusStatResponseBodyData {
	s.TxCount = &v
	return s
}

type GetGatewayStatusStatResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGatewayStatusStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGatewayStatusStatResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayStatusStatResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayStatusStatResponse) SetHeaders(v map[string]*string) *GetGatewayStatusStatResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayStatusStatResponse) SetStatusCode(v int32) *GetGatewayStatusStatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayStatusStatResponse) SetBody(v *GetGatewayStatusStatResponseBody) *GetGatewayStatusStatResponse {
	s.Body = v
	return s
}

type GetGatewayTransferPacketsDownloadUrlRequest struct {
	Ascending     *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	BeginMillis   *int64  `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DevEui        *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	EndMillis     *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	GwEui         *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	SortingField  *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s GetGatewayTransferPacketsDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTransferPacketsDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayTransferPacketsDownloadUrlRequest) SetAscending(v bool) *GetGatewayTransferPacketsDownloadUrlRequest {
	s.Ascending = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlRequest) SetBeginMillis(v int64) *GetGatewayTransferPacketsDownloadUrlRequest {
	s.BeginMillis = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlRequest) SetCategory(v string) *GetGatewayTransferPacketsDownloadUrlRequest {
	s.Category = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlRequest) SetDevEui(v string) *GetGatewayTransferPacketsDownloadUrlRequest {
	s.DevEui = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlRequest) SetEndMillis(v int64) *GetGatewayTransferPacketsDownloadUrlRequest {
	s.EndMillis = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlRequest) SetGwEui(v string) *GetGatewayTransferPacketsDownloadUrlRequest {
	s.GwEui = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlRequest) SetIotInstanceId(v string) *GetGatewayTransferPacketsDownloadUrlRequest {
	s.IotInstanceId = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlRequest) SetSortingField(v string) *GetGatewayTransferPacketsDownloadUrlRequest {
	s.SortingField = &v
	return s
}

type GetGatewayTransferPacketsDownloadUrlResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayTransferPacketsDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTransferPacketsDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayTransferPacketsDownloadUrlResponseBody) SetData(v string) *GetGatewayTransferPacketsDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlResponseBody) SetRequestId(v string) *GetGatewayTransferPacketsDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlResponseBody) SetSuccess(v bool) *GetGatewayTransferPacketsDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type GetGatewayTransferPacketsDownloadUrlResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGatewayTransferPacketsDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGatewayTransferPacketsDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTransferPacketsDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayTransferPacketsDownloadUrlResponse) SetHeaders(v map[string]*string) *GetGatewayTransferPacketsDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlResponse) SetStatusCode(v int32) *GetGatewayTransferPacketsDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayTransferPacketsDownloadUrlResponse) SetBody(v *GetGatewayTransferPacketsDownloadUrlResponseBody) *GetGatewayTransferPacketsDownloadUrlResponse {
	s.Body = v
	return s
}

type GetGatewayTupleOrderRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetGatewayTupleOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTupleOrderRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayTupleOrderRequest) SetOrderId(v string) *GetGatewayTupleOrderRequest {
	s.OrderId = &v
	return s
}

type GetGatewayTupleOrderResponseBody struct {
	Data      *GetGatewayTupleOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayTupleOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTupleOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayTupleOrderResponseBody) SetData(v *GetGatewayTupleOrderResponseBodyData) *GetGatewayTupleOrderResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayTupleOrderResponseBody) SetRequestId(v string) *GetGatewayTupleOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayTupleOrderResponseBody) SetSuccess(v bool) *GetGatewayTupleOrderResponseBody {
	s.Success = &v
	return s
}

type GetGatewayTupleOrderResponseBodyData struct {
	AcceptedMillis *int64  `json:"AcceptedMillis,omitempty" xml:"AcceptedMillis,omitempty"`
	CreatedMillis  *int64  `json:"CreatedMillis,omitempty" xml:"CreatedMillis,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderState     *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	RequiredCount  *int64  `json:"RequiredCount,omitempty" xml:"RequiredCount,omitempty"`
}

func (s GetGatewayTupleOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTupleOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayTupleOrderResponseBodyData) SetAcceptedMillis(v int64) *GetGatewayTupleOrderResponseBodyData {
	s.AcceptedMillis = &v
	return s
}

func (s *GetGatewayTupleOrderResponseBodyData) SetCreatedMillis(v int64) *GetGatewayTupleOrderResponseBodyData {
	s.CreatedMillis = &v
	return s
}

func (s *GetGatewayTupleOrderResponseBodyData) SetOrderId(v string) *GetGatewayTupleOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetGatewayTupleOrderResponseBodyData) SetOrderState(v string) *GetGatewayTupleOrderResponseBodyData {
	s.OrderState = &v
	return s
}

func (s *GetGatewayTupleOrderResponseBodyData) SetRequiredCount(v int64) *GetGatewayTupleOrderResponseBodyData {
	s.RequiredCount = &v
	return s
}

type GetGatewayTupleOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGatewayTupleOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGatewayTupleOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTupleOrderResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayTupleOrderResponse) SetHeaders(v map[string]*string) *GetGatewayTupleOrderResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayTupleOrderResponse) SetStatusCode(v int32) *GetGatewayTupleOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayTupleOrderResponse) SetBody(v *GetGatewayTupleOrderResponseBody) *GetGatewayTupleOrderResponse {
	s.Body = v
	return s
}

type GetGatewayTuplesDownloadUrlRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetGatewayTuplesDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTuplesDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayTuplesDownloadUrlRequest) SetOrderId(v string) *GetGatewayTuplesDownloadUrlRequest {
	s.OrderId = &v
	return s
}

type GetGatewayTuplesDownloadUrlResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayTuplesDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTuplesDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayTuplesDownloadUrlResponseBody) SetData(v string) *GetGatewayTuplesDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetGatewayTuplesDownloadUrlResponseBody) SetRequestId(v string) *GetGatewayTuplesDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayTuplesDownloadUrlResponseBody) SetSuccess(v bool) *GetGatewayTuplesDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type GetGatewayTuplesDownloadUrlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGatewayTuplesDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGatewayTuplesDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayTuplesDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayTuplesDownloadUrlResponse) SetHeaders(v map[string]*string) *GetGatewayTuplesDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayTuplesDownloadUrlResponse) SetStatusCode(v int32) *GetGatewayTuplesDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayTuplesDownloadUrlResponse) SetBody(v *GetGatewayTuplesDownloadUrlResponseBody) *GetGatewayTuplesDownloadUrlResponse {
	s.Body = v
	return s
}

type GetJoinPermissionAuthOrderRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetJoinPermissionAuthOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJoinPermissionAuthOrderRequest) GoString() string {
	return s.String()
}

func (s *GetJoinPermissionAuthOrderRequest) SetOrderId(v string) *GetJoinPermissionAuthOrderRequest {
	s.OrderId = &v
	return s
}

type GetJoinPermissionAuthOrderResponseBody struct {
	Data      *GetJoinPermissionAuthOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJoinPermissionAuthOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJoinPermissionAuthOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetJoinPermissionAuthOrderResponseBody) SetData(v *GetJoinPermissionAuthOrderResponseBodyData) *GetJoinPermissionAuthOrderResponseBody {
	s.Data = v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBody) SetRequestId(v string) *GetJoinPermissionAuthOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBody) SetSuccess(v bool) *GetJoinPermissionAuthOrderResponseBody {
	s.Success = &v
	return s
}

type GetJoinPermissionAuthOrderResponseBodyData struct {
	AcceptedMillis   *int64  `json:"AcceptedMillis,omitempty" xml:"AcceptedMillis,omitempty"`
	ApplyingMillis   *int64  `json:"ApplyingMillis,omitempty" xml:"ApplyingMillis,omitempty"`
	CanceledMillis   *int64  `json:"CanceledMillis,omitempty" xml:"CanceledMillis,omitempty"`
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderState       *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	OwnerAliyunId    *string `json:"OwnerAliyunId,omitempty" xml:"OwnerAliyunId,omitempty"`
	RejectedMillis   *int64  `json:"RejectedMillis,omitempty" xml:"RejectedMillis,omitempty"`
	RenterAliyunId   *string `json:"RenterAliyunId,omitempty" xml:"RenterAliyunId,omitempty"`
}

func (s GetJoinPermissionAuthOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJoinPermissionAuthOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJoinPermissionAuthOrderResponseBodyData) SetAcceptedMillis(v int64) *GetJoinPermissionAuthOrderResponseBodyData {
	s.AcceptedMillis = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBodyData) SetApplyingMillis(v int64) *GetJoinPermissionAuthOrderResponseBodyData {
	s.ApplyingMillis = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBodyData) SetCanceledMillis(v int64) *GetJoinPermissionAuthOrderResponseBodyData {
	s.CanceledMillis = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBodyData) SetJoinPermissionId(v string) *GetJoinPermissionAuthOrderResponseBodyData {
	s.JoinPermissionId = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBodyData) SetOrderId(v string) *GetJoinPermissionAuthOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBodyData) SetOrderState(v string) *GetJoinPermissionAuthOrderResponseBodyData {
	s.OrderState = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBodyData) SetOwnerAliyunId(v string) *GetJoinPermissionAuthOrderResponseBodyData {
	s.OwnerAliyunId = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBodyData) SetRejectedMillis(v int64) *GetJoinPermissionAuthOrderResponseBodyData {
	s.RejectedMillis = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponseBodyData) SetRenterAliyunId(v string) *GetJoinPermissionAuthOrderResponseBodyData {
	s.RenterAliyunId = &v
	return s
}

type GetJoinPermissionAuthOrderResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJoinPermissionAuthOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJoinPermissionAuthOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJoinPermissionAuthOrderResponse) GoString() string {
	return s.String()
}

func (s *GetJoinPermissionAuthOrderResponse) SetHeaders(v map[string]*string) *GetJoinPermissionAuthOrderResponse {
	s.Headers = v
	return s
}

func (s *GetJoinPermissionAuthOrderResponse) SetStatusCode(v int32) *GetJoinPermissionAuthOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJoinPermissionAuthOrderResponse) SetBody(v *GetJoinPermissionAuthOrderResponseBody) *GetJoinPermissionAuthOrderResponse {
	s.Body = v
	return s
}

type GetNodeRequest struct {
	DevEui *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
}

func (s GetNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeRequest) SetDevEui(v string) *GetNodeRequest {
	s.DevEui = &v
	return s
}

type GetNodeResponseBody struct {
	Data      *GetNodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBody) SetData(v *GetNodeResponseBodyData) *GetNodeResponseBody {
	s.Data = v
	return s
}

func (s *GetNodeResponseBody) SetRequestId(v string) *GetNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeResponseBody) SetSuccess(v bool) *GetNodeResponseBody {
	s.Success = &v
	return s
}

type GetNodeResponseBodyData struct {
	AuthTypes      *string `json:"AuthTypes,omitempty" xml:"AuthTypes,omitempty"`
	BoundMillis    *int64  `json:"BoundMillis,omitempty" xml:"BoundMillis,omitempty"`
	ClassMode      *string `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	DevAddr        *string `json:"DevAddr,omitempty" xml:"DevAddr,omitempty"`
	DevEui         *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	LastJoinMillis *int64  `json:"LastJoinMillis,omitempty" xml:"LastJoinMillis,omitempty"`
}

func (s GetNodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyData) SetAuthTypes(v string) *GetNodeResponseBodyData {
	s.AuthTypes = &v
	return s
}

func (s *GetNodeResponseBodyData) SetBoundMillis(v int64) *GetNodeResponseBodyData {
	s.BoundMillis = &v
	return s
}

func (s *GetNodeResponseBodyData) SetClassMode(v string) *GetNodeResponseBodyData {
	s.ClassMode = &v
	return s
}

func (s *GetNodeResponseBodyData) SetDevAddr(v string) *GetNodeResponseBodyData {
	s.DevAddr = &v
	return s
}

func (s *GetNodeResponseBodyData) SetDevEui(v string) *GetNodeResponseBodyData {
	s.DevEui = &v
	return s
}

func (s *GetNodeResponseBodyData) SetLastJoinMillis(v int64) *GetNodeResponseBodyData {
	s.LastJoinMillis = &v
	return s
}

type GetNodeResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponse) GoString() string {
	return s.String()
}

func (s *GetNodeResponse) SetHeaders(v map[string]*string) *GetNodeResponse {
	s.Headers = v
	return s
}

func (s *GetNodeResponse) SetStatusCode(v int32) *GetNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeResponse) SetBody(v *GetNodeResponseBody) *GetNodeResponse {
	s.Body = v
	return s
}

type GetNodeGroupRequest struct {
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	NodeGroupId   *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s GetNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *GetNodeGroupRequest) SetIotInstanceId(v string) *GetNodeGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *GetNodeGroupRequest) SetNodeGroupId(v string) *GetNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

type GetNodeGroupResponseBody struct {
	Data      *GetNodeGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponseBody) SetData(v *GetNodeGroupResponseBodyData) *GetNodeGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetNodeGroupResponseBody) SetRequestId(v string) *GetNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeGroupResponseBody) SetSuccess(v bool) *GetNodeGroupResponseBody {
	s.Success = &v
	return s
}

type GetNodeGroupResponseBodyData struct {
	ClassMode                   *string                                         `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	CreateMillis                *int64                                          `json:"CreateMillis,omitempty" xml:"CreateMillis,omitempty"`
	DataDispatchConfig          *GetNodeGroupResponseBodyDataDataDispatchConfig `json:"DataDispatchConfig,omitempty" xml:"DataDispatchConfig,omitempty" type:"Struct"`
	DataDispatchEnabled         *bool                                           `json:"DataDispatchEnabled,omitempty" xml:"DataDispatchEnabled,omitempty"`
	FreqBandPlanGroupId         *int64                                          `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	JoinEui                     *string                                         `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionEnabled       *bool                                           `json:"JoinPermissionEnabled,omitempty" xml:"JoinPermissionEnabled,omitempty"`
	JoinPermissionId            *string                                         `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName          *string                                         `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	JoinPermissionOwnerAliyunId *string                                         `json:"JoinPermissionOwnerAliyunId,omitempty" xml:"JoinPermissionOwnerAliyunId,omitempty"`
	JoinPermissionType          *string                                         `json:"JoinPermissionType,omitempty" xml:"JoinPermissionType,omitempty"`
	Locks                       []*GetNodeGroupResponseBodyDataLocks            `json:"Locks,omitempty" xml:"Locks,omitempty" type:"Repeated"`
	MulticastEnabled            *bool                                           `json:"MulticastEnabled,omitempty" xml:"MulticastEnabled,omitempty"`
	MulticastGroupId            *string                                         `json:"MulticastGroupId,omitempty" xml:"MulticastGroupId,omitempty"`
	MulticastNodeCapacity       *int32                                          `json:"MulticastNodeCapacity,omitempty" xml:"MulticastNodeCapacity,omitempty"`
	MulticastNodeCount          *int32                                          `json:"MulticastNodeCount,omitempty" xml:"MulticastNodeCount,omitempty"`
	NodeGroupId                 *string                                         `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName               *string                                         `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodesCnt                    *int64                                          `json:"NodesCnt,omitempty" xml:"NodesCnt,omitempty"`
	RxDailySum                  *string                                         `json:"RxDailySum,omitempty" xml:"RxDailySum,omitempty"`
	RxMonthSum                  *int64                                          `json:"RxMonthSum,omitempty" xml:"RxMonthSum,omitempty"`
	TxDailySum                  *int64                                          `json:"TxDailySum,omitempty" xml:"TxDailySum,omitempty"`
	TxMonthSum                  *int64                                          `json:"TxMonthSum,omitempty" xml:"TxMonthSum,omitempty"`
}

func (s GetNodeGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponseBodyData) SetClassMode(v string) *GetNodeGroupResponseBodyData {
	s.ClassMode = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetCreateMillis(v int64) *GetNodeGroupResponseBodyData {
	s.CreateMillis = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetDataDispatchConfig(v *GetNodeGroupResponseBodyDataDataDispatchConfig) *GetNodeGroupResponseBodyData {
	s.DataDispatchConfig = v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetDataDispatchEnabled(v bool) *GetNodeGroupResponseBodyData {
	s.DataDispatchEnabled = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetFreqBandPlanGroupId(v int64) *GetNodeGroupResponseBodyData {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetJoinEui(v string) *GetNodeGroupResponseBodyData {
	s.JoinEui = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetJoinPermissionEnabled(v bool) *GetNodeGroupResponseBodyData {
	s.JoinPermissionEnabled = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetJoinPermissionId(v string) *GetNodeGroupResponseBodyData {
	s.JoinPermissionId = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetJoinPermissionName(v string) *GetNodeGroupResponseBodyData {
	s.JoinPermissionName = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetJoinPermissionOwnerAliyunId(v string) *GetNodeGroupResponseBodyData {
	s.JoinPermissionOwnerAliyunId = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetJoinPermissionType(v string) *GetNodeGroupResponseBodyData {
	s.JoinPermissionType = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetLocks(v []*GetNodeGroupResponseBodyDataLocks) *GetNodeGroupResponseBodyData {
	s.Locks = v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetMulticastEnabled(v bool) *GetNodeGroupResponseBodyData {
	s.MulticastEnabled = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetMulticastGroupId(v string) *GetNodeGroupResponseBodyData {
	s.MulticastGroupId = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetMulticastNodeCapacity(v int32) *GetNodeGroupResponseBodyData {
	s.MulticastNodeCapacity = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetMulticastNodeCount(v int32) *GetNodeGroupResponseBodyData {
	s.MulticastNodeCount = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetNodeGroupId(v string) *GetNodeGroupResponseBodyData {
	s.NodeGroupId = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetNodeGroupName(v string) *GetNodeGroupResponseBodyData {
	s.NodeGroupName = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetNodesCnt(v int64) *GetNodeGroupResponseBodyData {
	s.NodesCnt = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetRxDailySum(v string) *GetNodeGroupResponseBodyData {
	s.RxDailySum = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetRxMonthSum(v int64) *GetNodeGroupResponseBodyData {
	s.RxMonthSum = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetTxDailySum(v int64) *GetNodeGroupResponseBodyData {
	s.TxDailySum = &v
	return s
}

func (s *GetNodeGroupResponseBodyData) SetTxMonthSum(v int64) *GetNodeGroupResponseBodyData {
	s.TxMonthSum = &v
	return s
}

type GetNodeGroupResponseBodyDataDataDispatchConfig struct {
	Destination *string                                                   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	IotProduct  *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct `json:"IotProduct,omitempty" xml:"IotProduct,omitempty" type:"Struct"`
	OnsTopics   *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics  `json:"OnsTopics,omitempty" xml:"OnsTopics,omitempty" type:"Struct"`
}

func (s GetNodeGroupResponseBodyDataDataDispatchConfig) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupResponseBodyDataDataDispatchConfig) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfig) SetDestination(v string) *GetNodeGroupResponseBodyDataDataDispatchConfig {
	s.Destination = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfig) SetIotProduct(v *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct) *GetNodeGroupResponseBodyDataDataDispatchConfig {
	s.IotProduct = v
	return s
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfig) SetOnsTopics(v *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics) *GetNodeGroupResponseBodyDataDataDispatchConfig {
	s.OnsTopics = v
	return s
}

type GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct struct {
	DebugSwitch *bool   `json:"DebugSwitch,omitempty" xml:"DebugSwitch,omitempty"`
	ProductKey  *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct) SetDebugSwitch(v bool) *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct {
	s.DebugSwitch = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct) SetProductKey(v string) *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct {
	s.ProductKey = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct) SetProductName(v string) *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct {
	s.ProductName = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct) SetProductType(v string) *GetNodeGroupResponseBodyDataDataDispatchConfigIotProduct {
	s.ProductType = &v
	return s
}

type GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics struct {
	DownlinkRegionName *string `json:"DownlinkRegionName,omitempty" xml:"DownlinkRegionName,omitempty"`
	DownlinkTopic      *string `json:"DownlinkTopic,omitempty" xml:"DownlinkTopic,omitempty"`
	UplinkRegionName   *string `json:"UplinkRegionName,omitempty" xml:"UplinkRegionName,omitempty"`
	UplinkTopic        *string `json:"UplinkTopic,omitempty" xml:"UplinkTopic,omitempty"`
}

func (s GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics) SetDownlinkRegionName(v string) *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics {
	s.DownlinkRegionName = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics) SetDownlinkTopic(v string) *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics {
	s.DownlinkTopic = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics) SetUplinkRegionName(v string) *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics {
	s.UplinkRegionName = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics) SetUplinkTopic(v string) *GetNodeGroupResponseBodyDataDataDispatchConfigOnsTopics {
	s.UplinkTopic = &v
	return s
}

type GetNodeGroupResponseBodyDataLocks struct {
	CreateMillis *int64  `json:"CreateMillis,omitempty" xml:"CreateMillis,omitempty"`
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	LockId       *string `json:"LockId,omitempty" xml:"LockId,omitempty"`
	LockType     *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s GetNodeGroupResponseBodyDataLocks) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupResponseBodyDataLocks) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponseBodyDataLocks) SetCreateMillis(v int64) *GetNodeGroupResponseBodyDataLocks {
	s.CreateMillis = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataLocks) SetEnabled(v bool) *GetNodeGroupResponseBodyDataLocks {
	s.Enabled = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataLocks) SetLockId(v string) *GetNodeGroupResponseBodyDataLocks {
	s.LockId = &v
	return s
}

func (s *GetNodeGroupResponseBodyDataLocks) SetLockType(v string) *GetNodeGroupResponseBodyDataLocks {
	s.LockType = &v
	return s
}

type GetNodeGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponse) SetHeaders(v map[string]*string) *GetNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *GetNodeGroupResponse) SetStatusCode(v int32) *GetNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeGroupResponse) SetBody(v *GetNodeGroupResponseBody) *GetNodeGroupResponse {
	s.Body = v
	return s
}

type GetNodeGroupTransferPacketsDownloadUrlRequest struct {
	Ascending     *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	BeginMillis   *int64  `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DevEui        *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	EndMillis     *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	NodeGroupId   *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	SortingField  *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s GetNodeGroupTransferPacketsDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupTransferPacketsDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetNodeGroupTransferPacketsDownloadUrlRequest) SetAscending(v bool) *GetNodeGroupTransferPacketsDownloadUrlRequest {
	s.Ascending = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlRequest) SetBeginMillis(v int64) *GetNodeGroupTransferPacketsDownloadUrlRequest {
	s.BeginMillis = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlRequest) SetCategory(v string) *GetNodeGroupTransferPacketsDownloadUrlRequest {
	s.Category = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlRequest) SetDevEui(v string) *GetNodeGroupTransferPacketsDownloadUrlRequest {
	s.DevEui = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlRequest) SetEndMillis(v int64) *GetNodeGroupTransferPacketsDownloadUrlRequest {
	s.EndMillis = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlRequest) SetIotInstanceId(v string) *GetNodeGroupTransferPacketsDownloadUrlRequest {
	s.IotInstanceId = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlRequest) SetNodeGroupId(v string) *GetNodeGroupTransferPacketsDownloadUrlRequest {
	s.NodeGroupId = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlRequest) SetSortingField(v string) *GetNodeGroupTransferPacketsDownloadUrlRequest {
	s.SortingField = &v
	return s
}

type GetNodeGroupTransferPacketsDownloadUrlResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeGroupTransferPacketsDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupTransferPacketsDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeGroupTransferPacketsDownloadUrlResponseBody) SetData(v string) *GetNodeGroupTransferPacketsDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlResponseBody) SetRequestId(v string) *GetNodeGroupTransferPacketsDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlResponseBody) SetSuccess(v bool) *GetNodeGroupTransferPacketsDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type GetNodeGroupTransferPacketsDownloadUrlResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeGroupTransferPacketsDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeGroupTransferPacketsDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeGroupTransferPacketsDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetNodeGroupTransferPacketsDownloadUrlResponse) SetHeaders(v map[string]*string) *GetNodeGroupTransferPacketsDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlResponse) SetStatusCode(v int32) *GetNodeGroupTransferPacketsDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeGroupTransferPacketsDownloadUrlResponse) SetBody(v *GetNodeGroupTransferPacketsDownloadUrlResponseBody) *GetNodeGroupTransferPacketsDownloadUrlResponse {
	s.Body = v
	return s
}

type GetNodeTransferPacketRequest struct {
	Base64EncodedMacPayload *string `json:"Base64EncodedMacPayload,omitempty" xml:"Base64EncodedMacPayload,omitempty"`
	DevEui                  *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	IotInstanceId           *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	LogMillis               *int64  `json:"LogMillis,omitempty" xml:"LogMillis,omitempty"`
}

func (s GetNodeTransferPacketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTransferPacketRequest) GoString() string {
	return s.String()
}

func (s *GetNodeTransferPacketRequest) SetBase64EncodedMacPayload(v string) *GetNodeTransferPacketRequest {
	s.Base64EncodedMacPayload = &v
	return s
}

func (s *GetNodeTransferPacketRequest) SetDevEui(v string) *GetNodeTransferPacketRequest {
	s.DevEui = &v
	return s
}

func (s *GetNodeTransferPacketRequest) SetIotInstanceId(v string) *GetNodeTransferPacketRequest {
	s.IotInstanceId = &v
	return s
}

func (s *GetNodeTransferPacketRequest) SetLogMillis(v int64) *GetNodeTransferPacketRequest {
	s.LogMillis = &v
	return s
}

type GetNodeTransferPacketResponseBody struct {
	Data      *GetNodeTransferPacketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeTransferPacketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTransferPacketResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeTransferPacketResponseBody) SetData(v *GetNodeTransferPacketResponseBodyData) *GetNodeTransferPacketResponseBody {
	s.Data = v
	return s
}

func (s *GetNodeTransferPacketResponseBody) SetRequestId(v string) *GetNodeTransferPacketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeTransferPacketResponseBody) SetSuccess(v bool) *GetNodeTransferPacketResponseBody {
	s.Success = &v
	return s
}

type GetNodeTransferPacketResponseBodyData struct {
	Base64EncodedMacPayload *string                  `json:"Base64EncodedMacPayload,omitempty" xml:"Base64EncodedMacPayload,omitempty"`
	ClassMode               *string                  `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	Datr                    *string                  `json:"Datr,omitempty" xml:"Datr,omitempty"`
	DevAddr                 *string                  `json:"DevAddr,omitempty" xml:"DevAddr,omitempty"`
	DevEui                  *string                  `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	FPort                   *int32                   `json:"FPort,omitempty" xml:"FPort,omitempty"`
	Freq                    *float32                 `json:"Freq,omitempty" xml:"Freq,omitempty"`
	FreqBandPlanGroupId     *int64                   `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	GwEui                   *string                  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	GwOwnerAliyunId         *string                  `json:"GwOwnerAliyunId,omitempty" xml:"GwOwnerAliyunId,omitempty"`
	HasData                 *bool                    `json:"HasData,omitempty" xml:"HasData,omitempty"`
	HasMacCommand           *bool                    `json:"HasMacCommand,omitempty" xml:"HasMacCommand,omitempty"`
	LogMillis               *int64                   `json:"LogMillis,omitempty" xml:"LogMillis,omitempty"`
	Lsnr                    *float32                 `json:"Lsnr,omitempty" xml:"Lsnr,omitempty"`
	MacCommandCIDs          []map[string]interface{} `json:"MacCommandCIDs,omitempty" xml:"MacCommandCIDs,omitempty" type:"Repeated"`
	MacPayloadSize          *int32                   `json:"MacPayloadSize,omitempty" xml:"MacPayloadSize,omitempty"`
	MessageType             *string                  `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	ProcessEvent            *string                  `json:"ProcessEvent,omitempty" xml:"ProcessEvent,omitempty"`
	Rssi                    *int32                   `json:"Rssi,omitempty" xml:"Rssi,omitempty"`
}

func (s GetNodeTransferPacketResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTransferPacketResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNodeTransferPacketResponseBodyData) SetBase64EncodedMacPayload(v string) *GetNodeTransferPacketResponseBodyData {
	s.Base64EncodedMacPayload = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetClassMode(v string) *GetNodeTransferPacketResponseBodyData {
	s.ClassMode = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetDatr(v string) *GetNodeTransferPacketResponseBodyData {
	s.Datr = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetDevAddr(v string) *GetNodeTransferPacketResponseBodyData {
	s.DevAddr = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetDevEui(v string) *GetNodeTransferPacketResponseBodyData {
	s.DevEui = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetFPort(v int32) *GetNodeTransferPacketResponseBodyData {
	s.FPort = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetFreq(v float32) *GetNodeTransferPacketResponseBodyData {
	s.Freq = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetFreqBandPlanGroupId(v int64) *GetNodeTransferPacketResponseBodyData {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetGwEui(v string) *GetNodeTransferPacketResponseBodyData {
	s.GwEui = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetGwOwnerAliyunId(v string) *GetNodeTransferPacketResponseBodyData {
	s.GwOwnerAliyunId = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetHasData(v bool) *GetNodeTransferPacketResponseBodyData {
	s.HasData = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetHasMacCommand(v bool) *GetNodeTransferPacketResponseBodyData {
	s.HasMacCommand = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetLogMillis(v int64) *GetNodeTransferPacketResponseBodyData {
	s.LogMillis = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetLsnr(v float32) *GetNodeTransferPacketResponseBodyData {
	s.Lsnr = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetMacCommandCIDs(v []map[string]interface{}) *GetNodeTransferPacketResponseBodyData {
	s.MacCommandCIDs = v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetMacPayloadSize(v int32) *GetNodeTransferPacketResponseBodyData {
	s.MacPayloadSize = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetMessageType(v string) *GetNodeTransferPacketResponseBodyData {
	s.MessageType = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetProcessEvent(v string) *GetNodeTransferPacketResponseBodyData {
	s.ProcessEvent = &v
	return s
}

func (s *GetNodeTransferPacketResponseBodyData) SetRssi(v int32) *GetNodeTransferPacketResponseBodyData {
	s.Rssi = &v
	return s
}

type GetNodeTransferPacketResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeTransferPacketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeTransferPacketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTransferPacketResponse) GoString() string {
	return s.String()
}

func (s *GetNodeTransferPacketResponse) SetHeaders(v map[string]*string) *GetNodeTransferPacketResponse {
	s.Headers = v
	return s
}

func (s *GetNodeTransferPacketResponse) SetStatusCode(v int32) *GetNodeTransferPacketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeTransferPacketResponse) SetBody(v *GetNodeTransferPacketResponseBody) *GetNodeTransferPacketResponse {
	s.Body = v
	return s
}

type GetNodeTransferPacketsDownloadUrlRequest struct {
	Ascending    *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	BeginMillis  *int64  `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DevEui       *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	EndMillis    *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	GwEui        *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	SortingField *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s GetNodeTransferPacketsDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTransferPacketsDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetNodeTransferPacketsDownloadUrlRequest) SetAscending(v bool) *GetNodeTransferPacketsDownloadUrlRequest {
	s.Ascending = &v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlRequest) SetBeginMillis(v int64) *GetNodeTransferPacketsDownloadUrlRequest {
	s.BeginMillis = &v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlRequest) SetCategory(v string) *GetNodeTransferPacketsDownloadUrlRequest {
	s.Category = &v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlRequest) SetDevEui(v string) *GetNodeTransferPacketsDownloadUrlRequest {
	s.DevEui = &v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlRequest) SetEndMillis(v int64) *GetNodeTransferPacketsDownloadUrlRequest {
	s.EndMillis = &v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlRequest) SetGwEui(v string) *GetNodeTransferPacketsDownloadUrlRequest {
	s.GwEui = &v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlRequest) SetSortingField(v string) *GetNodeTransferPacketsDownloadUrlRequest {
	s.SortingField = &v
	return s
}

type GetNodeTransferPacketsDownloadUrlResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeTransferPacketsDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTransferPacketsDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeTransferPacketsDownloadUrlResponseBody) SetData(v string) *GetNodeTransferPacketsDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlResponseBody) SetRequestId(v string) *GetNodeTransferPacketsDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlResponseBody) SetSuccess(v bool) *GetNodeTransferPacketsDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type GetNodeTransferPacketsDownloadUrlResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeTransferPacketsDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeTransferPacketsDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTransferPacketsDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetNodeTransferPacketsDownloadUrlResponse) SetHeaders(v map[string]*string) *GetNodeTransferPacketsDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlResponse) SetStatusCode(v int32) *GetNodeTransferPacketsDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeTransferPacketsDownloadUrlResponse) SetBody(v *GetNodeTransferPacketsDownloadUrlResponseBody) *GetNodeTransferPacketsDownloadUrlResponse {
	s.Body = v
	return s
}

type GetNodeTupleOrderRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetNodeTupleOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTupleOrderRequest) GoString() string {
	return s.String()
}

func (s *GetNodeTupleOrderRequest) SetOrderId(v string) *GetNodeTupleOrderRequest {
	s.OrderId = &v
	return s
}

type GetNodeTupleOrderResponseBody struct {
	Data      *GetNodeTupleOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeTupleOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTupleOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeTupleOrderResponseBody) SetData(v *GetNodeTupleOrderResponseBodyData) *GetNodeTupleOrderResponseBody {
	s.Data = v
	return s
}

func (s *GetNodeTupleOrderResponseBody) SetRequestId(v string) *GetNodeTupleOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeTupleOrderResponseBody) SetSuccess(v bool) *GetNodeTupleOrderResponseBody {
	s.Success = &v
	return s
}

type GetNodeTupleOrderResponseBodyData struct {
	AcceptedMillis *int64  `json:"AcceptedMillis,omitempty" xml:"AcceptedMillis,omitempty"`
	CreatedMillis  *int64  `json:"CreatedMillis,omitempty" xml:"CreatedMillis,omitempty"`
	IsKpm          *bool   `json:"IsKpm,omitempty" xml:"IsKpm,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderState     *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	RequiredCount  *int64  `json:"RequiredCount,omitempty" xml:"RequiredCount,omitempty"`
}

func (s GetNodeTupleOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTupleOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNodeTupleOrderResponseBodyData) SetAcceptedMillis(v int64) *GetNodeTupleOrderResponseBodyData {
	s.AcceptedMillis = &v
	return s
}

func (s *GetNodeTupleOrderResponseBodyData) SetCreatedMillis(v int64) *GetNodeTupleOrderResponseBodyData {
	s.CreatedMillis = &v
	return s
}

func (s *GetNodeTupleOrderResponseBodyData) SetIsKpm(v bool) *GetNodeTupleOrderResponseBodyData {
	s.IsKpm = &v
	return s
}

func (s *GetNodeTupleOrderResponseBodyData) SetOrderId(v string) *GetNodeTupleOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetNodeTupleOrderResponseBodyData) SetOrderState(v string) *GetNodeTupleOrderResponseBodyData {
	s.OrderState = &v
	return s
}

func (s *GetNodeTupleOrderResponseBodyData) SetRequiredCount(v int64) *GetNodeTupleOrderResponseBodyData {
	s.RequiredCount = &v
	return s
}

type GetNodeTupleOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeTupleOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeTupleOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTupleOrderResponse) GoString() string {
	return s.String()
}

func (s *GetNodeTupleOrderResponse) SetHeaders(v map[string]*string) *GetNodeTupleOrderResponse {
	s.Headers = v
	return s
}

func (s *GetNodeTupleOrderResponse) SetStatusCode(v int32) *GetNodeTupleOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeTupleOrderResponse) SetBody(v *GetNodeTupleOrderResponseBody) *GetNodeTupleOrderResponse {
	s.Body = v
	return s
}

type GetNodeTuplesDownloadUrlRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetNodeTuplesDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTuplesDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetNodeTuplesDownloadUrlRequest) SetOrderId(v string) *GetNodeTuplesDownloadUrlRequest {
	s.OrderId = &v
	return s
}

type GetNodeTuplesDownloadUrlResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeTuplesDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTuplesDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeTuplesDownloadUrlResponseBody) SetData(v string) *GetNodeTuplesDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetNodeTuplesDownloadUrlResponseBody) SetRequestId(v string) *GetNodeTuplesDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeTuplesDownloadUrlResponseBody) SetSuccess(v bool) *GetNodeTuplesDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type GetNodeTuplesDownloadUrlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeTuplesDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeTuplesDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeTuplesDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetNodeTuplesDownloadUrlResponse) SetHeaders(v map[string]*string) *GetNodeTuplesDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetNodeTuplesDownloadUrlResponse) SetStatusCode(v int32) *GetNodeTuplesDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeTuplesDownloadUrlResponse) SetBody(v *GetNodeTuplesDownloadUrlResponseBody) *GetNodeTuplesDownloadUrlResponse {
	s.Body = v
	return s
}

type GetNotificationRequest struct {
	NotificationId *string `json:"NotificationId,omitempty" xml:"NotificationId,omitempty"`
}

func (s GetNotificationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNotificationRequest) GoString() string {
	return s.String()
}

func (s *GetNotificationRequest) SetNotificationId(v string) *GetNotificationRequest {
	s.NotificationId = &v
	return s
}

type GetNotificationResponseBody struct {
	Data      *GetNotificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNotificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotificationResponseBody) SetData(v *GetNotificationResponseBodyData) *GetNotificationResponseBody {
	s.Data = v
	return s
}

func (s *GetNotificationResponseBody) SetRequestId(v string) *GetNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNotificationResponseBody) SetSuccess(v bool) *GetNotificationResponseBody {
	s.Success = &v
	return s
}

type GetNotificationResponseBodyData struct {
	Category               *string                                                `json:"Category,omitempty" xml:"Category,omitempty"`
	GatewayOfflineInfo     *GetNotificationResponseBodyDataGatewayOfflineInfo     `json:"GatewayOfflineInfo,omitempty" xml:"GatewayOfflineInfo,omitempty" type:"Struct"`
	HandleState            *string                                                `json:"HandleState,omitempty" xml:"HandleState,omitempty"`
	HandledMillis          *int64                                                 `json:"HandledMillis,omitempty" xml:"HandledMillis,omitempty"`
	JoinPermissionAuthInfo *GetNotificationResponseBodyDataJoinPermissionAuthInfo `json:"JoinPermissionAuthInfo,omitempty" xml:"JoinPermissionAuthInfo,omitempty" type:"Struct"`
	NoticeMillis           *int64                                                 `json:"NoticeMillis,omitempty" xml:"NoticeMillis,omitempty"`
	NotificationId         *string                                                `json:"NotificationId,omitempty" xml:"NotificationId,omitempty"`
}

func (s GetNotificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNotificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNotificationResponseBodyData) SetCategory(v string) *GetNotificationResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetNotificationResponseBodyData) SetGatewayOfflineInfo(v *GetNotificationResponseBodyDataGatewayOfflineInfo) *GetNotificationResponseBodyData {
	s.GatewayOfflineInfo = v
	return s
}

func (s *GetNotificationResponseBodyData) SetHandleState(v string) *GetNotificationResponseBodyData {
	s.HandleState = &v
	return s
}

func (s *GetNotificationResponseBodyData) SetHandledMillis(v int64) *GetNotificationResponseBodyData {
	s.HandledMillis = &v
	return s
}

func (s *GetNotificationResponseBodyData) SetJoinPermissionAuthInfo(v *GetNotificationResponseBodyDataJoinPermissionAuthInfo) *GetNotificationResponseBodyData {
	s.JoinPermissionAuthInfo = v
	return s
}

func (s *GetNotificationResponseBodyData) SetNoticeMillis(v int64) *GetNotificationResponseBodyData {
	s.NoticeMillis = &v
	return s
}

func (s *GetNotificationResponseBodyData) SetNotificationId(v string) *GetNotificationResponseBodyData {
	s.NotificationId = &v
	return s
}

type GetNotificationResponseBodyDataGatewayOfflineInfo struct {
	GwEui         *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	OfflineMillis *int64  `json:"OfflineMillis,omitempty" xml:"OfflineMillis,omitempty"`
}

func (s GetNotificationResponseBodyDataGatewayOfflineInfo) String() string {
	return tea.Prettify(s)
}

func (s GetNotificationResponseBodyDataGatewayOfflineInfo) GoString() string {
	return s.String()
}

func (s *GetNotificationResponseBodyDataGatewayOfflineInfo) SetGwEui(v string) *GetNotificationResponseBodyDataGatewayOfflineInfo {
	s.GwEui = &v
	return s
}

func (s *GetNotificationResponseBodyDataGatewayOfflineInfo) SetOfflineMillis(v int64) *GetNotificationResponseBodyDataGatewayOfflineInfo {
	s.OfflineMillis = &v
	return s
}

type GetNotificationResponseBodyDataJoinPermissionAuthInfo struct {
	AcceptedMillis     *int64  `json:"AcceptedMillis,omitempty" xml:"AcceptedMillis,omitempty"`
	ApplyingMillis     *int64  `json:"ApplyingMillis,omitempty" xml:"ApplyingMillis,omitempty"`
	CanceledMillis     *int64  `json:"CanceledMillis,omitempty" xml:"CanceledMillis,omitempty"`
	JoinEui            *string `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionId   *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName *string `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderState         *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	OwnerAliyunId      *string `json:"OwnerAliyunId,omitempty" xml:"OwnerAliyunId,omitempty"`
	RejectedMillis     *int64  `json:"RejectedMillis,omitempty" xml:"RejectedMillis,omitempty"`
	RenterAliyunId     *string `json:"RenterAliyunId,omitempty" xml:"RenterAliyunId,omitempty"`
}

func (s GetNotificationResponseBodyDataJoinPermissionAuthInfo) String() string {
	return tea.Prettify(s)
}

func (s GetNotificationResponseBodyDataJoinPermissionAuthInfo) GoString() string {
	return s.String()
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetAcceptedMillis(v int64) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.AcceptedMillis = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetApplyingMillis(v int64) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.ApplyingMillis = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetCanceledMillis(v int64) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.CanceledMillis = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetJoinEui(v string) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.JoinEui = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetJoinPermissionId(v string) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.JoinPermissionId = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetJoinPermissionName(v string) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.JoinPermissionName = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetOrderId(v string) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.OrderId = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetOrderState(v string) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.OrderState = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetOwnerAliyunId(v string) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.OwnerAliyunId = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetRejectedMillis(v int64) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.RejectedMillis = &v
	return s
}

func (s *GetNotificationResponseBodyDataJoinPermissionAuthInfo) SetRenterAliyunId(v string) *GetNotificationResponseBodyDataJoinPermissionAuthInfo {
	s.RenterAliyunId = &v
	return s
}

type GetNotificationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNotificationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNotificationResponse) GoString() string {
	return s.String()
}

func (s *GetNotificationResponse) SetHeaders(v map[string]*string) *GetNotificationResponse {
	s.Headers = v
	return s
}

func (s *GetNotificationResponse) SetStatusCode(v int32) *GetNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNotificationResponse) SetBody(v *GetNotificationResponseBody) *GetNotificationResponse {
	s.Body = v
	return s
}

type GetOwnedJoinPermissionRequest struct {
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
}

func (s GetOwnedJoinPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOwnedJoinPermissionRequest) GoString() string {
	return s.String()
}

func (s *GetOwnedJoinPermissionRequest) SetIotInstanceId(v string) *GetOwnedJoinPermissionRequest {
	s.IotInstanceId = &v
	return s
}

func (s *GetOwnedJoinPermissionRequest) SetJoinPermissionId(v string) *GetOwnedJoinPermissionRequest {
	s.JoinPermissionId = &v
	return s
}

type GetOwnedJoinPermissionResponseBody struct {
	Data      *GetOwnedJoinPermissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOwnedJoinPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOwnedJoinPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetOwnedJoinPermissionResponseBody) SetData(v *GetOwnedJoinPermissionResponseBodyData) *GetOwnedJoinPermissionResponseBody {
	s.Data = v
	return s
}

func (s *GetOwnedJoinPermissionResponseBody) SetRequestId(v string) *GetOwnedJoinPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBody) SetSuccess(v bool) *GetOwnedJoinPermissionResponseBody {
	s.Success = &v
	return s
}

type GetOwnedJoinPermissionResponseBodyData struct {
	AuthState               *string `json:"AuthState,omitempty" xml:"AuthState,omitempty"`
	BoundProductName        *string `json:"BoundProductName,omitempty" xml:"BoundProductName,omitempty"`
	ClassMode               *string `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	CreateMillis            *int64  `json:"CreateMillis,omitempty" xml:"CreateMillis,omitempty"`
	DataDispatchDestination *string `json:"DataDispatchDestination,omitempty" xml:"DataDispatchDestination,omitempty"`
	DataRate                *int64  `json:"DataRate,omitempty" xml:"DataRate,omitempty"`
	Enabled                 *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FreqBandPlanGroupId     *int64  `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	JoinEui                 *string `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionId        *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName      *string `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	MulticastEnabled        *bool   `json:"MulticastEnabled,omitempty" xml:"MulticastEnabled,omitempty"`
	MulticastNodeCapacity   *int32  `json:"MulticastNodeCapacity,omitempty" xml:"MulticastNodeCapacity,omitempty"`
	MulticastNodeCount      *int32  `json:"MulticastNodeCount,omitempty" xml:"MulticastNodeCount,omitempty"`
	NodesCnt                *int64  `json:"NodesCnt,omitempty" xml:"NodesCnt,omitempty"`
	RenterAliyunId          *string `json:"RenterAliyunId,omitempty" xml:"RenterAliyunId,omitempty"`
	RxDailySum              *int64  `json:"RxDailySum,omitempty" xml:"RxDailySum,omitempty"`
	RxDelay                 *int64  `json:"RxDelay,omitempty" xml:"RxDelay,omitempty"`
	RxMonthSum              *int64  `json:"RxMonthSum,omitempty" xml:"RxMonthSum,omitempty"`
	TxDailySum              *int64  `json:"TxDailySum,omitempty" xml:"TxDailySum,omitempty"`
	TxMonthSum              *int64  `json:"TxMonthSum,omitempty" xml:"TxMonthSum,omitempty"`
}

func (s GetOwnedJoinPermissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOwnedJoinPermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetAuthState(v string) *GetOwnedJoinPermissionResponseBodyData {
	s.AuthState = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetBoundProductName(v string) *GetOwnedJoinPermissionResponseBodyData {
	s.BoundProductName = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetClassMode(v string) *GetOwnedJoinPermissionResponseBodyData {
	s.ClassMode = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetCreateMillis(v int64) *GetOwnedJoinPermissionResponseBodyData {
	s.CreateMillis = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetDataDispatchDestination(v string) *GetOwnedJoinPermissionResponseBodyData {
	s.DataDispatchDestination = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetDataRate(v int64) *GetOwnedJoinPermissionResponseBodyData {
	s.DataRate = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetEnabled(v bool) *GetOwnedJoinPermissionResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetFreqBandPlanGroupId(v int64) *GetOwnedJoinPermissionResponseBodyData {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetJoinEui(v string) *GetOwnedJoinPermissionResponseBodyData {
	s.JoinEui = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetJoinPermissionId(v string) *GetOwnedJoinPermissionResponseBodyData {
	s.JoinPermissionId = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetJoinPermissionName(v string) *GetOwnedJoinPermissionResponseBodyData {
	s.JoinPermissionName = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetMulticastEnabled(v bool) *GetOwnedJoinPermissionResponseBodyData {
	s.MulticastEnabled = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetMulticastNodeCapacity(v int32) *GetOwnedJoinPermissionResponseBodyData {
	s.MulticastNodeCapacity = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetMulticastNodeCount(v int32) *GetOwnedJoinPermissionResponseBodyData {
	s.MulticastNodeCount = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetNodesCnt(v int64) *GetOwnedJoinPermissionResponseBodyData {
	s.NodesCnt = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetRenterAliyunId(v string) *GetOwnedJoinPermissionResponseBodyData {
	s.RenterAliyunId = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetRxDailySum(v int64) *GetOwnedJoinPermissionResponseBodyData {
	s.RxDailySum = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetRxDelay(v int64) *GetOwnedJoinPermissionResponseBodyData {
	s.RxDelay = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetRxMonthSum(v int64) *GetOwnedJoinPermissionResponseBodyData {
	s.RxMonthSum = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetTxDailySum(v int64) *GetOwnedJoinPermissionResponseBodyData {
	s.TxDailySum = &v
	return s
}

func (s *GetOwnedJoinPermissionResponseBodyData) SetTxMonthSum(v int64) *GetOwnedJoinPermissionResponseBodyData {
	s.TxMonthSum = &v
	return s
}

type GetOwnedJoinPermissionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOwnedJoinPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOwnedJoinPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOwnedJoinPermissionResponse) GoString() string {
	return s.String()
}

func (s *GetOwnedJoinPermissionResponse) SetHeaders(v map[string]*string) *GetOwnedJoinPermissionResponse {
	s.Headers = v
	return s
}

func (s *GetOwnedJoinPermissionResponse) SetStatusCode(v int32) *GetOwnedJoinPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOwnedJoinPermissionResponse) SetBody(v *GetOwnedJoinPermissionResponseBody) *GetOwnedJoinPermissionResponse {
	s.Body = v
	return s
}

type GetRentedJoinPermissionRequest struct {
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
}

func (s GetRentedJoinPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRentedJoinPermissionRequest) GoString() string {
	return s.String()
}

func (s *GetRentedJoinPermissionRequest) SetJoinPermissionId(v string) *GetRentedJoinPermissionRequest {
	s.JoinPermissionId = &v
	return s
}

type GetRentedJoinPermissionResponseBody struct {
	Data      *GetRentedJoinPermissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRentedJoinPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRentedJoinPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRentedJoinPermissionResponseBody) SetData(v *GetRentedJoinPermissionResponseBodyData) *GetRentedJoinPermissionResponseBody {
	s.Data = v
	return s
}

func (s *GetRentedJoinPermissionResponseBody) SetRequestId(v string) *GetRentedJoinPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBody) SetSuccess(v bool) *GetRentedJoinPermissionResponseBody {
	s.Success = &v
	return s
}

type GetRentedJoinPermissionResponseBodyData struct {
	BoundNodeGroupId    *string `json:"BoundNodeGroupId,omitempty" xml:"BoundNodeGroupId,omitempty"`
	BoundNodeGroupName  *string `json:"BoundNodeGroupName,omitempty" xml:"BoundNodeGroupName,omitempty"`
	ClassMode           *string `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	CreateMillis        *int64  `json:"CreateMillis,omitempty" xml:"CreateMillis,omitempty"`
	DataRate            *int64  `json:"DataRate,omitempty" xml:"DataRate,omitempty"`
	Enabled             *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FreqBandPlanGroupId *int64  `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	JoinEui             *string `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionId    *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName  *string `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	NodesCnt            *int64  `json:"NodesCnt,omitempty" xml:"NodesCnt,omitempty"`
	RxDailySum          *int64  `json:"RxDailySum,omitempty" xml:"RxDailySum,omitempty"`
	RxDelay             *int64  `json:"RxDelay,omitempty" xml:"RxDelay,omitempty"`
	RxMonthSum          *int64  `json:"RxMonthSum,omitempty" xml:"RxMonthSum,omitempty"`
	TxDailySum          *int64  `json:"TxDailySum,omitempty" xml:"TxDailySum,omitempty"`
	TxMonthSum          *int64  `json:"TxMonthSum,omitempty" xml:"TxMonthSum,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRentedJoinPermissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRentedJoinPermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRentedJoinPermissionResponseBodyData) SetBoundNodeGroupId(v string) *GetRentedJoinPermissionResponseBodyData {
	s.BoundNodeGroupId = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetBoundNodeGroupName(v string) *GetRentedJoinPermissionResponseBodyData {
	s.BoundNodeGroupName = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetClassMode(v string) *GetRentedJoinPermissionResponseBodyData {
	s.ClassMode = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetCreateMillis(v int64) *GetRentedJoinPermissionResponseBodyData {
	s.CreateMillis = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetDataRate(v int64) *GetRentedJoinPermissionResponseBodyData {
	s.DataRate = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetEnabled(v bool) *GetRentedJoinPermissionResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetFreqBandPlanGroupId(v int64) *GetRentedJoinPermissionResponseBodyData {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetJoinEui(v string) *GetRentedJoinPermissionResponseBodyData {
	s.JoinEui = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetJoinPermissionId(v string) *GetRentedJoinPermissionResponseBodyData {
	s.JoinPermissionId = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetJoinPermissionName(v string) *GetRentedJoinPermissionResponseBodyData {
	s.JoinPermissionName = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetNodesCnt(v int64) *GetRentedJoinPermissionResponseBodyData {
	s.NodesCnt = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetRxDailySum(v int64) *GetRentedJoinPermissionResponseBodyData {
	s.RxDailySum = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetRxDelay(v int64) *GetRentedJoinPermissionResponseBodyData {
	s.RxDelay = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetRxMonthSum(v int64) *GetRentedJoinPermissionResponseBodyData {
	s.RxMonthSum = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetTxDailySum(v int64) *GetRentedJoinPermissionResponseBodyData {
	s.TxDailySum = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetTxMonthSum(v int64) *GetRentedJoinPermissionResponseBodyData {
	s.TxMonthSum = &v
	return s
}

func (s *GetRentedJoinPermissionResponseBodyData) SetType(v string) *GetRentedJoinPermissionResponseBodyData {
	s.Type = &v
	return s
}

type GetRentedJoinPermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRentedJoinPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRentedJoinPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRentedJoinPermissionResponse) GoString() string {
	return s.String()
}

func (s *GetRentedJoinPermissionResponse) SetHeaders(v map[string]*string) *GetRentedJoinPermissionResponse {
	s.Headers = v
	return s
}

func (s *GetRentedJoinPermissionResponse) SetStatusCode(v int32) *GetRentedJoinPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRentedJoinPermissionResponse) SetBody(v *GetRentedJoinPermissionResponseBody) *GetRentedJoinPermissionResponse {
	s.Body = v
	return s
}

type GetUserLicenseResponseBody struct {
	Data      *GetUserLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserLicenseResponseBody) SetData(v *GetUserLicenseResponseBodyData) *GetUserLicenseResponseBody {
	s.Data = v
	return s
}

func (s *GetUserLicenseResponseBody) SetRequestId(v string) *GetUserLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserLicenseResponseBody) SetSuccess(v bool) *GetUserLicenseResponseBody {
	s.Success = &v
	return s
}

type GetUserLicenseResponseBodyData struct {
	GatewayCount                   *int64  `json:"GatewayCount,omitempty" xml:"GatewayCount,omitempty"`
	GatewayDingTalkCount           *int64  `json:"GatewayDingTalkCount,omitempty" xml:"GatewayDingTalkCount,omitempty"`
	GatewayDingTalkLimit           *int64  `json:"GatewayDingTalkLimit,omitempty" xml:"GatewayDingTalkLimit,omitempty"`
	GatewayFreeLimit               *int64  `json:"GatewayFreeLimit,omitempty" xml:"GatewayFreeLimit,omitempty"`
	GatewayLimit                   *int64  `json:"GatewayLimit,omitempty" xml:"GatewayLimit,omitempty"`
	GatewayPrePayCount             *int64  `json:"GatewayPrePayCount,omitempty" xml:"GatewayPrePayCount,omitempty"`
	GatewayProfessionalCount       *int64  `json:"GatewayProfessionalCount,omitempty" xml:"GatewayProfessionalCount,omitempty"`
	GatewayProfessionalLimit       *int64  `json:"GatewayProfessionalLimit,omitempty" xml:"GatewayProfessionalLimit,omitempty"`
	GatewayTupleCount              *int64  `json:"GatewayTupleCount,omitempty" xml:"GatewayTupleCount,omitempty"`
	GatewayTupleFreeLimit          *int64  `json:"GatewayTupleFreeLimit,omitempty" xml:"GatewayTupleFreeLimit,omitempty"`
	GatewayTupleHybridCount        *int64  `json:"GatewayTupleHybridCount,omitempty" xml:"GatewayTupleHybridCount,omitempty"`
	GatewayTupleHybridLimit        *int64  `json:"GatewayTupleHybridLimit,omitempty" xml:"GatewayTupleHybridLimit,omitempty"`
	GatewayTupleLimit              *int64  `json:"GatewayTupleLimit,omitempty" xml:"GatewayTupleLimit,omitempty"`
	GatewayTupleSingleChannelCount *int64  `json:"GatewayTupleSingleChannelCount,omitempty" xml:"GatewayTupleSingleChannelCount,omitempty"`
	GatewayTupleSingleChannelLimit *int64  `json:"GatewayTupleSingleChannelLimit,omitempty" xml:"GatewayTupleSingleChannelLimit,omitempty"`
	GatewayTupleStandardCount      *int64  `json:"GatewayTupleStandardCount,omitempty" xml:"GatewayTupleStandardCount,omitempty"`
	GatewayTupleStandardLimit      *int64  `json:"GatewayTupleStandardLimit,omitempty" xml:"GatewayTupleStandardLimit,omitempty"`
	LocalJoinPermissionCount       *int64  `json:"LocalJoinPermissionCount,omitempty" xml:"LocalJoinPermissionCount,omitempty"`
	LocalJoinPermissionFreeLimit   *int64  `json:"LocalJoinPermissionFreeLimit,omitempty" xml:"LocalJoinPermissionFreeLimit,omitempty"`
	LocalJoinPermissionLimit       *int64  `json:"LocalJoinPermissionLimit,omitempty" xml:"LocalJoinPermissionLimit,omitempty"`
	NodeCount                      *int64  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeFreeLimit                  *int64  `json:"NodeFreeLimit,omitempty" xml:"NodeFreeLimit,omitempty"`
	NodeLimit                      *int64  `json:"NodeLimit,omitempty" xml:"NodeLimit,omitempty"`
	NodeTupleCount                 *int64  `json:"NodeTupleCount,omitempty" xml:"NodeTupleCount,omitempty"`
	NodeTupleFreeLimit             *int64  `json:"NodeTupleFreeLimit,omitempty" xml:"NodeTupleFreeLimit,omitempty"`
	NodeTupleLimit                 *int64  `json:"NodeTupleLimit,omitempty" xml:"NodeTupleLimit,omitempty"`
	NodeTupleRelayCount            *int64  `json:"NodeTupleRelayCount,omitempty" xml:"NodeTupleRelayCount,omitempty"`
	NodeTupleRelayLimit            *int64  `json:"NodeTupleRelayLimit,omitempty" xml:"NodeTupleRelayLimit,omitempty"`
	NodeTupleStandardCount         *int64  `json:"NodeTupleStandardCount,omitempty" xml:"NodeTupleStandardCount,omitempty"`
	NodeTupleStandardLimit         *int64  `json:"NodeTupleStandardLimit,omitempty" xml:"NodeTupleStandardLimit,omitempty"`
	Oui                            *string `json:"Oui,omitempty" xml:"Oui,omitempty"`
	RelayCount                     *int64  `json:"RelayCount,omitempty" xml:"RelayCount,omitempty"`
	RelayLimit                     *int64  `json:"RelayLimit,omitempty" xml:"RelayLimit,omitempty"`
	RoamingJoinPermissionCount     *int64  `json:"RoamingJoinPermissionCount,omitempty" xml:"RoamingJoinPermissionCount,omitempty"`
	RoamingJoinPermissionFreeLimit *int64  `json:"RoamingJoinPermissionFreeLimit,omitempty" xml:"RoamingJoinPermissionFreeLimit,omitempty"`
	RoamingJoinPermissionLimit     *int64  `json:"RoamingJoinPermissionLimit,omitempty" xml:"RoamingJoinPermissionLimit,omitempty"`
}

func (s GetUserLicenseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserLicenseResponseBodyData) SetGatewayCount(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayDingTalkCount(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayDingTalkCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayDingTalkLimit(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayDingTalkLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayFreeLimit(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayFreeLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayLimit(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayPrePayCount(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayPrePayCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayProfessionalCount(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayProfessionalCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayProfessionalLimit(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayProfessionalLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayTupleCount(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayTupleCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayTupleFreeLimit(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayTupleFreeLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayTupleHybridCount(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayTupleHybridCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayTupleHybridLimit(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayTupleHybridLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayTupleLimit(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayTupleLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayTupleSingleChannelCount(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayTupleSingleChannelCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayTupleSingleChannelLimit(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayTupleSingleChannelLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayTupleStandardCount(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayTupleStandardCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetGatewayTupleStandardLimit(v int64) *GetUserLicenseResponseBodyData {
	s.GatewayTupleStandardLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetLocalJoinPermissionCount(v int64) *GetUserLicenseResponseBodyData {
	s.LocalJoinPermissionCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetLocalJoinPermissionFreeLimit(v int64) *GetUserLicenseResponseBodyData {
	s.LocalJoinPermissionFreeLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetLocalJoinPermissionLimit(v int64) *GetUserLicenseResponseBodyData {
	s.LocalJoinPermissionLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeCount(v int64) *GetUserLicenseResponseBodyData {
	s.NodeCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeFreeLimit(v int64) *GetUserLicenseResponseBodyData {
	s.NodeFreeLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeLimit(v int64) *GetUserLicenseResponseBodyData {
	s.NodeLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeTupleCount(v int64) *GetUserLicenseResponseBodyData {
	s.NodeTupleCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeTupleFreeLimit(v int64) *GetUserLicenseResponseBodyData {
	s.NodeTupleFreeLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeTupleLimit(v int64) *GetUserLicenseResponseBodyData {
	s.NodeTupleLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeTupleRelayCount(v int64) *GetUserLicenseResponseBodyData {
	s.NodeTupleRelayCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeTupleRelayLimit(v int64) *GetUserLicenseResponseBodyData {
	s.NodeTupleRelayLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeTupleStandardCount(v int64) *GetUserLicenseResponseBodyData {
	s.NodeTupleStandardCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetNodeTupleStandardLimit(v int64) *GetUserLicenseResponseBodyData {
	s.NodeTupleStandardLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetOui(v string) *GetUserLicenseResponseBodyData {
	s.Oui = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetRelayCount(v int64) *GetUserLicenseResponseBodyData {
	s.RelayCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetRelayLimit(v int64) *GetUserLicenseResponseBodyData {
	s.RelayLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetRoamingJoinPermissionCount(v int64) *GetUserLicenseResponseBodyData {
	s.RoamingJoinPermissionCount = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetRoamingJoinPermissionFreeLimit(v int64) *GetUserLicenseResponseBodyData {
	s.RoamingJoinPermissionFreeLimit = &v
	return s
}

func (s *GetUserLicenseResponseBodyData) SetRoamingJoinPermissionLimit(v int64) *GetUserLicenseResponseBodyData {
	s.RoamingJoinPermissionLimit = &v
	return s
}

type GetUserLicenseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetUserLicenseResponse) SetHeaders(v map[string]*string) *GetUserLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetUserLicenseResponse) SetStatusCode(v int32) *GetUserLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserLicenseResponse) SetBody(v *GetUserLicenseResponseBody) *GetUserLicenseResponse {
	s.Body = v
	return s
}

type ListActivatedFeaturesRequest struct {
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
}

func (s ListActivatedFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListActivatedFeaturesRequest) SetEnvironment(v string) *ListActivatedFeaturesRequest {
	s.Environment = &v
	return s
}

type ListActivatedFeaturesResponseBody struct {
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListActivatedFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListActivatedFeaturesResponseBody) SetData(v []*string) *ListActivatedFeaturesResponseBody {
	s.Data = v
	return s
}

func (s *ListActivatedFeaturesResponseBody) SetRequestId(v string) *ListActivatedFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActivatedFeaturesResponseBody) SetSuccess(v bool) *ListActivatedFeaturesResponseBody {
	s.Success = &v
	return s
}

type ListActivatedFeaturesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListActivatedFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListActivatedFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListActivatedFeaturesResponse) SetHeaders(v map[string]*string) *ListActivatedFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListActivatedFeaturesResponse) SetStatusCode(v int32) *ListActivatedFeaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActivatedFeaturesResponse) SetBody(v *ListActivatedFeaturesResponseBody) *ListActivatedFeaturesResponse {
	s.Body = v
	return s
}

type ListActiveGatewaysResponseBody struct {
	Data      []*ListActiveGatewaysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListActiveGatewaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActiveGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *ListActiveGatewaysResponseBody) SetData(v []*ListActiveGatewaysResponseBodyData) *ListActiveGatewaysResponseBody {
	s.Data = v
	return s
}

func (s *ListActiveGatewaysResponseBody) SetRequestId(v string) *ListActiveGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActiveGatewaysResponseBody) SetSuccess(v bool) *ListActiveGatewaysResponseBody {
	s.Success = &v
	return s
}

type ListActiveGatewaysResponseBodyData struct {
	Address             *string  `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressCode         *int64   `json:"AddressCode,omitempty" xml:"AddressCode,omitempty"`
	ChargeStatus        *string  `json:"ChargeStatus,omitempty" xml:"ChargeStatus,omitempty"`
	City                *string  `json:"City,omitempty" xml:"City,omitempty"`
	CommunicationMode   *string  `json:"CommunicationMode,omitempty" xml:"CommunicationMode,omitempty"`
	Description         *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	District            *string  `json:"District,omitempty" xml:"District,omitempty"`
	FreqBandPlanGroupId *int64   `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	GisCoordinateSystem *string  `json:"GisCoordinateSystem,omitempty" xml:"GisCoordinateSystem,omitempty"`
	GwEui               *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	Latitude            *float32 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude           *float32 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name                *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OnlineState         *string  `json:"OnlineState,omitempty" xml:"OnlineState,omitempty"`
}

func (s ListActiveGatewaysResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListActiveGatewaysResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListActiveGatewaysResponseBodyData) SetAddress(v string) *ListActiveGatewaysResponseBodyData {
	s.Address = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetAddressCode(v int64) *ListActiveGatewaysResponseBodyData {
	s.AddressCode = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetChargeStatus(v string) *ListActiveGatewaysResponseBodyData {
	s.ChargeStatus = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetCity(v string) *ListActiveGatewaysResponseBodyData {
	s.City = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetCommunicationMode(v string) *ListActiveGatewaysResponseBodyData {
	s.CommunicationMode = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetDescription(v string) *ListActiveGatewaysResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetDistrict(v string) *ListActiveGatewaysResponseBodyData {
	s.District = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetFreqBandPlanGroupId(v int64) *ListActiveGatewaysResponseBodyData {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetGisCoordinateSystem(v string) *ListActiveGatewaysResponseBodyData {
	s.GisCoordinateSystem = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetGwEui(v string) *ListActiveGatewaysResponseBodyData {
	s.GwEui = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetLatitude(v float32) *ListActiveGatewaysResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetLongitude(v float32) *ListActiveGatewaysResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetName(v string) *ListActiveGatewaysResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListActiveGatewaysResponseBodyData) SetOnlineState(v string) *ListActiveGatewaysResponseBodyData {
	s.OnlineState = &v
	return s
}

type ListActiveGatewaysResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListActiveGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListActiveGatewaysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActiveGatewaysResponse) GoString() string {
	return s.String()
}

func (s *ListActiveGatewaysResponse) SetHeaders(v map[string]*string) *ListActiveGatewaysResponse {
	s.Headers = v
	return s
}

func (s *ListActiveGatewaysResponse) SetStatusCode(v int32) *ListActiveGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActiveGatewaysResponse) SetBody(v *ListActiveGatewaysResponseBody) *ListActiveGatewaysResponse {
	s.Body = v
	return s
}

type ListFreqBandPlanGroupsResponseBody struct {
	Data      []*ListFreqBandPlanGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFreqBandPlanGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFreqBandPlanGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFreqBandPlanGroupsResponseBody) SetData(v []*ListFreqBandPlanGroupsResponseBodyData) *ListFreqBandPlanGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListFreqBandPlanGroupsResponseBody) SetRequestId(v string) *ListFreqBandPlanGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFreqBandPlanGroupsResponseBody) SetSuccess(v bool) *ListFreqBandPlanGroupsResponseBody {
	s.Success = &v
	return s
}

type ListFreqBandPlanGroupsResponseBodyData struct {
	BeginFrequency    *int64  `json:"BeginFrequency,omitempty" xml:"BeginFrequency,omitempty"`
	EndFrequency      *int64  `json:"EndFrequency,omitempty" xml:"EndFrequency,omitempty"`
	FrequencyRegionId *string `json:"FrequencyRegionId,omitempty" xml:"FrequencyRegionId,omitempty"`
	FrequencyType     *string `json:"FrequencyType,omitempty" xml:"FrequencyType,omitempty"`
	GroupId           *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ListFreqBandPlanGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFreqBandPlanGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFreqBandPlanGroupsResponseBodyData) SetBeginFrequency(v int64) *ListFreqBandPlanGroupsResponseBodyData {
	s.BeginFrequency = &v
	return s
}

func (s *ListFreqBandPlanGroupsResponseBodyData) SetEndFrequency(v int64) *ListFreqBandPlanGroupsResponseBodyData {
	s.EndFrequency = &v
	return s
}

func (s *ListFreqBandPlanGroupsResponseBodyData) SetFrequencyRegionId(v string) *ListFreqBandPlanGroupsResponseBodyData {
	s.FrequencyRegionId = &v
	return s
}

func (s *ListFreqBandPlanGroupsResponseBodyData) SetFrequencyType(v string) *ListFreqBandPlanGroupsResponseBodyData {
	s.FrequencyType = &v
	return s
}

func (s *ListFreqBandPlanGroupsResponseBodyData) SetGroupId(v int64) *ListFreqBandPlanGroupsResponseBodyData {
	s.GroupId = &v
	return s
}

type ListFreqBandPlanGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFreqBandPlanGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFreqBandPlanGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFreqBandPlanGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListFreqBandPlanGroupsResponse) SetHeaders(v map[string]*string) *ListFreqBandPlanGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListFreqBandPlanGroupsResponse) SetStatusCode(v int32) *ListFreqBandPlanGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFreqBandPlanGroupsResponse) SetBody(v *ListFreqBandPlanGroupsResponseBody) *ListFreqBandPlanGroupsResponse {
	s.Body = v
	return s
}

type ListGatewayOnlineRecordsRequest struct {
	Ascending    *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	GwEui        *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	Limit        *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OffSet       *int64  `json:"OffSet,omitempty" xml:"OffSet,omitempty"`
	SortingField *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListGatewayOnlineRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayOnlineRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayOnlineRecordsRequest) SetAscending(v bool) *ListGatewayOnlineRecordsRequest {
	s.Ascending = &v
	return s
}

func (s *ListGatewayOnlineRecordsRequest) SetGwEui(v string) *ListGatewayOnlineRecordsRequest {
	s.GwEui = &v
	return s
}

func (s *ListGatewayOnlineRecordsRequest) SetLimit(v int64) *ListGatewayOnlineRecordsRequest {
	s.Limit = &v
	return s
}

func (s *ListGatewayOnlineRecordsRequest) SetOffSet(v int64) *ListGatewayOnlineRecordsRequest {
	s.OffSet = &v
	return s
}

func (s *ListGatewayOnlineRecordsRequest) SetSortingField(v string) *ListGatewayOnlineRecordsRequest {
	s.SortingField = &v
	return s
}

type ListGatewayOnlineRecordsResponseBody struct {
	Data      *ListGatewayOnlineRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayOnlineRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayOnlineRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayOnlineRecordsResponseBody) SetData(v *ListGatewayOnlineRecordsResponseBodyData) *ListGatewayOnlineRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayOnlineRecordsResponseBody) SetRequestId(v string) *ListGatewayOnlineRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayOnlineRecordsResponseBody) SetSuccess(v bool) *ListGatewayOnlineRecordsResponseBody {
	s.Success = &v
	return s
}

type ListGatewayOnlineRecordsResponseBodyData struct {
	List       []*ListGatewayOnlineRecordsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGatewayOnlineRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayOnlineRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayOnlineRecordsResponseBodyData) SetList(v []*ListGatewayOnlineRecordsResponseBodyDataList) *ListGatewayOnlineRecordsResponseBodyData {
	s.List = v
	return s
}

func (s *ListGatewayOnlineRecordsResponseBodyData) SetTotalCount(v int64) *ListGatewayOnlineRecordsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListGatewayOnlineRecordsResponseBodyDataList struct {
	GwEui              *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	OnlineState        *string `json:"OnlineState,omitempty" xml:"OnlineState,omitempty"`
	StateChangedMillis *int64  `json:"StateChangedMillis,omitempty" xml:"StateChangedMillis,omitempty"`
}

func (s ListGatewayOnlineRecordsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayOnlineRecordsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListGatewayOnlineRecordsResponseBodyDataList) SetGwEui(v string) *ListGatewayOnlineRecordsResponseBodyDataList {
	s.GwEui = &v
	return s
}

func (s *ListGatewayOnlineRecordsResponseBodyDataList) SetOnlineState(v string) *ListGatewayOnlineRecordsResponseBodyDataList {
	s.OnlineState = &v
	return s
}

func (s *ListGatewayOnlineRecordsResponseBodyDataList) SetStateChangedMillis(v int64) *ListGatewayOnlineRecordsResponseBodyDataList {
	s.StateChangedMillis = &v
	return s
}

type ListGatewayOnlineRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGatewayOnlineRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGatewayOnlineRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayOnlineRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayOnlineRecordsResponse) SetHeaders(v map[string]*string) *ListGatewayOnlineRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayOnlineRecordsResponse) SetStatusCode(v int32) *ListGatewayOnlineRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayOnlineRecordsResponse) SetBody(v *ListGatewayOnlineRecordsResponseBody) *ListGatewayOnlineRecordsResponse {
	s.Body = v
	return s
}

type ListGatewayTransferFlowStatsRequest struct {
	BeginMillis      *int64  `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	EndMillis        *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	GwEui            *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	TimeIntervalUnit *string `json:"TimeIntervalUnit,omitempty" xml:"TimeIntervalUnit,omitempty"`
}

func (s ListGatewayTransferFlowStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTransferFlowStatsRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayTransferFlowStatsRequest) SetBeginMillis(v int64) *ListGatewayTransferFlowStatsRequest {
	s.BeginMillis = &v
	return s
}

func (s *ListGatewayTransferFlowStatsRequest) SetEndMillis(v int64) *ListGatewayTransferFlowStatsRequest {
	s.EndMillis = &v
	return s
}

func (s *ListGatewayTransferFlowStatsRequest) SetGwEui(v string) *ListGatewayTransferFlowStatsRequest {
	s.GwEui = &v
	return s
}

func (s *ListGatewayTransferFlowStatsRequest) SetIotInstanceId(v string) *ListGatewayTransferFlowStatsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListGatewayTransferFlowStatsRequest) SetTimeIntervalUnit(v string) *ListGatewayTransferFlowStatsRequest {
	s.TimeIntervalUnit = &v
	return s
}

type ListGatewayTransferFlowStatsResponseBody struct {
	Data      []*ListGatewayTransferFlowStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayTransferFlowStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTransferFlowStatsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayTransferFlowStatsResponseBody) SetData(v []*ListGatewayTransferFlowStatsResponseBodyData) *ListGatewayTransferFlowStatsResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayTransferFlowStatsResponseBody) SetRequestId(v string) *ListGatewayTransferFlowStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayTransferFlowStatsResponseBody) SetSuccess(v bool) *ListGatewayTransferFlowStatsResponseBody {
	s.Success = &v
	return s
}

type ListGatewayTransferFlowStatsResponseBodyData struct {
	DownlinkCount *int64  `json:"DownlinkCount,omitempty" xml:"DownlinkCount,omitempty"`
	StatMillis    *string `json:"StatMillis,omitempty" xml:"StatMillis,omitempty"`
	UplinkCount   *int64  `json:"UplinkCount,omitempty" xml:"UplinkCount,omitempty"`
}

func (s ListGatewayTransferFlowStatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTransferFlowStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayTransferFlowStatsResponseBodyData) SetDownlinkCount(v int64) *ListGatewayTransferFlowStatsResponseBodyData {
	s.DownlinkCount = &v
	return s
}

func (s *ListGatewayTransferFlowStatsResponseBodyData) SetStatMillis(v string) *ListGatewayTransferFlowStatsResponseBodyData {
	s.StatMillis = &v
	return s
}

func (s *ListGatewayTransferFlowStatsResponseBodyData) SetUplinkCount(v int64) *ListGatewayTransferFlowStatsResponseBodyData {
	s.UplinkCount = &v
	return s
}

type ListGatewayTransferFlowStatsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGatewayTransferFlowStatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGatewayTransferFlowStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTransferFlowStatsResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayTransferFlowStatsResponse) SetHeaders(v map[string]*string) *ListGatewayTransferFlowStatsResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayTransferFlowStatsResponse) SetStatusCode(v int32) *ListGatewayTransferFlowStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayTransferFlowStatsResponse) SetBody(v *ListGatewayTransferFlowStatsResponseBody) *ListGatewayTransferFlowStatsResponse {
	s.Body = v
	return s
}

type ListGatewayTransferPacketsRequest struct {
	Ascending     *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	BeginMillis   *int64  `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DevEui        *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	EndMillis     *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	GwEui         *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortingField  *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListGatewayTransferPacketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTransferPacketsRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayTransferPacketsRequest) SetAscending(v bool) *ListGatewayTransferPacketsRequest {
	s.Ascending = &v
	return s
}

func (s *ListGatewayTransferPacketsRequest) SetBeginMillis(v int64) *ListGatewayTransferPacketsRequest {
	s.BeginMillis = &v
	return s
}

func (s *ListGatewayTransferPacketsRequest) SetCategory(v string) *ListGatewayTransferPacketsRequest {
	s.Category = &v
	return s
}

func (s *ListGatewayTransferPacketsRequest) SetDevEui(v string) *ListGatewayTransferPacketsRequest {
	s.DevEui = &v
	return s
}

func (s *ListGatewayTransferPacketsRequest) SetEndMillis(v int64) *ListGatewayTransferPacketsRequest {
	s.EndMillis = &v
	return s
}

func (s *ListGatewayTransferPacketsRequest) SetGwEui(v string) *ListGatewayTransferPacketsRequest {
	s.GwEui = &v
	return s
}

func (s *ListGatewayTransferPacketsRequest) SetIotInstanceId(v string) *ListGatewayTransferPacketsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListGatewayTransferPacketsRequest) SetPageNumber(v int32) *ListGatewayTransferPacketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayTransferPacketsRequest) SetPageSize(v int32) *ListGatewayTransferPacketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayTransferPacketsRequest) SetSortingField(v string) *ListGatewayTransferPacketsRequest {
	s.SortingField = &v
	return s
}

type ListGatewayTransferPacketsResponseBody struct {
	Data      *ListGatewayTransferPacketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayTransferPacketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTransferPacketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayTransferPacketsResponseBody) SetData(v *ListGatewayTransferPacketsResponseBodyData) *ListGatewayTransferPacketsResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayTransferPacketsResponseBody) SetRequestId(v string) *ListGatewayTransferPacketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBody) SetSuccess(v bool) *ListGatewayTransferPacketsResponseBody {
	s.Success = &v
	return s
}

type ListGatewayTransferPacketsResponseBodyData struct {
	List       []*ListGatewayTransferPacketsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGatewayTransferPacketsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTransferPacketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayTransferPacketsResponseBodyData) SetList(v []*ListGatewayTransferPacketsResponseBodyDataList) *ListGatewayTransferPacketsResponseBodyData {
	s.List = v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyData) SetTotalCount(v int64) *ListGatewayTransferPacketsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListGatewayTransferPacketsResponseBodyDataList struct {
	Base64EncodedMacPayload *string                  `json:"Base64EncodedMacPayload,omitempty" xml:"Base64EncodedMacPayload,omitempty"`
	ClassMode               *string                  `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	Datr                    *string                  `json:"Datr,omitempty" xml:"Datr,omitempty"`
	DevAddr                 *string                  `json:"DevAddr,omitempty" xml:"DevAddr,omitempty"`
	DevEui                  *string                  `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	FPort                   *int32                   `json:"FPort,omitempty" xml:"FPort,omitempty"`
	Freq                    *string                  `json:"Freq,omitempty" xml:"Freq,omitempty"`
	GwEui                   *string                  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	HasData                 *bool                    `json:"HasData,omitempty" xml:"HasData,omitempty"`
	HasMacCommand           *bool                    `json:"HasMacCommand,omitempty" xml:"HasMacCommand,omitempty"`
	LogMillis               *string                  `json:"LogMillis,omitempty" xml:"LogMillis,omitempty"`
	Lsnr                    *float32                 `json:"Lsnr,omitempty" xml:"Lsnr,omitempty"`
	MacCommandCIDs          []map[string]interface{} `json:"MacCommandCIDs,omitempty" xml:"MacCommandCIDs,omitempty" type:"Repeated"`
	MacPayloadSize          *int64                   `json:"MacPayloadSize,omitempty" xml:"MacPayloadSize,omitempty"`
	MessageType             *string                  `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	NodeOwnerAliyunId       *string                  `json:"NodeOwnerAliyunId,omitempty" xml:"NodeOwnerAliyunId,omitempty"`
	ProcessEvent            *string                  `json:"ProcessEvent,omitempty" xml:"ProcessEvent,omitempty"`
	Rssi                    *int32                   `json:"Rssi,omitempty" xml:"Rssi,omitempty"`
}

func (s ListGatewayTransferPacketsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTransferPacketsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetBase64EncodedMacPayload(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.Base64EncodedMacPayload = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetClassMode(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.ClassMode = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetDatr(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.Datr = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetDevAddr(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.DevAddr = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetDevEui(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.DevEui = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetFPort(v int32) *ListGatewayTransferPacketsResponseBodyDataList {
	s.FPort = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetFreq(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.Freq = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetGwEui(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.GwEui = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetHasData(v bool) *ListGatewayTransferPacketsResponseBodyDataList {
	s.HasData = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetHasMacCommand(v bool) *ListGatewayTransferPacketsResponseBodyDataList {
	s.HasMacCommand = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetLogMillis(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.LogMillis = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetLsnr(v float32) *ListGatewayTransferPacketsResponseBodyDataList {
	s.Lsnr = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetMacCommandCIDs(v []map[string]interface{}) *ListGatewayTransferPacketsResponseBodyDataList {
	s.MacCommandCIDs = v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetMacPayloadSize(v int64) *ListGatewayTransferPacketsResponseBodyDataList {
	s.MacPayloadSize = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetMessageType(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.MessageType = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetNodeOwnerAliyunId(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.NodeOwnerAliyunId = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetProcessEvent(v string) *ListGatewayTransferPacketsResponseBodyDataList {
	s.ProcessEvent = &v
	return s
}

func (s *ListGatewayTransferPacketsResponseBodyDataList) SetRssi(v int32) *ListGatewayTransferPacketsResponseBodyDataList {
	s.Rssi = &v
	return s
}

type ListGatewayTransferPacketsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGatewayTransferPacketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGatewayTransferPacketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTransferPacketsResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayTransferPacketsResponse) SetHeaders(v map[string]*string) *ListGatewayTransferPacketsResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayTransferPacketsResponse) SetStatusCode(v int32) *ListGatewayTransferPacketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayTransferPacketsResponse) SetBody(v *ListGatewayTransferPacketsResponseBody) *ListGatewayTransferPacketsResponse {
	s.Body = v
	return s
}

type ListGatewayTupleOrdersRequest struct {
	Ascending    *bool     `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	Limit        *int64    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset       *int64    `json:"Offset,omitempty" xml:"Offset,omitempty"`
	SortingField *string   `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
	State        []*string `json:"State,omitempty" xml:"State,omitempty" type:"Repeated"`
}

func (s ListGatewayTupleOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTupleOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayTupleOrdersRequest) SetAscending(v bool) *ListGatewayTupleOrdersRequest {
	s.Ascending = &v
	return s
}

func (s *ListGatewayTupleOrdersRequest) SetLimit(v int64) *ListGatewayTupleOrdersRequest {
	s.Limit = &v
	return s
}

func (s *ListGatewayTupleOrdersRequest) SetOffset(v int64) *ListGatewayTupleOrdersRequest {
	s.Offset = &v
	return s
}

func (s *ListGatewayTupleOrdersRequest) SetSortingField(v string) *ListGatewayTupleOrdersRequest {
	s.SortingField = &v
	return s
}

func (s *ListGatewayTupleOrdersRequest) SetState(v []*string) *ListGatewayTupleOrdersRequest {
	s.State = v
	return s
}

type ListGatewayTupleOrdersResponseBody struct {
	Data      *ListGatewayTupleOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayTupleOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTupleOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayTupleOrdersResponseBody) SetData(v *ListGatewayTupleOrdersResponseBodyData) *ListGatewayTupleOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayTupleOrdersResponseBody) SetRequestId(v string) *ListGatewayTupleOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayTupleOrdersResponseBody) SetSuccess(v bool) *ListGatewayTupleOrdersResponseBody {
	s.Success = &v
	return s
}

type ListGatewayTupleOrdersResponseBodyData struct {
	List       []*ListGatewayTupleOrdersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGatewayTupleOrdersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTupleOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayTupleOrdersResponseBodyData) SetList(v []*ListGatewayTupleOrdersResponseBodyDataList) *ListGatewayTupleOrdersResponseBodyData {
	s.List = v
	return s
}

func (s *ListGatewayTupleOrdersResponseBodyData) SetTotalCount(v int64) *ListGatewayTupleOrdersResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListGatewayTupleOrdersResponseBodyDataList struct {
	AcceptedMillis *int64  `json:"AcceptedMillis,omitempty" xml:"AcceptedMillis,omitempty"`
	CreatedMillis  *int64  `json:"CreatedMillis,omitempty" xml:"CreatedMillis,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderState     *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	RequiredCount  *int64  `json:"RequiredCount,omitempty" xml:"RequiredCount,omitempty"`
	TupleType      *string `json:"TupleType,omitempty" xml:"TupleType,omitempty"`
}

func (s ListGatewayTupleOrdersResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTupleOrdersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListGatewayTupleOrdersResponseBodyDataList) SetAcceptedMillis(v int64) *ListGatewayTupleOrdersResponseBodyDataList {
	s.AcceptedMillis = &v
	return s
}

func (s *ListGatewayTupleOrdersResponseBodyDataList) SetCreatedMillis(v int64) *ListGatewayTupleOrdersResponseBodyDataList {
	s.CreatedMillis = &v
	return s
}

func (s *ListGatewayTupleOrdersResponseBodyDataList) SetOrderId(v string) *ListGatewayTupleOrdersResponseBodyDataList {
	s.OrderId = &v
	return s
}

func (s *ListGatewayTupleOrdersResponseBodyDataList) SetOrderState(v string) *ListGatewayTupleOrdersResponseBodyDataList {
	s.OrderState = &v
	return s
}

func (s *ListGatewayTupleOrdersResponseBodyDataList) SetRequiredCount(v int64) *ListGatewayTupleOrdersResponseBodyDataList {
	s.RequiredCount = &v
	return s
}

func (s *ListGatewayTupleOrdersResponseBodyDataList) SetTupleType(v string) *ListGatewayTupleOrdersResponseBodyDataList {
	s.TupleType = &v
	return s
}

type ListGatewayTupleOrdersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGatewayTupleOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGatewayTupleOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayTupleOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayTupleOrdersResponse) SetHeaders(v map[string]*string) *ListGatewayTupleOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayTupleOrdersResponse) SetStatusCode(v int32) *ListGatewayTupleOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayTupleOrdersResponse) SetBody(v *ListGatewayTupleOrdersResponseBody) *ListGatewayTupleOrdersResponse {
	s.Body = v
	return s
}

type ListGatewaysRequest struct {
	Ascending           *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	FreqBandPlanGroupId *int64  `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	FuzzyCity           *string `json:"FuzzyCity,omitempty" xml:"FuzzyCity,omitempty"`
	FuzzyGwEui          *string `json:"FuzzyGwEui,omitempty" xml:"FuzzyGwEui,omitempty"`
	FuzzyName           *string `json:"FuzzyName,omitempty" xml:"FuzzyName,omitempty"`
	IotInstanceId       *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	IsEnabled           *bool   `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	Limit               *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset              *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OnlineState         *string `json:"OnlineState,omitempty" xml:"OnlineState,omitempty"`
	SortingField        *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListGatewaysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysRequest) GoString() string {
	return s.String()
}

func (s *ListGatewaysRequest) SetAscending(v bool) *ListGatewaysRequest {
	s.Ascending = &v
	return s
}

func (s *ListGatewaysRequest) SetFreqBandPlanGroupId(v int64) *ListGatewaysRequest {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *ListGatewaysRequest) SetFuzzyCity(v string) *ListGatewaysRequest {
	s.FuzzyCity = &v
	return s
}

func (s *ListGatewaysRequest) SetFuzzyGwEui(v string) *ListGatewaysRequest {
	s.FuzzyGwEui = &v
	return s
}

func (s *ListGatewaysRequest) SetFuzzyName(v string) *ListGatewaysRequest {
	s.FuzzyName = &v
	return s
}

func (s *ListGatewaysRequest) SetIotInstanceId(v string) *ListGatewaysRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListGatewaysRequest) SetIsEnabled(v bool) *ListGatewaysRequest {
	s.IsEnabled = &v
	return s
}

func (s *ListGatewaysRequest) SetLimit(v int64) *ListGatewaysRequest {
	s.Limit = &v
	return s
}

func (s *ListGatewaysRequest) SetOffset(v int64) *ListGatewaysRequest {
	s.Offset = &v
	return s
}

func (s *ListGatewaysRequest) SetOnlineState(v string) *ListGatewaysRequest {
	s.OnlineState = &v
	return s
}

func (s *ListGatewaysRequest) SetSortingField(v string) *ListGatewaysRequest {
	s.SortingField = &v
	return s
}

type ListGatewaysResponseBody struct {
	Data      *ListGatewaysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBody) SetData(v *ListGatewaysResponseBodyData) *ListGatewaysResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewaysResponseBody) SetRequestId(v string) *ListGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewaysResponseBody) SetSuccess(v bool) *ListGatewaysResponseBody {
	s.Success = &v
	return s
}

type ListGatewaysResponseBodyData struct {
	List       []*ListGatewaysResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGatewaysResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyData) SetList(v []*ListGatewaysResponseBodyDataList) *ListGatewaysResponseBodyData {
	s.List = v
	return s
}

func (s *ListGatewaysResponseBodyData) SetTotalCount(v int64) *ListGatewaysResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListGatewaysResponseBodyDataList struct {
	Address                  *string  `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressCode              *int64   `json:"AddressCode,omitempty" xml:"AddressCode,omitempty"`
	AuthTypes                *string  `json:"AuthTypes,omitempty" xml:"AuthTypes,omitempty"`
	ChargeStatus             *string  `json:"ChargeStatus,omitempty" xml:"ChargeStatus,omitempty"`
	City                     *string  `json:"City,omitempty" xml:"City,omitempty"`
	ClassBSupported          *bool    `json:"ClassBSupported,omitempty" xml:"ClassBSupported,omitempty"`
	ClassBWorking            *bool    `json:"ClassBWorking,omitempty" xml:"ClassBWorking,omitempty"`
	CommunicationMode        *string  `json:"CommunicationMode,omitempty" xml:"CommunicationMode,omitempty"`
	Description              *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	District                 *string  `json:"District,omitempty" xml:"District,omitempty"`
	EmbeddedNsId             *string  `json:"EmbeddedNsId,omitempty" xml:"EmbeddedNsId,omitempty"`
	Enabled                  *bool    `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FreqBandPlanGroupId      *int64   `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	GisCoordinateSystem      *string  `json:"GisCoordinateSystem,omitempty" xml:"GisCoordinateSystem,omitempty"`
	GwEui                    *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	Latitude                 *float32 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude                *float32 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name                     *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OnlineState              *string  `json:"OnlineState,omitempty" xml:"OnlineState,omitempty"`
	OnlineStateChangedMillis *int64   `json:"OnlineStateChangedMillis,omitempty" xml:"OnlineStateChangedMillis,omitempty"`
	TimeCorrectable          *bool    `json:"TimeCorrectable,omitempty" xml:"TimeCorrectable,omitempty"`
}

func (s ListGatewaysResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponseBodyDataList) SetAddress(v string) *ListGatewaysResponseBodyDataList {
	s.Address = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetAddressCode(v int64) *ListGatewaysResponseBodyDataList {
	s.AddressCode = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetAuthTypes(v string) *ListGatewaysResponseBodyDataList {
	s.AuthTypes = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetChargeStatus(v string) *ListGatewaysResponseBodyDataList {
	s.ChargeStatus = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetCity(v string) *ListGatewaysResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetClassBSupported(v bool) *ListGatewaysResponseBodyDataList {
	s.ClassBSupported = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetClassBWorking(v bool) *ListGatewaysResponseBodyDataList {
	s.ClassBWorking = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetCommunicationMode(v string) *ListGatewaysResponseBodyDataList {
	s.CommunicationMode = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetDescription(v string) *ListGatewaysResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetDistrict(v string) *ListGatewaysResponseBodyDataList {
	s.District = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetEmbeddedNsId(v string) *ListGatewaysResponseBodyDataList {
	s.EmbeddedNsId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetEnabled(v bool) *ListGatewaysResponseBodyDataList {
	s.Enabled = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetFreqBandPlanGroupId(v int64) *ListGatewaysResponseBodyDataList {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetGisCoordinateSystem(v string) *ListGatewaysResponseBodyDataList {
	s.GisCoordinateSystem = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetGwEui(v string) *ListGatewaysResponseBodyDataList {
	s.GwEui = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetLatitude(v float32) *ListGatewaysResponseBodyDataList {
	s.Latitude = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetLongitude(v float32) *ListGatewaysResponseBodyDataList {
	s.Longitude = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetName(v string) *ListGatewaysResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetOnlineState(v string) *ListGatewaysResponseBodyDataList {
	s.OnlineState = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetOnlineStateChangedMillis(v int64) *ListGatewaysResponseBodyDataList {
	s.OnlineStateChangedMillis = &v
	return s
}

func (s *ListGatewaysResponseBodyDataList) SetTimeCorrectable(v bool) *ListGatewaysResponseBodyDataList {
	s.TimeCorrectable = &v
	return s
}

type ListGatewaysResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGatewaysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysResponse) GoString() string {
	return s.String()
}

func (s *ListGatewaysResponse) SetHeaders(v map[string]*string) *ListGatewaysResponse {
	s.Headers = v
	return s
}

func (s *ListGatewaysResponse) SetStatusCode(v int32) *ListGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewaysResponse) SetBody(v *ListGatewaysResponseBody) *ListGatewaysResponse {
	s.Body = v
	return s
}

type ListGatewaysGisInfoRequest struct {
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s ListGatewaysGisInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysGisInfoRequest) GoString() string {
	return s.String()
}

func (s *ListGatewaysGisInfoRequest) SetIotInstanceId(v string) *ListGatewaysGisInfoRequest {
	s.IotInstanceId = &v
	return s
}

type ListGatewaysGisInfoResponseBody struct {
	Data      []*ListGatewaysGisInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewaysGisInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysGisInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewaysGisInfoResponseBody) SetData(v []*ListGatewaysGisInfoResponseBodyData) *ListGatewaysGisInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewaysGisInfoResponseBody) SetRequestId(v string) *ListGatewaysGisInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBody) SetSuccess(v bool) *ListGatewaysGisInfoResponseBody {
	s.Success = &v
	return s
}

type ListGatewaysGisInfoResponseBodyData struct {
	AuthTypes           *string  `json:"AuthTypes,omitempty" xml:"AuthTypes,omitempty"`
	ChargeStatus        *string  `json:"ChargeStatus,omitempty" xml:"ChargeStatus,omitempty"`
	Enabled             *bool    `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FreqBandPlanGroupId *int64   `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	GisCoordinateSystem *string  `json:"GisCoordinateSystem,omitempty" xml:"GisCoordinateSystem,omitempty"`
	GisSourceType       *string  `json:"GisSourceType,omitempty" xml:"GisSourceType,omitempty"`
	GwEui               *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	Latitude            *float32 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude           *float32 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name                *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OnlineState         *string  `json:"OnlineState,omitempty" xml:"OnlineState,omitempty"`
}

func (s ListGatewaysGisInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysGisInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewaysGisInfoResponseBodyData) SetAuthTypes(v string) *ListGatewaysGisInfoResponseBodyData {
	s.AuthTypes = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetChargeStatus(v string) *ListGatewaysGisInfoResponseBodyData {
	s.ChargeStatus = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetEnabled(v bool) *ListGatewaysGisInfoResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetFreqBandPlanGroupId(v int64) *ListGatewaysGisInfoResponseBodyData {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetGisCoordinateSystem(v string) *ListGatewaysGisInfoResponseBodyData {
	s.GisCoordinateSystem = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetGisSourceType(v string) *ListGatewaysGisInfoResponseBodyData {
	s.GisSourceType = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetGwEui(v string) *ListGatewaysGisInfoResponseBodyData {
	s.GwEui = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetLatitude(v float32) *ListGatewaysGisInfoResponseBodyData {
	s.Latitude = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetLongitude(v float32) *ListGatewaysGisInfoResponseBodyData {
	s.Longitude = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetName(v string) *ListGatewaysGisInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListGatewaysGisInfoResponseBodyData) SetOnlineState(v string) *ListGatewaysGisInfoResponseBodyData {
	s.OnlineState = &v
	return s
}

type ListGatewaysGisInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGatewaysGisInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGatewaysGisInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewaysGisInfoResponse) GoString() string {
	return s.String()
}

func (s *ListGatewaysGisInfoResponse) SetHeaders(v map[string]*string) *ListGatewaysGisInfoResponse {
	s.Headers = v
	return s
}

func (s *ListGatewaysGisInfoResponse) SetStatusCode(v int32) *ListGatewaysGisInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewaysGisInfoResponse) SetBody(v *ListGatewaysGisInfoResponseBody) *ListGatewaysGisInfoResponse {
	s.Body = v
	return s
}

type ListNodeGroupTransferFlowStatsRequest struct {
	BeginMillis      *int64  `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	EndMillis        *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	NodeGroupId      *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	TimeIntervalUnit *string `json:"TimeIntervalUnit,omitempty" xml:"TimeIntervalUnit,omitempty"`
}

func (s ListNodeGroupTransferFlowStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupTransferFlowStatsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupTransferFlowStatsRequest) SetBeginMillis(v int64) *ListNodeGroupTransferFlowStatsRequest {
	s.BeginMillis = &v
	return s
}

func (s *ListNodeGroupTransferFlowStatsRequest) SetEndMillis(v int64) *ListNodeGroupTransferFlowStatsRequest {
	s.EndMillis = &v
	return s
}

func (s *ListNodeGroupTransferFlowStatsRequest) SetIotInstanceId(v string) *ListNodeGroupTransferFlowStatsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListNodeGroupTransferFlowStatsRequest) SetNodeGroupId(v string) *ListNodeGroupTransferFlowStatsRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodeGroupTransferFlowStatsRequest) SetTimeIntervalUnit(v string) *ListNodeGroupTransferFlowStatsRequest {
	s.TimeIntervalUnit = &v
	return s
}

type ListNodeGroupTransferFlowStatsResponseBody struct {
	Data      []*ListNodeGroupTransferFlowStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeGroupTransferFlowStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupTransferFlowStatsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupTransferFlowStatsResponseBody) SetData(v []*ListNodeGroupTransferFlowStatsResponseBodyData) *ListNodeGroupTransferFlowStatsResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeGroupTransferFlowStatsResponseBody) SetRequestId(v string) *ListNodeGroupTransferFlowStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeGroupTransferFlowStatsResponseBody) SetSuccess(v bool) *ListNodeGroupTransferFlowStatsResponseBody {
	s.Success = &v
	return s
}

type ListNodeGroupTransferFlowStatsResponseBodyData struct {
	DownlinkCount *int64 `json:"DownlinkCount,omitempty" xml:"DownlinkCount,omitempty"`
	StatMillis    *int64 `json:"StatMillis,omitempty" xml:"StatMillis,omitempty"`
	UplinkCount   *int64 `json:"UplinkCount,omitempty" xml:"UplinkCount,omitempty"`
}

func (s ListNodeGroupTransferFlowStatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupTransferFlowStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeGroupTransferFlowStatsResponseBodyData) SetDownlinkCount(v int64) *ListNodeGroupTransferFlowStatsResponseBodyData {
	s.DownlinkCount = &v
	return s
}

func (s *ListNodeGroupTransferFlowStatsResponseBodyData) SetStatMillis(v int64) *ListNodeGroupTransferFlowStatsResponseBodyData {
	s.StatMillis = &v
	return s
}

func (s *ListNodeGroupTransferFlowStatsResponseBodyData) SetUplinkCount(v int64) *ListNodeGroupTransferFlowStatsResponseBodyData {
	s.UplinkCount = &v
	return s
}

type ListNodeGroupTransferFlowStatsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeGroupTransferFlowStatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeGroupTransferFlowStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupTransferFlowStatsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupTransferFlowStatsResponse) SetHeaders(v map[string]*string) *ListNodeGroupTransferFlowStatsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupTransferFlowStatsResponse) SetStatusCode(v int32) *ListNodeGroupTransferFlowStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupTransferFlowStatsResponse) SetBody(v *ListNodeGroupTransferFlowStatsResponseBody) *ListNodeGroupTransferFlowStatsResponse {
	s.Body = v
	return s
}

type ListNodeGroupTransferPacketsRequest struct {
	Ascending     *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	BeginMillis   *int64  `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DevEui        *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	EndMillis     *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	NodeGroupId   *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortingField  *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListNodeGroupTransferPacketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupTransferPacketsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupTransferPacketsRequest) SetAscending(v bool) *ListNodeGroupTransferPacketsRequest {
	s.Ascending = &v
	return s
}

func (s *ListNodeGroupTransferPacketsRequest) SetBeginMillis(v int64) *ListNodeGroupTransferPacketsRequest {
	s.BeginMillis = &v
	return s
}

func (s *ListNodeGroupTransferPacketsRequest) SetCategory(v string) *ListNodeGroupTransferPacketsRequest {
	s.Category = &v
	return s
}

func (s *ListNodeGroupTransferPacketsRequest) SetDevEui(v string) *ListNodeGroupTransferPacketsRequest {
	s.DevEui = &v
	return s
}

func (s *ListNodeGroupTransferPacketsRequest) SetEndMillis(v int64) *ListNodeGroupTransferPacketsRequest {
	s.EndMillis = &v
	return s
}

func (s *ListNodeGroupTransferPacketsRequest) SetIotInstanceId(v string) *ListNodeGroupTransferPacketsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListNodeGroupTransferPacketsRequest) SetNodeGroupId(v string) *ListNodeGroupTransferPacketsRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodeGroupTransferPacketsRequest) SetPageNumber(v int32) *ListNodeGroupTransferPacketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeGroupTransferPacketsRequest) SetPageSize(v int32) *ListNodeGroupTransferPacketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeGroupTransferPacketsRequest) SetSortingField(v string) *ListNodeGroupTransferPacketsRequest {
	s.SortingField = &v
	return s
}

type ListNodeGroupTransferPacketsResponseBody struct {
	Data      *ListNodeGroupTransferPacketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeGroupTransferPacketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupTransferPacketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupTransferPacketsResponseBody) SetData(v *ListNodeGroupTransferPacketsResponseBodyData) *ListNodeGroupTransferPacketsResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBody) SetRequestId(v string) *ListNodeGroupTransferPacketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBody) SetSuccess(v bool) *ListNodeGroupTransferPacketsResponseBody {
	s.Success = &v
	return s
}

type ListNodeGroupTransferPacketsResponseBodyData struct {
	List       []*ListNodeGroupTransferPacketsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeGroupTransferPacketsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupTransferPacketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeGroupTransferPacketsResponseBodyData) SetList(v []*ListNodeGroupTransferPacketsResponseBodyDataList) *ListNodeGroupTransferPacketsResponseBodyData {
	s.List = v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyData) SetTotalCount(v int64) *ListNodeGroupTransferPacketsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListNodeGroupTransferPacketsResponseBodyDataList struct {
	Base64EncodedMacPayload *string  `json:"Base64EncodedMacPayload,omitempty" xml:"Base64EncodedMacPayload,omitempty"`
	ClassMode               *string  `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	Datr                    *string  `json:"Datr,omitempty" xml:"Datr,omitempty"`
	DevAddr                 *string  `json:"DevAddr,omitempty" xml:"DevAddr,omitempty"`
	DevEui                  *string  `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	FPort                   *int32   `json:"FPort,omitempty" xml:"FPort,omitempty"`
	FcntDown                *int64   `json:"FcntDown,omitempty" xml:"FcntDown,omitempty"`
	FcntUp                  *int64   `json:"FcntUp,omitempty" xml:"FcntUp,omitempty"`
	Freq                    *float32 `json:"Freq,omitempty" xml:"Freq,omitempty"`
	FreqBandPlanGroupId     *int64   `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	GwEui                   *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	GwOwnerAliyunId         *string  `json:"GwOwnerAliyunId,omitempty" xml:"GwOwnerAliyunId,omitempty"`
	HasData                 *bool    `json:"HasData,omitempty" xml:"HasData,omitempty"`
	HasMacCommand           *bool    `json:"HasMacCommand,omitempty" xml:"HasMacCommand,omitempty"`
	LogMillis               *int64   `json:"LogMillis,omitempty" xml:"LogMillis,omitempty"`
	Lsnr                    *float32 `json:"Lsnr,omitempty" xml:"Lsnr,omitempty"`
	MacCommandCIDs          *string  `json:"MacCommandCIDs,omitempty" xml:"MacCommandCIDs,omitempty"`
	MacPayloadSize          *int64   `json:"MacPayloadSize,omitempty" xml:"MacPayloadSize,omitempty"`
	MessageType             *string  `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	ProcessEvent            *string  `json:"ProcessEvent,omitempty" xml:"ProcessEvent,omitempty"`
	Rssi                    *int32   `json:"Rssi,omitempty" xml:"Rssi,omitempty"`
}

func (s ListNodeGroupTransferPacketsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupTransferPacketsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetBase64EncodedMacPayload(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.Base64EncodedMacPayload = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetClassMode(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.ClassMode = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetDatr(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.Datr = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetDevAddr(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.DevAddr = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetDevEui(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.DevEui = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetFPort(v int32) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.FPort = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetFcntDown(v int64) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.FcntDown = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetFcntUp(v int64) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.FcntUp = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetFreq(v float32) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.Freq = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetFreqBandPlanGroupId(v int64) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetGwEui(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.GwEui = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetGwOwnerAliyunId(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.GwOwnerAliyunId = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetHasData(v bool) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.HasData = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetHasMacCommand(v bool) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.HasMacCommand = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetLogMillis(v int64) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.LogMillis = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetLsnr(v float32) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.Lsnr = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetMacCommandCIDs(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.MacCommandCIDs = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetMacPayloadSize(v int64) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.MacPayloadSize = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetMessageType(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.MessageType = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetProcessEvent(v string) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.ProcessEvent = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponseBodyDataList) SetRssi(v int32) *ListNodeGroupTransferPacketsResponseBodyDataList {
	s.Rssi = &v
	return s
}

type ListNodeGroupTransferPacketsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeGroupTransferPacketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeGroupTransferPacketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupTransferPacketsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupTransferPacketsResponse) SetHeaders(v map[string]*string) *ListNodeGroupTransferPacketsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupTransferPacketsResponse) SetStatusCode(v int32) *ListNodeGroupTransferPacketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupTransferPacketsResponse) SetBody(v *ListNodeGroupTransferPacketsResponseBody) *ListNodeGroupTransferPacketsResponse {
	s.Body = v
	return s
}

type ListNodeGroupsRequest struct {
	Ascending     *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	FuzzyDevEui   *string `json:"FuzzyDevEui,omitempty" xml:"FuzzyDevEui,omitempty"`
	FuzzyJoinEui  *string `json:"FuzzyJoinEui,omitempty" xml:"FuzzyJoinEui,omitempty"`
	FuzzyName     *string `json:"FuzzyName,omitempty" xml:"FuzzyName,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Limit         *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset        *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	SortingField  *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListNodeGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsRequest) SetAscending(v bool) *ListNodeGroupsRequest {
	s.Ascending = &v
	return s
}

func (s *ListNodeGroupsRequest) SetFuzzyDevEui(v string) *ListNodeGroupsRequest {
	s.FuzzyDevEui = &v
	return s
}

func (s *ListNodeGroupsRequest) SetFuzzyJoinEui(v string) *ListNodeGroupsRequest {
	s.FuzzyJoinEui = &v
	return s
}

func (s *ListNodeGroupsRequest) SetFuzzyName(v string) *ListNodeGroupsRequest {
	s.FuzzyName = &v
	return s
}

func (s *ListNodeGroupsRequest) SetIotInstanceId(v string) *ListNodeGroupsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListNodeGroupsRequest) SetLimit(v int64) *ListNodeGroupsRequest {
	s.Limit = &v
	return s
}

func (s *ListNodeGroupsRequest) SetOffset(v int64) *ListNodeGroupsRequest {
	s.Offset = &v
	return s
}

func (s *ListNodeGroupsRequest) SetSortingField(v string) *ListNodeGroupsRequest {
	s.SortingField = &v
	return s
}

type ListNodeGroupsResponseBody struct {
	Data      *ListNodeGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBody) SetData(v *ListNodeGroupsResponseBodyData) *ListNodeGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeGroupsResponseBody) SetRequestId(v string) *ListNodeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeGroupsResponseBody) SetSuccess(v bool) *ListNodeGroupsResponseBody {
	s.Success = &v
	return s
}

type ListNodeGroupsResponseBodyData struct {
	List       []*ListNodeGroupsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBodyData) SetList(v []*ListNodeGroupsResponseBodyDataList) *ListNodeGroupsResponseBodyData {
	s.List = v
	return s
}

func (s *ListNodeGroupsResponseBodyData) SetTotalCount(v int64) *ListNodeGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListNodeGroupsResponseBodyDataList struct {
	ClassMode                   *string                                               `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	CreateMillis                *int64                                                `json:"CreateMillis,omitempty" xml:"CreateMillis,omitempty"`
	DataDispatchConfig          *ListNodeGroupsResponseBodyDataListDataDispatchConfig `json:"DataDispatchConfig,omitempty" xml:"DataDispatchConfig,omitempty" type:"Struct"`
	DataDispatchEnabled         *bool                                                 `json:"DataDispatchEnabled,omitempty" xml:"DataDispatchEnabled,omitempty"`
	FreqBandPlanGroupId         *int64                                                `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	JoinEui                     *string                                               `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionEnabled       *bool                                                 `json:"JoinPermissionEnabled,omitempty" xml:"JoinPermissionEnabled,omitempty"`
	JoinPermissionId            *string                                               `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName          *string                                               `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	JoinPermissionOwnerAliyunId *string                                               `json:"JoinPermissionOwnerAliyunId,omitempty" xml:"JoinPermissionOwnerAliyunId,omitempty"`
	JoinPermissionType          *string                                               `json:"JoinPermissionType,omitempty" xml:"JoinPermissionType,omitempty"`
	Locks                       []*ListNodeGroupsResponseBodyDataListLocks            `json:"Locks,omitempty" xml:"Locks,omitempty" type:"Repeated"`
	NodeGroupId                 *string                                               `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName               *string                                               `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodesCnt                    *int64                                                `json:"NodesCnt,omitempty" xml:"NodesCnt,omitempty"`
	RxDailySum                  *int64                                                `json:"RxDailySum,omitempty" xml:"RxDailySum,omitempty"`
	RxMonthSum                  *int64                                                `json:"RxMonthSum,omitempty" xml:"RxMonthSum,omitempty"`
	TxDailySum                  *int64                                                `json:"TxDailySum,omitempty" xml:"TxDailySum,omitempty"`
	TxMonthSum                  *int64                                                `json:"TxMonthSum,omitempty" xml:"TxMonthSum,omitempty"`
}

func (s ListNodeGroupsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBodyDataList) SetClassMode(v string) *ListNodeGroupsResponseBodyDataList {
	s.ClassMode = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetCreateMillis(v int64) *ListNodeGroupsResponseBodyDataList {
	s.CreateMillis = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetDataDispatchConfig(v *ListNodeGroupsResponseBodyDataListDataDispatchConfig) *ListNodeGroupsResponseBodyDataList {
	s.DataDispatchConfig = v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetDataDispatchEnabled(v bool) *ListNodeGroupsResponseBodyDataList {
	s.DataDispatchEnabled = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetFreqBandPlanGroupId(v int64) *ListNodeGroupsResponseBodyDataList {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetJoinEui(v string) *ListNodeGroupsResponseBodyDataList {
	s.JoinEui = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetJoinPermissionEnabled(v bool) *ListNodeGroupsResponseBodyDataList {
	s.JoinPermissionEnabled = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetJoinPermissionId(v string) *ListNodeGroupsResponseBodyDataList {
	s.JoinPermissionId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetJoinPermissionName(v string) *ListNodeGroupsResponseBodyDataList {
	s.JoinPermissionName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetJoinPermissionOwnerAliyunId(v string) *ListNodeGroupsResponseBodyDataList {
	s.JoinPermissionOwnerAliyunId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetJoinPermissionType(v string) *ListNodeGroupsResponseBodyDataList {
	s.JoinPermissionType = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetLocks(v []*ListNodeGroupsResponseBodyDataListLocks) *ListNodeGroupsResponseBodyDataList {
	s.Locks = v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetNodeGroupId(v string) *ListNodeGroupsResponseBodyDataList {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetNodeGroupName(v string) *ListNodeGroupsResponseBodyDataList {
	s.NodeGroupName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetNodesCnt(v int64) *ListNodeGroupsResponseBodyDataList {
	s.NodesCnt = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetRxDailySum(v int64) *ListNodeGroupsResponseBodyDataList {
	s.RxDailySum = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetRxMonthSum(v int64) *ListNodeGroupsResponseBodyDataList {
	s.RxMonthSum = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetTxDailySum(v int64) *ListNodeGroupsResponseBodyDataList {
	s.TxDailySum = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataList) SetTxMonthSum(v int64) *ListNodeGroupsResponseBodyDataList {
	s.TxMonthSum = &v
	return s
}

type ListNodeGroupsResponseBodyDataListDataDispatchConfig struct {
	Destination *string                                                         `json:"Destination,omitempty" xml:"Destination,omitempty"`
	IotProduct  *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct `json:"IotProduct,omitempty" xml:"IotProduct,omitempty" type:"Struct"`
	OnsTopics   *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics  `json:"OnsTopics,omitempty" xml:"OnsTopics,omitempty" type:"Struct"`
}

func (s ListNodeGroupsResponseBodyDataListDataDispatchConfig) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBodyDataListDataDispatchConfig) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfig) SetDestination(v string) *ListNodeGroupsResponseBodyDataListDataDispatchConfig {
	s.Destination = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfig) SetIotProduct(v *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct) *ListNodeGroupsResponseBodyDataListDataDispatchConfig {
	s.IotProduct = v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfig) SetOnsTopics(v *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics) *ListNodeGroupsResponseBodyDataListDataDispatchConfig {
	s.OnsTopics = v
	return s
}

type ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct struct {
	DebugSwitch *bool   `json:"DebugSwitch,omitempty" xml:"DebugSwitch,omitempty"`
	ProductKey  *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct) SetDebugSwitch(v bool) *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct {
	s.DebugSwitch = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct) SetProductKey(v string) *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct {
	s.ProductKey = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct) SetProductName(v string) *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct {
	s.ProductName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct) SetProductType(v string) *ListNodeGroupsResponseBodyDataListDataDispatchConfigIotProduct {
	s.ProductType = &v
	return s
}

type ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics struct {
	DownlinkRegionName *string `json:"DownlinkRegionName,omitempty" xml:"DownlinkRegionName,omitempty"`
	DownlinkTopic      *string `json:"DownlinkTopic,omitempty" xml:"DownlinkTopic,omitempty"`
	UplinkRegionName   *string `json:"UplinkRegionName,omitempty" xml:"UplinkRegionName,omitempty"`
	UplinkTopic        *string `json:"UplinkTopic,omitempty" xml:"UplinkTopic,omitempty"`
}

func (s ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics) SetDownlinkRegionName(v string) *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics {
	s.DownlinkRegionName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics) SetDownlinkTopic(v string) *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics {
	s.DownlinkTopic = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics) SetUplinkRegionName(v string) *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics {
	s.UplinkRegionName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics) SetUplinkTopic(v string) *ListNodeGroupsResponseBodyDataListDataDispatchConfigOnsTopics {
	s.UplinkTopic = &v
	return s
}

type ListNodeGroupsResponseBodyDataListLocks struct {
	CreateMillis *int64  `json:"CreateMillis,omitempty" xml:"CreateMillis,omitempty"`
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	LockId       *string `json:"LockId,omitempty" xml:"LockId,omitempty"`
	LockType     *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s ListNodeGroupsResponseBodyDataListLocks) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBodyDataListLocks) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBodyDataListLocks) SetCreateMillis(v int64) *ListNodeGroupsResponseBodyDataListLocks {
	s.CreateMillis = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListLocks) SetEnabled(v bool) *ListNodeGroupsResponseBodyDataListLocks {
	s.Enabled = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListLocks) SetLockId(v string) *ListNodeGroupsResponseBodyDataListLocks {
	s.LockId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyDataListLocks) SetLockType(v string) *ListNodeGroupsResponseBodyDataListLocks {
	s.LockType = &v
	return s
}

type ListNodeGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponse) SetHeaders(v map[string]*string) *ListNodeGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupsResponse) SetStatusCode(v int32) *ListNodeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupsResponse) SetBody(v *ListNodeGroupsResponseBody) *ListNodeGroupsResponse {
	s.Body = v
	return s
}

type ListNodeTransferPacketPathsRequest struct {
	Base64EncodedMacPayload *string `json:"Base64EncodedMacPayload,omitempty" xml:"Base64EncodedMacPayload,omitempty"`
	DevEui                  *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	IotInstanceId           *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	LogMillis               *int64  `json:"LogMillis,omitempty" xml:"LogMillis,omitempty"`
	PageNumber              *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNodeTransferPacketPathsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketPathsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketPathsRequest) SetBase64EncodedMacPayload(v string) *ListNodeTransferPacketPathsRequest {
	s.Base64EncodedMacPayload = &v
	return s
}

func (s *ListNodeTransferPacketPathsRequest) SetDevEui(v string) *ListNodeTransferPacketPathsRequest {
	s.DevEui = &v
	return s
}

func (s *ListNodeTransferPacketPathsRequest) SetIotInstanceId(v string) *ListNodeTransferPacketPathsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListNodeTransferPacketPathsRequest) SetLogMillis(v int64) *ListNodeTransferPacketPathsRequest {
	s.LogMillis = &v
	return s
}

func (s *ListNodeTransferPacketPathsRequest) SetPageNumber(v int32) *ListNodeTransferPacketPathsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeTransferPacketPathsRequest) SetPageSize(v int32) *ListNodeTransferPacketPathsRequest {
	s.PageSize = &v
	return s
}

type ListNodeTransferPacketPathsResponseBody struct {
	Data      *ListNodeTransferPacketPathsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeTransferPacketPathsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketPathsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketPathsResponseBody) SetData(v *ListNodeTransferPacketPathsResponseBodyData) *ListNodeTransferPacketPathsResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeTransferPacketPathsResponseBody) SetRequestId(v string) *ListNodeTransferPacketPathsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeTransferPacketPathsResponseBody) SetSuccess(v bool) *ListNodeTransferPacketPathsResponseBody {
	s.Success = &v
	return s
}

type ListNodeTransferPacketPathsResponseBodyData struct {
	List       []*ListNodeTransferPacketPathsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeTransferPacketPathsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketPathsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketPathsResponseBodyData) SetList(v []*ListNodeTransferPacketPathsResponseBodyDataList) *ListNodeTransferPacketPathsResponseBodyData {
	s.List = v
	return s
}

func (s *ListNodeTransferPacketPathsResponseBodyData) SetTotalCount(v int64) *ListNodeTransferPacketPathsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListNodeTransferPacketPathsResponseBodyDataList struct {
	BestPath *bool    `json:"BestPath,omitempty" xml:"BestPath,omitempty"`
	DevEui   *string  `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	GwEui    *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	Lsnr     *float32 `json:"Lsnr,omitempty" xml:"Lsnr,omitempty"`
	Rssi     *int32   `json:"Rssi,omitempty" xml:"Rssi,omitempty"`
}

func (s ListNodeTransferPacketPathsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketPathsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketPathsResponseBodyDataList) SetBestPath(v bool) *ListNodeTransferPacketPathsResponseBodyDataList {
	s.BestPath = &v
	return s
}

func (s *ListNodeTransferPacketPathsResponseBodyDataList) SetDevEui(v string) *ListNodeTransferPacketPathsResponseBodyDataList {
	s.DevEui = &v
	return s
}

func (s *ListNodeTransferPacketPathsResponseBodyDataList) SetGwEui(v string) *ListNodeTransferPacketPathsResponseBodyDataList {
	s.GwEui = &v
	return s
}

func (s *ListNodeTransferPacketPathsResponseBodyDataList) SetLsnr(v float32) *ListNodeTransferPacketPathsResponseBodyDataList {
	s.Lsnr = &v
	return s
}

func (s *ListNodeTransferPacketPathsResponseBodyDataList) SetRssi(v int32) *ListNodeTransferPacketPathsResponseBodyDataList {
	s.Rssi = &v
	return s
}

type ListNodeTransferPacketPathsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeTransferPacketPathsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeTransferPacketPathsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketPathsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketPathsResponse) SetHeaders(v map[string]*string) *ListNodeTransferPacketPathsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeTransferPacketPathsResponse) SetStatusCode(v int32) *ListNodeTransferPacketPathsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeTransferPacketPathsResponse) SetBody(v *ListNodeTransferPacketPathsResponseBody) *ListNodeTransferPacketPathsResponse {
	s.Body = v
	return s
}

type ListNodeTransferPacketsRequest struct {
	Ascending    *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	BeginMillis  *int64  `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DevEui       *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	EndMillis    *int64  `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	GwEui        *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortingField *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListNodeTransferPacketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketsRequest) SetAscending(v bool) *ListNodeTransferPacketsRequest {
	s.Ascending = &v
	return s
}

func (s *ListNodeTransferPacketsRequest) SetBeginMillis(v int64) *ListNodeTransferPacketsRequest {
	s.BeginMillis = &v
	return s
}

func (s *ListNodeTransferPacketsRequest) SetCategory(v string) *ListNodeTransferPacketsRequest {
	s.Category = &v
	return s
}

func (s *ListNodeTransferPacketsRequest) SetDevEui(v string) *ListNodeTransferPacketsRequest {
	s.DevEui = &v
	return s
}

func (s *ListNodeTransferPacketsRequest) SetEndMillis(v int64) *ListNodeTransferPacketsRequest {
	s.EndMillis = &v
	return s
}

func (s *ListNodeTransferPacketsRequest) SetGwEui(v string) *ListNodeTransferPacketsRequest {
	s.GwEui = &v
	return s
}

func (s *ListNodeTransferPacketsRequest) SetPageNumber(v int32) *ListNodeTransferPacketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeTransferPacketsRequest) SetPageSize(v int32) *ListNodeTransferPacketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeTransferPacketsRequest) SetSortingField(v string) *ListNodeTransferPacketsRequest {
	s.SortingField = &v
	return s
}

type ListNodeTransferPacketsResponseBody struct {
	Data      *ListNodeTransferPacketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeTransferPacketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketsResponseBody) SetData(v *ListNodeTransferPacketsResponseBodyData) *ListNodeTransferPacketsResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeTransferPacketsResponseBody) SetRequestId(v string) *ListNodeTransferPacketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeTransferPacketsResponseBody) SetSuccess(v bool) *ListNodeTransferPacketsResponseBody {
	s.Success = &v
	return s
}

type ListNodeTransferPacketsResponseBodyData struct {
	List       []*ListNodeTransferPacketsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeTransferPacketsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketsResponseBodyData) SetList(v []*ListNodeTransferPacketsResponseBodyDataList) *ListNodeTransferPacketsResponseBodyData {
	s.List = v
	return s
}

func (s *ListNodeTransferPacketsResponseBodyData) SetTotalCount(v int64) *ListNodeTransferPacketsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListNodeTransferPacketsResponseBodyDataList struct {
	ClassMode *string  `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	Datr      *string  `json:"Datr,omitempty" xml:"Datr,omitempty"`
	DevEui    *string  `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	Freq      *float32 `json:"Freq,omitempty" xml:"Freq,omitempty"`
	GwEui     *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	LogMillis *int64   `json:"LogMillis,omitempty" xml:"LogMillis,omitempty"`
	Rssi      *int32   `json:"Rssi,omitempty" xml:"Rssi,omitempty"`
	Snr       *float32 `json:"Snr,omitempty" xml:"Snr,omitempty"`
}

func (s ListNodeTransferPacketsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketsResponseBodyDataList) SetClassMode(v string) *ListNodeTransferPacketsResponseBodyDataList {
	s.ClassMode = &v
	return s
}

func (s *ListNodeTransferPacketsResponseBodyDataList) SetDatr(v string) *ListNodeTransferPacketsResponseBodyDataList {
	s.Datr = &v
	return s
}

func (s *ListNodeTransferPacketsResponseBodyDataList) SetDevEui(v string) *ListNodeTransferPacketsResponseBodyDataList {
	s.DevEui = &v
	return s
}

func (s *ListNodeTransferPacketsResponseBodyDataList) SetFreq(v float32) *ListNodeTransferPacketsResponseBodyDataList {
	s.Freq = &v
	return s
}

func (s *ListNodeTransferPacketsResponseBodyDataList) SetGwEui(v string) *ListNodeTransferPacketsResponseBodyDataList {
	s.GwEui = &v
	return s
}

func (s *ListNodeTransferPacketsResponseBodyDataList) SetLogMillis(v int64) *ListNodeTransferPacketsResponseBodyDataList {
	s.LogMillis = &v
	return s
}

func (s *ListNodeTransferPacketsResponseBodyDataList) SetRssi(v int32) *ListNodeTransferPacketsResponseBodyDataList {
	s.Rssi = &v
	return s
}

func (s *ListNodeTransferPacketsResponseBodyDataList) SetSnr(v float32) *ListNodeTransferPacketsResponseBodyDataList {
	s.Snr = &v
	return s
}

type ListNodeTransferPacketsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeTransferPacketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeTransferPacketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTransferPacketsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeTransferPacketsResponse) SetHeaders(v map[string]*string) *ListNodeTransferPacketsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeTransferPacketsResponse) SetStatusCode(v int32) *ListNodeTransferPacketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeTransferPacketsResponse) SetBody(v *ListNodeTransferPacketsResponseBody) *ListNodeTransferPacketsResponse {
	s.Body = v
	return s
}

type ListNodeTupleOrdersRequest struct {
	Ascending    *bool     `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	IsKpm        *bool     `json:"IsKpm,omitempty" xml:"IsKpm,omitempty"`
	Limit        *int64    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset       *int64    `json:"Offset,omitempty" xml:"Offset,omitempty"`
	SortingField *string   `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
	State        []*string `json:"State,omitempty" xml:"State,omitempty" type:"Repeated"`
}

func (s ListNodeTupleOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTupleOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListNodeTupleOrdersRequest) SetAscending(v bool) *ListNodeTupleOrdersRequest {
	s.Ascending = &v
	return s
}

func (s *ListNodeTupleOrdersRequest) SetIsKpm(v bool) *ListNodeTupleOrdersRequest {
	s.IsKpm = &v
	return s
}

func (s *ListNodeTupleOrdersRequest) SetLimit(v int64) *ListNodeTupleOrdersRequest {
	s.Limit = &v
	return s
}

func (s *ListNodeTupleOrdersRequest) SetOffset(v int64) *ListNodeTupleOrdersRequest {
	s.Offset = &v
	return s
}

func (s *ListNodeTupleOrdersRequest) SetSortingField(v string) *ListNodeTupleOrdersRequest {
	s.SortingField = &v
	return s
}

func (s *ListNodeTupleOrdersRequest) SetState(v []*string) *ListNodeTupleOrdersRequest {
	s.State = v
	return s
}

type ListNodeTupleOrdersResponseBody struct {
	Data      *ListNodeTupleOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeTupleOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTupleOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeTupleOrdersResponseBody) SetData(v *ListNodeTupleOrdersResponseBodyData) *ListNodeTupleOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeTupleOrdersResponseBody) SetRequestId(v string) *ListNodeTupleOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBody) SetSuccess(v bool) *ListNodeTupleOrdersResponseBody {
	s.Success = &v
	return s
}

type ListNodeTupleOrdersResponseBodyData struct {
	List       []*ListNodeTupleOrdersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeTupleOrdersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTupleOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeTupleOrdersResponseBodyData) SetList(v []*ListNodeTupleOrdersResponseBodyDataList) *ListNodeTupleOrdersResponseBodyData {
	s.List = v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyData) SetTotalCount(v int64) *ListNodeTupleOrdersResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListNodeTupleOrdersResponseBodyDataList struct {
	AcceptedMillis *int64  `json:"AcceptedMillis,omitempty" xml:"AcceptedMillis,omitempty"`
	CreatedMillis  *int64  `json:"CreatedMillis,omitempty" xml:"CreatedMillis,omitempty"`
	FailedCount    *int64  `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	IsKpm          *bool   `json:"IsKpm,omitempty" xml:"IsKpm,omitempty"`
	LoraVersion    *string `json:"LoraVersion,omitempty" xml:"LoraVersion,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderState     *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	RequiredCount  *int64  `json:"RequiredCount,omitempty" xml:"RequiredCount,omitempty"`
	SuccessCount   *int64  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	TupleType      *string `json:"TupleType,omitempty" xml:"TupleType,omitempty"`
}

func (s ListNodeTupleOrdersResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTupleOrdersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetAcceptedMillis(v int64) *ListNodeTupleOrdersResponseBodyDataList {
	s.AcceptedMillis = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetCreatedMillis(v int64) *ListNodeTupleOrdersResponseBodyDataList {
	s.CreatedMillis = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetFailedCount(v int64) *ListNodeTupleOrdersResponseBodyDataList {
	s.FailedCount = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetIsKpm(v bool) *ListNodeTupleOrdersResponseBodyDataList {
	s.IsKpm = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetLoraVersion(v string) *ListNodeTupleOrdersResponseBodyDataList {
	s.LoraVersion = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetOrderId(v string) *ListNodeTupleOrdersResponseBodyDataList {
	s.OrderId = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetOrderState(v string) *ListNodeTupleOrdersResponseBodyDataList {
	s.OrderState = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetRequiredCount(v int64) *ListNodeTupleOrdersResponseBodyDataList {
	s.RequiredCount = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetSuccessCount(v int64) *ListNodeTupleOrdersResponseBodyDataList {
	s.SuccessCount = &v
	return s
}

func (s *ListNodeTupleOrdersResponseBodyDataList) SetTupleType(v string) *ListNodeTupleOrdersResponseBodyDataList {
	s.TupleType = &v
	return s
}

type ListNodeTupleOrdersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeTupleOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeTupleOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeTupleOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListNodeTupleOrdersResponse) SetHeaders(v map[string]*string) *ListNodeTupleOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListNodeTupleOrdersResponse) SetStatusCode(v int32) *ListNodeTupleOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeTupleOrdersResponse) SetBody(v *ListNodeTupleOrdersResponseBody) *ListNodeTupleOrdersResponse {
	s.Body = v
	return s
}

type ListNodesByNodeGroupIdRequest struct {
	Ascending     *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	FuzzyDevEui   *string `json:"FuzzyDevEui,omitempty" xml:"FuzzyDevEui,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Limit         *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NodeGroupId   *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	Offset        *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	SortingField  *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListNodesByNodeGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByNodeGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListNodesByNodeGroupIdRequest) SetAscending(v bool) *ListNodesByNodeGroupIdRequest {
	s.Ascending = &v
	return s
}

func (s *ListNodesByNodeGroupIdRequest) SetFuzzyDevEui(v string) *ListNodesByNodeGroupIdRequest {
	s.FuzzyDevEui = &v
	return s
}

func (s *ListNodesByNodeGroupIdRequest) SetIotInstanceId(v string) *ListNodesByNodeGroupIdRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListNodesByNodeGroupIdRequest) SetLimit(v int64) *ListNodesByNodeGroupIdRequest {
	s.Limit = &v
	return s
}

func (s *ListNodesByNodeGroupIdRequest) SetNodeGroupId(v string) *ListNodesByNodeGroupIdRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodesByNodeGroupIdRequest) SetOffset(v int64) *ListNodesByNodeGroupIdRequest {
	s.Offset = &v
	return s
}

func (s *ListNodesByNodeGroupIdRequest) SetSortingField(v string) *ListNodesByNodeGroupIdRequest {
	s.SortingField = &v
	return s
}

type ListNodesByNodeGroupIdResponseBody struct {
	Data      *ListNodesByNodeGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodesByNodeGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByNodeGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesByNodeGroupIdResponseBody) SetData(v *ListNodesByNodeGroupIdResponseBodyData) *ListNodesByNodeGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBody) SetRequestId(v string) *ListNodesByNodeGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBody) SetSuccess(v bool) *ListNodesByNodeGroupIdResponseBody {
	s.Success = &v
	return s
}

type ListNodesByNodeGroupIdResponseBodyData struct {
	List       []*ListNodesByNodeGroupIdResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesByNodeGroupIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByNodeGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodesByNodeGroupIdResponseBodyData) SetList(v []*ListNodesByNodeGroupIdResponseBodyDataList) *ListNodesByNodeGroupIdResponseBodyData {
	s.List = v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyData) SetTotalCount(v int64) *ListNodesByNodeGroupIdResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListNodesByNodeGroupIdResponseBodyDataList struct {
	Appkey           *string `json:"Appkey,omitempty" xml:"Appkey,omitempty"`
	AuthTypes        *string `json:"AuthTypes,omitempty" xml:"AuthTypes,omitempty"`
	BoundMillis      *int64  `json:"BoundMillis,omitempty" xml:"BoundMillis,omitempty"`
	ClassMode        *string `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	DevAddr          *string `json:"DevAddr,omitempty" xml:"DevAddr,omitempty"`
	DevEui           *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	JoinEui          *string `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	LastJoinMillis   *int64  `json:"LastJoinMillis,omitempty" xml:"LastJoinMillis,omitempty"`
	MulticastGroupId *string `json:"MulticastGroupId,omitempty" xml:"MulticastGroupId,omitempty"`
	NodeType         *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s ListNodesByNodeGroupIdResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByNodeGroupIdResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetAppkey(v string) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.Appkey = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetAuthTypes(v string) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.AuthTypes = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetBoundMillis(v int64) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.BoundMillis = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetClassMode(v string) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.ClassMode = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetDevAddr(v string) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.DevAddr = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetDevEui(v string) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.DevEui = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetJoinEui(v string) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.JoinEui = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetLastJoinMillis(v int64) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.LastJoinMillis = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetMulticastGroupId(v string) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.MulticastGroupId = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponseBodyDataList) SetNodeType(v string) *ListNodesByNodeGroupIdResponseBodyDataList {
	s.NodeType = &v
	return s
}

type ListNodesByNodeGroupIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodesByNodeGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesByNodeGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByNodeGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListNodesByNodeGroupIdResponse) SetHeaders(v map[string]*string) *ListNodesByNodeGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListNodesByNodeGroupIdResponse) SetStatusCode(v int32) *ListNodesByNodeGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesByNodeGroupIdResponse) SetBody(v *ListNodesByNodeGroupIdResponseBody) *ListNodesByNodeGroupIdResponse {
	s.Body = v
	return s
}

type ListNodesByOwnedJoinPermissionIdRequest struct {
	Ascending        *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	FuzzyDevEui      *string `json:"FuzzyDevEui,omitempty" xml:"FuzzyDevEui,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	Limit            *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset           *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	SortingField     *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListNodesByOwnedJoinPermissionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByOwnedJoinPermissionIdRequest) GoString() string {
	return s.String()
}

func (s *ListNodesByOwnedJoinPermissionIdRequest) SetAscending(v bool) *ListNodesByOwnedJoinPermissionIdRequest {
	s.Ascending = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdRequest) SetFuzzyDevEui(v string) *ListNodesByOwnedJoinPermissionIdRequest {
	s.FuzzyDevEui = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdRequest) SetIotInstanceId(v string) *ListNodesByOwnedJoinPermissionIdRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdRequest) SetJoinPermissionId(v string) *ListNodesByOwnedJoinPermissionIdRequest {
	s.JoinPermissionId = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdRequest) SetLimit(v int64) *ListNodesByOwnedJoinPermissionIdRequest {
	s.Limit = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdRequest) SetOffset(v int64) *ListNodesByOwnedJoinPermissionIdRequest {
	s.Offset = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdRequest) SetSortingField(v string) *ListNodesByOwnedJoinPermissionIdRequest {
	s.SortingField = &v
	return s
}

type ListNodesByOwnedJoinPermissionIdResponseBody struct {
	Data      *ListNodesByOwnedJoinPermissionIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodesByOwnedJoinPermissionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByOwnedJoinPermissionIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBody) SetData(v *ListNodesByOwnedJoinPermissionIdResponseBodyData) *ListNodesByOwnedJoinPermissionIdResponseBody {
	s.Data = v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBody) SetRequestId(v string) *ListNodesByOwnedJoinPermissionIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBody) SetSuccess(v bool) *ListNodesByOwnedJoinPermissionIdResponseBody {
	s.Success = &v
	return s
}

type ListNodesByOwnedJoinPermissionIdResponseBodyData struct {
	List       []*ListNodesByOwnedJoinPermissionIdResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesByOwnedJoinPermissionIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByOwnedJoinPermissionIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBodyData) SetList(v []*ListNodesByOwnedJoinPermissionIdResponseBodyDataList) *ListNodesByOwnedJoinPermissionIdResponseBodyData {
	s.List = v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBodyData) SetTotalCount(v int64) *ListNodesByOwnedJoinPermissionIdResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListNodesByOwnedJoinPermissionIdResponseBodyDataList struct {
	BoundMillis    *int64  `json:"BoundMillis,omitempty" xml:"BoundMillis,omitempty"`
	ClassMode      *string `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	DevAddr        *string `json:"DevAddr,omitempty" xml:"DevAddr,omitempty"`
	DevEui         *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	LastJoinMillis *int64  `json:"LastJoinMillis,omitempty" xml:"LastJoinMillis,omitempty"`
}

func (s ListNodesByOwnedJoinPermissionIdResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByOwnedJoinPermissionIdResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBodyDataList) SetBoundMillis(v int64) *ListNodesByOwnedJoinPermissionIdResponseBodyDataList {
	s.BoundMillis = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBodyDataList) SetClassMode(v string) *ListNodesByOwnedJoinPermissionIdResponseBodyDataList {
	s.ClassMode = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBodyDataList) SetDevAddr(v string) *ListNodesByOwnedJoinPermissionIdResponseBodyDataList {
	s.DevAddr = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBodyDataList) SetDevEui(v string) *ListNodesByOwnedJoinPermissionIdResponseBodyDataList {
	s.DevEui = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdResponseBodyDataList) SetLastJoinMillis(v int64) *ListNodesByOwnedJoinPermissionIdResponseBodyDataList {
	s.LastJoinMillis = &v
	return s
}

type ListNodesByOwnedJoinPermissionIdResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodesByOwnedJoinPermissionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesByOwnedJoinPermissionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByOwnedJoinPermissionIdResponse) GoString() string {
	return s.String()
}

func (s *ListNodesByOwnedJoinPermissionIdResponse) SetHeaders(v map[string]*string) *ListNodesByOwnedJoinPermissionIdResponse {
	s.Headers = v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdResponse) SetStatusCode(v int32) *ListNodesByOwnedJoinPermissionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesByOwnedJoinPermissionIdResponse) SetBody(v *ListNodesByOwnedJoinPermissionIdResponseBody) *ListNodesByOwnedJoinPermissionIdResponse {
	s.Body = v
	return s
}

type ListNotificationsRequest struct {
	Ascending    *bool     `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	BeginMillis  *int64    `json:"BeginMillis,omitempty" xml:"BeginMillis,omitempty"`
	Category     []*string `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
	EndMillis    *int64    `json:"EndMillis,omitempty" xml:"EndMillis,omitempty"`
	HandleState  *string   `json:"HandleState,omitempty" xml:"HandleState,omitempty"`
	Limit        *int64    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset       *int64    `json:"Offset,omitempty" xml:"Offset,omitempty"`
	SortingField *string   `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListNotificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationsRequest) GoString() string {
	return s.String()
}

func (s *ListNotificationsRequest) SetAscending(v bool) *ListNotificationsRequest {
	s.Ascending = &v
	return s
}

func (s *ListNotificationsRequest) SetBeginMillis(v int64) *ListNotificationsRequest {
	s.BeginMillis = &v
	return s
}

func (s *ListNotificationsRequest) SetCategory(v []*string) *ListNotificationsRequest {
	s.Category = v
	return s
}

func (s *ListNotificationsRequest) SetEndMillis(v int64) *ListNotificationsRequest {
	s.EndMillis = &v
	return s
}

func (s *ListNotificationsRequest) SetHandleState(v string) *ListNotificationsRequest {
	s.HandleState = &v
	return s
}

func (s *ListNotificationsRequest) SetLimit(v int64) *ListNotificationsRequest {
	s.Limit = &v
	return s
}

func (s *ListNotificationsRequest) SetOffset(v int64) *ListNotificationsRequest {
	s.Offset = &v
	return s
}

func (s *ListNotificationsRequest) SetSortingField(v string) *ListNotificationsRequest {
	s.SortingField = &v
	return s
}

type ListNotificationsResponseBody struct {
	Data      *ListNotificationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNotificationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNotificationsResponseBody) SetData(v *ListNotificationsResponseBodyData) *ListNotificationsResponseBody {
	s.Data = v
	return s
}

func (s *ListNotificationsResponseBody) SetRequestId(v string) *ListNotificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNotificationsResponseBody) SetSuccess(v bool) *ListNotificationsResponseBody {
	s.Success = &v
	return s
}

type ListNotificationsResponseBodyData struct {
	List       []*ListNotificationsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNotificationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNotificationsResponseBodyData) SetList(v []*ListNotificationsResponseBodyDataList) *ListNotificationsResponseBodyData {
	s.List = v
	return s
}

func (s *ListNotificationsResponseBodyData) SetTotalCount(v int64) *ListNotificationsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListNotificationsResponseBodyDataList struct {
	Category               *string                                                      `json:"Category,omitempty" xml:"Category,omitempty"`
	GatewayDataflowLimit   *ListNotificationsResponseBodyDataListGatewayDataflowLimit   `json:"GatewayDataflowLimit,omitempty" xml:"GatewayDataflowLimit,omitempty" type:"Struct"`
	GatewayOfflineInfo     *ListNotificationsResponseBodyDataListGatewayOfflineInfo     `json:"GatewayOfflineInfo,omitempty" xml:"GatewayOfflineInfo,omitempty" type:"Struct"`
	HandleState            *string                                                      `json:"HandleState,omitempty" xml:"HandleState,omitempty"`
	HandledMillis          *int64                                                       `json:"HandledMillis,omitempty" xml:"HandledMillis,omitempty"`
	JoinPermissionAuthInfo *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo `json:"JoinPermissionAuthInfo,omitempty" xml:"JoinPermissionAuthInfo,omitempty" type:"Struct"`
	NoticeMillis           *int64                                                       `json:"NoticeMillis,omitempty" xml:"NoticeMillis,omitempty"`
	NotificationId         *string                                                      `json:"NotificationId,omitempty" xml:"NotificationId,omitempty"`
}

func (s ListNotificationsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNotificationsResponseBodyDataList) SetCategory(v string) *ListNotificationsResponseBodyDataList {
	s.Category = &v
	return s
}

func (s *ListNotificationsResponseBodyDataList) SetGatewayDataflowLimit(v *ListNotificationsResponseBodyDataListGatewayDataflowLimit) *ListNotificationsResponseBodyDataList {
	s.GatewayDataflowLimit = v
	return s
}

func (s *ListNotificationsResponseBodyDataList) SetGatewayOfflineInfo(v *ListNotificationsResponseBodyDataListGatewayOfflineInfo) *ListNotificationsResponseBodyDataList {
	s.GatewayOfflineInfo = v
	return s
}

func (s *ListNotificationsResponseBodyDataList) SetHandleState(v string) *ListNotificationsResponseBodyDataList {
	s.HandleState = &v
	return s
}

func (s *ListNotificationsResponseBodyDataList) SetHandledMillis(v int64) *ListNotificationsResponseBodyDataList {
	s.HandledMillis = &v
	return s
}

func (s *ListNotificationsResponseBodyDataList) SetJoinPermissionAuthInfo(v *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) *ListNotificationsResponseBodyDataList {
	s.JoinPermissionAuthInfo = v
	return s
}

func (s *ListNotificationsResponseBodyDataList) SetNoticeMillis(v int64) *ListNotificationsResponseBodyDataList {
	s.NoticeMillis = &v
	return s
}

func (s *ListNotificationsResponseBodyDataList) SetNotificationId(v string) *ListNotificationsResponseBodyDataList {
	s.NotificationId = &v
	return s
}

type ListNotificationsResponseBodyDataListGatewayDataflowLimit struct {
	AlarmDetail         *string `json:"AlarmDetail,omitempty" xml:"AlarmDetail,omitempty"`
	DataflowLimitMillis *int64  `json:"DataflowLimitMillis,omitempty" xml:"DataflowLimitMillis,omitempty"`
	GwEui               *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
}

func (s ListNotificationsResponseBodyDataListGatewayDataflowLimit) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationsResponseBodyDataListGatewayDataflowLimit) GoString() string {
	return s.String()
}

func (s *ListNotificationsResponseBodyDataListGatewayDataflowLimit) SetAlarmDetail(v string) *ListNotificationsResponseBodyDataListGatewayDataflowLimit {
	s.AlarmDetail = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListGatewayDataflowLimit) SetDataflowLimitMillis(v int64) *ListNotificationsResponseBodyDataListGatewayDataflowLimit {
	s.DataflowLimitMillis = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListGatewayDataflowLimit) SetGwEui(v string) *ListNotificationsResponseBodyDataListGatewayDataflowLimit {
	s.GwEui = &v
	return s
}

type ListNotificationsResponseBodyDataListGatewayOfflineInfo struct {
	GwEui         *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	OfflineMillis *int64  `json:"OfflineMillis,omitempty" xml:"OfflineMillis,omitempty"`
}

func (s ListNotificationsResponseBodyDataListGatewayOfflineInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationsResponseBodyDataListGatewayOfflineInfo) GoString() string {
	return s.String()
}

func (s *ListNotificationsResponseBodyDataListGatewayOfflineInfo) SetGwEui(v string) *ListNotificationsResponseBodyDataListGatewayOfflineInfo {
	s.GwEui = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListGatewayOfflineInfo) SetOfflineMillis(v int64) *ListNotificationsResponseBodyDataListGatewayOfflineInfo {
	s.OfflineMillis = &v
	return s
}

type ListNotificationsResponseBodyDataListJoinPermissionAuthInfo struct {
	AcceptedMillis     *int64  `json:"AcceptedMillis,omitempty" xml:"AcceptedMillis,omitempty"`
	ApplyingMillis     *int64  `json:"ApplyingMillis,omitempty" xml:"ApplyingMillis,omitempty"`
	CanceledMillis     *int64  `json:"CanceledMillis,omitempty" xml:"CanceledMillis,omitempty"`
	JoinEui            *string `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionId   *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName *string `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderState         *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	OwnerAliyunId      *string `json:"OwnerAliyunId,omitempty" xml:"OwnerAliyunId,omitempty"`
	RejectedMillis     *int64  `json:"RejectedMillis,omitempty" xml:"RejectedMillis,omitempty"`
	RenterAliyunId     *string `json:"RenterAliyunId,omitempty" xml:"RenterAliyunId,omitempty"`
}

func (s ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) GoString() string {
	return s.String()
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetAcceptedMillis(v int64) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.AcceptedMillis = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetApplyingMillis(v int64) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.ApplyingMillis = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetCanceledMillis(v int64) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.CanceledMillis = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetJoinEui(v string) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.JoinEui = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetJoinPermissionId(v string) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.JoinPermissionId = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetJoinPermissionName(v string) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.JoinPermissionName = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetOrderId(v string) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.OrderId = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetOrderState(v string) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.OrderState = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetOwnerAliyunId(v string) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.OwnerAliyunId = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetRejectedMillis(v int64) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.RejectedMillis = &v
	return s
}

func (s *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo) SetRenterAliyunId(v string) *ListNotificationsResponseBodyDataListJoinPermissionAuthInfo {
	s.RenterAliyunId = &v
	return s
}

type ListNotificationsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNotificationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNotificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationsResponse) GoString() string {
	return s.String()
}

func (s *ListNotificationsResponse) SetHeaders(v map[string]*string) *ListNotificationsResponse {
	s.Headers = v
	return s
}

func (s *ListNotificationsResponse) SetStatusCode(v int32) *ListNotificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNotificationsResponse) SetBody(v *ListNotificationsResponseBody) *ListNotificationsResponse {
	s.Body = v
	return s
}

type ListOwnedJoinPermissionsRequest struct {
	Ascending               *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	Enabled                 *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FuzzyJoinEui            *string `json:"FuzzyJoinEui,omitempty" xml:"FuzzyJoinEui,omitempty"`
	FuzzyJoinPermissionName *string `json:"FuzzyJoinPermissionName,omitempty" xml:"FuzzyJoinPermissionName,omitempty"`
	FuzzyRenterAliyunId     *string `json:"FuzzyRenterAliyunId,omitempty" xml:"FuzzyRenterAliyunId,omitempty"`
	IotInstanceId           *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Limit                   *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset                  *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	SortingField            *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
}

func (s ListOwnedJoinPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedJoinPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListOwnedJoinPermissionsRequest) SetAscending(v bool) *ListOwnedJoinPermissionsRequest {
	s.Ascending = &v
	return s
}

func (s *ListOwnedJoinPermissionsRequest) SetEnabled(v bool) *ListOwnedJoinPermissionsRequest {
	s.Enabled = &v
	return s
}

func (s *ListOwnedJoinPermissionsRequest) SetFuzzyJoinEui(v string) *ListOwnedJoinPermissionsRequest {
	s.FuzzyJoinEui = &v
	return s
}

func (s *ListOwnedJoinPermissionsRequest) SetFuzzyJoinPermissionName(v string) *ListOwnedJoinPermissionsRequest {
	s.FuzzyJoinPermissionName = &v
	return s
}

func (s *ListOwnedJoinPermissionsRequest) SetFuzzyRenterAliyunId(v string) *ListOwnedJoinPermissionsRequest {
	s.FuzzyRenterAliyunId = &v
	return s
}

func (s *ListOwnedJoinPermissionsRequest) SetIotInstanceId(v string) *ListOwnedJoinPermissionsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListOwnedJoinPermissionsRequest) SetLimit(v int64) *ListOwnedJoinPermissionsRequest {
	s.Limit = &v
	return s
}

func (s *ListOwnedJoinPermissionsRequest) SetOffset(v int64) *ListOwnedJoinPermissionsRequest {
	s.Offset = &v
	return s
}

func (s *ListOwnedJoinPermissionsRequest) SetSortingField(v string) *ListOwnedJoinPermissionsRequest {
	s.SortingField = &v
	return s
}

type ListOwnedJoinPermissionsResponseBody struct {
	Data      *ListOwnedJoinPermissionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOwnedJoinPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedJoinPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOwnedJoinPermissionsResponseBody) SetData(v *ListOwnedJoinPermissionsResponseBodyData) *ListOwnedJoinPermissionsResponseBody {
	s.Data = v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBody) SetRequestId(v string) *ListOwnedJoinPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBody) SetSuccess(v bool) *ListOwnedJoinPermissionsResponseBody {
	s.Success = &v
	return s
}

type ListOwnedJoinPermissionsResponseBodyData struct {
	List       []*ListOwnedJoinPermissionsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOwnedJoinPermissionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedJoinPermissionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOwnedJoinPermissionsResponseBodyData) SetList(v []*ListOwnedJoinPermissionsResponseBodyDataList) *ListOwnedJoinPermissionsResponseBodyData {
	s.List = v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyData) SetTotalCount(v int64) *ListOwnedJoinPermissionsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListOwnedJoinPermissionsResponseBodyDataList struct {
	AuthState           *string `json:"AuthState,omitempty" xml:"AuthState,omitempty"`
	ClassMode           *string `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	DataRate            *int64  `json:"DataRate,omitempty" xml:"DataRate,omitempty"`
	Enabled             *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FreqBandPlanGroupId *int64  `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	JoinEui             *string `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionId    *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName  *string `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	RenterAliyunId      *string `json:"RenterAliyunId,omitempty" xml:"RenterAliyunId,omitempty"`
	RxDelay             *int64  `json:"RxDelay,omitempty" xml:"RxDelay,omitempty"`
}

func (s ListOwnedJoinPermissionsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedJoinPermissionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetAuthState(v string) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.AuthState = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetClassMode(v string) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.ClassMode = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetDataRate(v int64) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.DataRate = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetEnabled(v bool) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.Enabled = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetFreqBandPlanGroupId(v int64) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetJoinEui(v string) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.JoinEui = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetJoinPermissionId(v string) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.JoinPermissionId = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetJoinPermissionName(v string) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.JoinPermissionName = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetRenterAliyunId(v string) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.RenterAliyunId = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponseBodyDataList) SetRxDelay(v int64) *ListOwnedJoinPermissionsResponseBodyDataList {
	s.RxDelay = &v
	return s
}

type ListOwnedJoinPermissionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOwnedJoinPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOwnedJoinPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOwnedJoinPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListOwnedJoinPermissionsResponse) SetHeaders(v map[string]*string) *ListOwnedJoinPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListOwnedJoinPermissionsResponse) SetStatusCode(v int32) *ListOwnedJoinPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOwnedJoinPermissionsResponse) SetBody(v *ListOwnedJoinPermissionsResponseBody) *ListOwnedJoinPermissionsResponse {
	s.Body = v
	return s
}

type ListRentedJoinPermissionsRequest struct {
	Ascending               *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	BoundNodeGroup          *bool   `json:"BoundNodeGroup,omitempty" xml:"BoundNodeGroup,omitempty"`
	Enabled                 *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FuzzyJoinEui            *string `json:"FuzzyJoinEui,omitempty" xml:"FuzzyJoinEui,omitempty"`
	FuzzyJoinPermissionName *string `json:"FuzzyJoinPermissionName,omitempty" xml:"FuzzyJoinPermissionName,omitempty"`
	FuzzyOwnerAliyunId      *string `json:"FuzzyOwnerAliyunId,omitempty" xml:"FuzzyOwnerAliyunId,omitempty"`
	IotInstanceId           *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Limit                   *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset                  *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	SortingField            *string `json:"SortingField,omitempty" xml:"SortingField,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRentedJoinPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRentedJoinPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListRentedJoinPermissionsRequest) SetAscending(v bool) *ListRentedJoinPermissionsRequest {
	s.Ascending = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetBoundNodeGroup(v bool) *ListRentedJoinPermissionsRequest {
	s.BoundNodeGroup = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetEnabled(v bool) *ListRentedJoinPermissionsRequest {
	s.Enabled = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetFuzzyJoinEui(v string) *ListRentedJoinPermissionsRequest {
	s.FuzzyJoinEui = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetFuzzyJoinPermissionName(v string) *ListRentedJoinPermissionsRequest {
	s.FuzzyJoinPermissionName = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetFuzzyOwnerAliyunId(v string) *ListRentedJoinPermissionsRequest {
	s.FuzzyOwnerAliyunId = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetIotInstanceId(v string) *ListRentedJoinPermissionsRequest {
	s.IotInstanceId = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetLimit(v int64) *ListRentedJoinPermissionsRequest {
	s.Limit = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetOffset(v int64) *ListRentedJoinPermissionsRequest {
	s.Offset = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetSortingField(v string) *ListRentedJoinPermissionsRequest {
	s.SortingField = &v
	return s
}

func (s *ListRentedJoinPermissionsRequest) SetType(v string) *ListRentedJoinPermissionsRequest {
	s.Type = &v
	return s
}

type ListRentedJoinPermissionsResponseBody struct {
	Data      *ListRentedJoinPermissionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRentedJoinPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRentedJoinPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRentedJoinPermissionsResponseBody) SetData(v *ListRentedJoinPermissionsResponseBodyData) *ListRentedJoinPermissionsResponseBody {
	s.Data = v
	return s
}

func (s *ListRentedJoinPermissionsResponseBody) SetRequestId(v string) *ListRentedJoinPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBody) SetSuccess(v bool) *ListRentedJoinPermissionsResponseBody {
	s.Success = &v
	return s
}

type ListRentedJoinPermissionsResponseBodyData struct {
	List       []*ListRentedJoinPermissionsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRentedJoinPermissionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRentedJoinPermissionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRentedJoinPermissionsResponseBodyData) SetList(v []*ListRentedJoinPermissionsResponseBodyDataList) *ListRentedJoinPermissionsResponseBodyData {
	s.List = v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyData) SetTotalCount(v int64) *ListRentedJoinPermissionsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListRentedJoinPermissionsResponseBodyDataList struct {
	BoundNodeGroup      *bool                                                            `json:"BoundNodeGroup,omitempty" xml:"BoundNodeGroup,omitempty"`
	BoundNodeGroupId    *string                                                          `json:"BoundNodeGroupId,omitempty" xml:"BoundNodeGroupId,omitempty"`
	BoundNodeGroupName  *string                                                          `json:"BoundNodeGroupName,omitempty" xml:"BoundNodeGroupName,omitempty"`
	ClassMode           *string                                                          `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	DataDispatchConfig  *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig `json:"DataDispatchConfig,omitempty" xml:"DataDispatchConfig,omitempty" type:"Struct"`
	DataRate            *string                                                          `json:"DataRate,omitempty" xml:"DataRate,omitempty"`
	Enabled             *bool                                                            `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	FreqBandPlanGroupId *string                                                          `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	JoinEui             *string                                                          `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionId    *string                                                          `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName  *string                                                          `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	OwnerAliyunId       *string                                                          `json:"OwnerAliyunId,omitempty" xml:"OwnerAliyunId,omitempty"`
	RxDelay             *string                                                          `json:"RxDelay,omitempty" xml:"RxDelay,omitempty"`
	Type                *string                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRentedJoinPermissionsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListRentedJoinPermissionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetBoundNodeGroup(v bool) *ListRentedJoinPermissionsResponseBodyDataList {
	s.BoundNodeGroup = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetBoundNodeGroupId(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.BoundNodeGroupId = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetBoundNodeGroupName(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.BoundNodeGroupName = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetClassMode(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.ClassMode = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetDataDispatchConfig(v *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig) *ListRentedJoinPermissionsResponseBodyDataList {
	s.DataDispatchConfig = v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetDataRate(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.DataRate = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetEnabled(v bool) *ListRentedJoinPermissionsResponseBodyDataList {
	s.Enabled = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetFreqBandPlanGroupId(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetJoinEui(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.JoinEui = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetJoinPermissionId(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.JoinPermissionId = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetJoinPermissionName(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.JoinPermissionName = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetOwnerAliyunId(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.OwnerAliyunId = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetRxDelay(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.RxDelay = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataList) SetType(v string) *ListRentedJoinPermissionsResponseBodyDataList {
	s.Type = &v
	return s
}

type ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig struct {
	Destination *string                                                                    `json:"Destination,omitempty" xml:"Destination,omitempty"`
	IotProduct  *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct `json:"IotProduct,omitempty" xml:"IotProduct,omitempty" type:"Struct"`
	OnsTopics   *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics  `json:"OnsTopics,omitempty" xml:"OnsTopics,omitempty" type:"Struct"`
}

func (s ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig) GoString() string {
	return s.String()
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig) SetDestination(v string) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig {
	s.Destination = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig) SetIotProduct(v *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig {
	s.IotProduct = v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig) SetOnsTopics(v *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfig {
	s.OnsTopics = v
	return s
}

type ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct struct {
	DebugSwitch *bool   `json:"DebugSwitch,omitempty" xml:"DebugSwitch,omitempty"`
	ProductKey  *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct) String() string {
	return tea.Prettify(s)
}

func (s ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct) GoString() string {
	return s.String()
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct) SetDebugSwitch(v bool) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct {
	s.DebugSwitch = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct) SetProductKey(v string) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct {
	s.ProductKey = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct) SetProductName(v string) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct {
	s.ProductName = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct) SetProductType(v string) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigIotProduct {
	s.ProductType = &v
	return s
}

type ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics struct {
	DownlinkRegionName *string `json:"DownlinkRegionName,omitempty" xml:"DownlinkRegionName,omitempty"`
	DownlinkTopic      *string `json:"DownlinkTopic,omitempty" xml:"DownlinkTopic,omitempty"`
	UplinkRegionName   *string `json:"UplinkRegionName,omitempty" xml:"UplinkRegionName,omitempty"`
	UplinkTopic        *string `json:"UplinkTopic,omitempty" xml:"UplinkTopic,omitempty"`
}

func (s ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics) String() string {
	return tea.Prettify(s)
}

func (s ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics) GoString() string {
	return s.String()
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics) SetDownlinkRegionName(v string) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics {
	s.DownlinkRegionName = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics) SetDownlinkTopic(v string) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics {
	s.DownlinkTopic = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics) SetUplinkRegionName(v string) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics {
	s.UplinkRegionName = &v
	return s
}

func (s *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics) SetUplinkTopic(v string) *ListRentedJoinPermissionsResponseBodyDataListDataDispatchConfigOnsTopics {
	s.UplinkTopic = &v
	return s
}

type ListRentedJoinPermissionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRentedJoinPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRentedJoinPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRentedJoinPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListRentedJoinPermissionsResponse) SetHeaders(v map[string]*string) *ListRentedJoinPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListRentedJoinPermissionsResponse) SetStatusCode(v int32) *ListRentedJoinPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRentedJoinPermissionsResponse) SetBody(v *ListRentedJoinPermissionsResponseBody) *ListRentedJoinPermissionsResponse {
	s.Body = v
	return s
}

type RejectJoinPermissionAuthOrderRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RejectJoinPermissionAuthOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectJoinPermissionAuthOrderRequest) GoString() string {
	return s.String()
}

func (s *RejectJoinPermissionAuthOrderRequest) SetOrderId(v string) *RejectJoinPermissionAuthOrderRequest {
	s.OrderId = &v
	return s
}

type RejectJoinPermissionAuthOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RejectJoinPermissionAuthOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectJoinPermissionAuthOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RejectJoinPermissionAuthOrderResponseBody) SetRequestId(v string) *RejectJoinPermissionAuthOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectJoinPermissionAuthOrderResponseBody) SetSuccess(v bool) *RejectJoinPermissionAuthOrderResponseBody {
	s.Success = &v
	return s
}

type RejectJoinPermissionAuthOrderResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RejectJoinPermissionAuthOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RejectJoinPermissionAuthOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectJoinPermissionAuthOrderResponse) GoString() string {
	return s.String()
}

func (s *RejectJoinPermissionAuthOrderResponse) SetHeaders(v map[string]*string) *RejectJoinPermissionAuthOrderResponse {
	s.Headers = v
	return s
}

func (s *RejectJoinPermissionAuthOrderResponse) SetStatusCode(v int32) *RejectJoinPermissionAuthOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectJoinPermissionAuthOrderResponse) SetBody(v *RejectJoinPermissionAuthOrderResponseBody) *RejectJoinPermissionAuthOrderResponse {
	s.Body = v
	return s
}

type RemoveNodeFromGroupRequest struct {
	DevEui      *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s RemoveNodeFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveNodeFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveNodeFromGroupRequest) SetDevEui(v string) *RemoveNodeFromGroupRequest {
	s.DevEui = &v
	return s
}

func (s *RemoveNodeFromGroupRequest) SetNodeGroupId(v string) *RemoveNodeFromGroupRequest {
	s.NodeGroupId = &v
	return s
}

type RemoveNodeFromGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveNodeFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveNodeFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveNodeFromGroupResponseBody) SetRequestId(v string) *RemoveNodeFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveNodeFromGroupResponseBody) SetSuccess(v bool) *RemoveNodeFromGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveNodeFromGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveNodeFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveNodeFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveNodeFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveNodeFromGroupResponse) SetHeaders(v map[string]*string) *RemoveNodeFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveNodeFromGroupResponse) SetStatusCode(v int32) *RemoveNodeFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveNodeFromGroupResponse) SetBody(v *RemoveNodeFromGroupResponseBody) *RemoveNodeFromGroupResponse {
	s.Body = v
	return s
}

type ReturnJoinPermissionRequest struct {
	JoinPermissionId   *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionType *string `json:"JoinPermissionType,omitempty" xml:"JoinPermissionType,omitempty"`
}

func (s ReturnJoinPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReturnJoinPermissionRequest) GoString() string {
	return s.String()
}

func (s *ReturnJoinPermissionRequest) SetJoinPermissionId(v string) *ReturnJoinPermissionRequest {
	s.JoinPermissionId = &v
	return s
}

func (s *ReturnJoinPermissionRequest) SetJoinPermissionType(v string) *ReturnJoinPermissionRequest {
	s.JoinPermissionType = &v
	return s
}

type ReturnJoinPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReturnJoinPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReturnJoinPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ReturnJoinPermissionResponseBody) SetRequestId(v string) *ReturnJoinPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReturnJoinPermissionResponseBody) SetSuccess(v bool) *ReturnJoinPermissionResponseBody {
	s.Success = &v
	return s
}

type ReturnJoinPermissionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReturnJoinPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReturnJoinPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReturnJoinPermissionResponse) GoString() string {
	return s.String()
}

func (s *ReturnJoinPermissionResponse) SetHeaders(v map[string]*string) *ReturnJoinPermissionResponse {
	s.Headers = v
	return s
}

func (s *ReturnJoinPermissionResponse) SetStatusCode(v int32) *ReturnJoinPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReturnJoinPermissionResponse) SetBody(v *ReturnJoinPermissionResponseBody) *ReturnJoinPermissionResponse {
	s.Body = v
	return s
}

type SendUnicastCommandRequest struct {
	CleanUp       *bool   `json:"CleanUp,omitempty" xml:"CleanUp,omitempty"`
	Confirmed     *bool   `json:"Confirmed,omitempty" xml:"Confirmed,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DevEui        *string `json:"DevEui,omitempty" xml:"DevEui,omitempty"`
	FPort         *int32  `json:"FPort,omitempty" xml:"FPort,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	MaxRetries    *int32  `json:"MaxRetries,omitempty" xml:"MaxRetries,omitempty"`
}

func (s SendUnicastCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s SendUnicastCommandRequest) GoString() string {
	return s.String()
}

func (s *SendUnicastCommandRequest) SetCleanUp(v bool) *SendUnicastCommandRequest {
	s.CleanUp = &v
	return s
}

func (s *SendUnicastCommandRequest) SetConfirmed(v bool) *SendUnicastCommandRequest {
	s.Confirmed = &v
	return s
}

func (s *SendUnicastCommandRequest) SetContent(v string) *SendUnicastCommandRequest {
	s.Content = &v
	return s
}

func (s *SendUnicastCommandRequest) SetDevEui(v string) *SendUnicastCommandRequest {
	s.DevEui = &v
	return s
}

func (s *SendUnicastCommandRequest) SetFPort(v int32) *SendUnicastCommandRequest {
	s.FPort = &v
	return s
}

func (s *SendUnicastCommandRequest) SetIotInstanceId(v string) *SendUnicastCommandRequest {
	s.IotInstanceId = &v
	return s
}

func (s *SendUnicastCommandRequest) SetMaxRetries(v int32) *SendUnicastCommandRequest {
	s.MaxRetries = &v
	return s
}

type SendUnicastCommandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendUnicastCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendUnicastCommandResponseBody) GoString() string {
	return s.String()
}

func (s *SendUnicastCommandResponseBody) SetRequestId(v string) *SendUnicastCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendUnicastCommandResponseBody) SetSuccess(v bool) *SendUnicastCommandResponseBody {
	s.Success = &v
	return s
}

type SendUnicastCommandResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendUnicastCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendUnicastCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s SendUnicastCommandResponse) GoString() string {
	return s.String()
}

func (s *SendUnicastCommandResponse) SetHeaders(v map[string]*string) *SendUnicastCommandResponse {
	s.Headers = v
	return s
}

func (s *SendUnicastCommandResponse) SetStatusCode(v int32) *SendUnicastCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *SendUnicastCommandResponse) SetBody(v *SendUnicastCommandResponseBody) *SendUnicastCommandResponse {
	s.Body = v
	return s
}

type SubmitGatewayTupleOrderRequest struct {
	RequiredCount *int64  `json:"RequiredCount,omitempty" xml:"RequiredCount,omitempty"`
	TupleType     *string `json:"TupleType,omitempty" xml:"TupleType,omitempty"`
}

func (s SubmitGatewayTupleOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitGatewayTupleOrderRequest) GoString() string {
	return s.String()
}

func (s *SubmitGatewayTupleOrderRequest) SetRequiredCount(v int64) *SubmitGatewayTupleOrderRequest {
	s.RequiredCount = &v
	return s
}

func (s *SubmitGatewayTupleOrderRequest) SetTupleType(v string) *SubmitGatewayTupleOrderRequest {
	s.TupleType = &v
	return s
}

type SubmitGatewayTupleOrderResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitGatewayTupleOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitGatewayTupleOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitGatewayTupleOrderResponseBody) SetData(v string) *SubmitGatewayTupleOrderResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitGatewayTupleOrderResponseBody) SetRequestId(v string) *SubmitGatewayTupleOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitGatewayTupleOrderResponseBody) SetSuccess(v bool) *SubmitGatewayTupleOrderResponseBody {
	s.Success = &v
	return s
}

type SubmitGatewayTupleOrderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitGatewayTupleOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitGatewayTupleOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitGatewayTupleOrderResponse) GoString() string {
	return s.String()
}

func (s *SubmitGatewayTupleOrderResponse) SetHeaders(v map[string]*string) *SubmitGatewayTupleOrderResponse {
	s.Headers = v
	return s
}

func (s *SubmitGatewayTupleOrderResponse) SetStatusCode(v int32) *SubmitGatewayTupleOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitGatewayTupleOrderResponse) SetBody(v *SubmitGatewayTupleOrderResponseBody) *SubmitGatewayTupleOrderResponse {
	s.Body = v
	return s
}

type SubmitJoinPermissionAuthOrderRequest struct {
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	RenterAliyunId   *string `json:"RenterAliyunId,omitempty" xml:"RenterAliyunId,omitempty"`
}

func (s SubmitJoinPermissionAuthOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitJoinPermissionAuthOrderRequest) GoString() string {
	return s.String()
}

func (s *SubmitJoinPermissionAuthOrderRequest) SetJoinPermissionId(v string) *SubmitJoinPermissionAuthOrderRequest {
	s.JoinPermissionId = &v
	return s
}

func (s *SubmitJoinPermissionAuthOrderRequest) SetRenterAliyunId(v string) *SubmitJoinPermissionAuthOrderRequest {
	s.RenterAliyunId = &v
	return s
}

type SubmitJoinPermissionAuthOrderResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitJoinPermissionAuthOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitJoinPermissionAuthOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitJoinPermissionAuthOrderResponseBody) SetData(v int64) *SubmitJoinPermissionAuthOrderResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitJoinPermissionAuthOrderResponseBody) SetRequestId(v string) *SubmitJoinPermissionAuthOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitJoinPermissionAuthOrderResponseBody) SetSuccess(v bool) *SubmitJoinPermissionAuthOrderResponseBody {
	s.Success = &v
	return s
}

type SubmitJoinPermissionAuthOrderResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitJoinPermissionAuthOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitJoinPermissionAuthOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitJoinPermissionAuthOrderResponse) GoString() string {
	return s.String()
}

func (s *SubmitJoinPermissionAuthOrderResponse) SetHeaders(v map[string]*string) *SubmitJoinPermissionAuthOrderResponse {
	s.Headers = v
	return s
}

func (s *SubmitJoinPermissionAuthOrderResponse) SetStatusCode(v int32) *SubmitJoinPermissionAuthOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitJoinPermissionAuthOrderResponse) SetBody(v *SubmitJoinPermissionAuthOrderResponseBody) *SubmitJoinPermissionAuthOrderResponse {
	s.Body = v
	return s
}

type SubmitNodeTupleOrderRequest struct {
	LoraVersion   *string `json:"LoraVersion,omitempty" xml:"LoraVersion,omitempty"`
	RequiredCount *int64  `json:"RequiredCount,omitempty" xml:"RequiredCount,omitempty"`
	TupleType     *string `json:"TupleType,omitempty" xml:"TupleType,omitempty"`
}

func (s SubmitNodeTupleOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitNodeTupleOrderRequest) GoString() string {
	return s.String()
}

func (s *SubmitNodeTupleOrderRequest) SetLoraVersion(v string) *SubmitNodeTupleOrderRequest {
	s.LoraVersion = &v
	return s
}

func (s *SubmitNodeTupleOrderRequest) SetRequiredCount(v int64) *SubmitNodeTupleOrderRequest {
	s.RequiredCount = &v
	return s
}

func (s *SubmitNodeTupleOrderRequest) SetTupleType(v string) *SubmitNodeTupleOrderRequest {
	s.TupleType = &v
	return s
}

type SubmitNodeTupleOrderResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitNodeTupleOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitNodeTupleOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitNodeTupleOrderResponseBody) SetData(v string) *SubmitNodeTupleOrderResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitNodeTupleOrderResponseBody) SetRequestId(v string) *SubmitNodeTupleOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitNodeTupleOrderResponseBody) SetSuccess(v bool) *SubmitNodeTupleOrderResponseBody {
	s.Success = &v
	return s
}

type SubmitNodeTupleOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitNodeTupleOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitNodeTupleOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitNodeTupleOrderResponse) GoString() string {
	return s.String()
}

func (s *SubmitNodeTupleOrderResponse) SetHeaders(v map[string]*string) *SubmitNodeTupleOrderResponse {
	s.Headers = v
	return s
}

func (s *SubmitNodeTupleOrderResponse) SetStatusCode(v int32) *SubmitNodeTupleOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitNodeTupleOrderResponse) SetBody(v *SubmitNodeTupleOrderResponseBody) *SubmitNodeTupleOrderResponse {
	s.Body = v
	return s
}

type UnbindJoinPermissionFromNodeGroupRequest struct {
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	NodeGroupId      *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s UnbindJoinPermissionFromNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindJoinPermissionFromNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *UnbindJoinPermissionFromNodeGroupRequest) SetJoinPermissionId(v string) *UnbindJoinPermissionFromNodeGroupRequest {
	s.JoinPermissionId = &v
	return s
}

func (s *UnbindJoinPermissionFromNodeGroupRequest) SetNodeGroupId(v string) *UnbindJoinPermissionFromNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

type UnbindJoinPermissionFromNodeGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindJoinPermissionFromNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindJoinPermissionFromNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindJoinPermissionFromNodeGroupResponseBody) SetRequestId(v string) *UnbindJoinPermissionFromNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindJoinPermissionFromNodeGroupResponseBody) SetSuccess(v bool) *UnbindJoinPermissionFromNodeGroupResponseBody {
	s.Success = &v
	return s
}

type UnbindJoinPermissionFromNodeGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindJoinPermissionFromNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindJoinPermissionFromNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindJoinPermissionFromNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *UnbindJoinPermissionFromNodeGroupResponse) SetHeaders(v map[string]*string) *UnbindJoinPermissionFromNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *UnbindJoinPermissionFromNodeGroupResponse) SetStatusCode(v int32) *UnbindJoinPermissionFromNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindJoinPermissionFromNodeGroupResponse) SetBody(v *UnbindJoinPermissionFromNodeGroupResponseBody) *UnbindJoinPermissionFromNodeGroupResponse {
	s.Body = v
	return s
}

type UpdateDataDispatchConfigRequest struct {
	DataDispatchDestination *string `json:"DataDispatchDestination,omitempty" xml:"DataDispatchDestination,omitempty"`
	DebugSwitch             *bool   `json:"DebugSwitch,omitempty" xml:"DebugSwitch,omitempty"`
	NodeGroupId             *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	ProductKey              *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProductName             *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductType             *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	UplinkRegionName        *string `json:"UplinkRegionName,omitempty" xml:"UplinkRegionName,omitempty"`
	UplinkTopic             *string `json:"UplinkTopic,omitempty" xml:"UplinkTopic,omitempty"`
}

func (s UpdateDataDispatchConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataDispatchConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataDispatchConfigRequest) SetDataDispatchDestination(v string) *UpdateDataDispatchConfigRequest {
	s.DataDispatchDestination = &v
	return s
}

func (s *UpdateDataDispatchConfigRequest) SetDebugSwitch(v bool) *UpdateDataDispatchConfigRequest {
	s.DebugSwitch = &v
	return s
}

func (s *UpdateDataDispatchConfigRequest) SetNodeGroupId(v string) *UpdateDataDispatchConfigRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateDataDispatchConfigRequest) SetProductKey(v string) *UpdateDataDispatchConfigRequest {
	s.ProductKey = &v
	return s
}

func (s *UpdateDataDispatchConfigRequest) SetProductName(v string) *UpdateDataDispatchConfigRequest {
	s.ProductName = &v
	return s
}

func (s *UpdateDataDispatchConfigRequest) SetProductType(v string) *UpdateDataDispatchConfigRequest {
	s.ProductType = &v
	return s
}

func (s *UpdateDataDispatchConfigRequest) SetUplinkRegionName(v string) *UpdateDataDispatchConfigRequest {
	s.UplinkRegionName = &v
	return s
}

func (s *UpdateDataDispatchConfigRequest) SetUplinkTopic(v string) *UpdateDataDispatchConfigRequest {
	s.UplinkTopic = &v
	return s
}

type UpdateDataDispatchConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataDispatchConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataDispatchConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataDispatchConfigResponseBody) SetRequestId(v string) *UpdateDataDispatchConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataDispatchConfigResponseBody) SetSuccess(v bool) *UpdateDataDispatchConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateDataDispatchConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDataDispatchConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDataDispatchConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataDispatchConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataDispatchConfigResponse) SetHeaders(v map[string]*string) *UpdateDataDispatchConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataDispatchConfigResponse) SetStatusCode(v int32) *UpdateDataDispatchConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataDispatchConfigResponse) SetBody(v *UpdateDataDispatchConfigResponseBody) *UpdateDataDispatchConfigResponse {
	s.Body = v
	return s
}

type UpdateDataDispatchEnablingStateRequest struct {
	DataDispatchEnabled *bool   `json:"DataDispatchEnabled,omitempty" xml:"DataDispatchEnabled,omitempty"`
	NodeGroupId         *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s UpdateDataDispatchEnablingStateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataDispatchEnablingStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataDispatchEnablingStateRequest) SetDataDispatchEnabled(v bool) *UpdateDataDispatchEnablingStateRequest {
	s.DataDispatchEnabled = &v
	return s
}

func (s *UpdateDataDispatchEnablingStateRequest) SetNodeGroupId(v string) *UpdateDataDispatchEnablingStateRequest {
	s.NodeGroupId = &v
	return s
}

type UpdateDataDispatchEnablingStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataDispatchEnablingStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataDispatchEnablingStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataDispatchEnablingStateResponseBody) SetRequestId(v string) *UpdateDataDispatchEnablingStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataDispatchEnablingStateResponseBody) SetSuccess(v bool) *UpdateDataDispatchEnablingStateResponseBody {
	s.Success = &v
	return s
}

type UpdateDataDispatchEnablingStateResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDataDispatchEnablingStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDataDispatchEnablingStateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataDispatchEnablingStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataDispatchEnablingStateResponse) SetHeaders(v map[string]*string) *UpdateDataDispatchEnablingStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataDispatchEnablingStateResponse) SetStatusCode(v int32) *UpdateDataDispatchEnablingStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataDispatchEnablingStateResponse) SetBody(v *UpdateDataDispatchEnablingStateResponseBody) *UpdateDataDispatchEnablingStateResponse {
	s.Body = v
	return s
}

type UpdateGatewayRequest struct {
	Address             *string  `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressCode         *int64   `json:"AddressCode,omitempty" xml:"AddressCode,omitempty"`
	City                *string  `json:"City,omitempty" xml:"City,omitempty"`
	CommunicationMode   *string  `json:"CommunicationMode,omitempty" xml:"CommunicationMode,omitempty"`
	Description         *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	District            *string  `json:"District,omitempty" xml:"District,omitempty"`
	FreqBandPlanGroupId *int64   `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	GisCoordinateSystem *string  `json:"GisCoordinateSystem,omitempty" xml:"GisCoordinateSystem,omitempty"`
	GwEui               *string  `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId       *string  `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	Latitude            *float32 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude           *float32 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name                *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRequest) SetAddress(v string) *UpdateGatewayRequest {
	s.Address = &v
	return s
}

func (s *UpdateGatewayRequest) SetAddressCode(v int64) *UpdateGatewayRequest {
	s.AddressCode = &v
	return s
}

func (s *UpdateGatewayRequest) SetCity(v string) *UpdateGatewayRequest {
	s.City = &v
	return s
}

func (s *UpdateGatewayRequest) SetCommunicationMode(v string) *UpdateGatewayRequest {
	s.CommunicationMode = &v
	return s
}

func (s *UpdateGatewayRequest) SetDescription(v string) *UpdateGatewayRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayRequest) SetDistrict(v string) *UpdateGatewayRequest {
	s.District = &v
	return s
}

func (s *UpdateGatewayRequest) SetFreqBandPlanGroupId(v int64) *UpdateGatewayRequest {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *UpdateGatewayRequest) SetGisCoordinateSystem(v string) *UpdateGatewayRequest {
	s.GisCoordinateSystem = &v
	return s
}

func (s *UpdateGatewayRequest) SetGwEui(v string) *UpdateGatewayRequest {
	s.GwEui = &v
	return s
}

func (s *UpdateGatewayRequest) SetIotInstanceId(v string) *UpdateGatewayRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UpdateGatewayRequest) SetLatitude(v float32) *UpdateGatewayRequest {
	s.Latitude = &v
	return s
}

func (s *UpdateGatewayRequest) SetLongitude(v float32) *UpdateGatewayRequest {
	s.Longitude = &v
	return s
}

func (s *UpdateGatewayRequest) SetName(v string) *UpdateGatewayRequest {
	s.Name = &v
	return s
}

type UpdateGatewayResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayResponseBody) SetRequestId(v string) *UpdateGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetSuccess(v bool) *UpdateGatewayResponseBody {
	s.Success = &v
	return s
}

type UpdateGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayResponse) SetHeaders(v map[string]*string) *UpdateGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayResponse) SetStatusCode(v int32) *UpdateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayResponse) SetBody(v *UpdateGatewayResponseBody) *UpdateGatewayResponse {
	s.Body = v
	return s
}

type UpdateGatewayEnablingStateRequest struct {
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GwEui         *string `json:"GwEui,omitempty" xml:"GwEui,omitempty"`
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
}

func (s UpdateGatewayEnablingStateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayEnablingStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayEnablingStateRequest) SetEnabled(v bool) *UpdateGatewayEnablingStateRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateGatewayEnablingStateRequest) SetGwEui(v string) *UpdateGatewayEnablingStateRequest {
	s.GwEui = &v
	return s
}

func (s *UpdateGatewayEnablingStateRequest) SetIotInstanceId(v string) *UpdateGatewayEnablingStateRequest {
	s.IotInstanceId = &v
	return s
}

type UpdateGatewayEnablingStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayEnablingStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayEnablingStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayEnablingStateResponseBody) SetRequestId(v string) *UpdateGatewayEnablingStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayEnablingStateResponseBody) SetSuccess(v bool) *UpdateGatewayEnablingStateResponseBody {
	s.Success = &v
	return s
}

type UpdateGatewayEnablingStateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGatewayEnablingStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGatewayEnablingStateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayEnablingStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayEnablingStateResponse) SetHeaders(v map[string]*string) *UpdateGatewayEnablingStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayEnablingStateResponse) SetStatusCode(v int32) *UpdateGatewayEnablingStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayEnablingStateResponse) SetBody(v *UpdateGatewayEnablingStateResponseBody) *UpdateGatewayEnablingStateResponse {
	s.Body = v
	return s
}

type UpdateNodeGroupRequest struct {
	NodeGroupId   *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
}

func (s UpdateNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupRequest) SetNodeGroupId(v string) *UpdateNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetNodeGroupName(v string) *UpdateNodeGroupRequest {
	s.NodeGroupName = &v
	return s
}

type UpdateNodeGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponseBody) SetRequestId(v string) *UpdateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeGroupResponseBody) SetSuccess(v bool) *UpdateNodeGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponse) SetHeaders(v map[string]*string) *UpdateNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeGroupResponse) SetStatusCode(v int32) *UpdateNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeGroupResponse) SetBody(v *UpdateNodeGroupResponseBody) *UpdateNodeGroupResponse {
	s.Body = v
	return s
}

type UpdateNotificationsHandleStateRequest struct {
	NotificationId    []*int64 `json:"NotificationId,omitempty" xml:"NotificationId,omitempty" type:"Repeated"`
	TargetHandleState *string  `json:"TargetHandleState,omitempty" xml:"TargetHandleState,omitempty"`
}

func (s UpdateNotificationsHandleStateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNotificationsHandleStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateNotificationsHandleStateRequest) SetNotificationId(v []*int64) *UpdateNotificationsHandleStateRequest {
	s.NotificationId = v
	return s
}

func (s *UpdateNotificationsHandleStateRequest) SetTargetHandleState(v string) *UpdateNotificationsHandleStateRequest {
	s.TargetHandleState = &v
	return s
}

type UpdateNotificationsHandleStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNotificationsHandleStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNotificationsHandleStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNotificationsHandleStateResponseBody) SetRequestId(v string) *UpdateNotificationsHandleStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNotificationsHandleStateResponseBody) SetSuccess(v bool) *UpdateNotificationsHandleStateResponseBody {
	s.Success = &v
	return s
}

type UpdateNotificationsHandleStateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateNotificationsHandleStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNotificationsHandleStateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNotificationsHandleStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateNotificationsHandleStateResponse) SetHeaders(v map[string]*string) *UpdateNotificationsHandleStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateNotificationsHandleStateResponse) SetStatusCode(v int32) *UpdateNotificationsHandleStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNotificationsHandleStateResponse) SetBody(v *UpdateNotificationsHandleStateResponseBody) *UpdateNotificationsHandleStateResponse {
	s.Body = v
	return s
}

type UpdateOwnedLocalJoinPermissionRequest struct {
	ClassMode           *string `json:"ClassMode,omitempty" xml:"ClassMode,omitempty"`
	DataRate            *string `json:"DataRate,omitempty" xml:"DataRate,omitempty"`
	FreqBandPlanGroupId *int64  `json:"FreqBandPlanGroupId,omitempty" xml:"FreqBandPlanGroupId,omitempty"`
	IotInstanceId       *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JoinEui             *string `json:"JoinEui,omitempty" xml:"JoinEui,omitempty"`
	JoinPermissionId    *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName  *string `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	RxDelay             *string `json:"RxDelay,omitempty" xml:"RxDelay,omitempty"`
}

func (s UpdateOwnedLocalJoinPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOwnedLocalJoinPermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateOwnedLocalJoinPermissionRequest) SetClassMode(v string) *UpdateOwnedLocalJoinPermissionRequest {
	s.ClassMode = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionRequest) SetDataRate(v string) *UpdateOwnedLocalJoinPermissionRequest {
	s.DataRate = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionRequest) SetFreqBandPlanGroupId(v int64) *UpdateOwnedLocalJoinPermissionRequest {
	s.FreqBandPlanGroupId = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionRequest) SetIotInstanceId(v string) *UpdateOwnedLocalJoinPermissionRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionRequest) SetJoinEui(v string) *UpdateOwnedLocalJoinPermissionRequest {
	s.JoinEui = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionRequest) SetJoinPermissionId(v string) *UpdateOwnedLocalJoinPermissionRequest {
	s.JoinPermissionId = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionRequest) SetJoinPermissionName(v string) *UpdateOwnedLocalJoinPermissionRequest {
	s.JoinPermissionName = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionRequest) SetRxDelay(v string) *UpdateOwnedLocalJoinPermissionRequest {
	s.RxDelay = &v
	return s
}

type UpdateOwnedLocalJoinPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOwnedLocalJoinPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOwnedLocalJoinPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOwnedLocalJoinPermissionResponseBody) SetRequestId(v string) *UpdateOwnedLocalJoinPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionResponseBody) SetSuccess(v bool) *UpdateOwnedLocalJoinPermissionResponseBody {
	s.Success = &v
	return s
}

type UpdateOwnedLocalJoinPermissionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOwnedLocalJoinPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOwnedLocalJoinPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOwnedLocalJoinPermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateOwnedLocalJoinPermissionResponse) SetHeaders(v map[string]*string) *UpdateOwnedLocalJoinPermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionResponse) SetStatusCode(v int32) *UpdateOwnedLocalJoinPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionResponse) SetBody(v *UpdateOwnedLocalJoinPermissionResponseBody) *UpdateOwnedLocalJoinPermissionResponse {
	s.Body = v
	return s
}

type UpdateOwnedLocalJoinPermissionEnablingStateRequest struct {
	Enabled          *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	IotInstanceId    *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
}

func (s UpdateOwnedLocalJoinPermissionEnablingStateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOwnedLocalJoinPermissionEnablingStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateOwnedLocalJoinPermissionEnablingStateRequest) SetEnabled(v bool) *UpdateOwnedLocalJoinPermissionEnablingStateRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionEnablingStateRequest) SetIotInstanceId(v string) *UpdateOwnedLocalJoinPermissionEnablingStateRequest {
	s.IotInstanceId = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionEnablingStateRequest) SetJoinPermissionId(v string) *UpdateOwnedLocalJoinPermissionEnablingStateRequest {
	s.JoinPermissionId = &v
	return s
}

type UpdateOwnedLocalJoinPermissionEnablingStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOwnedLocalJoinPermissionEnablingStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOwnedLocalJoinPermissionEnablingStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOwnedLocalJoinPermissionEnablingStateResponseBody) SetRequestId(v string) *UpdateOwnedLocalJoinPermissionEnablingStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionEnablingStateResponseBody) SetSuccess(v bool) *UpdateOwnedLocalJoinPermissionEnablingStateResponseBody {
	s.Success = &v
	return s
}

type UpdateOwnedLocalJoinPermissionEnablingStateResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateOwnedLocalJoinPermissionEnablingStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateOwnedLocalJoinPermissionEnablingStateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOwnedLocalJoinPermissionEnablingStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateOwnedLocalJoinPermissionEnablingStateResponse) SetHeaders(v map[string]*string) *UpdateOwnedLocalJoinPermissionEnablingStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionEnablingStateResponse) SetStatusCode(v int32) *UpdateOwnedLocalJoinPermissionEnablingStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOwnedLocalJoinPermissionEnablingStateResponse) SetBody(v *UpdateOwnedLocalJoinPermissionEnablingStateResponseBody) *UpdateOwnedLocalJoinPermissionEnablingStateResponse {
	s.Body = v
	return s
}

type UpdateRoamingJoinPermissionRequest struct {
	DataRate           *string `json:"DataRate,omitempty" xml:"DataRate,omitempty"`
	JoinPermissionId   *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
	JoinPermissionName *string `json:"JoinPermissionName,omitempty" xml:"JoinPermissionName,omitempty"`
	RxDelay            *string `json:"RxDelay,omitempty" xml:"RxDelay,omitempty"`
}

func (s UpdateRoamingJoinPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoamingJoinPermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoamingJoinPermissionRequest) SetDataRate(v string) *UpdateRoamingJoinPermissionRequest {
	s.DataRate = &v
	return s
}

func (s *UpdateRoamingJoinPermissionRequest) SetJoinPermissionId(v string) *UpdateRoamingJoinPermissionRequest {
	s.JoinPermissionId = &v
	return s
}

func (s *UpdateRoamingJoinPermissionRequest) SetJoinPermissionName(v string) *UpdateRoamingJoinPermissionRequest {
	s.JoinPermissionName = &v
	return s
}

func (s *UpdateRoamingJoinPermissionRequest) SetRxDelay(v string) *UpdateRoamingJoinPermissionRequest {
	s.RxDelay = &v
	return s
}

type UpdateRoamingJoinPermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRoamingJoinPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoamingJoinPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoamingJoinPermissionResponseBody) SetRequestId(v string) *UpdateRoamingJoinPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoamingJoinPermissionResponseBody) SetSuccess(v bool) *UpdateRoamingJoinPermissionResponseBody {
	s.Success = &v
	return s
}

type UpdateRoamingJoinPermissionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRoamingJoinPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoamingJoinPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoamingJoinPermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoamingJoinPermissionResponse) SetHeaders(v map[string]*string) *UpdateRoamingJoinPermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoamingJoinPermissionResponse) SetStatusCode(v int32) *UpdateRoamingJoinPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoamingJoinPermissionResponse) SetBody(v *UpdateRoamingJoinPermissionResponseBody) *UpdateRoamingJoinPermissionResponse {
	s.Body = v
	return s
}

type UpdateRoamingJoinPermissionEnablingStateRequest struct {
	Enabled          *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	JoinPermissionId *string `json:"JoinPermissionId,omitempty" xml:"JoinPermissionId,omitempty"`
}

func (s UpdateRoamingJoinPermissionEnablingStateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoamingJoinPermissionEnablingStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoamingJoinPermissionEnablingStateRequest) SetEnabled(v bool) *UpdateRoamingJoinPermissionEnablingStateRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateRoamingJoinPermissionEnablingStateRequest) SetJoinPermissionId(v string) *UpdateRoamingJoinPermissionEnablingStateRequest {
	s.JoinPermissionId = &v
	return s
}

type UpdateRoamingJoinPermissionEnablingStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRoamingJoinPermissionEnablingStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoamingJoinPermissionEnablingStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoamingJoinPermissionEnablingStateResponseBody) SetRequestId(v string) *UpdateRoamingJoinPermissionEnablingStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoamingJoinPermissionEnablingStateResponseBody) SetSuccess(v bool) *UpdateRoamingJoinPermissionEnablingStateResponseBody {
	s.Success = &v
	return s
}

type UpdateRoamingJoinPermissionEnablingStateResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRoamingJoinPermissionEnablingStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoamingJoinPermissionEnablingStateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoamingJoinPermissionEnablingStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoamingJoinPermissionEnablingStateResponse) SetHeaders(v map[string]*string) *UpdateRoamingJoinPermissionEnablingStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoamingJoinPermissionEnablingStateResponse) SetStatusCode(v int32) *UpdateRoamingJoinPermissionEnablingStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoamingJoinPermissionEnablingStateResponse) SetBody(v *UpdateRoamingJoinPermissionEnablingStateResponseBody) *UpdateRoamingJoinPermissionEnablingStateResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkwan"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AcceptJoinPermissionAuthOrderWithOptions(request *AcceptJoinPermissionAuthOrderRequest, runtime *util.RuntimeOptions) (_result *AcceptJoinPermissionAuthOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AcceptJoinPermissionAuthOrder"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptJoinPermissionAuthOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AcceptJoinPermissionAuthOrder(request *AcceptJoinPermissionAuthOrderRequest) (_result *AcceptJoinPermissionAuthOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptJoinPermissionAuthOrderResponse{}
	_body, _err := client.AcceptJoinPermissionAuthOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddNodeToGroupWithOptions(request *AddNodeToGroupRequest, runtime *util.RuntimeOptions) (_result *AddNodeToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PinCode)) {
		query["PinCode"] = request.PinCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddNodeToGroup"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddNodeToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddNodeToGroup(request *AddNodeToGroupRequest) (_result *AddNodeToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddNodeToGroupResponse{}
	_body, _err := client.AddNodeToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyRoamingJoinPermissionWithOptions(request *ApplyRoamingJoinPermissionRequest, runtime *util.RuntimeOptions) (_result *ApplyRoamingJoinPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassMode)) {
		query["ClassMode"] = request.ClassMode
	}

	if !tea.BoolValue(util.IsUnset(request.DataRate)) {
		query["DataRate"] = request.DataRate
	}

	if !tea.BoolValue(util.IsUnset(request.FreqBandPlanGroupId)) {
		query["FreqBandPlanGroupId"] = request.FreqBandPlanGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionName)) {
		query["JoinPermissionName"] = request.JoinPermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.RxDelay)) {
		query["RxDelay"] = request.RxDelay
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyRoamingJoinPermission"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyRoamingJoinPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyRoamingJoinPermission(request *ApplyRoamingJoinPermissionRequest) (_result *ApplyRoamingJoinPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyRoamingJoinPermissionResponse{}
	_body, _err := client.ApplyRoamingJoinPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindJoinPermissionToNodeGroupWithOptions(request *BindJoinPermissionToNodeGroupRequest, runtime *util.RuntimeOptions) (_result *BindJoinPermissionToNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindJoinPermissionToNodeGroup"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindJoinPermissionToNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindJoinPermissionToNodeGroup(request *BindJoinPermissionToNodeGroupRequest) (_result *BindJoinPermissionToNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindJoinPermissionToNodeGroupResponse{}
	_body, _err := client.BindJoinPermissionToNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelJoinPermissionAuthOrderWithOptions(request *CancelJoinPermissionAuthOrderRequest, runtime *util.RuntimeOptions) (_result *CancelJoinPermissionAuthOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelJoinPermissionAuthOrder"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelJoinPermissionAuthOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelJoinPermissionAuthOrder(request *CancelJoinPermissionAuthOrderRequest) (_result *CancelJoinPermissionAuthOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelJoinPermissionAuthOrderResponse{}
	_body, _err := client.CancelJoinPermissionAuthOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckCloudProductOpenStatusWithOptions(request *CheckCloudProductOpenStatusRequest, runtime *util.RuntimeOptions) (_result *CheckCloudProductOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckCloudProductOpenStatus"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckCloudProductOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckCloudProductOpenStatus(request *CheckCloudProductOpenStatusRequest) (_result *CheckCloudProductOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckCloudProductOpenStatusResponse{}
	_body, _err := client.CheckCloudProductOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckUserChargeStatusWithOptions(runtime *util.RuntimeOptions) (_result *CheckUserChargeStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CheckUserChargeStatus"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUserChargeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUserChargeStatus() (_result *CheckUserChargeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUserChargeStatusResponse{}
	_body, _err := client.CheckUserChargeStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountGatewayTupleOrdersWithOptions(request *CountGatewayTupleOrdersRequest, runtime *util.RuntimeOptions) (_result *CountGatewayTupleOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.States)) {
		query["States"] = request.States
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountGatewayTupleOrders"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountGatewayTupleOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountGatewayTupleOrders(request *CountGatewayTupleOrdersRequest) (_result *CountGatewayTupleOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountGatewayTupleOrdersResponse{}
	_body, _err := client.CountGatewayTupleOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountGatewaysWithOptions(request *CountGatewaysRequest, runtime *util.RuntimeOptions) (_result *CountGatewaysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FreqBandPlanGroupId)) {
		query["FreqBandPlanGroupId"] = request.FreqBandPlanGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyCity)) {
		query["FuzzyCity"] = request.FuzzyCity
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyGwEui)) {
		query["FuzzyGwEui"] = request.FuzzyGwEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyName)) {
		query["FuzzyName"] = request.FuzzyName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnabled)) {
		query["IsEnabled"] = request.IsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.OnlineState)) {
		query["OnlineState"] = request.OnlineState
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountGateways"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountGateways(request *CountGatewaysRequest) (_result *CountGatewaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountGatewaysResponse{}
	_body, _err := client.CountGatewaysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountNodeGroupsWithOptions(request *CountNodeGroupsRequest, runtime *util.RuntimeOptions) (_result *CountNodeGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FuzzyDevEui)) {
		query["FuzzyDevEui"] = request.FuzzyDevEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinEui)) {
		query["FuzzyJoinEui"] = request.FuzzyJoinEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyName)) {
		query["FuzzyName"] = request.FuzzyName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountNodeGroups"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountNodeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountNodeGroups(request *CountNodeGroupsRequest) (_result *CountNodeGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountNodeGroupsResponse{}
	_body, _err := client.CountNodeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountNodeTupleOrdersWithOptions(request *CountNodeTupleOrdersRequest, runtime *util.RuntimeOptions) (_result *CountNodeTupleOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsKpm)) {
		query["IsKpm"] = request.IsKpm
	}

	if !tea.BoolValue(util.IsUnset(request.States)) {
		query["States"] = request.States
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountNodeTupleOrders"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountNodeTupleOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountNodeTupleOrders(request *CountNodeTupleOrdersRequest) (_result *CountNodeTupleOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountNodeTupleOrdersResponse{}
	_body, _err := client.CountNodeTupleOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountNodesByNodeGroupIdWithOptions(request *CountNodesByNodeGroupIdRequest, runtime *util.RuntimeOptions) (_result *CountNodesByNodeGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FuzzyDevEui)) {
		query["FuzzyDevEui"] = request.FuzzyDevEui
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountNodesByNodeGroupId"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountNodesByNodeGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountNodesByNodeGroupId(request *CountNodesByNodeGroupIdRequest) (_result *CountNodesByNodeGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountNodesByNodeGroupIdResponse{}
	_body, _err := client.CountNodesByNodeGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountNodesByOwnedJoinPermissionIdWithOptions(request *CountNodesByOwnedJoinPermissionIdRequest, runtime *util.RuntimeOptions) (_result *CountNodesByOwnedJoinPermissionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FuzzyDevEui)) {
		query["FuzzyDevEui"] = request.FuzzyDevEui
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountNodesByOwnedJoinPermissionId"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountNodesByOwnedJoinPermissionIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountNodesByOwnedJoinPermissionId(request *CountNodesByOwnedJoinPermissionIdRequest) (_result *CountNodesByOwnedJoinPermissionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountNodesByOwnedJoinPermissionIdResponse{}
	_body, _err := client.CountNodesByOwnedJoinPermissionIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountNotificationsWithOptions(request *CountNotificationsRequest, runtime *util.RuntimeOptions) (_result *CountNotificationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.HandleState)) {
		query["HandleState"] = request.HandleState
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountNotifications"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountNotificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountNotifications(request *CountNotificationsRequest) (_result *CountNotificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountNotificationsResponse{}
	_body, _err := client.CountNotificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountOwnedJoinPermissionsWithOptions(request *CountOwnedJoinPermissionsRequest, runtime *util.RuntimeOptions) (_result *CountOwnedJoinPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinEui)) {
		query["FuzzyJoinEui"] = request.FuzzyJoinEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinPermissionName)) {
		query["FuzzyJoinPermissionName"] = request.FuzzyJoinPermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyRenterAliyunId)) {
		query["FuzzyRenterAliyunId"] = request.FuzzyRenterAliyunId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountOwnedJoinPermissions"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountOwnedJoinPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountOwnedJoinPermissions(request *CountOwnedJoinPermissionsRequest) (_result *CountOwnedJoinPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountOwnedJoinPermissionsResponse{}
	_body, _err := client.CountOwnedJoinPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CountRentedJoinPermissionsWithOptions(request *CountRentedJoinPermissionsRequest, runtime *util.RuntimeOptions) (_result *CountRentedJoinPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BoundNodeGroup)) {
		query["BoundNodeGroup"] = request.BoundNodeGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinEui)) {
		query["FuzzyJoinEui"] = request.FuzzyJoinEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinPermissionName)) {
		query["FuzzyJoinPermissionName"] = request.FuzzyJoinPermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyOwnerAliyunId)) {
		query["FuzzyOwnerAliyunId"] = request.FuzzyOwnerAliyunId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CountRentedJoinPermissions"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CountRentedJoinPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CountRentedJoinPermissions(request *CountRentedJoinPermissionsRequest) (_result *CountRentedJoinPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CountRentedJoinPermissionsResponse{}
	_body, _err := client.CountRentedJoinPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGatewayWithOptions(request *CreateGatewayRequest, runtime *util.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.AddressCode)) {
		query["AddressCode"] = request.AddressCode
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.CommunicationMode)) {
		query["CommunicationMode"] = request.CommunicationMode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.District)) {
		query["District"] = request.District
	}

	if !tea.BoolValue(util.IsUnset(request.FreqBandPlanGroupId)) {
		query["FreqBandPlanGroupId"] = request.FreqBandPlanGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GisCoordinateSystem)) {
		query["GisCoordinateSystem"] = request.GisCoordinateSystem
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Latitude)) {
		query["Latitude"] = request.Latitude
	}

	if !tea.BoolValue(util.IsUnset(request.Longitude)) {
		query["Longitude"] = request.Longitude
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PinCode)) {
		query["PinCode"] = request.PinCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGateway"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGateway(request *CreateGatewayRequest) (_result *CreateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CreateGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLocalJoinPermissionWithOptions(request *CreateLocalJoinPermissionRequest, runtime *util.RuntimeOptions) (_result *CreateLocalJoinPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassMode)) {
		query["ClassMode"] = request.ClassMode
	}

	if !tea.BoolValue(util.IsUnset(request.DataRate)) {
		query["DataRate"] = request.DataRate
	}

	if !tea.BoolValue(util.IsUnset(request.FreqBandPlanGroupId)) {
		query["FreqBandPlanGroupId"] = request.FreqBandPlanGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinEui)) {
		query["JoinEui"] = request.JoinEui
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionName)) {
		query["JoinPermissionName"] = request.JoinPermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.RxDelay)) {
		query["RxDelay"] = request.RxDelay
	}

	if !tea.BoolValue(util.IsUnset(request.UseDefaultJoinEui)) {
		query["UseDefaultJoinEui"] = request.UseDefaultJoinEui
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLocalJoinPermission"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLocalJoinPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLocalJoinPermission(request *CreateLocalJoinPermissionRequest) (_result *CreateLocalJoinPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLocalJoinPermissionResponse{}
	_body, _err := client.CreateLocalJoinPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNodeGroupWithOptions(request *CreateNodeGroupRequest, runtime *util.RuntimeOptions) (_result *CreateNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupName)) {
		query["NodeGroupName"] = request.NodeGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNodeGroup"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNodeGroup(request *CreateNodeGroupRequest) (_result *CreateNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodeGroupResponse{}
	_body, _err := client.CreateNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGatewayWithOptions(request *DeleteGatewayRequest, runtime *util.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGateway"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGateway(request *DeleteGatewayRequest) (_result *DeleteGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLocalJoinPermissionWithOptions(request *DeleteLocalJoinPermissionRequest, runtime *util.RuntimeOptions) (_result *DeleteLocalJoinPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLocalJoinPermission"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLocalJoinPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLocalJoinPermission(request *DeleteLocalJoinPermissionRequest) (_result *DeleteLocalJoinPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLocalJoinPermissionResponse{}
	_body, _err := client.DeleteLocalJoinPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNodeGroupWithOptions(request *DeleteNodeGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNodeGroup"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNodeGroup(request *DeleteNodeGroupRequest) (_result *DeleteNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodeGroupResponse{}
	_body, _err := client.DeleteNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFreqBandPlanGroupWithOptions(request *GetFreqBandPlanGroupRequest, runtime *util.RuntimeOptions) (_result *GetFreqBandPlanGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFreqBandPlanGroup"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFreqBandPlanGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFreqBandPlanGroup(request *GetFreqBandPlanGroupRequest) (_result *GetFreqBandPlanGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFreqBandPlanGroupResponse{}
	_body, _err := client.GetFreqBandPlanGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGatewayWithOptions(request *GetGatewayRequest, runtime *util.RuntimeOptions) (_result *GetGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGateway"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGateway(request *GetGatewayRequest) (_result *GetGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGatewayResponse{}
	_body, _err := client.GetGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGatewayPacketStatWithOptions(request *GetGatewayPacketStatRequest, runtime *util.RuntimeOptions) (_result *GetGatewayPacketStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGatewayPacketStat"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayPacketStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGatewayPacketStat(request *GetGatewayPacketStatRequest) (_result *GetGatewayPacketStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGatewayPacketStatResponse{}
	_body, _err := client.GetGatewayPacketStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGatewayStatusStatWithOptions(request *GetGatewayStatusStatRequest, runtime *util.RuntimeOptions) (_result *GetGatewayStatusStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGatewayStatusStat"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayStatusStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGatewayStatusStat(request *GetGatewayStatusStatRequest) (_result *GetGatewayStatusStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGatewayStatusStatResponse{}
	_body, _err := client.GetGatewayStatusStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGatewayTransferPacketsDownloadUrlWithOptions(request *GetGatewayTransferPacketsDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetGatewayTransferPacketsDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGatewayTransferPacketsDownloadUrl"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayTransferPacketsDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGatewayTransferPacketsDownloadUrl(request *GetGatewayTransferPacketsDownloadUrlRequest) (_result *GetGatewayTransferPacketsDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGatewayTransferPacketsDownloadUrlResponse{}
	_body, _err := client.GetGatewayTransferPacketsDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGatewayTupleOrderWithOptions(request *GetGatewayTupleOrderRequest, runtime *util.RuntimeOptions) (_result *GetGatewayTupleOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGatewayTupleOrder"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayTupleOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGatewayTupleOrder(request *GetGatewayTupleOrderRequest) (_result *GetGatewayTupleOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGatewayTupleOrderResponse{}
	_body, _err := client.GetGatewayTupleOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGatewayTuplesDownloadUrlWithOptions(request *GetGatewayTuplesDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetGatewayTuplesDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGatewayTuplesDownloadUrl"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayTuplesDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGatewayTuplesDownloadUrl(request *GetGatewayTuplesDownloadUrlRequest) (_result *GetGatewayTuplesDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGatewayTuplesDownloadUrlResponse{}
	_body, _err := client.GetGatewayTuplesDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJoinPermissionAuthOrderWithOptions(request *GetJoinPermissionAuthOrderRequest, runtime *util.RuntimeOptions) (_result *GetJoinPermissionAuthOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJoinPermissionAuthOrder"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJoinPermissionAuthOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJoinPermissionAuthOrder(request *GetJoinPermissionAuthOrderRequest) (_result *GetJoinPermissionAuthOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJoinPermissionAuthOrderResponse{}
	_body, _err := client.GetJoinPermissionAuthOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeWithOptions(request *GetNodeRequest, runtime *util.RuntimeOptions) (_result *GetNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNode"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNode(request *GetNodeRequest) (_result *GetNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeResponse{}
	_body, _err := client.GetNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeGroupWithOptions(request *GetNodeGroupRequest, runtime *util.RuntimeOptions) (_result *GetNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeGroup"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeGroup(request *GetNodeGroupRequest) (_result *GetNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeGroupResponse{}
	_body, _err := client.GetNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeGroupTransferPacketsDownloadUrlWithOptions(request *GetNodeGroupTransferPacketsDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetNodeGroupTransferPacketsDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeGroupTransferPacketsDownloadUrl"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeGroupTransferPacketsDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeGroupTransferPacketsDownloadUrl(request *GetNodeGroupTransferPacketsDownloadUrlRequest) (_result *GetNodeGroupTransferPacketsDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeGroupTransferPacketsDownloadUrlResponse{}
	_body, _err := client.GetNodeGroupTransferPacketsDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeTransferPacketWithOptions(request *GetNodeTransferPacketRequest, runtime *util.RuntimeOptions) (_result *GetNodeTransferPacketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Base64EncodedMacPayload)) {
		query["Base64EncodedMacPayload"] = request.Base64EncodedMacPayload
	}

	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LogMillis)) {
		query["LogMillis"] = request.LogMillis
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeTransferPacket"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeTransferPacketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeTransferPacket(request *GetNodeTransferPacketRequest) (_result *GetNodeTransferPacketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeTransferPacketResponse{}
	_body, _err := client.GetNodeTransferPacketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeTransferPacketsDownloadUrlWithOptions(request *GetNodeTransferPacketsDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetNodeTransferPacketsDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeTransferPacketsDownloadUrl"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeTransferPacketsDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeTransferPacketsDownloadUrl(request *GetNodeTransferPacketsDownloadUrlRequest) (_result *GetNodeTransferPacketsDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeTransferPacketsDownloadUrlResponse{}
	_body, _err := client.GetNodeTransferPacketsDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeTupleOrderWithOptions(request *GetNodeTupleOrderRequest, runtime *util.RuntimeOptions) (_result *GetNodeTupleOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeTupleOrder"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeTupleOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeTupleOrder(request *GetNodeTupleOrderRequest) (_result *GetNodeTupleOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeTupleOrderResponse{}
	_body, _err := client.GetNodeTupleOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeTuplesDownloadUrlWithOptions(request *GetNodeTuplesDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetNodeTuplesDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeTuplesDownloadUrl"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeTuplesDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeTuplesDownloadUrl(request *GetNodeTuplesDownloadUrlRequest) (_result *GetNodeTuplesDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeTuplesDownloadUrlResponse{}
	_body, _err := client.GetNodeTuplesDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNotificationWithOptions(request *GetNotificationRequest, runtime *util.RuntimeOptions) (_result *GetNotificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotificationId)) {
		query["NotificationId"] = request.NotificationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNotification"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNotificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNotification(request *GetNotificationRequest) (_result *GetNotificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNotificationResponse{}
	_body, _err := client.GetNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOwnedJoinPermissionWithOptions(request *GetOwnedJoinPermissionRequest, runtime *util.RuntimeOptions) (_result *GetOwnedJoinPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOwnedJoinPermission"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOwnedJoinPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOwnedJoinPermission(request *GetOwnedJoinPermissionRequest) (_result *GetOwnedJoinPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOwnedJoinPermissionResponse{}
	_body, _err := client.GetOwnedJoinPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRentedJoinPermissionWithOptions(request *GetRentedJoinPermissionRequest, runtime *util.RuntimeOptions) (_result *GetRentedJoinPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRentedJoinPermission"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRentedJoinPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRentedJoinPermission(request *GetRentedJoinPermissionRequest) (_result *GetRentedJoinPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRentedJoinPermissionResponse{}
	_body, _err := client.GetRentedJoinPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserLicenseWithOptions(runtime *util.RuntimeOptions) (_result *GetUserLicenseResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetUserLicense"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserLicense() (_result *GetUserLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserLicenseResponse{}
	_body, _err := client.GetUserLicenseWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListActivatedFeaturesWithOptions(request *ListActivatedFeaturesRequest, runtime *util.RuntimeOptions) (_result *ListActivatedFeaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListActivatedFeatures"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListActivatedFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListActivatedFeatures(request *ListActivatedFeaturesRequest) (_result *ListActivatedFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListActivatedFeaturesResponse{}
	_body, _err := client.ListActivatedFeaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListActiveGatewaysWithOptions(runtime *util.RuntimeOptions) (_result *ListActiveGatewaysResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListActiveGateways"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListActiveGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListActiveGateways() (_result *ListActiveGatewaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListActiveGatewaysResponse{}
	_body, _err := client.ListActiveGatewaysWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFreqBandPlanGroupsWithOptions(runtime *util.RuntimeOptions) (_result *ListFreqBandPlanGroupsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListFreqBandPlanGroups"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFreqBandPlanGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFreqBandPlanGroups() (_result *ListFreqBandPlanGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFreqBandPlanGroupsResponse{}
	_body, _err := client.ListFreqBandPlanGroupsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGatewayOnlineRecordsWithOptions(request *ListGatewayOnlineRecordsRequest, runtime *util.RuntimeOptions) (_result *ListGatewayOnlineRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OffSet)) {
		query["OffSet"] = request.OffSet
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGatewayOnlineRecords"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewayOnlineRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGatewayOnlineRecords(request *ListGatewayOnlineRecordsRequest) (_result *ListGatewayOnlineRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGatewayOnlineRecordsResponse{}
	_body, _err := client.ListGatewayOnlineRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGatewayTransferFlowStatsWithOptions(request *ListGatewayTransferFlowStatsRequest, runtime *util.RuntimeOptions) (_result *ListGatewayTransferFlowStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeIntervalUnit)) {
		query["TimeIntervalUnit"] = request.TimeIntervalUnit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGatewayTransferFlowStats"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewayTransferFlowStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGatewayTransferFlowStats(request *ListGatewayTransferFlowStatsRequest) (_result *ListGatewayTransferFlowStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGatewayTransferFlowStatsResponse{}
	_body, _err := client.ListGatewayTransferFlowStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGatewayTransferPacketsWithOptions(request *ListGatewayTransferPacketsRequest, runtime *util.RuntimeOptions) (_result *ListGatewayTransferPacketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGatewayTransferPackets"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewayTransferPacketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGatewayTransferPackets(request *ListGatewayTransferPacketsRequest) (_result *ListGatewayTransferPacketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGatewayTransferPacketsResponse{}
	_body, _err := client.ListGatewayTransferPacketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGatewayTupleOrdersWithOptions(request *ListGatewayTupleOrdersRequest, runtime *util.RuntimeOptions) (_result *ListGatewayTupleOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGatewayTupleOrders"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewayTupleOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGatewayTupleOrders(request *ListGatewayTupleOrdersRequest) (_result *ListGatewayTupleOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGatewayTupleOrdersResponse{}
	_body, _err := client.ListGatewayTupleOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGatewaysWithOptions(request *ListGatewaysRequest, runtime *util.RuntimeOptions) (_result *ListGatewaysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.FreqBandPlanGroupId)) {
		query["FreqBandPlanGroupId"] = request.FreqBandPlanGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyCity)) {
		query["FuzzyCity"] = request.FuzzyCity
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyGwEui)) {
		query["FuzzyGwEui"] = request.FuzzyGwEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyName)) {
		query["FuzzyName"] = request.FuzzyName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnabled)) {
		query["IsEnabled"] = request.IsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.OnlineState)) {
		query["OnlineState"] = request.OnlineState
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGateways"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGateways(request *ListGatewaysRequest) (_result *ListGatewaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGatewaysResponse{}
	_body, _err := client.ListGatewaysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGatewaysGisInfoWithOptions(request *ListGatewaysGisInfoRequest, runtime *util.RuntimeOptions) (_result *ListGatewaysGisInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGatewaysGisInfo"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewaysGisInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGatewaysGisInfo(request *ListGatewaysGisInfoRequest) (_result *ListGatewaysGisInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGatewaysGisInfoResponse{}
	_body, _err := client.ListGatewaysGisInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeGroupTransferFlowStatsWithOptions(request *ListNodeGroupTransferFlowStatsRequest, runtime *util.RuntimeOptions) (_result *ListNodeGroupTransferFlowStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeIntervalUnit)) {
		query["TimeIntervalUnit"] = request.TimeIntervalUnit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeGroupTransferFlowStats"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeGroupTransferFlowStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeGroupTransferFlowStats(request *ListNodeGroupTransferFlowStatsRequest) (_result *ListNodeGroupTransferFlowStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeGroupTransferFlowStatsResponse{}
	_body, _err := client.ListNodeGroupTransferFlowStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeGroupTransferPacketsWithOptions(request *ListNodeGroupTransferPacketsRequest, runtime *util.RuntimeOptions) (_result *ListNodeGroupTransferPacketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeGroupTransferPackets"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeGroupTransferPacketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeGroupTransferPackets(request *ListNodeGroupTransferPacketsRequest) (_result *ListNodeGroupTransferPacketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeGroupTransferPacketsResponse{}
	_body, _err := client.ListNodeGroupTransferPacketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeGroupsWithOptions(request *ListNodeGroupsRequest, runtime *util.RuntimeOptions) (_result *ListNodeGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyDevEui)) {
		query["FuzzyDevEui"] = request.FuzzyDevEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinEui)) {
		query["FuzzyJoinEui"] = request.FuzzyJoinEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyName)) {
		query["FuzzyName"] = request.FuzzyName
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeGroups"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeGroups(request *ListNodeGroupsRequest) (_result *ListNodeGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.ListNodeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeTransferPacketPathsWithOptions(request *ListNodeTransferPacketPathsRequest, runtime *util.RuntimeOptions) (_result *ListNodeTransferPacketPathsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Base64EncodedMacPayload)) {
		query["Base64EncodedMacPayload"] = request.Base64EncodedMacPayload
	}

	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LogMillis)) {
		query["LogMillis"] = request.LogMillis
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeTransferPacketPaths"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeTransferPacketPathsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeTransferPacketPaths(request *ListNodeTransferPacketPathsRequest) (_result *ListNodeTransferPacketPathsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeTransferPacketPathsResponse{}
	_body, _err := client.ListNodeTransferPacketPathsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeTransferPacketsWithOptions(request *ListNodeTransferPacketsRequest, runtime *util.RuntimeOptions) (_result *ListNodeTransferPacketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeTransferPackets"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeTransferPacketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeTransferPackets(request *ListNodeTransferPacketsRequest) (_result *ListNodeTransferPacketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeTransferPacketsResponse{}
	_body, _err := client.ListNodeTransferPacketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeTupleOrdersWithOptions(request *ListNodeTupleOrdersRequest, runtime *util.RuntimeOptions) (_result *ListNodeTupleOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.IsKpm)) {
		query["IsKpm"] = request.IsKpm
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeTupleOrders"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeTupleOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeTupleOrders(request *ListNodeTupleOrdersRequest) (_result *ListNodeTupleOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeTupleOrdersResponse{}
	_body, _err := client.ListNodeTupleOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesByNodeGroupIdWithOptions(request *ListNodesByNodeGroupIdRequest, runtime *util.RuntimeOptions) (_result *ListNodesByNodeGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyDevEui)) {
		query["FuzzyDevEui"] = request.FuzzyDevEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodesByNodeGroupId"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesByNodeGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodesByNodeGroupId(request *ListNodesByNodeGroupIdRequest) (_result *ListNodesByNodeGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesByNodeGroupIdResponse{}
	_body, _err := client.ListNodesByNodeGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesByOwnedJoinPermissionIdWithOptions(request *ListNodesByOwnedJoinPermissionIdRequest, runtime *util.RuntimeOptions) (_result *ListNodesByOwnedJoinPermissionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyDevEui)) {
		query["FuzzyDevEui"] = request.FuzzyDevEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodesByOwnedJoinPermissionId"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesByOwnedJoinPermissionIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodesByOwnedJoinPermissionId(request *ListNodesByOwnedJoinPermissionIdRequest) (_result *ListNodesByOwnedJoinPermissionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesByOwnedJoinPermissionIdResponse{}
	_body, _err := client.ListNodesByOwnedJoinPermissionIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNotificationsWithOptions(request *ListNotificationsRequest, runtime *util.RuntimeOptions) (_result *ListNotificationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.BeginMillis)) {
		query["BeginMillis"] = request.BeginMillis
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillis)) {
		query["EndMillis"] = request.EndMillis
	}

	if !tea.BoolValue(util.IsUnset(request.HandleState)) {
		query["HandleState"] = request.HandleState
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNotifications"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNotificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNotifications(request *ListNotificationsRequest) (_result *ListNotificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNotificationsResponse{}
	_body, _err := client.ListNotificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOwnedJoinPermissionsWithOptions(request *ListOwnedJoinPermissionsRequest, runtime *util.RuntimeOptions) (_result *ListOwnedJoinPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinEui)) {
		query["FuzzyJoinEui"] = request.FuzzyJoinEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinPermissionName)) {
		query["FuzzyJoinPermissionName"] = request.FuzzyJoinPermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyRenterAliyunId)) {
		query["FuzzyRenterAliyunId"] = request.FuzzyRenterAliyunId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOwnedJoinPermissions"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOwnedJoinPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOwnedJoinPermissions(request *ListOwnedJoinPermissionsRequest) (_result *ListOwnedJoinPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOwnedJoinPermissionsResponse{}
	_body, _err := client.ListOwnedJoinPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRentedJoinPermissionsWithOptions(request *ListRentedJoinPermissionsRequest, runtime *util.RuntimeOptions) (_result *ListRentedJoinPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.BoundNodeGroup)) {
		query["BoundNodeGroup"] = request.BoundNodeGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinEui)) {
		query["FuzzyJoinEui"] = request.FuzzyJoinEui
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyJoinPermissionName)) {
		query["FuzzyJoinPermissionName"] = request.FuzzyJoinPermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyOwnerAliyunId)) {
		query["FuzzyOwnerAliyunId"] = request.FuzzyOwnerAliyunId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.SortingField)) {
		query["SortingField"] = request.SortingField
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRentedJoinPermissions"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRentedJoinPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRentedJoinPermissions(request *ListRentedJoinPermissionsRequest) (_result *ListRentedJoinPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRentedJoinPermissionsResponse{}
	_body, _err := client.ListRentedJoinPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RejectJoinPermissionAuthOrderWithOptions(request *RejectJoinPermissionAuthOrderRequest, runtime *util.RuntimeOptions) (_result *RejectJoinPermissionAuthOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RejectJoinPermissionAuthOrder"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectJoinPermissionAuthOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RejectJoinPermissionAuthOrder(request *RejectJoinPermissionAuthOrderRequest) (_result *RejectJoinPermissionAuthOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectJoinPermissionAuthOrderResponse{}
	_body, _err := client.RejectJoinPermissionAuthOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveNodeFromGroupWithOptions(request *RemoveNodeFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveNodeFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveNodeFromGroup"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveNodeFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveNodeFromGroup(request *RemoveNodeFromGroupRequest) (_result *RemoveNodeFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveNodeFromGroupResponse{}
	_body, _err := client.RemoveNodeFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReturnJoinPermissionWithOptions(request *ReturnJoinPermissionRequest, runtime *util.RuntimeOptions) (_result *ReturnJoinPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionType)) {
		query["JoinPermissionType"] = request.JoinPermissionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReturnJoinPermission"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReturnJoinPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReturnJoinPermission(request *ReturnJoinPermissionRequest) (_result *ReturnJoinPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReturnJoinPermissionResponse{}
	_body, _err := client.ReturnJoinPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendUnicastCommandWithOptions(request *SendUnicastCommandRequest, runtime *util.RuntimeOptions) (_result *SendUnicastCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CleanUp)) {
		query["CleanUp"] = request.CleanUp
	}

	if !tea.BoolValue(util.IsUnset(request.Confirmed)) {
		query["Confirmed"] = request.Confirmed
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DevEui)) {
		query["DevEui"] = request.DevEui
	}

	if !tea.BoolValue(util.IsUnset(request.FPort)) {
		query["FPort"] = request.FPort
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxRetries)) {
		query["MaxRetries"] = request.MaxRetries
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendUnicastCommand"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendUnicastCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendUnicastCommand(request *SendUnicastCommandRequest) (_result *SendUnicastCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendUnicastCommandResponse{}
	_body, _err := client.SendUnicastCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitGatewayTupleOrderWithOptions(request *SubmitGatewayTupleOrderRequest, runtime *util.RuntimeOptions) (_result *SubmitGatewayTupleOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequiredCount)) {
		query["RequiredCount"] = request.RequiredCount
	}

	if !tea.BoolValue(util.IsUnset(request.TupleType)) {
		query["TupleType"] = request.TupleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitGatewayTupleOrder"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitGatewayTupleOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitGatewayTupleOrder(request *SubmitGatewayTupleOrderRequest) (_result *SubmitGatewayTupleOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitGatewayTupleOrderResponse{}
	_body, _err := client.SubmitGatewayTupleOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitJoinPermissionAuthOrderWithOptions(request *SubmitJoinPermissionAuthOrderRequest, runtime *util.RuntimeOptions) (_result *SubmitJoinPermissionAuthOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.RenterAliyunId)) {
		query["RenterAliyunId"] = request.RenterAliyunId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitJoinPermissionAuthOrder"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitJoinPermissionAuthOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitJoinPermissionAuthOrder(request *SubmitJoinPermissionAuthOrderRequest) (_result *SubmitJoinPermissionAuthOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitJoinPermissionAuthOrderResponse{}
	_body, _err := client.SubmitJoinPermissionAuthOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitNodeTupleOrderWithOptions(request *SubmitNodeTupleOrderRequest, runtime *util.RuntimeOptions) (_result *SubmitNodeTupleOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoraVersion)) {
		query["LoraVersion"] = request.LoraVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RequiredCount)) {
		query["RequiredCount"] = request.RequiredCount
	}

	if !tea.BoolValue(util.IsUnset(request.TupleType)) {
		query["TupleType"] = request.TupleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitNodeTupleOrder"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitNodeTupleOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitNodeTupleOrder(request *SubmitNodeTupleOrderRequest) (_result *SubmitNodeTupleOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitNodeTupleOrderResponse{}
	_body, _err := client.SubmitNodeTupleOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindJoinPermissionFromNodeGroupWithOptions(request *UnbindJoinPermissionFromNodeGroupRequest, runtime *util.RuntimeOptions) (_result *UnbindJoinPermissionFromNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindJoinPermissionFromNodeGroup"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindJoinPermissionFromNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindJoinPermissionFromNodeGroup(request *UnbindJoinPermissionFromNodeGroupRequest) (_result *UnbindJoinPermissionFromNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindJoinPermissionFromNodeGroupResponse{}
	_body, _err := client.UnbindJoinPermissionFromNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDataDispatchConfigWithOptions(request *UpdateDataDispatchConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateDataDispatchConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataDispatchDestination)) {
		query["DataDispatchDestination"] = request.DataDispatchDestination
	}

	if !tea.BoolValue(util.IsUnset(request.DebugSwitch)) {
		query["DebugSwitch"] = request.DebugSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.UplinkRegionName)) {
		query["UplinkRegionName"] = request.UplinkRegionName
	}

	if !tea.BoolValue(util.IsUnset(request.UplinkTopic)) {
		query["UplinkTopic"] = request.UplinkTopic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataDispatchConfig"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDataDispatchConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDataDispatchConfig(request *UpdateDataDispatchConfigRequest) (_result *UpdateDataDispatchConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDataDispatchConfigResponse{}
	_body, _err := client.UpdateDataDispatchConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDataDispatchEnablingStateWithOptions(request *UpdateDataDispatchEnablingStateRequest, runtime *util.RuntimeOptions) (_result *UpdateDataDispatchEnablingStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataDispatchEnabled)) {
		query["DataDispatchEnabled"] = request.DataDispatchEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataDispatchEnablingState"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDataDispatchEnablingStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDataDispatchEnablingState(request *UpdateDataDispatchEnablingStateRequest) (_result *UpdateDataDispatchEnablingStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDataDispatchEnablingStateResponse{}
	_body, _err := client.UpdateDataDispatchEnablingStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGatewayWithOptions(request *UpdateGatewayRequest, runtime *util.RuntimeOptions) (_result *UpdateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.AddressCode)) {
		query["AddressCode"] = request.AddressCode
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.CommunicationMode)) {
		query["CommunicationMode"] = request.CommunicationMode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.District)) {
		query["District"] = request.District
	}

	if !tea.BoolValue(util.IsUnset(request.FreqBandPlanGroupId)) {
		query["FreqBandPlanGroupId"] = request.FreqBandPlanGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GisCoordinateSystem)) {
		query["GisCoordinateSystem"] = request.GisCoordinateSystem
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Latitude)) {
		query["Latitude"] = request.Latitude
	}

	if !tea.BoolValue(util.IsUnset(request.Longitude)) {
		query["Longitude"] = request.Longitude
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGateway"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGateway(request *UpdateGatewayRequest) (_result *UpdateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGatewayResponse{}
	_body, _err := client.UpdateGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGatewayEnablingStateWithOptions(request *UpdateGatewayEnablingStateRequest, runtime *util.RuntimeOptions) (_result *UpdateGatewayEnablingStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.GwEui)) {
		query["GwEui"] = request.GwEui
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayEnablingState"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayEnablingStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGatewayEnablingState(request *UpdateGatewayEnablingStateRequest) (_result *UpdateGatewayEnablingStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGatewayEnablingStateResponse{}
	_body, _err := client.UpdateGatewayEnablingStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNodeGroupWithOptions(request *UpdateNodeGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupName)) {
		query["NodeGroupName"] = request.NodeGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNodeGroup"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNodeGroup(request *UpdateNodeGroupRequest) (_result *UpdateNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNodeGroupResponse{}
	_body, _err := client.UpdateNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNotificationsHandleStateWithOptions(request *UpdateNotificationsHandleStateRequest, runtime *util.RuntimeOptions) (_result *UpdateNotificationsHandleStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotificationId)) {
		query["NotificationId"] = request.NotificationId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetHandleState)) {
		query["TargetHandleState"] = request.TargetHandleState
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNotificationsHandleState"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNotificationsHandleStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNotificationsHandleState(request *UpdateNotificationsHandleStateRequest) (_result *UpdateNotificationsHandleStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNotificationsHandleStateResponse{}
	_body, _err := client.UpdateNotificationsHandleStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOwnedLocalJoinPermissionWithOptions(request *UpdateOwnedLocalJoinPermissionRequest, runtime *util.RuntimeOptions) (_result *UpdateOwnedLocalJoinPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassMode)) {
		query["ClassMode"] = request.ClassMode
	}

	if !tea.BoolValue(util.IsUnset(request.DataRate)) {
		query["DataRate"] = request.DataRate
	}

	if !tea.BoolValue(util.IsUnset(request.FreqBandPlanGroupId)) {
		query["FreqBandPlanGroupId"] = request.FreqBandPlanGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinEui)) {
		query["JoinEui"] = request.JoinEui
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionName)) {
		query["JoinPermissionName"] = request.JoinPermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.RxDelay)) {
		query["RxDelay"] = request.RxDelay
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOwnedLocalJoinPermission"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOwnedLocalJoinPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOwnedLocalJoinPermission(request *UpdateOwnedLocalJoinPermissionRequest) (_result *UpdateOwnedLocalJoinPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOwnedLocalJoinPermissionResponse{}
	_body, _err := client.UpdateOwnedLocalJoinPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOwnedLocalJoinPermissionEnablingStateWithOptions(request *UpdateOwnedLocalJoinPermissionEnablingStateRequest, runtime *util.RuntimeOptions) (_result *UpdateOwnedLocalJoinPermissionEnablingStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.IotInstanceId)) {
		query["IotInstanceId"] = request.IotInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOwnedLocalJoinPermissionEnablingState"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOwnedLocalJoinPermissionEnablingStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOwnedLocalJoinPermissionEnablingState(request *UpdateOwnedLocalJoinPermissionEnablingStateRequest) (_result *UpdateOwnedLocalJoinPermissionEnablingStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOwnedLocalJoinPermissionEnablingStateResponse{}
	_body, _err := client.UpdateOwnedLocalJoinPermissionEnablingStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoamingJoinPermissionWithOptions(request *UpdateRoamingJoinPermissionRequest, runtime *util.RuntimeOptions) (_result *UpdateRoamingJoinPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataRate)) {
		query["DataRate"] = request.DataRate
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionName)) {
		query["JoinPermissionName"] = request.JoinPermissionName
	}

	if !tea.BoolValue(util.IsUnset(request.RxDelay)) {
		query["RxDelay"] = request.RxDelay
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRoamingJoinPermission"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRoamingJoinPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRoamingJoinPermission(request *UpdateRoamingJoinPermissionRequest) (_result *UpdateRoamingJoinPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRoamingJoinPermissionResponse{}
	_body, _err := client.UpdateRoamingJoinPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoamingJoinPermissionEnablingStateWithOptions(request *UpdateRoamingJoinPermissionEnablingStateRequest, runtime *util.RuntimeOptions) (_result *UpdateRoamingJoinPermissionEnablingStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.JoinPermissionId)) {
		query["JoinPermissionId"] = request.JoinPermissionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRoamingJoinPermissionEnablingState"),
		Version:     tea.String("2019-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRoamingJoinPermissionEnablingStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRoamingJoinPermissionEnablingState(request *UpdateRoamingJoinPermissionEnablingStateRequest) (_result *UpdateRoamingJoinPermissionEnablingStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRoamingJoinPermissionEnablingStateResponse{}
	_body, _err := client.UpdateRoamingJoinPermissionEnablingStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
