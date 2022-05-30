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

type ChangeApplicationInfoRequest struct {
	// 阿里UID
	AliUid *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic|static
	AppTypeList *string `json:"AppTypeList,omitempty" xml:"AppTypeList,omitempty"`
	// [
	AppingList *string `json:"AppingList,omitempty" xml:"AppingList,omitempty"`
	// 商品code
	ItemCode *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
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
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
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
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
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

type GetFreeFlowStatusRequest struct {
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 免流产品ID
	FlowProductId *string `json:"FlowProductId,omitempty" xml:"FlowProductId,omitempty"`
	// C端手机号
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// 网络类型，如3G、4G、5G
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// 取值包括cm（中国移动）/ct（中国电信）/cu（中国联通）
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// 手机端私网ip地址
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// C端手机在运营商侧端伪码，如："75D35971BD"
	PseudoCode *string `json:"PseudoCode,omitempty" xml:"PseudoCode,omitempty"`
	// 手机端公网ip地址
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// 通过云通信获取的token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetFreeFlowStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowStatusRequest) GoString() string {
	return s.String()
}

func (s *GetFreeFlowStatusRequest) SetAliUid(v int64) *GetFreeFlowStatusRequest {
	s.AliUid = &v
	return s
}

func (s *GetFreeFlowStatusRequest) SetAppId(v string) *GetFreeFlowStatusRequest {
	s.AppId = &v
	return s
}

func (s *GetFreeFlowStatusRequest) SetFlowProductId(v string) *GetFreeFlowStatusRequest {
	s.FlowProductId = &v
	return s
}

func (s *GetFreeFlowStatusRequest) SetMobileNumber(v string) *GetFreeFlowStatusRequest {
	s.MobileNumber = &v
	return s
}

func (s *GetFreeFlowStatusRequest) SetNetType(v string) *GetFreeFlowStatusRequest {
	s.NetType = &v
	return s
}

func (s *GetFreeFlowStatusRequest) SetOperator(v string) *GetFreeFlowStatusRequest {
	s.Operator = &v
	return s
}

func (s *GetFreeFlowStatusRequest) SetPrivateIP(v string) *GetFreeFlowStatusRequest {
	s.PrivateIP = &v
	return s
}

func (s *GetFreeFlowStatusRequest) SetPseudoCode(v string) *GetFreeFlowStatusRequest {
	s.PseudoCode = &v
	return s
}

func (s *GetFreeFlowStatusRequest) SetPublicIP(v string) *GetFreeFlowStatusRequest {
	s.PublicIP = &v
	return s
}

func (s *GetFreeFlowStatusRequest) SetToken(v string) *GetFreeFlowStatusRequest {
	s.Token = &v
	return s
}

type GetFreeFlowStatusResponseBody struct {
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求链路ID，如POP请求进来的requestId，返回时原样返回
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务端处理耗时，ms
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFreeFlowStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetFreeFlowStatusResponseBody) SetCode(v string) *GetFreeFlowStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetFreeFlowStatusResponseBody) SetData(v interface{}) *GetFreeFlowStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetFreeFlowStatusResponseBody) SetMessage(v string) *GetFreeFlowStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetFreeFlowStatusResponseBody) SetRequestId(v string) *GetFreeFlowStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFreeFlowStatusResponseBody) SetRt(v int64) *GetFreeFlowStatusResponseBody {
	s.Rt = &v
	return s
}

func (s *GetFreeFlowStatusResponseBody) SetSuccess(v bool) *GetFreeFlowStatusResponseBody {
	s.Success = &v
	return s
}

type GetFreeFlowStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFreeFlowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFreeFlowStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFreeFlowStatusResponse) GoString() string {
	return s.String()
}

func (s *GetFreeFlowStatusResponse) SetHeaders(v map[string]*string) *GetFreeFlowStatusResponse {
	s.Headers = v
	return s
}

func (s *GetFreeFlowStatusResponse) SetStatusCode(v int32) *GetFreeFlowStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFreeFlowStatusResponse) SetBody(v *GetFreeFlowStatusResponseBody) *GetFreeFlowStatusResponse {
	s.Body = v
	return s
}

type GetFreeFlowUsageRequest struct {
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 当前页
	CurPageNum *int32 `json:"CurPageNum,omitempty" xml:"CurPageNum,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 查询月份
	Month *string `json:"Month,omitempty" xml:"Month,omitempty"`
	// 每页的记录数量
	NumPerPage *int32 `json:"NumPerPage,omitempty" xml:"NumPerPage,omitempty"`
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
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果
	Data *GetFreeFlowUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求链路ID，如POP请求进来的requestId，返回时原样返回
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务端处理耗时，ms
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// 当前页数
	CurPageNum *int32 `json:"CurPageNum,omitempty" xml:"CurPageNum,omitempty"`
	// C端用户列表
	CustomerList []*GetFreeFlowUsageResponseBodyDataCustomerList `json:"CustomerList,omitempty" xml:"CustomerList,omitempty" type:"Repeated"`
	// 是否有下一页
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// 是否有上一页
	HasPrev *bool `json:"HasPrev,omitempty" xml:"HasPrev,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 每页的记录条数
	NumPerPage *int32 `json:"NumPerPage,omitempty" xml:"NumPerPage,omitempty"`
	// 总记录条数
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// 总页数
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
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
	// 购买渠道
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// C端产品失效时间
	CustomerEndTime     *string `json:"CustomerEndTime,omitempty" xml:"CustomerEndTime,omitempty"`
	CustomerFlowOrderId *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
	// C端免流状态，取值包括create/working/expiration
	CustomerFlowStatus *string `json:"CustomerFlowStatus,omitempty" xml:"CustomerFlowStatus,omitempty"`
	// C端产品开通时间
	CustomerOpenTime *string `json:"CustomerOpenTime,omitempty" xml:"CustomerOpenTime,omitempty"`
	// C端产品生效时间
	CustomerStartTime *string `json:"CustomerStartTime,omitempty" xml:"CustomerStartTime,omitempty"`
	// 免流产品ID
	FlowProductId *string `json:"FlowProductId,omitempty" xml:"FlowProductId,omitempty"`
	// 免流产品名
	FlowProductName *string `json:"FlowProductName,omitempty" xml:"FlowProductName,omitempty"`
	// 是否包月，true或false
	IsLasting *bool `json:"IsLasting,omitempty" xml:"IsLasting,omitempty"`
	// C端手机号
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// 该流量包的计量单元数
	UnitNum *int32 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
	// 流量包价格
	UnitPrice *int32 `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
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
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果
	Data []*GetFreeFlowUsageStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求链路ID，如POP请求进来的requestId，返回时原样返回
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务端处理耗时，ms
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 该实例对应的包类型
	SpecType   *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TotalMoney *string `json:"TotalMoney,omitempty" xml:"TotalMoney,omitempty"`
	// 该实例的订单总数
	TotalOrderNumber *int64 `json:"TotalOrderNumber,omitempty" xml:"TotalOrderNumber,omitempty"`
	// 该实例的订单总计量单元数
	TotalUnitNumber *int64 `json:"TotalUnitNumber,omitempty" xml:"TotalUnitNumber,omitempty"`
	// 产品规格
	YunOutProduct *string `json:"YunOutProduct,omitempty" xml:"YunOutProduct,omitempty"`
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

type GetOrderFreeFlowProductStatusRequest struct {
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 在订购接口2.1.9中引擎侧生成的id
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
	// 结果码
	Code *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetOrderFreeFlowProductStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求链路ID，如POP请求进来的requestId，返回时原样返回
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务端处理耗时，ms
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// C端免流订单ID
	CustomerFlowOrderId   *string `json:"CustomerFlowOrderId,omitempty" xml:"CustomerFlowOrderId,omitempty"`
	CustomerFlowRequestId *string `json:"CustomerFlowRequestId,omitempty" xml:"CustomerFlowRequestId,omitempty"`
	// status为fail时，描述fail的具体原因
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// 执行中ordering、成功success、失败fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

type SaveApplicationInfoRequest struct {
	// 阿里UID
	AliUid *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic|static
	AppTypeList *string `json:"AppTypeList,omitempty" xml:"AppTypeList,omitempty"`
	// [
	AppingList *string `json:"AppingList,omitempty" xml:"AppingList,omitempty"`
	// 商品code
	ItemCode *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
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

type ValidateStatusRequest struct {
	// 阿里UID
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 应用名称
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// C端手机号
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// 取值包括cm（中国移动）/ct（中国电信）/cu（中国联通）
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// 运营商伪码
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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

func (s *ValidateStatusRequest) SetMobileNumber(v string) *ValidateStatusRequest {
	s.MobileNumber = &v
	return s
}

func (s *ValidateStatusRequest) SetOperator(v string) *ValidateStatusRequest {
	s.Operator = &v
	return s
}

func (s *ValidateStatusRequest) SetProductCode(v string) *ValidateStatusRequest {
	s.ProductCode = &v
	return s
}

type ValidateStatusResponseBody struct {
	// 结果码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果
	Data *ValidateStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 结果描述
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求链路ID，如POP请求进来的requestId，返回时原样返回
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务端处理耗时，ms
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// 是否处于免流状态，取值范围为true/false
	FreeFlow *bool `json:"FreeFlow,omitempty" xml:"FreeFlow,omitempty"`
	// 伪码
	PseudoCode *string `json:"PseudoCode,omitempty" xml:"PseudoCode,omitempty"`
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

func (client *Client) GetFreeFlowStatusWithOptions(request *GetFreeFlowStatusRequest, runtime *util.RuntimeOptions) (_result *GetFreeFlowStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFreeFlowStatus"),
		Version:     tea.String("2022-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFreeFlowStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFreeFlowStatus(request *GetFreeFlowStatusRequest) (_result *GetFreeFlowStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFreeFlowStatusResponse{}
	_body, _err := client.GetFreeFlowStatusWithOptions(request, runtime)
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
