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
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
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
