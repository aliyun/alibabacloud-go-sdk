// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryProcessorNersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ListQueryProcessorNersRequest
	GetDomain() *string
}

type ListQueryProcessorNersRequest struct {
	// The type of the industry.
	//
	// 	- ECOMMERCE
	//
	// example:
	//
	// ECOMMERCE
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
}

func (s ListQueryProcessorNersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorNersRequest) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersRequest) GetDomain() *string {
	return s.Domain
}

func (s *ListQueryProcessorNersRequest) SetDomain(v string) *ListQueryProcessorNersRequest {
	s.Domain = &v
	return s
}

func (s *ListQueryProcessorNersRequest) Validate() error {
	return dara.Validate(s)
}
