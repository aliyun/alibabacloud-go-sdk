// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotwordLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateHotwordLibraryRequest
	GetDescription() *string
	SetHotwordLibraryId(v string) *UpdateHotwordLibraryRequest
	GetHotwordLibraryId() *string
	SetHotwords(v []*Hotword) *UpdateHotwordLibraryRequest
	GetHotwords() []*Hotword
	SetName(v string) *UpdateHotwordLibraryRequest
	GetName() *string
}

type UpdateHotwordLibraryRequest struct {
	// example:
	//
	// 存放名人的词库
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *a0052ff71efbfd4e7e6c66*
	HotwordLibraryId *string    `json:"HotwordLibraryId,omitempty" xml:"HotwordLibraryId,omitempty"`
	Hotwords         []*Hotword `json:"Hotwords,omitempty" xml:"Hotwords,omitempty" type:"Repeated"`
	// example:
	//
	// my_hotwords
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateHotwordLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotwordLibraryRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotwordLibraryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateHotwordLibraryRequest) GetHotwordLibraryId() *string {
	return s.HotwordLibraryId
}

func (s *UpdateHotwordLibraryRequest) GetHotwords() []*Hotword {
	return s.Hotwords
}

func (s *UpdateHotwordLibraryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateHotwordLibraryRequest) SetDescription(v string) *UpdateHotwordLibraryRequest {
	s.Description = &v
	return s
}

func (s *UpdateHotwordLibraryRequest) SetHotwordLibraryId(v string) *UpdateHotwordLibraryRequest {
	s.HotwordLibraryId = &v
	return s
}

func (s *UpdateHotwordLibraryRequest) SetHotwords(v []*Hotword) *UpdateHotwordLibraryRequest {
	s.Hotwords = v
	return s
}

func (s *UpdateHotwordLibraryRequest) SetName(v string) *UpdateHotwordLibraryRequest {
	s.Name = &v
	return s
}

func (s *UpdateHotwordLibraryRequest) Validate() error {
	return dara.Validate(s)
}
