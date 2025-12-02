// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdatePlaybookRequest
	GetLang() *string
	SetPlaybookDescription(v string) *UpdatePlaybookRequest
	GetPlaybookDescription() *string
	SetPlaybookInputConfigs(v []*FieldInputConfig) *UpdatePlaybookRequest
	GetPlaybookInputConfigs() []*FieldInputConfig
	SetPlaybookName(v string) *UpdatePlaybookRequest
	GetPlaybookName() *string
	SetPlaybookOutputConfigs(v []*FieldOutputConfig) *UpdatePlaybookRequest
	GetPlaybookOutputConfigs() []*FieldOutputConfig
	SetPlaybookParamType(v string) *UpdatePlaybookRequest
	GetPlaybookParamType() *string
	SetPlaybookTaskFlow(v string) *UpdatePlaybookRequest
	GetPlaybookTaskFlow() *string
	SetPlaybookUuid(v string) *UpdatePlaybookRequest
	GetPlaybookUuid() *string
	SetRoleFor(v int64) *UpdatePlaybookRequest
	GetRoleFor() *int64
}

type UpdatePlaybookRequest struct {
	// The language type for requests and received messages. Values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Description of the playbook.
	//
	// example:
	//
	// alb block
	PlaybookDescription *string `json:"PlaybookDescription,omitempty" xml:"PlaybookDescription,omitempty"`
	// List of input parameter configurations for the playbook.
	PlaybookInputConfigs []*FieldInputConfig `json:"PlaybookInputConfigs,omitempty" xml:"PlaybookInputConfigs,omitempty" type:"Repeated"`
	// The name of the playbook.
	//
	// example:
	//
	// system_aliyun_alb_process_book
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// List of output parameter configurations for the playbook.
	PlaybookOutputConfigs []*FieldOutputConfig `json:"PlaybookOutputConfigs,omitempty" xml:"PlaybookOutputConfigs,omitempty" type:"Repeated"`
	// Type of input parameters for the playbook.
	//
	// - **template-ip**: IP entity.
	//
	// - **template-file**: File entity.
	//
	// - **template-process**: Process entity.
	//
	// - **template-host**: Host entity.
	//
	// - **template-domain**: Domain entity.
	//
	// - **template-container**: Container entity.
	//
	// - **template-incident**: Security incident.
	//
	// - **template-alert**: Security alert.
	//
	// - **custom**: Custom.
	//
	// example:
	//
	// template-ip
	PlaybookParamType *string `json:"PlaybookParamType,omitempty" xml:"PlaybookParamType,omitempty"`
	// Content of the playbook.
	//
	// example:
	//
	// [{"id":"58d87b7d-28d9-4f0e-b135-4adc4f1a70e4","zIndex":1,"data":{"nodeType":"startEvent","appType":"basic","nodeName":"start","icon":"icon-circle","description":"start"},"position":{"x":-440,"y":-170}},{"id":"5293c3f9-e1c9-4a49-b0eb-635067dc67e8","zIndex":1,"data":{"nodeType":"sequenceFlow","appType":"basic","icon":"icon-upper-right-arrow","isRequired":true},"source":{"cell":"58d87b7d-28d9-4f0e-b135-4adc4f1a70e4"},"target":{"cell":"f9d6d1f5-b0cd-45b6-93d0-02cd4b2a6fa2"},"vertices":[]},{"id":"317dd1be-2d20-460e-977e-1fc936ffb583","zIndex":1,"data":{"nodeType":"endEvent","appType":"basic","nodeName":"end","icon":"icon-radio-off-full","description":"end"},"position":{"x":-140,"y":-170}},{"id":"f9d6d1f5-b0cd-45b6-93d0-02cd4b2a6fa2","zIndex":1,"data":{"isDebug":false,"nodeType":"action","appType":"component","nodeName":"data","valueData":{"outputFields":"[{\\"fieldName\\":\\"ip\\",\\"fieldValue\\":\\"${event.ip}\\"}]"},"icon":"https://img.alicdn.com/imgextra/i2/O1CN01NCKmL026m1z8o0DeN_!!6000000007703-2-tps-248-248.png","description":"","advance":{"inputParamMode":false,"onError":"stop_cur_flow","rspStatusType":3,"rspStatusThreshold":0},"componentName":"DataFormat","actionName":"formatdata"},"position":{"x":-340,"y":-185}},{"id":"1c7f0021-fb93-4478-b10f-af78dd5a69d6","zIndex":1,"data":{"nodeType":"sequenceFlow","appType":"basic","icon":"icon-upper-right-arrow","isRequired":true},"source":{"cell":"f9d6d1f5-b0cd-45b6-93d0-02cd4b2a6fa2"},"target":{"cell":"317dd1be-2d20-460e-977e-1fc936ffb583"},"vertices":[]}]
	PlaybookTaskFlow *string `json:"PlaybookTaskFlow,omitempty" xml:"PlaybookTaskFlow,omitempty"`
	// UUID of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8f55e76d-b5d5-4720-9cd7-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The user ID for the administrator to switch to another member\\"s perspective.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdatePlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePlaybookRequest) GoString() string {
	return s.String()
}

func (s *UpdatePlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdatePlaybookRequest) GetPlaybookDescription() *string {
	return s.PlaybookDescription
}

func (s *UpdatePlaybookRequest) GetPlaybookInputConfigs() []*FieldInputConfig {
	return s.PlaybookInputConfigs
}

func (s *UpdatePlaybookRequest) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *UpdatePlaybookRequest) GetPlaybookOutputConfigs() []*FieldOutputConfig {
	return s.PlaybookOutputConfigs
}

func (s *UpdatePlaybookRequest) GetPlaybookParamType() *string {
	return s.PlaybookParamType
}

func (s *UpdatePlaybookRequest) GetPlaybookTaskFlow() *string {
	return s.PlaybookTaskFlow
}

func (s *UpdatePlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *UpdatePlaybookRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdatePlaybookRequest) SetLang(v string) *UpdatePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *UpdatePlaybookRequest) SetPlaybookDescription(v string) *UpdatePlaybookRequest {
	s.PlaybookDescription = &v
	return s
}

func (s *UpdatePlaybookRequest) SetPlaybookInputConfigs(v []*FieldInputConfig) *UpdatePlaybookRequest {
	s.PlaybookInputConfigs = v
	return s
}

func (s *UpdatePlaybookRequest) SetPlaybookName(v string) *UpdatePlaybookRequest {
	s.PlaybookName = &v
	return s
}

func (s *UpdatePlaybookRequest) SetPlaybookOutputConfigs(v []*FieldOutputConfig) *UpdatePlaybookRequest {
	s.PlaybookOutputConfigs = v
	return s
}

func (s *UpdatePlaybookRequest) SetPlaybookParamType(v string) *UpdatePlaybookRequest {
	s.PlaybookParamType = &v
	return s
}

func (s *UpdatePlaybookRequest) SetPlaybookTaskFlow(v string) *UpdatePlaybookRequest {
	s.PlaybookTaskFlow = &v
	return s
}

func (s *UpdatePlaybookRequest) SetPlaybookUuid(v string) *UpdatePlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *UpdatePlaybookRequest) SetRoleFor(v int64) *UpdatePlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdatePlaybookRequest) Validate() error {
	if s.PlaybookInputConfigs != nil {
		for _, item := range s.PlaybookInputConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PlaybookOutputConfigs != nil {
		for _, item := range s.PlaybookOutputConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
