// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotwordLibraryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateHotwordLibraryShrinkRequest
	GetDescription() *string
	SetHotwordLibraryId(v string) *UpdateHotwordLibraryShrinkRequest
	GetHotwordLibraryId() *string
	SetHotwordsShrink(v string) *UpdateHotwordLibraryShrinkRequest
	GetHotwordsShrink() *string
	SetName(v string) *UpdateHotwordLibraryShrinkRequest
	GetName() *string
}

type UpdateHotwordLibraryShrinkRequest struct {
	// The description of the hotword library. It can be up to 200 characters in length.
	//
	// example:
	//
	// 存放名人的词库
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the hotword library.
	//
	// This parameter is required.
	//
	// example:
	//
	// *a0052ff71efbfd4e7e6c66*
	HotwordLibraryId *string `json:"HotwordLibraryId,omitempty" xml:"HotwordLibraryId,omitempty"`
	// The hotword list. You can add up to 300 hotword entries to a single library.
	HotwordsShrink *string `json:"Hotwords,omitempty" xml:"Hotwords,omitempty"`
	// The name of the hotword library. It can be up to 100 characters in length.
	//
	// example:
	//
	// my_hotwords
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateHotwordLibraryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotwordLibraryShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotwordLibraryShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateHotwordLibraryShrinkRequest) GetHotwordLibraryId() *string {
	return s.HotwordLibraryId
}

func (s *UpdateHotwordLibraryShrinkRequest) GetHotwordsShrink() *string {
	return s.HotwordsShrink
}

func (s *UpdateHotwordLibraryShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateHotwordLibraryShrinkRequest) SetDescription(v string) *UpdateHotwordLibraryShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateHotwordLibraryShrinkRequest) SetHotwordLibraryId(v string) *UpdateHotwordLibraryShrinkRequest {
	s.HotwordLibraryId = &v
	return s
}

func (s *UpdateHotwordLibraryShrinkRequest) SetHotwordsShrink(v string) *UpdateHotwordLibraryShrinkRequest {
	s.HotwordsShrink = &v
	return s
}

func (s *UpdateHotwordLibraryShrinkRequest) SetName(v string) *UpdateHotwordLibraryShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateHotwordLibraryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
