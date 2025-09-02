// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunManualDagNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizDate(v string) *RunManualDagNodesRequest
	GetBizDate() *string
	SetDagParameters(v string) *RunManualDagNodesRequest
	GetDagParameters() *string
	SetEndBizDate(v string) *RunManualDagNodesRequest
	GetEndBizDate() *string
	SetExcludeNodeIds(v string) *RunManualDagNodesRequest
	GetExcludeNodeIds() *string
	SetFlowName(v string) *RunManualDagNodesRequest
	GetFlowName() *string
	SetIncludeNodeIds(v string) *RunManualDagNodesRequest
	GetIncludeNodeIds() *string
	SetNodeParameters(v string) *RunManualDagNodesRequest
	GetNodeParameters() *string
	SetProjectEnv(v string) *RunManualDagNodesRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *RunManualDagNodesRequest
	GetProjectId() *int64
	SetProjectName(v string) *RunManualDagNodesRequest
	GetProjectName() *string
	SetStartBizDate(v string) *RunManualDagNodesRequest
	GetStartBizDate() *string
}

type RunManualDagNodesRequest struct {
	// The data timestamp. The value of the data timestamp must be one or more days before the current date. For example, if the current date is November 11, 2020, set the value to 2020-11-10 00:00:00 or earlier. Configure this parameter in the YYYY-MM-DD 00:00:00 format. The StartBizDate parameter is used together with the EndBizDate parameter. You can configure only the BizDate parameter or the StartBizDate and EndBizDate parameters.
	//
	// example:
	//
	// 2020-11-11 00:00:00
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The parameters are synchronized to all the instances in the directed acyclic graph (DAG) of the workflow. If a workflow parameter specified in DagParameters is referenced as a scheduling parameter of a [node](https://help.aliyun.com/document_detail/147245.html), the value of the scheduling parameter is replaced with the value of the workflow parameter.
	//
	// example:
	//
	// {"kaaaa": "vaaaaa", "kbbbb": "vbbbbb"}
	DagParameters *string `json:"DagParameters,omitempty" xml:"DagParameters,omitempty"`
	// The end of the time range in which data generated needs to be processed. Configure this parameter in the yyyy-MM-dd HH:mm:ss format. The StartBizDate parameter is used together with the EndBizDate parameter. You can configure only the BizDate parameter or the StartBizDate and EndBizDate parameters.
	//
	// example:
	//
	// 2020-02-03 00:00:00
	EndBizDate *string `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
	// The IDs of the nodes that you do not need to run in the manually triggered workflow. DataWorks generates dry-run instances for all these nodes. After the dry-run instances are scheduled, the states of these instances are directly set to successful, but the scripts are not run. Separate multiple node IDs with commas (,). The ExcludeNodeIds parameter must be used together with the IncludeNodeIds parameter. This way, the settings of the ExcludeNodeIds parameter can take effect.
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
	// The IDs of the nodes that you need to run in the manually triggered workflow. Separate multiple node IDs with commas (,).
	//
	// example:
	//
	// 74324,74325
	IncludeNodeIds *string `json:"IncludeNodeIds,omitempty" xml:"IncludeNodeIds,omitempty"`
	// The scheduling parameters of nodes in the manually triggered workflow. Configure NodeParameters in the following JSON format: {"\\<ID of a node in the manually triggered workflow>": "Scheduling parameter settings of the node, which are in the same format as the parameter settings in the Scheduling Parameter section of the Properties tab on the DataStudio page", "\\<ID of a node in the manually triggered workflow>": "Scheduling parameter settings of the node, which are in the same format as the parameter settings in the Scheduling Parameter section of the Properties tab on the DataStudio page"}.
	//
	// example:
	//
	// {"20000123121": "key1=val2 key2=val2", "20000123124": "kkkk=vvvvv aaaa=bbbb"}
	NodeParameters *string `json:"NodeParameters,omitempty" xml:"NodeParameters,omitempty"`
	// The environment type of Operation Center. Valid values: PROD and DEV. The value PROD indicates the production environment. The value DEV indicates the development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD or DEV
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The ID of the workspace to which the manually triggered workflow belongs.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the workspace to which the manually triggered workflow belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_workspace
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The beginning of the time range in which data generated needs to be processed. Configure this parameter in the yyyy-MM-dd HH:mm:ss format. The StartBizDate parameter is used together with the EndBizDate parameter. You can configure only the BizDate parameter or the StartBizDate and EndBizDate parameters.
	//
	// example:
	//
	// 2020-02-02 00:00:00
	StartBizDate *string `json:"StartBizDate,omitempty" xml:"StartBizDate,omitempty"`
}

func (s RunManualDagNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RunManualDagNodesRequest) GoString() string {
	return s.String()
}

func (s *RunManualDagNodesRequest) GetBizDate() *string {
	return s.BizDate
}

func (s *RunManualDagNodesRequest) GetDagParameters() *string {
	return s.DagParameters
}

func (s *RunManualDagNodesRequest) GetEndBizDate() *string {
	return s.EndBizDate
}

func (s *RunManualDagNodesRequest) GetExcludeNodeIds() *string {
	return s.ExcludeNodeIds
}

func (s *RunManualDagNodesRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *RunManualDagNodesRequest) GetIncludeNodeIds() *string {
	return s.IncludeNodeIds
}

func (s *RunManualDagNodesRequest) GetNodeParameters() *string {
	return s.NodeParameters
}

func (s *RunManualDagNodesRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *RunManualDagNodesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RunManualDagNodesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *RunManualDagNodesRequest) GetStartBizDate() *string {
	return s.StartBizDate
}

func (s *RunManualDagNodesRequest) SetBizDate(v string) *RunManualDagNodesRequest {
	s.BizDate = &v
	return s
}

func (s *RunManualDagNodesRequest) SetDagParameters(v string) *RunManualDagNodesRequest {
	s.DagParameters = &v
	return s
}

func (s *RunManualDagNodesRequest) SetEndBizDate(v string) *RunManualDagNodesRequest {
	s.EndBizDate = &v
	return s
}

func (s *RunManualDagNodesRequest) SetExcludeNodeIds(v string) *RunManualDagNodesRequest {
	s.ExcludeNodeIds = &v
	return s
}

func (s *RunManualDagNodesRequest) SetFlowName(v string) *RunManualDagNodesRequest {
	s.FlowName = &v
	return s
}

func (s *RunManualDagNodesRequest) SetIncludeNodeIds(v string) *RunManualDagNodesRequest {
	s.IncludeNodeIds = &v
	return s
}

func (s *RunManualDagNodesRequest) SetNodeParameters(v string) *RunManualDagNodesRequest {
	s.NodeParameters = &v
	return s
}

func (s *RunManualDagNodesRequest) SetProjectEnv(v string) *RunManualDagNodesRequest {
	s.ProjectEnv = &v
	return s
}

func (s *RunManualDagNodesRequest) SetProjectId(v int64) *RunManualDagNodesRequest {
	s.ProjectId = &v
	return s
}

func (s *RunManualDagNodesRequest) SetProjectName(v string) *RunManualDagNodesRequest {
	s.ProjectName = &v
	return s
}

func (s *RunManualDagNodesRequest) SetStartBizDate(v string) *RunManualDagNodesRequest {
	s.StartBizDate = &v
	return s
}

func (s *RunManualDagNodesRequest) Validate() error {
	return dara.Validate(s)
}
