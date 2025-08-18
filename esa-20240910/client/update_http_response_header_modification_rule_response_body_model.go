// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpResponseHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHttpResponseHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type UpdateHttpResponseHeaderModificationRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHttpResponseHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpResponseHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpResponseHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpResponseHeaderModificationRuleResponseBody) SetRequestId(v string) *UpdateHttpResponseHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
