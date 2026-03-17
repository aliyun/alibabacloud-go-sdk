// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagExpressConnectInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagExpressConnectInterfaceResponseBody
	GetRequestId() *string
}

type ModifySagExpressConnectInterfaceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AFF7E5A6-6897-4FDC-A5A8-1978B5B3E545
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagExpressConnectInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagExpressConnectInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagExpressConnectInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagExpressConnectInterfaceResponseBody) SetRequestId(v string) *ModifySagExpressConnectInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
