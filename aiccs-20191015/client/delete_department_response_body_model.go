// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDepartmentResponseBody
	GetCode() *string
	SetData(v bool) *DeleteDepartmentResponseBody
	GetData() *bool
	SetHttpStatusCode(v int64) *DeleteDepartmentResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *DeleteDepartmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDepartmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDepartmentResponseBody
	GetSuccess() *bool
}

type DeleteDepartmentResponseBody struct {
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
	// seccessful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98B032F5-6473-4EAC-8BA8-C28993513A1F
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

func (s DeleteDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDepartmentResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteDepartmentResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DeleteDepartmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDepartmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDepartmentResponseBody) SetCode(v string) *DeleteDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetData(v bool) *DeleteDepartmentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetHttpStatusCode(v int64) *DeleteDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetMessage(v string) *DeleteDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetRequestId(v string) *DeleteDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetSuccess(v bool) *DeleteDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
