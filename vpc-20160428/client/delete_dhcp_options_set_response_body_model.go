// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDhcpOptionsSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDhcpOptionsSetResponseBody
	GetRequestId() *string
}

type DeleteDhcpOptionsSetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDhcpOptionsSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDhcpOptionsSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDhcpOptionsSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDhcpOptionsSetResponseBody) SetRequestId(v string) *DeleteDhcpOptionsSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDhcpOptionsSetResponseBody) Validate() error {
	return dara.Validate(s)
}
