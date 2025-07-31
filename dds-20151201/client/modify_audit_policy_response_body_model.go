// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAuditPolicyResponseBody
	GetRequestId() *string
}

type ModifyAuditPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC04D812-F18D-4568-9B88-F260D9590116
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAuditPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAuditPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAuditPolicyResponseBody) SetRequestId(v string) *ModifyAuditPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAuditPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
