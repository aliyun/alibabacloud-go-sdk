// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWafRulesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWafRulesetResponseBody
	GetRequestId() *string
}

type UpdateWafRulesetResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWafRulesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWafRulesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWafRulesetResponseBody) SetRequestId(v string) *UpdateWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWafRulesetResponseBody) Validate() error {
	return dara.Validate(s)
}
