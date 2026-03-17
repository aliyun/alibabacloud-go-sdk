// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSagExpressConnectInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSagExpressConnectInterfaceResponseBody
	GetRequestId() *string
}

type DeleteSagExpressConnectInterfaceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AFF7E5A6-6897-4FDC-A5A8-1978B5B3E545
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSagExpressConnectInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSagExpressConnectInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSagExpressConnectInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSagExpressConnectInterfaceResponseBody) SetRequestId(v string) *DeleteSagExpressConnectInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
