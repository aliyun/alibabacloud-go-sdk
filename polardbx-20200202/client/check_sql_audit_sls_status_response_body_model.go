// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSqlAuditSlsStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckSqlAuditSlsStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *CheckSqlAuditSlsStatusResponseBody
	GetStatus() *string
}

type CheckSqlAuditSlsStatusResponseBody struct {
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckSqlAuditSlsStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlAuditSlsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditSlsStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSqlAuditSlsStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CheckSqlAuditSlsStatusResponseBody) SetRequestId(v string) *CheckSqlAuditSlsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSqlAuditSlsStatusResponseBody) SetStatus(v string) *CheckSqlAuditSlsStatusResponseBody {
	s.Status = &v
	return s
}

func (s *CheckSqlAuditSlsStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
