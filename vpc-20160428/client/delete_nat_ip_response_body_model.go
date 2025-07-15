// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNatIpResponseBody
	GetRequestId() *string
}

type DeleteNatIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E9AD97A0-5338-43F8-8A80-5E274CCBA11B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNatIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNatIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNatIpResponseBody) SetRequestId(v string) *DeleteNatIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNatIpResponseBody) Validate() error {
	return dara.Validate(s)
}
