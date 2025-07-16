// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEsExecuteDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeEsExecuteDataRequest
	GetEndTime() *string
	SetRuleId(v string) *DescribeEsExecuteDataRequest
	GetRuleId() *string
	SetStartTime(v string) *DescribeEsExecuteDataRequest
	GetStartTime() *string
}

type DescribeEsExecuteDataRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-02-18T20:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the rule. You can call the [DescribeCdnDomainConfigs](https://help.aliyun.com/document_detail/90924.html) operation to query script IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 212896**
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-02-17T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEsExecuteDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEsExecuteDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEsExecuteDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEsExecuteDataRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeEsExecuteDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEsExecuteDataRequest) SetEndTime(v string) *DescribeEsExecuteDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEsExecuteDataRequest) SetRuleId(v string) *DescribeEsExecuteDataRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeEsExecuteDataRequest) SetStartTime(v string) *DescribeEsExecuteDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEsExecuteDataRequest) Validate() error {
	return dara.Validate(s)
}
