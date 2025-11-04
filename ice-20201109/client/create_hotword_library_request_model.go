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
	// The description of the hotword library. It can be up to 200 characters in length.
	//
	// example:
	//
	// 存放名人的词库
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hotword list. You can add up to 300 hotword entries to a single library.
	//
	// This parameter is required.
	Hotwords []*Hotword `json:"Hotwords,omitempty" xml:"Hotwords,omitempty" type:"Repeated"`
	// The name of the hotword library. It can be up to 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_hotwords
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The usage scenario of the hotword library. Valid values:
	//
	// · ASR: Automatic Speech Recognition
	//
	// · StructuredMediaAssets: structured media analysis
	//
	// · VideoTranslation: Video translation.
	//
	// This field cannot be modified after the hotword library is created.
	//
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
