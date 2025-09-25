// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibraryId(v string) *DeleteLibraryRequest
	GetLibraryId() *string
}

type DeleteLibraryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// skdfefxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s DeleteLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLibraryRequest) GoString() string {
	return s.String()
}

func (s *DeleteLibraryRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *DeleteLibraryRequest) SetLibraryId(v string) *DeleteLibraryRequest {
	s.LibraryId = &v
	return s
}

func (s *DeleteLibraryRequest) Validate() error {
	return dara.Validate(s)
}
