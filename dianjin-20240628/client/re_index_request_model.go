// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentId(v string) *ReIndexRequest
	GetDocumentId() *string
}

type ReIndexRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8326472354762354
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
}

func (s ReIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s ReIndexRequest) GoString() string {
	return s.String()
}

func (s *ReIndexRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *ReIndexRequest) SetDocumentId(v string) *ReIndexRequest {
	s.DocumentId = &v
	return s
}

func (s *ReIndexRequest) Validate() error {
	return dara.Validate(s)
}
