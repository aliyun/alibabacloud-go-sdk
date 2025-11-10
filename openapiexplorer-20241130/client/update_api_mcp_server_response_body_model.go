// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApiMcpServerResponseBody
	GetRequestId() *string
}

type UpdateApiMcpServerResponseBody struct {
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateApiMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApiMcpServerResponseBody) SetRequestId(v string) *UpdateApiMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApiMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
