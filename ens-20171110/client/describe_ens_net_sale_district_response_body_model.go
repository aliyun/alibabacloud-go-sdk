// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsNetSaleDistrictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnsNetSaleDistrictResponseBody
	GetCode() *int32
	SetEnsNetDistricts(v *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) *DescribeEnsNetSaleDistrictResponseBody
	GetEnsNetDistricts() *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts
	SetRequestId(v string) *DescribeEnsNetSaleDistrictResponseBody
	GetRequestId() *string
}

type DescribeEnsNetSaleDistrictResponseBody struct {
	// The returned service code. A value of 0 indicates that the operation was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the ISPs in the area.
	EnsNetDistricts *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts `json:"EnsNetDistricts,omitempty" xml:"EnsNetDistricts,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1707B55C-A12F-43EF-BC66-14FFDB9253C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsNetSaleDistrictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnsNetSaleDistrictResponseBody) GetEnsNetDistricts() *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts {
	return s.EnsNetDistricts
}

func (s *DescribeEnsNetSaleDistrictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsNetSaleDistrictResponseBody) SetCode(v int32) *DescribeEnsNetSaleDistrictResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBody) SetEnsNetDistricts(v *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) *DescribeEnsNetSaleDistrictResponseBody {
	s.EnsNetDistricts = v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBody) SetRequestId(v string) *DescribeEnsNetSaleDistrictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBody) Validate() error {
	if s.EnsNetDistricts != nil {
		if err := s.EnsNetDistricts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts struct {
	EnsNetDistrict []*DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict `json:"EnsNetDistrict,omitempty" xml:"EnsNetDistrict,omitempty" type:"Repeated"`
}

func (s DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) GetEnsNetDistrict() []*DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	return s.EnsNetDistrict
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) SetEnsNetDistrict(v []*DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts {
	s.EnsNetDistrict = v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistricts) Validate() error {
	if s.EnsNetDistrict != nil {
		for _, item := range s.EnsNetDistrict {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict struct {
	// The information about the ISP.
	//
	// example:
	//
	// {                         "count":2,                         "name":"telecommunication",                         "code":"telecom",                         "country":"cn"                     }
	EnsRegionIdCount *string `json:"EnsRegionIdCount,omitempty" xml:"EnsRegionIdCount,omitempty"`
	// The information about the instance.
	//
	// example:
	//
	// {                         "count":2,                         "code":"multiCarrier"                     }
	InstanceCount *string `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The region code.
	//
	// example:
	//
	// 100101
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// northEast
	NetDistrictEnName *string `json:"NetDistrictEnName,omitempty" xml:"NetDistrictEnName,omitempty"`
	// The parent code of the region.
	//
	// example:
	//
	// 100000
	NetDistrictFatherCode *string `json:"NetDistrictFatherCode,omitempty" xml:"NetDistrictFatherCode,omitempty"`
	// The region level. Valid values:
	//
	// 	- **Big**: area
	//
	// 	- **Middle**: province
	//
	// 	- **Small**: city
	//
	// example:
	//
	// Big
	NetDistrictLevel *string `json:"NetDistrictLevel,omitempty" xml:"NetDistrictLevel,omitempty"`
	// The Chinese name of the region.
	//
	// example:
	//
	// Northeast China
	NetDistrictName *string `json:"NetDistrictName,omitempty" xml:"NetDistrictName,omitempty"`
}

func (s DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetEnsRegionIdCount() *string {
	return s.EnsRegionIdCount
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetInstanceCount() *string {
	return s.InstanceCount
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictCode() *string {
	return s.NetDistrictCode
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictEnName() *string {
	return s.NetDistrictEnName
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictFatherCode() *string {
	return s.NetDistrictFatherCode
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictLevel() *string {
	return s.NetDistrictLevel
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictName() *string {
	return s.NetDistrictName
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetEnsRegionIdCount(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.EnsRegionIdCount = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetInstanceCount(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.InstanceCount = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictCode(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictEnName(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictEnName = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictFatherCode(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictFatherCode = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictLevel(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictLevel = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictName(v string) *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictName = &v
	return s
}

func (s *DescribeEnsNetSaleDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) Validate() error {
	return dara.Validate(s)
}
