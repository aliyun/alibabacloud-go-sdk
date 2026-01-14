// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceManagedControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateServiceManagedControlResponseBody
	GetRequestId() *string
}

type UpdateServiceManagedControlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B7770CB9-9745-4FE5-9EDA-D14B01A12A50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceManagedControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceManagedControlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceManagedControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceManagedControlResponseBody) SetRequestId(v string) *UpdateServiceManagedControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceManagedControlResponseBody) Validate() error {
	return dara.Validate(s)
}
