// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBResourceGroupResponseBody
	GetRequestId() *string
}

type DeleteDBResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A94B6C02-7BD4-5D67-9776-3AC8317E8DD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBResourceGroupResponseBody) SetRequestId(v string) *DeleteDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
