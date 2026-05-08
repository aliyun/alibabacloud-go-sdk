// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteContextStoreResponseBody
	GetRequestId() *string
}

type DeleteContextStoreResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteContextStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContextStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContextStoreResponseBody) SetRequestId(v string) *DeleteContextStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContextStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
