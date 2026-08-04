// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListMemberBalanceOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterListMemberBalanceOrdersResponseBodyData) *ModelRouterListMemberBalanceOrdersResponseBody
	GetData() *ModelRouterListMemberBalanceOrdersResponseBodyData
	SetErrCode(v string) *ModelRouterListMemberBalanceOrdersResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterListMemberBalanceOrdersResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterListMemberBalanceOrdersResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterListMemberBalanceOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterListMemberBalanceOrdersResponseBody
	GetSuccess() *bool
}

type ModelRouterListMemberBalanceOrdersResponseBody struct {
	// example:
	//
	// {}
	Data *ModelRouterListMemberBalanceOrdersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ModelRouterListMemberBalanceOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListMemberBalanceOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) GetData() *ModelRouterListMemberBalanceOrdersResponseBodyData {
	return s.Data
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) SetData(v *ModelRouterListMemberBalanceOrdersResponseBodyData) *ModelRouterListMemberBalanceOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) SetErrCode(v string) *ModelRouterListMemberBalanceOrdersResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) SetErrMessage(v string) *ModelRouterListMemberBalanceOrdersResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) SetHttpStatusCode(v int32) *ModelRouterListMemberBalanceOrdersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) SetRequestId(v string) *ModelRouterListMemberBalanceOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) SetSuccess(v bool) *ModelRouterListMemberBalanceOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterListMemberBalanceOrdersResponseBodyData struct {
	// example:
	//
	// []
	List []*BillOrderEntryDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 0
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterListMemberBalanceOrdersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListMemberBalanceOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterListMemberBalanceOrdersResponseBodyData) GetList() []*BillOrderEntryDTO {
	return s.List
}

func (s *ModelRouterListMemberBalanceOrdersResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterListMemberBalanceOrdersResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ModelRouterListMemberBalanceOrdersResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterListMemberBalanceOrdersResponseBodyData) SetList(v []*BillOrderEntryDTO) *ModelRouterListMemberBalanceOrdersResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBodyData) SetPage(v int32) *ModelRouterListMemberBalanceOrdersResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBodyData) SetSize(v int32) *ModelRouterListMemberBalanceOrdersResponseBodyData {
	s.Size = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBodyData) SetTotal(v int32) *ModelRouterListMemberBalanceOrdersResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersResponseBodyData) Validate() error {
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
