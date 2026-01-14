// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableAccelerateAreasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAreas(v []*ListAvailableAccelerateAreasResponseBodyAreas) *ListAvailableAccelerateAreasResponseBody
	GetAreas() []*ListAvailableAccelerateAreasResponseBodyAreas
	SetRequestId(v string) *ListAvailableAccelerateAreasResponseBody
	GetRequestId() *string
}

type ListAvailableAccelerateAreasResponseBody struct {
	// The information about acceleration areas.
	Areas []*ListAvailableAccelerateAreasResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A9B4E54C-9CCD-4002-91A9-D38C6C209192
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAvailableAccelerateAreasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponseBody) GetAreas() []*ListAvailableAccelerateAreasResponseBodyAreas {
	return s.Areas
}

func (s *ListAvailableAccelerateAreasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableAccelerateAreasResponseBody) SetAreas(v []*ListAvailableAccelerateAreasResponseBodyAreas) *ListAvailableAccelerateAreasResponseBody {
	s.Areas = v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBody) SetRequestId(v string) *ListAvailableAccelerateAreasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBody) Validate() error {
	if s.Areas != nil {
		for _, item := range s.Areas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvailableAccelerateAreasResponseBodyAreas struct {
	// The ID of the acceleration area.
	//
	// example:
	//
	// cn-huabei
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The acceleration area name.
	//
	// example:
	//
	// North China
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The information about acceleration regions.
	RegionList []*ListAvailableAccelerateAreasResponseBodyAreasRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s ListAvailableAccelerateAreasResponseBodyAreas) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) GetAreaId() *string {
	return s.AreaId
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) GetLocalName() *string {
	return s.LocalName
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) GetRegionList() []*ListAvailableAccelerateAreasResponseBodyAreasRegionList {
	return s.RegionList
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) SetAreaId(v string) *ListAvailableAccelerateAreasResponseBodyAreas {
	s.AreaId = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) SetLocalName(v string) *ListAvailableAccelerateAreasResponseBodyAreas {
	s.LocalName = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) SetRegionList(v []*ListAvailableAccelerateAreasResponseBodyAreasRegionList) *ListAvailableAccelerateAreasResponseBodyAreas {
	s.RegionList = v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) Validate() error {
	if s.RegionList != nil {
		for _, item := range s.RegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvailableAccelerateAreasResponseBodyAreasRegionList struct {
	// Indicates whether the region is in the Chinese mainland. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ChinaMainland *bool `json:"ChinaMainland,omitempty" xml:"ChinaMainland,omitempty"`
	// The line type of the elastic IP address (EIP) in the acceleration region. Valid values:
	//
	// 	- **BGP**: BGP (Multi-ISP) lines.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro lines.
	IspTypeList []*string `json:"IspTypeList,omitempty" xml:"IspTypeList,omitempty" type:"Repeated"`
	// The acceleration region name.
	//
	// example:
	//
	// China (Qingdao)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// Indicates whether multiple zones are supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	MultiAz *bool `json:"MultiAz,omitempty" xml:"MultiAz,omitempty"`
	// The ID of the acceleration region.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether IPv6 is supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SupportIpv6 *bool `json:"SupportIpv6,omitempty" xml:"SupportIpv6,omitempty"`
}

func (s ListAvailableAccelerateAreasResponseBodyAreasRegionList) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponseBodyAreasRegionList) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) GetChinaMainland() *bool {
	return s.ChinaMainland
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) GetIspTypeList() []*string {
	return s.IspTypeList
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) GetLocalName() *string {
	return s.LocalName
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) GetMultiAz() *bool {
	return s.MultiAz
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) GetSupportIpv6() *bool {
	return s.SupportIpv6
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) SetChinaMainland(v bool) *ListAvailableAccelerateAreasResponseBodyAreasRegionList {
	s.ChinaMainland = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) SetIspTypeList(v []*string) *ListAvailableAccelerateAreasResponseBodyAreasRegionList {
	s.IspTypeList = v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) SetLocalName(v string) *ListAvailableAccelerateAreasResponseBodyAreasRegionList {
	s.LocalName = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) SetMultiAz(v bool) *ListAvailableAccelerateAreasResponseBodyAreasRegionList {
	s.MultiAz = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) SetRegionId(v string) *ListAvailableAccelerateAreasResponseBodyAreasRegionList {
	s.RegionId = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) SetSupportIpv6(v bool) *ListAvailableAccelerateAreasResponseBodyAreasRegionList {
	s.SupportIpv6 = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) Validate() error {
	return dara.Validate(s)
}
