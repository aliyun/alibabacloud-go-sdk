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
	// - The list of tags to associate with the file. You can specify a maximum of 100 tags. The combined length of all tag values cannot exceed 700 characters.
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
