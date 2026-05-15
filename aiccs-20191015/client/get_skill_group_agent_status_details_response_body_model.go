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
	// example:
	//
	// 200
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSkillGroupAgentStatusDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 2000
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rows     *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
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
