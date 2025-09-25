// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentId(v string) *GetDocumentUrlRequest
	GetDocumentId() *string
}

type GetDocumentUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12681367362
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
}

func (s GetDocumentUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentUrlRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetDocumentUrlRequest) SetDocumentId(v string) *GetDocumentUrlRequest {
	s.DocumentId = &v
	return s
}

func (s *GetDocumentUrlRequest) Validate() error {
	return dara.Validate(s)
}
