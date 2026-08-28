// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayMaintenancePeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGatewayMaintenancePeriodResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateGatewayMaintenancePeriodResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayMaintenancePeriodResponseBody
	GetRequestId() *string
}

type UpdateGatewayMaintenancePeriodResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 8FA817D1-CCB0-5776-A604-8FC5DE6DACB9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayMaintenancePeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayMaintenancePeriodResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayMaintenancePeriodResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGatewayMaintenancePeriodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayMaintenancePeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayMaintenancePeriodResponseBody) SetCode(v string) *UpdateGatewayMaintenancePeriodResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayMaintenancePeriodResponseBody) SetMessage(v string) *UpdateGatewayMaintenancePeriodResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayMaintenancePeriodResponseBody) SetRequestId(v string) *UpdateGatewayMaintenancePeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayMaintenancePeriodResponseBody) Validate() error {
	return dara.Validate(s)
}
