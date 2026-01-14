// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskRecoverCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *TaskRecoverCallResponseBody
	GetCode() *int64
	SetMessage(v string) *TaskRecoverCallResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *TaskRecoverCallResponseBody
	GetModel() map[string]interface{}
	SetRequestId(v string) *TaskRecoverCallResponseBody
	GetRequestId() *string
	SetSuccess(v string) *TaskRecoverCallResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *TaskRecoverCallResponseBody
	GetTimestamp() *int64
}

type TaskRecoverCallResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskRecoverCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TaskRecoverCallResponseBody) GoString() string {
	return s.String()
}

func (s *TaskRecoverCallResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TaskRecoverCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TaskRecoverCallResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *TaskRecoverCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TaskRecoverCallResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *TaskRecoverCallResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TaskRecoverCallResponseBody) SetCode(v int64) *TaskRecoverCallResponseBody {
	s.Code = &v
	return s
}

func (s *TaskRecoverCallResponseBody) SetMessage(v string) *TaskRecoverCallResponseBody {
	s.Message = &v
	return s
}

func (s *TaskRecoverCallResponseBody) SetModel(v map[string]interface{}) *TaskRecoverCallResponseBody {
	s.Model = v
	return s
}

func (s *TaskRecoverCallResponseBody) SetRequestId(v string) *TaskRecoverCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskRecoverCallResponseBody) SetSuccess(v string) *TaskRecoverCallResponseBody {
	s.Success = &v
	return s
}

func (s *TaskRecoverCallResponseBody) SetTimestamp(v int64) *TaskRecoverCallResponseBody {
	s.Timestamp = &v
	return s
}

func (s *TaskRecoverCallResponseBody) Validate() error {
	return dara.Validate(s)
}
