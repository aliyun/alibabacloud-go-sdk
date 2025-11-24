// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateASMGatewayImportedServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateASMGatewayImportedServicesResponseBody
	GetRequestId() *string
}

type UpdateASMGatewayImportedServicesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 11fd0027-c27e-41bb-a565-75583054****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateASMGatewayImportedServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateASMGatewayImportedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayImportedServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateASMGatewayImportedServicesResponseBody) SetRequestId(v string) *UpdateASMGatewayImportedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesResponseBody) Validate() error {
	return dara.Validate(s)
}
