// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDetailReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentDetailReportResponseBody
	GetCode() *string
	SetData(v *GetAgentDetailReportResponseBodyData) *GetAgentDetailReportResponseBody
	GetData() *GetAgentDetailReportResponseBodyData
	SetMessage(v string) *GetAgentDetailReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentDetailReportResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetAgentDetailReportResponseBody
	GetSuccess() *string
}

type GetAgentDetailReportResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAgentDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetAgentDetailReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentDetailReportResponseBody) GetData() *GetAgentDetailReportResponseBodyData {
	return s.Data
}

func (s *GetAgentDetailReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentDetailReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentDetailReportResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetAgentDetailReportResponseBody) SetCode(v string) *GetAgentDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentDetailReportResponseBody) SetData(v *GetAgentDetailReportResponseBodyData) *GetAgentDetailReportResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentDetailReportResponseBody) SetMessage(v string) *GetAgentDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentDetailReportResponseBody) SetRequestId(v string) *GetAgentDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentDetailReportResponseBody) SetSuccess(v string) *GetAgentDetailReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentDetailReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentDetailReportResponseBodyData struct {
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
	// {"t_wait":379322.0,"hotline_time_outcall_avg":"32.00","n_ringing":0,"t_outbound_40":0.0,"hotline_time_incall_avg":-1,"t_calldialing":0.0,"n_inbound":276,"servicer_id":"-1","call_in_sep_sat_cnt":13,"request_cnt":231,"n_not_ready_99":811,"t_work_outbound":49338.0,"hotline_rate_handle_in_60s":"96%","n_not_ready_login":811,"t_not_ready_login":0.0,"n_work_inbound":176,"}
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 4
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetAgentDetailReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentDetailReportResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *GetAgentDetailReportResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAgentDetailReportResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetAgentDetailReportResponseBodyData) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *GetAgentDetailReportResponseBodyData) SetPageNum(v int64) *GetAgentDetailReportResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetAgentDetailReportResponseBodyData) SetPageSize(v int64) *GetAgentDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentDetailReportResponseBodyData) SetRows(v string) *GetAgentDetailReportResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetAgentDetailReportResponseBodyData) SetTotalNum(v int64) *GetAgentDetailReportResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetAgentDetailReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
