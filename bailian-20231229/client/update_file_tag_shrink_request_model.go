// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileTagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagsShrink(v string) *UpdateFileTagShrinkRequest
	GetTagsShrink() *string
}

type UpdateFileTagShrinkRequest struct {
	// - The list of tags associated with the file. You can specify up to 100 tags, and the total character length of all tags cannot exceed 700.
	//
	// This parameter is required.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateFileTagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileTagShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateFileTagShrinkRequest) SetTagsShrink(v string) *UpdateFileTagShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateFileTagShrinkRequest) Validate() error {
	return dara.Validate(s)
}
