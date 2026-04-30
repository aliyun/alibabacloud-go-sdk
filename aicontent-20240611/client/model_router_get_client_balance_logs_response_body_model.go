// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetClientBalanceLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterGetClientBalanceLogsResponseBodyData) *ModelRouterGetClientBalanceLogsResponseBody
	GetData() *ModelRouterGetClientBalanceLogsResponseBodyData
	SetErrCode(v string) *ModelRouterGetClientBalanceLogsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterGetClientBalanceLogsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterGetClientBalanceLogsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterGetClientBalanceLogsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterGetClientBalanceLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterGetClientBalanceLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterGetClientBalanceLogsResponseBody
	GetSuccess() *bool
}

type ModelRouterGetClientBalanceLogsResponseBody struct {
	Data *ModelRouterGetClientBalanceLogsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// None
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterGetClientBalanceLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetClientBalanceLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) GetData() *ModelRouterGetClientBalanceLogsResponseBodyData {
	return s.Data
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) SetData(v *ModelRouterGetClientBalanceLogsResponseBodyData) *ModelRouterGetClientBalanceLogsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) SetErrCode(v string) *ModelRouterGetClientBalanceLogsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) SetErrMessage(v string) *ModelRouterGetClientBalanceLogsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) SetHttpStatusCode(v int32) *ModelRouterGetClientBalanceLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) SetMaxResults(v int32) *ModelRouterGetClientBalanceLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) SetNextToken(v string) *ModelRouterGetClientBalanceLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) SetRequestId(v string) *ModelRouterGetClientBalanceLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) SetSuccess(v bool) *ModelRouterGetClientBalanceLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterGetClientBalanceLogsResponseBodyData struct {
	List []*ClientBalanceLogDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterGetClientBalanceLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetClientBalanceLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterGetClientBalanceLogsResponseBodyData) GetList() []*ClientBalanceLogDTO {
	return s.List
}

func (s *ModelRouterGetClientBalanceLogsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterGetClientBalanceLogsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterGetClientBalanceLogsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterGetClientBalanceLogsResponseBodyData) SetList(v []*ClientBalanceLogDTO) *ModelRouterGetClientBalanceLogsResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBodyData) SetPage(v int32) *ModelRouterGetClientBalanceLogsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBodyData) SetPageSize(v int32) *ModelRouterGetClientBalanceLogsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBodyData) SetTotal(v int32) *ModelRouterGetClientBalanceLogsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsResponseBodyData) Validate() error {
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
