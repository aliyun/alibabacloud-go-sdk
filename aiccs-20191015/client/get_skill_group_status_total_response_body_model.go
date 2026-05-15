// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupStatusTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillGroupStatusTotalResponseBody
	GetCode() *string
	SetData(v *GetSkillGroupStatusTotalResponseBodyData) *GetSkillGroupStatusTotalResponseBody
	GetData() *GetSkillGroupStatusTotalResponseBodyData
	SetMessage(v string) *GetSkillGroupStatusTotalResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSkillGroupStatusTotalResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSkillGroupStatusTotalResponseBody
	GetSuccess() *string
}

type GetSkillGroupStatusTotalResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSkillGroupStatusTotalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetSkillGroupStatusTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupStatusTotalResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillGroupStatusTotalResponseBody) GetData() *GetSkillGroupStatusTotalResponseBodyData {
	return s.Data
}

func (s *GetSkillGroupStatusTotalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillGroupStatusTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillGroupStatusTotalResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSkillGroupStatusTotalResponseBody) SetCode(v string) *GetSkillGroupStatusTotalResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBody) SetData(v *GetSkillGroupStatusTotalResponseBodyData) *GetSkillGroupStatusTotalResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBody) SetMessage(v string) *GetSkillGroupStatusTotalResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBody) SetRequestId(v string) *GetSkillGroupStatusTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBody) SetSuccess(v string) *GetSkillGroupStatusTotalResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillGroupStatusTotalResponseBodyData struct {
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

func (s GetSkillGroupStatusTotalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupStatusTotalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupStatusTotalResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *GetSkillGroupStatusTotalResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetSkillGroupStatusTotalResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetSkillGroupStatusTotalResponseBodyData) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *GetSkillGroupStatusTotalResponseBodyData) SetPageNum(v int64) *GetSkillGroupStatusTotalResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBodyData) SetPageSize(v int64) *GetSkillGroupStatusTotalResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBodyData) SetRows(v string) *GetSkillGroupStatusTotalResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBodyData) SetTotalNum(v int64) *GetSkillGroupStatusTotalResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupStatusTotalResponseBodyData) Validate() error {
	return dara.Validate(s)
}
