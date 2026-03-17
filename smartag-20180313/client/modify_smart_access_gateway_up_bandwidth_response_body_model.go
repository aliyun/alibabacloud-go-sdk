// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmartAccessGatewayUpBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySmartAccessGatewayUpBandwidthResponseBody
	GetRequestId() *string
}

type ModifySmartAccessGatewayUpBandwidthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EE837E9F-BD50-4C2B-9E47-260F9D848480
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySmartAccessGatewayUpBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySmartAccessGatewayUpBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySmartAccessGatewayUpBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySmartAccessGatewayUpBandwidthResponseBody) SetRequestId(v string) *ModifySmartAccessGatewayUpBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySmartAccessGatewayUpBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
