// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteACLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteACLResponseBody
	GetRequestId() *string
}

type DeleteACLResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 00546174-2CE6-4587-9550-04B6A3313938
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteACLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteACLResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteACLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteACLResponseBody) SetRequestId(v string) *DeleteACLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteACLResponseBody) Validate() error {
	return dara.Validate(s)
}
