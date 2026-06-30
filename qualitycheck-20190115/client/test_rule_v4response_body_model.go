// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestRuleV4ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TestRuleV4ResponseBody
	GetCode() *string
	SetData(v *TestRuleV4ResponseBodyData) *TestRuleV4ResponseBody
	GetData() *TestRuleV4ResponseBodyData
	SetHttpStatusCode(v int32) *TestRuleV4ResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *TestRuleV4ResponseBody
	GetMessage() *string
	SetRequestId(v string) *TestRuleV4ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TestRuleV4ResponseBody
	GetSuccess() *bool
}

type TestRuleV4ResponseBody struct {
	// Result code. 200 indicates success. Other values indicate failure. Callers can use this field to identify the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Full response body.
	Data *TestRuleV4ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error details when an error occurs. Returns successful when the request succeeds.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 96138D8D-XXXX-4E41-XXXX-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Callers can use this field to determine success: true means success; false or null means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestRuleV4ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBody) GetCode() *string {
	return s.Code
}

func (s *TestRuleV4ResponseBody) GetData() *TestRuleV4ResponseBodyData {
	return s.Data
}

func (s *TestRuleV4ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TestRuleV4ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TestRuleV4ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestRuleV4ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestRuleV4ResponseBody) SetCode(v string) *TestRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *TestRuleV4ResponseBody) SetData(v *TestRuleV4ResponseBodyData) *TestRuleV4ResponseBody {
	s.Data = v
	return s
}

func (s *TestRuleV4ResponseBody) SetHttpStatusCode(v int32) *TestRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TestRuleV4ResponseBody) SetMessage(v string) *TestRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *TestRuleV4ResponseBody) SetRequestId(v string) *TestRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestRuleV4ResponseBody) SetSuccess(v bool) *TestRuleV4ResponseBody {
	s.Success = &v
	return s
}

func (s *TestRuleV4ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestRuleV4ResponseBodyData struct {
	// Information about hit check items.
	HitRuleReviewInfoList []*TestRuleV4ResponseBodyDataHitRuleReviewInfoList `json:"HitRuleReviewInfoList,omitempty" xml:"HitRuleReviewInfoList,omitempty" type:"Repeated"`
	// List of hit advanced flow nodes.
	HitTaskFlowList []*TestRuleV4ResponseBodyDataHitTaskFlowList `json:"HitTaskFlowList,omitempty" xml:"HitTaskFlowList,omitempty" type:"Repeated"`
	// Information about rules that were not hit.
	UnhitRuleReviewInfoList []*TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList `json:"UnhitRuleReviewInfoList,omitempty" xml:"UnhitRuleReviewInfoList,omitempty" type:"Repeated"`
}

func (s TestRuleV4ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4ResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyData) GetHitRuleReviewInfoList() []*TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	return s.HitRuleReviewInfoList
}

func (s *TestRuleV4ResponseBodyData) GetHitTaskFlowList() []*TestRuleV4ResponseBodyDataHitTaskFlowList {
	return s.HitTaskFlowList
}

func (s *TestRuleV4ResponseBodyData) GetUnhitRuleReviewInfoList() []*TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList {
	return s.UnhitRuleReviewInfoList
}

func (s *TestRuleV4ResponseBodyData) SetHitRuleReviewInfoList(v []*TestRuleV4ResponseBodyDataHitRuleReviewInfoList) *TestRuleV4ResponseBodyData {
	s.HitRuleReviewInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyData) SetHitTaskFlowList(v []*TestRuleV4ResponseBodyDataHitTaskFlowList) *TestRuleV4ResponseBodyData {
	s.HitTaskFlowList = v
	return s
}

func (s *TestRuleV4ResponseBodyData) SetUnhitRuleReviewInfoList(v []*TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) *TestRuleV4ResponseBodyData {
	s.UnhitRuleReviewInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyData) Validate() error {
	if s.HitRuleReviewInfoList != nil {
		for _, item := range s.HitRuleReviewInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HitTaskFlowList != nil {
		for _, item := range s.HitTaskFlowList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UnhitRuleReviewInfoList != nil {
		for _, item := range s.UnhitRuleReviewInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoList struct {
	// ID of the hit branch.
	//
	// example:
	//
	// 1
	BranchHitId *int64 `json:"BranchHitId,omitempty" xml:"BranchHitId,omitempty"`
	// List of branch information.
	BranchInfoList []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList `json:"BranchInfoList,omitempty" xml:"BranchInfoList,omitempty" type:"Repeated"`
	// Information about hit conditions.
	ConditionHitInfoList []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList `json:"ConditionHitInfoList,omitempty" xml:"ConditionHitInfoList,omitempty" type:"Repeated"`
	// List of conditions.
	ConditionInfoList []*ConditionBasicInfo `json:"ConditionInfoList,omitempty" xml:"ConditionInfoList,omitempty" type:"Repeated"`
	// Judgement node name.
	//
	// example:
	//
	// 判断节点A
	JudgeNodeName *string `json:"JudgeNodeName,omitempty" xml:"JudgeNodeName,omitempty"`
	// Lambda expression. Example: a&\\&b.
	//
	// example:
	//
	// a&&b
	Lambda *string `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	// Whether the rule was hit.
	//
	// example:
	//
	// true
	Matched *bool `json:"Matched,omitempty" xml:"Matched,omitempty"`
	// Node type.
	//
	// example:
	//
	// 0
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// Check item ID.
	//
	// example:
	//
	// 451
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// Rule name.
	//
	// example:
	//
	// 规则A
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Whether scoring applies. Valid values: 1 (no scoring) and 3 (scoring).
	//
	// example:
	//
	// 1
	RuleScoreType *int32 `json:"RuleScoreType,omitempty" xml:"RuleScoreType,omitempty"`
	// Scoring type. Valid values: 0 (add or subtract points on hit) and 1 (one-time score on hit).
	//
	// example:
	//
	// 1
	ScoreNumType *int32 `json:"ScoreNumType,omitempty" xml:"ScoreNumType,omitempty"`
	// Flow ID.
	//
	// example:
	//
	// 1
	TaskFlowId *int64 `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoList) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetBranchHitId() *int64 {
	return s.BranchHitId
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetBranchInfoList() []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	return s.BranchInfoList
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetConditionHitInfoList() []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList {
	return s.ConditionHitInfoList
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetConditionInfoList() []*ConditionBasicInfo {
	return s.ConditionInfoList
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetJudgeNodeName() *string {
	return s.JudgeNodeName
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetLambda() *string {
	return s.Lambda
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetMatched() *bool {
	return s.Matched
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetNodeType() *string {
	return s.NodeType
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetRid() *int64 {
	return s.Rid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetRuleName() *string {
	return s.RuleName
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetRuleScoreType() *int32 {
	return s.RuleScoreType
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetScoreNumType() *int32 {
	return s.ScoreNumType
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) GetTaskFlowId() *int64 {
	return s.TaskFlowId
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetBranchHitId(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.BranchHitId = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetBranchInfoList(v []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.BranchInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetConditionHitInfoList(v []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.ConditionHitInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetConditionInfoList(v []*ConditionBasicInfo) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.ConditionInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetJudgeNodeName(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.JudgeNodeName = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetLambda(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.Lambda = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetMatched(v bool) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.Matched = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetNodeType(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.NodeType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetRid(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.Rid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetRuleName(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.RuleName = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetRuleScoreType(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.RuleScoreType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetScoreNumType(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.ScoreNumType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) SetTaskFlowId(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoList {
	s.TaskFlowId = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoList) Validate() error {
	if s.BranchInfoList != nil {
		for _, item := range s.BranchInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConditionHitInfoList != nil {
		for _, item := range s.ConditionHitInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConditionInfoList != nil {
		for _, item := range s.ConditionInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList struct {
	// Check item type.
	//
	// example:
	//
	// 0
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// Index number.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Lambda expression.
	//
	// example:
	//
	// a&&b
	Lambda *string `json:"Lambda,omitempty" xml:"Lambda,omitempty"`
	// Node name.
	//
	// example:
	//
	// 节点A
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// ID of the next flow node.
	//
	// example:
	//
	// 2
	NextNodeId *int64 `json:"NextNodeId,omitempty" xml:"NextNodeId,omitempty"`
	// Flow node condition.
	Situation *NextNodeSituations `json:"Situation,omitempty" xml:"Situation,omitempty"`
	// List of trigger IDs.
	Triggers []*string `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) GetCheckType() *int32 {
	return s.CheckType
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) GetIndex() *int32 {
	return s.Index
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) GetLambda() *string {
	return s.Lambda
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) GetName() *string {
	return s.Name
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) GetNextNodeId() *int64 {
	return s.NextNodeId
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) GetSituation() *NextNodeSituations {
	return s.Situation
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) GetTriggers() []*string {
	return s.Triggers
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetCheckType(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.CheckType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetIndex(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Index = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetLambda(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Lambda = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetName(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Name = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetNextNodeId(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.NextNodeId = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetSituation(v *NextNodeSituations) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Situation = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) SetTriggers(v []*string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList {
	s.Triggers = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListBranchInfoList) Validate() error {
	if s.Situation != nil {
		if err := s.Situation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList struct {
	// Condition ID.
	Cid []*string `json:"Cid,omitempty" xml:"Cid,omitempty" type:"Repeated"`
	// Key information that was hit. This information appears highlighted on the review page. Examples include keywords matched by a keyword-check operator or category information matched by an agent-model-check operator.
	KeyWords []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords `json:"KeyWords,omitempty" xml:"KeyWords,omitempty" type:"Repeated"`
	// Sentence details for the current hit check item.
	Phrase *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) GetCid() []*string {
	return s.Cid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) GetKeyWords() []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	return s.KeyWords
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) GetPhrase() *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	return s.Phrase
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) SetCid(v []*string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList {
	s.Cid = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) SetKeyWords(v []*TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList {
	s.KeyWords = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) SetPhrase(v *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList {
	s.Phrase = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoList) Validate() error {
	if s.KeyWords != nil {
		for _, item := range s.KeyWords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Phrase != nil {
		if err := s.Phrase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords struct {
	// Condition ID.
	//
	// example:
	//
	// 4
	Cid *string `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// Internal field. Ignore this field.
	//
	// example:
	//
	// 无
	CustomizeCode *string `json:"CustomizeCode,omitempty" xml:"CustomizeCode,omitempty"`
	// Starting character position for highlighting. Index starts at 0. Maximum value is the total number of characters in the sentence minus 1. The character at position from is included in the highlight.
	//
	// example:
	//
	// 1
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// Operator ID.
	//
	// example:
	//
	// 123
	Oid *string `json:"Oid,omitempty" xml:"Oid,omitempty"`
	// Key information matched by the operator.
	//
	// example:
	//
	// 无
	OperatorKey *string `json:"OperatorKey,omitempty" xml:"OperatorKey,omitempty"`
	// Index of this sentence in the full list of sentences. This is the index of the sentence in the dialogues array of the request parameters, starting from 0.
	//
	// example:
	//
	// 13
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// Similar phrase.
	//
	// example:
	//
	// 您好
	SimilarPhrase *string `json:"SimilarPhrase,omitempty" xml:"SimilarPhrase,omitempty"`
	// Internal use only. Ignore this field.
	//
	// example:
	//
	// 无
	Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The position of the character that follows the highlighted keyword. The character at the \\`to\\` position is not included in the highlight. For example, if \\`from\\` is 0 and \\`to\\` is 3, the highlighted keyword consists of the characters at indices 0, 1, and 2. The maximum value is one less than the total number of characters in the sentence.
	//
	// example:
	//
	// 3
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
	// Internal use only. Ignore this field.
	//
	// example:
	//
	// 无
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// Key information matched by the operator. For details, see **Detailed explanation of Val key information*	- in the response parameter description below.
	//
	// example:
	//
	// 你好
	Val *string `json:"Val,omitempty" xml:"Val,omitempty"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetCid() *string {
	return s.Cid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetCustomizeCode() *string {
	return s.CustomizeCode
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetFrom() *int32 {
	return s.From
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetOid() *string {
	return s.Oid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetOperatorKey() *string {
	return s.OperatorKey
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetPid() *int32 {
	return s.Pid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetSimilarPhrase() *string {
	return s.SimilarPhrase
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetTid() *string {
	return s.Tid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetTo() *int32 {
	return s.To
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetUuid() *string {
	return s.Uuid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) GetVal() *string {
	return s.Val
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetCid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Cid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetCustomizeCode(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.CustomizeCode = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetFrom(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.From = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetOid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Oid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetOperatorKey(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.OperatorKey = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetPid(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Pid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetSimilarPhrase(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.SimilarPhrase = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetTid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Tid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetTo(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.To = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetUuid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Uuid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) SetVal(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords {
	s.Val = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListKeyWords) Validate() error {
	return dara.Validate(s)
}

type TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase struct {
	// Start time offset of this sentence relative to the start of the full dialogue, in milliseconds. For example, if the total audio duration is 2 minutes and 10 seconds, and a customer starts speaking at 1 minute and 12 seconds and finishes at 1 minute and 20 seconds, then begin equals 72000 and end equals 80000.
	//
	// example:
	//
	// 72000
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// Start time of this sentence. Example: 2019-11-25 15:37:16.
	//
	// example:
	//
	// 2019-11-25 15:37:16
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// Channel ID.
	//
	// example:
	//
	// 0
	ChannelId *int32 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Internal use only. Ignore this field.
	//
	// example:
	//
	// 忽略
	EmotionFineGrainedValue *int32 `json:"EmotionFineGrainedValue,omitempty" xml:"EmotionFineGrainedValue,omitempty"`
	// Emotion intensity score. Value equals volume in decibels divided by 10. Valid range: [1, 10]. Higher values indicate stronger emotion.
	//
	// example:
	//
	// 7
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// End time offset of this sentence relative to the start of the full dialogue, in milliseconds. For example, if the total audio duration is 2 minutes and 10 seconds, and a customer starts speaking at 1 minute and 12 seconds and finishes at 1 minute and 20 seconds, then begin equals 72000 and end equals 80000.
	//
	// example:
	//
	// 80000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// Hit status. Valid values:
	//
	// - **0**: Not hit
	//
	// - **1**: Hit
	//
	// example:
	//
	// 1
	HitStatus *int32 `json:"HitStatus,omitempty" xml:"HitStatus,omitempty"`
	// Start time of this sentence in hh:mm:ss format.
	//
	// example:
	//
	// 10:00:00
	HourMinSec *string `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	// Role identifier. In offline voice scenarios, roles are limited to agent or customer. In offline text quality check scenarios, this field shows the identity value passed in during upload.
	//
	// example:
	//
	// 客服
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// Index of this sentence in the full list of sentences. This is the index of the sentence in the dialogues array of the request parameters, starting from 0.
	//
	// example:
	//
	// 3
	Pid *int32 `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// Internal use only. Ignore this field.
	//
	// example:
	//
	// 无
	RenterId *int64 `json:"RenterId,omitempty" xml:"RenterId,omitempty"`
	// Speaker role for this sentence. Valid values: agent and customer.
	//
	// example:
	//
	// 客服
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Internal use only. Ignore this field.
	//
	// example:
	//
	// 无
	Sid *int64 `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// Silence duration, in milliseconds.
	//
	// example:
	//
	// 1000
	SilenceDuration *int32 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	// Average speech rate for this sentence, in words per minute.
	//
	// example:
	//
	// 100
	SpeechRate *int32 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// Internal use only. Ignore this field.
	//
	// example:
	//
	// 无
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// Dialogue content.
	//
	// example:
	//
	// 你好，请问有什么可以帮您
	Words *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetBegin() *int64 {
	return s.Begin
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetBeginTime() *string {
	return s.BeginTime
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetChannelId() *int32 {
	return s.ChannelId
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetEmotionFineGrainedValue() *int32 {
	return s.EmotionFineGrainedValue
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetEnd() *int64 {
	return s.End
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetHitStatus() *int32 {
	return s.HitStatus
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetHourMinSec() *string {
	return s.HourMinSec
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetIdentity() *string {
	return s.Identity
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetPid() *int32 {
	return s.Pid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetRenterId() *int64 {
	return s.RenterId
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetRole() *string {
	return s.Role
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetSid() *int64 {
	return s.Sid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetSilenceDuration() *int32 {
	return s.SilenceDuration
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetUuid() *string {
	return s.Uuid
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) GetWords() *string {
	return s.Words
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetBegin(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Begin = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetBeginTime(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.BeginTime = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetChannelId(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.ChannelId = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetEmotionFineGrainedValue(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.EmotionFineGrainedValue = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetEmotionValue(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.EmotionValue = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetEnd(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.End = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetHitStatus(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.HitStatus = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetHourMinSec(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.HourMinSec = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetIdentity(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Identity = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetPid(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Pid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetRenterId(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.RenterId = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetRole(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Role = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetSid(v int64) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Sid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetSilenceDuration(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.SilenceDuration = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetSpeechRate(v int32) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.SpeechRate = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetUuid(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Uuid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) SetWords(v string) *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase {
	s.Words = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitRuleReviewInfoListConditionHitInfoListPhrase) Validate() error {
	return dara.Validate(s)
}

type TestRuleV4ResponseBodyDataHitTaskFlowList struct {
	// Flowchart canvas.
	GraphFlow *TaskGraphFlow `json:"GraphFlow,omitempty" xml:"GraphFlow,omitempty"`
	// Rule ID.
	//
	// example:
	//
	// 1
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// Flowchart type. Deprecated. Default value: 1.
	//
	// example:
	//
	// 无
	TaskFlowType *int32 `json:"TaskFlowType,omitempty" xml:"TaskFlowType,omitempty"`
}

func (s TestRuleV4ResponseBodyDataHitTaskFlowList) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataHitTaskFlowList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) GetGraphFlow() *TaskGraphFlow {
	return s.GraphFlow
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) GetRid() *int64 {
	return s.Rid
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) GetTaskFlowType() *int32 {
	return s.TaskFlowType
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) SetGraphFlow(v *TaskGraphFlow) *TestRuleV4ResponseBodyDataHitTaskFlowList {
	s.GraphFlow = v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) SetRid(v int64) *TestRuleV4ResponseBodyDataHitTaskFlowList {
	s.Rid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) SetTaskFlowType(v int32) *TestRuleV4ResponseBodyDataHitTaskFlowList {
	s.TaskFlowType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataHitTaskFlowList) Validate() error {
	if s.GraphFlow != nil {
		if err := s.GraphFlow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList struct {
	// List of conditions.
	ConditionInfoList []*ConditionBasicInfo `json:"ConditionInfoList,omitempty" xml:"ConditionInfoList,omitempty" type:"Repeated"`
	// Whether the rule was hit.
	//
	// example:
	//
	// true
	Matched *bool `json:"Matched,omitempty" xml:"Matched,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 2
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// Flowchart type. Deprecated. Default value: 1.
	//
	// example:
	//
	// 忽略
	TaskFlowType *int32 `json:"TaskFlowType,omitempty" xml:"TaskFlowType,omitempty"`
}

func (s TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) GoString() string {
	return s.String()
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) GetConditionInfoList() []*ConditionBasicInfo {
	return s.ConditionInfoList
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) GetMatched() *bool {
	return s.Matched
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) GetRid() *int64 {
	return s.Rid
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) GetTaskFlowType() *int32 {
	return s.TaskFlowType
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) SetConditionInfoList(v []*ConditionBasicInfo) *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList {
	s.ConditionInfoList = v
	return s
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) SetMatched(v bool) *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList {
	s.Matched = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) SetRid(v int64) *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList {
	s.Rid = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) SetTaskFlowType(v int32) *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList {
	s.TaskFlowType = &v
	return s
}

func (s *TestRuleV4ResponseBodyDataUnhitRuleReviewInfoList) Validate() error {
	if s.ConditionInfoList != nil {
		for _, item := range s.ConditionInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
