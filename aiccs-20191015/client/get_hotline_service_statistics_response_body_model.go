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
	// example:
	//
	// 200
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetHotlineServiceStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
