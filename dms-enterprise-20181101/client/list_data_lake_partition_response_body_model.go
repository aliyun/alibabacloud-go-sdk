// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataLakePartitionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakePartitionResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataLakePartitionResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakePartitionResponseBody
	GetNextToken() *string
	SetPartitionList(v []*DLPartition) *ListDataLakePartitionResponseBody
	GetPartitionList() []*DLPartition
	SetRequestId(v string) *ListDataLakePartitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakePartitionResponseBody
	GetSuccess() *bool
}

type ListDataLakePartitionResponseBody struct {
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
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataLakePartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakePartitionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakePartitionResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakePartitionResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakePartitionResponseBody) GetPartitionList() []*DLPartition {
	return s.PartitionList
}

func (s *ListDataLakePartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakePartitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakePartitionResponseBody) SetErrorCode(v string) *ListDataLakePartitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakePartitionResponseBody) SetErrorMessage(v string) *ListDataLakePartitionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakePartitionResponseBody) SetMaxResults(v int32) *ListDataLakePartitionResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakePartitionResponseBody) SetNextToken(v string) *ListDataLakePartitionResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataLakePartitionResponseBody) SetPartitionList(v []*DLPartition) *ListDataLakePartitionResponseBody {
	s.PartitionList = v
	return s
}

func (s *ListDataLakePartitionResponseBody) SetRequestId(v string) *ListDataLakePartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakePartitionResponseBody) SetSuccess(v bool) *ListDataLakePartitionResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakePartitionResponseBody) Validate() error {
	return dara.Validate(s)
}
