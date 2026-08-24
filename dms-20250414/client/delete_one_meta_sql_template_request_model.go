// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOneMetaSqlTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeUuid(v string) *DeleteOneMetaSqlTemplateRequest
	GetKnowledgeUuid() *string
}

type DeleteOneMetaSqlTemplateRequest struct {
	// This parameter is required.
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
}

func (s DeleteOneMetaSqlTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOneMetaSqlTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteOneMetaSqlTemplateRequest) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *DeleteOneMetaSqlTemplateRequest) SetKnowledgeUuid(v string) *DeleteOneMetaSqlTemplateRequest {
	s.KnowledgeUuid = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateRequest) Validate() error {
	return dara.Validate(s)
}
