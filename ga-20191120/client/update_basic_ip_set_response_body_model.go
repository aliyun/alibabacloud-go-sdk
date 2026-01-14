// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicIpSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateBasicIpSetResponseBody
	GetRequestId() *string
}

type UpdateBasicIpSetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6D2BFF54-6AF2-4679-88C4-2F2E187F16CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBasicIpSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicIpSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBasicIpSetResponseBody) SetRequestId(v string) *UpdateBasicIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBasicIpSetResponseBody) Validate() error {
	return dara.Validate(s)
}
