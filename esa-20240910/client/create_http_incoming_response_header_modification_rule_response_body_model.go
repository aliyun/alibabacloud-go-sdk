// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpIncomingResponseHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateHttpIncomingResponseHeaderModificationRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateHttpIncomingResponseHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type CreateHttpIncomingResponseHeaderModificationRuleResponseBody struct {
	ConfigId  *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHttpIncomingResponseHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingResponseHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponseBody) SetConfigId(v int64) *CreateHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponseBody) SetRequestId(v string) *CreateHttpIncomingResponseHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
