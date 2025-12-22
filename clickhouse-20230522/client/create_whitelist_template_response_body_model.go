// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhitelistTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateWhitelistTemplateResponseBodyData) *CreateWhitelistTemplateResponseBody
	GetData() *CreateWhitelistTemplateResponseBodyData
	SetRequestId(v string) *CreateWhitelistTemplateResponseBody
	GetRequestId() *string
}

type CreateWhitelistTemplateResponseBody struct {
	Data *CreateWhitelistTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWhitelistTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWhitelistTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWhitelistTemplateResponseBody) GetData() *CreateWhitelistTemplateResponseBodyData {
	return s.Data
}

func (s *CreateWhitelistTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWhitelistTemplateResponseBody) SetData(v *CreateWhitelistTemplateResponseBodyData) *CreateWhitelistTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateWhitelistTemplateResponseBody) SetRequestId(v string) *CreateWhitelistTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWhitelistTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWhitelistTemplateResponseBodyData struct {
	Templates []*CreateWhitelistTemplateResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s CreateWhitelistTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWhitelistTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWhitelistTemplateResponseBodyData) GetTemplates() []*CreateWhitelistTemplateResponseBodyDataTemplates {
	return s.Templates
}

func (s *CreateWhitelistTemplateResponseBodyData) SetTemplates(v []*CreateWhitelistTemplateResponseBodyDataTemplates) *CreateWhitelistTemplateResponseBodyData {
	s.Templates = v
	return s
}

func (s *CreateWhitelistTemplateResponseBodyData) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWhitelistTemplateResponseBodyDataTemplates struct {
	// example:
	//
	// 192.168.1.1,10.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// g-asdfwem
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// demo_template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateWhitelistTemplateResponseBodyDataTemplates) String() string {
	return dara.Prettify(s)
}

func (s CreateWhitelistTemplateResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *CreateWhitelistTemplateResponseBodyDataTemplates) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateWhitelistTemplateResponseBodyDataTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWhitelistTemplateResponseBodyDataTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateWhitelistTemplateResponseBodyDataTemplates) SetSecurityIPList(v string) *CreateWhitelistTemplateResponseBodyDataTemplates {
	s.SecurityIPList = &v
	return s
}

func (s *CreateWhitelistTemplateResponseBodyDataTemplates) SetTemplateId(v string) *CreateWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateId = &v
	return s
}

func (s *CreateWhitelistTemplateResponseBodyDataTemplates) SetTemplateName(v string) *CreateWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateName = &v
	return s
}

func (s *CreateWhitelistTemplateResponseBodyDataTemplates) Validate() error {
	return dara.Validate(s)
}
