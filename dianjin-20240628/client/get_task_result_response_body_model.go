// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetTaskResultResponseBody
	GetCost() *int64
	SetData(v map[string]interface{}) *GetTaskResultResponseBody
	GetData() map[string]interface{}
	SetDataType(v string) *GetTaskResultResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetTaskResultResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskResultResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetTaskResultResponseBody
	GetTime() *string
}

type GetTaskResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// {
	//
	//   "file_url": "https://finllmworks.oss-cn-zhangjiakou.aliyuncs.com/render_pdf/5336180997111160501.pdf"
	//
	// }
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
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
	// 9D5D6BB5-BEAE-53C8-A70A-7275CC1F856C
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

func (s GetTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetTaskResultResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetTaskResultResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetTaskResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskResultResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetTaskResultResponseBody) SetCost(v int64) *GetTaskResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetTaskResultResponseBody) SetData(v map[string]interface{}) *GetTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResultResponseBody) SetDataType(v string) *GetTaskResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetTaskResultResponseBody) SetErrCode(v string) *GetTaskResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetTaskResultResponseBody) SetMessage(v string) *GetTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResultResponseBody) SetRequestId(v string) *GetTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResultResponseBody) SetSuccess(v bool) *GetTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskResultResponseBody) SetTime(v string) *GetTaskResultResponseBody {
	s.Time = &v
	return s
}

func (s *GetTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}
