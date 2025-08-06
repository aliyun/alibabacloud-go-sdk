// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMediaLifecycleRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *CreateMediaLifecycleRuleResponseBody
	GetRuleId() *string
}

type CreateMediaLifecycleRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateMediaLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMediaLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMediaLifecycleRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateMediaLifecycleRuleResponseBody) SetRequestId(v string) *CreateMediaLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMediaLifecycleRuleResponseBody) SetRuleId(v string) *CreateMediaLifecycleRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateMediaLifecycleRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
