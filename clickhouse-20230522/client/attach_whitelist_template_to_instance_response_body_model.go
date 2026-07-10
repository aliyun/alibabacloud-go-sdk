// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachWhitelistTemplateToInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AttachWhitelistTemplateToInstanceResponseBodyData) *AttachWhitelistTemplateToInstanceResponseBody
	GetData() *AttachWhitelistTemplateToInstanceResponseBodyData
	SetRequestId(v string) *AttachWhitelistTemplateToInstanceResponseBody
	GetRequestId() *string
}

type AttachWhitelistTemplateToInstanceResponseBody struct {
	// The returned data.
	Data *AttachWhitelistTemplateToInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachWhitelistTemplateToInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) GetData() *AttachWhitelistTemplateToInstanceResponseBodyData {
	return s.Data
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) SetData(v *AttachWhitelistTemplateToInstanceResponseBodyData) *AttachWhitelistTemplateToInstanceResponseBody {
	s.Data = v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) SetRequestId(v string) *AttachWhitelistTemplateToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AttachWhitelistTemplateToInstanceResponseBodyData struct {
	// Instances that failed to be attached.
	AttachFailList []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList `json:"AttachFailList,omitempty" xml:"AttachFailList,omitempty" type:"Repeated"`
	// Instances to which the template was successfully attached.
	AttachSuccessedList []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList `json:"AttachSuccessedList,omitempty" xml:"AttachSuccessedList,omitempty" type:"Repeated"`
	// The status of the operation. A value of `ok` indicates success.
	//
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AttachWhitelistTemplateToInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) GetAttachFailList() []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList {
	return s.AttachFailList
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) GetAttachSuccessedList() []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList {
	return s.AttachSuccessedList
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) SetAttachFailList(v []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList) *AttachWhitelistTemplateToInstanceResponseBodyData {
	s.AttachFailList = v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) SetAttachSuccessedList(v []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList) *AttachWhitelistTemplateToInstanceResponseBodyData {
	s.AttachSuccessedList = v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) SetStatus(v string) *AttachWhitelistTemplateToInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyData) Validate() error {
	if s.AttachFailList != nil {
		for _, item := range s.AttachFailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AttachSuccessedList != nil {
		for _, item := range s.AttachSuccessedList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList struct {
	// The name of the instance.
	//
	// example:
	//
	// my-database
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The reason for the attachment failure.
	//
	// example:
	//
	// 处理异常
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList) GetReason() *string {
	return s.Reason
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList) SetDBInstanceId(v string) *AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList {
	s.DBInstanceId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList) SetReason(v string) *AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList {
	s.Reason = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachFailList) Validate() error {
	return dara.Validate(s)
}

type AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList struct {
	// The name of the instance.
	//
	// example:
	//
	// my-database
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The list of whitelist templates.
	Templates []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList) GetTemplates() []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates {
	return s.Templates
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList) SetDBInstanceId(v string) *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList {
	s.DBInstanceId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList) SetTemplates(v []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList {
	s.Templates = v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedList) Validate() error {
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

type AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates struct {
	// The list of attached instances.
	DbInstances []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Repeated"`
	// The IP address whitelist.
	//
	// example:
	//
	// 192.168.1.1,10.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The ID of the whitelist template.
	//
	// example:
	//
	// g-asdfwem
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the whitelist template.
	//
	// example:
	//
	// demo_template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) GetDbInstances() []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances {
	return s.DbInstances
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) SetDbInstances(v []*AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances) *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates {
	s.DbInstances = v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) SetSecurityIPList(v string) *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates {
	s.SecurityIPList = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) SetTemplateId(v string) *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates {
	s.TemplateId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) SetTemplateName(v string) *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates {
	s.TemplateName = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplates) Validate() error {
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

type AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances struct {
	// The ID of the instance.
	//
	// example:
	//
	// asdfwef
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances) SetDBInstanceId(v string) *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances {
	s.DBInstanceId = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponseBodyDataAttachSuccessedListTemplatesDbInstances) Validate() error {
	return dara.Validate(s)
}
