// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeAuditLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitializeAuditLogResponseBody
	GetRequestId() *string
}

type InitializeAuditLogResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4FE969D9-E1C7-5274-BE7D-8C3534587605
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeAuditLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeAuditLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeAuditLogResponseBody) SetRequestId(v string) *InitializeAuditLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeAuditLogResponseBody) Validate() error {
	return dara.Validate(s)
}
