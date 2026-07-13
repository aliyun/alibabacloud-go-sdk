// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceCategoryResponseBody
	GetRequestId() *string
}

type DeleteResourceCategoryResponseBody struct {
	// The unique identifier of the request.
	//
	// example:
	//
	// 8724BC18-904D-5A0D-BFF4-F0554F0037E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceCategoryResponseBody) SetRequestId(v string) *DeleteResourceCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
