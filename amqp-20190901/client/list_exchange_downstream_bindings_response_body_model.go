// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeDownstreamBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListExchangeDownstreamBindingsResponseBody
	GetCode() *int32
	SetData(v *ListExchangeDownstreamBindingsResponseBodyData) *ListExchangeDownstreamBindingsResponseBody
	GetData() *ListExchangeDownstreamBindingsResponseBodyData
	SetMessage(v string) *ListExchangeDownstreamBindingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListExchangeDownstreamBindingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExchangeDownstreamBindingsResponseBody
	GetSuccess() *bool
}

type ListExchangeDownstreamBindingsResponseBody struct {
	Code      *int32                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListExchangeDownstreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExchangeDownstreamBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeDownstreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangeDownstreamBindingsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListExchangeDownstreamBindingsResponseBody) GetData() *ListExchangeDownstreamBindingsResponseBodyData {
	return s.Data
}

func (s *ListExchangeDownstreamBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExchangeDownstreamBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExchangeDownstreamBindingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExchangeDownstreamBindingsResponseBody) SetCode(v int32) *ListExchangeDownstreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBody) SetData(v *ListExchangeDownstreamBindingsResponseBodyData) *ListExchangeDownstreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBody) SetMessage(v string) *ListExchangeDownstreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBody) SetRequestId(v string) *ListExchangeDownstreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBody) SetSuccess(v bool) *ListExchangeDownstreamBindingsResponseBody {
	s.Success = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExchangeDownstreamBindingsResponseBodyData struct {
	CurrentPage *int32                                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VoList      *ListExchangeDownstreamBindingsResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s ListExchangeDownstreamBindingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeDownstreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangeDownstreamBindingsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListExchangeDownstreamBindingsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExchangeDownstreamBindingsResponseBodyData) GetVoList() *ListExchangeDownstreamBindingsResponseBodyDataVoList {
	return s.VoList
}

func (s *ListExchangeDownstreamBindingsResponseBodyData) SetCurrentPage(v int32) *ListExchangeDownstreamBindingsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBodyData) SetPageSize(v int32) *ListExchangeDownstreamBindingsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBodyData) SetVoList(v *ListExchangeDownstreamBindingsResponseBodyDataVoList) *ListExchangeDownstreamBindingsResponseBodyData {
	s.VoList = v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExchangeDownstreamBindingsResponseBodyDataVoList struct {
	BindingVO []*ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO `json:"BindingVO,omitempty" xml:"BindingVO,omitempty" type:"Repeated"`
}

func (s ListExchangeDownstreamBindingsResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeDownstreamBindingsResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoList) GetBindingVO() []*ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO {
	return s.BindingVO
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoList) SetBindingVO(v []*ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) *ListExchangeDownstreamBindingsResponseBodyDataVoList {
	s.BindingVO = v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoList) Validate() error {
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

type ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO struct {
	Argument    *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	BindingKey  *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType *int32  `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	DstName     *string `json:"DstName,omitempty" xml:"DstName,omitempty"`
	SrcName     *string `json:"SrcName,omitempty" xml:"SrcName,omitempty"`
}

func (s ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) GoString() string {
	return s.String()
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) GetArgument() *string {
	return s.Argument
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) GetBindingKey() *string {
	return s.BindingKey
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) GetBindingType() *int32 {
	return s.BindingType
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) GetDstName() *string {
	return s.DstName
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) GetSrcName() *string {
	return s.SrcName
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) SetArgument(v string) *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO {
	s.Argument = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) SetBindingKey(v string) *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO {
	s.BindingKey = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) SetBindingType(v int32) *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO {
	s.BindingType = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) SetDstName(v string) *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO {
	s.DstName = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) SetSrcName(v string) *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO {
	s.SrcName = &v
	return s
}

func (s *ListExchangeDownstreamBindingsResponseBodyDataVoListBindingVO) Validate() error {
	return dara.Validate(s)
}
