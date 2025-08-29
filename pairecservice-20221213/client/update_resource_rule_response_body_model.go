// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceRuleResponseBody
	GetRequestId() *string
}

type UpdateResourceRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceRuleResponseBody) SetRequestId(v string) *UpdateResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
