// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DeletePromptRequest
	GetNamespaceId() *string
	SetPromptKey(v string) *DeletePromptRequest
	GetPromptKey() *string
}

type DeletePromptRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer-service-qa
	PromptKey *string `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
}

func (s DeletePromptRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePromptRequest) GoString() string {
	return s.String()
}

func (s *DeletePromptRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeletePromptRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *DeletePromptRequest) SetNamespaceId(v string) *DeletePromptRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeletePromptRequest) SetPromptKey(v string) *DeletePromptRequest {
	s.PromptKey = &v
	return s
}

func (s *DeletePromptRequest) Validate() error {
	return dara.Validate(s)
}
