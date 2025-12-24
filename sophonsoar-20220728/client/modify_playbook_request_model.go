// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyPlaybookRequest
	GetDescription() *string
	SetDisplayName(v string) *ModifyPlaybookRequest
	GetDisplayName() *string
	SetLang(v string) *ModifyPlaybookRequest
	GetLang() *string
	SetPlaybookUuid(v string) *ModifyPlaybookRequest
	GetPlaybookUuid() *string
	SetTaskflow(v string) *ModifyPlaybookRequest
	GetTaskflow() *string
}

type ModifyPlaybookRequest struct {
	// The description of the playbook.
	//
	// example:
	//
	// demo test task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun_waf_test_playbook
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-1586c35e61f8
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The XML configuration of the playbook.
	//
	// example:
	//
	// <?xml version="1.0" encoding="UTF-8"?><bpmn:definitions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" targetNamespace="http://bpmn.io/schema/bpmn" id="Definitions_1"><bpmn:process id="Process_1" isExecutable="false"><bpmn:startEvent id="StartEvent_1"/></bpmn:process><bpmndi:BPMNDiagram id="BPMNDiagram_1"><bpmndi:BPMNPlane id="BPMNPlane_1" bpmnElement="Process_1"><bpmndi:BPMNShape id="_BPMNShape_StartEvent_2" bpmnElement="StartEvent_1"><dc:Bounds height="36.0" width="36.0" x="173.0" y="102.0"/></bpmndi:BPMNShape></bpmndi:BPMNPlane></bpmndi:BPMNDiagram></bpmn:definitions>
	Taskflow *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
}

func (s ModifyPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyPlaybookRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ModifyPlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyPlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ModifyPlaybookRequest) GetTaskflow() *string {
	return s.Taskflow
}

func (s *ModifyPlaybookRequest) SetDescription(v string) *ModifyPlaybookRequest {
	s.Description = &v
	return s
}

func (s *ModifyPlaybookRequest) SetDisplayName(v string) *ModifyPlaybookRequest {
	s.DisplayName = &v
	return s
}

func (s *ModifyPlaybookRequest) SetLang(v string) *ModifyPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPlaybookRequest) SetPlaybookUuid(v string) *ModifyPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ModifyPlaybookRequest) SetTaskflow(v string) *ModifyPlaybookRequest {
	s.Taskflow = &v
	return s
}

func (s *ModifyPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
