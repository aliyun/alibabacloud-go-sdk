// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterModelEnhancementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterModelEnhancementRequest
	GetDBClusterId() *string
}

type DescribeDBClusterModelEnhancementRequest struct {
	// The ID of the PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterModelEnhancementRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterModelEnhancementRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterModelEnhancementRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterModelEnhancementRequest) SetDBClusterId(v string) *DescribeDBClusterModelEnhancementRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterModelEnhancementRequest) Validate() error {
	return dara.Validate(s)
}
