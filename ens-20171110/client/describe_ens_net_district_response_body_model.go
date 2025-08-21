// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsNetDistrictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnsNetDistrictResponseBody
	GetCode() *int32
	SetEnsNetDistricts(v *DescribeEnsNetDistrictResponseBodyEnsNetDistricts) *DescribeEnsNetDistrictResponseBody
	GetEnsNetDistricts() *DescribeEnsNetDistrictResponseBodyEnsNetDistricts
	SetRequestId(v string) *DescribeEnsNetDistrictResponseBody
	GetRequestId() *string
}

type DescribeEnsNetDistrictResponseBody struct {
	// The returned service code. A value of 0 indicates that the operation was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the regions.
	EnsNetDistricts *DescribeEnsNetDistrictResponseBodyEnsNetDistricts `json:"EnsNetDistricts,omitempty" xml:"EnsNetDistricts,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsNetDistrictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetDistrictResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnsNetDistrictResponseBody) GetEnsNetDistricts() *DescribeEnsNetDistrictResponseBodyEnsNetDistricts {
	return s.EnsNetDistricts
}

func (s *DescribeEnsNetDistrictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsNetDistrictResponseBody) SetCode(v int32) *DescribeEnsNetDistrictResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBody) SetEnsNetDistricts(v *DescribeEnsNetDistrictResponseBodyEnsNetDistricts) *DescribeEnsNetDistrictResponseBody {
	s.EnsNetDistricts = v
	return s
}

func (s *DescribeEnsNetDistrictResponseBody) SetRequestId(v string) *DescribeEnsNetDistrictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsNetDistrictResponseBodyEnsNetDistricts struct {
	EnsNetDistrict []*DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict `json:"EnsNetDistrict,omitempty" xml:"EnsNetDistrict,omitempty" type:"Repeated"`
}

func (s DescribeEnsNetDistrictResponseBodyEnsNetDistricts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetDistrictResponseBodyEnsNetDistricts) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistricts) GetEnsNetDistrict() []*DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	return s.EnsNetDistrict
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistricts) SetEnsNetDistrict(v []*DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) *DescribeEnsNetDistrictResponseBodyEnsNetDistricts {
	s.EnsNetDistrict = v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistricts) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict struct {
	// The number of nodes in the region.
	//
	// example:
	//
	// 2
	EnsRegionIdCount *string `json:"EnsRegionIdCount,omitempty" xml:"EnsRegionIdCount,omitempty"`
	// The code of the region.
	//
	// example:
	//
	// 100106
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// southWest
	NetDistrictEnName *string `json:"NetDistrictEnName,omitempty" xml:"NetDistrictEnName,omitempty"`
	// The parent code of the region.
	//
	// example:
	//
	// 100000
	NetDistrictFatherCode *string `json:"NetDistrictFatherCode,omitempty" xml:"NetDistrictFatherCode,omitempty"`
	// The level of the region.
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
	// The name of the region.
	//
	// example:
	//
	// southwest China
	NetDistrictName *string `json:"NetDistrictName,omitempty" xml:"NetDistrictName,omitempty"`
}

func (s DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetEnsRegionIdCount() *string {
	return s.EnsRegionIdCount
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictCode() *string {
	return s.NetDistrictCode
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictEnName() *string {
	return s.NetDistrictEnName
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictFatherCode() *string {
	return s.NetDistrictFatherCode
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictLevel() *string {
	return s.NetDistrictLevel
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) GetNetDistrictName() *string {
	return s.NetDistrictName
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetEnsRegionIdCount(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.EnsRegionIdCount = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictCode(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictCode = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictEnName(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictEnName = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictFatherCode(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictFatherCode = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictLevel(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictLevel = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) SetNetDistrictName(v string) *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict {
	s.NetDistrictName = &v
	return s
}

func (s *DescribeEnsNetDistrictResponseBodyEnsNetDistrictsEnsNetDistrict) Validate() error {
	return dara.Validate(s)
}
