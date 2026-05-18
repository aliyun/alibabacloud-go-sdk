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
	SetKbUuid(v string) *UpdateKnowledgeBaseRequest
	GetKbUuid() *string
	SetName(v string) *UpdateKnowledgeBaseRequest
	GetName() *string
}

type UpdateKnowledgeBaseRequest struct {
	// example:
	//
	// updated kb desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kb-HZ-zgr1***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// example:
	//
	// updated kb name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *UpdateKnowledgeBaseRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *UpdateKnowledgeBaseRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeBaseRequest) SetDescription(v string) *UpdateKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetKbUuid(v string) *UpdateKnowledgeBaseRequest {
	s.KbUuid = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetName(v string) *UpdateKnowledgeBaseRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
