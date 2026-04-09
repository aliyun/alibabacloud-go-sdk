// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApiMcpServerResponseBody
	GetRequestId() *string
}

type DeleteApiMcpServerResponseBody struct {
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteApiMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiMcpServerResponseBody) SetRequestId(v string) *DeleteApiMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
