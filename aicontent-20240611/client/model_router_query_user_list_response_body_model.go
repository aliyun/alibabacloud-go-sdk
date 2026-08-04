// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryUserListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryUserListResponseBodyData) *ModelRouterQueryUserListResponseBody
	GetData() *ModelRouterQueryUserListResponseBodyData
	SetErrCode(v string) *ModelRouterQueryUserListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryUserListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryUserListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryUserListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryUserListResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryUserListResponseBody struct {
	// The response data object.
	//
	// example:
	//
	// {}
	Data *ModelRouterQueryUserListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault code.
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

func (s ModelRouterQueryUserListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryUserListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryUserListResponseBody) GetData() *ModelRouterQueryUserListResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryUserListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryUserListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryUserListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryUserListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryUserListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryUserListResponseBody) SetData(v *ModelRouterQueryUserListResponseBodyData) *ModelRouterQueryUserListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryUserListResponseBody) SetErrCode(v string) *ModelRouterQueryUserListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryUserListResponseBody) SetErrMessage(v string) *ModelRouterQueryUserListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryUserListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryUserListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryUserListResponseBody) SetRequestId(v string) *ModelRouterQueryUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryUserListResponseBody) SetSuccess(v bool) *ModelRouterQueryUserListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryUserListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryUserListResponseBodyData struct {
	// The list of returned data.
	//
	// example:
	//
	// []
	List []*UserListItemDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
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
	// 0
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryUserListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryUserListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryUserListResponseBodyData) GetList() []*UserListItemDTO {
	return s.List
}

func (s *ModelRouterQueryUserListResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryUserListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryUserListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryUserListResponseBodyData) SetList(v []*UserListItemDTO) *ModelRouterQueryUserListResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryUserListResponseBodyData) SetPage(v int32) *ModelRouterQueryUserListResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryUserListResponseBodyData) SetPageSize(v int32) *ModelRouterQueryUserListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryUserListResponseBodyData) SetTotal(v int32) *ModelRouterQueryUserListResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryUserListResponseBodyData) Validate() error {
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
