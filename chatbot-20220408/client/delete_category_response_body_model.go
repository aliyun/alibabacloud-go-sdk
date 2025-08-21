// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCategoryResponseBody
	GetRequestId() *string
}

type DeleteCategoryResponseBody struct {
	// example:
	//
	// 4B3E0DE3-DC57-5BFE-88D4-ADD8ED024F55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCategoryResponseBody) SetRequestId(v string) *DeleteCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
