// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataLakeTableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakeTableResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataLakeTableResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeTableResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDataLakeTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakeTableResponseBody
	GetSuccess() *bool
	SetTableList(v []*DLTable) *ListDataLakeTableResponseBody
	GetTableList() []*DLTable
}

type ListDataLakeTableResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The number of records per page.
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
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 3D3FB827-E667-50DB-AD59-C83F8237FECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the tables.
	TableList []*DLTable `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Repeated"`
}

func (s ListDataLakeTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeTableResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeTableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakeTableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakeTableResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeTableResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakeTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakeTableResponseBody) GetTableList() []*DLTable {
	return s.TableList
}

func (s *ListDataLakeTableResponseBody) SetErrorCode(v string) *ListDataLakeTableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakeTableResponseBody) SetErrorMessage(v string) *ListDataLakeTableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakeTableResponseBody) SetMaxResults(v int32) *ListDataLakeTableResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeTableResponseBody) SetNextToken(v string) *ListDataLakeTableResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeTableResponseBody) SetRequestId(v string) *ListDataLakeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakeTableResponseBody) SetSuccess(v bool) *ListDataLakeTableResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakeTableResponseBody) SetTableList(v []*DLTable) *ListDataLakeTableResponseBody {
	s.TableList = v
	return s
}

func (s *ListDataLakeTableResponseBody) Validate() error {
	if s.TableList != nil {
		for _, item := range s.TableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
