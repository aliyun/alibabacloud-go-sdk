// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebApplicationRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DeleteWebApplicationRevisionRequest
	GetNamespaceId() *string
}

type DeleteWebApplicationRevisionRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteWebApplicationRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebApplicationRevisionRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebApplicationRevisionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteWebApplicationRevisionRequest) SetNamespaceId(v string) *DeleteWebApplicationRevisionRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteWebApplicationRevisionRequest) Validate() error {
	return dara.Validate(s)
}
