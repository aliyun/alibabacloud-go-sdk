// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubnetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetSubnetRequest
	GetRegionId() *string
	SetSubnetId(v string) *GetSubnetRequest
	GetSubnetId() *string
	SetVpdId(v string) *GetSubnetRequest
	GetVpdId() *string
}

type GetSubnetRequest struct {
	// The region ID of the data center.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Lingjun subnet instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-2avf0itf
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The ID of the CIDR block to which Lingjun belongs.
	//
	// example:
	//
	// vpd-cxcmdk1m
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s GetSubnetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubnetRequest) GoString() string {
	return s.String()
}

func (s *GetSubnetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSubnetRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *GetSubnetRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *GetSubnetRequest) SetRegionId(v string) *GetSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *GetSubnetRequest) SetSubnetId(v string) *GetSubnetRequest {
	s.SubnetId = &v
	return s
}

func (s *GetSubnetRequest) SetVpdId(v string) *GetSubnetRequest {
	s.VpdId = &v
	return s
}

func (s *GetSubnetRequest) Validate() error {
	return dara.Validate(s)
}
