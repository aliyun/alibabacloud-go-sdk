// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyIds(v []*int32) *DescribeModelOperatorUsageRequest
	GetApiKeyIds() []*int32
	SetEndTime(v string) *DescribeModelOperatorUsageRequest
	GetEndTime() *string
	SetGroupBy(v string) *DescribeModelOperatorUsageRequest
	GetGroupBy() *string
	SetKeys(v []*string) *DescribeModelOperatorUsageRequest
	GetKeys() []*string
	SetModelNames(v []*string) *DescribeModelOperatorUsageRequest
	GetModelNames() []*string
	SetPeriod(v int32) *DescribeModelOperatorUsageRequest
	GetPeriod() *int32
	SetStartTime(v string) *DescribeModelOperatorUsageRequest
	GetStartTime() *string
}

type DescribeModelOperatorUsageRequest struct {
	ApiKeyIds []*int32 `json:"ApiKeyIds,omitempty" xml:"ApiKeyIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-06-02T00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// model
	GroupBy    *string   `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	Keys       []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	ModelNames []*string `json:"ModelNames,omitempty" xml:"ModelNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 2026-06-01T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModelOperatorUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorUsageRequest) GetApiKeyIds() []*int32 {
	return s.ApiKeyIds
}

func (s *DescribeModelOperatorUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeModelOperatorUsageRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeModelOperatorUsageRequest) GetKeys() []*string {
	return s.Keys
}

func (s *DescribeModelOperatorUsageRequest) GetModelNames() []*string {
	return s.ModelNames
}

func (s *DescribeModelOperatorUsageRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeModelOperatorUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeModelOperatorUsageRequest) SetApiKeyIds(v []*int32) *DescribeModelOperatorUsageRequest {
	s.ApiKeyIds = v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetEndTime(v string) *DescribeModelOperatorUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetGroupBy(v string) *DescribeModelOperatorUsageRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetKeys(v []*string) *DescribeModelOperatorUsageRequest {
	s.Keys = v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetModelNames(v []*string) *DescribeModelOperatorUsageRequest {
	s.ModelNames = v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetPeriod(v int32) *DescribeModelOperatorUsageRequest {
	s.Period = &v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetStartTime(v string) *DescribeModelOperatorUsageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeModelOperatorUsageRequest) Validate() error {
	return dara.Validate(s)
}
