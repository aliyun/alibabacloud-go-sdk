// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeUpstreamBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListExchangeUpstreamBindingsResponseBody
	GetCode() *int32
	SetData(v *ListExchangeUpstreamBindingsResponseBodyData) *ListExchangeUpstreamBindingsResponseBody
	GetData() *ListExchangeUpstreamBindingsResponseBodyData
	SetMessage(v string) *ListExchangeUpstreamBindingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListExchangeUpstreamBindingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExchangeUpstreamBindingsResponseBody
	GetSuccess() *bool
}

type ListExchangeUpstreamBindingsResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListExchangeUpstreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExchangeUpstreamBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpstreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangeUpstreamBindingsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListExchangeUpstreamBindingsResponseBody) GetData() *ListExchangeUpstreamBindingsResponseBodyData {
	return s.Data
}

func (s *ListExchangeUpstreamBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExchangeUpstreamBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExchangeUpstreamBindingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExchangeUpstreamBindingsResponseBody) SetCode(v int32) *ListExchangeUpstreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBody) SetData(v *ListExchangeUpstreamBindingsResponseBodyData) *ListExchangeUpstreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBody) SetMessage(v string) *ListExchangeUpstreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBody) SetRequestId(v string) *ListExchangeUpstreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBody) SetSuccess(v bool) *ListExchangeUpstreamBindingsResponseBody {
	s.Success = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExchangeUpstreamBindingsResponseBodyData struct {
	CurrentPage *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VoList      *ListExchangeUpstreamBindingsResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s ListExchangeUpstreamBindingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpstreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangeUpstreamBindingsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListExchangeUpstreamBindingsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExchangeUpstreamBindingsResponseBodyData) GetVoList() *ListExchangeUpstreamBindingsResponseBodyDataVoList {
	return s.VoList
}

func (s *ListExchangeUpstreamBindingsResponseBodyData) SetCurrentPage(v int32) *ListExchangeUpstreamBindingsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBodyData) SetPageSize(v int32) *ListExchangeUpstreamBindingsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBodyData) SetVoList(v *ListExchangeUpstreamBindingsResponseBodyDataVoList) *ListExchangeUpstreamBindingsResponseBodyData {
	s.VoList = v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExchangeUpstreamBindingsResponseBodyDataVoList struct {
	BindingVO []*ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO `json:"BindingVO,omitempty" xml:"BindingVO,omitempty" type:"Repeated"`
}

func (s ListExchangeUpstreamBindingsResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpstreamBindingsResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoList) GetBindingVO() []*ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO {
	return s.BindingVO
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoList) SetBindingVO(v []*ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) *ListExchangeUpstreamBindingsResponseBodyDataVoList {
	s.BindingVO = v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoList) Validate() error {
	if s.BindingVO != nil {
		for _, item := range s.BindingVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO struct {
	Argument   *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	BindType   *int32  `json:"BindType,omitempty" xml:"BindType,omitempty"`
	BindingKey *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	DstName    *string `json:"DstName,omitempty" xml:"DstName,omitempty"`
	SrcName    *string `json:"SrcName,omitempty" xml:"SrcName,omitempty"`
}

func (s ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) GoString() string {
	return s.String()
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) GetArgument() *string {
	return s.Argument
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) GetBindType() *int32 {
	return s.BindType
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) GetBindingKey() *string {
	return s.BindingKey
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) GetDstName() *string {
	return s.DstName
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) GetSrcName() *string {
	return s.SrcName
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) SetArgument(v string) *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.Argument = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) SetBindType(v int32) *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.BindType = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) SetBindingKey(v string) *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.BindingKey = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) SetDstName(v string) *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.DstName = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) SetSrcName(v string) *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.SrcName = &v
	return s
}

func (s *ListExchangeUpstreamBindingsResponseBodyDataVoListBindingVO) Validate() error {
	return dara.Validate(s)
}
