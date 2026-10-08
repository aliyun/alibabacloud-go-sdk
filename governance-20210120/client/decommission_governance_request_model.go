// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecommissionGovernanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DecommissionGovernanceRequest
	GetRegionId() *string
}

type DecommissionGovernanceRequest struct {
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DecommissionGovernanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DecommissionGovernanceRequest) GoString() string {
	return s.String()
}

func (s *DecommissionGovernanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DecommissionGovernanceRequest) SetRegionId(v string) *DecommissionGovernanceRequest {
	s.RegionId = &v
	return s
}

func (s *DecommissionGovernanceRequest) Validate() error {
	return dara.Validate(s)
}
