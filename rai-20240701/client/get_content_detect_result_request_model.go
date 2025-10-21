// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContentDetectResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetContentDetectResultRequest
	GetRegionId() *string
	SetTaskId(v string) *GetContentDetectResultRequest
	GetTaskId() *string
}

type GetContentDetectResultRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetContentDetectResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContentDetectResultRequest) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetContentDetectResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetContentDetectResultRequest) SetRegionId(v string) *GetContentDetectResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetContentDetectResultRequest) SetTaskId(v string) *GetContentDetectResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetContentDetectResultRequest) Validate() error {
	return dara.Validate(s)
}
