// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWafRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWafRuleResponseBody
	GetRequestId() *string
}

type DeleteWafRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWafRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWafRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWafRuleResponseBody) SetRequestId(v string) *DeleteWafRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWafRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
