// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableIndustry(v *GetTextTemplateResponseBodyAvailableIndustry) *GetTextTemplateResponseBody
	GetAvailableIndustry() *GetTextTemplateResponseBodyAvailableIndustry
	SetRequestId(v string) *GetTextTemplateResponseBody
	GetRequestId() *string
}

type GetTextTemplateResponseBody struct {
	AvailableIndustry *GetTextTemplateResponseBodyAvailableIndustry `json:"availableIndustry,omitempty" xml:"availableIndustry,omitempty" type:"Struct"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTextTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTextTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponseBody) GetAvailableIndustry() *GetTextTemplateResponseBodyAvailableIndustry {
	return s.AvailableIndustry
}

func (s *GetTextTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTextTemplateResponseBody) SetAvailableIndustry(v *GetTextTemplateResponseBodyAvailableIndustry) *GetTextTemplateResponseBody {
	s.AvailableIndustry = v
	return s
}

func (s *GetTextTemplateResponseBody) SetRequestId(v string) *GetTextTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextTemplateResponseBody) Validate() error {
	if s.AvailableIndustry != nil {
		if err := s.AvailableIndustry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTextTemplateResponseBodyAvailableIndustry struct {
	// example:
	//
	// Car
	Name          *string                                                      `json:"name,omitempty" xml:"name,omitempty"`
	TextModeTypes []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypes `json:"textModeTypes,omitempty" xml:"textModeTypes,omitempty" type:"Repeated"`
}

func (s GetTextTemplateResponseBodyAvailableIndustry) String() string {
	return dara.Prettify(s)
}

func (s GetTextTemplateResponseBodyAvailableIndustry) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponseBodyAvailableIndustry) GetName() *string {
	return s.Name
}

func (s *GetTextTemplateResponseBodyAvailableIndustry) GetTextModeTypes() []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypes {
	return s.TextModeTypes
}

func (s *GetTextTemplateResponseBodyAvailableIndustry) SetName(v string) *GetTextTemplateResponseBodyAvailableIndustry {
	s.Name = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustry) SetTextModeTypes(v []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) *GetTextTemplateResponseBodyAvailableIndustry {
	s.TextModeTypes = v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustry) Validate() error {
	if s.TextModeTypes != nil {
		for _, item := range s.TextModeTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTextTemplateResponseBodyAvailableIndustryTextModeTypes struct {
	// example:
	//
	// Rewrite
	Name       *string                                                                `json:"name,omitempty" xml:"name,omitempty"`
	TextStyles []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles `json:"textStyles,omitempty" xml:"textStyles,omitempty" type:"Repeated"`
}

func (s GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) String() string {
	return dara.Prettify(s)
}

func (s GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) GetName() *string {
	return s.Name
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) GetTextStyles() []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles {
	return s.TextStyles
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) SetName(v string) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes {
	s.Name = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) SetTextStyles(v []*GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes {
	s.TextStyles = v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypes) Validate() error {
	if s.TextStyles != nil {
		for _, item := range s.TextStyles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// false
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// example:
	//
	// RED_BOOK
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 111
	TemplateKey *string `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
}

func (s GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) String() string {
	return dara.Prettify(s)
}

func (s GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) GoString() string {
	return s.String()
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) GetDesc() *string {
	return s.Desc
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) GetDisabled() *bool {
	return s.Disabled
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) GetName() *string {
	return s.Name
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) GetTemplateKey() *string {
	return s.TemplateKey
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) SetDesc(v string) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles {
	s.Desc = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) SetDisabled(v bool) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles {
	s.Disabled = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) SetName(v string) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles {
	s.Name = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) SetTemplateKey(v string) *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles {
	s.TemplateKey = &v
	return s
}

func (s *GetTextTemplateResponseBodyAvailableIndustryTextModeTypesTextStyles) Validate() error {
	return dara.Validate(s)
}
