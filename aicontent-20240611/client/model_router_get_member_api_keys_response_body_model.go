// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetMemberApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterGetMemberApiKeysResponseBodyData) *ModelRouterGetMemberApiKeysResponseBody
	GetData() *ModelRouterGetMemberApiKeysResponseBodyData
	SetErrCode(v string) *ModelRouterGetMemberApiKeysResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterGetMemberApiKeysResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterGetMemberApiKeysResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterGetMemberApiKeysResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterGetMemberApiKeysResponseBody
	GetSuccess() *bool
}

type ModelRouterGetMemberApiKeysResponseBody struct {
	// The response data object.
	//
	// example:
	//
	// {}
	Data *ModelRouterGetMemberApiKeysResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ModelRouterGetMemberApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberApiKeysResponseBody) GetData() *ModelRouterGetMemberApiKeysResponseBodyData {
	return s.Data
}

func (s *ModelRouterGetMemberApiKeysResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterGetMemberApiKeysResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterGetMemberApiKeysResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterGetMemberApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterGetMemberApiKeysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterGetMemberApiKeysResponseBody) SetData(v *ModelRouterGetMemberApiKeysResponseBodyData) *ModelRouterGetMemberApiKeysResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBody) SetErrCode(v string) *ModelRouterGetMemberApiKeysResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBody) SetErrMessage(v string) *ModelRouterGetMemberApiKeysResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBody) SetHttpStatusCode(v int32) *ModelRouterGetMemberApiKeysResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBody) SetRequestId(v string) *ModelRouterGetMemberApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBody) SetSuccess(v bool) *ModelRouterGetMemberApiKeysResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterGetMemberApiKeysResponseBodyData struct {
	// The list of returned data.
	//
	// example:
	//
	// []
	List []*MemberApiKeyDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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

func (s ModelRouterGetMemberApiKeysResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberApiKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberApiKeysResponseBodyData) GetList() []*MemberApiKeyDTO {
	return s.List
}

func (s *ModelRouterGetMemberApiKeysResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterGetMemberApiKeysResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterGetMemberApiKeysResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterGetMemberApiKeysResponseBodyData) SetList(v []*MemberApiKeyDTO) *ModelRouterGetMemberApiKeysResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBodyData) SetPage(v int32) *ModelRouterGetMemberApiKeysResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBodyData) SetPageSize(v int32) *ModelRouterGetMemberApiKeysResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBodyData) SetTotal(v int32) *ModelRouterGetMemberApiKeysResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterGetMemberApiKeysResponseBodyData) Validate() error {
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
