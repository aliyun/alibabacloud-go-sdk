// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDepartmentalLatitudeAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDepartmentalLatitudeAgentStatusResponseBody
	GetCode() *string
	SetData(v *GetDepartmentalLatitudeAgentStatusResponseBodyData) *GetDepartmentalLatitudeAgentStatusResponseBody
	GetData() *GetDepartmentalLatitudeAgentStatusResponseBodyData
	SetMessage(v string) *GetDepartmentalLatitudeAgentStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDepartmentalLatitudeAgentStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDepartmentalLatitudeAgentStatusResponseBody
	GetSuccess() *bool
}

type GetDepartmentalLatitudeAgentStatusResponseBody struct {
	// example:
	//
	// 200
	Code *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDepartmentalLatitudeAgentStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDepartmentalLatitudeAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) GetData() *GetDepartmentalLatitudeAgentStatusResponseBodyData {
	return s.Data
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetCode(v string) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetData(v *GetDepartmentalLatitudeAgentStatusResponseBodyData) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetMessage(v string) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetRequestId(v string) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) SetSuccess(v bool) *GetDepartmentalLatitudeAgentStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDepartmentalLatitudeAgentStatusResponseBodyData struct {
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

func (s GetDepartmentalLatitudeAgentStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) SetPageNum(v int32) *GetDepartmentalLatitudeAgentStatusResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) SetPageSize(v int32) *GetDepartmentalLatitudeAgentStatusResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) SetRows(v string) *GetDepartmentalLatitudeAgentStatusResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) SetTotalNum(v int32) *GetDepartmentalLatitudeAgentStatusResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
