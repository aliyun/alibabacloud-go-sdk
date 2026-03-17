// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSmartAccessGatewayClientUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetSmartAccessGatewayClientUserPasswordResponseBody
	GetRequestId() *string
}

type ResetSmartAccessGatewayClientUserPasswordResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BE1F7E80-4558-4021-B6D2-B94DA8AAAF81
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetSmartAccessGatewayClientUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetSmartAccessGatewayClientUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponseBody) SetRequestId(v string) *ResetSmartAccessGatewayClientUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
