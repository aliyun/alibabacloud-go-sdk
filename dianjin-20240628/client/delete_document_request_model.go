// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocIds(v []*string) *DeleteDocumentRequest
	GetDocIds() []*string
	SetLibraryId(v string) *DeleteDocumentRequest
	GetLibraryId() *string
}

type DeleteDocumentRequest struct {
	// This parameter is required.
	DocIds []*string `json:"docIds,omitempty" xml:"docIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s DeleteDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentRequest) GetDocIds() []*string {
	return s.DocIds
}

func (s *DeleteDocumentRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *DeleteDocumentRequest) SetDocIds(v []*string) *DeleteDocumentRequest {
	s.DocIds = v
	return s
}

func (s *DeleteDocumentRequest) SetLibraryId(v string) *DeleteDocumentRequest {
	s.LibraryId = &v
	return s
}

func (s *DeleteDocumentRequest) Validate() error {
	return dara.Validate(s)
}
