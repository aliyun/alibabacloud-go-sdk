// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnlockSmartAccessGatewayResponseBody
	GetRequestId() *string
}

type UnlockSmartAccessGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 05DC546B-DBF9-4028-88CD-1742AB4E014C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockSmartAccessGatewayResponseBody) SetRequestId(v string) *UnlockSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
