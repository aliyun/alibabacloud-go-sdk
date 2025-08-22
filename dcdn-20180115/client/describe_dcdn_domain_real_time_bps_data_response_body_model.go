// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDcdnDomainRealTimeBpsDataResponseBodyData) *DescribeDcdnDomainRealTimeBpsDataResponseBody
	GetData() *DescribeDcdnDomainRealTimeBpsDataResponseBodyData
	SetRequestId(v string) *DescribeDcdnDomainRealTimeBpsDataResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainRealTimeBpsDataResponseBody struct {
	// The returned data.
	Data *DescribeDcdnDomainRealTimeBpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B49E6DDA-F413-422B-B58E-2FA23F286726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBody) GetData() *DescribeDcdnDomainRealTimeBpsDataResponseBodyData {
	return s.Data
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBody) SetData(v *DescribeDcdnDomainRealTimeBpsDataResponseBodyData) *DescribeDcdnDomainRealTimeBpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeBpsDataResponseBodyData struct {
	BpsModel []*DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel `json:"BpsModel,omitempty" xml:"BpsModel,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyData) GetBpsModel() []*DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel {
	return s.BpsModel
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyData) SetBpsModel(v []*DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) *DescribeDcdnDomainRealTimeBpsDataResponseBodyData {
	s.BpsModel = v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel struct {
	// The bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 16710625.733333332
	Bps *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-02T11:05:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) SetBps(v float32) *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeBpsDataResponseBodyDataBpsModel) Validate() error {
	return dara.Validate(s)
}
