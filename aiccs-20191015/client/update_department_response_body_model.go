// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDepartmentResponseBody
	GetCode() *string
	SetData(v bool) *UpdateDepartmentResponseBody
	GetData() *bool
	SetHttpStatusCode(v int64) *UpdateDepartmentResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *UpdateDepartmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDepartmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDepartmentResponseBody
	GetSuccess() *bool
}

type UpdateDepartmentResponseBody struct {
	// The status code. A value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - **true**: Succeeded
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98B032F5-6473-4EAC-8BA8-C28993513A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API was invoked successfully. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDepartmentResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateDepartmentResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateDepartmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDepartmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDepartmentResponseBody) SetCode(v string) *UpdateDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetData(v bool) *UpdateDepartmentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetHttpStatusCode(v int64) *UpdateDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetMessage(v string) *UpdateDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetRequestId(v string) *UpdateDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetSuccess(v bool) *UpdateDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
