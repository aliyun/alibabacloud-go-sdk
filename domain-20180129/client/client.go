// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AcknowledgeTaskResultRequest struct {
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TaskDetailNo []*string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty" type:"Repeated"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s AcknowledgeTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s AcknowledgeTaskResultRequest) GoString() string {
	return s.String()
}

func (s *AcknowledgeTaskResultRequest) SetLang(v string) *AcknowledgeTaskResultRequest {
	s.Lang = &v
	return s
}

func (s *AcknowledgeTaskResultRequest) SetTaskDetailNo(v []*string) *AcknowledgeTaskResultRequest {
	s.TaskDetailNo = v
	return s
}

func (s *AcknowledgeTaskResultRequest) SetUserClientIp(v string) *AcknowledgeTaskResultRequest {
	s.UserClientIp = &v
	return s
}

type AcknowledgeTaskResultResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AcknowledgeTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcknowledgeTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *AcknowledgeTaskResultResponseBody) SetRequestId(v string) *AcknowledgeTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcknowledgeTaskResultResponseBody) SetResult(v int32) *AcknowledgeTaskResultResponseBody {
	s.Result = &v
	return s
}

type AcknowledgeTaskResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AcknowledgeTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AcknowledgeTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s AcknowledgeTaskResultResponse) GoString() string {
	return s.String()
}

func (s *AcknowledgeTaskResultResponse) SetHeaders(v map[string]*string) *AcknowledgeTaskResultResponse {
	s.Headers = v
	return s
}

func (s *AcknowledgeTaskResultResponse) SetStatusCode(v int32) *AcknowledgeTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *AcknowledgeTaskResultResponse) SetBody(v *AcknowledgeTaskResultResponseBody) *AcknowledgeTaskResultResponse {
	s.Body = v
	return s
}

type BatchFuzzyMatchDomainSensitiveWordRequest struct {
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s BatchFuzzyMatchDomainSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) SetKeyword(v string) *BatchFuzzyMatchDomainSensitiveWordRequest {
	s.Keyword = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) SetLang(v string) *BatchFuzzyMatchDomainSensitiveWordRequest {
	s.Lang = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) SetUserClientIp(v string) *BatchFuzzyMatchDomainSensitiveWordRequest {
	s.UserClientIp = &v
	return s
}

type BatchFuzzyMatchDomainSensitiveWordResponseBody struct {
	RequestId                    *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SensitiveWordMatchResultList *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList `json:"SensitiveWordMatchResultList,omitempty" xml:"SensitiveWordMatchResultList,omitempty" type:"Struct"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBody) SetRequestId(v string) *BatchFuzzyMatchDomainSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBody) SetSensitiveWordMatchResultList(v *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) *BatchFuzzyMatchDomainSensitiveWordResponseBody {
	s.SensitiveWordMatchResultList = v
	return s
}

type BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList struct {
	SensitiveWordMatchResult []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult `json:"SensitiveWordMatchResult,omitempty" xml:"SensitiveWordMatchResult,omitempty" type:"Repeated"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) String() string {
	return tea.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList) SetSensitiveWordMatchResult(v []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultList {
	s.SensitiveWordMatchResult = v
	return s
}

type BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult struct {
	Exist               *bool                                                                                                                  `json:"Exist,omitempty" xml:"Exist,omitempty"`
	Keyword             *string                                                                                                                `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	MatchedSentiveWords *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords `json:"MatchedSentiveWords,omitempty" xml:"MatchedSentiveWords,omitempty" type:"Struct"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) String() string {
	return tea.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) SetExist(v bool) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult {
	s.Exist = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) SetKeyword(v string) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult {
	s.Keyword = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult) SetMatchedSentiveWords(v *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResult {
	s.MatchedSentiveWords = v
	return s
}

type BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords struct {
	MatchedSensitiveWord []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord `json:"MatchedSensitiveWord,omitempty" xml:"MatchedSensitiveWord,omitempty" type:"Repeated"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) String() string {
	return tea.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords) SetMatchedSensitiveWord(v []*BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWords {
	s.MatchedSensitiveWord = v
	return s
}

type BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord struct {
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) String() string {
	return tea.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord) SetWord(v string) *BatchFuzzyMatchDomainSensitiveWordResponseBodySensitiveWordMatchResultListSensitiveWordMatchResultMatchedSentiveWordsMatchedSensitiveWord {
	s.Word = &v
	return s
}

type BatchFuzzyMatchDomainSensitiveWordResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchFuzzyMatchDomainSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchFuzzyMatchDomainSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) SetHeaders(v map[string]*string) *BatchFuzzyMatchDomainSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) SetStatusCode(v int32) *BatchFuzzyMatchDomainSensitiveWordResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordResponse) SetBody(v *BatchFuzzyMatchDomainSensitiveWordResponseBody) *BatchFuzzyMatchDomainSensitiveWordResponse {
	s.Body = v
	return s
}

type CancelDomainVerificationRequest struct {
	ActionType   *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CancelDomainVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDomainVerificationRequest) GoString() string {
	return s.String()
}

func (s *CancelDomainVerificationRequest) SetActionType(v string) *CancelDomainVerificationRequest {
	s.ActionType = &v
	return s
}

func (s *CancelDomainVerificationRequest) SetInstanceId(v string) *CancelDomainVerificationRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelDomainVerificationRequest) SetLang(v string) *CancelDomainVerificationRequest {
	s.Lang = &v
	return s
}

func (s *CancelDomainVerificationRequest) SetUserClientIp(v string) *CancelDomainVerificationRequest {
	s.UserClientIp = &v
	return s
}

type CancelDomainVerificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDomainVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDomainVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDomainVerificationResponseBody) SetRequestId(v string) *CancelDomainVerificationResponseBody {
	s.RequestId = &v
	return s
}

type CancelDomainVerificationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelDomainVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelDomainVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDomainVerificationResponse) GoString() string {
	return s.String()
}

func (s *CancelDomainVerificationResponse) SetHeaders(v map[string]*string) *CancelDomainVerificationResponse {
	s.Headers = v
	return s
}

func (s *CancelDomainVerificationResponse) SetStatusCode(v int32) *CancelDomainVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDomainVerificationResponse) SetBody(v *CancelDomainVerificationResponseBody) *CancelDomainVerificationResponse {
	s.Body = v
	return s
}

type CancelOperationAuditRequest struct {
	AuditRecordId *int64  `json:"AuditRecordId,omitempty" xml:"AuditRecordId,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CancelOperationAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOperationAuditRequest) GoString() string {
	return s.String()
}

func (s *CancelOperationAuditRequest) SetAuditRecordId(v int64) *CancelOperationAuditRequest {
	s.AuditRecordId = &v
	return s
}

func (s *CancelOperationAuditRequest) SetLang(v string) *CancelOperationAuditRequest {
	s.Lang = &v
	return s
}

type CancelOperationAuditResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOperationAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOperationAuditResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOperationAuditResponseBody) SetRequestId(v string) *CancelOperationAuditResponseBody {
	s.RequestId = &v
	return s
}

type CancelOperationAuditResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelOperationAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelOperationAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelOperationAuditResponse) GoString() string {
	return s.String()
}

func (s *CancelOperationAuditResponse) SetHeaders(v map[string]*string) *CancelOperationAuditResponse {
	s.Headers = v
	return s
}

func (s *CancelOperationAuditResponse) SetStatusCode(v int32) *CancelOperationAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOperationAuditResponse) SetBody(v *CancelOperationAuditResponseBody) *CancelOperationAuditResponse {
	s.Body = v
	return s
}

type CancelQualificationVerificationRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	QualificationType *string `json:"QualificationType,omitempty" xml:"QualificationType,omitempty"`
	UserClientIp      *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CancelQualificationVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelQualificationVerificationRequest) GoString() string {
	return s.String()
}

func (s *CancelQualificationVerificationRequest) SetInstanceId(v string) *CancelQualificationVerificationRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelQualificationVerificationRequest) SetLang(v string) *CancelQualificationVerificationRequest {
	s.Lang = &v
	return s
}

func (s *CancelQualificationVerificationRequest) SetQualificationType(v string) *CancelQualificationVerificationRequest {
	s.QualificationType = &v
	return s
}

func (s *CancelQualificationVerificationRequest) SetUserClientIp(v string) *CancelQualificationVerificationRequest {
	s.UserClientIp = &v
	return s
}

type CancelQualificationVerificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelQualificationVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelQualificationVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelQualificationVerificationResponseBody) SetRequestId(v string) *CancelQualificationVerificationResponseBody {
	s.RequestId = &v
	return s
}

type CancelQualificationVerificationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelQualificationVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelQualificationVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelQualificationVerificationResponse) GoString() string {
	return s.String()
}

func (s *CancelQualificationVerificationResponse) SetHeaders(v map[string]*string) *CancelQualificationVerificationResponse {
	s.Headers = v
	return s
}

func (s *CancelQualificationVerificationResponse) SetStatusCode(v int32) *CancelQualificationVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelQualificationVerificationResponse) SetBody(v *CancelQualificationVerificationResponseBody) *CancelQualificationVerificationResponse {
	s.Body = v
	return s
}

type CancelTaskRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TaskNo       *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CancelTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTaskRequest) SetLang(v string) *CancelTaskRequest {
	s.Lang = &v
	return s
}

func (s *CancelTaskRequest) SetTaskNo(v string) *CancelTaskRequest {
	s.TaskNo = &v
	return s
}

func (s *CancelTaskRequest) SetUserClientIp(v string) *CancelTaskRequest {
	s.UserClientIp = &v
	return s
}

type CancelTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelTaskResponse) SetHeaders(v map[string]*string) *CancelTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelTaskResponse) SetStatusCode(v int32) *CancelTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelTaskResponse) SetBody(v *CancelTaskResponseBody) *CancelTaskResponse {
	s.Body = v
	return s
}

type CheckDomainRequest struct {
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FeeCommand  *string `json:"FeeCommand,omitempty" xml:"FeeCommand,omitempty"`
	FeeCurrency *string `json:"FeeCurrency,omitempty" xml:"FeeCurrency,omitempty"`
	FeePeriod   *int32  `json:"FeePeriod,omitempty" xml:"FeePeriod,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CheckDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainRequest) SetDomainName(v string) *CheckDomainRequest {
	s.DomainName = &v
	return s
}

func (s *CheckDomainRequest) SetFeeCommand(v string) *CheckDomainRequest {
	s.FeeCommand = &v
	return s
}

func (s *CheckDomainRequest) SetFeeCurrency(v string) *CheckDomainRequest {
	s.FeeCurrency = &v
	return s
}

func (s *CheckDomainRequest) SetFeePeriod(v int32) *CheckDomainRequest {
	s.FeePeriod = &v
	return s
}

func (s *CheckDomainRequest) SetLang(v string) *CheckDomainRequest {
	s.Lang = &v
	return s
}

type CheckDomainResponseBody struct {
	Avail        *string `json:"Avail,omitempty" xml:"Avail,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DynamicCheck *bool   `json:"DynamicCheck,omitempty" xml:"DynamicCheck,omitempty"`
	Premium      *string `json:"Premium,omitempty" xml:"Premium,omitempty"`
	Price        *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBody) SetAvail(v string) *CheckDomainResponseBody {
	s.Avail = &v
	return s
}

func (s *CheckDomainResponseBody) SetDomainName(v string) *CheckDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *CheckDomainResponseBody) SetDynamicCheck(v bool) *CheckDomainResponseBody {
	s.DynamicCheck = &v
	return s
}

func (s *CheckDomainResponseBody) SetPremium(v string) *CheckDomainResponseBody {
	s.Premium = &v
	return s
}

func (s *CheckDomainResponseBody) SetPrice(v int64) *CheckDomainResponseBody {
	s.Price = &v
	return s
}

func (s *CheckDomainResponseBody) SetReason(v string) *CheckDomainResponseBody {
	s.Reason = &v
	return s
}

func (s *CheckDomainResponseBody) SetRequestId(v string) *CheckDomainResponseBody {
	s.RequestId = &v
	return s
}

type CheckDomainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainResponse) GoString() string {
	return s.String()
}

func (s *CheckDomainResponse) SetHeaders(v map[string]*string) *CheckDomainResponse {
	s.Headers = v
	return s
}

func (s *CheckDomainResponse) SetStatusCode(v int32) *CheckDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDomainResponse) SetBody(v *CheckDomainResponseBody) *CheckDomainResponse {
	s.Body = v
	return s
}

type CheckDomainSunriseClaimRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CheckDomainSunriseClaimRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainSunriseClaimRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainSunriseClaimRequest) SetDomainName(v string) *CheckDomainSunriseClaimRequest {
	s.DomainName = &v
	return s
}

func (s *CheckDomainSunriseClaimRequest) SetLang(v string) *CheckDomainSunriseClaimRequest {
	s.Lang = &v
	return s
}

func (s *CheckDomainSunriseClaimRequest) SetUserClientIp(v string) *CheckDomainSunriseClaimRequest {
	s.UserClientIp = &v
	return s
}

type CheckDomainSunriseClaimResponseBody struct {
	ClaimKey  *string `json:"ClaimKey,omitempty" xml:"ClaimKey,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckDomainSunriseClaimResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainSunriseClaimResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainSunriseClaimResponseBody) SetClaimKey(v string) *CheckDomainSunriseClaimResponseBody {
	s.ClaimKey = &v
	return s
}

func (s *CheckDomainSunriseClaimResponseBody) SetRequestId(v string) *CheckDomainSunriseClaimResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDomainSunriseClaimResponseBody) SetResult(v int32) *CheckDomainSunriseClaimResponseBody {
	s.Result = &v
	return s
}

type CheckDomainSunriseClaimResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckDomainSunriseClaimResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckDomainSunriseClaimResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainSunriseClaimResponse) GoString() string {
	return s.String()
}

func (s *CheckDomainSunriseClaimResponse) SetHeaders(v map[string]*string) *CheckDomainSunriseClaimResponse {
	s.Headers = v
	return s
}

func (s *CheckDomainSunriseClaimResponse) SetStatusCode(v int32) *CheckDomainSunriseClaimResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDomainSunriseClaimResponse) SetBody(v *CheckDomainSunriseClaimResponseBody) *CheckDomainSunriseClaimResponse {
	s.Body = v
	return s
}

type CheckMaxYearOfServerLockRequest struct {
	CheckAction  *string `json:"CheckAction,omitempty" xml:"CheckAction,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CheckMaxYearOfServerLockRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckMaxYearOfServerLockRequest) GoString() string {
	return s.String()
}

func (s *CheckMaxYearOfServerLockRequest) SetCheckAction(v string) *CheckMaxYearOfServerLockRequest {
	s.CheckAction = &v
	return s
}

func (s *CheckMaxYearOfServerLockRequest) SetDomainName(v string) *CheckMaxYearOfServerLockRequest {
	s.DomainName = &v
	return s
}

func (s *CheckMaxYearOfServerLockRequest) SetLang(v string) *CheckMaxYearOfServerLockRequest {
	s.Lang = &v
	return s
}

func (s *CheckMaxYearOfServerLockRequest) SetUserClientIp(v string) *CheckMaxYearOfServerLockRequest {
	s.UserClientIp = &v
	return s
}

type CheckMaxYearOfServerLockResponseBody struct {
	MaxYear   *int32  `json:"MaxYear,omitempty" xml:"MaxYear,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckMaxYearOfServerLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckMaxYearOfServerLockResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMaxYearOfServerLockResponseBody) SetMaxYear(v int32) *CheckMaxYearOfServerLockResponseBody {
	s.MaxYear = &v
	return s
}

func (s *CheckMaxYearOfServerLockResponseBody) SetRequestId(v string) *CheckMaxYearOfServerLockResponseBody {
	s.RequestId = &v
	return s
}

type CheckMaxYearOfServerLockResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckMaxYearOfServerLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckMaxYearOfServerLockResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMaxYearOfServerLockResponse) GoString() string {
	return s.String()
}

func (s *CheckMaxYearOfServerLockResponse) SetHeaders(v map[string]*string) *CheckMaxYearOfServerLockResponse {
	s.Headers = v
	return s
}

func (s *CheckMaxYearOfServerLockResponse) SetStatusCode(v int32) *CheckMaxYearOfServerLockResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMaxYearOfServerLockResponse) SetBody(v *CheckMaxYearOfServerLockResponseBody) *CheckMaxYearOfServerLockResponse {
	s.Body = v
	return s
}

type CheckProcessingServerLockApplyRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FeePeriod    *int32  `json:"FeePeriod,omitempty" xml:"FeePeriod,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CheckProcessingServerLockApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckProcessingServerLockApplyRequest) GoString() string {
	return s.String()
}

func (s *CheckProcessingServerLockApplyRequest) SetDomainName(v string) *CheckProcessingServerLockApplyRequest {
	s.DomainName = &v
	return s
}

func (s *CheckProcessingServerLockApplyRequest) SetFeePeriod(v int32) *CheckProcessingServerLockApplyRequest {
	s.FeePeriod = &v
	return s
}

func (s *CheckProcessingServerLockApplyRequest) SetLang(v string) *CheckProcessingServerLockApplyRequest {
	s.Lang = &v
	return s
}

func (s *CheckProcessingServerLockApplyRequest) SetUserClientIp(v string) *CheckProcessingServerLockApplyRequest {
	s.UserClientIp = &v
	return s
}

type CheckProcessingServerLockApplyResponseBody struct {
	Exists    *bool   `json:"Exists,omitempty" xml:"Exists,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckProcessingServerLockApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckProcessingServerLockApplyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckProcessingServerLockApplyResponseBody) SetExists(v bool) *CheckProcessingServerLockApplyResponseBody {
	s.Exists = &v
	return s
}

func (s *CheckProcessingServerLockApplyResponseBody) SetRequestId(v string) *CheckProcessingServerLockApplyResponseBody {
	s.RequestId = &v
	return s
}

type CheckProcessingServerLockApplyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckProcessingServerLockApplyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckProcessingServerLockApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckProcessingServerLockApplyResponse) GoString() string {
	return s.String()
}

func (s *CheckProcessingServerLockApplyResponse) SetHeaders(v map[string]*string) *CheckProcessingServerLockApplyResponse {
	s.Headers = v
	return s
}

func (s *CheckProcessingServerLockApplyResponse) SetStatusCode(v int32) *CheckProcessingServerLockApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckProcessingServerLockApplyResponse) SetBody(v *CheckProcessingServerLockApplyResponseBody) *CheckProcessingServerLockApplyResponse {
	s.Body = v
	return s
}

type CheckTransferInFeasibilityRequest struct {
	DomainName                *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang                      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TransferAuthorizationCode *string `json:"TransferAuthorizationCode,omitempty" xml:"TransferAuthorizationCode,omitempty"`
	UserClientIp              *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CheckTransferInFeasibilityRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckTransferInFeasibilityRequest) GoString() string {
	return s.String()
}

func (s *CheckTransferInFeasibilityRequest) SetDomainName(v string) *CheckTransferInFeasibilityRequest {
	s.DomainName = &v
	return s
}

func (s *CheckTransferInFeasibilityRequest) SetLang(v string) *CheckTransferInFeasibilityRequest {
	s.Lang = &v
	return s
}

func (s *CheckTransferInFeasibilityRequest) SetTransferAuthorizationCode(v string) *CheckTransferInFeasibilityRequest {
	s.TransferAuthorizationCode = &v
	return s
}

func (s *CheckTransferInFeasibilityRequest) SetUserClientIp(v string) *CheckTransferInFeasibilityRequest {
	s.UserClientIp = &v
	return s
}

type CheckTransferInFeasibilityResponseBody struct {
	CanTransfer *bool   `json:"CanTransfer,omitempty" xml:"CanTransfer,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ProductId   *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckTransferInFeasibilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckTransferInFeasibilityResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTransferInFeasibilityResponseBody) SetCanTransfer(v bool) *CheckTransferInFeasibilityResponseBody {
	s.CanTransfer = &v
	return s
}

func (s *CheckTransferInFeasibilityResponseBody) SetCode(v string) *CheckTransferInFeasibilityResponseBody {
	s.Code = &v
	return s
}

func (s *CheckTransferInFeasibilityResponseBody) SetMessage(v string) *CheckTransferInFeasibilityResponseBody {
	s.Message = &v
	return s
}

func (s *CheckTransferInFeasibilityResponseBody) SetProductId(v string) *CheckTransferInFeasibilityResponseBody {
	s.ProductId = &v
	return s
}

func (s *CheckTransferInFeasibilityResponseBody) SetRequestId(v string) *CheckTransferInFeasibilityResponseBody {
	s.RequestId = &v
	return s
}

type CheckTransferInFeasibilityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckTransferInFeasibilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckTransferInFeasibilityResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckTransferInFeasibilityResponse) GoString() string {
	return s.String()
}

func (s *CheckTransferInFeasibilityResponse) SetHeaders(v map[string]*string) *CheckTransferInFeasibilityResponse {
	s.Headers = v
	return s
}

func (s *CheckTransferInFeasibilityResponse) SetStatusCode(v int32) *CheckTransferInFeasibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckTransferInFeasibilityResponse) SetBody(v *CheckTransferInFeasibilityResponseBody) *CheckTransferInFeasibilityResponse {
	s.Body = v
	return s
}

type ConfirmTransferInEmailRequest struct {
	DomainName   []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Email        *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ConfirmTransferInEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmTransferInEmailRequest) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailRequest) SetDomainName(v []*string) *ConfirmTransferInEmailRequest {
	s.DomainName = v
	return s
}

func (s *ConfirmTransferInEmailRequest) SetEmail(v string) *ConfirmTransferInEmailRequest {
	s.Email = &v
	return s
}

func (s *ConfirmTransferInEmailRequest) SetLang(v string) *ConfirmTransferInEmailRequest {
	s.Lang = &v
	return s
}

func (s *ConfirmTransferInEmailRequest) SetUserClientIp(v string) *ConfirmTransferInEmailRequest {
	s.UserClientIp = &v
	return s
}

type ConfirmTransferInEmailResponseBody struct {
	FailList    *ConfirmTransferInEmailResponseBodyFailList    `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Struct"`
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList *ConfirmTransferInEmailResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Struct"`
}

func (s ConfirmTransferInEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmTransferInEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailResponseBody) SetFailList(v *ConfirmTransferInEmailResponseBodyFailList) *ConfirmTransferInEmailResponseBody {
	s.FailList = v
	return s
}

func (s *ConfirmTransferInEmailResponseBody) SetRequestId(v string) *ConfirmTransferInEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmTransferInEmailResponseBody) SetSuccessList(v *ConfirmTransferInEmailResponseBodySuccessList) *ConfirmTransferInEmailResponseBody {
	s.SuccessList = v
	return s
}

type ConfirmTransferInEmailResponseBodyFailList struct {
	FailDomain []*string `json:"FailDomain,omitempty" xml:"FailDomain,omitempty" type:"Repeated"`
}

func (s ConfirmTransferInEmailResponseBodyFailList) String() string {
	return tea.Prettify(s)
}

func (s ConfirmTransferInEmailResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailResponseBodyFailList) SetFailDomain(v []*string) *ConfirmTransferInEmailResponseBodyFailList {
	s.FailDomain = v
	return s
}

type ConfirmTransferInEmailResponseBodySuccessList struct {
	SuccessDomain []*string `json:"SuccessDomain,omitempty" xml:"SuccessDomain,omitempty" type:"Repeated"`
}

func (s ConfirmTransferInEmailResponseBodySuccessList) String() string {
	return tea.Prettify(s)
}

func (s ConfirmTransferInEmailResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailResponseBodySuccessList) SetSuccessDomain(v []*string) *ConfirmTransferInEmailResponseBodySuccessList {
	s.SuccessDomain = v
	return s
}

type ConfirmTransferInEmailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfirmTransferInEmailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmTransferInEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmTransferInEmailResponse) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailResponse) SetHeaders(v map[string]*string) *ConfirmTransferInEmailResponse {
	s.Headers = v
	return s
}

func (s *ConfirmTransferInEmailResponse) SetStatusCode(v int32) *ConfirmTransferInEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmTransferInEmailResponse) SetBody(v *ConfirmTransferInEmailResponseBody) *ConfirmTransferInEmailResponse {
	s.Body = v
	return s
}

type DeleteContactTemplatesRequest struct {
	RegistrantProfileIds *string `json:"RegistrantProfileIds,omitempty" xml:"RegistrantProfileIds,omitempty"`
	UserClientIp         *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteContactTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactTemplatesRequest) SetRegistrantProfileIds(v string) *DeleteContactTemplatesRequest {
	s.RegistrantProfileIds = &v
	return s
}

func (s *DeleteContactTemplatesRequest) SetUserClientIp(v string) *DeleteContactTemplatesRequest {
	s.UserClientIp = &v
	return s
}

type DeleteContactTemplatesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContactTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactTemplatesResponseBody) SetRequestId(v string) *DeleteContactTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteContactTemplatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContactTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContactTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactTemplatesResponse) SetHeaders(v map[string]*string) *DeleteContactTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactTemplatesResponse) SetStatusCode(v int32) *DeleteContactTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactTemplatesResponse) SetBody(v *DeleteContactTemplatesResponseBody) *DeleteContactTemplatesResponse {
	s.Body = v
	return s
}

type DeleteDomainGroupRequest struct {
	DomainGroupId *int64  `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp  *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteDomainGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupRequest) SetDomainGroupId(v int64) *DeleteDomainGroupRequest {
	s.DomainGroupId = &v
	return s
}

func (s *DeleteDomainGroupRequest) SetLang(v string) *DeleteDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDomainGroupRequest) SetUserClientIp(v string) *DeleteDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

type DeleteDomainGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupResponseBody) SetRequestId(v string) *DeleteDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupResponse) SetHeaders(v map[string]*string) *DeleteDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainGroupResponse) SetStatusCode(v int32) *DeleteDomainGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainGroupResponse) SetBody(v *DeleteDomainGroupResponseBody) *DeleteDomainGroupResponse {
	s.Body = v
	return s
}

type DeleteEmailVerificationRequest struct {
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteEmailVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationRequest) SetEmail(v string) *DeleteEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *DeleteEmailVerificationRequest) SetLang(v string) *DeleteEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *DeleteEmailVerificationRequest) SetUserClientIp(v string) *DeleteEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

type DeleteEmailVerificationResponseBody struct {
	FailList    []*DeleteEmailVerificationResponseBodyFailList    `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Repeated"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList []*DeleteEmailVerificationResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Repeated"`
}

func (s DeleteEmailVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationResponseBody) SetFailList(v []*DeleteEmailVerificationResponseBodyFailList) *DeleteEmailVerificationResponseBody {
	s.FailList = v
	return s
}

func (s *DeleteEmailVerificationResponseBody) SetRequestId(v string) *DeleteEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEmailVerificationResponseBody) SetSuccessList(v []*DeleteEmailVerificationResponseBodySuccessList) *DeleteEmailVerificationResponseBody {
	s.SuccessList = v
	return s
}

type DeleteEmailVerificationResponseBodyFailList struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteEmailVerificationResponseBodyFailList) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmailVerificationResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationResponseBodyFailList) SetCode(v string) *DeleteEmailVerificationResponseBodyFailList {
	s.Code = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodyFailList) SetEmail(v string) *DeleteEmailVerificationResponseBodyFailList {
	s.Email = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodyFailList) SetMessage(v string) *DeleteEmailVerificationResponseBodyFailList {
	s.Message = &v
	return s
}

type DeleteEmailVerificationResponseBodySuccessList struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteEmailVerificationResponseBodySuccessList) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmailVerificationResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationResponseBodySuccessList) SetCode(v string) *DeleteEmailVerificationResponseBodySuccessList {
	s.Code = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodySuccessList) SetEmail(v string) *DeleteEmailVerificationResponseBodySuccessList {
	s.Email = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodySuccessList) SetMessage(v string) *DeleteEmailVerificationResponseBodySuccessList {
	s.Message = &v
	return s
}

type DeleteEmailVerificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEmailVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationResponse) SetHeaders(v map[string]*string) *DeleteEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *DeleteEmailVerificationResponse) SetStatusCode(v int32) *DeleteEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEmailVerificationResponse) SetBody(v *DeleteEmailVerificationResponseBody) *DeleteEmailVerificationResponse {
	s.Body = v
	return s
}

type DeleteRegistrantProfileRequest struct {
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegistrantProfileId *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	UserClientIp        *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteRegistrantProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRegistrantProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteRegistrantProfileRequest) SetLang(v string) *DeleteRegistrantProfileRequest {
	s.Lang = &v
	return s
}

func (s *DeleteRegistrantProfileRequest) SetRegistrantProfileId(v int64) *DeleteRegistrantProfileRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *DeleteRegistrantProfileRequest) SetUserClientIp(v string) *DeleteRegistrantProfileRequest {
	s.UserClientIp = &v
	return s
}

type DeleteRegistrantProfileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRegistrantProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRegistrantProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRegistrantProfileResponseBody) SetRequestId(v string) *DeleteRegistrantProfileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRegistrantProfileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRegistrantProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRegistrantProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRegistrantProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteRegistrantProfileResponse) SetHeaders(v map[string]*string) *DeleteRegistrantProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteRegistrantProfileResponse) SetStatusCode(v int32) *DeleteRegistrantProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRegistrantProfileResponse) SetBody(v *DeleteRegistrantProfileResponseBody) *DeleteRegistrantProfileResponse {
	s.Body = v
	return s
}

type EmailVerifiedRequest struct {
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s EmailVerifiedRequest) String() string {
	return tea.Prettify(s)
}

func (s EmailVerifiedRequest) GoString() string {
	return s.String()
}

func (s *EmailVerifiedRequest) SetEmail(v string) *EmailVerifiedRequest {
	s.Email = &v
	return s
}

func (s *EmailVerifiedRequest) SetLang(v string) *EmailVerifiedRequest {
	s.Lang = &v
	return s
}

func (s *EmailVerifiedRequest) SetUserClientIp(v string) *EmailVerifiedRequest {
	s.UserClientIp = &v
	return s
}

type EmailVerifiedResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EmailVerifiedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EmailVerifiedResponseBody) GoString() string {
	return s.String()
}

func (s *EmailVerifiedResponseBody) SetRequestId(v string) *EmailVerifiedResponseBody {
	s.RequestId = &v
	return s
}

type EmailVerifiedResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EmailVerifiedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EmailVerifiedResponse) String() string {
	return tea.Prettify(s)
}

func (s EmailVerifiedResponse) GoString() string {
	return s.String()
}

func (s *EmailVerifiedResponse) SetHeaders(v map[string]*string) *EmailVerifiedResponse {
	s.Headers = v
	return s
}

func (s *EmailVerifiedResponse) SetStatusCode(v int32) *EmailVerifiedResponse {
	s.StatusCode = &v
	return s
}

func (s *EmailVerifiedResponse) SetBody(v *EmailVerifiedResponseBody) *EmailVerifiedResponse {
	s.Body = v
	return s
}

type FuzzyMatchDomainSensitiveWordRequest struct {
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s FuzzyMatchDomainSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordRequest) SetKeyword(v string) *FuzzyMatchDomainSensitiveWordRequest {
	s.Keyword = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordRequest) SetLang(v string) *FuzzyMatchDomainSensitiveWordRequest {
	s.Lang = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordRequest) SetUserClientIp(v string) *FuzzyMatchDomainSensitiveWordRequest {
	s.UserClientIp = &v
	return s
}

type FuzzyMatchDomainSensitiveWordResponseBody struct {
	Exist               *bool                                                         `json:"Exist,omitempty" xml:"Exist,omitempty"`
	Keyword             *string                                                       `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	MatchedSentiveWords *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords `json:"MatchedSentiveWords,omitempty" xml:"MatchedSentiveWords,omitempty" type:"Struct"`
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FuzzyMatchDomainSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) SetExist(v bool) *FuzzyMatchDomainSensitiveWordResponseBody {
	s.Exist = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) SetKeyword(v string) *FuzzyMatchDomainSensitiveWordResponseBody {
	s.Keyword = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) SetMatchedSentiveWords(v *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) *FuzzyMatchDomainSensitiveWordResponseBody {
	s.MatchedSentiveWords = v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponseBody) SetRequestId(v string) *FuzzyMatchDomainSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

type FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords struct {
	MatchedSensitiveWord []*FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord `json:"MatchedSensitiveWord,omitempty" xml:"MatchedSensitiveWord,omitempty" type:"Repeated"`
}

func (s FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) String() string {
	return tea.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords) SetMatchedSensitiveWord(v []*FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWords {
	s.MatchedSensitiveWord = v
	return s
}

type FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord struct {
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) String() string {
	return tea.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord) SetWord(v string) *FuzzyMatchDomainSensitiveWordResponseBodyMatchedSentiveWordsMatchedSensitiveWord {
	s.Word = &v
	return s
}

type FuzzyMatchDomainSensitiveWordResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FuzzyMatchDomainSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FuzzyMatchDomainSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordResponse) SetHeaders(v map[string]*string) *FuzzyMatchDomainSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponse) SetStatusCode(v int32) *FuzzyMatchDomainSensitiveWordResponse {
	s.StatusCode = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordResponse) SetBody(v *FuzzyMatchDomainSensitiveWordResponseBody) *FuzzyMatchDomainSensitiveWordResponse {
	s.Body = v
	return s
}

type GetOperationOssUploadPolicyRequest struct {
	AuditType *int32  `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetOperationOssUploadPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOperationOssUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetOperationOssUploadPolicyRequest) SetAuditType(v int32) *GetOperationOssUploadPolicyRequest {
	s.AuditType = &v
	return s
}

func (s *GetOperationOssUploadPolicyRequest) SetLang(v string) *GetOperationOssUploadPolicyRequest {
	s.Lang = &v
	return s
}

type GetOperationOssUploadPolicyResponseBody struct {
	Accessid      *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	ExpireTime    *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FileDir       *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	Host          *string `json:"Host,omitempty" xml:"Host,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature     *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetOperationOssUploadPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOperationOssUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationOssUploadPolicyResponseBody) SetAccessid(v string) *GetOperationOssUploadPolicyResponseBody {
	s.Accessid = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetEncodedPolicy(v string) *GetOperationOssUploadPolicyResponseBody {
	s.EncodedPolicy = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetExpireTime(v string) *GetOperationOssUploadPolicyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetFileDir(v string) *GetOperationOssUploadPolicyResponseBody {
	s.FileDir = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetHost(v string) *GetOperationOssUploadPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetRequestId(v string) *GetOperationOssUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetSignature(v string) *GetOperationOssUploadPolicyResponseBody {
	s.Signature = &v
	return s
}

type GetOperationOssUploadPolicyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOperationOssUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOperationOssUploadPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOperationOssUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetOperationOssUploadPolicyResponse) SetHeaders(v map[string]*string) *GetOperationOssUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetOperationOssUploadPolicyResponse) SetStatusCode(v int32) *GetOperationOssUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponse) SetBody(v *GetOperationOssUploadPolicyResponseBody) *GetOperationOssUploadPolicyResponse {
	s.Body = v
	return s
}

type GetQualificationUploadPolicyRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s GetQualificationUploadPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualificationUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetQualificationUploadPolicyRequest) SetLang(v string) *GetQualificationUploadPolicyRequest {
	s.Lang = &v
	return s
}

func (s *GetQualificationUploadPolicyRequest) SetUserClientIp(v string) *GetQualificationUploadPolicyRequest {
	s.UserClientIp = &v
	return s
}

type GetQualificationUploadPolicyResponseBody struct {
	Accessid  *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Prefix    *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetQualificationUploadPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualificationUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualificationUploadPolicyResponseBody) SetAccessid(v string) *GetQualificationUploadPolicyResponseBody {
	s.Accessid = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetDir(v string) *GetQualificationUploadPolicyResponseBody {
	s.Dir = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetExpire(v string) *GetQualificationUploadPolicyResponseBody {
	s.Expire = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetHost(v string) *GetQualificationUploadPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetPolicy(v string) *GetQualificationUploadPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetPrefix(v string) *GetQualificationUploadPolicyResponseBody {
	s.Prefix = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetRequestId(v string) *GetQualificationUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetSignature(v string) *GetQualificationUploadPolicyResponseBody {
	s.Signature = &v
	return s
}

type GetQualificationUploadPolicyResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQualificationUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQualificationUploadPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualificationUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetQualificationUploadPolicyResponse) SetHeaders(v map[string]*string) *GetQualificationUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetQualificationUploadPolicyResponse) SetStatusCode(v int32) *GetQualificationUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualificationUploadPolicyResponse) SetBody(v *GetQualificationUploadPolicyResponseBody) *GetQualificationUploadPolicyResponse {
	s.Body = v
	return s
}

type ListEmailVerificationRequest struct {
	BeginCreateTime    *int64  `json:"BeginCreateTime,omitempty" xml:"BeginCreateTime,omitempty"`
	Email              *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EndCreateTime      *int64  `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum            *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserClientIp       *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	VerificationStatus *int32  `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s ListEmailVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *ListEmailVerificationRequest) SetBeginCreateTime(v int64) *ListEmailVerificationRequest {
	s.BeginCreateTime = &v
	return s
}

func (s *ListEmailVerificationRequest) SetEmail(v string) *ListEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *ListEmailVerificationRequest) SetEndCreateTime(v int64) *ListEmailVerificationRequest {
	s.EndCreateTime = &v
	return s
}

func (s *ListEmailVerificationRequest) SetLang(v string) *ListEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *ListEmailVerificationRequest) SetPageNum(v int32) *ListEmailVerificationRequest {
	s.PageNum = &v
	return s
}

func (s *ListEmailVerificationRequest) SetPageSize(v int32) *ListEmailVerificationRequest {
	s.PageSize = &v
	return s
}

func (s *ListEmailVerificationRequest) SetUserClientIp(v string) *ListEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *ListEmailVerificationRequest) SetVerificationStatus(v int32) *ListEmailVerificationRequest {
	s.VerificationStatus = &v
	return s
}

type ListEmailVerificationResponseBody struct {
	CurrentPageNum *int32                                   `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListEmailVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextPage       *bool                                    `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                                    `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                                   `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                   `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListEmailVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ListEmailVerificationResponseBody) SetCurrentPageNum(v int32) *ListEmailVerificationResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetData(v []*ListEmailVerificationResponseBodyData) *ListEmailVerificationResponseBody {
	s.Data = v
	return s
}

func (s *ListEmailVerificationResponseBody) SetNextPage(v bool) *ListEmailVerificationResponseBody {
	s.NextPage = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetPageSize(v int32) *ListEmailVerificationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetPrePage(v bool) *ListEmailVerificationResponseBody {
	s.PrePage = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetRequestId(v string) *ListEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetTotalItemNum(v int32) *ListEmailVerificationResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListEmailVerificationResponseBody) SetTotalPageNum(v int32) *ListEmailVerificationResponseBody {
	s.TotalPageNum = &v
	return s
}

type ListEmailVerificationResponseBodyData struct {
	ConfirmIp           *string `json:"ConfirmIp,omitempty" xml:"ConfirmIp,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EmailVerificationNo *string `json:"EmailVerificationNo,omitempty" xml:"EmailVerificationNo,omitempty"`
	GmtCreate           *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	SendIp              *string `json:"SendIp,omitempty" xml:"SendIp,omitempty"`
	TokenSendTime       *string `json:"TokenSendTime,omitempty" xml:"TokenSendTime,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VerificationStatus  *int32  `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
	VerificationTime    *string `json:"VerificationTime,omitempty" xml:"VerificationTime,omitempty"`
}

func (s ListEmailVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEmailVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEmailVerificationResponseBodyData) SetConfirmIp(v string) *ListEmailVerificationResponseBodyData {
	s.ConfirmIp = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetEmail(v string) *ListEmailVerificationResponseBodyData {
	s.Email = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetEmailVerificationNo(v string) *ListEmailVerificationResponseBodyData {
	s.EmailVerificationNo = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetGmtCreate(v string) *ListEmailVerificationResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetGmtModified(v string) *ListEmailVerificationResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetSendIp(v string) *ListEmailVerificationResponseBodyData {
	s.SendIp = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetTokenSendTime(v string) *ListEmailVerificationResponseBodyData {
	s.TokenSendTime = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetUserId(v string) *ListEmailVerificationResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetVerificationStatus(v int32) *ListEmailVerificationResponseBodyData {
	s.VerificationStatus = &v
	return s
}

func (s *ListEmailVerificationResponseBodyData) SetVerificationTime(v string) *ListEmailVerificationResponseBodyData {
	s.VerificationTime = &v
	return s
}

type ListEmailVerificationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEmailVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *ListEmailVerificationResponse) SetHeaders(v map[string]*string) *ListEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *ListEmailVerificationResponse) SetStatusCode(v int32) *ListEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEmailVerificationResponse) SetBody(v *ListEmailVerificationResponseBody) *ListEmailVerificationResponse {
	s.Body = v
	return s
}

type ListServerLockRequest struct {
	BeginStartDate   *int64  `json:"BeginStartDate,omitempty" xml:"BeginStartDate,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndExpireDate    *int64  `json:"EndExpireDate,omitempty" xml:"EndExpireDate,omitempty"`
	EndStartDate     *int64  `json:"EndStartDate,omitempty" xml:"EndStartDate,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	LockProductId    *string `json:"LockProductId,omitempty" xml:"LockProductId,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServerLockStatus *int32  `json:"ServerLockStatus,omitempty" xml:"ServerLockStatus,omitempty"`
	StartExpireDate  *int64  `json:"StartExpireDate,omitempty" xml:"StartExpireDate,omitempty"`
	UserClientIp     *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ListServerLockRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerLockRequest) GoString() string {
	return s.String()
}

func (s *ListServerLockRequest) SetBeginStartDate(v int64) *ListServerLockRequest {
	s.BeginStartDate = &v
	return s
}

func (s *ListServerLockRequest) SetDomainName(v string) *ListServerLockRequest {
	s.DomainName = &v
	return s
}

func (s *ListServerLockRequest) SetEndExpireDate(v int64) *ListServerLockRequest {
	s.EndExpireDate = &v
	return s
}

func (s *ListServerLockRequest) SetEndStartDate(v int64) *ListServerLockRequest {
	s.EndStartDate = &v
	return s
}

func (s *ListServerLockRequest) SetLang(v string) *ListServerLockRequest {
	s.Lang = &v
	return s
}

func (s *ListServerLockRequest) SetLockProductId(v string) *ListServerLockRequest {
	s.LockProductId = &v
	return s
}

func (s *ListServerLockRequest) SetPageNum(v int32) *ListServerLockRequest {
	s.PageNum = &v
	return s
}

func (s *ListServerLockRequest) SetPageSize(v int32) *ListServerLockRequest {
	s.PageSize = &v
	return s
}

func (s *ListServerLockRequest) SetServerLockStatus(v int32) *ListServerLockRequest {
	s.ServerLockStatus = &v
	return s
}

func (s *ListServerLockRequest) SetStartExpireDate(v int64) *ListServerLockRequest {
	s.StartExpireDate = &v
	return s
}

func (s *ListServerLockRequest) SetUserClientIp(v string) *ListServerLockRequest {
	s.UserClientIp = &v
	return s
}

type ListServerLockResponseBody struct {
	CurrentPageNum *int32                            `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListServerLockResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextPage       *bool                             `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                             `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                            `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                            `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListServerLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerLockResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerLockResponseBody) SetCurrentPageNum(v int32) *ListServerLockResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListServerLockResponseBody) SetData(v []*ListServerLockResponseBodyData) *ListServerLockResponseBody {
	s.Data = v
	return s
}

func (s *ListServerLockResponseBody) SetNextPage(v bool) *ListServerLockResponseBody {
	s.NextPage = &v
	return s
}

func (s *ListServerLockResponseBody) SetPageSize(v int32) *ListServerLockResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServerLockResponseBody) SetPrePage(v bool) *ListServerLockResponseBody {
	s.PrePage = &v
	return s
}

func (s *ListServerLockResponseBody) SetRequestId(v string) *ListServerLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerLockResponseBody) SetTotalItemNum(v int32) *ListServerLockResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListServerLockResponseBody) SetTotalPageNum(v int32) *ListServerLockResponseBody {
	s.TotalPageNum = &v
	return s
}

type ListServerLockResponseBodyData struct {
	DomainInstanceId *string `json:"DomainInstanceId,omitempty" xml:"DomainInstanceId,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ExpireDate       *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LockInstanceId   *string `json:"LockInstanceId,omitempty" xml:"LockInstanceId,omitempty"`
	LockProductId    *string `json:"LockProductId,omitempty" xml:"LockProductId,omitempty"`
	ServerLockStatus *string `json:"ServerLockStatus,omitempty" xml:"ServerLockStatus,omitempty"`
	StartDate        *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListServerLockResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServerLockResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServerLockResponseBodyData) SetDomainInstanceId(v string) *ListServerLockResponseBodyData {
	s.DomainInstanceId = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetDomainName(v string) *ListServerLockResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetExpireDate(v string) *ListServerLockResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetGmtCreate(v string) *ListServerLockResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetGmtModified(v string) *ListServerLockResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetLockInstanceId(v string) *ListServerLockResponseBodyData {
	s.LockInstanceId = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetLockProductId(v string) *ListServerLockResponseBodyData {
	s.LockProductId = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetServerLockStatus(v string) *ListServerLockResponseBodyData {
	s.ServerLockStatus = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetStartDate(v string) *ListServerLockResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetUserId(v string) *ListServerLockResponseBodyData {
	s.UserId = &v
	return s
}

type ListServerLockResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServerLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServerLockResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerLockResponse) GoString() string {
	return s.String()
}

func (s *ListServerLockResponse) SetHeaders(v map[string]*string) *ListServerLockResponse {
	s.Headers = v
	return s
}

func (s *ListServerLockResponse) SetStatusCode(v int32) *ListServerLockResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerLockResponse) SetBody(v *ListServerLockResponseBody) *ListServerLockResponse {
	s.Body = v
	return s
}

type LookupTmchNoticeRequest struct {
	ClaimKey     *string `json:"ClaimKey,omitempty" xml:"ClaimKey,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s LookupTmchNoticeRequest) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeRequest) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeRequest) SetClaimKey(v string) *LookupTmchNoticeRequest {
	s.ClaimKey = &v
	return s
}

func (s *LookupTmchNoticeRequest) SetLang(v string) *LookupTmchNoticeRequest {
	s.Lang = &v
	return s
}

func (s *LookupTmchNoticeRequest) SetUserClientIp(v string) *LookupTmchNoticeRequest {
	s.UserClientIp = &v
	return s
}

type LookupTmchNoticeResponseBody struct {
	Claims    *LookupTmchNoticeResponseBodyClaims `json:"Claims,omitempty" xml:"Claims,omitempty" type:"Struct"`
	Id        *int64                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Label     *string                             `json:"Label,omitempty" xml:"Label,omitempty"`
	NotAfter  *string                             `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	NotBefore *string                             `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LookupTmchNoticeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBody) SetClaims(v *LookupTmchNoticeResponseBodyClaims) *LookupTmchNoticeResponseBody {
	s.Claims = v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetId(v int64) *LookupTmchNoticeResponseBody {
	s.Id = &v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetLabel(v string) *LookupTmchNoticeResponseBody {
	s.Label = &v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetNotAfter(v string) *LookupTmchNoticeResponseBody {
	s.NotAfter = &v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetNotBefore(v string) *LookupTmchNoticeResponseBody {
	s.NotBefore = &v
	return s
}

func (s *LookupTmchNoticeResponseBody) SetRequestId(v string) *LookupTmchNoticeResponseBody {
	s.RequestId = &v
	return s
}

type LookupTmchNoticeResponseBodyClaims struct {
	Claim []*LookupTmchNoticeResponseBodyClaimsClaim `json:"Claim,omitempty" xml:"Claim,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaims) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaims) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaims) SetClaim(v []*LookupTmchNoticeResponseBodyClaimsClaim) *LookupTmchNoticeResponseBodyClaims {
	s.Claim = v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaim struct {
	ClassDescs       *LookupTmchNoticeResponseBodyClaimsClaimClassDescs `json:"ClassDescs,omitempty" xml:"ClassDescs,omitempty" type:"Struct"`
	Contacts         *LookupTmchNoticeResponseBodyClaimsClaimContacts   `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	GoodsAndServices *string                                            `json:"GoodsAndServices,omitempty" xml:"GoodsAndServices,omitempty"`
	Holders          *LookupTmchNoticeResponseBodyClaimsClaimHolders    `json:"Holders,omitempty" xml:"Holders,omitempty" type:"Struct"`
	JurDesc          *LookupTmchNoticeResponseBodyClaimsClaimJurDesc    `json:"JurDesc,omitempty" xml:"JurDesc,omitempty" type:"Struct"`
	MarkName         *string                                            `json:"MarkName,omitempty" xml:"MarkName,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaim) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaim) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetClassDescs(v *LookupTmchNoticeResponseBodyClaimsClaimClassDescs) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.ClassDescs = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetContacts(v *LookupTmchNoticeResponseBodyClaimsClaimContacts) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.Contacts = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetGoodsAndServices(v string) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.GoodsAndServices = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetHolders(v *LookupTmchNoticeResponseBodyClaimsClaimHolders) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.Holders = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetJurDesc(v *LookupTmchNoticeResponseBodyClaimsClaimJurDesc) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.JurDesc = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaim) SetMarkName(v string) *LookupTmchNoticeResponseBodyClaimsClaim {
	s.MarkName = &v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimClassDescs struct {
	ClassDesc []*LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc `json:"ClassDesc,omitempty" xml:"ClassDesc,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimClassDescs) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimClassDescs) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescs) SetClassDesc(v []*LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) *LookupTmchNoticeResponseBodyClaimsClaimClassDescs {
	s.ClassDesc = v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc struct {
	ClassNum *int32  `json:"ClassNum,omitempty" xml:"ClassNum,omitempty"`
	Desc     *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) SetClassNum(v int32) *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc {
	s.ClassNum = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc) SetDesc(v string) *LookupTmchNoticeResponseBodyClaimsClaimClassDescsClassDesc {
	s.Desc = &v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimContacts struct {
	Contact []*LookupTmchNoticeResponseBodyClaimsClaimContactsContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContacts) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContacts) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContacts) SetContact(v []*LookupTmchNoticeResponseBodyClaimsClaimContactsContact) *LookupTmchNoticeResponseBodyClaimsClaimContacts {
	s.Contact = v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimContactsContact struct {
	Addr  *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Struct"`
	Email *string                                                     `json:"Email,omitempty" xml:"Email,omitempty"`
	Fax   *string                                                     `json:"Fax,omitempty" xml:"Fax,omitempty"`
	Name  *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Org   *string                                                     `json:"Org,omitempty" xml:"Org,omitempty"`
	Type  *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	Voice *string                                                     `json:"Voice,omitempty" xml:"Voice,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContact) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContact) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetAddr(v *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Addr = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetEmail(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Email = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetFax(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Fax = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetName(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Name = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetOrg(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Org = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetType(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Type = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContact) SetVoice(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContact {
	s.Voice = &v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr struct {
	Cc     *string                                                           `json:"Cc,omitempty" xml:"Cc,omitempty"`
	City   *string                                                           `json:"City,omitempty" xml:"City,omitempty"`
	Pc     *string                                                           `json:"Pc,omitempty" xml:"Pc,omitempty"`
	Sp     *string                                                           `json:"Sp,omitempty" xml:"Sp,omitempty"`
	Street *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet `json:"Street,omitempty" xml:"Street,omitempty" type:"Struct"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetCc(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.Cc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetCity(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.City = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetPc(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.Pc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetSp(v string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.Sp = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr) SetStreet(v *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddr {
	s.Street = v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet struct {
	Street []*string `json:"Street,omitempty" xml:"Street,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet) SetStreet(v []*string) *LookupTmchNoticeResponseBodyClaimsClaimContactsContactAddrStreet {
	s.Street = v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimHolders struct {
	Holder []*LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder `json:"Holder,omitempty" xml:"Holder,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHolders) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHolders) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHolders) SetHolder(v []*LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) *LookupTmchNoticeResponseBodyClaimsClaimHolders {
	s.Holder = v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder struct {
	Addr        *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Struct"`
	Entitlement *string                                                   `json:"Entitlement,omitempty" xml:"Entitlement,omitempty"`
	Org         *string                                                   `json:"Org,omitempty" xml:"Org,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) SetAddr(v *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder {
	s.Addr = v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) SetEntitlement(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder {
	s.Entitlement = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder) SetOrg(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolder {
	s.Org = &v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr struct {
	Cc     *string                                                         `json:"Cc,omitempty" xml:"Cc,omitempty"`
	City   *string                                                         `json:"City,omitempty" xml:"City,omitempty"`
	Pc     *string                                                         `json:"Pc,omitempty" xml:"Pc,omitempty"`
	Sp     *string                                                         `json:"Sp,omitempty" xml:"Sp,omitempty"`
	Street *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet `json:"Street,omitempty" xml:"Street,omitempty" type:"Struct"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetCc(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.Cc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetCity(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.City = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetPc(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.Pc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetSp(v string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.Sp = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr) SetStreet(v *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddr {
	s.Street = v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet struct {
	Street []*string `json:"Street,omitempty" xml:"Street,omitempty" type:"Repeated"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet) SetStreet(v []*string) *LookupTmchNoticeResponseBodyClaimsClaimHoldersHolderAddrStreet {
	s.Street = v
	return s
}

type LookupTmchNoticeResponseBodyClaimsClaimJurDesc struct {
	Desc  *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	JurCC *string `json:"JurCC,omitempty" xml:"JurCC,omitempty"`
}

func (s LookupTmchNoticeResponseBodyClaimsClaimJurDesc) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponseBodyClaimsClaimJurDesc) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimJurDesc) SetDesc(v string) *LookupTmchNoticeResponseBodyClaimsClaimJurDesc {
	s.Desc = &v
	return s
}

func (s *LookupTmchNoticeResponseBodyClaimsClaimJurDesc) SetJurCC(v string) *LookupTmchNoticeResponseBodyClaimsClaimJurDesc {
	s.JurCC = &v
	return s
}

type LookupTmchNoticeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LookupTmchNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LookupTmchNoticeResponse) String() string {
	return tea.Prettify(s)
}

func (s LookupTmchNoticeResponse) GoString() string {
	return s.String()
}

func (s *LookupTmchNoticeResponse) SetHeaders(v map[string]*string) *LookupTmchNoticeResponse {
	s.Headers = v
	return s
}

func (s *LookupTmchNoticeResponse) SetStatusCode(v int32) *LookupTmchNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *LookupTmchNoticeResponse) SetBody(v *LookupTmchNoticeResponseBody) *LookupTmchNoticeResponse {
	s.Body = v
	return s
}

type PollTaskResultRequest struct {
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskNo           *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskResultStatus *int32  `json:"TaskResultStatus,omitempty" xml:"TaskResultStatus,omitempty"`
	UserClientIp     *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s PollTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s PollTaskResultRequest) GoString() string {
	return s.String()
}

func (s *PollTaskResultRequest) SetDomainName(v string) *PollTaskResultRequest {
	s.DomainName = &v
	return s
}

func (s *PollTaskResultRequest) SetInstanceId(v string) *PollTaskResultRequest {
	s.InstanceId = &v
	return s
}

func (s *PollTaskResultRequest) SetLang(v string) *PollTaskResultRequest {
	s.Lang = &v
	return s
}

func (s *PollTaskResultRequest) SetPageNum(v int32) *PollTaskResultRequest {
	s.PageNum = &v
	return s
}

func (s *PollTaskResultRequest) SetPageSize(v int32) *PollTaskResultRequest {
	s.PageSize = &v
	return s
}

func (s *PollTaskResultRequest) SetTaskNo(v string) *PollTaskResultRequest {
	s.TaskNo = &v
	return s
}

func (s *PollTaskResultRequest) SetTaskResultStatus(v int32) *PollTaskResultRequest {
	s.TaskResultStatus = &v
	return s
}

func (s *PollTaskResultRequest) SetUserClientIp(v string) *PollTaskResultRequest {
	s.UserClientIp = &v
	return s
}

type PollTaskResultResponseBody struct {
	CurrentPageNum *int32                          `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *PollTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                           `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                           `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                          `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                          `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s PollTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PollTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *PollTaskResultResponseBody) SetCurrentPageNum(v int32) *PollTaskResultResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *PollTaskResultResponseBody) SetData(v *PollTaskResultResponseBodyData) *PollTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *PollTaskResultResponseBody) SetNextPage(v bool) *PollTaskResultResponseBody {
	s.NextPage = &v
	return s
}

func (s *PollTaskResultResponseBody) SetPageSize(v int32) *PollTaskResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *PollTaskResultResponseBody) SetPrePage(v bool) *PollTaskResultResponseBody {
	s.PrePage = &v
	return s
}

func (s *PollTaskResultResponseBody) SetRequestId(v string) *PollTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *PollTaskResultResponseBody) SetTotalItemNum(v int32) *PollTaskResultResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *PollTaskResultResponseBody) SetTotalPageNum(v int32) *PollTaskResultResponseBody {
	s.TotalPageNum = &v
	return s
}

type PollTaskResultResponseBodyData struct {
	TaskDetail []*PollTaskResultResponseBodyDataTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Repeated"`
}

func (s PollTaskResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PollTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *PollTaskResultResponseBodyData) SetTaskDetail(v []*PollTaskResultResponseBodyDataTaskDetail) *PollTaskResultResponseBodyData {
	s.TaskDetail = v
	return s
}

type PollTaskResultResponseBodyDataTaskDetail struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrorMsg            *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskDetailNo        *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskResult          *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	TryCount            *int32  `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s PollTaskResultResponseBodyDataTaskDetail) String() string {
	return tea.Prettify(s)
}

func (s PollTaskResultResponseBodyDataTaskDetail) GoString() string {
	return s.String()
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetCreateTime(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.CreateTime = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetDomainName(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.DomainName = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetErrorMsg(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.ErrorMsg = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetInstanceId(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.InstanceId = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskDetailNo(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskDetailNo = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskNo(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskNo = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskResult(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskResult = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskStatus(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskStatus = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskStatusCode(v int32) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskStatusCode = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskType(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskType = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskTypeDescription(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskTypeDescription = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTryCount(v int32) *PollTaskResultResponseBodyDataTaskDetail {
	s.TryCount = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetUpdateTime(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.UpdateTime = &v
	return s
}

type PollTaskResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PollTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PollTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s PollTaskResultResponse) GoString() string {
	return s.String()
}

func (s *PollTaskResultResponse) SetHeaders(v map[string]*string) *PollTaskResultResponse {
	s.Headers = v
	return s
}

func (s *PollTaskResultResponse) SetStatusCode(v int32) *PollTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *PollTaskResultResponse) SetBody(v *PollTaskResultResponseBody) *PollTaskResultResponse {
	s.Body = v
	return s
}

type QueryAdvancedDomainListRequest struct {
	DomainGroupId         *int64  `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainNameSort        *bool   `json:"DomainNameSort,omitempty" xml:"DomainNameSort,omitempty"`
	DomainStatus          *int32  `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	EndExpirationDate     *int64  `json:"EndExpirationDate,omitempty" xml:"EndExpirationDate,omitempty"`
	EndLength             *int32  `json:"EndLength,omitempty" xml:"EndLength,omitempty"`
	EndRegistrationDate   *int64  `json:"EndRegistrationDate,omitempty" xml:"EndRegistrationDate,omitempty"`
	Excluded              *string `json:"Excluded,omitempty" xml:"Excluded,omitempty"`
	ExcludedPrefix        *bool   `json:"ExcludedPrefix,omitempty" xml:"ExcludedPrefix,omitempty"`
	ExcludedSuffix        *bool   `json:"ExcludedSuffix,omitempty" xml:"ExcludedSuffix,omitempty"`
	ExpirationDateSort    *bool   `json:"ExpirationDateSort,omitempty" xml:"ExpirationDateSort,omitempty"`
	Form                  *int32  `json:"Form,omitempty" xml:"Form,omitempty"`
	KeyWord               *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	KeyWordPrefix         *bool   `json:"KeyWordPrefix,omitempty" xml:"KeyWordPrefix,omitempty"`
	KeyWordSuffix         *bool   `json:"KeyWordSuffix,omitempty" xml:"KeyWordSuffix,omitempty"`
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum               *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductDomainType     *string `json:"ProductDomainType,omitempty" xml:"ProductDomainType,omitempty"`
	ProductDomainTypeSort *bool   `json:"ProductDomainTypeSort,omitempty" xml:"ProductDomainTypeSort,omitempty"`
	RegistrationDateSort  *bool   `json:"RegistrationDateSort,omitempty" xml:"RegistrationDateSort,omitempty"`
	StartExpirationDate   *int64  `json:"StartExpirationDate,omitempty" xml:"StartExpirationDate,omitempty"`
	StartLength           *int32  `json:"StartLength,omitempty" xml:"StartLength,omitempty"`
	StartRegistrationDate *int64  `json:"StartRegistrationDate,omitempty" xml:"StartRegistrationDate,omitempty"`
	Suffixs               *string `json:"Suffixs,omitempty" xml:"Suffixs,omitempty"`
	TradeType             *int32  `json:"TradeType,omitempty" xml:"TradeType,omitempty"`
	UserClientIp          *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryAdvancedDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedDomainListRequest) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListRequest) SetDomainGroupId(v int64) *QueryAdvancedDomainListRequest {
	s.DomainGroupId = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetDomainNameSort(v bool) *QueryAdvancedDomainListRequest {
	s.DomainNameSort = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetDomainStatus(v int32) *QueryAdvancedDomainListRequest {
	s.DomainStatus = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetEndExpirationDate(v int64) *QueryAdvancedDomainListRequest {
	s.EndExpirationDate = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetEndLength(v int32) *QueryAdvancedDomainListRequest {
	s.EndLength = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetEndRegistrationDate(v int64) *QueryAdvancedDomainListRequest {
	s.EndRegistrationDate = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetExcluded(v string) *QueryAdvancedDomainListRequest {
	s.Excluded = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetExcludedPrefix(v bool) *QueryAdvancedDomainListRequest {
	s.ExcludedPrefix = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetExcludedSuffix(v bool) *QueryAdvancedDomainListRequest {
	s.ExcludedSuffix = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetExpirationDateSort(v bool) *QueryAdvancedDomainListRequest {
	s.ExpirationDateSort = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetForm(v int32) *QueryAdvancedDomainListRequest {
	s.Form = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetKeyWord(v string) *QueryAdvancedDomainListRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetKeyWordPrefix(v bool) *QueryAdvancedDomainListRequest {
	s.KeyWordPrefix = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetKeyWordSuffix(v bool) *QueryAdvancedDomainListRequest {
	s.KeyWordSuffix = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetLang(v string) *QueryAdvancedDomainListRequest {
	s.Lang = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetPageNum(v int32) *QueryAdvancedDomainListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetPageSize(v int32) *QueryAdvancedDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetProductDomainType(v string) *QueryAdvancedDomainListRequest {
	s.ProductDomainType = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetProductDomainTypeSort(v bool) *QueryAdvancedDomainListRequest {
	s.ProductDomainTypeSort = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetRegistrationDateSort(v bool) *QueryAdvancedDomainListRequest {
	s.RegistrationDateSort = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetStartExpirationDate(v int64) *QueryAdvancedDomainListRequest {
	s.StartExpirationDate = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetStartLength(v int32) *QueryAdvancedDomainListRequest {
	s.StartLength = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetStartRegistrationDate(v int64) *QueryAdvancedDomainListRequest {
	s.StartRegistrationDate = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetSuffixs(v string) *QueryAdvancedDomainListRequest {
	s.Suffixs = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetTradeType(v int32) *QueryAdvancedDomainListRequest {
	s.TradeType = &v
	return s
}

func (s *QueryAdvancedDomainListRequest) SetUserClientIp(v string) *QueryAdvancedDomainListRequest {
	s.UserClientIp = &v
	return s
}

type QueryAdvancedDomainListResponseBody struct {
	CurrentPageNum *int32                                   `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryAdvancedDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                                    `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                                    `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                                   `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                   `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryAdvancedDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBody) SetCurrentPageNum(v int32) *QueryAdvancedDomainListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetData(v *QueryAdvancedDomainListResponseBodyData) *QueryAdvancedDomainListResponseBody {
	s.Data = v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetNextPage(v bool) *QueryAdvancedDomainListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetPageSize(v int32) *QueryAdvancedDomainListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetPrePage(v bool) *QueryAdvancedDomainListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetRequestId(v string) *QueryAdvancedDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetTotalItemNum(v int32) *QueryAdvancedDomainListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBody) SetTotalPageNum(v int32) *QueryAdvancedDomainListResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryAdvancedDomainListResponseBodyData struct {
	Domain []*QueryAdvancedDomainListResponseBodyDataDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s QueryAdvancedDomainListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBodyData) SetDomain(v []*QueryAdvancedDomainListResponseBodyDataDomain) *QueryAdvancedDomainListResponseBodyData {
	s.Domain = v
	return s
}

type QueryAdvancedDomainListResponseBodyDataDomain struct {
	DnsList                  *QueryAdvancedDomainListResponseBodyDataDomainDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	DomainAuditStatus        *string                                               `json:"DomainAuditStatus,omitempty" xml:"DomainAuditStatus,omitempty"`
	DomainGroupId            *string                                               `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName          *string                                               `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	DomainName               *string                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus             *string                                               `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType               *string                                               `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	Email                    *string                                               `json:"Email,omitempty" xml:"Email,omitempty"`
	ExpirationCurrDateDiff   *int32                                                `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	ExpirationDate           *string                                               `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong       *int64                                                `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	ExpirationDateStatus     *string                                               `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	InstanceId               *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Premium                  *bool                                                 `json:"Premium,omitempty" xml:"Premium,omitempty"`
	ProductId                *string                                               `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RegistrantOrganization   *string                                               `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantType           *string                                               `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	RegistrationDate         *string                                               `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	RegistrationDateLong     *int64                                                `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	Remark                   *string                                               `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ZhRegistrantOrganization *string                                               `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryAdvancedDomainListResponseBodyDataDomain) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDnsList(v *QueryAdvancedDomainListResponseBodyDataDomainDnsList) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DnsList = v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainAuditStatus(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainAuditStatus = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainGroupId(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainGroupId = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainGroupName(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainGroupName = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainName(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainStatus(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainStatus = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetDomainType(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.DomainType = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetEmail(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.Email = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetExpirationCurrDateDiff(v int32) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ExpirationCurrDateDiff = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetExpirationDate(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ExpirationDate = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetExpirationDateLong(v int64) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetExpirationDateStatus(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ExpirationDateStatus = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetInstanceId(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.InstanceId = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetPremium(v bool) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.Premium = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetProductId(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ProductId = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRegistrantOrganization(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRegistrantType(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.RegistrantType = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRegistrationDate(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.RegistrationDate = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRegistrationDateLong(v int64) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.RegistrationDateLong = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetRemark(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.Remark = &v
	return s
}

func (s *QueryAdvancedDomainListResponseBodyDataDomain) SetZhRegistrantOrganization(v string) *QueryAdvancedDomainListResponseBodyDataDomain {
	s.ZhRegistrantOrganization = &v
	return s
}

type QueryAdvancedDomainListResponseBodyDataDomainDnsList struct {
	Dns []*string `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Repeated"`
}

func (s QueryAdvancedDomainListResponseBodyDataDomainDnsList) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedDomainListResponseBodyDataDomainDnsList) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponseBodyDataDomainDnsList) SetDns(v []*string) *QueryAdvancedDomainListResponseBodyDataDomainDnsList {
	s.Dns = v
	return s
}

type QueryAdvancedDomainListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAdvancedDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAdvancedDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvancedDomainListResponse) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponse) SetHeaders(v map[string]*string) *QueryAdvancedDomainListResponse {
	s.Headers = v
	return s
}

func (s *QueryAdvancedDomainListResponse) SetStatusCode(v int32) *QueryAdvancedDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAdvancedDomainListResponse) SetBody(v *QueryAdvancedDomainListResponseBody) *QueryAdvancedDomainListResponse {
	s.Body = v
	return s
}

type QueryArtExtensionRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryArtExtensionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryArtExtensionRequest) GoString() string {
	return s.String()
}

func (s *QueryArtExtensionRequest) SetDomainName(v string) *QueryArtExtensionRequest {
	s.DomainName = &v
	return s
}

func (s *QueryArtExtensionRequest) SetLang(v string) *QueryArtExtensionRequest {
	s.Lang = &v
	return s
}

func (s *QueryArtExtensionRequest) SetUserClientIp(v string) *QueryArtExtensionRequest {
	s.UserClientIp = &v
	return s
}

type QueryArtExtensionResponseBody struct {
	DateOrPeriod            *string `json:"DateOrPeriod,omitempty" xml:"DateOrPeriod,omitempty"`
	Dimensions              *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Features                *string `json:"Features,omitempty" xml:"Features,omitempty"`
	InscriptionsAndMarkings *string `json:"InscriptionsAndMarkings,omitempty" xml:"InscriptionsAndMarkings,omitempty"`
	Maker                   *string `json:"Maker,omitempty" xml:"Maker,omitempty"`
	MaterialsAndTechniques  *string `json:"MaterialsAndTechniques,omitempty" xml:"MaterialsAndTechniques,omitempty"`
	ObjectType              *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Reference               *string `json:"Reference,omitempty" xml:"Reference,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Subject                 *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Title                   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryArtExtensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryArtExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryArtExtensionResponseBody) SetDateOrPeriod(v string) *QueryArtExtensionResponseBody {
	s.DateOrPeriod = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetDimensions(v string) *QueryArtExtensionResponseBody {
	s.Dimensions = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetFeatures(v string) *QueryArtExtensionResponseBody {
	s.Features = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetInscriptionsAndMarkings(v string) *QueryArtExtensionResponseBody {
	s.InscriptionsAndMarkings = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetMaker(v string) *QueryArtExtensionResponseBody {
	s.Maker = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetMaterialsAndTechniques(v string) *QueryArtExtensionResponseBody {
	s.MaterialsAndTechniques = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetObjectType(v string) *QueryArtExtensionResponseBody {
	s.ObjectType = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetReference(v string) *QueryArtExtensionResponseBody {
	s.Reference = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetRequestId(v string) *QueryArtExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetSubject(v string) *QueryArtExtensionResponseBody {
	s.Subject = &v
	return s
}

func (s *QueryArtExtensionResponseBody) SetTitle(v string) *QueryArtExtensionResponseBody {
	s.Title = &v
	return s
}

type QueryArtExtensionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryArtExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryArtExtensionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryArtExtensionResponse) GoString() string {
	return s.String()
}

func (s *QueryArtExtensionResponse) SetHeaders(v map[string]*string) *QueryArtExtensionResponse {
	s.Headers = v
	return s
}

func (s *QueryArtExtensionResponse) SetStatusCode(v int32) *QueryArtExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryArtExtensionResponse) SetBody(v *QueryArtExtensionResponseBody) *QueryArtExtensionResponse {
	s.Body = v
	return s
}

type QueryChangeLogListRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndDate      *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate    *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryChangeLogListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeLogListRequest) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListRequest) SetDomainName(v string) *QueryChangeLogListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryChangeLogListRequest) SetEndDate(v int64) *QueryChangeLogListRequest {
	s.EndDate = &v
	return s
}

func (s *QueryChangeLogListRequest) SetLang(v string) *QueryChangeLogListRequest {
	s.Lang = &v
	return s
}

func (s *QueryChangeLogListRequest) SetPageNum(v int32) *QueryChangeLogListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryChangeLogListRequest) SetPageSize(v int32) *QueryChangeLogListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryChangeLogListRequest) SetStartDate(v int64) *QueryChangeLogListRequest {
	s.StartDate = &v
	return s
}

func (s *QueryChangeLogListRequest) SetUserClientIp(v string) *QueryChangeLogListRequest {
	s.UserClientIp = &v
	return s
}

type QueryChangeLogListResponseBody struct {
	CurrentPageNum *int32                              `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryChangeLogListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                               `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                               `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultLimit    *bool                               `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	TotalItemNum   *int32                              `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                              `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryChangeLogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeLogListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListResponseBody) SetCurrentPageNum(v int32) *QueryChangeLogListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetData(v *QueryChangeLogListResponseBodyData) *QueryChangeLogListResponseBody {
	s.Data = v
	return s
}

func (s *QueryChangeLogListResponseBody) SetNextPage(v bool) *QueryChangeLogListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetPageSize(v int32) *QueryChangeLogListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetPrePage(v bool) *QueryChangeLogListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetRequestId(v string) *QueryChangeLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetResultLimit(v bool) *QueryChangeLogListResponseBody {
	s.ResultLimit = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetTotalItemNum(v int32) *QueryChangeLogListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryChangeLogListResponseBody) SetTotalPageNum(v int32) *QueryChangeLogListResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryChangeLogListResponseBodyData struct {
	ChangeLog []*QueryChangeLogListResponseBodyDataChangeLog `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty" type:"Repeated"`
}

func (s QueryChangeLogListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeLogListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListResponseBodyData) SetChangeLog(v []*QueryChangeLogListResponseBodyDataChangeLog) *QueryChangeLogListResponseBodyData {
	s.ChangeLog = v
	return s
}

type QueryChangeLogListResponseBodyDataChangeLog struct {
	Details            *string `json:"Details,omitempty" xml:"Details,omitempty"`
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Operation          *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	OperationIPAddress *string `json:"OperationIPAddress,omitempty" xml:"OperationIPAddress,omitempty"`
	Result             *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Time               *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryChangeLogListResponseBodyDataChangeLog) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeLogListResponseBodyDataChangeLog) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetDetails(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.Details = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetDomainName(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.DomainName = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetOperation(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.Operation = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetOperationIPAddress(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.OperationIPAddress = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetResult(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.Result = &v
	return s
}

func (s *QueryChangeLogListResponseBodyDataChangeLog) SetTime(v string) *QueryChangeLogListResponseBodyDataChangeLog {
	s.Time = &v
	return s
}

type QueryChangeLogListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryChangeLogListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryChangeLogListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryChangeLogListResponse) GoString() string {
	return s.String()
}

func (s *QueryChangeLogListResponse) SetHeaders(v map[string]*string) *QueryChangeLogListResponse {
	s.Headers = v
	return s
}

func (s *QueryChangeLogListResponse) SetStatusCode(v int32) *QueryChangeLogListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChangeLogListResponse) SetBody(v *QueryChangeLogListResponseBody) *QueryChangeLogListResponse {
	s.Body = v
	return s
}

type QueryContactInfoRequest struct {
	ContactType  *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryContactInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryContactInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryContactInfoRequest) SetContactType(v string) *QueryContactInfoRequest {
	s.ContactType = &v
	return s
}

func (s *QueryContactInfoRequest) SetDomainName(v string) *QueryContactInfoRequest {
	s.DomainName = &v
	return s
}

func (s *QueryContactInfoRequest) SetLang(v string) *QueryContactInfoRequest {
	s.Lang = &v
	return s
}

func (s *QueryContactInfoRequest) SetUserClientIp(v string) *QueryContactInfoRequest {
	s.UserClientIp = &v
	return s
}

type QueryContactInfoResponseBody struct {
	Address                  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	City                     *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	CreateDate               *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Email                    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PostalCode               *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                 *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName           *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization   *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TelArea                  *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                   *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	ZhAddress                *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryContactInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContactInfoResponseBody) SetAddress(v string) *QueryContactInfoResponseBody {
	s.Address = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetCity(v string) *QueryContactInfoResponseBody {
	s.City = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetCountry(v string) *QueryContactInfoResponseBody {
	s.Country = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetCreateDate(v string) *QueryContactInfoResponseBody {
	s.CreateDate = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetEmail(v string) *QueryContactInfoResponseBody {
	s.Email = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetPostalCode(v string) *QueryContactInfoResponseBody {
	s.PostalCode = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetProvince(v string) *QueryContactInfoResponseBody {
	s.Province = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetRegistrantName(v string) *QueryContactInfoResponseBody {
	s.RegistrantName = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetRegistrantOrganization(v string) *QueryContactInfoResponseBody {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetRequestId(v string) *QueryContactInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetTelArea(v string) *QueryContactInfoResponseBody {
	s.TelArea = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetTelExt(v string) *QueryContactInfoResponseBody {
	s.TelExt = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetTelephone(v string) *QueryContactInfoResponseBody {
	s.Telephone = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhAddress(v string) *QueryContactInfoResponseBody {
	s.ZhAddress = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhCity(v string) *QueryContactInfoResponseBody {
	s.ZhCity = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhProvince(v string) *QueryContactInfoResponseBody {
	s.ZhProvince = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhRegistrantName(v string) *QueryContactInfoResponseBody {
	s.ZhRegistrantName = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhRegistrantOrganization(v string) *QueryContactInfoResponseBody {
	s.ZhRegistrantOrganization = &v
	return s
}

type QueryContactInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryContactInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryContactInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryContactInfoResponse) SetHeaders(v map[string]*string) *QueryContactInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryContactInfoResponse) SetStatusCode(v int32) *QueryContactInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContactInfoResponse) SetBody(v *QueryContactInfoResponseBody) *QueryContactInfoResponse {
	s.Body = v
	return s
}

type QueryDSRecordRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDSRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDSRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryDSRecordRequest) SetDomainName(v string) *QueryDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDSRecordRequest) SetLang(v string) *QueryDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *QueryDSRecordRequest) SetUserClientIp(v string) *QueryDSRecordRequest {
	s.UserClientIp = &v
	return s
}

type QueryDSRecordResponseBody struct {
	DSRecordList []*QueryDSRecordResponseBodyDSRecordList `json:"DSRecordList,omitempty" xml:"DSRecordList,omitempty" type:"Repeated"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDSRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDSRecordResponseBody) SetDSRecordList(v []*QueryDSRecordResponseBodyDSRecordList) *QueryDSRecordResponseBody {
	s.DSRecordList = v
	return s
}

func (s *QueryDSRecordResponseBody) SetRequestId(v string) *QueryDSRecordResponseBody {
	s.RequestId = &v
	return s
}

type QueryDSRecordResponseBodyDSRecordList struct {
	Algorithm  *int32  `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Digest     *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	DigestType *int32  `json:"DigestType,omitempty" xml:"DigestType,omitempty"`
	KeyTag     *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
}

func (s QueryDSRecordResponseBodyDSRecordList) String() string {
	return tea.Prettify(s)
}

func (s QueryDSRecordResponseBodyDSRecordList) GoString() string {
	return s.String()
}

func (s *QueryDSRecordResponseBodyDSRecordList) SetAlgorithm(v int32) *QueryDSRecordResponseBodyDSRecordList {
	s.Algorithm = &v
	return s
}

func (s *QueryDSRecordResponseBodyDSRecordList) SetDigest(v string) *QueryDSRecordResponseBodyDSRecordList {
	s.Digest = &v
	return s
}

func (s *QueryDSRecordResponseBodyDSRecordList) SetDigestType(v int32) *QueryDSRecordResponseBodyDSRecordList {
	s.DigestType = &v
	return s
}

func (s *QueryDSRecordResponseBodyDSRecordList) SetKeyTag(v int32) *QueryDSRecordResponseBodyDSRecordList {
	s.KeyTag = &v
	return s
}

type QueryDSRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDSRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDSRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryDSRecordResponse) SetHeaders(v map[string]*string) *QueryDSRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryDSRecordResponse) SetStatusCode(v int32) *QueryDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDSRecordResponse) SetBody(v *QueryDSRecordResponseBody) *QueryDSRecordResponse {
	s.Body = v
	return s
}

type QueryDnsHostRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDnsHostRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDnsHostRequest) GoString() string {
	return s.String()
}

func (s *QueryDnsHostRequest) SetInstanceId(v string) *QueryDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryDnsHostRequest) SetLang(v string) *QueryDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *QueryDnsHostRequest) SetUserClientIp(v string) *QueryDnsHostRequest {
	s.UserClientIp = &v
	return s
}

type QueryDnsHostResponseBody struct {
	DnsHostList []*QueryDnsHostResponseBodyDnsHostList `json:"DnsHostList,omitempty" xml:"DnsHostList,omitempty" type:"Repeated"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDnsHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDnsHostResponseBody) SetDnsHostList(v []*QueryDnsHostResponseBodyDnsHostList) *QueryDnsHostResponseBody {
	s.DnsHostList = v
	return s
}

func (s *QueryDnsHostResponseBody) SetRequestId(v string) *QueryDnsHostResponseBody {
	s.RequestId = &v
	return s
}

type QueryDnsHostResponseBodyDnsHostList struct {
	DnsName *string   `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	IpList  []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
}

func (s QueryDnsHostResponseBodyDnsHostList) String() string {
	return tea.Prettify(s)
}

func (s QueryDnsHostResponseBodyDnsHostList) GoString() string {
	return s.String()
}

func (s *QueryDnsHostResponseBodyDnsHostList) SetDnsName(v string) *QueryDnsHostResponseBodyDnsHostList {
	s.DnsName = &v
	return s
}

func (s *QueryDnsHostResponseBodyDnsHostList) SetIpList(v []*string) *QueryDnsHostResponseBodyDnsHostList {
	s.IpList = v
	return s
}

type QueryDnsHostResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDnsHostResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDnsHostResponse) GoString() string {
	return s.String()
}

func (s *QueryDnsHostResponse) SetHeaders(v map[string]*string) *QueryDnsHostResponse {
	s.Headers = v
	return s
}

func (s *QueryDnsHostResponse) SetStatusCode(v int32) *QueryDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDnsHostResponse) SetBody(v *QueryDnsHostResponseBody) *QueryDnsHostResponse {
	s.Body = v
	return s
}

type QueryDomainAdminDivisionRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainAdminDivisionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainAdminDivisionRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionRequest) SetLang(v string) *QueryDomainAdminDivisionRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainAdminDivisionRequest) SetUserClientIp(v string) *QueryDomainAdminDivisionRequest {
	s.UserClientIp = &v
	return s
}

type QueryDomainAdminDivisionResponseBody struct {
	AdminDivisions *QueryDomainAdminDivisionResponseBodyAdminDivisions `json:"AdminDivisions,omitempty" xml:"AdminDivisions,omitempty" type:"Struct"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDomainAdminDivisionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBody) SetAdminDivisions(v *QueryDomainAdminDivisionResponseBodyAdminDivisions) *QueryDomainAdminDivisionResponseBody {
	s.AdminDivisions = v
	return s
}

func (s *QueryDomainAdminDivisionResponseBody) SetRequestId(v string) *QueryDomainAdminDivisionResponseBody {
	s.RequestId = &v
	return s
}

type QueryDomainAdminDivisionResponseBodyAdminDivisions struct {
	AdminDivision []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision `json:"AdminDivision,omitempty" xml:"AdminDivision,omitempty" type:"Repeated"`
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisions) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisions) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisions) SetAdminDivision(v []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) *QueryDomainAdminDivisionResponseBodyAdminDivisions {
	s.AdminDivision = v
	return s
}

type QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision struct {
	Children     *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Struct"`
	DivisionName *string                                                                  `json:"DivisionName,omitempty" xml:"DivisionName,omitempty"`
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) SetChildren(v *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision {
	s.Children = v
	return s
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision) SetDivisionName(v string) *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivision {
	s.DivisionName = &v
	return s
}

type QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren struct {
	Children []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren) SetChildren(v []*QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildren {
	s.Children = v
	return s
}

type QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren struct {
	ChildDivisionName *string `json:"ChildDivisionName,omitempty" xml:"ChildDivisionName,omitempty"`
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren) SetChildDivisionName(v string) *QueryDomainAdminDivisionResponseBodyAdminDivisionsAdminDivisionChildrenChildren {
	s.ChildDivisionName = &v
	return s
}

type QueryDomainAdminDivisionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDomainAdminDivisionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDomainAdminDivisionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainAdminDivisionResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionResponse) SetHeaders(v map[string]*string) *QueryDomainAdminDivisionResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainAdminDivisionResponse) SetStatusCode(v int32) *QueryDomainAdminDivisionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainAdminDivisionResponse) SetBody(v *QueryDomainAdminDivisionResponseBody) *QueryDomainAdminDivisionResponse {
	s.Body = v
	return s
}

type QueryDomainByDomainNameRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainByDomainNameRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByDomainNameRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameRequest) SetDomainName(v string) *QueryDomainByDomainNameRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDomainByDomainNameRequest) SetLang(v string) *QueryDomainByDomainNameRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainByDomainNameRequest) SetUserClientIp(v string) *QueryDomainByDomainNameRequest {
	s.UserClientIp = &v
	return s
}

type QueryDomainByDomainNameResponseBody struct {
	DnsList                      *QueryDomainByDomainNameResponseBodyDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	DomainGroupId                *int64                                      `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName              *string                                     `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	DomainName                   *string                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameProxyService       *bool                                       `json:"DomainNameProxyService,omitempty" xml:"DomainNameProxyService,omitempty"`
	DomainNameVerificationStatus *string                                     `json:"DomainNameVerificationStatus,omitempty" xml:"DomainNameVerificationStatus,omitempty"`
	DomainStatus                 *string                                     `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType                   *string                                     `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	Email                        *string                                     `json:"Email,omitempty" xml:"Email,omitempty"`
	EmailVerificationClientHold  *bool                                       `json:"EmailVerificationClientHold,omitempty" xml:"EmailVerificationClientHold,omitempty"`
	EmailVerificationStatus      *int32                                      `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	ExpirationCurrDateDiff       *int32                                      `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	ExpirationDate               *string                                     `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong           *int64                                      `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	ExpirationDateStatus         *string                                     `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	InstanceId                   *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Premium                      *bool                                       `json:"Premium,omitempty" xml:"Premium,omitempty"`
	RealNameStatus               *string                                     `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	RegistrantName               *string                                     `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization       *string                                     `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantType               *string                                     `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	RegistrantUpdatingStatus     *string                                     `json:"RegistrantUpdatingStatus,omitempty" xml:"RegistrantUpdatingStatus,omitempty"`
	RegistrationDate             *string                                     `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	RegistrationDateLong         *int64                                      `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	Remark                       *string                                     `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RequestId                    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransferOutStatus            *string                                     `json:"TransferOutStatus,omitempty" xml:"TransferOutStatus,omitempty"`
	TransferProhibitionLock      *string                                     `json:"TransferProhibitionLock,omitempty" xml:"TransferProhibitionLock,omitempty"`
	UpdateProhibitionLock        *string                                     `json:"UpdateProhibitionLock,omitempty" xml:"UpdateProhibitionLock,omitempty"`
	UserId                       *string                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ZhRegistrantName             *string                                     `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization     *string                                     `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryDomainByDomainNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByDomainNameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameResponseBody) SetDnsList(v *QueryDomainByDomainNameResponseBodyDnsList) *QueryDomainByDomainNameResponseBody {
	s.DnsList = v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainGroupId(v int64) *QueryDomainByDomainNameResponseBody {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainGroupName(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainGroupName = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainName(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainNameProxyService(v bool) *QueryDomainByDomainNameResponseBody {
	s.DomainNameProxyService = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainNameVerificationStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainNameVerificationStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetDomainType(v string) *QueryDomainByDomainNameResponseBody {
	s.DomainType = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetEmail(v string) *QueryDomainByDomainNameResponseBody {
	s.Email = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetEmailVerificationClientHold(v bool) *QueryDomainByDomainNameResponseBody {
	s.EmailVerificationClientHold = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetEmailVerificationStatus(v int32) *QueryDomainByDomainNameResponseBody {
	s.EmailVerificationStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetExpirationCurrDateDiff(v int32) *QueryDomainByDomainNameResponseBody {
	s.ExpirationCurrDateDiff = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetExpirationDate(v string) *QueryDomainByDomainNameResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetExpirationDateLong(v int64) *QueryDomainByDomainNameResponseBody {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetExpirationDateStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.ExpirationDateStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetInstanceId(v string) *QueryDomainByDomainNameResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetPremium(v bool) *QueryDomainByDomainNameResponseBody {
	s.Premium = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRealNameStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.RealNameStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrantName(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrantName = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrantOrganization(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrantType(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrantType = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrantUpdatingStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrantUpdatingStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrationDate(v string) *QueryDomainByDomainNameResponseBody {
	s.RegistrationDate = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRegistrationDateLong(v int64) *QueryDomainByDomainNameResponseBody {
	s.RegistrationDateLong = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRemark(v string) *QueryDomainByDomainNameResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetRequestId(v string) *QueryDomainByDomainNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetTransferOutStatus(v string) *QueryDomainByDomainNameResponseBody {
	s.TransferOutStatus = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetTransferProhibitionLock(v string) *QueryDomainByDomainNameResponseBody {
	s.TransferProhibitionLock = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetUpdateProhibitionLock(v string) *QueryDomainByDomainNameResponseBody {
	s.UpdateProhibitionLock = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetUserId(v string) *QueryDomainByDomainNameResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetZhRegistrantName(v string) *QueryDomainByDomainNameResponseBody {
	s.ZhRegistrantName = &v
	return s
}

func (s *QueryDomainByDomainNameResponseBody) SetZhRegistrantOrganization(v string) *QueryDomainByDomainNameResponseBody {
	s.ZhRegistrantOrganization = &v
	return s
}

type QueryDomainByDomainNameResponseBodyDnsList struct {
	Dns []*string `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Repeated"`
}

func (s QueryDomainByDomainNameResponseBodyDnsList) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByDomainNameResponseBodyDnsList) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameResponseBodyDnsList) SetDns(v []*string) *QueryDomainByDomainNameResponseBodyDnsList {
	s.Dns = v
	return s
}

type QueryDomainByDomainNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDomainByDomainNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDomainByDomainNameResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByDomainNameResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameResponse) SetHeaders(v map[string]*string) *QueryDomainByDomainNameResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainByDomainNameResponse) SetStatusCode(v int32) *QueryDomainByDomainNameResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainByDomainNameResponse) SetBody(v *QueryDomainByDomainNameResponseBody) *QueryDomainByDomainNameResponse {
	s.Body = v
	return s
}

type QueryDomainByInstanceIdRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainByInstanceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdRequest) SetInstanceId(v string) *QueryDomainByInstanceIdRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainByInstanceIdRequest) SetLang(v string) *QueryDomainByInstanceIdRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainByInstanceIdRequest) SetUserClientIp(v string) *QueryDomainByInstanceIdRequest {
	s.UserClientIp = &v
	return s
}

type QueryDomainByInstanceIdResponseBody struct {
	DnsList                      *QueryDomainByInstanceIdResponseBodyDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	DomainGroupId                *int64                                      `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName              *string                                     `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	DomainName                   *string                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameProxyService       *bool                                       `json:"DomainNameProxyService,omitempty" xml:"DomainNameProxyService,omitempty"`
	DomainNameVerificationStatus *string                                     `json:"DomainNameVerificationStatus,omitempty" xml:"DomainNameVerificationStatus,omitempty"`
	DomainStatus                 *string                                     `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType                   *string                                     `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	Email                        *string                                     `json:"Email,omitempty" xml:"Email,omitempty"`
	EmailVerificationClientHold  *bool                                       `json:"EmailVerificationClientHold,omitempty" xml:"EmailVerificationClientHold,omitempty"`
	EmailVerificationStatus      *int32                                      `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	ExpirationCurrDateDiff       *int32                                      `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	ExpirationDate               *string                                     `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong           *int64                                      `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	ExpirationDateStatus         *string                                     `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	InstanceId                   *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Premium                      *bool                                       `json:"Premium,omitempty" xml:"Premium,omitempty"`
	RealNameStatus               *string                                     `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	RegistrantName               *string                                     `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization       *string                                     `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantType               *string                                     `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	RegistrantUpdatingStatus     *string                                     `json:"RegistrantUpdatingStatus,omitempty" xml:"RegistrantUpdatingStatus,omitempty"`
	RegistrationDate             *string                                     `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	RegistrationDateLong         *int64                                      `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	Remark                       *string                                     `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RequestId                    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransferOutStatus            *string                                     `json:"TransferOutStatus,omitempty" xml:"TransferOutStatus,omitempty"`
	TransferProhibitionLock      *string                                     `json:"TransferProhibitionLock,omitempty" xml:"TransferProhibitionLock,omitempty"`
	UpdateProhibitionLock        *string                                     `json:"UpdateProhibitionLock,omitempty" xml:"UpdateProhibitionLock,omitempty"`
	UserId                       *string                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ZhRegistrantName             *string                                     `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization     *string                                     `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryDomainByInstanceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdResponseBody) SetDnsList(v *QueryDomainByInstanceIdResponseBodyDnsList) *QueryDomainByInstanceIdResponseBody {
	s.DnsList = v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainGroupId(v int64) *QueryDomainByInstanceIdResponseBody {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainGroupName(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainGroupName = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainName(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainNameProxyService(v bool) *QueryDomainByInstanceIdResponseBody {
	s.DomainNameProxyService = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainNameVerificationStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainNameVerificationStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetDomainType(v string) *QueryDomainByInstanceIdResponseBody {
	s.DomainType = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetEmail(v string) *QueryDomainByInstanceIdResponseBody {
	s.Email = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetEmailVerificationClientHold(v bool) *QueryDomainByInstanceIdResponseBody {
	s.EmailVerificationClientHold = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetEmailVerificationStatus(v int32) *QueryDomainByInstanceIdResponseBody {
	s.EmailVerificationStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetExpirationCurrDateDiff(v int32) *QueryDomainByInstanceIdResponseBody {
	s.ExpirationCurrDateDiff = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetExpirationDate(v string) *QueryDomainByInstanceIdResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetExpirationDateLong(v int64) *QueryDomainByInstanceIdResponseBody {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetExpirationDateStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.ExpirationDateStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetInstanceId(v string) *QueryDomainByInstanceIdResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetPremium(v bool) *QueryDomainByInstanceIdResponseBody {
	s.Premium = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRealNameStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.RealNameStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrantName(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrantName = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrantOrganization(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrantType(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrantType = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrantUpdatingStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrantUpdatingStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrationDate(v string) *QueryDomainByInstanceIdResponseBody {
	s.RegistrationDate = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRegistrationDateLong(v int64) *QueryDomainByInstanceIdResponseBody {
	s.RegistrationDateLong = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRemark(v string) *QueryDomainByInstanceIdResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetRequestId(v string) *QueryDomainByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetTransferOutStatus(v string) *QueryDomainByInstanceIdResponseBody {
	s.TransferOutStatus = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetTransferProhibitionLock(v string) *QueryDomainByInstanceIdResponseBody {
	s.TransferProhibitionLock = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetUpdateProhibitionLock(v string) *QueryDomainByInstanceIdResponseBody {
	s.UpdateProhibitionLock = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetUserId(v string) *QueryDomainByInstanceIdResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetZhRegistrantName(v string) *QueryDomainByInstanceIdResponseBody {
	s.ZhRegistrantName = &v
	return s
}

func (s *QueryDomainByInstanceIdResponseBody) SetZhRegistrantOrganization(v string) *QueryDomainByInstanceIdResponseBody {
	s.ZhRegistrantOrganization = &v
	return s
}

type QueryDomainByInstanceIdResponseBodyDnsList struct {
	Dns []*string `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Repeated"`
}

func (s QueryDomainByInstanceIdResponseBodyDnsList) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByInstanceIdResponseBodyDnsList) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdResponseBodyDnsList) SetDns(v []*string) *QueryDomainByInstanceIdResponseBodyDnsList {
	s.Dns = v
	return s
}

type QueryDomainByInstanceIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDomainByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDomainByInstanceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainByInstanceIdResponse) SetHeaders(v map[string]*string) *QueryDomainByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainByInstanceIdResponse) SetStatusCode(v int32) *QueryDomainByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainByInstanceIdResponse) SetBody(v *QueryDomainByInstanceIdResponseBody) *QueryDomainByInstanceIdResponse {
	s.Body = v
	return s
}

type QueryDomainGroupListRequest struct {
	DomainGroupName   *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ShowDeletingGroup *bool   `json:"ShowDeletingGroup,omitempty" xml:"ShowDeletingGroup,omitempty"`
	UserClientIp      *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainGroupListRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListRequest) SetDomainGroupName(v string) *QueryDomainGroupListRequest {
	s.DomainGroupName = &v
	return s
}

func (s *QueryDomainGroupListRequest) SetLang(v string) *QueryDomainGroupListRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainGroupListRequest) SetShowDeletingGroup(v bool) *QueryDomainGroupListRequest {
	s.ShowDeletingGroup = &v
	return s
}

func (s *QueryDomainGroupListRequest) SetUserClientIp(v string) *QueryDomainGroupListRequest {
	s.UserClientIp = &v
	return s
}

type QueryDomainGroupListResponseBody struct {
	Data      *QueryDomainGroupListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDomainGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListResponseBody) SetData(v *QueryDomainGroupListResponseBodyData) *QueryDomainGroupListResponseBody {
	s.Data = v
	return s
}

func (s *QueryDomainGroupListResponseBody) SetRequestId(v string) *QueryDomainGroupListResponseBody {
	s.RequestId = &v
	return s
}

type QueryDomainGroupListResponseBodyData struct {
	DomainGroup []*QueryDomainGroupListResponseBodyDataDomainGroup `json:"DomainGroup,omitempty" xml:"DomainGroup,omitempty" type:"Repeated"`
}

func (s QueryDomainGroupListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainGroupListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListResponseBodyData) SetDomainGroup(v []*QueryDomainGroupListResponseBodyDataDomainGroup) *QueryDomainGroupListResponseBodyData {
	s.DomainGroup = v
	return s
}

type QueryDomainGroupListResponseBodyDataDomainGroup struct {
	BeingDeleted      *bool   `json:"BeingDeleted,omitempty" xml:"BeingDeleted,omitempty"`
	CreationDate      *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	DomainGroupId     *string `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName   *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	DomainGroupStatus *string `json:"DomainGroupStatus,omitempty" xml:"DomainGroupStatus,omitempty"`
	ModificationDate  *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	TotalNumber       *int32  `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s QueryDomainGroupListResponseBodyDataDomainGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainGroupListResponseBodyDataDomainGroup) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetBeingDeleted(v bool) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.BeingDeleted = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetCreationDate(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.CreationDate = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetDomainGroupId(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetDomainGroupName(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.DomainGroupName = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetDomainGroupStatus(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.DomainGroupStatus = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetModificationDate(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.ModificationDate = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetTotalNumber(v int32) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.TotalNumber = &v
	return s
}

type QueryDomainGroupListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDomainGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDomainGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainGroupListResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListResponse) SetHeaders(v map[string]*string) *QueryDomainGroupListResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainGroupListResponse) SetStatusCode(v int32) *QueryDomainGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainGroupListResponse) SetBody(v *QueryDomainGroupListResponseBody) *QueryDomainGroupListResponse {
	s.Body = v
	return s
}

type QueryDomainListRequest struct {
	DomainGroupId         *string `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndExpirationDate     *int64  `json:"EndExpirationDate,omitempty" xml:"EndExpirationDate,omitempty"`
	EndRegistrationDate   *int64  `json:"EndRegistrationDate,omitempty" xml:"EndRegistrationDate,omitempty"`
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OrderByType           *string `json:"OrderByType,omitempty" xml:"OrderByType,omitempty"`
	OrderKeyType          *string `json:"OrderKeyType,omitempty" xml:"OrderKeyType,omitempty"`
	PageNum               *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductDomainType     *string `json:"ProductDomainType,omitempty" xml:"ProductDomainType,omitempty"`
	QueryType             *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	StartExpirationDate   *int64  `json:"StartExpirationDate,omitempty" xml:"StartExpirationDate,omitempty"`
	StartRegistrationDate *int64  `json:"StartRegistrationDate,omitempty" xml:"StartRegistrationDate,omitempty"`
	UserClientIp          *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainListRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainListRequest) SetDomainGroupId(v string) *QueryDomainListRequest {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainListRequest) SetDomainName(v string) *QueryDomainListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDomainListRequest) SetEndExpirationDate(v int64) *QueryDomainListRequest {
	s.EndExpirationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetEndRegistrationDate(v int64) *QueryDomainListRequest {
	s.EndRegistrationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetLang(v string) *QueryDomainListRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainListRequest) SetOrderByType(v string) *QueryDomainListRequest {
	s.OrderByType = &v
	return s
}

func (s *QueryDomainListRequest) SetOrderKeyType(v string) *QueryDomainListRequest {
	s.OrderKeyType = &v
	return s
}

func (s *QueryDomainListRequest) SetPageNum(v int32) *QueryDomainListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryDomainListRequest) SetPageSize(v int32) *QueryDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDomainListRequest) SetProductDomainType(v string) *QueryDomainListRequest {
	s.ProductDomainType = &v
	return s
}

func (s *QueryDomainListRequest) SetQueryType(v string) *QueryDomainListRequest {
	s.QueryType = &v
	return s
}

func (s *QueryDomainListRequest) SetStartExpirationDate(v int64) *QueryDomainListRequest {
	s.StartExpirationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetStartRegistrationDate(v int64) *QueryDomainListRequest {
	s.StartRegistrationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetUserClientIp(v string) *QueryDomainListRequest {
	s.UserClientIp = &v
	return s
}

type QueryDomainListResponseBody struct {
	CurrentPageNum *int32                           `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                            `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                            `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                           `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                           `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBody) SetCurrentPageNum(v int32) *QueryDomainListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryDomainListResponseBody) SetData(v *QueryDomainListResponseBodyData) *QueryDomainListResponseBody {
	s.Data = v
	return s
}

func (s *QueryDomainListResponseBody) SetNextPage(v bool) *QueryDomainListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryDomainListResponseBody) SetPageSize(v int32) *QueryDomainListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDomainListResponseBody) SetPrePage(v bool) *QueryDomainListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryDomainListResponseBody) SetRequestId(v string) *QueryDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainListResponseBody) SetTotalItemNum(v int32) *QueryDomainListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryDomainListResponseBody) SetTotalPageNum(v int32) *QueryDomainListResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryDomainListResponseBodyData struct {
	Domain []*QueryDomainListResponseBodyDataDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s QueryDomainListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBodyData) SetDomain(v []*QueryDomainListResponseBodyDataDomain) *QueryDomainListResponseBodyData {
	s.Domain = v
	return s
}

type QueryDomainListResponseBodyDataDomain struct {
	DomainAuditStatus      *string `json:"DomainAuditStatus,omitempty" xml:"DomainAuditStatus,omitempty"`
	DomainGroupId          *string `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName        *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	DomainName             *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus           *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType             *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	ExpirationCurrDateDiff *int32  `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	ExpirationDate         *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong     *int64  `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	ExpirationDateStatus   *string `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Premium                *bool   `json:"Premium,omitempty" xml:"Premium,omitempty"`
	ProductId              *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RegistrantType         *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	RegistrationDate       *string `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	RegistrationDateLong   *int64  `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	Remark                 *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s QueryDomainListResponseBodyDataDomain) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainListResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainAuditStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainAuditStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainGroupId(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainGroupName(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainGroupName = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainName(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetDomainType(v string) *QueryDomainListResponseBodyDataDomain {
	s.DomainType = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetExpirationCurrDateDiff(v int32) *QueryDomainListResponseBodyDataDomain {
	s.ExpirationCurrDateDiff = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetExpirationDate(v string) *QueryDomainListResponseBodyDataDomain {
	s.ExpirationDate = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetExpirationDateLong(v int64) *QueryDomainListResponseBodyDataDomain {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetExpirationDateStatus(v string) *QueryDomainListResponseBodyDataDomain {
	s.ExpirationDateStatus = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetInstanceId(v string) *QueryDomainListResponseBodyDataDomain {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetPremium(v bool) *QueryDomainListResponseBodyDataDomain {
	s.Premium = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetProductId(v string) *QueryDomainListResponseBodyDataDomain {
	s.ProductId = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRegistrantType(v string) *QueryDomainListResponseBodyDataDomain {
	s.RegistrantType = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRegistrationDate(v string) *QueryDomainListResponseBodyDataDomain {
	s.RegistrationDate = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRegistrationDateLong(v int64) *QueryDomainListResponseBodyDataDomain {
	s.RegistrationDateLong = &v
	return s
}

func (s *QueryDomainListResponseBodyDataDomain) SetRemark(v string) *QueryDomainListResponseBodyDataDomain {
	s.Remark = &v
	return s
}

type QueryDomainListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainListResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainListResponse) SetHeaders(v map[string]*string) *QueryDomainListResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainListResponse) SetStatusCode(v int32) *QueryDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainListResponse) SetBody(v *QueryDomainListResponseBody) *QueryDomainListResponse {
	s.Body = v
	return s
}

type QueryDomainRealNameVerificationInfoRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FetchImage   *bool   `json:"FetchImage,omitempty" xml:"FetchImage,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainRealNameVerificationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainRealNameVerificationInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainRealNameVerificationInfoRequest) SetDomainName(v string) *QueryDomainRealNameVerificationInfoRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoRequest) SetFetchImage(v bool) *QueryDomainRealNameVerificationInfoRequest {
	s.FetchImage = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoRequest) SetLang(v string) *QueryDomainRealNameVerificationInfoRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoRequest) SetUserClientIp(v string) *QueryDomainRealNameVerificationInfoRequest {
	s.UserClientIp = &v
	return s
}

type QueryDomainRealNameVerificationInfoResponseBody struct {
	DomainName             *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	IdentityCredential     *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	IdentityCredentialNo   *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	IdentityCredentialUrl  *string `json:"IdentityCredentialUrl,omitempty" xml:"IdentityCredentialUrl,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubmissionDate         *string `json:"SubmissionDate,omitempty" xml:"SubmissionDate,omitempty"`
}

func (s QueryDomainRealNameVerificationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainRealNameVerificationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetDomainName(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetIdentityCredential(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.IdentityCredential = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetIdentityCredentialNo(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.IdentityCredentialNo = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetIdentityCredentialType(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.IdentityCredentialType = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetIdentityCredentialUrl(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.IdentityCredentialUrl = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetInstanceId(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetRequestId(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponseBody) SetSubmissionDate(v string) *QueryDomainRealNameVerificationInfoResponseBody {
	s.SubmissionDate = &v
	return s
}

type QueryDomainRealNameVerificationInfoResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDomainRealNameVerificationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDomainRealNameVerificationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainRealNameVerificationInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainRealNameVerificationInfoResponse) SetHeaders(v map[string]*string) *QueryDomainRealNameVerificationInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponse) SetStatusCode(v int32) *QueryDomainRealNameVerificationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoResponse) SetBody(v *QueryDomainRealNameVerificationInfoResponseBody) *QueryDomainRealNameVerificationInfoResponse {
	s.Body = v
	return s
}

type QueryDomainSuffixRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainSuffixRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainSuffixRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainSuffixRequest) SetLang(v string) *QueryDomainSuffixRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainSuffixRequest) SetUserClientIp(v string) *QueryDomainSuffixRequest {
	s.UserClientIp = &v
	return s
}

type QueryDomainSuffixResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuffixList *QueryDomainSuffixResponseBodySuffixList `json:"SuffixList,omitempty" xml:"SuffixList,omitempty" type:"Struct"`
}

func (s QueryDomainSuffixResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainSuffixResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainSuffixResponseBody) SetRequestId(v string) *QueryDomainSuffixResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainSuffixResponseBody) SetSuffixList(v *QueryDomainSuffixResponseBodySuffixList) *QueryDomainSuffixResponseBody {
	s.SuffixList = v
	return s
}

type QueryDomainSuffixResponseBodySuffixList struct {
	Suffix []*string `json:"Suffix,omitempty" xml:"Suffix,omitempty" type:"Repeated"`
}

func (s QueryDomainSuffixResponseBodySuffixList) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainSuffixResponseBodySuffixList) GoString() string {
	return s.String()
}

func (s *QueryDomainSuffixResponseBodySuffixList) SetSuffix(v []*string) *QueryDomainSuffixResponseBodySuffixList {
	s.Suffix = v
	return s
}

type QueryDomainSuffixResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDomainSuffixResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDomainSuffixResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainSuffixResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainSuffixResponse) SetHeaders(v map[string]*string) *QueryDomainSuffixResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainSuffixResponse) SetStatusCode(v int32) *QueryDomainSuffixResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainSuffixResponse) SetBody(v *QueryDomainSuffixResponseBody) *QueryDomainSuffixResponse {
	s.Body = v
	return s
}

type QueryEmailVerificationRequest struct {
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryEmailVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *QueryEmailVerificationRequest) SetEmail(v string) *QueryEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *QueryEmailVerificationRequest) SetLang(v string) *QueryEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *QueryEmailVerificationRequest) SetUserClientIp(v string) *QueryEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

type QueryEmailVerificationResponseBody struct {
	ConfirmIp           *string `json:"ConfirmIp,omitempty" xml:"ConfirmIp,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EmailVerificationNo *string `json:"EmailVerificationNo,omitempty" xml:"EmailVerificationNo,omitempty"`
	GmtCreate           *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SendIp              *string `json:"SendIp,omitempty" xml:"SendIp,omitempty"`
	TokenSendTime       *string `json:"TokenSendTime,omitempty" xml:"TokenSendTime,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VerificationStatus  *int32  `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
	VerificationTime    *string `json:"VerificationTime,omitempty" xml:"VerificationTime,omitempty"`
}

func (s QueryEmailVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmailVerificationResponseBody) SetConfirmIp(v string) *QueryEmailVerificationResponseBody {
	s.ConfirmIp = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetEmail(v string) *QueryEmailVerificationResponseBody {
	s.Email = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetEmailVerificationNo(v string) *QueryEmailVerificationResponseBody {
	s.EmailVerificationNo = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetGmtCreate(v string) *QueryEmailVerificationResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetGmtModified(v string) *QueryEmailVerificationResponseBody {
	s.GmtModified = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetRequestId(v string) *QueryEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetSendIp(v string) *QueryEmailVerificationResponseBody {
	s.SendIp = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetTokenSendTime(v string) *QueryEmailVerificationResponseBody {
	s.TokenSendTime = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetUserId(v string) *QueryEmailVerificationResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetVerificationStatus(v int32) *QueryEmailVerificationResponseBody {
	s.VerificationStatus = &v
	return s
}

func (s *QueryEmailVerificationResponseBody) SetVerificationTime(v string) *QueryEmailVerificationResponseBody {
	s.VerificationTime = &v
	return s
}

type QueryEmailVerificationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEmailVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *QueryEmailVerificationResponse) SetHeaders(v map[string]*string) *QueryEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *QueryEmailVerificationResponse) SetStatusCode(v int32) *QueryEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEmailVerificationResponse) SetBody(v *QueryEmailVerificationResponseBody) *QueryEmailVerificationResponse {
	s.Body = v
	return s
}

type QueryEnsAssociationRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryEnsAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEnsAssociationRequest) GoString() string {
	return s.String()
}

func (s *QueryEnsAssociationRequest) SetDomainName(v string) *QueryEnsAssociationRequest {
	s.DomainName = &v
	return s
}

func (s *QueryEnsAssociationRequest) SetLang(v string) *QueryEnsAssociationRequest {
	s.Lang = &v
	return s
}

func (s *QueryEnsAssociationRequest) SetUserClientIp(v string) *QueryEnsAssociationRequest {
	s.UserClientIp = &v
	return s
}

type QueryEnsAssociationResponseBody struct {
	Address   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryEnsAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEnsAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEnsAssociationResponseBody) SetAddress(v string) *QueryEnsAssociationResponseBody {
	s.Address = &v
	return s
}

func (s *QueryEnsAssociationResponseBody) SetRequestId(v string) *QueryEnsAssociationResponseBody {
	s.RequestId = &v
	return s
}

type QueryEnsAssociationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEnsAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEnsAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEnsAssociationResponse) GoString() string {
	return s.String()
}

func (s *QueryEnsAssociationResponse) SetHeaders(v map[string]*string) *QueryEnsAssociationResponse {
	s.Headers = v
	return s
}

func (s *QueryEnsAssociationResponse) SetStatusCode(v int32) *QueryEnsAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEnsAssociationResponse) SetBody(v *QueryEnsAssociationResponseBody) *QueryEnsAssociationResponse {
	s.Body = v
	return s
}

type QueryFailReasonForDomainRealNameVerificationRequest struct {
	DomainName                 *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang                       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RealNameVerificationAction *string `json:"RealNameVerificationAction,omitempty" xml:"RealNameVerificationAction,omitempty"`
	UserClientIp               *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryFailReasonForDomainRealNameVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFailReasonForDomainRealNameVerificationRequest) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) SetDomainName(v string) *QueryFailReasonForDomainRealNameVerificationRequest {
	s.DomainName = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) SetLang(v string) *QueryFailReasonForDomainRealNameVerificationRequest {
	s.Lang = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) SetRealNameVerificationAction(v string) *QueryFailReasonForDomainRealNameVerificationRequest {
	s.RealNameVerificationAction = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationRequest) SetUserClientIp(v string) *QueryFailReasonForDomainRealNameVerificationRequest {
	s.UserClientIp = &v
	return s
}

type QueryFailReasonForDomainRealNameVerificationResponseBody struct {
	Data      []*QueryFailReasonForDomainRealNameVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFailReasonForDomainRealNameVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFailReasonForDomainRealNameVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBody) SetData(v []*QueryFailReasonForDomainRealNameVerificationResponseBodyData) *QueryFailReasonForDomainRealNameVerificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBody) SetRequestId(v string) *QueryFailReasonForDomainRealNameVerificationResponseBody {
	s.RequestId = &v
	return s
}

type QueryFailReasonForDomainRealNameVerificationResponseBodyData struct {
	Date                         *string `json:"Date,omitempty" xml:"Date,omitempty"`
	DomainNameVerificationStatus *string `json:"DomainNameVerificationStatus,omitempty" xml:"DomainNameVerificationStatus,omitempty"`
	FailReason                   *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s QueryFailReasonForDomainRealNameVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFailReasonForDomainRealNameVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBodyData) SetDate(v string) *QueryFailReasonForDomainRealNameVerificationResponseBodyData {
	s.Date = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBodyData) SetDomainNameVerificationStatus(v string) *QueryFailReasonForDomainRealNameVerificationResponseBodyData {
	s.DomainNameVerificationStatus = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponseBodyData) SetFailReason(v string) *QueryFailReasonForDomainRealNameVerificationResponseBodyData {
	s.FailReason = &v
	return s
}

type QueryFailReasonForDomainRealNameVerificationResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryFailReasonForDomainRealNameVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFailReasonForDomainRealNameVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFailReasonForDomainRealNameVerificationResponse) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) SetHeaders(v map[string]*string) *QueryFailReasonForDomainRealNameVerificationResponse {
	s.Headers = v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) SetStatusCode(v int32) *QueryFailReasonForDomainRealNameVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFailReasonForDomainRealNameVerificationResponse) SetBody(v *QueryFailReasonForDomainRealNameVerificationResponseBody) *QueryFailReasonForDomainRealNameVerificationResponse {
	s.Body = v
	return s
}

type QueryFailReasonForRegistrantProfileRealNameVerificationRequest struct {
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegistrantProfileID *int64  `json:"RegistrantProfileID,omitempty" xml:"RegistrantProfileID,omitempty"`
	UserClientIp        *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationRequest) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) SetLang(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationRequest {
	s.Lang = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) SetRegistrantProfileID(v int64) *QueryFailReasonForRegistrantProfileRealNameVerificationRequest {
	s.RegistrantProfileID = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) SetUserClientIp(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationRequest {
	s.UserClientIp = &v
	return s
}

type QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody struct {
	Data      []*QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) SetData(v []*QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) SetRequestId(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody {
	s.RequestId = &v
	return s
}

type QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData struct {
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty"`
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) SetDate(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData {
	s.Date = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData) SetFailReason(v string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBodyData {
	s.FailReason = &v
	return s
}

type QueryFailReasonForRegistrantProfileRealNameVerificationResponse struct {
	Headers    map[string]*string                                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFailReasonForRegistrantProfileRealNameVerificationResponse) GoString() string {
	return s.String()
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) SetHeaders(v map[string]*string) *QueryFailReasonForRegistrantProfileRealNameVerificationResponse {
	s.Headers = v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) SetStatusCode(v int32) *QueryFailReasonForRegistrantProfileRealNameVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFailReasonForRegistrantProfileRealNameVerificationResponse) SetBody(v *QueryFailReasonForRegistrantProfileRealNameVerificationResponseBody) *QueryFailReasonForRegistrantProfileRealNameVerificationResponse {
	s.Body = v
	return s
}

type QueryFailingReasonListForQualificationRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Limit             *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	QualificationType *string `json:"QualificationType,omitempty" xml:"QualificationType,omitempty"`
	UserClientIp      *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryFailingReasonListForQualificationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFailingReasonListForQualificationRequest) GoString() string {
	return s.String()
}

func (s *QueryFailingReasonListForQualificationRequest) SetInstanceId(v string) *QueryFailingReasonListForQualificationRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryFailingReasonListForQualificationRequest) SetLang(v string) *QueryFailingReasonListForQualificationRequest {
	s.Lang = &v
	return s
}

func (s *QueryFailingReasonListForQualificationRequest) SetLimit(v int32) *QueryFailingReasonListForQualificationRequest {
	s.Limit = &v
	return s
}

func (s *QueryFailingReasonListForQualificationRequest) SetQualificationType(v string) *QueryFailingReasonListForQualificationRequest {
	s.QualificationType = &v
	return s
}

func (s *QueryFailingReasonListForQualificationRequest) SetUserClientIp(v string) *QueryFailingReasonListForQualificationRequest {
	s.UserClientIp = &v
	return s
}

type QueryFailingReasonListForQualificationResponseBody struct {
	Data      []*QueryFailingReasonListForQualificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFailingReasonListForQualificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFailingReasonListForQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFailingReasonListForQualificationResponseBody) SetData(v []*QueryFailingReasonListForQualificationResponseBodyData) *QueryFailingReasonListForQualificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryFailingReasonListForQualificationResponseBody) SetRequestId(v string) *QueryFailingReasonListForQualificationResponseBody {
	s.RequestId = &v
	return s
}

type QueryFailingReasonListForQualificationResponseBodyData struct {
	Date       *string `json:"Date,omitempty" xml:"Date,omitempty"`
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s QueryFailingReasonListForQualificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFailingReasonListForQualificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFailingReasonListForQualificationResponseBodyData) SetDate(v string) *QueryFailingReasonListForQualificationResponseBodyData {
	s.Date = &v
	return s
}

func (s *QueryFailingReasonListForQualificationResponseBodyData) SetFailReason(v string) *QueryFailingReasonListForQualificationResponseBodyData {
	s.FailReason = &v
	return s
}

type QueryFailingReasonListForQualificationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryFailingReasonListForQualificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFailingReasonListForQualificationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFailingReasonListForQualificationResponse) GoString() string {
	return s.String()
}

func (s *QueryFailingReasonListForQualificationResponse) SetHeaders(v map[string]*string) *QueryFailingReasonListForQualificationResponse {
	s.Headers = v
	return s
}

func (s *QueryFailingReasonListForQualificationResponse) SetStatusCode(v int32) *QueryFailingReasonListForQualificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFailingReasonListForQualificationResponse) SetBody(v *QueryFailingReasonListForQualificationResponseBody) *QueryFailingReasonListForQualificationResponse {
	s.Body = v
	return s
}

type QueryLocalEnsAssociationRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryLocalEnsAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLocalEnsAssociationRequest) GoString() string {
	return s.String()
}

func (s *QueryLocalEnsAssociationRequest) SetDomainName(v string) *QueryLocalEnsAssociationRequest {
	s.DomainName = &v
	return s
}

func (s *QueryLocalEnsAssociationRequest) SetLang(v string) *QueryLocalEnsAssociationRequest {
	s.Lang = &v
	return s
}

func (s *QueryLocalEnsAssociationRequest) SetUserClientIp(v string) *QueryLocalEnsAssociationRequest {
	s.UserClientIp = &v
	return s
}

type QueryLocalEnsAssociationResponseBody struct {
	Address   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryLocalEnsAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLocalEnsAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLocalEnsAssociationResponseBody) SetAddress(v string) *QueryLocalEnsAssociationResponseBody {
	s.Address = &v
	return s
}

func (s *QueryLocalEnsAssociationResponseBody) SetRequestId(v string) *QueryLocalEnsAssociationResponseBody {
	s.RequestId = &v
	return s
}

type QueryLocalEnsAssociationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryLocalEnsAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryLocalEnsAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLocalEnsAssociationResponse) GoString() string {
	return s.String()
}

func (s *QueryLocalEnsAssociationResponse) SetHeaders(v map[string]*string) *QueryLocalEnsAssociationResponse {
	s.Headers = v
	return s
}

func (s *QueryLocalEnsAssociationResponse) SetStatusCode(v int32) *QueryLocalEnsAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLocalEnsAssociationResponse) SetBody(v *QueryLocalEnsAssociationResponseBody) *QueryLocalEnsAssociationResponse {
	s.Body = v
	return s
}

type QueryOperationAuditInfoDetailRequest struct {
	AuditRecordId *int64  `json:"AuditRecordId,omitempty" xml:"AuditRecordId,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s QueryOperationAuditInfoDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOperationAuditInfoDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoDetailRequest) SetAuditRecordId(v int64) *QueryOperationAuditInfoDetailRequest {
	s.AuditRecordId = &v
	return s
}

func (s *QueryOperationAuditInfoDetailRequest) SetLang(v string) *QueryOperationAuditInfoDetailRequest {
	s.Lang = &v
	return s
}

type QueryOperationAuditInfoDetailResponseBody struct {
	AuditInfo    *string `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty"`
	AuditStatus  *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuditType    *int32  `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UpdateTime   *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryOperationAuditInfoDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOperationAuditInfoDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetAuditInfo(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.AuditInfo = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetAuditStatus(v int32) *QueryOperationAuditInfoDetailResponseBody {
	s.AuditStatus = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetAuditType(v int32) *QueryOperationAuditInfoDetailResponseBody {
	s.AuditType = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetBusinessName(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.BusinessName = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetCreateTime(v int64) *QueryOperationAuditInfoDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetDomainName(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetId(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.Id = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetRemark(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetRequestId(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetUpdateTime(v int64) *QueryOperationAuditInfoDetailResponseBody {
	s.UpdateTime = &v
	return s
}

type QueryOperationAuditInfoDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOperationAuditInfoDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOperationAuditInfoDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOperationAuditInfoDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoDetailResponse) SetHeaders(v map[string]*string) *QueryOperationAuditInfoDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryOperationAuditInfoDetailResponse) SetStatusCode(v int32) *QueryOperationAuditInfoDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponse) SetBody(v *QueryOperationAuditInfoDetailResponseBody) *QueryOperationAuditInfoDetailResponse {
	s.Body = v
	return s
}

type QueryOperationAuditInfoListRequest struct {
	AuditStatus *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuditType   *int32  `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryOperationAuditInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOperationAuditInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoListRequest) SetAuditStatus(v int32) *QueryOperationAuditInfoListRequest {
	s.AuditStatus = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetAuditType(v int32) *QueryOperationAuditInfoListRequest {
	s.AuditType = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetDomainName(v string) *QueryOperationAuditInfoListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetLang(v string) *QueryOperationAuditInfoListRequest {
	s.Lang = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetPageNum(v int32) *QueryOperationAuditInfoListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryOperationAuditInfoListRequest) SetPageSize(v int32) *QueryOperationAuditInfoListRequest {
	s.PageSize = &v
	return s
}

type QueryOperationAuditInfoListResponseBody struct {
	CurrentPageNum *int32                                         `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryOperationAuditInfoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextPage       *bool                                          `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                                          `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                                         `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                         `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryOperationAuditInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOperationAuditInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoListResponseBody) SetCurrentPageNum(v int32) *QueryOperationAuditInfoListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetData(v []*QueryOperationAuditInfoListResponseBodyData) *QueryOperationAuditInfoListResponseBody {
	s.Data = v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetNextPage(v bool) *QueryOperationAuditInfoListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetPageSize(v int32) *QueryOperationAuditInfoListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetPrePage(v bool) *QueryOperationAuditInfoListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetRequestId(v string) *QueryOperationAuditInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetTotalItemNum(v int32) *QueryOperationAuditInfoListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBody) SetTotalPageNum(v int32) *QueryOperationAuditInfoListResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryOperationAuditInfoListResponseBodyData struct {
	AuditInfo    *string `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty"`
	AuditStatus  *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuditType    *int32  `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UpdateTime   *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryOperationAuditInfoListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOperationAuditInfoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetAuditInfo(v string) *QueryOperationAuditInfoListResponseBodyData {
	s.AuditInfo = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetAuditStatus(v int32) *QueryOperationAuditInfoListResponseBodyData {
	s.AuditStatus = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetAuditType(v int32) *QueryOperationAuditInfoListResponseBodyData {
	s.AuditType = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetBusinessName(v string) *QueryOperationAuditInfoListResponseBodyData {
	s.BusinessName = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetCreateTime(v int64) *QueryOperationAuditInfoListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetDomainName(v string) *QueryOperationAuditInfoListResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetId(v int64) *QueryOperationAuditInfoListResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetRemark(v string) *QueryOperationAuditInfoListResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QueryOperationAuditInfoListResponseBodyData) SetUpdateTime(v int64) *QueryOperationAuditInfoListResponseBodyData {
	s.UpdateTime = &v
	return s
}

type QueryOperationAuditInfoListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOperationAuditInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOperationAuditInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOperationAuditInfoListResponse) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoListResponse) SetHeaders(v map[string]*string) *QueryOperationAuditInfoListResponse {
	s.Headers = v
	return s
}

func (s *QueryOperationAuditInfoListResponse) SetStatusCode(v int32) *QueryOperationAuditInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOperationAuditInfoListResponse) SetBody(v *QueryOperationAuditInfoListResponseBody) *QueryOperationAuditInfoListResponse {
	s.Body = v
	return s
}

type QueryQualificationDetailRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	QualificationType *string `json:"QualificationType,omitempty" xml:"QualificationType,omitempty"`
	UserClientIp      *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryQualificationDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryQualificationDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailRequest) SetInstanceId(v string) *QueryQualificationDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryQualificationDetailRequest) SetLang(v string) *QueryQualificationDetailRequest {
	s.Lang = &v
	return s
}

func (s *QueryQualificationDetailRequest) SetQualificationType(v string) *QueryQualificationDetailRequest {
	s.QualificationType = &v
	return s
}

func (s *QueryQualificationDetailRequest) SetUserClientIp(v string) *QueryQualificationDetailRequest {
	s.UserClientIp = &v
	return s
}

type QueryQualificationDetailResponseBody struct {
	AuditStatus *int32                                           `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	Credentials *QueryQualificationDetailResponseBodyCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrackId     *string                                          `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s QueryQualificationDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryQualificationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailResponseBody) SetAuditStatus(v int32) *QueryQualificationDetailResponseBody {
	s.AuditStatus = &v
	return s
}

func (s *QueryQualificationDetailResponseBody) SetCredentials(v *QueryQualificationDetailResponseBodyCredentials) *QueryQualificationDetailResponseBody {
	s.Credentials = v
	return s
}

func (s *QueryQualificationDetailResponseBody) SetRequestId(v string) *QueryQualificationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryQualificationDetailResponseBody) SetTrackId(v string) *QueryQualificationDetailResponseBody {
	s.TrackId = &v
	return s
}

type QueryQualificationDetailResponseBodyCredentials struct {
	QualificationCredential []*QueryQualificationDetailResponseBodyCredentialsQualificationCredential `json:"QualificationCredential,omitempty" xml:"QualificationCredential,omitempty" type:"Repeated"`
}

func (s QueryQualificationDetailResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s QueryQualificationDetailResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailResponseBodyCredentials) SetQualificationCredential(v []*QueryQualificationDetailResponseBodyCredentialsQualificationCredential) *QueryQualificationDetailResponseBodyCredentials {
	s.QualificationCredential = v
	return s
}

type QueryQualificationDetailResponseBodyCredentialsQualificationCredential struct {
	CredentialNo   *string `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	CredentialUrl  *string `json:"CredentialUrl,omitempty" xml:"CredentialUrl,omitempty"`
}

func (s QueryQualificationDetailResponseBodyCredentialsQualificationCredential) String() string {
	return tea.Prettify(s)
}

func (s QueryQualificationDetailResponseBodyCredentialsQualificationCredential) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) SetCredentialNo(v string) *QueryQualificationDetailResponseBodyCredentialsQualificationCredential {
	s.CredentialNo = &v
	return s
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) SetCredentialType(v string) *QueryQualificationDetailResponseBodyCredentialsQualificationCredential {
	s.CredentialType = &v
	return s
}

func (s *QueryQualificationDetailResponseBodyCredentialsQualificationCredential) SetCredentialUrl(v string) *QueryQualificationDetailResponseBodyCredentialsQualificationCredential {
	s.CredentialUrl = &v
	return s
}

type QueryQualificationDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryQualificationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryQualificationDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryQualificationDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryQualificationDetailResponse) SetHeaders(v map[string]*string) *QueryQualificationDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryQualificationDetailResponse) SetStatusCode(v int32) *QueryQualificationDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryQualificationDetailResponse) SetBody(v *QueryQualificationDetailResponseBody) *QueryQualificationDetailResponse {
	s.Body = v
	return s
}

type QueryRegistrantProfileRealNameVerificationInfoRequest struct {
	FetchImage          *bool   `json:"FetchImage,omitempty" xml:"FetchImage,omitempty"`
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegistrantProfileId *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	UserClientIp        *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryRegistrantProfileRealNameVerificationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRegistrantProfileRealNameVerificationInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetFetchImage(v bool) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.FetchImage = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetLang(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.Lang = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetRegistrantProfileId(v int64) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetUserClientIp(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.UserClientIp = &v
	return s
}

type QueryRegistrantProfileRealNameVerificationInfoResponseBody struct {
	IdentityCredential     *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	IdentityCredentialNo   *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	IdentityCredentialUrl  *string `json:"IdentityCredentialUrl,omitempty" xml:"IdentityCredentialUrl,omitempty"`
	ModificationDate       *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	RegistrantProfileId    *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubmissionDate         *string `json:"SubmissionDate,omitempty" xml:"SubmissionDate,omitempty"`
}

func (s QueryRegistrantProfileRealNameVerificationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRegistrantProfileRealNameVerificationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetIdentityCredential(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.IdentityCredential = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetIdentityCredentialNo(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.IdentityCredentialNo = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetIdentityCredentialType(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.IdentityCredentialType = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetIdentityCredentialUrl(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.IdentityCredentialUrl = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetModificationDate(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.ModificationDate = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetRegistrantProfileId(v int64) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetRequestId(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponseBody) SetSubmissionDate(v string) *QueryRegistrantProfileRealNameVerificationInfoResponseBody {
	s.SubmissionDate = &v
	return s
}

type QueryRegistrantProfileRealNameVerificationInfoResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRegistrantProfileRealNameVerificationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRegistrantProfileRealNameVerificationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRegistrantProfileRealNameVerificationInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) SetHeaders(v map[string]*string) *QueryRegistrantProfileRealNameVerificationInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) SetStatusCode(v int32) *QueryRegistrantProfileRealNameVerificationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoResponse) SetBody(v *QueryRegistrantProfileRealNameVerificationInfoResponseBody) *QueryRegistrantProfileRealNameVerificationInfoResponse {
	s.Body = v
	return s
}

type QueryRegistrantProfilesRequest struct {
	DefaultRegistrantProfile *bool   `json:"DefaultRegistrantProfile,omitempty" xml:"DefaultRegistrantProfile,omitempty"`
	Email                    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang                     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum                  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize                 *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RealNameStatus           *string `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	RegistrantOrganization   *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantProfileId      *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RegistrantProfileType    *string `json:"RegistrantProfileType,omitempty" xml:"RegistrantProfileType,omitempty"`
	RegistrantType           *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	UserClientIp             *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryRegistrantProfilesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRegistrantProfilesRequest) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesRequest) SetDefaultRegistrantProfile(v bool) *QueryRegistrantProfilesRequest {
	s.DefaultRegistrantProfile = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetEmail(v string) *QueryRegistrantProfilesRequest {
	s.Email = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetLang(v string) *QueryRegistrantProfilesRequest {
	s.Lang = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetPageNum(v int32) *QueryRegistrantProfilesRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetPageSize(v int32) *QueryRegistrantProfilesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRealNameStatus(v string) *QueryRegistrantProfilesRequest {
	s.RealNameStatus = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRegistrantOrganization(v string) *QueryRegistrantProfilesRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRegistrantProfileId(v int64) *QueryRegistrantProfilesRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRegistrantProfileType(v string) *QueryRegistrantProfilesRequest {
	s.RegistrantProfileType = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetRegistrantType(v string) *QueryRegistrantProfilesRequest {
	s.RegistrantType = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetUserClientIp(v string) *QueryRegistrantProfilesRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryRegistrantProfilesRequest) SetZhRegistrantOrganization(v string) *QueryRegistrantProfilesRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

type QueryRegistrantProfilesResponseBody struct {
	CurrentPageNum     *int32                                                 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	NextPage           *bool                                                  `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize           *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage            *bool                                                  `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RegistrantProfiles *QueryRegistrantProfilesResponseBodyRegistrantProfiles `json:"RegistrantProfiles,omitempty" xml:"RegistrantProfiles,omitempty" type:"Struct"`
	RequestId          *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum       *int32                                                 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum       *int32                                                 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryRegistrantProfilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRegistrantProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponseBody) SetCurrentPageNum(v int32) *QueryRegistrantProfilesResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetNextPage(v bool) *QueryRegistrantProfilesResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetPageSize(v int32) *QueryRegistrantProfilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetPrePage(v bool) *QueryRegistrantProfilesResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetRegistrantProfiles(v *QueryRegistrantProfilesResponseBodyRegistrantProfiles) *QueryRegistrantProfilesResponseBody {
	s.RegistrantProfiles = v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetRequestId(v string) *QueryRegistrantProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetTotalItemNum(v int32) *QueryRegistrantProfilesResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBody) SetTotalPageNum(v int32) *QueryRegistrantProfilesResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryRegistrantProfilesResponseBodyRegistrantProfiles struct {
	RegistrantProfile []*QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile `json:"RegistrantProfile,omitempty" xml:"RegistrantProfile,omitempty" type:"Repeated"`
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfiles) String() string {
	return tea.Prettify(s)
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfiles) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfiles) SetRegistrantProfile(v []*QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) *QueryRegistrantProfilesResponseBodyRegistrantProfiles {
	s.RegistrantProfile = v
	return s
}

type QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile struct {
	Address                  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	City                     *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DefaultRegistrantProfile *bool   `json:"DefaultRegistrantProfile,omitempty" xml:"DefaultRegistrantProfile,omitempty"`
	Email                    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EmailVerificationStatus  *int32  `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	PostalCode               *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                 *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RealNameStatus           *string `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
	RegistrantName           *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization   *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantProfileId      *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RegistrantProfileType    *string `json:"RegistrantProfileType,omitempty" xml:"RegistrantProfileType,omitempty"`
	RegistrantType           *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	TelArea                  *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                   *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	UpdateTime               *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ZhAddress                *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) String() string {
	return tea.Prettify(s)
}

func (s QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetAddress(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Address = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCity(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.City = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCountry(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Country = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetCreateTime(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.CreateTime = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetDefaultRegistrantProfile(v bool) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.DefaultRegistrantProfile = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetEmail(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Email = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetEmailVerificationStatus(v int32) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.EmailVerificationStatus = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetPostalCode(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.PostalCode = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetProvince(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Province = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRealNameStatus(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RealNameStatus = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantName(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantName = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantOrganization(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantProfileId(v int64) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantProfileType(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantProfileType = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetRegistrantType(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.RegistrantType = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetTelArea(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.TelArea = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetTelExt(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.TelExt = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetTelephone(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.Telephone = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetUpdateTime(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.UpdateTime = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhAddress(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhAddress = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhCity(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhCity = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhProvince(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhProvince = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhRegistrantName(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhRegistrantName = &v
	return s
}

func (s *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile) SetZhRegistrantOrganization(v string) *QueryRegistrantProfilesResponseBodyRegistrantProfilesRegistrantProfile {
	s.ZhRegistrantOrganization = &v
	return s
}

type QueryRegistrantProfilesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRegistrantProfilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRegistrantProfilesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRegistrantProfilesResponse) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfilesResponse) SetHeaders(v map[string]*string) *QueryRegistrantProfilesResponse {
	s.Headers = v
	return s
}

func (s *QueryRegistrantProfilesResponse) SetStatusCode(v int32) *QueryRegistrantProfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRegistrantProfilesResponse) SetBody(v *QueryRegistrantProfilesResponseBody) *QueryRegistrantProfilesResponse {
	s.Body = v
	return s
}

type QueryServerLockRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryServerLockRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServerLockRequest) GoString() string {
	return s.String()
}

func (s *QueryServerLockRequest) SetInstanceId(v string) *QueryServerLockRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryServerLockRequest) SetLang(v string) *QueryServerLockRequest {
	s.Lang = &v
	return s
}

func (s *QueryServerLockRequest) SetUserClientIp(v string) *QueryServerLockRequest {
	s.UserClientIp = &v
	return s
}

type QueryServerLockResponseBody struct {
	DomainInstanceId *string `json:"DomainInstanceId,omitempty" xml:"DomainInstanceId,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ExpireDate       *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LockInstanceId   *string `json:"LockInstanceId,omitempty" xml:"LockInstanceId,omitempty"`
	LockProductId    *string `json:"LockProductId,omitempty" xml:"LockProductId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerLockStatus *int32  `json:"ServerLockStatus,omitempty" xml:"ServerLockStatus,omitempty"`
	StartDate        *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryServerLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServerLockResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServerLockResponseBody) SetDomainInstanceId(v string) *QueryServerLockResponseBody {
	s.DomainInstanceId = &v
	return s
}

func (s *QueryServerLockResponseBody) SetDomainName(v string) *QueryServerLockResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryServerLockResponseBody) SetExpireDate(v string) *QueryServerLockResponseBody {
	s.ExpireDate = &v
	return s
}

func (s *QueryServerLockResponseBody) SetGmtCreate(v string) *QueryServerLockResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *QueryServerLockResponseBody) SetGmtModified(v string) *QueryServerLockResponseBody {
	s.GmtModified = &v
	return s
}

func (s *QueryServerLockResponseBody) SetLockInstanceId(v string) *QueryServerLockResponseBody {
	s.LockInstanceId = &v
	return s
}

func (s *QueryServerLockResponseBody) SetLockProductId(v string) *QueryServerLockResponseBody {
	s.LockProductId = &v
	return s
}

func (s *QueryServerLockResponseBody) SetRequestId(v string) *QueryServerLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServerLockResponseBody) SetServerLockStatus(v int32) *QueryServerLockResponseBody {
	s.ServerLockStatus = &v
	return s
}

func (s *QueryServerLockResponseBody) SetStartDate(v string) *QueryServerLockResponseBody {
	s.StartDate = &v
	return s
}

func (s *QueryServerLockResponseBody) SetUserId(v string) *QueryServerLockResponseBody {
	s.UserId = &v
	return s
}

type QueryServerLockResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryServerLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServerLockResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServerLockResponse) GoString() string {
	return s.String()
}

func (s *QueryServerLockResponse) SetHeaders(v map[string]*string) *QueryServerLockResponse {
	s.Headers = v
	return s
}

func (s *QueryServerLockResponse) SetStatusCode(v int32) *QueryServerLockResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryServerLockResponse) SetBody(v *QueryServerLockResponseBody) *QueryServerLockResponse {
	s.Body = v
	return s
}

type QueryTaskDetailHistoryRequest struct {
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameCursor   *string `json:"DomainNameCursor,omitempty" xml:"DomainNameCursor,omitempty"`
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskDetailNoCursor *string `json:"TaskDetailNoCursor,omitempty" xml:"TaskDetailNoCursor,omitempty"`
	TaskNo             *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskStatus         *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	UserClientIp       *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTaskDetailHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailHistoryRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryRequest) SetDomainName(v string) *QueryTaskDetailHistoryRequest {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetDomainNameCursor(v string) *QueryTaskDetailHistoryRequest {
	s.DomainNameCursor = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetLang(v string) *QueryTaskDetailHistoryRequest {
	s.Lang = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetPageSize(v int32) *QueryTaskDetailHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetTaskDetailNoCursor(v string) *QueryTaskDetailHistoryRequest {
	s.TaskDetailNoCursor = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetTaskNo(v string) *QueryTaskDetailHistoryRequest {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetTaskStatus(v int32) *QueryTaskDetailHistoryRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryRequest) SetUserClientIp(v string) *QueryTaskDetailHistoryRequest {
	s.UserClientIp = &v
	return s
}

type QueryTaskDetailHistoryResponseBody struct {
	CurrentPageCursor *QueryTaskDetailHistoryResponseBodyCurrentPageCursor `json:"CurrentPageCursor,omitempty" xml:"CurrentPageCursor,omitempty" type:"Struct"`
	NextPageCursor    *QueryTaskDetailHistoryResponseBodyNextPageCursor    `json:"NextPageCursor,omitempty" xml:"NextPageCursor,omitempty" type:"Struct"`
	Objects           []*QueryTaskDetailHistoryResponseBodyObjects         `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	PageSize          *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePageCursor     *QueryTaskDetailHistoryResponseBodyPrePageCursor     `json:"PrePageCursor,omitempty" xml:"PrePageCursor,omitempty" type:"Struct"`
	RequestId         *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBody) SetCurrentPageCursor(v *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) *QueryTaskDetailHistoryResponseBody {
	s.CurrentPageCursor = v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetNextPageCursor(v *QueryTaskDetailHistoryResponseBodyNextPageCursor) *QueryTaskDetailHistoryResponseBody {
	s.NextPageCursor = v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetObjects(v []*QueryTaskDetailHistoryResponseBodyObjects) *QueryTaskDetailHistoryResponseBody {
	s.Objects = v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetPageSize(v int32) *QueryTaskDetailHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetPrePageCursor(v *QueryTaskDetailHistoryResponseBodyPrePageCursor) *QueryTaskDetailHistoryResponseBody {
	s.PrePageCursor = v
	return s
}

func (s *QueryTaskDetailHistoryResponseBody) SetRequestId(v string) *QueryTaskDetailHistoryResponseBody {
	s.RequestId = &v
	return s
}

type QueryTaskDetailHistoryResponseBodyCurrentPageCursor struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrorMsg            *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskDetailNo        *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	TryCount            *int32  `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBodyCurrentPageCursor) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBodyCurrentPageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetCreateTime(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetDomainName(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetErrorMsg(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetInstanceId(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskDetailNo(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskNo(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskStatus(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskStatusCode(v int32) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskType(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTaskTypeDescription(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetTryCount(v int32) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyCurrentPageCursor) SetUpdateTime(v string) *QueryTaskDetailHistoryResponseBodyCurrentPageCursor {
	s.UpdateTime = &v
	return s
}

type QueryTaskDetailHistoryResponseBodyNextPageCursor struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrorMsg            *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskDetailNo        *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	TryCount            *int32  `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBodyNextPageCursor) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBodyNextPageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetCreateTime(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetDomainName(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetErrorMsg(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetInstanceId(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskDetailNo(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskNo(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskStatus(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskStatusCode(v int32) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskType(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTaskTypeDescription(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetTryCount(v int32) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyNextPageCursor) SetUpdateTime(v string) *QueryTaskDetailHistoryResponseBodyNextPageCursor {
	s.UpdateTime = &v
	return s
}

type QueryTaskDetailHistoryResponseBodyObjects struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrorMsg            *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskDetailNo        *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	TryCount            *int32  `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBodyObjects) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBodyObjects) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetCreateTime(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetDomainName(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetErrorMsg(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetInstanceId(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskDetailNo(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskNo(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskStatus(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskStatusCode(v int32) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskType(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTaskTypeDescription(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetTryCount(v int32) *QueryTaskDetailHistoryResponseBodyObjects {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyObjects) SetUpdateTime(v string) *QueryTaskDetailHistoryResponseBodyObjects {
	s.UpdateTime = &v
	return s
}

type QueryTaskDetailHistoryResponseBodyPrePageCursor struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrorMsg            *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskDetailNo        *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	TryCount            *int32  `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailHistoryResponseBodyPrePageCursor) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailHistoryResponseBodyPrePageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetCreateTime(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetDomainName(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetErrorMsg(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetInstanceId(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskDetailNo(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskNo(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskStatus(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskStatusCode(v int32) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskType(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTaskTypeDescription(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetTryCount(v int32) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailHistoryResponseBodyPrePageCursor) SetUpdateTime(v string) *QueryTaskDetailHistoryResponseBodyPrePageCursor {
	s.UpdateTime = &v
	return s
}

type QueryTaskDetailHistoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTaskDetailHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTaskDetailHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailHistoryResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailHistoryResponse) SetHeaders(v map[string]*string) *QueryTaskDetailHistoryResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskDetailHistoryResponse) SetStatusCode(v int32) *QueryTaskDetailHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskDetailHistoryResponse) SetBody(v *QueryTaskDetailHistoryResponseBody) *QueryTaskDetailHistoryResponse {
	s.Body = v
	return s
}

type QueryTaskDetailListRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskNo       *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskStatus   *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTaskDetailListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailListRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListRequest) SetDomainName(v string) *QueryTaskDetailListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetInstanceId(v string) *QueryTaskDetailListRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetLang(v string) *QueryTaskDetailListRequest {
	s.Lang = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetPageNum(v int32) *QueryTaskDetailListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetPageSize(v int32) *QueryTaskDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetTaskNo(v string) *QueryTaskDetailListRequest {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetTaskStatus(v int32) *QueryTaskDetailListRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailListRequest) SetUserClientIp(v string) *QueryTaskDetailListRequest {
	s.UserClientIp = &v
	return s
}

type QueryTaskDetailListResponseBody struct {
	CurrentPageNum *int32                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryTaskDetailListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                                `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                                `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                               `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                               `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryTaskDetailListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListResponseBody) SetCurrentPageNum(v int32) *QueryTaskDetailListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetData(v *QueryTaskDetailListResponseBodyData) *QueryTaskDetailListResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetNextPage(v bool) *QueryTaskDetailListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetPageSize(v int32) *QueryTaskDetailListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetPrePage(v bool) *QueryTaskDetailListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetRequestId(v string) *QueryTaskDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetTotalItemNum(v int32) *QueryTaskDetailListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTaskDetailListResponseBody) SetTotalPageNum(v int32) *QueryTaskDetailListResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryTaskDetailListResponseBodyData struct {
	TaskDetail []*QueryTaskDetailListResponseBodyDataTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Repeated"`
}

func (s QueryTaskDetailListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListResponseBodyData) SetTaskDetail(v []*QueryTaskDetailListResponseBodyDataTaskDetail) *QueryTaskDetailListResponseBodyData {
	s.TaskDetail = v
	return s
}

type QueryTaskDetailListResponseBodyDataTaskDetail struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrorMsg            *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TaskDetailNo        *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskResult          *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	TryCount            *int32  `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryTaskDetailListResponseBodyDataTaskDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailListResponseBodyDataTaskDetail) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetCreateTime(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetDomainName(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.DomainName = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetErrorMsg(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetInstanceId(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.InstanceId = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskDetailNo(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskDetailNo = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskNo(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskResult(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskResult = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskStatus(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskStatusCode(v int32) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskType(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskType = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTaskTypeDescription(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetTryCount(v int32) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.TryCount = &v
	return s
}

func (s *QueryTaskDetailListResponseBodyDataTaskDetail) SetUpdateTime(v string) *QueryTaskDetailListResponseBodyDataTaskDetail {
	s.UpdateTime = &v
	return s
}

type QueryTaskDetailListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTaskDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTaskDetailListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskDetailListResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskDetailListResponse) SetHeaders(v map[string]*string) *QueryTaskDetailListResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskDetailListResponse) SetStatusCode(v int32) *QueryTaskDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskDetailListResponse) SetBody(v *QueryTaskDetailListResponseBody) *QueryTaskDetailListResponse {
	s.Body = v
	return s
}

type QueryTaskInfoHistoryRequest struct {
	BeginCreateTime  *int64  `json:"BeginCreateTime,omitempty" xml:"BeginCreateTime,omitempty"`
	CreateTimeCursor *int64  `json:"CreateTimeCursor,omitempty" xml:"CreateTimeCursor,omitempty"`
	EndCreateTime    *int64  `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskNoCursor     *string `json:"TaskNoCursor,omitempty" xml:"TaskNoCursor,omitempty"`
	UserClientIp     *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTaskInfoHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoHistoryRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryRequest) SetBeginCreateTime(v int64) *QueryTaskInfoHistoryRequest {
	s.BeginCreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetCreateTimeCursor(v int64) *QueryTaskInfoHistoryRequest {
	s.CreateTimeCursor = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetEndCreateTime(v int64) *QueryTaskInfoHistoryRequest {
	s.EndCreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetLang(v string) *QueryTaskInfoHistoryRequest {
	s.Lang = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetPageSize(v int32) *QueryTaskInfoHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetTaskNoCursor(v string) *QueryTaskInfoHistoryRequest {
	s.TaskNoCursor = &v
	return s
}

func (s *QueryTaskInfoHistoryRequest) SetUserClientIp(v string) *QueryTaskInfoHistoryRequest {
	s.UserClientIp = &v
	return s
}

type QueryTaskInfoHistoryResponseBody struct {
	CurrentPageCursor *QueryTaskInfoHistoryResponseBodyCurrentPageCursor `json:"CurrentPageCursor,omitempty" xml:"CurrentPageCursor,omitempty" type:"Struct"`
	NextPageCursor    *QueryTaskInfoHistoryResponseBodyNextPageCursor    `json:"NextPageCursor,omitempty" xml:"NextPageCursor,omitempty" type:"Struct"`
	Objects           []*QueryTaskInfoHistoryResponseBodyObjects         `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	PageSize          *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePageCursor     *QueryTaskInfoHistoryResponseBodyPrePageCursor     `json:"PrePageCursor,omitempty" xml:"PrePageCursor,omitempty" type:"Struct"`
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBody) SetCurrentPageCursor(v *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) *QueryTaskInfoHistoryResponseBody {
	s.CurrentPageCursor = v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetNextPageCursor(v *QueryTaskInfoHistoryResponseBodyNextPageCursor) *QueryTaskInfoHistoryResponseBody {
	s.NextPageCursor = v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetObjects(v []*QueryTaskInfoHistoryResponseBodyObjects) *QueryTaskInfoHistoryResponseBody {
	s.Objects = v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetPageSize(v int32) *QueryTaskInfoHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetPrePageCursor(v *QueryTaskInfoHistoryResponseBodyPrePageCursor) *QueryTaskInfoHistoryResponseBody {
	s.PrePageCursor = v
	return s
}

func (s *QueryTaskInfoHistoryResponseBody) SetRequestId(v string) *QueryTaskInfoHistoryResponseBody {
	s.RequestId = &v
	return s
}

type QueryTaskInfoHistoryResponseBodyCurrentPageCursor struct {
	Clientip            *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeLong      *int64  `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskNum             *int32  `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBodyCurrentPageCursor) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBodyCurrentPageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetClientip(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.Clientip = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetCreateTime(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetCreateTimeLong(v int64) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.CreateTimeLong = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskNo(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskNum(v int32) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskStatus(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskStatusCode(v int32) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskType(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyCurrentPageCursor) SetTaskTypeDescription(v string) *QueryTaskInfoHistoryResponseBodyCurrentPageCursor {
	s.TaskTypeDescription = &v
	return s
}

type QueryTaskInfoHistoryResponseBodyNextPageCursor struct {
	Clientip            *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeLong      *int64  `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskNum             *int32  `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBodyNextPageCursor) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBodyNextPageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetClientip(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.Clientip = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetCreateTime(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetCreateTimeLong(v int64) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.CreateTimeLong = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskNo(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskNum(v int32) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskStatus(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskStatusCode(v int32) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskType(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyNextPageCursor) SetTaskTypeDescription(v string) *QueryTaskInfoHistoryResponseBodyNextPageCursor {
	s.TaskTypeDescription = &v
	return s
}

type QueryTaskInfoHistoryResponseBodyObjects struct {
	Clientip            *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeLong      *int64  `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskNum             *int32  `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBodyObjects) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBodyObjects) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetClientip(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.Clientip = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetCreateTime(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetCreateTimeLong(v int64) *QueryTaskInfoHistoryResponseBodyObjects {
	s.CreateTimeLong = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskNo(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskNum(v int32) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskStatus(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskStatusCode(v int32) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskType(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskType = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyObjects) SetTaskTypeDescription(v string) *QueryTaskInfoHistoryResponseBodyObjects {
	s.TaskTypeDescription = &v
	return s
}

type QueryTaskInfoHistoryResponseBodyPrePageCursor struct {
	Clientip            *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeLong      *int64  `json:"CreateTimeLong,omitempty" xml:"CreateTimeLong,omitempty"`
	TaskNo              *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskNum             *int32  `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode      *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskInfoHistoryResponseBodyPrePageCursor) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoHistoryResponseBodyPrePageCursor) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetClientip(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.Clientip = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetCreateTime(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetCreateTimeLong(v int64) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.CreateTimeLong = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskNo(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskNum(v int32) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskStatus(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskStatusCode(v int32) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskType(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskType = &v
	return s
}

func (s *QueryTaskInfoHistoryResponseBodyPrePageCursor) SetTaskTypeDescription(v string) *QueryTaskInfoHistoryResponseBodyPrePageCursor {
	s.TaskTypeDescription = &v
	return s
}

type QueryTaskInfoHistoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTaskInfoHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTaskInfoHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskInfoHistoryResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskInfoHistoryResponse) SetHeaders(v map[string]*string) *QueryTaskInfoHistoryResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskInfoHistoryResponse) SetStatusCode(v int32) *QueryTaskInfoHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskInfoHistoryResponse) SetBody(v *QueryTaskInfoHistoryResponseBody) *QueryTaskInfoHistoryResponse {
	s.Body = v
	return s
}

type QueryTaskListRequest struct {
	BeginCreateTime *int64  `json:"BeginCreateTime,omitempty" xml:"BeginCreateTime,omitempty"`
	EndCreateTime   *int64  `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum         *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserClientIp    *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskListRequest) SetBeginCreateTime(v int64) *QueryTaskListRequest {
	s.BeginCreateTime = &v
	return s
}

func (s *QueryTaskListRequest) SetEndCreateTime(v int64) *QueryTaskListRequest {
	s.EndCreateTime = &v
	return s
}

func (s *QueryTaskListRequest) SetLang(v string) *QueryTaskListRequest {
	s.Lang = &v
	return s
}

func (s *QueryTaskListRequest) SetPageNum(v int32) *QueryTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTaskListRequest) SetPageSize(v int32) *QueryTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskListRequest) SetUserClientIp(v string) *QueryTaskListRequest {
	s.UserClientIp = &v
	return s
}

type QueryTaskListResponseBody struct {
	CurrentPageNum *int32                         `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryTaskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                          `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                          `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                         `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                         `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBody) SetCurrentPageNum(v int32) *QueryTaskListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTaskListResponseBody) SetData(v *QueryTaskListResponseBodyData) *QueryTaskListResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskListResponseBody) SetNextPage(v bool) *QueryTaskListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryTaskListResponseBody) SetPageSize(v int32) *QueryTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskListResponseBody) SetPrePage(v bool) *QueryTaskListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryTaskListResponseBody) SetRequestId(v string) *QueryTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskListResponseBody) SetTotalItemNum(v int32) *QueryTaskListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTaskListResponseBody) SetTotalPageNum(v int32) *QueryTaskListResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryTaskListResponseBodyData struct {
	TaskInfo []*QueryTaskListResponseBodyDataTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Repeated"`
}

func (s QueryTaskListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBodyData) SetTaskInfo(v []*QueryTaskListResponseBodyDataTaskInfo) *QueryTaskListResponseBodyData {
	s.TaskInfo = v
	return s
}

type QueryTaskListResponseBodyDataTaskInfo struct {
	Clientip             *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TaskCancelStatus     *string `json:"TaskCancelStatus,omitempty" xml:"TaskCancelStatus,omitempty"`
	TaskCancelStatusCode *int32  `json:"TaskCancelStatusCode,omitempty" xml:"TaskCancelStatusCode,omitempty"`
	TaskNo               *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskNum              *int32  `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	TaskStatus           *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode       *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType             *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription  *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskListResponseBodyDataTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListResponseBodyDataTaskInfo) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetClientip(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.Clientip = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetCreateTime(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskCancelStatus(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskCancelStatus = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskCancelStatusCode(v int32) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskCancelStatusCode = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskNo(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskNum(v int32) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskStatus(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskStatusCode(v int32) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskType(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskType = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskTypeDescription(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskTypeDescription = &v
	return s
}

type QueryTaskListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponse) SetHeaders(v map[string]*string) *QueryTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskListResponse) SetStatusCode(v int32) *QueryTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskListResponse) SetBody(v *QueryTaskListResponseBody) *QueryTaskListResponse {
	s.Body = v
	return s
}

type QueryTransferInByInstanceIdRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTransferInByInstanceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferInByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *QueryTransferInByInstanceIdRequest) SetInstanceId(v string) *QueryTransferInByInstanceIdRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTransferInByInstanceIdRequest) SetLang(v string) *QueryTransferInByInstanceIdRequest {
	s.Lang = &v
	return s
}

func (s *QueryTransferInByInstanceIdRequest) SetUserClientIp(v string) *QueryTransferInByInstanceIdRequest {
	s.UserClientIp = &v
	return s
}

type QueryTransferInByInstanceIdResponseBody struct {
	DomainName                                  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Email                                       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExpirationDate                              *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong                          *int64  `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	InstanceId                                  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModificationDate                            *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	ModificationDateLong                        *int64  `json:"ModificationDateLong,omitempty" xml:"ModificationDateLong,omitempty"`
	NeedMailCheck                               *bool   `json:"NeedMailCheck,omitempty" xml:"NeedMailCheck,omitempty"`
	ProgressBarType                             *int32  `json:"ProgressBarType,omitempty" xml:"ProgressBarType,omitempty"`
	RequestId                                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                                  *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultDate                                  *string `json:"ResultDate,omitempty" xml:"ResultDate,omitempty"`
	ResultDateLong                              *int64  `json:"ResultDateLong,omitempty" xml:"ResultDateLong,omitempty"`
	ResultMsg                                   *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	SimpleTransferInStatus                      *string `json:"SimpleTransferInStatus,omitempty" xml:"SimpleTransferInStatus,omitempty"`
	Status                                      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmissionDate                              *string `json:"SubmissionDate,omitempty" xml:"SubmissionDate,omitempty"`
	SubmissionDateLong                          *int64  `json:"SubmissionDateLong,omitempty" xml:"SubmissionDateLong,omitempty"`
	TransferAuthorizationCodeSubmissionDate     *string `json:"TransferAuthorizationCodeSubmissionDate,omitempty" xml:"TransferAuthorizationCodeSubmissionDate,omitempty"`
	TransferAuthorizationCodeSubmissionDateLong *int64  `json:"TransferAuthorizationCodeSubmissionDateLong,omitempty" xml:"TransferAuthorizationCodeSubmissionDateLong,omitempty"`
	UserId                                      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WhoisMailStatus                             *bool   `json:"WhoisMailStatus,omitempty" xml:"WhoisMailStatus,omitempty"`
}

func (s QueryTransferInByInstanceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferInByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTransferInByInstanceIdResponseBody) SetDomainName(v string) *QueryTransferInByInstanceIdResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetEmail(v string) *QueryTransferInByInstanceIdResponseBody {
	s.Email = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetExpirationDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetExpirationDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetInstanceId(v string) *QueryTransferInByInstanceIdResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetModificationDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ModificationDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetModificationDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.ModificationDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetNeedMailCheck(v bool) *QueryTransferInByInstanceIdResponseBody {
	s.NeedMailCheck = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetProgressBarType(v int32) *QueryTransferInByInstanceIdResponseBody {
	s.ProgressBarType = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetRequestId(v string) *QueryTransferInByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetResultCode(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetResultDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ResultDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetResultDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.ResultDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetResultMsg(v string) *QueryTransferInByInstanceIdResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetSimpleTransferInStatus(v string) *QueryTransferInByInstanceIdResponseBody {
	s.SimpleTransferInStatus = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetStatus(v int32) *QueryTransferInByInstanceIdResponseBody {
	s.Status = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetSubmissionDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.SubmissionDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetSubmissionDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.SubmissionDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetTransferAuthorizationCodeSubmissionDate(v string) *QueryTransferInByInstanceIdResponseBody {
	s.TransferAuthorizationCodeSubmissionDate = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetTransferAuthorizationCodeSubmissionDateLong(v int64) *QueryTransferInByInstanceIdResponseBody {
	s.TransferAuthorizationCodeSubmissionDateLong = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetUserId(v string) *QueryTransferInByInstanceIdResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponseBody) SetWhoisMailStatus(v bool) *QueryTransferInByInstanceIdResponseBody {
	s.WhoisMailStatus = &v
	return s
}

type QueryTransferInByInstanceIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTransferInByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTransferInByInstanceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferInByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *QueryTransferInByInstanceIdResponse) SetHeaders(v map[string]*string) *QueryTransferInByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *QueryTransferInByInstanceIdResponse) SetStatusCode(v int32) *QueryTransferInByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponse) SetBody(v *QueryTransferInByInstanceIdResponseBody) *QueryTransferInByInstanceIdResponse {
	s.Body = v
	return s
}

type QueryTransferInListRequest struct {
	DomainName             *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang                   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum                *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize               *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SimpleTransferInStatus *string `json:"SimpleTransferInStatus,omitempty" xml:"SimpleTransferInStatus,omitempty"`
	SubmissionEndDate      *int64  `json:"SubmissionEndDate,omitempty" xml:"SubmissionEndDate,omitempty"`
	SubmissionStartDate    *int64  `json:"SubmissionStartDate,omitempty" xml:"SubmissionStartDate,omitempty"`
	UserClientIp           *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTransferInListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferInListRequest) GoString() string {
	return s.String()
}

func (s *QueryTransferInListRequest) SetDomainName(v string) *QueryTransferInListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryTransferInListRequest) SetLang(v string) *QueryTransferInListRequest {
	s.Lang = &v
	return s
}

func (s *QueryTransferInListRequest) SetPageNum(v int32) *QueryTransferInListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTransferInListRequest) SetPageSize(v int32) *QueryTransferInListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTransferInListRequest) SetSimpleTransferInStatus(v string) *QueryTransferInListRequest {
	s.SimpleTransferInStatus = &v
	return s
}

func (s *QueryTransferInListRequest) SetSubmissionEndDate(v int64) *QueryTransferInListRequest {
	s.SubmissionEndDate = &v
	return s
}

func (s *QueryTransferInListRequest) SetSubmissionStartDate(v int64) *QueryTransferInListRequest {
	s.SubmissionStartDate = &v
	return s
}

func (s *QueryTransferInListRequest) SetUserClientIp(v string) *QueryTransferInListRequest {
	s.UserClientIp = &v
	return s
}

type QueryTransferInListResponseBody struct {
	CurrentPageNum *int32                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryTransferInListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                                `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                                `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                               `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                               `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryTransferInListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferInListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTransferInListResponseBody) SetCurrentPageNum(v int32) *QueryTransferInListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetData(v *QueryTransferInListResponseBodyData) *QueryTransferInListResponseBody {
	s.Data = v
	return s
}

func (s *QueryTransferInListResponseBody) SetNextPage(v bool) *QueryTransferInListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetPageSize(v int32) *QueryTransferInListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetPrePage(v bool) *QueryTransferInListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetRequestId(v string) *QueryTransferInListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetTotalItemNum(v int32) *QueryTransferInListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTransferInListResponseBody) SetTotalPageNum(v int32) *QueryTransferInListResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryTransferInListResponseBodyData struct {
	TransferInInfo []*QueryTransferInListResponseBodyDataTransferInInfo `json:"TransferInInfo,omitempty" xml:"TransferInInfo,omitempty" type:"Repeated"`
}

func (s QueryTransferInListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferInListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTransferInListResponseBodyData) SetTransferInInfo(v []*QueryTransferInListResponseBodyDataTransferInInfo) *QueryTransferInListResponseBodyData {
	s.TransferInInfo = v
	return s
}

type QueryTransferInListResponseBodyDataTransferInInfo struct {
	DomainName                                  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Email                                       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExpirationDate                              *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong                          *int64  `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	InstanceId                                  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModificationDate                            *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	ModificationDateLong                        *int64  `json:"ModificationDateLong,omitempty" xml:"ModificationDateLong,omitempty"`
	NeedMailCheck                               *bool   `json:"NeedMailCheck,omitempty" xml:"NeedMailCheck,omitempty"`
	ProgressBarType                             *int32  `json:"ProgressBarType,omitempty" xml:"ProgressBarType,omitempty"`
	ResultCode                                  *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultDate                                  *string `json:"ResultDate,omitempty" xml:"ResultDate,omitempty"`
	ResultDateLong                              *int64  `json:"ResultDateLong,omitempty" xml:"ResultDateLong,omitempty"`
	ResultMsg                                   *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	SimpleTransferInStatus                      *string `json:"SimpleTransferInStatus,omitempty" xml:"SimpleTransferInStatus,omitempty"`
	Status                                      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmissionDate                              *string `json:"SubmissionDate,omitempty" xml:"SubmissionDate,omitempty"`
	SubmissionDateLong                          *int64  `json:"SubmissionDateLong,omitempty" xml:"SubmissionDateLong,omitempty"`
	TransferAuthorizationCodeSubmissionDate     *string `json:"TransferAuthorizationCodeSubmissionDate,omitempty" xml:"TransferAuthorizationCodeSubmissionDate,omitempty"`
	TransferAuthorizationCodeSubmissionDateLong *int64  `json:"TransferAuthorizationCodeSubmissionDateLong,omitempty" xml:"TransferAuthorizationCodeSubmissionDateLong,omitempty"`
	UserId                                      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WhoisMailStatus                             *bool   `json:"WhoisMailStatus,omitempty" xml:"WhoisMailStatus,omitempty"`
}

func (s QueryTransferInListResponseBodyDataTransferInInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferInListResponseBodyDataTransferInInfo) GoString() string {
	return s.String()
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetDomainName(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.DomainName = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetEmail(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.Email = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetExpirationDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ExpirationDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetExpirationDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ExpirationDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetInstanceId(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.InstanceId = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetModificationDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ModificationDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetModificationDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ModificationDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetNeedMailCheck(v bool) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.NeedMailCheck = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetProgressBarType(v int32) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ProgressBarType = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetResultCode(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ResultCode = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetResultDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ResultDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetResultDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ResultDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetResultMsg(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.ResultMsg = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetSimpleTransferInStatus(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.SimpleTransferInStatus = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetStatus(v int32) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.Status = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetSubmissionDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.SubmissionDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetSubmissionDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.SubmissionDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetTransferAuthorizationCodeSubmissionDate(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.TransferAuthorizationCodeSubmissionDate = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetTransferAuthorizationCodeSubmissionDateLong(v int64) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.TransferAuthorizationCodeSubmissionDateLong = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetUserId(v string) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.UserId = &v
	return s
}

func (s *QueryTransferInListResponseBodyDataTransferInInfo) SetWhoisMailStatus(v bool) *QueryTransferInListResponseBodyDataTransferInInfo {
	s.WhoisMailStatus = &v
	return s
}

type QueryTransferInListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTransferInListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTransferInListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferInListResponse) GoString() string {
	return s.String()
}

func (s *QueryTransferInListResponse) SetHeaders(v map[string]*string) *QueryTransferInListResponse {
	s.Headers = v
	return s
}

func (s *QueryTransferInListResponse) SetStatusCode(v int32) *QueryTransferInListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTransferInListResponse) SetBody(v *QueryTransferInListResponseBody) *QueryTransferInListResponse {
	s.Body = v
	return s
}

type QueryTransferOutInfoRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTransferOutInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferOutInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryTransferOutInfoRequest) SetDomainName(v string) *QueryTransferOutInfoRequest {
	s.DomainName = &v
	return s
}

func (s *QueryTransferOutInfoRequest) SetLang(v string) *QueryTransferOutInfoRequest {
	s.Lang = &v
	return s
}

func (s *QueryTransferOutInfoRequest) SetUserClientIp(v string) *QueryTransferOutInfoRequest {
	s.UserClientIp = &v
	return s
}

type QueryTransferOutInfoResponseBody struct {
	Email                             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExpirationDate                    *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	PendingRequestDate                *string `json:"PendingRequestDate,omitempty" xml:"PendingRequestDate,omitempty"`
	RequestId                         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                        *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMsg                         *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Status                            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TransferAuthorizationCodeSendDate *string `json:"TransferAuthorizationCodeSendDate,omitempty" xml:"TransferAuthorizationCodeSendDate,omitempty"`
}

func (s QueryTransferOutInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferOutInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTransferOutInfoResponseBody) SetEmail(v string) *QueryTransferOutInfoResponseBody {
	s.Email = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetExpirationDate(v string) *QueryTransferOutInfoResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetPendingRequestDate(v string) *QueryTransferOutInfoResponseBody {
	s.PendingRequestDate = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetRequestId(v string) *QueryTransferOutInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetResultCode(v string) *QueryTransferOutInfoResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetResultMsg(v string) *QueryTransferOutInfoResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetStatus(v int32) *QueryTransferOutInfoResponseBody {
	s.Status = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetTransferAuthorizationCodeSendDate(v string) *QueryTransferOutInfoResponseBody {
	s.TransferAuthorizationCodeSendDate = &v
	return s
}

type QueryTransferOutInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTransferOutInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTransferOutInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferOutInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryTransferOutInfoResponse) SetHeaders(v map[string]*string) *QueryTransferOutInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryTransferOutInfoResponse) SetStatusCode(v int32) *QueryTransferOutInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTransferOutInfoResponse) SetBody(v *QueryTransferOutInfoResponseBody) *QueryTransferOutInfoResponse {
	s.Body = v
	return s
}

type RegistrantProfileRealNameVerificationRequest struct {
	IdentityCredential     *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	IdentityCredentialNo   *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	IdentityCredentialType *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	Lang                   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegistrantProfileID    *int64  `json:"RegistrantProfileID,omitempty" xml:"RegistrantProfileID,omitempty"`
	UserClientIp           *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s RegistrantProfileRealNameVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RegistrantProfileRealNameVerificationRequest) GoString() string {
	return s.String()
}

func (s *RegistrantProfileRealNameVerificationRequest) SetIdentityCredential(v string) *RegistrantProfileRealNameVerificationRequest {
	s.IdentityCredential = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetIdentityCredentialNo(v string) *RegistrantProfileRealNameVerificationRequest {
	s.IdentityCredentialNo = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetIdentityCredentialType(v string) *RegistrantProfileRealNameVerificationRequest {
	s.IdentityCredentialType = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetLang(v string) *RegistrantProfileRealNameVerificationRequest {
	s.Lang = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetRegistrantProfileID(v int64) *RegistrantProfileRealNameVerificationRequest {
	s.RegistrantProfileID = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationRequest) SetUserClientIp(v string) *RegistrantProfileRealNameVerificationRequest {
	s.UserClientIp = &v
	return s
}

type RegistrantProfileRealNameVerificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegistrantProfileRealNameVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegistrantProfileRealNameVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *RegistrantProfileRealNameVerificationResponseBody) SetRequestId(v string) *RegistrantProfileRealNameVerificationResponseBody {
	s.RequestId = &v
	return s
}

type RegistrantProfileRealNameVerificationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegistrantProfileRealNameVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegistrantProfileRealNameVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RegistrantProfileRealNameVerificationResponse) GoString() string {
	return s.String()
}

func (s *RegistrantProfileRealNameVerificationResponse) SetHeaders(v map[string]*string) *RegistrantProfileRealNameVerificationResponse {
	s.Headers = v
	return s
}

func (s *RegistrantProfileRealNameVerificationResponse) SetStatusCode(v int32) *RegistrantProfileRealNameVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RegistrantProfileRealNameVerificationResponse) SetBody(v *RegistrantProfileRealNameVerificationResponseBody) *RegistrantProfileRealNameVerificationResponse {
	s.Body = v
	return s
}

type ResendEmailVerificationRequest struct {
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ResendEmailVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationRequest) SetEmail(v string) *ResendEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *ResendEmailVerificationRequest) SetLang(v string) *ResendEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *ResendEmailVerificationRequest) SetUserClientIp(v string) *ResendEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

type ResendEmailVerificationResponseBody struct {
	FailList    []*ResendEmailVerificationResponseBodyFailList    `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Repeated"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList []*ResendEmailVerificationResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Repeated"`
}

func (s ResendEmailVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationResponseBody) SetFailList(v []*ResendEmailVerificationResponseBodyFailList) *ResendEmailVerificationResponseBody {
	s.FailList = v
	return s
}

func (s *ResendEmailVerificationResponseBody) SetRequestId(v string) *ResendEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendEmailVerificationResponseBody) SetSuccessList(v []*ResendEmailVerificationResponseBodySuccessList) *ResendEmailVerificationResponseBody {
	s.SuccessList = v
	return s
}

type ResendEmailVerificationResponseBodyFailList struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ResendEmailVerificationResponseBodyFailList) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailVerificationResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationResponseBodyFailList) SetCode(v string) *ResendEmailVerificationResponseBodyFailList {
	s.Code = &v
	return s
}

func (s *ResendEmailVerificationResponseBodyFailList) SetEmail(v string) *ResendEmailVerificationResponseBodyFailList {
	s.Email = &v
	return s
}

func (s *ResendEmailVerificationResponseBodyFailList) SetMessage(v string) *ResendEmailVerificationResponseBodyFailList {
	s.Message = &v
	return s
}

type ResendEmailVerificationResponseBodySuccessList struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ResendEmailVerificationResponseBodySuccessList) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailVerificationResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationResponseBodySuccessList) SetCode(v string) *ResendEmailVerificationResponseBodySuccessList {
	s.Code = &v
	return s
}

func (s *ResendEmailVerificationResponseBodySuccessList) SetEmail(v string) *ResendEmailVerificationResponseBodySuccessList {
	s.Email = &v
	return s
}

func (s *ResendEmailVerificationResponseBodySuccessList) SetMessage(v string) *ResendEmailVerificationResponseBodySuccessList {
	s.Message = &v
	return s
}

type ResendEmailVerificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResendEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResendEmailVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationResponse) SetHeaders(v map[string]*string) *ResendEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *ResendEmailVerificationResponse) SetStatusCode(v int32) *ResendEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendEmailVerificationResponse) SetBody(v *ResendEmailVerificationResponseBody) *ResendEmailVerificationResponse {
	s.Body = v
	return s
}

type ResetQualificationVerificationRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ResetQualificationVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetQualificationVerificationRequest) GoString() string {
	return s.String()
}

func (s *ResetQualificationVerificationRequest) SetInstanceId(v string) *ResetQualificationVerificationRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetQualificationVerificationRequest) SetLang(v string) *ResetQualificationVerificationRequest {
	s.Lang = &v
	return s
}

func (s *ResetQualificationVerificationRequest) SetUserClientIp(v string) *ResetQualificationVerificationRequest {
	s.UserClientIp = &v
	return s
}

type ResetQualificationVerificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetQualificationVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetQualificationVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ResetQualificationVerificationResponseBody) SetRequestId(v string) *ResetQualificationVerificationResponseBody {
	s.RequestId = &v
	return s
}

type ResetQualificationVerificationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetQualificationVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetQualificationVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetQualificationVerificationResponse) GoString() string {
	return s.String()
}

func (s *ResetQualificationVerificationResponse) SetHeaders(v map[string]*string) *ResetQualificationVerificationResponse {
	s.Headers = v
	return s
}

func (s *ResetQualificationVerificationResponse) SetStatusCode(v int32) *ResetQualificationVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetQualificationVerificationResponse) SetBody(v *ResetQualificationVerificationResponseBody) *ResetQualificationVerificationResponse {
	s.Body = v
	return s
}

type SaveBatchDomainRemarkRequest struct {
	InstanceIds  *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchDomainRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchDomainRemarkRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchDomainRemarkRequest) SetInstanceIds(v string) *SaveBatchDomainRemarkRequest {
	s.InstanceIds = &v
	return s
}

func (s *SaveBatchDomainRemarkRequest) SetLang(v string) *SaveBatchDomainRemarkRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchDomainRemarkRequest) SetRemark(v string) *SaveBatchDomainRemarkRequest {
	s.Remark = &v
	return s
}

func (s *SaveBatchDomainRemarkRequest) SetUserClientIp(v string) *SaveBatchDomainRemarkRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchDomainRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveBatchDomainRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchDomainRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchDomainRemarkResponseBody) SetRequestId(v string) *SaveBatchDomainRemarkResponseBody {
	s.RequestId = &v
	return s
}

type SaveBatchDomainRemarkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchDomainRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchDomainRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchDomainRemarkResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchDomainRemarkResponse) SetHeaders(v map[string]*string) *SaveBatchDomainRemarkResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchDomainRemarkResponse) SetStatusCode(v int32) *SaveBatchDomainRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchDomainRemarkResponse) SetBody(v *SaveBatchDomainRemarkResponseBody) *SaveBatchDomainRemarkResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForCreatingOrderActivateRequest struct {
	CouponNo           *string                                                           `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	Lang               *string                                                           `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OrderActivateParam []*SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam `json:"OrderActivateParam,omitempty" xml:"OrderActivateParam,omitempty" type:"Repeated"`
	PromotionNo        *string                                                           `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	UseCoupon          *bool                                                             `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	UsePromotion       *bool                                                             `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	UserClientIp       *string                                                           `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderActivateRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderActivateRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetCouponNo(v string) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetLang(v string) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetOrderActivateParam(v []*SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.OrderActivateParam = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequest) SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderActivateRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam struct {
	Address                   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AliyunDns                 *bool   `json:"AliyunDns,omitempty" xml:"AliyunDns,omitempty"`
	City                      *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                   *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Dns1                      *string `json:"Dns1,omitempty" xml:"Dns1,omitempty"`
	Dns2                      *string `json:"Dns2,omitempty" xml:"Dns2,omitempty"`
	DomainName                *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Email                     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EnableDomainProxy         *bool   `json:"EnableDomainProxy,omitempty" xml:"EnableDomainProxy,omitempty"`
	PermitPremiumActivation   *bool   `json:"PermitPremiumActivation,omitempty" xml:"PermitPremiumActivation,omitempty"`
	PostalCode                *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                  *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName            *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization    *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantProfileId       *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RegistrantType            *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	SubscriptionDuration      *int32  `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
	TelArea                   *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                    *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                 *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	TrademarkDomainActivation *bool   `json:"TrademarkDomainActivation,omitempty" xml:"TrademarkDomainActivation,omitempty"`
	ZhAddress                 *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                    *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince                *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName          *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization  *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetAddress(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Address = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetAliyunDns(v bool) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.AliyunDns = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetCity(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.City = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetCountry(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Country = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetDns1(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Dns1 = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetDns2(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Dns2 = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetDomainName(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.DomainName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetEmail(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Email = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetEnableDomainProxy(v bool) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.EnableDomainProxy = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetPermitPremiumActivation(v bool) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.PermitPremiumActivation = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetPostalCode(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.PostalCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetProvince(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Province = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetRegistrantName(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.RegistrantName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetRegistrantOrganization(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetRegistrantProfileId(v int64) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetRegistrantType(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.RegistrantType = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetSubscriptionDuration(v int32) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.SubscriptionDuration = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetTelArea(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.TelArea = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetTelExt(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.TelExt = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetTelephone(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.Telephone = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetTrademarkDomainActivation(v bool) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.TrademarkDomainActivation = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhAddress(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhAddress = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhCity(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhCity = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhProvince(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhProvince = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhRegistrantName(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam) SetZhRegistrantOrganization(v string) *SaveBatchTaskForCreatingOrderActivateRequestOrderActivateParam {
	s.ZhRegistrantOrganization = &v
	return s
}

type SaveBatchTaskForCreatingOrderActivateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderActivateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderActivateResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderActivateResponseBody) SetRequestId(v string) *SaveBatchTaskForCreatingOrderActivateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateResponseBody) SetTaskNo(v string) *SaveBatchTaskForCreatingOrderActivateResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForCreatingOrderActivateResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForCreatingOrderActivateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForCreatingOrderActivateResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderActivateResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderActivateResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderActivateResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) SetBody(v *SaveBatchTaskForCreatingOrderActivateResponseBody) *SaveBatchTaskForCreatingOrderActivateResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForCreatingOrderRedeemRequest struct {
	CouponNo         *string                                                       `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	Lang             *string                                                       `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OrderRedeemParam []*SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam `json:"OrderRedeemParam,omitempty" xml:"OrderRedeemParam,omitempty" type:"Repeated"`
	PromotionNo      *string                                                       `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	UseCoupon        *bool                                                         `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	UsePromotion     *bool                                                         `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	UserClientIp     *string                                                       `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRedeemRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRedeemRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetCouponNo(v string) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetLang(v string) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetOrderRedeemParam(v []*SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.OrderRedeemParam = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequest) SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderRedeemRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam struct {
	CurrentExpirationDate *int64  `json:"CurrentExpirationDate,omitempty" xml:"CurrentExpirationDate,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) SetCurrentExpirationDate(v int64) *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam {
	s.CurrentExpirationDate = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam) SetDomainName(v string) *SaveBatchTaskForCreatingOrderRedeemRequestOrderRedeemParam {
	s.DomainName = &v
	return s
}

type SaveBatchTaskForCreatingOrderRedeemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRedeemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRedeemResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponseBody) SetRequestId(v string) *SaveBatchTaskForCreatingOrderRedeemResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponseBody) SetTaskNo(v string) *SaveBatchTaskForCreatingOrderRedeemResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForCreatingOrderRedeemResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForCreatingOrderRedeemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForCreatingOrderRedeemResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRedeemResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderRedeemResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderRedeemResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) SetBody(v *SaveBatchTaskForCreatingOrderRedeemResponseBody) *SaveBatchTaskForCreatingOrderRedeemResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForCreatingOrderRenewRequest struct {
	CouponNo        *string                                                     `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	Lang            *string                                                     `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OrderRenewParam []*SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam `json:"OrderRenewParam,omitempty" xml:"OrderRenewParam,omitempty" type:"Repeated"`
	PromotionNo     *string                                                     `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	UseCoupon       *bool                                                       `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	UsePromotion    *bool                                                       `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	UserClientIp    *string                                                     `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRenewRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRenewRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetCouponNo(v string) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetLang(v string) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetOrderRenewParam(v []*SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.OrderRenewParam = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequest) SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderRenewRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam struct {
	CurrentExpirationDate *int64  `json:"CurrentExpirationDate,omitempty" xml:"CurrentExpirationDate,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SubscriptionDuration  *int32  `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) SetCurrentExpirationDate(v int64) *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam {
	s.CurrentExpirationDate = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) SetDomainName(v string) *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam {
	s.DomainName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam) SetSubscriptionDuration(v int32) *SaveBatchTaskForCreatingOrderRenewRequestOrderRenewParam {
	s.SubscriptionDuration = &v
	return s
}

type SaveBatchTaskForCreatingOrderRenewResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRenewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRenewResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRenewResponseBody) SetRequestId(v string) *SaveBatchTaskForCreatingOrderRenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewResponseBody) SetTaskNo(v string) *SaveBatchTaskForCreatingOrderRenewResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForCreatingOrderRenewResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForCreatingOrderRenewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForCreatingOrderRenewResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRenewResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderRenewResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderRenewResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) SetBody(v *SaveBatchTaskForCreatingOrderRenewResponseBody) *SaveBatchTaskForCreatingOrderRenewResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForCreatingOrderTransferRequest struct {
	CouponNo           *string                                                           `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	Lang               *string                                                           `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OrderTransferParam []*SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam `json:"OrderTransferParam,omitempty" xml:"OrderTransferParam,omitempty" type:"Repeated"`
	PromotionNo        *string                                                           `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	UseCoupon          *bool                                                             `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	UsePromotion       *bool                                                             `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	UserClientIp       *string                                                           `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderTransferRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetCouponNo(v string) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetLang(v string) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetOrderTransferParam(v []*SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.OrderTransferParam = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetPromotionNo(v string) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetUseCoupon(v bool) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetUsePromotion(v bool) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequest) SetUserClientIp(v string) *SaveBatchTaskForCreatingOrderTransferRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam struct {
	AuthorizationCode     *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PermitPremiumTransfer *bool   `json:"PermitPremiumTransfer,omitempty" xml:"PermitPremiumTransfer,omitempty"`
	RegistrantProfileId   *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) SetAuthorizationCode(v string) *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam {
	s.AuthorizationCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) SetDomainName(v string) *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam {
	s.DomainName = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) SetPermitPremiumTransfer(v bool) *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam {
	s.PermitPremiumTransfer = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam) SetRegistrantProfileId(v int64) *SaveBatchTaskForCreatingOrderTransferRequestOrderTransferParam {
	s.RegistrantProfileId = &v
	return s
}

type SaveBatchTaskForCreatingOrderTransferResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderTransferResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderTransferResponseBody) SetRequestId(v string) *SaveBatchTaskForCreatingOrderTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferResponseBody) SetTaskNo(v string) *SaveBatchTaskForCreatingOrderTransferResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForCreatingOrderTransferResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForCreatingOrderTransferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForCreatingOrderTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderTransferResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderTransferResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) SetBody(v *SaveBatchTaskForCreatingOrderTransferResponseBody) *SaveBatchTaskForCreatingOrderTransferResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForDomainNameProxyServiceRequest struct {
	DomainName   []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Status       *bool     `json:"Status,omitempty" xml:"Status,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForDomainNameProxyServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForDomainNameProxyServiceRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetDomainName(v []*string) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetLang(v string) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetStatus(v bool) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.Status = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetUserClientIp(v string) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchTaskForDomainNameProxyServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForDomainNameProxyServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForDomainNameProxyServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponseBody) SetRequestId(v string) *SaveBatchTaskForDomainNameProxyServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponseBody) SetTaskNo(v string) *SaveBatchTaskForDomainNameProxyServiceResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForDomainNameProxyServiceResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForDomainNameProxyServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForDomainNameProxyServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForDomainNameProxyServiceResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForDomainNameProxyServiceResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) SetStatusCode(v int32) *SaveBatchTaskForDomainNameProxyServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceResponse) SetBody(v *SaveBatchTaskForDomainNameProxyServiceResponseBody) *SaveBatchTaskForDomainNameProxyServiceResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForModifyingDomainDnsRequest struct {
	AliyunDns        *bool     `json:"AliyunDns,omitempty" xml:"AliyunDns,omitempty"`
	DomainName       []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	DomainNameServer []*string `json:"DomainNameServer,omitempty" xml:"DomainNameServer,omitempty" type:"Repeated"`
	Lang             *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp     *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForModifyingDomainDnsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForModifyingDomainDnsRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetAliyunDns(v bool) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.AliyunDns = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetDomainName(v []*string) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetDomainNameServer(v []*string) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.DomainNameServer = v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetLang(v string) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetUserClientIp(v string) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchTaskForModifyingDomainDnsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForModifyingDomainDnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForModifyingDomainDnsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForModifyingDomainDnsResponseBody) SetRequestId(v string) *SaveBatchTaskForModifyingDomainDnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsResponseBody) SetTaskNo(v string) *SaveBatchTaskForModifyingDomainDnsResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForModifyingDomainDnsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForModifyingDomainDnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForModifyingDomainDnsResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForModifyingDomainDnsResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForModifyingDomainDnsResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) SetStatusCode(v int32) *SaveBatchTaskForModifyingDomainDnsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsResponse) SetBody(v *SaveBatchTaskForModifyingDomainDnsResponseBody) *SaveBatchTaskForModifyingDomainDnsResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForReserveDropListDomainRequest struct {
	ContactTemplateId *string                                                `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	Domains           []*SaveBatchTaskForReserveDropListDomainRequestDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s SaveBatchTaskForReserveDropListDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForReserveDropListDomainRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForReserveDropListDomainRequest) SetContactTemplateId(v string) *SaveBatchTaskForReserveDropListDomainRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainRequest) SetDomains(v []*SaveBatchTaskForReserveDropListDomainRequestDomains) *SaveBatchTaskForReserveDropListDomainRequest {
	s.Domains = v
	return s
}

type SaveBatchTaskForReserveDropListDomainRequestDomains struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SaveBatchTaskForReserveDropListDomainRequestDomains) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForReserveDropListDomainRequestDomains) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForReserveDropListDomainRequestDomains) SetDomainName(v string) *SaveBatchTaskForReserveDropListDomainRequestDomains {
	s.DomainName = &v
	return s
}

type SaveBatchTaskForReserveDropListDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForReserveDropListDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForReserveDropListDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForReserveDropListDomainResponseBody) SetRequestId(v string) *SaveBatchTaskForReserveDropListDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainResponseBody) SetTaskNo(v string) *SaveBatchTaskForReserveDropListDomainResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForReserveDropListDomainResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForReserveDropListDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForReserveDropListDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForReserveDropListDomainResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForReserveDropListDomainResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) SetStatusCode(v int32) *SaveBatchTaskForReserveDropListDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainResponse) SetBody(v *SaveBatchTaskForReserveDropListDomainResponseBody) *SaveBatchTaskForReserveDropListDomainResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForTransferProhibitionLockRequest struct {
	DomainName   []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Status       *bool     `json:"Status,omitempty" xml:"Status,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForTransferProhibitionLockRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForTransferProhibitionLockRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) SetDomainName(v []*string) *SaveBatchTaskForTransferProhibitionLockRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) SetLang(v string) *SaveBatchTaskForTransferProhibitionLockRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) SetStatus(v bool) *SaveBatchTaskForTransferProhibitionLockRequest {
	s.Status = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockRequest) SetUserClientIp(v string) *SaveBatchTaskForTransferProhibitionLockRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchTaskForTransferProhibitionLockResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForTransferProhibitionLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForTransferProhibitionLockResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferProhibitionLockResponseBody) SetRequestId(v string) *SaveBatchTaskForTransferProhibitionLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockResponseBody) SetTaskNo(v string) *SaveBatchTaskForTransferProhibitionLockResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForTransferProhibitionLockResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForTransferProhibitionLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForTransferProhibitionLockResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForTransferProhibitionLockResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForTransferProhibitionLockResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) SetStatusCode(v int32) *SaveBatchTaskForTransferProhibitionLockResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) SetBody(v *SaveBatchTaskForTransferProhibitionLockResponseBody) *SaveBatchTaskForTransferProhibitionLockResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForUpdateProhibitionLockRequest struct {
	DomainName   []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Status       *bool     `json:"Status,omitempty" xml:"Status,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForUpdateProhibitionLockRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForUpdateProhibitionLockRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) SetDomainName(v []*string) *SaveBatchTaskForUpdateProhibitionLockRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) SetLang(v string) *SaveBatchTaskForUpdateProhibitionLockRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) SetStatus(v bool) *SaveBatchTaskForUpdateProhibitionLockRequest {
	s.Status = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockRequest) SetUserClientIp(v string) *SaveBatchTaskForUpdateProhibitionLockRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchTaskForUpdateProhibitionLockResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForUpdateProhibitionLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForUpdateProhibitionLockResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponseBody) SetRequestId(v string) *SaveBatchTaskForUpdateProhibitionLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponseBody) SetTaskNo(v string) *SaveBatchTaskForUpdateProhibitionLockResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForUpdateProhibitionLockResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForUpdateProhibitionLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForUpdateProhibitionLockResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForUpdateProhibitionLockResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForUpdateProhibitionLockResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) SetStatusCode(v int32) *SaveBatchTaskForUpdateProhibitionLockResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) SetBody(v *SaveBatchTaskForUpdateProhibitionLockResponseBody) *SaveBatchTaskForUpdateProhibitionLockResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForUpdatingContactInfoByNewContactRequest struct {
	Address                  *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	City                     *string   `json:"City,omitempty" xml:"City,omitempty"`
	ContactType              *string   `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	Country                  *string   `json:"Country,omitempty" xml:"Country,omitempty"`
	DomainName               []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Email                    *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang                     *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PostalCode               *string   `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                 *string   `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName           *string   `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization   *string   `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantType           *string   `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	TelArea                  *string   `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                   *string   `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                *string   `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	TransferOutProhibited    *bool     `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	UserClientIp             *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZhAddress                *string   `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string   `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string   `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string   `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string   `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetAddress(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Address = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetCity(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.City = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetContactType(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ContactType = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetCountry(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Country = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetDomainName(v []*string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetEmail(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Email = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetLang(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetPostalCode(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetProvince(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Province = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetRegistrantName(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetRegistrantOrganization(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetRegistrantType(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.RegistrantType = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetTelArea(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.TelArea = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetTelExt(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.TelExt = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetTelephone(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.Telephone = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetTransferOutProhibited(v bool) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.TransferOutProhibited = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetUserClientIp(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhAddress(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhAddress = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhCity(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhCity = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhProvince(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhProvince = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhRegistrantName(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) SetZhRegistrantOrganization(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

type SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) SetRequestId(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) SetTaskNo(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForUpdatingContactInfoByNewContactResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForUpdatingContactInfoByNewContactResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) SetStatusCode(v int32) *SaveBatchTaskForUpdatingContactInfoByNewContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) SetBody(v *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) *SaveBatchTaskForUpdatingContactInfoByNewContactResponse {
	s.Body = v
	return s
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest struct {
	ContactType           *string   `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	DomainName            []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Lang                  *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegistrantProfileId   *int64    `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	TransferOutProhibited *bool     `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	UserClientIp          *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetContactType(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.ContactType = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetDomainName(v []*string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetLang(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetRegistrantProfileId(v int64) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetTransferOutProhibited(v bool) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.TransferOutProhibited = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) SetUserClientIp(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest {
	s.UserClientIp = &v
	return s
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) SetRequestId(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) SetTaskNo(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody {
	s.TaskNo = &v
	return s
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse struct {
	Headers    map[string]*string                                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) SetStatusCode(v int32) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) SetBody(v *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse {
	s.Body = v
	return s
}

type SaveDomainGroupRequest struct {
	DomainGroupId   *int64  `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp    *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveDomainGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *SaveDomainGroupRequest) SetDomainGroupId(v int64) *SaveDomainGroupRequest {
	s.DomainGroupId = &v
	return s
}

func (s *SaveDomainGroupRequest) SetDomainGroupName(v string) *SaveDomainGroupRequest {
	s.DomainGroupName = &v
	return s
}

func (s *SaveDomainGroupRequest) SetLang(v string) *SaveDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *SaveDomainGroupRequest) SetUserClientIp(v string) *SaveDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

type SaveDomainGroupResponseBody struct {
	BeingDeleted      *bool   `json:"BeingDeleted,omitempty" xml:"BeingDeleted,omitempty"`
	CreationDate      *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	DomainGroupId     *int64  `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName   *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	DomainGroupStatus *string `json:"DomainGroupStatus,omitempty" xml:"DomainGroupStatus,omitempty"`
	ModificationDate  *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNumber       *int32  `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s SaveDomainGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SaveDomainGroupResponseBody) SetBeingDeleted(v bool) *SaveDomainGroupResponseBody {
	s.BeingDeleted = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetCreationDate(v string) *SaveDomainGroupResponseBody {
	s.CreationDate = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetDomainGroupId(v int64) *SaveDomainGroupResponseBody {
	s.DomainGroupId = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetDomainGroupName(v string) *SaveDomainGroupResponseBody {
	s.DomainGroupName = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetDomainGroupStatus(v string) *SaveDomainGroupResponseBody {
	s.DomainGroupStatus = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetModificationDate(v string) *SaveDomainGroupResponseBody {
	s.ModificationDate = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetRequestId(v string) *SaveDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetTotalNumber(v int32) *SaveDomainGroupResponseBody {
	s.TotalNumber = &v
	return s
}

type SaveDomainGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveDomainGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *SaveDomainGroupResponse) SetHeaders(v map[string]*string) *SaveDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *SaveDomainGroupResponse) SetStatusCode(v int32) *SaveDomainGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveDomainGroupResponse) SetBody(v *SaveDomainGroupResponseBody) *SaveDomainGroupResponse {
	s.Body = v
	return s
}

type SaveRegistrantProfileRequest struct {
	Address                  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	City                     *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	DefaultRegistrantProfile *bool   `json:"DefaultRegistrantProfile,omitempty" xml:"DefaultRegistrantProfile,omitempty"`
	Email                    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang                     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PostalCode               *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                 *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName           *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization   *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantProfileId      *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RegistrantProfileType    *string `json:"RegistrantProfileType,omitempty" xml:"RegistrantProfileType,omitempty"`
	RegistrantType           *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	TelArea                  *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                   *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	UserClientIp             *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZhAddress                *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s SaveRegistrantProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveRegistrantProfileRequest) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileRequest) SetAddress(v string) *SaveRegistrantProfileRequest {
	s.Address = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetCity(v string) *SaveRegistrantProfileRequest {
	s.City = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetCountry(v string) *SaveRegistrantProfileRequest {
	s.Country = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetDefaultRegistrantProfile(v bool) *SaveRegistrantProfileRequest {
	s.DefaultRegistrantProfile = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetEmail(v string) *SaveRegistrantProfileRequest {
	s.Email = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetLang(v string) *SaveRegistrantProfileRequest {
	s.Lang = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetPostalCode(v string) *SaveRegistrantProfileRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetProvince(v string) *SaveRegistrantProfileRequest {
	s.Province = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantName(v string) *SaveRegistrantProfileRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantOrganization(v string) *SaveRegistrantProfileRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantProfileId(v int64) *SaveRegistrantProfileRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantProfileType(v string) *SaveRegistrantProfileRequest {
	s.RegistrantProfileType = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetRegistrantType(v string) *SaveRegistrantProfileRequest {
	s.RegistrantType = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetTelArea(v string) *SaveRegistrantProfileRequest {
	s.TelArea = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetTelExt(v string) *SaveRegistrantProfileRequest {
	s.TelExt = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetTelephone(v string) *SaveRegistrantProfileRequest {
	s.Telephone = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetUserClientIp(v string) *SaveRegistrantProfileRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetZhAddress(v string) *SaveRegistrantProfileRequest {
	s.ZhAddress = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetZhCity(v string) *SaveRegistrantProfileRequest {
	s.ZhCity = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetZhProvince(v string) *SaveRegistrantProfileRequest {
	s.ZhProvince = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetZhRegistrantName(v string) *SaveRegistrantProfileRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveRegistrantProfileRequest) SetZhRegistrantOrganization(v string) *SaveRegistrantProfileRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

type SaveRegistrantProfileResponseBody struct {
	RegistrantProfileId *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveRegistrantProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveRegistrantProfileResponseBody) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileResponseBody) SetRegistrantProfileId(v int64) *SaveRegistrantProfileResponseBody {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveRegistrantProfileResponseBody) SetRequestId(v string) *SaveRegistrantProfileResponseBody {
	s.RequestId = &v
	return s
}

type SaveRegistrantProfileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveRegistrantProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveRegistrantProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveRegistrantProfileResponse) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileResponse) SetHeaders(v map[string]*string) *SaveRegistrantProfileResponse {
	s.Headers = v
	return s
}

func (s *SaveRegistrantProfileResponse) SetStatusCode(v int32) *SaveRegistrantProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveRegistrantProfileResponse) SetBody(v *SaveRegistrantProfileResponseBody) *SaveRegistrantProfileResponse {
	s.Body = v
	return s
}

type SaveRegistrantProfileRealNameVerificationRequest struct {
	Address                  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	City                     *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Email                    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	IdentityCredential       *string `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	IdentityCredentialNo     *string `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	IdentityCredentialType   *string `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	Lang                     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PostalCode               *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                 *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName           *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization   *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantProfileId      *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RegistrantProfileType    *string `json:"RegistrantProfileType,omitempty" xml:"RegistrantProfileType,omitempty"`
	RegistrantType           *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	TelArea                  *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                   *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	UserClientIp             *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZhAddress                *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s SaveRegistrantProfileRealNameVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveRegistrantProfileRealNameVerificationRequest) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetAddress(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Address = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetCity(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.City = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetCountry(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Country = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetEmail(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Email = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetIdentityCredential(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.IdentityCredential = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetIdentityCredentialNo(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.IdentityCredentialNo = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetIdentityCredentialType(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.IdentityCredentialType = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetLang(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Lang = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetPostalCode(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetProvince(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Province = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantName(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantOrganization(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantProfileId(v int64) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantProfileType(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantProfileType = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetRegistrantType(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.RegistrantType = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetTelArea(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.TelArea = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetTelExt(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.TelExt = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetTelephone(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.Telephone = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetUserClientIp(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhAddress(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhAddress = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhCity(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhCity = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhProvince(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhProvince = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhRegistrantName(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationRequest) SetZhRegistrantOrganization(v string) *SaveRegistrantProfileRealNameVerificationRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

type SaveRegistrantProfileRealNameVerificationResponseBody struct {
	RegistrantProfileId *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveRegistrantProfileRealNameVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveRegistrantProfileRealNameVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileRealNameVerificationResponseBody) SetRegistrantProfileId(v int64) *SaveRegistrantProfileRealNameVerificationResponseBody {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationResponseBody) SetRequestId(v string) *SaveRegistrantProfileRealNameVerificationResponseBody {
	s.RequestId = &v
	return s
}

type SaveRegistrantProfileRealNameVerificationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveRegistrantProfileRealNameVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveRegistrantProfileRealNameVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveRegistrantProfileRealNameVerificationResponse) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) SetHeaders(v map[string]*string) *SaveRegistrantProfileRealNameVerificationResponse {
	s.Headers = v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) SetStatusCode(v int32) *SaveRegistrantProfileRealNameVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveRegistrantProfileRealNameVerificationResponse) SetBody(v *SaveRegistrantProfileRealNameVerificationResponseBody) *SaveRegistrantProfileRealNameVerificationResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForAddingDSRecordRequest struct {
	Algorithm    *int32  `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Digest       *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	DigestType   *int32  `json:"DigestType,omitempty" xml:"DigestType,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	KeyTag       *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForAddingDSRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForAddingDSRecordRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetAlgorithm(v int32) *SaveSingleTaskForAddingDSRecordRequest {
	s.Algorithm = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetDigest(v string) *SaveSingleTaskForAddingDSRecordRequest {
	s.Digest = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetDigestType(v int32) *SaveSingleTaskForAddingDSRecordRequest {
	s.DigestType = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetDomainName(v string) *SaveSingleTaskForAddingDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetKeyTag(v int32) *SaveSingleTaskForAddingDSRecordRequest {
	s.KeyTag = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetLang(v string) *SaveSingleTaskForAddingDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordRequest) SetUserClientIp(v string) *SaveSingleTaskForAddingDSRecordRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForAddingDSRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForAddingDSRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForAddingDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAddingDSRecordResponseBody) SetRequestId(v string) *SaveSingleTaskForAddingDSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordResponseBody) SetTaskNo(v string) *SaveSingleTaskForAddingDSRecordResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForAddingDSRecordResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForAddingDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForAddingDSRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForAddingDSRecordResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAddingDSRecordResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForAddingDSRecordResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordResponse) SetStatusCode(v int32) *SaveSingleTaskForAddingDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordResponse) SetBody(v *SaveSingleTaskForAddingDSRecordResponseBody) *SaveSingleTaskForAddingDSRecordResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForApprovingTransferOutRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForApprovingTransferOutRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForApprovingTransferOutRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) SetDomainName(v string) *SaveSingleTaskForApprovingTransferOutRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) SetLang(v string) *SaveSingleTaskForApprovingTransferOutRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutRequest) SetUserClientIp(v string) *SaveSingleTaskForApprovingTransferOutRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForApprovingTransferOutResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForApprovingTransferOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForApprovingTransferOutResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForApprovingTransferOutResponseBody) SetRequestId(v string) *SaveSingleTaskForApprovingTransferOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutResponseBody) SetTaskNo(v string) *SaveSingleTaskForApprovingTransferOutResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForApprovingTransferOutResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForApprovingTransferOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForApprovingTransferOutResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForApprovingTransferOutResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForApprovingTransferOutResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) SetStatusCode(v int32) *SaveSingleTaskForApprovingTransferOutResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) SetBody(v *SaveSingleTaskForApprovingTransferOutResponseBody) *SaveSingleTaskForApprovingTransferOutResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForAssociatingEnsRequest struct {
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForAssociatingEnsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForAssociatingEnsRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAssociatingEnsRequest) SetAddress(v string) *SaveSingleTaskForAssociatingEnsRequest {
	s.Address = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsRequest) SetDomainName(v string) *SaveSingleTaskForAssociatingEnsRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsRequest) SetLang(v string) *SaveSingleTaskForAssociatingEnsRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsRequest) SetUserClientIp(v string) *SaveSingleTaskForAssociatingEnsRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForAssociatingEnsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForAssociatingEnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForAssociatingEnsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAssociatingEnsResponseBody) SetRequestId(v string) *SaveSingleTaskForAssociatingEnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsResponseBody) SetTaskNo(v string) *SaveSingleTaskForAssociatingEnsResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForAssociatingEnsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForAssociatingEnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForAssociatingEnsResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForAssociatingEnsResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAssociatingEnsResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForAssociatingEnsResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsResponse) SetStatusCode(v int32) *SaveSingleTaskForAssociatingEnsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsResponse) SetBody(v *SaveSingleTaskForAssociatingEnsResponseBody) *SaveSingleTaskForAssociatingEnsResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForCancelingTransferInRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForCancelingTransferInRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferInRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferInRequest) SetDomainName(v string) *SaveSingleTaskForCancelingTransferInRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInRequest) SetLang(v string) *SaveSingleTaskForCancelingTransferInRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInRequest) SetUserClientIp(v string) *SaveSingleTaskForCancelingTransferInRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForCancelingTransferInResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCancelingTransferInResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferInResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferInResponseBody) SetRequestId(v string) *SaveSingleTaskForCancelingTransferInResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInResponseBody) SetTaskNo(v string) *SaveSingleTaskForCancelingTransferInResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForCancelingTransferInResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForCancelingTransferInResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForCancelingTransferInResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferInResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferInResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCancelingTransferInResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInResponse) SetStatusCode(v int32) *SaveSingleTaskForCancelingTransferInResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInResponse) SetBody(v *SaveSingleTaskForCancelingTransferInResponseBody) *SaveSingleTaskForCancelingTransferInResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForCancelingTransferOutRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForCancelingTransferOutRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferOutRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) SetDomainName(v string) *SaveSingleTaskForCancelingTransferOutRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) SetLang(v string) *SaveSingleTaskForCancelingTransferOutRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutRequest) SetUserClientIp(v string) *SaveSingleTaskForCancelingTransferOutRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForCancelingTransferOutResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCancelingTransferOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferOutResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferOutResponseBody) SetRequestId(v string) *SaveSingleTaskForCancelingTransferOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutResponseBody) SetTaskNo(v string) *SaveSingleTaskForCancelingTransferOutResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForCancelingTransferOutResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForCancelingTransferOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForCancelingTransferOutResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferOutResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCancelingTransferOutResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) SetStatusCode(v int32) *SaveSingleTaskForCancelingTransferOutResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) SetBody(v *SaveSingleTaskForCancelingTransferOutResponseBody) *SaveSingleTaskForCancelingTransferOutResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForCreatingDnsHostRequest struct {
	DnsName      *string   `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip           []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForCreatingDnsHostRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingDnsHostRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetDnsName(v string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.DnsName = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetInstanceId(v string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetIp(v []*string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.Ip = v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetLang(v string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForCreatingDnsHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCreatingDnsHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingDnsHostResponseBody) SetRequestId(v string) *SaveSingleTaskForCreatingDnsHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostResponseBody) SetTaskNo(v string) *SaveSingleTaskForCreatingDnsHostResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForCreatingDnsHostResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForCreatingDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForCreatingDnsHostResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingDnsHostResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingDnsHostResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostResponse) SetBody(v *SaveSingleTaskForCreatingDnsHostResponseBody) *SaveSingleTaskForCreatingDnsHostResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForCreatingOrderActivateRequest struct {
	Address                   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AliyunDns                 *bool   `json:"AliyunDns,omitempty" xml:"AliyunDns,omitempty"`
	City                      *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                   *string `json:"Country,omitempty" xml:"Country,omitempty"`
	CouponNo                  *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	Dns1                      *string `json:"Dns1,omitempty" xml:"Dns1,omitempty"`
	Dns2                      *string `json:"Dns2,omitempty" xml:"Dns2,omitempty"`
	DomainName                *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Email                     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EnableDomainProxy         *bool   `json:"EnableDomainProxy,omitempty" xml:"EnableDomainProxy,omitempty"`
	Lang                      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PermitPremiumActivation   *bool   `json:"PermitPremiumActivation,omitempty" xml:"PermitPremiumActivation,omitempty"`
	PostalCode                *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	PromotionNo               *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	Province                  *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName            *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization    *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantProfileId       *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	RegistrantType            *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	SubscriptionDuration      *int32  `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
	TelArea                   *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                    *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                 *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	TrademarkDomainActivation *bool   `json:"TrademarkDomainActivation,omitempty" xml:"TrademarkDomainActivation,omitempty"`
	UseCoupon                 *bool   `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	UsePromotion              *bool   `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	UserClientIp              *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZhAddress                 *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                    *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince                *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName          *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization  *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderActivateRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderActivateRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetAddress(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Address = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetAliyunDns(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.AliyunDns = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetCity(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.City = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetCountry(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Country = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetCouponNo(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetDns1(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Dns1 = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetDns2(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Dns2 = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetDomainName(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetEmail(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Email = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetEnableDomainProxy(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.EnableDomainProxy = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetLang(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetPermitPremiumActivation(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.PermitPremiumActivation = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetPostalCode(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetProvince(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Province = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetRegistrantName(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetRegistrantOrganization(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetRegistrantProfileId(v int64) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetRegistrantType(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.RegistrantType = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetSubscriptionDuration(v int32) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.SubscriptionDuration = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetTelArea(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.TelArea = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetTelExt(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.TelExt = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetTelephone(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.Telephone = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetTrademarkDomainActivation(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.TrademarkDomainActivation = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhAddress(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhAddress = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhCity(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhCity = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhProvince(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhProvince = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhRegistrantName(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateRequest) SetZhRegistrantOrganization(v string) *SaveSingleTaskForCreatingOrderActivateRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

type SaveSingleTaskForCreatingOrderActivateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderActivateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderActivateResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderActivateResponseBody) SetRequestId(v string) *SaveSingleTaskForCreatingOrderActivateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateResponseBody) SetTaskNo(v string) *SaveSingleTaskForCreatingOrderActivateResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForCreatingOrderActivateResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForCreatingOrderActivateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForCreatingOrderActivateResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderActivateResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderActivateResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderActivateResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateResponse) SetBody(v *SaveSingleTaskForCreatingOrderActivateResponseBody) *SaveSingleTaskForCreatingOrderActivateResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForCreatingOrderRedeemRequest struct {
	CouponNo              *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	CurrentExpirationDate *int64  `json:"CurrentExpirationDate,omitempty" xml:"CurrentExpirationDate,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PromotionNo           *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	UseCoupon             *bool   `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	UsePromotion          *bool   `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	UserClientIp          *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderRedeemRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRedeemRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetCouponNo(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetCurrentExpirationDate(v int64) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.CurrentExpirationDate = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetDomainName(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetLang(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderRedeemRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForCreatingOrderRedeemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderRedeemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRedeemResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponseBody) SetRequestId(v string) *SaveSingleTaskForCreatingOrderRedeemResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponseBody) SetTaskNo(v string) *SaveSingleTaskForCreatingOrderRedeemResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForCreatingOrderRedeemResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForCreatingOrderRedeemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForCreatingOrderRedeemResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRedeemResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderRedeemResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderRedeemResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) SetBody(v *SaveSingleTaskForCreatingOrderRedeemResponseBody) *SaveSingleTaskForCreatingOrderRedeemResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForCreatingOrderRenewRequest struct {
	CouponNo              *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	CurrentExpirationDate *int64  `json:"CurrentExpirationDate,omitempty" xml:"CurrentExpirationDate,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PromotionNo           *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	SubscriptionDuration  *int32  `json:"SubscriptionDuration,omitempty" xml:"SubscriptionDuration,omitempty"`
	UseCoupon             *bool   `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	UsePromotion          *bool   `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	UserClientIp          *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderRenewRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRenewRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetCouponNo(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetCurrentExpirationDate(v int64) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.CurrentExpirationDate = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetDomainName(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetLang(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetSubscriptionDuration(v int32) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.SubscriptionDuration = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderRenewRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForCreatingOrderRenewResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderRenewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRenewResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRenewResponseBody) SetRequestId(v string) *SaveSingleTaskForCreatingOrderRenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewResponseBody) SetTaskNo(v string) *SaveSingleTaskForCreatingOrderRenewResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForCreatingOrderRenewResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForCreatingOrderRenewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForCreatingOrderRenewResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRenewResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderRenewResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderRenewResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRenewResponse) SetBody(v *SaveSingleTaskForCreatingOrderRenewResponseBody) *SaveSingleTaskForCreatingOrderRenewResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForCreatingOrderTransferRequest struct {
	AuthorizationCode     *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	CouponNo              *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PermitPremiumTransfer *bool   `json:"PermitPremiumTransfer,omitempty" xml:"PermitPremiumTransfer,omitempty"`
	PromotionNo           *string `json:"PromotionNo,omitempty" xml:"PromotionNo,omitempty"`
	RegistrantProfileId   *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	UseCoupon             *bool   `json:"UseCoupon,omitempty" xml:"UseCoupon,omitempty"`
	UsePromotion          *bool   `json:"UsePromotion,omitempty" xml:"UsePromotion,omitempty"`
	UserClientIp          *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderTransferRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetAuthorizationCode(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.AuthorizationCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetCouponNo(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.CouponNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetDomainName(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetLang(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetPermitPremiumTransfer(v bool) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.PermitPremiumTransfer = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetPromotionNo(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.PromotionNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetRegistrantProfileId(v int64) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetUseCoupon(v bool) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.UseCoupon = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetUsePromotion(v bool) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.UsePromotion = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingOrderTransferRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForCreatingOrderTransferResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderTransferResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderTransferResponseBody) SetRequestId(v string) *SaveSingleTaskForCreatingOrderTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferResponseBody) SetTaskNo(v string) *SaveSingleTaskForCreatingOrderTransferResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForCreatingOrderTransferResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForCreatingOrderTransferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForCreatingOrderTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderTransferResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderTransferResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) SetBody(v *SaveSingleTaskForCreatingOrderTransferResponseBody) *SaveSingleTaskForCreatingOrderTransferResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForDeletingDSRecordRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	KeyTag       *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForDeletingDSRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDeletingDSRecordRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) SetDomainName(v string) *SaveSingleTaskForDeletingDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) SetKeyTag(v int32) *SaveSingleTaskForDeletingDSRecordRequest {
	s.KeyTag = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) SetLang(v string) *SaveSingleTaskForDeletingDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordRequest) SetUserClientIp(v string) *SaveSingleTaskForDeletingDSRecordRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForDeletingDSRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForDeletingDSRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDeletingDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDSRecordResponseBody) SetRequestId(v string) *SaveSingleTaskForDeletingDSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordResponseBody) SetTaskNo(v string) *SaveSingleTaskForDeletingDSRecordResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForDeletingDSRecordResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForDeletingDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForDeletingDSRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDeletingDSRecordResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForDeletingDSRecordResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) SetStatusCode(v int32) *SaveSingleTaskForDeletingDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) SetBody(v *SaveSingleTaskForDeletingDSRecordResponseBody) *SaveSingleTaskForDeletingDSRecordResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForDeletingDnsHostRequest struct {
	DnsName      *string `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForDeletingDnsHostRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDeletingDnsHostRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) SetDnsName(v string) *SaveSingleTaskForDeletingDnsHostRequest {
	s.DnsName = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) SetInstanceId(v string) *SaveSingleTaskForDeletingDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) SetLang(v string) *SaveSingleTaskForDeletingDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) SetUserClientIp(v string) *SaveSingleTaskForDeletingDnsHostRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForDeletingDnsHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForDeletingDnsHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDeletingDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDnsHostResponseBody) SetRequestId(v string) *SaveSingleTaskForDeletingDnsHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostResponseBody) SetTaskNo(v string) *SaveSingleTaskForDeletingDnsHostResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForDeletingDnsHostResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForDeletingDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForDeletingDnsHostResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDeletingDnsHostResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForDeletingDnsHostResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) SetStatusCode(v int32) *SaveSingleTaskForDeletingDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostResponse) SetBody(v *SaveSingleTaskForDeletingDnsHostResponseBody) *SaveSingleTaskForDeletingDnsHostResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForDisassociatingEnsRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForDisassociatingEnsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDisassociatingEnsRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) SetDomainName(v string) *SaveSingleTaskForDisassociatingEnsRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) SetLang(v string) *SaveSingleTaskForDisassociatingEnsRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsRequest) SetUserClientIp(v string) *SaveSingleTaskForDisassociatingEnsRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForDisassociatingEnsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForDisassociatingEnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDisassociatingEnsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDisassociatingEnsResponseBody) SetRequestId(v string) *SaveSingleTaskForDisassociatingEnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsResponseBody) SetTaskNo(v string) *SaveSingleTaskForDisassociatingEnsResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForDisassociatingEnsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForDisassociatingEnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForDisassociatingEnsResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDisassociatingEnsResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForDisassociatingEnsResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) SetStatusCode(v int32) *SaveSingleTaskForDisassociatingEnsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsResponse) SetBody(v *SaveSingleTaskForDisassociatingEnsResponseBody) *SaveSingleTaskForDisassociatingEnsResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForDomainNameProxyServiceRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Status       *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForDomainNameProxyServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDomainNameProxyServiceRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) SetDomainName(v string) *SaveSingleTaskForDomainNameProxyServiceRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) SetLang(v string) *SaveSingleTaskForDomainNameProxyServiceRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) SetStatus(v bool) *SaveSingleTaskForDomainNameProxyServiceRequest {
	s.Status = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) SetUserClientIp(v string) *SaveSingleTaskForDomainNameProxyServiceRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForDomainNameProxyServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForDomainNameProxyServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDomainNameProxyServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponseBody) SetRequestId(v string) *SaveSingleTaskForDomainNameProxyServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponseBody) SetTaskNo(v string) *SaveSingleTaskForDomainNameProxyServiceResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForDomainNameProxyServiceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForDomainNameProxyServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForDomainNameProxyServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForDomainNameProxyServiceResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForDomainNameProxyServiceResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) SetStatusCode(v int32) *SaveSingleTaskForDomainNameProxyServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceResponse) SetBody(v *SaveSingleTaskForDomainNameProxyServiceResponseBody) *SaveSingleTaskForDomainNameProxyServiceResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForModifyingDSRecordRequest struct {
	Algorithm    *int32  `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Digest       *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	DigestType   *int32  `json:"DigestType,omitempty" xml:"DigestType,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	KeyTag       *int32  `json:"KeyTag,omitempty" xml:"KeyTag,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForModifyingDSRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForModifyingDSRecordRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetAlgorithm(v int32) *SaveSingleTaskForModifyingDSRecordRequest {
	s.Algorithm = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetDigest(v string) *SaveSingleTaskForModifyingDSRecordRequest {
	s.Digest = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetDigestType(v int32) *SaveSingleTaskForModifyingDSRecordRequest {
	s.DigestType = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetDomainName(v string) *SaveSingleTaskForModifyingDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetKeyTag(v int32) *SaveSingleTaskForModifyingDSRecordRequest {
	s.KeyTag = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetLang(v string) *SaveSingleTaskForModifyingDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordRequest) SetUserClientIp(v string) *SaveSingleTaskForModifyingDSRecordRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForModifyingDSRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForModifyingDSRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForModifyingDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDSRecordResponseBody) SetRequestId(v string) *SaveSingleTaskForModifyingDSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordResponseBody) SetTaskNo(v string) *SaveSingleTaskForModifyingDSRecordResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForModifyingDSRecordResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForModifyingDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForModifyingDSRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForModifyingDSRecordResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForModifyingDSRecordResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) SetStatusCode(v int32) *SaveSingleTaskForModifyingDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) SetBody(v *SaveSingleTaskForModifyingDSRecordResponseBody) *SaveSingleTaskForModifyingDSRecordResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForModifyingDnsHostRequest struct {
	DnsName      *string   `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip           []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForModifyingDnsHostRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForModifyingDnsHostRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetDnsName(v string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.DnsName = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetInstanceId(v string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetIp(v []*string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.Ip = v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetLang(v string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetUserClientIp(v string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForModifyingDnsHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForModifyingDnsHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForModifyingDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDnsHostResponseBody) SetRequestId(v string) *SaveSingleTaskForModifyingDnsHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostResponseBody) SetTaskNo(v string) *SaveSingleTaskForModifyingDnsHostResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForModifyingDnsHostResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForModifyingDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForModifyingDnsHostResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForModifyingDnsHostResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForModifyingDnsHostResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) SetStatusCode(v int32) *SaveSingleTaskForModifyingDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostResponse) SetBody(v *SaveSingleTaskForModifyingDnsHostResponseBody) *SaveSingleTaskForModifyingDnsHostResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForQueryingTransferAuthorizationCodeRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) SetDomainName(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) SetLang(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) SetUserClientIp(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) SetRequestId(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) SetTaskNo(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForQueryingTransferAuthorizationCodeResponse struct {
	Headers    map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) SetStatusCode(v int32) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) SetBody(v *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForSaveArtExtensionRequest struct {
	DateOrPeriod            *string `json:"DateOrPeriod,omitempty" xml:"DateOrPeriod,omitempty"`
	Dimensions              *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	DomainName              *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Features                *string `json:"Features,omitempty" xml:"Features,omitempty"`
	InscriptionsAndMarkings *string `json:"InscriptionsAndMarkings,omitempty" xml:"InscriptionsAndMarkings,omitempty"`
	Lang                    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Maker                   *string `json:"Maker,omitempty" xml:"Maker,omitempty"`
	MaterialsAndTechniques  *string `json:"MaterialsAndTechniques,omitempty" xml:"MaterialsAndTechniques,omitempty"`
	ObjectType              *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Reference               *string `json:"Reference,omitempty" xml:"Reference,omitempty"`
	Subject                 *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Title                   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SaveSingleTaskForSaveArtExtensionRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForSaveArtExtensionRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetDateOrPeriod(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.DateOrPeriod = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetDimensions(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Dimensions = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetDomainName(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetFeatures(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Features = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetInscriptionsAndMarkings(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.InscriptionsAndMarkings = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetLang(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetMaker(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Maker = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetMaterialsAndTechniques(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.MaterialsAndTechniques = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetObjectType(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.ObjectType = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetReference(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Reference = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetSubject(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Subject = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionRequest) SetTitle(v string) *SaveSingleTaskForSaveArtExtensionRequest {
	s.Title = &v
	return s
}

type SaveSingleTaskForSaveArtExtensionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForSaveArtExtensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForSaveArtExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSaveArtExtensionResponseBody) SetRequestId(v string) *SaveSingleTaskForSaveArtExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionResponseBody) SetTaskNo(v string) *SaveSingleTaskForSaveArtExtensionResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForSaveArtExtensionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForSaveArtExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForSaveArtExtensionResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForSaveArtExtensionResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForSaveArtExtensionResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) SetStatusCode(v int32) *SaveSingleTaskForSaveArtExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) SetBody(v *SaveSingleTaskForSaveArtExtensionResponseBody) *SaveSingleTaskForSaveArtExtensionResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForSynchronizingDSRecordRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDSRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDSRecordRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) SetDomainName(v string) *SaveSingleTaskForSynchronizingDSRecordRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) SetLang(v string) *SaveSingleTaskForSynchronizingDSRecordRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordRequest) SetUserClientIp(v string) *SaveSingleTaskForSynchronizingDSRecordRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForSynchronizingDSRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDSRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponseBody) SetRequestId(v string) *SaveSingleTaskForSynchronizingDSRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponseBody) SetTaskNo(v string) *SaveSingleTaskForSynchronizingDSRecordResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForSynchronizingDSRecordResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForSynchronizingDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForSynchronizingDSRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDSRecordResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForSynchronizingDSRecordResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) SetStatusCode(v int32) *SaveSingleTaskForSynchronizingDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) SetBody(v *SaveSingleTaskForSynchronizingDSRecordResponseBody) *SaveSingleTaskForSynchronizingDSRecordResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForSynchronizingDnsHostRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDnsHostRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDnsHostRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) SetInstanceId(v string) *SaveSingleTaskForSynchronizingDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) SetLang(v string) *SaveSingleTaskForSynchronizingDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostRequest) SetUserClientIp(v string) *SaveSingleTaskForSynchronizingDnsHostRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForSynchronizingDnsHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDnsHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponseBody) SetRequestId(v string) *SaveSingleTaskForSynchronizingDnsHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponseBody) SetTaskNo(v string) *SaveSingleTaskForSynchronizingDnsHostResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForSynchronizingDnsHostResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForSynchronizingDnsHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForSynchronizingDnsHostResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDnsHostResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForSynchronizingDnsHostResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) SetStatusCode(v int32) *SaveSingleTaskForSynchronizingDnsHostResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponse) SetBody(v *SaveSingleTaskForSynchronizingDnsHostResponseBody) *SaveSingleTaskForSynchronizingDnsHostResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForTransferProhibitionLockRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Status       *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForTransferProhibitionLockRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForTransferProhibitionLockRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) SetDomainName(v string) *SaveSingleTaskForTransferProhibitionLockRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) SetLang(v string) *SaveSingleTaskForTransferProhibitionLockRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) SetStatus(v bool) *SaveSingleTaskForTransferProhibitionLockRequest {
	s.Status = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockRequest) SetUserClientIp(v string) *SaveSingleTaskForTransferProhibitionLockRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForTransferProhibitionLockResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForTransferProhibitionLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForTransferProhibitionLockResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForTransferProhibitionLockResponseBody) SetRequestId(v string) *SaveSingleTaskForTransferProhibitionLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockResponseBody) SetTaskNo(v string) *SaveSingleTaskForTransferProhibitionLockResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForTransferProhibitionLockResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForTransferProhibitionLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForTransferProhibitionLockResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForTransferProhibitionLockResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForTransferProhibitionLockResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) SetStatusCode(v int32) *SaveSingleTaskForTransferProhibitionLockResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForTransferProhibitionLockResponse) SetBody(v *SaveSingleTaskForTransferProhibitionLockResponseBody) *SaveSingleTaskForTransferProhibitionLockResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForUpdateProhibitionLockRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Status       *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForUpdateProhibitionLockRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForUpdateProhibitionLockRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) SetDomainName(v string) *SaveSingleTaskForUpdateProhibitionLockRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) SetLang(v string) *SaveSingleTaskForUpdateProhibitionLockRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) SetStatus(v bool) *SaveSingleTaskForUpdateProhibitionLockRequest {
	s.Status = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockRequest) SetUserClientIp(v string) *SaveSingleTaskForUpdateProhibitionLockRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForUpdateProhibitionLockResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForUpdateProhibitionLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForUpdateProhibitionLockResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponseBody) SetRequestId(v string) *SaveSingleTaskForUpdateProhibitionLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponseBody) SetTaskNo(v string) *SaveSingleTaskForUpdateProhibitionLockResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForUpdateProhibitionLockResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForUpdateProhibitionLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForUpdateProhibitionLockResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForUpdateProhibitionLockResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForUpdateProhibitionLockResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) SetStatusCode(v int32) *SaveSingleTaskForUpdateProhibitionLockResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForUpdateProhibitionLockResponse) SetBody(v *SaveSingleTaskForUpdateProhibitionLockResponseBody) *SaveSingleTaskForUpdateProhibitionLockResponse {
	s.Body = v
	return s
}

type SaveSingleTaskForUpdatingContactInfoRequest struct {
	AddTransferLock     *bool   `json:"AddTransferLock,omitempty" xml:"AddTransferLock,omitempty"`
	ContactType         *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegistrantProfileId *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	UserClientIp        *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForUpdatingContactInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForUpdatingContactInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetAddTransferLock(v bool) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.AddTransferLock = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetContactType(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.ContactType = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetDomainName(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetInstanceId(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetLang(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetRegistrantProfileId(v int64) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoRequest) SetUserClientIp(v string) *SaveSingleTaskForUpdatingContactInfoRequest {
	s.UserClientIp = &v
	return s
}

type SaveSingleTaskForUpdatingContactInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForUpdatingContactInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForUpdatingContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdatingContactInfoResponseBody) SetRequestId(v string) *SaveSingleTaskForUpdatingContactInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoResponseBody) SetTaskNo(v string) *SaveSingleTaskForUpdatingContactInfoResponseBody {
	s.TaskNo = &v
	return s
}

type SaveSingleTaskForUpdatingContactInfoResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSingleTaskForUpdatingContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSingleTaskForUpdatingContactInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSingleTaskForUpdatingContactInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForUpdatingContactInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) SetStatusCode(v int32) *SaveSingleTaskForUpdatingContactInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoResponse) SetBody(v *SaveSingleTaskForUpdatingContactInfoResponseBody) *SaveSingleTaskForUpdatingContactInfoResponse {
	s.Body = v
	return s
}

type SaveTaskForSubmittingDomainDeleteRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForSubmittingDomainDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForSubmittingDomainDeleteRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) SetInstanceId(v string) *SaveTaskForSubmittingDomainDeleteRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) SetLang(v string) *SaveTaskForSubmittingDomainDeleteRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) SetUserClientIp(v string) *SaveTaskForSubmittingDomainDeleteRequest {
	s.UserClientIp = &v
	return s
}

type SaveTaskForSubmittingDomainDeleteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForSubmittingDomainDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForSubmittingDomainDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainDeleteResponseBody) SetRequestId(v string) *SaveTaskForSubmittingDomainDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteResponseBody) SetTaskNo(v string) *SaveTaskForSubmittingDomainDeleteResponseBody {
	s.TaskNo = &v
	return s
}

type SaveTaskForSubmittingDomainDeleteResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveTaskForSubmittingDomainDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTaskForSubmittingDomainDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForSubmittingDomainDeleteResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainDeleteResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) SetStatusCode(v int32) *SaveTaskForSubmittingDomainDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteResponse) SetBody(v *SaveTaskForSubmittingDomainDeleteResponseBody) *SaveTaskForSubmittingDomainDeleteResponse {
	s.Body = v
	return s
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest struct {
	DomainName             []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	IdentityCredential     *string   `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	IdentityCredentialNo   *string   `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	IdentityCredentialType *string   `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	Lang                   *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp           *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetDomainName(v []*string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.DomainName = v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetIdentityCredential(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.IdentityCredential = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetIdentityCredentialNo(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.IdentityCredentialNo = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetIdentityCredentialType(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.IdentityCredentialType = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetLang(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) SetUserClientIp(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest {
	s.UserClientIp = &v
	return s
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) SetRequestId(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) SetTaskNo(v string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody {
	s.TaskNo = &v
	return s
}

type SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse struct {
	Headers    map[string]*string                                                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) SetStatusCode(v int32) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse) SetBody(v *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponseBody) *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse {
	s.Body = v
	return s
}

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest struct {
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegistrantProfileId *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	UserClientIp        *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetDomainName(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.DomainName = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetInstanceId(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetLang(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetRegistrantProfileId(v int64) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) SetUserClientIp(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest {
	s.UserClientIp = &v
	return s
}

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) SetRequestId(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) SetTaskNo(v string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody {
	s.TaskNo = &v
	return s
}

type SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse struct {
	Headers    map[string]*string                                                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) SetHeaders(v map[string]*string) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) SetStatusCode(v int32) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse) SetBody(v *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponseBody) *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse {
	s.Body = v
	return s
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest struct {
	Address                  *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	City                     *string   `json:"City,omitempty" xml:"City,omitempty"`
	Country                  *string   `json:"Country,omitempty" xml:"Country,omitempty"`
	DomainName               []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Email                    *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	IdentityCredential       *string   `json:"IdentityCredential,omitempty" xml:"IdentityCredential,omitempty"`
	IdentityCredentialNo     *string   `json:"IdentityCredentialNo,omitempty" xml:"IdentityCredentialNo,omitempty"`
	IdentityCredentialType   *string   `json:"IdentityCredentialType,omitempty" xml:"IdentityCredentialType,omitempty"`
	Lang                     *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PostalCode               *string   `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                 *string   `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName           *string   `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization   *string   `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantType           *string   `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	TelArea                  *string   `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                   *string   `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                *string   `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	TransferOutProhibited    *bool     `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	UserClientIp             *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZhAddress                *string   `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string   `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string   `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string   `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string   `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetAddress(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Address = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetCity(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.City = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetCountry(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Country = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetDomainName(v []*string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.DomainName = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetEmail(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Email = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetIdentityCredential(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.IdentityCredential = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetIdentityCredentialNo(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.IdentityCredentialNo = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetIdentityCredentialType(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.IdentityCredentialType = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetLang(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetPostalCode(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetProvince(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Province = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetRegistrantName(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.RegistrantName = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetRegistrantOrganization(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetRegistrantType(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.RegistrantType = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetTelArea(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.TelArea = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetTelExt(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.TelExt = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetTelephone(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.Telephone = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetTransferOutProhibited(v bool) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.TransferOutProhibited = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetUserClientIp(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhAddress(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhAddress = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhCity(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhCity = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhProvince(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhProvince = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhRegistrantName(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) SetZhRegistrantOrganization(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) SetRequestId(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) SetTaskNo(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody {
	s.TaskNo = &v
	return s
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) SetHeaders(v map[string]*string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) SetStatusCode(v int32) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse) SetBody(v *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse {
	s.Body = v
	return s
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest struct {
	DomainName            []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	Lang                  *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegistrantProfileId   *int64    `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	TransferOutProhibited *bool     `json:"TransferOutProhibited,omitempty" xml:"TransferOutProhibited,omitempty"`
	UserClientIp          *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetDomainName(v []*string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.DomainName = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetLang(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetRegistrantProfileId(v int64) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetTransferOutProhibited(v bool) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.TransferOutProhibited = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) SetUserClientIp(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest {
	s.UserClientIp = &v
	return s
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) SetRequestId(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) SetTaskNo(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody {
	s.TaskNo = &v
	return s
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse struct {
	Headers    map[string]*string                                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) SetHeaders(v map[string]*string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) SetStatusCode(v int32) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) SetBody(v *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse {
	s.Body = v
	return s
}

type ScrollDomainListRequest struct {
	DomainGroupId         *int64  `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainStatus          *int32  `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	EndExpirationDate     *int64  `json:"EndExpirationDate,omitempty" xml:"EndExpirationDate,omitempty"`
	EndLength             *int32  `json:"EndLength,omitempty" xml:"EndLength,omitempty"`
	EndRegistrationDate   *int64  `json:"EndRegistrationDate,omitempty" xml:"EndRegistrationDate,omitempty"`
	Excluded              *string `json:"Excluded,omitempty" xml:"Excluded,omitempty"`
	ExcludedPrefix        *bool   `json:"ExcludedPrefix,omitempty" xml:"ExcludedPrefix,omitempty"`
	ExcludedSuffix        *bool   `json:"ExcludedSuffix,omitempty" xml:"ExcludedSuffix,omitempty"`
	Form                  *int32  `json:"Form,omitempty" xml:"Form,omitempty"`
	KeyWord               *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	KeyWordPrefix         *bool   `json:"KeyWordPrefix,omitempty" xml:"KeyWordPrefix,omitempty"`
	KeyWordSuffix         *bool   `json:"KeyWordSuffix,omitempty" xml:"KeyWordSuffix,omitempty"`
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductDomainType     *string `json:"ProductDomainType,omitempty" xml:"ProductDomainType,omitempty"`
	ScrollId              *string `json:"ScrollId,omitempty" xml:"ScrollId,omitempty"`
	StartExpirationDate   *int64  `json:"StartExpirationDate,omitempty" xml:"StartExpirationDate,omitempty"`
	StartLength           *int32  `json:"StartLength,omitempty" xml:"StartLength,omitempty"`
	StartRegistrationDate *int64  `json:"StartRegistrationDate,omitempty" xml:"StartRegistrationDate,omitempty"`
	Suffixs               *string `json:"Suffixs,omitempty" xml:"Suffixs,omitempty"`
	TradeType             *int32  `json:"TradeType,omitempty" xml:"TradeType,omitempty"`
	UserClientIp          *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ScrollDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s ScrollDomainListRequest) GoString() string {
	return s.String()
}

func (s *ScrollDomainListRequest) SetDomainGroupId(v int64) *ScrollDomainListRequest {
	s.DomainGroupId = &v
	return s
}

func (s *ScrollDomainListRequest) SetDomainStatus(v int32) *ScrollDomainListRequest {
	s.DomainStatus = &v
	return s
}

func (s *ScrollDomainListRequest) SetEndExpirationDate(v int64) *ScrollDomainListRequest {
	s.EndExpirationDate = &v
	return s
}

func (s *ScrollDomainListRequest) SetEndLength(v int32) *ScrollDomainListRequest {
	s.EndLength = &v
	return s
}

func (s *ScrollDomainListRequest) SetEndRegistrationDate(v int64) *ScrollDomainListRequest {
	s.EndRegistrationDate = &v
	return s
}

func (s *ScrollDomainListRequest) SetExcluded(v string) *ScrollDomainListRequest {
	s.Excluded = &v
	return s
}

func (s *ScrollDomainListRequest) SetExcludedPrefix(v bool) *ScrollDomainListRequest {
	s.ExcludedPrefix = &v
	return s
}

func (s *ScrollDomainListRequest) SetExcludedSuffix(v bool) *ScrollDomainListRequest {
	s.ExcludedSuffix = &v
	return s
}

func (s *ScrollDomainListRequest) SetForm(v int32) *ScrollDomainListRequest {
	s.Form = &v
	return s
}

func (s *ScrollDomainListRequest) SetKeyWord(v string) *ScrollDomainListRequest {
	s.KeyWord = &v
	return s
}

func (s *ScrollDomainListRequest) SetKeyWordPrefix(v bool) *ScrollDomainListRequest {
	s.KeyWordPrefix = &v
	return s
}

func (s *ScrollDomainListRequest) SetKeyWordSuffix(v bool) *ScrollDomainListRequest {
	s.KeyWordSuffix = &v
	return s
}

func (s *ScrollDomainListRequest) SetLang(v string) *ScrollDomainListRequest {
	s.Lang = &v
	return s
}

func (s *ScrollDomainListRequest) SetPageSize(v int32) *ScrollDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *ScrollDomainListRequest) SetProductDomainType(v string) *ScrollDomainListRequest {
	s.ProductDomainType = &v
	return s
}

func (s *ScrollDomainListRequest) SetScrollId(v string) *ScrollDomainListRequest {
	s.ScrollId = &v
	return s
}

func (s *ScrollDomainListRequest) SetStartExpirationDate(v int64) *ScrollDomainListRequest {
	s.StartExpirationDate = &v
	return s
}

func (s *ScrollDomainListRequest) SetStartLength(v int32) *ScrollDomainListRequest {
	s.StartLength = &v
	return s
}

func (s *ScrollDomainListRequest) SetStartRegistrationDate(v int64) *ScrollDomainListRequest {
	s.StartRegistrationDate = &v
	return s
}

func (s *ScrollDomainListRequest) SetSuffixs(v string) *ScrollDomainListRequest {
	s.Suffixs = &v
	return s
}

func (s *ScrollDomainListRequest) SetTradeType(v int32) *ScrollDomainListRequest {
	s.TradeType = &v
	return s
}

func (s *ScrollDomainListRequest) SetUserClientIp(v string) *ScrollDomainListRequest {
	s.UserClientIp = &v
	return s
}

type ScrollDomainListResponseBody struct {
	Data         *ScrollDomainListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	PageSize     *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScrollId     *string                           `json:"ScrollId,omitempty" xml:"ScrollId,omitempty"`
	TotalItemNum *int32                            `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
}

func (s ScrollDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScrollDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBody) SetData(v *ScrollDomainListResponseBodyData) *ScrollDomainListResponseBody {
	s.Data = v
	return s
}

func (s *ScrollDomainListResponseBody) SetPageSize(v int32) *ScrollDomainListResponseBody {
	s.PageSize = &v
	return s
}

func (s *ScrollDomainListResponseBody) SetRequestId(v string) *ScrollDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScrollDomainListResponseBody) SetScrollId(v string) *ScrollDomainListResponseBody {
	s.ScrollId = &v
	return s
}

func (s *ScrollDomainListResponseBody) SetTotalItemNum(v int32) *ScrollDomainListResponseBody {
	s.TotalItemNum = &v
	return s
}

type ScrollDomainListResponseBodyData struct {
	Domain []*ScrollDomainListResponseBodyDataDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s ScrollDomainListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ScrollDomainListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBodyData) SetDomain(v []*ScrollDomainListResponseBodyDataDomain) *ScrollDomainListResponseBodyData {
	s.Domain = v
	return s
}

type ScrollDomainListResponseBodyDataDomain struct {
	DnsList                  *ScrollDomainListResponseBodyDataDomainDnsList `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Struct"`
	DomainAuditStatus        *string                                        `json:"DomainAuditStatus,omitempty" xml:"DomainAuditStatus,omitempty"`
	DomainGroupId            *string                                        `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName          *string                                        `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	DomainName               *string                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus             *string                                        `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainType               *string                                        `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	Email                    *string                                        `json:"Email,omitempty" xml:"Email,omitempty"`
	ExpirationCurrDateDiff   *int32                                         `json:"ExpirationCurrDateDiff,omitempty" xml:"ExpirationCurrDateDiff,omitempty"`
	ExpirationDate           *string                                        `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	ExpirationDateLong       *int64                                         `json:"ExpirationDateLong,omitempty" xml:"ExpirationDateLong,omitempty"`
	ExpirationDateStatus     *string                                        `json:"ExpirationDateStatus,omitempty" xml:"ExpirationDateStatus,omitempty"`
	InstanceId               *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Premium                  *bool                                          `json:"Premium,omitempty" xml:"Premium,omitempty"`
	ProductId                *string                                        `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RegistrantOrganization   *string                                        `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantType           *string                                        `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	RegistrationDate         *string                                        `json:"RegistrationDate,omitempty" xml:"RegistrationDate,omitempty"`
	RegistrationDateLong     *int64                                         `json:"RegistrationDateLong,omitempty" xml:"RegistrationDateLong,omitempty"`
	Remark                   *string                                        `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ZhRegistrantOrganization *string                                        `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s ScrollDomainListResponseBodyDataDomain) String() string {
	return tea.Prettify(s)
}

func (s ScrollDomainListResponseBodyDataDomain) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDnsList(v *ScrollDomainListResponseBodyDataDomainDnsList) *ScrollDomainListResponseBodyDataDomain {
	s.DnsList = v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainAuditStatus(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainAuditStatus = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainGroupId(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainGroupId = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainGroupName(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainGroupName = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainName(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainName = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainStatus(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainStatus = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetDomainType(v string) *ScrollDomainListResponseBodyDataDomain {
	s.DomainType = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetEmail(v string) *ScrollDomainListResponseBodyDataDomain {
	s.Email = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetExpirationCurrDateDiff(v int32) *ScrollDomainListResponseBodyDataDomain {
	s.ExpirationCurrDateDiff = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetExpirationDate(v string) *ScrollDomainListResponseBodyDataDomain {
	s.ExpirationDate = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetExpirationDateLong(v int64) *ScrollDomainListResponseBodyDataDomain {
	s.ExpirationDateLong = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetExpirationDateStatus(v string) *ScrollDomainListResponseBodyDataDomain {
	s.ExpirationDateStatus = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetInstanceId(v string) *ScrollDomainListResponseBodyDataDomain {
	s.InstanceId = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetPremium(v bool) *ScrollDomainListResponseBodyDataDomain {
	s.Premium = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetProductId(v string) *ScrollDomainListResponseBodyDataDomain {
	s.ProductId = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRegistrantOrganization(v string) *ScrollDomainListResponseBodyDataDomain {
	s.RegistrantOrganization = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRegistrantType(v string) *ScrollDomainListResponseBodyDataDomain {
	s.RegistrantType = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRegistrationDate(v string) *ScrollDomainListResponseBodyDataDomain {
	s.RegistrationDate = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRegistrationDateLong(v int64) *ScrollDomainListResponseBodyDataDomain {
	s.RegistrationDateLong = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetRemark(v string) *ScrollDomainListResponseBodyDataDomain {
	s.Remark = &v
	return s
}

func (s *ScrollDomainListResponseBodyDataDomain) SetZhRegistrantOrganization(v string) *ScrollDomainListResponseBodyDataDomain {
	s.ZhRegistrantOrganization = &v
	return s
}

type ScrollDomainListResponseBodyDataDomainDnsList struct {
	Dns []*string `json:"Dns,omitempty" xml:"Dns,omitempty" type:"Repeated"`
}

func (s ScrollDomainListResponseBodyDataDomainDnsList) String() string {
	return tea.Prettify(s)
}

func (s ScrollDomainListResponseBodyDataDomainDnsList) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponseBodyDataDomainDnsList) SetDns(v []*string) *ScrollDomainListResponseBodyDataDomainDnsList {
	s.Dns = v
	return s
}

type ScrollDomainListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScrollDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScrollDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s ScrollDomainListResponse) GoString() string {
	return s.String()
}

func (s *ScrollDomainListResponse) SetHeaders(v map[string]*string) *ScrollDomainListResponse {
	s.Headers = v
	return s
}

func (s *ScrollDomainListResponse) SetStatusCode(v int32) *ScrollDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *ScrollDomainListResponse) SetBody(v *ScrollDomainListResponseBody) *ScrollDomainListResponse {
	s.Body = v
	return s
}

type SetDefaultRegistrantProfileRequest struct {
	RegistrantProfileId *int64  `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	UserClientIp        *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SetDefaultRegistrantProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultRegistrantProfileRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultRegistrantProfileRequest) SetRegistrantProfileId(v int64) *SetDefaultRegistrantProfileRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SetDefaultRegistrantProfileRequest) SetUserClientIp(v string) *SetDefaultRegistrantProfileRequest {
	s.UserClientIp = &v
	return s
}

type SetDefaultRegistrantProfileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultRegistrantProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultRegistrantProfileResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultRegistrantProfileResponseBody) SetRequestId(v string) *SetDefaultRegistrantProfileResponseBody {
	s.RequestId = &v
	return s
}

type SetDefaultRegistrantProfileResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDefaultRegistrantProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDefaultRegistrantProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultRegistrantProfileResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultRegistrantProfileResponse) SetHeaders(v map[string]*string) *SetDefaultRegistrantProfileResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultRegistrantProfileResponse) SetStatusCode(v int32) *SetDefaultRegistrantProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultRegistrantProfileResponse) SetBody(v *SetDefaultRegistrantProfileResponseBody) *SetDefaultRegistrantProfileResponse {
	s.Body = v
	return s
}

type SubmitEmailVerificationRequest struct {
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SendIfExist  *bool   `json:"SendIfExist,omitempty" xml:"SendIfExist,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SubmitEmailVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationRequest) SetEmail(v string) *SubmitEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *SubmitEmailVerificationRequest) SetLang(v string) *SubmitEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *SubmitEmailVerificationRequest) SetSendIfExist(v bool) *SubmitEmailVerificationRequest {
	s.SendIfExist = &v
	return s
}

func (s *SubmitEmailVerificationRequest) SetUserClientIp(v string) *SubmitEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

type SubmitEmailVerificationResponseBody struct {
	ExistList   []*SubmitEmailVerificationResponseBodyExistList   `json:"ExistList,omitempty" xml:"ExistList,omitempty" type:"Repeated"`
	FailList    []*SubmitEmailVerificationResponseBodyFailList    `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Repeated"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList []*SubmitEmailVerificationResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Repeated"`
}

func (s SubmitEmailVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponseBody) SetExistList(v []*SubmitEmailVerificationResponseBodyExistList) *SubmitEmailVerificationResponseBody {
	s.ExistList = v
	return s
}

func (s *SubmitEmailVerificationResponseBody) SetFailList(v []*SubmitEmailVerificationResponseBodyFailList) *SubmitEmailVerificationResponseBody {
	s.FailList = v
	return s
}

func (s *SubmitEmailVerificationResponseBody) SetRequestId(v string) *SubmitEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitEmailVerificationResponseBody) SetSuccessList(v []*SubmitEmailVerificationResponseBodySuccessList) *SubmitEmailVerificationResponseBody {
	s.SuccessList = v
	return s
}

type SubmitEmailVerificationResponseBodyExistList struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SubmitEmailVerificationResponseBodyExistList) String() string {
	return tea.Prettify(s)
}

func (s SubmitEmailVerificationResponseBodyExistList) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponseBodyExistList) SetCode(v string) *SubmitEmailVerificationResponseBodyExistList {
	s.Code = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyExistList) SetEmail(v string) *SubmitEmailVerificationResponseBodyExistList {
	s.Email = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyExistList) SetMessage(v string) *SubmitEmailVerificationResponseBodyExistList {
	s.Message = &v
	return s
}

type SubmitEmailVerificationResponseBodyFailList struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SubmitEmailVerificationResponseBodyFailList) String() string {
	return tea.Prettify(s)
}

func (s SubmitEmailVerificationResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponseBodyFailList) SetCode(v string) *SubmitEmailVerificationResponseBodyFailList {
	s.Code = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyFailList) SetEmail(v string) *SubmitEmailVerificationResponseBodyFailList {
	s.Email = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodyFailList) SetMessage(v string) *SubmitEmailVerificationResponseBodyFailList {
	s.Message = &v
	return s
}

type SubmitEmailVerificationResponseBodySuccessList struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SubmitEmailVerificationResponseBodySuccessList) String() string {
	return tea.Prettify(s)
}

func (s SubmitEmailVerificationResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponseBodySuccessList) SetCode(v string) *SubmitEmailVerificationResponseBodySuccessList {
	s.Code = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodySuccessList) SetEmail(v string) *SubmitEmailVerificationResponseBodySuccessList {
	s.Email = &v
	return s
}

func (s *SubmitEmailVerificationResponseBodySuccessList) SetMessage(v string) *SubmitEmailVerificationResponseBodySuccessList {
	s.Message = &v
	return s
}

type SubmitEmailVerificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitEmailVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponse) SetHeaders(v map[string]*string) *SubmitEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *SubmitEmailVerificationResponse) SetStatusCode(v int32) *SubmitEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitEmailVerificationResponse) SetBody(v *SubmitEmailVerificationResponseBody) *SubmitEmailVerificationResponse {
	s.Body = v
	return s
}

type SubmitOperationAuditInfoRequest struct {
	AuditInfo  *string `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty"`
	AuditType  *int32  `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s SubmitOperationAuditInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitOperationAuditInfoRequest) GoString() string {
	return s.String()
}

func (s *SubmitOperationAuditInfoRequest) SetAuditInfo(v string) *SubmitOperationAuditInfoRequest {
	s.AuditInfo = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetAuditType(v int32) *SubmitOperationAuditInfoRequest {
	s.AuditType = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetDomainName(v string) *SubmitOperationAuditInfoRequest {
	s.DomainName = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetId(v int64) *SubmitOperationAuditInfoRequest {
	s.Id = &v
	return s
}

func (s *SubmitOperationAuditInfoRequest) SetLang(v string) *SubmitOperationAuditInfoRequest {
	s.Lang = &v
	return s
}

type SubmitOperationAuditInfoResponseBody struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitOperationAuditInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitOperationAuditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitOperationAuditInfoResponseBody) SetId(v int64) *SubmitOperationAuditInfoResponseBody {
	s.Id = &v
	return s
}

func (s *SubmitOperationAuditInfoResponseBody) SetRequestId(v string) *SubmitOperationAuditInfoResponseBody {
	s.RequestId = &v
	return s
}

type SubmitOperationAuditInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitOperationAuditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitOperationAuditInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitOperationAuditInfoResponse) GoString() string {
	return s.String()
}

func (s *SubmitOperationAuditInfoResponse) SetHeaders(v map[string]*string) *SubmitOperationAuditInfoResponse {
	s.Headers = v
	return s
}

func (s *SubmitOperationAuditInfoResponse) SetStatusCode(v int32) *SubmitOperationAuditInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitOperationAuditInfoResponse) SetBody(v *SubmitOperationAuditInfoResponseBody) *SubmitOperationAuditInfoResponse {
	s.Body = v
	return s
}

type SubmitOperationCredentialsRequest struct {
	AuditRecordId *int64  `json:"AuditRecordId,omitempty" xml:"AuditRecordId,omitempty"`
	AuditType     *int32  `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	Credentials   *string `json:"Credentials,omitempty" xml:"Credentials,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegType       *int32  `json:"RegType,omitempty" xml:"RegType,omitempty"`
}

func (s SubmitOperationCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitOperationCredentialsRequest) GoString() string {
	return s.String()
}

func (s *SubmitOperationCredentialsRequest) SetAuditRecordId(v int64) *SubmitOperationCredentialsRequest {
	s.AuditRecordId = &v
	return s
}

func (s *SubmitOperationCredentialsRequest) SetAuditType(v int32) *SubmitOperationCredentialsRequest {
	s.AuditType = &v
	return s
}

func (s *SubmitOperationCredentialsRequest) SetCredentials(v string) *SubmitOperationCredentialsRequest {
	s.Credentials = &v
	return s
}

func (s *SubmitOperationCredentialsRequest) SetLang(v string) *SubmitOperationCredentialsRequest {
	s.Lang = &v
	return s
}

func (s *SubmitOperationCredentialsRequest) SetRegType(v int32) *SubmitOperationCredentialsRequest {
	s.RegType = &v
	return s
}

type SubmitOperationCredentialsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitOperationCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitOperationCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitOperationCredentialsResponseBody) SetRequestId(v string) *SubmitOperationCredentialsResponseBody {
	s.RequestId = &v
	return s
}

type SubmitOperationCredentialsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitOperationCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitOperationCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitOperationCredentialsResponse) GoString() string {
	return s.String()
}

func (s *SubmitOperationCredentialsResponse) SetHeaders(v map[string]*string) *SubmitOperationCredentialsResponse {
	s.Headers = v
	return s
}

func (s *SubmitOperationCredentialsResponse) SetStatusCode(v int32) *SubmitOperationCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitOperationCredentialsResponse) SetBody(v *SubmitOperationCredentialsResponseBody) *SubmitOperationCredentialsResponse {
	s.Body = v
	return s
}

type TransferInCheckMailTokenRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Token        *string `json:"Token,omitempty" xml:"Token,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s TransferInCheckMailTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferInCheckMailTokenRequest) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenRequest) SetLang(v string) *TransferInCheckMailTokenRequest {
	s.Lang = &v
	return s
}

func (s *TransferInCheckMailTokenRequest) SetToken(v string) *TransferInCheckMailTokenRequest {
	s.Token = &v
	return s
}

func (s *TransferInCheckMailTokenRequest) SetUserClientIp(v string) *TransferInCheckMailTokenRequest {
	s.UserClientIp = &v
	return s
}

type TransferInCheckMailTokenResponseBody struct {
	FailList    *TransferInCheckMailTokenResponseBodyFailList    `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Struct"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList *TransferInCheckMailTokenResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Struct"`
}

func (s TransferInCheckMailTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferInCheckMailTokenResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenResponseBody) SetFailList(v *TransferInCheckMailTokenResponseBodyFailList) *TransferInCheckMailTokenResponseBody {
	s.FailList = v
	return s
}

func (s *TransferInCheckMailTokenResponseBody) SetRequestId(v string) *TransferInCheckMailTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferInCheckMailTokenResponseBody) SetSuccessList(v *TransferInCheckMailTokenResponseBodySuccessList) *TransferInCheckMailTokenResponseBody {
	s.SuccessList = v
	return s
}

type TransferInCheckMailTokenResponseBodyFailList struct {
	FailDomain []*string `json:"FailDomain,omitempty" xml:"FailDomain,omitempty" type:"Repeated"`
}

func (s TransferInCheckMailTokenResponseBodyFailList) String() string {
	return tea.Prettify(s)
}

func (s TransferInCheckMailTokenResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenResponseBodyFailList) SetFailDomain(v []*string) *TransferInCheckMailTokenResponseBodyFailList {
	s.FailDomain = v
	return s
}

type TransferInCheckMailTokenResponseBodySuccessList struct {
	SuccessDomain []*string `json:"SuccessDomain,omitempty" xml:"SuccessDomain,omitempty" type:"Repeated"`
}

func (s TransferInCheckMailTokenResponseBodySuccessList) String() string {
	return tea.Prettify(s)
}

func (s TransferInCheckMailTokenResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenResponseBodySuccessList) SetSuccessDomain(v []*string) *TransferInCheckMailTokenResponseBodySuccessList {
	s.SuccessDomain = v
	return s
}

type TransferInCheckMailTokenResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferInCheckMailTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferInCheckMailTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferInCheckMailTokenResponse) GoString() string {
	return s.String()
}

func (s *TransferInCheckMailTokenResponse) SetHeaders(v map[string]*string) *TransferInCheckMailTokenResponse {
	s.Headers = v
	return s
}

func (s *TransferInCheckMailTokenResponse) SetStatusCode(v int32) *TransferInCheckMailTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferInCheckMailTokenResponse) SetBody(v *TransferInCheckMailTokenResponseBody) *TransferInCheckMailTokenResponse {
	s.Body = v
	return s
}

type TransferInReenterTransferAuthorizationCodeRequest struct {
	DomainName                *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang                      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TransferAuthorizationCode *string `json:"TransferAuthorizationCode,omitempty" xml:"TransferAuthorizationCode,omitempty"`
	UserClientIp              *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s TransferInReenterTransferAuthorizationCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferInReenterTransferAuthorizationCodeRequest) GoString() string {
	return s.String()
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) SetDomainName(v string) *TransferInReenterTransferAuthorizationCodeRequest {
	s.DomainName = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) SetLang(v string) *TransferInReenterTransferAuthorizationCodeRequest {
	s.Lang = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) SetTransferAuthorizationCode(v string) *TransferInReenterTransferAuthorizationCodeRequest {
	s.TransferAuthorizationCode = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) SetUserClientIp(v string) *TransferInReenterTransferAuthorizationCodeRequest {
	s.UserClientIp = &v
	return s
}

type TransferInReenterTransferAuthorizationCodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferInReenterTransferAuthorizationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferInReenterTransferAuthorizationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInReenterTransferAuthorizationCodeResponseBody) SetRequestId(v string) *TransferInReenterTransferAuthorizationCodeResponseBody {
	s.RequestId = &v
	return s
}

type TransferInReenterTransferAuthorizationCodeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferInReenterTransferAuthorizationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferInReenterTransferAuthorizationCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferInReenterTransferAuthorizationCodeResponse) GoString() string {
	return s.String()
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) SetHeaders(v map[string]*string) *TransferInReenterTransferAuthorizationCodeResponse {
	s.Headers = v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) SetStatusCode(v int32) *TransferInReenterTransferAuthorizationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) SetBody(v *TransferInReenterTransferAuthorizationCodeResponseBody) *TransferInReenterTransferAuthorizationCodeResponse {
	s.Body = v
	return s
}

type TransferInRefetchWhoisEmailRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s TransferInRefetchWhoisEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferInRefetchWhoisEmailRequest) GoString() string {
	return s.String()
}

func (s *TransferInRefetchWhoisEmailRequest) SetDomainName(v string) *TransferInRefetchWhoisEmailRequest {
	s.DomainName = &v
	return s
}

func (s *TransferInRefetchWhoisEmailRequest) SetLang(v string) *TransferInRefetchWhoisEmailRequest {
	s.Lang = &v
	return s
}

func (s *TransferInRefetchWhoisEmailRequest) SetUserClientIp(v string) *TransferInRefetchWhoisEmailRequest {
	s.UserClientIp = &v
	return s
}

type TransferInRefetchWhoisEmailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferInRefetchWhoisEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferInRefetchWhoisEmailResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInRefetchWhoisEmailResponseBody) SetRequestId(v string) *TransferInRefetchWhoisEmailResponseBody {
	s.RequestId = &v
	return s
}

type TransferInRefetchWhoisEmailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferInRefetchWhoisEmailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferInRefetchWhoisEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferInRefetchWhoisEmailResponse) GoString() string {
	return s.String()
}

func (s *TransferInRefetchWhoisEmailResponse) SetHeaders(v map[string]*string) *TransferInRefetchWhoisEmailResponse {
	s.Headers = v
	return s
}

func (s *TransferInRefetchWhoisEmailResponse) SetStatusCode(v int32) *TransferInRefetchWhoisEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferInRefetchWhoisEmailResponse) SetBody(v *TransferInRefetchWhoisEmailResponseBody) *TransferInRefetchWhoisEmailResponse {
	s.Body = v
	return s
}

type TransferInResendMailTokenRequest struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s TransferInResendMailTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferInResendMailTokenRequest) GoString() string {
	return s.String()
}

func (s *TransferInResendMailTokenRequest) SetDomainName(v string) *TransferInResendMailTokenRequest {
	s.DomainName = &v
	return s
}

func (s *TransferInResendMailTokenRequest) SetLang(v string) *TransferInResendMailTokenRequest {
	s.Lang = &v
	return s
}

func (s *TransferInResendMailTokenRequest) SetUserClientIp(v string) *TransferInResendMailTokenRequest {
	s.UserClientIp = &v
	return s
}

type TransferInResendMailTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferInResendMailTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferInResendMailTokenResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInResendMailTokenResponseBody) SetRequestId(v string) *TransferInResendMailTokenResponseBody {
	s.RequestId = &v
	return s
}

type TransferInResendMailTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferInResendMailTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferInResendMailTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferInResendMailTokenResponse) GoString() string {
	return s.String()
}

func (s *TransferInResendMailTokenResponse) SetHeaders(v map[string]*string) *TransferInResendMailTokenResponse {
	s.Headers = v
	return s
}

func (s *TransferInResendMailTokenResponse) SetStatusCode(v int32) *TransferInResendMailTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferInResendMailTokenResponse) SetBody(v *TransferInResendMailTokenResponseBody) *TransferInResendMailTokenResponse {
	s.Body = v
	return s
}

type UpdateDomainToDomainGroupRequest struct {
	DataSource    *int32    `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	DomainGroupId *int64    `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainName    []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	FileToUpload  *string   `json:"FileToUpload,omitempty" xml:"FileToUpload,omitempty"`
	Lang          *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Replace       *bool     `json:"Replace,omitempty" xml:"Replace,omitempty"`
	UserClientIp  *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s UpdateDomainToDomainGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainToDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainToDomainGroupRequest) SetDataSource(v int32) *UpdateDomainToDomainGroupRequest {
	s.DataSource = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetDomainGroupId(v int64) *UpdateDomainToDomainGroupRequest {
	s.DomainGroupId = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetDomainName(v []*string) *UpdateDomainToDomainGroupRequest {
	s.DomainName = v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetFileToUpload(v string) *UpdateDomainToDomainGroupRequest {
	s.FileToUpload = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetLang(v string) *UpdateDomainToDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetReplace(v bool) *UpdateDomainToDomainGroupRequest {
	s.Replace = &v
	return s
}

func (s *UpdateDomainToDomainGroupRequest) SetUserClientIp(v string) *UpdateDomainToDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

type UpdateDomainToDomainGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainToDomainGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainToDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainToDomainGroupResponseBody) SetRequestId(v string) *UpdateDomainToDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDomainToDomainGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDomainToDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDomainToDomainGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainToDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainToDomainGroupResponse) SetHeaders(v map[string]*string) *UpdateDomainToDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainToDomainGroupResponse) SetStatusCode(v int32) *UpdateDomainToDomainGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainToDomainGroupResponse) SetBody(v *UpdateDomainToDomainGroupResponseBody) *UpdateDomainToDomainGroupResponse {
	s.Body = v
	return s
}

type VerifyContactFieldRequest struct {
	Address                  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	City                     *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	DomainName               *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Email                    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang                     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PostalCode               *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province                 *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName           *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization   *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RegistrantType           *string `json:"RegistrantType,omitempty" xml:"RegistrantType,omitempty"`
	TelArea                  *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                   *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone                *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	UserClientIp             *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZhAddress                *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s VerifyContactFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyContactFieldRequest) GoString() string {
	return s.String()
}

func (s *VerifyContactFieldRequest) SetAddress(v string) *VerifyContactFieldRequest {
	s.Address = &v
	return s
}

func (s *VerifyContactFieldRequest) SetCity(v string) *VerifyContactFieldRequest {
	s.City = &v
	return s
}

func (s *VerifyContactFieldRequest) SetCountry(v string) *VerifyContactFieldRequest {
	s.Country = &v
	return s
}

func (s *VerifyContactFieldRequest) SetDomainName(v string) *VerifyContactFieldRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyContactFieldRequest) SetEmail(v string) *VerifyContactFieldRequest {
	s.Email = &v
	return s
}

func (s *VerifyContactFieldRequest) SetLang(v string) *VerifyContactFieldRequest {
	s.Lang = &v
	return s
}

func (s *VerifyContactFieldRequest) SetPostalCode(v string) *VerifyContactFieldRequest {
	s.PostalCode = &v
	return s
}

func (s *VerifyContactFieldRequest) SetProvince(v string) *VerifyContactFieldRequest {
	s.Province = &v
	return s
}

func (s *VerifyContactFieldRequest) SetRegistrantName(v string) *VerifyContactFieldRequest {
	s.RegistrantName = &v
	return s
}

func (s *VerifyContactFieldRequest) SetRegistrantOrganization(v string) *VerifyContactFieldRequest {
	s.RegistrantOrganization = &v
	return s
}

func (s *VerifyContactFieldRequest) SetRegistrantType(v string) *VerifyContactFieldRequest {
	s.RegistrantType = &v
	return s
}

func (s *VerifyContactFieldRequest) SetTelArea(v string) *VerifyContactFieldRequest {
	s.TelArea = &v
	return s
}

func (s *VerifyContactFieldRequest) SetTelExt(v string) *VerifyContactFieldRequest {
	s.TelExt = &v
	return s
}

func (s *VerifyContactFieldRequest) SetTelephone(v string) *VerifyContactFieldRequest {
	s.Telephone = &v
	return s
}

func (s *VerifyContactFieldRequest) SetUserClientIp(v string) *VerifyContactFieldRequest {
	s.UserClientIp = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhAddress(v string) *VerifyContactFieldRequest {
	s.ZhAddress = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhCity(v string) *VerifyContactFieldRequest {
	s.ZhCity = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhProvince(v string) *VerifyContactFieldRequest {
	s.ZhProvince = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhRegistrantName(v string) *VerifyContactFieldRequest {
	s.ZhRegistrantName = &v
	return s
}

func (s *VerifyContactFieldRequest) SetZhRegistrantOrganization(v string) *VerifyContactFieldRequest {
	s.ZhRegistrantOrganization = &v
	return s
}

type VerifyContactFieldResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyContactFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyContactFieldResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyContactFieldResponseBody) SetRequestId(v string) *VerifyContactFieldResponseBody {
	s.RequestId = &v
	return s
}

type VerifyContactFieldResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyContactFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyContactFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyContactFieldResponse) GoString() string {
	return s.String()
}

func (s *VerifyContactFieldResponse) SetHeaders(v map[string]*string) *VerifyContactFieldResponse {
	s.Headers = v
	return s
}

func (s *VerifyContactFieldResponse) SetStatusCode(v int32) *VerifyContactFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyContactFieldResponse) SetBody(v *VerifyContactFieldResponseBody) *VerifyContactFieldResponse {
	s.Body = v
	return s
}

type VerifyEmailRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Token        *string `json:"Token,omitempty" xml:"Token,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s VerifyEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyEmailRequest) GoString() string {
	return s.String()
}

func (s *VerifyEmailRequest) SetLang(v string) *VerifyEmailRequest {
	s.Lang = &v
	return s
}

func (s *VerifyEmailRequest) SetToken(v string) *VerifyEmailRequest {
	s.Token = &v
	return s
}

func (s *VerifyEmailRequest) SetUserClientIp(v string) *VerifyEmailRequest {
	s.UserClientIp = &v
	return s
}

type VerifyEmailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyEmailResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyEmailResponseBody) SetRequestId(v string) *VerifyEmailResponseBody {
	s.RequestId = &v
	return s
}

type VerifyEmailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyEmailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyEmailResponse) GoString() string {
	return s.String()
}

func (s *VerifyEmailResponse) SetHeaders(v map[string]*string) *VerifyEmailResponse {
	s.Headers = v
	return s
}

func (s *VerifyEmailResponse) SetStatusCode(v int32) *VerifyEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyEmailResponse) SetBody(v *VerifyEmailResponseBody) *VerifyEmailResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("domain"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AcknowledgeTaskResultWithOptions(request *AcknowledgeTaskResultRequest, runtime *util.RuntimeOptions) (_result *AcknowledgeTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDetailNo)) {
		query["TaskDetailNo"] = request.TaskDetailNo
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AcknowledgeTaskResult"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcknowledgeTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AcknowledgeTaskResult(request *AcknowledgeTaskResultRequest) (_result *AcknowledgeTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcknowledgeTaskResultResponse{}
	_body, _err := client.AcknowledgeTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchFuzzyMatchDomainSensitiveWordWithOptions(request *BatchFuzzyMatchDomainSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *BatchFuzzyMatchDomainSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchFuzzyMatchDomainSensitiveWord"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchFuzzyMatchDomainSensitiveWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchFuzzyMatchDomainSensitiveWord(request *BatchFuzzyMatchDomainSensitiveWordRequest) (_result *BatchFuzzyMatchDomainSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchFuzzyMatchDomainSensitiveWordResponse{}
	_body, _err := client.BatchFuzzyMatchDomainSensitiveWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelDomainVerificationWithOptions(request *CancelDomainVerificationRequest, runtime *util.RuntimeOptions) (_result *CancelDomainVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDomainVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDomainVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelDomainVerification(request *CancelDomainVerificationRequest) (_result *CancelDomainVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelDomainVerificationResponse{}
	_body, _err := client.CancelDomainVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelOperationAuditWithOptions(request *CancelOperationAuditRequest, runtime *util.RuntimeOptions) (_result *CancelOperationAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditRecordId)) {
		query["AuditRecordId"] = request.AuditRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelOperationAudit"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelOperationAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelOperationAudit(request *CancelOperationAuditRequest) (_result *CancelOperationAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelOperationAuditResponse{}
	_body, _err := client.CancelOperationAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelQualificationVerificationWithOptions(request *CancelQualificationVerificationRequest, runtime *util.RuntimeOptions) (_result *CancelQualificationVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationType)) {
		query["QualificationType"] = request.QualificationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelQualificationVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelQualificationVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelQualificationVerification(request *CancelQualificationVerificationRequest) (_result *CancelQualificationVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelQualificationVerificationResponse{}
	_body, _err := client.CancelQualificationVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelTaskWithOptions(request *CancelTaskRequest, runtime *util.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TaskNo)) {
		query["TaskNo"] = request.TaskNo
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelTask"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelTask(request *CancelTaskRequest) (_result *CancelTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDomainWithOptions(request *CheckDomainRequest, runtime *util.RuntimeOptions) (_result *CheckDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FeeCommand)) {
		query["FeeCommand"] = request.FeeCommand
	}

	if !tea.BoolValue(util.IsUnset(request.FeeCurrency)) {
		query["FeeCurrency"] = request.FeeCurrency
	}

	if !tea.BoolValue(util.IsUnset(request.FeePeriod)) {
		query["FeePeriod"] = request.FeePeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDomain"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDomain(request *CheckDomainRequest) (_result *CheckDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDomainResponse{}
	_body, _err := client.CheckDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDomainSunriseClaimWithOptions(request *CheckDomainSunriseClaimRequest, runtime *util.RuntimeOptions) (_result *CheckDomainSunriseClaimResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDomainSunriseClaim"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDomainSunriseClaimResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDomainSunriseClaim(request *CheckDomainSunriseClaimRequest) (_result *CheckDomainSunriseClaimResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDomainSunriseClaimResponse{}
	_body, _err := client.CheckDomainSunriseClaimWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckMaxYearOfServerLockWithOptions(request *CheckMaxYearOfServerLockRequest, runtime *util.RuntimeOptions) (_result *CheckMaxYearOfServerLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckAction)) {
		query["CheckAction"] = request.CheckAction
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckMaxYearOfServerLock"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckMaxYearOfServerLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckMaxYearOfServerLock(request *CheckMaxYearOfServerLockRequest) (_result *CheckMaxYearOfServerLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMaxYearOfServerLockResponse{}
	_body, _err := client.CheckMaxYearOfServerLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckProcessingServerLockApplyWithOptions(request *CheckProcessingServerLockApplyRequest, runtime *util.RuntimeOptions) (_result *CheckProcessingServerLockApplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FeePeriod)) {
		query["FeePeriod"] = request.FeePeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckProcessingServerLockApply"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckProcessingServerLockApplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckProcessingServerLockApply(request *CheckProcessingServerLockApplyRequest) (_result *CheckProcessingServerLockApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckProcessingServerLockApplyResponse{}
	_body, _err := client.CheckProcessingServerLockApplyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckTransferInFeasibilityWithOptions(request *CheckTransferInFeasibilityRequest, runtime *util.RuntimeOptions) (_result *CheckTransferInFeasibilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TransferAuthorizationCode)) {
		query["TransferAuthorizationCode"] = request.TransferAuthorizationCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckTransferInFeasibility"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckTransferInFeasibilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckTransferInFeasibility(request *CheckTransferInFeasibilityRequest) (_result *CheckTransferInFeasibilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckTransferInFeasibilityResponse{}
	_body, _err := client.CheckTransferInFeasibilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmTransferInEmailWithOptions(request *ConfirmTransferInEmailRequest, runtime *util.RuntimeOptions) (_result *ConfirmTransferInEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmTransferInEmail"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmTransferInEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmTransferInEmail(request *ConfirmTransferInEmailRequest) (_result *ConfirmTransferInEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmTransferInEmailResponse{}
	_body, _err := client.ConfirmTransferInEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactTemplatesWithOptions(request *DeleteContactTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteContactTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileIds)) {
		query["RegistrantProfileIds"] = request.RegistrantProfileIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteContactTemplates"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContactTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContactTemplates(request *DeleteContactTemplatesRequest) (_result *DeleteContactTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContactTemplatesResponse{}
	_body, _err := client.DeleteContactTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainGroupWithOptions(request *DeleteDomainGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainGroupId)) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomainGroup"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomainGroup(request *DeleteDomainGroupRequest) (_result *DeleteDomainGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainGroupResponse{}
	_body, _err := client.DeleteDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEmailVerificationWithOptions(request *DeleteEmailVerificationRequest, runtime *util.RuntimeOptions) (_result *DeleteEmailVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEmailVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEmailVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEmailVerification(request *DeleteEmailVerificationRequest) (_result *DeleteEmailVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEmailVerificationResponse{}
	_body, _err := client.DeleteEmailVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRegistrantProfileWithOptions(request *DeleteRegistrantProfileRequest, runtime *util.RuntimeOptions) (_result *DeleteRegistrantProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRegistrantProfile"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRegistrantProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRegistrantProfile(request *DeleteRegistrantProfileRequest) (_result *DeleteRegistrantProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRegistrantProfileResponse{}
	_body, _err := client.DeleteRegistrantProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EmailVerifiedWithOptions(request *EmailVerifiedRequest, runtime *util.RuntimeOptions) (_result *EmailVerifiedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EmailVerified"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EmailVerifiedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EmailVerified(request *EmailVerifiedRequest) (_result *EmailVerifiedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EmailVerifiedResponse{}
	_body, _err := client.EmailVerifiedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FuzzyMatchDomainSensitiveWordWithOptions(request *FuzzyMatchDomainSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *FuzzyMatchDomainSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FuzzyMatchDomainSensitiveWord"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FuzzyMatchDomainSensitiveWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FuzzyMatchDomainSensitiveWord(request *FuzzyMatchDomainSensitiveWordRequest) (_result *FuzzyMatchDomainSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FuzzyMatchDomainSensitiveWordResponse{}
	_body, _err := client.FuzzyMatchDomainSensitiveWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOperationOssUploadPolicyWithOptions(request *GetOperationOssUploadPolicyRequest, runtime *util.RuntimeOptions) (_result *GetOperationOssUploadPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditType)) {
		query["AuditType"] = request.AuditType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOperationOssUploadPolicy"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOperationOssUploadPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOperationOssUploadPolicy(request *GetOperationOssUploadPolicyRequest) (_result *GetOperationOssUploadPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOperationOssUploadPolicyResponse{}
	_body, _err := client.GetOperationOssUploadPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQualificationUploadPolicyWithOptions(request *GetQualificationUploadPolicyRequest, runtime *util.RuntimeOptions) (_result *GetQualificationUploadPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQualificationUploadPolicy"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQualificationUploadPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQualificationUploadPolicy(request *GetQualificationUploadPolicyRequest) (_result *GetQualificationUploadPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQualificationUploadPolicyResponse{}
	_body, _err := client.GetQualificationUploadPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEmailVerificationWithOptions(request *ListEmailVerificationRequest, runtime *util.RuntimeOptions) (_result *ListEmailVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginCreateTime)) {
		query["BeginCreateTime"] = request.BeginCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EndCreateTime)) {
		query["EndCreateTime"] = request.EndCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationStatus)) {
		query["VerificationStatus"] = request.VerificationStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEmailVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEmailVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEmailVerification(request *ListEmailVerificationRequest) (_result *ListEmailVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEmailVerificationResponse{}
	_body, _err := client.ListEmailVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServerLockWithOptions(request *ListServerLockRequest, runtime *util.RuntimeOptions) (_result *ListServerLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginStartDate)) {
		query["BeginStartDate"] = request.BeginStartDate
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndExpireDate)) {
		query["EndExpireDate"] = request.EndExpireDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndStartDate)) {
		query["EndStartDate"] = request.EndStartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LockProductId)) {
		query["LockProductId"] = request.LockProductId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServerLockStatus)) {
		query["ServerLockStatus"] = request.ServerLockStatus
	}

	if !tea.BoolValue(util.IsUnset(request.StartExpireDate)) {
		query["StartExpireDate"] = request.StartExpireDate
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServerLock"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServerLock(request *ListServerLockRequest) (_result *ListServerLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerLockResponse{}
	_body, _err := client.ListServerLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LookupTmchNoticeWithOptions(request *LookupTmchNoticeRequest, runtime *util.RuntimeOptions) (_result *LookupTmchNoticeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClaimKey)) {
		query["ClaimKey"] = request.ClaimKey
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LookupTmchNotice"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LookupTmchNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LookupTmchNotice(request *LookupTmchNoticeRequest) (_result *LookupTmchNoticeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LookupTmchNoticeResponse{}
	_body, _err := client.LookupTmchNoticeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PollTaskResultWithOptions(request *PollTaskResultRequest, runtime *util.RuntimeOptions) (_result *PollTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskNo)) {
		query["TaskNo"] = request.TaskNo
	}

	if !tea.BoolValue(util.IsUnset(request.TaskResultStatus)) {
		query["TaskResultStatus"] = request.TaskResultStatus
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PollTaskResult"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PollTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PollTaskResult(request *PollTaskResultRequest) (_result *PollTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PollTaskResultResponse{}
	_body, _err := client.PollTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAdvancedDomainListWithOptions(request *QueryAdvancedDomainListRequest, runtime *util.RuntimeOptions) (_result *QueryAdvancedDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainGroupId)) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainNameSort)) {
		query["DomainNameSort"] = request.DomainNameSort
	}

	if !tea.BoolValue(util.IsUnset(request.DomainStatus)) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EndExpirationDate)) {
		query["EndExpirationDate"] = request.EndExpirationDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndLength)) {
		query["EndLength"] = request.EndLength
	}

	if !tea.BoolValue(util.IsUnset(request.EndRegistrationDate)) {
		query["EndRegistrationDate"] = request.EndRegistrationDate
	}

	if !tea.BoolValue(util.IsUnset(request.Excluded)) {
		query["Excluded"] = request.Excluded
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludedPrefix)) {
		query["ExcludedPrefix"] = request.ExcludedPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludedSuffix)) {
		query["ExcludedSuffix"] = request.ExcludedSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.ExpirationDateSort)) {
		query["ExpirationDateSort"] = request.ExpirationDateSort
	}

	if !tea.BoolValue(util.IsUnset(request.Form)) {
		query["Form"] = request.Form
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWordPrefix)) {
		query["KeyWordPrefix"] = request.KeyWordPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWordSuffix)) {
		query["KeyWordSuffix"] = request.KeyWordSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductDomainType)) {
		query["ProductDomainType"] = request.ProductDomainType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductDomainTypeSort)) {
		query["ProductDomainTypeSort"] = request.ProductDomainTypeSort
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrationDateSort)) {
		query["RegistrationDateSort"] = request.RegistrationDateSort
	}

	if !tea.BoolValue(util.IsUnset(request.StartExpirationDate)) {
		query["StartExpirationDate"] = request.StartExpirationDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartLength)) {
		query["StartLength"] = request.StartLength
	}

	if !tea.BoolValue(util.IsUnset(request.StartRegistrationDate)) {
		query["StartRegistrationDate"] = request.StartRegistrationDate
	}

	if !tea.BoolValue(util.IsUnset(request.Suffixs)) {
		query["Suffixs"] = request.Suffixs
	}

	if !tea.BoolValue(util.IsUnset(request.TradeType)) {
		query["TradeType"] = request.TradeType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAdvancedDomainList"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAdvancedDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAdvancedDomainList(request *QueryAdvancedDomainListRequest) (_result *QueryAdvancedDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAdvancedDomainListResponse{}
	_body, _err := client.QueryAdvancedDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryArtExtensionWithOptions(request *QueryArtExtensionRequest, runtime *util.RuntimeOptions) (_result *QueryArtExtensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryArtExtension"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryArtExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryArtExtension(request *QueryArtExtensionRequest) (_result *QueryArtExtensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryArtExtensionResponse{}
	_body, _err := client.QueryArtExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryChangeLogListWithOptions(request *QueryChangeLogListRequest, runtime *util.RuntimeOptions) (_result *QueryChangeLogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryChangeLogList"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryChangeLogListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryChangeLogList(request *QueryChangeLogListRequest) (_result *QueryChangeLogListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryChangeLogListResponse{}
	_body, _err := client.QueryChangeLogListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryContactInfoWithOptions(request *QueryContactInfoRequest, runtime *util.RuntimeOptions) (_result *QueryContactInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactType)) {
		query["ContactType"] = request.ContactType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryContactInfo"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryContactInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryContactInfo(request *QueryContactInfoRequest) (_result *QueryContactInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryContactInfoResponse{}
	_body, _err := client.QueryContactInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDSRecordWithOptions(request *QueryDSRecordRequest, runtime *util.RuntimeOptions) (_result *QueryDSRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDSRecord"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDSRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDSRecord(request *QueryDSRecordRequest) (_result *QueryDSRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDSRecordResponse{}
	_body, _err := client.QueryDSRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDnsHostWithOptions(request *QueryDnsHostRequest, runtime *util.RuntimeOptions) (_result *QueryDnsHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDnsHost"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDnsHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDnsHost(request *QueryDnsHostRequest) (_result *QueryDnsHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDnsHostResponse{}
	_body, _err := client.QueryDnsHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDomainAdminDivisionWithOptions(request *QueryDomainAdminDivisionRequest, runtime *util.RuntimeOptions) (_result *QueryDomainAdminDivisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDomainAdminDivision"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDomainAdminDivisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDomainAdminDivision(request *QueryDomainAdminDivisionRequest) (_result *QueryDomainAdminDivisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDomainAdminDivisionResponse{}
	_body, _err := client.QueryDomainAdminDivisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDomainByDomainNameWithOptions(request *QueryDomainByDomainNameRequest, runtime *util.RuntimeOptions) (_result *QueryDomainByDomainNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDomainByDomainName"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDomainByDomainNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDomainByDomainName(request *QueryDomainByDomainNameRequest) (_result *QueryDomainByDomainNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDomainByDomainNameResponse{}
	_body, _err := client.QueryDomainByDomainNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDomainByInstanceIdWithOptions(request *QueryDomainByInstanceIdRequest, runtime *util.RuntimeOptions) (_result *QueryDomainByInstanceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDomainByInstanceId"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDomainByInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDomainByInstanceId(request *QueryDomainByInstanceIdRequest) (_result *QueryDomainByInstanceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDomainByInstanceIdResponse{}
	_body, _err := client.QueryDomainByInstanceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDomainGroupListWithOptions(request *QueryDomainGroupListRequest, runtime *util.RuntimeOptions) (_result *QueryDomainGroupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainGroupName)) {
		query["DomainGroupName"] = request.DomainGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDeletingGroup)) {
		query["ShowDeletingGroup"] = request.ShowDeletingGroup
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDomainGroupList"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDomainGroupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDomainGroupList(request *QueryDomainGroupListRequest) (_result *QueryDomainGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDomainGroupListResponse{}
	_body, _err := client.QueryDomainGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDomainListWithOptions(request *QueryDomainListRequest, runtime *util.RuntimeOptions) (_result *QueryDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainGroupId)) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndExpirationDate)) {
		query["EndExpirationDate"] = request.EndExpirationDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndRegistrationDate)) {
		query["EndRegistrationDate"] = request.EndRegistrationDate
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OrderByType)) {
		query["OrderByType"] = request.OrderByType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderKeyType)) {
		query["OrderKeyType"] = request.OrderKeyType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductDomainType)) {
		query["ProductDomainType"] = request.ProductDomainType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.StartExpirationDate)) {
		query["StartExpirationDate"] = request.StartExpirationDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartRegistrationDate)) {
		query["StartRegistrationDate"] = request.StartRegistrationDate
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDomainList"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDomainList(request *QueryDomainListRequest) (_result *QueryDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDomainListResponse{}
	_body, _err := client.QueryDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDomainRealNameVerificationInfoWithOptions(request *QueryDomainRealNameVerificationInfoRequest, runtime *util.RuntimeOptions) (_result *QueryDomainRealNameVerificationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FetchImage)) {
		query["FetchImage"] = request.FetchImage
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDomainRealNameVerificationInfo"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDomainRealNameVerificationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDomainRealNameVerificationInfo(request *QueryDomainRealNameVerificationInfoRequest) (_result *QueryDomainRealNameVerificationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDomainRealNameVerificationInfoResponse{}
	_body, _err := client.QueryDomainRealNameVerificationInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDomainSuffixWithOptions(request *QueryDomainSuffixRequest, runtime *util.RuntimeOptions) (_result *QueryDomainSuffixResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDomainSuffix"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDomainSuffixResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDomainSuffix(request *QueryDomainSuffixRequest) (_result *QueryDomainSuffixResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDomainSuffixResponse{}
	_body, _err := client.QueryDomainSuffixWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEmailVerificationWithOptions(request *QueryEmailVerificationRequest, runtime *util.RuntimeOptions) (_result *QueryEmailVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEmailVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEmailVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEmailVerification(request *QueryEmailVerificationRequest) (_result *QueryEmailVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEmailVerificationResponse{}
	_body, _err := client.QueryEmailVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEnsAssociationWithOptions(request *QueryEnsAssociationRequest, runtime *util.RuntimeOptions) (_result *QueryEnsAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEnsAssociation"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEnsAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEnsAssociation(request *QueryEnsAssociationRequest) (_result *QueryEnsAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEnsAssociationResponse{}
	_body, _err := client.QueryEnsAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFailReasonForDomainRealNameVerificationWithOptions(request *QueryFailReasonForDomainRealNameVerificationRequest, runtime *util.RuntimeOptions) (_result *QueryFailReasonForDomainRealNameVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RealNameVerificationAction)) {
		query["RealNameVerificationAction"] = request.RealNameVerificationAction
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFailReasonForDomainRealNameVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFailReasonForDomainRealNameVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFailReasonForDomainRealNameVerification(request *QueryFailReasonForDomainRealNameVerificationRequest) (_result *QueryFailReasonForDomainRealNameVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFailReasonForDomainRealNameVerificationResponse{}
	_body, _err := client.QueryFailReasonForDomainRealNameVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFailReasonForRegistrantProfileRealNameVerificationWithOptions(request *QueryFailReasonForRegistrantProfileRealNameVerificationRequest, runtime *util.RuntimeOptions) (_result *QueryFailReasonForRegistrantProfileRealNameVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileID)) {
		query["RegistrantProfileID"] = request.RegistrantProfileID
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFailReasonForRegistrantProfileRealNameVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFailReasonForRegistrantProfileRealNameVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFailReasonForRegistrantProfileRealNameVerification(request *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) (_result *QueryFailReasonForRegistrantProfileRealNameVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFailReasonForRegistrantProfileRealNameVerificationResponse{}
	_body, _err := client.QueryFailReasonForRegistrantProfileRealNameVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFailingReasonListForQualificationWithOptions(request *QueryFailingReasonListForQualificationRequest, runtime *util.RuntimeOptions) (_result *QueryFailingReasonListForQualificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationType)) {
		query["QualificationType"] = request.QualificationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFailingReasonListForQualification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFailingReasonListForQualificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFailingReasonListForQualification(request *QueryFailingReasonListForQualificationRequest) (_result *QueryFailingReasonListForQualificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFailingReasonListForQualificationResponse{}
	_body, _err := client.QueryFailingReasonListForQualificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryLocalEnsAssociationWithOptions(request *QueryLocalEnsAssociationRequest, runtime *util.RuntimeOptions) (_result *QueryLocalEnsAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryLocalEnsAssociation"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLocalEnsAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryLocalEnsAssociation(request *QueryLocalEnsAssociationRequest) (_result *QueryLocalEnsAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryLocalEnsAssociationResponse{}
	_body, _err := client.QueryLocalEnsAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOperationAuditInfoDetailWithOptions(request *QueryOperationAuditInfoDetailRequest, runtime *util.RuntimeOptions) (_result *QueryOperationAuditInfoDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditRecordId)) {
		query["AuditRecordId"] = request.AuditRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOperationAuditInfoDetail"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOperationAuditInfoDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOperationAuditInfoDetail(request *QueryOperationAuditInfoDetailRequest) (_result *QueryOperationAuditInfoDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOperationAuditInfoDetailResponse{}
	_body, _err := client.QueryOperationAuditInfoDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOperationAuditInfoListWithOptions(request *QueryOperationAuditInfoListRequest, runtime *util.RuntimeOptions) (_result *QueryOperationAuditInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AuditType)) {
		query["AuditType"] = request.AuditType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOperationAuditInfoList"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOperationAuditInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOperationAuditInfoList(request *QueryOperationAuditInfoListRequest) (_result *QueryOperationAuditInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOperationAuditInfoListResponse{}
	_body, _err := client.QueryOperationAuditInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryQualificationDetailWithOptions(request *QueryQualificationDetailRequest, runtime *util.RuntimeOptions) (_result *QueryQualificationDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationType)) {
		query["QualificationType"] = request.QualificationType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryQualificationDetail"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryQualificationDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryQualificationDetail(request *QueryQualificationDetailRequest) (_result *QueryQualificationDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryQualificationDetailResponse{}
	_body, _err := client.QueryQualificationDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRegistrantProfileRealNameVerificationInfoWithOptions(request *QueryRegistrantProfileRealNameVerificationInfoRequest, runtime *util.RuntimeOptions) (_result *QueryRegistrantProfileRealNameVerificationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FetchImage)) {
		query["FetchImage"] = request.FetchImage
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRegistrantProfileRealNameVerificationInfo"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRegistrantProfileRealNameVerificationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRegistrantProfileRealNameVerificationInfo(request *QueryRegistrantProfileRealNameVerificationInfoRequest) (_result *QueryRegistrantProfileRealNameVerificationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRegistrantProfileRealNameVerificationInfoResponse{}
	_body, _err := client.QueryRegistrantProfileRealNameVerificationInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRegistrantProfilesWithOptions(request *QueryRegistrantProfilesRequest, runtime *util.RuntimeOptions) (_result *QueryRegistrantProfilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefaultRegistrantProfile)) {
		query["DefaultRegistrantProfile"] = request.DefaultRegistrantProfile
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RealNameStatus)) {
		query["RealNameStatus"] = request.RealNameStatus
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantOrganization)) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileType)) {
		query["RegistrantProfileType"] = request.RegistrantProfileType
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantType)) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantOrganization)) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRegistrantProfiles"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRegistrantProfilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRegistrantProfiles(request *QueryRegistrantProfilesRequest) (_result *QueryRegistrantProfilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRegistrantProfilesResponse{}
	_body, _err := client.QueryRegistrantProfilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServerLockWithOptions(request *QueryServerLockRequest, runtime *util.RuntimeOptions) (_result *QueryServerLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryServerLock"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryServerLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServerLock(request *QueryServerLockRequest) (_result *QueryServerLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServerLockResponse{}
	_body, _err := client.QueryServerLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTaskDetailHistoryWithOptions(request *QueryTaskDetailHistoryRequest, runtime *util.RuntimeOptions) (_result *QueryTaskDetailHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainNameCursor)) {
		query["DomainNameCursor"] = request.DomainNameCursor
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDetailNoCursor)) {
		query["TaskDetailNoCursor"] = request.TaskDetailNoCursor
	}

	if !tea.BoolValue(util.IsUnset(request.TaskNo)) {
		query["TaskNo"] = request.TaskNo
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTaskDetailHistory"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskDetailHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTaskDetailHistory(request *QueryTaskDetailHistoryRequest) (_result *QueryTaskDetailHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskDetailHistoryResponse{}
	_body, _err := client.QueryTaskDetailHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTaskDetailListWithOptions(request *QueryTaskDetailListRequest, runtime *util.RuntimeOptions) (_result *QueryTaskDetailListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskNo)) {
		query["TaskNo"] = request.TaskNo
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTaskDetailList"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskDetailListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTaskDetailList(request *QueryTaskDetailListRequest) (_result *QueryTaskDetailListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskDetailListResponse{}
	_body, _err := client.QueryTaskDetailListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTaskInfoHistoryWithOptions(request *QueryTaskInfoHistoryRequest, runtime *util.RuntimeOptions) (_result *QueryTaskInfoHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginCreateTime)) {
		query["BeginCreateTime"] = request.BeginCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeCursor)) {
		query["CreateTimeCursor"] = request.CreateTimeCursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndCreateTime)) {
		query["EndCreateTime"] = request.EndCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskNoCursor)) {
		query["TaskNoCursor"] = request.TaskNoCursor
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTaskInfoHistory"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskInfoHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTaskInfoHistory(request *QueryTaskInfoHistoryRequest) (_result *QueryTaskInfoHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskInfoHistoryResponse{}
	_body, _err := client.QueryTaskInfoHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTaskListWithOptions(request *QueryTaskListRequest, runtime *util.RuntimeOptions) (_result *QueryTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginCreateTime)) {
		query["BeginCreateTime"] = request.BeginCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndCreateTime)) {
		query["EndCreateTime"] = request.EndCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTaskList"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTaskList(request *QueryTaskListRequest) (_result *QueryTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskListResponse{}
	_body, _err := client.QueryTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTransferInByInstanceIdWithOptions(request *QueryTransferInByInstanceIdRequest, runtime *util.RuntimeOptions) (_result *QueryTransferInByInstanceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTransferInByInstanceId"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTransferInByInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTransferInByInstanceId(request *QueryTransferInByInstanceIdRequest) (_result *QueryTransferInByInstanceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTransferInByInstanceIdResponse{}
	_body, _err := client.QueryTransferInByInstanceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTransferInListWithOptions(request *QueryTransferInListRequest, runtime *util.RuntimeOptions) (_result *QueryTransferInListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SimpleTransferInStatus)) {
		query["SimpleTransferInStatus"] = request.SimpleTransferInStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SubmissionEndDate)) {
		query["SubmissionEndDate"] = request.SubmissionEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.SubmissionStartDate)) {
		query["SubmissionStartDate"] = request.SubmissionStartDate
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTransferInList"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTransferInListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTransferInList(request *QueryTransferInListRequest) (_result *QueryTransferInListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTransferInListResponse{}
	_body, _err := client.QueryTransferInListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTransferOutInfoWithOptions(request *QueryTransferOutInfoRequest, runtime *util.RuntimeOptions) (_result *QueryTransferOutInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTransferOutInfo"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTransferOutInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTransferOutInfo(request *QueryTransferOutInfoRequest) (_result *QueryTransferOutInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTransferOutInfoResponse{}
	_body, _err := client.QueryTransferOutInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegistrantProfileRealNameVerificationWithOptions(request *RegistrantProfileRealNameVerificationRequest, runtime *util.RuntimeOptions) (_result *RegistrantProfileRealNameVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentityCredentialNo)) {
		query["IdentityCredentialNo"] = request.IdentityCredentialNo
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityCredentialType)) {
		query["IdentityCredentialType"] = request.IdentityCredentialType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileID)) {
		query["RegistrantProfileID"] = request.RegistrantProfileID
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentityCredential)) {
		body["IdentityCredential"] = request.IdentityCredential
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RegistrantProfileRealNameVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegistrantProfileRealNameVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegistrantProfileRealNameVerification(request *RegistrantProfileRealNameVerificationRequest) (_result *RegistrantProfileRealNameVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegistrantProfileRealNameVerificationResponse{}
	_body, _err := client.RegistrantProfileRealNameVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResendEmailVerificationWithOptions(request *ResendEmailVerificationRequest, runtime *util.RuntimeOptions) (_result *ResendEmailVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResendEmailVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResendEmailVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResendEmailVerification(request *ResendEmailVerificationRequest) (_result *ResendEmailVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResendEmailVerificationResponse{}
	_body, _err := client.ResendEmailVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetQualificationVerificationWithOptions(request *ResetQualificationVerificationRequest, runtime *util.RuntimeOptions) (_result *ResetQualificationVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetQualificationVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetQualificationVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetQualificationVerification(request *ResetQualificationVerificationRequest) (_result *ResetQualificationVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetQualificationVerificationResponse{}
	_body, _err := client.ResetQualificationVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchDomainRemarkWithOptions(request *SaveBatchDomainRemarkRequest, runtime *util.RuntimeOptions) (_result *SaveBatchDomainRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchDomainRemark"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchDomainRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchDomainRemark(request *SaveBatchDomainRemarkRequest) (_result *SaveBatchDomainRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchDomainRemarkResponse{}
	_body, _err := client.SaveBatchDomainRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForCreatingOrderActivateWithOptions(request *SaveBatchTaskForCreatingOrderActivateRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForCreatingOrderActivateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponNo)) {
		query["CouponNo"] = request.CouponNo
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OrderActivateParam)) {
		query["OrderActivateParam"] = request.OrderActivateParam
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionNo)) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !tea.BoolValue(util.IsUnset(request.UseCoupon)) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotion)) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForCreatingOrderActivate"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForCreatingOrderActivateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForCreatingOrderActivate(request *SaveBatchTaskForCreatingOrderActivateRequest) (_result *SaveBatchTaskForCreatingOrderActivateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForCreatingOrderActivateResponse{}
	_body, _err := client.SaveBatchTaskForCreatingOrderActivateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForCreatingOrderRedeemWithOptions(request *SaveBatchTaskForCreatingOrderRedeemRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForCreatingOrderRedeemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponNo)) {
		query["CouponNo"] = request.CouponNo
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OrderRedeemParam)) {
		query["OrderRedeemParam"] = request.OrderRedeemParam
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionNo)) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !tea.BoolValue(util.IsUnset(request.UseCoupon)) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotion)) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForCreatingOrderRedeem"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForCreatingOrderRedeemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForCreatingOrderRedeem(request *SaveBatchTaskForCreatingOrderRedeemRequest) (_result *SaveBatchTaskForCreatingOrderRedeemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForCreatingOrderRedeemResponse{}
	_body, _err := client.SaveBatchTaskForCreatingOrderRedeemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForCreatingOrderRenewWithOptions(request *SaveBatchTaskForCreatingOrderRenewRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForCreatingOrderRenewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponNo)) {
		query["CouponNo"] = request.CouponNo
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OrderRenewParam)) {
		query["OrderRenewParam"] = request.OrderRenewParam
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionNo)) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !tea.BoolValue(util.IsUnset(request.UseCoupon)) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotion)) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForCreatingOrderRenew"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForCreatingOrderRenewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForCreatingOrderRenew(request *SaveBatchTaskForCreatingOrderRenewRequest) (_result *SaveBatchTaskForCreatingOrderRenewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForCreatingOrderRenewResponse{}
	_body, _err := client.SaveBatchTaskForCreatingOrderRenewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForCreatingOrderTransferWithOptions(request *SaveBatchTaskForCreatingOrderTransferRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForCreatingOrderTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponNo)) {
		query["CouponNo"] = request.CouponNo
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OrderTransferParam)) {
		query["OrderTransferParam"] = request.OrderTransferParam
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionNo)) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !tea.BoolValue(util.IsUnset(request.UseCoupon)) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotion)) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForCreatingOrderTransfer"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForCreatingOrderTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForCreatingOrderTransfer(request *SaveBatchTaskForCreatingOrderTransferRequest) (_result *SaveBatchTaskForCreatingOrderTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForCreatingOrderTransferResponse{}
	_body, _err := client.SaveBatchTaskForCreatingOrderTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForDomainNameProxyServiceWithOptions(request *SaveBatchTaskForDomainNameProxyServiceRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForDomainNameProxyServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForDomainNameProxyService"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForDomainNameProxyServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForDomainNameProxyService(request *SaveBatchTaskForDomainNameProxyServiceRequest) (_result *SaveBatchTaskForDomainNameProxyServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForDomainNameProxyServiceResponse{}
	_body, _err := client.SaveBatchTaskForDomainNameProxyServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForModifyingDomainDnsWithOptions(request *SaveBatchTaskForModifyingDomainDnsRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForModifyingDomainDnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunDns)) {
		query["AliyunDns"] = request.AliyunDns
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainNameServer)) {
		query["DomainNameServer"] = request.DomainNameServer
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForModifyingDomainDns"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForModifyingDomainDnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForModifyingDomainDns(request *SaveBatchTaskForModifyingDomainDnsRequest) (_result *SaveBatchTaskForModifyingDomainDnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForModifyingDomainDnsResponse{}
	_body, _err := client.SaveBatchTaskForModifyingDomainDnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForReserveDropListDomainWithOptions(request *SaveBatchTaskForReserveDropListDomainRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForReserveDropListDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactTemplateId)) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Domains)) {
		query["Domains"] = request.Domains
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForReserveDropListDomain"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForReserveDropListDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForReserveDropListDomain(request *SaveBatchTaskForReserveDropListDomainRequest) (_result *SaveBatchTaskForReserveDropListDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForReserveDropListDomainResponse{}
	_body, _err := client.SaveBatchTaskForReserveDropListDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForTransferProhibitionLockWithOptions(request *SaveBatchTaskForTransferProhibitionLockRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForTransferProhibitionLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForTransferProhibitionLock"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForTransferProhibitionLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForTransferProhibitionLock(request *SaveBatchTaskForTransferProhibitionLockRequest) (_result *SaveBatchTaskForTransferProhibitionLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForTransferProhibitionLockResponse{}
	_body, _err := client.SaveBatchTaskForTransferProhibitionLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForUpdateProhibitionLockWithOptions(request *SaveBatchTaskForUpdateProhibitionLockRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForUpdateProhibitionLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForUpdateProhibitionLock"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForUpdateProhibitionLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForUpdateProhibitionLock(request *SaveBatchTaskForUpdateProhibitionLockRequest) (_result *SaveBatchTaskForUpdateProhibitionLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForUpdateProhibitionLockResponse{}
	_body, _err := client.SaveBatchTaskForUpdateProhibitionLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForUpdatingContactInfoByNewContactWithOptions(request *SaveBatchTaskForUpdatingContactInfoByNewContactRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForUpdatingContactInfoByNewContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.ContactType)) {
		query["ContactType"] = request.ContactType
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PostalCode)) {
		query["PostalCode"] = request.PostalCode
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantName)) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantOrganization)) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantType)) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !tea.BoolValue(util.IsUnset(request.TelArea)) {
		query["TelArea"] = request.TelArea
	}

	if !tea.BoolValue(util.IsUnset(request.TelExt)) {
		query["TelExt"] = request.TelExt
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		query["Telephone"] = request.Telephone
	}

	if !tea.BoolValue(util.IsUnset(request.TransferOutProhibited)) {
		query["TransferOutProhibited"] = request.TransferOutProhibited
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZhAddress)) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ZhCity)) {
		query["ZhCity"] = request.ZhCity
	}

	if !tea.BoolValue(util.IsUnset(request.ZhProvince)) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantName)) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantOrganization)) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForUpdatingContactInfoByNewContact"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForUpdatingContactInfoByNewContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForUpdatingContactInfoByNewContact(request *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) (_result *SaveBatchTaskForUpdatingContactInfoByNewContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForUpdatingContactInfoByNewContactResponse{}
	_body, _err := client.SaveBatchTaskForUpdatingContactInfoByNewContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdWithOptions(request *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest, runtime *util.RuntimeOptions) (_result *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactType)) {
		query["ContactType"] = request.ContactType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.TransferOutProhibited)) {
		query["TransferOutProhibited"] = request.TransferOutProhibited
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId(request *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) (_result *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse{}
	_body, _err := client.SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveDomainGroupWithOptions(request *SaveDomainGroupRequest, runtime *util.RuntimeOptions) (_result *SaveDomainGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainGroupId)) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainGroupName)) {
		query["DomainGroupName"] = request.DomainGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveDomainGroup"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveDomainGroup(request *SaveDomainGroupRequest) (_result *SaveDomainGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveDomainGroupResponse{}
	_body, _err := client.SaveDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveRegistrantProfileWithOptions(request *SaveRegistrantProfileRequest, runtime *util.RuntimeOptions) (_result *SaveRegistrantProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRegistrantProfile)) {
		query["DefaultRegistrantProfile"] = request.DefaultRegistrantProfile
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PostalCode)) {
		query["PostalCode"] = request.PostalCode
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantName)) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantOrganization)) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileType)) {
		query["RegistrantProfileType"] = request.RegistrantProfileType
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantType)) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !tea.BoolValue(util.IsUnset(request.TelArea)) {
		query["TelArea"] = request.TelArea
	}

	if !tea.BoolValue(util.IsUnset(request.TelExt)) {
		query["TelExt"] = request.TelExt
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		query["Telephone"] = request.Telephone
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZhAddress)) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ZhCity)) {
		query["ZhCity"] = request.ZhCity
	}

	if !tea.BoolValue(util.IsUnset(request.ZhProvince)) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantName)) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantOrganization)) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveRegistrantProfile"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveRegistrantProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveRegistrantProfile(request *SaveRegistrantProfileRequest) (_result *SaveRegistrantProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveRegistrantProfileResponse{}
	_body, _err := client.SaveRegistrantProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveRegistrantProfileRealNameVerificationWithOptions(request *SaveRegistrantProfileRealNameVerificationRequest, runtime *util.RuntimeOptions) (_result *SaveRegistrantProfileRealNameVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityCredential)) {
		query["IdentityCredential"] = request.IdentityCredential
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityCredentialNo)) {
		query["IdentityCredentialNo"] = request.IdentityCredentialNo
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityCredentialType)) {
		query["IdentityCredentialType"] = request.IdentityCredentialType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PostalCode)) {
		query["PostalCode"] = request.PostalCode
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantName)) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantOrganization)) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileType)) {
		query["RegistrantProfileType"] = request.RegistrantProfileType
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantType)) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !tea.BoolValue(util.IsUnset(request.TelArea)) {
		query["TelArea"] = request.TelArea
	}

	if !tea.BoolValue(util.IsUnset(request.TelExt)) {
		query["TelExt"] = request.TelExt
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		query["Telephone"] = request.Telephone
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZhAddress)) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ZhCity)) {
		query["ZhCity"] = request.ZhCity
	}

	if !tea.BoolValue(util.IsUnset(request.ZhProvince)) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantName)) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantOrganization)) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveRegistrantProfileRealNameVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveRegistrantProfileRealNameVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveRegistrantProfileRealNameVerification(request *SaveRegistrantProfileRealNameVerificationRequest) (_result *SaveRegistrantProfileRealNameVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveRegistrantProfileRealNameVerificationResponse{}
	_body, _err := client.SaveRegistrantProfileRealNameVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForAddingDSRecordWithOptions(request *SaveSingleTaskForAddingDSRecordRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForAddingDSRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.Digest)) {
		query["Digest"] = request.Digest
	}

	if !tea.BoolValue(util.IsUnset(request.DigestType)) {
		query["DigestType"] = request.DigestType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.KeyTag)) {
		query["KeyTag"] = request.KeyTag
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForAddingDSRecord"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForAddingDSRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForAddingDSRecord(request *SaveSingleTaskForAddingDSRecordRequest) (_result *SaveSingleTaskForAddingDSRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForAddingDSRecordResponse{}
	_body, _err := client.SaveSingleTaskForAddingDSRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForApprovingTransferOutWithOptions(request *SaveSingleTaskForApprovingTransferOutRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForApprovingTransferOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForApprovingTransferOut"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForApprovingTransferOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForApprovingTransferOut(request *SaveSingleTaskForApprovingTransferOutRequest) (_result *SaveSingleTaskForApprovingTransferOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForApprovingTransferOutResponse{}
	_body, _err := client.SaveSingleTaskForApprovingTransferOutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForAssociatingEnsWithOptions(request *SaveSingleTaskForAssociatingEnsRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForAssociatingEnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForAssociatingEns"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForAssociatingEnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForAssociatingEns(request *SaveSingleTaskForAssociatingEnsRequest) (_result *SaveSingleTaskForAssociatingEnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForAssociatingEnsResponse{}
	_body, _err := client.SaveSingleTaskForAssociatingEnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForCancelingTransferInWithOptions(request *SaveSingleTaskForCancelingTransferInRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForCancelingTransferInResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForCancelingTransferIn"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForCancelingTransferInResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForCancelingTransferIn(request *SaveSingleTaskForCancelingTransferInRequest) (_result *SaveSingleTaskForCancelingTransferInResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForCancelingTransferInResponse{}
	_body, _err := client.SaveSingleTaskForCancelingTransferInWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForCancelingTransferOutWithOptions(request *SaveSingleTaskForCancelingTransferOutRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForCancelingTransferOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForCancelingTransferOut"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForCancelingTransferOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForCancelingTransferOut(request *SaveSingleTaskForCancelingTransferOutRequest) (_result *SaveSingleTaskForCancelingTransferOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForCancelingTransferOutResponse{}
	_body, _err := client.SaveSingleTaskForCancelingTransferOutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingDnsHostWithOptions(request *SaveSingleTaskForCreatingDnsHostRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForCreatingDnsHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DnsName)) {
		query["DnsName"] = request.DnsName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForCreatingDnsHost"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForCreatingDnsHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingDnsHost(request *SaveSingleTaskForCreatingDnsHostRequest) (_result *SaveSingleTaskForCreatingDnsHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForCreatingDnsHostResponse{}
	_body, _err := client.SaveSingleTaskForCreatingDnsHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingOrderActivateWithOptions(request *SaveSingleTaskForCreatingOrderActivateRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForCreatingOrderActivateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunDns)) {
		query["AliyunDns"] = request.AliyunDns
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.CouponNo)) {
		query["CouponNo"] = request.CouponNo
	}

	if !tea.BoolValue(util.IsUnset(request.Dns1)) {
		query["Dns1"] = request.Dns1
	}

	if !tea.BoolValue(util.IsUnset(request.Dns2)) {
		query["Dns2"] = request.Dns2
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDomainProxy)) {
		query["EnableDomainProxy"] = request.EnableDomainProxy
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PermitPremiumActivation)) {
		query["PermitPremiumActivation"] = request.PermitPremiumActivation
	}

	if !tea.BoolValue(util.IsUnset(request.PostalCode)) {
		query["PostalCode"] = request.PostalCode
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionNo)) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantName)) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantOrganization)) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantType)) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionDuration)) {
		query["SubscriptionDuration"] = request.SubscriptionDuration
	}

	if !tea.BoolValue(util.IsUnset(request.TelArea)) {
		query["TelArea"] = request.TelArea
	}

	if !tea.BoolValue(util.IsUnset(request.TelExt)) {
		query["TelExt"] = request.TelExt
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		query["Telephone"] = request.Telephone
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkDomainActivation)) {
		query["TrademarkDomainActivation"] = request.TrademarkDomainActivation
	}

	if !tea.BoolValue(util.IsUnset(request.UseCoupon)) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotion)) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZhAddress)) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ZhCity)) {
		query["ZhCity"] = request.ZhCity
	}

	if !tea.BoolValue(util.IsUnset(request.ZhProvince)) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantName)) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantOrganization)) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForCreatingOrderActivate"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForCreatingOrderActivateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingOrderActivate(request *SaveSingleTaskForCreatingOrderActivateRequest) (_result *SaveSingleTaskForCreatingOrderActivateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForCreatingOrderActivateResponse{}
	_body, _err := client.SaveSingleTaskForCreatingOrderActivateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingOrderRedeemWithOptions(request *SaveSingleTaskForCreatingOrderRedeemRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForCreatingOrderRedeemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponNo)) {
		query["CouponNo"] = request.CouponNo
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentExpirationDate)) {
		query["CurrentExpirationDate"] = request.CurrentExpirationDate
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionNo)) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !tea.BoolValue(util.IsUnset(request.UseCoupon)) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotion)) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForCreatingOrderRedeem"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForCreatingOrderRedeemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingOrderRedeem(request *SaveSingleTaskForCreatingOrderRedeemRequest) (_result *SaveSingleTaskForCreatingOrderRedeemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForCreatingOrderRedeemResponse{}
	_body, _err := client.SaveSingleTaskForCreatingOrderRedeemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingOrderRenewWithOptions(request *SaveSingleTaskForCreatingOrderRenewRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForCreatingOrderRenewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CouponNo)) {
		query["CouponNo"] = request.CouponNo
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentExpirationDate)) {
		query["CurrentExpirationDate"] = request.CurrentExpirationDate
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionNo)) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionDuration)) {
		query["SubscriptionDuration"] = request.SubscriptionDuration
	}

	if !tea.BoolValue(util.IsUnset(request.UseCoupon)) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotion)) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForCreatingOrderRenew"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForCreatingOrderRenewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingOrderRenew(request *SaveSingleTaskForCreatingOrderRenewRequest) (_result *SaveSingleTaskForCreatingOrderRenewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForCreatingOrderRenewResponse{}
	_body, _err := client.SaveSingleTaskForCreatingOrderRenewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingOrderTransferWithOptions(request *SaveSingleTaskForCreatingOrderTransferRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForCreatingOrderTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationCode)) {
		query["AuthorizationCode"] = request.AuthorizationCode
	}

	if !tea.BoolValue(util.IsUnset(request.CouponNo)) {
		query["CouponNo"] = request.CouponNo
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PermitPremiumTransfer)) {
		query["PermitPremiumTransfer"] = request.PermitPremiumTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionNo)) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.UseCoupon)) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.UsePromotion)) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForCreatingOrderTransfer"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForCreatingOrderTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForCreatingOrderTransfer(request *SaveSingleTaskForCreatingOrderTransferRequest) (_result *SaveSingleTaskForCreatingOrderTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForCreatingOrderTransferResponse{}
	_body, _err := client.SaveSingleTaskForCreatingOrderTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForDeletingDSRecordWithOptions(request *SaveSingleTaskForDeletingDSRecordRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForDeletingDSRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.KeyTag)) {
		query["KeyTag"] = request.KeyTag
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForDeletingDSRecord"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForDeletingDSRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForDeletingDSRecord(request *SaveSingleTaskForDeletingDSRecordRequest) (_result *SaveSingleTaskForDeletingDSRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForDeletingDSRecordResponse{}
	_body, _err := client.SaveSingleTaskForDeletingDSRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForDeletingDnsHostWithOptions(request *SaveSingleTaskForDeletingDnsHostRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForDeletingDnsHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DnsName)) {
		query["DnsName"] = request.DnsName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForDeletingDnsHost"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForDeletingDnsHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForDeletingDnsHost(request *SaveSingleTaskForDeletingDnsHostRequest) (_result *SaveSingleTaskForDeletingDnsHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForDeletingDnsHostResponse{}
	_body, _err := client.SaveSingleTaskForDeletingDnsHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForDisassociatingEnsWithOptions(request *SaveSingleTaskForDisassociatingEnsRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForDisassociatingEnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForDisassociatingEns"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForDisassociatingEnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForDisassociatingEns(request *SaveSingleTaskForDisassociatingEnsRequest) (_result *SaveSingleTaskForDisassociatingEnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForDisassociatingEnsResponse{}
	_body, _err := client.SaveSingleTaskForDisassociatingEnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForDomainNameProxyServiceWithOptions(request *SaveSingleTaskForDomainNameProxyServiceRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForDomainNameProxyServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForDomainNameProxyService"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForDomainNameProxyServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForDomainNameProxyService(request *SaveSingleTaskForDomainNameProxyServiceRequest) (_result *SaveSingleTaskForDomainNameProxyServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForDomainNameProxyServiceResponse{}
	_body, _err := client.SaveSingleTaskForDomainNameProxyServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForModifyingDSRecordWithOptions(request *SaveSingleTaskForModifyingDSRecordRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForModifyingDSRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.Digest)) {
		query["Digest"] = request.Digest
	}

	if !tea.BoolValue(util.IsUnset(request.DigestType)) {
		query["DigestType"] = request.DigestType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.KeyTag)) {
		query["KeyTag"] = request.KeyTag
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForModifyingDSRecord"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForModifyingDSRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForModifyingDSRecord(request *SaveSingleTaskForModifyingDSRecordRequest) (_result *SaveSingleTaskForModifyingDSRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForModifyingDSRecordResponse{}
	_body, _err := client.SaveSingleTaskForModifyingDSRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForModifyingDnsHostWithOptions(request *SaveSingleTaskForModifyingDnsHostRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForModifyingDnsHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DnsName)) {
		query["DnsName"] = request.DnsName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForModifyingDnsHost"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForModifyingDnsHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForModifyingDnsHost(request *SaveSingleTaskForModifyingDnsHostRequest) (_result *SaveSingleTaskForModifyingDnsHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForModifyingDnsHostResponse{}
	_body, _err := client.SaveSingleTaskForModifyingDnsHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForQueryingTransferAuthorizationCodeWithOptions(request *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForQueryingTransferAuthorizationCode"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForQueryingTransferAuthorizationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForQueryingTransferAuthorizationCode(request *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) (_result *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForQueryingTransferAuthorizationCodeResponse{}
	_body, _err := client.SaveSingleTaskForQueryingTransferAuthorizationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForSaveArtExtensionWithOptions(request *SaveSingleTaskForSaveArtExtensionRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForSaveArtExtensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateOrPeriod)) {
		query["DateOrPeriod"] = request.DateOrPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Features)) {
		query["Features"] = request.Features
	}

	if !tea.BoolValue(util.IsUnset(request.InscriptionsAndMarkings)) {
		query["InscriptionsAndMarkings"] = request.InscriptionsAndMarkings
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Maker)) {
		query["Maker"] = request.Maker
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialsAndTechniques)) {
		query["MaterialsAndTechniques"] = request.MaterialsAndTechniques
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.Reference)) {
		query["Reference"] = request.Reference
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForSaveArtExtension"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForSaveArtExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForSaveArtExtension(request *SaveSingleTaskForSaveArtExtensionRequest) (_result *SaveSingleTaskForSaveArtExtensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForSaveArtExtensionResponse{}
	_body, _err := client.SaveSingleTaskForSaveArtExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForSynchronizingDSRecordWithOptions(request *SaveSingleTaskForSynchronizingDSRecordRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForSynchronizingDSRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForSynchronizingDSRecord"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForSynchronizingDSRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForSynchronizingDSRecord(request *SaveSingleTaskForSynchronizingDSRecordRequest) (_result *SaveSingleTaskForSynchronizingDSRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForSynchronizingDSRecordResponse{}
	_body, _err := client.SaveSingleTaskForSynchronizingDSRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForSynchronizingDnsHostWithOptions(request *SaveSingleTaskForSynchronizingDnsHostRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForSynchronizingDnsHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForSynchronizingDnsHost"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForSynchronizingDnsHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForSynchronizingDnsHost(request *SaveSingleTaskForSynchronizingDnsHostRequest) (_result *SaveSingleTaskForSynchronizingDnsHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForSynchronizingDnsHostResponse{}
	_body, _err := client.SaveSingleTaskForSynchronizingDnsHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForTransferProhibitionLockWithOptions(request *SaveSingleTaskForTransferProhibitionLockRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForTransferProhibitionLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForTransferProhibitionLock"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForTransferProhibitionLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForTransferProhibitionLock(request *SaveSingleTaskForTransferProhibitionLockRequest) (_result *SaveSingleTaskForTransferProhibitionLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForTransferProhibitionLockResponse{}
	_body, _err := client.SaveSingleTaskForTransferProhibitionLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForUpdateProhibitionLockWithOptions(request *SaveSingleTaskForUpdateProhibitionLockRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForUpdateProhibitionLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForUpdateProhibitionLock"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForUpdateProhibitionLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForUpdateProhibitionLock(request *SaveSingleTaskForUpdateProhibitionLockRequest) (_result *SaveSingleTaskForUpdateProhibitionLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForUpdateProhibitionLockResponse{}
	_body, _err := client.SaveSingleTaskForUpdateProhibitionLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSingleTaskForUpdatingContactInfoWithOptions(request *SaveSingleTaskForUpdatingContactInfoRequest, runtime *util.RuntimeOptions) (_result *SaveSingleTaskForUpdatingContactInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddTransferLock)) {
		query["AddTransferLock"] = request.AddTransferLock
	}

	if !tea.BoolValue(util.IsUnset(request.ContactType)) {
		query["ContactType"] = request.ContactType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSingleTaskForUpdatingContactInfo"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSingleTaskForUpdatingContactInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSingleTaskForUpdatingContactInfo(request *SaveSingleTaskForUpdatingContactInfoRequest) (_result *SaveSingleTaskForUpdatingContactInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSingleTaskForUpdatingContactInfoResponse{}
	_body, _err := client.SaveSingleTaskForUpdatingContactInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTaskForSubmittingDomainDeleteWithOptions(request *SaveTaskForSubmittingDomainDeleteRequest, runtime *util.RuntimeOptions) (_result *SaveTaskForSubmittingDomainDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveTaskForSubmittingDomainDelete"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTaskForSubmittingDomainDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTaskForSubmittingDomainDelete(request *SaveTaskForSubmittingDomainDeleteRequest) (_result *SaveTaskForSubmittingDomainDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTaskForSubmittingDomainDeleteResponse{}
	_body, _err := client.SaveTaskForSubmittingDomainDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialWithOptions(request *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest, runtime *util.RuntimeOptions) (_result *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityCredentialNo)) {
		query["IdentityCredentialNo"] = request.IdentityCredentialNo
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityCredentialType)) {
		query["IdentityCredentialType"] = request.IdentityCredentialType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentityCredential)) {
		body["IdentityCredential"] = request.IdentityCredential
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredential"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredential(request *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest) (_result *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse{}
	_body, _err := client.SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDWithOptions(request *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest, runtime *util.RuntimeOptions) (_result *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileID"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileID(request *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest) (_result *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse{}
	_body, _err := client.SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTaskForUpdatingRegistrantInfoByIdentityCredentialWithOptions(request *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest, runtime *util.RuntimeOptions) (_result *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityCredentialNo)) {
		query["IdentityCredentialNo"] = request.IdentityCredentialNo
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityCredentialType)) {
		query["IdentityCredentialType"] = request.IdentityCredentialType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PostalCode)) {
		query["PostalCode"] = request.PostalCode
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantName)) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantOrganization)) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantType)) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !tea.BoolValue(util.IsUnset(request.TelArea)) {
		query["TelArea"] = request.TelArea
	}

	if !tea.BoolValue(util.IsUnset(request.TelExt)) {
		query["TelExt"] = request.TelExt
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		query["Telephone"] = request.Telephone
	}

	if !tea.BoolValue(util.IsUnset(request.TransferOutProhibited)) {
		query["TransferOutProhibited"] = request.TransferOutProhibited
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZhAddress)) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ZhCity)) {
		query["ZhCity"] = request.ZhCity
	}

	if !tea.BoolValue(util.IsUnset(request.ZhProvince)) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantName)) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantOrganization)) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentityCredential)) {
		body["IdentityCredential"] = request.IdentityCredential
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveTaskForUpdatingRegistrantInfoByIdentityCredential"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTaskForUpdatingRegistrantInfoByIdentityCredential(request *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest) (_result *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse{}
	_body, _err := client.SaveTaskForUpdatingRegistrantInfoByIdentityCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDWithOptions(request *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest, runtime *util.RuntimeOptions) (_result *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.TransferOutProhibited)) {
		query["TransferOutProhibited"] = request.TransferOutProhibited
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveTaskForUpdatingRegistrantInfoByRegistrantProfileID"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTaskForUpdatingRegistrantInfoByRegistrantProfileID(request *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) (_result *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse{}
	_body, _err := client.SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScrollDomainListWithOptions(request *ScrollDomainListRequest, runtime *util.RuntimeOptions) (_result *ScrollDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainGroupId)) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainStatus)) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EndExpirationDate)) {
		query["EndExpirationDate"] = request.EndExpirationDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndLength)) {
		query["EndLength"] = request.EndLength
	}

	if !tea.BoolValue(util.IsUnset(request.EndRegistrationDate)) {
		query["EndRegistrationDate"] = request.EndRegistrationDate
	}

	if !tea.BoolValue(util.IsUnset(request.Excluded)) {
		query["Excluded"] = request.Excluded
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludedPrefix)) {
		query["ExcludedPrefix"] = request.ExcludedPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludedSuffix)) {
		query["ExcludedSuffix"] = request.ExcludedSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.Form)) {
		query["Form"] = request.Form
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWordPrefix)) {
		query["KeyWordPrefix"] = request.KeyWordPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWordSuffix)) {
		query["KeyWordSuffix"] = request.KeyWordSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductDomainType)) {
		query["ProductDomainType"] = request.ProductDomainType
	}

	if !tea.BoolValue(util.IsUnset(request.ScrollId)) {
		query["ScrollId"] = request.ScrollId
	}

	if !tea.BoolValue(util.IsUnset(request.StartExpirationDate)) {
		query["StartExpirationDate"] = request.StartExpirationDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartLength)) {
		query["StartLength"] = request.StartLength
	}

	if !tea.BoolValue(util.IsUnset(request.StartRegistrationDate)) {
		query["StartRegistrationDate"] = request.StartRegistrationDate
	}

	if !tea.BoolValue(util.IsUnset(request.Suffixs)) {
		query["Suffixs"] = request.Suffixs
	}

	if !tea.BoolValue(util.IsUnset(request.TradeType)) {
		query["TradeType"] = request.TradeType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ScrollDomainList"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScrollDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScrollDomainList(request *ScrollDomainListRequest) (_result *ScrollDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScrollDomainListResponse{}
	_body, _err := client.ScrollDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDefaultRegistrantProfileWithOptions(request *SetDefaultRegistrantProfileRequest, runtime *util.RuntimeOptions) (_result *SetDefaultRegistrantProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegistrantProfileId)) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDefaultRegistrantProfile"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDefaultRegistrantProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDefaultRegistrantProfile(request *SetDefaultRegistrantProfileRequest) (_result *SetDefaultRegistrantProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultRegistrantProfileResponse{}
	_body, _err := client.SetDefaultRegistrantProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitEmailVerificationWithOptions(request *SubmitEmailVerificationRequest, runtime *util.RuntimeOptions) (_result *SubmitEmailVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SendIfExist)) {
		query["SendIfExist"] = request.SendIfExist
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitEmailVerification"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitEmailVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitEmailVerification(request *SubmitEmailVerificationRequest) (_result *SubmitEmailVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitEmailVerificationResponse{}
	_body, _err := client.SubmitEmailVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitOperationAuditInfoWithOptions(request *SubmitOperationAuditInfoRequest, runtime *util.RuntimeOptions) (_result *SubmitOperationAuditInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditInfo)) {
		query["AuditInfo"] = request.AuditInfo
	}

	if !tea.BoolValue(util.IsUnset(request.AuditType)) {
		query["AuditType"] = request.AuditType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitOperationAuditInfo"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitOperationAuditInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitOperationAuditInfo(request *SubmitOperationAuditInfoRequest) (_result *SubmitOperationAuditInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitOperationAuditInfoResponse{}
	_body, _err := client.SubmitOperationAuditInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitOperationCredentialsWithOptions(request *SubmitOperationCredentialsRequest, runtime *util.RuntimeOptions) (_result *SubmitOperationCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditRecordId)) {
		query["AuditRecordId"] = request.AuditRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.AuditType)) {
		query["AuditType"] = request.AuditType
	}

	if !tea.BoolValue(util.IsUnset(request.Credentials)) {
		query["Credentials"] = request.Credentials
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegType)) {
		query["RegType"] = request.RegType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitOperationCredentials"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitOperationCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitOperationCredentials(request *SubmitOperationCredentialsRequest) (_result *SubmitOperationCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitOperationCredentialsResponse{}
	_body, _err := client.SubmitOperationCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferInCheckMailTokenWithOptions(request *TransferInCheckMailTokenRequest, runtime *util.RuntimeOptions) (_result *TransferInCheckMailTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferInCheckMailToken"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferInCheckMailTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferInCheckMailToken(request *TransferInCheckMailTokenRequest) (_result *TransferInCheckMailTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferInCheckMailTokenResponse{}
	_body, _err := client.TransferInCheckMailTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferInReenterTransferAuthorizationCodeWithOptions(request *TransferInReenterTransferAuthorizationCodeRequest, runtime *util.RuntimeOptions) (_result *TransferInReenterTransferAuthorizationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TransferAuthorizationCode)) {
		query["TransferAuthorizationCode"] = request.TransferAuthorizationCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferInReenterTransferAuthorizationCode"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferInReenterTransferAuthorizationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferInReenterTransferAuthorizationCode(request *TransferInReenterTransferAuthorizationCodeRequest) (_result *TransferInReenterTransferAuthorizationCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferInReenterTransferAuthorizationCodeResponse{}
	_body, _err := client.TransferInReenterTransferAuthorizationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferInRefetchWhoisEmailWithOptions(request *TransferInRefetchWhoisEmailRequest, runtime *util.RuntimeOptions) (_result *TransferInRefetchWhoisEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferInRefetchWhoisEmail"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferInRefetchWhoisEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferInRefetchWhoisEmail(request *TransferInRefetchWhoisEmailRequest) (_result *TransferInRefetchWhoisEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferInRefetchWhoisEmailResponse{}
	_body, _err := client.TransferInRefetchWhoisEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferInResendMailTokenWithOptions(request *TransferInResendMailTokenRequest, runtime *util.RuntimeOptions) (_result *TransferInResendMailTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferInResendMailToken"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferInResendMailTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferInResendMailToken(request *TransferInResendMailTokenRequest) (_result *TransferInResendMailTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferInResendMailTokenResponse{}
	_body, _err := client.TransferInResendMailTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDomainToDomainGroupWithOptions(request *UpdateDomainToDomainGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateDomainToDomainGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		query["DataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.DomainGroupId)) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Replace)) {
		query["Replace"] = request.Replace
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileToUpload)) {
		body["FileToUpload"] = request.FileToUpload
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDomainToDomainGroup"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDomainToDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDomainToDomainGroup(request *UpdateDomainToDomainGroupRequest) (_result *UpdateDomainToDomainGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDomainToDomainGroupResponse{}
	_body, _err := client.UpdateDomainToDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyContactFieldWithOptions(request *VerifyContactFieldRequest, runtime *util.RuntimeOptions) (_result *VerifyContactFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PostalCode)) {
		query["PostalCode"] = request.PostalCode
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantName)) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantOrganization)) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrantType)) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !tea.BoolValue(util.IsUnset(request.TelArea)) {
		query["TelArea"] = request.TelArea
	}

	if !tea.BoolValue(util.IsUnset(request.TelExt)) {
		query["TelExt"] = request.TelExt
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		query["Telephone"] = request.Telephone
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZhAddress)) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ZhCity)) {
		query["ZhCity"] = request.ZhCity
	}

	if !tea.BoolValue(util.IsUnset(request.ZhProvince)) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantName)) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !tea.BoolValue(util.IsUnset(request.ZhRegistrantOrganization)) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyContactField"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyContactFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyContactField(request *VerifyContactFieldRequest) (_result *VerifyContactFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyContactFieldResponse{}
	_body, _err := client.VerifyContactFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyEmailWithOptions(request *VerifyEmailRequest, runtime *util.RuntimeOptions) (_result *VerifyEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyEmail"),
		Version:     tea.String("2018-01-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyEmail(request *VerifyEmailRequest) (_result *VerifyEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyEmailResponse{}
	_body, _err := client.VerifyEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
