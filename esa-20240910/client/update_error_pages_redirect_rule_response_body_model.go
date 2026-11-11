// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErrorPagesRedirectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateErrorPagesRedirectRuleResponseBody
	GetRequestId() *string
}

type UpdateErrorPagesRedirectRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErrorPagesRedirectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateErrorPagesRedirectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErrorPagesRedirectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateErrorPagesRedirectRuleResponseBody) SetRequestId(v string) *UpdateErrorPagesRedirectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
