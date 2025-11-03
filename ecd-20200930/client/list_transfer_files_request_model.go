// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransferFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransferFilesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransferFilesRequest
	GetNextToken() *string
	SetTaskId(v string) *ListTransferFilesRequest
	GetTaskId() *string
}

type ListTransferFilesRequest struct {
	// The number of entries to return on each page.
	//
	// Maximum value: 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the transmission task.
	//
	// example:
	//
	// trt-03tdwg4tcuwdzv****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTransferFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransferFilesRequest) GoString() string {
	return s.String()
}

func (s *ListTransferFilesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransferFilesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransferFilesRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTransferFilesRequest) SetMaxResults(v int32) *ListTransferFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransferFilesRequest) SetNextToken(v string) *ListTransferFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransferFilesRequest) SetTaskId(v string) *ListTransferFilesRequest {
	s.TaskId = &v
	return s
}

func (s *ListTransferFilesRequest) Validate() error {
	return dara.Validate(s)
}
