// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRedirectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateRedirectRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateRedirectRuleResponseBody
	GetRequestId() *string
}

type CreateRedirectRuleResponseBody struct {
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
	// 1FCB0DA6-9B6D-509D-B91C-B9B9F0780D0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRedirectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRedirectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRedirectRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateRedirectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRedirectRuleResponseBody) SetConfigId(v int64) *CreateRedirectRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateRedirectRuleResponseBody) SetRequestId(v string) *CreateRedirectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRedirectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
