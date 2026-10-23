// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListActionRecordsResponseBody
	GetRequestId() *string
	SetResult(v []*ListActionRecordsResponseBodyResult) *ListActionRecordsResponseBody
	GetResult() []*ListActionRecordsResponseBodyResult
}

type ListActionRecordsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of instance change records.
	Result []*ListActionRecordsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListActionRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListActionRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActionRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListActionRecordsResponseBody) GetResult() []*ListActionRecordsResponseBodyResult {
	return s.Result
}

func (s *ListActionRecordsResponseBody) SetRequestId(v string) *ListActionRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActionRecordsResponseBody) SetResult(v []*ListActionRecordsResponseBodyResult) *ListActionRecordsResponseBody {
	s.Result = v
	return s
}

func (s *ListActionRecordsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListActionRecordsResponseBodyResult struct {
	// The name of the operation that triggered the change.
	//
	// example:
	//
	// RestartInstance
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// The parameters used for the change.
	ActionParams           map[string]interface{} `json:"actionParams,omitempty" xml:"actionParams,omitempty"`
	ActionResultAccessList []*string              `json:"actionResultAccessList,omitempty" xml:"actionResultAccessList,omitempty" type:"Repeated"`
	// The end time of the change. The value is a UNIX timestamp in milliseconds.
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The ID of the Elasticsearch instance that was changed.
	//
	// example:
	//
	// es-cn-nif24adwc0082****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance information after the change, represented as a JSON string.
	MetaNow *string `json:"metaNow,omitempty" xml:"metaNow,omitempty"`
	// The instance information before the change, represented as a JSON string.
	MetaOld *string `json:"metaOld,omitempty" xml:"metaOld,omitempty"`
	// The ID of the Alibaba Cloud account to which the account that triggered the change belongs.
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// The progress of the change, represented as a percentage string.
	//
	// example:
	//
	// 100%
	Process *string `json:"process,omitempty" xml:"process,omitempty"`
	// The field differences before and after the change.
	RecordDiff map[string]interface{} `json:"recordDiff,omitempty" xml:"recordDiff,omitempty"`
	RecordIds  []*string              `json:"recordIds,omitempty" xml:"recordIds,omitempty" type:"Repeated"`
	// The ID of the request that triggered the change.
	//
	// example:
	//
	// request-123
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The start time of the change. The value is a UNIX timestamp in milliseconds.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The current state of the change.
	//
	// example:
	//
	// FINISHED
	StateType *string `json:"stateType,omitempty" xml:"stateType,omitempty"`
	// The list of status details for each phase of the change.
	StatusInfo []*ListActionRecordsResponseBodyResultStatusInfo `json:"statusInfo,omitempty" xml:"statusInfo,omitempty" type:"Repeated"`
	// The ID of the user that triggered the change.
	//
	// example:
	//
	// user-123
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserInfo *string `json:"userInfo,omitempty" xml:"userInfo,omitempty"`
	// The type of the account that triggered the change.
	//
	// example:
	//
	// customer
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s ListActionRecordsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListActionRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListActionRecordsResponseBodyResult) GetActionName() *string {
	return s.ActionName
}

func (s *ListActionRecordsResponseBodyResult) GetActionParams() map[string]interface{} {
	return s.ActionParams
}

func (s *ListActionRecordsResponseBodyResult) GetActionResultAccessList() []*string {
	return s.ActionResultAccessList
}

func (s *ListActionRecordsResponseBodyResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListActionRecordsResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListActionRecordsResponseBodyResult) GetMetaNow() *string {
	return s.MetaNow
}

func (s *ListActionRecordsResponseBodyResult) GetMetaOld() *string {
	return s.MetaOld
}

func (s *ListActionRecordsResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListActionRecordsResponseBodyResult) GetProcess() *string {
	return s.Process
}

func (s *ListActionRecordsResponseBodyResult) GetRecordDiff() map[string]interface{} {
	return s.RecordDiff
}

func (s *ListActionRecordsResponseBodyResult) GetRecordIds() []*string {
	return s.RecordIds
}

func (s *ListActionRecordsResponseBodyResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListActionRecordsResponseBodyResult) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListActionRecordsResponseBodyResult) GetStateType() *string {
	return s.StateType
}

func (s *ListActionRecordsResponseBodyResult) GetStatusInfo() []*ListActionRecordsResponseBodyResultStatusInfo {
	return s.StatusInfo
}

func (s *ListActionRecordsResponseBodyResult) GetUserId() *string {
	return s.UserId
}

func (s *ListActionRecordsResponseBodyResult) GetUserInfo() *string {
	return s.UserInfo
}

func (s *ListActionRecordsResponseBodyResult) GetUserType() *string {
	return s.UserType
}

func (s *ListActionRecordsResponseBodyResult) SetActionName(v string) *ListActionRecordsResponseBodyResult {
	s.ActionName = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetActionParams(v map[string]interface{}) *ListActionRecordsResponseBodyResult {
	s.ActionParams = v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetActionResultAccessList(v []*string) *ListActionRecordsResponseBodyResult {
	s.ActionResultAccessList = v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetEndTime(v int64) *ListActionRecordsResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetInstanceId(v string) *ListActionRecordsResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetMetaNow(v string) *ListActionRecordsResponseBodyResult {
	s.MetaNow = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetMetaOld(v string) *ListActionRecordsResponseBodyResult {
	s.MetaOld = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetOwnerId(v string) *ListActionRecordsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetProcess(v string) *ListActionRecordsResponseBodyResult {
	s.Process = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetRecordDiff(v map[string]interface{}) *ListActionRecordsResponseBodyResult {
	s.RecordDiff = v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetRecordIds(v []*string) *ListActionRecordsResponseBodyResult {
	s.RecordIds = v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetRequestId(v string) *ListActionRecordsResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetStartTime(v int64) *ListActionRecordsResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetStateType(v string) *ListActionRecordsResponseBodyResult {
	s.StateType = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetStatusInfo(v []*ListActionRecordsResponseBodyResultStatusInfo) *ListActionRecordsResponseBodyResult {
	s.StatusInfo = v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetUserId(v string) *ListActionRecordsResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetUserInfo(v string) *ListActionRecordsResponseBodyResult {
	s.UserInfo = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) SetUserType(v string) *ListActionRecordsResponseBodyResult {
	s.UserType = &v
	return s
}

func (s *ListActionRecordsResponseBodyResult) Validate() error {
	if s.StatusInfo != nil {
		for _, item := range s.StatusInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListActionRecordsResponseBodyResultStatusInfo struct {
	// The number of completed nodes in the phase.
	//
	// example:
	//
	// 3
	CompleteNodeCount *int32 `json:"completeNodeCount,omitempty" xml:"completeNodeCount,omitempty"`
	// The end time of the phase. The value is a UNIX timestamp in milliseconds.
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The exception information for the phase.
	Exception *string `json:"exception,omitempty" xml:"exception,omitempty"`
	// The time elapsed for the phase, in milliseconds.
	//
	// example:
	//
	// 1000
	LatencyMills *int64 `json:"latencyMills,omitempty" xml:"latencyMills,omitempty"`
	// The total number of nodes involved in the phase.
	//
	// example:
	//
	// 5
	NodeCount *int32 `json:"nodeCount,omitempty" xml:"nodeCount,omitempty"`
	// The completion progress of the phase, represented as a percentage string.
	//
	// example:
	//
	// 100%
	Process *string `json:"process,omitempty" xml:"process,omitempty"`
	// The start time of the phase. The value is a UNIX timestamp in milliseconds.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The current state of the phase.
	//
	// example:
	//
	// FINISHED
	StateType *string `json:"stateType,omitempty" xml:"stateType,omitempty"`
	// The name of the phase.
	//
	// example:
	//
	// rolling
	SubState *string `json:"subState,omitempty" xml:"subState,omitempty"`
	// The list of status details for the sub-phases included in the current phase.
	SubStatusInfo []*ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo `json:"subStatusInfo,omitempty" xml:"subStatusInfo,omitempty" type:"Repeated"`
}

func (s ListActionRecordsResponseBodyResultStatusInfo) String() string {
	return dara.Prettify(s)
}

func (s ListActionRecordsResponseBodyResultStatusInfo) GoString() string {
	return s.String()
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetCompleteNodeCount() *int32 {
	return s.CompleteNodeCount
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetException() *string {
	return s.Exception
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetLatencyMills() *int64 {
	return s.LatencyMills
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetProcess() *string {
	return s.Process
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetStateType() *string {
	return s.StateType
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetSubState() *string {
	return s.SubState
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) GetSubStatusInfo() []*ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	return s.SubStatusInfo
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetCompleteNodeCount(v int32) *ListActionRecordsResponseBodyResultStatusInfo {
	s.CompleteNodeCount = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetEndTime(v int64) *ListActionRecordsResponseBodyResultStatusInfo {
	s.EndTime = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetException(v string) *ListActionRecordsResponseBodyResultStatusInfo {
	s.Exception = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetLatencyMills(v int64) *ListActionRecordsResponseBodyResultStatusInfo {
	s.LatencyMills = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetNodeCount(v int32) *ListActionRecordsResponseBodyResultStatusInfo {
	s.NodeCount = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetProcess(v string) *ListActionRecordsResponseBodyResultStatusInfo {
	s.Process = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetStartTime(v int64) *ListActionRecordsResponseBodyResultStatusInfo {
	s.StartTime = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetStateType(v string) *ListActionRecordsResponseBodyResultStatusInfo {
	s.StateType = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetSubState(v string) *ListActionRecordsResponseBodyResultStatusInfo {
	s.SubState = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) SetSubStatusInfo(v []*ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) *ListActionRecordsResponseBodyResultStatusInfo {
	s.SubStatusInfo = v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfo) Validate() error {
	if s.SubStatusInfo != nil {
		for _, item := range s.SubStatusInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo struct {
	// The number of completed nodes in the sub-phase.
	//
	// example:
	//
	// 3
	CompleteNodeCount *int32 `json:"completeNodeCount,omitempty" xml:"completeNodeCount,omitempty"`
	// The end time of the sub-phase. The value is a UNIX timestamp in milliseconds.
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The exception information for the sub-phase.
	Exception *string `json:"exception,omitempty" xml:"exception,omitempty"`
	// The time elapsed for the sub-phase, in milliseconds.
	//
	// example:
	//
	// 1000
	LatencyMills *int64 `json:"latencyMills,omitempty" xml:"latencyMills,omitempty"`
	// The total number of nodes involved in the sub-phase.
	//
	// example:
	//
	// 5
	NodeCount *int32 `json:"nodeCount,omitempty" xml:"nodeCount,omitempty"`
	// The completion progress of the sub-phase, represented as a percentage string.
	//
	// example:
	//
	// 100%
	Process *string `json:"process,omitempty" xml:"process,omitempty"`
	// The start time of the sub-phase. The value is a UNIX timestamp in milliseconds.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The current state of the sub-phase.
	//
	// example:
	//
	// FINISHED
	StateType *string `json:"stateType,omitempty" xml:"stateType,omitempty"`
	// The name of the sub-phase.
	//
	// example:
	//
	// clusterStateWait
	SubState *string `json:"subState,omitempty" xml:"subState,omitempty"`
}

func (s ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) String() string {
	return dara.Prettify(s)
}

func (s ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GoString() string {
	return s.String()
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GetCompleteNodeCount() *int32 {
	return s.CompleteNodeCount
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GetException() *string {
	return s.Exception
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GetLatencyMills() *int64 {
	return s.LatencyMills
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GetProcess() *string {
	return s.Process
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GetStateType() *string {
	return s.StateType
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) GetSubState() *string {
	return s.SubState
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) SetCompleteNodeCount(v int32) *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	s.CompleteNodeCount = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) SetEndTime(v int64) *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	s.EndTime = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) SetException(v string) *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	s.Exception = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) SetLatencyMills(v int64) *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	s.LatencyMills = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) SetNodeCount(v int32) *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	s.NodeCount = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) SetProcess(v string) *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	s.Process = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) SetStartTime(v int64) *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	s.StartTime = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) SetStateType(v string) *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	s.StateType = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) SetSubState(v string) *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo {
	s.SubState = &v
	return s
}

func (s *ListActionRecordsResponseBodyResultStatusInfoSubStatusInfo) Validate() error {
	return dara.Validate(s)
}
