// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInstanceKeyPairResponseBody
	GetRequestId() *string
}

type DeleteInstanceKeyPairResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceKeyPairResponseBody) SetRequestId(v string) *DeleteInstanceKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}
