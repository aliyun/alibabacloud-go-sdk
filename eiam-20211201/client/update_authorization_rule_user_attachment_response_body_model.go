// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleUserAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAuthorizationRuleUserAttachmentResponseBody
	GetRequestId() *string
}

type UpdateAuthorizationRuleUserAttachmentResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuthorizationRuleUserAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleUserAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleUserAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthorizationRuleUserAttachmentResponseBody) SetRequestId(v string) *UpdateAuthorizationRuleUserAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
