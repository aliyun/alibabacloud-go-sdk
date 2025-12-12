// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhitelistTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetWhitelistTemplateResponseBodyData) *GetWhitelistTemplateResponseBody
	GetData() *GetWhitelistTemplateResponseBodyData
	SetRequestId(v string) *GetWhitelistTemplateResponseBody
	GetRequestId() *string
}

type GetWhitelistTemplateResponseBody struct {
	Data *GetWhitelistTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWhitelistTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWhitelistTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetWhitelistTemplateResponseBody) GetData() *GetWhitelistTemplateResponseBodyData {
	return s.Data
}

func (s *GetWhitelistTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWhitelistTemplateResponseBody) SetData(v *GetWhitelistTemplateResponseBodyData) *GetWhitelistTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetWhitelistTemplateResponseBody) SetRequestId(v string) *GetWhitelistTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWhitelistTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWhitelistTemplateResponseBodyData struct {
	Templates []*GetWhitelistTemplateResponseBodyDataTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s GetWhitelistTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWhitelistTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWhitelistTemplateResponseBodyData) GetTemplates() []*GetWhitelistTemplateResponseBodyDataTemplates {
	return s.Templates
}

func (s *GetWhitelistTemplateResponseBodyData) SetTemplates(v []*GetWhitelistTemplateResponseBodyDataTemplates) *GetWhitelistTemplateResponseBodyData {
	s.Templates = v
	return s
}

func (s *GetWhitelistTemplateResponseBodyData) Validate() error {
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

type GetWhitelistTemplateResponseBodyDataTemplates struct {
	DbInstances []*GetWhitelistTemplateResponseBodyDataTemplatesDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Repeated"`
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

func (s GetWhitelistTemplateResponseBodyDataTemplates) String() string {
	return dara.Prettify(s)
}

func (s GetWhitelistTemplateResponseBodyDataTemplates) GoString() string {
	return s.String()
}

func (s *GetWhitelistTemplateResponseBodyDataTemplates) GetDbInstances() []*GetWhitelistTemplateResponseBodyDataTemplatesDbInstances {
	return s.DbInstances
}

func (s *GetWhitelistTemplateResponseBodyDataTemplates) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *GetWhitelistTemplateResponseBodyDataTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetWhitelistTemplateResponseBodyDataTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetWhitelistTemplateResponseBodyDataTemplates) SetDbInstances(v []*GetWhitelistTemplateResponseBodyDataTemplatesDbInstances) *GetWhitelistTemplateResponseBodyDataTemplates {
	s.DbInstances = v
	return s
}

func (s *GetWhitelistTemplateResponseBodyDataTemplates) SetSecurityIPList(v string) *GetWhitelistTemplateResponseBodyDataTemplates {
	s.SecurityIPList = &v
	return s
}

func (s *GetWhitelistTemplateResponseBodyDataTemplates) SetTemplateId(v string) *GetWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateId = &v
	return s
}

func (s *GetWhitelistTemplateResponseBodyDataTemplates) SetTemplateName(v string) *GetWhitelistTemplateResponseBodyDataTemplates {
	s.TemplateName = &v
	return s
}

func (s *GetWhitelistTemplateResponseBodyDataTemplates) Validate() error {
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

type GetWhitelistTemplateResponseBodyDataTemplatesDbInstances struct {
	// example:
	//
	// cc-asdfwef
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s GetWhitelistTemplateResponseBodyDataTemplatesDbInstances) String() string {
	return dara.Prettify(s)
}

func (s GetWhitelistTemplateResponseBodyDataTemplatesDbInstances) GoString() string {
	return s.String()
}

func (s *GetWhitelistTemplateResponseBodyDataTemplatesDbInstances) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetWhitelistTemplateResponseBodyDataTemplatesDbInstances) SetDBInstanceId(v string) *GetWhitelistTemplateResponseBodyDataTemplatesDbInstances {
	s.DBInstanceId = &v
	return s
}

func (s *GetWhitelistTemplateResponseBodyDataTemplatesDbInstances) Validate() error {
	return dara.Validate(s)
}
