// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreatePlaybookRequest
	GetLang() *string
	SetPlaybookDescription(v string) *CreatePlaybookRequest
	GetPlaybookDescription() *string
	SetPlaybookInputConfigs(v []*FieldInputConfig) *CreatePlaybookRequest
	GetPlaybookInputConfigs() []*FieldInputConfig
	SetPlaybookName(v string) *CreatePlaybookRequest
	GetPlaybookName() *string
	SetPlaybookOutputConfigs(v []*FieldOutputConfig) *CreatePlaybookRequest
	GetPlaybookOutputConfigs() []*FieldOutputConfig
	SetPlaybookParamType(v string) *CreatePlaybookRequest
	GetPlaybookParamType() *string
	SetPlaybookTaskFlow(v string) *CreatePlaybookRequest
	GetPlaybookTaskFlow() *string
	SetRoleFor(v int64) *CreatePlaybookRequest
	GetRoleFor() *int64
	SetSrcPlaybookUuid(v string) *CreatePlaybookRequest
	GetSrcPlaybookUuid() *string
}

type CreatePlaybookRequest struct {
	// Language type for receiving messages. Values:
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
	// waf ip block
	PlaybookDescription *string `json:"PlaybookDescription,omitempty" xml:"PlaybookDescription,omitempty"`
	// Input parameter configuration of the playbook.
	PlaybookInputConfigs []*FieldInputConfig `json:"PlaybookInputConfigs,omitempty" xml:"PlaybookInputConfigs,omitempty" type:"Repeated"`
	// Name of the playbook, without special characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// waftest
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// Output parameter configuration of the playbook.
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
	// custom
	PlaybookParamType *string `json:"PlaybookParamType,omitempty" xml:"PlaybookParamType,omitempty"`
	// Workflow of the playbook.
	//
	// example:
	//
	// [{"id":"58d87b7d-28d9-4f0e-b135-4adc4f1a70e4","zIndex":1,"data":{"nodeType":"startEvent","appType":"basic","nodeName":"start","icon":"icon-circle","description":"start"},"position":{"x":-440,"y":-170}},{"id":"5293c3f9-e1c9-4a49-b0eb-635067dc67e8","zIndex":1,"data":{"nodeType":"sequenceFlow","appType":"basic","icon":"icon-upper-right-arrow","isRequired":true},"source":{"cell":"58d87b7d-28d9-4f0e-b135-4adc4f1a70e4"},"target":{"cell":"f9d6d1f5-b0cd-45b6-93d0-02cd4b2a6fa2"},"vertices":[]},{"id":"317dd1be-2d20-460e-977e-1fc936ffb583","zIndex":1,"data":{"nodeType":"endEvent","appType":"basic","nodeName":"end","icon":"icon-radio-off-full","description":"end"},"position":{"x":-140,"y":-170}},{"id":"f9d6d1f5-b0cd-45b6-93d0-02cd4b2a6fa2","zIndex":1,"data":{"isDebug":false,"nodeType":"action","appType":"component","nodeName":"data","valueData":{"outputFields":"[{\\"fieldName\\":\\"ip\\",\\"fieldValue\\":\\"${event.ip}\\"}]"},"icon":"https://img.alicdn.com/imgextra/i2/O1CN01NCKmL026m1z8o0DeN_!!6000000007703-2-tps-248-248.png","description":"","advance":{"inputParamMode":false,"onError":"stop_cur_flow","rspStatusType":3,"rspStatusThreshold":0},"componentName":"DataFormat","actionName":"formatdata"},"position":{"x":-340,"y":-185}},{"id":"1c7f0021-fb93-4478-b10f-af78dd5a69d6","zIndex":1,"data":{"nodeType":"sequenceFlow","appType":"basic","icon":"icon-upper-right-arrow","isRequired":true},"source":{"cell":"f9d6d1f5-b0cd-45b6-93d0-02cd4b2a6fa2"},"target":{"cell":"317dd1be-2d20-460e-977e-1fc936ffb583"},"vertices":[]}]
	PlaybookTaskFlow *string `json:"PlaybookTaskFlow,omitempty" xml:"PlaybookTaskFlow,omitempty"`
	// Resource directory member account ID.
	//
	// example:
	//
	// 170*********3093
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// In a copy scenario, the UUID of the source playbook needs to be filled in. When this parameter has a value, all other parameters except the playbook name and description are invalid.
	//
	// example:
	//
	// 1B5A9144-****-****-A466-****9D64AA99
	SrcPlaybookUuid *string `json:"SrcPlaybookUuid,omitempty" xml:"SrcPlaybookUuid,omitempty"`
}

func (s CreatePlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePlaybookRequest) GoString() string {
	return s.String()
}

func (s *CreatePlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *CreatePlaybookRequest) GetPlaybookDescription() *string {
	return s.PlaybookDescription
}

func (s *CreatePlaybookRequest) GetPlaybookInputConfigs() []*FieldInputConfig {
	return s.PlaybookInputConfigs
}

func (s *CreatePlaybookRequest) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *CreatePlaybookRequest) GetPlaybookOutputConfigs() []*FieldOutputConfig {
	return s.PlaybookOutputConfigs
}

func (s *CreatePlaybookRequest) GetPlaybookParamType() *string {
	return s.PlaybookParamType
}

func (s *CreatePlaybookRequest) GetPlaybookTaskFlow() *string {
	return s.PlaybookTaskFlow
}

func (s *CreatePlaybookRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreatePlaybookRequest) GetSrcPlaybookUuid() *string {
	return s.SrcPlaybookUuid
}

func (s *CreatePlaybookRequest) SetLang(v string) *CreatePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *CreatePlaybookRequest) SetPlaybookDescription(v string) *CreatePlaybookRequest {
	s.PlaybookDescription = &v
	return s
}

func (s *CreatePlaybookRequest) SetPlaybookInputConfigs(v []*FieldInputConfig) *CreatePlaybookRequest {
	s.PlaybookInputConfigs = v
	return s
}

func (s *CreatePlaybookRequest) SetPlaybookName(v string) *CreatePlaybookRequest {
	s.PlaybookName = &v
	return s
}

func (s *CreatePlaybookRequest) SetPlaybookOutputConfigs(v []*FieldOutputConfig) *CreatePlaybookRequest {
	s.PlaybookOutputConfigs = v
	return s
}

func (s *CreatePlaybookRequest) SetPlaybookParamType(v string) *CreatePlaybookRequest {
	s.PlaybookParamType = &v
	return s
}

func (s *CreatePlaybookRequest) SetPlaybookTaskFlow(v string) *CreatePlaybookRequest {
	s.PlaybookTaskFlow = &v
	return s
}

func (s *CreatePlaybookRequest) SetRoleFor(v int64) *CreatePlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *CreatePlaybookRequest) SetSrcPlaybookUuid(v string) *CreatePlaybookRequest {
	s.SrcPlaybookUuid = &v
	return s
}

func (s *CreatePlaybookRequest) Validate() error {
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
