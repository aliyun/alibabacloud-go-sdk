// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAICallDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeAICallDetailsRequest
	GetBizType() *string
	SetEndTime(v string) *DescribeAICallDetailsRequest
	GetEndTime() *string
	SetGatewayRegion(v string) *DescribeAICallDetailsRequest
	GetGatewayRegion() *string
	SetModelId(v string) *DescribeAICallDetailsRequest
	GetModelId() *string
	SetPageNumber(v int32) *DescribeAICallDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAICallDetailsRequest
	GetPageSize() *int32
	SetSearchRequestId(v string) *DescribeAICallDetailsRequest
	GetSearchRequestId() *string
	SetStartTime(v string) *DescribeAICallDetailsRequest
	GetStartTime() *string
	SetWyIds(v []*string) *DescribeAICallDetailsRequest
	GetWyIds() []*string
}

type DescribeAICallDetailsRequest struct {
	// 业务版本，沿用售卖 BizType：ENTERPRISE 或 1 为企业版；BUSINESS 或 20 为商业版。未传默认企业版。
	//
	// example:
	//
	// ENTERPRISE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// 带时区的 ISO 8601 结束时间，整秒精度，默认当前时间；应大于开始时间且不晚于当前时间，区间左闭右开。
	//
	// example:
	//
	// 2026-11-06T00:00:00+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 网关日志地域：cn-shanghai（国内）、cn-hongkong（香港）、ap-southeast-1（新加坡）、us-east-1（美国）；兼容 cn-hangzhou，仅查询路由等同 cn-shanghai；杭州预发实际数据地域为 cn-hangzhou。不传时使用当前站点默认地域；预发只查询预发日志，国内映射杭州预发，美国预发未配置时返回 ServiceUnavailable。
	//
	// example:
	//
	// cn-shanghai
	GatewayRegion *string `json:"GatewayRegion,omitempty" xml:"GatewayRegion,omitempty"`
	// 可选请求模型 ID；不传或空白时不筛选，非空时精确匹配 custom_log.requestedModel 或 ai_log.request_model 任一字段，与 SearchRequestId、WyIds 等条件取交集。不按 ai_log.model 筛选，不支持模糊匹配。
	//
	// example:
	//
	// example-model
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// 页码，从 1 开始；分页偏移超过查询上限时返回 InvalidParameter。
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回数量，最大 50。
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 可选请求 ID；不传或空白时不筛选，非空时精确匹配 SLS request_id，与 ModelId、WyIds 等条件取交集。不得包含控制字符或管道符。
	//
	// example:
	//
	// model-call-example-001
	SearchRequestId *string `json:"SearchRequestId,omitempty" xml:"SearchRequestId,omitempty"`
	// 带时区的 ISO 8601 开始时间，整秒精度，默认 UTC+8 今日零点；不得早于当前时间按 UTC+8 回退 3 个自然月，区间不超过 30 天。
	//
	// example:
	//
	// 2026-11-05T00:00:00+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 可选用户 ID 列表，最多 50 项；不传或空数组时不筛选。非空时精确匹配 custom_log.wyId，列表内取 OR，与租户、BizType、SearchRequestId 和 ModelId 条件取 AND。每项去除首尾空白后为 1 至 128 个字符，不允许控制字符或管道符。
	//
	// example:
	//
	// ["wy-example-001"]
	WyIds []*string `json:"WyIds,omitempty" xml:"WyIds,omitempty" type:"Repeated"`
}

func (s DescribeAICallDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAICallDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAICallDetailsRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeAICallDetailsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAICallDetailsRequest) GetGatewayRegion() *string {
	return s.GatewayRegion
}

func (s *DescribeAICallDetailsRequest) GetModelId() *string {
	return s.ModelId
}

func (s *DescribeAICallDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAICallDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAICallDetailsRequest) GetSearchRequestId() *string {
	return s.SearchRequestId
}

func (s *DescribeAICallDetailsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAICallDetailsRequest) GetWyIds() []*string {
	return s.WyIds
}

func (s *DescribeAICallDetailsRequest) SetBizType(v string) *DescribeAICallDetailsRequest {
	s.BizType = &v
	return s
}

func (s *DescribeAICallDetailsRequest) SetEndTime(v string) *DescribeAICallDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAICallDetailsRequest) SetGatewayRegion(v string) *DescribeAICallDetailsRequest {
	s.GatewayRegion = &v
	return s
}

func (s *DescribeAICallDetailsRequest) SetModelId(v string) *DescribeAICallDetailsRequest {
	s.ModelId = &v
	return s
}

func (s *DescribeAICallDetailsRequest) SetPageNumber(v int32) *DescribeAICallDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAICallDetailsRequest) SetPageSize(v int32) *DescribeAICallDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAICallDetailsRequest) SetSearchRequestId(v string) *DescribeAICallDetailsRequest {
	s.SearchRequestId = &v
	return s
}

func (s *DescribeAICallDetailsRequest) SetStartTime(v string) *DescribeAICallDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAICallDetailsRequest) SetWyIds(v []*string) *DescribeAICallDetailsRequest {
	s.WyIds = v
	return s
}

func (s *DescribeAICallDetailsRequest) Validate() error {
	return dara.Validate(s)
}
