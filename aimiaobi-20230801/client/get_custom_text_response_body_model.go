// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomTextResponseBody
	GetCode() *string
	SetData(v *GetCustomTextResponseBodyData) *GetCustomTextResponseBody
	GetData() *GetCustomTextResponseBodyData
	SetHttpStatusCode(v int32) *GetCustomTextResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCustomTextResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomTextResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomTextResponseBody
	GetSuccess() *bool
}

type GetCustomTextResponseBody struct {
	// example:
	//
	// NoData
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCustomTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetCustomTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomTextResponseBody) GetData() *GetCustomTextResponseBodyData {
	return s.Data
}

func (s *GetCustomTextResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCustomTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomTextResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomTextResponseBody) SetCode(v string) *GetCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomTextResponseBody) SetData(v *GetCustomTextResponseBodyData) *GetCustomTextResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomTextResponseBody) SetHttpStatusCode(v int32) *GetCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomTextResponseBody) SetMessage(v string) *GetCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomTextResponseBody) SetRequestId(v string) *GetCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomTextResponseBody) SetSuccess(v bool) *GetCustomTextResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomTextResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCustomTextResponseBodyData struct {
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
	// 34
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

func (s GetCustomTextResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomTextResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetCustomTextResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCustomTextResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetCustomTextResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetCustomTextResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetCustomTextResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetCustomTextResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *GetCustomTextResponseBodyData) SetContent(v string) *GetCustomTextResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetCreateTime(v string) *GetCustomTextResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetCreateUser(v string) *GetCustomTextResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetId(v int64) *GetCustomTextResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetTitle(v string) *GetCustomTextResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetUpdateTime(v string) *GetCustomTextResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetUpdateUser(v string) *GetCustomTextResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *GetCustomTextResponseBodyData) Validate() error {
	return dara.Validate(s)
}
