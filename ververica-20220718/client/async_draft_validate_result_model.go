// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncDraftValidateResult interface {
	dara.Model
	String() string
	GoString() string
	SetDraftValidationDetail(v *DraftValidationDetail) *AsyncDraftValidateResult
	GetDraftValidationDetail() *DraftValidationDetail
	SetMessage(v string) *AsyncDraftValidateResult
	GetMessage() *string
	SetSuccess(v bool) *AsyncDraftValidateResult
	GetSuccess() *bool
	SetTicketStatus(v string) *AsyncDraftValidateResult
	GetTicketStatus() *string
}

type AsyncDraftValidateResult struct {
	DraftValidationDetail *DraftValidationDetail `json:"draftValidationDetail,omitempty" xml:"draftValidationDetail,omitempty"`
	Message               *string                `json:"message,omitempty" xml:"message,omitempty"`
	Success               *bool                  `json:"success,omitempty" xml:"success,omitempty"`
	TicketStatus          *string                `json:"ticketStatus,omitempty" xml:"ticketStatus,omitempty"`
}

func (s AsyncDraftValidateResult) String() string {
	return dara.Prettify(s)
}

func (s AsyncDraftValidateResult) GoString() string {
	return s.String()
}

func (s *AsyncDraftValidateResult) GetDraftValidationDetail() *DraftValidationDetail {
	return s.DraftValidationDetail
}

func (s *AsyncDraftValidateResult) GetMessage() *string {
	return s.Message
}

func (s *AsyncDraftValidateResult) GetSuccess() *bool {
	return s.Success
}

func (s *AsyncDraftValidateResult) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *AsyncDraftValidateResult) SetDraftValidationDetail(v *DraftValidationDetail) *AsyncDraftValidateResult {
	s.DraftValidationDetail = v
	return s
}

func (s *AsyncDraftValidateResult) SetMessage(v string) *AsyncDraftValidateResult {
	s.Message = &v
	return s
}

func (s *AsyncDraftValidateResult) SetSuccess(v bool) *AsyncDraftValidateResult {
	s.Success = &v
	return s
}

func (s *AsyncDraftValidateResult) SetTicketStatus(v string) *AsyncDraftValidateResult {
	s.TicketStatus = &v
	return s
}

func (s *AsyncDraftValidateResult) Validate() error {
	if s.DraftValidationDetail != nil {
		if err := s.DraftValidationDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
