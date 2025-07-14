// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWebApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *StartWebApplicationRequest
	GetNamespaceId() *string
}

type StartWebApplicationRequest struct {
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s StartWebApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartWebApplicationRequest) GoString() string {
	return s.String()
}

func (s *StartWebApplicationRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *StartWebApplicationRequest) SetNamespaceId(v string) *StartWebApplicationRequest {
	s.NamespaceId = &v
	return s
}

func (s *StartWebApplicationRequest) Validate() error {
	return dara.Validate(s)
}
