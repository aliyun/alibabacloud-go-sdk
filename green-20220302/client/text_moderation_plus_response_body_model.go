// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextModerationPlusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TextModerationPlusResponseBody
	GetCode() *int32
	SetData(v *TextModerationPlusResponseBodyData) *TextModerationPlusResponseBody
	GetData() *TextModerationPlusResponseBodyData
	SetMessage(v string) *TextModerationPlusResponseBody
	GetMessage() *string
	SetRequestId(v string) *TextModerationPlusResponseBody
	GetRequestId() *string
}

type TextModerationPlusResponseBody struct {
	// The returned HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation results.
	Data *TextModerationPlusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextModerationPlusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TextModerationPlusResponseBody) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TextModerationPlusResponseBody) GetData() *TextModerationPlusResponseBodyData {
	return s.Data
}

func (s *TextModerationPlusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TextModerationPlusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TextModerationPlusResponseBody) SetCode(v int32) *TextModerationPlusResponseBody {
	s.Code = &v
	return s
}

func (s *TextModerationPlusResponseBody) SetData(v *TextModerationPlusResponseBodyData) *TextModerationPlusResponseBody {
	s.Data = v
	return s
}

func (s *TextModerationPlusResponseBody) SetMessage(v string) *TextModerationPlusResponseBody {
	s.Message = &v
	return s
}

func (s *TextModerationPlusResponseBody) SetRequestId(v string) *TextModerationPlusResponseBody {
	s.RequestId = &v
	return s
}

func (s *TextModerationPlusResponseBody) Validate() error {
	return dara.Validate(s)
}

type TextModerationPlusResponseBodyData struct {
	// The suggestion.
	Advice []*TextModerationPlusResponseBodyDataAdvice `json:"Advice,omitempty" xml:"Advice,omitempty" type:"Repeated"`
	// The level of prompt attack
	//
	// example:
	//
	// none
	AttackLevel *string `json:"AttackLevel,omitempty" xml:"AttackLevel,omitempty"`
	// The result of prompt attack detect
	AttackResult []*TextModerationPlusResponseBodyDataAttackResult `json:"AttackResult,omitempty" xml:"AttackResult,omitempty" type:"Repeated"`
	// The id of data
	//
	// example:
	//
	// text1234
	DataId           *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	ManualTaskId     *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// The results.
	Result []*TextModerationPlusResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The score.
	//
	// example:
	//
	// 1
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The level of sensitivity data
	//
	// example:
	//
	// S0
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// The result of sensitivity data detect
	SensitiveResult   []*TextModerationPlusResponseBodyDataSensitiveResult `json:"SensitiveResult,omitempty" xml:"SensitiveResult,omitempty" type:"Repeated"`
	TranslatedContent *string                                              `json:"TranslatedContent,omitempty" xml:"TranslatedContent,omitempty"`
}

func (s TextModerationPlusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TextModerationPlusResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyData) GetAdvice() []*TextModerationPlusResponseBodyDataAdvice {
	return s.Advice
}

func (s *TextModerationPlusResponseBodyData) GetAttackLevel() *string {
	return s.AttackLevel
}

func (s *TextModerationPlusResponseBodyData) GetAttackResult() []*TextModerationPlusResponseBodyDataAttackResult {
	return s.AttackResult
}

func (s *TextModerationPlusResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *TextModerationPlusResponseBodyData) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *TextModerationPlusResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *TextModerationPlusResponseBodyData) GetResult() []*TextModerationPlusResponseBodyDataResult {
	return s.Result
}

func (s *TextModerationPlusResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *TextModerationPlusResponseBodyData) GetScore() *float32 {
	return s.Score
}

func (s *TextModerationPlusResponseBodyData) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *TextModerationPlusResponseBodyData) GetSensitiveResult() []*TextModerationPlusResponseBodyDataSensitiveResult {
	return s.SensitiveResult
}

func (s *TextModerationPlusResponseBodyData) GetTranslatedContent() *string {
	return s.TranslatedContent
}

func (s *TextModerationPlusResponseBodyData) SetAdvice(v []*TextModerationPlusResponseBodyDataAdvice) *TextModerationPlusResponseBodyData {
	s.Advice = v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetAttackLevel(v string) *TextModerationPlusResponseBodyData {
	s.AttackLevel = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetAttackResult(v []*TextModerationPlusResponseBodyDataAttackResult) *TextModerationPlusResponseBodyData {
	s.AttackResult = v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetDataId(v string) *TextModerationPlusResponseBodyData {
	s.DataId = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetDetectedLanguage(v string) *TextModerationPlusResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetManualTaskId(v string) *TextModerationPlusResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetResult(v []*TextModerationPlusResponseBodyDataResult) *TextModerationPlusResponseBodyData {
	s.Result = v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetRiskLevel(v string) *TextModerationPlusResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetScore(v float32) *TextModerationPlusResponseBodyData {
	s.Score = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetSensitiveLevel(v string) *TextModerationPlusResponseBodyData {
	s.SensitiveLevel = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetSensitiveResult(v []*TextModerationPlusResponseBodyDataSensitiveResult) *TextModerationPlusResponseBodyData {
	s.SensitiveResult = v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetTranslatedContent(v string) *TextModerationPlusResponseBodyData {
	s.TranslatedContent = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type TextModerationPlusResponseBodyDataAdvice struct {
	// The answer.
	//
	// example:
	//
	// XXX
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// Hit Label
	//
	// example:
	//
	// xxx
	HitLabel *string `json:"HitLabel,omitempty" xml:"HitLabel,omitempty"`
	// Hit Library Name
	//
	// example:
	//
	// xxx
	HitLibName *string `json:"HitLibName,omitempty" xml:"HitLibName,omitempty"`
}

func (s TextModerationPlusResponseBodyDataAdvice) String() string {
	return dara.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataAdvice) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataAdvice) GetAnswer() *string {
	return s.Answer
}

func (s *TextModerationPlusResponseBodyDataAdvice) GetHitLabel() *string {
	return s.HitLabel
}

func (s *TextModerationPlusResponseBodyDataAdvice) GetHitLibName() *string {
	return s.HitLibName
}

func (s *TextModerationPlusResponseBodyDataAdvice) SetAnswer(v string) *TextModerationPlusResponseBodyDataAdvice {
	s.Answer = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAdvice) SetHitLabel(v string) *TextModerationPlusResponseBodyDataAdvice {
	s.HitLabel = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAdvice) SetHitLibName(v string) *TextModerationPlusResponseBodyDataAdvice {
	s.HitLibName = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAdvice) Validate() error {
	return dara.Validate(s)
}

type TextModerationPlusResponseBodyDataAttackResult struct {
	// The level of prompt attack
	//
	// example:
	//
	// none
	AttackLevel *string `json:"AttackLevel,omitempty" xml:"AttackLevel,omitempty"`
	// The confidence
	//
	// example:
	//
	// 0
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Description
	//
	// example:
	//
	// safe
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label
	//
	// example:
	//
	// safe
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s TextModerationPlusResponseBodyDataAttackResult) String() string {
	return dara.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataAttackResult) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataAttackResult) GetAttackLevel() *string {
	return s.AttackLevel
}

func (s *TextModerationPlusResponseBodyDataAttackResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *TextModerationPlusResponseBodyDataAttackResult) GetDescription() *string {
	return s.Description
}

func (s *TextModerationPlusResponseBodyDataAttackResult) GetLabel() *string {
	return s.Label
}

func (s *TextModerationPlusResponseBodyDataAttackResult) SetAttackLevel(v string) *TextModerationPlusResponseBodyDataAttackResult {
	s.AttackLevel = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAttackResult) SetConfidence(v float32) *TextModerationPlusResponseBodyDataAttackResult {
	s.Confidence = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAttackResult) SetDescription(v string) *TextModerationPlusResponseBodyDataAttackResult {
	s.Description = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAttackResult) SetLabel(v string) *TextModerationPlusResponseBodyDataAttackResult {
	s.Label = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAttackResult) Validate() error {
	return dara.Validate(s)
}

type TextModerationPlusResponseBodyDataResult struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The custom term hit by the moderated content.
	CustomizedHit []*TextModerationPlusResponseBodyDataResultCustomizedHit `json:"CustomizedHit,omitempty" xml:"CustomizedHit,omitempty" type:"Repeated"`
	// The description of the label.
	//
	// example:
	//
	// none
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The term hit by the moderated content.
	//
	// example:
	//
	// XXX
	RiskWords *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
}

func (s TextModerationPlusResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *TextModerationPlusResponseBodyDataResult) GetCustomizedHit() []*TextModerationPlusResponseBodyDataResultCustomizedHit {
	return s.CustomizedHit
}

func (s *TextModerationPlusResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *TextModerationPlusResponseBodyDataResult) GetLabel() *string {
	return s.Label
}

func (s *TextModerationPlusResponseBodyDataResult) GetRiskWords() *string {
	return s.RiskWords
}

func (s *TextModerationPlusResponseBodyDataResult) SetConfidence(v float32) *TextModerationPlusResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) SetCustomizedHit(v []*TextModerationPlusResponseBodyDataResultCustomizedHit) *TextModerationPlusResponseBodyDataResult {
	s.CustomizedHit = v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) SetDescription(v string) *TextModerationPlusResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) SetLabel(v string) *TextModerationPlusResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) SetRiskWords(v string) *TextModerationPlusResponseBodyDataResult {
	s.RiskWords = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type TextModerationPlusResponseBodyDataResultCustomizedHit struct {
	// The terms that are hit. Multiple terms are separated by commas (,).
	//
	// example:
	//
	// xxx
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// The library name.
	//
	// example:
	//
	// test
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s TextModerationPlusResponseBodyDataResultCustomizedHit) String() string {
	return dara.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataResultCustomizedHit) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataResultCustomizedHit) GetKeyWords() *string {
	return s.KeyWords
}

func (s *TextModerationPlusResponseBodyDataResultCustomizedHit) GetLibName() *string {
	return s.LibName
}

func (s *TextModerationPlusResponseBodyDataResultCustomizedHit) SetKeyWords(v string) *TextModerationPlusResponseBodyDataResultCustomizedHit {
	s.KeyWords = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResultCustomizedHit) SetLibName(v string) *TextModerationPlusResponseBodyDataResultCustomizedHit {
	s.LibName = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResultCustomizedHit) Validate() error {
	return dara.Validate(s)
}

type TextModerationPlusResponseBodyDataSensitiveResult struct {
	// Description
	//
	// example:
	//
	// xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label
	//
	// example:
	//
	// 1234
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sensitive data.
	SensitiveData []*string `json:"SensitiveData,omitempty" xml:"SensitiveData,omitempty" type:"Repeated"`
	// The level of sensitivity data
	//
	// example:
	//
	// S1
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
}

func (s TextModerationPlusResponseBodyDataSensitiveResult) String() string {
	return dara.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataSensitiveResult) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) GetDescription() *string {
	return s.Description
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) GetLabel() *string {
	return s.Label
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) GetSensitiveData() []*string {
	return s.SensitiveData
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) SetDescription(v string) *TextModerationPlusResponseBodyDataSensitiveResult {
	s.Description = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) SetLabel(v string) *TextModerationPlusResponseBodyDataSensitiveResult {
	s.Label = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) SetSensitiveData(v []*string) *TextModerationPlusResponseBodyDataSensitiveResult {
	s.SensitiveData = v
	return s
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) SetSensitiveLevel(v string) *TextModerationPlusResponseBodyDataSensitiveResult {
	s.SensitiveLevel = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) Validate() error {
	return dara.Validate(s)
}
