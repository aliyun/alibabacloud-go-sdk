// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateContextStoreResponseBody
	GetRequestId() *string
}

type UpdateContextStoreResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateContextStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextStoreResponseBody) SetRequestId(v string) *UpdateContextStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
