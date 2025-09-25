// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *UpdateDocumentResponseBody
	GetCost() *int64
	SetData(v string) *UpdateDocumentResponseBody
	GetData() *string
	SetDataType(v string) *UpdateDocumentResponseBody
	GetDataType() *string
	SetErrCode(v string) *UpdateDocumentResponseBody
	GetErrCode() *string
	SetMessage(v string) *UpdateDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDocumentResponseBody
	GetSuccess() *bool
	SetTime(v string) *UpdateDocumentResponseBody
	GetTime() *string
}

type UpdateDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// null
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
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
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

func (s UpdateDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocumentResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *UpdateDocumentResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateDocumentResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *UpdateDocumentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDocumentResponseBody) GetTime() *string {
	return s.Time
}

func (s *UpdateDocumentResponseBody) SetCost(v int64) *UpdateDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetData(v string) *UpdateDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetDataType(v string) *UpdateDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetErrCode(v string) *UpdateDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetMessage(v string) *UpdateDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetRequestId(v string) *UpdateDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetSuccess(v bool) *UpdateDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetTime(v string) *UpdateDocumentResponseBody {
	s.Time = &v
	return s
}

func (s *UpdateDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
