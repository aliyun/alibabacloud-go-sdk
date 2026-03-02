// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcInvokeReview interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *PbcInvokeReview
	GetApplicant() *string
	SetGmtCreate(v string) *PbcInvokeReview
	GetGmtCreate() *string
	SetId(v int64) *PbcInvokeReview
	GetId() *int64
	SetInvokeId(v int64) *PbcInvokeReview
	GetInvokeId() *int64
	SetInvokePbcId(v int64) *PbcInvokeReview
	GetInvokePbcId() *int64
	SetInvokePbcName(v string) *PbcInvokeReview
	GetInvokePbcName() *string
	SetPbcId(v int64) *PbcInvokeReview
	GetPbcId() *int64
	SetPbcName(v string) *PbcInvokeReview
	GetPbcName() *string
	SetRequestId(v string) *PbcInvokeReview
	GetRequestId() *string
	SetReviewer(v string) *PbcInvokeReview
	GetReviewer() *string
	SetReviewerId(v string) *PbcInvokeReview
	GetReviewerId() *string
	SetStatus(v string) *PbcInvokeReview
	GetStatus() *string
	SetUsage(v string) *PbcInvokeReview
	GetUsage() *string
}

type PbcInvokeReview struct {
	Applicant     *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	InvokeId      *int64  `json:"invokeId,omitempty" xml:"invokeId,omitempty"`
	InvokePbcId   *int64  `json:"invokePbcId,omitempty" xml:"invokePbcId,omitempty"`
	InvokePbcName *string `json:"invokePbcName,omitempty" xml:"invokePbcName,omitempty"`
	PbcId         *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	PbcName       *string `json:"pbcName,omitempty" xml:"pbcName,omitempty"`
	RequestId     *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Reviewer      *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	ReviewerId    *string `json:"reviewerId,omitempty" xml:"reviewerId,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
	Usage         *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s PbcInvokeReview) String() string {
	return dara.Prettify(s)
}

func (s PbcInvokeReview) GoString() string {
	return s.String()
}

func (s *PbcInvokeReview) GetApplicant() *string {
	return s.Applicant
}

func (s *PbcInvokeReview) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PbcInvokeReview) GetId() *int64 {
	return s.Id
}

func (s *PbcInvokeReview) GetInvokeId() *int64 {
	return s.InvokeId
}

func (s *PbcInvokeReview) GetInvokePbcId() *int64 {
	return s.InvokePbcId
}

func (s *PbcInvokeReview) GetInvokePbcName() *string {
	return s.InvokePbcName
}

func (s *PbcInvokeReview) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PbcInvokeReview) GetPbcName() *string {
	return s.PbcName
}

func (s *PbcInvokeReview) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcInvokeReview) GetReviewer() *string {
	return s.Reviewer
}

func (s *PbcInvokeReview) GetReviewerId() *string {
	return s.ReviewerId
}

func (s *PbcInvokeReview) GetStatus() *string {
	return s.Status
}

func (s *PbcInvokeReview) GetUsage() *string {
	return s.Usage
}

func (s *PbcInvokeReview) SetApplicant(v string) *PbcInvokeReview {
	s.Applicant = &v
	return s
}

func (s *PbcInvokeReview) SetGmtCreate(v string) *PbcInvokeReview {
	s.GmtCreate = &v
	return s
}

func (s *PbcInvokeReview) SetId(v int64) *PbcInvokeReview {
	s.Id = &v
	return s
}

func (s *PbcInvokeReview) SetInvokeId(v int64) *PbcInvokeReview {
	s.InvokeId = &v
	return s
}

func (s *PbcInvokeReview) SetInvokePbcId(v int64) *PbcInvokeReview {
	s.InvokePbcId = &v
	return s
}

func (s *PbcInvokeReview) SetInvokePbcName(v string) *PbcInvokeReview {
	s.InvokePbcName = &v
	return s
}

func (s *PbcInvokeReview) SetPbcId(v int64) *PbcInvokeReview {
	s.PbcId = &v
	return s
}

func (s *PbcInvokeReview) SetPbcName(v string) *PbcInvokeReview {
	s.PbcName = &v
	return s
}

func (s *PbcInvokeReview) SetRequestId(v string) *PbcInvokeReview {
	s.RequestId = &v
	return s
}

func (s *PbcInvokeReview) SetReviewer(v string) *PbcInvokeReview {
	s.Reviewer = &v
	return s
}

func (s *PbcInvokeReview) SetReviewerId(v string) *PbcInvokeReview {
	s.ReviewerId = &v
	return s
}

func (s *PbcInvokeReview) SetStatus(v string) *PbcInvokeReview {
	s.Status = &v
	return s
}

func (s *PbcInvokeReview) SetUsage(v string) *PbcInvokeReview {
	s.Usage = &v
	return s
}

func (s *PbcInvokeReview) Validate() error {
	return dara.Validate(s)
}
