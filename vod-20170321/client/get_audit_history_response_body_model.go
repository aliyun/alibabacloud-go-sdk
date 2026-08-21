// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHistories(v []*GetAuditHistoryResponseBodyHistories) *GetAuditHistoryResponseBody
	GetHistories() []*GetAuditHistoryResponseBodyHistories
	SetRequestId(v string) *GetAuditHistoryResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetAuditHistoryResponseBody
	GetStatus() *string
	SetTotal(v int64) *GetAuditHistoryResponseBody
	GetTotal() *int64
}

type GetAuditHistoryResponseBody struct {
	// The list of review history records.
	Histories []*GetAuditHistoryResponseBodyHistories `json:"Histories,omitempty" xml:"Histories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-43*****D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The review result. Indicates the result of the current manual review. Valid values:
	//
	// - **Normal**: the content is normal.
	//
	// - **Blocked**: the content is blocked.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of review history records.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAuditHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuditHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditHistoryResponseBody) GetHistories() []*GetAuditHistoryResponseBodyHistories {
	return s.Histories
}

func (s *GetAuditHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuditHistoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAuditHistoryResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetAuditHistoryResponseBody) SetHistories(v []*GetAuditHistoryResponseBodyHistories) *GetAuditHistoryResponseBody {
	s.Histories = v
	return s
}

func (s *GetAuditHistoryResponseBody) SetRequestId(v string) *GetAuditHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditHistoryResponseBody) SetStatus(v string) *GetAuditHistoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetAuditHistoryResponseBody) SetTotal(v int64) *GetAuditHistoryResponseBody {
	s.Total = &v
	return s
}

func (s *GetAuditHistoryResponseBody) Validate() error {
	if s.Histories != nil {
		for _, item := range s.Histories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAuditHistoryResponseBodyHistories struct {
	// The reviewer.
	//
	// example:
	//
	// auditor
	Auditor *string `json:"Auditor,omitempty" xml:"Auditor,omitempty"`
	// The review details, which are the specific comments provided by the reviewer.
	//
	// example:
	//
	// Contains nudity
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the record was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The reason for rejection. If the review result is rejection, the reason must be provided.
	//
	// example:
	//
	// Pornographic video
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The review result. Valid values:
	//
	// - **Normal**
	//
	// - **Blocked**
	//
	// example:
	//
	// Blocked
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAuditHistoryResponseBodyHistories) String() string {
	return dara.Prettify(s)
}

func (s GetAuditHistoryResponseBodyHistories) GoString() string {
	return s.String()
}

func (s *GetAuditHistoryResponseBodyHistories) GetAuditor() *string {
	return s.Auditor
}

func (s *GetAuditHistoryResponseBodyHistories) GetComment() *string {
	return s.Comment
}

func (s *GetAuditHistoryResponseBodyHistories) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetAuditHistoryResponseBodyHistories) GetReason() *string {
	return s.Reason
}

func (s *GetAuditHistoryResponseBodyHistories) GetStatus() *string {
	return s.Status
}

func (s *GetAuditHistoryResponseBodyHistories) SetAuditor(v string) *GetAuditHistoryResponseBodyHistories {
	s.Auditor = &v
	return s
}

func (s *GetAuditHistoryResponseBodyHistories) SetComment(v string) *GetAuditHistoryResponseBodyHistories {
	s.Comment = &v
	return s
}

func (s *GetAuditHistoryResponseBodyHistories) SetCreationTime(v string) *GetAuditHistoryResponseBodyHistories {
	s.CreationTime = &v
	return s
}

func (s *GetAuditHistoryResponseBodyHistories) SetReason(v string) *GetAuditHistoryResponseBodyHistories {
	s.Reason = &v
	return s
}

func (s *GetAuditHistoryResponseBodyHistories) SetStatus(v string) *GetAuditHistoryResponseBodyHistories {
	s.Status = &v
	return s
}

func (s *GetAuditHistoryResponseBodyHistories) Validate() error {
	return dara.Validate(s)
}
