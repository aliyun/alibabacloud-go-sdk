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
	// The value of the business date must be less than or equal to the current date minus one day. For example, if today is November 11, 2020, the business date must be 2020-11-10 00:00:00 or an earlier date. The hour, minute, and second fields of the business date must all be set to 00.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-11 00:00:00
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The workflow parameters. These parameters are synchronized to all instances of the current DAG. If the scheduling parameters of an internal node reference the workflow parameters in DagParameters, the corresponding parameter values of the node are replaced with the workflow parameters in DagParameters.
	//
	// example:
	//
	// {"kaaaa": "vaaaaa", "kbbbb": "vbbbbb"}
	DagParameters *string `json:"DagParameters,omitempty" xml:"DagParameters,omitempty"`
	// The list of IDs of the nodes that do not need to be run.
	//
	// example:
	//
	// 123,456
	ExcludeNodeIds *string `json:"ExcludeNodeIds,omitempty" xml:"ExcludeNodeIds,omitempty"`
	// The name of the manual workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_workflow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The list of IDs of the nodes to be run.
	//
	// example:
	//
	// 74324,74325
	IncludeNodeIds *string `json:"IncludeNodeIds,omitempty" xml:"IncludeNodeIds,omitempty"`
	// The node parameter information passed when the manual workflow is executed, in JSON format:
	//
	// `
	//
	// {
	//
	//      "<A node ID inside the manual workflow>": "The scheduling parameter information of the node, consistent with the parameter format in the data development scheduling configuration",
	//
	//      "<A node ID inside the manual workflow>": "The scheduling parameter information of the node, consistent with the parameter format in the data development scheduling configuration"
	//
	// }
	//
	// `
	//
	// example:
	//
	// {"20000123121": "key1=val2 key2=val2", "20000123124": "kkkk=vvvvv aaaa=bbbb"}
	NodeParameters *string `json:"NodeParameters,omitempty" xml:"NodeParameters,omitempty"`
	// The environment identifier of the Scheduling Operation Center. PROD indicates the production environment, and DEV indicates the development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD or DEV
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The English name of the workspace to which the manual workflow belongs.
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
