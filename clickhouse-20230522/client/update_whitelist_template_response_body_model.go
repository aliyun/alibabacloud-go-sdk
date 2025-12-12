// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhitelistTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateWhitelistTemplateResponseBodyData) *UpdateWhitelistTemplateResponseBody
	GetData() *UpdateWhitelistTemplateResponseBodyData
	SetRequestId(v string) *UpdateWhitelistTemplateResponseBody
	GetRequestId() *string
}

type UpdateWhitelistTemplateResponseBody struct {
	Data *UpdateWhitelistTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWhitelistTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhitelistTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWhitelistTemplateResponseBody) GetData() *UpdateWhitelistTemplateResponseBodyData {
	return s.Data
}

func (s *UpdateWhitelistTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWhitelistTemplateResponseBody) SetData(v *UpdateWhitelistTemplateResponseBodyData) *UpdateWhitelistTemplateResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWhitelistTemplateResponseBody) SetRequestId(v string) *UpdateWhitelistTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWhitelistTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWhitelistTemplateResponseBodyData struct {
	Templates []*UpdateWhitelistTemplateResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s UpdateWhitelistTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhitelistTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateWhitelistTemplateResponseBodyData) GetTemplates() []*UpdateWhitelistTemplateResponseBodyDataTemplates {
	return s.Templates
}

func (s *UpdateWhitelistTemplateResponseBodyData) SetTemplates(v []*UpdateWhitelistTemplateResponseBodyDataTemplates) *UpdateWhitelistTemplateResponseBodyData {
	s.Templates = v
	return s
}

func (s *UpdateWhitelistTemplateResponseBodyData) Validate() error {
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

type UpdateWhitelistTemplateResponseBodyDataTemplates struct {
	DbInstances []*UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Repeated"`
	// example:
	//
	// 192.168.1.1,10.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// g-asdfwem
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateWhitelistTemplateResponseBodyDataTemplates) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhitelistTemplateResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplates) GetDbInstances() []*UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances {
	return s.DbInstances
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplates) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplates) SetDbInstances(v []*UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances) *UpdateWhitelistTemplateResponseBodyDataTemplates {
	s.DbInstances = v
	return s
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplates) SetSecurityIPList(v string) *UpdateWhitelistTemplateResponseBodyDataTemplates {
	s.SecurityIPList = &v
	return s
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplates) SetTemplateId(v string) *UpdateWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateId = &v
	return s
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplates) Validate() error {
	if s.DbInstances != nil {
		for _, item := range s.DbInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances struct {
	// example:
	//
	// asdfwef
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
}

func (s UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances) GoString() string {
	return s.String()
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances) SetDbInstanceName(v string) *UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances {
	s.DbInstanceName = &v
	return s
}

func (s *UpdateWhitelistTemplateResponseBodyDataTemplatesDbInstances) Validate() error {
	return dara.Validate(s)
}
