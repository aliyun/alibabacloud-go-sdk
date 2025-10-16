// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclChecksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclType(v string) *DescribeAclChecksRequest
	GetAclType() *string
	SetLang(v string) *DescribeAclChecksRequest
	GetLang() *string
}

type DescribeAclChecksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// VPC
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAclChecksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclChecksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclChecksRequest) GetAclType() *string {
	return s.AclType
}

func (s *DescribeAclChecksRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAclChecksRequest) SetAclType(v string) *DescribeAclChecksRequest {
	s.AclType = &v
	return s
}

func (s *DescribeAclChecksRequest) SetLang(v string) *DescribeAclChecksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAclChecksRequest) Validate() error {
	return dara.Validate(s)
}
