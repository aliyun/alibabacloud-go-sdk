// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DeleteWebApplicationRequest
	GetNamespaceId() *string
}

type DeleteWebApplicationRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteWebApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebApplicationRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteWebApplicationRequest) SetNamespaceId(v string) *DeleteWebApplicationRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteWebApplicationRequest) Validate() error {
	return dara.Validate(s)
}
