// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopAttackListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDomainTopAttackListRequest
	GetEndTime() *int64
	SetInterval(v int64) *DescribeDomainTopAttackListRequest
	GetInterval() *int64
	SetResourceGroupId(v string) *DescribeDomainTopAttackListRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainTopAttackListRequest
	GetStartTime() *int64
}

type DescribeDomainTopAttackListRequest struct {
	// The end of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the hour or half hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query data. Unit: seconds.
	//
	// - Valid values: 300, 3600, 14400, 28800, and 86400.
	//
	// - If you do not specify this parameter or specify an unsupported value, the default value 3600 is used.
	//
	// example:
	//
	// 3600
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro and Anti-DDoS Premium instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the hour or half hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainTopAttackListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopAttackListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopAttackListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainTopAttackListRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeDomainTopAttackListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainTopAttackListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainTopAttackListRequest) SetEndTime(v int64) *DescribeDomainTopAttackListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopAttackListRequest) SetInterval(v int64) *DescribeDomainTopAttackListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainTopAttackListRequest) SetResourceGroupId(v string) *DescribeDomainTopAttackListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainTopAttackListRequest) SetStartTime(v int64) *DescribeDomainTopAttackListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopAttackListRequest) Validate() error {
	return dara.Validate(s)
}
