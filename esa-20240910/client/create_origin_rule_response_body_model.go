// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateOriginRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateOriginRuleResponseBody
	GetRequestId() *string
}

type CreateOriginRuleResponseBody struct {
	// Configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOriginRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOriginRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateOriginRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOriginRuleResponseBody) SetConfigId(v int64) *CreateOriginRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateOriginRuleResponseBody) SetRequestId(v string) *CreateOriginRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOriginRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
