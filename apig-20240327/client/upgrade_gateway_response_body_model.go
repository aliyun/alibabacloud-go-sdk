// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpgradeGatewayResponseBody
	GetCode() *string
	SetMessage(v string) *UpgradeGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeGatewayResponseBody
	GetRequestId() *string
}

type UpgradeGatewayResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 043360DA-ED3B-5386-9B7A-D94DECF99A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpgradeGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeGatewayResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeGatewayResponseBody) SetCode(v string) *UpgradeGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeGatewayResponseBody) SetMessage(v string) *UpgradeGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeGatewayResponseBody) SetRequestId(v string) *UpgradeGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
