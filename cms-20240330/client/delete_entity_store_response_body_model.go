// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEntityStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEntityStoreResponseBody
	GetRequestId() *string
}

type DeleteEntityStoreResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteEntityStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEntityStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEntityStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEntityStoreResponseBody) SetRequestId(v string) *DeleteEntityStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEntityStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
