// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceRuleResponseBody
	GetRequestId() *string
}

type DeleteResourceRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceRuleResponseBody) SetRequestId(v string) *DeleteResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
