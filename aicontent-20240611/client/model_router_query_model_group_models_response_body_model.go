// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryModelGroupModelsResponseBodyData) *ModelRouterQueryModelGroupModelsResponseBody
	GetData() *ModelRouterQueryModelGroupModelsResponseBodyData
	SetErrCode(v string) *ModelRouterQueryModelGroupModelsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelGroupModelsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupModelsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryModelGroupModelsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryModelGroupModelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryModelGroupModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryModelGroupModelsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryModelGroupModelsResponseBody struct {
	// The response struct.
	Data *ModelRouterQueryModelGroupModelsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message code.
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
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This field is not used.
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

func (s ModelRouterQueryModelGroupModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) GetData() *ModelRouterQueryModelGroupModelsResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) SetData(v *ModelRouterQueryModelGroupModelsResponseBodyData) *ModelRouterQueryModelGroupModelsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) SetErrCode(v string) *ModelRouterQueryModelGroupModelsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) SetErrMessage(v string) *ModelRouterQueryModelGroupModelsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupModelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) SetMaxResults(v int32) *ModelRouterQueryModelGroupModelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) SetNextToken(v string) *ModelRouterQueryModelGroupModelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) SetRequestId(v string) *ModelRouterQueryModelGroupModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) SetSuccess(v bool) *ModelRouterQueryModelGroupModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryModelGroupModelsResponseBodyData struct {
	// The elements.
	List []*ModelGroupModelDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The requested page number.
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

func (s ModelRouterQueryModelGroupModelsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupModelsResponseBodyData) GetList() []*ModelGroupModelDTO {
	return s.List
}

func (s *ModelRouterQueryModelGroupModelsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryModelGroupModelsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelGroupModelsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryModelGroupModelsResponseBodyData) SetList(v []*ModelGroupModelDTO) *ModelRouterQueryModelGroupModelsResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBodyData) SetPage(v int32) *ModelRouterQueryModelGroupModelsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBodyData) SetPageSize(v int32) *ModelRouterQueryModelGroupModelsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBodyData) SetTotal(v int32) *ModelRouterQueryModelGroupModelsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsResponseBodyData) Validate() error {
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
