// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlaybookUuid(v string) *VerifyPlaybookRequest
	GetPlaybookUuid() *string
	SetTaskFlow(v string) *VerifyPlaybookRequest
	GetTaskFlow() *string
}

type VerifyPlaybookRequest struct {
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the playbook UUID.
	//
	// example:
	//
	// 9fcd3829-80ff-4681-be1e-4d2662c35fed
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The XML configuration of the playbook.
	//
	// example:
	//
	// <?xml version="1.0" encoding="UTF-8"?><bpmn:definitions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" targetNamespace="http://bpmn.io/schema/bpmn" id="Definitions_1"><bpmn:process id="Process_1" isExecutable="false"><bpmn:startEvent id="StartEvent_1"/></bpmn:process><bpmndi:BPMNDiagram id="BPMNDiagram_1"><bpmndi:BPMNPlane id="BPMNPlane_1" bpmnElement="Process_1"><bpmndi:BPMNShape id="_BPMNShape_StartEvent_2" bpmnElement="StartEvent_1"><dc:Bounds height="36.0" width="36.0" x="173.0" y="102.0"/></bpmndi:BPMNShape></bpmndi:BPMNPlane></bpmndi:BPMNDiagram></bpmn:definitions>
	TaskFlow *string `json:"TaskFlow,omitempty" xml:"TaskFlow,omitempty"`
}

func (s VerifyPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *VerifyPlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *VerifyPlaybookRequest) GetTaskFlow() *string {
	return s.TaskFlow
}

func (s *VerifyPlaybookRequest) SetPlaybookUuid(v string) *VerifyPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *VerifyPlaybookRequest) SetTaskFlow(v string) *VerifyPlaybookRequest {
	s.TaskFlow = &v
	return s
}

func (s *VerifyPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
