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

type ChangeApplicationInfoRequest struct {
	AliUid      *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppTypeList *string `json:"AppTypeList,omitempty" xml:"AppTypeList,omitempty"`
	AppingList  *string `json:"AppingList,omitempty" xml:"AppingList,omitempty"`
	ItemCode    *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
}

func (s ChangeApplicationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplicationInfoRequest) GoString() string {
	return s.String()
}

func (s *ChangeApplicationInfoRequest) SetAliUid(v int64) *ChangeApplicationInfoRequest {
	s.AliUid = &v
	return s
}

func (s *ChangeApplicationInfoRequest) SetAppId(v string) *ChangeApplicationInfoRequest {
	s.AppId = &v
	return s
}

func (s *ChangeApplicationInfoRequest) SetAppName(v string) *ChangeApplicationInfoRequest {
	s.AppName = &v
	return s
}

func (s *ChangeApplicationInfoRequest) SetAppTypeList(v string) *ChangeApplicationInfoRequest {
	s.AppTypeList = &v
	return s
}

func (s *ChangeApplicationInfoRequest) SetAppingList(v string) *ChangeApplicationInfoRequest {
	s.AppingList = &v
	return s
}

func (s *ChangeApplicationInfoRequest) SetItemCode(v string) *ChangeApplicationInfoRequest {
	s.ItemCode = &v
	return s
}

type ChangeApplicationInfoResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeApplicationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplicationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeApplicationInfoResponseBody) SetAppId(v string) *ChangeApplicationInfoResponseBody {
	s.AppId = &v
	return s
}

func (s *ChangeApplicationInfoResponseBody) SetCode(v string) *ChangeApplicationInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeApplicationInfoResponseBody) SetMessage(v string) *ChangeApplicationInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeApplicationInfoResponseBody) SetRequestId(v string) *ChangeApplicationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeApplicationInfoResponseBody) SetSuccess(v bool) *ChangeApplicationInfoResponseBody {
	s.Success = &v
	return s
}

type ChangeApplicationInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeApplicationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeApplicationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplicationInfoResponse) GoString() string {
	return s.String()
}

func (s *ChangeApplicationInfoResponse) SetHeaders(v map[string]*string) *ChangeApplicationInfoResponse {
	s.Headers = v
	return s
}

func (s *ChangeApplicationInfoResponse) SetStatusCode(v int32) *ChangeApplicationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeApplicationInfoResponse) SetBody(v *ChangeApplicationInfoResponseBody) *ChangeApplicationInfoResponse {
	s.Body = v
	return s
}

type ChargeFlowRequest struct {
	ChannelCode *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemCode    *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OrderTime   *string `json:"OrderTime,omitempty" xml:"OrderTime,omitempty"`
	OutBizNo    *string `json:"OutBizNo,omitempty" xml:"OutBizNo,omitempty"`
	UId         *int64  `json:"UId,omitempty" xml:"UId,omitempty"`
}

func (s ChargeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ChargeFlowRequest) GoString() string {
	return s.String()
}

func (s *ChargeFlowRequest) SetChannelCode(v string) *ChargeFlowRequest {
	s.ChannelCode = &v
	return s
}

func (s *ChargeFlowRequest) SetInstanceId(v string) *ChargeFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *ChargeFlowRequest) SetItemCode(v string) *ChargeFlowRequest {
	s.ItemCode = &v
	return s
}

func (s *ChargeFlowRequest) SetMobile(v string) *ChargeFlowRequest {
	s.Mobile = &v
	return s
}

func (s *ChargeFlowRequest) SetOrderTime(v string) *ChargeFlowRequest {
	s.OrderTime = &v
	return s
}

func (s *ChargeFlowRequest) SetOutBizNo(v string) *ChargeFlowRequest {
	s.OutBizNo = &v
	return s
}

func (s *ChargeFlowRequest) SetUId(v int64) *ChargeFlowRequest {
	s.UId = &v
	return s
}

type ChargeFlowResponseBody struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ChargeFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                      `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChargeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChargeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ChargeFlowResponseBody) SetCode(v string) *ChargeFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ChargeFlowResponseBody) SetData(v *ChargeFlowResponseBodyData) *ChargeFlowResponseBody {
	s.Data = v
	return s
}

func (s *ChargeFlowResponseBody) SetMessage(v string) *ChargeFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ChargeFlowResponseBody) SetRequestId(v string) *ChargeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChargeFlowResponseBody) SetRt(v int64) *ChargeFlowResponseBody {
	s.Rt = &v
	return s
}

func (s *ChargeFlowResponseBody) SetSuccess(v bool) *ChargeFlowResponseBody {
	s.Success = &v
	return s
}

type ChargeFlowResponseBodyData struct {
	BizCode               *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	CustomerFlowOrderId   *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ChargeFlowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChargeFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChargeFlowResponseBodyData) SetBizCode(v string) *ChargeFlowResponseBodyData {
	s.BizCode = &v
	return s
}

func (s *ChargeFlowResponseBodyData) SetCustomerFlowOrderId(v string) *ChargeFlowResponseBodyData {
	s.CustomerFlowOrderId = &v
	return s
}

func (s *ChargeFlowResponseBodyData) SetCustomerFlowRequestId(v string) *ChargeFlowResponseBodyData {
	s.CustomerFlowRequestId = &v
	return s
}

func (s *ChargeFlowResponseBodyData) SetStatus(v string) *ChargeFlowResponseBodyData {
	s.Status = &v
	return s
}

type ChargeFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChargeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChargeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ChargeFlowResponse) GoString() string {
	return s.String()
}

func (s *ChargeFlowResponse) SetHeaders(v map[string]*string) *ChargeFlowResponse {
	s.Headers = v
	return s
}

func (s *ChargeFlowResponse) SetStatusCode(v int32) *ChargeFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ChargeFlowResponse) SetBody(v *ChargeFlowResponseBody) *ChargeFlowResponse {
	s.Body = v
	return s
}

type CreateApplicationInfoRequest struct {
	AliUid      *int64                                    `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppName     *string                                   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppTypeList []*string                                 `json:"AppTypeList,omitempty" xml:"AppTypeList,omitempty" type:"Repeated"`
	AppingList  []*CreateApplicationInfoRequestAppingList `json:"AppingList,omitempty" xml:"AppingList,omitempty" type:"Repeated"`
	ItemCode    *string                                   `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
}

func (s CreateApplicationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationInfoRequest) SetAliUid(v int64) *CreateApplicationInfoRequest {
	s.AliUid = &v
	return s
}

func (s *CreateApplicationInfoRequest) SetAppName(v string) *CreateApplicationInfoRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationInfoRequest) SetAppTypeList(v []*string) *CreateApplicationInfoRequest {
	s.AppTypeList = v
	return s
}

func (s *CreateApplicationInfoRequest) SetAppingList(v []*CreateApplicationInfoRequestAppingList) *CreateApplicationInfoRequest {
	s.AppingList = v
	return s
}

func (s *CreateApplicationInfoRequest) SetItemCode(v string) *CreateApplicationInfoRequest {
	s.ItemCode = &v
	return s
}

type CreateApplicationInfoRequestAppingList struct {
	ExtId           *int64    `json:"ExtId,omitempty" xml:"ExtId,omitempty"`
	FlowIp          []*string `json:"FlowIp,omitempty" xml:"FlowIp,omitempty" type:"Repeated"`
	FlowUrl         []*string `json:"FlowUrl,omitempty" xml:"FlowUrl,omitempty" type:"Repeated"`
	OriginalIpList  []*string `json:"OriginalIpList,omitempty" xml:"OriginalIpList,omitempty" type:"Repeated"`
	OriginalUrlList []*string `json:"OriginalUrlList,omitempty" xml:"OriginalUrlList,omitempty" type:"Repeated"`
}

func (s CreateApplicationInfoRequestAppingList) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationInfoRequestAppingList) GoString() string {
	return s.String()
}

func (s *CreateApplicationInfoRequestAppingList) SetExtId(v int64) *CreateApplicationInfoRequestAppingList {
	s.ExtId = &v
	return s
}

func (s *CreateApplicationInfoRequestAppingList) SetFlowIp(v []*string) *CreateApplicationInfoRequestAppingList {
	s.FlowIp = v
	return s
}

func (s *CreateApplicationInfoRequestAppingList) SetFlowUrl(v []*string) *CreateApplicationInfoRequestAppingList {
	s.FlowUrl = v
	return s
}

func (s *CreateApplicationInfoRequestAppingList) SetOriginalIpList(v []*string) *CreateApplicationInfoRequestAppingList {
	s.OriginalIpList = v
	return s
}

func (s *CreateApplicationInfoRequestAppingList) SetOriginalUrlList(v []*string) *CreateApplicationInfoRequestAppingList {
	s.OriginalUrlList = v
	return s
}

type CreateApplicationInfoResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateApplicationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationInfoResponseBody) SetAppId(v string) *CreateApplicationInfoResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateApplicationInfoResponseBody) SetCode(v string) *CreateApplicationInfoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationInfoResponseBody) SetMessage(v string) *CreateApplicationInfoResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationInfoResponseBody) SetRequestId(v string) *CreateApplicationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationInfoResponseBody) SetSuccess(v bool) *CreateApplicationInfoResponseBody {
	s.Success = &v
	return s
}

type CreateApplicationInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApplicationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationInfoResponse) SetHeaders(v map[string]*string) *CreateApplicationInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationInfoResponse) SetStatusCode(v int32) *CreateApplicationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationInfoResponse) SetBody(v *CreateApplicationInfoResponseBody) *CreateApplicationInfoResponse {
	s.Body = v
	return s
}

type GetAliyunXgipTokenResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64  `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAliyunXgipTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAliyunXgipTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAliyunXgipTokenResponseBody) SetCode(v string) *GetAliyunXgipTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetAliyunXgipTokenResponseBody) SetData(v string) *GetAliyunXgipTokenResponseBody {
	s.Data = &v
	return s
}

func (s *GetAliyunXgipTokenResponseBody) SetMessage(v string) *GetAliyunXgipTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetAliyunXgipTokenResponseBody) SetRequestId(v string) *GetAliyunXgipTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAliyunXgipTokenResponseBody) SetRt(v int64) *GetAliyunXgipTokenResponseBody {
	s.Rt = &v
	return s
}

func (s *GetAliyunXgipTokenResponseBody) SetSuccess(v bool) *GetAliyunXgipTokenResponseBody {
	s.Success = &v
	return s
}

type GetAliyunXgipTokenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAliyunXgipTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAliyunXgipTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAliyunXgipTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAliyunXgipTokenResponse) SetHeaders(v map[string]*string) *GetAliyunXgipTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAliyunXgipTokenResponse) SetStatusCode(v int32) *GetAliyunXgipTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAliyunXgipTokenResponse) SetBody(v *GetAliyunXgipTokenResponseBody) *GetAliyunXgipTokenResponse {
	s.Body = v
	return s
}

type GetApplicationRequest struct {
	AliUid   *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppCode  *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	ItemCode *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) SetAliUid(v int64) *GetApplicationRequest {
	s.AliUid = &v
	return s
}

func (s *GetApplicationRequest) SetAppCode(v string) *GetApplicationRequest {
	s.AppCode = &v
	return s
}

func (s *GetApplicationRequest) SetItemCode(v string) *GetApplicationRequest {
	s.ItemCode = &v
	return s
}

type GetApplicationResponseBody struct {
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                   `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) SetCode(v string) *GetApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationResponseBody) SetData(v []map[string]interface{}) *GetApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationResponseBody) SetMessage(v string) *GetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) SetRt(v int64) *GetApplicationResponseBody {
	s.Rt = &v
	return s
}

func (s *GetApplicationResponseBody) SetSuccess(v bool) *GetApplicationResponseBody {
	s.Success = &v
	return s
}

type GetApplicationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) SetHeaders(v map[string]*string) *GetApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationResponse) SetStatusCode(v int32) *GetApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationResponse) SetBody(v *GetApplicationResponseBody) *GetApplicationResponse {
	s.Body = v
	return s
}

type GetFreeFlowInstanceRequest struct {
	Aliuid     *int64  `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemCode   *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
}

func (s GetFreeFlowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetFreeFlowInstanceRequest) SetAliuid(v int64) *GetFreeFlowInstanceRequest {
	s.Aliuid = &v
	return s
}

func (s *GetFreeFlowInstanceRequest) SetAppId(v string) *GetFreeFlowInstanceRequest {
	s.AppId = &v
	return s
}

func (s *GetFreeFlowInstanceRequest) SetInstanceId(v string) *GetFreeFlowInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFreeFlowInstanceRequest) SetItemCode(v string) *GetFreeFlowInstanceRequest {
	s.ItemCode = &v
	return s
}

type GetFreeFlowInstanceResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetFreeFlowInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                                 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFreeFlowInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFreeFlowInstanceResponseBody) SetCode(v string) *GetFreeFlowInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBody) SetData(v []*GetFreeFlowInstanceResponseBodyData) *GetFreeFlowInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetFreeFlowInstanceResponseBody) SetMessage(v string) *GetFreeFlowInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBody) SetRequestId(v string) *GetFreeFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBody) SetRt(v int64) *GetFreeFlowInstanceResponseBody {
	s.Rt = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBody) SetSuccess(v bool) *GetFreeFlowInstanceResponseBody {
	s.Success = &v
	return s
}

type GetFreeFlowInstanceResponseBodyData struct {
	AppCode        *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceMemo   *string `json:"InstanceMemo,omitempty" xml:"InstanceMemo,omitempty"`
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	OpenTime       *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	SpecType       *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetFreeFlowInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFreeFlowInstanceResponseBodyData) SetAppCode(v string) *GetFreeFlowInstanceResponseBodyData {
	s.AppCode = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBodyData) SetAppName(v string) *GetFreeFlowInstanceResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBodyData) SetEndTime(v string) *GetFreeFlowInstanceResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBodyData) SetInstanceMemo(v string) *GetFreeFlowInstanceResponseBodyData {
	s.InstanceMemo = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBodyData) SetInstanceStatus(v string) *GetFreeFlowInstanceResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBodyData) SetOpenTime(v string) *GetFreeFlowInstanceResponseBodyData {
	s.OpenTime = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBodyData) SetSpecType(v string) *GetFreeFlowInstanceResponseBodyData {
	s.SpecType = &v
	return s
}

func (s *GetFreeFlowInstanceResponseBodyData) SetStartTime(v string) *GetFreeFlowInstanceResponseBodyData {
	s.StartTime = &v
	return s
}

type GetFreeFlowInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFreeFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFreeFlowInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetFreeFlowInstanceResponse) SetHeaders(v map[string]*string) *GetFreeFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetFreeFlowInstanceResponse) SetStatusCode(v int32) *GetFreeFlowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFreeFlowInstanceResponse) SetBody(v *GetFreeFlowInstanceResponseBody) *GetFreeFlowInstanceResponse {
	s.Body = v
	return s
}

type GetFreeFlowProductListRequest struct {
	AliUid     *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetFreeFlowProductListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowProductListRequest) GoString() string {
	return s.String()
}

func (s *GetFreeFlowProductListRequest) SetAliUid(v int64) *GetFreeFlowProductListRequest {
	s.AliUid = &v
	return s
}

func (s *GetFreeFlowProductListRequest) SetInstanceId(v string) *GetFreeFlowProductListRequest {
	s.InstanceId = &v
	return s
}

type GetFreeFlowProductListResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetFreeFlowProductListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                                    `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFreeFlowProductListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowProductListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFreeFlowProductListResponseBody) SetCode(v string) *GetFreeFlowProductListResponseBody {
	s.Code = &v
	return s
}

func (s *GetFreeFlowProductListResponseBody) SetData(v []*GetFreeFlowProductListResponseBodyData) *GetFreeFlowProductListResponseBody {
	s.Data = v
	return s
}

func (s *GetFreeFlowProductListResponseBody) SetMessage(v string) *GetFreeFlowProductListResponseBody {
	s.Message = &v
	return s
}

func (s *GetFreeFlowProductListResponseBody) SetRequestId(v string) *GetFreeFlowProductListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFreeFlowProductListResponseBody) SetRt(v int64) *GetFreeFlowProductListResponseBody {
	s.Rt = &v
	return s
}

func (s *GetFreeFlowProductListResponseBody) SetSuccess(v bool) *GetFreeFlowProductListResponseBody {
	s.Success = &v
	return s
}

type GetFreeFlowProductListResponseBodyData struct {
	Configured        *bool   `json:"Configured,omitempty" xml:"Configured,omitempty"`
	FlowProductAmount *string `json:"FlowProductAmount,omitempty" xml:"FlowProductAmount,omitempty"`
	FlowProductId     *string `json:"FlowProductId,omitempty" xml:"FlowProductId,omitempty"`
	FlowProductName   *string `json:"FlowProductName,omitempty" xml:"FlowProductName,omitempty"`
	FlowProductPeriod *string `json:"FlowProductPeriod,omitempty" xml:"FlowProductPeriod,omitempty"`
	FlowType          *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	Operator          *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	SpecType          *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Spid              *string `json:"Spid,omitempty" xml:"Spid,omitempty"`
	UnitPrice         *int32  `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
}

func (s GetFreeFlowProductListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowProductListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFreeFlowProductListResponseBodyData) SetConfigured(v bool) *GetFreeFlowProductListResponseBodyData {
	s.Configured = &v
	return s
}

func (s *GetFreeFlowProductListResponseBodyData) SetFlowProductAmount(v string) *GetFreeFlowProductListResponseBodyData {
	s.FlowProductAmount = &v
	return s
}

func (s *GetFreeFlowProductListResponseBodyData) SetFlowProductId(v string) *GetFreeFlowProductListResponseBodyData {
	s.FlowProductId = &v
	return s
}

func (s *GetFreeFlowProductListResponseBodyData) SetFlowProductName(v string) *GetFreeFlowProductListResponseBodyData {
	s.FlowProductName = &v
	return s
}

func (s *GetFreeFlowProductListResponseBodyData) SetFlowProductPeriod(v string) *GetFreeFlowProductListResponseBodyData {
	s.FlowProductPeriod = &v
	return s
}

func (s *GetFreeFlowProductListResponseBodyData) SetFlowType(v string) *GetFreeFlowProductListResponseBodyData {
	s.FlowType = &v
	return s
}

func (s *GetFreeFlowProductListResponseBodyData) SetOperator(v string) *GetFreeFlowProductListResponseBodyData {
	s.Operator = &v
	return s
}

func (s *GetFreeFlowProductListResponseBodyData) SetSpecType(v string) *GetFreeFlowProductListResponseBodyData {
	s.SpecType = &v
	return s
}

func (s *GetFreeFlowProductListResponseBodyData) SetSpid(v string) *GetFreeFlowProductListResponseBodyData {
	s.Spid = &v
	return s
}

func (s *GetFreeFlowProductListResponseBodyData) SetUnitPrice(v int32) *GetFreeFlowProductListResponseBodyData {
	s.UnitPrice = &v
	return s
}

type GetFreeFlowProductListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFreeFlowProductListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFreeFlowProductListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowProductListResponse) GoString() string {
	return s.String()
}

func (s *GetFreeFlowProductListResponse) SetHeaders(v map[string]*string) *GetFreeFlowProductListResponse {
	s.Headers = v
	return s
}

func (s *GetFreeFlowProductListResponse) SetStatusCode(v int32) *GetFreeFlowProductListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFreeFlowProductListResponse) SetBody(v *GetFreeFlowProductListResponseBody) *GetFreeFlowProductListResponse {
	s.Body = v
	return s
}

type GetFreeFlowUsageRequest struct {
	AliUid     *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CurPageNum *int32  `json:"CurPageNum,omitempty" xml:"CurPageNum,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Month      *string `json:"Month,omitempty" xml:"Month,omitempty"`
	NumPerPage *int32  `json:"NumPerPage,omitempty" xml:"NumPerPage,omitempty"`
}

func (s GetFreeFlowUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowUsageRequest) GoString() string {
	return s.String()
}

func (s *GetFreeFlowUsageRequest) SetAliUid(v int64) *GetFreeFlowUsageRequest {
	s.AliUid = &v
	return s
}

func (s *GetFreeFlowUsageRequest) SetCurPageNum(v int32) *GetFreeFlowUsageRequest {
	s.CurPageNum = &v
	return s
}

func (s *GetFreeFlowUsageRequest) SetInstanceId(v string) *GetFreeFlowUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFreeFlowUsageRequest) SetMonth(v string) *GetFreeFlowUsageRequest {
	s.Month = &v
	return s
}

func (s *GetFreeFlowUsageRequest) SetNumPerPage(v int32) *GetFreeFlowUsageRequest {
	s.NumPerPage = &v
	return s
}

type GetFreeFlowUsageResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetFreeFlowUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                            `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFreeFlowUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetFreeFlowUsageResponseBody) SetCode(v string) *GetFreeFlowUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetFreeFlowUsageResponseBody) SetData(v *GetFreeFlowUsageResponseBodyData) *GetFreeFlowUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetFreeFlowUsageResponseBody) SetMessage(v string) *GetFreeFlowUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetFreeFlowUsageResponseBody) SetRequestId(v string) *GetFreeFlowUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFreeFlowUsageResponseBody) SetRt(v int64) *GetFreeFlowUsageResponseBody {
	s.Rt = &v
	return s
}

func (s *GetFreeFlowUsageResponseBody) SetSuccess(v bool) *GetFreeFlowUsageResponseBody {
	s.Success = &v
	return s
}

type GetFreeFlowUsageResponseBodyData struct {
	CurPageNum   *int32                                          `json:"CurPageNum,omitempty" xml:"CurPageNum,omitempty"`
	CustomerList []*GetFreeFlowUsageResponseBodyDataCustomerList `json:"CustomerList,omitempty" xml:"CustomerList,omitempty" type:"Repeated"`
	HasNext      *bool                                           `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	HasPrev      *bool                                           `json:"HasPrev,omitempty" xml:"HasPrev,omitempty"`
	InstanceId   *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NumPerPage   *int32                                          `json:"NumPerPage,omitempty" xml:"NumPerPage,omitempty"`
	TotalNum     *int32                                          `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPageNum *int32                                          `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s GetFreeFlowUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFreeFlowUsageResponseBodyData) SetCurPageNum(v int32) *GetFreeFlowUsageResponseBodyData {
	s.CurPageNum = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyData) SetCustomerList(v []*GetFreeFlowUsageResponseBodyDataCustomerList) *GetFreeFlowUsageResponseBodyData {
	s.CustomerList = v
	return s
}

func (s *GetFreeFlowUsageResponseBodyData) SetHasNext(v bool) *GetFreeFlowUsageResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyData) SetHasPrev(v bool) *GetFreeFlowUsageResponseBodyData {
	s.HasPrev = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyData) SetInstanceId(v string) *GetFreeFlowUsageResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyData) SetNumPerPage(v int32) *GetFreeFlowUsageResponseBodyData {
	s.NumPerPage = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyData) SetTotalNum(v int32) *GetFreeFlowUsageResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyData) SetTotalPageNum(v int32) *GetFreeFlowUsageResponseBodyData {
	s.TotalPageNum = &v
	return s
}

type GetFreeFlowUsageResponseBodyDataCustomerList struct {
	ChannelId           *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CustomerEndTime     *string `json:"CustomerEndTime,omitempty" xml:"CustomerEndTime,omitempty"`
	CustomerFlowOrderId *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
	CustomerFlowStatus  *string `json:"CustomerFlowStatus,omitempty" xml:"CustomerFlowStatus,omitempty"`
	CustomerOpenTime    *string `json:"CustomerOpenTime,omitempty" xml:"CustomerOpenTime,omitempty"`
	CustomerStartTime   *string `json:"CustomerStartTime,omitempty" xml:"CustomerStartTime,omitempty"`
	FlowProductId       *string `json:"FlowProductId,omitempty" xml:"FlowProductId,omitempty"`
	FlowProductName     *string `json:"FlowProductName,omitempty" xml:"FlowProductName,omitempty"`
	IsLasting           *bool   `json:"IsLasting,omitempty" xml:"IsLasting,omitempty"`
	MobileNumber        *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	UnitNum             *int32  `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
	UnitPrice           *int32  `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
}

func (s GetFreeFlowUsageResponseBodyDataCustomerList) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowUsageResponseBodyDataCustomerList) GoString() string {
	return s.String()
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetChannelId(v string) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.ChannelId = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetCustomerEndTime(v string) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.CustomerEndTime = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetCustomerFlowOrderId(v string) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.CustomerFlowOrderId = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetCustomerFlowStatus(v string) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.CustomerFlowStatus = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetCustomerOpenTime(v string) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.CustomerOpenTime = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetCustomerStartTime(v string) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.CustomerStartTime = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetFlowProductId(v string) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.FlowProductId = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetFlowProductName(v string) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.FlowProductName = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetIsLasting(v bool) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.IsLasting = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetMobileNumber(v string) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.MobileNumber = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetUnitNum(v int32) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.UnitNum = &v
	return s
}

func (s *GetFreeFlowUsageResponseBodyDataCustomerList) SetUnitPrice(v int32) *GetFreeFlowUsageResponseBodyDataCustomerList {
	s.UnitPrice = &v
	return s
}

type GetFreeFlowUsageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFreeFlowUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFreeFlowUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowUsageResponse) GoString() string {
	return s.String()
}

func (s *GetFreeFlowUsageResponse) SetHeaders(v map[string]*string) *GetFreeFlowUsageResponse {
	s.Headers = v
	return s
}

func (s *GetFreeFlowUsageResponse) SetStatusCode(v int32) *GetFreeFlowUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFreeFlowUsageResponse) SetBody(v *GetFreeFlowUsageResponseBody) *GetFreeFlowUsageResponse {
	s.Body = v
	return s
}

type GetFreeFlowUsageStatisticRequest struct {
	AliUid     *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Month      *int64  `json:"Month,omitempty" xml:"Month,omitempty"`
}

func (s GetFreeFlowUsageStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowUsageStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetFreeFlowUsageStatisticRequest) SetAliUid(v int64) *GetFreeFlowUsageStatisticRequest {
	s.AliUid = &v
	return s
}

func (s *GetFreeFlowUsageStatisticRequest) SetAppId(v string) *GetFreeFlowUsageStatisticRequest {
	s.AppId = &v
	return s
}

func (s *GetFreeFlowUsageStatisticRequest) SetAppName(v string) *GetFreeFlowUsageStatisticRequest {
	s.AppName = &v
	return s
}

func (s *GetFreeFlowUsageStatisticRequest) SetInstanceId(v string) *GetFreeFlowUsageStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFreeFlowUsageStatisticRequest) SetMonth(v int64) *GetFreeFlowUsageStatisticRequest {
	s.Month = &v
	return s
}

type GetFreeFlowUsageStatisticResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetFreeFlowUsageStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                                       `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFreeFlowUsageStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowUsageStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetFreeFlowUsageStatisticResponseBody) SetCode(v string) *GetFreeFlowUsageStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBody) SetData(v []*GetFreeFlowUsageStatisticResponseBodyData) *GetFreeFlowUsageStatisticResponseBody {
	s.Data = v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBody) SetMessage(v string) *GetFreeFlowUsageStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBody) SetRequestId(v string) *GetFreeFlowUsageStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBody) SetRt(v int64) *GetFreeFlowUsageStatisticResponseBody {
	s.Rt = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBody) SetSuccess(v bool) *GetFreeFlowUsageStatisticResponseBody {
	s.Success = &v
	return s
}

type GetFreeFlowUsageStatisticResponseBodyData struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SpecType         *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TotalMoney       *string `json:"TotalMoney,omitempty" xml:"TotalMoney,omitempty"`
	TotalOrderNumber *int64  `json:"TotalOrderNumber,omitempty" xml:"TotalOrderNumber,omitempty"`
	TotalUnitNumber  *int64  `json:"TotalUnitNumber,omitempty" xml:"TotalUnitNumber,omitempty"`
	YunOutProduct    *string `json:"YunOutProduct,omitempty" xml:"YunOutProduct,omitempty"`
}

func (s GetFreeFlowUsageStatisticResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowUsageStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFreeFlowUsageStatisticResponseBodyData) SetInstanceId(v string) *GetFreeFlowUsageStatisticResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBodyData) SetSpecType(v string) *GetFreeFlowUsageStatisticResponseBodyData {
	s.SpecType = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBodyData) SetTotalMoney(v string) *GetFreeFlowUsageStatisticResponseBodyData {
	s.TotalMoney = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBodyData) SetTotalOrderNumber(v int64) *GetFreeFlowUsageStatisticResponseBodyData {
	s.TotalOrderNumber = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBodyData) SetTotalUnitNumber(v int64) *GetFreeFlowUsageStatisticResponseBodyData {
	s.TotalUnitNumber = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponseBodyData) SetYunOutProduct(v string) *GetFreeFlowUsageStatisticResponseBodyData {
	s.YunOutProduct = &v
	return s
}

type GetFreeFlowUsageStatisticResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFreeFlowUsageStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFreeFlowUsageStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowUsageStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetFreeFlowUsageStatisticResponse) SetHeaders(v map[string]*string) *GetFreeFlowUsageStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetFreeFlowUsageStatisticResponse) SetStatusCode(v int32) *GetFreeFlowUsageStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFreeFlowUsageStatisticResponse) SetBody(v *GetFreeFlowUsageStatisticResponseBody) *GetFreeFlowUsageStatisticResponse {
	s.Body = v
	return s
}

type GetInventoryInfoRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemCode   *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	Mobile     *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UId        *int64  `json:"UId,omitempty" xml:"UId,omitempty"`
}

func (s GetInventoryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryInfoRequest) GoString() string {
	return s.String()
}

func (s *GetInventoryInfoRequest) SetInstanceId(v string) *GetInventoryInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInventoryInfoRequest) SetItemCode(v string) *GetInventoryInfoRequest {
	s.ItemCode = &v
	return s
}

func (s *GetInventoryInfoRequest) SetMobile(v string) *GetInventoryInfoRequest {
	s.Mobile = &v
	return s
}

func (s *GetInventoryInfoRequest) SetUId(v int64) *GetInventoryInfoRequest {
	s.UId = &v
	return s
}

type GetInventoryInfoResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetInventoryInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                            `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInventoryInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetInventoryInfoResponseBody) SetCode(v string) *GetInventoryInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetInventoryInfoResponseBody) SetData(v *GetInventoryInfoResponseBodyData) *GetInventoryInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetInventoryInfoResponseBody) SetMessage(v string) *GetInventoryInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetInventoryInfoResponseBody) SetRequestId(v string) *GetInventoryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInventoryInfoResponseBody) SetRt(v int64) *GetInventoryInfoResponseBody {
	s.Rt = &v
	return s
}

func (s *GetInventoryInfoResponseBody) SetSuccess(v bool) *GetInventoryInfoResponseBody {
	s.Success = &v
	return s
}

type GetInventoryInfoResponseBodyData struct {
	Inventory         *int64 `json:"Inventory,omitempty" xml:"Inventory,omitempty"`
	ResidualInventory *int64 `json:"ResidualInventory,omitempty" xml:"ResidualInventory,omitempty"`
	UsedStock         *int64 `json:"UsedStock,omitempty" xml:"UsedStock,omitempty"`
}

func (s GetInventoryInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInventoryInfoResponseBodyData) SetInventory(v int64) *GetInventoryInfoResponseBodyData {
	s.Inventory = &v
	return s
}

func (s *GetInventoryInfoResponseBodyData) SetResidualInventory(v int64) *GetInventoryInfoResponseBodyData {
	s.ResidualInventory = &v
	return s
}

func (s *GetInventoryInfoResponseBodyData) SetUsedStock(v int64) *GetInventoryInfoResponseBodyData {
	s.UsedStock = &v
	return s
}

type GetInventoryInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInventoryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInventoryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInventoryInfoResponse) GoString() string {
	return s.String()
}

func (s *GetInventoryInfoResponse) SetHeaders(v map[string]*string) *GetInventoryInfoResponse {
	s.Headers = v
	return s
}

func (s *GetInventoryInfoResponse) SetStatusCode(v int32) *GetInventoryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInventoryInfoResponse) SetBody(v *GetInventoryInfoResponseBody) *GetInventoryInfoResponse {
	s.Body = v
	return s
}

type GetItemInstListRequest struct {
	Current    *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemCode   *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	Mobile     *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UId        *int64  `json:"UId,omitempty" xml:"UId,omitempty"`
}

func (s GetItemInstListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetItemInstListRequest) GoString() string {
	return s.String()
}

func (s *GetItemInstListRequest) SetCurrent(v int32) *GetItemInstListRequest {
	s.Current = &v
	return s
}

func (s *GetItemInstListRequest) SetInstanceId(v string) *GetItemInstListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetItemInstListRequest) SetItemCode(v string) *GetItemInstListRequest {
	s.ItemCode = &v
	return s
}

func (s *GetItemInstListRequest) SetMobile(v string) *GetItemInstListRequest {
	s.Mobile = &v
	return s
}

func (s *GetItemInstListRequest) SetPageSize(v int32) *GetItemInstListRequest {
	s.PageSize = &v
	return s
}

func (s *GetItemInstListRequest) SetStatus(v string) *GetItemInstListRequest {
	s.Status = &v
	return s
}

func (s *GetItemInstListRequest) SetUId(v int64) *GetItemInstListRequest {
	s.UId = &v
	return s
}

type GetItemInstListResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetItemInstListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                           `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetItemInstListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetItemInstListResponseBody) GoString() string {
	return s.String()
}

func (s *GetItemInstListResponseBody) SetCode(v string) *GetItemInstListResponseBody {
	s.Code = &v
	return s
}

func (s *GetItemInstListResponseBody) SetData(v *GetItemInstListResponseBodyData) *GetItemInstListResponseBody {
	s.Data = v
	return s
}

func (s *GetItemInstListResponseBody) SetMessage(v string) *GetItemInstListResponseBody {
	s.Message = &v
	return s
}

func (s *GetItemInstListResponseBody) SetRequestId(v string) *GetItemInstListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetItemInstListResponseBody) SetRt(v int64) *GetItemInstListResponseBody {
	s.Rt = &v
	return s
}

func (s *GetItemInstListResponseBody) SetSuccess(v bool) *GetItemInstListResponseBody {
	s.Success = &v
	return s
}

type GetItemInstListResponseBodyData struct {
	List     []*GetItemInstListResponseBodyDataList   `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageInfo *GetItemInstListResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s GetItemInstListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetItemInstListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetItemInstListResponseBodyData) SetList(v []*GetItemInstListResponseBodyDataList) *GetItemInstListResponseBodyData {
	s.List = v
	return s
}

func (s *GetItemInstListResponseBodyData) SetPageInfo(v *GetItemInstListResponseBodyDataPageInfo) *GetItemInstListResponseBodyData {
	s.PageInfo = v
	return s
}

type GetItemInstListResponseBodyDataList struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProductId   *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetItemInstListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetItemInstListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetItemInstListResponseBodyDataList) SetCreateTime(v string) *GetItemInstListResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *GetItemInstListResponseBodyDataList) SetExpireTime(v string) *GetItemInstListResponseBodyDataList {
	s.ExpireTime = &v
	return s
}

func (s *GetItemInstListResponseBodyDataList) SetInstanceId(v string) *GetItemInstListResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *GetItemInstListResponseBodyDataList) SetProductId(v string) *GetItemInstListResponseBodyDataList {
	s.ProductId = &v
	return s
}

func (s *GetItemInstListResponseBodyDataList) SetProductName(v string) *GetItemInstListResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *GetItemInstListResponseBodyDataList) SetStatus(v int32) *GetItemInstListResponseBodyDataList {
	s.Status = &v
	return s
}

type GetItemInstListResponseBodyDataPageInfo struct {
	Current  *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetItemInstListResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetItemInstListResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *GetItemInstListResponseBodyDataPageInfo) SetCurrent(v int32) *GetItemInstListResponseBodyDataPageInfo {
	s.Current = &v
	return s
}

func (s *GetItemInstListResponseBodyDataPageInfo) SetPageSize(v int32) *GetItemInstListResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetItemInstListResponseBodyDataPageInfo) SetTotal(v int64) *GetItemInstListResponseBodyDataPageInfo {
	s.Total = &v
	return s
}

type GetItemInstListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetItemInstListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetItemInstListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetItemInstListResponse) GoString() string {
	return s.String()
}

func (s *GetItemInstListResponse) SetHeaders(v map[string]*string) *GetItemInstListResponse {
	s.Headers = v
	return s
}

func (s *GetItemInstListResponse) SetStatusCode(v int32) *GetItemInstListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetItemInstListResponse) SetBody(v *GetItemInstListResponseBody) *GetItemInstListResponse {
	s.Body = v
	return s
}

type GetItemListRequest struct {
	AliUid  *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
}

func (s GetItemListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetItemListRequest) GoString() string {
	return s.String()
}

func (s *GetItemListRequest) SetAliUid(v int64) *GetItemListRequest {
	s.AliUid = &v
	return s
}

func (s *GetItemListRequest) SetBizCode(v string) *GetItemListRequest {
	s.BizCode = &v
	return s
}

type GetItemListResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetItemListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                         `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetItemListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetItemListResponseBody) GoString() string {
	return s.String()
}

func (s *GetItemListResponseBody) SetCode(v string) *GetItemListResponseBody {
	s.Code = &v
	return s
}

func (s *GetItemListResponseBody) SetData(v []*GetItemListResponseBodyData) *GetItemListResponseBody {
	s.Data = v
	return s
}

func (s *GetItemListResponseBody) SetMessage(v string) *GetItemListResponseBody {
	s.Message = &v
	return s
}

func (s *GetItemListResponseBody) SetRequestId(v string) *GetItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetItemListResponseBody) SetRt(v int64) *GetItemListResponseBody {
	s.Rt = &v
	return s
}

func (s *GetItemListResponseBody) SetSuccess(v bool) *GetItemListResponseBody {
	s.Success = &v
	return s
}

type GetItemListResponseBodyData struct {
	BizCode    *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	BizType    *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ItemBuyUrl *string `json:"ItemBuyUrl,omitempty" xml:"ItemBuyUrl,omitempty"`
	ItemCode   *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	ItemName   *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetItemListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetItemListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetItemListResponseBodyData) SetBizCode(v string) *GetItemListResponseBodyData {
	s.BizCode = &v
	return s
}

func (s *GetItemListResponseBodyData) SetBizType(v string) *GetItemListResponseBodyData {
	s.BizType = &v
	return s
}

func (s *GetItemListResponseBodyData) SetItemBuyUrl(v string) *GetItemListResponseBodyData {
	s.ItemBuyUrl = &v
	return s
}

func (s *GetItemListResponseBodyData) SetItemCode(v string) *GetItemListResponseBodyData {
	s.ItemCode = &v
	return s
}

func (s *GetItemListResponseBodyData) SetItemName(v string) *GetItemListResponseBodyData {
	s.ItemName = &v
	return s
}

func (s *GetItemListResponseBodyData) SetName(v string) *GetItemListResponseBodyData {
	s.Name = &v
	return s
}

type GetItemListResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetItemListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetItemListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetItemListResponse) GoString() string {
	return s.String()
}

func (s *GetItemListResponse) SetHeaders(v map[string]*string) *GetItemListResponse {
	s.Headers = v
	return s
}

func (s *GetItemListResponse) SetStatusCode(v int32) *GetItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetItemListResponse) SetBody(v *GetItemListResponseBody) *GetItemListResponse {
	s.Body = v
	return s
}

type GetOrderFreeFlowProductStatusRequest struct {
	AliUid              *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CustomerFlowOrderId *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
}

func (s GetOrderFreeFlowProductStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrderFreeFlowProductStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOrderFreeFlowProductStatusRequest) SetAliUid(v int64) *GetOrderFreeFlowProductStatusRequest {
	s.AliUid = &v
	return s
}

func (s *GetOrderFreeFlowProductStatusRequest) SetCustomerFlowOrderId(v string) *GetOrderFreeFlowProductStatusRequest {
	s.CustomerFlowOrderId = &v
	return s
}

type GetOrderFreeFlowProductStatusResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetOrderFreeFlowProductStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                                         `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrderFreeFlowProductStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrderFreeFlowProductStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderFreeFlowProductStatusResponseBody) SetCode(v string) *GetOrderFreeFlowProductStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponseBody) SetData(v *GetOrderFreeFlowProductStatusResponseBodyData) *GetOrderFreeFlowProductStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponseBody) SetMessage(v string) *GetOrderFreeFlowProductStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponseBody) SetRequestId(v string) *GetOrderFreeFlowProductStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponseBody) SetRt(v int64) *GetOrderFreeFlowProductStatusResponseBody {
	s.Rt = &v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponseBody) SetSuccess(v bool) *GetOrderFreeFlowProductStatusResponseBody {
	s.Success = &v
	return s
}

type GetOrderFreeFlowProductStatusResponseBodyData struct {
	CustomerFlowOrderId   *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	Error                 *string `json:"Error,omitempty" xml:"Error,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOrderFreeFlowProductStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOrderFreeFlowProductStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrderFreeFlowProductStatusResponseBodyData) SetCustomerFlowOrderId(v string) *GetOrderFreeFlowProductStatusResponseBodyData {
	s.CustomerFlowOrderId = &v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponseBodyData) SetCustomerFlowRequestId(v string) *GetOrderFreeFlowProductStatusResponseBodyData {
	s.CustomerFlowRequestId = &v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponseBodyData) SetError(v string) *GetOrderFreeFlowProductStatusResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponseBodyData) SetStatus(v string) *GetOrderFreeFlowProductStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetOrderFreeFlowProductStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrderFreeFlowProductStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrderFreeFlowProductStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrderFreeFlowProductStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOrderFreeFlowProductStatusResponse) SetHeaders(v map[string]*string) *GetOrderFreeFlowProductStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponse) SetStatusCode(v int32) *GetOrderFreeFlowProductStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderFreeFlowProductStatusResponse) SetBody(v *GetOrderFreeFlowProductStatusResponseBody) *GetOrderFreeFlowProductStatusResponse {
	s.Body = v
	return s
}

type GetOrderItemListResponseBody struct {
	Code      *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64      `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrderItemListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrderItemListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderItemListResponseBody) SetCode(v string) *GetOrderItemListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrderItemListResponseBody) SetData(v interface{}) *GetOrderItemListResponseBody {
	s.Data = v
	return s
}

func (s *GetOrderItemListResponseBody) SetMessage(v string) *GetOrderItemListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrderItemListResponseBody) SetRequestId(v string) *GetOrderItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderItemListResponseBody) SetRt(v int64) *GetOrderItemListResponseBody {
	s.Rt = &v
	return s
}

func (s *GetOrderItemListResponseBody) SetSuccess(v bool) *GetOrderItemListResponseBody {
	s.Success = &v
	return s
}

type GetOrderItemListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrderItemListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrderItemListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrderItemListResponse) GoString() string {
	return s.String()
}

func (s *GetOrderItemListResponse) SetHeaders(v map[string]*string) *GetOrderItemListResponse {
	s.Headers = v
	return s
}

func (s *GetOrderItemListResponse) SetStatusCode(v int32) *GetOrderItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderItemListResponse) SetBody(v *GetOrderItemListResponseBody) *GetOrderItemListResponse {
	s.Body = v
	return s
}

type GetQosFlowUsageRequest struct {
	AliUid     *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CurPageNum *int32  `json:"CurPageNum,omitempty" xml:"CurPageNum,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Month      *int64  `json:"Month,omitempty" xml:"Month,omitempty"`
	NumPerPage *int32  `json:"NumPerPage,omitempty" xml:"NumPerPage,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type       *bool   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQosFlowUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQosFlowUsageRequest) GoString() string {
	return s.String()
}

func (s *GetQosFlowUsageRequest) SetAliUid(v int64) *GetQosFlowUsageRequest {
	s.AliUid = &v
	return s
}

func (s *GetQosFlowUsageRequest) SetCurPageNum(v int32) *GetQosFlowUsageRequest {
	s.CurPageNum = &v
	return s
}

func (s *GetQosFlowUsageRequest) SetEndTime(v string) *GetQosFlowUsageRequest {
	s.EndTime = &v
	return s
}

func (s *GetQosFlowUsageRequest) SetInstanceId(v string) *GetQosFlowUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQosFlowUsageRequest) SetMonth(v int64) *GetQosFlowUsageRequest {
	s.Month = &v
	return s
}

func (s *GetQosFlowUsageRequest) SetNumPerPage(v int32) *GetQosFlowUsageRequest {
	s.NumPerPage = &v
	return s
}

func (s *GetQosFlowUsageRequest) SetStartTime(v string) *GetQosFlowUsageRequest {
	s.StartTime = &v
	return s
}

func (s *GetQosFlowUsageRequest) SetType(v bool) *GetQosFlowUsageRequest {
	s.Type = &v
	return s
}

type GetQosFlowUsageResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQosFlowUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                           `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQosFlowUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQosFlowUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetQosFlowUsageResponseBody) SetCode(v string) *GetQosFlowUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetQosFlowUsageResponseBody) SetData(v *GetQosFlowUsageResponseBodyData) *GetQosFlowUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetQosFlowUsageResponseBody) SetMessage(v string) *GetQosFlowUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetQosFlowUsageResponseBody) SetRequestId(v string) *GetQosFlowUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQosFlowUsageResponseBody) SetRt(v int64) *GetQosFlowUsageResponseBody {
	s.Rt = &v
	return s
}

func (s *GetQosFlowUsageResponseBody) SetSuccess(v bool) *GetQosFlowUsageResponseBody {
	s.Success = &v
	return s
}

type GetQosFlowUsageResponseBodyData struct {
	CurPageNum   *int32                                         `json:"CurPageNum,omitempty" xml:"CurPageNum,omitempty"`
	CustomerList []*GetQosFlowUsageResponseBodyDataCustomerList `json:"CustomerList,omitempty" xml:"CustomerList,omitempty" type:"Repeated"`
	HasNext      *bool                                          `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	HasPrev      *bool                                          `json:"HasPrev,omitempty" xml:"HasPrev,omitempty"`
	InstanceId   *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NumPerPage   *int32                                         `json:"NumPerPage,omitempty" xml:"NumPerPage,omitempty"`
	TotalNum     *int32                                         `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPageNum *int32                                         `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s GetQosFlowUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQosFlowUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQosFlowUsageResponseBodyData) SetCurPageNum(v int32) *GetQosFlowUsageResponseBodyData {
	s.CurPageNum = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyData) SetCustomerList(v []*GetQosFlowUsageResponseBodyDataCustomerList) *GetQosFlowUsageResponseBodyData {
	s.CustomerList = v
	return s
}

func (s *GetQosFlowUsageResponseBodyData) SetHasNext(v bool) *GetQosFlowUsageResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyData) SetHasPrev(v bool) *GetQosFlowUsageResponseBodyData {
	s.HasPrev = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyData) SetInstanceId(v string) *GetQosFlowUsageResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyData) SetNumPerPage(v int32) *GetQosFlowUsageResponseBodyData {
	s.NumPerPage = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyData) SetTotalNum(v int32) *GetQosFlowUsageResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyData) SetTotalPageNum(v int32) *GetQosFlowUsageResponseBodyData {
	s.TotalPageNum = &v
	return s
}

type GetQosFlowUsageResponseBodyDataCustomerList struct {
	AliUid       *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DsDay        *int64  `json:"DsDay,omitempty" xml:"DsDay,omitempty"`
	DsMonth      *int64  `json:"DsMonth,omitempty" xml:"DsMonth,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemCode     *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Operator     *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OrderNum     *int32  `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	ProductId    *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName  *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Provice      *string `json:"Provice,omitempty" xml:"Provice,omitempty"`
	QosRequestId *string `json:"QosRequestId,omitempty" xml:"QosRequestId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SpecType     *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalPrice   *int32  `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	TotolTime    *int64  `json:"TotolTime,omitempty" xml:"TotolTime,omitempty"`
}

func (s GetQosFlowUsageResponseBodyDataCustomerList) String() string {
	return tea.Prettify(s)
}

func (s GetQosFlowUsageResponseBodyDataCustomerList) GoString() string {
	return s.String()
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetAliUid(v int64) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.AliUid = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetAppId(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.AppId = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetDsDay(v int64) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.DsDay = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetDsMonth(v int64) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.DsMonth = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetEndTime(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.EndTime = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetInstanceId(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.InstanceId = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetItemCode(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.ItemCode = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetMessage(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.Message = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetName(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.Name = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetOperator(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.Operator = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetOrderNum(v int32) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.OrderNum = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetProductId(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.ProductId = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetProductName(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.ProductName = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetProvice(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.Provice = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetQosRequestId(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.QosRequestId = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetRemark(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.Remark = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetSpecType(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.SpecType = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetStartTime(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.StartTime = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetStatus(v string) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.Status = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetTotalPrice(v int32) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.TotalPrice = &v
	return s
}

func (s *GetQosFlowUsageResponseBodyDataCustomerList) SetTotolTime(v int64) *GetQosFlowUsageResponseBodyDataCustomerList {
	s.TotolTime = &v
	return s
}

type GetQosFlowUsageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQosFlowUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQosFlowUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQosFlowUsageResponse) GoString() string {
	return s.String()
}

func (s *GetQosFlowUsageResponse) SetHeaders(v map[string]*string) *GetQosFlowUsageResponse {
	s.Headers = v
	return s
}

func (s *GetQosFlowUsageResponse) SetStatusCode(v int32) *GetQosFlowUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQosFlowUsageResponse) SetBody(v *GetQosFlowUsageResponseBody) *GetQosFlowUsageResponse {
	s.Body = v
	return s
}

type GetQosUsageStatisticRequest struct {
	AliUid     *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CurPageNum *int32  `json:"CurPageNum,omitempty" xml:"CurPageNum,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Month      *int64  `json:"Month,omitempty" xml:"Month,omitempty"`
	NumPerPage *int32  `json:"NumPerPage,omitempty" xml:"NumPerPage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetQosUsageStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQosUsageStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetQosUsageStatisticRequest) SetAliUid(v int64) *GetQosUsageStatisticRequest {
	s.AliUid = &v
	return s
}

func (s *GetQosUsageStatisticRequest) SetCurPageNum(v int32) *GetQosUsageStatisticRequest {
	s.CurPageNum = &v
	return s
}

func (s *GetQosUsageStatisticRequest) SetEndTime(v string) *GetQosUsageStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *GetQosUsageStatisticRequest) SetInstanceId(v string) *GetQosUsageStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQosUsageStatisticRequest) SetMonth(v int64) *GetQosUsageStatisticRequest {
	s.Month = &v
	return s
}

func (s *GetQosUsageStatisticRequest) SetNumPerPage(v int32) *GetQosUsageStatisticRequest {
	s.NumPerPage = &v
	return s
}

func (s *GetQosUsageStatisticRequest) SetRequestId(v string) *GetQosUsageStatisticRequest {
	s.RequestId = &v
	return s
}

func (s *GetQosUsageStatisticRequest) SetStartTime(v string) *GetQosUsageStatisticRequest {
	s.StartTime = &v
	return s
}

type GetQosUsageStatisticResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetQosUsageStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                                `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQosUsageStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQosUsageStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetQosUsageStatisticResponseBody) SetCode(v string) *GetQosUsageStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *GetQosUsageStatisticResponseBody) SetData(v *GetQosUsageStatisticResponseBodyData) *GetQosUsageStatisticResponseBody {
	s.Data = v
	return s
}

func (s *GetQosUsageStatisticResponseBody) SetMessage(v string) *GetQosUsageStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *GetQosUsageStatisticResponseBody) SetRequestId(v string) *GetQosUsageStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQosUsageStatisticResponseBody) SetRt(v int64) *GetQosUsageStatisticResponseBody {
	s.Rt = &v
	return s
}

func (s *GetQosUsageStatisticResponseBody) SetSuccess(v bool) *GetQosUsageStatisticResponseBody {
	s.Success = &v
	return s
}

type GetQosUsageStatisticResponseBodyData struct {
	CurPageNum                  *int32                                                             `json:"CurPageNum,omitempty" xml:"CurPageNum,omitempty"`
	GetQosUsageStatisticResList []*GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList `json:"GetQosUsageStatisticResList,omitempty" xml:"GetQosUsageStatisticResList,omitempty" type:"Repeated"`
	NumPerPage                  *int32                                                             `json:"NumPerPage,omitempty" xml:"NumPerPage,omitempty"`
	TotalNum                    *int32                                                             `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetQosUsageStatisticResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQosUsageStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQosUsageStatisticResponseBodyData) SetCurPageNum(v int32) *GetQosUsageStatisticResponseBodyData {
	s.CurPageNum = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyData) SetGetQosUsageStatisticResList(v []*GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) *GetQosUsageStatisticResponseBodyData {
	s.GetQosUsageStatisticResList = v
	return s
}

func (s *GetQosUsageStatisticResponseBodyData) SetNumPerPage(v int32) *GetQosUsageStatisticResponseBodyData {
	s.NumPerPage = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyData) SetTotalNum(v int32) *GetQosUsageStatisticResponseBodyData {
	s.TotalNum = &v
	return s
}

type GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList struct {
	BillCount     *int64  `json:"BillCount,omitempty" xml:"BillCount,omitempty"`
	DsDay         *int64  `json:"DsDay,omitempty" xml:"DsDay,omitempty"`
	FailCount     *int64  `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Month         *int64  `json:"Month,omitempty" xml:"Month,omitempty"`
	Operator      *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	ProductName   *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Provice       *string `json:"Provice,omitempty" xml:"Provice,omitempty"`
	SpecType      *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	SucessCount   *int64  `json:"SucessCount,omitempty" xml:"SucessCount,omitempty"`
	TotalCount    *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	YunOutProduct *string `json:"YunOutProduct,omitempty" xml:"YunOutProduct,omitempty"`
}

func (s GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) String() string {
	return tea.Prettify(s)
}

func (s GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) GoString() string {
	return s.String()
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetBillCount(v int64) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.BillCount = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetDsDay(v int64) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.DsDay = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetFailCount(v int64) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.FailCount = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetInstanceId(v string) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.InstanceId = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetMonth(v int64) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.Month = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetOperator(v string) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.Operator = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetProductName(v string) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.ProductName = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetProvice(v string) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.Provice = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetSpecType(v string) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.SpecType = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetSucessCount(v int64) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.SucessCount = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetTotalCount(v int64) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.TotalCount = &v
	return s
}

func (s *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList) SetYunOutProduct(v string) *GetQosUsageStatisticResponseBodyDataGetQosUsageStatisticResList {
	s.YunOutProduct = &v
	return s
}

type GetQosUsageStatisticResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQosUsageStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQosUsageStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQosUsageStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetQosUsageStatisticResponse) SetHeaders(v map[string]*string) *GetQosUsageStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetQosUsageStatisticResponse) SetStatusCode(v int32) *GetQosUsageStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQosUsageStatisticResponse) SetBody(v *GetQosUsageStatisticResponseBody) *GetQosUsageStatisticResponse {
	s.Body = v
	return s
}

type GetUatDataListRequest struct {
	AliUid  *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DsMonth *int64  `json:"DsMonth,omitempty" xml:"DsMonth,omitempty"`
	ItemId  *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s GetUatDataListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUatDataListRequest) GoString() string {
	return s.String()
}

func (s *GetUatDataListRequest) SetAliUid(v int64) *GetUatDataListRequest {
	s.AliUid = &v
	return s
}

func (s *GetUatDataListRequest) SetDsMonth(v int64) *GetUatDataListRequest {
	s.DsMonth = &v
	return s
}

func (s *GetUatDataListRequest) SetItemId(v string) *GetUatDataListRequest {
	s.ItemId = &v
	return s
}

type GetUatDataListResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetUatDataListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                            `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUatDataListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUatDataListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUatDataListResponseBody) SetCode(v string) *GetUatDataListResponseBody {
	s.Code = &v
	return s
}

func (s *GetUatDataListResponseBody) SetData(v []*GetUatDataListResponseBodyData) *GetUatDataListResponseBody {
	s.Data = v
	return s
}

func (s *GetUatDataListResponseBody) SetMessage(v string) *GetUatDataListResponseBody {
	s.Message = &v
	return s
}

func (s *GetUatDataListResponseBody) SetRequestId(v string) *GetUatDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUatDataListResponseBody) SetRt(v int64) *GetUatDataListResponseBody {
	s.Rt = &v
	return s
}

func (s *GetUatDataListResponseBody) SetSuccess(v bool) *GetUatDataListResponseBody {
	s.Success = &v
	return s
}

type GetUatDataListResponseBodyData struct {
	DsList   []*GetUatDataListResponseBodyDataDsList `json:"DsList,omitempty" xml:"DsList,omitempty" type:"Repeated"`
	SpecType *string                                 `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
}

func (s GetUatDataListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUatDataListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUatDataListResponseBodyData) SetDsList(v []*GetUatDataListResponseBodyDataDsList) *GetUatDataListResponseBodyData {
	s.DsList = v
	return s
}

func (s *GetUatDataListResponseBodyData) SetSpecType(v string) *GetUatDataListResponseBodyData {
	s.SpecType = &v
	return s
}

type GetUatDataListResponseBodyDataDsList struct {
	DsData *int64 `json:"DsData,omitempty" xml:"DsData,omitempty"`
	DsDay  *int64 `json:"DsDay,omitempty" xml:"DsDay,omitempty"`
}

func (s GetUatDataListResponseBodyDataDsList) String() string {
	return tea.Prettify(s)
}

func (s GetUatDataListResponseBodyDataDsList) GoString() string {
	return s.String()
}

func (s *GetUatDataListResponseBodyDataDsList) SetDsData(v int64) *GetUatDataListResponseBodyDataDsList {
	s.DsData = &v
	return s
}

func (s *GetUatDataListResponseBodyDataDsList) SetDsDay(v int64) *GetUatDataListResponseBodyDataDsList {
	s.DsDay = &v
	return s
}

type GetUatDataListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUatDataListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUatDataListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUatDataListResponse) GoString() string {
	return s.String()
}

func (s *GetUatDataListResponse) SetHeaders(v map[string]*string) *GetUatDataListResponse {
	s.Headers = v
	return s
}

func (s *GetUatDataListResponse) SetStatusCode(v int32) *GetUatDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUatDataListResponse) SetBody(v *GetUatDataListResponseBody) *GetUatDataListResponse {
	s.Body = v
	return s
}

type GetUatSpecCtDataRequest struct {
	AliUid  *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DsMonth *int64  `json:"DsMonth,omitempty" xml:"DsMonth,omitempty"`
	ItemId  *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s GetUatSpecCtDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUatSpecCtDataRequest) GoString() string {
	return s.String()
}

func (s *GetUatSpecCtDataRequest) SetAliUid(v int64) *GetUatSpecCtDataRequest {
	s.AliUid = &v
	return s
}

func (s *GetUatSpecCtDataRequest) SetDsMonth(v int64) *GetUatSpecCtDataRequest {
	s.DsMonth = &v
	return s
}

func (s *GetUatSpecCtDataRequest) SetItemId(v string) *GetUatSpecCtDataRequest {
	s.ItemId = &v
	return s
}

type GetUatSpecCtDataResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetUatSpecCtDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                              `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUatSpecCtDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUatSpecCtDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetUatSpecCtDataResponseBody) SetCode(v string) *GetUatSpecCtDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetUatSpecCtDataResponseBody) SetData(v []*GetUatSpecCtDataResponseBodyData) *GetUatSpecCtDataResponseBody {
	s.Data = v
	return s
}

func (s *GetUatSpecCtDataResponseBody) SetMessage(v string) *GetUatSpecCtDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetUatSpecCtDataResponseBody) SetRequestId(v string) *GetUatSpecCtDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUatSpecCtDataResponseBody) SetRt(v int64) *GetUatSpecCtDataResponseBody {
	s.Rt = &v
	return s
}

func (s *GetUatSpecCtDataResponseBody) SetSuccess(v bool) *GetUatSpecCtDataResponseBody {
	s.Success = &v
	return s
}

type GetUatSpecCtDataResponseBodyData struct {
	MonthCount *int64  `json:"MonthCount,omitempty" xml:"MonthCount,omitempty"`
	SpecType   *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
}

func (s GetUatSpecCtDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUatSpecCtDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUatSpecCtDataResponseBodyData) SetMonthCount(v int64) *GetUatSpecCtDataResponseBodyData {
	s.MonthCount = &v
	return s
}

func (s *GetUatSpecCtDataResponseBodyData) SetSpecType(v string) *GetUatSpecCtDataResponseBodyData {
	s.SpecType = &v
	return s
}

type GetUatSpecCtDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUatSpecCtDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUatSpecCtDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUatSpecCtDataResponse) GoString() string {
	return s.String()
}

func (s *GetUatSpecCtDataResponse) SetHeaders(v map[string]*string) *GetUatSpecCtDataResponse {
	s.Headers = v
	return s
}

func (s *GetUatSpecCtDataResponse) SetStatusCode(v int32) *GetUatSpecCtDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUatSpecCtDataResponse) SetBody(v *GetUatSpecCtDataResponseBody) *GetUatSpecCtDataResponse {
	s.Body = v
	return s
}

type MockGetOrderFreeFlowProductStatusRequest struct {
	AliUid              *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CustomerFlowOrderId *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
}

func (s MockGetOrderFreeFlowProductStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s MockGetOrderFreeFlowProductStatusRequest) GoString() string {
	return s.String()
}

func (s *MockGetOrderFreeFlowProductStatusRequest) SetAliUid(v int64) *MockGetOrderFreeFlowProductStatusRequest {
	s.AliUid = &v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusRequest) SetCustomerFlowOrderId(v string) *MockGetOrderFreeFlowProductStatusRequest {
	s.CustomerFlowOrderId = &v
	return s
}

type MockGetOrderFreeFlowProductStatusResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MockGetOrderFreeFlowProductStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                                             `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MockGetOrderFreeFlowProductStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MockGetOrderFreeFlowProductStatusResponseBody) GoString() string {
	return s.String()
}

func (s *MockGetOrderFreeFlowProductStatusResponseBody) SetCode(v string) *MockGetOrderFreeFlowProductStatusResponseBody {
	s.Code = &v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponseBody) SetData(v *MockGetOrderFreeFlowProductStatusResponseBodyData) *MockGetOrderFreeFlowProductStatusResponseBody {
	s.Data = v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponseBody) SetMessage(v string) *MockGetOrderFreeFlowProductStatusResponseBody {
	s.Message = &v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponseBody) SetRequestId(v string) *MockGetOrderFreeFlowProductStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponseBody) SetRt(v int64) *MockGetOrderFreeFlowProductStatusResponseBody {
	s.Rt = &v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponseBody) SetSuccess(v bool) *MockGetOrderFreeFlowProductStatusResponseBody {
	s.Success = &v
	return s
}

type MockGetOrderFreeFlowProductStatusResponseBodyData struct {
	CustomerFlowOrderId   *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	Error                 *string `json:"Error,omitempty" xml:"Error,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s MockGetOrderFreeFlowProductStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MockGetOrderFreeFlowProductStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *MockGetOrderFreeFlowProductStatusResponseBodyData) SetCustomerFlowOrderId(v string) *MockGetOrderFreeFlowProductStatusResponseBodyData {
	s.CustomerFlowOrderId = &v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponseBodyData) SetCustomerFlowRequestId(v string) *MockGetOrderFreeFlowProductStatusResponseBodyData {
	s.CustomerFlowRequestId = &v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponseBodyData) SetError(v string) *MockGetOrderFreeFlowProductStatusResponseBodyData {
	s.Error = &v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponseBodyData) SetStatus(v string) *MockGetOrderFreeFlowProductStatusResponseBodyData {
	s.Status = &v
	return s
}

type MockGetOrderFreeFlowProductStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MockGetOrderFreeFlowProductStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MockGetOrderFreeFlowProductStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s MockGetOrderFreeFlowProductStatusResponse) GoString() string {
	return s.String()
}

func (s *MockGetOrderFreeFlowProductStatusResponse) SetHeaders(v map[string]*string) *MockGetOrderFreeFlowProductStatusResponse {
	s.Headers = v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponse) SetStatusCode(v int32) *MockGetOrderFreeFlowProductStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *MockGetOrderFreeFlowProductStatusResponse) SetBody(v *MockGetOrderFreeFlowProductStatusResponseBody) *MockGetOrderFreeFlowProductStatusResponse {
	s.Body = v
	return s
}

type MockOrderFreeFlowProductRequest struct {
	AliUid                *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ChannelId             *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	FlowProductId         *string `json:"FlowProductId,omitempty" xml:"FlowProductId,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lasting               *string `json:"Lasting,omitempty" xml:"Lasting,omitempty"`
	MobileNumber          *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	NotifyUrl             *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	Operator              *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	PurchaseOrderId       *string `json:"PurchaseOrderId,omitempty" xml:"PurchaseOrderId,omitempty"`
}

func (s MockOrderFreeFlowProductRequest) String() string {
	return tea.Prettify(s)
}

func (s MockOrderFreeFlowProductRequest) GoString() string {
	return s.String()
}

func (s *MockOrderFreeFlowProductRequest) SetAliUid(v int64) *MockOrderFreeFlowProductRequest {
	s.AliUid = &v
	return s
}

func (s *MockOrderFreeFlowProductRequest) SetChannelId(v string) *MockOrderFreeFlowProductRequest {
	s.ChannelId = &v
	return s
}

func (s *MockOrderFreeFlowProductRequest) SetCustomerFlowRequestId(v string) *MockOrderFreeFlowProductRequest {
	s.CustomerFlowRequestId = &v
	return s
}

func (s *MockOrderFreeFlowProductRequest) SetFlowProductId(v string) *MockOrderFreeFlowProductRequest {
	s.FlowProductId = &v
	return s
}

func (s *MockOrderFreeFlowProductRequest) SetInstanceId(v string) *MockOrderFreeFlowProductRequest {
	s.InstanceId = &v
	return s
}

func (s *MockOrderFreeFlowProductRequest) SetLasting(v string) *MockOrderFreeFlowProductRequest {
	s.Lasting = &v
	return s
}

func (s *MockOrderFreeFlowProductRequest) SetMobileNumber(v string) *MockOrderFreeFlowProductRequest {
	s.MobileNumber = &v
	return s
}

func (s *MockOrderFreeFlowProductRequest) SetNotifyUrl(v string) *MockOrderFreeFlowProductRequest {
	s.NotifyUrl = &v
	return s
}

func (s *MockOrderFreeFlowProductRequest) SetOperator(v string) *MockOrderFreeFlowProductRequest {
	s.Operator = &v
	return s
}

func (s *MockOrderFreeFlowProductRequest) SetPurchaseOrderId(v string) *MockOrderFreeFlowProductRequest {
	s.PurchaseOrderId = &v
	return s
}

type MockOrderFreeFlowProductResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MockOrderFreeFlowProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                                    `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MockOrderFreeFlowProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MockOrderFreeFlowProductResponseBody) GoString() string {
	return s.String()
}

func (s *MockOrderFreeFlowProductResponseBody) SetCode(v string) *MockOrderFreeFlowProductResponseBody {
	s.Code = &v
	return s
}

func (s *MockOrderFreeFlowProductResponseBody) SetData(v *MockOrderFreeFlowProductResponseBodyData) *MockOrderFreeFlowProductResponseBody {
	s.Data = v
	return s
}

func (s *MockOrderFreeFlowProductResponseBody) SetMessage(v string) *MockOrderFreeFlowProductResponseBody {
	s.Message = &v
	return s
}

func (s *MockOrderFreeFlowProductResponseBody) SetRequestId(v string) *MockOrderFreeFlowProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *MockOrderFreeFlowProductResponseBody) SetRt(v int64) *MockOrderFreeFlowProductResponseBody {
	s.Rt = &v
	return s
}

func (s *MockOrderFreeFlowProductResponseBody) SetSuccess(v bool) *MockOrderFreeFlowProductResponseBody {
	s.Success = &v
	return s
}

type MockOrderFreeFlowProductResponseBodyData struct {
	BizCode               *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	CustomerFlowOrderId   *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s MockOrderFreeFlowProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MockOrderFreeFlowProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *MockOrderFreeFlowProductResponseBodyData) SetBizCode(v string) *MockOrderFreeFlowProductResponseBodyData {
	s.BizCode = &v
	return s
}

func (s *MockOrderFreeFlowProductResponseBodyData) SetCustomerFlowOrderId(v string) *MockOrderFreeFlowProductResponseBodyData {
	s.CustomerFlowOrderId = &v
	return s
}

func (s *MockOrderFreeFlowProductResponseBodyData) SetCustomerFlowRequestId(v string) *MockOrderFreeFlowProductResponseBodyData {
	s.CustomerFlowRequestId = &v
	return s
}

func (s *MockOrderFreeFlowProductResponseBodyData) SetStatus(v string) *MockOrderFreeFlowProductResponseBodyData {
	s.Status = &v
	return s
}

type MockOrderFreeFlowProductResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MockOrderFreeFlowProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MockOrderFreeFlowProductResponse) String() string {
	return tea.Prettify(s)
}

func (s MockOrderFreeFlowProductResponse) GoString() string {
	return s.String()
}

func (s *MockOrderFreeFlowProductResponse) SetHeaders(v map[string]*string) *MockOrderFreeFlowProductResponse {
	s.Headers = v
	return s
}

func (s *MockOrderFreeFlowProductResponse) SetStatusCode(v int32) *MockOrderFreeFlowProductResponse {
	s.StatusCode = &v
	return s
}

func (s *MockOrderFreeFlowProductResponse) SetBody(v *MockOrderFreeFlowProductResponseBody) *MockOrderFreeFlowProductResponse {
	s.Body = v
	return s
}

type ModifyApplicationRequest struct {
	AliUid      *int64                                `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppCode     *string                               `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName     *string                               `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppTypeList []*string                             `json:"AppTypeList,omitempty" xml:"AppTypeList,omitempty" type:"Repeated"`
	AppingList  []*ModifyApplicationRequestAppingList `json:"AppingList,omitempty" xml:"AppingList,omitempty" type:"Repeated"`
}

func (s ModifyApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApplicationRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationRequest) SetAliUid(v int64) *ModifyApplicationRequest {
	s.AliUid = &v
	return s
}

func (s *ModifyApplicationRequest) SetAppCode(v string) *ModifyApplicationRequest {
	s.AppCode = &v
	return s
}

func (s *ModifyApplicationRequest) SetAppName(v string) *ModifyApplicationRequest {
	s.AppName = &v
	return s
}

func (s *ModifyApplicationRequest) SetAppTypeList(v []*string) *ModifyApplicationRequest {
	s.AppTypeList = v
	return s
}

func (s *ModifyApplicationRequest) SetAppingList(v []*ModifyApplicationRequestAppingList) *ModifyApplicationRequest {
	s.AppingList = v
	return s
}

type ModifyApplicationRequestAppingList struct {
	ExtId           *int64    `json:"ExtId,omitempty" xml:"ExtId,omitempty"`
	FlowIp          []*string `json:"FlowIp,omitempty" xml:"FlowIp,omitempty" type:"Repeated"`
	FlowUrl         []*string `json:"FlowUrl,omitempty" xml:"FlowUrl,omitempty" type:"Repeated"`
	OriginalIpList  []*string `json:"OriginalIpList,omitempty" xml:"OriginalIpList,omitempty" type:"Repeated"`
	OriginalUrlList []*string `json:"OriginalUrlList,omitempty" xml:"OriginalUrlList,omitempty" type:"Repeated"`
}

func (s ModifyApplicationRequestAppingList) String() string {
	return tea.Prettify(s)
}

func (s ModifyApplicationRequestAppingList) GoString() string {
	return s.String()
}

func (s *ModifyApplicationRequestAppingList) SetExtId(v int64) *ModifyApplicationRequestAppingList {
	s.ExtId = &v
	return s
}

func (s *ModifyApplicationRequestAppingList) SetFlowIp(v []*string) *ModifyApplicationRequestAppingList {
	s.FlowIp = v
	return s
}

func (s *ModifyApplicationRequestAppingList) SetFlowUrl(v []*string) *ModifyApplicationRequestAppingList {
	s.FlowUrl = v
	return s
}

func (s *ModifyApplicationRequestAppingList) SetOriginalIpList(v []*string) *ModifyApplicationRequestAppingList {
	s.OriginalIpList = v
	return s
}

func (s *ModifyApplicationRequestAppingList) SetOriginalUrlList(v []*string) *ModifyApplicationRequestAppingList {
	s.OriginalUrlList = v
	return s
}

type ModifyApplicationResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationResponseBody) SetAppId(v string) *ModifyApplicationResponseBody {
	s.AppId = &v
	return s
}

func (s *ModifyApplicationResponseBody) SetCode(v string) *ModifyApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyApplicationResponseBody) SetMessage(v string) *ModifyApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyApplicationResponseBody) SetRequestId(v string) *ModifyApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationResponseBody) SetSuccess(v bool) *ModifyApplicationResponseBody {
	s.Success = &v
	return s
}

type ModifyApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyApplicationResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationResponse) SetHeaders(v map[string]*string) *ModifyApplicationResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationResponse) SetStatusCode(v int32) *ModifyApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationResponse) SetBody(v *ModifyApplicationResponseBody) *ModifyApplicationResponse {
	s.Body = v
	return s
}

type OrderFreeFlowProductRequest struct {
	AliUid                *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ChannelId             *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	FlowProductId         *string `json:"FlowProductId,omitempty" xml:"FlowProductId,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lasting               *string `json:"Lasting,omitempty" xml:"Lasting,omitempty"`
	MobileNumber          *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	NotifyUrl             *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	Operator              *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	PurchaseOrderId       *string `json:"PurchaseOrderId,omitempty" xml:"PurchaseOrderId,omitempty"`
}

func (s OrderFreeFlowProductRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderFreeFlowProductRequest) GoString() string {
	return s.String()
}

func (s *OrderFreeFlowProductRequest) SetAliUid(v int64) *OrderFreeFlowProductRequest {
	s.AliUid = &v
	return s
}

func (s *OrderFreeFlowProductRequest) SetChannelId(v string) *OrderFreeFlowProductRequest {
	s.ChannelId = &v
	return s
}

func (s *OrderFreeFlowProductRequest) SetCustomerFlowRequestId(v string) *OrderFreeFlowProductRequest {
	s.CustomerFlowRequestId = &v
	return s
}

func (s *OrderFreeFlowProductRequest) SetFlowProductId(v string) *OrderFreeFlowProductRequest {
	s.FlowProductId = &v
	return s
}

func (s *OrderFreeFlowProductRequest) SetInstanceId(v string) *OrderFreeFlowProductRequest {
	s.InstanceId = &v
	return s
}

func (s *OrderFreeFlowProductRequest) SetLasting(v string) *OrderFreeFlowProductRequest {
	s.Lasting = &v
	return s
}

func (s *OrderFreeFlowProductRequest) SetMobileNumber(v string) *OrderFreeFlowProductRequest {
	s.MobileNumber = &v
	return s
}

func (s *OrderFreeFlowProductRequest) SetNotifyUrl(v string) *OrderFreeFlowProductRequest {
	s.NotifyUrl = &v
	return s
}

func (s *OrderFreeFlowProductRequest) SetOperator(v string) *OrderFreeFlowProductRequest {
	s.Operator = &v
	return s
}

func (s *OrderFreeFlowProductRequest) SetPurchaseOrderId(v string) *OrderFreeFlowProductRequest {
	s.PurchaseOrderId = &v
	return s
}

type OrderFreeFlowProductResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *OrderFreeFlowProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                                `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OrderFreeFlowProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrderFreeFlowProductResponseBody) GoString() string {
	return s.String()
}

func (s *OrderFreeFlowProductResponseBody) SetCode(v string) *OrderFreeFlowProductResponseBody {
	s.Code = &v
	return s
}

func (s *OrderFreeFlowProductResponseBody) SetData(v *OrderFreeFlowProductResponseBodyData) *OrderFreeFlowProductResponseBody {
	s.Data = v
	return s
}

func (s *OrderFreeFlowProductResponseBody) SetMessage(v string) *OrderFreeFlowProductResponseBody {
	s.Message = &v
	return s
}

func (s *OrderFreeFlowProductResponseBody) SetRequestId(v string) *OrderFreeFlowProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *OrderFreeFlowProductResponseBody) SetRt(v int64) *OrderFreeFlowProductResponseBody {
	s.Rt = &v
	return s
}

func (s *OrderFreeFlowProductResponseBody) SetSuccess(v bool) *OrderFreeFlowProductResponseBody {
	s.Success = &v
	return s
}

type OrderFreeFlowProductResponseBodyData struct {
	BizCode               *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	CustomerFlowOrderId   *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s OrderFreeFlowProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OrderFreeFlowProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *OrderFreeFlowProductResponseBodyData) SetBizCode(v string) *OrderFreeFlowProductResponseBodyData {
	s.BizCode = &v
	return s
}

func (s *OrderFreeFlowProductResponseBodyData) SetCustomerFlowOrderId(v string) *OrderFreeFlowProductResponseBodyData {
	s.CustomerFlowOrderId = &v
	return s
}

func (s *OrderFreeFlowProductResponseBodyData) SetCustomerFlowRequestId(v string) *OrderFreeFlowProductResponseBodyData {
	s.CustomerFlowRequestId = &v
	return s
}

func (s *OrderFreeFlowProductResponseBodyData) SetStatus(v string) *OrderFreeFlowProductResponseBodyData {
	s.Status = &v
	return s
}

type OrderFreeFlowProductResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OrderFreeFlowProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OrderFreeFlowProductResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderFreeFlowProductResponse) GoString() string {
	return s.String()
}

func (s *OrderFreeFlowProductResponse) SetHeaders(v map[string]*string) *OrderFreeFlowProductResponse {
	s.Headers = v
	return s
}

func (s *OrderFreeFlowProductResponse) SetStatusCode(v int32) *OrderFreeFlowProductResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderFreeFlowProductResponse) SetBody(v *OrderFreeFlowProductResponseBody) *OrderFreeFlowProductResponse {
	s.Body = v
	return s
}

type OrderQosProductRequest struct {
	AliUid       *int64    `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ChannelId    *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	IPv6         *string   `json:"IPv6,omitempty" xml:"IPv6,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpType       *string   `json:"IpType,omitempty" xml:"IpType,omitempty"`
	MobileNumber *string   `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	Operator     *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	PrivateIpv4  *string   `json:"PrivateIpv4,omitempty" xml:"PrivateIpv4,omitempty"`
	ProductId    *string   `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Provice      *string   `json:"Provice,omitempty" xml:"Provice,omitempty"`
	PublicIpv4   *string   `json:"PublicIpv4,omitempty" xml:"PublicIpv4,omitempty"`
	QosRequestId *string   `json:"QosRequestId,omitempty" xml:"QosRequestId,omitempty"`
	TargetIpList []*string `json:"TargetIpList,omitempty" xml:"TargetIpList,omitempty" type:"Repeated"`
	Token        *string   `json:"Token,omitempty" xml:"Token,omitempty"`
	UnitNum      *int32    `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s OrderQosProductRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderQosProductRequest) GoString() string {
	return s.String()
}

func (s *OrderQosProductRequest) SetAliUid(v int64) *OrderQosProductRequest {
	s.AliUid = &v
	return s
}

func (s *OrderQosProductRequest) SetChannelId(v string) *OrderQosProductRequest {
	s.ChannelId = &v
	return s
}

func (s *OrderQosProductRequest) SetIPv6(v string) *OrderQosProductRequest {
	s.IPv6 = &v
	return s
}

func (s *OrderQosProductRequest) SetInstanceId(v string) *OrderQosProductRequest {
	s.InstanceId = &v
	return s
}

func (s *OrderQosProductRequest) SetIpType(v string) *OrderQosProductRequest {
	s.IpType = &v
	return s
}

func (s *OrderQosProductRequest) SetMobileNumber(v string) *OrderQosProductRequest {
	s.MobileNumber = &v
	return s
}

func (s *OrderQosProductRequest) SetOperator(v string) *OrderQosProductRequest {
	s.Operator = &v
	return s
}

func (s *OrderQosProductRequest) SetPrivateIpv4(v string) *OrderQosProductRequest {
	s.PrivateIpv4 = &v
	return s
}

func (s *OrderQosProductRequest) SetProductId(v string) *OrderQosProductRequest {
	s.ProductId = &v
	return s
}

func (s *OrderQosProductRequest) SetProvice(v string) *OrderQosProductRequest {
	s.Provice = &v
	return s
}

func (s *OrderQosProductRequest) SetPublicIpv4(v string) *OrderQosProductRequest {
	s.PublicIpv4 = &v
	return s
}

func (s *OrderQosProductRequest) SetQosRequestId(v string) *OrderQosProductRequest {
	s.QosRequestId = &v
	return s
}

func (s *OrderQosProductRequest) SetTargetIpList(v []*string) *OrderQosProductRequest {
	s.TargetIpList = v
	return s
}

func (s *OrderQosProductRequest) SetToken(v string) *OrderQosProductRequest {
	s.Token = &v
	return s
}

func (s *OrderQosProductRequest) SetUnitNum(v int32) *OrderQosProductRequest {
	s.UnitNum = &v
	return s
}

type OrderQosProductResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64  `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OrderQosProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrderQosProductResponseBody) GoString() string {
	return s.String()
}

func (s *OrderQosProductResponseBody) SetCode(v string) *OrderQosProductResponseBody {
	s.Code = &v
	return s
}

func (s *OrderQosProductResponseBody) SetData(v string) *OrderQosProductResponseBody {
	s.Data = &v
	return s
}

func (s *OrderQosProductResponseBody) SetMessage(v string) *OrderQosProductResponseBody {
	s.Message = &v
	return s
}

func (s *OrderQosProductResponseBody) SetRequestId(v string) *OrderQosProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *OrderQosProductResponseBody) SetRt(v int64) *OrderQosProductResponseBody {
	s.Rt = &v
	return s
}

func (s *OrderQosProductResponseBody) SetSuccess(v bool) *OrderQosProductResponseBody {
	s.Success = &v
	return s
}

type OrderQosProductResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OrderQosProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OrderQosProductResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderQosProductResponse) GoString() string {
	return s.String()
}

func (s *OrderQosProductResponse) SetHeaders(v map[string]*string) *OrderQosProductResponse {
	s.Headers = v
	return s
}

func (s *OrderQosProductResponse) SetStatusCode(v int32) *OrderQosProductResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderQosProductResponse) SetBody(v *OrderQosProductResponseBody) *OrderQosProductResponse {
	s.Body = v
	return s
}

type QueryOrderListRequest struct {
	Current    *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemCode   *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	Mobile     *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OrderTime  *string `json:"OrderTime,omitempty" xml:"OrderTime,omitempty"`
	OutBizNo   *string `json:"OutBizNo,omitempty" xml:"OutBizNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UId        *int64  `json:"UId,omitempty" xml:"UId,omitempty"`
}

func (s QueryOrderListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderListRequest) SetCurrent(v int32) *QueryOrderListRequest {
	s.Current = &v
	return s
}

func (s *QueryOrderListRequest) SetInstanceId(v string) *QueryOrderListRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryOrderListRequest) SetItemCode(v string) *QueryOrderListRequest {
	s.ItemCode = &v
	return s
}

func (s *QueryOrderListRequest) SetMobile(v string) *QueryOrderListRequest {
	s.Mobile = &v
	return s
}

func (s *QueryOrderListRequest) SetOrderTime(v string) *QueryOrderListRequest {
	s.OrderTime = &v
	return s
}

func (s *QueryOrderListRequest) SetOutBizNo(v string) *QueryOrderListRequest {
	s.OutBizNo = &v
	return s
}

func (s *QueryOrderListRequest) SetPageSize(v int32) *QueryOrderListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrderListRequest) SetUId(v int64) *QueryOrderListRequest {
	s.UId = &v
	return s
}

type QueryOrderListResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryOrderListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                          `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryOrderListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBody) SetCode(v string) *QueryOrderListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderListResponseBody) SetData(v *QueryOrderListResponseBodyData) *QueryOrderListResponseBody {
	s.Data = v
	return s
}

func (s *QueryOrderListResponseBody) SetMessage(v string) *QueryOrderListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderListResponseBody) SetRequestId(v string) *QueryOrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderListResponseBody) SetRt(v int64) *QueryOrderListResponseBody {
	s.Rt = &v
	return s
}

func (s *QueryOrderListResponseBody) SetSuccess(v bool) *QueryOrderListResponseBody {
	s.Success = &v
	return s
}

type QueryOrderListResponseBodyData struct {
	List     []*QueryOrderListResponseBodyDataList   `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageInfo *QueryOrderListResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s QueryOrderListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyData) SetList(v []*QueryOrderListResponseBodyDataList) *QueryOrderListResponseBodyData {
	s.List = v
	return s
}

func (s *QueryOrderListResponseBodyData) SetPageInfo(v *QueryOrderListResponseBodyDataPageInfo) *QueryOrderListResponseBodyData {
	s.PageInfo = v
	return s
}

type QueryOrderListResponseBodyDataList struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Operator    *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderTime   *string `json:"OrderTime,omitempty" xml:"OrderTime,omitempty"`
	OutBizNo    *string `json:"OutBizNo,omitempty" xml:"OutBizNo,omitempty"`
	ProductId   *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryOrderListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyDataList) SetInstanceId(v string) *QueryOrderListResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *QueryOrderListResponseBodyDataList) SetMessage(v string) *QueryOrderListResponseBodyDataList {
	s.Message = &v
	return s
}

func (s *QueryOrderListResponseBodyDataList) SetMobile(v string) *QueryOrderListResponseBodyDataList {
	s.Mobile = &v
	return s
}

func (s *QueryOrderListResponseBodyDataList) SetOperator(v string) *QueryOrderListResponseBodyDataList {
	s.Operator = &v
	return s
}

func (s *QueryOrderListResponseBodyDataList) SetOrderId(v string) *QueryOrderListResponseBodyDataList {
	s.OrderId = &v
	return s
}

func (s *QueryOrderListResponseBodyDataList) SetOrderTime(v string) *QueryOrderListResponseBodyDataList {
	s.OrderTime = &v
	return s
}

func (s *QueryOrderListResponseBodyDataList) SetOutBizNo(v string) *QueryOrderListResponseBodyDataList {
	s.OutBizNo = &v
	return s
}

func (s *QueryOrderListResponseBodyDataList) SetProductId(v string) *QueryOrderListResponseBodyDataList {
	s.ProductId = &v
	return s
}

func (s *QueryOrderListResponseBodyDataList) SetProductName(v string) *QueryOrderListResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *QueryOrderListResponseBodyDataList) SetStatus(v string) *QueryOrderListResponseBodyDataList {
	s.Status = &v
	return s
}

type QueryOrderListResponseBodyDataPageInfo struct {
	Current  *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryOrderListResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyDataPageInfo) SetCurrent(v int32) *QueryOrderListResponseBodyDataPageInfo {
	s.Current = &v
	return s
}

func (s *QueryOrderListResponseBodyDataPageInfo) SetPageSize(v int32) *QueryOrderListResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *QueryOrderListResponseBodyDataPageInfo) SetTotal(v int64) *QueryOrderListResponseBodyDataPageInfo {
	s.Total = &v
	return s
}

type QueryOrderListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOrderListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponse) SetHeaders(v map[string]*string) *QueryOrderListResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderListResponse) SetStatusCode(v int32) *QueryOrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderListResponse) SetBody(v *QueryOrderListResponseBody) *QueryOrderListResponse {
	s.Body = v
	return s
}

type SaveApplicationInfoRequest struct {
	AliUid      *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppTypeList *string `json:"AppTypeList,omitempty" xml:"AppTypeList,omitempty"`
	AppingList  *string `json:"AppingList,omitempty" xml:"AppingList,omitempty"`
	ItemCode    *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
}

func (s SaveApplicationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveApplicationInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveApplicationInfoRequest) SetAliUid(v int64) *SaveApplicationInfoRequest {
	s.AliUid = &v
	return s
}

func (s *SaveApplicationInfoRequest) SetAppId(v string) *SaveApplicationInfoRequest {
	s.AppId = &v
	return s
}

func (s *SaveApplicationInfoRequest) SetAppName(v string) *SaveApplicationInfoRequest {
	s.AppName = &v
	return s
}

func (s *SaveApplicationInfoRequest) SetAppTypeList(v string) *SaveApplicationInfoRequest {
	s.AppTypeList = &v
	return s
}

func (s *SaveApplicationInfoRequest) SetAppingList(v string) *SaveApplicationInfoRequest {
	s.AppingList = &v
	return s
}

func (s *SaveApplicationInfoRequest) SetItemCode(v string) *SaveApplicationInfoRequest {
	s.ItemCode = &v
	return s
}

type SaveApplicationInfoResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveApplicationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveApplicationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveApplicationInfoResponseBody) SetAppId(v string) *SaveApplicationInfoResponseBody {
	s.AppId = &v
	return s
}

func (s *SaveApplicationInfoResponseBody) SetCode(v string) *SaveApplicationInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SaveApplicationInfoResponseBody) SetMessage(v string) *SaveApplicationInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SaveApplicationInfoResponseBody) SetRequestId(v string) *SaveApplicationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveApplicationInfoResponseBody) SetSuccess(v bool) *SaveApplicationInfoResponseBody {
	s.Success = &v
	return s
}

type SaveApplicationInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveApplicationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveApplicationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveApplicationInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveApplicationInfoResponse) SetHeaders(v map[string]*string) *SaveApplicationInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveApplicationInfoResponse) SetStatusCode(v int32) *SaveApplicationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveApplicationInfoResponse) SetBody(v *SaveApplicationInfoResponseBody) *SaveApplicationInfoResponse {
	s.Body = v
	return s
}

type SdkGetInventoryInfoRequest struct {
	ChannelCode *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemCode    *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OutBizNo    *string `json:"OutBizNo,omitempty" xml:"OutBizNo,omitempty"`
	UId         *int64  `json:"UId,omitempty" xml:"UId,omitempty"`
}

func (s SdkGetInventoryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SdkGetInventoryInfoRequest) GoString() string {
	return s.String()
}

func (s *SdkGetInventoryInfoRequest) SetChannelCode(v string) *SdkGetInventoryInfoRequest {
	s.ChannelCode = &v
	return s
}

func (s *SdkGetInventoryInfoRequest) SetInstanceId(v string) *SdkGetInventoryInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *SdkGetInventoryInfoRequest) SetItemCode(v string) *SdkGetInventoryInfoRequest {
	s.ItemCode = &v
	return s
}

func (s *SdkGetInventoryInfoRequest) SetMobile(v string) *SdkGetInventoryInfoRequest {
	s.Mobile = &v
	return s
}

func (s *SdkGetInventoryInfoRequest) SetOutBizNo(v string) *SdkGetInventoryInfoRequest {
	s.OutBizNo = &v
	return s
}

func (s *SdkGetInventoryInfoRequest) SetUId(v int64) *SdkGetInventoryInfoRequest {
	s.UId = &v
	return s
}

type SdkGetInventoryInfoResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SdkGetInventoryInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                               `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SdkGetInventoryInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SdkGetInventoryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SdkGetInventoryInfoResponseBody) SetCode(v string) *SdkGetInventoryInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SdkGetInventoryInfoResponseBody) SetData(v *SdkGetInventoryInfoResponseBodyData) *SdkGetInventoryInfoResponseBody {
	s.Data = v
	return s
}

func (s *SdkGetInventoryInfoResponseBody) SetMessage(v string) *SdkGetInventoryInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SdkGetInventoryInfoResponseBody) SetRequestId(v string) *SdkGetInventoryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdkGetInventoryInfoResponseBody) SetRt(v int64) *SdkGetInventoryInfoResponseBody {
	s.Rt = &v
	return s
}

func (s *SdkGetInventoryInfoResponseBody) SetSuccess(v bool) *SdkGetInventoryInfoResponseBody {
	s.Success = &v
	return s
}

type SdkGetInventoryInfoResponseBodyData struct {
	Inventory         *int64 `json:"Inventory,omitempty" xml:"Inventory,omitempty"`
	ResidualInventory *int64 `json:"ResidualInventory,omitempty" xml:"ResidualInventory,omitempty"`
	UsedStock         *int64 `json:"UsedStock,omitempty" xml:"UsedStock,omitempty"`
}

func (s SdkGetInventoryInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SdkGetInventoryInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *SdkGetInventoryInfoResponseBodyData) SetInventory(v int64) *SdkGetInventoryInfoResponseBodyData {
	s.Inventory = &v
	return s
}

func (s *SdkGetInventoryInfoResponseBodyData) SetResidualInventory(v int64) *SdkGetInventoryInfoResponseBodyData {
	s.ResidualInventory = &v
	return s
}

func (s *SdkGetInventoryInfoResponseBodyData) SetUsedStock(v int64) *SdkGetInventoryInfoResponseBodyData {
	s.UsedStock = &v
	return s
}

type SdkGetInventoryInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SdkGetInventoryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SdkGetInventoryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SdkGetInventoryInfoResponse) GoString() string {
	return s.String()
}

func (s *SdkGetInventoryInfoResponse) SetHeaders(v map[string]*string) *SdkGetInventoryInfoResponse {
	s.Headers = v
	return s
}

func (s *SdkGetInventoryInfoResponse) SetStatusCode(v int32) *SdkGetInventoryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SdkGetInventoryInfoResponse) SetBody(v *SdkGetInventoryInfoResponseBody) *SdkGetInventoryInfoResponse {
	s.Body = v
	return s
}

type SdkOrderQosProductRequest struct {
	AliUid       *int64    `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ChannelId    *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CtToken      *string   `json:"CtToken,omitempty" xml:"CtToken,omitempty"`
	IPv6         *string   `json:"IPv6,omitempty" xml:"IPv6,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpType       *string   `json:"IpType,omitempty" xml:"IpType,omitempty"`
	MobileNumber *string   `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	Operator     *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	PrivateIpv4  *string   `json:"PrivateIpv4,omitempty" xml:"PrivateIpv4,omitempty"`
	ProductId    *string   `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Provice      *string   `json:"Provice,omitempty" xml:"Provice,omitempty"`
	PublicIpv4   *string   `json:"PublicIpv4,omitempty" xml:"PublicIpv4,omitempty"`
	QosRequestId *string   `json:"QosRequestId,omitempty" xml:"QosRequestId,omitempty"`
	TargetIpList []*string `json:"TargetIpList,omitempty" xml:"TargetIpList,omitempty" type:"Repeated"`
	Token        *string   `json:"Token,omitempty" xml:"Token,omitempty"`
	UnitNum      *int32    `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s SdkOrderQosProductRequest) String() string {
	return tea.Prettify(s)
}

func (s SdkOrderQosProductRequest) GoString() string {
	return s.String()
}

func (s *SdkOrderQosProductRequest) SetAliUid(v int64) *SdkOrderQosProductRequest {
	s.AliUid = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetChannelId(v string) *SdkOrderQosProductRequest {
	s.ChannelId = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetCtToken(v string) *SdkOrderQosProductRequest {
	s.CtToken = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetIPv6(v string) *SdkOrderQosProductRequest {
	s.IPv6 = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetInstanceId(v string) *SdkOrderQosProductRequest {
	s.InstanceId = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetIpType(v string) *SdkOrderQosProductRequest {
	s.IpType = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetMobileNumber(v string) *SdkOrderQosProductRequest {
	s.MobileNumber = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetOperator(v string) *SdkOrderQosProductRequest {
	s.Operator = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetPrivateIpv4(v string) *SdkOrderQosProductRequest {
	s.PrivateIpv4 = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetProductId(v string) *SdkOrderQosProductRequest {
	s.ProductId = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetProvice(v string) *SdkOrderQosProductRequest {
	s.Provice = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetPublicIpv4(v string) *SdkOrderQosProductRequest {
	s.PublicIpv4 = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetQosRequestId(v string) *SdkOrderQosProductRequest {
	s.QosRequestId = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetTargetIpList(v []*string) *SdkOrderQosProductRequest {
	s.TargetIpList = v
	return s
}

func (s *SdkOrderQosProductRequest) SetToken(v string) *SdkOrderQosProductRequest {
	s.Token = &v
	return s
}

func (s *SdkOrderQosProductRequest) SetUnitNum(v int32) *SdkOrderQosProductRequest {
	s.UnitNum = &v
	return s
}

type SdkOrderQosProductResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64  `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SdkOrderQosProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SdkOrderQosProductResponseBody) GoString() string {
	return s.String()
}

func (s *SdkOrderQosProductResponseBody) SetCode(v string) *SdkOrderQosProductResponseBody {
	s.Code = &v
	return s
}

func (s *SdkOrderQosProductResponseBody) SetData(v string) *SdkOrderQosProductResponseBody {
	s.Data = &v
	return s
}

func (s *SdkOrderQosProductResponseBody) SetMessage(v string) *SdkOrderQosProductResponseBody {
	s.Message = &v
	return s
}

func (s *SdkOrderQosProductResponseBody) SetRequestId(v string) *SdkOrderQosProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdkOrderQosProductResponseBody) SetRt(v int64) *SdkOrderQosProductResponseBody {
	s.Rt = &v
	return s
}

func (s *SdkOrderQosProductResponseBody) SetSuccess(v bool) *SdkOrderQosProductResponseBody {
	s.Success = &v
	return s
}

type SdkOrderQosProductResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SdkOrderQosProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SdkOrderQosProductResponse) String() string {
	return tea.Prettify(s)
}

func (s SdkOrderQosProductResponse) GoString() string {
	return s.String()
}

func (s *SdkOrderQosProductResponse) SetHeaders(v map[string]*string) *SdkOrderQosProductResponse {
	s.Headers = v
	return s
}

func (s *SdkOrderQosProductResponse) SetStatusCode(v int32) *SdkOrderQosProductResponse {
	s.StatusCode = &v
	return s
}

func (s *SdkOrderQosProductResponse) SetBody(v *SdkOrderQosProductResponseBody) *SdkOrderQosProductResponse {
	s.Body = v
	return s
}

type SdkValidateStatusRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CredentialType  *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	CredentialValue *string `json:"CredentialValue,omitempty" xml:"CredentialValue,omitempty"`
	Operator        *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s SdkValidateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SdkValidateStatusRequest) GoString() string {
	return s.String()
}

func (s *SdkValidateStatusRequest) SetAppId(v string) *SdkValidateStatusRequest {
	s.AppId = &v
	return s
}

func (s *SdkValidateStatusRequest) SetCredentialType(v string) *SdkValidateStatusRequest {
	s.CredentialType = &v
	return s
}

func (s *SdkValidateStatusRequest) SetCredentialValue(v string) *SdkValidateStatusRequest {
	s.CredentialValue = &v
	return s
}

func (s *SdkValidateStatusRequest) SetOperator(v string) *SdkValidateStatusRequest {
	s.Operator = &v
	return s
}

func (s *SdkValidateStatusRequest) SetToken(v string) *SdkValidateStatusRequest {
	s.Token = &v
	return s
}

type SdkValidateStatusResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SdkValidateStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                             `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SdkValidateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SdkValidateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SdkValidateStatusResponseBody) SetCode(v string) *SdkValidateStatusResponseBody {
	s.Code = &v
	return s
}

func (s *SdkValidateStatusResponseBody) SetData(v *SdkValidateStatusResponseBodyData) *SdkValidateStatusResponseBody {
	s.Data = v
	return s
}

func (s *SdkValidateStatusResponseBody) SetMessage(v string) *SdkValidateStatusResponseBody {
	s.Message = &v
	return s
}

func (s *SdkValidateStatusResponseBody) SetRequestId(v string) *SdkValidateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdkValidateStatusResponseBody) SetRt(v int64) *SdkValidateStatusResponseBody {
	s.Rt = &v
	return s
}

func (s *SdkValidateStatusResponseBody) SetSuccess(v bool) *SdkValidateStatusResponseBody {
	s.Success = &v
	return s
}

type SdkValidateStatusResponseBodyData struct {
	AppExtPopList []*SdkValidateStatusResponseBodyDataAppExtPopList `json:"AppExtPopList,omitempty" xml:"AppExtPopList,omitempty" type:"Repeated"`
	FreeFlow      *bool                                             `json:"FreeFlow,omitempty" xml:"FreeFlow,omitempty"`
	PseudoCode    *string                                           `json:"PseudoCode,omitempty" xml:"PseudoCode,omitempty"`
}

func (s SdkValidateStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SdkValidateStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *SdkValidateStatusResponseBodyData) SetAppExtPopList(v []*SdkValidateStatusResponseBodyDataAppExtPopList) *SdkValidateStatusResponseBodyData {
	s.AppExtPopList = v
	return s
}

func (s *SdkValidateStatusResponseBodyData) SetFreeFlow(v bool) *SdkValidateStatusResponseBodyData {
	s.FreeFlow = &v
	return s
}

func (s *SdkValidateStatusResponseBodyData) SetPseudoCode(v string) *SdkValidateStatusResponseBodyData {
	s.PseudoCode = &v
	return s
}

type SdkValidateStatusResponseBodyDataAppExtPopList struct {
	ExtId           *int64    `json:"ExtId,omitempty" xml:"ExtId,omitempty"`
	FlowIp          []*string `json:"FlowIp,omitempty" xml:"FlowIp,omitempty" type:"Repeated"`
	FlowUrl         []*string `json:"FlowUrl,omitempty" xml:"FlowUrl,omitempty" type:"Repeated"`
	OriginalIpList  []*string `json:"OriginalIpList,omitempty" xml:"OriginalIpList,omitempty" type:"Repeated"`
	OriginalUrlList []*string `json:"OriginalUrlList,omitempty" xml:"OriginalUrlList,omitempty" type:"Repeated"`
}

func (s SdkValidateStatusResponseBodyDataAppExtPopList) String() string {
	return tea.Prettify(s)
}

func (s SdkValidateStatusResponseBodyDataAppExtPopList) GoString() string {
	return s.String()
}

func (s *SdkValidateStatusResponseBodyDataAppExtPopList) SetExtId(v int64) *SdkValidateStatusResponseBodyDataAppExtPopList {
	s.ExtId = &v
	return s
}

func (s *SdkValidateStatusResponseBodyDataAppExtPopList) SetFlowIp(v []*string) *SdkValidateStatusResponseBodyDataAppExtPopList {
	s.FlowIp = v
	return s
}

func (s *SdkValidateStatusResponseBodyDataAppExtPopList) SetFlowUrl(v []*string) *SdkValidateStatusResponseBodyDataAppExtPopList {
	s.FlowUrl = v
	return s
}

func (s *SdkValidateStatusResponseBodyDataAppExtPopList) SetOriginalIpList(v []*string) *SdkValidateStatusResponseBodyDataAppExtPopList {
	s.OriginalIpList = v
	return s
}

func (s *SdkValidateStatusResponseBodyDataAppExtPopList) SetOriginalUrlList(v []*string) *SdkValidateStatusResponseBodyDataAppExtPopList {
	s.OriginalUrlList = v
	return s
}

type SdkValidateStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SdkValidateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SdkValidateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SdkValidateStatusResponse) GoString() string {
	return s.String()
}

func (s *SdkValidateStatusResponse) SetHeaders(v map[string]*string) *SdkValidateStatusResponse {
	s.Headers = v
	return s
}

func (s *SdkValidateStatusResponse) SetStatusCode(v int32) *SdkValidateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SdkValidateStatusResponse) SetBody(v *SdkValidateStatusResponseBody) *SdkValidateStatusResponse {
	s.Body = v
	return s
}

type ValidControllerAuthorRequest struct {
	AliUid   *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ItemCode *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
}

func (s ValidControllerAuthorRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidControllerAuthorRequest) GoString() string {
	return s.String()
}

func (s *ValidControllerAuthorRequest) SetAliUid(v int64) *ValidControllerAuthorRequest {
	s.AliUid = &v
	return s
}

func (s *ValidControllerAuthorRequest) SetItemCode(v string) *ValidControllerAuthorRequest {
	s.ItemCode = &v
	return s
}

type ValidControllerAuthorResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64  `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidControllerAuthorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidControllerAuthorResponseBody) GoString() string {
	return s.String()
}

func (s *ValidControllerAuthorResponseBody) SetCode(v string) *ValidControllerAuthorResponseBody {
	s.Code = &v
	return s
}

func (s *ValidControllerAuthorResponseBody) SetData(v bool) *ValidControllerAuthorResponseBody {
	s.Data = &v
	return s
}

func (s *ValidControllerAuthorResponseBody) SetMessage(v string) *ValidControllerAuthorResponseBody {
	s.Message = &v
	return s
}

func (s *ValidControllerAuthorResponseBody) SetRequestId(v string) *ValidControllerAuthorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidControllerAuthorResponseBody) SetRt(v int64) *ValidControllerAuthorResponseBody {
	s.Rt = &v
	return s
}

func (s *ValidControllerAuthorResponseBody) SetSuccess(v bool) *ValidControllerAuthorResponseBody {
	s.Success = &v
	return s
}

type ValidControllerAuthorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidControllerAuthorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidControllerAuthorResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidControllerAuthorResponse) GoString() string {
	return s.String()
}

func (s *ValidControllerAuthorResponse) SetHeaders(v map[string]*string) *ValidControllerAuthorResponse {
	s.Headers = v
	return s
}

func (s *ValidControllerAuthorResponse) SetStatusCode(v int32) *ValidControllerAuthorResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidControllerAuthorResponse) SetBody(v *ValidControllerAuthorResponseBody) *ValidControllerAuthorResponse {
	s.Body = v
	return s
}

type ValidFlowRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemCode   *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	Mobile     *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UId        *int64  `json:"UId,omitempty" xml:"UId,omitempty"`
}

func (s ValidFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidFlowRequest) GoString() string {
	return s.String()
}

func (s *ValidFlowRequest) SetInstanceId(v string) *ValidFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *ValidFlowRequest) SetItemCode(v string) *ValidFlowRequest {
	s.ItemCode = &v
	return s
}

func (s *ValidFlowRequest) SetMobile(v string) *ValidFlowRequest {
	s.Mobile = &v
	return s
}

func (s *ValidFlowRequest) SetUId(v int64) *ValidFlowRequest {
	s.UId = &v
	return s
}

type ValidFlowResponseBody struct {
	Code      *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64      `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ValidFlowResponseBody) SetCode(v string) *ValidFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ValidFlowResponseBody) SetData(v interface{}) *ValidFlowResponseBody {
	s.Data = v
	return s
}

func (s *ValidFlowResponseBody) SetMessage(v string) *ValidFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ValidFlowResponseBody) SetRequestId(v string) *ValidFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidFlowResponseBody) SetRt(v int64) *ValidFlowResponseBody {
	s.Rt = &v
	return s
}

func (s *ValidFlowResponseBody) SetSuccess(v bool) *ValidFlowResponseBody {
	s.Success = &v
	return s
}

type ValidFlowResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidFlowResponse) GoString() string {
	return s.String()
}

func (s *ValidFlowResponse) SetHeaders(v map[string]*string) *ValidFlowResponse {
	s.Headers = v
	return s
}

func (s *ValidFlowResponse) SetStatusCode(v int32) *ValidFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidFlowResponse) SetBody(v *ValidFlowResponseBody) *ValidFlowResponse {
	s.Body = v
	return s
}

type ValidateStatusRequest struct {
	AliUid          *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CredentialType  *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	CredentialValue *string `json:"CredentialValue,omitempty" xml:"CredentialValue,omitempty"`
	Operator        *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s ValidateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateStatusRequest) GoString() string {
	return s.String()
}

func (s *ValidateStatusRequest) SetAliUid(v int64) *ValidateStatusRequest {
	s.AliUid = &v
	return s
}

func (s *ValidateStatusRequest) SetAppId(v string) *ValidateStatusRequest {
	s.AppId = &v
	return s
}

func (s *ValidateStatusRequest) SetCredentialType(v string) *ValidateStatusRequest {
	s.CredentialType = &v
	return s
}

func (s *ValidateStatusRequest) SetCredentialValue(v string) *ValidateStatusRequest {
	s.CredentialValue = &v
	return s
}

func (s *ValidateStatusRequest) SetOperator(v string) *ValidateStatusRequest {
	s.Operator = &v
	return s
}

type ValidateStatusResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ValidateStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rt        *int64                          `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateStatusResponseBody) SetCode(v string) *ValidateStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateStatusResponseBody) SetData(v *ValidateStatusResponseBodyData) *ValidateStatusResponseBody {
	s.Data = v
	return s
}

func (s *ValidateStatusResponseBody) SetMessage(v string) *ValidateStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateStatusResponseBody) SetRequestId(v string) *ValidateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateStatusResponseBody) SetRt(v int64) *ValidateStatusResponseBody {
	s.Rt = &v
	return s
}

func (s *ValidateStatusResponseBody) SetSuccess(v bool) *ValidateStatusResponseBody {
	s.Success = &v
	return s
}

type ValidateStatusResponseBodyData struct {
	AppExtPopList []*ValidateStatusResponseBodyDataAppExtPopList `json:"AppExtPopList,omitempty" xml:"AppExtPopList,omitempty" type:"Repeated"`
	FreeFlow      *bool                                          `json:"FreeFlow,omitempty" xml:"FreeFlow,omitempty"`
	PseudoCode    *string                                        `json:"PseudoCode,omitempty" xml:"PseudoCode,omitempty"`
}

func (s ValidateStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ValidateStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValidateStatusResponseBodyData) SetAppExtPopList(v []*ValidateStatusResponseBodyDataAppExtPopList) *ValidateStatusResponseBodyData {
	s.AppExtPopList = v
	return s
}

func (s *ValidateStatusResponseBodyData) SetFreeFlow(v bool) *ValidateStatusResponseBodyData {
	s.FreeFlow = &v
	return s
}

func (s *ValidateStatusResponseBodyData) SetPseudoCode(v string) *ValidateStatusResponseBodyData {
	s.PseudoCode = &v
	return s
}

type ValidateStatusResponseBodyDataAppExtPopList struct {
	ExtId           *int64    `json:"ExtId,omitempty" xml:"ExtId,omitempty"`
	FlowIp          []*string `json:"FlowIp,omitempty" xml:"FlowIp,omitempty" type:"Repeated"`
	FlowUrl         []*string `json:"FlowUrl,omitempty" xml:"FlowUrl,omitempty" type:"Repeated"`
	OriginalIpList  []*string `json:"OriginalIpList,omitempty" xml:"OriginalIpList,omitempty" type:"Repeated"`
	OriginalUrlList []*string `json:"OriginalUrlList,omitempty" xml:"OriginalUrlList,omitempty" type:"Repeated"`
}

func (s ValidateStatusResponseBodyDataAppExtPopList) String() string {
	return tea.Prettify(s)
}

func (s ValidateStatusResponseBodyDataAppExtPopList) GoString() string {
	return s.String()
}

func (s *ValidateStatusResponseBodyDataAppExtPopList) SetExtId(v int64) *ValidateStatusResponseBodyDataAppExtPopList {
	s.ExtId = &v
	return s
}

func (s *ValidateStatusResponseBodyDataAppExtPopList) SetFlowIp(v []*string) *ValidateStatusResponseBodyDataAppExtPopList {
	s.FlowIp = v
	return s
}

func (s *ValidateStatusResponseBodyDataAppExtPopList) SetFlowUrl(v []*string) *ValidateStatusResponseBodyDataAppExtPopList {
	s.FlowUrl = v
	return s
}

func (s *ValidateStatusResponseBodyDataAppExtPopList) SetOriginalIpList(v []*string) *ValidateStatusResponseBodyDataAppExtPopList {
	s.OriginalIpList = v
	return s
}

func (s *ValidateStatusResponseBodyDataAppExtPopList) SetOriginalUrlList(v []*string) *ValidateStatusResponseBodyDataAppExtPopList {
	s.OriginalUrlList = v
	return s
}

type ValidateStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateStatusResponse) GoString() string {
	return s.String()
}

func (s *ValidateStatusResponse) SetHeaders(v map[string]*string) *ValidateStatusResponse {
	s.Headers = v
	return s
}

func (s *ValidateStatusResponse) SetStatusCode(v int32) *ValidateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateStatusResponse) SetBody(v *ValidateStatusResponseBody) *ValidateStatusResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("xgippop"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ChangeApplicationInfoWithOptions(request *ChangeApplicationInfoRequest, runtime *util.RuntimeOptions) (_result *ChangeApplicationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppTypeList)) {
		body["AppTypeList"] = request.AppTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.AppingList)) {
		body["AppingList"] = request.AppingList
	}

	if !tea.BoolValue(util.IsUnset(request.ItemCode)) {
		body["ItemCode"] = request.ItemCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeApplicationInfo"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeApplicationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeApplicationInfo(request *ChangeApplicationInfoRequest) (_result *ChangeApplicationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeApplicationInfoResponse{}
	_body, _err := client.ChangeApplicationInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChargeFlowWithOptions(request *ChargeFlowRequest, runtime *util.RuntimeOptions) (_result *ChargeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelCode)) {
		query["ChannelCode"] = request.ChannelCode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemCode)) {
		query["ItemCode"] = request.ItemCode
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.OrderTime)) {
		query["OrderTime"] = request.OrderTime
	}

	if !tea.BoolValue(util.IsUnset(request.OutBizNo)) {
		query["OutBizNo"] = request.OutBizNo
	}

	if !tea.BoolValue(util.IsUnset(request.UId)) {
		query["UId"] = request.UId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChargeFlow"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChargeFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChargeFlow(request *ChargeFlowRequest) (_result *ChargeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChargeFlowResponse{}
	_body, _err := client.ChargeFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationInfoWithOptions(request *CreateApplicationInfoRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppTypeList)) {
		body["AppTypeList"] = request.AppTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.AppingList)) {
		body["AppingList"] = request.AppingList
	}

	if !tea.BoolValue(util.IsUnset(request.ItemCode)) {
		body["ItemCode"] = request.ItemCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplicationInfo"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplicationInfo(request *CreateApplicationInfoRequest) (_result *CreateApplicationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationInfoResponse{}
	_body, _err := client.CreateApplicationInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAliyunXgipTokenWithOptions(runtime *util.RuntimeOptions) (_result *GetAliyunXgipTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAliyunXgipToken"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAliyunXgipTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAliyunXgipToken() (_result *GetAliyunXgipTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAliyunXgipTokenResponse{}
	_body, _err := client.GetAliyunXgipTokenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *util.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplication"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFreeFlowInstanceWithOptions(request *GetFreeFlowInstanceRequest, runtime *util.RuntimeOptions) (_result *GetFreeFlowInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFreeFlowInstance"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFreeFlowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFreeFlowInstance(request *GetFreeFlowInstanceRequest) (_result *GetFreeFlowInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFreeFlowInstanceResponse{}
	_body, _err := client.GetFreeFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFreeFlowProductListWithOptions(request *GetFreeFlowProductListRequest, runtime *util.RuntimeOptions) (_result *GetFreeFlowProductListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFreeFlowProductList"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFreeFlowProductListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFreeFlowProductList(request *GetFreeFlowProductListRequest) (_result *GetFreeFlowProductListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFreeFlowProductListResponse{}
	_body, _err := client.GetFreeFlowProductListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFreeFlowUsageWithOptions(request *GetFreeFlowUsageRequest, runtime *util.RuntimeOptions) (_result *GetFreeFlowUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFreeFlowUsage"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFreeFlowUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFreeFlowUsage(request *GetFreeFlowUsageRequest) (_result *GetFreeFlowUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFreeFlowUsageResponse{}
	_body, _err := client.GetFreeFlowUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFreeFlowUsageStatisticWithOptions(request *GetFreeFlowUsageStatisticRequest, runtime *util.RuntimeOptions) (_result *GetFreeFlowUsageStatisticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFreeFlowUsageStatistic"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFreeFlowUsageStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFreeFlowUsageStatistic(request *GetFreeFlowUsageStatisticRequest) (_result *GetFreeFlowUsageStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFreeFlowUsageStatisticResponse{}
	_body, _err := client.GetFreeFlowUsageStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInventoryInfoWithOptions(request *GetInventoryInfoRequest, runtime *util.RuntimeOptions) (_result *GetInventoryInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInventoryInfo"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInventoryInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInventoryInfo(request *GetInventoryInfoRequest) (_result *GetInventoryInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInventoryInfoResponse{}
	_body, _err := client.GetInventoryInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetItemInstListWithOptions(request *GetItemInstListRequest, runtime *util.RuntimeOptions) (_result *GetItemInstListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetItemInstList"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetItemInstListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetItemInstList(request *GetItemInstListRequest) (_result *GetItemInstListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetItemInstListResponse{}
	_body, _err := client.GetItemInstListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetItemListWithOptions(request *GetItemListRequest, runtime *util.RuntimeOptions) (_result *GetItemListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetItemList"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetItemList(request *GetItemListRequest) (_result *GetItemListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetItemListResponse{}
	_body, _err := client.GetItemListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrderFreeFlowProductStatusWithOptions(request *GetOrderFreeFlowProductStatusRequest, runtime *util.RuntimeOptions) (_result *GetOrderFreeFlowProductStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrderFreeFlowProductStatus"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrderFreeFlowProductStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrderFreeFlowProductStatus(request *GetOrderFreeFlowProductStatusRequest) (_result *GetOrderFreeFlowProductStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOrderFreeFlowProductStatusResponse{}
	_body, _err := client.GetOrderFreeFlowProductStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrderItemListWithOptions(runtime *util.RuntimeOptions) (_result *GetOrderItemListResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetOrderItemList"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrderItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrderItemList() (_result *GetOrderItemListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOrderItemListResponse{}
	_body, _err := client.GetOrderItemListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQosFlowUsageWithOptions(request *GetQosFlowUsageRequest, runtime *util.RuntimeOptions) (_result *GetQosFlowUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQosFlowUsage"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQosFlowUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQosFlowUsage(request *GetQosFlowUsageRequest) (_result *GetQosFlowUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQosFlowUsageResponse{}
	_body, _err := client.GetQosFlowUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQosUsageStatisticWithOptions(request *GetQosUsageStatisticRequest, runtime *util.RuntimeOptions) (_result *GetQosUsageStatisticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQosUsageStatistic"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQosUsageStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQosUsageStatistic(request *GetQosUsageStatisticRequest) (_result *GetQosUsageStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQosUsageStatisticResponse{}
	_body, _err := client.GetQosUsageStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUatDataListWithOptions(request *GetUatDataListRequest, runtime *util.RuntimeOptions) (_result *GetUatDataListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUatDataList"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUatDataListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUatDataList(request *GetUatDataListRequest) (_result *GetUatDataListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUatDataListResponse{}
	_body, _err := client.GetUatDataListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUatSpecCtDataWithOptions(request *GetUatSpecCtDataRequest, runtime *util.RuntimeOptions) (_result *GetUatSpecCtDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUatSpecCtData"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUatSpecCtDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUatSpecCtData(request *GetUatSpecCtDataRequest) (_result *GetUatSpecCtDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUatSpecCtDataResponse{}
	_body, _err := client.GetUatSpecCtDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MockGetOrderFreeFlowProductStatusWithOptions(request *MockGetOrderFreeFlowProductStatusRequest, runtime *util.RuntimeOptions) (_result *MockGetOrderFreeFlowProductStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MockGetOrderFreeFlowProductStatus"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MockGetOrderFreeFlowProductStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MockGetOrderFreeFlowProductStatus(request *MockGetOrderFreeFlowProductStatusRequest) (_result *MockGetOrderFreeFlowProductStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MockGetOrderFreeFlowProductStatusResponse{}
	_body, _err := client.MockGetOrderFreeFlowProductStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MockOrderFreeFlowProductWithOptions(request *MockOrderFreeFlowProductRequest, runtime *util.RuntimeOptions) (_result *MockOrderFreeFlowProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerFlowRequestId)) {
		query["CustomerFlowRequestId"] = request.CustomerFlowRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowProductId)) {
		query["FlowProductId"] = request.FlowProductId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lasting)) {
		query["Lasting"] = request.Lasting
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNumber)) {
		query["MobileNumber"] = request.MobileNumber
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyUrl)) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseOrderId)) {
		query["PurchaseOrderId"] = request.PurchaseOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MockOrderFreeFlowProduct"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MockOrderFreeFlowProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MockOrderFreeFlowProduct(request *MockOrderFreeFlowProductRequest) (_result *MockOrderFreeFlowProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MockOrderFreeFlowProductResponse{}
	_body, _err := client.MockOrderFreeFlowProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyApplicationWithOptions(request *ModifyApplicationRequest, runtime *util.RuntimeOptions) (_result *ModifyApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		body["AppCode"] = request.AppCode
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppTypeList)) {
		body["AppTypeList"] = request.AppTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.AppingList)) {
		body["AppingList"] = request.AppingList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApplication"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApplication(request *ModifyApplicationRequest) (_result *ModifyApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyApplicationResponse{}
	_body, _err := client.ModifyApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OrderFreeFlowProductWithOptions(request *OrderFreeFlowProductRequest, runtime *util.RuntimeOptions) (_result *OrderFreeFlowProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerFlowRequestId)) {
		query["CustomerFlowRequestId"] = request.CustomerFlowRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.FlowProductId)) {
		query["FlowProductId"] = request.FlowProductId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lasting)) {
		query["Lasting"] = request.Lasting
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNumber)) {
		query["MobileNumber"] = request.MobileNumber
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyUrl)) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseOrderId)) {
		query["PurchaseOrderId"] = request.PurchaseOrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OrderFreeFlowProduct"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OrderFreeFlowProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OrderFreeFlowProduct(request *OrderFreeFlowProductRequest) (_result *OrderFreeFlowProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OrderFreeFlowProductResponse{}
	_body, _err := client.OrderFreeFlowProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OrderQosProductWithOptions(request *OrderQosProductRequest, runtime *util.RuntimeOptions) (_result *OrderQosProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Provice)) {
		query["Provice"] = request.Provice
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		body["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.IPv6)) {
		body["IPv6"] = request.IPv6
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpType)) {
		body["IpType"] = request.IpType
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNumber)) {
		body["MobileNumber"] = request.MobileNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpv4)) {
		body["PrivateIpv4"] = request.PrivateIpv4
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIpv4)) {
		body["PublicIpv4"] = request.PublicIpv4
	}

	if !tea.BoolValue(util.IsUnset(request.QosRequestId)) {
		body["QosRequestId"] = request.QosRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetIpList)) {
		body["TargetIpList"] = request.TargetIpList
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UnitNum)) {
		body["UnitNum"] = request.UnitNum
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OrderQosProduct"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OrderQosProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OrderQosProduct(request *OrderQosProductRequest) (_result *OrderQosProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OrderQosProductResponse{}
	_body, _err := client.OrderQosProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderListWithOptions(request *QueryOrderListRequest, runtime *util.RuntimeOptions) (_result *QueryOrderListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrderList"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderList(request *QueryOrderListRequest) (_result *QueryOrderListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderListResponse{}
	_body, _err := client.QueryOrderListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveApplicationInfoWithOptions(request *SaveApplicationInfoRequest, runtime *util.RuntimeOptions) (_result *SaveApplicationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppTypeList)) {
		body["AppTypeList"] = request.AppTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.AppingList)) {
		body["AppingList"] = request.AppingList
	}

	if !tea.BoolValue(util.IsUnset(request.ItemCode)) {
		body["ItemCode"] = request.ItemCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveApplicationInfo"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveApplicationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveApplicationInfo(request *SaveApplicationInfoRequest) (_result *SaveApplicationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveApplicationInfoResponse{}
	_body, _err := client.SaveApplicationInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SdkGetInventoryInfoWithOptions(request *SdkGetInventoryInfoRequest, runtime *util.RuntimeOptions) (_result *SdkGetInventoryInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SdkGetInventoryInfo"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SdkGetInventoryInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SdkGetInventoryInfo(request *SdkGetInventoryInfoRequest) (_result *SdkGetInventoryInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SdkGetInventoryInfoResponse{}
	_body, _err := client.SdkGetInventoryInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SdkOrderQosProductWithOptions(request *SdkOrderQosProductRequest, runtime *util.RuntimeOptions) (_result *SdkOrderQosProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CtToken)) {
		query["CtToken"] = request.CtToken
	}

	if !tea.BoolValue(util.IsUnset(request.IPv6)) {
		query["IPv6"] = request.IPv6
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpType)) {
		query["IpType"] = request.IpType
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNumber)) {
		query["MobileNumber"] = request.MobileNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpv4)) {
		query["PrivateIpv4"] = request.PrivateIpv4
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Provice)) {
		query["Provice"] = request.Provice
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIpv4)) {
		query["PublicIpv4"] = request.PublicIpv4
	}

	if !tea.BoolValue(util.IsUnset(request.QosRequestId)) {
		query["QosRequestId"] = request.QosRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetIpList)) {
		query["TargetIpList"] = request.TargetIpList
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UnitNum)) {
		query["UnitNum"] = request.UnitNum
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SdkOrderQosProduct"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SdkOrderQosProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SdkOrderQosProduct(request *SdkOrderQosProductRequest) (_result *SdkOrderQosProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SdkOrderQosProductResponse{}
	_body, _err := client.SdkOrderQosProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SdkValidateStatusWithOptions(request *SdkValidateStatusRequest, runtime *util.RuntimeOptions) (_result *SdkValidateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SdkValidateStatus"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SdkValidateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SdkValidateStatus(request *SdkValidateStatusRequest) (_result *SdkValidateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SdkValidateStatusResponse{}
	_body, _err := client.SdkValidateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidControllerAuthorWithOptions(request *ValidControllerAuthorRequest, runtime *util.RuntimeOptions) (_result *ValidControllerAuthorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidControllerAuthor"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidControllerAuthorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidControllerAuthor(request *ValidControllerAuthorRequest) (_result *ValidControllerAuthorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidControllerAuthorResponse{}
	_body, _err := client.ValidControllerAuthorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidFlowWithOptions(request *ValidFlowRequest, runtime *util.RuntimeOptions) (_result *ValidFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidFlow"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidFlow(request *ValidFlowRequest) (_result *ValidFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidFlowResponse{}
	_body, _err := client.ValidFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateStatusWithOptions(request *ValidateStatusRequest, runtime *util.RuntimeOptions) (_result *ValidateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateStatus"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateStatus(request *ValidateStatusRequest) (_result *ValidateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateStatusResponse{}
	_body, _err := client.ValidateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
