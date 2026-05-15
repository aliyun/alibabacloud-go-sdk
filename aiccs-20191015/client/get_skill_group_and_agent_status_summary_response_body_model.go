// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupAndAgentStatusSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody
	GetCode() *string
	SetData(v *GetSkillGroupAndAgentStatusSummaryResponseBodyData) *GetSkillGroupAndAgentStatusSummaryResponseBody
	GetData() *GetSkillGroupAndAgentStatusSummaryResponseBodyData
	SetMessage(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody
	GetSuccess() *string
}

type GetSkillGroupAndAgentStatusSummaryResponseBody struct {
	// example:
	//
	// 200
	Code *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSkillGroupAndAgentStatusSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// message
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

func (s GetSkillGroupAndAgentStatusSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) GetData() *GetSkillGroupAndAgentStatusSummaryResponseBodyData {
	return s.Data
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetCode(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetData(v *GetSkillGroupAndAgentStatusSummaryResponseBodyData) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetMessage(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetRequestId(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) SetSuccess(v string) *GetSkillGroupAndAgentStatusSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillGroupAndAgentStatusSummaryResponseBodyData struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 2000
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rows     *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// example:
	//
	// 4
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetSkillGroupAndAgentStatusSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) SetPageNum(v int32) *GetSkillGroupAndAgentStatusSummaryResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) SetPageSize(v int32) *GetSkillGroupAndAgentStatusSummaryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) SetRows(v string) *GetSkillGroupAndAgentStatusSummaryResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) SetTotalNum(v int32) *GetSkillGroupAndAgentStatusSummaryResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
