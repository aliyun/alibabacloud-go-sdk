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
	// The business edition. Valid values: ENTERPRISE or 1 for Enterprise Edition, and BUSINESS or 20 for Business Edition. Default value: Enterprise Edition.
	//
	// example:
	//
	// ENTERPRISE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The end time of the query, exclusive. Specify the time in ISO 8601 format with time zone, accurate to seconds. Default value: the current time. The end time must be later than StartTime and cannot be later than the current time.
	//
	// example:
	//
	// 2026-11-06T00:00:00+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region of the AI gateway logs. Valid values: cn-shanghai (Shanghai), cn-hongkong (Hong Kong), ap-southeast-1 (Singapore), and us-east-1 (US East). The cn-hangzhou value is also supported and routed to cn-shanghai in the production environment. If this parameter is not specified, the default value is cn-shanghai for the China site and ap-southeast-1 for the international site.
	//
	// example:
	//
	// cn-shanghai
	GatewayRegion *string `json:"GatewayRegion,omitempty" xml:"GatewayRegion,omitempty"`
	// The requested model ID during invocation for exact match. Fuzzy match is not supported. This ID may be different from the actual executed model Model returned in the response. If this parameter is not specified or is left empty, no filtering is performed. The value can be up to 128 characters in length and cannot contain control characters or vertical bars.
	//
	// example:
	//
	// example-model
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1. The product of PageNumber and PageSize cannot exceed 1,000,000.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The model invocation request ID for exact match. This ID is different from the RequestId of the current API operation. If this parameter is not specified or is left empty, no filtering is performed. The value can be up to 128 characters in length and cannot contain control characters or vertical bars. This condition is combined with other filter conditions by using the AND operator.
	//
	// example:
	//
	// model-call-example-001
	SearchRequestId *string `json:"SearchRequestId,omitempty" xml:"SearchRequestId,omitempty"`
	// The start time of the query, inclusive. Specify the time in ISO 8601 format with time zone, accurate to seconds. Default value: 00:00:00 (UTC+8) of the current day. The start time cannot be earlier than three calendar months before the current time, and the query time range cannot exceed 30 days.
	//
	// example:
	//
	// 2026-11-05T00:00:00+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of user IDs. You can specify up to 50 user IDs. If this parameter is not specified or an empty array is passed, no filtering is performed. The user IDs in the list are combined by using the OR operator, and this condition is combined with other filter conditions by using the AND operator.
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
