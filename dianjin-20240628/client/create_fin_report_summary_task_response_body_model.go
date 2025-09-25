// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFinReportSummaryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *CreateFinReportSummaryTaskResponseBody
	GetCost() *int64
	SetData(v string) *CreateFinReportSummaryTaskResponseBody
	GetData() *string
	SetDataType(v string) *CreateFinReportSummaryTaskResponseBody
	GetDataType() *string
	SetErrCode(v string) *CreateFinReportSummaryTaskResponseBody
	GetErrCode() *string
	SetMessage(v string) *CreateFinReportSummaryTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFinReportSummaryTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFinReportSummaryTaskResponseBody
	GetSuccess() *bool
	SetTime(v string) *CreateFinReportSummaryTaskResponseBody
	GetTime() *string
}

type CreateFinReportSummaryTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 3284627354
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateFinReportSummaryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFinReportSummaryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFinReportSummaryTaskResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *CreateFinReportSummaryTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateFinReportSummaryTaskResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *CreateFinReportSummaryTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateFinReportSummaryTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFinReportSummaryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFinReportSummaryTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFinReportSummaryTaskResponseBody) GetTime() *string {
	return s.Time
}

func (s *CreateFinReportSummaryTaskResponseBody) SetCost(v int64) *CreateFinReportSummaryTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetData(v string) *CreateFinReportSummaryTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetDataType(v string) *CreateFinReportSummaryTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetErrCode(v string) *CreateFinReportSummaryTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetMessage(v string) *CreateFinReportSummaryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetRequestId(v string) *CreateFinReportSummaryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetSuccess(v bool) *CreateFinReportSummaryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetTime(v string) *CreateFinReportSummaryTaskResponseBody {
	s.Time = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
