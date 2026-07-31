// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListBalanceOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterListBalanceOrdersResponseBodyData) *ModelRouterListBalanceOrdersResponseBody
	GetData() *ModelRouterListBalanceOrdersResponseBodyData
	SetErrCode(v string) *ModelRouterListBalanceOrdersResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterListBalanceOrdersResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterListBalanceOrdersResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterListBalanceOrdersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterListBalanceOrdersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterListBalanceOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterListBalanceOrdersResponseBody
	GetSuccess() *bool
}

type ModelRouterListBalanceOrdersResponseBody struct {
	// The data object.
	Data *ModelRouterListBalanceOrdersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault error message encoding.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The maximum number of results.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterListBalanceOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListBalanceOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterListBalanceOrdersResponseBody) GetData() *ModelRouterListBalanceOrdersResponseBodyData {
	return s.Data
}

func (s *ModelRouterListBalanceOrdersResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterListBalanceOrdersResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterListBalanceOrdersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterListBalanceOrdersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterListBalanceOrdersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterListBalanceOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterListBalanceOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterListBalanceOrdersResponseBody) SetData(v *ModelRouterListBalanceOrdersResponseBodyData) *ModelRouterListBalanceOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBody) SetErrCode(v string) *ModelRouterListBalanceOrdersResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBody) SetErrMessage(v string) *ModelRouterListBalanceOrdersResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBody) SetHttpStatusCode(v int32) *ModelRouterListBalanceOrdersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBody) SetMaxResults(v int32) *ModelRouterListBalanceOrdersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBody) SetNextToken(v string) *ModelRouterListBalanceOrdersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBody) SetRequestId(v string) *ModelRouterListBalanceOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBody) SetSuccess(v bool) *ModelRouterListBalanceOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterListBalanceOrdersResponseBodyData struct {
	// The balance change log entries.
	List []*BalanceOrderDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The requested page.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 5
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterListBalanceOrdersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListBalanceOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterListBalanceOrdersResponseBodyData) GetList() []*BalanceOrderDTO {
	return s.List
}

func (s *ModelRouterListBalanceOrdersResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterListBalanceOrdersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterListBalanceOrdersResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterListBalanceOrdersResponseBodyData) SetList(v []*BalanceOrderDTO) *ModelRouterListBalanceOrdersResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBodyData) SetPage(v int32) *ModelRouterListBalanceOrdersResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBodyData) SetPageSize(v int32) *ModelRouterListBalanceOrdersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBodyData) SetTotal(v int32) *ModelRouterListBalanceOrdersResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterListBalanceOrdersResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
