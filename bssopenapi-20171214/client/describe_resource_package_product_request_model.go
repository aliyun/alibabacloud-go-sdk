// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcePackageProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *DescribeResourcePackageProductRequest
	GetProductCode() *string
}

type DescribeResourcePackageProductRequest struct {
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ossbag
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeResourcePackageProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeResourcePackageProductRequest) SetProductCode(v string) *DescribeResourcePackageProductRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeResourcePackageProductRequest) Validate() error {
	return dara.Validate(s)
}
