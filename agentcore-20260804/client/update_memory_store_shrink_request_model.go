// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryStoreShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateMemoryStoreShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateMemoryStoreShrinkRequest
	GetClientToken() *string
}

type UpdateMemoryStoreShrinkRequest struct {
	// The update MemoryStore request.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client request token.
	//
	// example:
	//
	// memory-store-update-20260909-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateMemoryStoreShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryStoreShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryStoreShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateMemoryStoreShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateMemoryStoreShrinkRequest) SetBodyShrink(v string) *UpdateMemoryStoreShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateMemoryStoreShrinkRequest) SetClientToken(v string) *UpdateMemoryStoreShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateMemoryStoreShrinkRequest) Validate() error {
	return dara.Validate(s)
}
