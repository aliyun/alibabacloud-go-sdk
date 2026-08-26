// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateKnowledgeBaseRequest
	GetDescription() *string
	SetKnowledgeBaseType(v string) *CreateKnowledgeBaseRequest
	GetKnowledgeBaseType() *string
	SetKnowledgeSpaceId(v string) *CreateKnowledgeBaseRequest
	GetKnowledgeSpaceId() *string
	SetName(v string) *CreateKnowledgeBaseRequest
	GetName() *string
	SetRegionId(v string) *CreateKnowledgeBaseRequest
	GetRegionId() *string
	SetSearchMode(v string) *CreateKnowledgeBaseRequest
	GetSearchMode() *string
}

type CreateKnowledgeBaseRequest struct {
	// The description of the knowledge base.
	//
	// example:
	//
	// testkbDesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the knowledge base: PERSONAL or PUBLIC.
	//
	// example:
	//
	// PUBLIC
	KnowledgeBaseType *string `json:"KnowledgeBaseType,omitempty" xml:"KnowledgeBaseType,omitempty"`
	// The unique identifier of the knowledge space.
	//
	// This parameter is required.
	//
	// example:
	//
	// pks-xxxxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// The name of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// testkb
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The search mode. Valid values:
	//
	// 	- balanced (default): balanced mode
	//
	// 	- precise: precise mode
	//
	// 	- semantic: semantic mode
	//
	// 	- knn: KNN mode
	//
	// 	- rrf: reciprocal rank fusion
	//
	// example:
	//
	// balanced
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
}

func (s CreateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseRequest) GetKnowledgeBaseType() *string {
	return s.KnowledgeBaseType
}

func (s *CreateKnowledgeBaseRequest) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *CreateKnowledgeBaseRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateKnowledgeBaseRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *CreateKnowledgeBaseRequest) SetDescription(v string) *CreateKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetKnowledgeBaseType(v string) *CreateKnowledgeBaseRequest {
	s.KnowledgeBaseType = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetKnowledgeSpaceId(v string) *CreateKnowledgeBaseRequest {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetName(v string) *CreateKnowledgeBaseRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetRegionId(v string) *CreateKnowledgeBaseRequest {
	s.RegionId = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetSearchMode(v string) *CreateKnowledgeBaseRequest {
	s.SearchMode = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
