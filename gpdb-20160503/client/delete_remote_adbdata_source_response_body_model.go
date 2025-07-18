// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRemoteADBDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRemoteADBDataSourceResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *DeleteRemoteADBDataSourceResponseBody
	GetTaskId() *int32
}

type DeleteRemoteADBDataSourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// e9d60eb1-e90d-4bc6-a470-c8b767460858
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 90000
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteRemoteADBDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemoteADBDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRemoteADBDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRemoteADBDataSourceResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteRemoteADBDataSourceResponseBody) SetRequestId(v string) *DeleteRemoteADBDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRemoteADBDataSourceResponseBody) SetTaskId(v int32) *DeleteRemoteADBDataSourceResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteRemoteADBDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
