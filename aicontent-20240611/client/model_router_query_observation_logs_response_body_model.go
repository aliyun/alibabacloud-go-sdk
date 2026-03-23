// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryObservationLogsResponseBodyData) *ModelRouterQueryObservationLogsResponseBody
	GetData() *ModelRouterQueryObservationLogsResponseBodyData
	SetErrCode(v string) *ModelRouterQueryObservationLogsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryObservationLogsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryObservationLogsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryObservationLogsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryObservationLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryObservationLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryObservationLogsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryObservationLogsResponseBody struct {
	// example:
	//
	// []
	Data *ModelRouterQueryObservationLogsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
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

func (s ModelRouterQueryObservationLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetData() *ModelRouterQueryObservationLogsResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetData(v *ModelRouterQueryObservationLogsResponseBodyData) *ModelRouterQueryObservationLogsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetErrCode(v string) *ModelRouterQueryObservationLogsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetErrMessage(v string) *ModelRouterQueryObservationLogsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryObservationLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetMaxResults(v int32) *ModelRouterQueryObservationLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetNextToken(v string) *ModelRouterQueryObservationLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetRequestId(v string) *ModelRouterQueryObservationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetSuccess(v bool) *ModelRouterQueryObservationLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryObservationLogsResponseBodyData struct {
	List []*RequestLogDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// None
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryObservationLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) GetList() []*RequestLogDTO {
	return s.List
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) GetNextToken() *int32 {
	return s.NextToken
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) SetList(v []*RequestLogDTO) *ModelRouterQueryObservationLogsResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) SetMaxResults(v int32) *ModelRouterQueryObservationLogsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) SetNextToken(v int32) *ModelRouterQueryObservationLogsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) SetPage(v int32) *ModelRouterQueryObservationLogsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) SetPageSize(v int32) *ModelRouterQueryObservationLogsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) SetTotal(v int32) *ModelRouterQueryObservationLogsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBodyData) Validate() error {
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
