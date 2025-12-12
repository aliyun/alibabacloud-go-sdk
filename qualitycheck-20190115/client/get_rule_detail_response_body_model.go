// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRuleDetailResponseBody
	GetCode() *string
	SetData(v *GetRuleDetailResponseBodyData) *GetRuleDetailResponseBody
	GetData() *GetRuleDetailResponseBodyData
	SetMessage(v string) *GetRuleDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRuleDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRuleDetailResponseBody
	GetSuccess() *bool
}

type GetRuleDetailResponseBody struct {
	// example:
	//
	// 200
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetRuleDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRuleDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRuleDetailResponseBody) GetData() *GetRuleDetailResponseBodyData {
	return s.Data
}

func (s *GetRuleDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRuleDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuleDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRuleDetailResponseBody) SetCode(v string) *GetRuleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetData(v *GetRuleDetailResponseBodyData) *GetRuleDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleDetailResponseBody) SetMessage(v string) *GetRuleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetRequestId(v string) *GetRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleDetailResponseBody) SetSuccess(v bool) *GetRuleDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetRuleDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleDetailResponseBodyData struct {
	Conditions *GetRuleDetailResponseBodyDataConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Struct"`
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rules    *GetRuleDetailResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyData) GetConditions() *GetRuleDetailResponseBodyDataConditions {
	return s.Conditions
}

func (s *GetRuleDetailResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetRuleDetailResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetRuleDetailResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRuleDetailResponseBodyData) GetRules() *GetRuleDetailResponseBodyDataRules {
	return s.Rules
}

func (s *GetRuleDetailResponseBodyData) SetConditions(v *GetRuleDetailResponseBodyDataConditions) *GetRuleDetailResponseBodyData {
	s.Conditions = v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetCount(v int32) *GetRuleDetailResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetPageNumber(v int32) *GetRuleDetailResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetPageSize(v int32) *GetRuleDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRuleDetailResponseBodyData) SetRules(v *GetRuleDetailResponseBodyDataRules) *GetRuleDetailResponseBodyData {
	s.Rules = v
	return s
}

func (s *GetRuleDetailResponseBodyData) Validate() error {
	if s.Conditions != nil {
		if err := s.Conditions.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataConditions struct {
	ConditionBasicInfo []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfo `json:"ConditionBasicInfo,omitempty" xml:"ConditionBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditions) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditions) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditions) GetConditionBasicInfo() []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	return s.ConditionBasicInfo
}

func (s *GetRuleDetailResponseBodyDataConditions) SetConditionBasicInfo(v []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) *GetRuleDetailResponseBodyDataConditions {
	s.ConditionBasicInfo = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditions) Validate() error {
	if s.ConditionBasicInfo != nil {
		for _, item := range s.ConditionBasicInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfo struct {
	CheckRange *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange `json:"CheckRange,omitempty" xml:"CheckRange,omitempty" type:"Struct"`
	// example:
	//
	// 7
	ConditionInfoCid *string `json:"ConditionInfoCid,omitempty" xml:"ConditionInfoCid,omitempty"`
	// example:
	//
	// 7
	OperLambda *string                                                             `json:"OperLambda,omitempty" xml:"OperLambda,omitempty"`
	Operators  *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) GetCheckRange() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	return s.CheckRange
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) GetConditionInfoCid() *string {
	return s.ConditionInfoCid
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) GetOperLambda() *string {
	return s.OperLambda
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) GetOperators() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators {
	return s.Operators
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetCheckRange(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.CheckRange = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetConditionInfoCid(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.ConditionInfoCid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetOperLambda(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.OperLambda = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) SetOperators(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo {
	s.Operators = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfo) Validate() error {
	if s.CheckRange != nil {
		if err := s.CheckRange.Validate(); err != nil {
			return err
		}
	}
	if s.Operators != nil {
		if err := s.Operators.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange struct {
	// example:
	//
	// true
	Absolute *bool                                                                      `json:"Absolute,omitempty" xml:"Absolute,omitempty"`
	Anchor   *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor `json:"Anchor,omitempty" xml:"Anchor,omitempty" type:"Struct"`
	Range    *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange  `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
	Role     *string                                                                    `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) GetAbsolute() *bool {
	return s.Absolute
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) GetAnchor() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	return s.Anchor
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) GetRange() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange {
	return s.Range
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) GetRole() *string {
	return s.Role
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetAbsolute(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Absolute = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetAnchor(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Anchor = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetRange(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Range = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) SetRole(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange {
	s.Role = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRange) Validate() error {
	if s.Anchor != nil {
		if err := s.Anchor.Validate(); err != nil {
			return err
		}
	}
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor struct {
	// example:
	//
	// 7
	AnchorCid *string `json:"AnchorCid,omitempty" xml:"AnchorCid,omitempty"`
	// example:
	//
	// 1
	HitTime *int32 `json:"HitTime,omitempty" xml:"HitTime,omitempty"`
	// example:
	//
	// AFTER
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) GetAnchorCid() *string {
	return s.AnchorCid
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) GetHitTime() *int32 {
	return s.HitTime
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) GetLocation() *string {
	return s.Location
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) SetAnchorCid(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	s.AnchorCid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) SetHitTime(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	s.HitTime = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) SetLocation(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor {
	s.Location = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeAnchor) Validate() error {
	return dara.Validate(s)
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange struct {
	// example:
	//
	// 1
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 10
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) GetFrom() *int32 {
	return s.From
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) GetTo() *int32 {
	return s.To
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) SetFrom(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange {
	s.From = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) SetTo(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange {
	s.To = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoCheckRangeRange) Validate() error {
	return dara.Validate(s)
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators struct {
	OperatorBasicInfo []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo `json:"OperatorBasicInfo,omitempty" xml:"OperatorBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) GetOperatorBasicInfo() []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	return s.OperatorBasicInfo
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) SetOperatorBasicInfo(v []*GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators {
	s.OperatorBasicInfo = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperators) Validate() error {
	if s.OperatorBasicInfo != nil {
		for _, item := range s.OperatorBasicInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo struct {
	// example:
	//
	// 8
	Oid *string `json:"Oid,omitempty" xml:"Oid,omitempty"`
	// example:
	//
	// operator demo
	OperName *string                                                                                   `json:"OperName,omitempty" xml:"OperName,omitempty"`
	Param    *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// example:
	//
	// REGULAR_EXPRESSION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) GetOid() *string {
	return s.Oid
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) GetOperName() *string {
	return s.OperName
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) GetParam() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	return s.Param
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) GetType() *string {
	return s.Type
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetOid(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.Oid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetOperName(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.OperName = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetParam(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.Param = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) SetType(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo {
	s.Type = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfo) Validate() error {
	if s.Param != nil {
		if err := s.Param.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam struct {
	AntModelInfo *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo `json:"AntModelInfo,omitempty" xml:"AntModelInfo,omitempty" type:"Struct"`
	// example:
	//
	// true
	Average *bool `json:"Average,omitempty" xml:"Average,omitempty"`
	// example:
	//
	// DIALOGUE
	BeginType *string `json:"BeginType,omitempty" xml:"BeginType,omitempty"`
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// example:
	//
	// gt
	CompareOperator *string `json:"CompareOperator,omitempty" xml:"CompareOperator,omitempty"`
	// example:
	//
	// true
	ContextChatMatch *bool `json:"ContextChatMatch,omitempty" xml:"ContextChatMatch,omitempty"`
	// example:
	//
	// 1000
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// example:
	//
	// true
	DifferentRole *bool                                                                                             `json:"DifferentRole,omitempty" xml:"DifferentRole,omitempty"`
	Excludes      *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Struct"`
	// example:
	//
	// 3
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// true
	FromEnd *bool `json:"FromEnd,omitempty" xml:"FromEnd,omitempty"`
	// example:
	//
	// 1
	HitTime *int32 `json:"HitTime,omitempty" xml:"HitTime,omitempty"`
	// example:
	//
	// true
	InSentence *bool `json:"InSentence,omitempty" xml:"InSentence,omitempty"`
	// example:
	//
	// 5000
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// true
	KeywordExtension *bool `json:"KeywordExtension,omitempty" xml:"KeywordExtension,omitempty"`
	// example:
	//
	// 3
	KeywordMatchSize *int32 `json:"KeywordMatchSize,omitempty" xml:"KeywordMatchSize,omitempty"`
	// example:
	//
	// 8
	MaxEmotionChangeValue *int32 `json:"MaxEmotionChangeValue,omitempty" xml:"MaxEmotionChangeValue,omitempty"`
	// example:
	//
	// 4
	MinWordSize  *int32                                                                                                `json:"MinWordSize,omitempty" xml:"MinWordSize,omitempty"`
	NotRegex     *string                                                                                               `json:"NotRegex,omitempty" xml:"NotRegex,omitempty"`
	OperKeyWords *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords `json:"OperKeyWords,omitempty" xml:"OperKeyWords,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	Phrase     *string                                                                                             `json:"Phrase,omitempty" xml:"Phrase,omitempty"`
	Pvalues    *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues    `json:"Pvalues,omitempty" xml:"Pvalues,omitempty" type:"Struct"`
	References *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences `json:"References,omitempty" xml:"References,omitempty" type:"Struct"`
	Regex      *string                                                                                             `json:"Regex,omitempty" xml:"Regex,omitempty"`
	// example:
	//
	// 80
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 90
	SimilarityThreshold *float32                                                                                                    `json:"Similarity_threshold,omitempty" xml:"Similarity_threshold,omitempty"`
	SimilarlySentences  *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences `json:"SimilarlySentences,omitempty" xml:"SimilarlySentences,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Target     *int32  `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetRole *string `json:"TargetRole,omitempty" xml:"TargetRole,omitempty"`
	// example:
	//
	// 4
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// 280
	VelocityInMint *int32 `json:"VelocityInMint,omitempty" xml:"VelocityInMint,omitempty"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetAntModelInfo() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo {
	return s.AntModelInfo
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetAverage() *bool {
	return s.Average
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetBeginType() *string {
	return s.BeginType
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetCompareOperator() *string {
	return s.CompareOperator
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetContextChatMatch() *bool {
	return s.ContextChatMatch
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetDifferentRole() *bool {
	return s.DifferentRole
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetExcludes() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes {
	return s.Excludes
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetFrom() *int32 {
	return s.From
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetFromEnd() *bool {
	return s.FromEnd
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetHitTime() *int32 {
	return s.HitTime
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetInSentence() *bool {
	return s.InSentence
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetInterval() *int32 {
	return s.Interval
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetKeywordExtension() *bool {
	return s.KeywordExtension
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetKeywordMatchSize() *int32 {
	return s.KeywordMatchSize
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetMaxEmotionChangeValue() *int32 {
	return s.MaxEmotionChangeValue
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetMinWordSize() *int32 {
	return s.MinWordSize
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetNotRegex() *string {
	return s.NotRegex
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetOperKeyWords() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords {
	return s.OperKeyWords
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetPhrase() *string {
	return s.Phrase
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetPvalues() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues {
	return s.Pvalues
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetReferences() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences {
	return s.References
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetRegex() *string {
	return s.Regex
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetScore() *int32 {
	return s.Score
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetSimilarityThreshold() *float32 {
	return s.SimilarityThreshold
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetSimilarlySentences() *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences {
	return s.SimilarlySentences
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetTarget() *int32 {
	return s.Target
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetTargetRole() *string {
	return s.TargetRole
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetThreshold() *float32 {
	return s.Threshold
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) GetVelocityInMint() *int32 {
	return s.VelocityInMint
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetAntModelInfo(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.AntModelInfo = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetAverage(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Average = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetBeginType(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.BeginType = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetCheckType(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.CheckType = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetCompareOperator(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.CompareOperator = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetContextChatMatch(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.ContextChatMatch = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetDelayTime(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.DelayTime = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetDifferentRole(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.DifferentRole = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetExcludes(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Excludes = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetFrom(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.From = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetFromEnd(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.FromEnd = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetHitTime(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.HitTime = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetInSentence(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.InSentence = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetInterval(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Interval = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetKeywordExtension(v bool) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.KeywordExtension = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetKeywordMatchSize(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.KeywordMatchSize = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetMaxEmotionChangeValue(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.MaxEmotionChangeValue = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetMinWordSize(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.MinWordSize = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetNotRegex(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.NotRegex = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetOperKeyWords(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.OperKeyWords = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetPhrase(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Phrase = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetPvalues(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Pvalues = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetReferences(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.References = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetRegex(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Regex = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetScore(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Score = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetSimilarityThreshold(v float32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.SimilarityThreshold = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetSimilarlySentences(v *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.SimilarlySentences = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetTarget(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Target = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetTargetRole(v string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.TargetRole = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetThreshold(v float32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.Threshold = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) SetVelocityInMint(v int32) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam {
	s.VelocityInMint = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParam) Validate() error {
	if s.AntModelInfo != nil {
		if err := s.AntModelInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Excludes != nil {
		if err := s.Excludes.Validate(); err != nil {
			return err
		}
	}
	if s.OperKeyWords != nil {
		if err := s.OperKeyWords.Validate(); err != nil {
			return err
		}
	}
	if s.Pvalues != nil {
		if err := s.Pvalues.Validate(); err != nil {
			return err
		}
	}
	if s.References != nil {
		if err := s.References.Validate(); err != nil {
			return err
		}
	}
	if s.SimilarlySentences != nil {
		if err := s.SimilarlySentences.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo struct {
	AntModelInfo []*string `json:"AntModelInfo,omitempty" xml:"AntModelInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) GetAntModelInfo() []*string {
	return s.AntModelInfo
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) SetAntModelInfo(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo {
	s.AntModelInfo = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamAntModelInfo) Validate() error {
	return dara.Validate(s)
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes struct {
	Excludes []*string `json:"Excludes,omitempty" xml:"Excludes,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) GetExcludes() []*string {
	return s.Excludes
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) SetExcludes(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes {
	s.Excludes = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamExcludes) Validate() error {
	return dara.Validate(s)
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords struct {
	OperKeyWord []*string `json:"OperKeyWord,omitempty" xml:"OperKeyWord,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) GetOperKeyWord() []*string {
	return s.OperKeyWord
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) SetOperKeyWord(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords {
	s.OperKeyWord = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamOperKeyWords) Validate() error {
	return dara.Validate(s)
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues struct {
	Pvalues []*string `json:"Pvalues,omitempty" xml:"Pvalues,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) GetPvalues() []*string {
	return s.Pvalues
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) SetPvalues(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues {
	s.Pvalues = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamPvalues) Validate() error {
	return dara.Validate(s)
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences struct {
	Reference []*string `json:"Reference,omitempty" xml:"Reference,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) GetReference() []*string {
	return s.Reference
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) SetReference(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences {
	s.Reference = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamReferences) Validate() error {
	return dara.Validate(s)
}

type GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences struct {
	SimilarlySentence []*string `json:"SimilarlySentence,omitempty" xml:"SimilarlySentence,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) GetSimilarlySentence() []*string {
	return s.SimilarlySentence
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) SetSimilarlySentence(v []*string) *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences {
	s.SimilarlySentence = v
	return s
}

func (s *GetRuleDetailResponseBodyDataConditionsConditionBasicInfoOperatorsOperatorBasicInfoParamSimilarlySentences) Validate() error {
	return dara.Validate(s)
}

type GetRuleDetailResponseBodyDataRules struct {
	RuleBasicInfo []*GetRuleDetailResponseBodyDataRulesRuleBasicInfo `json:"RuleBasicInfo,omitempty" xml:"RuleBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRules) GetRuleBasicInfo() []*GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	return s.RuleBasicInfo
}

func (s *GetRuleDetailResponseBodyDataRules) SetRuleBasicInfo(v []*GetRuleDetailResponseBodyDataRulesRuleBasicInfo) *GetRuleDetailResponseBodyDataRules {
	s.RuleBasicInfo = v
	return s
}

func (s *GetRuleDetailResponseBodyDataRules) Validate() error {
	if s.RuleBasicInfo != nil {
		for _, item := range s.RuleBasicInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfo struct {
	BusinessCategories *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories `json:"BusinessCategories,omitempty" xml:"BusinessCategories,omitempty" type:"Struct"`
	// example:
	//
	// 4
	Rid *string `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 7&&!8
	RuleLambda *string                                                  `json:"RuleLambda,omitempty" xml:"RuleLambda,omitempty"`
	Triggers   *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Struct"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) GetBusinessCategories() *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories {
	return s.BusinessCategories
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) GetRid() *string {
	return s.Rid
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) GetRuleLambda() *string {
	return s.RuleLambda
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) GetTriggers() *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers {
	return s.Triggers
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetBusinessCategories(v *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.BusinessCategories = v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetRid(v string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.Rid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetRuleLambda(v string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.RuleLambda = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) SetTriggers(v *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) *GetRuleDetailResponseBodyDataRulesRuleBasicInfo {
	s.Triggers = v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfo) Validate() error {
	if s.BusinessCategories != nil {
		if err := s.BusinessCategories.Validate(); err != nil {
			return err
		}
	}
	if s.Triggers != nil {
		if err := s.Triggers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories struct {
	BusinessCategoryBasicInfo []*GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo `json:"BusinessCategoryBasicInfo,omitempty" xml:"BusinessCategoryBasicInfo,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) GetBusinessCategoryBasicInfo() []*GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	return s.BusinessCategoryBasicInfo
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) SetBusinessCategoryBasicInfo(v []*GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories {
	s.BusinessCategoryBasicInfo = v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategories) Validate() error {
	if s.BusinessCategoryBasicInfo != nil {
		for _, item := range s.BusinessCategoryBasicInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo struct {
	// example:
	//
	// 264971810
	Bid          *int32  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// 1
	ServiceType *int32 `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) GetBid() *int32 {
	return s.Bid
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) GetBusinessName() *string {
	return s.BusinessName
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) GetServiceType() *int32 {
	return s.ServiceType
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) SetBid(v int32) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	s.Bid = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) SetBusinessName(v string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	s.BusinessName = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) SetServiceType(v int32) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo {
	s.ServiceType = &v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoBusinessCategoriesBusinessCategoryBasicInfo) Validate() error {
	return dara.Validate(s)
}

type GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers struct {
	Trigger []*string `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Repeated"`
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) GetTrigger() []*string {
	return s.Trigger
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) SetTrigger(v []*string) *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers {
	s.Trigger = v
	return s
}

func (s *GetRuleDetailResponseBodyDataRulesRuleBasicInfoTriggers) Validate() error {
	return dara.Validate(s)
}
