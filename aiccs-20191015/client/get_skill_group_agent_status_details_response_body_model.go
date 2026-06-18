// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupAgentStatusDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillGroupAgentStatusDetailsResponseBody
	GetCode() *string
	SetData(v *GetSkillGroupAgentStatusDetailsResponseBodyData) *GetSkillGroupAgentStatusDetailsResponseBody
	GetData() *GetSkillGroupAgentStatusDetailsResponseBodyData
	SetMessage(v string) *GetSkillGroupAgentStatusDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSkillGroupAgentStatusDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSkillGroupAgentStatusDetailsResponseBody
	GetSuccess() *string
}

type GetSkillGroupAgentStatusDetailsResponseBody struct {
	// The status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data list.
	Data *GetSkillGroupAgentStatusDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the API call succeeded. Valid values:
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

func (s GetSkillGroupAgentStatusDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) GetData() *GetSkillGroupAgentStatusDetailsResponseBodyData {
	return s.Data
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetCode(v string) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetData(v *GetSkillGroupAgentStatusDetailsResponseBodyData) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetMessage(v string) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetRequestId(v string) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) SetSuccess(v string) *GetSkillGroupAgentStatusDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillGroupAgentStatusDetailsResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of items per page.
	//
	// example:
	//
	// 2000
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A JSON string of type List<Map>.
	//
	// example:
	//
	// {"tenant_id": "905", "tenant_name": "非单元测试化BU", "department_id": "-1", "svc_avai_capacity": 7,      "svc_current_capacity": 0, "department_name": "-1",  "svc_online_cnt": 3, "svc_ask_offline_servicer_num": 0,  "date_id": "20210401",  "svc_max_capacity": 7, "svc_rest_cnt": 0, "current_speci_ability_num": 0, "servicer_rest_cnt": 122, "max_speci_ability_num": 3,      "svc_offline_cnt": 0,  "svc_no_client_cnt": 3, "svc_ask_rest_servicer_num": 0}
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 4
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetSkillGroupAgentStatusDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAgentStatusDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) SetPageNum(v int64) *GetSkillGroupAgentStatusDetailsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) SetPageSize(v int64) *GetSkillGroupAgentStatusDetailsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) SetRows(v string) *GetSkillGroupAgentStatusDetailsResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) SetTotalNum(v int64) *GetSkillGroupAgentStatusDetailsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupAgentStatusDetailsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
