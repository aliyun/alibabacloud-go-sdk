// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNormalizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNormalizationRuleResponseBody
	GetRequestId() *string
}

type UpdateNormalizationRuleResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNormalizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNormalizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNormalizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNormalizationRuleResponseBody) SetRequestId(v string) *UpdateNormalizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNormalizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
