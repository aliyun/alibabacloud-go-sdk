// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisposeServiceWorkOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisposeServiceWorkOrderResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DisposeServiceWorkOrderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisposeServiceWorkOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisposeServiceWorkOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisposeServiceWorkOrderResponseBody
	GetSuccess() *bool
}

type DisposeServiceWorkOrderResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message of the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ED520610-6231-5D80-BADD-A8CDC7BBC809
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisposeServiceWorkOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisposeServiceWorkOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DisposeServiceWorkOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisposeServiceWorkOrderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisposeServiceWorkOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisposeServiceWorkOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisposeServiceWorkOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisposeServiceWorkOrderResponseBody) SetCode(v string) *DisposeServiceWorkOrderResponseBody {
	s.Code = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetHttpStatusCode(v int32) *DisposeServiceWorkOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetMessage(v string) *DisposeServiceWorkOrderResponseBody {
	s.Message = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetRequestId(v string) *DisposeServiceWorkOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetSuccess(v bool) *DisposeServiceWorkOrderResponseBody {
	s.Success = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
