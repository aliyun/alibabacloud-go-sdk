// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentId(v string) *PreviewDocumentRequest
	GetDocumentId() *string
}

type PreviewDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8326472354762354
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
}

func (s PreviewDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewDocumentRequest) GoString() string {
	return s.String()
}

func (s *PreviewDocumentRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *PreviewDocumentRequest) SetDocumentId(v string) *PreviewDocumentRequest {
	s.DocumentId = &v
	return s
}

func (s *PreviewDocumentRequest) Validate() error {
	return dara.Validate(s)
}
