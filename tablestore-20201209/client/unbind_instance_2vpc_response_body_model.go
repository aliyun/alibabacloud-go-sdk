// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindInstance2VpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindInstance2VpcResponseBody
	GetRequestId() *string
}

type UnbindInstance2VpcResponseBody struct {
	// example:
	//
	// 18DD77BF-F967-576D-80D1-79121399AB53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindInstance2VpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindInstance2VpcResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindInstance2VpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindInstance2VpcResponseBody) SetRequestId(v string) *UnbindInstance2VpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindInstance2VpcResponseBody) Validate() error {
	return dara.Validate(s)
}
