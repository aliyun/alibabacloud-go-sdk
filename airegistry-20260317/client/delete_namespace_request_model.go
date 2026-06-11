// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DeleteNamespaceRequest
	GetNamespaceId() *string
}

type DeleteNamespaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteNamespaceRequest) SetNamespaceId(v string) *DeleteNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
