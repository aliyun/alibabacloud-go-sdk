// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeByteHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDomainRealTimeByteHitRateDataResponseBodyData) *DescribeDomainRealTimeByteHitRateDataResponseBody
	GetData() *DescribeDomainRealTimeByteHitRateDataResponseBodyData
	SetRequestId(v string) *DescribeDomainRealTimeByteHitRateDataResponseBody
	GetRequestId() *string
}

type DescribeDomainRealTimeByteHitRateDataResponseBody struct {
	// The data returned.
	Data *DescribeDomainRealTimeByteHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 70A26B11-3673-479C-AEA8-E03FC5D3496D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBody) GetData() *DescribeDomainRealTimeByteHitRateDataResponseBodyData {
	return s.Data
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBody) SetData(v *DescribeDomainRealTimeByteHitRateDataResponseBodyData) *DescribeDomainRealTimeByteHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeByteHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainRealTimeByteHitRateDataResponseBodyData struct {
	ByteHitRateDataModel []*DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel `json:"ByteHitRateDataModel,omitempty" xml:"ByteHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyData) GetByteHitRateDataModel() []*DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	return s.ByteHitRateDataModel
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyData) SetByteHitRateDataModel(v []*DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) *DescribeDomainRealTimeByteHitRateDataResponseBodyData {
	s.ByteHitRateDataModel = v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel struct {
	// The byte hit ratio. The byte hit ratio is measured in percentage.
	//
	// example:
	//
	// 0.8956940476262277
	ByteHitRate *float32 `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GetByteHitRate() *float32 {
	return s.ByteHitRate
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetByteHitRate(v float32) *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetTimeStamp(v string) *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) Validate() error {
	return dara.Validate(s)
}
