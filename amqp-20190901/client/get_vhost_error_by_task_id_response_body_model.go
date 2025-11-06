// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVhostErrorByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetVhostErrorByTaskIdResponseBody
	GetCode() *int32
	SetData(v *GetVhostErrorByTaskIdResponseBodyData) *GetVhostErrorByTaskIdResponseBody
	GetData() *GetVhostErrorByTaskIdResponseBodyData
	SetMessage(v string) *GetVhostErrorByTaskIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVhostErrorByTaskIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVhostErrorByTaskIdResponseBody
	GetSuccess() *bool
}

type GetVhostErrorByTaskIdResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetVhostErrorByTaskIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVhostErrorByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVhostErrorByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetVhostErrorByTaskIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetVhostErrorByTaskIdResponseBody) GetData() *GetVhostErrorByTaskIdResponseBodyData {
	return s.Data
}

func (s *GetVhostErrorByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVhostErrorByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVhostErrorByTaskIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVhostErrorByTaskIdResponseBody) SetCode(v int32) *GetVhostErrorByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBody) SetData(v *GetVhostErrorByTaskIdResponseBodyData) *GetVhostErrorByTaskIdResponseBody {
	s.Data = v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBody) SetMessage(v string) *GetVhostErrorByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBody) SetRequestId(v string) *GetVhostErrorByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBody) SetSuccess(v bool) *GetVhostErrorByTaskIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVhostErrorByTaskIdResponseBodyData struct {
	CurrentPage *int32                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      *GetVhostErrorByTaskIdResponseBodyDataVoList `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Struct"`
}

func (s GetVhostErrorByTaskIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVhostErrorByTaskIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVhostErrorByTaskIdResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetVhostErrorByTaskIdResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetVhostErrorByTaskIdResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetVhostErrorByTaskIdResponseBodyData) GetVoList() *GetVhostErrorByTaskIdResponseBodyDataVoList {
	return s.VoList
}

func (s *GetVhostErrorByTaskIdResponseBodyData) SetCurrentPage(v int32) *GetVhostErrorByTaskIdResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBodyData) SetPageSize(v int32) *GetVhostErrorByTaskIdResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBodyData) SetTotalCount(v int64) *GetVhostErrorByTaskIdResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBodyData) SetVoList(v *GetVhostErrorByTaskIdResponseBodyDataVoList) *GetVhostErrorByTaskIdResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBodyData) Validate() error {
	if s.VoList != nil {
		if err := s.VoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVhostErrorByTaskIdResponseBodyDataVoList struct {
	VhostErrorDO []*GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO `json:"VhostErrorDO,omitempty" xml:"VhostErrorDO,omitempty" type:"Repeated"`
}

func (s GetVhostErrorByTaskIdResponseBodyDataVoList) String() string {
	return dara.Prettify(s)
}

func (s GetVhostErrorByTaskIdResponseBodyDataVoList) GoString() string {
	return s.String()
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoList) GetVhostErrorDO() []*GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO {
	return s.VhostErrorDO
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoList) SetVhostErrorDO(v []*GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) *GetVhostErrorByTaskIdResponseBodyDataVoList {
	s.VhostErrorDO = v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoList) Validate() error {
	if s.VhostErrorDO != nil {
		for _, item := range s.VhostErrorDO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO struct {
	ErrorMessage *int32  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	TaskId       *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VhostName    *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) String() string {
	return dara.Prettify(s)
}

func (s GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) GoString() string {
	return s.String()
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) GetErrorMessage() *int32 {
	return s.ErrorMessage
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) GetVhostName() *string {
	return s.VhostName
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) SetErrorMessage(v int32) *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO {
	s.ErrorMessage = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) SetTaskId(v int64) *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO {
	s.TaskId = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) SetVhostName(v string) *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO {
	s.VhostName = &v
	return s
}

func (s *GetVhostErrorByTaskIdResponseBodyDataVoListVhostErrorDO) Validate() error {
	return dara.Validate(s)
}
