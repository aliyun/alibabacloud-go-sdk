// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDiagnosisResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeInstanceDiagnosisResultResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*DescribeInstanceDiagnosisResultResponseBodyData) *DescribeInstanceDiagnosisResultResponseBody
	GetData() []*DescribeInstanceDiagnosisResultResponseBodyData
	SetErrCode(v string) *DescribeInstanceDiagnosisResultResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeInstanceDiagnosisResultResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeInstanceDiagnosisResultResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeInstanceDiagnosisResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceDiagnosisResultResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeInstanceDiagnosisResultResponseBody
	GetTotal() *int32
}

type DescribeInstanceDiagnosisResultResponseBody struct {
	// AccessDeniedDetail
	//
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*DescribeInstanceDiagnosisResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceDiagnosisResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDiagnosisResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDiagnosisResultResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeInstanceDiagnosisResultResponseBody) GetData() []*DescribeInstanceDiagnosisResultResponseBodyData {
	return s.Data
}

func (s *DescribeInstanceDiagnosisResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeInstanceDiagnosisResultResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeInstanceDiagnosisResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInstanceDiagnosisResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceDiagnosisResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceDiagnosisResultResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeInstanceDiagnosisResultResponseBody) SetAccessDeniedDetail(v string) *DescribeInstanceDiagnosisResultResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBody) SetData(v []*DescribeInstanceDiagnosisResultResponseBodyData) *DescribeInstanceDiagnosisResultResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBody) SetErrCode(v string) *DescribeInstanceDiagnosisResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBody) SetErrMessage(v string) *DescribeInstanceDiagnosisResultResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceDiagnosisResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBody) SetRequestId(v string) *DescribeInstanceDiagnosisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBody) SetSuccess(v bool) *DescribeInstanceDiagnosisResultResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBody) SetTotal(v int32) *DescribeInstanceDiagnosisResultResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceDiagnosisResultResponseBodyData struct {
	BestPractice *string `json:"BestPractice,omitempty" xml:"BestPractice,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// table_analysis
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// example:
	//
	// 2026-03-08T02:12:32Z
	EvaluationTime *int64 `json:"EvaluationTime,omitempty" xml:"EvaluationTime,omitempty"`
	// example:
	//
	// 10.0
	FullScore *float64 `json:"FullScore,omitempty" xml:"FullScore,omitempty"`
	// example:
	//
	// c-b25e21e243889XXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// for autotest
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// 02cf887a
	ItemId   *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// example:
	//
	// 2026-03-08
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
	// example:
	//
	// 10.0
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// healthy
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DescribeInstanceDiagnosisResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDiagnosisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetBestPractice() *string {
	return s.BestPractice
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetEvaluationTime() *int64 {
	return s.EvaluationTime
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetFullScore() *float64 {
	return s.FullScore
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetIntroduction() *string {
	return s.Introduction
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetItemId() *string {
	return s.ItemId
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetItemName() *string {
	return s.ItemName
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetReportDate() *string {
	return s.ReportDate
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetScore() *float64 {
	return s.Score
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) GetSuggestion() *string {
	return s.Suggestion
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetBestPractice(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.BestPractice = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetDescription(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetDimension(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.Dimension = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetEvaluationTime(v int64) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.EvaluationTime = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetFullScore(v float64) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.FullScore = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetInstanceId(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetIntroduction(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.Introduction = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetItemId(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.ItemId = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetItemName(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.ItemName = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetReportDate(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.ReportDate = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetScore(v float64) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.Score = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetStatus(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) SetSuggestion(v string) *DescribeInstanceDiagnosisResultResponseBodyData {
	s.Suggestion = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
