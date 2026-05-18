// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentName(v string) *UpdateDocumentRequest
	GetDocumentName() *string
	SetKbUuid(v string) *UpdateDocumentRequest
	GetKbUuid() *string
	SetNewDescription(v string) *UpdateDocumentRequest
	GetNewDescription() *string
}

type UpdateDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test.md
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test123
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
}

func (s UpdateDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocumentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocumentRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *UpdateDocumentRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *UpdateDocumentRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdateDocumentRequest) SetDocumentName(v string) *UpdateDocumentRequest {
	s.DocumentName = &v
	return s
}

func (s *UpdateDocumentRequest) SetKbUuid(v string) *UpdateDocumentRequest {
	s.KbUuid = &v
	return s
}

func (s *UpdateDocumentRequest) SetNewDescription(v string) *UpdateDocumentRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateDocumentRequest) Validate() error {
	return dara.Validate(s)
}
