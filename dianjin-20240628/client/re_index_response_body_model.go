// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *ReIndexResponseBody
	GetCost() *int64
	SetData(v string) *ReIndexResponseBody
	GetData() *string
	SetDataType(v string) *ReIndexResponseBody
	GetDataType() *string
	SetErrCode(v string) *ReIndexResponseBody
	GetErrCode() *string
	SetMessage(v string) *ReIndexResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReIndexResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReIndexResponseBody
	GetSuccess() *bool
	SetTime(v string) *ReIndexResponseBody
	GetTime() *string
}

type ReIndexResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// True
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
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s ReIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReIndexResponseBody) GoString() string {
	return s.String()
}

func (s *ReIndexResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *ReIndexResponseBody) GetData() *string {
	return s.Data
}

func (s *ReIndexResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *ReIndexResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ReIndexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReIndexResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReIndexResponseBody) GetTime() *string {
	return s.Time
}

func (s *ReIndexResponseBody) SetCost(v int64) *ReIndexResponseBody {
	s.Cost = &v
	return s
}

func (s *ReIndexResponseBody) SetData(v string) *ReIndexResponseBody {
	s.Data = &v
	return s
}

func (s *ReIndexResponseBody) SetDataType(v string) *ReIndexResponseBody {
	s.DataType = &v
	return s
}

func (s *ReIndexResponseBody) SetErrCode(v string) *ReIndexResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ReIndexResponseBody) SetMessage(v string) *ReIndexResponseBody {
	s.Message = &v
	return s
}

func (s *ReIndexResponseBody) SetRequestId(v string) *ReIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReIndexResponseBody) SetSuccess(v bool) *ReIndexResponseBody {
	s.Success = &v
	return s
}

func (s *ReIndexResponseBody) SetTime(v string) *ReIndexResponseBody {
	s.Time = &v
	return s
}

func (s *ReIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
