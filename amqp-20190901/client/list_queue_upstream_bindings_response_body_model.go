// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueUpstreamBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListQueueUpstreamBindingsResponseBody
	GetCode() *int32
	SetData(v *ListQueueUpstreamBindingsResponseBodyData) *ListQueueUpstreamBindingsResponseBody
	GetData() *ListQueueUpstreamBindingsResponseBodyData
	SetMessage(v string) *ListQueueUpstreamBindingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListQueueUpstreamBindingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQueueUpstreamBindingsResponseBody
	GetSuccess() *bool
}

type ListQueueUpstreamBindingsResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListQueueUpstreamBindingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQueueUpstreamBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpstreamBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueUpstreamBindingsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListQueueUpstreamBindingsResponseBody) GetData() *ListQueueUpstreamBindingsResponseBodyData {
	return s.Data
}

func (s *ListQueueUpstreamBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQueueUpstreamBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueueUpstreamBindingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQueueUpstreamBindingsResponseBody) SetCode(v int32) *ListQueueUpstreamBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBody) SetData(v *ListQueueUpstreamBindingsResponseBodyData) *ListQueueUpstreamBindingsResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBody) SetMessage(v string) *ListQueueUpstreamBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBody) SetRequestId(v string) *ListQueueUpstreamBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBody) SetSuccess(v bool) *ListQueueUpstreamBindingsResponseBody {
	s.Success = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueueUpstreamBindingsResponseBodyData struct {
	CurrentPage *int32                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VoList      *ListQueueUpstreamBindingsResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s ListQueueUpstreamBindingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpstreamBindingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueUpstreamBindingsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListQueueUpstreamBindingsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQueueUpstreamBindingsResponseBodyData) GetVoList() *ListQueueUpstreamBindingsResponseBodyDataVoList {
	return s.VoList
}

func (s *ListQueueUpstreamBindingsResponseBodyData) SetCurrentPage(v int32) *ListQueueUpstreamBindingsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBodyData) SetPageSize(v int32) *ListQueueUpstreamBindingsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBodyData) SetVoList(v *ListQueueUpstreamBindingsResponseBodyDataVoList) *ListQueueUpstreamBindingsResponseBodyData {
	s.VoList = v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueueUpstreamBindingsResponseBodyDataVoList struct {
	BindingVO []*ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO `json:"BindingVO,omitempty" xml:"BindingVO,omitempty" type:"Repeated"`
}

func (s ListQueueUpstreamBindingsResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpstreamBindingsResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoList) GetBindingVO() []*ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO {
	return s.BindingVO
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoList) SetBindingVO(v []*ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) *ListQueueUpstreamBindingsResponseBodyDataVoList {
	s.BindingVO = v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoList) Validate() error {
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

type ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO struct {
	Argument    *string `json:"Argument,omitempty" xml:"Argument,omitempty"`
	BindingKey  *string `json:"BindingKey,omitempty" xml:"BindingKey,omitempty"`
	BindingType *int32  `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	DstName     *string `json:"DstName,omitempty" xml:"DstName,omitempty"`
	SrcName     *string `json:"SrcName,omitempty" xml:"SrcName,omitempty"`
}

func (s ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) GoString() string {
	return s.String()
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) GetArgument() *string {
	return s.Argument
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) GetBindingKey() *string {
	return s.BindingKey
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) GetBindingType() *int32 {
	return s.BindingType
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) GetDstName() *string {
	return s.DstName
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) GetSrcName() *string {
	return s.SrcName
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) SetArgument(v string) *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.Argument = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) SetBindingKey(v string) *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.BindingKey = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) SetBindingType(v int32) *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.BindingType = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) SetDstName(v string) *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.DstName = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) SetSrcName(v string) *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO {
	s.SrcName = &v
	return s
}

func (s *ListQueueUpstreamBindingsResponseBodyDataVoListBindingVO) Validate() error {
	return dara.Validate(s)
}
