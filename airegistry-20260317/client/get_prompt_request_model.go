// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *GetPromptRequest
	GetNamespaceId() *string
	SetPromptKey(v string) *GetPromptRequest
	GetPromptKey() *string
}

type GetPromptRequest struct {
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

func (s GetPromptRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPromptRequest) GoString() string {
	return s.String()
}

func (s *GetPromptRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetPromptRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *GetPromptRequest) SetNamespaceId(v string) *GetPromptRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetPromptRequest) SetPromptKey(v string) *GetPromptRequest {
	s.PromptKey = &v
	return s
}

func (s *GetPromptRequest) Validate() error {
	return dara.Validate(s)
}
