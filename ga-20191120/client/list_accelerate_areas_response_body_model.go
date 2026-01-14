// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccelerateAreasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAreas(v []*ListAccelerateAreasResponseBodyAreas) *ListAccelerateAreasResponseBody
	GetAreas() []*ListAccelerateAreasResponseBodyAreas
	SetRequestId(v string) *ListAccelerateAreasResponseBody
	GetRequestId() *string
}

type ListAccelerateAreasResponseBody struct {
	// The information about the areas.
	Areas []*ListAccelerateAreasResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccelerateAreasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerateAreasResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponseBody) GetAreas() []*ListAccelerateAreasResponseBodyAreas {
	return s.Areas
}

func (s *ListAccelerateAreasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccelerateAreasResponseBody) SetAreas(v []*ListAccelerateAreasResponseBodyAreas) *ListAccelerateAreasResponseBody {
	s.Areas = v
	return s
}

func (s *ListAccelerateAreasResponseBody) SetRequestId(v string) *ListAccelerateAreasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccelerateAreasResponseBody) Validate() error {
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

type ListAccelerateAreasResponseBodyAreas struct {
	// The area ID.
	//
	// example:
	//
	// cn-huabei
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The name of the area.
	//
	// example:
	//
	// North China
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The regions in the acceleration area.
	RegionList []*ListAccelerateAreasResponseBodyAreasRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s ListAccelerateAreasResponseBodyAreas) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerateAreasResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponseBodyAreas) GetAreaId() *string {
	return s.AreaId
}

func (s *ListAccelerateAreasResponseBodyAreas) GetLocalName() *string {
	return s.LocalName
}

func (s *ListAccelerateAreasResponseBodyAreas) GetRegionList() []*ListAccelerateAreasResponseBodyAreasRegionList {
	return s.RegionList
}

func (s *ListAccelerateAreasResponseBodyAreas) SetAreaId(v string) *ListAccelerateAreasResponseBodyAreas {
	s.AreaId = &v
	return s
}

func (s *ListAccelerateAreasResponseBodyAreas) SetLocalName(v string) *ListAccelerateAreasResponseBodyAreas {
	s.LocalName = &v
	return s
}

func (s *ListAccelerateAreasResponseBodyAreas) SetRegionList(v []*ListAccelerateAreasResponseBodyAreasRegionList) *ListAccelerateAreasResponseBodyAreas {
	s.RegionList = v
	return s
}

func (s *ListAccelerateAreasResponseBodyAreas) Validate() error {
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

type ListAccelerateAreasResponseBodyAreasRegionList struct {
	// The name of the region.
	//
	// example:
	//
	// China (Qingdao)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAccelerateAreasResponseBodyAreasRegionList) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerateAreasResponseBodyAreasRegionList) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponseBodyAreasRegionList) GetLocalName() *string {
	return s.LocalName
}

func (s *ListAccelerateAreasResponseBodyAreasRegionList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccelerateAreasResponseBodyAreasRegionList) SetLocalName(v string) *ListAccelerateAreasResponseBodyAreasRegionList {
	s.LocalName = &v
	return s
}

func (s *ListAccelerateAreasResponseBodyAreasRegionList) SetRegionId(v string) *ListAccelerateAreasResponseBodyAreasRegionList {
	s.RegionId = &v
	return s
}

func (s *ListAccelerateAreasResponseBodyAreasRegionList) Validate() error {
	return dara.Validate(s)
}
