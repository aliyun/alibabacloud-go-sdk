// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMcpServerResponseBody
	GetRequestId() *string
}

type DeleteMcpServerResponseBody struct {
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcpServerResponseBody) SetRequestId(v string) *DeleteMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
