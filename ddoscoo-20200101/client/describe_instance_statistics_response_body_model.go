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
	// The statistics on the instance.
	InstanceStatistics []*DescribeInstanceStatisticsResponseBodyInstanceStatistics `json:"InstanceStatistics,omitempty" xml:"InstanceStatistics,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 642319A9-D1F2-4459-A447-E57CFC599FDE
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
	if s.InstanceStatistics != nil {
		for _, item := range s.InstanceStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceStatisticsResponseBodyInstanceStatistics struct {
	// The number of advanced mitigation sessions that are used in this month.
	//
	// >  This parameter is returned only if Anti-DDoS Proxy (Outside Chinese Mainland) instances are queried.
	//
	// example:
	//
	// 1
	DefenseCountUsage *int32 `json:"DefenseCountUsage,omitempty" xml:"DefenseCountUsage,omitempty"`
	// The number of domain names that are protected by the instance.
	//
	// example:
	//
	// 1
	DomainUsage *int32 `json:"DomainUsage,omitempty" xml:"DomainUsage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of ports that are protected by the instance.
	//
	// example:
	//
	// 2
	PortUsage *int32 `json:"PortUsage,omitempty" xml:"PortUsage,omitempty"`
	// The number of websites that are protected by the instance.
	//
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
