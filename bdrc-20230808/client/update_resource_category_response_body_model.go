// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceCategoryResponseBody
	GetRequestId() *string
}

type UpdateResourceCategoryResponseBody struct {
	// The unique ID of the request.
	//
	// example:
	//
	// 700683DE-0154-56D4-8D76-3B7A2C2C7DF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceCategoryResponseBody) SetRequestId(v string) *UpdateResourceCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
