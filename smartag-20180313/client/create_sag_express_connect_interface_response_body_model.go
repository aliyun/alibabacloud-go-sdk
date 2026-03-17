// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSagExpressConnectInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSagExpressConnectInterfaceResponseBody
	GetRequestId() *string
}

type CreateSagExpressConnectInterfaceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AFF7E5A6-6897-4FDC-A5A8-1978B5B3E545
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSagExpressConnectInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSagExpressConnectInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSagExpressConnectInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSagExpressConnectInterfaceResponseBody) SetRequestId(v string) *CreateSagExpressConnectInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
