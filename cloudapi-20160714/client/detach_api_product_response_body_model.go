// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachApiProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachApiProductResponseBody
	GetRequestId() *string
}

type DetachApiProductResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98E4A7DC-1EA6-5E6A-ACFE-91B60CE7D4BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachApiProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachApiProductResponseBody) GoString() string {
	return s.String()
}

func (s *DetachApiProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachApiProductResponseBody) SetRequestId(v string) *DetachApiProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachApiProductResponseBody) Validate() error {
	return dara.Validate(s)
}
