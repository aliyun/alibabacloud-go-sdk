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
	// The list of areas.
	Areas []*ListAvailableAccelerateAreasResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
	// The request ID.
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
	// The ID of the area.
	//
	// example:
	//
	// cn-huabei
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The name of the area.
	//
	// example:
	//
	// China North
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The list of regions.
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
	// - **true**: The region is in the Chinese mainland.
	//
	// - **false**: The region is not in the Chinese mainland.
	//
	// example:
	//
	// true
	ChinaMainland *bool `json:"ChinaMainland,omitempty" xml:"ChinaMainland,omitempty"`
	// The line type of the public IP address in the acceleration region.
	//
	// - **BGP*	- (default): BGP (Multi-ISP) line.
	//
	// - **BGP_PRO**: BGP (Multi-ISP) Pro line.
	IspTypeList []*string `json:"IspTypeList,omitempty" xml:"IspTypeList,omitempty" type:"Repeated"`
	// The name of the region.
	//
	// example:
	//
	// China (Qingdao)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// Indicates whether multi-zone deployment is supported. Valid values:
	//
	// - **true**: Multi-zone deployment is supported.
	//
	// - **false**: Multi-zone deployment is not supported.
	//
	// example:
	//
	// true
	MultiAz *bool `json:"MultiAz,omitempty" xml:"MultiAz,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether IPv6 is supported. Valid values:
	//
	// - **true**: IPv6 is supported.
	//
	// - **false**: IPv6 is not supported.
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
