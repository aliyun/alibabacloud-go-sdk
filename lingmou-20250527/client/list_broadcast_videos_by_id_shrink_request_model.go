// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastVideosByIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoIdsShrink(v string) *ListBroadcastVideosByIdShrinkRequest
	GetVideoIdsShrink() *string
}

type ListBroadcastVideosByIdShrinkRequest struct {
	VideoIdsShrink *string `json:"videoIds,omitempty" xml:"videoIds,omitempty"`
}

func (s ListBroadcastVideosByIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastVideosByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListBroadcastVideosByIdShrinkRequest) GetVideoIdsShrink() *string {
	return s.VideoIdsShrink
}

func (s *ListBroadcastVideosByIdShrinkRequest) SetVideoIdsShrink(v string) *ListBroadcastVideosByIdShrinkRequest {
	s.VideoIdsShrink = &v
	return s
}

func (s *ListBroadcastVideosByIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
