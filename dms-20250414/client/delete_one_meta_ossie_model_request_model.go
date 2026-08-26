// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOneMetaOssieModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeUuid(v string) *DeleteOneMetaOssieModelRequest
	GetKnowledgeUuid() *string
}

type DeleteOneMetaOssieModelRequest struct {
	// The UUID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86c5c290052147c***
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
}

func (s DeleteOneMetaOssieModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOneMetaOssieModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteOneMetaOssieModelRequest) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *DeleteOneMetaOssieModelRequest) SetKnowledgeUuid(v string) *DeleteOneMetaOssieModelRequest {
	s.KnowledgeUuid = &v
	return s
}

func (s *DeleteOneMetaOssieModelRequest) Validate() error {
	return dara.Validate(s)
}
