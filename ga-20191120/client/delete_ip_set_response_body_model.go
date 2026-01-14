// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpSetResponseBody
	GetRequestId() *string
}

type DeleteIpSetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD61839A-5CC5-404B-8C6E-56066F0C432D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpSetResponseBody) SetRequestId(v string) *DeleteIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpSetResponseBody) Validate() error {
	return dara.Validate(s)
}
