// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePackageStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstancePackageStateResponseBody
	GetRequestId() *string
}

type UpdateInstancePackageStateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2597E94B-5346-42D1-BB58-XXXXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstancePackageStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePackageStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstancePackageStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstancePackageStateResponseBody) SetRequestId(v string) *UpdateInstancePackageStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstancePackageStateResponseBody) Validate() error {
	return dara.Validate(s)
}
