// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasProvinceStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVerifyPersonasProvinceStatisticsResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject) *DescribeVerifyPersonasProvinceStatisticsResponseBody
	GetResultObject() *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject
}

type DescribeVerifyPersonasProvinceStatisticsResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// D9821F95-CC18-5439-BB1C-21A0FF0C2003
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Query result.
	ResultObject *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeVerifyPersonasProvinceStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasProvinceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBody) GetResultObject() *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBody) SetRequestId(v string) *DescribeVerifyPersonasProvinceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBody) SetResultObject(v *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject) *DescribeVerifyPersonasProvinceStatisticsResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject struct {
	// Total number of devices.
	//
	// example:
	//
	// 30
	AllUserCnt *int64 `json:"AllUserCnt,omitempty" xml:"AllUserCnt,omitempty"`
	// Data items.
	Items []*DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject) GetAllUserCnt() *int64 {
	return s.AllUserCnt
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject) GetItems() []*DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems {
	return s.Items
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject) SetAllUserCnt(v int64) *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject {
	s.AllUserCnt = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject) SetItems(v []*DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject {
	s.Items = v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObject) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems struct {
	// Total number of devices in the province.
	//
	// example:
	//
	// 5
	ProvinceCnt *int64 `json:"ProvinceCnt,omitempty" xml:"ProvinceCnt,omitempty"`
	// Province name.
	//
	// example:
	//
	// 浙江
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// Percentage of the total for this province.
	//
	// example:
	//
	// 35.71
	ProvinceRate *string `json:"ProvinceRate,omitempty" xml:"ProvinceRate,omitempty"`
}

func (s DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) GetProvinceCnt() *int64 {
	return s.ProvinceCnt
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) GetProvinceRate() *string {
	return s.ProvinceRate
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) SetProvinceCnt(v int64) *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems {
	s.ProvinceCnt = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) SetProvinceName(v string) *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems {
	s.ProvinceName = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) SetProvinceRate(v string) *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems {
	s.ProvinceRate = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsResponseBodyResultObjectItems) Validate() error {
	return dara.Validate(s)
}
