// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManualDagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizDate(v string) *CreateManualDagRequest
	GetBizDate() *string
	SetDagParameters(v string) *CreateManualDagRequest
	GetDagParameters() *string
	SetExcludeNodeIds(v string) *CreateManualDagRequest
	GetExcludeNodeIds() *string
	SetFlowName(v string) *CreateManualDagRequest
	GetFlowName() *string
	SetIncludeNodeIds(v string) *CreateManualDagRequest
	GetIncludeNodeIds() *string
	SetNodeParameters(v string) *CreateManualDagRequest
	GetNodeParameters() *string
	SetProjectEnv(v string) *CreateManualDagRequest
	GetProjectEnv() *string
	SetProjectName(v string) *CreateManualDagRequest
	GetProjectName() *string
}

type CreateManualDagRequest struct {
	// The data timestamp. The value of the data timestamp must be one or more days before the current date. For example, if the current date is November 11, 2020, set the value to 2020-11-10 00:00:00 or earlier. Configure this parameter in the YYYY-MM-DD 00:00:00 format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-11 00:00:00
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The parameters of the manually triggered workflow, which are synchronized to all the instances in the directed acyclic graph (DAG) of the workflow. If a workflow parameter specified in DagParameters is referenced as a scheduling parameter of a node, the value of the scheduling parameter is replaced with the value of the workflow parameter.
	//
	// example:
	//
	// {"kaaaa": "vaaaaa", "kbbbb": "vbbbbb"}
	DagParameters *string `json:"DagParameters,omitempty" xml:"DagParameters,omitempty"`
	// The IDs of the nodes that do not need to be run.
	//
	// example:
	//
	// 123,456
	ExcludeNodeIds *string `json:"ExcludeNodeIds,omitempty" xml:"ExcludeNodeIds,omitempty"`
	// The name of the manually triggered workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_workflow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The IDs of the nodes that you want to run.
	//
	// example:
	//
	// 74324,74325
	IncludeNodeIds *string `json:"IncludeNodeIds,omitempty" xml:"IncludeNodeIds,omitempty"`
	// The parameters transmitted between nodes in the manually triggered workflow. The parameters are in the following JSON format: `{ "<ID of a node in the manually triggered workflow>": "Scheduling parameter settings of the node, which are in the same format as the parameters in the Scheduling Parameter section on the Properties tab of the DataStudio page", "<ID of a node in the manually triggered workflow>": "Scheduling parameter settings of the node, which are in the same format as the parameters in the Scheduling Parameter section on the Properties tab of the DataStudio page" }`
	//
	// example:
	//
	// {"20000123121": "key1=val2 key2=val2", "20000123124": "kkkk=vvvvv aaaa=bbbb"}
	NodeParameters *string `json:"NodeParameters,omitempty" xml:"NodeParameters,omitempty"`
	// The environment type of Operation Center. Valid values: PROD and DEV.
	//
	// This parameter is required.
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The name of the workspace to which the manually triggered workflow belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_workspace
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateManualDagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateManualDagRequest) GoString() string {
	return s.String()
}

func (s *CreateManualDagRequest) GetBizDate() *string {
	return s.BizDate
}

func (s *CreateManualDagRequest) GetDagParameters() *string {
	return s.DagParameters
}

func (s *CreateManualDagRequest) GetExcludeNodeIds() *string {
	return s.ExcludeNodeIds
}

func (s *CreateManualDagRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateManualDagRequest) GetIncludeNodeIds() *string {
	return s.IncludeNodeIds
}

func (s *CreateManualDagRequest) GetNodeParameters() *string {
	return s.NodeParameters
}

func (s *CreateManualDagRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *CreateManualDagRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateManualDagRequest) SetBizDate(v string) *CreateManualDagRequest {
	s.BizDate = &v
	return s
}

func (s *CreateManualDagRequest) SetDagParameters(v string) *CreateManualDagRequest {
	s.DagParameters = &v
	return s
}

func (s *CreateManualDagRequest) SetExcludeNodeIds(v string) *CreateManualDagRequest {
	s.ExcludeNodeIds = &v
	return s
}

func (s *CreateManualDagRequest) SetFlowName(v string) *CreateManualDagRequest {
	s.FlowName = &v
	return s
}

func (s *CreateManualDagRequest) SetIncludeNodeIds(v string) *CreateManualDagRequest {
	s.IncludeNodeIds = &v
	return s
}

func (s *CreateManualDagRequest) SetNodeParameters(v string) *CreateManualDagRequest {
	s.NodeParameters = &v
	return s
}

func (s *CreateManualDagRequest) SetProjectEnv(v string) *CreateManualDagRequest {
	s.ProjectEnv = &v
	return s
}

func (s *CreateManualDagRequest) SetProjectName(v string) *CreateManualDagRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateManualDagRequest) Validate() error {
	return dara.Validate(s)
}
