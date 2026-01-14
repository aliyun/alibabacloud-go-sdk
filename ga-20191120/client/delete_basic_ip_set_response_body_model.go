// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicIpSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBasicIpSetResponseBody
	GetRequestId() *string
}

type DeleteBasicIpSetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6D2BFF54-6AF2-4679-88C4-2F2E187F16CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicIpSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicIpSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBasicIpSetResponseBody) SetRequestId(v string) *DeleteBasicIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBasicIpSetResponseBody) Validate() error {
	return dara.Validate(s)
}
