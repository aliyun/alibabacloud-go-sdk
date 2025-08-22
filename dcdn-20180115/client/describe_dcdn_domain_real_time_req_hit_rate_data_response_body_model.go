// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeReqHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody
	GetData() *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData
	SetRequestId(v string) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainRealTimeReqHitRateDataResponseBody struct {
	// The list of byte hit ratios.
	Data *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) GetData() *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData {
	return s.Data
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) SetData(v *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData struct {
	ReqHitRateDataModel []*DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel `json:"ReqHitRateDataModel,omitempty" xml:"ReqHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) GetReqHitRateDataModel() []*DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	return s.ReqHitRateDataModel
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) SetReqHitRateDataModel(v []*DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData {
	s.ReqHitRateDataModel = v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel struct {
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
	// 2016-10-20T04:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GetReqHitRate() *float32 {
	return s.ReqHitRate
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetReqHitRate(v float32) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) Validate() error {
	return dara.Validate(s)
}
