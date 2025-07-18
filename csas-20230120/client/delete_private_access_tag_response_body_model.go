// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateAccessTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrivateAccessTagResponseBody
	GetRequestId() *string
}

type DeletePrivateAccessTagResponseBody struct {
	// example:
	//
	// FD724DBC-CD76-5235-BF76-59C51B73296D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateAccessTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateAccessTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivateAccessTagResponseBody) SetRequestId(v string) *DeletePrivateAccessTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivateAccessTagResponseBody) Validate() error {
	return dara.Validate(s)
}
