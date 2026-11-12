// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomImageResponseBody
	GetRequestId() *string
}

type DeleteCustomImageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomImageResponseBody) SetRequestId(v string) *DeleteCustomImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomImageResponseBody) Validate() error {
	return dara.Validate(s)
}
