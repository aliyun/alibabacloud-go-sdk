// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudNotePhrasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateCloudNotePhrasesShrinkRequest
	GetAppId() *string
	SetPhraseShrink(v string) *CreateCloudNotePhrasesShrinkRequest
	GetPhraseShrink() *string
}

type CreateCloudNotePhrasesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	PhraseShrink *string `json:"Phrase,omitempty" xml:"Phrase,omitempty"`
}

func (s CreateCloudNotePhrasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudNotePhrasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudNotePhrasesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateCloudNotePhrasesShrinkRequest) GetPhraseShrink() *string {
	return s.PhraseShrink
}

func (s *CreateCloudNotePhrasesShrinkRequest) SetAppId(v string) *CreateCloudNotePhrasesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateCloudNotePhrasesShrinkRequest) SetPhraseShrink(v string) *CreateCloudNotePhrasesShrinkRequest {
	s.PhraseShrink = &v
	return s
}

func (s *CreateCloudNotePhrasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
