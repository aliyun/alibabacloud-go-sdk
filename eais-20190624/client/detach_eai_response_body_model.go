// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEaiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachEaiResponseBody
	GetRequestId() *string
}

type DetachEaiResponseBody struct {
	// example:
	//
	// 04DEB304-2436-4CB9-BB63-468BCEA03D9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachEaiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachEaiResponseBody) GoString() string {
	return s.String()
}

func (s *DetachEaiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachEaiResponseBody) SetRequestId(v string) *DetachEaiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachEaiResponseBody) Validate() error {
	return dara.Validate(s)
}
