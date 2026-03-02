// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcReview interface {
	dara.Model
	String() string
	GoString() string
	SetDeveloperId(v int64) *PbcReview
	GetDeveloperId() *int64
	SetDeveloperName(v string) *PbcReview
	GetDeveloperName() *string
	SetFeedbackPbcApi(v string) *PbcReview
	GetFeedbackPbcApi() *string
	SetFeedbackPbcInstruction(v string) *PbcReview
	GetFeedbackPbcInstruction() *string
	SetFeedbackPbcSchema(v string) *PbcReview
	GetFeedbackPbcSchema() *string
	SetFeedbackSecurity(v string) *PbcReview
	GetFeedbackSecurity() *string
	SetFeedbackServiceAvailable(v string) *PbcReview
	GetFeedbackServiceAvailable() *string
	SetId(v int64) *PbcReview
	GetId() *int64
	SetMarketId(v int64) *PbcReview
	GetMarketId() *int64
	SetMarketName(v string) *PbcReview
	GetMarketName() *string
	SetPbcName(v string) *PbcReview
	GetPbcName() *string
	SetPbcUrl(v string) *PbcReview
	GetPbcUrl() *string
	SetPbcVersion(v string) *PbcReview
	GetPbcVersion() *string
	SetRemainTime(v string) *PbcReview
	GetRemainTime() *string
	SetRequestId(v string) *PbcReview
	GetRequestId() *string
	SetReviewerId(v int64) *PbcReview
	GetReviewerId() *int64
	SetStatus(v string) *PbcReview
	GetStatus() *string
}

type PbcReview struct {
	// example:
	//
	// 1
	DeveloperId   *int64  `json:"developerId,omitempty" xml:"developerId,omitempty"`
	DeveloperName *string `json:"developerName,omitempty" xml:"developerName,omitempty"`
	// example:
	//
	// 符合规范
	FeedbackPbcApi *string `json:"feedbackPbcApi,omitempty" xml:"feedbackPbcApi,omitempty"`
	// example:
	//
	// 符合规范
	FeedbackPbcInstruction *string `json:"feedbackPbcInstruction,omitempty" xml:"feedbackPbcInstruction,omitempty"`
	// example:
	//
	// 符合pbc规格规范
	FeedbackPbcSchema *string `json:"feedbackPbcSchema,omitempty" xml:"feedbackPbcSchema,omitempty"`
	// example:
	//
	// 服务安全
	FeedbackSecurity *string `json:"feedbackSecurity,omitempty" xml:"feedbackSecurity,omitempty"`
	// example:
	//
	// 服务可用
	FeedbackServiceAvailable *string `json:"feedbackServiceAvailable,omitempty" xml:"feedbackServiceAvailable,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MarketId   *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	MarketName *string `json:"marketName,omitempty" xml:"marketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 基础商品
	PbcName *string `json:"pbcName,omitempty" xml:"pbcName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://catalog.e2.aliyun.com/pbc/product
	PbcUrl *string `json:"pbcUrl,omitempty" xml:"pbcUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-04-01T00:00:00.000Z
	PbcVersion *string `json:"pbcVersion,omitempty" xml:"pbcVersion,omitempty"`
	// example:
	//
	// 4
	RemainTime *string `json:"remainTime,omitempty" xml:"remainTime,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewerId *int64 `json:"reviewerId,omitempty" xml:"reviewerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// REVIEWING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PbcReview) String() string {
	return dara.Prettify(s)
}

func (s PbcReview) GoString() string {
	return s.String()
}

func (s *PbcReview) GetDeveloperId() *int64 {
	return s.DeveloperId
}

func (s *PbcReview) GetDeveloperName() *string {
	return s.DeveloperName
}

func (s *PbcReview) GetFeedbackPbcApi() *string {
	return s.FeedbackPbcApi
}

func (s *PbcReview) GetFeedbackPbcInstruction() *string {
	return s.FeedbackPbcInstruction
}

func (s *PbcReview) GetFeedbackPbcSchema() *string {
	return s.FeedbackPbcSchema
}

func (s *PbcReview) GetFeedbackSecurity() *string {
	return s.FeedbackSecurity
}

func (s *PbcReview) GetFeedbackServiceAvailable() *string {
	return s.FeedbackServiceAvailable
}

func (s *PbcReview) GetId() *int64 {
	return s.Id
}

func (s *PbcReview) GetMarketId() *int64 {
	return s.MarketId
}

func (s *PbcReview) GetMarketName() *string {
	return s.MarketName
}

func (s *PbcReview) GetPbcName() *string {
	return s.PbcName
}

func (s *PbcReview) GetPbcUrl() *string {
	return s.PbcUrl
}

func (s *PbcReview) GetPbcVersion() *string {
	return s.PbcVersion
}

func (s *PbcReview) GetRemainTime() *string {
	return s.RemainTime
}

func (s *PbcReview) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcReview) GetReviewerId() *int64 {
	return s.ReviewerId
}

func (s *PbcReview) GetStatus() *string {
	return s.Status
}

func (s *PbcReview) SetDeveloperId(v int64) *PbcReview {
	s.DeveloperId = &v
	return s
}

func (s *PbcReview) SetDeveloperName(v string) *PbcReview {
	s.DeveloperName = &v
	return s
}

func (s *PbcReview) SetFeedbackPbcApi(v string) *PbcReview {
	s.FeedbackPbcApi = &v
	return s
}

func (s *PbcReview) SetFeedbackPbcInstruction(v string) *PbcReview {
	s.FeedbackPbcInstruction = &v
	return s
}

func (s *PbcReview) SetFeedbackPbcSchema(v string) *PbcReview {
	s.FeedbackPbcSchema = &v
	return s
}

func (s *PbcReview) SetFeedbackSecurity(v string) *PbcReview {
	s.FeedbackSecurity = &v
	return s
}

func (s *PbcReview) SetFeedbackServiceAvailable(v string) *PbcReview {
	s.FeedbackServiceAvailable = &v
	return s
}

func (s *PbcReview) SetId(v int64) *PbcReview {
	s.Id = &v
	return s
}

func (s *PbcReview) SetMarketId(v int64) *PbcReview {
	s.MarketId = &v
	return s
}

func (s *PbcReview) SetMarketName(v string) *PbcReview {
	s.MarketName = &v
	return s
}

func (s *PbcReview) SetPbcName(v string) *PbcReview {
	s.PbcName = &v
	return s
}

func (s *PbcReview) SetPbcUrl(v string) *PbcReview {
	s.PbcUrl = &v
	return s
}

func (s *PbcReview) SetPbcVersion(v string) *PbcReview {
	s.PbcVersion = &v
	return s
}

func (s *PbcReview) SetRemainTime(v string) *PbcReview {
	s.RemainTime = &v
	return s
}

func (s *PbcReview) SetRequestId(v string) *PbcReview {
	s.RequestId = &v
	return s
}

func (s *PbcReview) SetReviewerId(v int64) *PbcReview {
	s.ReviewerId = &v
	return s
}

func (s *PbcReview) SetStatus(v string) *PbcReview {
	s.Status = &v
	return s
}

func (s *PbcReview) Validate() error {
	return dara.Validate(s)
}
