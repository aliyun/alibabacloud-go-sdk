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
	SetRegionId(v string) *DetachCcnInstanceFromCenRequest
	GetRegionId() *string
}

type DetachCcnInstanceFromCenRequest struct {
	// The ID of the Cloud Connect Network (CCN) that is bound to the CEN instance.
	//
	// example:
	//
	// ccn-isdjvvkexkrpk*****
	CcnId *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance from which you want to revoke the authorization.
	//
	// example:
	//
	// cen-9j8gkkj7z9vie9a8z9
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
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

func (s *DetachCcnInstanceFromCenRequest) SetRegionId(v string) *DetachCcnInstanceFromCenRequest {
	s.RegionId = &v
	return s
}

func (s *DetachCcnInstanceFromCenRequest) Validate() error {
	return dara.Validate(s)
}
