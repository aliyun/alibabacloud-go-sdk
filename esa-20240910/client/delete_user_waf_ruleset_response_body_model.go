// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserWafRulesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserWafRulesetResponseBody
	GetRequestId() *string
}

type DeleteUserWafRulesetResponseBody struct {
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserWafRulesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserWafRulesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserWafRulesetResponseBody) SetRequestId(v string) *DeleteUserWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserWafRulesetResponseBody) Validate() error {
	return dara.Validate(s)
}
