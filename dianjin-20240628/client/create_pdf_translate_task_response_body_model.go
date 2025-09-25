// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdfTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *CreatePdfTranslateTaskResponseBody
	GetCost() *int64
	SetData(v string) *CreatePdfTranslateTaskResponseBody
	GetData() *string
	SetDataType(v string) *CreatePdfTranslateTaskResponseBody
	GetDataType() *string
	SetErrCode(v string) *CreatePdfTranslateTaskResponseBody
	GetErrCode() *string
	SetMessage(v string) *CreatePdfTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePdfTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePdfTranslateTaskResponseBody
	GetSuccess() *bool
	SetTime(v string) *CreatePdfTranslateTaskResponseBody
	GetTime() *string
}

type CreatePdfTranslateTaskResponseBody struct {
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

func (s CreatePdfTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePdfTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePdfTranslateTaskResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *CreatePdfTranslateTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreatePdfTranslateTaskResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *CreatePdfTranslateTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreatePdfTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePdfTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePdfTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePdfTranslateTaskResponseBody) GetTime() *string {
	return s.Time
}

func (s *CreatePdfTranslateTaskResponseBody) SetCost(v int64) *CreatePdfTranslateTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetData(v string) *CreatePdfTranslateTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetDataType(v string) *CreatePdfTranslateTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetErrCode(v string) *CreatePdfTranslateTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetMessage(v string) *CreatePdfTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetRequestId(v string) *CreatePdfTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetSuccess(v bool) *CreatePdfTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetTime(v string) *CreatePdfTranslateTaskResponseBody {
	s.Time = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
