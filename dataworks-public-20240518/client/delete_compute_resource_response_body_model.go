// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteComputeResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteComputeResourceResponseBody
	GetSuccess() *bool
}

type DeleteComputeResourceResponseBody struct {
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// B56432E0-2112-5C97-88D0-AA0AE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded.
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteComputeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComputeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComputeResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteComputeResourceResponseBody) SetRequestId(v string) *DeleteComputeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComputeResourceResponseBody) SetSuccess(v bool) *DeleteComputeResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteComputeResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
