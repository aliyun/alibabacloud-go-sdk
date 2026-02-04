// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastAudiosByIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioIdsShrink(v string) *ListBroadcastAudiosByIdShrinkRequest
	GetAudioIdsShrink() *string
}

type ListBroadcastAudiosByIdShrinkRequest struct {
	AudioIdsShrink *string `json:"audioIds,omitempty" xml:"audioIds,omitempty"`
}

func (s ListBroadcastAudiosByIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastAudiosByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListBroadcastAudiosByIdShrinkRequest) GetAudioIdsShrink() *string {
	return s.AudioIdsShrink
}

func (s *ListBroadcastAudiosByIdShrinkRequest) SetAudioIdsShrink(v string) *ListBroadcastAudiosByIdShrinkRequest {
	s.AudioIdsShrink = &v
	return s
}

func (s *ListBroadcastAudiosByIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
