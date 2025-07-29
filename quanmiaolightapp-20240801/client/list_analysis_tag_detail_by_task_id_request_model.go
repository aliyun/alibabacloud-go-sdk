// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnalysisTagDetailByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAnalysisTagDetailByTaskIdRequest
	GetNextToken() *string
	SetTaskId(v string) *ListAnalysisTagDetailByTaskIdRequest
	GetTaskId() *string
}

type ListAnalysisTagDetailByTaskIdRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// JlroP3CjgQh5PQDlH3ArzADkBTPZgVqo+64jhZRglNq0mEYoV5SlGb/Juvo8CdfYE9rlwEr2pIJQwdaYotak9g==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListAnalysisTagDetailByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetNextToken(v string) *ListAnalysisTagDetailByTaskIdRequest {
	s.NextToken = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetTaskId(v string) *ListAnalysisTagDetailByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
