// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPromptVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *SubmitPromptVersionRequest
	GetNamespaceId() *string
	SetPromptKey(v string) *SubmitPromptVersionRequest
	GetPromptKey() *string
	SetPromptVersion(v string) *SubmitPromptVersionRequest
	GetPromptVersion() *string
}

type SubmitPromptVersionRequest struct {
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
	// example:
	//
	// 0.0.1
	PromptVersion *string `json:"PromptVersion,omitempty" xml:"PromptVersion,omitempty"`
}

func (s SubmitPromptVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPromptVersionRequest) GoString() string {
	return s.String()
}

func (s *SubmitPromptVersionRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *SubmitPromptVersionRequest) GetPromptKey() *string {
	return s.PromptKey
}

func (s *SubmitPromptVersionRequest) GetPromptVersion() *string {
	return s.PromptVersion
}

func (s *SubmitPromptVersionRequest) SetNamespaceId(v string) *SubmitPromptVersionRequest {
	s.NamespaceId = &v
	return s
}

func (s *SubmitPromptVersionRequest) SetPromptKey(v string) *SubmitPromptVersionRequest {
	s.PromptKey = &v
	return s
}

func (s *SubmitPromptVersionRequest) SetPromptVersion(v string) *SubmitPromptVersionRequest {
	s.PromptVersion = &v
	return s
}

func (s *SubmitPromptVersionRequest) Validate() error {
	return dara.Validate(s)
}
