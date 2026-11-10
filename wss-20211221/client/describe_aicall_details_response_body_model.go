// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAICallDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallDetails(v []*DescribeAICallDetailsResponseBodyCallDetails) *DescribeAICallDetailsResponseBody
	GetCallDetails() []*DescribeAICallDetailsResponseBodyCallDetails
	SetPageNumber(v int32) *DescribeAICallDetailsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAICallDetailsResponseBody
	GetPageSize() *int32
	SetQueryEndTime(v string) *DescribeAICallDetailsResponseBody
	GetQueryEndTime() *string
	SetQueryStartTime(v string) *DescribeAICallDetailsResponseBody
	GetQueryStartTime() *string
	SetRequestId(v string) *DescribeAICallDetailsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAICallDetailsResponseBody
	GetTotalCount() *int64
}

type DescribeAICallDetailsResponseBody struct {
	// 调用明细列表。
	//
	// example:
	//
	// []
	CallDetails []*DescribeAICallDetailsResponseBodyCallDetails `json:"CallDetails,omitempty" xml:"CallDetails,omitempty" type:"Repeated"`
	// 当前页码。
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 当前页大小。
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 实际查询终点。翻页时应复用此值。
	//
	// example:
	//
	// 2026-11-06T00:00:00+08:00
	QueryEndTime *string `json:"QueryEndTime,omitempty" xml:"QueryEndTime,omitempty"`
	// 实际查询起点。翻页时应复用此值。
	//
	// example:
	//
	// 2026-11-05T00:00:00+08:00
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// POP 请求 ID。
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 满足条件的总记录数。
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAICallDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAICallDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAICallDetailsResponseBody) GetCallDetails() []*DescribeAICallDetailsResponseBodyCallDetails {
	return s.CallDetails
}

func (s *DescribeAICallDetailsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAICallDetailsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAICallDetailsResponseBody) GetQueryEndTime() *string {
	return s.QueryEndTime
}

func (s *DescribeAICallDetailsResponseBody) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeAICallDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAICallDetailsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAICallDetailsResponseBody) SetCallDetails(v []*DescribeAICallDetailsResponseBodyCallDetails) *DescribeAICallDetailsResponseBody {
	s.CallDetails = v
	return s
}

func (s *DescribeAICallDetailsResponseBody) SetPageNumber(v int32) *DescribeAICallDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAICallDetailsResponseBody) SetPageSize(v int32) *DescribeAICallDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAICallDetailsResponseBody) SetQueryEndTime(v string) *DescribeAICallDetailsResponseBody {
	s.QueryEndTime = &v
	return s
}

func (s *DescribeAICallDetailsResponseBody) SetQueryStartTime(v string) *DescribeAICallDetailsResponseBody {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeAICallDetailsResponseBody) SetRequestId(v string) *DescribeAICallDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAICallDetailsResponseBody) SetTotalCount(v int64) *DescribeAICallDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAICallDetailsResponseBody) Validate() error {
	if s.CallDetails != nil {
		for _, item := range s.CallDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAICallDetailsResponseBodyCallDetails struct {
	// 售卖 deduction 记录中的原始 agent_type；未匹配或 ADB 查询失败时不返回。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// AIComputer
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// 调用状态：SUCCESS、RATE_LIMITED、TIMEOUT 或 FAILED。
	//
	// example:
	//
	// SUCCESS
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// 日志时间，ISO 8601。
	//
	// example:
	//
	// 2026-11-05T12:30:00+08:00
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// 首包时延，单位毫秒；失败或无首包数据时为空。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 620
	FirstTokenLatencyMs *int64 `json:"FirstTokenLatencyMs,omitempty" xml:"FirstTokenLatencyMs,omitempty"`
	// 本条调用日志的实际数据地域，取值为 cn-shanghai、cn-hangzhou、cn-hongkong、ap-southeast-1 或 us-east-1；杭州预发返回 cn-hangzhou。
	//
	// example:
	//
	// cn-shanghai
	GatewayRegion *string `json:"GatewayRegion,omitempty" xml:"GatewayRegion,omitempty"`
	// AI 网关返回的 HTTP 状态码。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 原始输入 token 数，前端据此转换上下文规模。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1200
	InputToken *int64 `json:"InputToken,omitempty" xml:"InputToken,omitempty"`
	// 产品实例 ID。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// example-instance-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实际调用模型。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// example-model
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 按请求 ID 和输入输出 token 匹配的实际抵扣积分，十进制字符串；积分查询失败或超时时为空。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0.28896
	Points *string `json:"Points,omitempty" xml:"Points,omitempty"`
	// 本条调用已匹配抵扣记录的积分构成，与 Points 来自同一条记录。缺少有效明细时不返回此对象；对象存在时完整返回五项积分，五项之和等于 Points，真实零值返回字符串 0。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"NormalInput":"0.02","Output":"0.1","ImplicitCacheHit":"0","ExplicitCacheCreation":"0.15","ExplicitCacheHit":"0.01896"}
	PointsBreakdown *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown `json:"PointsBreakdown,omitempty" xml:"PointsBreakdown,omitempty" type:"Struct"`
	// AI 网关请求 ID，与外层 POP 请求 ID 不同。
	//
	// example:
	//
	// model-call-example-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 用户 WyId，用户名由前端组装。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// wy-example-001
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s DescribeAICallDetailsResponseBodyCallDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeAICallDetailsResponseBodyCallDetails) GoString() string {
	return s.String()
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetCallStatus() *string {
	return s.CallStatus
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetCallTime() *string {
	return s.CallTime
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetFirstTokenLatencyMs() *int64 {
	return s.FirstTokenLatencyMs
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetGatewayRegion() *string {
	return s.GatewayRegion
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetInputToken() *int64 {
	return s.InputToken
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetModel() *string {
	return s.Model
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetPoints() *string {
	return s.Points
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetPointsBreakdown() *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown {
	return s.PointsBreakdown
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) GetWyId() *string {
	return s.WyId
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetAgentType(v string) *DescribeAICallDetailsResponseBodyCallDetails {
	s.AgentType = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetCallStatus(v string) *DescribeAICallDetailsResponseBodyCallDetails {
	s.CallStatus = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetCallTime(v string) *DescribeAICallDetailsResponseBodyCallDetails {
	s.CallTime = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetFirstTokenLatencyMs(v int64) *DescribeAICallDetailsResponseBodyCallDetails {
	s.FirstTokenLatencyMs = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetGatewayRegion(v string) *DescribeAICallDetailsResponseBodyCallDetails {
	s.GatewayRegion = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetHttpStatusCode(v int32) *DescribeAICallDetailsResponseBodyCallDetails {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetInputToken(v int64) *DescribeAICallDetailsResponseBodyCallDetails {
	s.InputToken = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetInstanceId(v string) *DescribeAICallDetailsResponseBodyCallDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetModel(v string) *DescribeAICallDetailsResponseBodyCallDetails {
	s.Model = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetPoints(v string) *DescribeAICallDetailsResponseBodyCallDetails {
	s.Points = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetPointsBreakdown(v *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) *DescribeAICallDetailsResponseBodyCallDetails {
	s.PointsBreakdown = v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetRequestId(v string) *DescribeAICallDetailsResponseBodyCallDetails {
	s.RequestId = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) SetWyId(v string) *DescribeAICallDetailsResponseBodyCallDetails {
	s.WyId = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetails) Validate() error {
	if s.PointsBreakdown != nil {
		if err := s.PointsBreakdown.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown struct {
	// 显式缓存创建积分，单位为积分，使用十进制字符串表示；真实零值为 0。
	//
	// example:
	//
	// 0.15
	ExplicitCacheCreation *string `json:"ExplicitCacheCreation,omitempty" xml:"ExplicitCacheCreation,omitempty"`
	// 显式缓存命中积分，单位为积分，使用十进制字符串表示；真实零值为 0。
	//
	// example:
	//
	// 0.01896
	ExplicitCacheHit *string `json:"ExplicitCacheHit,omitempty" xml:"ExplicitCacheHit,omitempty"`
	// 隐式缓存命中积分，单位为积分，使用十进制字符串表示；真实零值为 0。
	//
	// example:
	//
	// 0
	ImplicitCacheHit *string `json:"ImplicitCacheHit,omitempty" xml:"ImplicitCacheHit,omitempty"`
	// 普通输入积分，单位为积分，使用十进制字符串表示；真实零值为 0。
	//
	// example:
	//
	// 0.02
	NormalInput *string `json:"NormalInput,omitempty" xml:"NormalInput,omitempty"`
	// 输出积分，单位为积分，使用十进制字符串表示；真实零值为 0。
	//
	// example:
	//
	// 0.1
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) String() string {
	return dara.Prettify(s)
}

func (s DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) GoString() string {
	return s.String()
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) GetExplicitCacheCreation() *string {
	return s.ExplicitCacheCreation
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) GetExplicitCacheHit() *string {
	return s.ExplicitCacheHit
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) GetImplicitCacheHit() *string {
	return s.ImplicitCacheHit
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) GetNormalInput() *string {
	return s.NormalInput
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) GetOutput() *string {
	return s.Output
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) SetExplicitCacheCreation(v string) *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown {
	s.ExplicitCacheCreation = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) SetExplicitCacheHit(v string) *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown {
	s.ExplicitCacheHit = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) SetImplicitCacheHit(v string) *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown {
	s.ImplicitCacheHit = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) SetNormalInput(v string) *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown {
	s.NormalInput = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) SetOutput(v string) *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown {
	s.Output = &v
	return s
}

func (s *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown) Validate() error {
	return dara.Validate(s)
}
