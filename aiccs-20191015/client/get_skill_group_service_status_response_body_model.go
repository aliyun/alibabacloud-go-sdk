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
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSkillGroupServiceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
