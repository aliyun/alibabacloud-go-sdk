// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodEsTemplateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataItems(v []*ListVodEsTemplateInfoResponseBodyDataItems) *ListVodEsTemplateInfoResponseBody
	GetDataItems() []*ListVodEsTemplateInfoResponseBodyDataItems
	SetPageNumber(v int32) *ListVodEsTemplateInfoResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVodEsTemplateInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListVodEsTemplateInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVodEsTemplateInfoResponseBody
	GetTotalCount() *int32
}

type ListVodEsTemplateInfoResponseBody struct {
	DataItems  []*ListVodEsTemplateInfoResponseBodyDataItems `json:"DataItems,omitempty" xml:"DataItems,omitempty" type:"Repeated"`
	PageNumber *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVodEsTemplateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVodEsTemplateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodEsTemplateInfoResponseBody) GetDataItems() []*ListVodEsTemplateInfoResponseBodyDataItems {
	return s.DataItems
}

func (s *ListVodEsTemplateInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVodEsTemplateInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVodEsTemplateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVodEsTemplateInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVodEsTemplateInfoResponseBody) SetDataItems(v []*ListVodEsTemplateInfoResponseBodyDataItems) *ListVodEsTemplateInfoResponseBody {
	s.DataItems = v
	return s
}

func (s *ListVodEsTemplateInfoResponseBody) SetPageNumber(v int32) *ListVodEsTemplateInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVodEsTemplateInfoResponseBody) SetPageSize(v int32) *ListVodEsTemplateInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVodEsTemplateInfoResponseBody) SetRequestId(v string) *ListVodEsTemplateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodEsTemplateInfoResponseBody) SetTotalCount(v int32) *ListVodEsTemplateInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVodEsTemplateInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVodEsTemplateInfoResponseBodyDataItems struct {
	Id       *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	TmplDef  *string `json:"TmplDef,omitempty" xml:"TmplDef,omitempty"`
	TmplDesc *string `json:"TmplDesc,omitempty" xml:"TmplDesc,omitempty"`
	TmplName *string `json:"TmplName,omitempty" xml:"TmplName,omitempty"`
}

func (s ListVodEsTemplateInfoResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListVodEsTemplateInfoResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListVodEsTemplateInfoResponseBodyDataItems) GetId() *int32 {
	return s.Id
}

func (s *ListVodEsTemplateInfoResponseBodyDataItems) GetTmplDef() *string {
	return s.TmplDef
}

func (s *ListVodEsTemplateInfoResponseBodyDataItems) GetTmplDesc() *string {
	return s.TmplDesc
}

func (s *ListVodEsTemplateInfoResponseBodyDataItems) GetTmplName() *string {
	return s.TmplName
}

func (s *ListVodEsTemplateInfoResponseBodyDataItems) SetId(v int32) *ListVodEsTemplateInfoResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListVodEsTemplateInfoResponseBodyDataItems) SetTmplDef(v string) *ListVodEsTemplateInfoResponseBodyDataItems {
	s.TmplDef = &v
	return s
}

func (s *ListVodEsTemplateInfoResponseBodyDataItems) SetTmplDesc(v string) *ListVodEsTemplateInfoResponseBodyDataItems {
	s.TmplDesc = &v
	return s
}

func (s *ListVodEsTemplateInfoResponseBodyDataItems) SetTmplName(v string) *ListVodEsTemplateInfoResponseBodyDataItems {
	s.TmplName = &v
	return s
}

func (s *ListVodEsTemplateInfoResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
