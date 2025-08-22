// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeByteHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody
	GetData() *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData
	SetRequestId(v string) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainRealTimeByteHitRateDataResponseBody struct {
	// The list of byte hit ratios.
	Data *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) GetData() *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData {
	return s.Data
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) SetData(v *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData struct {
	ByteHitRateDataModel []*DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel `json:"ByteHitRateDataModel,omitempty" xml:"ByteHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) GetByteHitRateDataModel() []*DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	return s.ByteHitRateDataModel
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) SetByteHitRateDataModel(v []*DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData {
	s.ByteHitRateDataModel = v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel struct {
	// The byte hit ratio.
	//
	// example:
	//
	// 0.8956940476262277
	ByteHitRate *float32 `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-10-20T04:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GetByteHitRate() *float32 {
	return s.ByteHitRate
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetByteHitRate(v float32) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) Validate() error {
	return dara.Validate(s)
}
