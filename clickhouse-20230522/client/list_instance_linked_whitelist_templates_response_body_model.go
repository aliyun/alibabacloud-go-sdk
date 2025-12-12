// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceLinkedWhitelistTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListInstanceLinkedWhitelistTemplatesResponseBodyData) *ListInstanceLinkedWhitelistTemplatesResponseBody
	GetData() *ListInstanceLinkedWhitelistTemplatesResponseBodyData
	SetRequestId(v string) *ListInstanceLinkedWhitelistTemplatesResponseBody
	GetRequestId() *string
}

type ListInstanceLinkedWhitelistTemplatesResponseBody struct {
	Data *ListInstanceLinkedWhitelistTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceLinkedWhitelistTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLinkedWhitelistTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBody) GetData() *ListInstanceLinkedWhitelistTemplatesResponseBodyData {
	return s.Data
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBody) SetData(v *ListInstanceLinkedWhitelistTemplatesResponseBodyData) *ListInstanceLinkedWhitelistTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBody) SetRequestId(v string) *ListInstanceLinkedWhitelistTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceLinkedWhitelistTemplatesResponseBodyData struct {
	Templates []*ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListInstanceLinkedWhitelistTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLinkedWhitelistTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyData) GetTemplates() []*ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates {
	return s.Templates
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyData) SetTemplates(v []*ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) *ListInstanceLinkedWhitelistTemplatesResponseBodyData {
	s.Templates = v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyData) Validate() error {
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

type ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates struct {
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

func (s ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) SetSecurityIPList(v string) *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates {
	s.SecurityIPList = &v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) SetTemplateId(v string) *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) SetTemplateName(v string) *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListInstanceLinkedWhitelistTemplatesResponseBodyDataTemplates) Validate() error {
	return dara.Validate(s)
}
