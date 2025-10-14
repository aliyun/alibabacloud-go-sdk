// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNormalizationRuleVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNormalizationRuleVersionResponseBody
	GetRequestId() *string
}

type DeleteNormalizationRuleVersionResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNormalizationRuleVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNormalizationRuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNormalizationRuleVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNormalizationRuleVersionResponseBody) SetRequestId(v string) *DeleteNormalizationRuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNormalizationRuleVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
