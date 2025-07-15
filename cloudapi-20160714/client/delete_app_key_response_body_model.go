// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppKeyResponseBody
	GetRequestId() *string
}

type DeleteAppKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 79EF055D-AC00-5161-8F35-6A36AAED7422
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppKeyResponseBody) SetRequestId(v string) *DeleteAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
