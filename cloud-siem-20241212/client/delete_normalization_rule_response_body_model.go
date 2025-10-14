// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNormalizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNormalizationRuleResponseBody
	GetRequestId() *string
}

type DeleteNormalizationRuleResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNormalizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNormalizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNormalizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNormalizationRuleResponseBody) SetRequestId(v string) *DeleteNormalizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNormalizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
