// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteClientResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteClientResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteClientResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteClientResponseBody
	GetSuccess() *bool
}

type DeleteClientResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C51A9094-64B7-5DC0-B9FE-5FC1AC7E081D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteClientResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClientResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteClientResponseBody) SetCode(v string) *DeleteClientResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteClientResponseBody) SetMessage(v string) *DeleteClientResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClientResponseBody) SetRequestId(v string) *DeleteClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClientResponseBody) SetSuccess(v bool) *DeleteClientResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteClientResponseBody) Validate() error {
	return dara.Validate(s)
}
