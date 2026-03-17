// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootSmartAccessGatewayResponseBody
	GetRequestId() *string
}

type RebootSmartAccessGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B1792769-5CC3-4D6F-A5A5-E6408EBFBAD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *RebootSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootSmartAccessGatewayResponseBody) SetRequestId(v string) *RebootSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
