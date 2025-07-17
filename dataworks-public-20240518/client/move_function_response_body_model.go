// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MoveFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MoveFunctionResponseBody
	GetSuccess() *bool
}

type MoveFunctionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 48C0E2F0-52BA-5888-BDFA-28F1B9AFXXXX
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

func (s MoveFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *MoveFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveFunctionResponseBody) SetRequestId(v string) *MoveFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveFunctionResponseBody) SetSuccess(v bool) *MoveFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *MoveFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
