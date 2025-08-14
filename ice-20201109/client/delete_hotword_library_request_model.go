// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotwordLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotwordLibraryId(v string) *DeleteHotwordLibraryRequest
	GetHotwordLibraryId() *string
}

type DeleteHotwordLibraryRequest struct {
	// example:
	//
	// ****cdb3e74639973036bc84****
	HotwordLibraryId *string `json:"HotwordLibraryId,omitempty" xml:"HotwordLibraryId,omitempty"`
}

func (s DeleteHotwordLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotwordLibraryRequest) GoString() string {
	return s.String()
}

func (s *DeleteHotwordLibraryRequest) GetHotwordLibraryId() *string {
	return s.HotwordLibraryId
}

func (s *DeleteHotwordLibraryRequest) SetHotwordLibraryId(v string) *DeleteHotwordLibraryRequest {
	s.HotwordLibraryId = &v
	return s
}

func (s *DeleteHotwordLibraryRequest) Validate() error {
	return dara.Validate(s)
}
