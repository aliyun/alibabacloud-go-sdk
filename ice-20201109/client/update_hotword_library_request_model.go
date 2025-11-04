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
	Hotwords []*Hotword `json:"Hotwords,omitempty" xml:"Hotwords,omitempty" type:"Repeated"`
	// The name of the hotword library. It can be up to 100 characters in length.
	//
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
	if s.Hotwords != nil {
		for _, item := range s.Hotwords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
