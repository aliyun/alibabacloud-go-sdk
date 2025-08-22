// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomDomainSampleRateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *DeleteCustomDomainSampleRateRequest
	GetDomainNames() *string
}

type DeleteCustomDomainSampleRateRequest struct {
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
}

func (s DeleteCustomDomainSampleRateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomDomainSampleRateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomDomainSampleRateRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *DeleteCustomDomainSampleRateRequest) SetDomainNames(v string) *DeleteCustomDomainSampleRateRequest {
	s.DomainNames = &v
	return s
}

func (s *DeleteCustomDomainSampleRateRequest) Validate() error {
	return dara.Validate(s)
}
