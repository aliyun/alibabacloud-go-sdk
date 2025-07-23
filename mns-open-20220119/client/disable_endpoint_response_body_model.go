// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DisableEndpointResponseBody
	GetCode() *int64
	SetMessage(v string) *DisableEndpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableEndpointResponseBody
	GetRequestId() *string
	SetStatus(v string) *DisableEndpointResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DisableEndpointResponseBody
	GetSuccess() *bool
}

type DisableEndpointResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DisableEndpointResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DisableEndpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableEndpointResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DisableEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableEndpointResponseBody) SetCode(v int64) *DisableEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *DisableEndpointResponseBody) SetMessage(v string) *DisableEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *DisableEndpointResponseBody) SetRequestId(v string) *DisableEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableEndpointResponseBody) SetStatus(v string) *DisableEndpointResponseBody {
	s.Status = &v
	return s
}

func (s *DisableEndpointResponseBody) SetSuccess(v bool) *DisableEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *DisableEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
