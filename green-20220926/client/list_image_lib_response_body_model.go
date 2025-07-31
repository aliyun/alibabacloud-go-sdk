// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListImageLibResponseBody
	GetCode() *int32
	SetHttpStatusCode(v int32) *ListImageLibResponseBody
	GetHttpStatusCode() *int32
	SetLibList(v []*ListImageLibResponseBodyLibList) *ListImageLibResponseBody
	GetLibList() []*ListImageLibResponseBodyLibList
	SetMsg(v string) *ListImageLibResponseBody
	GetMsg() *string
	SetRequestId(v string) *ListImageLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImageLibResponseBody
	GetSuccess() *bool
}

type ListImageLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	LibList        []*ListImageLibResponseBodyLibList `json:"LibList,omitempty" xml:"LibList,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListImageLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListImageLibResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListImageLibResponseBody) GetLibList() []*ListImageLibResponseBodyLibList {
	return s.LibList
}

func (s *ListImageLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListImageLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImageLibResponseBody) SetCode(v int32) *ListImageLibResponseBody {
	s.Code = &v
	return s
}

func (s *ListImageLibResponseBody) SetHttpStatusCode(v int32) *ListImageLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListImageLibResponseBody) SetLibList(v []*ListImageLibResponseBodyLibList) *ListImageLibResponseBody {
	s.LibList = v
	return s
}

func (s *ListImageLibResponseBody) SetMsg(v string) *ListImageLibResponseBody {
	s.Msg = &v
	return s
}

func (s *ListImageLibResponseBody) SetRequestId(v string) *ListImageLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageLibResponseBody) SetSuccess(v bool) *ListImageLibResponseBody {
	s.Success = &v
	return s
}

func (s *ListImageLibResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListImageLibResponseBodyLibList struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1
	FreeInspection *int32 `json:"FreeInspection,omitempty" xml:"FreeInspection,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1
	ImageNum *int64 `json:"ImageNum,omitempty" xml:"ImageNum,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId   *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ListImageLibResponseBodyLibList) String() string {
	return dara.Prettify(s)
}

func (s ListImageLibResponseBodyLibList) GoString() string {
	return s.String()
}

func (s *ListImageLibResponseBodyLibList) GetComment() *string {
	return s.Comment
}

func (s *ListImageLibResponseBodyLibList) GetFreeInspection() *int32 {
	return s.FreeInspection
}

func (s *ListImageLibResponseBodyLibList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListImageLibResponseBodyLibList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListImageLibResponseBodyLibList) GetImageNum() *int64 {
	return s.ImageNum
}

func (s *ListImageLibResponseBodyLibList) GetLibId() *string {
	return s.LibId
}

func (s *ListImageLibResponseBodyLibList) GetLibName() *string {
	return s.LibName
}

func (s *ListImageLibResponseBodyLibList) SetComment(v string) *ListImageLibResponseBodyLibList {
	s.Comment = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetFreeInspection(v int32) *ListImageLibResponseBodyLibList {
	s.FreeInspection = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetGmtCreate(v string) *ListImageLibResponseBodyLibList {
	s.GmtCreate = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetGmtModified(v string) *ListImageLibResponseBodyLibList {
	s.GmtModified = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetImageNum(v int64) *ListImageLibResponseBodyLibList {
	s.ImageNum = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetLibId(v string) *ListImageLibResponseBodyLibList {
	s.LibId = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetLibName(v string) *ListImageLibResponseBodyLibList {
	s.LibName = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) Validate() error {
	return dara.Validate(s)
}
