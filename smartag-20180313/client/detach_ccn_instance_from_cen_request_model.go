// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCcnInstanceFromCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *DetachCcnInstanceFromCenRequest
	GetCcnId() *string
	SetCenId(v string) *DetachCcnInstanceFromCenRequest
	GetCenId() *string
	SetRegionBindingList(v []*DetachCcnInstanceFromCenRequestRegionBindingList) *DetachCcnInstanceFromCenRequest
	GetRegionBindingList() []*DetachCcnInstanceFromCenRequestRegionBindingList
	SetRegionId(v string) *DetachCcnInstanceFromCenRequest
	GetRegionId() *string
}

type DetachCcnInstanceFromCenRequest struct {
	// The ID of the attached Cloud Connect Network (CCN).
	//
	// example:
	//
	// ccn-isdjvvkexkrpk*****
	CcnId *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	// The ID of the CEN instance from which you want to revoke authorization.
	//
	// example:
	//
	// cen-9j8gkkj7z9vie9a8z9
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The list of transit routers to be disassociated from the CCN.
	RegionBindingList []*DetachCcnInstanceFromCenRequestRegionBindingList `json:"RegionBindingList,omitempty" xml:"RegionBindingList,omitempty" type:"Repeated"`
	// The region ID of the Smart Access Gateway instance. You can call the DescribeRegions operation to query the regions supported by Smart Access Gateway and the corresponding region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachCcnInstanceFromCenRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachCcnInstanceFromCenRequest) GoString() string {
	return s.String()
}

func (s *DetachCcnInstanceFromCenRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *DetachCcnInstanceFromCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *DetachCcnInstanceFromCenRequest) GetRegionBindingList() []*DetachCcnInstanceFromCenRequestRegionBindingList {
	return s.RegionBindingList
}

func (s *DetachCcnInstanceFromCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachCcnInstanceFromCenRequest) SetCcnId(v string) *DetachCcnInstanceFromCenRequest {
	s.CcnId = &v
	return s
}

func (s *DetachCcnInstanceFromCenRequest) SetCenId(v string) *DetachCcnInstanceFromCenRequest {
	s.CenId = &v
	return s
}

func (s *DetachCcnInstanceFromCenRequest) SetRegionBindingList(v []*DetachCcnInstanceFromCenRequestRegionBindingList) *DetachCcnInstanceFromCenRequest {
	s.RegionBindingList = v
	return s
}

func (s *DetachCcnInstanceFromCenRequest) SetRegionId(v string) *DetachCcnInstanceFromCenRequest {
	s.RegionId = &v
	return s
}

func (s *DetachCcnInstanceFromCenRequest) Validate() error {
	if s.RegionBindingList != nil {
		for _, item := range s.RegionBindingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachCcnInstanceFromCenRequestRegionBindingList struct {
	// The region ID of the transit router to be disassociated.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the forward routing transit router.
	//
	// example:
	//
	// tr-uf6mja2ssto2twh8u03xw
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DetachCcnInstanceFromCenRequestRegionBindingList) String() string {
	return dara.Prettify(s)
}

func (s DetachCcnInstanceFromCenRequestRegionBindingList) GoString() string {
	return s.String()
}

func (s *DetachCcnInstanceFromCenRequestRegionBindingList) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachCcnInstanceFromCenRequestRegionBindingList) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DetachCcnInstanceFromCenRequestRegionBindingList) SetRegionId(v string) *DetachCcnInstanceFromCenRequestRegionBindingList {
	s.RegionId = &v
	return s
}

func (s *DetachCcnInstanceFromCenRequestRegionBindingList) SetTransitRouterId(v string) *DetachCcnInstanceFromCenRequestRegionBindingList {
	s.TransitRouterId = &v
	return s
}

func (s *DetachCcnInstanceFromCenRequestRegionBindingList) Validate() error {
	return dara.Validate(s)
}
