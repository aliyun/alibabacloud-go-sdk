// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocsSummaryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *CreateDocsSummaryTaskResponseBody
	GetCost() *int64
	SetData(v string) *CreateDocsSummaryTaskResponseBody
	GetData() *string
	SetDataType(v string) *CreateDocsSummaryTaskResponseBody
	GetDataType() *string
	SetErrCode(v string) *CreateDocsSummaryTaskResponseBody
	GetErrCode() *string
	SetMessage(v string) *CreateDocsSummaryTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDocsSummaryTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDocsSummaryTaskResponseBody
	GetSuccess() *bool
	SetTime(v string) *CreateDocsSummaryTaskResponseBody
	GetTime() *string
}

type CreateDocsSummaryTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 765675376
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

func (s CreateDocsSummaryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDocsSummaryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocsSummaryTaskResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *CreateDocsSummaryTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateDocsSummaryTaskResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *CreateDocsSummaryTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateDocsSummaryTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDocsSummaryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDocsSummaryTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDocsSummaryTaskResponseBody) GetTime() *string {
	return s.Time
}

func (s *CreateDocsSummaryTaskResponseBody) SetCost(v int64) *CreateDocsSummaryTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetData(v string) *CreateDocsSummaryTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetDataType(v string) *CreateDocsSummaryTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetErrCode(v string) *CreateDocsSummaryTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetMessage(v string) *CreateDocsSummaryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetRequestId(v string) *CreateDocsSummaryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetSuccess(v bool) *CreateDocsSummaryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetTime(v string) *CreateDocsSummaryTaskResponseBody {
	s.Time = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
