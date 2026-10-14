// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticTableBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AgenticTableBaseInfo) *ListAgenticTableBaseInfoResponseBody
	GetData() []*AgenticTableBaseInfo
	SetErrorCode(v string) *ListAgenticTableBaseInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAgenticTableBaseInfoResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListAgenticTableBaseInfoResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgenticTableBaseInfoResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgenticTableBaseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgenticTableBaseInfoResponseBody
	GetSuccess() *bool
}

type ListAgenticTableBaseInfoResponseBody struct {
	// The list of basic table information.
	Data []*AgenticTableBaseInfo `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. This parameter is returned when an exception occurs.
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
	// The token for the next query. If this parameter is empty, no more data is available. If this parameter is not empty, set this parameter to the NextToken value in the next request to query the next page.
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

func (s ListAgenticTableBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticTableBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgenticTableBaseInfoResponseBody) GetData() []*AgenticTableBaseInfo {
	return s.Data
}

func (s *ListAgenticTableBaseInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAgenticTableBaseInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAgenticTableBaseInfoResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgenticTableBaseInfoResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgenticTableBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgenticTableBaseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgenticTableBaseInfoResponseBody) SetData(v []*AgenticTableBaseInfo) *ListAgenticTableBaseInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListAgenticTableBaseInfoResponseBody) SetErrorCode(v string) *ListAgenticTableBaseInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAgenticTableBaseInfoResponseBody) SetErrorMessage(v string) *ListAgenticTableBaseInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAgenticTableBaseInfoResponseBody) SetMaxResults(v int32) *ListAgenticTableBaseInfoResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgenticTableBaseInfoResponseBody) SetNextToken(v string) *ListAgenticTableBaseInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgenticTableBaseInfoResponseBody) SetRequestId(v string) *ListAgenticTableBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgenticTableBaseInfoResponseBody) SetSuccess(v bool) *ListAgenticTableBaseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgenticTableBaseInfoResponseBody) Validate() error {
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
