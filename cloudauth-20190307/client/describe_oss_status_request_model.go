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
	// The service code. Valid values:
	//
	// - antcloudauth: Chinese financial-grade ID Verification.
	//
	// - cloudauthst (discontinued): ID Verification Enhanced Edition.
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
