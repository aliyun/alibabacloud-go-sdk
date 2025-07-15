// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectTrafficQosRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExpressConnectTrafficQosRuleResponseBody
	GetRequestId() *string
}

type DeleteExpressConnectTrafficQosRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExpressConnectTrafficQosRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectTrafficQosRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectTrafficQosRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExpressConnectTrafficQosRuleResponseBody) SetRequestId(v string) *DeleteExpressConnectTrafficQosRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
