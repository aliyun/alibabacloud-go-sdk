// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMsgIdListByQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetMsgIdListByQueueResponseBody
	GetCode() *int32
	SetData(v *GetMsgIdListByQueueResponseBodyData) *GetMsgIdListByQueueResponseBody
	GetData() *GetMsgIdListByQueueResponseBodyData
	SetMessage(v string) *GetMsgIdListByQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMsgIdListByQueueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMsgIdListByQueueResponseBody
	GetSuccess() *bool
}

type GetMsgIdListByQueueResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMsgIdListByQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMsgIdListByQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMsgIdListByQueueResponseBody) GoString() string {
	return s.String()
}

func (s *GetMsgIdListByQueueResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetMsgIdListByQueueResponseBody) GetData() *GetMsgIdListByQueueResponseBodyData {
	return s.Data
}

func (s *GetMsgIdListByQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMsgIdListByQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMsgIdListByQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMsgIdListByQueueResponseBody) SetCode(v int32) *GetMsgIdListByQueueResponseBody {
	s.Code = &v
	return s
}

func (s *GetMsgIdListByQueueResponseBody) SetData(v *GetMsgIdListByQueueResponseBodyData) *GetMsgIdListByQueueResponseBody {
	s.Data = v
	return s
}

func (s *GetMsgIdListByQueueResponseBody) SetMessage(v string) *GetMsgIdListByQueueResponseBody {
	s.Message = &v
	return s
}

func (s *GetMsgIdListByQueueResponseBody) SetRequestId(v string) *GetMsgIdListByQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMsgIdListByQueueResponseBody) SetSuccess(v bool) *GetMsgIdListByQueueResponseBody {
	s.Success = &v
	return s
}

func (s *GetMsgIdListByQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMsgIdListByQueueResponseBodyData struct {
	CurrentPage *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VoList      []*string `json:"VoList,omitempty" xml:"VoList,omitempty" type:"Repeated"`
}

func (s GetMsgIdListByQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMsgIdListByQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMsgIdListByQueueResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetMsgIdListByQueueResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMsgIdListByQueueResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetMsgIdListByQueueResponseBodyData) GetVoList() []*string {
	return s.VoList
}

func (s *GetMsgIdListByQueueResponseBodyData) SetCurrentPage(v int32) *GetMsgIdListByQueueResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetMsgIdListByQueueResponseBodyData) SetPageSize(v int32) *GetMsgIdListByQueueResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMsgIdListByQueueResponseBodyData) SetTotalCount(v int32) *GetMsgIdListByQueueResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetMsgIdListByQueueResponseBodyData) SetVoList(v []*string) *GetMsgIdListByQueueResponseBodyData {
	s.VoList = v
	return s
}

func (s *GetMsgIdListByQueueResponseBodyData) Validate() error {
	return dara.Validate(s)
}
