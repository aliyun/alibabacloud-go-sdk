// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelInputContentDetectResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetModelInputContentDetectResultRequest
	GetRegionId() *string
	SetTaskId(v string) *GetModelInputContentDetectResultRequest
	GetTaskId() *string
}

type GetModelInputContentDetectResultRequest struct {
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

func (s GetModelInputContentDetectResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelInputContentDetectResultRequest) GoString() string {
	return s.String()
}

func (s *GetModelInputContentDetectResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetModelInputContentDetectResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetModelInputContentDetectResultRequest) SetRegionId(v string) *GetModelInputContentDetectResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetModelInputContentDetectResultRequest) SetTaskId(v string) *GetModelInputContentDetectResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetModelInputContentDetectResultRequest) Validate() error {
	return dara.Validate(s)
}
