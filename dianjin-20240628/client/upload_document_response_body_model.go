// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *UploadDocumentResponseBody
	GetCost() *int64
	SetData(v string) *UploadDocumentResponseBody
	GetData() *string
	SetDataType(v string) *UploadDocumentResponseBody
	GetDataType() *string
	SetErrCode(v string) *UploadDocumentResponseBody
	GetErrCode() *string
	SetMessage(v string) *UploadDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadDocumentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadDocumentResponseBody
	GetSuccess() *bool
	SetTime(v string) *UploadDocumentResponseBody
	GetTime() *string
}

type UploadDocumentResponseBody struct {
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
	// ff3fef67-48d9-4379-a237-9ba8143fe739
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

func (s UploadDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDocumentResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *UploadDocumentResponseBody) GetData() *string {
	return s.Data
}

func (s *UploadDocumentResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *UploadDocumentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UploadDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadDocumentResponseBody) GetTime() *string {
	return s.Time
}

func (s *UploadDocumentResponseBody) SetCost(v int64) *UploadDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *UploadDocumentResponseBody) SetData(v string) *UploadDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *UploadDocumentResponseBody) SetDataType(v string) *UploadDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *UploadDocumentResponseBody) SetErrCode(v string) *UploadDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UploadDocumentResponseBody) SetMessage(v string) *UploadDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDocumentResponseBody) SetRequestId(v string) *UploadDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDocumentResponseBody) SetSuccess(v bool) *UploadDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDocumentResponseBody) SetTime(v string) *UploadDocumentResponseBody {
	s.Time = &v
	return s
}

func (s *UploadDocumentResponseBody) Validate() error {
	return dara.Validate(s)
}
