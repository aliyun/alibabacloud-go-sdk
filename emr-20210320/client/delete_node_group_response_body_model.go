// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNodeGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNodeGroupResponseBody
	GetSuccess() *bool
}

type DeleteNodeGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 64EBE3F5-0498-1578-BEC0-6ACE364E912D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation to release the node group was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNodeGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNodeGroupResponseBody) SetRequestId(v string) *DeleteNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeGroupResponseBody) SetSuccess(v bool) *DeleteNodeGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
