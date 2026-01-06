// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListMultimodalSearchTaskResultRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *ListMultimodalSearchTaskResultRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMultimodalSearchTaskResultRequest
	GetPageSize() *int32
	SetTaskId(v string) *ListMultimodalSearchTaskResultRequest
	GetTaskId() *string
}

type ListMultimodalSearchTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ds-*****ab0
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListMultimodalSearchTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchTaskResultRequest) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchTaskResultRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListMultimodalSearchTaskResultRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMultimodalSearchTaskResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMultimodalSearchTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListMultimodalSearchTaskResultRequest) SetDBClusterId(v string) *ListMultimodalSearchTaskResultRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListMultimodalSearchTaskResultRequest) SetPageNumber(v int32) *ListMultimodalSearchTaskResultRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMultimodalSearchTaskResultRequest) SetPageSize(v int32) *ListMultimodalSearchTaskResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListMultimodalSearchTaskResultRequest) SetTaskId(v string) *ListMultimodalSearchTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *ListMultimodalSearchTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
