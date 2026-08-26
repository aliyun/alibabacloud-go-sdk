// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOneMetaOssieModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocFormat(v string) *GetOneMetaOssieModelRequest
	GetDocFormat() *string
	SetKnowledgeUuid(v string) *GetOneMetaOssieModelRequest
	GetKnowledgeUuid() *string
}

type GetOneMetaOssieModelRequest struct {
	// The document type of the semantic model. Valid values: JSON and YAML.
	//
	// This parameter is required.
	//
	// example:
	//
	// JSON
	DocFormat *string `json:"DocFormat,omitempty" xml:"DocFormat,omitempty"`
	// The UUID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86c5c290052147c***
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
}

func (s GetOneMetaOssieModelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOneMetaOssieModelRequest) GoString() string {
	return s.String()
}

func (s *GetOneMetaOssieModelRequest) GetDocFormat() *string {
	return s.DocFormat
}

func (s *GetOneMetaOssieModelRequest) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *GetOneMetaOssieModelRequest) SetDocFormat(v string) *GetOneMetaOssieModelRequest {
	s.DocFormat = &v
	return s
}

func (s *GetOneMetaOssieModelRequest) SetKnowledgeUuid(v string) *GetOneMetaOssieModelRequest {
	s.KnowledgeUuid = &v
	return s
}

func (s *GetOneMetaOssieModelRequest) Validate() error {
	return dara.Validate(s)
}
