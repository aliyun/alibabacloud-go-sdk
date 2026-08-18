// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetMemberBalanceLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterGetMemberBalanceLogsResponseBodyData) *ModelRouterGetMemberBalanceLogsResponseBody
	GetData() *ModelRouterGetMemberBalanceLogsResponseBodyData
	SetErrCode(v string) *ModelRouterGetMemberBalanceLogsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterGetMemberBalanceLogsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterGetMemberBalanceLogsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterGetMemberBalanceLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterGetMemberBalanceLogsResponseBody
	GetSuccess() *bool
}

type ModelRouterGetMemberBalanceLogsResponseBody struct {
	// The response data object.
	//
	// example:
	//
	// {}
	Data *ModelRouterGetMemberBalanceLogsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ModelRouterGetMemberBalanceLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberBalanceLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) GetData() *ModelRouterGetMemberBalanceLogsResponseBodyData {
	return s.Data
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) SetData(v *ModelRouterGetMemberBalanceLogsResponseBodyData) *ModelRouterGetMemberBalanceLogsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) SetErrCode(v string) *ModelRouterGetMemberBalanceLogsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) SetErrMessage(v string) *ModelRouterGetMemberBalanceLogsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) SetHttpStatusCode(v int32) *ModelRouterGetMemberBalanceLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) SetRequestId(v string) *ModelRouterGetMemberBalanceLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) SetSuccess(v bool) *ModelRouterGetMemberBalanceLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterGetMemberBalanceLogsResponseBodyData struct {
	// The list of balance change logs.
	//
	// example:
	//
	// []
	List []*ClientBalanceLogDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 0
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterGetMemberBalanceLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberBalanceLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberBalanceLogsResponseBodyData) GetList() []*ClientBalanceLogDTO {
	return s.List
}

func (s *ModelRouterGetMemberBalanceLogsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterGetMemberBalanceLogsResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ModelRouterGetMemberBalanceLogsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterGetMemberBalanceLogsResponseBodyData) SetList(v []*ClientBalanceLogDTO) *ModelRouterGetMemberBalanceLogsResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBodyData) SetPage(v int32) *ModelRouterGetMemberBalanceLogsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBodyData) SetSize(v int32) *ModelRouterGetMemberBalanceLogsResponseBodyData {
	s.Size = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBodyData) SetTotal(v int32) *ModelRouterGetMemberBalanceLogsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsResponseBodyData) Validate() error {
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
