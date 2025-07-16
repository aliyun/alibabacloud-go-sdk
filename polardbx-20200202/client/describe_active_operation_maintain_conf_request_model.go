// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationMaintainConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeActiveOperationMaintainConfRequest
	GetRegionId() *string
}

type DescribeActiveOperationMaintainConfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeActiveOperationMaintainConfRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationMaintainConfRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintainConfRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeActiveOperationMaintainConfRequest) SetRegionId(v string) *DescribeActiveOperationMaintainConfRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfRequest) Validate() error {
	return dara.Validate(s)
}
