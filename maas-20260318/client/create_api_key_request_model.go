// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateApiKeyRequest
	GetDescription() *string
	SetWorkspaceId(v string) *CreateApiKeyRequest
	GetWorkspaceId() *string
}

type CreateApiKeyRequest struct {
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-y3kv9qctnlytgmga
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateApiKeyRequest) SetDescription(v string) *CreateApiKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateApiKeyRequest) SetWorkspaceId(v string) *CreateApiKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
