// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceCode(v string) *DescribeOssStatusRequest
	GetServiceCode() *string
}

type DescribeOssStatusRequest struct {
	// Service code:
	//
	// - antcloudauth: Financial-grade real-person authentication
	//
	// - cloudauthst (discontinued): Enhanced real-person authentication
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s DescribeOssStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssStatusRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeOssStatusRequest) SetServiceCode(v string) *DescribeOssStatusRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeOssStatusRequest) Validate() error {
	return dara.Validate(s)
}
