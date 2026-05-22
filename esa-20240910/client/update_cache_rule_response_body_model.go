// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCacheRuleResponseBody
	GetRequestId() *string
}

type UpdateCacheRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCacheRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCacheRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCacheRuleResponseBody) SetRequestId(v string) *UpdateCacheRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCacheRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
