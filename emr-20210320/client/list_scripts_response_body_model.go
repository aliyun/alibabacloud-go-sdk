// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListScriptsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListScriptsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListScriptsResponseBody
	GetRequestId() *string
	SetScripts(v []*ListScriptsResponseBodyScripts) *ListScriptsResponseBody
	GetScripts() []*ListScriptsResponseBodyScripts
	SetTotalCount(v int32) *ListScriptsResponseBody
	GetTotalCount() *int32
}

type ListScriptsResponseBody struct {
	// The maximum number of entries returned in the request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// dd6b1b2a-5837-5237-abe4-ff0c89568982
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of scripts.
	Scripts []*ListScriptsResponseBodyScripts `json:"Scripts,omitempty" xml:"Scripts,omitempty" type:"Repeated"`
	// The total number of entries that meet the request conditions.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListScriptsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptsResponseBody) GetScripts() []*ListScriptsResponseBodyScripts {
	return s.Scripts
}

func (s *ListScriptsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScriptsResponseBody) SetMaxResults(v int32) *ListScriptsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListScriptsResponseBody) SetNextToken(v string) *ListScriptsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListScriptsResponseBody) SetRequestId(v string) *ListScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptsResponseBody) SetScripts(v []*ListScriptsResponseBodyScripts) *ListScriptsResponseBody {
	s.Scripts = v
	return s
}

func (s *ListScriptsResponseBody) SetTotalCount(v int32) *ListScriptsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScriptsResponseBody) Validate() error {
	if s.Scripts != nil {
		for _, item := range s.Scripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptsResponseBodyScripts struct {
	// The name of the API.
	//
	// example:
	//
	// ListScripts
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The time when the execution ended. This parameter is returned only when ScriptType is set to NORMAL.
	//
	// example:
	//
	// 1639715635819
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The policy used to handle an execution failure. Valid values:
	//
	// - FAILED_CONTINUE: Continue the execution.
	//
	// - FAILED_BLOCK: Block the execution.
	//
	// example:
	//
	// FAILED_CONTINUE
	ExecutionFailStrategy *string `json:"ExecutionFailStrategy,omitempty" xml:"ExecutionFailStrategy,omitempty"`
	// The time to execute the script. Valid values:
	//
	// - BEFORE_INSTALL: before application installation.
	//
	// - AFTER_STARTED: after application startup.
	//
	// example:
	//
	// BEFORE_INSTALL
	ExecutionMoment *string `json:"ExecutionMoment,omitempty" xml:"ExecutionMoment,omitempty"`
	// The execution status of the script. This parameter is returned only when `ScriptType` is set to `NORMAL`. Valid values:
	//
	// - SCRIPT_COMPLETED: The script is successfully executed.
	//
	// - SCRIPT_SUBMISSION_FAILED: The script fails to be executed.
	//
	// - SCRIPT_RUNNING: The script is being executed.
	//
	// example:
	//
	// SCRIPT_COMPLETED
	ExecutionState *string `json:"ExecutionState,omitempty" xml:"ExecutionState,omitempty"`
	// The time when the script was last updated.
	//
	// example:
	//
	// 1639714634819
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The node selector.
	NodeSelector *NodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The runtime parameters of the script.
	//
	// example:
	//
	// --mode=client -h -p
	ScriptArgs *string `json:"ScriptArgs,omitempty" xml:"ScriptArgs,omitempty"`
	// The script ID.
	//
	// example:
	//
	// cs-bf25219d103043a0820613e32781****
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The script name.
	//
	// example:
	//
	// check_env
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The script path.
	//
	// example:
	//
	// oss://bucket1/check_evn.sh
	ScriptPath *string `json:"ScriptPath,omitempty" xml:"ScriptPath,omitempty"`
	// The time when the execution started. This parameter is returned only when ScriptType is set to NORMAL.
	//
	// example:
	//
	// 1639714634000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListScriptsResponseBodyScripts) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBodyScripts) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBodyScripts) GetAction() *string {
	return s.Action
}

func (s *ListScriptsResponseBodyScripts) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListScriptsResponseBodyScripts) GetExecutionFailStrategy() *string {
	return s.ExecutionFailStrategy
}

func (s *ListScriptsResponseBodyScripts) GetExecutionMoment() *string {
	return s.ExecutionMoment
}

func (s *ListScriptsResponseBodyScripts) GetExecutionState() *string {
	return s.ExecutionState
}

func (s *ListScriptsResponseBodyScripts) GetLastUpdateTime() *int64 {
	return s.LastUpdateTime
}

func (s *ListScriptsResponseBodyScripts) GetNodeSelector() *NodeSelector {
	return s.NodeSelector
}

func (s *ListScriptsResponseBodyScripts) GetRegionId() *string {
	return s.RegionId
}

func (s *ListScriptsResponseBodyScripts) GetScriptArgs() *string {
	return s.ScriptArgs
}

func (s *ListScriptsResponseBodyScripts) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptsResponseBodyScripts) GetScriptName() *string {
	return s.ScriptName
}

func (s *ListScriptsResponseBodyScripts) GetScriptPath() *string {
	return s.ScriptPath
}

func (s *ListScriptsResponseBodyScripts) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListScriptsResponseBodyScripts) SetAction(v string) *ListScriptsResponseBodyScripts {
	s.Action = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetEndTime(v int64) *ListScriptsResponseBodyScripts {
	s.EndTime = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetExecutionFailStrategy(v string) *ListScriptsResponseBodyScripts {
	s.ExecutionFailStrategy = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetExecutionMoment(v string) *ListScriptsResponseBodyScripts {
	s.ExecutionMoment = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetExecutionState(v string) *ListScriptsResponseBodyScripts {
	s.ExecutionState = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetLastUpdateTime(v int64) *ListScriptsResponseBodyScripts {
	s.LastUpdateTime = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetNodeSelector(v *NodeSelector) *ListScriptsResponseBodyScripts {
	s.NodeSelector = v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetRegionId(v string) *ListScriptsResponseBodyScripts {
	s.RegionId = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetScriptArgs(v string) *ListScriptsResponseBodyScripts {
	s.ScriptArgs = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetScriptId(v string) *ListScriptsResponseBodyScripts {
	s.ScriptId = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetScriptName(v string) *ListScriptsResponseBodyScripts {
	s.ScriptName = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetScriptPath(v string) *ListScriptsResponseBodyScripts {
	s.ScriptPath = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetStartTime(v int64) *ListScriptsResponseBodyScripts {
	s.StartTime = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) Validate() error {
	if s.NodeSelector != nil {
		if err := s.NodeSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}
