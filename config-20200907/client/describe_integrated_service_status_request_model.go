// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntegratedServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceCode(v string) *DescribeIntegratedServiceStatusRequest
	GetServiceCode() *string
}

type DescribeIntegratedServiceStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cadt
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s DescribeIntegratedServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntegratedServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeIntegratedServiceStatusRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeIntegratedServiceStatusRequest) SetServiceCode(v string) *DescribeIntegratedServiceStatusRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeIntegratedServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
