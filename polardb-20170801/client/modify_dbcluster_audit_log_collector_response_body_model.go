// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAuditLogCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterAuditLogCollectorResponseBody
	GetRequestId() *string
}

type ModifyDBClusterAuditLogCollectorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C0ACF0-DD29-4B67-9190-B7A48C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterAuditLogCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAuditLogCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAuditLogCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterAuditLogCollectorResponseBody) SetRequestId(v string) *ModifyDBClusterAuditLogCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorResponseBody) Validate() error {
	return dara.Validate(s)
}
