// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBindingErrorByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetBindingErrorByTaskIdResponseBody
	GetCode() *int32
	SetData(v *GetBindingErrorByTaskIdResponseBodyData) *GetBindingErrorByTaskIdResponseBody
	GetData() *GetBindingErrorByTaskIdResponseBodyData
	SetMessage(v string) *GetBindingErrorByTaskIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBindingErrorByTaskIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBindingErrorByTaskIdResponseBody
	GetSuccess() *bool
}

type GetBindingErrorByTaskIdResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetBindingErrorByTaskIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBindingErrorByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBindingErrorByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetBindingErrorByTaskIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetBindingErrorByTaskIdResponseBody) GetData() *GetBindingErrorByTaskIdResponseBodyData {
	return s.Data
}

func (s *GetBindingErrorByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBindingErrorByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBindingErrorByTaskIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBindingErrorByTaskIdResponseBody) SetCode(v int32) *GetBindingErrorByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBody) SetData(v *GetBindingErrorByTaskIdResponseBodyData) *GetBindingErrorByTaskIdResponseBody {
	s.Data = v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBody) SetMessage(v string) *GetBindingErrorByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBody) SetRequestId(v string) *GetBindingErrorByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBody) SetSuccess(v bool) *GetBindingErrorByTaskIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBindingErrorByTaskIdResponseBodyData struct {
	CurrentPage *int32                                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      *GetBindingErrorByTaskIdResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s GetBindingErrorByTaskIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBindingErrorByTaskIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBindingErrorByTaskIdResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetBindingErrorByTaskIdResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetBindingErrorByTaskIdResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetBindingErrorByTaskIdResponseBodyData) GetVoList() *GetBindingErrorByTaskIdResponseBodyDataVoList {
	return s.VoList
}

func (s *GetBindingErrorByTaskIdResponseBodyData) SetCurrentPage(v int32) *GetBindingErrorByTaskIdResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyData) SetPageSize(v int32) *GetBindingErrorByTaskIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyData) SetTotalCount(v int64) *GetBindingErrorByTaskIdResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyData) SetVoList(v *GetBindingErrorByTaskIdResponseBodyDataVoList) *GetBindingErrorByTaskIdResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBindingErrorByTaskIdResponseBodyDataVoList struct {
	BindingErrorDO []*GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO `json:"BindingErrorDO,omitempty" xml:"BindingErrorDO,omitempty" type:"Repeated"`
}

func (s GetBindingErrorByTaskIdResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s GetBindingErrorByTaskIdResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoList) GetBindingErrorDO() []*GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO {
	return s.BindingErrorDO
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoList) SetBindingErrorDO(v []*GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) *GetBindingErrorByTaskIdResponseBodyDataVoList {
	s.BindingErrorDO = v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoList) Validate() error {
	if s.BindingErrorDO != nil {
		for _, item := range s.BindingErrorDO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO struct {
	Destination     *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	ErrorMessage    *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RoutingKey      *string `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty"`
	Src             *string `json:"Src,omitempty" xml:"Src,omitempty"`
	TaskId          *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VhostName       *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) String() string {
	return dara.Prettify(s)
}

func (s GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) GoString() string {
	return s.String()
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) GetDestination() *string {
	return s.Destination
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) GetDestinationType() *string {
	return s.DestinationType
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) GetRoutingKey() *string {
	return s.RoutingKey
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) GetSrc() *string {
	return s.Src
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) GetVhostName() *string {
	return s.VhostName
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) SetDestination(v string) *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO {
	s.Destination = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) SetDestinationType(v string) *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO {
	s.DestinationType = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) SetErrorMessage(v string) *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO {
	s.ErrorMessage = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) SetRoutingKey(v string) *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO {
	s.RoutingKey = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) SetSrc(v string) *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO {
	s.Src = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) SetTaskId(v int64) *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO {
	s.TaskId = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) SetVhostName(v string) *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO {
	s.VhostName = &v
	return s
}

func (s *GetBindingErrorByTaskIdResponseBodyDataVoListBindingErrorDO) Validate() error {
	return dara.Validate(s)
}
