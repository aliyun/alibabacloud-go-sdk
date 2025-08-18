// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWafRulesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateWafRulesetResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateWafRulesetResponseBody
	GetRequestId() *string
}

type CreateWafRulesetResponseBody struct {
	// Ruleset ID.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWafRulesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWafRulesetResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateWafRulesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWafRulesetResponseBody) SetId(v int64) *CreateWafRulesetResponseBody {
	s.Id = &v
	return s
}

func (s *CreateWafRulesetResponseBody) SetRequestId(v string) *CreateWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWafRulesetResponseBody) Validate() error {
	return dara.Validate(s)
}
