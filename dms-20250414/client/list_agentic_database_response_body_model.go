// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AgenticDatabase) *ListAgenticDatabaseResponseBody
	GetData() []*AgenticDatabase
	SetErrorCode(v string) *ListAgenticDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAgenticDatabaseResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListAgenticDatabaseResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgenticDatabaseResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgenticDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgenticDatabaseResponseBody
	GetSuccess() *bool
}

type ListAgenticDatabaseResponseBody struct {
	// The list of databases.
	Data []*AgenticDatabase `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned when a request exception occurs.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The maximum number of entries returned per page in the current query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. An empty value indicates that no more data is available. If the value is not empty, use it as the NextToken parameter in the next request to query the next page.
	//
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAgenticDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgenticDatabaseResponseBody) GetData() []*AgenticDatabase {
	return s.Data
}

func (s *ListAgenticDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAgenticDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAgenticDatabaseResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgenticDatabaseResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgenticDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgenticDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgenticDatabaseResponseBody) SetData(v []*AgenticDatabase) *ListAgenticDatabaseResponseBody {
	s.Data = v
	return s
}

func (s *ListAgenticDatabaseResponseBody) SetErrorCode(v string) *ListAgenticDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAgenticDatabaseResponseBody) SetErrorMessage(v string) *ListAgenticDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAgenticDatabaseResponseBody) SetMaxResults(v int32) *ListAgenticDatabaseResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgenticDatabaseResponseBody) SetNextToken(v string) *ListAgenticDatabaseResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgenticDatabaseResponseBody) SetRequestId(v string) *ListAgenticDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgenticDatabaseResponseBody) SetSuccess(v bool) *ListAgenticDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgenticDatabaseResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
