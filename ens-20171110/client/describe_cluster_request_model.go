// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterRequest
	GetClusterId() *string
}

type DescribeClusterRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// c8f0377146d104687ac562eef9403****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterRequest) Validate() error {
	return dara.Validate(s)
}
