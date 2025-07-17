// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataLakePartitionNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakePartitionNameResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataLakePartitionNameResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakePartitionNameResponseBody
	GetNextToken() *string
	SetPartitionNameList(v []*string) *ListDataLakePartitionNameResponseBody
	GetPartitionNameList() []*string
	SetRequestId(v string) *ListDataLakePartitionNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakePartitionNameResponseBody
	GetSuccess() *bool
}

type ListDataLakePartitionNameResponseBody struct {
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
	NextToken         *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PartitionNameList []*string `json:"PartitionNameList,omitempty" xml:"PartitionNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataLakePartitionNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionNameResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakePartitionNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakePartitionNameResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakePartitionNameResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakePartitionNameResponseBody) GetPartitionNameList() []*string {
	return s.PartitionNameList
}

func (s *ListDataLakePartitionNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakePartitionNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakePartitionNameResponseBody) SetErrorCode(v string) *ListDataLakePartitionNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakePartitionNameResponseBody) SetErrorMessage(v string) *ListDataLakePartitionNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakePartitionNameResponseBody) SetMaxResults(v int32) *ListDataLakePartitionNameResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakePartitionNameResponseBody) SetNextToken(v string) *ListDataLakePartitionNameResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataLakePartitionNameResponseBody) SetPartitionNameList(v []*string) *ListDataLakePartitionNameResponseBody {
	s.PartitionNameList = v
	return s
}

func (s *ListDataLakePartitionNameResponseBody) SetRequestId(v string) *ListDataLakePartitionNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakePartitionNameResponseBody) SetSuccess(v bool) *ListDataLakePartitionNameResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakePartitionNameResponseBody) Validate() error {
	return dara.Validate(s)
}
