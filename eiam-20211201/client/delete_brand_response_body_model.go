// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBrandResponseBody
	GetRequestId() *string
}

type DeleteBrandResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBrandResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBrandResponseBody) SetRequestId(v string) *DeleteBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBrandResponseBody) Validate() error {
	return dara.Validate(s)
}
