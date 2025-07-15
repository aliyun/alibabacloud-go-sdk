// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *DescribeSystemParametersRequest
	GetSecurityToken() *string
}

type DescribeSystemParametersRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeSystemParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeSystemParametersRequest) SetSecurityToken(v string) *DescribeSystemParametersRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSystemParametersRequest) Validate() error {
	return dara.Validate(s)
}
