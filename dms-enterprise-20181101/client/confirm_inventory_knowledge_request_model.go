// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmInventoryKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityId(v int64) *ConfirmInventoryKnowledgeRequest
	GetEntityId() *int64
	SetJobId(v int64) *ConfirmInventoryKnowledgeRequest
	GetJobId() *int64
	SetKnowledgeType(v string) *ConfirmInventoryKnowledgeRequest
	GetKnowledgeType() *string
}

type ConfirmInventoryKnowledgeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2001
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	KnowledgeType *string `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
}

func (s ConfirmInventoryKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmInventoryKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *ConfirmInventoryKnowledgeRequest) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ConfirmInventoryKnowledgeRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ConfirmInventoryKnowledgeRequest) GetKnowledgeType() *string {
	return s.KnowledgeType
}

func (s *ConfirmInventoryKnowledgeRequest) SetEntityId(v int64) *ConfirmInventoryKnowledgeRequest {
	s.EntityId = &v
	return s
}

func (s *ConfirmInventoryKnowledgeRequest) SetJobId(v int64) *ConfirmInventoryKnowledgeRequest {
	s.JobId = &v
	return s
}

func (s *ConfirmInventoryKnowledgeRequest) SetKnowledgeType(v string) *ConfirmInventoryKnowledgeRequest {
	s.KnowledgeType = &v
	return s
}

func (s *ConfirmInventoryKnowledgeRequest) Validate() error {
	return dara.Validate(s)
}
