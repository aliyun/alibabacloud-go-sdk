// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserWafRulesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserWafRulesetResponseBody
	GetRequestId() *string
}

type UpdateUserWafRulesetResponseBody struct {
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserWafRulesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserWafRulesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserWafRulesetResponseBody) SetRequestId(v string) *UpdateUserWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserWafRulesetResponseBody) Validate() error {
	return dara.Validate(s)
}
