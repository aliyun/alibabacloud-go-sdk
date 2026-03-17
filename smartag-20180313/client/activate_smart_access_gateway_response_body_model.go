// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActivateSmartAccessGatewayResponseBody
	GetRequestId() *string
}

type ActivateSmartAccessGatewayResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E223E535-AE11-4158-B00F-DC107887A909
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateSmartAccessGatewayResponseBody) SetRequestId(v string) *ActivateSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
