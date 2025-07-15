// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectTrafficQosRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyExpressConnectTrafficQosRuleResponseBody
	GetRequestId() *string
}

type ModifyExpressConnectTrafficQosRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B48B4B9-1EAD-469F-B488-594DAB4B6A1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyExpressConnectTrafficQosRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressConnectTrafficQosRuleResponseBody) SetRequestId(v string) *ModifyExpressConnectTrafficQosRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
