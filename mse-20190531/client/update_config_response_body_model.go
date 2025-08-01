// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateConfigResponseBody
	GetCode() *int32
	SetHttpStatusCode(v int32) *UpdateConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateConfigResponseBody
	GetSuccess() *bool
}

type UpdateConfigResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8BD1E58D-0755-42AC-A599-E6B55112****
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

func (s UpdateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConfigResponseBody) SetCode(v int32) *UpdateConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConfigResponseBody) SetHttpStatusCode(v int32) *UpdateConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConfigResponseBody) SetMessage(v string) *UpdateConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConfigResponseBody) SetRequestId(v string) *UpdateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigResponseBody) SetSuccess(v bool) *UpdateConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
