// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCcnInstanceToCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *AttachCcnInstanceToCenRequest
	GetCcnId() *string
	SetCenId(v string) *AttachCcnInstanceToCenRequest
	GetCenId() *string
	SetRegionId(v string) *AttachCcnInstanceToCenRequest
	GetRegionId() *string
	SetSubnet(v string) *AttachCcnInstanceToCenRequest
	GetSubnet() *string
}

type AttachCcnInstanceToCenRequest struct {
	// The ID of the Cloud Connect Network (CCN) instance to attach.
	//
	// example:
	//
	// ccn-isdjvvkexkrpk*****
	CcnId *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	// The ID of the CEN instance to authorize.
	//
	// example:
	//
	// cen-joimmi1s2ob3rdxw79
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The region ID of the Cloud Connect Network (CCN) instance. You can invoke the DescribeRegions operation to query the regions supported by Smart Access Gateway and the corresponding region IDs.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Internet CIDR block used when the Cloud Connect Network (CCN) instance is attached to the CEN instance.
	//
	// example:
	//
	// 172.16.55.0/24
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s AttachCcnInstanceToCenRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachCcnInstanceToCenRequest) GoString() string {
	return s.String()
}

func (s *AttachCcnInstanceToCenRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *AttachCcnInstanceToCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *AttachCcnInstanceToCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachCcnInstanceToCenRequest) GetSubnet() *string {
	return s.Subnet
}

func (s *AttachCcnInstanceToCenRequest) SetCcnId(v string) *AttachCcnInstanceToCenRequest {
	s.CcnId = &v
	return s
}

func (s *AttachCcnInstanceToCenRequest) SetCenId(v string) *AttachCcnInstanceToCenRequest {
	s.CenId = &v
	return s
}

func (s *AttachCcnInstanceToCenRequest) SetRegionId(v string) *AttachCcnInstanceToCenRequest {
	s.RegionId = &v
	return s
}

func (s *AttachCcnInstanceToCenRequest) SetSubnet(v string) *AttachCcnInstanceToCenRequest {
	s.Subnet = &v
	return s
}

func (s *AttachCcnInstanceToCenRequest) Validate() error {
	return dara.Validate(s)
}
