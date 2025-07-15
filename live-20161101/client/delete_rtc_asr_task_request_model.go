// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtcAsrTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteRtcAsrTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteRtcAsrTaskRequest
	GetRegionId() *string
	SetTaskId(v string) *DeleteRtcAsrTaskRequest
	GetTaskId() *string
}

type DeleteRtcAsrTaskRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID. You can obtain the ID from the response to the [CreateRtcAsrTask](https://help.aliyun.com/document_detail/2848217.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// asr-51c72******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteRtcAsrTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtcAsrTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteRtcAsrTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRtcAsrTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRtcAsrTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteRtcAsrTaskRequest) SetOwnerId(v int64) *DeleteRtcAsrTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRtcAsrTaskRequest) SetRegionId(v string) *DeleteRtcAsrTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRtcAsrTaskRequest) SetTaskId(v string) *DeleteRtcAsrTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteRtcAsrTaskRequest) Validate() error {
	return dara.Validate(s)
}
