// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAuditLogConfigResponseBody
	GetRequestId() *string
}

type ModifyAuditLogConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8D0C0AFC-E9CD-47A4-8395-5C31BF9B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAuditLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAuditLogConfigResponseBody) SetRequestId(v string) *ModifyAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAuditLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
