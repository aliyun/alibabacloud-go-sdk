// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachCenResponseBody
	GetRequestId() *string
}

type DetachCenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachCenResponseBody) GoString() string {
	return s.String()
}

func (s *DetachCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachCenResponseBody) SetRequestId(v string) *DetachCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachCenResponseBody) Validate() error {
	return dara.Validate(s)
}
