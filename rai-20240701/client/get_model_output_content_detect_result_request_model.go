// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelOutputContentDetectResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetModelOutputContentDetectResultRequest
	GetRegionId() *string
	SetTaskId(v string) *GetModelOutputContentDetectResultRequest
	GetTaskId() *string
}

type GetModelOutputContentDetectResultRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetModelOutputContentDetectResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelOutputContentDetectResultRequest) GoString() string {
	return s.String()
}

func (s *GetModelOutputContentDetectResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetModelOutputContentDetectResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetModelOutputContentDetectResultRequest) SetRegionId(v string) *GetModelOutputContentDetectResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetModelOutputContentDetectResultRequest) SetTaskId(v string) *GetModelOutputContentDetectResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetModelOutputContentDetectResultRequest) Validate() error {
	return dara.Validate(s)
}
