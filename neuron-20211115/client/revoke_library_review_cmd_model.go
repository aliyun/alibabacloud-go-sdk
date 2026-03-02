// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeLibraryReviewCmd interface {
	dara.Model
	String() string
	GoString() string
	SetLibraryId(v int64) *RevokeLibraryReviewCmd
	GetLibraryId() *int64
}

type RevokeLibraryReviewCmd struct {
	LibraryId *int64 `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s RevokeLibraryReviewCmd) String() string {
	return dara.Prettify(s)
}

func (s RevokeLibraryReviewCmd) GoString() string {
	return s.String()
}

func (s *RevokeLibraryReviewCmd) GetLibraryId() *int64 {
	return s.LibraryId
}

func (s *RevokeLibraryReviewCmd) SetLibraryId(v int64) *RevokeLibraryReviewCmd {
	s.LibraryId = &v
	return s
}

func (s *RevokeLibraryReviewCmd) Validate() error {
	return dara.Validate(s)
}
