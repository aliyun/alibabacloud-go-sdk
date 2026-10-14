// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAgenticDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AgenticDatabase) *SearchAgenticDatabaseResponseBody
	GetData() []*AgenticDatabase
	SetErrorCode(v string) *SearchAgenticDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchAgenticDatabaseResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *SearchAgenticDatabaseResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchAgenticDatabaseResponseBody
	GetNextToken() *string
	SetRequestId(v string) *SearchAgenticDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchAgenticDatabaseResponseBody
	GetSuccess() *bool
}

type SearchAgenticDatabaseResponseBody struct {
	// The list of databases.
	Data []*AgenticDatabase `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. This parameter is returned when a request exception occurs.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. This parameter is returned when the call fails.
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
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchAgenticDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAgenticDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAgenticDatabaseResponseBody) GetData() []*AgenticDatabase {
	return s.Data
}

func (s *SearchAgenticDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchAgenticDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchAgenticDatabaseResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchAgenticDatabaseResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchAgenticDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchAgenticDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchAgenticDatabaseResponseBody) SetData(v []*AgenticDatabase) *SearchAgenticDatabaseResponseBody {
	s.Data = v
	return s
}

func (s *SearchAgenticDatabaseResponseBody) SetErrorCode(v string) *SearchAgenticDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchAgenticDatabaseResponseBody) SetErrorMessage(v string) *SearchAgenticDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchAgenticDatabaseResponseBody) SetMaxResults(v int32) *SearchAgenticDatabaseResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchAgenticDatabaseResponseBody) SetNextToken(v string) *SearchAgenticDatabaseResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchAgenticDatabaseResponseBody) SetRequestId(v string) *SearchAgenticDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAgenticDatabaseResponseBody) SetSuccess(v bool) *SearchAgenticDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *SearchAgenticDatabaseResponseBody) Validate() error {
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
