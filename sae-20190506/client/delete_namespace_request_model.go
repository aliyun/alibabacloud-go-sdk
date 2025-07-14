// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNameSpaceShortId(v string) *DeleteNamespaceRequest
	GetNameSpaceShortId() *string
	SetNamespaceId(v string) *DeleteNamespaceRequest
	GetNamespaceId() *string
}

type DeleteNamespaceRequest struct {
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// cn-beijing:test
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *DeleteNamespaceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteNamespaceRequest) SetNameSpaceShortId(v string) *DeleteNamespaceRequest {
	s.NameSpaceShortId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetNamespaceId(v string) *DeleteNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
