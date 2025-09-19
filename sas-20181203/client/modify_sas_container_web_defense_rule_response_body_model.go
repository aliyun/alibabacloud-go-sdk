// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySasContainerWebDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySasContainerWebDefenseRuleResponseBody
	GetRequestId() *string
}

type ModifySasContainerWebDefenseRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A447E4E3-42A3-58B7-A7D4-2287745BEFDC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySasContainerWebDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySasContainerWebDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySasContainerWebDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySasContainerWebDefenseRuleResponseBody) SetRequestId(v string) *ModifySasContainerWebDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySasContainerWebDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
