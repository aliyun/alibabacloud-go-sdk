// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhitelistTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteWhitelistTemplateResponseBodyData) *DeleteWhitelistTemplateResponseBody
	GetData() *DeleteWhitelistTemplateResponseBodyData
	SetRequestId(v string) *DeleteWhitelistTemplateResponseBody
	GetRequestId() *string
}

type DeleteWhitelistTemplateResponseBody struct {
	Data *DeleteWhitelistTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWhitelistTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhitelistTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistTemplateResponseBody) GetData() *DeleteWhitelistTemplateResponseBodyData {
	return s.Data
}

func (s *DeleteWhitelistTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWhitelistTemplateResponseBody) SetData(v *DeleteWhitelistTemplateResponseBodyData) *DeleteWhitelistTemplateResponseBody {
	s.Data = v
	return s
}

func (s *DeleteWhitelistTemplateResponseBody) SetRequestId(v string) *DeleteWhitelistTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWhitelistTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteWhitelistTemplateResponseBodyData struct {
	Templates []*DeleteWhitelistTemplateResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DeleteWhitelistTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhitelistTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistTemplateResponseBodyData) GetTemplates() []*DeleteWhitelistTemplateResponseBodyDataTemplates {
	return s.Templates
}

func (s *DeleteWhitelistTemplateResponseBodyData) SetTemplates(v []*DeleteWhitelistTemplateResponseBodyDataTemplates) *DeleteWhitelistTemplateResponseBodyData {
	s.Templates = v
	return s
}

func (s *DeleteWhitelistTemplateResponseBodyData) Validate() error {
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

type DeleteWhitelistTemplateResponseBodyDataTemplates struct {
	DbInstances []*DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Repeated"`
	// example:
	//
	// 192.168.1.1,10.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// g-asdfwem
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteWhitelistTemplateResponseBodyDataTemplates) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhitelistTemplateResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplates) GetDbInstances() []*DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances {
	return s.DbInstances
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplates) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplates) SetDbInstances(v []*DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances) *DeleteWhitelistTemplateResponseBodyDataTemplates {
	s.DbInstances = v
	return s
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplates) SetSecurityIPList(v string) *DeleteWhitelistTemplateResponseBodyDataTemplates {
	s.SecurityIPList = &v
	return s
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplates) SetTemplateId(v string) *DeleteWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateId = &v
	return s
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplates) Validate() error {
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

type DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances struct {
	// example:
	//
	// asdfwef
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
}

func (s DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances) SetDbInstanceName(v string) *DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances {
	s.DbInstanceName = &v
	return s
}

func (s *DeleteWhitelistTemplateResponseBodyDataTemplatesDbInstances) Validate() error {
	return dara.Validate(s)
}
