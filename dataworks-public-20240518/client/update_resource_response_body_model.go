// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateResourceResponseBody
	GetSuccess() *bool
}

type UpdateResourceResponseBody struct {
	// The request ID. You can troubleshoot issues based on the ID.
	//
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA81XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetSuccess(v bool) *UpdateResourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
