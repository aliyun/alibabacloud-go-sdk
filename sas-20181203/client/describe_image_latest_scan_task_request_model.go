// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageLatestScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigest(v string) *DescribeImageLatestScanTaskRequest
	GetDigest() *string
}

type DescribeImageLatestScanTaskRequest struct {
	// The digest value of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8f0fbdb41d3d1ade4ffdf21558443f4c03342010563bb8c43ccc09594d50****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
}

func (s DescribeImageLatestScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageLatestScanTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageLatestScanTaskRequest) GetDigest() *string {
	return s.Digest
}

func (s *DescribeImageLatestScanTaskRequest) SetDigest(v string) *DescribeImageLatestScanTaskRequest {
	s.Digest = &v
	return s
}

func (s *DescribeImageLatestScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
