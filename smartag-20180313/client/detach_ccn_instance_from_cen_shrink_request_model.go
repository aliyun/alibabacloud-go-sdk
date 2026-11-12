// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCcnInstanceFromCenShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *DetachCcnInstanceFromCenShrinkRequest
	GetCcnId() *string
	SetCenId(v string) *DetachCcnInstanceFromCenShrinkRequest
	GetCenId() *string
	SetRegionBindingListShrink(v string) *DetachCcnInstanceFromCenShrinkRequest
	GetRegionBindingListShrink() *string
	SetRegionId(v string) *DetachCcnInstanceFromCenShrinkRequest
	GetRegionId() *string
}

type DetachCcnInstanceFromCenShrinkRequest struct {
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
	RegionBindingListShrink *string `json:"RegionBindingList,omitempty" xml:"RegionBindingList,omitempty"`
	// The region ID of the Smart Access Gateway instance. You can call the DescribeRegions operation to query the regions supported by Smart Access Gateway and the corresponding region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachCcnInstanceFromCenShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachCcnInstanceFromCenShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachCcnInstanceFromCenShrinkRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *DetachCcnInstanceFromCenShrinkRequest) GetCenId() *string {
	return s.CenId
}

func (s *DetachCcnInstanceFromCenShrinkRequest) GetRegionBindingListShrink() *string {
	return s.RegionBindingListShrink
}

func (s *DetachCcnInstanceFromCenShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachCcnInstanceFromCenShrinkRequest) SetCcnId(v string) *DetachCcnInstanceFromCenShrinkRequest {
	s.CcnId = &v
	return s
}

func (s *DetachCcnInstanceFromCenShrinkRequest) SetCenId(v string) *DetachCcnInstanceFromCenShrinkRequest {
	s.CenId = &v
	return s
}

func (s *DetachCcnInstanceFromCenShrinkRequest) SetRegionBindingListShrink(v string) *DetachCcnInstanceFromCenShrinkRequest {
	s.RegionBindingListShrink = &v
	return s
}

func (s *DetachCcnInstanceFromCenShrinkRequest) SetRegionId(v string) *DetachCcnInstanceFromCenShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DetachCcnInstanceFromCenShrinkRequest) Validate() error {
	return dara.Validate(s)
}
