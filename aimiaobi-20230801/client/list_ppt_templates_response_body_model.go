// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPptTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPptTemplatesResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListPptTemplatesResponseBody
	GetCurrent() *int32
	SetData(v []*ListPptTemplatesResponseBodyData) *ListPptTemplatesResponseBody
	GetData() []*ListPptTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *ListPptTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListPptTemplatesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListPptTemplatesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListPptTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPptTemplatesResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListPptTemplatesResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListPptTemplatesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListPptTemplatesResponseBody
	GetTotal() *int32
}

type ListPptTemplatesResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                              `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListPptTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// +CBOXvu2YLxC6DOua8Qupg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPptTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPptTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPptTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPptTemplatesResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListPptTemplatesResponseBody) GetData() []*ListPptTemplatesResponseBodyData {
	return s.Data
}

func (s *ListPptTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPptTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPptTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPptTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPptTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPptTemplatesResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListPptTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPptTemplatesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListPptTemplatesResponseBody) SetCode(v string) *ListPptTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPptTemplatesResponseBody) SetCurrent(v int32) *ListPptTemplatesResponseBody {
	s.Current = &v
	return s
}

func (s *ListPptTemplatesResponseBody) SetData(v []*ListPptTemplatesResponseBodyData) *ListPptTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListPptTemplatesResponseBody) SetHttpStatusCode(v int32) *ListPptTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPptTemplatesResponseBody) SetMaxResults(v int32) *ListPptTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPptTemplatesResponseBody) SetMessage(v string) *ListPptTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPptTemplatesResponseBody) SetNextToken(v string) *ListPptTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPptTemplatesResponseBody) SetRequestId(v string) *ListPptTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPptTemplatesResponseBody) SetSize(v int32) *ListPptTemplatesResponseBody {
	s.Size = &v
	return s
}

func (s *ListPptTemplatesResponseBody) SetSuccess(v bool) *ListPptTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPptTemplatesResponseBody) SetTotal(v int32) *ListPptTemplatesResponseBody {
	s.Total = &v
	return s
}

func (s *ListPptTemplatesResponseBody) Validate() error {
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

type ListPptTemplatesResponseBodyData struct {
	// example:
	//
	// http://xxx.com/a.png
	CoverImg *string `json:"CoverImg,omitempty" xml:"CoverImg,omitempty"`
	// ID
	//
	// example:
	//
	// 10
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListPptTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPptTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPptTemplatesResponseBodyData) GetCoverImg() *string {
	return s.CoverImg
}

func (s *ListPptTemplatesResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListPptTemplatesResponseBodyData) SetCoverImg(v string) *ListPptTemplatesResponseBodyData {
	s.CoverImg = &v
	return s
}

func (s *ListPptTemplatesResponseBodyData) SetId(v int64) *ListPptTemplatesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListPptTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
