// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachWhitelistTemplateToInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetachWhitelistTemplateToInstanceResponseBodyData) *DetachWhitelistTemplateToInstanceResponseBody
	GetData() *DetachWhitelistTemplateToInstanceResponseBodyData
	SetRequestId(v string) *DetachWhitelistTemplateToInstanceResponseBody
	GetRequestId() *string
}

type DetachWhitelistTemplateToInstanceResponseBody struct {
	Data *DetachWhitelistTemplateToInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachWhitelistTemplateToInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) GetData() *DetachWhitelistTemplateToInstanceResponseBodyData {
	return s.Data
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) SetData(v *DetachWhitelistTemplateToInstanceResponseBodyData) *DetachWhitelistTemplateToInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) SetRequestId(v string) *DetachWhitelistTemplateToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachWhitelistTemplateToInstanceResponseBodyData struct {
	DetachFailList      []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList      `json:"DetachFailList,omitempty" xml:"DetachFailList,omitempty" type:"Repeated"`
	DetachSuccessedList []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList `json:"DetachSuccessedList,omitempty" xml:"DetachSuccessedList,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DetachWhitelistTemplateToInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) GetDetachFailList() []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList {
	return s.DetachFailList
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) GetDetachSuccessedList() []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList {
	return s.DetachSuccessedList
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) SetDetachFailList(v []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList) *DetachWhitelistTemplateToInstanceResponseBodyData {
	s.DetachFailList = v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) SetDetachSuccessedList(v []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList) *DetachWhitelistTemplateToInstanceResponseBodyData {
	s.DetachSuccessedList = v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) SetStatus(v string) *DetachWhitelistTemplateToInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyData) Validate() error {
	if s.DetachFailList != nil {
		for _, item := range s.DetachFailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DetachSuccessedList != nil {
		for _, item := range s.DetachSuccessedList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList struct {
	// example:
	//
	// my-database
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 处理异常
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList) GetReason() *string {
	return s.Reason
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList) SetDBInstanceId(v string) *DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList {
	s.DBInstanceId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList) SetReason(v string) *DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList {
	s.Reason = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachFailList) Validate() error {
	return dara.Validate(s)
}

type DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList struct {
	// example:
	//
	// my-database
	DBInstanceId *string                                                                          `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Templates    []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList) GetTemplates() []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates {
	return s.Templates
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList) SetDBInstanceId(v string) *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList {
	s.DBInstanceId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList) SetTemplates(v []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList {
	s.Templates = v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedList) Validate() error {
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

type DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates struct {
	DbInstances []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Repeated"`
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

func (s DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) GetDbInstances() []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances {
	return s.DbInstances
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) SetDbInstances(v []*DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances) *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates {
	s.DbInstances = v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) SetSecurityIPList(v string) *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates {
	s.SecurityIPList = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) SetTemplateId(v string) *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates {
	s.TemplateId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) SetTemplateName(v string) *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates {
	s.TemplateName = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplates) Validate() error {
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

type DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances struct {
	// example:
	//
	// asdfwef
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances) String() string {
	return dara.Prettify(s)
}

func (s DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances) GoString() string {
	return s.String()
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances) SetDBInstanceId(v string) *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances {
	s.DBInstanceId = &v
	return s
}

func (s *DetachWhitelistTemplateToInstanceResponseBodyDataDetachSuccessedListTemplatesDbInstances) Validate() error {
	return dara.Validate(s)
}
