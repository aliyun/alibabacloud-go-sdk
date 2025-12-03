// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEaisEiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachEaisEiResponseBody
	GetRequestId() *string
}

type DetachEaisEiResponseBody struct {
	// example:
	//
	// 04DEB304-2436-4CB9-BB63-468BCEA0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachEaisEiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *DetachEaisEiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachEaisEiResponseBody) SetRequestId(v string) *DetachEaisEiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachEaisEiResponseBody) Validate() error {
	return dara.Validate(s)
}
