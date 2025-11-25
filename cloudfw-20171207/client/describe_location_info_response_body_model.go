// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocationInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCnCount(v int32) *DescribeLocationInfoResponseBody
	GetCnCount() *int32
	SetCnList(v []*DescribeLocationInfoResponseBodyCnList) *DescribeLocationInfoResponseBody
	GetCnList() []*DescribeLocationInfoResponseBodyCnList
	SetCnProvinceList(v []*DescribeLocationInfoResponseBodyCnProvinceList) *DescribeLocationInfoResponseBody
	GetCnProvinceList() []*DescribeLocationInfoResponseBodyCnProvinceList
	SetOverseasCount(v int32) *DescribeLocationInfoResponseBody
	GetOverseasCount() *int32
	SetOverseasList(v []*DescribeLocationInfoResponseBodyOverseasList) *DescribeLocationInfoResponseBody
	GetOverseasList() []*DescribeLocationInfoResponseBodyOverseasList
	SetRequestId(v string) *DescribeLocationInfoResponseBody
	GetRequestId() *string
}

type DescribeLocationInfoResponseBody struct {
	// example:
	//
	// 340
	CnCount        *int32                                            `json:"CnCount,omitempty" xml:"CnCount,omitempty"`
	CnList         []*DescribeLocationInfoResponseBodyCnList         `json:"CnList,omitempty" xml:"CnList,omitempty" type:"Repeated"`
	CnProvinceList []*DescribeLocationInfoResponseBodyCnProvinceList `json:"CnProvinceList,omitempty" xml:"CnProvinceList,omitempty" type:"Repeated"`
	// example:
	//
	// 238
	OverseasCount *int32                                          `json:"OverseasCount,omitempty" xml:"OverseasCount,omitempty"`
	OverseasList  []*DescribeLocationInfoResponseBodyOverseasList `json:"OverseasList,omitempty" xml:"OverseasList,omitempty" type:"Repeated"`
	// example:
	//
	// CD3BE433-FAB0-55D8-918A-69B306****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLocationInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLocationInfoResponseBody) GetCnCount() *int32 {
	return s.CnCount
}

func (s *DescribeLocationInfoResponseBody) GetCnList() []*DescribeLocationInfoResponseBodyCnList {
	return s.CnList
}

func (s *DescribeLocationInfoResponseBody) GetCnProvinceList() []*DescribeLocationInfoResponseBodyCnProvinceList {
	return s.CnProvinceList
}

func (s *DescribeLocationInfoResponseBody) GetOverseasCount() *int32 {
	return s.OverseasCount
}

func (s *DescribeLocationInfoResponseBody) GetOverseasList() []*DescribeLocationInfoResponseBodyOverseasList {
	return s.OverseasList
}

func (s *DescribeLocationInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLocationInfoResponseBody) SetCnCount(v int32) *DescribeLocationInfoResponseBody {
	s.CnCount = &v
	return s
}

func (s *DescribeLocationInfoResponseBody) SetCnList(v []*DescribeLocationInfoResponseBodyCnList) *DescribeLocationInfoResponseBody {
	s.CnList = v
	return s
}

func (s *DescribeLocationInfoResponseBody) SetCnProvinceList(v []*DescribeLocationInfoResponseBodyCnProvinceList) *DescribeLocationInfoResponseBody {
	s.CnProvinceList = v
	return s
}

func (s *DescribeLocationInfoResponseBody) SetOverseasCount(v int32) *DescribeLocationInfoResponseBody {
	s.OverseasCount = &v
	return s
}

func (s *DescribeLocationInfoResponseBody) SetOverseasList(v []*DescribeLocationInfoResponseBodyOverseasList) *DescribeLocationInfoResponseBody {
	s.OverseasList = v
	return s
}

func (s *DescribeLocationInfoResponseBody) SetRequestId(v string) *DescribeLocationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLocationInfoResponseBody) Validate() error {
	if s.CnList != nil {
		for _, item := range s.CnList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CnProvinceList != nil {
		for _, item := range s.CnProvinceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OverseasList != nil {
		for _, item := range s.OverseasList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLocationInfoResponseBodyCnList struct {
	// example:
	//
	// 3301
	LocId   *string `json:"LocId,omitempty" xml:"LocId,omitempty"`
	LocName *string `json:"LocName,omitempty" xml:"LocName,omitempty"`
}

func (s DescribeLocationInfoResponseBodyCnList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocationInfoResponseBodyCnList) GoString() string {
	return s.String()
}

func (s *DescribeLocationInfoResponseBodyCnList) GetLocId() *string {
	return s.LocId
}

func (s *DescribeLocationInfoResponseBodyCnList) GetLocName() *string {
	return s.LocName
}

func (s *DescribeLocationInfoResponseBodyCnList) SetLocId(v string) *DescribeLocationInfoResponseBodyCnList {
	s.LocId = &v
	return s
}

func (s *DescribeLocationInfoResponseBodyCnList) SetLocName(v string) *DescribeLocationInfoResponseBodyCnList {
	s.LocName = &v
	return s
}

func (s *DescribeLocationInfoResponseBodyCnList) Validate() error {
	return dara.Validate(s)
}

type DescribeLocationInfoResponseBodyCnProvinceList struct {
	Cities       []*DescribeLocationInfoResponseBodyCnProvinceListCities `json:"Cities,omitempty" xml:"Cities,omitempty" type:"Repeated"`
	ProvinceName *string                                                 `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
}

func (s DescribeLocationInfoResponseBodyCnProvinceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocationInfoResponseBodyCnProvinceList) GoString() string {
	return s.String()
}

func (s *DescribeLocationInfoResponseBodyCnProvinceList) GetCities() []*DescribeLocationInfoResponseBodyCnProvinceListCities {
	return s.Cities
}

func (s *DescribeLocationInfoResponseBodyCnProvinceList) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *DescribeLocationInfoResponseBodyCnProvinceList) SetCities(v []*DescribeLocationInfoResponseBodyCnProvinceListCities) *DescribeLocationInfoResponseBodyCnProvinceList {
	s.Cities = v
	return s
}

func (s *DescribeLocationInfoResponseBodyCnProvinceList) SetProvinceName(v string) *DescribeLocationInfoResponseBodyCnProvinceList {
	s.ProvinceName = &v
	return s
}

func (s *DescribeLocationInfoResponseBodyCnProvinceList) Validate() error {
	if s.Cities != nil {
		for _, item := range s.Cities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLocationInfoResponseBodyCnProvinceListCities struct {
	// example:
	//
	// 123
	LocId   *string `json:"LocId,omitempty" xml:"LocId,omitempty"`
	LocName *string `json:"LocName,omitempty" xml:"LocName,omitempty"`
}

func (s DescribeLocationInfoResponseBodyCnProvinceListCities) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocationInfoResponseBodyCnProvinceListCities) GoString() string {
	return s.String()
}

func (s *DescribeLocationInfoResponseBodyCnProvinceListCities) GetLocId() *string {
	return s.LocId
}

func (s *DescribeLocationInfoResponseBodyCnProvinceListCities) GetLocName() *string {
	return s.LocName
}

func (s *DescribeLocationInfoResponseBodyCnProvinceListCities) SetLocId(v string) *DescribeLocationInfoResponseBodyCnProvinceListCities {
	s.LocId = &v
	return s
}

func (s *DescribeLocationInfoResponseBodyCnProvinceListCities) SetLocName(v string) *DescribeLocationInfoResponseBodyCnProvinceListCities {
	s.LocName = &v
	return s
}

func (s *DescribeLocationInfoResponseBodyCnProvinceListCities) Validate() error {
	return dara.Validate(s)
}

type DescribeLocationInfoResponseBodyOverseasList struct {
	// example:
	//
	// US
	LocId   *string `json:"LocId,omitempty" xml:"LocId,omitempty"`
	LocName *string `json:"LocName,omitempty" xml:"LocName,omitempty"`
}

func (s DescribeLocationInfoResponseBodyOverseasList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocationInfoResponseBodyOverseasList) GoString() string {
	return s.String()
}

func (s *DescribeLocationInfoResponseBodyOverseasList) GetLocId() *string {
	return s.LocId
}

func (s *DescribeLocationInfoResponseBodyOverseasList) GetLocName() *string {
	return s.LocName
}

func (s *DescribeLocationInfoResponseBodyOverseasList) SetLocId(v string) *DescribeLocationInfoResponseBodyOverseasList {
	s.LocId = &v
	return s
}

func (s *DescribeLocationInfoResponseBodyOverseasList) SetLocName(v string) *DescribeLocationInfoResponseBodyOverseasList {
	s.LocName = &v
	return s
}

func (s *DescribeLocationInfoResponseBodyOverseasList) Validate() error {
	return dara.Validate(s)
}
