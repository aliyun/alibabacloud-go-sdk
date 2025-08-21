// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRegionIdIpv6InfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeEnsRegionIdIpv6InfoRequest
	GetEnsRegionId() *string
}

type DescribeEnsRegionIdIpv6InfoRequest struct {
	// The ID of the node. You can specify only one node ID in a call.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu-xxxx-4
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeEnsRegionIdIpv6InfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsRegionIdIpv6InfoRequest) SetEnsRegionId(v string) *DescribeEnsRegionIdIpv6InfoRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoRequest) Validate() error {
	return dara.Validate(s)
}
