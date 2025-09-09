// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSqlAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableSqlAuditResponseBody
	GetRequestId() *string
	SetResult(v bool) *DisableSqlAuditResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DisableSqlAuditResponseBody
	GetSuccess() *bool
}

type DisableSqlAuditResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E273749A-9A33-44CF-ABE7-0CB19C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return result.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSqlAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSqlAuditResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSqlAuditResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DisableSqlAuditResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableSqlAuditResponseBody) SetRequestId(v string) *DisableSqlAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSqlAuditResponseBody) SetResult(v bool) *DisableSqlAuditResponseBody {
	s.Result = &v
	return s
}

func (s *DisableSqlAuditResponseBody) SetSuccess(v bool) *DisableSqlAuditResponseBody {
	s.Success = &v
	return s
}

func (s *DisableSqlAuditResponseBody) Validate() error {
	return dara.Validate(s)
}
