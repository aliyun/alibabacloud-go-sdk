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
	// The business date. The value must be less than or equal to the current date minus 1 day. For example, if today is November 11, 2020, the business date must be 00:00:00 on November 10, 2020 or an earlier date. The hour, minute, and second values of the business date must all be set to 00.
	//
	// This parameter is used together with the StartBizDate and EndBizDate parameters. You can configure only one of BizDate or the StartBizDate and EndBizDate pair.
	//
	// Format: `yyyy-MM-dd HH:mm:ss`. Example: `2020-11-11 00:00:00`.
	//
	// example:
	//
	// 2020-11-11 00:00:00
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// This parameter is synchronized to all instances of the current dagrun. If the scheduling parameters of internal nodes ([supported node types](https://help.aliyun.com/document_detail/147245.html)) reference workflow parameters in DagParameters, the corresponding parameter values of the nodes are replaced with the workflow parameters in DagParameters.
	//
	// example:
	//
	// {"kaaaa": "vaaaaa", "kbbbb": "vbbbbb"}
	DagParameters *string `json:"DagParameters,omitempty" xml:"DagParameters,omitempty"`
	// The business end date. Format: yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is used together with the StartBizDate parameter. You can configure only one of the StartBizDate and EndBizDate pair or the BizDate parameter.
	//
	// example:
	//
	// 2020-02-03 00:00:00
	EndBizDate *string `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
	// The IDs of nodes that you do not want to run within the workflow. The specified nodes generate dry-run instances during execution. After a dry-run instance is scheduled, it immediately succeeds without executing the script content. Separate multiple node IDs with commas (,).
	//
	// The ExcludeNodeIds parameter takes effect only when used together with the IncludeNodeIds parameter.
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
	// The IDs of specific nodes to run within the manual workflow. Separate multiple node IDs with commas (,).
	//
	// example:
	//
	// 74324,74325
	IncludeNodeIds *string `json:"IncludeNodeIds,omitempty" xml:"IncludeNodeIds,omitempty"`
	// The node parameter information passed when the manual workflow is executed. This corresponds to the **scheduling parameters*	- configured in the **Properties*	- of nodes within the manual workflow.
	//
	// A JSON format: { "<Node ID within the manual workflow>": "Scheduling parameter information of the node, in the same format as the parameters in the data development scheduling configuration", "<Node ID within the manual workflow>": "Scheduling parameter information of the node, in the same format as the parameters in the data development scheduling configuration" }
	//
	// example:
	//
	// {"20000123121": "key1=val2 key2=val2", "20000123124": "kkkk=vvvvv aaaa=bbbb"}
	NodeParameters *string `json:"NodeParameters,omitempty" xml:"NodeParameters,omitempty"`
	// The environment identifier of the Operation Center. PROD indicates the production environment. DEV indicates the development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD or DEV
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the workspace to which the manual workflow belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_workspace
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The business start date. Format: yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is used together with the EndBizDate parameter. You can configure only one of the StartBizDate and EndBizDate pair or the BizDate parameter.
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
