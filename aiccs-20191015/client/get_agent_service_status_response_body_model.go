// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentServiceStatusResponseBody
	GetCode() *string
	SetData(v *GetAgentServiceStatusResponseBodyData) *GetAgentServiceStatusResponseBody
	GetData() *GetAgentServiceStatusResponseBodyData
	SetMessage(v string) *GetAgentServiceStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentServiceStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentServiceStatusResponseBody
	GetSuccess() *bool
}

type GetAgentServiceStatusResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAgentServiceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentServiceStatusResponseBody) GetData() *GetAgentServiceStatusResponseBodyData {
	return s.Data
}

func (s *GetAgentServiceStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentServiceStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentServiceStatusResponseBody) SetCode(v string) *GetAgentServiceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentServiceStatusResponseBody) SetData(v *GetAgentServiceStatusResponseBodyData) *GetAgentServiceStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentServiceStatusResponseBody) SetMessage(v string) *GetAgentServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentServiceStatusResponseBody) SetRequestId(v string) *GetAgentServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentServiceStatusResponseBody) SetSuccess(v bool) *GetAgentServiceStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentServiceStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentServiceStatusResponseBodyData struct {
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 2000
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {"online_40s_transfer_ready_cnt":81,"minute_id":"-1","online_unsatis_cnt":0,"online_simple_cnt":0,"average_queue_time":-1,"service_pickup":"2086","total_waiting_time":"981","online_service_time_len":58208,"online_direct_give_up_len":0,"break_ratio":"2%"}
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 4
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetAgentServiceStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentServiceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentServiceStatusResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *GetAgentServiceStatusResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAgentServiceStatusResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetAgentServiceStatusResponseBodyData) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *GetAgentServiceStatusResponseBodyData) SetPageNum(v int64) *GetAgentServiceStatusResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetAgentServiceStatusResponseBodyData) SetPageSize(v int64) *GetAgentServiceStatusResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentServiceStatusResponseBodyData) SetRows(v string) *GetAgentServiceStatusResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetAgentServiceStatusResponseBodyData) SetTotalNum(v int64) *GetAgentServiceStatusResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetAgentServiceStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
