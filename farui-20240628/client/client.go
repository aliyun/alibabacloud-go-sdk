// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type CreateTextFileRequest struct {
	// example:
	//
	// e9a93201-7e96-4dc1-9678-2832fc132d08
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 1714476549
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TextFileName *string `json:"TextFileName,omitempty" xml:"TextFileName,omitempty"`
	TextFileUrl  *string `json:"TextFileUrl,omitempty" xml:"TextFileUrl,omitempty"`
}

func (s CreateTextFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileRequest) GoString() string {
	return s.String()
}

func (s *CreateTextFileRequest) SetClientToken(v string) *CreateTextFileRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTextFileRequest) SetCreateTime(v string) *CreateTextFileRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateTextFileRequest) SetTextFileName(v string) *CreateTextFileRequest {
	s.TextFileName = &v
	return s
}

func (s *CreateTextFileRequest) SetTextFileUrl(v string) *CreateTextFileRequest {
	s.TextFileUrl = &v
	return s
}

type CreateTextFileAdvanceRequest struct {
	// example:
	//
	// e9a93201-7e96-4dc1-9678-2832fc132d08
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 1714476549
	CreateTime        *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TextFileName      *string   `json:"TextFileName,omitempty" xml:"TextFileName,omitempty"`
	TextFileUrlObject io.Reader `json:"TextFileUrl,omitempty" xml:"TextFileUrl,omitempty"`
}

func (s CreateTextFileAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateTextFileAdvanceRequest) SetClientToken(v string) *CreateTextFileAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetCreateTime(v string) *CreateTextFileAdvanceRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetTextFileName(v string) *CreateTextFileAdvanceRequest {
	s.TextFileName = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetTextFileUrlObject(v io.Reader) *CreateTextFileAdvanceRequest {
	s.TextFileUrlObject = v
	return s
}

type CreateTextFileResponseBody struct {
	// example:
	//
	// Request.Signature.Error
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTextFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 81E6F6D2-8ACB-5BDA-9C7C-4D6268CD9652
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTextFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTextFileResponseBody) SetCode(v string) *CreateTextFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTextFileResponseBody) SetData(v *CreateTextFileResponseBodyData) *CreateTextFileResponseBody {
	s.Data = v
	return s
}

func (s *CreateTextFileResponseBody) SetHttpStatusCode(v int64) *CreateTextFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTextFileResponseBody) SetMessage(v string) *CreateTextFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTextFileResponseBody) SetRequestId(v string) *CreateTextFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTextFileResponseBody) SetSuccess(v bool) *CreateTextFileResponseBody {
	s.Success = &v
	return s
}

type CreateTextFileResponseBodyData struct {
	// example:
	//
	// 36d6447d277c4a1c9fd0def1d16341f1
	TextFileId   *string `json:"TextFileId,omitempty" xml:"TextFileId,omitempty"`
	TextFileName *string `json:"TextFileName,omitempty" xml:"TextFileName,omitempty"`
	TextFileUrl  *string `json:"TextFileUrl,omitempty" xml:"TextFileUrl,omitempty"`
}

func (s CreateTextFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTextFileResponseBodyData) SetTextFileId(v string) *CreateTextFileResponseBodyData {
	s.TextFileId = &v
	return s
}

func (s *CreateTextFileResponseBodyData) SetTextFileName(v string) *CreateTextFileResponseBodyData {
	s.TextFileName = &v
	return s
}

func (s *CreateTextFileResponseBodyData) SetTextFileUrl(v string) *CreateTextFileResponseBodyData {
	s.TextFileUrl = &v
	return s
}

type CreateTextFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTextFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTextFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileResponse) GoString() string {
	return s.String()
}

func (s *CreateTextFileResponse) SetHeaders(v map[string]*string) *CreateTextFileResponse {
	s.Headers = v
	return s
}

func (s *CreateTextFileResponse) SetStatusCode(v int32) *CreateTextFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTextFileResponse) SetBody(v *CreateTextFileResponseBody) *CreateTextFileResponse {
	s.Body = v
	return s
}

type RunContractResultGenerationRequest struct {
	// example:
	//
	// farui
	AppId     *string                                      `json:"appId,omitempty" xml:"appId,omitempty"`
	Assistant *RunContractResultGenerationRequestAssistant `json:"assistant,omitempty" xml:"assistant,omitempty" type:"Struct"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s RunContractResultGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequest) SetAppId(v string) *RunContractResultGenerationRequest {
	s.AppId = &v
	return s
}

func (s *RunContractResultGenerationRequest) SetAssistant(v *RunContractResultGenerationRequestAssistant) *RunContractResultGenerationRequest {
	s.Assistant = v
	return s
}

func (s *RunContractResultGenerationRequest) SetStream(v bool) *RunContractResultGenerationRequest {
	s.Stream = &v
	return s
}

type RunContractResultGenerationRequestAssistant struct {
	MetaData *RunContractResultGenerationRequestAssistantMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
	// example:
	//
	// contract_examime
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RunContractResultGenerationRequestAssistant) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistant) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistant) SetMetaData(v *RunContractResultGenerationRequestAssistantMetaData) *RunContractResultGenerationRequestAssistant {
	s.MetaData = v
	return s
}

func (s *RunContractResultGenerationRequestAssistant) SetType(v string) *RunContractResultGenerationRequestAssistant {
	s.Type = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistant) SetVersion(v string) *RunContractResultGenerationRequestAssistant {
	s.Version = &v
	return s
}

type RunContractResultGenerationRequestAssistantMetaData struct {
	CustomRuleConfig *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig `json:"customRuleConfig,omitempty" xml:"customRuleConfig,omitempty" type:"Struct"`
	// example:
	//
	// 9a6b1ba60d9944249363ec3cc1529b7b
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 1
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// b265b416-ca1f-425d-9340-c968f39624e1
	RuleTaskId *string                                                     `json:"ruleTaskId,omitempty" xml:"ruleTaskId,omitempty"`
	Rules      []*RunContractResultGenerationRequestAssistantMetaDataRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s RunContractResultGenerationRequestAssistantMetaData) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistantMetaData) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetCustomRuleConfig(v *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) *RunContractResultGenerationRequestAssistantMetaData {
	s.CustomRuleConfig = v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetFileId(v string) *RunContractResultGenerationRequestAssistantMetaData {
	s.FileId = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetPosition(v string) *RunContractResultGenerationRequestAssistantMetaData {
	s.Position = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetRuleTaskId(v string) *RunContractResultGenerationRequestAssistantMetaData {
	s.RuleTaskId = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaData) SetRules(v []*RunContractResultGenerationRequestAssistantMetaDataRules) *RunContractResultGenerationRequestAssistantMetaData {
	s.Rules = v
	return s
}

type RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig struct {
	CustomRules []*RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules `json:"customRules,omitempty" xml:"customRules,omitempty" type:"Repeated"`
}

func (s RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig) SetCustomRules(v []*RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfig {
	s.CustomRules = v
	return s
}

type RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules struct {
	// example:
	//
	// high
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	RuleDesc  *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty"`
	RuleTitle *string `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
}

func (s RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) SetRiskLevel(v string) *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules {
	s.RiskLevel = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) SetRuleDesc(v string) *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules {
	s.RuleDesc = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules) SetRuleTitle(v string) *RunContractResultGenerationRequestAssistantMetaDataCustomRuleConfigCustomRules {
	s.RuleTitle = &v
	return s
}

type RunContractResultGenerationRequestAssistantMetaDataRules struct {
	// example:
	//
	// medium
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// example:
	//
	// 2.1
	RuleSequence *string `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	RuleTag      *string `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	RuleTitle    *string `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
}

func (s RunContractResultGenerationRequestAssistantMetaDataRules) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationRequestAssistantMetaDataRules) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) SetRiskLevel(v string) *RunContractResultGenerationRequestAssistantMetaDataRules {
	s.RiskLevel = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) SetRuleSequence(v string) *RunContractResultGenerationRequestAssistantMetaDataRules {
	s.RuleSequence = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) SetRuleTag(v string) *RunContractResultGenerationRequestAssistantMetaDataRules {
	s.RuleTag = &v
	return s
}

func (s *RunContractResultGenerationRequestAssistantMetaDataRules) SetRuleTitle(v string) *RunContractResultGenerationRequestAssistantMetaDataRules {
	s.RuleTitle = &v
	return s
}

type RunContractResultGenerationShrinkRequest struct {
	// example:
	//
	// farui
	AppId           *string `json:"appId,omitempty" xml:"appId,omitempty"`
	AssistantShrink *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s RunContractResultGenerationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationShrinkRequest) SetAppId(v string) *RunContractResultGenerationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunContractResultGenerationShrinkRequest) SetAssistantShrink(v string) *RunContractResultGenerationShrinkRequest {
	s.AssistantShrink = &v
	return s
}

func (s *RunContractResultGenerationShrinkRequest) SetStream(v bool) *RunContractResultGenerationShrinkRequest {
	s.Stream = &v
	return s
}

type RunContractResultGenerationResponseBody struct {
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Message *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Output  *RunContractResultGenerationResponseBodyOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// example:
	//
	// 744419D0-671A-5997-9840-E8AE48356194
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Usage   *RunContractResultGenerationResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
}

func (s RunContractResultGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBody) SetCode(v string) *RunContractResultGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetMessage(v string) *RunContractResultGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetOutput(v *RunContractResultGenerationResponseBodyOutput) *RunContractResultGenerationResponseBody {
	s.Output = v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetRequestId(v string) *RunContractResultGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetSuccess(v bool) *RunContractResultGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetUsage(v *RunContractResultGenerationResponseBodyUsage) *RunContractResultGenerationResponseBody {
	s.Usage = v
	return s
}

func (s *RunContractResultGenerationResponseBody) SetHttpStatusCode(v string) *RunContractResultGenerationResponseBody {
	s.HttpStatusCode = &v
	return s
}

type RunContractResultGenerationResponseBodyOutput struct {
	Result *RunContractResultGenerationResponseBodyOutputResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// eaa56e1e-e205-4f5e-926e-5e2269ae7f68
	ResultTaskId *string `json:"resultTaskId,omitempty" xml:"resultTaskId,omitempty"`
}

func (s RunContractResultGenerationResponseBodyOutput) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationResponseBodyOutput) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBodyOutput) SetResult(v *RunContractResultGenerationResponseBodyOutputResult) *RunContractResultGenerationResponseBodyOutput {
	s.Result = v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutput) SetResultTaskId(v string) *RunContractResultGenerationResponseBodyOutput {
	s.ResultTaskId = &v
	return s
}

type RunContractResultGenerationResponseBodyOutputResult struct {
	ExamineBrief  *string `json:"examineBrief,omitempty" xml:"examineBrief,omitempty"`
	ExamineResult *string `json:"examineResult,omitempty" xml:"examineResult,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// example:
	//
	// 1.1
	RuleSequence *string                                                        `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	RuleTag      *string                                                        `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	RuleTitle    *string                                                        `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
	SubRisks     []*RunContractResultGenerationResponseBodyOutputResultSubRisks `json:"subRisks,omitempty" xml:"subRisks,omitempty" type:"Repeated"`
}

func (s RunContractResultGenerationResponseBodyOutputResult) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationResponseBodyOutputResult) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetExamineBrief(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.ExamineBrief = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetExamineResult(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.ExamineResult = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetRiskLevel(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.RiskLevel = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetRuleSequence(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.RuleSequence = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetRuleTag(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.RuleTag = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetRuleTitle(v string) *RunContractResultGenerationResponseBodyOutputResult {
	s.RuleTitle = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResult) SetSubRisks(v []*RunContractResultGenerationResponseBodyOutputResultSubRisks) *RunContractResultGenerationResponseBodyOutputResult {
	s.SubRisks = v
	return s
}

type RunContractResultGenerationResponseBodyOutputResultSubRisks struct {
	OriginalContent *string `json:"originalContent,omitempty" xml:"originalContent,omitempty"`
	ResultContent   *string `json:"resultContent,omitempty" xml:"resultContent,omitempty"`
	ResultType      *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
	RiskBrief       *string `json:"riskBrief,omitempty" xml:"riskBrief,omitempty"`
	RiskClause      *string `json:"riskClause,omitempty" xml:"riskClause,omitempty"`
	RiskExplain     *string `json:"riskExplain,omitempty" xml:"riskExplain,omitempty"`
}

func (s RunContractResultGenerationResponseBodyOutputResultSubRisks) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationResponseBodyOutputResultSubRisks) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetOriginalContent(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.OriginalContent = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetResultContent(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.ResultContent = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetResultType(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.ResultType = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetRiskBrief(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.RiskBrief = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetRiskClause(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.RiskClause = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyOutputResultSubRisks) SetRiskExplain(v string) *RunContractResultGenerationResponseBodyOutputResultSubRisks {
	s.RiskExplain = &v
	return s
}

type RunContractResultGenerationResponseBodyUsage struct {
	// example:
	//
	// 5
	Input *int64 `json:"input,omitempty" xml:"input,omitempty"`
	// example:
	//
	// page
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s RunContractResultGenerationResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponseBodyUsage) SetInput(v int64) *RunContractResultGenerationResponseBodyUsage {
	s.Input = &v
	return s
}

func (s *RunContractResultGenerationResponseBodyUsage) SetUnit(v string) *RunContractResultGenerationResponseBodyUsage {
	s.Unit = &v
	return s
}

type RunContractResultGenerationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunContractResultGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunContractResultGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunContractResultGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponse) SetHeaders(v map[string]*string) *RunContractResultGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunContractResultGenerationResponse) SetStatusCode(v int32) *RunContractResultGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunContractResultGenerationResponse) SetBody(v *RunContractResultGenerationResponseBody) *RunContractResultGenerationResponse {
	s.Body = v
	return s
}

type RunContractRuleGenerationRequest struct {
	// example:
	//
	// farui
	AppId     *string                                    `json:"appId,omitempty" xml:"appId,omitempty"`
	Assistant *RunContractRuleGenerationRequestAssistant `json:"assistant,omitempty" xml:"assistant,omitempty" type:"Struct"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s RunContractRuleGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunContractRuleGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationRequest) SetAppId(v string) *RunContractRuleGenerationRequest {
	s.AppId = &v
	return s
}

func (s *RunContractRuleGenerationRequest) SetAssistant(v *RunContractRuleGenerationRequestAssistant) *RunContractRuleGenerationRequest {
	s.Assistant = v
	return s
}

func (s *RunContractRuleGenerationRequest) SetStream(v bool) *RunContractRuleGenerationRequest {
	s.Stream = &v
	return s
}

type RunContractRuleGenerationRequestAssistant struct {
	MetaData *RunContractRuleGenerationRequestAssistantMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" type:"Struct"`
	// example:
	//
	// contract_examime
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RunContractRuleGenerationRequestAssistant) String() string {
	return tea.Prettify(s)
}

func (s RunContractRuleGenerationRequestAssistant) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationRequestAssistant) SetMetaData(v *RunContractRuleGenerationRequestAssistantMetaData) *RunContractRuleGenerationRequestAssistant {
	s.MetaData = v
	return s
}

func (s *RunContractRuleGenerationRequestAssistant) SetType(v string) *RunContractRuleGenerationRequestAssistant {
	s.Type = &v
	return s
}

func (s *RunContractRuleGenerationRequestAssistant) SetVersion(v string) *RunContractRuleGenerationRequestAssistant {
	s.Version = &v
	return s
}

type RunContractRuleGenerationRequestAssistantMetaData struct {
	// example:
	//
	// 9a6b1ba60d9944249363ec3cc1529b7b
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 1
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
}

func (s RunContractRuleGenerationRequestAssistantMetaData) String() string {
	return tea.Prettify(s)
}

func (s RunContractRuleGenerationRequestAssistantMetaData) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationRequestAssistantMetaData) SetFileId(v string) *RunContractRuleGenerationRequestAssistantMetaData {
	s.FileId = &v
	return s
}

func (s *RunContractRuleGenerationRequestAssistantMetaData) SetPosition(v string) *RunContractRuleGenerationRequestAssistantMetaData {
	s.Position = &v
	return s
}

type RunContractRuleGenerationShrinkRequest struct {
	// example:
	//
	// farui
	AppId           *string `json:"appId,omitempty" xml:"appId,omitempty"`
	AssistantShrink *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s RunContractRuleGenerationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunContractRuleGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationShrinkRequest) SetAppId(v string) *RunContractRuleGenerationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunContractRuleGenerationShrinkRequest) SetAssistantShrink(v string) *RunContractRuleGenerationShrinkRequest {
	s.AssistantShrink = &v
	return s
}

func (s *RunContractRuleGenerationShrinkRequest) SetStream(v bool) *RunContractRuleGenerationShrinkRequest {
	s.Stream = &v
	return s
}

type RunContractRuleGenerationResponseBody struct {
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Message *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Output  *RunContractRuleGenerationResponseBodyOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// example:
	//
	// 744419D0-671A-5997-9840-E8AE48356194
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Usage   *RunContractRuleGenerationResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
}

func (s RunContractRuleGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunContractRuleGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponseBody) SetCode(v string) *RunContractRuleGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetMessage(v string) *RunContractRuleGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetOutput(v *RunContractRuleGenerationResponseBodyOutput) *RunContractRuleGenerationResponseBody {
	s.Output = v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetRequestId(v string) *RunContractRuleGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetSuccess(v bool) *RunContractRuleGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetUsage(v *RunContractRuleGenerationResponseBodyUsage) *RunContractRuleGenerationResponseBody {
	s.Usage = v
	return s
}

func (s *RunContractRuleGenerationResponseBody) SetHttpStatusCode(v int32) *RunContractRuleGenerationResponseBody {
	s.HttpStatusCode = &v
	return s
}

type RunContractRuleGenerationResponseBodyOutput struct {
	// example:
	//
	// b265b416-ca1f-425d-9340-c968f39624e9
	RuleTaskId *string                                             `json:"ruleTaskId,omitempty" xml:"ruleTaskId,omitempty"`
	Rules      []*RunContractRuleGenerationResponseBodyOutputRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s RunContractRuleGenerationResponseBodyOutput) String() string {
	return tea.Prettify(s)
}

func (s RunContractRuleGenerationResponseBodyOutput) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponseBodyOutput) SetRuleTaskId(v string) *RunContractRuleGenerationResponseBodyOutput {
	s.RuleTaskId = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutput) SetRules(v []*RunContractRuleGenerationResponseBodyOutputRules) *RunContractRuleGenerationResponseBodyOutput {
	s.Rules = v
	return s
}

type RunContractRuleGenerationResponseBodyOutputRules struct {
	// example:
	//
	// medium
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// example:
	//
	// 1.1
	RuleSequence *string `json:"ruleSequence,omitempty" xml:"ruleSequence,omitempty"`
	RuleTag      *string `json:"ruleTag,omitempty" xml:"ruleTag,omitempty"`
	RuleTitle    *string `json:"ruleTitle,omitempty" xml:"ruleTitle,omitempty"`
}

func (s RunContractRuleGenerationResponseBodyOutputRules) String() string {
	return tea.Prettify(s)
}

func (s RunContractRuleGenerationResponseBodyOutputRules) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) SetRiskLevel(v string) *RunContractRuleGenerationResponseBodyOutputRules {
	s.RiskLevel = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) SetRuleSequence(v string) *RunContractRuleGenerationResponseBodyOutputRules {
	s.RuleSequence = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) SetRuleTag(v string) *RunContractRuleGenerationResponseBodyOutputRules {
	s.RuleTag = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyOutputRules) SetRuleTitle(v string) *RunContractRuleGenerationResponseBodyOutputRules {
	s.RuleTitle = &v
	return s
}

type RunContractRuleGenerationResponseBodyUsage struct {
	// example:
	//
	// 5
	Input *int64 `json:"input,omitempty" xml:"input,omitempty"`
	// example:
	//
	// page
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s RunContractRuleGenerationResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s RunContractRuleGenerationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponseBodyUsage) SetInput(v int64) *RunContractRuleGenerationResponseBodyUsage {
	s.Input = &v
	return s
}

func (s *RunContractRuleGenerationResponseBodyUsage) SetUnit(v string) *RunContractRuleGenerationResponseBodyUsage {
	s.Unit = &v
	return s
}

type RunContractRuleGenerationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunContractRuleGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunContractRuleGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunContractRuleGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationResponse) SetHeaders(v map[string]*string) *RunContractRuleGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunContractRuleGenerationResponse) SetStatusCode(v int32) *RunContractRuleGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunContractRuleGenerationResponse) SetBody(v *RunContractRuleGenerationResponseBody) *RunContractRuleGenerationResponse {
	s.Body = v
	return s
}

type RunLegalAdviceConsultationRequest struct {
	// example:
	//
	// farui
	AppId     *string                                     `json:"appId,omitempty" xml:"appId,omitempty"`
	Assistant *RunLegalAdviceConsultationRequestAssistant `json:"assistant,omitempty" xml:"assistant,omitempty" type:"Struct"`
	// example:
	//
	// true
	Stream *bool                                    `json:"stream,omitempty" xml:"stream,omitempty"`
	Thread *RunLegalAdviceConsultationRequestThread `json:"thread,omitempty" xml:"thread,omitempty" type:"Struct"`
}

func (s RunLegalAdviceConsultationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationRequest) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequest) SetAppId(v string) *RunLegalAdviceConsultationRequest {
	s.AppId = &v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetAssistant(v *RunLegalAdviceConsultationRequestAssistant) *RunLegalAdviceConsultationRequest {
	s.Assistant = v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetStream(v bool) *RunLegalAdviceConsultationRequest {
	s.Stream = &v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetThread(v *RunLegalAdviceConsultationRequestThread) *RunLegalAdviceConsultationRequest {
	s.Thread = v
	return s
}

type RunLegalAdviceConsultationRequestAssistant struct {
	// example:
	//
	// assitant_abc_123
	Id       *string            `json:"id,omitempty" xml:"id,omitempty"`
	MetaData map[string]*string `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// example:
	//
	// legal_advice_consult
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RunLegalAdviceConsultationRequestAssistant) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestAssistant) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetId(v string) *RunLegalAdviceConsultationRequestAssistant {
	s.Id = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetMetaData(v map[string]*string) *RunLegalAdviceConsultationRequestAssistant {
	s.MetaData = v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetType(v string) *RunLegalAdviceConsultationRequestAssistant {
	s.Type = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetVersion(v string) *RunLegalAdviceConsultationRequestAssistant {
	s.Version = &v
	return s
}

type RunLegalAdviceConsultationRequestThread struct {
	Messages []*RunLegalAdviceConsultationRequestThreadMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s RunLegalAdviceConsultationRequestThread) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestThread) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestThread) SetMessages(v []*RunLegalAdviceConsultationRequestThreadMessages) *RunLegalAdviceConsultationRequestThread {
	s.Messages = v
	return s
}

type RunLegalAdviceConsultationRequestThreadMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunLegalAdviceConsultationRequestThreadMessages) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestThreadMessages) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestThreadMessages) SetContent(v string) *RunLegalAdviceConsultationRequestThreadMessages {
	s.Content = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestThreadMessages) SetRole(v string) *RunLegalAdviceConsultationRequestThreadMessages {
	s.Role = &v
	return s
}

type RunLegalAdviceConsultationShrinkRequest struct {
	// example:
	//
	// farui
	AppId           *string `json:"appId,omitempty" xml:"appId,omitempty"`
	AssistantShrink *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// true
	Stream       *bool   `json:"stream,omitempty" xml:"stream,omitempty"`
	ThreadShrink *string `json:"thread,omitempty" xml:"thread,omitempty"`
}

func (s RunLegalAdviceConsultationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetAppId(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetAssistantShrink(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.AssistantShrink = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetStream(v bool) *RunLegalAdviceConsultationShrinkRequest {
	s.Stream = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetThreadShrink(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.ThreadShrink = &v
	return s
}

type RunLegalAdviceConsultationResponseBody struct {
	// example:
	//
	// Request.Signature.Error
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 744419D0-671A-5997-9840-E8AE48356194
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseMarkdown *string `json:"ResponseMarkdown,omitempty" xml:"ResponseMarkdown,omitempty"`
	// example:
	//
	// 1
	Round  *int32  `json:"Round,omitempty" xml:"Round,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Usage   *RunLegalAdviceConsultationResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
}

func (s RunLegalAdviceConsultationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationResponseBody) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationResponseBody) SetCode(v string) *RunLegalAdviceConsultationResponseBody {
	s.Code = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetMessage(v string) *RunLegalAdviceConsultationResponseBody {
	s.Message = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetRequestId(v string) *RunLegalAdviceConsultationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetResponseMarkdown(v string) *RunLegalAdviceConsultationResponseBody {
	s.ResponseMarkdown = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetRound(v int32) *RunLegalAdviceConsultationResponseBody {
	s.Round = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetStatus(v string) *RunLegalAdviceConsultationResponseBody {
	s.Status = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetSuccess(v bool) *RunLegalAdviceConsultationResponseBody {
	s.Success = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetUsage(v *RunLegalAdviceConsultationResponseBodyUsage) *RunLegalAdviceConsultationResponseBody {
	s.Usage = v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetHttpStatusCode(v string) *RunLegalAdviceConsultationResponseBody {
	s.HttpStatusCode = &v
	return s
}

type RunLegalAdviceConsultationResponseBodyUsage struct {
	// example:
	//
	// 500
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 700
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 1200
	TotalTokens *int32 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunLegalAdviceConsultationResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) SetInputTokens(v int32) *RunLegalAdviceConsultationResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) SetOutputTokens(v int32) *RunLegalAdviceConsultationResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) SetTotalTokens(v int32) *RunLegalAdviceConsultationResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

type RunLegalAdviceConsultationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunLegalAdviceConsultationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunLegalAdviceConsultationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationResponse) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationResponse) SetHeaders(v map[string]*string) *RunLegalAdviceConsultationResponse {
	s.Headers = v
	return s
}

func (s *RunLegalAdviceConsultationResponse) SetStatusCode(v int32) *RunLegalAdviceConsultationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunLegalAdviceConsultationResponse) SetBody(v *RunLegalAdviceConsultationResponseBody) *RunLegalAdviceConsultationResponse {
	s.Body = v
	return s
}

type RunSearchCaseFullTextRequest struct {
	// example:
	//
	// farui
	AppId           *string                                      `json:"appId,omitempty" xml:"appId,omitempty"`
	FilterCondition *RunSearchCaseFullTextRequestFilterCondition `json:"filterCondition,omitempty" xml:"filterCondition,omitempty" type:"Struct"`
	// This parameter is required.
	PageParam *RunSearchCaseFullTextRequestPageParam `json:"pageParam,omitempty" xml:"pageParam,omitempty" type:"Struct"`
	// This parameter is required.
	Query               *string                             `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywords       []*string                           `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty" type:"Repeated"`
	ReferLevel          *string                             `json:"referLevel,omitempty" xml:"referLevel,omitempty"`
	SortKeyAndDirection map[string]*string                  `json:"sortKeyAndDirection,omitempty" xml:"sortKeyAndDirection,omitempty"`
	Thread              *RunSearchCaseFullTextRequestThread `json:"thread,omitempty" xml:"thread,omitempty" type:"Struct"`
}

func (s RunSearchCaseFullTextRequest) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextRequest) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequest) SetAppId(v string) *RunSearchCaseFullTextRequest {
	s.AppId = &v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetFilterCondition(v *RunSearchCaseFullTextRequestFilterCondition) *RunSearchCaseFullTextRequest {
	s.FilterCondition = v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetPageParam(v *RunSearchCaseFullTextRequestPageParam) *RunSearchCaseFullTextRequest {
	s.PageParam = v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetQuery(v string) *RunSearchCaseFullTextRequest {
	s.Query = &v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetQueryKeywords(v []*string) *RunSearchCaseFullTextRequest {
	s.QueryKeywords = v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetReferLevel(v string) *RunSearchCaseFullTextRequest {
	s.ReferLevel = &v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetSortKeyAndDirection(v map[string]*string) *RunSearchCaseFullTextRequest {
	s.SortKeyAndDirection = v
	return s
}

func (s *RunSearchCaseFullTextRequest) SetThread(v *RunSearchCaseFullTextRequestThread) *RunSearchCaseFullTextRequest {
	s.Thread = v
	return s
}

type RunSearchCaseFullTextRequestFilterCondition struct {
	CaseNo    *string `json:"caseNo,omitempty" xml:"caseNo,omitempty"`
	CaseTitle *string `json:"caseTitle,omitempty" xml:"caseTitle,omitempty"`
}

func (s RunSearchCaseFullTextRequestFilterCondition) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextRequestFilterCondition) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequestFilterCondition) SetCaseNo(v string) *RunSearchCaseFullTextRequestFilterCondition {
	s.CaseNo = &v
	return s
}

func (s *RunSearchCaseFullTextRequestFilterCondition) SetCaseTitle(v string) *RunSearchCaseFullTextRequestFilterCondition {
	s.CaseTitle = &v
	return s
}

type RunSearchCaseFullTextRequestPageParam struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s RunSearchCaseFullTextRequestPageParam) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextRequestPageParam) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequestPageParam) SetPageNumber(v int32) *RunSearchCaseFullTextRequestPageParam {
	s.PageNumber = &v
	return s
}

func (s *RunSearchCaseFullTextRequestPageParam) SetPageSize(v int32) *RunSearchCaseFullTextRequestPageParam {
	s.PageSize = &v
	return s
}

type RunSearchCaseFullTextRequestThread struct {
	Messages []*RunSearchCaseFullTextRequestThreadMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s RunSearchCaseFullTextRequestThread) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextRequestThread) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequestThread) SetMessages(v []*RunSearchCaseFullTextRequestThreadMessages) *RunSearchCaseFullTextRequestThread {
	s.Messages = v
	return s
}

type RunSearchCaseFullTextRequestThreadMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunSearchCaseFullTextRequestThreadMessages) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextRequestThreadMessages) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextRequestThreadMessages) SetContent(v string) *RunSearchCaseFullTextRequestThreadMessages {
	s.Content = &v
	return s
}

func (s *RunSearchCaseFullTextRequestThreadMessages) SetRole(v string) *RunSearchCaseFullTextRequestThreadMessages {
	s.Role = &v
	return s
}

type RunSearchCaseFullTextShrinkRequest struct {
	// example:
	//
	// farui
	AppId                 *string `json:"appId,omitempty" xml:"appId,omitempty"`
	FilterConditionShrink *string `json:"filterCondition,omitempty" xml:"filterCondition,omitempty"`
	// This parameter is required.
	PageParamShrink *string `json:"pageParam,omitempty" xml:"pageParam,omitempty"`
	// This parameter is required.
	Query                     *string `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywordsShrink       *string `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty"`
	ReferLevel                *string `json:"referLevel,omitempty" xml:"referLevel,omitempty"`
	SortKeyAndDirectionShrink *string `json:"sortKeyAndDirection,omitempty" xml:"sortKeyAndDirection,omitempty"`
	ThreadShrink              *string `json:"thread,omitempty" xml:"thread,omitempty"`
}

func (s RunSearchCaseFullTextShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextShrinkRequest) SetAppId(v string) *RunSearchCaseFullTextShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetFilterConditionShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetPageParamShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.PageParamShrink = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetQuery(v string) *RunSearchCaseFullTextShrinkRequest {
	s.Query = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetQueryKeywordsShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.QueryKeywordsShrink = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetReferLevel(v string) *RunSearchCaseFullTextShrinkRequest {
	s.ReferLevel = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetSortKeyAndDirectionShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.SortKeyAndDirectionShrink = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetThreadShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.ThreadShrink = &v
	return s
}

type RunSearchCaseFullTextResponseBody struct {
	// example:
	//
	// null
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *RunSearchCaseFullTextResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C844BE6B-33A9-5AC4-A1AE-97B131849E0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RunSearchCaseFullTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBody) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBody) SetCode(v string) *RunSearchCaseFullTextResponseBody {
	s.Code = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetData(v *RunSearchCaseFullTextResponseBodyData) *RunSearchCaseFullTextResponseBody {
	s.Data = v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetHttpStatusCode(v int64) *RunSearchCaseFullTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetMessage(v string) *RunSearchCaseFullTextResponseBody {
	s.Message = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetRequestId(v string) *RunSearchCaseFullTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBody) SetSuccess(v bool) *RunSearchCaseFullTextResponseBody {
	s.Success = &v
	return s
}

type RunSearchCaseFullTextResponseBodyData struct {
	CaseLevel  *string                                            `json:"caseLevel,omitempty" xml:"caseLevel,omitempty"`
	CaseResult []*RunSearchCaseFullTextResponseBodyDataCaseResult `json:"caseResult,omitempty" xml:"caseResult,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize      *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query         *string   `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywords []*string `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s RunSearchCaseFullTextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBodyData) SetCaseLevel(v string) *RunSearchCaseFullTextResponseBodyData {
	s.CaseLevel = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetCaseResult(v []*RunSearchCaseFullTextResponseBodyDataCaseResult) *RunSearchCaseFullTextResponseBodyData {
	s.CaseResult = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetCurrentPage(v int32) *RunSearchCaseFullTextResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetPageSize(v int32) *RunSearchCaseFullTextResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetQuery(v string) *RunSearchCaseFullTextResponseBodyData {
	s.Query = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetQueryKeywords(v []*string) *RunSearchCaseFullTextResponseBodyData {
	s.QueryKeywords = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyData) SetTotalCount(v int64) *RunSearchCaseFullTextResponseBodyData {
	s.TotalCount = &v
	return s
}

type RunSearchCaseFullTextResponseBodyDataCaseResult struct {
	CaseDomain *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain `json:"caseDomain,omitempty" xml:"caseDomain,omitempty" type:"Struct"`
	Mode       *string                                                    `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// 0.88
	Similarity *string `json:"similarity,omitempty" xml:"similarity,omitempty"`
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResult) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResult) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) SetCaseDomain(v *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) *RunSearchCaseFullTextResponseBodyDataCaseResult {
	s.CaseDomain = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) SetMode(v string) *RunSearchCaseFullTextResponseBodyDataCaseResult {
	s.Mode = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResult) SetSimilarity(v string) *RunSearchCaseFullTextResponseBodyDataCaseResult {
	s.Similarity = &v
	return s
}

type RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain struct {
	AbstractObj           *string                                                              `json:"abstractObj,omitempty" xml:"abstractObj,omitempty"`
	AppliedLaws           *string                                                              `json:"appliedLaws,omitempty" xml:"appliedLaws,omitempty"`
	BasicCase             *string                                                              `json:"basicCase,omitempty" xml:"basicCase,omitempty"`
	CaseBasic             *string                                                              `json:"caseBasic,omitempty" xml:"caseBasic,omitempty"`
	CaseCause             *string                                                              `json:"caseCause,omitempty" xml:"caseCause,omitempty"`
	CaseFeature           *string                                                              `json:"caseFeature,omitempty" xml:"caseFeature,omitempty"`
	CaseId                *string                                                              `json:"caseId,omitempty" xml:"caseId,omitempty"`
	CaseNo                *string                                                              `json:"caseNo,omitempty" xml:"caseNo,omitempty"`
	CaseSummary           *string                                                              `json:"caseSummary,omitempty" xml:"caseSummary,omitempty"`
	CaseTitle             *string                                                              `json:"caseTitle,omitempty" xml:"caseTitle,omitempty"`
	CaseType              *string                                                              `json:"caseType,omitempty" xml:"caseType,omitempty"`
	CloseCaseCause        *string                                                              `json:"closeCaseCause,omitempty" xml:"closeCaseCause,omitempty"`
	CourtFindOut          *string                                                              `json:"courtFindOut,omitempty" xml:"courtFindOut,omitempty"`
	CourtThink            *string                                                              `json:"courtThink,omitempty" xml:"courtThink,omitempty"`
	DataFrom              *string                                                              `json:"dataFrom,omitempty" xml:"dataFrom,omitempty"`
	DisputeFocus          *string                                                              `json:"disputeFocus,omitempty" xml:"disputeFocus,omitempty"`
	DisputeFocusTag       []*string                                                            `json:"disputeFocusTag,omitempty" xml:"disputeFocusTag,omitempty" type:"Repeated"`
	Disputedpoints        *string                                                              `json:"disputedpoints,omitempty" xml:"disputedpoints,omitempty"`
	DocumentType          *string                                                              `json:"documentType,omitempty" xml:"documentType,omitempty"`
	JudgReason            *string                                                              `json:"judgReason,omitempty" xml:"judgReason,omitempty"`
	Keyfacts              *string                                                              `json:"keyfacts,omitempty" xml:"keyfacts,omitempty"`
	LegalBasis            *string                                                              `json:"legalBasis,omitempty" xml:"legalBasis,omitempty"`
	Litigants             *string                                                              `json:"litigants,omitempty" xml:"litigants,omitempty"`
	LitigationParticipant *string                                                              `json:"litigationParticipant,omitempty" xml:"litigationParticipant,omitempty"`
	OpenCaseCause         *string                                                              `json:"openCaseCause,omitempty" xml:"openCaseCause,omitempty"`
	PreTrialProcess       *string                                                              `json:"preTrialProcess,omitempty" xml:"preTrialProcess,omitempty"`
	ReferLevel            *string                                                              `json:"referLevel,omitempty" xml:"referLevel,omitempty"`
	RefereeGist           *string                                                              `json:"refereeGist,omitempty" xml:"refereeGist,omitempty"`
	SourceContent         *string                                                              `json:"sourceContent,omitempty" xml:"sourceContent,omitempty"`
	TrialCourt            *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt `json:"trialCourt,omitempty" xml:"trialCourt,omitempty" type:"Struct"`
	// example:
	//
	// 2018-09-27
	TrialDate    *string `json:"trialDate,omitempty" xml:"trialDate,omitempty"`
	TrialLevel   *string `json:"trialLevel,omitempty" xml:"trialLevel,omitempty"`
	TrialProcess *string `json:"trialProcess,omitempty" xml:"trialProcess,omitempty"`
	TrialProgram *string `json:"trialProgram,omitempty" xml:"trialProgram,omitempty"`
	Verdict      *string `json:"verdict,omitempty" xml:"verdict,omitempty"`
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetAbstractObj(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.AbstractObj = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetAppliedLaws(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.AppliedLaws = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetBasicCase(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.BasicCase = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseBasic(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseBasic = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseCause(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseCause = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseFeature(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseFeature = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseId(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseId = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseNo(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseNo = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseSummary(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseSummary = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseTitle(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseTitle = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCaseType(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CaseType = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCloseCaseCause(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CloseCaseCause = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCourtFindOut(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CourtFindOut = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetCourtThink(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.CourtThink = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDataFrom(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.DataFrom = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDisputeFocus(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.DisputeFocus = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDisputeFocusTag(v []*string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.DisputeFocusTag = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDisputedpoints(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.Disputedpoints = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetDocumentType(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.DocumentType = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetJudgReason(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.JudgReason = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetKeyfacts(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.Keyfacts = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetLegalBasis(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.LegalBasis = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetLitigants(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.Litigants = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetLitigationParticipant(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.LitigationParticipant = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetOpenCaseCause(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.OpenCaseCause = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetPreTrialProcess(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.PreTrialProcess = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetReferLevel(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.ReferLevel = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetRefereeGist(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.RefereeGist = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetSourceContent(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.SourceContent = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialCourt(v *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialCourt = v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialDate(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialDate = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialLevel(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialLevel = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialProcess(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialProcess = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetTrialProgram(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.TrialProgram = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain) SetVerdict(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomain {
	s.Verdict = &v
	return s
}

type RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt struct {
	City        *string `json:"city,omitempty" xml:"city,omitempty"`
	CommonLevel *string `json:"commonLevel,omitempty" xml:"commonLevel,omitempty"`
	Country     *string `json:"country,omitempty" xml:"country,omitempty"`
	County      *string `json:"county,omitempty" xml:"county,omitempty"`
	District    *string `json:"district,omitempty" xml:"district,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Province    *string `json:"province,omitempty" xml:"province,omitempty"`
	// example:
	//
	// “”
	SpecialLevel *string `json:"specialLevel,omitempty" xml:"specialLevel,omitempty"`
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetCity(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.City = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetCommonLevel(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.CommonLevel = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetCountry(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.Country = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetCounty(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.County = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetDistrict(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.District = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetName(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.Name = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetProvince(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.Province = &v
	return s
}

func (s *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt) SetSpecialLevel(v string) *RunSearchCaseFullTextResponseBodyDataCaseResultCaseDomainTrialCourt {
	s.SpecialLevel = &v
	return s
}

type RunSearchCaseFullTextResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSearchCaseFullTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSearchCaseFullTextResponse) String() string {
	return tea.Prettify(s)
}

func (s RunSearchCaseFullTextResponse) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponse) SetHeaders(v map[string]*string) *RunSearchCaseFullTextResponse {
	s.Headers = v
	return s
}

func (s *RunSearchCaseFullTextResponse) SetStatusCode(v int32) *RunSearchCaseFullTextResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSearchCaseFullTextResponse) SetBody(v *RunSearchCaseFullTextResponseBody) *RunSearchCaseFullTextResponse {
	s.Body = v
	return s
}

type RunSearchLawQueryRequest struct {
	// example:
	//
	// farui
	AppId           *string                                  `json:"appId,omitempty" xml:"appId,omitempty"`
	FilterCondition *RunSearchLawQueryRequestFilterCondition `json:"filterCondition,omitempty" xml:"filterCondition,omitempty" type:"Struct"`
	PageParam       *RunSearchLawQueryRequestPageParam       `json:"pageParam,omitempty" xml:"pageParam,omitempty" type:"Struct"`
	// This parameter is required.
	Query         *string                         `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywords []*string                       `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty" type:"Repeated"`
	Thread        *RunSearchLawQueryRequestThread `json:"thread,omitempty" xml:"thread,omitempty" type:"Struct"`
}

func (s RunSearchLawQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryRequest) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequest) SetAppId(v string) *RunSearchLawQueryRequest {
	s.AppId = &v
	return s
}

func (s *RunSearchLawQueryRequest) SetFilterCondition(v *RunSearchLawQueryRequestFilterCondition) *RunSearchLawQueryRequest {
	s.FilterCondition = v
	return s
}

func (s *RunSearchLawQueryRequest) SetPageParam(v *RunSearchLawQueryRequestPageParam) *RunSearchLawQueryRequest {
	s.PageParam = v
	return s
}

func (s *RunSearchLawQueryRequest) SetQuery(v string) *RunSearchLawQueryRequest {
	s.Query = &v
	return s
}

func (s *RunSearchLawQueryRequest) SetQueryKeywords(v []*string) *RunSearchLawQueryRequest {
	s.QueryKeywords = v
	return s
}

func (s *RunSearchLawQueryRequest) SetThread(v *RunSearchLawQueryRequestThread) *RunSearchLawQueryRequest {
	s.Thread = v
	return s
}

type RunSearchLawQueryRequestFilterCondition struct {
	LawName *string `json:"lawName,omitempty" xml:"lawName,omitempty"`
}

func (s RunSearchLawQueryRequestFilterCondition) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryRequestFilterCondition) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequestFilterCondition) SetLawName(v string) *RunSearchLawQueryRequestFilterCondition {
	s.LawName = &v
	return s
}

type RunSearchLawQueryRequestPageParam struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s RunSearchLawQueryRequestPageParam) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryRequestPageParam) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequestPageParam) SetPageNumber(v int32) *RunSearchLawQueryRequestPageParam {
	s.PageNumber = &v
	return s
}

func (s *RunSearchLawQueryRequestPageParam) SetPageSize(v int32) *RunSearchLawQueryRequestPageParam {
	s.PageSize = &v
	return s
}

type RunSearchLawQueryRequestThread struct {
	Messages []*RunSearchLawQueryRequestThreadMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s RunSearchLawQueryRequestThread) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryRequestThread) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequestThread) SetMessages(v []*RunSearchLawQueryRequestThreadMessages) *RunSearchLawQueryRequestThread {
	s.Messages = v
	return s
}

type RunSearchLawQueryRequestThreadMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunSearchLawQueryRequestThreadMessages) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryRequestThreadMessages) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryRequestThreadMessages) SetContent(v string) *RunSearchLawQueryRequestThreadMessages {
	s.Content = &v
	return s
}

func (s *RunSearchLawQueryRequestThreadMessages) SetRole(v string) *RunSearchLawQueryRequestThreadMessages {
	s.Role = &v
	return s
}

type RunSearchLawQueryShrinkRequest struct {
	// example:
	//
	// farui
	AppId                 *string `json:"appId,omitempty" xml:"appId,omitempty"`
	FilterConditionShrink *string `json:"filterCondition,omitempty" xml:"filterCondition,omitempty"`
	PageParamShrink       *string `json:"pageParam,omitempty" xml:"pageParam,omitempty"`
	// This parameter is required.
	Query               *string `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywordsShrink *string `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty"`
	ThreadShrink        *string `json:"thread,omitempty" xml:"thread,omitempty"`
}

func (s RunSearchLawQueryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryShrinkRequest) SetAppId(v string) *RunSearchLawQueryShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetFilterConditionShrink(v string) *RunSearchLawQueryShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetPageParamShrink(v string) *RunSearchLawQueryShrinkRequest {
	s.PageParamShrink = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetQuery(v string) *RunSearchLawQueryShrinkRequest {
	s.Query = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetQueryKeywordsShrink(v string) *RunSearchLawQueryShrinkRequest {
	s.QueryKeywordsShrink = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetThreadShrink(v string) *RunSearchLawQueryShrinkRequest {
	s.ThreadShrink = &v
	return s
}

type RunSearchLawQueryResponseBody struct {
	// example:
	//
	// Ok
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *RunSearchLawQueryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 05062567-EB51-50F6-AF56-0BE44955848D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RunSearchLawQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryResponseBody) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBody) SetCode(v string) *RunSearchLawQueryResponseBody {
	s.Code = &v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetData(v *RunSearchLawQueryResponseBodyData) *RunSearchLawQueryResponseBody {
	s.Data = v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetHttpStatusCode(v int64) *RunSearchLawQueryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetMessage(v string) *RunSearchLawQueryResponseBody {
	s.Message = &v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetRequestId(v string) *RunSearchLawQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSearchLawQueryResponseBody) SetSuccess(v bool) *RunSearchLawQueryResponseBody {
	s.Success = &v
	return s
}

type RunSearchLawQueryResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int32                                        `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	LawResult   []*RunSearchLawQueryResponseBodyDataLawResult `json:"lawResult,omitempty" xml:"lawResult,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	PageSize            *int32                                                `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query               *string                                               `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywords       []*string                                             `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty" type:"Repeated"`
	SortKeyAndDirection *RunSearchLawQueryResponseBodyDataSortKeyAndDirection `json:"sortKeyAndDirection,omitempty" xml:"sortKeyAndDirection,omitempty" type:"Struct"`
	// example:
	//
	// 0
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s RunSearchLawQueryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBodyData) SetCurrentPage(v int32) *RunSearchLawQueryResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetLawResult(v []*RunSearchLawQueryResponseBodyDataLawResult) *RunSearchLawQueryResponseBodyData {
	s.LawResult = v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetPageSize(v int32) *RunSearchLawQueryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetQuery(v string) *RunSearchLawQueryResponseBodyData {
	s.Query = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetQueryKeywords(v []*string) *RunSearchLawQueryResponseBodyData {
	s.QueryKeywords = v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetSortKeyAndDirection(v *RunSearchLawQueryResponseBodyDataSortKeyAndDirection) *RunSearchLawQueryResponseBodyData {
	s.SortKeyAndDirection = v
	return s
}

func (s *RunSearchLawQueryResponseBodyData) SetTotalCount(v int64) *RunSearchLawQueryResponseBodyData {
	s.TotalCount = &v
	return s
}

type RunSearchLawQueryResponseBodyDataLawResult struct {
	LawDomain *RunSearchLawQueryResponseBodyDataLawResultLawDomain `json:"lawDomain,omitempty" xml:"lawDomain,omitempty" type:"Struct"`
	// example:
	//
	// 0.0050
	Similarity *string `json:"similarity,omitempty" xml:"similarity,omitempty"`
}

func (s RunSearchLawQueryResponseBodyDataLawResult) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryResponseBodyDataLawResult) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBodyDataLawResult) SetLawDomain(v *RunSearchLawQueryResponseBodyDataLawResultLawDomain) *RunSearchLawQueryResponseBodyDataLawResult {
	s.LawDomain = v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResult) SetSimilarity(v string) *RunSearchLawQueryResponseBodyDataLawResult {
	s.Similarity = &v
	return s
}

type RunSearchLawQueryResponseBodyDataLawResultLawDomain struct {
	AbolitionBasis         *string `json:"abolitionBasis,omitempty" xml:"abolitionBasis,omitempty"`
	ImplementYearMonthDate *string `json:"implementYearMonthDate,omitempty" xml:"implementYearMonthDate,omitempty"`
	// example:
	//
	// null
	InvalidBasis *string `json:"invalidBasis,omitempty" xml:"invalidBasis,omitempty"`
	// example:
	//
	// ""
	IssuingNo    *string `json:"issuingNo,omitempty" xml:"issuingNo,omitempty"`
	IssuingOrgan *string `json:"issuingOrgan,omitempty" xml:"issuingOrgan,omitempty"`
	// example:
	//
	// b2274825c8c3bc2343ca73680243ddc8
	LawId *string `json:"lawId,omitempty" xml:"lawId,omitempty"`
	// example:
	//
	// ccc209683be1509676174fd6890f24b8
	LawItemId        *string `json:"lawItemId,omitempty" xml:"lawItemId,omitempty"`
	LawName          *string `json:"lawName,omitempty" xml:"lawName,omitempty"`
	LawOrder         *string `json:"lawOrder,omitempty" xml:"lawOrder,omitempty"`
	LawSourceContent *string `json:"lawSourceContent,omitempty" xml:"lawSourceContent,omitempty"`
	LawTitle         *string `json:"lawTitle,omitempty" xml:"lawTitle,omitempty"`
	// example:
	//
	// "[]"
	ModifyBasis          *string `json:"modifyBasis,omitempty" xml:"modifyBasis,omitempty"`
	PotencyLevel         *string `json:"potencyLevel,omitempty" xml:"potencyLevel,omitempty"`
	ReleaseYearMonthDate *string `json:"releaseYearMonthDate,omitempty" xml:"releaseYearMonthDate,omitempty"`
	// example:
	//
	// null
	ThematicClassify *string `json:"thematicClassify,omitempty" xml:"thematicClassify,omitempty"`
	Timeliness       *string `json:"timeliness,omitempty" xml:"timeliness,omitempty"`
}

func (s RunSearchLawQueryResponseBodyDataLawResultLawDomain) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryResponseBodyDataLawResultLawDomain) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetAbolitionBasis(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.AbolitionBasis = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetImplementYearMonthDate(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.ImplementYearMonthDate = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetInvalidBasis(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.InvalidBasis = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetIssuingNo(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.IssuingNo = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetIssuingOrgan(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.IssuingOrgan = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawId(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawId = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawItemId(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawItemId = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawName(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawName = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawOrder(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawOrder = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawSourceContent(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawSourceContent = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetLawTitle(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.LawTitle = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetModifyBasis(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.ModifyBasis = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetPotencyLevel(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.PotencyLevel = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetReleaseYearMonthDate(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.ReleaseYearMonthDate = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetThematicClassify(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.ThematicClassify = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataLawResultLawDomain) SetTimeliness(v string) *RunSearchLawQueryResponseBodyDataLawResultLawDomain {
	s.Timeliness = &v
	return s
}

type RunSearchLawQueryResponseBodyDataSortKeyAndDirection struct {
	// example:
	//
	// desc
	ReleaseYearMonthDate *string `json:"releaseYearMonthDate,omitempty" xml:"releaseYearMonthDate,omitempty"`
	// example:
	//
	// desc
	Similarity *string `json:"similarity,omitempty" xml:"similarity,omitempty"`
}

func (s RunSearchLawQueryResponseBodyDataSortKeyAndDirection) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryResponseBodyDataSortKeyAndDirection) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponseBodyDataSortKeyAndDirection) SetReleaseYearMonthDate(v string) *RunSearchLawQueryResponseBodyDataSortKeyAndDirection {
	s.ReleaseYearMonthDate = &v
	return s
}

func (s *RunSearchLawQueryResponseBodyDataSortKeyAndDirection) SetSimilarity(v string) *RunSearchLawQueryResponseBodyDataSortKeyAndDirection {
	s.Similarity = &v
	return s
}

type RunSearchLawQueryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSearchLawQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSearchLawQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s RunSearchLawQueryResponse) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryResponse) SetHeaders(v map[string]*string) *RunSearchLawQueryResponse {
	s.Headers = v
	return s
}

func (s *RunSearchLawQueryResponse) SetStatusCode(v int32) *RunSearchLawQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSearchLawQueryResponse) SetBody(v *RunSearchLawQueryResponseBody) *RunSearchLawQueryResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("farui"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传合同文件
//
// @param request - CreateTextFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTextFileResponse
func (client *Client) CreateTextFileWithOptions(WorkspaceId *string, request *CreateTextFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTextFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["CreateTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.TextFileName)) {
		body["TextFileName"] = request.TextFileName
	}

	if !tea.BoolValue(util.IsUnset(request.TextFileUrl)) {
		body["TextFileUrl"] = request.TextFileUrl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTextFile"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/data/textFile"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTextFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传合同文件
//
// @param request - CreateTextFileRequest
//
// @return CreateTextFileResponse
func (client *Client) CreateTextFile(WorkspaceId *string, request *CreateTextFileRequest) (_result *CreateTextFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTextFileResponse{}
	_body, _err := client.CreateTextFileWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTextFileAdvance(WorkspaceId *string, request *CreateTextFileAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTextFileResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("FaRui"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	createTextFileReq := &CreateTextFileRequest{}
	openapiutil.Convert(request, createTextFileReq)
	if !tea.BoolValue(util.IsUnset(request.TextFileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.TextFileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		createTextFileReq.TextFileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	createTextFileResp, _err := client.CreateTextFileWithOptions(WorkspaceId, createTextFileReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createTextFileResp
	return _result, _err
}

// Summary:
//
// 生成合同审查结果
//
// @param tmpReq - RunContractResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractResultGenerationResponse
func (client *Client) RunContractResultGenerationWithOptions(workspaceId *string, tmpReq *RunContractResultGenerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunContractResultGenerationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunContractResultGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Assistant)) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, tea.String("assistant"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AssistantShrink)) {
		body["assistant"] = request.AssistantShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunContractResultGeneration"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/farui/contract/result/genarate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunContractResultGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成合同审查结果
//
// @param request - RunContractResultGenerationRequest
//
// @return RunContractResultGenerationResponse
func (client *Client) RunContractResultGeneration(workspaceId *string, request *RunContractResultGenerationRequest) (_result *RunContractResultGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunContractResultGenerationResponse{}
	_body, _err := client.RunContractResultGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成合同审查规则
//
// @param tmpReq - RunContractRuleGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContractRuleGenerationResponse
func (client *Client) RunContractRuleGenerationWithOptions(workspaceId *string, tmpReq *RunContractRuleGenerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunContractRuleGenerationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunContractRuleGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Assistant)) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, tea.String("assistant"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AssistantShrink)) {
		body["assistant"] = request.AssistantShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunContractRuleGeneration"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/farui/contract/rule/genarate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunContractRuleGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成合同审查规则
//
// @param request - RunContractRuleGenerationRequest
//
// @return RunContractRuleGenerationResponse
func (client *Client) RunContractRuleGeneration(workspaceId *string, request *RunContractRuleGenerationRequest) (_result *RunContractRuleGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunContractRuleGenerationResponse{}
	_body, _err := client.RunContractRuleGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 法律咨询
//
// @param tmpReq - RunLegalAdviceConsultationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLegalAdviceConsultationResponse
func (client *Client) RunLegalAdviceConsultationWithOptions(workspaceId *string, tmpReq *RunLegalAdviceConsultationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunLegalAdviceConsultationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunLegalAdviceConsultationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Assistant)) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, tea.String("assistant"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Thread)) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, tea.String("thread"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AssistantShrink)) {
		body["assistant"] = request.AssistantShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.ThreadShrink)) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunLegalAdviceConsultation"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/farui/legalAdvice/consult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunLegalAdviceConsultationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 法律咨询
//
// @param request - RunLegalAdviceConsultationRequest
//
// @return RunLegalAdviceConsultationResponse
func (client *Client) RunLegalAdviceConsultation(workspaceId *string, request *RunLegalAdviceConsultationRequest) (_result *RunLegalAdviceConsultationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunLegalAdviceConsultationResponse{}
	_body, _err := client.RunLegalAdviceConsultationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 案例检索
//
// @param tmpReq - RunSearchCaseFullTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchCaseFullTextResponse
func (client *Client) RunSearchCaseFullTextWithOptions(workspaceId *string, tmpReq *RunSearchCaseFullTextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunSearchCaseFullTextResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunSearchCaseFullTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FilterCondition)) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, tea.String("filterCondition"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PageParam)) {
		request.PageParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageParam, tea.String("pageParam"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.QueryKeywords)) {
		request.QueryKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryKeywords, tea.String("queryKeywords"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SortKeyAndDirection)) {
		request.SortKeyAndDirectionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SortKeyAndDirection, tea.String("sortKeyAndDirection"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Thread)) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, tea.String("thread"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.FilterConditionShrink)) {
		body["filterCondition"] = request.FilterConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageParamShrink)) {
		body["pageParam"] = request.PageParamShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywordsShrink)) {
		body["queryKeywords"] = request.QueryKeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ReferLevel)) {
		body["referLevel"] = request.ReferLevel
	}

	if !tea.BoolValue(util.IsUnset(request.SortKeyAndDirectionShrink)) {
		body["sortKeyAndDirection"] = request.SortKeyAndDirectionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ThreadShrink)) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunSearchCaseFullText"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/farui/search/case/fulltext"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunSearchCaseFullTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 案例检索
//
// @param request - RunSearchCaseFullTextRequest
//
// @return RunSearchCaseFullTextResponse
func (client *Client) RunSearchCaseFullText(workspaceId *string, request *RunSearchCaseFullTextRequest) (_result *RunSearchCaseFullTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunSearchCaseFullTextResponse{}
	_body, _err := client.RunSearchCaseFullTextWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 法规搜索
//
// @param tmpReq - RunSearchLawQueryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSearchLawQueryResponse
func (client *Client) RunSearchLawQueryWithOptions(workspaceId *string, tmpReq *RunSearchLawQueryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunSearchLawQueryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunSearchLawQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FilterCondition)) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, tea.String("filterCondition"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PageParam)) {
		request.PageParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageParam, tea.String("pageParam"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.QueryKeywords)) {
		request.QueryKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryKeywords, tea.String("queryKeywords"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Thread)) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, tea.String("thread"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.FilterConditionShrink)) {
		body["filterCondition"] = request.FilterConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageParamShrink)) {
		body["pageParam"] = request.PageParamShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKeywordsShrink)) {
		body["queryKeywords"] = request.QueryKeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ThreadShrink)) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunSearchLawQuery"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/farui/search/law/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunSearchLawQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 法规搜索
//
// @param request - RunSearchLawQueryRequest
//
// @return RunSearchLawQueryResponse
func (client *Client) RunSearchLawQuery(workspaceId *string, request *RunSearchLawQueryRequest) (_result *RunSearchLawQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunSearchLawQueryResponse{}
	_body, _err := client.RunSearchLawQueryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
