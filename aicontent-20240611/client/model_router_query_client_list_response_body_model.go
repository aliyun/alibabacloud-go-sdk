// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryClientListResponseBodyData) *ModelRouterQueryClientListResponseBody
	GetData() *ModelRouterQueryClientListResponseBodyData
	SetErrCode(v string) *ModelRouterQueryClientListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryClientListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryClientListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryClientListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryClientListResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryClientListResponseBody struct {
	// example:
	//
	// []
	Data *ModelRouterQueryClientListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ModelRouterQueryClientListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientListResponseBody) GetData() *ModelRouterQueryClientListResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryClientListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryClientListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryClientListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryClientListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryClientListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryClientListResponseBody) SetData(v *ModelRouterQueryClientListResponseBodyData) *ModelRouterQueryClientListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetErrCode(v string) *ModelRouterQueryClientListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetErrMessage(v string) *ModelRouterQueryClientListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryClientListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetRequestId(v string) *ModelRouterQueryClientListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetSuccess(v bool) *ModelRouterQueryClientListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryClientListResponseBodyData struct {
	List []*ClientDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Page      *int32  `json:"page,omitempty" xml:"page,omitempty"`
	PageSize  *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total     *int32  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryClientListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientListResponseBodyData) GetList() []*ClientDTO {
	return s.List
}

func (s *ModelRouterQueryClientListResponseBodyData) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ModelRouterQueryClientListResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryClientListResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryClientListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryClientListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryClientListResponseBodyData) SetList(v []*ClientDTO) *ModelRouterQueryClientListResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryClientListResponseBodyData) SetMaxResults(v string) *ModelRouterQueryClientListResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBodyData) SetNextToken(v string) *ModelRouterQueryClientListResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBodyData) SetPage(v int32) *ModelRouterQueryClientListResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBodyData) SetPageSize(v int32) *ModelRouterQueryClientListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBodyData) SetTotal(v int32) *ModelRouterQueryClientListResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBodyData) Validate() error {
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
