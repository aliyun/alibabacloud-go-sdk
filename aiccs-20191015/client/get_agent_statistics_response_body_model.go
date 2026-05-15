// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentStatisticsResponseBody
	GetCode() *string
	SetData(v *GetAgentStatisticsResponseBodyData) *GetAgentStatisticsResponseBody
	GetData() *GetAgentStatisticsResponseBodyData
	SetMessage(v string) *GetAgentStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetAgentStatisticsResponseBody
	GetSuccess() *string
}

type GetAgentStatisticsResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAgentStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentStatisticsResponseBody) GetData() *GetAgentStatisticsResponseBodyData {
	return s.Data
}

func (s *GetAgentStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentStatisticsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetAgentStatisticsResponseBody) SetCode(v string) *GetAgentStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentStatisticsResponseBody) SetData(v *GetAgentStatisticsResponseBodyData) *GetAgentStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentStatisticsResponseBody) SetMessage(v string) *GetAgentStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentStatisticsResponseBody) SetRequestId(v string) *GetAgentStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentStatisticsResponseBody) SetSuccess(v string) *GetAgentStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentStatisticsResponseBodyData struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 2000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {"n_resttype_1":15,"t_outcall_speak":829747,"average_incoming_time":"8451.29","n_resttype_3":0,"minute_id":"-1","n_conference_speak":0,"n_resttype_2":0,"n_resttype_5":0,"n_resttype_4":0,"n_resttype_7":0,"n_resttype_6":0,"n_resttype_9":0,"n_resttype_8":0,"n_outcall_dial":25,"total_break_time":"58555","n_internal_speak":0,"n_send_step_transfer":7,"n_consulted_internal_speak":0}
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 4
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetAgentStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetAgentStatisticsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentStatisticsResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetAgentStatisticsResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetAgentStatisticsResponseBodyData) SetPageNum(v int32) *GetAgentStatisticsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetAgentStatisticsResponseBodyData) SetPageSize(v int32) *GetAgentStatisticsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentStatisticsResponseBodyData) SetRows(v string) *GetAgentStatisticsResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetAgentStatisticsResponseBodyData) SetTotalNum(v int32) *GetAgentStatisticsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetAgentStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
