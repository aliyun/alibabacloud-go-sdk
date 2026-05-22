// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpRequestHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHttpRequestHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type DeleteHttpRequestHeaderModificationRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHttpRequestHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpRequestHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpRequestHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpRequestHeaderModificationRuleResponseBody) SetRequestId(v string) *DeleteHttpRequestHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpRequestHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
