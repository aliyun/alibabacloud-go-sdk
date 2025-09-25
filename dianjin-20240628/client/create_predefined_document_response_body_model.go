// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePredefinedDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *CreatePredefinedDocumentResponseBody
	GetCost() *int64
	SetData(v string) *CreatePredefinedDocumentResponseBody
	GetData() *string
	SetDataType(v string) *CreatePredefinedDocumentResponseBody
	GetDataType() *string
	SetErrCode(v string) *CreatePredefinedDocumentResponseBody
	GetErrCode() *string
	SetMessage(v string) *CreatePredefinedDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePredefinedDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePredefinedDocumentResponseBody
	GetSuccess() *bool
	SetTime(v string) *CreatePredefinedDocumentResponseBody
	GetTime() *string
}

type CreatePredefinedDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 1782981430906818562
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
	// 0a06dfe617018288881568684e2937
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

func (s CreatePredefinedDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePredefinedDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *CreatePredefinedDocumentResponseBody) GetData() *string {
	return s.Data
}

func (s *CreatePredefinedDocumentResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *CreatePredefinedDocumentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreatePredefinedDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePredefinedDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePredefinedDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePredefinedDocumentResponseBody) GetTime() *string {
	return s.Time
}

func (s *CreatePredefinedDocumentResponseBody) SetCost(v int64) *CreatePredefinedDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetData(v string) *CreatePredefinedDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetDataType(v string) *CreatePredefinedDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetErrCode(v string) *CreatePredefinedDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetMessage(v string) *CreatePredefinedDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetRequestId(v string) *CreatePredefinedDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetSuccess(v bool) *CreatePredefinedDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetTime(v string) *CreatePredefinedDocumentResponseBody {
	s.Time = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
