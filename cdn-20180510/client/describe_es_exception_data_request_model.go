// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEsExceptionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeEsExceptionDataRequest
	GetEndTime() *string
	SetRuleId(v string) *DescribeEsExceptionDataRequest
	GetRuleId() *string
	SetStartTime(v string) *DescribeEsExceptionDataRequest
	GetStartTime() *string
}

type DescribeEsExceptionDataRequest struct {
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-02-18T20:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The script ID. You can call the [DescribeCdnDomainConfigs](https://help.aliyun.com/document_detail/90924.html) operation to query script IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 212896**
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-02-17T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEsExceptionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEsExceptionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEsExceptionDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEsExceptionDataRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeEsExceptionDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEsExceptionDataRequest) SetEndTime(v string) *DescribeEsExceptionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEsExceptionDataRequest) SetRuleId(v string) *DescribeEsExceptionDataRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeEsExceptionDataRequest) SetStartTime(v string) *DescribeEsExceptionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEsExceptionDataRequest) Validate() error {
	return dara.Validate(s)
}
