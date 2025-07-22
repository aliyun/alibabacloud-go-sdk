// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudNotePhrasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyCloudNotePhrasesShrinkRequest
	GetAppId() *string
	SetPhraseShrink(v string) *ModifyCloudNotePhrasesShrinkRequest
	GetPhraseShrink() *string
}

type ModifyCloudNotePhrasesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	PhraseShrink *string `json:"Phrase,omitempty" xml:"Phrase,omitempty"`
}

func (s ModifyCloudNotePhrasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudNotePhrasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudNotePhrasesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyCloudNotePhrasesShrinkRequest) GetPhraseShrink() *string {
	return s.PhraseShrink
}

func (s *ModifyCloudNotePhrasesShrinkRequest) SetAppId(v string) *ModifyCloudNotePhrasesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyCloudNotePhrasesShrinkRequest) SetPhraseShrink(v string) *ModifyCloudNotePhrasesShrinkRequest {
	s.PhraseShrink = &v
	return s
}

func (s *ModifyCloudNotePhrasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
