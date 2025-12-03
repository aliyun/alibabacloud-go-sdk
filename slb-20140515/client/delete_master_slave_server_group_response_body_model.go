// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMasterSlaveServerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMasterSlaveServerGroupResponseBody
	GetRequestId() *string
}

type DeleteMasterSlaveServerGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMasterSlaveServerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMasterSlaveServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMasterSlaveServerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMasterSlaveServerGroupResponseBody) SetRequestId(v string) *DeleteMasterSlaveServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMasterSlaveServerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
