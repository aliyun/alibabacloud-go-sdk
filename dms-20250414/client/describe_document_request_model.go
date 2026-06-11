// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentName(v string) *DescribeDocumentRequest
	GetDocumentName() *string
	SetKbUuid(v string) *DescribeDocumentRequest
	GetKbUuid() *string
}

type DescribeDocumentRequest struct {
	// The name of the document.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.md
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
}

func (s DescribeDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocumentRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocumentRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *DescribeDocumentRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DescribeDocumentRequest) SetDocumentName(v string) *DescribeDocumentRequest {
	s.DocumentName = &v
	return s
}

func (s *DescribeDocumentRequest) SetKbUuid(v string) *DescribeDocumentRequest {
	s.KbUuid = &v
	return s
}

func (s *DescribeDocumentRequest) Validate() error {
	return dara.Validate(s)
}
