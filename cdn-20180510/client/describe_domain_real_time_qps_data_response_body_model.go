// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeQpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDomainRealTimeQpsDataResponseBodyData) *DescribeDomainRealTimeQpsDataResponseBody
	GetData() *DescribeDomainRealTimeQpsDataResponseBodyData
	SetRequestId(v string) *DescribeDomainRealTimeQpsDataResponseBody
	GetRequestId() *string
}

type DescribeDomainRealTimeQpsDataResponseBody struct {
	// The data entries returned.
	Data *DescribeDomainRealTimeQpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 32DC9806-E9F9-4490-BBDC-B3A9E32FCC1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainRealTimeQpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataResponseBody) GetData() *DescribeDomainRealTimeQpsDataResponseBodyData {
	return s.Data
}

func (s *DescribeDomainRealTimeQpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealTimeQpsDataResponseBody) SetData(v *DescribeDomainRealTimeQpsDataResponseBodyData) *DescribeDomainRealTimeQpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainRealTimeQpsDataResponseBodyData struct {
	QpsModel []*DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel `json:"QpsModel,omitempty" xml:"QpsModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeQpsDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyData) GetQpsModel() []*DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel {
	return s.QpsModel
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyData) SetQpsModel(v []*DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) *DescribeDomainRealTimeQpsDataResponseBodyData {
	s.QpsModel = v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyData) Validate() error {
	if s.QpsModel != nil {
		for _, item := range s.QpsModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel struct {
	// The number of queries per second.
	//
	// example:
	//
	// 1851.25
	Qps *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The timestamp of the data returned. The time follows the yyyy-MM-ddTHH:mm:ssZ format in the ISO 8601 standard and is in UTC.
	//
	// example:
	//
	// 2019-12-02T11:25:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) GetQps() *float32 {
	return s.Qps
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) SetQps(v float32) *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.Qps = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) SetTimeStamp(v string) *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) Validate() error {
	return dara.Validate(s)
}
