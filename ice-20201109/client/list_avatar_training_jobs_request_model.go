// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarTrainingJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListAvatarTrainingJobsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListAvatarTrainingJobsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListAvatarTrainingJobsRequest
	GetStatus() *string
}

type ListAvatarTrainingJobsRequest struct {
	// 	- The page number.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 	- The number of entries per page.
	//
	// 	- Default value: 10.
	//
	// 	- Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 	- The job state.
	//
	// 	- Valid values: Init, Queuing, Training, Success, and Fail.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAvatarTrainingJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarTrainingJobsRequest) GoString() string {
	return s.String()
}

func (s *ListAvatarTrainingJobsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListAvatarTrainingJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAvatarTrainingJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAvatarTrainingJobsRequest) SetPageNo(v int32) *ListAvatarTrainingJobsRequest {
	s.PageNo = &v
	return s
}

func (s *ListAvatarTrainingJobsRequest) SetPageSize(v int32) *ListAvatarTrainingJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAvatarTrainingJobsRequest) SetStatus(v string) *ListAvatarTrainingJobsRequest {
	s.Status = &v
	return s
}

func (s *ListAvatarTrainingJobsRequest) Validate() error {
	return dara.Validate(s)
}
