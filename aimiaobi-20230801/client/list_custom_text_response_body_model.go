// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCustomTextResponseBody
	GetCode() *string
	SetData(v []*ListCustomTextResponseBodyData) *ListCustomTextResponseBody
	GetData() []*ListCustomTextResponseBodyData
	SetHttpStatusCode(v int32) *ListCustomTextResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCustomTextResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCustomTextResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomTextResponseBody
	GetSuccess() *bool
}

type ListCustomTextResponseBody struct {
	// example:
	//
	// NoData
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListCustomTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCustomTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCustomTextResponseBody) GetData() []*ListCustomTextResponseBodyData {
	return s.Data
}

func (s *ListCustomTextResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCustomTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCustomTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomTextResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomTextResponseBody) SetCode(v string) *ListCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomTextResponseBody) SetData(v []*ListCustomTextResponseBodyData) *ListCustomTextResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomTextResponseBody) SetHttpStatusCode(v int32) *ListCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCustomTextResponseBody) SetMessage(v string) *ListCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomTextResponseBody) SetRequestId(v string) *ListCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomTextResponseBody) SetSuccess(v bool) *ListCustomTextResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomTextResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomTextResponseBodyData struct {
	// example:
	//
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 创建用户
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 40
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 修改用户
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s ListCustomTextResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomTextResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListCustomTextResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCustomTextResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListCustomTextResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListCustomTextResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListCustomTextResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCustomTextResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *ListCustomTextResponseBodyData) SetContent(v string) *ListCustomTextResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetCreateTime(v string) *ListCustomTextResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetCreateUser(v string) *ListCustomTextResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetId(v int64) *ListCustomTextResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetTitle(v string) *ListCustomTextResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetUpdateTime(v string) *ListCustomTextResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetUpdateUser(v string) *ListCustomTextResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *ListCustomTextResponseBodyData) Validate() error {
	return dara.Validate(s)
}
