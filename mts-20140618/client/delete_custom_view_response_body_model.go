// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomViewResponseBody
	GetRequestId() *string
}

type DeleteCustomViewResponseBody struct {
	// example:
	//
	// 580e8ce3-3b80-44c5-9f3f-36ac3cc5bdd5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomViewResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomViewResponseBody) SetRequestId(v string) *DeleteCustomViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomViewResponseBody) Validate() error {
	return dara.Validate(s)
}
