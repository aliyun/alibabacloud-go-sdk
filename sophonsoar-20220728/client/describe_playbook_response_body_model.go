// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlaybook(v *DescribePlaybookResponseBodyPlaybook) *DescribePlaybookResponseBody
	GetPlaybook() *DescribePlaybookResponseBodyPlaybook
	SetRequestId(v string) *DescribePlaybookResponseBody
	GetRequestId() *string
}

type DescribePlaybookResponseBody struct {
	// The configuration of the playbook.
	Playbook *DescribePlaybookResponseBodyPlaybook `json:"Playbook,omitempty" xml:"Playbook,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2989BC59-E9F0-5C83-B453-B368857649C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookResponseBody) GetPlaybook() *DescribePlaybookResponseBodyPlaybook {
	return s.Playbook
}

func (s *DescribePlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlaybookResponseBody) SetPlaybook(v *DescribePlaybookResponseBodyPlaybook) *DescribePlaybookResponseBody {
	s.Playbook = v
	return s
}

func (s *DescribePlaybookResponseBody) SetRequestId(v string) *DescribePlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlaybookResponseBody) Validate() error {
	if s.Playbook != nil {
		if err := s.Playbook.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePlaybookResponseBodyPlaybook struct {
	// The description of the playbook.
	//
	// example:
	//
	// demo playbook
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the playbook.
	//
	// example:
	//
	// demo_test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The number of times that the playbook failed to be run.
	//
	// example:
	//
	// 1
	FailExeNum *int32 `json:"FailExeNum,omitempty" xml:"FailExeNum,omitempty"`
	// The creation time of the playbook. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1665288858000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time of the playbook. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1677482519000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The input parameter configuration of the playbook. The value is a JSON array.
	//
	// >  For more information, see [DescribePlaybookInputOutput](~~DescribePlaybookInputOutput~~).
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "typeName": "String",
	//
	//         "dataClass": "normal",
	//
	//         "dataType": "String",
	//
	//         "description": "period",
	//
	//         "example": "",
	//
	//         "name": "period",
	//
	//         "required": false
	//
	//     }
	//
	// ]
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The time when the playbook was last run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1665288858000
	LastExeTime *int64 `json:"LastExeTime,omitempty" xml:"LastExeTime,omitempty"`
	// The status of the playbook. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 0
	OnlineActive *bool `json:"OnlineActive,omitempty" xml:"OnlineActive,omitempty"`
	// The MD5 hash value in the latest published version of the playbook.
	//
	// example:
	//
	// asdfsdfe232-e2b2-44fd-b2cc-xxxxx
	OnlineReleaseTaskflowMd5 *string `json:"OnlineReleaseTaskflowMd5,omitempty" xml:"OnlineReleaseTaskflowMd5,omitempty"`
	// The type of the playbook. Valid values:
	//
	// 	- **preset**: predefined playbook
	//
	// 	- **user**: custom playbook
	//
	// example:
	//
	// preset
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// 8db257d3-e2b2-44fd-b2cc-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The number of times that the playbook was successfully run.
	//
	// example:
	//
	// 100
	SuccessExeNum *int32 `json:"SuccessExeNum,omitempty" xml:"SuccessExeNum,omitempty"`
	// The XML configuration of the playbook.
	//
	// example:
	//
	// <?xml version="1.0" encoding="UTF-8"?><bpmn:definitions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" targetNamespace="http://bpmn.io/schema/bpmn" id="Definitions_1"><bpmn:process id="Process_1" isExecutable="false"><bpmn:startEvent id="StartEvent_1"/></bpmn:process><bpmndi:BPMNDiagram id="BPMNDiagram_1"><bpmndi:BPMNPlane id="BPMNPlane_1" bpmnElement="Process_1"><bpmndi:BPMNShape id="_BPMNShape_StartEvent_2" bpmnElement="StartEvent_1"><dc:Bounds height="36.0" width="36.0" x="173.0" y="102.0"/></bpmndi:BPMNShape></bpmndi:BPMNPlane></bpmndi:BPMNDiagram></bpmn:definitions>
	Taskflow *string `json:"Taskflow,omitempty" xml:"Taskflow,omitempty"`
	// The playbook configuration type.
	//
	// 	- **xml**: XML format.
	//
	// 	- **x6**: JSON format.
	//
	// example:
	//
	// xml
	TaskflowType *string `json:"TaskflowType,omitempty" xml:"TaskflowType,omitempty"`
}

func (s DescribePlaybookResponseBodyPlaybook) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookResponseBodyPlaybook) GoString() string {
	return s.String()
}

func (s *DescribePlaybookResponseBodyPlaybook) GetDescription() *string {
	return s.Description
}

func (s *DescribePlaybookResponseBodyPlaybook) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribePlaybookResponseBodyPlaybook) GetFailExeNum() *int32 {
	return s.FailExeNum
}

func (s *DescribePlaybookResponseBodyPlaybook) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribePlaybookResponseBodyPlaybook) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribePlaybookResponseBodyPlaybook) GetInputParams() *string {
	return s.InputParams
}

func (s *DescribePlaybookResponseBodyPlaybook) GetLastExeTime() *int64 {
	return s.LastExeTime
}

func (s *DescribePlaybookResponseBodyPlaybook) GetOnlineActive() *bool {
	return s.OnlineActive
}

func (s *DescribePlaybookResponseBodyPlaybook) GetOnlineReleaseTaskflowMd5() *string {
	return s.OnlineReleaseTaskflowMd5
}

func (s *DescribePlaybookResponseBodyPlaybook) GetOwnType() *string {
	return s.OwnType
}

func (s *DescribePlaybookResponseBodyPlaybook) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribePlaybookResponseBodyPlaybook) GetSuccessExeNum() *int32 {
	return s.SuccessExeNum
}

func (s *DescribePlaybookResponseBodyPlaybook) GetTaskflow() *string {
	return s.Taskflow
}

func (s *DescribePlaybookResponseBodyPlaybook) GetTaskflowType() *string {
	return s.TaskflowType
}

func (s *DescribePlaybookResponseBodyPlaybook) SetDescription(v string) *DescribePlaybookResponseBodyPlaybook {
	s.Description = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetDisplayName(v string) *DescribePlaybookResponseBodyPlaybook {
	s.DisplayName = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetFailExeNum(v int32) *DescribePlaybookResponseBodyPlaybook {
	s.FailExeNum = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetGmtCreate(v string) *DescribePlaybookResponseBodyPlaybook {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetGmtModified(v string) *DescribePlaybookResponseBodyPlaybook {
	s.GmtModified = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetInputParams(v string) *DescribePlaybookResponseBodyPlaybook {
	s.InputParams = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetLastExeTime(v int64) *DescribePlaybookResponseBodyPlaybook {
	s.LastExeTime = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetOnlineActive(v bool) *DescribePlaybookResponseBodyPlaybook {
	s.OnlineActive = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetOnlineReleaseTaskflowMd5(v string) *DescribePlaybookResponseBodyPlaybook {
	s.OnlineReleaseTaskflowMd5 = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetOwnType(v string) *DescribePlaybookResponseBodyPlaybook {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetPlaybookUuid(v string) *DescribePlaybookResponseBodyPlaybook {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetSuccessExeNum(v int32) *DescribePlaybookResponseBodyPlaybook {
	s.SuccessExeNum = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetTaskflow(v string) *DescribePlaybookResponseBodyPlaybook {
	s.Taskflow = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) SetTaskflowType(v string) *DescribePlaybookResponseBodyPlaybook {
	s.TaskflowType = &v
	return s
}

func (s *DescribePlaybookResponseBodyPlaybook) Validate() error {
	return dara.Validate(s)
}
