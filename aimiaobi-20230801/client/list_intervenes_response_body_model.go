// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIntervenesResponseBody
	GetCode() *string
	SetData(v *ListIntervenesResponseBodyData) *ListIntervenesResponseBody
	GetData() *ListIntervenesResponseBodyData
	SetHttpStatusCode(v int32) *ListIntervenesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListIntervenesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIntervenesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIntervenesResponseBody
	GetSuccess() *bool
}

type ListIntervenesResponseBody struct {
	// example:
	//
	// 0
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListIntervenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 428DCC0D-3C63-5306-BD1B-124396AB97BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIntervenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntervenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIntervenesResponseBody) GetData() *ListIntervenesResponseBodyData {
	return s.Data
}

func (s *ListIntervenesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIntervenesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIntervenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntervenesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIntervenesResponseBody) SetCode(v string) *ListIntervenesResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervenesResponseBody) SetData(v *ListIntervenesResponseBodyData) *ListIntervenesResponseBody {
	s.Data = v
	return s
}

func (s *ListIntervenesResponseBody) SetHttpStatusCode(v int32) *ListIntervenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervenesResponseBody) SetMessage(v string) *ListIntervenesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervenesResponseBody) SetRequestId(v string) *ListIntervenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervenesResponseBody) SetSuccess(v bool) *ListIntervenesResponseBody {
	s.Success = &v
	return s
}

func (s *ListIntervenesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIntervenesResponseBodyData struct {
	Code          *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	InterveneList []*ListIntervenesResponseBodyDataInterveneList `json:"InterveneList,omitempty" xml:"InterveneList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListIntervenesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIntervenesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *ListIntervenesResponseBodyData) GetInterveneList() []*ListIntervenesResponseBodyDataInterveneList {
	return s.InterveneList
}

func (s *ListIntervenesResponseBodyData) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListIntervenesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntervenesResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListIntervenesResponseBodyData) SetCode(v int32) *ListIntervenesResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListIntervenesResponseBodyData) SetInterveneList(v []*ListIntervenesResponseBodyDataInterveneList) *ListIntervenesResponseBodyData {
	s.InterveneList = v
	return s
}

func (s *ListIntervenesResponseBodyData) SetPageIndex(v int32) *ListIntervenesResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListIntervenesResponseBodyData) SetPageSize(v int32) *ListIntervenesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIntervenesResponseBodyData) SetTotalSize(v int64) *ListIntervenesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListIntervenesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListIntervenesResponseBodyDataInterveneList struct {
	// id
	//
	// example:
	//
	// 36559
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListIntervenesResponseBodyDataInterveneList) String() string {
	return dara.Prettify(s)
}

func (s ListIntervenesResponseBodyDataInterveneList) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponseBodyDataInterveneList) GetId() *string {
	return s.Id
}

func (s *ListIntervenesResponseBodyDataInterveneList) GetQuery() *string {
	return s.Query
}

func (s *ListIntervenesResponseBodyDataInterveneList) SetId(v string) *ListIntervenesResponseBodyDataInterveneList {
	s.Id = &v
	return s
}

func (s *ListIntervenesResponseBodyDataInterveneList) SetQuery(v string) *ListIntervenesResponseBodyDataInterveneList {
	s.Query = &v
	return s
}

func (s *ListIntervenesResponseBodyDataInterveneList) Validate() error {
	return dara.Validate(s)
}
