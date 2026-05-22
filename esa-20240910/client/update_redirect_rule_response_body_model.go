// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRedirectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRedirectRuleResponseBody
	GetRequestId() *string
}

type UpdateRedirectRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRedirectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRedirectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRedirectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRedirectRuleResponseBody) SetRequestId(v string) *UpdateRedirectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRedirectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
