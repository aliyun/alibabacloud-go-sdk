// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetDocumentUrlResponseBody
	GetCost() *int64
	SetData(v string) *GetDocumentUrlResponseBody
	GetData() *string
	SetDataType(v string) *GetDocumentUrlResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetDocumentUrlResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetDocumentUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocumentUrlResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetDocumentUrlResponseBody
	GetTime() *string
}

type GetDocumentUrlResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// https://path_to_file
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
	// 66249B43-8C2B-5EE7-AE78-B382306621C6
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

func (s GetDocumentUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentUrlResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetDocumentUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetDocumentUrlResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetDocumentUrlResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDocumentUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocumentUrlResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetDocumentUrlResponseBody) SetCost(v int64) *GetDocumentUrlResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetData(v string) *GetDocumentUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetDataType(v string) *GetDocumentUrlResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetErrCode(v string) *GetDocumentUrlResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetMessage(v string) *GetDocumentUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetRequestId(v string) *GetDocumentUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetSuccess(v bool) *GetDocumentUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetTime(v string) *GetDocumentUrlResponseBody {
	s.Time = &v
	return s
}

func (s *GetDocumentUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
