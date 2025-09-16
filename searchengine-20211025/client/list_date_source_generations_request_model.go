// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDateSourceGenerationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ListDateSourceGenerationsRequest
	GetDomainName() *string
	SetValidStatus(v bool) *ListDateSourceGenerationsRequest
	GetValidStatus() *bool
}

type ListDateSourceGenerationsRequest struct {
	// The data center where the data source is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// bj_vpc_domain_1
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// Specifies the index versions to be returned. Valid values:
	//
	// 1.  true (default): returns the index versions that are complete and not expired.
	//
	// 2.  false: returns all index versions.
	//
	// example:
	//
	// true
	ValidStatus *bool `json:"validStatus,omitempty" xml:"validStatus,omitempty"`
}

func (s ListDateSourceGenerationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDateSourceGenerationsRequest) GoString() string {
	return s.String()
}

func (s *ListDateSourceGenerationsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ListDateSourceGenerationsRequest) GetValidStatus() *bool {
	return s.ValidStatus
}

func (s *ListDateSourceGenerationsRequest) SetDomainName(v string) *ListDateSourceGenerationsRequest {
	s.DomainName = &v
	return s
}

func (s *ListDateSourceGenerationsRequest) SetValidStatus(v bool) *ListDateSourceGenerationsRequest {
	s.ValidStatus = &v
	return s
}

func (s *ListDateSourceGenerationsRequest) Validate() error {
	return dara.Validate(s)
}
