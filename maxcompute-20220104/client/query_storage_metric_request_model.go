// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStorageMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectList(v []*string) *QueryStorageMetricRequest
	GetProjectList() []*string
	SetTypeList(v []*string) *QueryStorageMetricRequest
	GetTypeList() []*string
	SetEndTime(v int64) *QueryStorageMetricRequest
	GetEndTime() *int64
	SetStartTime(v int64) *QueryStorageMetricRequest
	GetStartTime() *int64
}

type QueryStorageMetricRequest struct {
	ProjectList []*string `json:"projectList,omitempty" xml:"projectList,omitempty" type:"Repeated"`
	TypeList    []*string `json:"typeList,omitempty" xml:"typeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1735536322
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1735534322
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s QueryStorageMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryStorageMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryStorageMetricRequest) GetProjectList() []*string {
	return s.ProjectList
}

func (s *QueryStorageMetricRequest) GetTypeList() []*string {
	return s.TypeList
}

func (s *QueryStorageMetricRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryStorageMetricRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryStorageMetricRequest) SetProjectList(v []*string) *QueryStorageMetricRequest {
	s.ProjectList = v
	return s
}

func (s *QueryStorageMetricRequest) SetTypeList(v []*string) *QueryStorageMetricRequest {
	s.TypeList = v
	return s
}

func (s *QueryStorageMetricRequest) SetEndTime(v int64) *QueryStorageMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryStorageMetricRequest) SetStartTime(v int64) *QueryStorageMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryStorageMetricRequest) Validate() error {
	return dara.Validate(s)
}
