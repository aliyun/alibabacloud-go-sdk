// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWafRulesetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWafRulesetResponseBody
	GetRequestId() *string
}

type DeleteWafRulesetResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWafRulesetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWafRulesetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWafRulesetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWafRulesetResponseBody) SetRequestId(v string) *DeleteWafRulesetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWafRulesetResponseBody) Validate() error {
	return dara.Validate(s)
}
