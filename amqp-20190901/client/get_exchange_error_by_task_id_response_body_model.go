// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExchangeErrorByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetExchangeErrorByTaskIdResponseBody
	GetCode() *int32
	SetData(v *GetExchangeErrorByTaskIdResponseBodyData) *GetExchangeErrorByTaskIdResponseBody
	GetData() *GetExchangeErrorByTaskIdResponseBodyData
	SetMessage(v string) *GetExchangeErrorByTaskIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExchangeErrorByTaskIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetExchangeErrorByTaskIdResponseBody
	GetSuccess() *bool
}

type GetExchangeErrorByTaskIdResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetExchangeErrorByTaskIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetExchangeErrorByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeErrorByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetExchangeErrorByTaskIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetExchangeErrorByTaskIdResponseBody) GetData() *GetExchangeErrorByTaskIdResponseBodyData {
	return s.Data
}

func (s *GetExchangeErrorByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExchangeErrorByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExchangeErrorByTaskIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetExchangeErrorByTaskIdResponseBody) SetCode(v int32) *GetExchangeErrorByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBody) SetData(v *GetExchangeErrorByTaskIdResponseBodyData) *GetExchangeErrorByTaskIdResponseBody {
	s.Data = v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBody) SetMessage(v string) *GetExchangeErrorByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBody) SetRequestId(v string) *GetExchangeErrorByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBody) SetSuccess(v bool) *GetExchangeErrorByTaskIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExchangeErrorByTaskIdResponseBodyData struct {
	CurrentPage *int32                                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      *GetExchangeErrorByTaskIdResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s GetExchangeErrorByTaskIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeErrorByTaskIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExchangeErrorByTaskIdResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetExchangeErrorByTaskIdResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetExchangeErrorByTaskIdResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetExchangeErrorByTaskIdResponseBodyData) GetVoList() *GetExchangeErrorByTaskIdResponseBodyDataVoList {
	return s.VoList
}

func (s *GetExchangeErrorByTaskIdResponseBodyData) SetCurrentPage(v int32) *GetExchangeErrorByTaskIdResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyData) SetPageSize(v int32) *GetExchangeErrorByTaskIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyData) SetTotalCount(v int64) *GetExchangeErrorByTaskIdResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyData) SetVoList(v *GetExchangeErrorByTaskIdResponseBodyDataVoList) *GetExchangeErrorByTaskIdResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExchangeErrorByTaskIdResponseBodyDataVoList struct {
	ExchangeErrorDO []*GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO `json:"ExchangeErrorDO,omitempty" xml:"ExchangeErrorDO,omitempty" type:"Repeated"`
}

func (s GetExchangeErrorByTaskIdResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeErrorByTaskIdResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoList) GetExchangeErrorDO() []*GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO {
	return s.ExchangeErrorDO
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoList) SetExchangeErrorDO(v []*GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) *GetExchangeErrorByTaskIdResponseBodyDataVoList {
	s.ExchangeErrorDO = v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoList) Validate() error {
	if s.ExchangeErrorDO != nil {
		for _, item := range s.ExchangeErrorDO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO struct {
	ErrorMessage *int32  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	ExchangeType *string `json:"ExchangeType,omitempty" xml:"ExchangeType,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VhostName    *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) GoString() string {
	return s.String()
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) GetErrorMessage() *int32 {
	return s.ErrorMessage
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) GetExchangeType() *string {
	return s.ExchangeType
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) GetVhostName() *string {
	return s.VhostName
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) SetErrorMessage(v int32) *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO {
	s.ErrorMessage = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) SetExchangeName(v string) *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO {
	s.ExchangeName = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) SetExchangeType(v string) *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO {
	s.ExchangeType = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) SetTaskId(v int64) *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO {
	s.TaskId = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) SetVhostName(v string) *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO {
	s.VhostName = &v
	return s
}

func (s *GetExchangeErrorByTaskIdResponseBodyDataVoListExchangeErrorDO) Validate() error {
	return dara.Validate(s)
}
