// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLibraryChatGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *RunLibraryChatGenerationResponseBody
	GetCost() *int64
	SetData(v interface{}) *RunLibraryChatGenerationResponseBody
	GetData() interface{}
	SetDataType(v string) *RunLibraryChatGenerationResponseBody
	GetDataType() *string
	SetErrCode(v string) *RunLibraryChatGenerationResponseBody
	GetErrCode() *string
	SetMessage(v string) *RunLibraryChatGenerationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunLibraryChatGenerationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunLibraryChatGenerationResponseBody
	GetSuccess() *bool
	SetTime(v string) *RunLibraryChatGenerationResponseBody
	GetTime() *string
}

type RunLibraryChatGenerationResponseBody struct {
	// example:
	//
	// null
	Cost *int64      `json:"cost,omitempty" xml:"cost,omitempty"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
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

func (s RunLibraryChatGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunLibraryChatGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *RunLibraryChatGenerationResponseBody) GetData() interface{} {
	return s.Data
}

func (s *RunLibraryChatGenerationResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *RunLibraryChatGenerationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RunLibraryChatGenerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunLibraryChatGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunLibraryChatGenerationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunLibraryChatGenerationResponseBody) GetTime() *string {
	return s.Time
}

func (s *RunLibraryChatGenerationResponseBody) SetCost(v int64) *RunLibraryChatGenerationResponseBody {
	s.Cost = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetData(v interface{}) *RunLibraryChatGenerationResponseBody {
	s.Data = v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetDataType(v string) *RunLibraryChatGenerationResponseBody {
	s.DataType = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetErrCode(v string) *RunLibraryChatGenerationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetMessage(v string) *RunLibraryChatGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetRequestId(v string) *RunLibraryChatGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetSuccess(v bool) *RunLibraryChatGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetTime(v string) *RunLibraryChatGenerationResponseBody {
	s.Time = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) Validate() error {
	return dara.Validate(s)
}
