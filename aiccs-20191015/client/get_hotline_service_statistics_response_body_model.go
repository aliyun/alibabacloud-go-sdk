// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineServiceStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineServiceStatisticsResponseBody
	GetCode() *string
	SetData(v *GetHotlineServiceStatisticsResponseBodyData) *GetHotlineServiceStatisticsResponseBody
	GetData() *GetHotlineServiceStatisticsResponseBodyData
	SetMessage(v string) *GetHotlineServiceStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineServiceStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetHotlineServiceStatisticsResponseBody
	GetSuccess() *string
}

type GetHotlineServiceStatisticsResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data list.
	Data *GetHotlineServiceStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Status description.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API invocation succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineServiceStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineServiceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineServiceStatisticsResponseBody) GetData() *GetHotlineServiceStatisticsResponseBodyData {
	return s.Data
}

func (s *GetHotlineServiceStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineServiceStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineServiceStatisticsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetHotlineServiceStatisticsResponseBody) SetCode(v string) *GetHotlineServiceStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBody) SetData(v *GetHotlineServiceStatisticsResponseBodyData) *GetHotlineServiceStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBody) SetMessage(v string) *GetHotlineServiceStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBody) SetRequestId(v string) *GetHotlineServiceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBody) SetSuccess(v string) *GetHotlineServiceStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotlineServiceStatisticsResponseBodyData struct {
	// Current page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Page size.
	//
	// example:
	//
	// 2000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A JSON string of type List<Map>.
	//
	// example:
	//
	// {"tenant_id":"905","n_resttype_now":0,"tenant_name":"非单元测试化BU","n_resttype1_now":0,"group_name":"-1","department_id":"-1","department_name":"-1","n_resttype98_now":0,"n_online_now":0,"date_id":"20210401","n_resttype3_now":0,"n_resttype5_now":0,"n_busy_now":0,"n_resttype2_now":0,"group_id":"-1","n_idle_now":0,"n_resttype4_now":0,"n_ack_now":0,"n_resttype99_now":0}
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 4
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetHotlineServiceStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineServiceStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetHotlineServiceStatisticsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHotlineServiceStatisticsResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetHotlineServiceStatisticsResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetHotlineServiceStatisticsResponseBodyData) SetPageNum(v int32) *GetHotlineServiceStatisticsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBodyData) SetPageSize(v int32) *GetHotlineServiceStatisticsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBodyData) SetRows(v string) *GetHotlineServiceStatisticsResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBodyData) SetTotalNum(v int32) *GetHotlineServiceStatisticsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
