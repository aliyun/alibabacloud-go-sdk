// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetTaskStatusResponseBody
	GetCost() *int64
	SetData(v string) *GetTaskStatusResponseBody
	GetData() *string
	SetDataType(v string) *GetTaskStatusResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetTaskStatusResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskStatusResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetTaskStatusResponseBody
	GetTime() *string
}

type GetTaskStatusResponseBody struct {
	Cost      *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	DataType  *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	ErrCode   *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	Time      *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetTaskStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *GetTaskStatusResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetTaskStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskStatusResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetTaskStatusResponseBody) SetCost(v int64) *GetTaskStatusResponseBody {
	s.Cost = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetData(v string) *GetTaskStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetDataType(v string) *GetTaskStatusResponseBody {
	s.DataType = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetErrCode(v string) *GetTaskStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetMessage(v string) *GetTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccess(v bool) *GetTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetTime(v string) *GetTaskStatusResponseBody {
	s.Time = &v
	return s
}

func (s *GetTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
