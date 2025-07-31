// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditLogFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAuditLogFilterResponseBody
	GetRequestId() *string
}

type ModifyAuditLogFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E209BE2B-F264-4B9D-81F6-A5A5FB1FBF28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAuditLogFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditLogFilterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAuditLogFilterResponseBody) SetRequestId(v string) *ModifyAuditLogFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAuditLogFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
