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

type CreateApplicationInfoRequest struct {
	// 阿里UID
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic（动态业务）/static（静态业务
	AppTypeList []*string                                 `json:"AppTypeList,omitempty" xml:"AppTypeList,omitempty" type:"Repeated"`
	AppingList  []*CreateApplicationInfoRequestAppingList `json:"AppingList,omitempty" xml:"AppingList,omitempty" type:"Repeated"`
	// 商品code
	ItemCode *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
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
	ExtId *int64 `json:"ExtId,omitempty" xml:"ExtId,omitempty"`
	// cdn ip
	FlowIp []*string `json:"FlowIp,omitempty" xml:"FlowIp,omitempty" type:"Repeated"`
	// cdn 域名信息
	FlowUrl []*string `json:"FlowUrl,omitempty" xml:"FlowUrl,omitempty" type:"Repeated"`
	// 业务方ip集合
	OriginalIpList []*string `json:"OriginalIpList,omitempty" xml:"OriginalIpList,omitempty" type:"Repeated"`
	// 业务方域名集合
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
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

type GetApplicationRequest struct {
	// 阿里UID
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 应用ID
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
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

type GetApplicationResponseBody struct {
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果
	Data []*GetApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求链路ID，如POP请求进来的requestId，返回时原样返回
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务端处理耗时，ms
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *GetApplicationResponseBody) SetData(v []*GetApplicationResponseBodyData) *GetApplicationResponseBody {
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

type GetApplicationResponseBodyData struct {
	// 用户编号
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 应用编号
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用名称
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// //dynamic（动态业务）/static（静态业务）  // 前端传递，参数3
	AppTypeList *string                                     `json:"AppTypeList,omitempty" xml:"AppTypeList,omitempty"`
	AppingList  []*GetApplicationResponseBodyDataAppingList `json:"AppingList,omitempty" xml:"AppingList,omitempty" type:"Repeated"`
	EndTime     *string                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Database Column Remarks:
	ItemCode *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	// 开通时间
	OpenTime  *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyData) SetAliUid(v int64) *GetApplicationResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetAppId(v string) *GetApplicationResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetAppName(v string) *GetApplicationResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetAppStatus(v string) *GetApplicationResponseBodyData {
	s.AppStatus = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetAppTypeList(v string) *GetApplicationResponseBodyData {
	s.AppTypeList = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetAppingList(v []*GetApplicationResponseBodyDataAppingList) *GetApplicationResponseBodyData {
	s.AppingList = v
	return s
}

func (s *GetApplicationResponseBodyData) SetEndTime(v string) *GetApplicationResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetItemCode(v string) *GetApplicationResponseBodyData {
	s.ItemCode = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetOpenTime(v string) *GetApplicationResponseBodyData {
	s.OpenTime = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetStartTime(v string) *GetApplicationResponseBodyData {
	s.StartTime = &v
	return s
}

type GetApplicationResponseBodyDataAppingList struct {
	ExtId *int64 `json:"ExtId,omitempty" xml:"ExtId,omitempty"`
	// cdn ip
	FlowIp []*string `json:"FlowIp,omitempty" xml:"FlowIp,omitempty" type:"Repeated"`
	// cdn 域名信息
	FlowUrl []*string `json:"FlowUrl,omitempty" xml:"FlowUrl,omitempty" type:"Repeated"`
	// 业务方ip集合
	OriginalIpList []*string `json:"OriginalIpList,omitempty" xml:"OriginalIpList,omitempty" type:"Repeated"`
	// 业务方域名集合
	OriginalUrlList []*string `json:"OriginalUrlList,omitempty" xml:"OriginalUrlList,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyDataAppingList) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyDataAppingList) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataAppingList) SetExtId(v int64) *GetApplicationResponseBodyDataAppingList {
	s.ExtId = &v
	return s
}

func (s *GetApplicationResponseBodyDataAppingList) SetFlowIp(v []*string) *GetApplicationResponseBodyDataAppingList {
	s.FlowIp = v
	return s
}

func (s *GetApplicationResponseBodyDataAppingList) SetFlowUrl(v []*string) *GetApplicationResponseBodyDataAppingList {
	s.FlowUrl = v
	return s
}

func (s *GetApplicationResponseBodyDataAppingList) SetOriginalIpList(v []*string) *GetApplicationResponseBodyDataAppingList {
	s.OriginalIpList = v
	return s
}

func (s *GetApplicationResponseBodyDataAppingList) SetOriginalUrlList(v []*string) *GetApplicationResponseBodyDataAppingList {
	s.OriginalUrlList = v
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
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemCode   *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
}

func (s GetFreeFlowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetFreeFlowInstanceRequest) SetAliUid(v int64) *GetFreeFlowInstanceRequest {
	s.AliUid = &v
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
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果
	Data []*GetFreeFlowInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求链路ID，如POP请求进来的requestId，返回时原样返回
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务端处理耗时，ms
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// APP编号
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// APP名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 产品失效时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 实例名称
	InstanceMemo *string `json:"InstanceMemo,omitempty" xml:"InstanceMemo,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// 产品开通时间
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	// 规格类型
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// 产品生效时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// 用户编号
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 实例ID
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
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果
	Data []*GetFreeFlowProductListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求链路ID，如POP请求进来的requestId，返回时原样返回
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务端处理耗时，ms
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// 产品当前状态，true为已配置，false为未配置
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
	// 产品包含的流量大小
	FlowProductAmount *string `json:"FlowProductAmount,omitempty" xml:"FlowProductAmount,omitempty"`
	// 免流产品ID
	FlowProductId *string `json:"FlowProductId,omitempty" xml:"FlowProductId,omitempty"`
	// 流量包名称
	FlowProductName *string `json:"FlowProductName,omitempty" xml:"FlowProductName,omitempty"`
	// 产品周期
	FlowProductPeriod *string `json:"FlowProductPeriod,omitempty" xml:"FlowProductPeriod,omitempty"`
	// 取值包括free（定向流量）/normal（通用流量）
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// 取值包括cm（中国移动）/ct（中国电信）/cu（中国联通）
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// 注意这里返回的全量套餐里，只能包含SpecType与该InstanceId的SpecType相同的规格
	SpecType  *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Spid      *string `json:"Spid,omitempty" xml:"Spid,omitempty"`
	UnitPrice *int32  `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
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

type ModifyApplicationRequest struct {
	// AliUid
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// AppId
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic（动态业务）/static（静态业务
	AppTypeList []*string `json:"AppTypeList,omitempty" xml:"AppTypeList,omitempty" type:"Repeated"`
	// 扩展属性 源域名和源ip信息保存
	AppingList []*ModifyApplicationRequestAppingList `json:"AppingList,omitempty" xml:"AppingList,omitempty" type:"Repeated"`
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
	ExtId *int64 `json:"ExtId,omitempty" xml:"ExtId,omitempty"`
	// cdn ip
	FlowIp []*string `json:"FlowIp,omitempty" xml:"FlowIp,omitempty" type:"Repeated"`
	// cdn 域名信息
	FlowUrl []*string `json:"FlowUrl,omitempty" xml:"FlowUrl,omitempty" type:"Repeated"`
	// 业务方ip集合
	OriginalIpList []*string `json:"OriginalIpList,omitempty" xml:"OriginalIpList,omitempty" type:"Repeated"`
	// 业务方域名集合
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
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 渠道ID
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// 客户侧生成的QoS请求ID，需要保证请求幂等性，确保不同请求间该参数值唯一
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	// 免流产品ID
	FlowProductId *string `json:"FlowProductId,omitempty" xml:"FlowProductId,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 是否包月，true为包月，false为不包月
	Lasting *string `json:"Lasting,omitempty" xml:"Lasting,omitempty"`
	// C端手机号
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// 客户提供的回调通知接口url
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// 取值包括cm（中国移动）/ct（中国电信）/cu（中国联通）
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// 支付订单ID
	PurchaseOrderId *string `json:"PurchaseOrderId,omitempty" xml:"PurchaseOrderId,omitempty"`
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
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果
	Data *OrderFreeFlowProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求链路ID，如POP请求进来的requestId，返回时原样返回
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务端处理耗时，ms
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// C端免流订单ID
	CustomerFlowOrderId   *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	// 执行中ordering、成功success、失败fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s OrderFreeFlowProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OrderFreeFlowProductResponseBodyData) GoString() string {
	return s.String()
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

func (client *Client) CreateApplicationInfoWithOptions(request *CreateApplicationInfoRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppTypeList)) {
		query["AppTypeList"] = request.AppTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.AppingList)) {
		query["AppingList"] = request.AppingList
	}

	if !tea.BoolValue(util.IsUnset(request.ItemCode)) {
		query["ItemCode"] = request.ItemCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
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
