// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeReqHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDomainRealTimeReqHitRateDataResponseBodyData) *DescribeDomainRealTimeReqHitRateDataResponseBody
	GetData() *DescribeDomainRealTimeReqHitRateDataResponseBodyData
	SetRequestId(v string) *DescribeDomainRealTimeReqHitRateDataResponseBody
	GetRequestId() *string
}

type DescribeDomainRealTimeReqHitRateDataResponseBody struct {
	// The data returned.
	Data *DescribeDomainRealTimeReqHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 70A26B11-3673-479C-AEA8-E03FC5D3496D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBody) GetData() *DescribeDomainRealTimeReqHitRateDataResponseBodyData {
	return s.Data
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBody) SetData(v *DescribeDomainRealTimeReqHitRateDataResponseBodyData) *DescribeDomainRealTimeReqHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainRealTimeReqHitRateDataResponseBodyData struct {
	ReqHitRateDataModel []*DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel `json:"ReqHitRateDataModel,omitempty" xml:"ReqHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyData) GetReqHitRateDataModel() []*DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	return s.ReqHitRateDataModel
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyData) SetReqHitRateDataModel(v []*DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) *DescribeDomainRealTimeReqHitRateDataResponseBodyData {
	s.ReqHitRateDataModel = v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyData) Validate() error {
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

type DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel struct {
	// The request hit ratio.
	//
	// example:
	//
	// 0.8956940476262277
	ReqHitRate *float32 `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-02T11:26:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GetReqHitRate() *float32 {
	return s.ReqHitRate
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetReqHitRate(v float32) *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetTimeStamp(v string) *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) Validate() error {
	return dara.Validate(s)
}
