// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmbeddedStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEmbeddedStatusResponseBody
	GetRequestId() *string
	SetResult(v int32) *UpdateEmbeddedStatusResponseBody
	GetResult() *int32
	SetSuccess(v bool) *UpdateEmbeddedStatusResponseBody
	GetSuccess() *bool
}

type UpdateEmbeddedStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D78*********DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of works that are opened or closed.
	//
	// example:
	//
	// 1
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEmbeddedStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmbeddedStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEmbeddedStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEmbeddedStatusResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *UpdateEmbeddedStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEmbeddedStatusResponseBody) SetRequestId(v string) *UpdateEmbeddedStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEmbeddedStatusResponseBody) SetResult(v int32) *UpdateEmbeddedStatusResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateEmbeddedStatusResponseBody) SetSuccess(v bool) *UpdateEmbeddedStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEmbeddedStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
