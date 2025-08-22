// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAppNamesResponseBody
	GetCode() *int32
	SetData(v []*ListAppNamesResponseBodyData) *ListAppNamesResponseBody
	GetData() []*ListAppNamesResponseBodyData
	SetMessage(v string) *ListAppNamesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAppNamesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAppNamesResponseBody
	GetSuccess() *bool
}

type ListAppNamesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// .
	Data []*ListAppNamesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3881C59F-59F1-5B2E-8110-7D689CA9B207
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppNamesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAppNamesResponseBody) GetData() []*ListAppNamesResponseBodyData {
	return s.Data
}

func (s *ListAppNamesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppNamesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAppNamesResponseBody) SetCode(v int32) *ListAppNamesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppNamesResponseBody) SetData(v []*ListAppNamesResponseBodyData) *ListAppNamesResponseBody {
	s.Data = v
	return s
}

func (s *ListAppNamesResponseBody) SetMessage(v string) *ListAppNamesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppNamesResponseBody) SetRequestId(v string) *ListAppNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppNamesResponseBody) SetSuccess(v bool) *ListAppNamesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppNamesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAppNamesResponseBodyData struct {
	// example:
	//
	// test-app
	AppGroupId *string `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 15
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListAppNamesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAppNamesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppNamesResponseBodyData) GetAppGroupId() *string {
	return s.AppGroupId
}

func (s *ListAppNamesResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *ListAppNamesResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListAppNamesResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListAppNamesResponseBodyData) SetAppGroupId(v string) *ListAppNamesResponseBodyData {
	s.AppGroupId = &v
	return s
}

func (s *ListAppNamesResponseBodyData) SetAppName(v string) *ListAppNamesResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListAppNamesResponseBodyData) SetId(v int64) *ListAppNamesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAppNamesResponseBodyData) SetTitle(v string) *ListAppNamesResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListAppNamesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
