// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVServerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVServerGroupResponseBody
	GetRequestId() *string
}

type DeleteVServerGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVServerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVServerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVServerGroupResponseBody) SetRequestId(v string) *DeleteVServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVServerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
