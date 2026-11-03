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
	// The short namespace ID. This ID does not require a region prefix. Use this parameter.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The long namespace ID. If specified, this parameter takes precedence and NameSpaceShortId is ignored. This parameter is retained for backward compatibility. Use the short namespace ID to simplify the request.
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
