// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiDestinationName(v string) *UpdateApiDestinationResponseBody
	GetApiDestinationName() *string
	SetCode(v string) *UpdateApiDestinationResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateApiDestinationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApiDestinationResponseBody
	GetRequestId() *string
}

type UpdateApiDestinationResponseBody struct {
	// api-destination-name
	//
	// example:
	//
	// api-destination-name
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The response code. If the request is successful, Success is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request is successful, success is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 382E6272-8E9C-5681-AC96-A8AF0BFAC1A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApiDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationResponseBody) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *UpdateApiDestinationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApiDestinationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApiDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApiDestinationResponseBody) SetApiDestinationName(v string) *UpdateApiDestinationResponseBody {
	s.ApiDestinationName = &v
	return s
}

func (s *UpdateApiDestinationResponseBody) SetCode(v string) *UpdateApiDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApiDestinationResponseBody) SetMessage(v string) *UpdateApiDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApiDestinationResponseBody) SetRequestId(v string) *UpdateApiDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApiDestinationResponseBody) Validate() error {
	return dara.Validate(s)
}
