// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSqlAuditEnableStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckSqlAuditEnableStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *CheckSqlAuditEnableStatusResponseBody
	GetStatus() *string
	SetSuccess(v bool) *CheckSqlAuditEnableStatusResponseBody
	GetSuccess() *bool
}

type CheckSqlAuditEnableStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FF13E47D-4E38-4A5A-BA68-32A554******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the SQL audit feature. Valid values:
	//
	// 	- enabled: The SQL audit feature is enabled.
	//
	// 	- disabled: The SQL audit feature is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckSqlAuditEnableStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlAuditEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSqlAuditEnableStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CheckSqlAuditEnableStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetRequestId(v string) *CheckSqlAuditEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetStatus(v string) *CheckSqlAuditEnableStatusResponseBody {
	s.Status = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponseBody) SetSuccess(v bool) *CheckSqlAuditEnableStatusResponseBody {
	s.Success = &v
	return s
}

func (s *CheckSqlAuditEnableStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
