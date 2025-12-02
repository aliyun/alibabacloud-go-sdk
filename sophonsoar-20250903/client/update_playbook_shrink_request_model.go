// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePlaybookShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdatePlaybookShrinkRequest
	GetLang() *string
	SetPlaybookDescription(v string) *UpdatePlaybookShrinkRequest
	GetPlaybookDescription() *string
	SetPlaybookInputConfigsShrink(v string) *UpdatePlaybookShrinkRequest
	GetPlaybookInputConfigsShrink() *string
	SetPlaybookName(v string) *UpdatePlaybookShrinkRequest
	GetPlaybookName() *string
	SetPlaybookOutputConfigsShrink(v string) *UpdatePlaybookShrinkRequest
	GetPlaybookOutputConfigsShrink() *string
	SetPlaybookParamType(v string) *UpdatePlaybookShrinkRequest
	GetPlaybookParamType() *string
	SetPlaybookTaskFlow(v string) *UpdatePlaybookShrinkRequest
	GetPlaybookTaskFlow() *string
	SetPlaybookUuid(v string) *UpdatePlaybookShrinkRequest
	GetPlaybookUuid() *string
	SetRoleFor(v int64) *UpdatePlaybookShrinkRequest
	GetRoleFor() *int64
}

type UpdatePlaybookShrinkRequest struct {
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
	PlaybookInputConfigsShrink *string `json:"PlaybookInputConfigs,omitempty" xml:"PlaybookInputConfigs,omitempty"`
	// The name of the playbook.
	//
	// example:
	//
	// system_aliyun_alb_process_book
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// List of output parameter configurations for the playbook.
	PlaybookOutputConfigsShrink *string `json:"PlaybookOutputConfigs,omitempty" xml:"PlaybookOutputConfigs,omitempty"`
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

func (s UpdatePlaybookShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePlaybookShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePlaybookShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdatePlaybookShrinkRequest) GetPlaybookDescription() *string {
	return s.PlaybookDescription
}

func (s *UpdatePlaybookShrinkRequest) GetPlaybookInputConfigsShrink() *string {
	return s.PlaybookInputConfigsShrink
}

func (s *UpdatePlaybookShrinkRequest) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *UpdatePlaybookShrinkRequest) GetPlaybookOutputConfigsShrink() *string {
	return s.PlaybookOutputConfigsShrink
}

func (s *UpdatePlaybookShrinkRequest) GetPlaybookParamType() *string {
	return s.PlaybookParamType
}

func (s *UpdatePlaybookShrinkRequest) GetPlaybookTaskFlow() *string {
	return s.PlaybookTaskFlow
}

func (s *UpdatePlaybookShrinkRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *UpdatePlaybookShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdatePlaybookShrinkRequest) SetLang(v string) *UpdatePlaybookShrinkRequest {
	s.Lang = &v
	return s
}

func (s *UpdatePlaybookShrinkRequest) SetPlaybookDescription(v string) *UpdatePlaybookShrinkRequest {
	s.PlaybookDescription = &v
	return s
}

func (s *UpdatePlaybookShrinkRequest) SetPlaybookInputConfigsShrink(v string) *UpdatePlaybookShrinkRequest {
	s.PlaybookInputConfigsShrink = &v
	return s
}

func (s *UpdatePlaybookShrinkRequest) SetPlaybookName(v string) *UpdatePlaybookShrinkRequest {
	s.PlaybookName = &v
	return s
}

func (s *UpdatePlaybookShrinkRequest) SetPlaybookOutputConfigsShrink(v string) *UpdatePlaybookShrinkRequest {
	s.PlaybookOutputConfigsShrink = &v
	return s
}

func (s *UpdatePlaybookShrinkRequest) SetPlaybookParamType(v string) *UpdatePlaybookShrinkRequest {
	s.PlaybookParamType = &v
	return s
}

func (s *UpdatePlaybookShrinkRequest) SetPlaybookTaskFlow(v string) *UpdatePlaybookShrinkRequest {
	s.PlaybookTaskFlow = &v
	return s
}

func (s *UpdatePlaybookShrinkRequest) SetPlaybookUuid(v string) *UpdatePlaybookShrinkRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *UpdatePlaybookShrinkRequest) SetRoleFor(v int64) *UpdatePlaybookShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdatePlaybookShrinkRequest) Validate() error {
	return dara.Validate(s)
}
