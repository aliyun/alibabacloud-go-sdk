// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMediaCategoryResponseBody
	GetRequestId() *string
}

type UpdateMediaCategoryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E3931857-E3D3-4D6E-9C7B-D2C09441BD01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaCategoryResponseBody) SetRequestId(v string) *UpdateMediaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
