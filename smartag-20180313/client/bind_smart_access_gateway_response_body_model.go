// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindSmartAccessGatewayResponseBody
	GetRequestId() *string
}

type BindSmartAccessGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8A3FF8DD-B27D-4ED2-B032-5EF90B38195D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *BindSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindSmartAccessGatewayResponseBody) SetRequestId(v string) *BindSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
