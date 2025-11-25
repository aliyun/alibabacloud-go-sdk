// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormalPrivateDnsCount(v int64) *DescribePrivateDnsStatisticsResponseBody
	GetAbnormalPrivateDnsCount() *int64
	SetCreatedPrivateDnsCount(v int64) *DescribePrivateDnsStatisticsResponseBody
	GetCreatedPrivateDnsCount() *int64
	SetDomainNameTotalCount(v int64) *DescribePrivateDnsStatisticsResponseBody
	GetDomainNameTotalCount() *int64
	SetNewDomainNameTotalCount(v int64) *DescribePrivateDnsStatisticsResponseBody
	GetNewDomainNameTotalCount() *int64
	SetNormalPrivateDnsCount(v int64) *DescribePrivateDnsStatisticsResponseBody
	GetNormalPrivateDnsCount() *int64
	SetPrivateDnsRegionList(v []*DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) *DescribePrivateDnsStatisticsResponseBody
	GetPrivateDnsRegionList() []*DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList
	SetRequestId(v string) *DescribePrivateDnsStatisticsResponseBody
	GetRequestId() *string
}

type DescribePrivateDnsStatisticsResponseBody struct {
	// example:
	//
	// 12
	AbnormalPrivateDnsCount *int64 `json:"AbnormalPrivateDnsCount,omitempty" xml:"AbnormalPrivateDnsCount,omitempty"`
	// example:
	//
	// 6
	CreatedPrivateDnsCount *int64 `json:"CreatedPrivateDnsCount,omitempty" xml:"CreatedPrivateDnsCount,omitempty"`
	// example:
	//
	// 5
	DomainNameTotalCount *int64 `json:"DomainNameTotalCount,omitempty" xml:"DomainNameTotalCount,omitempty"`
	// example:
	//
	// 2
	NewDomainNameTotalCount *int64 `json:"NewDomainNameTotalCount,omitempty" xml:"NewDomainNameTotalCount,omitempty"`
	// example:
	//
	// 21
	NormalPrivateDnsCount *int64                                                          `json:"NormalPrivateDnsCount,omitempty" xml:"NormalPrivateDnsCount,omitempty"`
	PrivateDnsRegionList  []*DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList `json:"PrivateDnsRegionList,omitempty" xml:"PrivateDnsRegionList,omitempty" type:"Repeated"`
	// example:
	//
	// 5716ED52-1B82-5DE1-8695-EFEC453D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrivateDnsStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsStatisticsResponseBody) GetAbnormalPrivateDnsCount() *int64 {
	return s.AbnormalPrivateDnsCount
}

func (s *DescribePrivateDnsStatisticsResponseBody) GetCreatedPrivateDnsCount() *int64 {
	return s.CreatedPrivateDnsCount
}

func (s *DescribePrivateDnsStatisticsResponseBody) GetDomainNameTotalCount() *int64 {
	return s.DomainNameTotalCount
}

func (s *DescribePrivateDnsStatisticsResponseBody) GetNewDomainNameTotalCount() *int64 {
	return s.NewDomainNameTotalCount
}

func (s *DescribePrivateDnsStatisticsResponseBody) GetNormalPrivateDnsCount() *int64 {
	return s.NormalPrivateDnsCount
}

func (s *DescribePrivateDnsStatisticsResponseBody) GetPrivateDnsRegionList() []*DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList {
	return s.PrivateDnsRegionList
}

func (s *DescribePrivateDnsStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrivateDnsStatisticsResponseBody) SetAbnormalPrivateDnsCount(v int64) *DescribePrivateDnsStatisticsResponseBody {
	s.AbnormalPrivateDnsCount = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBody) SetCreatedPrivateDnsCount(v int64) *DescribePrivateDnsStatisticsResponseBody {
	s.CreatedPrivateDnsCount = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBody) SetDomainNameTotalCount(v int64) *DescribePrivateDnsStatisticsResponseBody {
	s.DomainNameTotalCount = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBody) SetNewDomainNameTotalCount(v int64) *DescribePrivateDnsStatisticsResponseBody {
	s.NewDomainNameTotalCount = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBody) SetNormalPrivateDnsCount(v int64) *DescribePrivateDnsStatisticsResponseBody {
	s.NormalPrivateDnsCount = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBody) SetPrivateDnsRegionList(v []*DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) *DescribePrivateDnsStatisticsResponseBody {
	s.PrivateDnsRegionList = v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBody) SetRequestId(v string) *DescribePrivateDnsStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBody) Validate() error {
	if s.PrivateDnsRegionList != nil {
		for _, item := range s.PrivateDnsRegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList struct {
	// example:
	//
	// 10
	DomainNameCount *int64 `json:"DomainNameCount,omitempty" xml:"DomainNameCount,omitempty"`
	// example:
	//
	// 1
	NewDomainNameCount *int64 `json:"NewDomainNameCount,omitempty" xml:"NewDomainNameCount,omitempty"`
	// example:
	//
	// 1
	PrivateDnsCount *int64 `json:"PrivateDnsCount,omitempty" xml:"PrivateDnsCount,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) GetDomainNameCount() *int64 {
	return s.DomainNameCount
}

func (s *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) GetNewDomainNameCount() *int64 {
	return s.NewDomainNameCount
}

func (s *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) GetPrivateDnsCount() *int64 {
	return s.PrivateDnsCount
}

func (s *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) SetDomainNameCount(v int64) *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList {
	s.DomainNameCount = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) SetNewDomainNameCount(v int64) *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList {
	s.NewDomainNameCount = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) SetPrivateDnsCount(v int64) *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList {
	s.PrivateDnsCount = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) SetRegionNo(v string) *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList {
	s.RegionNo = &v
	return s
}

func (s *DescribePrivateDnsStatisticsResponseBodyPrivateDnsRegionList) Validate() error {
	return dara.Validate(s)
}
