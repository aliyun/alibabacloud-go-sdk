// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataMaskingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataMaskingRuleResponseBody
	GetRequestId() *string
}

type DeleteDataMaskingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataMaskingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataMaskingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataMaskingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataMaskingRuleResponseBody) SetRequestId(v string) *DeleteDataMaskingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataMaskingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
