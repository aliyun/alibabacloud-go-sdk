// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomResponseCodeRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomResponseCodeRuleResponseBody
	GetRequestId() *string
}

type UpdateCustomResponseCodeRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomResponseCodeRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomResponseCodeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomResponseCodeRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomResponseCodeRuleResponseBody) SetRequestId(v string) *UpdateCustomResponseCodeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomResponseCodeRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
