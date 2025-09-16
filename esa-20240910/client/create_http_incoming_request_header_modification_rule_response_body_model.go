// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpIncomingRequestHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateHttpIncomingRequestHeaderModificationRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateHttpIncomingRequestHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type CreateHttpIncomingRequestHeaderModificationRuleResponseBody struct {
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// 7C414690-9D7B-5D66-9CD9-AD0B3F25ED49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHttpIncomingRequestHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingRequestHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponseBody) SetConfigId(v int64) *CreateHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponseBody) SetRequestId(v string) *CreateHttpIncomingRequestHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
