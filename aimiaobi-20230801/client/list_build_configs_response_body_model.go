// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBuildConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBuildConfigsResponseBody
	GetCode() *string
	SetData(v []*ListBuildConfigsResponseBodyData) *ListBuildConfigsResponseBody
	GetData() []*ListBuildConfigsResponseBodyData
	SetHttpStatusCode(v int32) *ListBuildConfigsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBuildConfigsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBuildConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBuildConfigsResponseBody
	GetSuccess() *bool
}

type ListBuildConfigsResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListBuildConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DA021073-17CE-5CCF-9FEB-93226C766887
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBuildConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBuildConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBuildConfigsResponseBody) GetData() []*ListBuildConfigsResponseBodyData {
	return s.Data
}

func (s *ListBuildConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBuildConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBuildConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBuildConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBuildConfigsResponseBody) SetCode(v string) *ListBuildConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetData(v []*ListBuildConfigsResponseBodyData) *ListBuildConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListBuildConfigsResponseBody) SetHttpStatusCode(v int32) *ListBuildConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetMessage(v string) *ListBuildConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetRequestId(v string) *ListBuildConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetSuccess(v bool) *ListBuildConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListBuildConfigsResponseBody) Validate() error {
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

type ListBuildConfigsResponseBodyData struct {
	// example:
	//
	// true
	BuildIn *bool `json:"BuildIn,omitempty" xml:"BuildIn,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser *string                                     `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	Id         *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Keywords   []*ListBuildConfigsResponseBodyDataKeywords `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// example:
	//
	// writingStyle
	Tag            *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// example:
	//
	// media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s ListBuildConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBuildConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponseBodyData) GetBuildIn() *bool {
	return s.BuildIn
}

func (s *ListBuildConfigsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBuildConfigsResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListBuildConfigsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListBuildConfigsResponseBodyData) GetKeywords() []*ListBuildConfigsResponseBodyDataKeywords {
	return s.Keywords
}

func (s *ListBuildConfigsResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *ListBuildConfigsResponseBodyData) GetTagDescription() *string {
	return s.TagDescription
}

func (s *ListBuildConfigsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListBuildConfigsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListBuildConfigsResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *ListBuildConfigsResponseBodyData) SetBuildIn(v bool) *ListBuildConfigsResponseBodyData {
	s.BuildIn = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetCreateTime(v string) *ListBuildConfigsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetCreateUser(v string) *ListBuildConfigsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetId(v int64) *ListBuildConfigsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetKeywords(v []*ListBuildConfigsResponseBodyDataKeywords) *ListBuildConfigsResponseBodyData {
	s.Keywords = v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetTag(v string) *ListBuildConfigsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetTagDescription(v string) *ListBuildConfigsResponseBodyData {
	s.TagDescription = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetType(v string) *ListBuildConfigsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetUpdateTime(v string) *ListBuildConfigsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetUpdateUser(v string) *ListBuildConfigsResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) Validate() error {
	if s.Keywords != nil {
		for _, item := range s.Keywords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBuildConfigsResponseBodyDataKeywords struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListBuildConfigsResponseBodyDataKeywords) String() string {
	return dara.Prettify(s)
}

func (s ListBuildConfigsResponseBodyDataKeywords) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponseBodyDataKeywords) GetDescription() *string {
	return s.Description
}

func (s *ListBuildConfigsResponseBodyDataKeywords) GetKey() *string {
	return s.Key
}

func (s *ListBuildConfigsResponseBodyDataKeywords) SetDescription(v string) *ListBuildConfigsResponseBodyDataKeywords {
	s.Description = &v
	return s
}

func (s *ListBuildConfigsResponseBodyDataKeywords) SetKey(v string) *ListBuildConfigsResponseBodyDataKeywords {
	s.Key = &v
	return s
}

func (s *ListBuildConfigsResponseBodyDataKeywords) Validate() error {
	return dara.Validate(s)
}
