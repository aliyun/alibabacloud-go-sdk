// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerKnowledgeBaseSyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *TriggerKnowledgeBaseSyncRequest
	GetKnowledgeBaseId() *string
	SetLinkId(v string) *TriggerKnowledgeBaseSyncRequest
	GetLinkId() *string
	SetRegionId(v string) *TriggerKnowledgeBaseSyncRequest
	GetRegionId() *string
}

type TriggerKnowledgeBaseSyncRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// example:
	//
	// pkbl-xxxxx
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s TriggerKnowledgeBaseSyncRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerKnowledgeBaseSyncRequest) GoString() string {
	return s.String()
}

func (s *TriggerKnowledgeBaseSyncRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *TriggerKnowledgeBaseSyncRequest) GetLinkId() *string {
	return s.LinkId
}

func (s *TriggerKnowledgeBaseSyncRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TriggerKnowledgeBaseSyncRequest) SetKnowledgeBaseId(v string) *TriggerKnowledgeBaseSyncRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *TriggerKnowledgeBaseSyncRequest) SetLinkId(v string) *TriggerKnowledgeBaseSyncRequest {
	s.LinkId = &v
	return s
}

func (s *TriggerKnowledgeBaseSyncRequest) SetRegionId(v string) *TriggerKnowledgeBaseSyncRequest {
	s.RegionId = &v
	return s
}

func (s *TriggerKnowledgeBaseSyncRequest) Validate() error {
	return dara.Validate(s)
}
