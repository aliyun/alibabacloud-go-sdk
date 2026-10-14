// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AgenticCatalog) *ListAgenticCatalogResponseBody
	GetData() []*AgenticCatalog
	SetErrorCode(v string) *ListAgenticCatalogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAgenticCatalogResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListAgenticCatalogResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgenticCatalogResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgenticCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgenticCatalogResponseBody
	GetSuccess() *bool
}

type ListAgenticCatalogResponseBody struct {
	// The list of catalogs.
	Data []*AgenticCatalog `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// The token for the next query. An empty value indicates that no more data is available. If the value is not empty, use this value as the NextToken parameter in the next request to query the next page.
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

func (s ListAgenticCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgenticCatalogResponseBody) GetData() []*AgenticCatalog {
	return s.Data
}

func (s *ListAgenticCatalogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAgenticCatalogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAgenticCatalogResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgenticCatalogResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgenticCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgenticCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgenticCatalogResponseBody) SetData(v []*AgenticCatalog) *ListAgenticCatalogResponseBody {
	s.Data = v
	return s
}

func (s *ListAgenticCatalogResponseBody) SetErrorCode(v string) *ListAgenticCatalogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAgenticCatalogResponseBody) SetErrorMessage(v string) *ListAgenticCatalogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAgenticCatalogResponseBody) SetMaxResults(v int32) *ListAgenticCatalogResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgenticCatalogResponseBody) SetNextToken(v string) *ListAgenticCatalogResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgenticCatalogResponseBody) SetRequestId(v string) *ListAgenticCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgenticCatalogResponseBody) SetSuccess(v bool) *ListAgenticCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgenticCatalogResponseBody) Validate() error {
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
