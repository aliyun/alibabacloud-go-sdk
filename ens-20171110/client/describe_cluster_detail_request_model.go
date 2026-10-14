// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterDetailRequest
	GetClusterId() *string
}

type DescribeClusterDetailRequest struct {
	// The cluster ID.
	//
	// Only ECK cluster IDs with the eck- prefix are accepted. Legacy short hexadecimal cluster IDs are not supported. Call DescribeClustersV1 to obtain the ID. When listing clusters, pass in the Name parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// eck-11111111
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterDetailRequest) SetClusterId(v string) *DescribeClusterDetailRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterDetailRequest) Validate() error {
	return dara.Validate(s)
}
