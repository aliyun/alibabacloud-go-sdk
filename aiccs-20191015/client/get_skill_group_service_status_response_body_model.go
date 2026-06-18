// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillGroupServiceStatusResponseBody
	GetCode() *string
	SetData(v *GetSkillGroupServiceStatusResponseBodyData) *GetSkillGroupServiceStatusResponseBody
	GetData() *GetSkillGroupServiceStatusResponseBodyData
	SetMessage(v string) *GetSkillGroupServiceStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSkillGroupServiceStatusResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSkillGroupServiceStatusResponseBody
	GetSuccess() *string
}

type GetSkillGroupServiceStatusResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of data.
	Data *GetSkillGroupServiceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API was invoked successfully. Valid values:
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

func (s GetSkillGroupServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillGroupServiceStatusResponseBody) GetData() *GetSkillGroupServiceStatusResponseBodyData {
	return s.Data
}

func (s *GetSkillGroupServiceStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillGroupServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillGroupServiceStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSkillGroupServiceStatusResponseBody) SetCode(v string) *GetSkillGroupServiceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBody) SetData(v *GetSkillGroupServiceStatusResponseBodyData) *GetSkillGroupServiceStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBody) SetMessage(v string) *GetSkillGroupServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBody) SetRequestId(v string) *GetSkillGroupServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBody) SetSuccess(v string) *GetSkillGroupServiceStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillGroupServiceStatusResponseBodyData struct {
	// Current page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Page size. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 2000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Information in the form of a JSON string of type List<Map>.
	//
	// example:
	//
	// {"online_40s_transfer_ready_cnt":382,"minute_id":"-1","online_unsatis_cnt":0,"online_simple_cnt":0,"average_queue_time":"0.39","service_pickup":"7752","online_service_time_len":220753,"online_direct_give_up_len":1187,"channel_instance_name":"-1","servicer_id":"-1","tenant_name":"非单元测试化BU","group_name":"-1","online_30s_transfer_ready_cnt":382}
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 4
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetSkillGroupServiceStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceStatusResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetSkillGroupServiceStatusResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupServiceStatusResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetSkillGroupServiceStatusResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetSkillGroupServiceStatusResponseBodyData) SetPageNum(v int32) *GetSkillGroupServiceStatusResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBodyData) SetPageSize(v int32) *GetSkillGroupServiceStatusResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBodyData) SetRows(v string) *GetSkillGroupServiceStatusResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBodyData) SetTotalNum(v int32) *GetSkillGroupServiceStatusResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupServiceStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
