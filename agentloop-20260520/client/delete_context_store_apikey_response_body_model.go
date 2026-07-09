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
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
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
