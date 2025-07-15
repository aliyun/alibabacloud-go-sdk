// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoriesByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCategoriesByTaskIdResponseBody
	GetCode() *string
	SetData(v []*GetCategoriesByTaskIdResponseBodyData) *GetCategoriesByTaskIdResponseBody
	GetData() []*GetCategoriesByTaskIdResponseBodyData
	SetHttpStatusCode(v int32) *GetCategoriesByTaskIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCategoriesByTaskIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCategoriesByTaskIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCategoriesByTaskIdResponseBody
	GetSuccess() *bool
}

type GetCategoriesByTaskIdResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetCategoriesByTaskIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s GetCategoriesByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetCategoriesByTaskIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCategoriesByTaskIdResponseBody) GetData() []*GetCategoriesByTaskIdResponseBodyData {
	return s.Data
}

func (s *GetCategoriesByTaskIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCategoriesByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCategoriesByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCategoriesByTaskIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCategoriesByTaskIdResponseBody) SetCode(v string) *GetCategoriesByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetCategoriesByTaskIdResponseBody) SetData(v []*GetCategoriesByTaskIdResponseBodyData) *GetCategoriesByTaskIdResponseBody {
	s.Data = v
	return s
}

func (s *GetCategoriesByTaskIdResponseBody) SetHttpStatusCode(v int32) *GetCategoriesByTaskIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCategoriesByTaskIdResponseBody) SetMessage(v string) *GetCategoriesByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetCategoriesByTaskIdResponseBody) SetRequestId(v string) *GetCategoriesByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCategoriesByTaskIdResponseBody) SetSuccess(v bool) *GetCategoriesByTaskIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetCategoriesByTaskIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCategoriesByTaskIdResponseBodyData struct {
	Category *string                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	Children []*GetCategoriesByTaskIdResponseBodyDataChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetCategoriesByTaskIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesByTaskIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCategoriesByTaskIdResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *GetCategoriesByTaskIdResponseBodyData) GetChildren() []*GetCategoriesByTaskIdResponseBodyDataChildren {
	return s.Children
}

func (s *GetCategoriesByTaskIdResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetCategoriesByTaskIdResponseBodyData) SetCategory(v string) *GetCategoriesByTaskIdResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetCategoriesByTaskIdResponseBodyData) SetChildren(v []*GetCategoriesByTaskIdResponseBodyDataChildren) *GetCategoriesByTaskIdResponseBodyData {
	s.Children = v
	return s
}

func (s *GetCategoriesByTaskIdResponseBodyData) SetCount(v int32) *GetCategoriesByTaskIdResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetCategoriesByTaskIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetCategoriesByTaskIdResponseBodyDataChildren struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s GetCategoriesByTaskIdResponseBodyDataChildren) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesByTaskIdResponseBodyDataChildren) GoString() string {
	return s.String()
}

func (s *GetCategoriesByTaskIdResponseBodyDataChildren) GetCategory() *string {
	return s.Category
}

func (s *GetCategoriesByTaskIdResponseBodyDataChildren) SetCategory(v string) *GetCategoriesByTaskIdResponseBodyDataChildren {
	s.Category = &v
	return s
}

func (s *GetCategoriesByTaskIdResponseBodyDataChildren) Validate() error {
	return dara.Validate(s)
}
