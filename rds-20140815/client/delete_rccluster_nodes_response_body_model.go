// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCClusterNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRCClusterNodesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DeleteRCClusterNodesResponseBody
	GetTaskId() *string
}

type DeleteRCClusterNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7E0970A1-0434-5C83-B560-613EBA11B525
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 238028563
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteRCClusterNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRCClusterNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRCClusterNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteRCClusterNodesResponseBody) SetRequestId(v string) *DeleteRCClusterNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRCClusterNodesResponseBody) SetTaskId(v string) *DeleteRCClusterNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteRCClusterNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
