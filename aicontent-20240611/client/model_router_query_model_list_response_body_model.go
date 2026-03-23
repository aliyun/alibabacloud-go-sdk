// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryModelListResponseBodyData) *ModelRouterQueryModelListResponseBody
	GetData() *ModelRouterQueryModelListResponseBodyData
	SetErrCode(v string) *ModelRouterQueryModelListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryModelListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryModelListResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryModelListResponseBody struct {
	// example:
	//
	// []
	Data *ModelRouterQueryModelListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryModelListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelListResponseBody) GetData() *ModelRouterQueryModelListResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryModelListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelListResponseBody) SetData(v *ModelRouterQueryModelListResponseBodyData) *ModelRouterQueryModelListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetErrCode(v string) *ModelRouterQueryModelListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetErrMessage(v string) *ModelRouterQueryModelListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetRequestId(v string) *ModelRouterQueryModelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetSuccess(v bool) *ModelRouterQueryModelListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryModelListResponseBodyData struct {
	List []*ModelDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// None
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
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
	// 1
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 5
	Total *string `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryModelListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelListResponseBodyData) GetList() []*ModelDTO {
	return s.List
}

func (s *ModelRouterQueryModelListResponseBodyData) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ModelRouterQueryModelListResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelListResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryModelListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelListResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *ModelRouterQueryModelListResponseBodyData) SetList(v []*ModelDTO) *ModelRouterQueryModelListResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryModelListResponseBodyData) SetMaxResults(v string) *ModelRouterQueryModelListResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBodyData) SetNextToken(v string) *ModelRouterQueryModelListResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBodyData) SetPage(v int32) *ModelRouterQueryModelListResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBodyData) SetPageSize(v int32) *ModelRouterQueryModelListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBodyData) SetTotal(v string) *ModelRouterQueryModelListResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBodyData) Validate() error {
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
