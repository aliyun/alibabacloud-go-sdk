// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpResponseHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateHttpResponseHeaderModificationRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateHttpResponseHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type CreateHttpResponseHeaderModificationRuleResponseBody struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHttpResponseHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpResponseHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpResponseHeaderModificationRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateHttpResponseHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpResponseHeaderModificationRuleResponseBody) SetConfigId(v int64) *CreateHttpResponseHeaderModificationRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleResponseBody) SetRequestId(v string) *CreateHttpResponseHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpResponseHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
