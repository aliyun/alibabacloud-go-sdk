// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppRagConfigRequest
	GetAppId() *string
	SetPromptStrategy(v string) *UpdateMmAppRagConfigRequest
	GetPromptStrategy() *string
	SetRetrieveMaxLength(v int32) *UpdateMmAppRagConfigRequest
	GetRetrieveMaxLength() *int32
	SetTopK(v int32) *UpdateMmAppRagConfigRequest
	GetTopK() *int32
	SetWorkspaceId(v string) *UpdateMmAppRagConfigRequest
	GetWorkspaceId() *string
}

type UpdateMmAppRagConfigRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	PromptStrategy *string `json:"PromptStrategy,omitempty" xml:"PromptStrategy,omitempty"`
	// This parameter is required.
	RetrieveMaxLength *int32 `json:"RetrieveMaxLength,omitempty" xml:"RetrieveMaxLength,omitempty"`
	TopK              *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMmAppRagConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppRagConfigRequest) GetPromptStrategy() *string {
	return s.PromptStrategy
}

func (s *UpdateMmAppRagConfigRequest) GetRetrieveMaxLength() *int32 {
	return s.RetrieveMaxLength
}

func (s *UpdateMmAppRagConfigRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *UpdateMmAppRagConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppRagConfigRequest) SetAppId(v string) *UpdateMmAppRagConfigRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppRagConfigRequest) SetPromptStrategy(v string) *UpdateMmAppRagConfigRequest {
	s.PromptStrategy = &v
	return s
}

func (s *UpdateMmAppRagConfigRequest) SetRetrieveMaxLength(v int32) *UpdateMmAppRagConfigRequest {
	s.RetrieveMaxLength = &v
	return s
}

func (s *UpdateMmAppRagConfigRequest) SetTopK(v int32) *UpdateMmAppRagConfigRequest {
	s.TopK = &v
	return s
}

func (s *UpdateMmAppRagConfigRequest) SetWorkspaceId(v string) *UpdateMmAppRagConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppRagConfigRequest) Validate() error {
	return dara.Validate(s)
}
