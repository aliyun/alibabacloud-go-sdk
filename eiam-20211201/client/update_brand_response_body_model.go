// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateBrandResponseBody
	GetRequestId() *string
}

type UpdateBrandResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBrandResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBrandResponseBody) SetRequestId(v string) *UpdateBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBrandResponseBody) Validate() error {
	return dara.Validate(s)
}
