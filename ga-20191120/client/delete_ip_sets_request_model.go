// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpSetIds(v []*string) *DeleteIpSetsRequest
	GetIpSetIds() []*string
	SetRegionId(v string) *DeleteIpSetsRequest
	GetRegionId() *string
}

type DeleteIpSetsRequest struct {
	// The IDs of the acceleration regions that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ips-bp11c9mpphtb1xkds****
	IpSetIds []*string `json:"IpSetIds,omitempty" xml:"IpSetIds,omitempty" type:"Repeated"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpSetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpSetsRequest) GetIpSetIds() []*string {
	return s.IpSetIds
}

func (s *DeleteIpSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpSetsRequest) SetIpSetIds(v []*string) *DeleteIpSetsRequest {
	s.IpSetIds = v
	return s
}

func (s *DeleteIpSetsRequest) SetRegionId(v string) *DeleteIpSetsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpSetsRequest) Validate() error {
	return dara.Validate(s)
}
