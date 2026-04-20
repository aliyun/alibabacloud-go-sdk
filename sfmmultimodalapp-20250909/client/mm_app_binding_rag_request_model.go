// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMmAppBindingRagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *MmAppBindingRagRequest
	GetAppId() *string
	SetKnowledgeBaseCodeList(v []*string) *MmAppBindingRagRequest
	GetKnowledgeBaseCodeList() []*string
	SetWorkspaceId(v string) *MmAppBindingRagRequest
	GetWorkspaceId() *string
}

type MmAppBindingRagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_a2eb4e04b48041108edb1f6de815
	AppId                 *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	KnowledgeBaseCodeList []*string `json:"KnowledgeBaseCodeList,omitempty" xml:"KnowledgeBaseCodeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-6uhm7nfev4k8pwcz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s MmAppBindingRagRequest) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingRagRequest) GoString() string {
	return s.String()
}

func (s *MmAppBindingRagRequest) GetAppId() *string {
	return s.AppId
}

func (s *MmAppBindingRagRequest) GetKnowledgeBaseCodeList() []*string {
	return s.KnowledgeBaseCodeList
}

func (s *MmAppBindingRagRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *MmAppBindingRagRequest) SetAppId(v string) *MmAppBindingRagRequest {
	s.AppId = &v
	return s
}

func (s *MmAppBindingRagRequest) SetKnowledgeBaseCodeList(v []*string) *MmAppBindingRagRequest {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *MmAppBindingRagRequest) SetWorkspaceId(v string) *MmAppBindingRagRequest {
	s.WorkspaceId = &v
	return s
}

func (s *MmAppBindingRagRequest) Validate() error {
	return dara.Validate(s)
}
