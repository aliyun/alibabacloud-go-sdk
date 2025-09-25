// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibraryId(v string) *GetLibraryRequest
	GetLibraryId() *string
}

type GetLibraryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cjshcxxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s GetLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryRequest) GoString() string {
	return s.String()
}

func (s *GetLibraryRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GetLibraryRequest) SetLibraryId(v string) *GetLibraryRequest {
	s.LibraryId = &v
	return s
}

func (s *GetLibraryRequest) Validate() error {
	return dara.Validate(s)
}
