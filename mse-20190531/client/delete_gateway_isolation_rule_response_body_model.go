// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIsolationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteGatewayIsolationRuleResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteGatewayIsolationRuleResponseBody
	GetRequestId() *string
}

type DeleteGatewayIsolationRuleResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 94B12406-E44D-57C9-BF93-A8B35BFF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayIsolationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIsolationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIsolationRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGatewayIsolationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayIsolationRuleResponseBody) SetData(v bool) *DeleteGatewayIsolationRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayIsolationRuleResponseBody) SetRequestId(v string) *DeleteGatewayIsolationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayIsolationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
