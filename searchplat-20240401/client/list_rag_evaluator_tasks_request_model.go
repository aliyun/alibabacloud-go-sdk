// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRagEvaluatorTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *ListRagEvaluatorTasksRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListRagEvaluatorTasksRequest
	GetPageSize() *string
}

type ListRagEvaluatorTasksRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListRagEvaluatorTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRagEvaluatorTasksRequest) GoString() string {
	return s.String()
}

func (s *ListRagEvaluatorTasksRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListRagEvaluatorTasksRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListRagEvaluatorTasksRequest) SetPageNumber(v string) *ListRagEvaluatorTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRagEvaluatorTasksRequest) SetPageSize(v string) *ListRagEvaluatorTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListRagEvaluatorTasksRequest) Validate() error {
	return dara.Validate(s)
}
