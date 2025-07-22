// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoLiveStreamRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAutoLiveStreamRuleResponseBody
	GetRequestId() *string
	SetRuleId(v int64) *CreateAutoLiveStreamRuleResponseBody
	GetRuleId() *int64
}

type CreateAutoLiveStreamRuleResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateAutoLiveStreamRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoLiveStreamRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoLiveStreamRuleResponseBody) GetRuleId() *int64 {
	return s.RuleId
}

func (s *CreateAutoLiveStreamRuleResponseBody) SetRequestId(v string) *CreateAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoLiveStreamRuleResponseBody) SetRuleId(v int64) *CreateAutoLiveStreamRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateAutoLiveStreamRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
