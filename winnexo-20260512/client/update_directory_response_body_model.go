// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDirectoryResponseBody
	GetCode() *string
	SetDirectoryId(v string) *UpdateDirectoryResponseBody
	GetDirectoryId() *string
	SetMessage(v string) *UpdateDirectoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDirectoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDirectoryResponseBody
	GetSuccess() *bool
}

type UpdateDirectoryResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDirectoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDirectoryResponseBody) SetCode(v string) *UpdateDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDirectoryResponseBody) SetDirectoryId(v string) *UpdateDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDirectoryResponseBody) SetMessage(v string) *UpdateDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDirectoryResponseBody) SetRequestId(v string) *UpdateDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDirectoryResponseBody) SetSuccess(v bool) *UpdateDirectoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
