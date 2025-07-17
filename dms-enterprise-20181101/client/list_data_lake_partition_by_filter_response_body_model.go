// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionByFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataLakePartitionByFilterResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakePartitionByFilterResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataLakePartitionByFilterResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakePartitionByFilterResponseBody
	GetNextToken() *string
	SetPartitionList(v []*DLPartition) *ListDataLakePartitionByFilterResponseBody
	GetPartitionList() []*DLPartition
	SetRequestId(v string) *ListDataLakePartitionByFilterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakePartitionByFilterResponseBody
	GetSuccess() *bool
}

type ListDataLakePartitionByFilterResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken     *string        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PartitionList []*DLPartition `json:"PartitionList,omitempty" xml:"PartitionList,omitempty" type:"Repeated"`
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataLakePartitionByFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionByFilterResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionByFilterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakePartitionByFilterResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakePartitionByFilterResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakePartitionByFilterResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakePartitionByFilterResponseBody) GetPartitionList() []*DLPartition {
	return s.PartitionList
}

func (s *ListDataLakePartitionByFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakePartitionByFilterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakePartitionByFilterResponseBody) SetErrorCode(v string) *ListDataLakePartitionByFilterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakePartitionByFilterResponseBody) SetErrorMessage(v string) *ListDataLakePartitionByFilterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakePartitionByFilterResponseBody) SetMaxResults(v int32) *ListDataLakePartitionByFilterResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakePartitionByFilterResponseBody) SetNextToken(v string) *ListDataLakePartitionByFilterResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataLakePartitionByFilterResponseBody) SetPartitionList(v []*DLPartition) *ListDataLakePartitionByFilterResponseBody {
	s.PartitionList = v
	return s
}

func (s *ListDataLakePartitionByFilterResponseBody) SetRequestId(v string) *ListDataLakePartitionByFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakePartitionByFilterResponseBody) SetSuccess(v bool) *ListDataLakePartitionByFilterResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakePartitionByFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
