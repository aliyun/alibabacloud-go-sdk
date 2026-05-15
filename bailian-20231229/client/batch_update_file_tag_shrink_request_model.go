// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFileTagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileInfosShrink(v string) *BatchUpdateFileTagShrinkRequest
	GetFileInfosShrink() *string
	SetUpdateMode(v string) *BatchUpdateFileTagShrinkRequest
	GetUpdateMode() *string
}

type BatchUpdateFileTagShrinkRequest struct {
	// This parameter is required.
	FileInfosShrink *string `json:"FileInfos,omitempty" xml:"FileInfos,omitempty"`
	// example:
	//
	// OVERWRITE
	UpdateMode *string `json:"UpdateMode,omitempty" xml:"UpdateMode,omitempty"`
}

func (s BatchUpdateFileTagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFileTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFileTagShrinkRequest) GetFileInfosShrink() *string {
	return s.FileInfosShrink
}

func (s *BatchUpdateFileTagShrinkRequest) GetUpdateMode() *string {
	return s.UpdateMode
}

func (s *BatchUpdateFileTagShrinkRequest) SetFileInfosShrink(v string) *BatchUpdateFileTagShrinkRequest {
	s.FileInfosShrink = &v
	return s
}

func (s *BatchUpdateFileTagShrinkRequest) SetUpdateMode(v string) *BatchUpdateFileTagShrinkRequest {
	s.UpdateMode = &v
	return s
}

func (s *BatchUpdateFileTagShrinkRequest) Validate() error {
	return dara.Validate(s)
}
