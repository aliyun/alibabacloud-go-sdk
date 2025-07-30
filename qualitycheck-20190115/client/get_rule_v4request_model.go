// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleV4Request interface {
	dara.Model
	String() string
	GoString() string
	SetRuleId(v int64) *GetRuleV4Request
	GetRuleId() *int64
}

type GetRuleV4Request struct {
	// This parameter is required.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetRuleV4Request) String() string {
	return dara.Prettify(s)
}

func (s GetRuleV4Request) GoString() string {
	return s.String()
}

func (s *GetRuleV4Request) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetRuleV4Request) SetRuleId(v int64) *GetRuleV4Request {
	s.RuleId = &v
	return s
}

func (s *GetRuleV4Request) Validate() error {
	return dara.Validate(s)
}
