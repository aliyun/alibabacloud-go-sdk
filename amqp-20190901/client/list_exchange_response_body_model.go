// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExchangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListExchangeResponseBody
	GetCode() *int32
	SetData(v *ListExchangeResponseBodyData) *ListExchangeResponseBody
	GetData() *ListExchangeResponseBodyData
	SetMessage(v string) *ListExchangeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListExchangeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExchangeResponseBody
	GetSuccess() *bool
}

type ListExchangeResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListExchangeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExchangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeResponseBody) GoString() string {
	return s.String()
}

func (s *ListExchangeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListExchangeResponseBody) GetData() *ListExchangeResponseBodyData {
	return s.Data
}

func (s *ListExchangeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExchangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExchangeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExchangeResponseBody) SetCode(v int32) *ListExchangeResponseBody {
	s.Code = &v
	return s
}

func (s *ListExchangeResponseBody) SetData(v *ListExchangeResponseBodyData) *ListExchangeResponseBody {
	s.Data = v
	return s
}

func (s *ListExchangeResponseBody) SetMessage(v string) *ListExchangeResponseBody {
	s.Message = &v
	return s
}

func (s *ListExchangeResponseBody) SetRequestId(v string) *ListExchangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExchangeResponseBody) SetSuccess(v bool) *ListExchangeResponseBody {
	s.Success = &v
	return s
}

func (s *ListExchangeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExchangeResponseBodyData struct {
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      *ListExchangeResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s ListExchangeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExchangeResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListExchangeResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExchangeResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListExchangeResponseBodyData) GetVoList() *ListExchangeResponseBodyDataVoList {
	return s.VoList
}

func (s *ListExchangeResponseBodyData) SetCurrentPage(v int32) *ListExchangeResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListExchangeResponseBodyData) SetPageSize(v int32) *ListExchangeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListExchangeResponseBodyData) SetTotalCount(v int64) *ListExchangeResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListExchangeResponseBodyData) SetVoList(v *ListExchangeResponseBodyDataVoList) *ListExchangeResponseBodyData {
	s.VoList = v
	return s
}

func (s *ListExchangeResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExchangeResponseBodyDataVoList struct {
	ExchangVO []*ListExchangeResponseBodyDataVoListExchangVO `json:"ExchangVO,omitempty" xml:"ExchangVO,omitempty" type:"Repeated"`
}

func (s ListExchangeResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *ListExchangeResponseBodyDataVoList) GetExchangVO() []*ListExchangeResponseBodyDataVoListExchangVO {
	return s.ExchangVO
}

func (s *ListExchangeResponseBodyDataVoList) SetExchangVO(v []*ListExchangeResponseBodyDataVoListExchangVO) *ListExchangeResponseBodyDataVoList {
	s.ExchangVO = v
	return s
}

func (s *ListExchangeResponseBodyDataVoList) Validate() error {
	if s.ExchangVO != nil {
		for _, item := range s.ExchangVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExchangeResponseBodyDataVoListExchangVO struct {
	Attributes   *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	AutoDelete   *bool   `json:"AutoDelete,omitempty" xml:"AutoDelete,omitempty"`
	CanDelete    *bool   `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExchangeType *int32  `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	Internal     *bool   `json:"Internal,omitempty" xml:"Internal,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VhostName    *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s ListExchangeResponseBodyDataVoListExchangVO) String() string {
	return dara.Prettify(s)
}

func (s ListExchangeResponseBodyDataVoListExchangVO) GoString() string {
	return s.String()
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) GetAttributes() *string {
	return s.Attributes
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) GetAutoDelete() *bool {
	return s.AutoDelete
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) GetExchangeType() *int32 {
	return s.ExchangeType
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) GetInternal() *bool {
	return s.Internal
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) GetName() *string {
	return s.Name
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) GetVhostName() *string {
	return s.VhostName
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) SetAttributes(v string) *ListExchangeResponseBodyDataVoListExchangVO {
	s.Attributes = &v
	return s
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) SetAutoDelete(v bool) *ListExchangeResponseBodyDataVoListExchangVO {
	s.AutoDelete = &v
	return s
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) SetCanDelete(v bool) *ListExchangeResponseBodyDataVoListExchangVO {
	s.CanDelete = &v
	return s
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) SetCreateTime(v int64) *ListExchangeResponseBodyDataVoListExchangVO {
	s.CreateTime = &v
	return s
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) SetExchangeType(v int32) *ListExchangeResponseBodyDataVoListExchangVO {
	s.ExchangeType = &v
	return s
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) SetInternal(v bool) *ListExchangeResponseBodyDataVoListExchangVO {
	s.Internal = &v
	return s
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) SetName(v string) *ListExchangeResponseBodyDataVoListExchangVO {
	s.Name = &v
	return s
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) SetVhostName(v string) *ListExchangeResponseBodyDataVoListExchangVO {
	s.VhostName = &v
	return s
}

func (s *ListExchangeResponseBodyDataVoListExchangVO) Validate() error {
	return dara.Validate(s)
}
