// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnnualDocSummaryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *CreateAnnualDocSummaryTaskResponseBody
	GetCost() *int64
	SetData(v string) *CreateAnnualDocSummaryTaskResponseBody
	GetData() *string
	SetDataType(v string) *CreateAnnualDocSummaryTaskResponseBody
	GetDataType() *string
	SetErrCode(v string) *CreateAnnualDocSummaryTaskResponseBody
	GetErrCode() *string
	SetMessage(v string) *CreateAnnualDocSummaryTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAnnualDocSummaryTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAnnualDocSummaryTaskResponseBody
	GetSuccess() *bool
	SetTime(v string) *CreateAnnualDocSummaryTaskResponseBody
	GetTime() *string
}

type CreateAnnualDocSummaryTaskResponseBody struct {
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
	// 32FFC91D-0A9F-585A-B84F-8A54C5187035
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateAnnualDocSummaryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAnnualDocSummaryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnnualDocSummaryTaskResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *CreateAnnualDocSummaryTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateAnnualDocSummaryTaskResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *CreateAnnualDocSummaryTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateAnnualDocSummaryTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAnnualDocSummaryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAnnualDocSummaryTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAnnualDocSummaryTaskResponseBody) GetTime() *string {
	return s.Time
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetCost(v int64) *CreateAnnualDocSummaryTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetData(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetDataType(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetErrCode(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetMessage(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetRequestId(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetSuccess(v bool) *CreateAnnualDocSummaryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetTime(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.Time = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
