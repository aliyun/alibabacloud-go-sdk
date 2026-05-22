// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOriginRuleResponseBody
	GetRequestId() *string
}

type DeleteOriginRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOriginRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOriginRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOriginRuleResponseBody) SetRequestId(v string) *DeleteOriginRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOriginRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
