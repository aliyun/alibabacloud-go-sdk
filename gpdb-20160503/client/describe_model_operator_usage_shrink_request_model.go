// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorUsageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyIdsShrink(v string) *DescribeModelOperatorUsageShrinkRequest
	GetApiKeyIdsShrink() *string
	SetEndTime(v string) *DescribeModelOperatorUsageShrinkRequest
	GetEndTime() *string
	SetGroupBy(v string) *DescribeModelOperatorUsageShrinkRequest
	GetGroupBy() *string
	SetKeysShrink(v string) *DescribeModelOperatorUsageShrinkRequest
	GetKeysShrink() *string
	SetModelNamesShrink(v string) *DescribeModelOperatorUsageShrinkRequest
	GetModelNamesShrink() *string
	SetPeriod(v int32) *DescribeModelOperatorUsageShrinkRequest
	GetPeriod() *int32
	SetStartTime(v string) *DescribeModelOperatorUsageShrinkRequest
	GetStartTime() *string
}

type DescribeModelOperatorUsageShrinkRequest struct {
	ApiKeyIdsShrink *string `json:"ApiKeyIds,omitempty" xml:"ApiKeyIds,omitempty"`
	// example:
	//
	// 2026-06-02T00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// model
	GroupBy          *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	KeysShrink       *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	ModelNamesShrink *string `json:"ModelNames,omitempty" xml:"ModelNames,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 2026-06-01T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModelOperatorUsageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetApiKeyIdsShrink() *string {
	return s.ApiKeyIdsShrink
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetKeysShrink() *string {
	return s.KeysShrink
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetModelNamesShrink() *string {
	return s.ModelNamesShrink
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetApiKeyIdsShrink(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.ApiKeyIdsShrink = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetEndTime(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetGroupBy(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetKeysShrink(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetModelNamesShrink(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.ModelNamesShrink = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetPeriod(v int32) *DescribeModelOperatorUsageShrinkRequest {
	s.Period = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetStartTime(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
