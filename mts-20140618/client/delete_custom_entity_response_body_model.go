// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomEntityResponseBody
	GetRequestId() *string
}

type DeleteCustomEntityResponseBody struct {
	// example:
	//
	// 580e8ce3-3b80-44c5-9f3f-36ac3cc5bdd5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomEntityResponseBody) SetRequestId(v string) *DeleteCustomEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
