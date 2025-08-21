// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionIspsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeRegionIspsRequest
	GetEnsRegionId() *string
}

type DescribeRegionIspsRequest struct {
	// The ID of the node. You can specify only one node ID in a call.
	//
	// example:
	//
	// cn-dalian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeRegionIspsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionIspsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionIspsRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeRegionIspsRequest) SetEnsRegionId(v string) *DescribeRegionIspsRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeRegionIspsRequest) Validate() error {
	return dara.Validate(s)
}
