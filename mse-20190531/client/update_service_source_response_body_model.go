// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateServiceSourceResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateServiceSourceResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateServiceSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateServiceSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateServiceSourceResponseBody
	GetSuccess() *bool
}

type UpdateServiceSourceResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 63
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6AB2B0B6-4A86-54E2-A340-FC47A22EE659
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateServiceSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceSourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateServiceSourceResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateServiceSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateServiceSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateServiceSourceResponseBody) SetCode(v int32) *UpdateServiceSourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceSourceResponseBody) SetData(v int64) *UpdateServiceSourceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateServiceSourceResponseBody) SetHttpStatusCode(v int32) *UpdateServiceSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateServiceSourceResponseBody) SetMessage(v string) *UpdateServiceSourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceSourceResponseBody) SetRequestId(v string) *UpdateServiceSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceSourceResponseBody) SetSuccess(v bool) *UpdateServiceSourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateServiceSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
