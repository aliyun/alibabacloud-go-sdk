// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTenantBindNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumber(v string) *DescribeTenantBindNumberRequest
	GetNumber() *string
}

type DescribeTenantBindNumberRequest struct {
	// example:
	//
	// 95187
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribeTenantBindNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTenantBindNumberRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantBindNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *DescribeTenantBindNumberRequest) SetNumber(v string) *DescribeTenantBindNumberRequest {
	s.Number = &v
	return s
}

func (s *DescribeTenantBindNumberRequest) Validate() error {
	return dara.Validate(s)
}
