// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeEnsRegionsRequest
	GetEnsRegionId() *string
}

type DescribeEnsRegionsRequest struct {
	// The ID of the node.
	//
	// By default, all available node IDs are returned.
	//
	// example:
	//
	// cn-dalian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeEnsRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionsRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsRegionsRequest) SetEnsRegionId(v string) *DescribeEnsRegionsRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionsRequest) Validate() error {
	return dara.Validate(s)
}
