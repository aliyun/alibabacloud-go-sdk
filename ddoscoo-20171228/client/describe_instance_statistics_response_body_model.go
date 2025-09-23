// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceStatistics(v []*DescribeInstanceStatisticsResponseBodyInstanceStatistics) *DescribeInstanceStatisticsResponseBody
	GetInstanceStatistics() []*DescribeInstanceStatisticsResponseBodyInstanceStatistics
	SetRequestId(v string) *DescribeInstanceStatisticsResponseBody
	GetRequestId() *string
}

type DescribeInstanceStatisticsResponseBody struct {
	InstanceStatistics []*DescribeInstanceStatisticsResponseBodyInstanceStatistics `json:"InstanceStatistics,omitempty" xml:"InstanceStatistics,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBody) GetInstanceStatistics() []*DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	return s.InstanceStatistics
}

func (s *DescribeInstanceStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceStatisticsResponseBody) SetInstanceStatistics(v []*DescribeInstanceStatisticsResponseBodyInstanceStatistics) *DescribeInstanceStatisticsResponseBody {
	s.InstanceStatistics = v
	return s
}

func (s *DescribeInstanceStatisticsResponseBody) SetRequestId(v string) *DescribeInstanceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceStatisticsResponseBodyInstanceStatistics struct {
	// example:
	//
	// 1
	DefenseCountUsage *int32 `json:"DefenseCountUsage,omitempty" xml:"DefenseCountUsage,omitempty"`
	// example:
	//
	// 10
	DomainUsage *int32 `json:"DomainUsage,omitempty" xml:"DomainUsage,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	PortUsage *int32 `json:"PortUsage,omitempty" xml:"PortUsage,omitempty"`
	// example:
	//
	// 1
	SiteUsage *int32 `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
}

func (s DescribeInstanceStatisticsResponseBodyInstanceStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBodyInstanceStatistics) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) GetDefenseCountUsage() *int32 {
	return s.DefenseCountUsage
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) GetDomainUsage() *int32 {
	return s.DomainUsage
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) GetPortUsage() *int32 {
	return s.PortUsage
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) GetSiteUsage() *int32 {
	return s.SiteUsage
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetDefenseCountUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.DefenseCountUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetDomainUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.DomainUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetInstanceId(v string) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetPortUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.PortUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetSiteUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.SiteUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) Validate() error {
	return dara.Validate(s)
}
