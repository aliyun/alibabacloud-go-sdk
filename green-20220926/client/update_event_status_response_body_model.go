// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEventStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEventStatusResponseBody
	GetSuccess() *bool
}

type UpdateEventStatusResponseBody struct {
	// The ID assigned by the backend to uniquely identify a request. You can use this ID for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEventStatusResponseBody) SetRequestId(v string) *UpdateEventStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventStatusResponseBody) SetSuccess(v bool) *UpdateEventStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEventStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
