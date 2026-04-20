// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMmAppBindingRagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *MmAppBindingRagShrinkRequest
	GetAppId() *string
	SetKnowledgeBaseCodeListShrink(v string) *MmAppBindingRagShrinkRequest
	GetKnowledgeBaseCodeListShrink() *string
	SetWorkspaceId(v string) *MmAppBindingRagShrinkRequest
	GetWorkspaceId() *string
}

type MmAppBindingRagShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_a2eb4e04b48041108edb1f6de815
	AppId                       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	KnowledgeBaseCodeListShrink *string `json:"KnowledgeBaseCodeList,omitempty" xml:"KnowledgeBaseCodeList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-6uhm7nfev4k8pwcz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s MmAppBindingRagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingRagShrinkRequest) GoString() string {
	return s.String()
}

func (s *MmAppBindingRagShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *MmAppBindingRagShrinkRequest) GetKnowledgeBaseCodeListShrink() *string {
	return s.KnowledgeBaseCodeListShrink
}

func (s *MmAppBindingRagShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *MmAppBindingRagShrinkRequest) SetAppId(v string) *MmAppBindingRagShrinkRequest {
	s.AppId = &v
	return s
}

func (s *MmAppBindingRagShrinkRequest) SetKnowledgeBaseCodeListShrink(v string) *MmAppBindingRagShrinkRequest {
	s.KnowledgeBaseCodeListShrink = &v
	return s
}

func (s *MmAppBindingRagShrinkRequest) SetWorkspaceId(v string) *MmAppBindingRagShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *MmAppBindingRagShrinkRequest) Validate() error {
	return dara.Validate(s)
}
