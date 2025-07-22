// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudNotePhrasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteCloudNotePhrasesShrinkRequest
	GetAppId() *string
	SetPhraseShrink(v string) *DeleteCloudNotePhrasesShrinkRequest
	GetPhraseShrink() *string
}

type DeleteCloudNotePhrasesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	PhraseShrink *string `json:"Phrase,omitempty" xml:"Phrase,omitempty"`
}

func (s DeleteCloudNotePhrasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudNotePhrasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudNotePhrasesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteCloudNotePhrasesShrinkRequest) GetPhraseShrink() *string {
	return s.PhraseShrink
}

func (s *DeleteCloudNotePhrasesShrinkRequest) SetAppId(v string) *DeleteCloudNotePhrasesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteCloudNotePhrasesShrinkRequest) SetPhraseShrink(v string) *DeleteCloudNotePhrasesShrinkRequest {
	s.PhraseShrink = &v
	return s
}

func (s *DeleteCloudNotePhrasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
