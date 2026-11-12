// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryStoreShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateMemoryStoreShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateMemoryStoreShrinkRequest
	GetClientToken() *string
}

type CreateMemoryStoreShrinkRequest struct {
	// The request to create a memory store.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client request token.
	//
	// example:
	//
	// memory-store-create-20260909-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateMemoryStoreShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateMemoryStoreShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMemoryStoreShrinkRequest) SetBodyShrink(v string) *CreateMemoryStoreShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateMemoryStoreShrinkRequest) SetClientToken(v string) *CreateMemoryStoreShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMemoryStoreShrinkRequest) Validate() error {
	return dara.Validate(s)
}
