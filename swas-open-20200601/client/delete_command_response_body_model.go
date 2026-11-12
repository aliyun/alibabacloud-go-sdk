// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCommandResponseBody
	GetRequestId() *string
}

type DeleteCommandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommandResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCommandResponseBody) SetRequestId(v string) *DeleteCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
