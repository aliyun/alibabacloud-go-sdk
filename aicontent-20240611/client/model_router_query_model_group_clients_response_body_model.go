// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryModelGroupClientsResponseBodyData) *ModelRouterQueryModelGroupClientsResponseBody
	GetData() *ModelRouterQueryModelGroupClientsResponseBodyData
	SetErrCode(v string) *ModelRouterQueryModelGroupClientsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelGroupClientsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupClientsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryModelGroupClientsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryModelGroupClientsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryModelGroupClientsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryModelGroupClientsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryModelGroupClientsResponseBody struct {
	// The model usage table structure.
	Data *ModelRouterQueryModelGroupClientsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault message code.
	//
	// example:
	//
	// null
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
	// maxResults
	//
	// example:
	//
	// 10
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

func (s ModelRouterQueryModelGroupClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) GetData() *ModelRouterQueryModelGroupClientsResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) SetData(v *ModelRouterQueryModelGroupClientsResponseBodyData) *ModelRouterQueryModelGroupClientsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) SetErrCode(v string) *ModelRouterQueryModelGroupClientsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) SetErrMessage(v string) *ModelRouterQueryModelGroupClientsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupClientsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) SetMaxResults(v int32) *ModelRouterQueryModelGroupClientsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) SetNextToken(v string) *ModelRouterQueryModelGroupClientsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) SetRequestId(v string) *ModelRouterQueryModelGroupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) SetSuccess(v bool) *ModelRouterQueryModelGroupClientsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryModelGroupClientsResponseBodyData struct {
	// The list of departments.
	List []*ModelGroupClientDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// None
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryModelGroupClientsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupClientsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupClientsResponseBodyData) GetList() []*ModelGroupClientDTO {
	return s.List
}

func (s *ModelRouterQueryModelGroupClientsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryModelGroupClientsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelGroupClientsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryModelGroupClientsResponseBodyData) SetList(v []*ModelGroupClientDTO) *ModelRouterQueryModelGroupClientsResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBodyData) SetPage(v int32) *ModelRouterQueryModelGroupClientsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBodyData) SetPageSize(v int32) *ModelRouterQueryModelGroupClientsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBodyData) SetTotal(v int32) *ModelRouterQueryModelGroupClientsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsResponseBodyData) Validate() error {
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
