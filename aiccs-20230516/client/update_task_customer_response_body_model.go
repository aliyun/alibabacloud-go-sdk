// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskCustomerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateTaskCustomerResponseBody
	GetCode() *int64
	SetMessage(v string) *UpdateTaskCustomerResponseBody
	GetMessage() *string
	SetModel(v *UpdateTaskCustomerResponseBodyModel) *UpdateTaskCustomerResponseBody
	GetModel() *UpdateTaskCustomerResponseBodyModel
	SetRequestId(v string) *UpdateTaskCustomerResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateTaskCustomerResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *UpdateTaskCustomerResponseBody
	GetTimestamp() *int64
}

type UpdateTaskCustomerResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *UpdateTaskCustomerResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
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

func (s UpdateTaskCustomerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateTaskCustomerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTaskCustomerResponseBody) GetModel() *UpdateTaskCustomerResponseBodyModel {
	return s.Model
}

func (s *UpdateTaskCustomerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskCustomerResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateTaskCustomerResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UpdateTaskCustomerResponseBody) SetCode(v int64) *UpdateTaskCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetMessage(v string) *UpdateTaskCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetModel(v *UpdateTaskCustomerResponseBodyModel) *UpdateTaskCustomerResponseBody {
	s.Model = v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetRequestId(v string) *UpdateTaskCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetSuccess(v string) *UpdateTaskCustomerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetTimestamp(v int64) *UpdateTaskCustomerResponseBody {
	s.Timestamp = &v
	return s
}

func (s *UpdateTaskCustomerResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTaskCustomerResponseBodyModel struct {
	// 错误手机列表
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s UpdateTaskCustomerResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskCustomerResponseBodyModel) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerResponseBodyModel) GetUnHandleNumbers() []*string {
	return s.UnHandleNumbers
}

func (s *UpdateTaskCustomerResponseBodyModel) SetUnHandleNumbers(v []*string) *UpdateTaskCustomerResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

func (s *UpdateTaskCustomerResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
