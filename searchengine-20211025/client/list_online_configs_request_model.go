// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnlineConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ListOnlineConfigsRequest
	GetDomain() *string
}

type ListOnlineConfigsRequest struct {
	// The name of the domain
	//
	// This parameter is required.
	//
	// example:
	//
	// sz_vpc_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
}

func (s ListOnlineConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsRequest) GetDomain() *string {
	return s.Domain
}

func (s *ListOnlineConfigsRequest) SetDomain(v string) *ListOnlineConfigsRequest {
	s.Domain = &v
	return s
}

func (s *ListOnlineConfigsRequest) Validate() error {
	return dara.Validate(s)
}
