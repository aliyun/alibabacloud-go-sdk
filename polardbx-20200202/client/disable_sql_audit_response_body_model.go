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
}

type DisableSqlAuditResponseBody struct {
	// example:
	//
	// DC3DAE3E-0F8A-4596-9104-F5167C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DisableSqlAuditResponseBody) SetRequestId(v string) *DisableSqlAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSqlAuditResponseBody) Validate() error {
	return dara.Validate(s)
}
