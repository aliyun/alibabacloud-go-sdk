// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotwordLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateHotwordLibraryRequest
	GetDescription() *string
	SetHotwords(v []*Hotword) *CreateHotwordLibraryRequest
	GetHotwords() []*Hotword
	SetName(v string) *CreateHotwordLibraryRequest
	GetName() *string
	SetUsageScenario(v string) *CreateHotwordLibraryRequest
	GetUsageScenario() *string
}

type CreateHotwordLibraryRequest struct {
	// example:
	//
	// 存放名人的词库
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Hotwords []*Hotword `json:"Hotwords,omitempty" xml:"Hotwords,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// my_hotwords
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ASR
	UsageScenario *string `json:"UsageScenario,omitempty" xml:"UsageScenario,omitempty"`
}

func (s CreateHotwordLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHotwordLibraryRequest) GoString() string {
	return s.String()
}

func (s *CreateHotwordLibraryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHotwordLibraryRequest) GetHotwords() []*Hotword {
	return s.Hotwords
}

func (s *CreateHotwordLibraryRequest) GetName() *string {
	return s.Name
}

func (s *CreateHotwordLibraryRequest) GetUsageScenario() *string {
	return s.UsageScenario
}

func (s *CreateHotwordLibraryRequest) SetDescription(v string) *CreateHotwordLibraryRequest {
	s.Description = &v
	return s
}

func (s *CreateHotwordLibraryRequest) SetHotwords(v []*Hotword) *CreateHotwordLibraryRequest {
	s.Hotwords = v
	return s
}

func (s *CreateHotwordLibraryRequest) SetName(v string) *CreateHotwordLibraryRequest {
	s.Name = &v
	return s
}

func (s *CreateHotwordLibraryRequest) SetUsageScenario(v string) *CreateHotwordLibraryRequest {
	s.UsageScenario = &v
	return s
}

func (s *CreateHotwordLibraryRequest) Validate() error {
	return dara.Validate(s)
}
