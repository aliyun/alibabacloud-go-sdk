// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWafRulesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateUserWafRulesetResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateUserWafRulesetResponseBody
	GetRequestId() *string
	SetRuleIds(v []*int64) *CreateUserWafRulesetResponseBody
	GetRuleIds() []*int64
}

type CreateUserWafRulesetResponseBody struct {
	// example:
	//
	// 665d3af3621bccf3fe29e1a4
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleIds   []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s CreateUserWafRulesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserWafRulesetResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateUserWafRulesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserWafRulesetResponseBody) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *CreateUserWafRulesetResponseBody) SetId(v int64) *CreateUserWafRulesetResponseBody {
	s.Id = &v
	return s
}

func (s *CreateUserWafRulesetResponseBody) SetRequestId(v string) *CreateUserWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserWafRulesetResponseBody) SetRuleIds(v []*int64) *CreateUserWafRulesetResponseBody {
	s.RuleIds = v
	return s
}

func (s *CreateUserWafRulesetResponseBody) Validate() error {
	return dara.Validate(s)
}
