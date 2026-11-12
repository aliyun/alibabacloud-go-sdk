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
	// The list of AI invocation details on the current page. If no records exist, an empty array is returned.
	//
	// example:
	//
	// []
	CallDetails []*DescribeAICallDetailsResponseBodyCallDetails `json:"CallDetails,omitempty" xml:"CallDetails,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The actual end time of the query, in ISO 8601 format and UTC+08:00 time zone. Reuse this value when you turn pages.
	//
	// example:
	//
	// 2026-11-06T00:00:00+08:00
	QueryEndTime *string `json:"QueryEndTime,omitempty" xml:"QueryEndTime,omitempty"`
	// The actual start time of the query, in ISO 8601 format and UTC+08:00 time zone. Reuse this value when you turn pages.
	//
	// example:
	//
	// 2026-11-05T00:00:00+08:00
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	// The request ID of the current API operation, which is used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of invocation records that meet the query conditions, including records that do not match any deducted points.
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
	// The product or agent type associated with the invocation, which retains the original value from the deduction record. If no match is found or the query fails, this parameter is not returned.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// AIComputer
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The invocation status. Valid values: SUCCESS (successful, HTTP 2xx), RATE_LIMITED (throttled, HTTP 429), TIMEOUT (timed out, HTTP 504), and FAILED (other failure statuses).
	//
	// example:
	//
	// SUCCESS
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// The time of the model invocation log, in ISO 8601 format and UTC+08:00 time zone.
	//
	// example:
	//
	// 2026-11-05T12:30:00+08:00
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	// The latency of the first token, in milliseconds. If the invocation fails or no first token data exists, this parameter is empty.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 620
	FirstTokenLatencyMs *int64 `json:"FirstTokenLatencyMs,omitempty" xml:"FirstTokenLatencyMs,omitempty"`
	// The actual data region of the invocation log.
	//
	// example:
	//
	// cn-shanghai
	GatewayRegion *string `json:"GatewayRegion,omitempty" xml:"GatewayRegion,omitempty"`
	// The HTTP status code returned by the AI gateway.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The number of original input tokens. If no valid data exists, this parameter is empty.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1200
	InputToken *int64 `json:"InputToken,omitempty" xml:"InputToken,omitempty"`
	// The instance ID of the product associated with the model invocation.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// example-instance-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the actually invoked model. If the actual model is missing, the requested model is used as a fallback.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// example-model
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The actual deducted points matched, returned as a decimal string. A true zero value is returned as 0. If no match is found or the query fails, this parameter is not returned. Do not treat a missing value as zero.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0.28896
	Points *string `json:"Points,omitempty" xml:"Points,omitempty"`
	// The point breakdown of the matched deduction record for this invocation, which comes from the same record as Points. If valid details are missing, this object is not returned. If the object exists, all five point types are returned. The sum of the five point types equals Points, and a true zero value is returned as the string 0.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"NormalInput":"0.02","Output":"0.1","ImplicitCacheHit":"0","ExplicitCacheCreation":"0.15","ExplicitCacheHit":"0.01896"}
	PointsBreakdown *DescribeAICallDetailsResponseBodyCallDetailsPointsBreakdown `json:"PointsBreakdown,omitempty" xml:"PointsBreakdown,omitempty" type:"Struct"`
	// The request ID of the AI gateway invocation, which is different from the request ID of the outer API operation.
	//
	// example:
	//
	// model-call-example-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the user who initiated the model invocation. If no valid data exists, this parameter is empty.
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
	// The explicit cache creation points, in points, represented as a decimal string. A true zero value is 0.
	//
	// example:
	//
	// 0.15
	ExplicitCacheCreation *string `json:"ExplicitCacheCreation,omitempty" xml:"ExplicitCacheCreation,omitempty"`
	// The explicit cache hit points, in points, represented as a decimal string. A true zero value is 0.
	//
	// example:
	//
	// 0.01896
	ExplicitCacheHit *string `json:"ExplicitCacheHit,omitempty" xml:"ExplicitCacheHit,omitempty"`
	// The implicit cache hit points, in points, represented as a decimal string. A true zero value is 0.
	//
	// example:
	//
	// 0
	ImplicitCacheHit *string `json:"ImplicitCacheHit,omitempty" xml:"ImplicitCacheHit,omitempty"`
	// The normal input points, in points, represented as a decimal string. A true zero value is 0.
	//
	// example:
	//
	// 0.02
	NormalInput *string `json:"NormalInput,omitempty" xml:"NormalInput,omitempty"`
	// The output points, in points, represented as a decimal string. A true zero value is 0.
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
