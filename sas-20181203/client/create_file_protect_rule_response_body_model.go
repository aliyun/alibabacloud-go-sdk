// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileProtectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateFileProtectRuleResponseBody
	GetRequestId() *string
	SetRuleId(v int64) *CreateFileProtectRuleResponseBody
	GetRuleId() *int64
}

type CreateFileProtectRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC7C8984-D108-516F-9D36-3DF1D1228CCA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 123
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateFileProtectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileProtectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileProtectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileProtectRuleResponseBody) GetRuleId() *int64 {
	return s.RuleId
}

func (s *CreateFileProtectRuleResponseBody) SetRequestId(v string) *CreateFileProtectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileProtectRuleResponseBody) SetRuleId(v int64) *CreateFileProtectRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateFileProtectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
