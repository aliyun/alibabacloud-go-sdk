// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDomainRealTimeBpsDataResponseBodyData) *DescribeDomainRealTimeBpsDataResponseBody
	GetData() *DescribeDomainRealTimeBpsDataResponseBodyData
	SetRequestId(v string) *DescribeDomainRealTimeBpsDataResponseBody
	GetRequestId() *string
}

type DescribeDomainRealTimeBpsDataResponseBody struct {
	// The data returned.
	Data *DescribeDomainRealTimeBpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B49E6DDA-F413-422B-B58E-2FA23F286726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainRealTimeBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataResponseBody) GetData() *DescribeDomainRealTimeBpsDataResponseBodyData {
	return s.Data
}

func (s *DescribeDomainRealTimeBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealTimeBpsDataResponseBody) SetData(v *DescribeDomainRealTimeBpsDataResponseBodyData) *DescribeDomainRealTimeBpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainRealTimeBpsDataResponseBodyData struct {
	BpsModel []*DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel `json:"BpsModel,omitempty" xml:"BpsModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeBpsDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyData) GetBpsModel() []*DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel {
	return s.BpsModel
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyData) SetBpsModel(v []*DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) *DescribeDomainRealTimeBpsDataResponseBodyData {
	s.BpsModel = v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyData) Validate() error {
	if s.BpsModel != nil {
		for _, item := range s.BpsModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel struct {
	// The bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 16710625.733333332
	Bps *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The timestamp of the data returned. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-30T05:41:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) SetBps(v float32) *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.Bps = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) SetTimeStamp(v string) *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) Validate() error {
	return dara.Validate(s)
}
