// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossCloudRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrossCloudRegionList(v []*DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) *DescribeCrossCloudRegionResponseBody
	GetCrossCloudRegionList() []*DescribeCrossCloudRegionResponseBodyCrossCloudRegionList
	SetRequestId(v string) *DescribeCrossCloudRegionResponseBody
	GetRequestId() *string
}

type DescribeCrossCloudRegionResponseBody struct {
	CrossCloudRegionList []*DescribeCrossCloudRegionResponseBodyCrossCloudRegionList `json:"CrossCloudRegionList,omitempty" xml:"CrossCloudRegionList,omitempty" type:"Repeated"`
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCrossCloudRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudRegionResponseBody) GetCrossCloudRegionList() []*DescribeCrossCloudRegionResponseBodyCrossCloudRegionList {
	return s.CrossCloudRegionList
}

func (s *DescribeCrossCloudRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCrossCloudRegionResponseBody) SetCrossCloudRegionList(v []*DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) *DescribeCrossCloudRegionResponseBody {
	s.CrossCloudRegionList = v
	return s
}

func (s *DescribeCrossCloudRegionResponseBody) SetRequestId(v string) *DescribeCrossCloudRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrossCloudRegionResponseBody) Validate() error {
	if s.CrossCloudRegionList != nil {
		for _, item := range s.CrossCloudRegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCrossCloudRegionResponseBodyCrossCloudRegionList struct {
	// example:
	//
	// cn-east-3
	CrossCloudRegionId *string `json:"CrossCloudRegionId,omitempty" xml:"CrossCloudRegionId,omitempty"`
	// example:
	//
	// cn-east-3
	CrossCloudRegionName *string                                                                       `json:"CrossCloudRegionName,omitempty" xml:"CrossCloudRegionName,omitempty"`
	CrossCloudZoneList   []*DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList `json:"CrossCloudZoneList,omitempty" xml:"CrossCloudZoneList,omitempty" type:"Repeated"`
	// example:
	//
	// pj-87681rbcef6******
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) GetCrossCloudRegionId() *string {
	return s.CrossCloudRegionId
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) GetCrossCloudRegionName() *string {
	return s.CrossCloudRegionName
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) GetCrossCloudZoneList() []*DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList {
	return s.CrossCloudZoneList
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) SetCrossCloudRegionId(v string) *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList {
	s.CrossCloudRegionId = &v
	return s
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) SetCrossCloudRegionName(v string) *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList {
	s.CrossCloudRegionName = &v
	return s
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) SetCrossCloudZoneList(v []*DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList) *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList {
	s.CrossCloudZoneList = v
	return s
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) SetProjectId(v string) *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList {
	s.ProjectId = &v
	return s
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionList) Validate() error {
	if s.CrossCloudZoneList != nil {
		for _, item := range s.CrossCloudZoneList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList struct {
	// example:
	//
	// cn-east-3-1
	CrossCloudZoneId *string `json:"CrossCloudZoneId,omitempty" xml:"CrossCloudZoneId,omitempty"`
	// example:
	//
	// cn-east-3-1
	CrossCloudZoneName *string `json:"CrossCloudZoneName,omitempty" xml:"CrossCloudZoneName,omitempty"`
}

func (s DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList) GoString() string {
	return s.String()
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList) GetCrossCloudZoneId() *string {
	return s.CrossCloudZoneId
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList) GetCrossCloudZoneName() *string {
	return s.CrossCloudZoneName
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList) SetCrossCloudZoneId(v string) *DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList {
	s.CrossCloudZoneId = &v
	return s
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList) SetCrossCloudZoneName(v string) *DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList {
	s.CrossCloudZoneName = &v
	return s
}

func (s *DescribeCrossCloudRegionResponseBodyCrossCloudRegionListCrossCloudZoneList) Validate() error {
	return dara.Validate(s)
}
