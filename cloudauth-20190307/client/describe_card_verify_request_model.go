// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCardVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertifyId(v string) *DescribeCardVerifyRequest
	GetCertifyId() *string
}

type DescribeCardVerifyRequest struct {
	// Authentication request ID.
	//
	// You must first call the initialization interface InitCardVerify to submit an authentication request in order to get the authentication request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0xxxx
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
}

func (s DescribeCardVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCardVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribeCardVerifyRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DescribeCardVerifyRequest) SetCertifyId(v string) *DescribeCardVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *DescribeCardVerifyRequest) Validate() error {
	return dara.Validate(s)
}
