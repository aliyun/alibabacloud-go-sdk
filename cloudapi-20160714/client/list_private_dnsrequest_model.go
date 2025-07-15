// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateDNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntranetDomain(v string) *ListPrivateDNSRequest
	GetIntranetDomain() *string
	SetPageNumber(v int32) *ListPrivateDNSRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPrivateDNSRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *ListPrivateDNSRequest
	GetSecurityToken() *string
	SetType(v string) *ListPrivateDNSRequest
	GetType() *string
}

type ListPrivateDNSRequest struct {
	// The internal domain name.
	//
	// example:
	//
	// api.demo.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The internal domain name resolution type. Valid values:
	//
	// 	- VPC: resolution for virtual private cloud (VPC) access authorizations. A resolution of this type can be bound only to traditional dedicated instances.
	//
	// 	- A: resolution that supports A records. A resolution of this type can be bound only to VPC integration dedicated instances.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrivateDNSRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateDNSRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateDNSRequest) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *ListPrivateDNSRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPrivateDNSRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPrivateDNSRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListPrivateDNSRequest) GetType() *string {
	return s.Type
}

func (s *ListPrivateDNSRequest) SetIntranetDomain(v string) *ListPrivateDNSRequest {
	s.IntranetDomain = &v
	return s
}

func (s *ListPrivateDNSRequest) SetPageNumber(v int32) *ListPrivateDNSRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPrivateDNSRequest) SetPageSize(v int32) *ListPrivateDNSRequest {
	s.PageSize = &v
	return s
}

func (s *ListPrivateDNSRequest) SetSecurityToken(v string) *ListPrivateDNSRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListPrivateDNSRequest) SetType(v string) *ListPrivateDNSRequest {
	s.Type = &v
	return s
}

func (s *ListPrivateDNSRequest) Validate() error {
	return dara.Validate(s)
}
