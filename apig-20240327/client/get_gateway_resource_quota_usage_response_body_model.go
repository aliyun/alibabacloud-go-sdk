// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayResourceQuotaUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayResourceQuotaUsageResponseBody
	GetCode() *string
	SetData(v *GetGatewayResourceQuotaUsageResponseBodyData) *GetGatewayResourceQuotaUsageResponseBody
	GetData() *GetGatewayResourceQuotaUsageResponseBodyData
	SetMessage(v string) *GetGatewayResourceQuotaUsageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayResourceQuotaUsageResponseBody
	GetRequestId() *string
}

type GetGatewayResourceQuotaUsageResponseBody struct {
	// 业务响应码。成功时为 Ok；失败时为具体错误码，应结合 HTTP 状态码处理。
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 目标网关及其九项资源配额观测结果。示例数值只说明格式，实际有效上限以本次响应为准。
	//
	// example:
	//
	// {"gatewayId":"gw-d5be0s9q5z6f1234567g","observedAt":"2026-09-17T08:00:00Z","items":[{"quotaKey":"Route","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":12,"limit":800},{"quotaKey":"ConsumerAuthorizationRule","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":20,"limit":1000},{"quotaKey":"McpServer","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":false},{"quotaKey":"Domain","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":2,"limit":100},{"quotaKey":"Service","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":5,"limit":800},{"quotaKey":"ConsumerQuotaRule","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":0,"limit":1000},{"quotaKey":"K8sServiceSource","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":1,"limit":3},{"quotaKey":"InstalledPlugin","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":2,"limit":5},{"quotaKey":"CustomPlugin","usedScope":"ACCOUNT_REGION","limitScope":"ACCOUNT_GATEWAY_TYPE","applicable":false}]}
	Data *GetGatewayResourceQuotaUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 失败时返回的错误说明，成功响应通常省略本字段。示例为错误信息，不是成功响应。
	//
	// example:
	//
	// The specified gateway does not exist.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求的唯一标识，用于排查问题。
	//
	// example:
	//
	// D0A6A1A0-8793-4C10-AB1A-2C03E770A912
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayResourceQuotaUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResourceQuotaUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayResourceQuotaUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayResourceQuotaUsageResponseBody) GetData() *GetGatewayResourceQuotaUsageResponseBodyData {
	return s.Data
}

func (s *GetGatewayResourceQuotaUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayResourceQuotaUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayResourceQuotaUsageResponseBody) SetCode(v string) *GetGatewayResourceQuotaUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBody) SetData(v *GetGatewayResourceQuotaUsageResponseBodyData) *GetGatewayResourceQuotaUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBody) SetMessage(v string) *GetGatewayResourceQuotaUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBody) SetRequestId(v string) *GetGatewayResourceQuotaUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayResourceQuotaUsageResponseBodyData struct {
	// 本次查询的目标网关唯一标识，与请求路径 gatewayId 一致。
	//
	// example:
	//
	// gw-d5be0s9q5z6f1234567g
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// 固定返回九项，每个 quotaKey 仅出现一次。不适用或暂不展示的项仍保留，applicable=false 且省略 used/limit。不返回剩余额度或百分比；示例数值不是固定默认上限。
	//
	// example:
	//
	// [{"quotaKey":"Route","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":12,"limit":800},{"quotaKey":"ConsumerAuthorizationRule","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":20,"limit":1000},{"quotaKey":"McpServer","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":false},{"quotaKey":"Domain","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":2,"limit":100},{"quotaKey":"Service","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":5,"limit":800},{"quotaKey":"ConsumerQuotaRule","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":0,"limit":1000},{"quotaKey":"K8sServiceSource","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":1,"limit":3},{"quotaKey":"InstalledPlugin","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":2,"limit":5},{"quotaKey":"CustomPlugin","usedScope":"ACCOUNT_REGION","limitScope":"ACCOUNT_GATEWAY_TYPE","applicable":false}]
	Items []*GetGatewayResourceQuotaUsageResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 服务端完成本次统计的 UTC 时间，格式为 RFC 3339，可包含小数秒。各来源独立读取，不保证跨来源瞬时原子快照。
	//
	// example:
	//
	// 2026-09-17T08:00:00Z
	ObservedAt *string `json:"observedAt,omitempty" xml:"observedAt,omitempty"`
}

func (s GetGatewayResourceQuotaUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResourceQuotaUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) GetItems() []*GetGatewayResourceQuotaUsageResponseBodyDataItems {
	return s.Items
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) SetGatewayId(v string) *GetGatewayResourceQuotaUsageResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) SetItems(v []*GetGatewayResourceQuotaUsageResponseBodyDataItems) *GetGatewayResourceQuotaUsageResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) SetObservedAt(v string) *GetGatewayResourceQuotaUsageResponseBodyData {
	s.ObservedAt = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGatewayResourceQuotaUsageResponseBodyDataItems struct {
	// true 时返回 used 和 limit；false 时省略二者。CustomPlugin 当前固定为 false，仅表示暂时隐藏配额展示，不影响插件上传、安装或既有配额校验。其他项按网关能力及有效额度判定，不能仅凭本字段推断写入操作是否允许。
	//
	// example:
	//
	// true
	Applicable *bool `json:"applicable,omitempty" xml:"applicable,omitempty"`
	// 当前生效的非负整数配额上限，单位与 used 相同，已考虑现有配置、加白及适用的购买额度。仅 applicable=true 时返回，包括合法零值；示例不是所有网关的固定上限。
	//
	// example:
	//
	// 800
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// GATEWAY 表示当前网关；ACCOUNT_GATEWAY_TYPE 表示当前账号、地域及网关类型的范围。仅 CustomPlugin 保留 ACCOUNT_GATEWAY_TYPE 标识；该项当前不读取或返回上限。
	//
	// example:
	//
	// GATEWAY
	LimitScope *string `json:"limitScope,omitempty" xml:"limitScope,omitempty"`
	// 指标标识：Route（路由）、ConsumerAuthorizationRule（消费者授权规则）、McpServer（MCP Server）、Domain（域名）、Service（服务）、ConsumerQuotaRule（消费者配额规则）、K8sServiceSource（K8s 服务来源）、InstalledPlugin（已安装插件）、CustomPlugin（自定义插件）。
	//
	// example:
	//
	// Route
	QuotaKey *string `json:"quotaKey,omitempty" xml:"quotaKey,omitempty"`
	// 实际占用数量，为非负整数，单位与 quotaKey 对应。仅 applicable=true 时返回；零值为真实零用量，历史超额可大于 limit，不截断。
	//
	// example:
	//
	// 12
	Used *int64 `json:"used,omitempty" xml:"used,omitempty"`
	// GATEWAY 表示当前网关；ACCOUNT_REGION 表示当前账号在当前地域的共享范围。仅 CustomPlugin 保留 ACCOUNT_REGION 标识；该项当前不读取或返回用量。
	//
	// example:
	//
	// GATEWAY
	UsedScope *string `json:"usedScope,omitempty" xml:"usedScope,omitempty"`
}

func (s GetGatewayResourceQuotaUsageResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResourceQuotaUsageResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetApplicable() *bool {
	return s.Applicable
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetLimit() *int64 {
	return s.Limit
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetLimitScope() *string {
	return s.LimitScope
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetQuotaKey() *string {
	return s.QuotaKey
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetUsed() *int64 {
	return s.Used
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetUsedScope() *string {
	return s.UsedScope
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetApplicable(v bool) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.Applicable = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetLimit(v int64) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.Limit = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetLimitScope(v string) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.LimitScope = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetQuotaKey(v string) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.QuotaKey = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetUsed(v int64) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.Used = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetUsedScope(v string) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.UsedScope = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
