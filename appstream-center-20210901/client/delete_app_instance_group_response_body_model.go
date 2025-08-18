// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstanceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppInstanceGroupResponseBody
	GetRequestId() *string
}

type DeleteAppInstanceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppInstanceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppInstanceGroupResponseBody) SetRequestId(v string) *DeleteAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppInstanceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
