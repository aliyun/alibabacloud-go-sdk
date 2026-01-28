// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleGroupAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAuthorizationRuleGroupAttachmentResponseBody
	GetRequestId() *string
}

type UpdateAuthorizationRuleGroupAttachmentResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuthorizationRuleGroupAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleGroupAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponseBody) SetRequestId(v string) *UpdateAuthorizationRuleGroupAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
