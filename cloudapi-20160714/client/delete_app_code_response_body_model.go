// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppCodeResponseBody
	GetRequestId() *string
}

type DeleteAppCodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E8515BA6-81CD-4191-A7CF-C4FCDD3C0D99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppCodeResponseBody) SetRequestId(v string) *DeleteAppCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
