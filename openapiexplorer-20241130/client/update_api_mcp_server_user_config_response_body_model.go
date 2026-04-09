// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiMcpServerUserConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApiMcpServerUserConfigResponseBody
	GetRequestId() *string
}

type UpdateApiMcpServerUserConfigResponseBody struct {
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateApiMcpServerUserConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerUserConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApiMcpServerUserConfigResponseBody) SetRequestId(v string) *UpdateApiMcpServerUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApiMcpServerUserConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
