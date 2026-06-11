// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentName(v string) *DeleteDocumentRequest
	GetDocumentName() *string
	SetKbUuid(v string) *DeleteDocumentRequest
	GetKbUuid() *string
}

type DeleteDocumentRequest struct {
	// The document name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.md
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
}

func (s DeleteDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *DeleteDocumentRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DeleteDocumentRequest) SetDocumentName(v string) *DeleteDocumentRequest {
	s.DocumentName = &v
	return s
}

func (s *DeleteDocumentRequest) SetKbUuid(v string) *DeleteDocumentRequest {
	s.KbUuid = &v
	return s
}

func (s *DeleteDocumentRequest) Validate() error {
	return dara.Validate(s)
}
