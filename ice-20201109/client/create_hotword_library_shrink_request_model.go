// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotwordLibraryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateHotwordLibraryShrinkRequest
	GetDescription() *string
	SetHotwordsShrink(v string) *CreateHotwordLibraryShrinkRequest
	GetHotwordsShrink() *string
	SetName(v string) *CreateHotwordLibraryShrinkRequest
	GetName() *string
	SetUsageScenario(v string) *CreateHotwordLibraryShrinkRequest
	GetUsageScenario() *string
}

type CreateHotwordLibraryShrinkRequest struct {
	// example:
	//
	// 存放名人的词库
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	HotwordsShrink *string `json:"Hotwords,omitempty" xml:"Hotwords,omitempty"`
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

func (s CreateHotwordLibraryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHotwordLibraryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHotwordLibraryShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHotwordLibraryShrinkRequest) GetHotwordsShrink() *string {
	return s.HotwordsShrink
}

func (s *CreateHotwordLibraryShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateHotwordLibraryShrinkRequest) GetUsageScenario() *string {
	return s.UsageScenario
}

func (s *CreateHotwordLibraryShrinkRequest) SetDescription(v string) *CreateHotwordLibraryShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateHotwordLibraryShrinkRequest) SetHotwordsShrink(v string) *CreateHotwordLibraryShrinkRequest {
	s.HotwordsShrink = &v
	return s
}

func (s *CreateHotwordLibraryShrinkRequest) SetName(v string) *CreateHotwordLibraryShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateHotwordLibraryShrinkRequest) SetUsageScenario(v string) *CreateHotwordLibraryShrinkRequest {
	s.UsageScenario = &v
	return s
}

func (s *CreateHotwordLibraryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
