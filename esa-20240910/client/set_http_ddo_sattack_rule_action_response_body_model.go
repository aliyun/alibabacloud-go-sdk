// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackRuleActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetHttpDDoSAttackRuleActionResponseBody
	GetRequestId() *string
}

type SetHttpDDoSAttackRuleActionResponseBody struct {
	// Request ID
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetHttpDDoSAttackRuleActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackRuleActionResponseBody) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackRuleActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetHttpDDoSAttackRuleActionResponseBody) SetRequestId(v string) *SetHttpDDoSAttackRuleActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetHttpDDoSAttackRuleActionResponseBody) Validate() error {
	return dara.Validate(s)
}
