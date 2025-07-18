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
	// ABB39CC3-4488-4857-905D-2E4A051D0521
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
