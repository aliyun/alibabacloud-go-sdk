// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateContextStoreResponseBody
	GetRequestId() *string
}

type CreateContextStoreResponseBody struct {
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateContextStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContextStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContextStoreResponseBody) SetRequestId(v string) *CreateContextStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContextStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
