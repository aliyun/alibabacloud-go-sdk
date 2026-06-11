// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListScriptsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListScriptsResponseBody
	GetRequestId() *string
	SetScripts(v *ListScriptsResponseBodyScripts) *ListScriptsResponseBody
	GetScripts() *ListScriptsResponseBodyScripts
	SetSuccess(v bool) *ListScriptsResponseBody
	GetSuccess() *bool
}

type ListScriptsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The paginated list of scripts.
	Scripts *ListScriptsResponseBodyScripts `json:"Scripts,omitempty" xml:"Scripts,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptsResponseBody) GetScripts() *ListScriptsResponseBodyScripts {
	return s.Scripts
}

func (s *ListScriptsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScriptsResponseBody) SetCode(v string) *ListScriptsResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptsResponseBody) SetHttpStatusCode(v int32) *ListScriptsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptsResponseBody) SetMessage(v string) *ListScriptsResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptsResponseBody) SetRequestId(v string) *ListScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptsResponseBody) SetScripts(v *ListScriptsResponseBodyScripts) *ListScriptsResponseBody {
	s.Scripts = v
	return s
}

func (s *ListScriptsResponseBody) SetSuccess(v bool) *ListScriptsResponseBody {
	s.Success = &v
	return s
}

func (s *ListScriptsResponseBody) Validate() error {
	if s.Scripts != nil {
		if err := s.Scripts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScriptsResponseBodyScripts struct {
	// The list of scripts.
	List []*ListScriptsResponseBodyScriptsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of scripts.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScriptsResponseBodyScripts) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBodyScripts) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBodyScripts) GetList() []*ListScriptsResponseBodyScriptsList {
	return s.List
}

func (s *ListScriptsResponseBodyScripts) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptsResponseBodyScripts) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptsResponseBodyScripts) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScriptsResponseBodyScripts) SetList(v []*ListScriptsResponseBodyScriptsList) *ListScriptsResponseBodyScripts {
	s.List = v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetPageNumber(v int32) *ListScriptsResponseBodyScripts {
	s.PageNumber = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetPageSize(v int32) *ListScriptsResponseBodyScripts {
	s.PageSize = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) SetTotalCount(v int32) *ListScriptsResponseBodyScripts {
	s.TotalCount = &v
	return s
}

func (s *ListScriptsResponseBodyScripts) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptsResponseBodyScriptsList struct {
	// The agent access key.
	//
	// example:
	//
	// 14791f5f226b4878b3d9b676a0291234
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Indicates whether the agent is an LLM agent.
	//
	// example:
	//
	// false
	AgentLlm *bool `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1578965079000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The debug status.
	//
	// example:
	//
	// DRAFTED
	DebugStatus *string `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	// Indicates whether emotion detection is enabled. This parameter is always false for LLM-based scripts.
	//
	// example:
	//
	// false
	EmotionEnable *bool `json:"EmotionEnable,omitempty" xml:"EmotionEnable,omitempty"`
	// The industry.
	//
	// example:
	//
	// 金融
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Indicates whether the debug version of the script is a draft.
	//
	// example:
	//
	// true
	IsDebugDrafted *bool `json:"IsDebugDrafted,omitempty" xml:"IsDebugDrafted,omitempty"`
	// Indicates whether the script is a draft.
	//
	// example:
	//
	// true
	IsDrafted *bool `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	// Indicates whether the script is for a preset scene.
	//
	// example:
	//
	// false
	IsPreset *bool `json:"IsPreset,omitempty" xml:"IsPreset,omitempty"`
	// Indicates whether long wait is enabled. This parameter is always false for LLM-based scripts.
	//
	// example:
	//
	// false
	LongWaitEnable *bool `json:"LongWaitEnable,omitempty" xml:"LongWaitEnable,omitempty"`
	// Indicates whether mini playback is enabled.
	//
	// example:
	//
	// false
	MiniPlaybackEnable *bool `json:"MiniPlaybackEnable,omitempty" xml:"MiniPlaybackEnable,omitempty"`
	// Indicates whether graceful barge-in is enabled. This parameter is always false for LLM-based scripts.
	//
	// example:
	//
	// false
	NewBargeInEnable *bool `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	// The NLU access type, which is set to Managed for LLM scenarios and is empty for small model scenarios.
	//
	// example:
	//
	// 空或者Managed
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// The NLU engine, which is set to Prompts for LLM scenarios and is empty for small model scenarios.
	//
	// example:
	//
	// 空或者Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// The Function Compute configuration for function calling mode.
	NluProfile *ListScriptsResponseBodyScriptsListNluProfile `json:"NluProfile,omitempty" xml:"NluProfile,omitempty" type:"Struct"`
	// The rejection reason.
	//
	// example:
	//
	// 话术用语不合规
	RejectReason *string `json:"RejectReason,omitempty" xml:"RejectReason,omitempty"`
	// The scene.
	//
	// example:
	//
	// 电销
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The script description.
	//
	// example:
	//
	// 电话销售话术
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// The script ID.
	//
	// example:
	//
	// 8d6a6e41-8093-49af-a9d1-0281878758ac
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The script name.
	//
	// example:
	//
	// 电话销售话术
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The script status.
	//
	// example:
	//
	// DRAFTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1578965079000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The agent ID.
	//
	// example:
	//
	// 1160195
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
}

func (s ListScriptsResponseBodyScriptsList) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBodyScriptsList) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBodyScriptsList) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListScriptsResponseBodyScriptsList) GetAgentLlm() *bool {
	return s.AgentLlm
}

func (s *ListScriptsResponseBodyScriptsList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListScriptsResponseBodyScriptsList) GetDebugStatus() *string {
	return s.DebugStatus
}

func (s *ListScriptsResponseBodyScriptsList) GetEmotionEnable() *bool {
	return s.EmotionEnable
}

func (s *ListScriptsResponseBodyScriptsList) GetIndustry() *string {
	return s.Industry
}

func (s *ListScriptsResponseBodyScriptsList) GetIsDebugDrafted() *bool {
	return s.IsDebugDrafted
}

func (s *ListScriptsResponseBodyScriptsList) GetIsDrafted() *bool {
	return s.IsDrafted
}

func (s *ListScriptsResponseBodyScriptsList) GetIsPreset() *bool {
	return s.IsPreset
}

func (s *ListScriptsResponseBodyScriptsList) GetLongWaitEnable() *bool {
	return s.LongWaitEnable
}

func (s *ListScriptsResponseBodyScriptsList) GetMiniPlaybackEnable() *bool {
	return s.MiniPlaybackEnable
}

func (s *ListScriptsResponseBodyScriptsList) GetNewBargeInEnable() *bool {
	return s.NewBargeInEnable
}

func (s *ListScriptsResponseBodyScriptsList) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *ListScriptsResponseBodyScriptsList) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ListScriptsResponseBodyScriptsList) GetNluProfile() *ListScriptsResponseBodyScriptsListNluProfile {
	return s.NluProfile
}

func (s *ListScriptsResponseBodyScriptsList) GetRejectReason() *string {
	return s.RejectReason
}

func (s *ListScriptsResponseBodyScriptsList) GetScene() *string {
	return s.Scene
}

func (s *ListScriptsResponseBodyScriptsList) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *ListScriptsResponseBodyScriptsList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptsResponseBodyScriptsList) GetScriptName() *string {
	return s.ScriptName
}

func (s *ListScriptsResponseBodyScriptsList) GetStatus() *string {
	return s.Status
}

func (s *ListScriptsResponseBodyScriptsList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListScriptsResponseBodyScriptsList) GetAgentId() *int64 {
	return s.AgentId
}

func (s *ListScriptsResponseBodyScriptsList) SetAgentKey(v string) *ListScriptsResponseBodyScriptsList {
	s.AgentKey = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetAgentLlm(v bool) *ListScriptsResponseBodyScriptsList {
	s.AgentLlm = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetCreateTime(v int64) *ListScriptsResponseBodyScriptsList {
	s.CreateTime = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetDebugStatus(v string) *ListScriptsResponseBodyScriptsList {
	s.DebugStatus = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetEmotionEnable(v bool) *ListScriptsResponseBodyScriptsList {
	s.EmotionEnable = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetIndustry(v string) *ListScriptsResponseBodyScriptsList {
	s.Industry = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetIsDebugDrafted(v bool) *ListScriptsResponseBodyScriptsList {
	s.IsDebugDrafted = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetIsDrafted(v bool) *ListScriptsResponseBodyScriptsList {
	s.IsDrafted = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetIsPreset(v bool) *ListScriptsResponseBodyScriptsList {
	s.IsPreset = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetLongWaitEnable(v bool) *ListScriptsResponseBodyScriptsList {
	s.LongWaitEnable = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetMiniPlaybackEnable(v bool) *ListScriptsResponseBodyScriptsList {
	s.MiniPlaybackEnable = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetNewBargeInEnable(v bool) *ListScriptsResponseBodyScriptsList {
	s.NewBargeInEnable = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetNluAccessType(v string) *ListScriptsResponseBodyScriptsList {
	s.NluAccessType = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetNluEngine(v string) *ListScriptsResponseBodyScriptsList {
	s.NluEngine = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetNluProfile(v *ListScriptsResponseBodyScriptsListNluProfile) *ListScriptsResponseBodyScriptsList {
	s.NluProfile = v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetRejectReason(v string) *ListScriptsResponseBodyScriptsList {
	s.RejectReason = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetScene(v string) *ListScriptsResponseBodyScriptsList {
	s.Scene = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetScriptDescription(v string) *ListScriptsResponseBodyScriptsList {
	s.ScriptDescription = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetScriptId(v string) *ListScriptsResponseBodyScriptsList {
	s.ScriptId = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetScriptName(v string) *ListScriptsResponseBodyScriptsList {
	s.ScriptName = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetStatus(v string) *ListScriptsResponseBodyScriptsList {
	s.Status = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetUpdateTime(v int64) *ListScriptsResponseBodyScriptsList {
	s.UpdateTime = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) SetAgentId(v int64) *ListScriptsResponseBodyScriptsList {
	s.AgentId = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsList) Validate() error {
	if s.NluProfile != nil {
		if err := s.NluProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScriptsResponseBodyScriptsListNluProfile struct {
	// The function name.
	//
	// example:
	//
	// sanfang_test
	FcFunction *string `json:"FcFunction,omitempty" xml:"FcFunction,omitempty"`
	// The URL of the function trigger.
	//
	// example:
	//
	// http://sanfang_test-xxxxxx.cn-shanghai-vpc.fcapp.run
	FcHttpTriggerUrl *string `json:"FcHttpTriggerUrl,omitempty" xml:"FcHttpTriggerUrl,omitempty"`
	// The function region.
	//
	// example:
	//
	// cn-shanghai
	FcRegion *string `json:"FcRegion,omitempty" xml:"FcRegion,omitempty"`
}

func (s ListScriptsResponseBodyScriptsListNluProfile) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBodyScriptsListNluProfile) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBodyScriptsListNluProfile) GetFcFunction() *string {
	return s.FcFunction
}

func (s *ListScriptsResponseBodyScriptsListNluProfile) GetFcHttpTriggerUrl() *string {
	return s.FcHttpTriggerUrl
}

func (s *ListScriptsResponseBodyScriptsListNluProfile) GetFcRegion() *string {
	return s.FcRegion
}

func (s *ListScriptsResponseBodyScriptsListNluProfile) SetFcFunction(v string) *ListScriptsResponseBodyScriptsListNluProfile {
	s.FcFunction = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsListNluProfile) SetFcHttpTriggerUrl(v string) *ListScriptsResponseBodyScriptsListNluProfile {
	s.FcHttpTriggerUrl = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsListNluProfile) SetFcRegion(v string) *ListScriptsResponseBodyScriptsListNluProfile {
	s.FcRegion = &v
	return s
}

func (s *ListScriptsResponseBodyScriptsListNluProfile) Validate() error {
	return dara.Validate(s)
}
