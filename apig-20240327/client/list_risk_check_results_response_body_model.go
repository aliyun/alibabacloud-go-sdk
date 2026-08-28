// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRiskCheckResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRiskCheckResultsResponseBody
	GetCode() *string
	SetData(v *ListRiskCheckResultsResponseBodyData) *ListRiskCheckResultsResponseBody
	GetData() *ListRiskCheckResultsResponseBodyData
	SetMaxResults(v int32) *ListRiskCheckResultsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListRiskCheckResultsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListRiskCheckResultsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRiskCheckResultsResponseBody
	GetRequestId() *string
}

type ListRiskCheckResultsResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListRiskCheckResultsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// token-xxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BF76AA7C-2C1E-5C3F-B366-5EC07F9662DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRiskCheckResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRiskCheckResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRiskCheckResultsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRiskCheckResultsResponseBody) GetData() *ListRiskCheckResultsResponseBodyData {
	return s.Data
}

func (s *ListRiskCheckResultsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRiskCheckResultsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRiskCheckResultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRiskCheckResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRiskCheckResultsResponseBody) SetCode(v string) *ListRiskCheckResultsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRiskCheckResultsResponseBody) SetData(v *ListRiskCheckResultsResponseBodyData) *ListRiskCheckResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListRiskCheckResultsResponseBody) SetMaxResults(v int32) *ListRiskCheckResultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRiskCheckResultsResponseBody) SetMessage(v string) *ListRiskCheckResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRiskCheckResultsResponseBody) SetNextToken(v string) *ListRiskCheckResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRiskCheckResultsResponseBody) SetRequestId(v string) *ListRiskCheckResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRiskCheckResultsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRiskCheckResultsResponseBodyData struct {
	Items []*RiskCheckResults `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListRiskCheckResultsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRiskCheckResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRiskCheckResultsResponseBodyData) GetItems() []*RiskCheckResults {
	return s.Items
}

func (s *ListRiskCheckResultsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRiskCheckResultsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRiskCheckResultsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListRiskCheckResultsResponseBodyData) SetItems(v []*RiskCheckResults) *ListRiskCheckResultsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListRiskCheckResultsResponseBodyData) SetPageNumber(v int32) *ListRiskCheckResultsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRiskCheckResultsResponseBodyData) SetPageSize(v int32) *ListRiskCheckResultsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRiskCheckResultsResponseBodyData) SetTotalSize(v int32) *ListRiskCheckResultsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListRiskCheckResultsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
