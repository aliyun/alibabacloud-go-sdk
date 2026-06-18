// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupServiceCapabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillGroupServiceCapabilityResponseBody
	GetCode() *string
	SetData(v *GetSkillGroupServiceCapabilityResponseBodyData) *GetSkillGroupServiceCapabilityResponseBody
	GetData() *GetSkillGroupServiceCapabilityResponseBodyData
	SetMessage(v string) *GetSkillGroupServiceCapabilityResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSkillGroupServiceCapabilityResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSkillGroupServiceCapabilityResponseBody
	GetSuccess() *string
}

type GetSkillGroupServiceCapabilityResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data list.
	Data *GetSkillGroupServiceCapabilityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Status code description.
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

func (s GetSkillGroupServiceCapabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillGroupServiceCapabilityResponseBody) GetData() *GetSkillGroupServiceCapabilityResponseBodyData {
	return s.Data
}

func (s *GetSkillGroupServiceCapabilityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillGroupServiceCapabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillGroupServiceCapabilityResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetCode(v string) *GetSkillGroupServiceCapabilityResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetData(v *GetSkillGroupServiceCapabilityResponseBodyData) *GetSkillGroupServiceCapabilityResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetMessage(v string) *GetSkillGroupServiceCapabilityResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetRequestId(v string) *GetSkillGroupServiceCapabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBody) SetSuccess(v string) *GetSkillGroupServiceCapabilityResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillGroupServiceCapabilityResponseBodyData struct {
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
	// {
	//
	//       "tenant_id": "905",
	//
	//       "tenant_name": "非单元测试化BU",
	//
	//       "group_name": "-1",
	//
	//       "department_id": "-1",
	//
	//       "department_name": "-1",
	//
	//       "svc_online_cnt": 0,
	//
	//       "svc_ask_offline_servicer_num": 0,
	//
	//       "date_id": "20210326",
	//
	//       "group_id": "-1",
	//
	//       "svc_rest_cnt": 0,
	//
	//       "servicer_rest_cnt": 234,
	//
	//       "svc_offline_cnt": 0,
	//
	//       "svc_no_client_cnt": 0,
	//
	//       "svc_ask_rest_servicer_num": 0
	//
	// }
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 4
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetSkillGroupServiceCapabilityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupServiceCapabilityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) SetPageNum(v int32) *GetSkillGroupServiceCapabilityResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) SetPageSize(v int32) *GetSkillGroupServiceCapabilityResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) SetRows(v string) *GetSkillGroupServiceCapabilityResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) SetTotalNum(v int64) *GetSkillGroupServiceCapabilityResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupServiceCapabilityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
