// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DebugPlaybookRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DebugPlaybookRequest
	GetPlaybookUuid() *string
	SetRecord(v string) *DebugPlaybookRequest
	GetRecord() *string
	SetTaskflow(v string) *DebugPlaybookRequest
	GetTaskflow() *string
}

type DebugPlaybookRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f916b93e-e814-459f-9662-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The input parameters that you use to debug the playbook. You can define the parameters based on your business requirements.
	//
	// example:
	//
	// {
	//
	//    "param1":"a",
	//
	//    "param2":"b"
	//
	// }
	Record *string `json:"Record,omitempty" xml:"Record,omitempty"`
	// The XML configuration of the playbook.
	//
	// >  You can call the [DescribePlaybook](~~DescribePlaybook~~) operation to query the XML configuration of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// <?xml version="1.0" encoding="UTF-8"?><bpmn:definitions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" targetNamespace="http://bpmn.io/schema/bpmn" id="Definitions_1"><bpmn:process id="Process_1" isExecutable="false"><bpmn:startEvent id="StartEvent_1"/></bpmn:process><bpmndi:BPMNDiagram id="BPMNDiagram_1"><bpmndi:BPMNPlane id="BPMNPlane_1" bpmnElement="Process_1"><bpmndi:BPMNShape id="_BPMNShape_StartEvent_2" bpmnElement="StartEvent_1"><dc:Bounds height="36.0" width="36.0" x="173.0" y="102.0"/></bpmndi:BPMNShape></bpmndi:BPMNPlane></bpmndi:BPMNDiagram></bpmn:definitions>
	Taskflow *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
}

func (s DebugPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s DebugPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DebugPlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *DebugPlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DebugPlaybookRequest) GetRecord() *string {
	return s.Record
}

func (s *DebugPlaybookRequest) GetTaskflow() *string {
	return s.Taskflow
}

func (s *DebugPlaybookRequest) SetLang(v string) *DebugPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DebugPlaybookRequest) SetPlaybookUuid(v string) *DebugPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DebugPlaybookRequest) SetRecord(v string) *DebugPlaybookRequest {
	s.Record = &v
	return s
}

func (s *DebugPlaybookRequest) SetTaskflow(v string) *DebugPlaybookRequest {
	s.Taskflow = &v
	return s
}

func (s *DebugPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
