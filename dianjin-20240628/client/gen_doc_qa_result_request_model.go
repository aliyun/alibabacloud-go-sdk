// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenDocQaResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *GenDocQaResultRequest
	GetDocId() *string
	SetLibraryId(v string) *GenDocQaResultRequest
	GetLibraryId() *string
	SetRequestId(v string) *GenDocQaResultRequest
	GetRequestId() *string
}

type GenDocQaResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 182364872346
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdgdsfg
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenDocQaResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GenDocQaResultRequest) GoString() string {
	return s.String()
}

func (s *GenDocQaResultRequest) GetDocId() *string {
	return s.DocId
}

func (s *GenDocQaResultRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GenDocQaResultRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GenDocQaResultRequest) SetDocId(v string) *GenDocQaResultRequest {
	s.DocId = &v
	return s
}

func (s *GenDocQaResultRequest) SetLibraryId(v string) *GenDocQaResultRequest {
	s.LibraryId = &v
	return s
}

func (s *GenDocQaResultRequest) SetRequestId(v string) *GenDocQaResultRequest {
	s.RequestId = &v
	return s
}

func (s *GenDocQaResultRequest) Validate() error {
	return dara.Validate(s)
}
