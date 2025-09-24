// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListQuotasRequest
	GetName() *string
}

type ListQuotasRequest struct {
	// The quota name. Fuzzy search is supported.
	//
	// example:
	//
	// quota-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasRequest) GetName() *string {
	return s.Name
}

func (s *ListQuotasRequest) SetName(v string) *ListQuotasRequest {
	s.Name = &v
	return s
}

func (s *ListQuotasRequest) Validate() error {
	return dara.Validate(s)
}
