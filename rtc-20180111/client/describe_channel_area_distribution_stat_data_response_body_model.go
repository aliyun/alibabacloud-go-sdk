// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelAreaDistributionStatDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAreaStatList(v []*DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) *DescribeChannelAreaDistributionStatDataResponseBody
	GetAreaStatList() []*DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList
	SetRequestId(v string) *DescribeChannelAreaDistributionStatDataResponseBody
	GetRequestId() *string
}

type DescribeChannelAreaDistributionStatDataResponseBody struct {
	// The regional statistics list.
	AreaStatList []*DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList `json:"AreaStatList,omitempty" xml:"AreaStatList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataResponseBody) GetAreaStatList() []*DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	return s.AreaStatList
}

func (s *DescribeChannelAreaDistributionStatDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChannelAreaDistributionStatDataResponseBody) SetAreaStatList(v []*DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) *DescribeChannelAreaDistributionStatDataResponseBody {
	s.AreaStatList = v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBody) SetRequestId(v string) *DescribeChannelAreaDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBody) Validate() error {
	if s.AreaStatList != nil {
		for _, item := range s.AreaStatList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList struct {
	// The region name, such as China_Zhejiang Province_Hangzhou.
	//
	// example:
	//
	// China_Zhejiang Province_Hangzhou
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The number of communicating users.
	//
	// example:
	//
	// 1
	CallUserCount *int32 `json:"CallUserCount,omitempty" xml:"CallUserCount,omitempty"`
	// The high-quality transmission rate, represented as a decimal. For example, 0.9999 indicates a high-quality transmission rate of 99.99%.
	//
	// example:
	//
	// 0.9999
	HighQualityTransmissionRate *string `json:"HighQualityTransmissionRate,omitempty" xml:"HighQualityTransmissionRate,omitempty"`
	// The number of publishers.
	//
	// example:
	//
	// 1
	PubUserCount *int32 `json:"PubUserCount,omitempty" xml:"PubUserCount,omitempty"`
	// The number of subscribers.
	//
	// example:
	//
	// 1
	SubUserCount *int32 `json:"SubUserCount,omitempty" xml:"SubUserCount,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) GetAreaName() *string {
	return s.AreaName
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) GetCallUserCount() *int32 {
	return s.CallUserCount
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) GetHighQualityTransmissionRate() *string {
	return s.HighQualityTransmissionRate
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) GetPubUserCount() *int32 {
	return s.PubUserCount
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) GetSubUserCount() *int32 {
	return s.SubUserCount
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetAreaName(v string) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.AreaName = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetCallUserCount(v int32) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.CallUserCount = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetHighQualityTransmissionRate(v string) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.HighQualityTransmissionRate = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetPubUserCount(v int32) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.PubUserCount = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetSubUserCount(v int32) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.SubUserCount = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) Validate() error {
	return dara.Validate(s)
}
