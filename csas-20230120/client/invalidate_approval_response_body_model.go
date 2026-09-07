// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidateApprovalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalId(v string) *InvalidateApprovalResponseBody
	GetApprovalId() *string
	SetEffectStatus(v string) *InvalidateApprovalResponseBody
	GetEffectStatus() *string
	SetReportType(v string) *InvalidateApprovalResponseBody
	GetReportType() *string
	SetRequestId(v string) *InvalidateApprovalResponseBody
	GetRequestId() *string
}

type InvalidateApprovalResponseBody struct {
	// The ID of the invalidated approval instance.
	//
	// example:
	//
	// approval-6b5188a28634****
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// The effective status of the approval. When the invalidation succeeds, the value is fixed as Expired, which indicates that the approval has been invalidated.
	//
	// example:
	//
	// Expired
	EffectStatus *string `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// The approval type. Valid values:
	//
	// 	- ApprovalReport: approval.
	//
	// 	- BackendReport: backend approval.
	//
	// example:
	//
	// BackendReport
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvalidateApprovalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvalidateApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidateApprovalResponseBody) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *InvalidateApprovalResponseBody) GetEffectStatus() *string {
	return s.EffectStatus
}

func (s *InvalidateApprovalResponseBody) GetReportType() *string {
	return s.ReportType
}

func (s *InvalidateApprovalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvalidateApprovalResponseBody) SetApprovalId(v string) *InvalidateApprovalResponseBody {
	s.ApprovalId = &v
	return s
}

func (s *InvalidateApprovalResponseBody) SetEffectStatus(v string) *InvalidateApprovalResponseBody {
	s.EffectStatus = &v
	return s
}

func (s *InvalidateApprovalResponseBody) SetReportType(v string) *InvalidateApprovalResponseBody {
	s.ReportType = &v
	return s
}

func (s *InvalidateApprovalResponseBody) SetRequestId(v string) *InvalidateApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvalidateApprovalResponseBody) Validate() error {
	return dara.Validate(s)
}
