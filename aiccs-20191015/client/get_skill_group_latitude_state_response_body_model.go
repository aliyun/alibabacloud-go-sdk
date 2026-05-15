// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupLatitudeStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillGroupLatitudeStateResponseBody
	GetCode() *string
	SetData(v *GetSkillGroupLatitudeStateResponseBodyData) *GetSkillGroupLatitudeStateResponseBody
	GetData() *GetSkillGroupLatitudeStateResponseBodyData
	SetMessage(v string) *GetSkillGroupLatitudeStateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSkillGroupLatitudeStateResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSkillGroupLatitudeStateResponseBody
	GetSuccess() *string
}

type GetSkillGroupLatitudeStateResponseBody struct {
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSkillGroupLatitudeStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetSkillGroupLatitudeStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupLatitudeStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillGroupLatitudeStateResponseBody) GetData() *GetSkillGroupLatitudeStateResponseBodyData {
	return s.Data
}

func (s *GetSkillGroupLatitudeStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillGroupLatitudeStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillGroupLatitudeStateResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetCode(v string) *GetSkillGroupLatitudeStateResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetData(v *GetSkillGroupLatitudeStateResponseBodyData) *GetSkillGroupLatitudeStateResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetMessage(v string) *GetSkillGroupLatitudeStateResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetRequestId(v string) *GetSkillGroupLatitudeStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBody) SetSuccess(v string) *GetSkillGroupLatitudeStateResponseBody {
	s.Success = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillGroupLatitudeStateResponseBodyData struct {
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

func (s GetSkillGroupLatitudeStateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupLatitudeStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) SetPageNum(v int32) *GetSkillGroupLatitudeStateResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) SetPageSize(v int32) *GetSkillGroupLatitudeStateResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) SetRows(v string) *GetSkillGroupLatitudeStateResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) SetTotalNum(v int32) *GetSkillGroupLatitudeStateResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetSkillGroupLatitudeStateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
