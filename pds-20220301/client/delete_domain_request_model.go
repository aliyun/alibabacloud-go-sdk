// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *DeleteDomainRequest
	GetDomainId() *string
}

type DeleteDomainRequest struct {
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *DeleteDomainRequest) SetDomainId(v string) *DeleteDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteDomainRequest) Validate() error {
	return dara.Validate(s)
}
