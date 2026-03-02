// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryReview interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *LibraryReview
	GetApplicant() *string
	SetArtificatId(v string) *LibraryReview
	GetArtificatId() *string
	SetDeveloperName(v string) *LibraryReview
	GetDeveloperName() *string
	SetFeedbackLibraryInstruction(v string) *LibraryReview
	GetFeedbackLibraryInstruction() *string
	SetFeedbackLibrarySchema(v string) *LibraryReview
	GetFeedbackLibrarySchema() *string
	SetGmtCreate(v string) *LibraryReview
	GetGmtCreate() *string
	SetGroupId(v string) *LibraryReview
	GetGroupId() *string
	SetId(v int64) *LibraryReview
	GetId() *int64
	SetLibraryName(v string) *LibraryReview
	GetLibraryName() *string
	SetLibraryUrl(v string) *LibraryReview
	GetLibraryUrl() *string
	SetMarketId(v string) *LibraryReview
	GetMarketId() *string
	SetMarketName(v string) *LibraryReview
	GetMarketName() *string
	SetRemainTime(v string) *LibraryReview
	GetRemainTime() *string
	SetRequestId(v string) *LibraryReview
	GetRequestId() *string
	SetReviewerId(v string) *LibraryReview
	GetReviewerId() *string
	SetStatus(v string) *LibraryReview
	GetStatus() *string
}

type LibraryReview struct {
	Applicant                  *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	ArtificatId                *string `json:"artificatId,omitempty" xml:"artificatId,omitempty"`
	DeveloperName              *string `json:"developerName,omitempty" xml:"developerName,omitempty"`
	FeedbackLibraryInstruction *string `json:"feedbackLibraryInstruction,omitempty" xml:"feedbackLibraryInstruction,omitempty"`
	FeedbackLibrarySchema      *string `json:"feedbackLibrarySchema,omitempty" xml:"feedbackLibrarySchema,omitempty"`
	GmtCreate                  *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GroupId                    *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Id                         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LibraryName                *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	LibraryUrl                 *string `json:"libraryUrl,omitempty" xml:"libraryUrl,omitempty"`
	MarketId                   *string `json:"marketId,omitempty" xml:"marketId,omitempty"`
	MarketName                 *string `json:"marketName,omitempty" xml:"marketName,omitempty"`
	RemainTime                 *string `json:"remainTime,omitempty" xml:"remainTime,omitempty"`
	RequestId                  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ReviewerId                 *string `json:"reviewerId,omitempty" xml:"reviewerId,omitempty"`
	Status                     *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s LibraryReview) String() string {
	return dara.Prettify(s)
}

func (s LibraryReview) GoString() string {
	return s.String()
}

func (s *LibraryReview) GetApplicant() *string {
	return s.Applicant
}

func (s *LibraryReview) GetArtificatId() *string {
	return s.ArtificatId
}

func (s *LibraryReview) GetDeveloperName() *string {
	return s.DeveloperName
}

func (s *LibraryReview) GetFeedbackLibraryInstruction() *string {
	return s.FeedbackLibraryInstruction
}

func (s *LibraryReview) GetFeedbackLibrarySchema() *string {
	return s.FeedbackLibrarySchema
}

func (s *LibraryReview) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *LibraryReview) GetGroupId() *string {
	return s.GroupId
}

func (s *LibraryReview) GetId() *int64 {
	return s.Id
}

func (s *LibraryReview) GetLibraryName() *string {
	return s.LibraryName
}

func (s *LibraryReview) GetLibraryUrl() *string {
	return s.LibraryUrl
}

func (s *LibraryReview) GetMarketId() *string {
	return s.MarketId
}

func (s *LibraryReview) GetMarketName() *string {
	return s.MarketName
}

func (s *LibraryReview) GetRemainTime() *string {
	return s.RemainTime
}

func (s *LibraryReview) GetRequestId() *string {
	return s.RequestId
}

func (s *LibraryReview) GetReviewerId() *string {
	return s.ReviewerId
}

func (s *LibraryReview) GetStatus() *string {
	return s.Status
}

func (s *LibraryReview) SetApplicant(v string) *LibraryReview {
	s.Applicant = &v
	return s
}

func (s *LibraryReview) SetArtificatId(v string) *LibraryReview {
	s.ArtificatId = &v
	return s
}

func (s *LibraryReview) SetDeveloperName(v string) *LibraryReview {
	s.DeveloperName = &v
	return s
}

func (s *LibraryReview) SetFeedbackLibraryInstruction(v string) *LibraryReview {
	s.FeedbackLibraryInstruction = &v
	return s
}

func (s *LibraryReview) SetFeedbackLibrarySchema(v string) *LibraryReview {
	s.FeedbackLibrarySchema = &v
	return s
}

func (s *LibraryReview) SetGmtCreate(v string) *LibraryReview {
	s.GmtCreate = &v
	return s
}

func (s *LibraryReview) SetGroupId(v string) *LibraryReview {
	s.GroupId = &v
	return s
}

func (s *LibraryReview) SetId(v int64) *LibraryReview {
	s.Id = &v
	return s
}

func (s *LibraryReview) SetLibraryName(v string) *LibraryReview {
	s.LibraryName = &v
	return s
}

func (s *LibraryReview) SetLibraryUrl(v string) *LibraryReview {
	s.LibraryUrl = &v
	return s
}

func (s *LibraryReview) SetMarketId(v string) *LibraryReview {
	s.MarketId = &v
	return s
}

func (s *LibraryReview) SetMarketName(v string) *LibraryReview {
	s.MarketName = &v
	return s
}

func (s *LibraryReview) SetRemainTime(v string) *LibraryReview {
	s.RemainTime = &v
	return s
}

func (s *LibraryReview) SetRequestId(v string) *LibraryReview {
	s.RequestId = &v
	return s
}

func (s *LibraryReview) SetReviewerId(v string) *LibraryReview {
	s.ReviewerId = &v
	return s
}

func (s *LibraryReview) SetStatus(v string) *LibraryReview {
	s.Status = &v
	return s
}

func (s *LibraryReview) Validate() error {
	return dara.Validate(s)
}
