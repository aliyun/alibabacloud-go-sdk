// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerlessClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeServerlessClusterRequest
	GetClusterId() *string
	SetZoneId(v string) *DescribeServerlessClusterRequest
	GetZoneId() *string
}

type DescribeServerlessClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16f1441y6p2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sdh0b7f4k5f****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeServerlessClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerlessClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeServerlessClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeServerlessClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeServerlessClusterRequest) SetClusterId(v string) *DescribeServerlessClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeServerlessClusterRequest) SetZoneId(v string) *DescribeServerlessClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeServerlessClusterRequest) Validate() error {
	return dara.Validate(s)
}
