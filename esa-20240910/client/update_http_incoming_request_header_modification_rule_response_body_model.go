// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpIncomingRequestHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type UpdateHttpIncomingRequestHeaderModificationRuleResponseBody struct {
	// example:
	//
	// BFEF3861-8BB7-5B63-954C-6575EA7FB2CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody) SetRequestId(v string) *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
