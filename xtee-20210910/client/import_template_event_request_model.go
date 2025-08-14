// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportTemplateEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ImportTemplateEventRequest
	GetLang() *string
	SetEventTemplateIds(v string) *ImportTemplateEventRequest
	GetEventTemplateIds() *string
	SetRegId(v string) *ImportTemplateEventRequest
	GetRegId() *string
}

type ImportTemplateEventRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The template ID of the event.
	//
	// example:
	//
	// register
	EventTemplateIds *string `json:"eventTemplateIds,omitempty" xml:"eventTemplateIds,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s ImportTemplateEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportTemplateEventRequest) GoString() string {
	return s.String()
}

func (s *ImportTemplateEventRequest) GetLang() *string {
	return s.Lang
}

func (s *ImportTemplateEventRequest) GetEventTemplateIds() *string {
	return s.EventTemplateIds
}

func (s *ImportTemplateEventRequest) GetRegId() *string {
	return s.RegId
}

func (s *ImportTemplateEventRequest) SetLang(v string) *ImportTemplateEventRequest {
	s.Lang = &v
	return s
}

func (s *ImportTemplateEventRequest) SetEventTemplateIds(v string) *ImportTemplateEventRequest {
	s.EventTemplateIds = &v
	return s
}

func (s *ImportTemplateEventRequest) SetRegId(v string) *ImportTemplateEventRequest {
	s.RegId = &v
	return s
}

func (s *ImportTemplateEventRequest) Validate() error {
	return dara.Validate(s)
}
