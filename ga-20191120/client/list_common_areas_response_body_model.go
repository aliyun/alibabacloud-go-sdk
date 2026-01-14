// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonAreasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAreas(v []*ListCommonAreasResponseBodyAreas) *ListCommonAreasResponseBody
	GetAreas() []*ListCommonAreasResponseBodyAreas
	SetRequestId(v string) *ListCommonAreasResponseBody
	GetRequestId() *string
}

type ListCommonAreasResponseBody struct {
	// The information about the areas.
	Areas []*ListCommonAreasResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCommonAreasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCommonAreasResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonAreasResponseBody) GetAreas() []*ListCommonAreasResponseBodyAreas {
	return s.Areas
}

func (s *ListCommonAreasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCommonAreasResponseBody) SetAreas(v []*ListCommonAreasResponseBodyAreas) *ListCommonAreasResponseBody {
	s.Areas = v
	return s
}

func (s *ListCommonAreasResponseBody) SetRequestId(v string) *ListCommonAreasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonAreasResponseBody) Validate() error {
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

type ListCommonAreasResponseBodyAreas struct {
	// The area ID.
	//
	// example:
	//
	// cn-huabei
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The area name.
	//
	// example:
	//
	// North China
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The information about the regions.
	RegionList []*ListCommonAreasResponseBodyAreasRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s ListCommonAreasResponseBodyAreas) String() string {
	return dara.Prettify(s)
}

func (s ListCommonAreasResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *ListCommonAreasResponseBodyAreas) GetAreaId() *string {
	return s.AreaId
}

func (s *ListCommonAreasResponseBodyAreas) GetLocalName() *string {
	return s.LocalName
}

func (s *ListCommonAreasResponseBodyAreas) GetRegionList() []*ListCommonAreasResponseBodyAreasRegionList {
	return s.RegionList
}

func (s *ListCommonAreasResponseBodyAreas) SetAreaId(v string) *ListCommonAreasResponseBodyAreas {
	s.AreaId = &v
	return s
}

func (s *ListCommonAreasResponseBodyAreas) SetLocalName(v string) *ListCommonAreasResponseBodyAreas {
	s.LocalName = &v
	return s
}

func (s *ListCommonAreasResponseBodyAreas) SetRegionList(v []*ListCommonAreasResponseBodyAreasRegionList) *ListCommonAreasResponseBodyAreas {
	s.RegionList = v
	return s
}

func (s *ListCommonAreasResponseBodyAreas) Validate() error {
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

type ListCommonAreasResponseBodyAreasRegionList struct {
	// The region name.
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

func (s ListCommonAreasResponseBodyAreasRegionList) String() string {
	return dara.Prettify(s)
}

func (s ListCommonAreasResponseBodyAreasRegionList) GoString() string {
	return s.String()
}

func (s *ListCommonAreasResponseBodyAreasRegionList) GetLocalName() *string {
	return s.LocalName
}

func (s *ListCommonAreasResponseBodyAreasRegionList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCommonAreasResponseBodyAreasRegionList) SetLocalName(v string) *ListCommonAreasResponseBodyAreasRegionList {
	s.LocalName = &v
	return s
}

func (s *ListCommonAreasResponseBodyAreasRegionList) SetRegionId(v string) *ListCommonAreasResponseBodyAreasRegionList {
	s.RegionId = &v
	return s
}

func (s *ListCommonAreasResponseBodyAreasRegionList) Validate() error {
	return dara.Validate(s)
}
