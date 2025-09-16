// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpIncomingResponseHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type UpdateHttpIncomingResponseHeaderModificationRuleResponseBody struct {
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHttpIncomingResponseHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingResponseHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody) SetRequestId(v string) *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
