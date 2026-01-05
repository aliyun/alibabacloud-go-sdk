// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateComputeResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateComputeResourceResponseBody
	GetSuccess() *bool
}

type UpdateComputeResourceResponseBody struct {
	// The request ID. Used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateComputeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateComputeResourceResponseBody) SetRequestId(v string) *UpdateComputeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeResourceResponseBody) SetSuccess(v bool) *UpdateComputeResourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateComputeResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
