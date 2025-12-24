// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ConvertPlaybookRequest
	GetLang() *string
	SetRoleFor(v int64) *ConvertPlaybookRequest
	GetRoleFor() *int64
	SetRoleType(v string) *ConvertPlaybookRequest
	GetRoleType() *string
	SetTaskflow(v string) *ConvertPlaybookRequest
	GetTaskflow() *string
}

type ConvertPlaybookRequest struct {
	// Language type for request and response messages. Values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// User ID for the administrator to switch to another member\\"s perspective.
	//
	// example:
	//
	// 13760*****718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// View type. Values:
	//
	// - 0: Current Alibaba Cloud account view.
	//
	// - 1: View for all accounts under the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// XML configuration information for playbook orchestration.
	//
	// This parameter is required.
	//
	// example:
	//
	// <?xml version=\\"1.0\\" encoding=\\"UTF-8\\"?>
	//
	// <bpmn:definitions xmlns:xsi=\\"http://www.w3.org/2001/XMLSchema-instance\\" xmlns:bpmn=\\"http://www.omg.org/spec/BPMN/20100524/MODEL\\" xmlns:bpmndi=\\"http://www.omg.org/spec/BPMN/20100524/DI\\" xmlns:dc=\\"http://www.omg.org/spec/DD/20100524/DC\\" id=\\"Definitions_1\\" targetNamespace=\\"http://bpmn.io/schema/bpmn\\">
	//
	//   <bpmn:process id=\\"Process_1\\" isExecutable=\\"false\\">
	//
	//     <bpmn:startEvent id=\\"StartEvent_1\\" />
	//
	//   </bpmn:process>
	//
	//   <bpmndi:BPMNDiagram id=\\"BPMNDiagram_1\\">
	//
	//      <bpmndi:BPMNPlane id=\\"BPMNPlane_1\\" bpmnElement=\\"Process_1\\">
	//
	//            <bpmndi:BPMNShape id=\\"_BPMNShape_StartEvent_2\\" bpmnElement=\\"StartEvent_1\\">
	//
	//                    <dc:Bounds x=\\"173\\" y=\\"102\\" width=\\"36\\" height=\\"36\\" />
	//
	//             </bpmndi:BPMNShape>
	//
	//      </bpmndi:BPMNPlane>
	//
	//   </bpmndi:BPMNDiagram>
	//
	// </bpmn:definitions>
	Taskflow *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
}

func (s ConvertPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertPlaybookRequest) GoString() string {
	return s.String()
}

func (s *ConvertPlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *ConvertPlaybookRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ConvertPlaybookRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *ConvertPlaybookRequest) GetTaskflow() *string {
	return s.Taskflow
}

func (s *ConvertPlaybookRequest) SetLang(v string) *ConvertPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *ConvertPlaybookRequest) SetRoleFor(v int64) *ConvertPlaybookRequest {
	s.RoleFor = &v
	return s
}

func (s *ConvertPlaybookRequest) SetRoleType(v string) *ConvertPlaybookRequest {
	s.RoleType = &v
	return s
}

func (s *ConvertPlaybookRequest) SetTaskflow(v string) *ConvertPlaybookRequest {
	s.Taskflow = &v
	return s
}

func (s *ConvertPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
