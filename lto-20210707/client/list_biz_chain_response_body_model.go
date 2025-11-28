// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBizChainResponseBody
	GetCode() *string
	SetData(v *ListBizChainResponseBodyData) *ListBizChainResponseBody
	GetData() *ListBizChainResponseBodyData
	SetHttpStatusCode(v int32) *ListBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBizChainResponseBody
	GetSuccess() *bool
}

type ListBizChainResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListBizChainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *ListBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBizChainResponseBody) GetData() *ListBizChainResponseBodyData {
	return s.Data
}

func (s *ListBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBizChainResponseBody) SetCode(v string) *ListBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *ListBizChainResponseBody) SetData(v *ListBizChainResponseBodyData) *ListBizChainResponseBody {
	s.Data = v
	return s
}

func (s *ListBizChainResponseBody) SetHttpStatusCode(v int32) *ListBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBizChainResponseBody) SetMessage(v string) *ListBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *ListBizChainResponseBody) SetRequestId(v string) *ListBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBizChainResponseBody) SetSuccess(v bool) *ListBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *ListBizChainResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBizChainResponseBodyData struct {
	Num      *int32                                  `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListBizChainResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListBizChainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBizChainResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListBizChainResponseBodyData) GetPageData() []*ListBizChainResponseBodyDataPageData {
	return s.PageData
}

func (s *ListBizChainResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListBizChainResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListBizChainResponseBodyData) SetNum(v int32) *ListBizChainResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListBizChainResponseBodyData) SetPageData(v []*ListBizChainResponseBodyDataPageData) *ListBizChainResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListBizChainResponseBodyData) SetSize(v int32) *ListBizChainResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListBizChainResponseBodyData) SetTotal(v int32) *ListBizChainResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListBizChainResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBizChainResponseBodyDataPageData struct {
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBizChainResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListBizChainResponseBodyDataPageData) GetBizChainId() *string {
	return s.BizChainId
}

func (s *ListBizChainResponseBodyDataPageData) GetName() *string {
	return s.Name
}

func (s *ListBizChainResponseBodyDataPageData) GetRemark() *string {
	return s.Remark
}

func (s *ListBizChainResponseBodyDataPageData) GetType() *string {
	return s.Type
}

func (s *ListBizChainResponseBodyDataPageData) SetBizChainId(v string) *ListBizChainResponseBodyDataPageData {
	s.BizChainId = &v
	return s
}

func (s *ListBizChainResponseBodyDataPageData) SetName(v string) *ListBizChainResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListBizChainResponseBodyDataPageData) SetRemark(v string) *ListBizChainResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListBizChainResponseBodyDataPageData) SetType(v string) *ListBizChainResponseBodyDataPageData {
	s.Type = &v
	return s
}

func (s *ListBizChainResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
