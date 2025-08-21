// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCategoryResponseBody
	GetRequestId() *string
}

type UpdateCategoryResponseBody struct {
	// example:
	//
	// F79E7305-5314-5069-A701-9591AD051902
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCategoryResponseBody) SetRequestId(v string) *UpdateCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
