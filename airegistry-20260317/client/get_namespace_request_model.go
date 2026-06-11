// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *GetNamespaceRequest
	GetNamespaceId() *string
}

type GetNamespaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetNamespaceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetNamespaceRequest) SetNamespaceId(v string) *GetNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
