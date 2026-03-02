// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultDomainName(v string) *SetDefaultDomainRequest
	GetDefaultDomainName() *string
}

type SetDefaultDomainRequest struct {
	// This parameter is required.
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty"`
}

func (s SetDefaultDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultDomainRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainRequest) GetDefaultDomainName() *string {
	return s.DefaultDomainName
}

func (s *SetDefaultDomainRequest) SetDefaultDomainName(v string) *SetDefaultDomainRequest {
	s.DefaultDomainName = &v
	return s
}

func (s *SetDefaultDomainRequest) Validate() error {
	return dara.Validate(s)
}
