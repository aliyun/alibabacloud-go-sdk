// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataByTimeStampResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataList(v *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) *DescribeDomainBpsDataByTimeStampResponseBody
	GetBpsDataList() *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList
	SetDomainName(v string) *DescribeDomainBpsDataByTimeStampResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeDomainBpsDataByTimeStampResponseBody
	GetRequestId() *string
	SetTimeStamp(v string) *DescribeDomainBpsDataByTimeStampResponseBody
	GetTimeStamp() *string
}

type DescribeDomainBpsDataByTimeStampResponseBody struct {
	// A list of bandwidth values by ISP and region.
	BpsDataList *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList `json:"BpsDataList,omitempty" xml:"BpsDataList,omitempty" type:"Struct"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The point in time.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainBpsDataByTimeStampResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) GetBpsDataList() *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList {
	return s.BpsDataList
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetBpsDataList(v *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.BpsDataList = v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetDomainName(v string) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetRequestId(v string) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetTimeStamp(v string) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) Validate() error {
	if s.BpsDataList != nil {
		if err := s.BpsDataList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList struct {
	BpsDataModel []*DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel `json:"BpsDataModel,omitempty" xml:"BpsDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) GetBpsDataModel() []*DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	return s.BpsDataModel
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) SetBpsDataModel(v []*DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList {
	s.BpsDataModel = v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) Validate() error {
	if s.BpsDataModel != nil {
		for _, item := range s.BpsDataModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel struct {
	// The bandwidth value.
	//
	// example:
	//
	// 52119553
	Bps *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The name of the ISP.
	//
	// example:
	//
	// unicom
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// Liaoning
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) GetBps() *int64 {
	return s.Bps
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) GetIspName() *string {
	return s.IspName
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) GetLocationName() *string {
	return s.LocationName
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetBps(v int64) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.Bps = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetIspName(v string) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.IspName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetLocationName(v string) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.LocationName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetTimeStamp(v string) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) Validate() error {
	return dara.Validate(s)
}
