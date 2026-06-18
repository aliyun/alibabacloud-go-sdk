// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileIdsShrink(v string) *DeleteFilesShrinkRequest
	GetFileIdsShrink() *string
}

type DeleteFilesShrinkRequest struct {
	// The list of IDs of the files to be deleted. A maximum of 20 files can be deleted in a single request.
	//
	// This parameter is required.
	FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
}

func (s DeleteFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilesShrinkRequest) GetFileIdsShrink() *string {
	return s.FileIdsShrink
}

func (s *DeleteFilesShrinkRequest) SetFileIdsShrink(v string) *DeleteFilesShrinkRequest {
	s.FileIdsShrink = &v
	return s
}

func (s *DeleteFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
