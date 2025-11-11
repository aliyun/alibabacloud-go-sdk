// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterveneCntResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInterveneCntResponseBody
	GetCode() *string
	SetData(v *ListInterveneCntResponseBodyData) *ListInterveneCntResponseBody
	GetData() *ListInterveneCntResponseBodyData
	SetHttpStatusCode(v int32) *ListInterveneCntResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInterveneCntResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInterveneCntResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInterveneCntResponseBody
	GetSuccess() *bool
}

type ListInterveneCntResponseBody struct {
	// example:
	//
	// 0
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListInterveneCntResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInterveneCntResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneCntResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterveneCntResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInterveneCntResponseBody) GetData() *ListInterveneCntResponseBodyData {
	return s.Data
}

func (s *ListInterveneCntResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInterveneCntResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInterveneCntResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInterveneCntResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInterveneCntResponseBody) SetCode(v string) *ListInterveneCntResponseBody {
	s.Code = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetData(v *ListInterveneCntResponseBodyData) *ListInterveneCntResponseBody {
	s.Data = v
	return s
}

func (s *ListInterveneCntResponseBody) SetHttpStatusCode(v int32) *ListInterveneCntResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetMessage(v string) *ListInterveneCntResponseBody {
	s.Message = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetRequestId(v string) *ListInterveneCntResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetSuccess(v bool) *ListInterveneCntResponseBody {
	s.Success = &v
	return s
}

func (s *ListInterveneCntResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInterveneCntResponseBodyData struct {
	CntList []interface{} `json:"CntList,omitempty" xml:"CntList,omitempty" type:"Repeated"`
	Code    *int32        `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10
	PageCnt *int32 `json:"PageCnt,omitempty" xml:"PageCnt,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneCntResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneCntResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInterveneCntResponseBodyData) GetCntList() []interface{} {
	return s.CntList
}

func (s *ListInterveneCntResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *ListInterveneCntResponseBodyData) GetPageCnt() *int32 {
	return s.PageCnt
}

func (s *ListInterveneCntResponseBodyData) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListInterveneCntResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterveneCntResponseBodyData) SetCntList(v []interface{}) *ListInterveneCntResponseBodyData {
	s.CntList = v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetCode(v int32) *ListInterveneCntResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetPageCnt(v int32) *ListInterveneCntResponseBodyData {
	s.PageCnt = &v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetPageIndex(v int32) *ListInterveneCntResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetPageSize(v int32) *ListInterveneCntResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInterveneCntResponseBodyData) Validate() error {
	return dara.Validate(s)
}
