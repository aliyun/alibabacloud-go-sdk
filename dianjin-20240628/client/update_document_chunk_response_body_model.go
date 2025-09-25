// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocumentChunkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *UpdateDocumentChunkResponseBody
	GetCost() *int64
	SetData(v string) *UpdateDocumentChunkResponseBody
	GetData() *string
	SetDataType(v string) *UpdateDocumentChunkResponseBody
	GetDataType() *string
	SetErrCode(v string) *UpdateDocumentChunkResponseBody
	GetErrCode() *string
	SetMessage(v string) *UpdateDocumentChunkResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDocumentChunkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDocumentChunkResponseBody
	GetSuccess() *bool
	SetTime(v string) *UpdateDocumentChunkResponseBody
	GetTime() *string
}

type UpdateDocumentChunkResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// SUCCESS
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
	// 003D019A-1BB3-53EC-A0D2-CE76DA5D73B1
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

func (s UpdateDocumentChunkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocumentChunkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocumentChunkResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *UpdateDocumentChunkResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateDocumentChunkResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *UpdateDocumentChunkResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateDocumentChunkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDocumentChunkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDocumentChunkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDocumentChunkResponseBody) GetTime() *string {
	return s.Time
}

func (s *UpdateDocumentChunkResponseBody) SetCost(v int64) *UpdateDocumentChunkResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetData(v string) *UpdateDocumentChunkResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetDataType(v string) *UpdateDocumentChunkResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetErrCode(v string) *UpdateDocumentChunkResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetMessage(v string) *UpdateDocumentChunkResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetRequestId(v string) *UpdateDocumentChunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetSuccess(v bool) *UpdateDocumentChunkResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetTime(v string) *UpdateDocumentChunkResponseBody {
	s.Time = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) Validate() error {
	return dara.Validate(s)
}
