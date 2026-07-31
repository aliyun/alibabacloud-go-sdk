// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelId(v string) *ListVoiceRequest
	GetModelId() *string
	SetWorkspaceId(v string) *ListVoiceRequest
	GetWorkspaceId() *string
}

type ListVoiceRequest struct {
	// This parameter is required.
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceRequest) GoString() string {
	return s.String()
}

func (s *ListVoiceRequest) GetModelId() *string {
	return s.ModelId
}

func (s *ListVoiceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListVoiceRequest) SetModelId(v string) *ListVoiceRequest {
	s.ModelId = &v
	return s
}

func (s *ListVoiceRequest) SetWorkspaceId(v string) *ListVoiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListVoiceRequest) Validate() error {
	return dara.Validate(s)
}
