// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCustomDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DeleteWebCustomDomainRequest
	GetNamespaceId() *string
}

type DeleteWebCustomDomainRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteWebCustomDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebCustomDomainRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteWebCustomDomainRequest) SetNamespaceId(v string) *DeleteWebCustomDomainRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteWebCustomDomainRequest) Validate() error {
	return dara.Validate(s)
}
