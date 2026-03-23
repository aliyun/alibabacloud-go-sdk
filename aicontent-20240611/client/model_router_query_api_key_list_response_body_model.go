// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryApiKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryApiKeyListResponseBodyData) *ModelRouterQueryApiKeyListResponseBody
	GetData() *ModelRouterQueryApiKeyListResponseBodyData
	SetErrCode(v string) *ModelRouterQueryApiKeyListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryApiKeyListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryApiKeyListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryApiKeyListResponseBody
	GetMaxResults() *int32
	SetRequestId(v string) *ModelRouterQueryApiKeyListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryApiKeyListResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryApiKeyListResponseBody struct {
	// example:
	//
	// []
	Data *ModelRouterQueryApiKeyListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// maxResults
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryApiKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryApiKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetData() *ModelRouterQueryApiKeyListResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetData(v *ModelRouterQueryApiKeyListResponseBodyData) *ModelRouterQueryApiKeyListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetErrCode(v string) *ModelRouterQueryApiKeyListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetErrMessage(v string) *ModelRouterQueryApiKeyListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryApiKeyListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetMaxResults(v int32) *ModelRouterQueryApiKeyListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetRequestId(v string) *ModelRouterQueryApiKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetSuccess(v bool) *ModelRouterQueryApiKeyListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryApiKeyListResponseBodyData struct {
	List      []*ApiKeyDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	MaxResult *string      `json:"maxResult,omitempty" xml:"maxResult,omitempty"`
	// example:
	//
	// None
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 0
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryApiKeyListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryApiKeyListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) GetList() []*ApiKeyDTO {
	return s.List
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) GetMaxResult() *string {
	return s.MaxResult
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) SetList(v []*ApiKeyDTO) *ModelRouterQueryApiKeyListResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) SetMaxResult(v string) *ModelRouterQueryApiKeyListResponseBodyData {
	s.MaxResult = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) SetNextToken(v string) *ModelRouterQueryApiKeyListResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) SetPage(v int32) *ModelRouterQueryApiKeyListResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) SetPageSize(v int32) *ModelRouterQueryApiKeyListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) SetTotal(v int32) *ModelRouterQueryApiKeyListResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBodyData) Validate() error {
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
