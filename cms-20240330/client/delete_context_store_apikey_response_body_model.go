// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextStoreAPIKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteContextStoreAPIKeyResponseBody
	GetRequestId() *string
}

type DeleteContextStoreAPIKeyResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteContextStoreAPIKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextStoreAPIKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContextStoreAPIKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContextStoreAPIKeyResponseBody) SetRequestId(v string) *DeleteContextStoreAPIKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContextStoreAPIKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
