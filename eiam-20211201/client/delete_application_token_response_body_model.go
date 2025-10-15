// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApplicationTokenResponseBody
	GetRequestId() *string
}

type DeleteApplicationTokenResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationTokenResponseBody) SetRequestId(v string) *DeleteApplicationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
