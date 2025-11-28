// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeStatInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeDn(v string) *DescribeEdgeStatInfoRequest
	GetEdgeDn() *string
	SetEdgePk(v string) *DescribeEdgeStatInfoRequest
	GetEdgePk() *string
	SetRegionId(v string) *DescribeEdgeStatInfoRequest
	GetRegionId() *string
}

type DescribeEdgeStatInfoRequest struct {
	EdgeDn *string `json:"EdgeDn,omitempty" xml:"EdgeDn,omitempty"`
	// This parameter is required.
	EdgePk   *string `json:"EdgePk,omitempty" xml:"EdgePk,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEdgeStatInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeStatInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeEdgeStatInfoRequest) GetEdgeDn() *string {
	return s.EdgeDn
}

func (s *DescribeEdgeStatInfoRequest) GetEdgePk() *string {
	return s.EdgePk
}

func (s *DescribeEdgeStatInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEdgeStatInfoRequest) SetEdgeDn(v string) *DescribeEdgeStatInfoRequest {
	s.EdgeDn = &v
	return s
}

func (s *DescribeEdgeStatInfoRequest) SetEdgePk(v string) *DescribeEdgeStatInfoRequest {
	s.EdgePk = &v
	return s
}

func (s *DescribeEdgeStatInfoRequest) SetRegionId(v string) *DescribeEdgeStatInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEdgeStatInfoRequest) Validate() error {
	return dara.Validate(s)
}
