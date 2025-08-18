// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompressionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateCompressionRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateCompressionRuleResponseBody
	GetRequestId() *string
}

type CreateCompressionRuleResponseBody struct {
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
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCompressionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCompressionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCompressionRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateCompressionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCompressionRuleResponseBody) SetConfigId(v int64) *CreateCompressionRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateCompressionRuleResponseBody) SetRequestId(v string) *CreateCompressionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCompressionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
