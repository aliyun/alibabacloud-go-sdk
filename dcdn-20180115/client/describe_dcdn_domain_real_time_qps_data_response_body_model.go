// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeQpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDcdnDomainRealTimeQpsDataResponseBodyData) *DescribeDcdnDomainRealTimeQpsDataResponseBody
	GetData() *DescribeDcdnDomainRealTimeQpsDataResponseBodyData
	SetRequestId(v string) *DescribeDcdnDomainRealTimeQpsDataResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainRealTimeQpsDataResponseBody struct {
	// The information about the backup set.
	Data *DescribeDcdnDomainRealTimeQpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 32DC9806-E9F9-4490-BBDC-B3A9E32FCC1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBody) GetData() *DescribeDcdnDomainRealTimeQpsDataResponseBodyData {
	return s.Data
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBody) SetData(v *DescribeDcdnDomainRealTimeQpsDataResponseBodyData) *DescribeDcdnDomainRealTimeQpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeQpsDataResponseBodyData struct {
	QpsModel []*DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel `json:"QpsModel,omitempty" xml:"QpsModel,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyData) GetQpsModel() []*DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel {
	return s.QpsModel
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyData) SetQpsModel(v []*DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) *DescribeDcdnDomainRealTimeQpsDataResponseBodyData {
	s.QpsModel = v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel struct {
	// The number of queries per second (QPS).
	//
	// example:
	//
	// 1851.25
	Qps *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-02T11:26:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) GetQps() *float32 {
	return s.Qps
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) SetQps(v float32) *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.Qps = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeQpsDataResponseBodyDataQpsModel) Validate() error {
	return dara.Validate(s)
}
