// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIsolationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateGatewayIsolationRuleResponseBody
	GetData() *int64
	SetRequestId(v string) *CreateGatewayIsolationRuleResponseBody
	GetRequestId() *string
}

type CreateGatewayIsolationRuleResponseBody struct {
	// example:
	//
	// 608
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 94B12406-E44D-57C9-BF93-A8B35BFF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGatewayIsolationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayIsolationRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateGatewayIsolationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayIsolationRuleResponseBody) SetData(v int64) *CreateGatewayIsolationRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGatewayIsolationRuleResponseBody) SetRequestId(v string) *CreateGatewayIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayIsolationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
