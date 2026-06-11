// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseUploadSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKbUuid(v string) *DescribeKnowledgeBaseUploadSignatureRequest
	GetKbUuid() *string
}

type DescribeKnowledgeBaseUploadSignatureRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
}

func (s DescribeKnowledgeBaseUploadSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseUploadSignatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseUploadSignatureRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DescribeKnowledgeBaseUploadSignatureRequest) SetKbUuid(v string) *DescribeKnowledgeBaseUploadSignatureRequest {
	s.KbUuid = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureRequest) Validate() error {
	return dara.Validate(s)
}
