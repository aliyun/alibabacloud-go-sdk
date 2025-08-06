// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppInfoResponseBody
	GetRequestId() *string
}

type DeleteAppInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppInfoResponseBody) SetRequestId(v string) *DeleteAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
