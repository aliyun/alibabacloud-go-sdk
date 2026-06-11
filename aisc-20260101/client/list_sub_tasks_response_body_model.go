// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListSubTasksResponseBodyData) *ListSubTasksResponseBody
	GetData() []*ListSubTasksResponseBodyData
	SetPageInfo(v *ListSubTasksResponseBodyPageInfo) *ListSubTasksResponseBody
	GetPageInfo() *ListSubTasksResponseBodyPageInfo
	SetRequestId(v string) *ListSubTasksResponseBody
	GetRequestId() *string
}

type ListSubTasksResponseBody struct {
	Data     []*ListSubTasksResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageInfo *ListSubTasksResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 9FDE3D6F-26BD-5937-B0E5-8F47962B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSubTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBody) GetData() []*ListSubTasksResponseBodyData {
	return s.Data
}

func (s *ListSubTasksResponseBody) GetPageInfo() *ListSubTasksResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListSubTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubTasksResponseBody) SetData(v []*ListSubTasksResponseBodyData) *ListSubTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListSubTasksResponseBody) SetPageInfo(v *ListSubTasksResponseBodyPageInfo) *ListSubTasksResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListSubTasksResponseBody) SetRequestId(v string) *ListSubTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubTasksResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSubTasksResponseBodyData struct {
	// example:
	//
	// 03d1f08455e965cac0351eaa59256fd9
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// example:
	//
	// 4190063324899520
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// c7acb2f1264e4467887ef8f4c36c44ca1
	Target            *string                                        `json:"Target,omitempty" xml:"Target,omitempty"`
	TaskResultMessage *ListSubTasksResponseBodyDataTaskResultMessage `json:"TaskResultMessage,omitempty" xml:"TaskResultMessage,omitempty" type:"Struct"`
	// example:
	//
	// success
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s ListSubTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyData) GetFileHash() *string {
	return s.FileHash
}

func (s *ListSubTasksResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListSubTasksResponseBodyData) GetTarget() *string {
	return s.Target
}

func (s *ListSubTasksResponseBodyData) GetTaskResultMessage() *ListSubTasksResponseBodyDataTaskResultMessage {
	return s.TaskResultMessage
}

func (s *ListSubTasksResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListSubTasksResponseBodyData) SetFileHash(v string) *ListSubTasksResponseBodyData {
	s.FileHash = &v
	return s
}

func (s *ListSubTasksResponseBodyData) SetId(v int64) *ListSubTasksResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListSubTasksResponseBodyData) SetTarget(v string) *ListSubTasksResponseBodyData {
	s.Target = &v
	return s
}

func (s *ListSubTasksResponseBodyData) SetTaskResultMessage(v *ListSubTasksResponseBodyDataTaskResultMessage) *ListSubTasksResponseBodyData {
	s.TaskResultMessage = v
	return s
}

func (s *ListSubTasksResponseBodyData) SetTaskStatus(v string) *ListSubTasksResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListSubTasksResponseBodyData) Validate() error {
	if s.TaskResultMessage != nil {
		if err := s.TaskResultMessage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSubTasksResponseBodyDataTaskResultMessage struct {
	SkillCheckResult *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult `json:"SkillCheckResult,omitempty" xml:"SkillCheckResult,omitempty" type:"Struct"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessage) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessage) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessage) GetSkillCheckResult() *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult {
	return s.SkillCheckResult
}

func (s *ListSubTasksResponseBodyDataTaskResultMessage) SetSkillCheckResult(v *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult) *ListSubTasksResponseBodyDataTaskResultMessage {
	s.SkillCheckResult = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessage) Validate() error {
	if s.SkillCheckResult != nil {
		if err := s.SkillCheckResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult struct {
	RiskInfo []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty" type:"Repeated"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult) GetRiskInfo() []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo {
	return s.RiskInfo
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult) SetRiskInfo(v []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult {
	s.RiskInfo = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResult) Validate() error {
	if s.RiskInfo != nil {
		for _, item := range s.RiskInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo struct {
	Ext *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt `json:"Ext,omitempty" xml:"Ext,omitempty" type:"Struct"`
	// example:
	//
	// /home/97e55e6af371836f/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// file
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) GetExt() *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt {
	return s.Ext
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) GetPath() *string {
	return s.Path
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) GetResultType() *string {
	return s.ResultType
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) SetExt(v *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo {
	s.Ext = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) SetPath(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo {
	s.Path = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) SetResultType(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo {
	s.ResultType = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfo) Validate() error {
	if s.Ext != nil {
		if err := s.Ext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt struct {
	Config    *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig    `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	Guardrail *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail `json:"Guardrail,omitempty" xml:"Guardrail,omitempty" type:"Struct"`
	Sensitive *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive `json:"Sensitive,omitempty" xml:"Sensitive,omitempty" type:"Struct"`
	Virus     []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus   `json:"Virus,omitempty" xml:"Virus,omitempty" type:"Repeated"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) GetConfig() *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig {
	return s.Config
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) GetGuardrail() *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail {
	return s.Guardrail
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) GetSensitive() *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive {
	return s.Sensitive
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) GetVirus() []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus {
	return s.Virus
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) SetConfig(v *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt {
	s.Config = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) SetGuardrail(v *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt {
	s.Guardrail = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) SetSensitive(v *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt {
	s.Sensitive = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) SetVirus(v []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt {
	s.Virus = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExt) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.Guardrail != nil {
		if err := s.Guardrail.Validate(); err != nil {
			return err
		}
	}
	if s.Sensitive != nil {
		if err := s.Sensitive.Validate(); err != nil {
			return err
		}
	}
	if s.Virus != nil {
		for _, item := range s.Virus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig struct {
	Detail []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig) GetDetail() []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail {
	return s.Detail
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig) SetDetail(v []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig {
	s.Detail = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfig) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail struct {
	// example:
	//
	// allowed-tools: Bash(agent-browser:*)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// The skill configuration allows Bash execution via agent-browser:	- pattern without requiring user confirmation. This enables potentially dangerous command execution through the browser automation CLI.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Dangerous Tools Without Confirmation
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// example:
	//
	// 2555
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) GetContent() *string {
	return s.Content
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) GetDescription() *string {
	return s.Description
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) GetItemName() *string {
	return s.ItemName
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) GetLine() *string {
	return s.Line
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) SetContent(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail {
	s.Content = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) SetDescription(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail {
	s.Description = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) SetItemName(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail {
	s.ItemName = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) SetLine(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail {
	s.Line = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtConfigDetail) Validate() error {
	return dara.Validate(s)
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail struct {
	Detail []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail) GetDetail() []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail {
	return s.Detail
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail) SetDetail(v []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail {
	s.Detail = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail) SetSuggestion(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail {
	s.Suggestion = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrail) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail struct {
	// example:
	//
	// high
	Level  *string                                                                                          `json:"Level,omitempty" xml:"Level,omitempty"`
	Result []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// example:
	//
	// promptAttack
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) GetLevel() *string {
	return s.Level
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) GetResult() []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult {
	return s.Result
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) GetType() *string {
	return s.Type
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) SetLevel(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail {
	s.Level = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) SetResult(v []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail {
	s.Result = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) SetSuggestion(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail {
	s.Suggestion = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) SetType(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail {
	s.Type = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetail) Validate() error {
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

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult struct {
	// example:
	//
	// 25
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// Suspicious attacks.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// attack
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) GetDescription() *string {
	return s.Description
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) GetLabel() *string {
	return s.Label
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) GetLevel() *string {
	return s.Level
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) SetConfidence(v float32) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult {
	s.Confidence = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) SetDescription(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult {
	s.Description = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) SetLabel(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult {
	s.Label = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) SetLevel(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult {
	s.Level = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtGuardrailDetailResult) Validate() error {
	return dara.Validate(s)
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive struct {
	Detail []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive) GetDetail() []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail {
	return s.Detail
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive) SetDetail(v []*ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive {
	s.Detail = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitive) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail struct {
	// example:
	//
	// aliyun_ak_24
	Desc   *string   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Result []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail) GetDesc() *string {
	return s.Desc
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail) GetResult() []*string {
	return s.Result
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail) SetDesc(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail {
	s.Desc = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail) SetResult(v []*string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail {
	s.Result = v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtSensitiveDetail) Validate() error {
	return dara.Validate(s)
}

type ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus struct {
	// example:
	//
	// {}
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// Backdoor
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) GetExt() *string {
	return s.Ext
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) GetScore() *int32 {
	return s.Score
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) GetType() *string {
	return s.Type
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) SetExt(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus {
	s.Ext = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) SetScore(v int32) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus {
	s.Score = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) SetType(v string) *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus {
	s.Type = &v
	return s
}

func (s *ListSubTasksResponseBodyDataTaskResultMessageSkillCheckResultRiskInfoExtVirus) Validate() error {
	return dara.Validate(s)
}

type ListSubTasksResponseBodyPageInfo struct {
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSubTasksResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponseBodyPageInfo) GetCount() *string {
	return s.Count
}

func (s *ListSubTasksResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSubTasksResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubTasksResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSubTasksResponseBodyPageInfo) SetCount(v string) *ListSubTasksResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListSubTasksResponseBodyPageInfo) SetCurrentPage(v int32) *ListSubTasksResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListSubTasksResponseBodyPageInfo) SetPageSize(v int32) *ListSubTasksResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListSubTasksResponseBodyPageInfo) SetTotalCount(v int32) *ListSubTasksResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListSubTasksResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
