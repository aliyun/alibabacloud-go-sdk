// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDebugDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v string) *GetJobDebugDataRequest
	GetCursor() *string
	SetEndTime(v int64) *GetJobDebugDataRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetJobDebugDataRequest
	GetInstanceId() *string
	SetJobName(v string) *GetJobDebugDataRequest
	GetJobName() *string
	SetLimit(v int32) *GetJobDebugDataRequest
	GetLimit() *int32
	SetRegionId(v string) *GetJobDebugDataRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetJobDebugDataRequest
	GetStartTime() *int64
}

type GetJobDebugDataRequest struct {
	Cursor  *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	EndTime *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Limit   *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// This parameter is required.
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetJobDebugDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobDebugDataRequest) GoString() string {
	return s.String()
}

func (s *GetJobDebugDataRequest) GetCursor() *string {
	return s.Cursor
}

func (s *GetJobDebugDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetJobDebugDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetJobDebugDataRequest) GetJobName() *string {
	return s.JobName
}

func (s *GetJobDebugDataRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *GetJobDebugDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetJobDebugDataRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetJobDebugDataRequest) SetCursor(v string) *GetJobDebugDataRequest {
	s.Cursor = &v
	return s
}

func (s *GetJobDebugDataRequest) SetEndTime(v int64) *GetJobDebugDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetJobDebugDataRequest) SetInstanceId(v string) *GetJobDebugDataRequest {
	s.InstanceId = &v
	return s
}

func (s *GetJobDebugDataRequest) SetJobName(v string) *GetJobDebugDataRequest {
	s.JobName = &v
	return s
}

func (s *GetJobDebugDataRequest) SetLimit(v int32) *GetJobDebugDataRequest {
	s.Limit = &v
	return s
}

func (s *GetJobDebugDataRequest) SetRegionId(v string) *GetJobDebugDataRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobDebugDataRequest) SetStartTime(v int64) *GetJobDebugDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetJobDebugDataRequest) Validate() error {
	return dara.Validate(s)
}
