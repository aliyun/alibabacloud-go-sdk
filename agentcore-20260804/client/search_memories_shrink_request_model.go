// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoriesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *SearchMemoriesShrinkRequest
	GetBodyShrink() *string
}

type SearchMemoriesShrinkRequest struct {
	// The semantic retrieval request.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMemoriesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchMemoriesShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *SearchMemoriesShrinkRequest) SetBodyShrink(v string) *SearchMemoriesShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *SearchMemoriesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
