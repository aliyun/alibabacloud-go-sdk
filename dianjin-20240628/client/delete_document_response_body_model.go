// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *DeleteDocumentResponseBody
	GetCost() *int64
	SetData(v bool) *DeleteDocumentResponseBody
	GetData() *bool
	SetDataType(v string) *DeleteDocumentResponseBody
	GetDataType() *string
	SetErrCode(v string) *DeleteDocumentResponseBody
	GetErrCode() *string
	SetMessage(v string) *DeleteDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDocumentResponseBody
	GetSuccess() *bool
	SetTime(v string) *DeleteDocumentResponseBody
	GetTime() *string
}

type DeleteDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
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
	// 67C7021A-D268-553D-8C15-A087B9604028
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

func (s DeleteDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *DeleteDocumentResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDocumentResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *DeleteDocumentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDocumentResponseBody) GetTime() *string {
	return s.Time
}

func (s *DeleteDocumentResponseBody) SetCost(v int64) *DeleteDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetData(v bool) *DeleteDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetDataType(v string) *DeleteDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetErrCode(v string) *DeleteDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetMessage(v string) *DeleteDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetRequestId(v string) *DeleteDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetSuccess(v bool) *DeleteDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetTime(v string) *DeleteDocumentResponseBody {
	s.Time = &v
	return s
}

func (s *DeleteDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
