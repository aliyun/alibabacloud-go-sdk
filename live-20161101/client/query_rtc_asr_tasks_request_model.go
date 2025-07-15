// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRtcAsrTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryRtcAsrTasksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryRtcAsrTasksRequest
	GetRegionId() *string
	SetTaskId(v string) *QueryRtcAsrTasksRequest
	GetTaskId() *string
}

type QueryRtcAsrTasksRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the task that you want to query. If you do not specify this parameter, all running tasks under your UID are queried.
	//
	// example:
	//
	// asr-a6ac15e0-9118-4b4c-9e64-306163a0****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryRtcAsrTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRtcAsrTasksRequest) GoString() string {
	return s.String()
}

func (s *QueryRtcAsrTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryRtcAsrTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryRtcAsrTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryRtcAsrTasksRequest) SetOwnerId(v int64) *QueryRtcAsrTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRtcAsrTasksRequest) SetRegionId(v string) *QueryRtcAsrTasksRequest {
	s.RegionId = &v
	return s
}

func (s *QueryRtcAsrTasksRequest) SetTaskId(v string) *QueryRtcAsrTasksRequest {
	s.TaskId = &v
	return s
}

func (s *QueryRtcAsrTasksRequest) Validate() error {
	return dara.Validate(s)
}
