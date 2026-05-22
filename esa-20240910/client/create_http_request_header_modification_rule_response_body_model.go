// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpRequestHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateHttpRequestHeaderModificationRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateHttpRequestHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type CreateHttpRequestHeaderModificationRuleResponseBody struct {
	ConfigId  *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHttpRequestHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpRequestHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpRequestHeaderModificationRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateHttpRequestHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpRequestHeaderModificationRuleResponseBody) SetConfigId(v int64) *CreateHttpRequestHeaderModificationRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleResponseBody) SetRequestId(v string) *CreateHttpRequestHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
