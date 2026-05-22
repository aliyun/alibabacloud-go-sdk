// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRedirectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRedirectRuleResponseBody
	GetRequestId() *string
}

type DeleteRedirectRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRedirectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRedirectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRedirectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRedirectRuleResponseBody) SetRequestId(v string) *DeleteRedirectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRedirectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
