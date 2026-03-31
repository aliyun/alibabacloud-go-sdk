// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDefenseRuleResponseBody
	GetRequestId() *string
	SetRuleIds(v string) *CreateDefenseRuleResponseBody
	GetRuleIds() *string
}

type CreateDefenseRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 26E46541-7AAB-5565-801D-F14DBDC5F186
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the protection rules. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 22215,23354,462165
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
}

func (s CreateDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDefenseRuleResponseBody) GetRuleIds() *string {
	return s.RuleIds
}

func (s *CreateDefenseRuleResponseBody) SetRequestId(v string) *CreateDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefenseRuleResponseBody) SetRuleIds(v string) *CreateDefenseRuleResponseBody {
	s.RuleIds = &v
	return s
}

func (s *CreateDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
