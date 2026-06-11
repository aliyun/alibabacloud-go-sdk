// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *GetPromptVersionRequest
	GetNamespaceId() *string
	SetPromptKey(v string) *GetPromptVersionRequest
	GetPromptKey() *string
	SetPromptVersion(v string) *GetPromptVersionRequest
	GetPromptVersion() *string
}

type GetPromptVersionRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	PromptVersion *string `json:"PromptVersion,omitempty" xml:"PromptVersion,omitempty"`
}

func (s GetPromptVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPromptVersionRequest) GoString() string {
	return s.String()
}

func (s *GetPromptVersionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetPromptVersionRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *GetPromptVersionRequest) GetPromptVersion() *string {
	return s.PromptVersion
}

func (s *GetPromptVersionRequest) SetNamespaceId(v string) *GetPromptVersionRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetPromptVersionRequest) SetPromptKey(v string) *GetPromptVersionRequest {
	s.PromptKey = &v
	return s
}

func (s *GetPromptVersionRequest) SetPromptVersion(v string) *GetPromptVersionRequest {
	s.PromptVersion = &v
	return s
}

func (s *GetPromptVersionRequest) Validate() error {
	return dara.Validate(s)
}
