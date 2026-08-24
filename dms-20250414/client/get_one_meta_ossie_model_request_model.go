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
	// This parameter is required.
	DocFormat *string `json:"DocFormat,omitempty" xml:"DocFormat,omitempty"`
	// This parameter is required.
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
