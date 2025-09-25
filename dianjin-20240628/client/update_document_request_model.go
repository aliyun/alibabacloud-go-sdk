// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *UpdateDocumentRequest
	GetDocId() *string
	SetLibraryId(v string) *UpdateDocumentRequest
	GetLibraryId() *string
	SetMeta(v map[string]interface{}) *UpdateDocumentRequest
	GetMeta() map[string]interface{}
	SetTitle(v string) *UpdateDocumentRequest
	GetTitle() *string
}

type UpdateDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// {
	//
	//         "businessId": "12321"
	//
	//     }
	Meta map[string]interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocumentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocumentRequest) GetDocId() *string {
	return s.DocId
}

func (s *UpdateDocumentRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *UpdateDocumentRequest) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *UpdateDocumentRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateDocumentRequest) SetDocId(v string) *UpdateDocumentRequest {
	s.DocId = &v
	return s
}

func (s *UpdateDocumentRequest) SetLibraryId(v string) *UpdateDocumentRequest {
	s.LibraryId = &v
	return s
}

func (s *UpdateDocumentRequest) SetMeta(v map[string]interface{}) *UpdateDocumentRequest {
	s.Meta = v
	return s
}

func (s *UpdateDocumentRequest) SetTitle(v string) *UpdateDocumentRequest {
	s.Title = &v
	return s
}

func (s *UpdateDocumentRequest) Validate() error {
	return dara.Validate(s)
}
