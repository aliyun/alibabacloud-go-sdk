// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *RebuildTaskResponseBody
	GetCost() *int64
	SetData(v []map[string]interface{}) *RebuildTaskResponseBody
	GetData() []map[string]interface{}
	SetDataType(v string) *RebuildTaskResponseBody
	GetDataType() *string
	SetErrCode(v string) *RebuildTaskResponseBody
	GetErrCode() *string
	SetMessage(v string) *RebuildTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *RebuildTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RebuildTaskResponseBody
	GetSuccess() *bool
	SetTime(v string) *RebuildTaskResponseBody
	GetTime() *string
}

type RebuildTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64                   `json:"cost,omitempty" xml:"cost,omitempty"`
	Data []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
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

func (s RebuildTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildTaskResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *RebuildTaskResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *RebuildTaskResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *RebuildTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RebuildTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RebuildTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebuildTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RebuildTaskResponseBody) GetTime() *string {
	return s.Time
}

func (s *RebuildTaskResponseBody) SetCost(v int64) *RebuildTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *RebuildTaskResponseBody) SetData(v []map[string]interface{}) *RebuildTaskResponseBody {
	s.Data = v
	return s
}

func (s *RebuildTaskResponseBody) SetDataType(v string) *RebuildTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *RebuildTaskResponseBody) SetErrCode(v string) *RebuildTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RebuildTaskResponseBody) SetMessage(v string) *RebuildTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RebuildTaskResponseBody) SetRequestId(v string) *RebuildTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildTaskResponseBody) SetSuccess(v bool) *RebuildTaskResponseBody {
	s.Success = &v
	return s
}

func (s *RebuildTaskResponseBody) SetTime(v string) *RebuildTaskResponseBody {
	s.Time = &v
	return s
}

func (s *RebuildTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
