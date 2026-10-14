// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAgenticCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*AgenticCatalog) *SearchAgenticCatalogResponseBody
	GetData() []*AgenticCatalog
	SetErrorCode(v string) *SearchAgenticCatalogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchAgenticCatalogResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *SearchAgenticCatalogResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchAgenticCatalogResponseBody
	GetNextToken() *string
	SetRequestId(v string) *SearchAgenticCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchAgenticCatalogResponseBody
	GetSuccess() *bool
}

type SearchAgenticCatalogResponseBody struct {
	// The list of catalogs.
	Data []*AgenticCatalog `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code, which is returned when a request exception occurs.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message, which is returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The maximum number of entries returned per page in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. An empty value indicates that no more data is available. If the value is not empty, set the NextToken parameter to this value in the next request to query the next page.
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

func (s SearchAgenticCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAgenticCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAgenticCatalogResponseBody) GetData() []*AgenticCatalog {
	return s.Data
}

func (s *SearchAgenticCatalogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchAgenticCatalogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchAgenticCatalogResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchAgenticCatalogResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchAgenticCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchAgenticCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchAgenticCatalogResponseBody) SetData(v []*AgenticCatalog) *SearchAgenticCatalogResponseBody {
	s.Data = v
	return s
}

func (s *SearchAgenticCatalogResponseBody) SetErrorCode(v string) *SearchAgenticCatalogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchAgenticCatalogResponseBody) SetErrorMessage(v string) *SearchAgenticCatalogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchAgenticCatalogResponseBody) SetMaxResults(v int32) *SearchAgenticCatalogResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchAgenticCatalogResponseBody) SetNextToken(v string) *SearchAgenticCatalogResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchAgenticCatalogResponseBody) SetRequestId(v string) *SearchAgenticCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAgenticCatalogResponseBody) SetSuccess(v bool) *SearchAgenticCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *SearchAgenticCatalogResponseBody) Validate() error {
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
