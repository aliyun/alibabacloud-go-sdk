// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclType(v string) *DescribeAclAppsRequest
	GetAclType() *string
	SetLang(v string) *DescribeAclAppsRequest
	GetLang() *string
	SetPageNo(v int32) *DescribeAclAppsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAclAppsRequest
	GetPageSize() *int32
	SetPopular(v int32) *DescribeAclAppsRequest
	GetPopular() *int32
	SetProtocols(v []*string) *DescribeAclAppsRequest
	GetProtocols() []*string
}

type DescribeAclAppsRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	Popular   *int32    `json:"Popular,omitempty" xml:"Popular,omitempty"`
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
}

func (s DescribeAclAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclAppsRequest) GetAclType() *string {
	return s.AclType
}

func (s *DescribeAclAppsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAclAppsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAclAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAclAppsRequest) GetPopular() *int32 {
	return s.Popular
}

func (s *DescribeAclAppsRequest) GetProtocols() []*string {
	return s.Protocols
}

func (s *DescribeAclAppsRequest) SetAclType(v string) *DescribeAclAppsRequest {
	s.AclType = &v
	return s
}

func (s *DescribeAclAppsRequest) SetLang(v string) *DescribeAclAppsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAclAppsRequest) SetPageNo(v int32) *DescribeAclAppsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeAclAppsRequest) SetPageSize(v int32) *DescribeAclAppsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAclAppsRequest) SetPopular(v int32) *DescribeAclAppsRequest {
	s.Popular = &v
	return s
}

func (s *DescribeAclAppsRequest) SetProtocols(v []*string) *DescribeAclAppsRequest {
	s.Protocols = v
	return s
}

func (s *DescribeAclAppsRequest) Validate() error {
	return dara.Validate(s)
}
