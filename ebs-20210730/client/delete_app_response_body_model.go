// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppResponseBody
	GetRequestId() *string
}

type DeleteAppResponseBody struct {
	// example:
	//
	// 1F58EB6E-7B7A-5437-9612-F14D50822A61
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) Validate() error {
	return dara.Validate(s)
}
