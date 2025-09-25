// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokePluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *InvokePluginResponseBody
	GetCost() *int64
	SetData(v map[string]interface{}) *InvokePluginResponseBody
	GetData() map[string]interface{}
	SetDataType(v string) *InvokePluginResponseBody
	GetDataType() *string
	SetErrCode(v string) *InvokePluginResponseBody
	GetErrCode() *string
	SetMessage(v string) *InvokePluginResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvokePluginResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvokePluginResponseBody
	GetSuccess() *bool
	SetTime(v string) *InvokePluginResponseBody
	GetTime() *string
}

type InvokePluginResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// {\\"jobWaiting\\": [0, 0], \\"timestamps\\": [1713383820, 1713383880], \\"jobUsage\\": [0, 0], \\"quotaUsage\\": [123, 32]}
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
	// 915AAAB9-4908-5224-9E53-9E9D7D0AA94B
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

func (s InvokePluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokePluginResponseBody) GoString() string {
	return s.String()
}

func (s *InvokePluginResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *InvokePluginResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *InvokePluginResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *InvokePluginResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *InvokePluginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvokePluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokePluginResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvokePluginResponseBody) GetTime() *string {
	return s.Time
}

func (s *InvokePluginResponseBody) SetCost(v int64) *InvokePluginResponseBody {
	s.Cost = &v
	return s
}

func (s *InvokePluginResponseBody) SetData(v map[string]interface{}) *InvokePluginResponseBody {
	s.Data = v
	return s
}

func (s *InvokePluginResponseBody) SetDataType(v string) *InvokePluginResponseBody {
	s.DataType = &v
	return s
}

func (s *InvokePluginResponseBody) SetErrCode(v string) *InvokePluginResponseBody {
	s.ErrCode = &v
	return s
}

func (s *InvokePluginResponseBody) SetMessage(v string) *InvokePluginResponseBody {
	s.Message = &v
	return s
}

func (s *InvokePluginResponseBody) SetRequestId(v string) *InvokePluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokePluginResponseBody) SetSuccess(v bool) *InvokePluginResponseBody {
	s.Success = &v
	return s
}

func (s *InvokePluginResponseBody) SetTime(v string) *InvokePluginResponseBody {
	s.Time = &v
	return s
}

func (s *InvokePluginResponseBody) Validate() error {
	return dara.Validate(s)
}
