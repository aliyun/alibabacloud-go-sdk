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
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The information about the token.
	//
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queried partitions.
	PartitionList []*DLPartition `json:"PartitionList,omitempty" xml:"PartitionList,omitempty" type:"Repeated"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request succeeded.
	//
	// 	- **false**: The request failed.
	//
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
	if s.PartitionList != nil {
		for _, item := range s.PartitionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
