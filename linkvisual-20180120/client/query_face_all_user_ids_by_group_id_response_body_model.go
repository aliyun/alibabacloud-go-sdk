// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllUserIdsByGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceAllUserIdsByGroupIdResponseBody
	GetCode() *string
	SetData(v *QueryFaceAllUserIdsByGroupIdResponseBodyData) *QueryFaceAllUserIdsByGroupIdResponseBody
	GetData() *QueryFaceAllUserIdsByGroupIdResponseBodyData
	SetErrorMessage(v string) *QueryFaceAllUserIdsByGroupIdResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryFaceAllUserIdsByGroupIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFaceAllUserIdsByGroupIdResponseBody
	GetSuccess() *bool
}

type QueryFaceAllUserIdsByGroupIdResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryFaceAllUserIdsByGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A13279A-5640-45E7-87AA-83750541AD0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) GetData() *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	return s.Data
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetCode(v string) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetData(v *QueryFaceAllUserIdsByGroupIdResponseBodyData) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetErrorMessage(v string) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetRequestId(v string) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) SetSuccess(v bool) *QueryFaceAllUserIdsByGroupIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceAllUserIdsByGroupIdResponseBodyData struct {
	// The list of users.
	List []*QueryFaceAllUserIdsByGroupIdResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of users.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) GetList() []*QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	return s.List
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetList(v []*QueryFaceAllUserIdsByGroupIdResponseBodyDataList) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.List = v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetPage(v int32) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetPageSize(v int32) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) SetTotal(v int32) *QueryFaceAllUserIdsByGroupIdResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyData) Validate() error {
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

type QueryFaceAllUserIdsByGroupIdResponseBodyDataList struct {
	// The user ID in the caller service.
	//
	// example:
	//
	// TestCustomUserId
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// The username.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user parameters.
	//
	// example:
	//
	// {"age":28, "sex":"male"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The user ID.
	//
	// example:
	//
	// TestUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserIdsByGroupIdResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) GetCustomUserId() *string {
	return s.CustomUserId
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) GetParams() *string {
	return s.Params
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) GetUserId() *string {
	return s.UserId
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetCustomUserId(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.CustomUserId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetName(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetParams(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.Params = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) SetUserId(v string) *QueryFaceAllUserIdsByGroupIdResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *QueryFaceAllUserIdsByGroupIdResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
