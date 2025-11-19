// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeReqHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeVodDomainRealTimeReqHitRateDataResponseBodyData) *DescribeVodDomainRealTimeReqHitRateDataResponseBody
	GetData() *DescribeVodDomainRealTimeReqHitRateDataResponseBodyData
	SetRequestId(v string) *DescribeVodDomainRealTimeReqHitRateDataResponseBody
	GetRequestId() *string
}

type DescribeVodDomainRealTimeReqHitRateDataResponseBody struct {
	// The returned results.
	Data *DescribeVodDomainRealTimeReqHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 70A26B11-3673-479C-AEA8-E03FC5D3496D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainRealTimeReqHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBody) GetData() *DescribeVodDomainRealTimeReqHitRateDataResponseBodyData {
	return s.Data
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBody) SetData(v *DescribeVodDomainRealTimeReqHitRateDataResponseBodyData) *DescribeVodDomainRealTimeReqHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBody) SetRequestId(v string) *DescribeVodDomainRealTimeReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainRealTimeReqHitRateDataResponseBodyData struct {
	ReqHitRateDataModel []*DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel `json:"ReqHitRateDataModel,omitempty" xml:"ReqHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainRealTimeReqHitRateDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeReqHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBodyData) GetReqHitRateDataModel() []*DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	return s.ReqHitRateDataModel
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBodyData) SetReqHitRateDataModel(v []*DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) *DescribeVodDomainRealTimeReqHitRateDataResponseBodyData {
	s.ReqHitRateDataModel = v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBodyData) Validate() error {
	if s.ReqHitRateDataModel != nil {
		for _, item := range s.ReqHitRateDataModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel struct {
	// The cache hit ratio that is calculated based on requests. The cache hit ratio is measured in percentage.
	//
	// example:
	//
	// 0.8956940476262277
	ReqHitRate *float32 `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-02T11:26:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GetReqHitRate() *float32 {
	return s.ReqHitRate
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetReqHitRate(v float32) *DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetTimeStamp(v string) *DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) Validate() error {
	return dara.Validate(s)
}
