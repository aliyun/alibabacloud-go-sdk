// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateKnowledgeBaseRequest
	GetDescription() *string
	SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseRequest
	GetKnowledgeBaseId() *string
	SetName(v string) *UpdateKnowledgeBaseRequest
	GetName() *string
	SetRegionId(v string) *UpdateKnowledgeBaseRequest
	GetRegionId() *string
	SetSearchMode(v string) *UpdateKnowledgeBaseRequest
	GetSearchMode() *string
}

type UpdateKnowledgeBaseRequest struct {
	// The description of the knowledge base.
	//
	// example:
	//
	// test desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The name of the knowledge base.
	//
	// example:
	//
	// testName
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
	// 	- balanced (default): balanced mode.
	//
	// 	- precise: precise mode.
	//
	// 	- semantic: semantic mode.
	//
	// 	- knn: KNN mode.
	//
	// 	- rrf: reciprocal rank fusion (RRF) mode.
	//
	// example:
	//
	// balanced
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
}

func (s UpdateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeBaseRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKnowledgeBaseRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *UpdateKnowledgeBaseRequest) SetDescription(v string) *UpdateKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetName(v string) *UpdateKnowledgeBaseRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetRegionId(v string) *UpdateKnowledgeBaseRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetSearchMode(v string) *UpdateKnowledgeBaseRequest {
	s.SearchMode = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
